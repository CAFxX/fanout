package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int64; F1 uint; F2 int32 }
type GenS5 struct { F0 uint32 }
func genFn0(p0 int64) int8 {
    var v6 bool = false
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 255, F1: 0}
        for fi9 := 0; fi9 < 15; fi9++ {
            switch fi7 {
            case 4294967295:
                v10 := []int{(^(fi9)), (fi9 % ((fi7) | 1))}
                _ = v10
            }
            v11 := GenS1{F0: 127, F1: 0}
            v12 := GenS1{F0: 135, F1: 1}
            var v13 uint32 = 65535
            if (2 != 1) {
                var v14 uint = 68
                fmt.Println(fi7, v6, v14)
                _ = v14
            }
            if (true && (237 >= p0)) {
                v6 = v6
                if (!true) {
                    var v15 int16 = -647
                    v13 = (v13 & v13)
                    if (false || (true || (v6 && (false || (100 <= 42))))) {
                        v16 := []int{-512}
                        if len(v16) >= 2 {
                            v16 = v16[1:len(v16)]
                        }
                        fmt.Println(v6, fi9)
                        _ = v16
                    }
                    var v17 uint8 = 0
                    _ = v15
                    _ = v17
                } else {
                    p0 = 9223372036854775807
                    p0 -= 1
                    var v18 uint64 = 2
                    _ = v18
                }
                p0 = -9223372036854775808
            } else {
                var v19 bool = true
                var v20 string = ("" + "foo bar")
                _ = v19
                _ = v20
            }
            _ = fi9
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(fi7)
        _ = fi7
        _ = v8
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            switch fi21 {
            case 42:
                p0 = 1
            default:
                var v23 int16 = -725
                _ = v23
            }
            if (128 > 230) {
                v24 := GenS2{F0: fi22}
                _ = v24
            } else {
                v25 := GenS3{F0: 1}
                p0 = (p0 / (((p0 | (p0 % ((p0) | 1)))) | 1))
                _ = v25
            }
            _ = fi22
        }
        _ = fi21
    }
    _ = p0
    _ = v6
    return -128
}

func main() {
    var v1 float32 = (float32((1.5 - 1e10)) * ((1e10 + 1e10) + 1.5))
    v2 := make(map[string]bool)
    var v3 float32 = (0.1 + (float32(float64(v1)) - (v1 * v1)))
    v4 := []string{"foo bar"}
    v5 := []string{("foo bar" + "foo bar"), "\n\t"}
    v26 := GenS4{F0: -9223372036854775808, F1: 489, F2: 1}
    v27 := GenS5{F0: 1}
    var v28 uint64 = 275
    for fi29 := 0; fi29 < 14; fi29++ {
        if ("hello" == "hello") {
            v4 = append(v4, ("foo bar" + ""))
        }
        v28 = v28
        _ = fi29
    }
    var v30 uint8 = 255
    v3 = v1
    var v31 int8 = -106
    var v32 float32 = float32((0.0 * (3.14159 - -2.25)))
    for fi33 := 0; fi33 < 9; fi33++ {
        fmt.Println(v3)
        _ = len(v2)
        _ = fi33
    }
    v3 -= 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
