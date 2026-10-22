package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 uint; F1 uint64; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int16; F1 int64 }
func genFn0(p0 uint64, p1 int8) uint64 {
    v6 := GenS1{F0: p0}
    fmt.Println(p0)
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    return p0
}
func genFn1(p0 uint32, p1 uint8, p2 uint32) int {
    fmt.Println(p0, p2)
    var v7 int = -109
    v8 := []int{int(p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint64 = 0
    v2 := make(map[string]uint32)
    v3 := make(map[int]bool)
    var v4 float64 = 1.5
    v5 := make(map[int]uint32)
    v9 := GenS2{F0: 300, F1: 2}
    delete(v2, (("a" + (("" + ("a" + "")) + "")) + "\n\t"))
    switch 32767 {
    case 128:
        var v10 int32 = 2147483647
        _ = v10
    case 32768:
        v11 := GenS2{F0: 0, F1: 18446744073709551615}
        _ = v11
    case -585:
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 float32 = 3.14159
            for fi14 := 0; fi14 < 20; fi14++ {
                var v15 uint32 = 127
                v1 -= genFn0((v1 ^ 1), 94)
                v16 := GenS3{F0: 239, F1: (^(v1)), F2: 346}
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 uint = 42
            _ = fi12
            _ = v13
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := GenS4{F0: 440}
            _ = fi18
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 2; fi20++ {
        v4 -= ((v4 - (float64(1e10) / 1e10)) * ((v4 + ((3.14159 * (v4 + (0.0 + v4))) - v4)) / 3.14159))
        for fi21 := 0; fi21 < 19; fi21++ {
            var v22 int16 = -939
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    _ = len(v3)
    v4 = v4
    var v23 string = "a"
    v4 = ((0.1 * (3.14159 - ((v4 - (float64(0.1) * v4)) / v4))) / ((1.5 + v4) / (v4 * float64((1.5 * 0.1)))))
    for fi24 := 0; fi24 < 2; fi24++ {
        if (111 < 199) {
            var v25 uint = 594
            v26 := GenS5{F0: 1000, F1: 1}
            v27 := GenS5{F0: -934, F1: 365}
            _ = v25
            _ = v26
            _ = v27
        }
        v23 = "foo bar"
        _ = fi24
    }
    v1 = v1
    switch 0 {
    case 4294967295:
        for fi28 := 0; fi28 < 19; fi28++ {
            v1 -= 42
            v23 = "a"
            _ = fi28
        }
    case 32767:
        var v29 bool = (-1 <= -1)
        _ = v29
    }
    var v30 int8 = -128
    v31 := GenS2{F0: -1, F1: 348}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v9
}
