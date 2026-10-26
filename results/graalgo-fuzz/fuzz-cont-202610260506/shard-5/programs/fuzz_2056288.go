package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0(p0 int64) int16 {
    var v6 int8 = -117
    var v7 int16 = -1
    switch "foo bar" {
    case "":
        for fi8 := 0; fi8 < 4; fi8++ {
            v9 := GenS1{F0: 1, F1: 65535, F2: -2147483648}
            var v10 int16 = (v7 ^ v7)
            var v11 float64 = (1e10 / 1.5)
            _ = fi8
            _ = v10
            _ = v11
            _ = v9
        }
    case "foo bar":
        var v12 int32 = -1
        _ = v12
    case "a":
        var v13 bool = false
        _ = v13
    default:
        v14 := []int{}
        _ = v14
    }
    _ = p0
    _ = v6
    _ = v7
    return v7
}

func main() {
    v1 := []string{"hello", ("hello" + "hello"), (("hello" + (("hello" + ("hello" + "")) + ("hello" + ("foo bar" + "a")))) + "")}
    var v2 uint64 = 1
    var v3 float64 = 0.1
    var v4 uint = 65535
    var v5 int = -262
    v4 += v4
    for fi15 := 0; fi15 < 15; fi15++ {
        v3 -= ((v3 + (((float64(1e10) / v3) * v3) / v3)) - (1e10 * float64(1.5)))
        v3 += (3.14159 / 3.14159)
        _ = fi15
    }
    var v16 uint32 = 626
    var v17 int64 = -1
    v2 = (v2 & 0)
    fmt.Println(v5)
    v2 = (v2 / ((1) | 1))
    v4 = 18446744073709551615
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
