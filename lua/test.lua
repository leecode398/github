_M.REDIRECT_STATUS = {
    [301] = "Moved Permanently",
    [302] = "Found",
    [303] = "See Other",
    [307] = "Temporary Redirect",
    [308] = "Permanent Redirect",
}

function _M.request_raw(url, method, data, headers, timeout, proxy, ssl_verify, redirect)
    local httpc = http.new()
    if timeout then
        httpc:set_timeout(timeout * 1000)
    else
        httpc:set_timeout(60 * 1000)
    end
    
    if not redirect then
        redirect = 5
    end
    if redirect < 0 then
        return nil, "exceeded maximum redirect times"
    end
    
    local parsed_uri, err = httpc:parse_uri(url, false)
    if not parsed_uri then
        return nil, "failed to parse url " .. url .. ": " .. err
    end
    local scheme, host, port, path, query = unpack(parsed_uri)
    
    -- We always add "Host" field to the headers.
    local last_headers = {}
    if headers then
        for key, value in pairs(headers) do
            last_headers[key] = value
        end
    end
    if not last_headers["Host"] and not last_headers["host"] then
        local host_header = host
        if port then
            if (scheme == "https" and port == 443) or (scheme == "http" and port == 80) then
                host_header = host_header
            else
                host_header = host_header .. ":" .. port
            end
        end
        last_headers["Host"] = host_header
    end
    
    local params = {
        method = method,
        headers = last_headers,
        body = data,
        path = path, -- Note: use partial request-URL for compatibility
        query = query
    }

    if proxy then
        local parsed_proxy, err = httpc:parse_uri(proxy, false)
        if not parsed_proxy then
            return nil, "failed to parse proxy " .. proxy .. ": " .. err
        end
        local _, proxy_host, proxy_port = unpack(parsed_proxy)
        local ok, err = httpc:connect(proxy_host, proxy_port)
        if not ok then
            return nil, "failed to connect " .. proxy .. ": " .. err
        end
        
        -- HTTPS proxy
        -- https://tools.ietf.org/html/rfc2817#section-5
        -- https://github.com/pintsized/lua-resty-http/issues/63
        if scheme == "https" then
            local res, err = httpc:request({
                version = 1.1,
                method = "CONNECT",
                headers = {["Host"] = last_headers["Host"]},
                path = host .. ":" .. port, -- Note: must include port
            })
            if not res then
                return nil, err
            end
            ngx.log(ngx.NOTICE, res.status .. " " .. res.reason)
            
            if res.status < 200 or res.status >= 300 then
                ngx.log(ngx.ERR, res.status .. " " .. res.reason)
                return nil, "connect proxy failed: " .. res.status .. " " .. res.reason
            end
            
            -- Note: sending a CONNECT request creates a TCP tunnel, the proxy will not
            -- parse data sending by client after tunnel created. In other words, the proxy
            -- will not modify request-URL in start line, so we had better use partial
            -- request-URL for some web server can only recognise partial request-URL.
        else
            -- Note: proxy get origin server host by parsing request-URL in start line.
            params.path = ngx.re.gsub(url, "\\s", "%20", "jo")
        end
    else
        local ok, err = httpc:connect(host, port)
        if not ok then
            return nil, "failed to connect " .. url .. ": " .. err
        end
    end

    if scheme == "https" then
        local verify = true
        if ssl_verify == false then
            verify = false
        end
        local ok, err = httpc:ssl_handshake(nil, host, verify)
        if not ok then
            local ok2, err2 = httpc:close()
            if not ok2 then
                ngx.log(ngx.ERR, err2)
            end
            return nil, "failed to ssl_handshake: " .. err
        end
    end

    local res, err = httpc:request(params)
    if not res then
        local ok2, err2 = httpc:close()
        if not ok2 then
            ngx.log(ngx.ERR, err2)
        end
        return nil, err
    end

    local redirect_url = res.headers["Location"] or res.headers["location"]
    if _M.REDIRECT_STATUS[res.status] and redirect_url and redirect_url ~= "" then
        ngx.log(ngx.NOTICE, res.status .. " " .. res.reason .. ": " .. redirect_url)
        if string.char(string.byte(redirect_url, 1)) == "/" and string.char(string.byte(redirect_url, 2)) ~= "/" then
            redirect_url = scheme .. "://" .. host .. ":" .. port .. redirect_url
        end
        
        if scheme == "https" then
            local ok, err = httpc:close()
            if not ok then
                ngx.log(ngx.ERR, err)
            end
        else
            local ok, err = httpc:set_keepalive()
            if not ok then
                ngx.log(ngx.ERR, "failed to set_keepalive: " .. err)
            end
        end
        return _M.request_raw(redirect_url, method, data, headers, timeout, proxy, ssl_verify, redirect - 1)
    end

    local head_method = method == "head" or method == "HEAD"
    local content = ""
    if res.has_body and not head_method then
        local body, err = res:read_body()
        if not body then
            local ok2, err2 = httpc:close()
            if not ok2 then
                ngx.log(ngx.ERR, err2)
            end
            return nil, err
        end
        content = body
    end

    if scheme == "https" then
        local ok, err = httpc:close()
        if not ok then
            ngx.log(ngx.ERR, err)
        end
    else
        local ok, err = httpc:set_keepalive()
        if not ok then
            ngx.log(ngx.ERR, "failed to set_keepalive: " .. err)
        end
    end

    return {
        status = res.status,
        reason = res.reason,
        headers = res.headers,
        content = content,
    }
end

