package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint8 }
type GenS2 struct { F0 int32; F1 int; F2 int8 }
type GenS3 struct { F0 int8; F1 uint32; F2 int64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int }
func genFn0(p0 int, p1 string, p2 int16) bool {
    p0 = p0
    var v6 bool = (1.5 < 1e10)
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            var v8 int32 = 0
            for fi9 := 0; fi9 < 15; fi9++ {
                v6 = v6
                fmt.Println(fi7, v6, p2)
                v10 := GenS1{F0: 128, F1: ((!(!(false && v6))) && (0.1 <= 3.14159)), F2: 2}
                _ = fi9
                _ = v10
            }
            fmt.Println(p1, fi7, p0)
            _ = fi7
            _ = v8
        }
        p0 -= -321
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ((0.0 != 0.1) && v6)
}
func genFn1(p0 uint8, p1 string, p2 int32) int64 {
    p2 = 1
    if false {
        p1 = p1
    }
    _ = p0
    _ = p1
    _ = p2
    return -499
}

func main() {
    v1 := []int{825, 776}
    var v2 float32 = -2.25
    var v3 int = 182
    v4 := make(map[int]uint32)
    v5 := []string{"hello", "\n\t", ("a" + (("" + "\n\t") + "hello")), "foo bar"}
    v11 := GenS2{F0: 41, F1: (-(v3)), F2: 1}
    v3 = (v3 % ((v3) | 1))
    v12 := GenS3{F0: 1, F1: 786, F2: 42}
    v13 := GenS1{F0: 608, F1: (0 <= -4), F2: 128}
    v14 := GenS2{F0: 1000, F1: v3, F2: 0}
    _ = len(v4)
    v15 := GenS3{F0: 100, F1: 1, F2: -1}
    if false {
        if false {
            if true {
                v16 := GenS2{F0: 0, F1: (v3 ^ (v3 | v3)), F2: 42}
                var v17 uint = 65535
                _ = v16
                _ = v17
            } else {
                var v18 int64 = 307
                v19 := GenS4{F0: 2}
                switch -9223372036854775808 {
                case -383:
                    v20 := GenS3{F0: -1, F1: 2147483647, F2: 0}
                    _ = v20
                }
                _ = v18
                _ = v19
            }
            v2 = (v2 * v2)
        }
        v21 := GenS2{F0: 814, F1: (^(v3)), F2: -128}
        _ = v21
    }
    switch ("hello" + "") {
    case "a":
        v4[123] = 0
    case "\n\t":
        if (((("hello" + "\n\t") + "") + "\n\t") != "") {
            v3 -= v3
            v3 += 127
        }
    case "hello":
        fmt.Println(v2, v2)
    }
    switch v3 {
    case -1:
        for fi22 := 0; fi22 < 18; fi22++ {
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := GenS5{F0: false, F1: -391}
                var v25 float32 = (float32((float64(1e10) / -2.25)) * (v2 / v2))
                _ = fi23
                _ = v24
                _ = v25
            }
            v2 = float32(0.0)
            var v26 uint32 = 4294967295
            _ = fi22
            _ = v26
        }
    case 414:
        var v27 int64 = -1
        _ = v27
    case -9223372036854775808:
        v28 := GenS1{F0: v3, F1: ((!(904 > -9223372036854775808)) && false), F2: 1}
        _ = v28
    default:
        fmt.Println(v2)
    }
    v3 = (v3 | (v3 / (((v3 * (v3 ^ 128))) | 1)))
    if ((true || (1.5 <= (1e10 * (0.1 / 1.5)))) || ("" <= "\n\t")) {
        delete(v4, 32768)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
