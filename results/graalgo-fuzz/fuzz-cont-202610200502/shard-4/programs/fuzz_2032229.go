package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
func genFn0(p0 int8) string {
    for fi6 := 0; fi6 < 20; fi6++ {
        p0 = p0
        var v7 uint8 = 2
        p0 = 112
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 11; fi8++ {
        fmt.Println(p0)
        switch "hello" {
        case "hello":
            fmt.Println(p0, fi8)
        default:
            var v9 uint32 = 16
            _ = v9
        }
        var v10 string = (("" + ("hello" + (("a" + "hello") + ((("a" + "\n\t") + ("" + ("hello" + ""))) + "")))) + "")
        _ = fi8
        _ = v10
    }
    _ = p0
    return "hello"
}
func genFn1(p0 int64) uint16 {
    v11 := []int{127, int(p0), 4294967295}
    for fi12 := 0; fi12 < 16; fi12++ {
        var v13 uint = 42
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v11
    return 359
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[int]int)
    var v3 float32 = (-2.25 / -2.25)
    var v4 uint16 = 32767
    var v5 int32 = 793
    v14 := GenS1{F0: int16(v5), F1: 103}
    if ((1e10 - (0.0 * float64(v3))) != (0.1 * 1.5)) {
        v1 += (v1 >> ((789) & 63))
        v1 += uint32(v4)
        v15 := []int{128, 2147483647}
        _ = v15
    }
    v4 -= v4
    v16 := v2[1]
    v3 = float32(-2.25)
    for fi17 := 0; fi17 < 18; fi17++ {
        v1 -= v1
        var v18 int8 = -71
        _ = fi17
        _ = v18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v3 += (v3 + 3.14159)
        _ = fi19
    }
    fmt.Println(v5, v1)
    for fi20 := 0; fi20 < 1; fi20++ {
        v16 = (fi20 % ((127) | 1))
        v21 := []int{(fi20 >> ((255) & 63)), 127, 128}
        _ = fi20
        _ = v21
    }
    v3 = 0.0
    var v22 int8 = -1
    switch "foo bar" {
    case "hello":
        fmt.Println(v1, v16, v22)
    case "\n\t":
        v16 = -2147483648
    default:
        if false {
            var v23 bool = (((true || (true || ((v5 - (^(v5))) < (^(v5))))) || false) && true)
            fmt.Println(v1, v1)
            v22 = (v22 / (((^(v22))) | 1))
            _ = v23
        }
    }
    v24 := []int{(v16 - v16), v16, v16, v16}
    var v25 uint64 = 0
    var v26 int64 = 0
    v2[(^(v16))] = (-(v16))
    v26 += v26
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v22
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
