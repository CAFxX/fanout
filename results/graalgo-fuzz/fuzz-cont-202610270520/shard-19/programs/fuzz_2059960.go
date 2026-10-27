package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 uint16; F2 string }
type GenS4 struct { F0 uint8; F1 int32; F2 uint }
func genFn0(p0 uint8) uint {
    switch "" {
    case "":
        p0 -= p0
    case "hello":
        v5 := GenS1{F0: -1}
        _ = v5
    default:
        v6 := []int{127, 4294967295}
        _ = v6
    }
    if false {
        for fi7 := 0; fi7 < 19; fi7++ {
            if (!(false && true)) {
                v8 := []int{4294967295}
                _ = v8
            }
            for fi9 := 0; fi9 < 5; fi9++ {
                v10 := GenS1{F0: -298}
                var v11 int8 = int8(p0)
                p0 -= p0
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = fi7
        }
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            switch (("hello" + "") + (("\n\t" + "hello") + "")) {
            case "foo bar":
                p0 = 57
            default:
                if true {
                    p0 -= (p0 - 0)
                } else {
                    fmt.Println(p0)
                    var v13 bool = false
                    _ = v13
                }
            }
            for fi14 := 0; fi14 < 5; fi14++ {
                p0 += p0
                p0 = 1
                _ = fi14
            }
            _ = fi12
        }
        v15 := []int{-627, 686, 1}
        _ = v15
    }
    _ = p0
    return 65535
}
func genFn1(p0 uint64, p1 int64, p2 uint32) uint8 {
    for fi16 := 0; fi16 < 17; fi16++ {
        v17 := GenS1{F0: p1}
        v18 := []int{(fi16 | (fi16 / ((fi16) | 1))), 970}
        var v19 int = fi16
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = p2
    return 74
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[string]uint32)
    v3 := make(map[string]uint32)
    var v4 int16 = -32768
    var v20 int8 = 113
    v1 = v1
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS1{F0: 668}
        v4 += (v4 / (((-(v4))) | 1))
        v23 := GenS2{F0: 917}
        _ = fi21
        _ = v22
        _ = v23
    }
    v20 = (v20 ^ v20)
    switch v1 {
    case "a":
        var v24 int32 = 490
        _ = v24
    case "":
        v25 := GenS1{F0: -9223372036854775808}
        _ = v25
    case "foo bar":
        var v26 int64 = 516
        _ = v26
    default:
        fmt.Println(v1)
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v28 := GenS3{F0: v20, F1: 65535, F2: ((v1 + v1) + v1)}
        var v29 int = (fi27 % ((fi27) | 1))
        switch (("a" + v1) + ("a" + "hello")) {
        case "hello":
            v30 := GenS3{F0: v20, F1: 568, F2: "foo bar"}
            _ = v30
        default:
            if true {
                v31 := GenS4{F0: 127, F1: 340, F2: 529}
                var v32 int64 = -83
                v32 -= int64(v20)
                _ = v31
                _ = v32
            }
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    _ = len(v3)
    var v33 int = 42
    if true {
        v34 := GenS2{F0: 65535}
        for fi35 := 0; fi35 < 14; fi35++ {
            var v36 string = v1
            _ = len(v3)
            _ = fi35
            _ = v36
        }
        _ = v34
    } else {
        switch "hello" {
        case "\n\t":
            var v37 int = (v33 | (v33 % ((v33) | 1)))
            _ = v37
        default:
            v38 := GenS4{F0: 1, F1: 879, F2: 0}
            _ = v38
        }
    }
    var v39 uint = 587
    v40 := []int{(v33 << ((v39) & 63)), (v33 ^ v33)}
    v41 := GenS2{F0: 255}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v33
    _ = v39
    _ = v4
    _ = v40
    _ = v41
}
