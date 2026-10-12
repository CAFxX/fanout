package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
type GenS2 struct { F0 uint8; F1 string; F2 bool }
type GenS3 struct { F0 bool; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint; F1 uint8; F2 uint8 }
func genFn0(p0 uint, p1 uint16) bool {
    var v7 string = (("foo bar" + ("hello" + "a")) + "hello")
    v8 := GenS1{F0: (p0 & p0), F1: 252, F2: p0}
    p0 -= (^(p0))
    var v9 string = v7
    v10 := GenS2{F0: 2, F1: ((("foo bar" + ("\n\t" + ((v9 + v9) + v9))) + v7) + v9), F2: (!(-1 <= 647))}
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return (0.0 >= (-2.25 / ((3.14159 + (1.5 - (-2.25 * (0.1 * float64(3.14159))))) / 1.5)))
}

func main() {
    var v1 int8 = 100
    var v2 uint64 = 2
    v3 := []int{32767, 32768, 42, 32767}
    v4 := []int{127, 32767, -429, 749}
    v5 := make(map[int]bool)
    var v6 uint16 = 949
    v11 := GenS3{F0: false, F1: 629, F2: 4294967295}
    v6 += v6
    for fi12 := 0; fi12 < 18; fi12++ {
        v6 = (^(v6))
        var v13 uint32 = 42
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 int32 = 1
            v6 = (v6 >> ((18446744073709551615) & 63))
            var v16 int8 = (v1 - v1)
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    switch 186 {
    case -2147483648:
        var v17 int8 = -1
        _ = v17
    case 128:
        v18 := GenS2{F0: 2, F1: "hello", F2: (32768 <= 0)}
        _ = v18
    }
    v19 := GenS2{F0: 127, F1: "hello", F2: (true || false)}
    var v20 int8 = v1
    var v21 int16 = 1
    if (-32768 <= v21) {
        for fi22 := 0; fi22 < 14; fi22++ {
            fmt.Println(v6)
            _ = fi22
        }
        fmt.Println(v1, v2)
        v23 := GenS4{F0: 65535, F1: 52, F2: 184}
        _ = v23
    } else {
        fmt.Println(v6)
        var v24 int8 = v1
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
