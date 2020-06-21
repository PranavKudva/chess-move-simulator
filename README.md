# Chess Move Simulator  
  
Golang based solution to simulate movement of a piece placed on chess board to finite set of cells as per their predefined moves.
  
### Prerequisites  
  
Golang 1.14  

### Third party libraries used

- Testify (github.com/stretchr/testify/assert) for assertions in tests.

## Running the tests  
  
To run tests in all the packages, use:  
  
```  
go test ./...  
```  
  
### Building the binary  
  
To build the binary which will act as a commandline tool, use:  
  
```  
go build .  
```  
  
## Running the binary  

To build the binary which will act as a commandline tool, use:  
  
```  
./chess <piece_type> <cell_number>  
```  
  
## Demo  
  
Assuming the below chessboard:  
  
```  
A1 A2 A3 A4 A5 A6 A7 A8  
B1 B2 B3 B4 B5 B6 B7 B8  
C1 C2 C3 C4 C5 C6 C7 C8  
D1 D2 D3 D4 D5 D6 D7 D8  
E1 E2 E3 E4 E5 E6 E7 E8  
F1 F2 F3 F4 F5 F6 F7 F8  
G1 G2 G3 G4 G5 G6 G7 G8  
H1 H2 H3 H4 H5 H6 H7 H8  
```  
  
Run using each of piece types at cell E3 produces:  
  
```  
  
./chess King E3
E2, E4, D3, F3, D2, F2, D4, F4
  
./chess Queen E3
E2, E1, E4, E5, E6, E7, E8, E9, D3, C3, B3, A3, F3, G3, H3, D2, C1, F2, G1, D4, C5, B6, A7, F4, G5, H6

./chess Bishop E3
D2, C1, F2, G1, D4, C5, B6, A7, F4, G5, H6
  
./chess Rook E3
E2, E1, E4, E5, E6, E7, E8, E9, D3, C3, B3, A3, F3, G3, H3
  
./chess Pawn E3
D3, F3
  
./chess Horse E3
D1, F1, D5, F5, C2, C4, G2, G4  
  
```
