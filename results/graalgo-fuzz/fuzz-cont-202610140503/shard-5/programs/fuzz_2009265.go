package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 int16) uint16 {
    var v5 int16 = p1
    var v6 uint64 = 803
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return uint16(v6)
}
func genFn1(p0 int8, p1 int64) int16 {
    p0 = (p0 ^ p0)
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return -1
}

func main() {
    v1 := []string{"hello"}
    var v2 bool = (128 <= 91)
    var v3 float32 = (float32(3.14159) / -2.25)
    var v4 float64 = 0.0
    var v7 uint64 = 131
    var v8 float64 = 0.0
    v9 := GenS1{F0: 100}
    if (0.0 >= v4) {
        v7 = v7
    }
    if (!(32767 < 32767)) {
        v10 := GenS1{F0: 0}
        fmt.Println(v7, v4, v8)
        _ = v10
    } else {
        if len(v1) > 0 {
            v1[0] = "hello"
        }
        v8 -= ((v4 - ((v8 - v4) / v8)) - float64((v3 + 0.0)))
    }
    v1 = append(v1, ("a" + ("foo bar" + "foo bar")))
    _ = len(v1)
    switch 835 {
    case 127:
        switch 127 {
        case 1:
            v2 = (!(!(-2147483648 >= 9223372036854775807)))
        default:
            v11 := GenS1{F0: 100}
            _ = v11
        }
    case -171:
        switch ("\n\t" + "foo bar") {
        case "\n\t":
            v12 := GenS1{F0: -128}
            _ = v12
        default:
            var v13 int64 = -162
            _ = v13
        }
    default:
        var v14 int64 = -531
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
