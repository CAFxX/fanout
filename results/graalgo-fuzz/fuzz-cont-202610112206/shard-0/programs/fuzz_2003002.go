package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 string; F2 string }
func genFn0() int16 {
    var v5 int64 = 42
    v6 := GenS1{F0: 42, F1: 127, F2: 4294967295}
    v5 = (-(v5))
    v5 = (-(v5))
    v5 -= v5
    _ = v5
    _ = v6
    return 949
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = 0
    v3 := make(map[int]string)
    var v4 bool = false
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := []int{(fi7 + -9223372036854775808), fi7, (-(fi7)), 767}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: 0}
    for fi10 := 0; fi10 < 14; fi10++ {
        switch 9223372036854775807 {
        case -9223372036854775808:
            var v11 uint32 = 1
            _ = v11
        case 703:
            var v12 float64 = float64(float32(0.0))
            _ = v12
        }
        if v4 {
            v13 := GenS3{F0: (fi10 | fi10), F1: ("" + "\n\t"), F2: (((((("hello" + ("hello" + "foo bar")) + "a") + "foo bar") + ("hello" + ("\n\t" + "\n\t"))) + "hello") + "foo bar")}
            v14 := []int{-693, 0}
            var v15 string = "foo bar"
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := []int{(fi10 % ((127) | 1)), (fi10 * (fi10 >> ((195) & 63))), fi10}
        _ = fi10
        _ = v16
    }
    v2 = 577
    for fi17 := 0; fi17 < 10; fi17++ {
        var v18 int = 9223372036854775807
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS2{F0: -325}
            _ = fi19
            _ = v20
        }
        v21 := []int{fi17, (v18 - 948)}
        _ = fi17
        _ = v18
        _ = v21
    }
    var v22 uint8 = 255
    v3[int(v22)] = ("" + "")
    var v23 uint = uint(v22)
    for fi24 := 0; fi24 < 6; fi24++ {
        v23 -= uint(v22)
        v25 := []int{(fi24 - fi24), (fi24 | (^(fi24)))}
        _ = fi24
        _ = v25
    }
    v26 := []int{0, -9223372036854775808, 9223372036854775807, 42}
    v27 := GenS3{F0: 267, F1: ("hello" + ("a" + "\n\t")), F2: "foo bar"}
    if len(v26) > 0 {
        v26[0] = 127
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v9
}
