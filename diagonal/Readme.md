# Understanding this pattern

1. Draw the pattern

//to do - add for 2x2 (and maybe more details to show other sizes)... omg a component to show this would be amazing.
2. Remove the connecting lines when the algorithm "turns"
    - You now have your "groups"

3. There is only one other way to to a diagonal iteration, you just swap the direction of the lines 
    - so in summary you can go at either: 45/225 or 135/315)

ok now we can do our iteration. We know we need to cover the entire rectangle, so we need one pass of both the x and the y axis, the direction we go will determine the way our indices change

Lets start with a 2x2 0going left then down

```go

```

Notice that given any corner, the next set of indices (and so on), will always be the same, the only difference will be the order within the set.

| **Corner** | **Direction** | **Diagonal Order** | 
|---|---|---|

| **Top Left** | Down-Right (135/315)  | `[1], [4, 2], [7, 5, 3]` | 
| **Top Right** | Down-Left (225/45) | `[7], [4, 8], [1, 5, 9]` |
| **Bottom Left** | Up-Right (225/45)  | `[7], [8, 4], [9, 5, 1]` | 
| **Bottom Right** |  Up-Left (135/315)    | `[1], [2, 4], [3, 5, 7]` | 

