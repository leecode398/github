CXX = g++
LDFLAGS =  bitree.cpp stack.cpp

SRC = 
OBJ = $(SRC:.cc=.o)
EXEC = tree

all: $(EXEC)

$(EXEC): $(OBJ)
	$(CXX) $(LDFLAGS) -o $@ $(OBJ) $(LBLIBS)

clean:
	rm -rf $(OBJ) $(EXEC)