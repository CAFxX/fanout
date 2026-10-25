package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 bool) int {
    v4 := []int{0, -251, 2147483647}
    _ = len(v4)
    var v5 int = 9223372036854775807
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (v5 | (v5 & (-(v5))))
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = false
    var v3 int = 128
    v3 -= v3
    var v6 uint8 = 42
    v2 = (v3 < v3)
    v2 = false
    switch ((((("\n\t" + ("" + "foo bar")) + "\n\t") + "hello") + "foo bar") + "") {
    case "":
        fmt.Println(v3, v3)
    case "foo bar":
        var v7 int64 = 1
        _ = v7
    default:
        v8 := GenS1{F0: v6}
        _ = v8
    }
    v9 := []int{v3, (v3 * (v3 + v3))}
    fmt.Println(v2, v2, v3)
    fmt.Println(v6)
    if v2 {
        v2 = v2
        var v10 uint = 149
        v11 := GenS2{F0: 0}
        _ = v10
        _ = v11
    }
    var v12 string = "a"
    _ = len(v1)
    _ = len(v9)
    v6 -= v6
    for fi13 := 0; fi13 < 1; fi13++ {
        v3 = v3
        _ = fi13
    }
    var v14 int16 = 42
    var v15 uint8 = 128
    var v16 uint8 = v15
    var v17 int64 = 879
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
