CXX = clang++
LDFLAGS =  bitree.cpp

SRC = 
OBJ = $(SRC:.cc=.o)
EXEC = tree

all: $(EXEC)

$(EXEC): $(OBJ)
	$(CXX) $(LDFLAGS) -o $@ $(OBJ) $(LBLIBS)

clean:
	rm -rf $(OBJ) $(EXEC)
