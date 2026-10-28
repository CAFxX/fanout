package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
func genFn0(p0 int64, p1 uint16) bool {
    p1 = (p1 % ((741) | 1))
    p1 = p1
    _ = p0
    _ = p1
    return (2147483647 <= 628)
}
func genFn1(p0 int32) int {
    v7 := GenS1{F0: 9223372036854775807}
    v8 := []int{42, 943, 4294967295, 127}
    p0 = p0
    var v9 int64 = 42
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int16 = 0
    var v2 uint64 = 299
    var v3 uint16 = 42
    var v4 bool = true
    v5 := make(map[uint32]int)
    var v6 bool = (false && (65535 >= 291))
    var v10 int = -876
    v11 := []int{}
    v5[2147483647] = (v10 * v10)
    v10 -= v10
    v4 = v6
    switch "foo bar" {
    case "hello":
        v3 -= (v3 / ((v3) | 1))
    case "":
        switch -9223372036854775808 {
        case -947:
            switch "hello" {
            case "\n\t":
                fmt.Println(v3, v10, v10)
            default:
                v12 := GenS1{F0: -754}
                _ = v12
            }
        case 128:
            v6 = true
        }
    case "a":
        fmt.Println(v2, v6)
    }
    var v13 float64 = 1e10
    for fi14 := 0; fi14 < 7; fi14++ {
        var v15 string = ("foo bar" + "foo bar")
        var v16 int = fi14
        _ = fi14
        _ = v15
        _ = v16
    }
    v2 = 18446744073709551615
    v17 := v5[940]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
