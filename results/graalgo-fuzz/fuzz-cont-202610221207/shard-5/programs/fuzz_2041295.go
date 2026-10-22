package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 int32 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 int8) uint8 {
    v6 := GenS1{F0: 255}
    p0 += (p0 ^ (^(p0)))
    v7 := GenS2{F0: 352, F1: 255, F2: 958}
    _ = p0
    _ = v6
    _ = v7
    return 2
}
func genFn1() string {
    v8 := GenS3{F0: 4294967295, F1: 2147483647, F2: 18446744073709551615}
    switch 1 {
    case 4294967295:
        var v9 uint64 = 731
        _ = v9
    case -32:
        v10 := []int{}
        _ = v10
    default:
        var v11 int64 = -780
        _ = v11
    }
    _ = v8
    return "hello"
}

func main() {
    var v1 int16 = -32768
    var v2 int8 = -128
    var v3 int = 250
    v4 := make(map[int]string)
    var v5 uint = 2
    switch ("a" + "foo bar") {
    case "\n\t":
        fmt.Println(v3, v1, v5)
    case "":
        for fi12 := 0; fi12 < 16; fi12++ {
            v5 -= 4294967295
            _ = fi12
        }
    case "foo bar":
        v13 := []int{v3}
        _ = v13
    default:
        v4[(v3 ^ (v3 / ((v3) | 1)))] = ""
    }
    v1 += -32768
    for fi14 := 0; fi14 < 16; fi14++ {
        v1 += v1
        _ = fi14
    }
    fmt.Println(v3, v5)
    switch (v3 >> (((^(v5))) & 63)) {
    case -22:
        switch (v3 >> ((v5) & 63)) {
        case 2147483647:
            var v15 float64 = (3.14159 - 0.1)
            _ = v15
        default:
            var v16 bool = (-9223372036854775808 == 42)
            _ = v16
        }
    case 2:
        var v17 uint32 = 1
        _ = v17
    default:
        v5 -= 638
    }
    var v18 int = v3
    v3 -= -4
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
