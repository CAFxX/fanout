package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 uint16, p1 int64, p2 int16) uint {
    p1 += p1
    switch "a" {
    case "hello":
        v7 := GenS1{F0: (p1 >> ((42) & 63)), F1: (p2 - -538)}
        _ = v7
    default:
        var v8 string = (("hello" + ("foo bar" + "hello")) + "a")
        _ = v8
    }
    if ((2 != 1) && true) {
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint64 = 0
            v11 := GenS2{F0: int16(p0)}
            _ = fi9
            _ = v10
            _ = v11
        }
        p0 = p0
    } else {
        v12 := []int{128}
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 103
    v3 := make(map[string]bool)
    var v4 uint32 = 2147483647
    var v5 int8 = 100
    var v6 float32 = -2.25
    if ((v4 + (^(v4))) > v4) {
        v13 := v1[1]
        _ = v13
    }
    fmt.Println(v4, v4)
    var v14 bool = ((^(v4)) == 4294967295)
    var v15 float64 = ((1e10 / (-2.25 * -2.25)) + ((0.1 + (((3.14159 - 0.0) * 0.0) * 1e10)) * 0.1))
    fmt.Println(v6, v4, v5)
    var v16 float32 = 0.1
    v16 = v16
    v2 -= v2
    var v17 int16 = -1
    var v18 int32 = 42
    delete(v1, 128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
