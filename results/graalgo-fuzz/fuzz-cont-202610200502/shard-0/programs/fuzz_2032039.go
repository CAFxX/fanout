package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 2
    var v3 int32 = -2147483648
    var v4 uint64 = 312
    var v5 int8 = 38
    for fi6 := 0; fi6 < 12; fi6++ {
        var v7 int32 = -1
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v4)
            _ = fi8
        }
        _ = fi6
        _ = v7
    }
    v3 += v3
    v2 -= 53
    var v9 uint = uint(v2)
    switch ("hello" + ("a" + ((("a" + ("a" + "hello")) + ("" + "a")) + ("\n\t" + "hello")))) {
    case "foo bar":
        if false {
            v5 += v5
        } else {
            if (0 != v2) {
                var v10 int64 = 42
                _ = v10
            } else {
                fmt.Println(v5, v5)
                v9 = uint(v2)
            }
        }
    case "\n\t":
        v9 -= v9
    }
    fmt.Println(v9, v9)
    switch "hello" {
    case "hello":
        var v11 uint32 = 932
        _ = v11
    default:
        v12 := GenS1{F0: -9223372036854775808, F1: 2147483647}
        _ = v12
    }
    switch 1 {
    case 32767:
        for fi13 := 0; fi13 < 2; fi13++ {
            var v14 uint8 = v2
            var v15 float32 = 1e10
            v4 = v4
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        v1[0] = (525 >= v4)
    }
    v16 := []int{9223372036854775807, int(v2)}
    for fi17 := 0; fi17 < 11; fi17++ {
        var v18 float64 = 0.0
        v19 := GenS2{F0: (-(v3)), F1: -1}
        _ = fi17
        _ = v18
        _ = v19
    }
    fmt.Println(v3, v2, v4)
    v9 = v9
    fmt.Println(v9)
    v4 = v4
    for fi20 := 0; fi20 < 18; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            v2 = v2
            _ = fi21
        }
        v4 = v4
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
