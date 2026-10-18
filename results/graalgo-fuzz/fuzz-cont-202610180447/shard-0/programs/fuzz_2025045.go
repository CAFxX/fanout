package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 bool }
func genFn0(p0 uint8, p1 uint64) int8 {
    fmt.Println(p1)
    fmt.Println(p0, p1)
    var v6 float32 = (3.14159 * -2.25)
    fmt.Println(v6)
    _ = p0
    _ = p1
    _ = v6
    return -128
}
func genFn1(p0 int, p1 int64) int {
    var v7 int64 = (p1 & (p1 * (p1 % ((p1) | 1))))
    v7 = (p1 ^ 1)
    _ = p0
    _ = p1
    _ = v7
    return (p0 * 9223372036854775807)
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint32 = 943
    v3 := []string{("a" + "a")}
    var v4 uint = 1
    v5 := make(map[int]string)
    fmt.Println(v4, v2)
    fmt.Println(v2, v2)
    switch ("foo bar" + "") {
    case "hello":
        for fi8 := 0; fi8 < 15; fi8++ {
            v4 += (^(v4))
            v4 += 2
            v4 = v4
            _ = fi8
        }
    default:
        v9 := GenS1{F0: true}
        _ = v9
    }
    v10 := GenS2{F0: 2147483647, F1: (0 <= 0)}
    for fi11 := 0; fi11 < 11; fi11++ {
        if ((1.5 - (3.14159 * 1.5)) != 3.14159) {
            for fi12 := 0; fi12 < 11; fi12++ {
                v4 -= 306
                _ = fi12
            }
            var v13 uint8 = 1
            v4 -= v4
            _ = v13
        } else {
            v4 -= v4
        }
        _ = fi11
    }
    switch -37 {
    case 128:
        v4 -= 2
    default:
        for fi14 := 0; fi14 < 17; fi14++ {
            var v15 uint8 = 2
            fmt.Println(fi14)
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v2)
    var v16 uint = (v4 | (v4 ^ v4))
    v2 += 2
    var v17 uint64 = 582
    if (!(827 <= 42)) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v18 float32 = -2.25
    v19 := GenS2{F0: 128, F1: (true && true)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
