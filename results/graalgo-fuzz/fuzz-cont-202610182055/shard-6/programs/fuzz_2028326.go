package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 int16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int8; F1 uint8; F2 uint32 }
func genFn0(p0 uint8, p1 int64) uint8 {
    switch "a" {
    case "hello":
        var v4 int64 = (p1 - 881)
        _ = v4
    }
    var v5 int32 = 357
    _ = p0
    _ = p1
    _ = v5
    return p0
}
func genFn1(p0 int16) uint16 {
    var v6 float64 = 0.0
    var v7 uint64 = 1
    p0 = 0
    var v8 float32 = -2.25
    var v9 uint32 = 1
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint16 = 1
    v3 := make(map[string]uint32)
    _ = len(v1)
    if false {
        var v10 uint64 = 4294967295
        v11 := []int{}
        v10 = (v10 * v10)
        _ = v10
        _ = v11
    }
    switch 9223372036854775807 {
    case 128:
        if true {
            if (138 != 42) {
                var v12 uint64 = 42
                switch 745 {
                case 32768:
                    v12 -= (v12 / ((v12) | 1))
                case -149:
                    if false {
                        switch "foo bar" {
                        case "hello":
                            var v13 uint16 = 906
                            _ = v13
                        }
                        var v14 int64 = -332
                        v14 -= (v14 / ((-1) | 1))
                        _ = v14
                    } else {
                        var v15 uint64 = v12
                        _ = v15
                    }
                default:
                    v16 := GenS1{F0: 70, F1: ((false || false) && (!((!false) || (float64(-2.25) < (3.14159 * 0.0))))), F2: -958}
                    _ = v16
                }
                var v17 float32 = ((1e10 / 3.14159) / 3.14159)
                _ = v12
                _ = v17
            }
            if true {
                if (true && false) {
                    v18 := []int{-9223372036854775808, 9223372036854775807, 32768, -1}
                    fmt.Println(v2)
                    v19 := GenS2{F0: 0, F1: 450, F2: ((((("" + ("a" + "")) + ("a" + ("a" + ("a" + "")))) + (("hello" + "") + "a")) + ((("a" + (("a" + "\n\t") + "")) + "") + "foo bar")) + "")}
                    _ = v18
                    _ = v19
                }
                var v20 bool = true
                _ = v20
            }
            v2 = (v2 * (v2 % (((v2 - v2)) | 1)))
        } else {
            var v21 uint16 = (^(v2))
            _ = len(v3)
            _ = v21
        }
    case 127:
        v2 += v2
    case 134:
        v22 := GenS1{F0: v2, F1: false, F2: 42}
        _ = v22
    default:
        if (!(!true)) {
            var v23 uint32 = 42
            _ = v23
        }
    }
    v2 -= (v2 / ((0) | 1))
    v2 += genFn1(1)
    for fi24 := 0; fi24 < 12; fi24++ {
        if true {
            if (v2 <= (^(v2))) {
                for fi25 := 0; fi25 < 16; fi25++ {
                    var v26 int64 = 1
                    _ = fi25
                    _ = v26
                }
            } else {
                var v27 float32 = (0.0 + ((1e10 + -2.25) * 1.5))
                v27 += (v27 + ((1.5 * ((1.5 + v27) - (float32(1e10) - 0.0))) + 0.1))
                _ = v27
            }
        } else {
            v2 += 1
            v28 := GenS3{F0: 934}
            for fi29 := 0; fi29 < 12; fi29++ {
                v2 -= 42
                var v30 int16 = 379
                _ = fi29
                _ = v30
            }
            _ = v28
        }
        _ = len(v3)
        _ = fi24
    }
    v2 -= 37
    fmt.Println(v2)
    for fi31 := 0; fi31 < 1; fi31++ {
        if ((1 <= 244) || (3.14159 < 1e10)) {
            v2 = 696
        }
        _ = fi31
    }
    if true {
        v32 := []int{32768, -259, 366, 42}
        _ = v32
    } else {
        var v33 bool = true
        _ = v33
    }
    if false {
        v34 := GenS3{F0: 442}
        var v35 int8 = -128
        v36 := GenS4{F0: -2147483648, F1: (-699 < 42)}
        _ = v34
        _ = v35
        _ = v36
    }
    switch 0 {
    case 42:
        switch -922 {
        case -890:
            fmt.Println(v2)
        default:
            v2 += (v2 + (v2 - v2))
        }
    case -1:
        fmt.Println(v2)
    case -578:
        _ = len(v3)
    }
    var v37 float32 = (1.5 + ((((float32(0.1) - ((float32(-2.25) + (0.1 + (0.1 + 0.0))) * float32(1.5))) * float32((0.0 / 1.5))) - 0.0) + 1.5))
    for fi38 := 0; fi38 < 4; fi38++ {
        v39 := []int{(fi38 / ((fi38) | 1)), -1, fi38, 1}
        v40 := GenS5{F0: -112, F1: 0, F2: 42}
        v41 := v1[222]
        _ = fi38
        _ = v39
        _ = v40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
}
