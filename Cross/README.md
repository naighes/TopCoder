# Cross
URL: (https://community.topcoder.com/stat?c=problem_statement&pm=14521&rd=16851)
## Problem tatement
There is a rectangular board that is divided into n rows by m columns of cells. Each cell is either black or white. You are given the description of the board in the `String[]` board. Each character in board represents one cell. More precisely, the character `board[i][j]` represents the cell at coordinates (row `i`, column `j`). The character `'#'` represents a black cell, the character `'.'` is a white cell. 
You want to find a cross on this board. A cross is a configuration of 5 black cells placed like this:

```
 #
###
 #
```

Formally, there is a cross centered at `(x, y)` if the following five cells are all black: `(x, y)`, `(x+1, y)`, `(x-1, y)`, `(x, y-1)`, and `(x, y+1)`. Note that other cells adjacent to the cross may also be black, it is still a valid cross. 

Return `"Exists"` if there is at least one cross on the given board. Otherwise, return `"Does not exist"`.

## Examples
0. `{".##", "###","##."}`
**Returns**: `"Exist"`
1. `{".##", "###", "#.."}`
**Returns**: `"Does not exist"`
This time there is no cross.
2. `{"######", "######", "######", "######", "######", "######", "######"}`
**Returns**: `"Exist"`
3. `{".#.#", "#.#.", ".#.#", "#.#."}`
**Returns**: `"Does not exist"`
4. `{".#.#", "###.", ".###", "#.#."}`
**Returns**: `"Exist"`
