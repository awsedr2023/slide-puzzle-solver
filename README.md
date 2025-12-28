# Slide Puzzle Solver

[English](#english) | [日本語](#japanese)

<a name="english"></a>
## English

A sliding puzzle solver written in Go. This tool finds the shortest path to solve a sliding puzzle of arbitrary size (e.g., 8-puzzle, 15-puzzle) using IDA* (Iterative Deepening A*) algorithm.

### Features

- Solves N x M sliding puzzles (minimum 2x2).
- Finds the shortest solution (optimal).
- Detects unsolvable configurations.
- Supports custom start and goal states.

### Usage

#### Build

```bash
go build -o slide-puzzle-solver cmd/solver/main.go
```

#### Run

The blank tile is represented by the largest number (`rows * cols`). For a 3x3 puzzle, the blank tile is `9`.

The input slice represents the puzzle board from left to right, top to bottom. For example, the arguments `1 8 2 4 3 5 7 6 9` for a 3x3 puzzle correspond to the following board:

```
1 8 2
4 3 5
7 6 9
```
(Where `9` is the blank tile)

**Standard Goal:**

To solve a 3x3 puzzle where the goal is `1 2 3 4 5 6 7 8 9`:

```bash
./slide-puzzle-solver -rows 3 -cols 3 1 8 2 4 3 5 7 6 9
```

**Custom Goal:**

You can provide both start and goal states. The first `rows * cols` numbers are the start state, and the next `rows * cols` numbers are the goal state.

```bash
./slide-puzzle-solver -rows 2 -cols 2 4 1 3 2 1 2 3 4
```

### Library Usage
You can also use this package as a library in your Go programs. 
#### Installation 
```bash
go get github.com/awsedr2023/slide-puzzle-solver
```
#### Example
The `Solve` function returns a slice of blank tile indices representing the path from start to goal, including the initial position.

```go
package main

import (
    "fmt"
    "log"

    "github.com/awsedr2023/slide-puzzle-solver/solver"
)

func main() {
    rows, cols := 3, 3
    start := []int{1, 2, 3, 4, 5, 6, 8, 7, 9} // 9 is the blank tile
    goal := solver.StandardGoal(rows, cols)

    path, err := solver.Solve(start, goal, rows, cols)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("Solved in %d moves\n", len(path)-1)
}
```

---

<a name="japanese"></a>
## 日本語

Go言語で書かれたスライドパズルソルバーです。IDA*アルゴリズムを使用して、任意のサイズのスライドパズル（8パズル、15パズルなど）の最短解を見つけます。

### 特徴

- N x M のスライドパズルを解くことができます（最小サイズは 2x2）。
- 最短手順（最適解）を見つけます。
- 解けない配置（Unsolvable）を検出します。
- 任意の開始状態とゴール状態を指定可能です。

### 使い方

#### ビルド

```bash
go build -o slide-puzzle-solver cmd/solver/main.go
```

#### 実行

空のマスは最大値（`行数 * 列数`）で表されます。3x3パズルの場合、空マスは `9` です。

入力のスライスは、パズルの盤面を左から右、上から下の順に並べたものに対応します。
例えば、3x3パズルで `1 8 2 4 3 5 7 6 9` という引数を指定した場合、以下の盤面に対応します。

```
1 8 2
4 3 5
7 6 9
```
（`9`が空マスです）

**標準ゴール:**

ゴールが `1 2 3 4 5 6 7 8 9` となる3x3パズルを解く場合：

```bash
./slide-puzzle-solver -rows 3 -cols 3 1 8 2 4 3 5 7 6 9
```

**カスタムゴール:**

開始状態とゴール状態の両方を指定することもできます。最初の `行数 * 列数` 個の数字が開始状態、次の `行数 * 列数` 個の数字がゴール状態です。

```bash
./slide-puzzle-solver -rows 2 -cols 2 4 1 3 2 1 2 3 4
```

### ライブラリとしての使用
このパッケージは、Goプログラム内でライブラリとして使用することもできます。
#### インストール
```bash
go get github.com/awsedr2023/slide-puzzle-solver
```
#### 使用例
`Solve` 関数は `[]int` 型の移動履歴を返します。これには初期位置の空マスインデックスも含まれます。

```go
package main

import (
    "fmt"
    "log"

    "github.com/awsedr2023/slide-puzzle-solver/solver"
)

func main() {
    rows, cols := 3, 3
    start := []int{1, 2, 3, 4, 5, 6, 8, 7, 9} // 9が空マス
    goal := solver.StandardGoal(rows, cols)

    path, err := solver.Solve(start, goal, rows, cols)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Printf("%d手で解けました\n", len(path)-1)
}
```