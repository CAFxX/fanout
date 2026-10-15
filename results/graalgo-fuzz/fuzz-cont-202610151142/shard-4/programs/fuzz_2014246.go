package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 int8, p2 int64) uint64 {
    var v7 uint64 = 42
    p1 -= int8(p2)
    var v8 string = p0
    if true {
        var v9 uint8 = 255
        var v10 bool = ((v7 ^ v7) >= 951)
        _ = v10
        _ = v9
    } else {
        v11 := GenS1{F0: 128, F1: true}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return v7
}

func main() {
    var v1 uint = 717
    var v2 bool = false
    var v3 int64 = 903
    var v4 bool = false
    var v5 float32 = float32(1.5)
    v6 := make(map[string]bool)
    if false {
        if v4 {
            for fi12 := 0; fi12 < 6; fi12++ {
                v13 := []int{fi12, (fi12 >> (((^(v1))) & 63))}
                _ = fi12
                _ = v13
            }
            if false {
                v14 := []int{}
                _ = v14
            }
            var v15 int32 = -1
            _ = v15
        }
    }
    v3 -= v3
    for fi16 := 0; fi16 < 1; fi16++ {
        v3 -= (v3 ^ v3)
        var v17 uint32 = 2
        for fi18 := 0; fi18 < 2; fi18++ {
            var v19 uint8 = 42
            v20 := GenS2{F0: -32768}
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    fmt.Println(v4)
    var v21 uint16 = 255
    for fi22 := 0; fi22 < 13; fi22++ {
        var v23 int = (fi22 % (((fi22 / ((fi22) | 1))) | 1))
        if ((genFn0(("foo bar" + "hello"), -96, v3) >= 18446744073709551615) || v2) {
            v24 := []int{4294967295}
            switch fi22 {
            case 32767:
                var v25 float64 = 0.1
                _ = v25
            case 127:
                switch "\n\t" {
                case "foo bar":
                    if v4 {
                        _ = len(v24)
                        v2 = false
                    } else {
                        v24 = append(v24, (fi22 ^ (v23 | 32767)))
                        if len(v24) > 0 {
                            v24[0] = v23
                        }
                        switch "foo bar" {
                        case "\n\t":
                            v26 := GenS1{F0: 110, F1: (v5 < (1e10 + float32(1.5)))}
                            _ = v26
                        case "hello":
                            v27 := GenS3{F0: v23}
                            _ = v27
                        }
                    }
                case "":
                    var v28 int = 9223372036854775807
                    _ = v28
                case "hello":
                    v4 = false
                default:
                    v6[((("" + "") + "") + "")] = ((42 != 194) || v4)
                }
            }
            _ = v24
        } else {
            v29 := GenS1{F0: uint8(v21), F1: ((float64(1.5) / (float64((v5 * ((v5 * float32(((float64(((1e10 / v5) / 1.5)) / (-2.25 / 3.14159)) - (0.1 - 1e10)))) + ((float32(-2.25) / (v5 * 0.1)) / 1.5)))) + ((1e10 * 0.1) - 1e10))) == (1e10 + 3.14159))}
            v1 = v1
            for fi30 := 0; fi30 < 6; fi30++ {
                var v31 uint64 = genFn0("hello", -128, 9223372036854775807)
                fmt.Println(v31)
                v3 = v3
                _ = fi30
                _ = v31
            }
            _ = v29
        }
        v2 = (58 != 127)
        _ = fi22
        _ = v23
    }
    switch "foo bar" {
    case "a":
        fmt.Println(v5, v3, v3)
    case "hello":
        v32 := GenS1{F0: 253, F1: (!("a" <= (("foo bar" + "\n\t") + "foo bar")))}
        _ = v32
    }
    v33 := []int{-1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
