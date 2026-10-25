package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int16; F1 bool; F2 uint16 }
func genFn0(p0 uint64, p1 uint32, p2 bool) int32 {
    switch ("a" + "a") {
    case "":
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 int = (-(fi6))
            var v8 bool = ((2 == uint(p0)) && p2)
            p0 += p0
            _ = fi6
            _ = v7
            _ = v8
        }
    case "hello":
        var v9 uint16 = 2
        _ = v9
    default:
        var v10 uint8 = 1
        _ = v10
    }
    var v11 int32 = 864
    v11 = v11
    fmt.Println(p0, p2, p0)
    for fi12 := 0; fi12 < 1; fi12++ {
        var v13 uint32 = (p1 + (p1 << ((42) & 63)))
        p1 = 1
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return (v11 >> ((2) & 63))
}
func genFn1(p0 int) int8 {
    fmt.Println(p0)
    var v14 uint8 = 2
    if (538 >= 9223372036854775807) {
        fmt.Println(v14, p0)
        v15 := GenS1{F0: 32767, F1: p0}
        v16 := GenS2{F0: -665, F1: true, F2: 32767}
        _ = v15
        _ = v16
    }
    var v17 int32 = 0
    _ = p0
    _ = v14
    _ = v17
    return 42
}

func main() {
    var v1 int32 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]int)
    var v4 float64 = 1.5
    v5 := make(map[uint32]int)
    var v18 bool = true
    var v19 int32 = 217
    var v20 int = 1
    var v21 int = (v20 << ((18446744073709551615) & 63))
    if (v4 >= (v4 + 0.0)) {
        if v18 {
            var v22 int32 = 695
            _ = v22
        }
        v23 := GenS2{F0: 166, F1: ("foo bar" >= "foo bar"), F2: 2}
        _ = v23
    }
    v24 := GenS1{F0: -621, F1: (-(v21))}
    if v18 {
        var v25 string = "a"
        _ = v25
    }
    fmt.Println(v18, v18)
    if false {
        v3[-1] = v21
    } else {
        v26 := GenS1{F0: 477, F1: -1}
        switch v20 {
        case -319:
            var v27 float64 = ((((0.0 + 1.5) + (v4 * 1.5)) + -2.25) * (v4 / (0.1 / 0.1)))
            _ = v27
        case -385:
            v28 := []int{(^(v20)), (-(v20)), v21}
            _ = v28
        default:
            v1 -= v19
        }
        _ = len(v5)
        _ = v26
    }
    fmt.Println(v4, v4, v21)
    fmt.Println(v1, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
