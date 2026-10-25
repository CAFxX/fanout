package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 uint64) uint16 {
    v5 := []int{-9223372036854775808, 420, 90, -282}
    var v6 int64 = 912
    v7 := GenS1{F0: 0}
    var v8 string = ""
    v5 = append(v5, -1)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 bool = true
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 uint16 = 819
    v9 := v3[0]
    var v10 uint = 2
    _ = len(v3)
    v1 = (1000 == -1)
    fmt.Println(v1)
    v11 := GenS1{F0: 0}
    if (((("\n\t" + "foo bar") + "\n\t") + ("foo bar" + "\n\t")) <= "\n\t") {
        switch 0 {
        case 1:
            var v12 int64 = 758
            _ = v12
        case 2147483647:
            var v13 int8 = -128
            _ = v13
        case -9223372036854775808:
            _ = len(v2)
        }
    } else {
        v14 := v2["a"]
        fmt.Println(v4)
        _ = v14
    }
    v15 := []int{32767, 0, 32767}
    var v16 int32 = 42
    var v17 uint8 = 128
    v2["a"] = 0
    var v18 float32 = (1.5 - (-2.25 + -2.25))
    fmt.Println(v4, v10, v10)
    v10 -= v10
    v15 = append(v15, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
