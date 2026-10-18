package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int16, p1 int, p2 int) string {
    var v6 int16 = p0
    var v7 int64 = 1
    v6 = (p0 >> ((164) & 63))
    v8 := []int{(p2 * (p1 | p2)), p1, p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 uint16) int32 {
    var v9 int32 = -2147483648
    switch 944 {
    case -9223372036854775808:
        if true {
            v10 := []int{}
            _ = v10
        } else {
            fmt.Println(v9, p0)
        }
    }
    v11 := []int{}
    _ = p0
    _ = v11
    _ = v9
    return v9
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    v3 := make(map[string]bool)
    var v4 float32 = ((3.14159 * 1.5) + 0.0)
    var v5 string = ("hello" + "a")
    fmt.Println(v5, v4, v4)
    v5 = v5
    fmt.Println(v5, v1, v1)
    v12 := []int{}
    v4 = (3.14159 * 1e10)
    v5 = (("\n\t" + (v5 + "\n\t")) + v5)
    for fi13 := 0; fi13 < 12; fi13++ {
        v4 = v4
        _ = fi13
    }
    _ = len(v12)
    var v14 bool = (0 <= 42)
    if (v14 || v14) {
        v5 = (v5 + v5)
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        var v16 float64 = 0.1
        fmt.Println(v14, fi15)
        var v17 int64 = -9223372036854775808
        _ = fi15
        _ = v16
        _ = v17
    }
    v18 := GenS1{F0: 2147483648}
    switch 42 {
    case 32767:
        var v19 int64 = 42
        _ = v19
    case 2147483647:
        if (128 <= -658) {
            v20 := GenS1{F0: 886}
            _ = v20
        }
    case 128:
        fmt.Println(v4, v4)
    default:
        fmt.Println(v1)
    }
    for fi21 := 0; fi21 < 20; fi21++ {
        v5 = v5
        _ = len(v3)
        _ = fi21
    }
    v1 = v1
    for fi22 := 0; fi22 < 12; fi22++ {
        var v23 int16 = 0
        _ = fi22
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
