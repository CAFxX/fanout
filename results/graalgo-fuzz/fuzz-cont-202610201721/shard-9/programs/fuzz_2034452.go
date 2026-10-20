package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0() string {
    var v4 int8 = -128
    if false {
        var v5 int16 = 401
        v6 := GenS1{F0: 42, F1: (!(128 > 1)), F2: 194}
        switch "hello" {
        case "":
            var v7 uint32 = 42
            _ = v7
        default:
            var v8 bool = true
            _ = v8
        }
        _ = v5
        _ = v6
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        var v10 uint64 = 42
        v10 += 1
        if (128 < 154) {
            v11 := GenS1{F0: 977, F1: ("" >= "a"), F2: 2147483647}
            _ = v11
        } else {
            var v12 uint8 = 128
            for fi13 := 0; fi13 < 14; fi13++ {
                v14 := []int{128, (fi13 | (fi9 << ((2) & 63))), fi13}
                v12 = (v12 + v12)
                if false {
                    var v15 int16 = 336
                    v14 = append(v14, (fi13 - fi13))
                    fmt.Println(v10)
                    _ = v15
                }
                _ = fi13
                _ = v14
            }
            var v16 int16 = -29
            _ = v12
            _ = v16
        }
        _ = fi9
        _ = v10
    }
    var v17 float64 = -2.25
    v4 -= v4
    _ = v17
    _ = v4
    return ("a" + "")
}
func genFn1(p0 int64, p1 bool, p2 uint64) uint64 {
    var v18 int64 = 1
    var v19 uint32 = uint32(p2)
    var v20 uint8 = 128
    v20 = (v20 / (((v20 ^ 0)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    _ = v20
    return p2
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := make(map[int]string)
    v21 := []int{32768, 9223372036854775807}
    var v22 uint16 = 293
    v22 = v22
    if false {
        v22 += v22
    }
    v22 = 32767
    v22 += v22
    v23 := GenS1{F0: 223, F1: true, F2: 0}
    if ((1 <= uint64(v22)) || (!(false && false))) {
        if true {
            v24 := v3[-539]
            var v25 uint = 1
            _ = v24
            _ = v25
        }
        var v26 int64 = 42
        _ = v26
    }
    if (!(false || true)) {
        v27 := GenS2{F0: 65535}
        _ = v27
    } else {
        v22 -= (v22 - 65535)
        v28 := GenS1{F0: 1, F1: true, F2: 74}
        _ = v28
    }
    var v29 float32 = float32(((1e10 * 3.14159) + 0.1))
    v30 := GenS3{F0: uint(v22), F1: v22}
    switch -1 {
    case 0:
        var v31 int8 = 0
        _ = v31
    case 2147483647:
        switch "\n\t" {
        case "a":
            switch 127 {
            case 42:
                if (((true || false) && true) && true) {
                    if false {
                        v32 := GenS3{F0: 4294967295, F1: (v22 % ((v22) | 1))}
                        if false {
                            if len(v21) > 0 {
                                v21[0] = -805
                            }
                        }
                        _ = v32
                    }
                    v29 += v29
                }
            case 127:
                v33 := GenS1{F0: 255, F1: (-132 == -1), F2: 65535}
                _ = v33
            default:
                switch "hello" {
                case "\n\t":
                    _ = len(v21)
                case "hello":
                    if true {
                        var v34 float64 = (-2.25 + ((0.0 * 0.0) + float64(0.0)))
                        var v35 int8 = 42
                        v22 = v22
                        _ = v34
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 10; fi36++ {
                            v22 += (v22 % ((65535) | 1))
                            _ = fi36
                        }
                    }
                case "":
                    _ = len(v21)
                }
            }
        default:
            v2["foo bar"] = (false || false)
        }
    }
    var v37 string = ((("a" + "hello") + "a") + "\n\t")
    if ((v37 + v37) != v37) {
        switch 0 {
        case 0:
            v22 += v22
        default:
            for fi38 := 0; fi38 < 9; fi38++ {
                var v39 uint32 = 0
                _ = fi38
                _ = v39
            }
        }
        if ((128 < 153) || (v37 == (v37 + ""))) {
            var v40 int16 = 1000
            _ = v40
        }
    }
    var v41 uint8 = 104
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v29
    _ = v3
    _ = v30
    _ = v37
    _ = v41
}
