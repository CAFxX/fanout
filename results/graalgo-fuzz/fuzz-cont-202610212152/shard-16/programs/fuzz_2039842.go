package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
func genFn0() uint32 {
    v7 := GenS1{F0: 255}
    var v8 int8 = -128
    v8 = v8
    v8 = -5
    v8 = 42
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 int64, p1 uint64) int64 {
    if (-2.25 >= 3.14159) {
        fmt.Println(p0, p0)
        v9 := []int{127, 30, 32767, -1}
        _ = v9
    } else {
        v10 := GenS1{F0: 0}
        if false {
            p1 = 4294967295
            var v11 string = ("hello" + "a")
            var v12 uint8 = 127
            _ = v11
            _ = v12
        }
        var v13 int64 = 35
        _ = v10
        _ = v13
    }
    p0 = p0
    p0 = (p0 | (^(p0)))
    switch 4294967295 {
    case 1:
        switch "foo bar" {
        case "":
            v14 := GenS2{F0: 32767}
            _ = v14
        case "a":
            var v15 uint16 = 65535
            _ = v15
        default:
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v16 int32 = -2147483648
                _ = v16
            }
        }
    case 0:
        v17 := GenS1{F0: 499}
        _ = v17
    case -704:
        p1 -= p1
    }
    _ = p0
    _ = p1
    return (-(p0))
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]int)
    v3 := []int{-370, 0, -9223372036854775808, 2147483647}
    var v4 int64 = 42
    var v5 uint8 = 0
    var v6 uint8 = v5
    v6 = 1
    v18 := GenS1{F0: uint16(v6)}
    var v19 uint32 = 4294967295
    v4 = 1
    var v20 float32 = float32(0.0)
    v5 -= (v6 * v6)
    var v21 float64 = (0.0 + ((3.14159 + ((0.1 * 0.0) + 0.1)) / ((0.0 - 1e10) / 1.5)))
    v5 = v6
    v21 = (3.14159 - v21)
    fmt.Println(v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
