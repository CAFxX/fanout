package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32; F1 int; F2 uint16 }
func genFn0(p0 int32, p1 uint32) uint {
    var v4 int64 = 9223372036854775807
    var v5 int32 = (p0 + -1)
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = (-(v4))
        v7 := []int{-262, (fi6 % ((fi6) | 1)), fi6}
        var v8 int16 = 1
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 667
}
func genFn1(p0 int64, p1 int32) uint8 {
    for fi9 := 0; fi9 < 6; fi9++ {
        p1 -= (p1 - p1)
        switch (("\n\t" + "a") + "") {
        case "foo bar":
            if true {
                var v10 int16 = 1
                p0 -= p0
                _ = v10
            }
        default:
            p0 = -1
        }
        var v11 int16 = 32767
        _ = fi9
        _ = v11
    }
    var v12 uint32 = 391
    p1 = p1
    _ = p0
    _ = p1
    _ = v12
    return 1
}

func main() {
    var v1 float64 = 1.5
    var v2 uint = 739
    var v3 uint32 = 42
    v13 := GenS1{F0: -36, F1: (v2 ^ (v2 & genFn0(-572, (v3 + v3))))}
    v14 := GenS1{F0: 100, F1: v2}
    switch "a" {
    case "hello":
        if (((2 <= genFn1(807, 271)) || false) && (true && (((42 != 59) || false) && true))) {
            v2 += (v2 >> ((v2) & 63))
            var v15 float32 = (-2.25 - float32((0.1 + v1)))
            _ = v15
        } else {
            var v16 int16 = 66
            _ = v16
        }
    case "":
        v3 += (v3 & 234)
    default:
        if (2 < v3) {
            for fi17 := 0; fi17 < 4; fi17++ {
                v18 := GenS1{F0: -128, F1: v2}
                v2 += v2
                _ = fi17
                _ = v18
            }
            var v19 int64 = -9223372036854775808
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        var v21 string = ("foo bar" + "")
        v22 := []int{(fi20 % ((fi20) | 1)), fi20}
        v22 = append(v22, (fi20 / ((0) | 1)))
        _ = fi20
        _ = v21
        _ = v22
    }
    if (354 == 0) {
        v23 := GenS2{F0: 210, F1: 32767, F2: 773}
        _ = v23
    }
    v2 = (v2 >> ((v2) & 63))
    switch 128 {
    case 982:
        v24 := []int{-1, -449, -575, 42}
        _ = v24
    case -668:
        var v25 uint64 = 1
        _ = v25
    }
    var v26 float64 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
}
