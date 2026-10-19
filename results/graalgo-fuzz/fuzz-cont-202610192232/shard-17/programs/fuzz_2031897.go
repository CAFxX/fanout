package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int64; F2 uint64 }
func genFn0(p0 uint64, p1 uint) int16 {
    v6 := GenS1{F0: ""}
    var v7 uint16 = 854
    p0 += (p0 << ((p1) & 63))
    v8 := []int{-2147483648, 127, -1}
    v9 := GenS2{F0: 520, F1: -356, F2: 117}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 496
}

func main() {
    var v1 int = 0
    var v2 int16 = -32768
    v3 := make(map[string]bool)
    v4 := make(map[string]bool)
    var v5 uint64 = 32
    switch v1 {
    case -510:
        var v10 int16 = v2
        _ = v10
    default:
        var v11 int16 = v2
        _ = v11
    }
    var v12 uint = uint(v5)
    v12 = (^(v12))
    fmt.Println(v12, v2)
    var v13 bool = true
    var v14 uint64 = 22
    switch "\n\t" {
    case "foo bar":
        v3["foo bar"] = v13
    default:
        _ = len(v4)
    }
    switch ("foo bar" + "\n\t") {
    case "hello":
        var v15 int16 = v2
        _ = v15
    default:
        fmt.Println(v1, v1)
    }
    fmt.Println(v13, v12)
    if (v13 || (("\n\t" + (("\n\t" + ("a" + "hello")) + "foo bar")) <= "")) {
        v16 := v3["\n\t"]
        v17 := GenS2{F0: 2147483648, F1: 0, F2: 1}
        var v18 int = 32768
        _ = v16
        _ = v17
        _ = v18
    }
    v5 = (^(v14))
    v19 := []int{v1, (v1 >> (((^(v12))) & 63)), v1}
    _ = len(v19)
    var v20 uint64 = 1
    for fi21 := 0; fi21 < 20; fi21++ {
        switch 2147483647 {
        case -2147483648:
            v22 := GenS1{F0: "\n\t"}
            _ = v22
        default:
            fmt.Println(v1, v2)
        }
        _ = fi21
    }
    fmt.Println(v14, v12, v5)
    fmt.Println(v20, v14)
    var v23 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
