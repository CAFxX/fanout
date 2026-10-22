package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 int8; F1 uint16; F2 int8 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 int32, p1 int, p2 uint64) uint32 {
    fmt.Println(p2, p2, p2)
    v6 := GenS1{F0: (359 <= -1), F1: 437}
    if (!false) {
        fmt.Println(p2)
        v7 := []int{p1}
        switch p1 {
        case -2147483648:
            p1 += p1
        case 127:
            var v8 uint16 = 2
            _ = v8
        case 0:
            p0 = -824
        default:
            if (127 != -128) {
                p0 += p0
                var v9 float64 = ((1.5 - 1e10) - 0.0)
                _ = v9
            }
        }
        _ = v7
    } else {
        for fi10 := 0; fi10 < 2; fi10++ {
            v11 := []int{(^(p1)), p1, (fi10 * (p1 | p1)), (p1 & fi10)}
            var v12 uint32 = 2147483647
            _ = fi10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 7; fi13++ {
            if true {
                v14 := []int{}
                var v15 int64 = 9223372036854775807
                _ = v14
                _ = v15
            }
            _ = fi13
        }
    }
    v16 := []int{(p1 | p1), -9223372036854775808, (p1 % ((p1) | 1))}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v6
    return 0
}
func genFn1() string {
    if false {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint8 = 8
            switch 127 {
            case 975:
                v19 := []int{-2147483648}
                _ = v19
            }
            var v20 uint64 = 4294967295
            _ = fi17
            _ = v18
            _ = v20
        }
        var v21 string = ""
        _ = v21
    }
    v22 := GenS1{F0: true, F1: -457}
    var v23 int32 = 2147483647
    v24 := GenS2{F0: 4294967295, F1: 306, F2: 0}
    _ = v22
    _ = v23
    _ = v24
    return "hello"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int = 852
    var v3 int = v2
    v4 := make(map[int]bool)
    v5 := make(map[uint32]string)
    var v25 bool = true
    var v26 int16 = -814
    var v27 float32 = 1.5
    switch "foo bar" {
    case "a":
        v28 := GenS3{F0: -56, F1: 255, F2: 100}
        _ = v28
    default:
        for fi29 := 0; fi29 < 5; fi29++ {
            fmt.Println(v25, v2, v2)
            fmt.Println(v3)
            for fi30 := 0; fi30 < 20; fi30++ {
                v3 = -409
                var v31 float64 = float64(v27)
                fmt.Println(v2, fi29, v31)
                _ = fi30
                _ = v31
            }
            _ = fi29
        }
    }
    var v32 float32 = 0.1
    if ("a" == ("" + "\n\t")) {
        fmt.Println(v25)
    }
    v27 -= ((v32 + (v27 - (v27 + -2.25))) - (v32 * v32))
    v32 = v32
    v25 = ((v26 > v26) || ((v27 == v32) || (float64((v27 + -2.25)) >= -2.25)))
    delete(v5, 706)
    if v25 {
        if (v25 && v25) {
            var v33 string = "hello"
            _ = v33
        } else {
            if false {
                v34 := GenS4{F0: 146}
                v2 = v2
                var v35 string = "\n\t"
                _ = v34
                _ = v35
            }
        }
        var v36 uint32 = 8
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
