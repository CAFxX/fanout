package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS5 struct { F0 uint; F1 uint; F2 int8 }
func genFn0(p0 string, p1 int16) uint16 {
    p1 += (p1 / ((p1) | 1))
    fmt.Println(p1)
    v4 := GenS1{F0: p1}
    _ = p0
    _ = p1
    _ = v4
    return 0
}
func genFn1(p0 uint64, p1 int8) int16 {
    for fi5 := 0; fi5 < 5; fi5++ {
        var v6 int32 = -914
        for fi7 := 0; fi7 < 3; fi7++ {
            if (127 < p1) {
                v8 := GenS2{F0: p1, F1: 65535, F2: p1}
                _ = v8
            }
            if true {
                if true {
                    if true {
                        fmt.Println(p0, fi5, p0)
                        p1 = (p1 ^ p1)
                        fmt.Println(p1, fi5)
                    }
                }
                switch fi7 {
                case 577:
                    v9 := []int{}
                    _ = v9
                case 32768:
                    if (true && false) {
                        v10 := GenS3{F0: 2, F1: p1}
                        _ = v10
                    }
                case 4294967295:
                    switch (-(fi7)) {
                    case 1:
                        v11 := GenS1{F0: 1000}
                        _ = v11
                    case 128:
                        p1 = (p1 + p1)
                    }
                default:
                    v6 = v6
                }
                if (false || ((true && false) && (true && (32767 <= -1)))) {
                    v12 := []int{1}
                    _ = v12
                }
                if false {
                    v13 := GenS4{F0: 42, F1: 2, F2: (v6 - v6)}
                    var v14 uint16 = 65535
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS5{F0: 65535, F1: 431, F2: (p1 * p1)}
                    v16 := []int{(fi5 >> ((255) & 63))}
                    var v17 uint64 = p0
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
            var v18 uint8 = 255
            _ = fi7
            _ = v18
        }
        _ = fi5
        _ = v6
    }
    v19 := GenS4{F0: 65535, F1: 245, F2: 42}
    fmt.Println(p0, p0)
    var v20 int8 = p1
    var v21 uint64 = p0
    _ = p0
    _ = p1
    _ = v19
    _ = v20
    _ = v21
    return -32768
}

func main() {
    var v1 int = -114
    var v2 int32 = 1
    var v3 uint8 = 1
    if (917 < 18446744073709551615) {
        var v22 uint = 1
        v23 := GenS3{F0: 65535, F1: -109}
        var v24 string = "a"
        _ = v22
        _ = v23
        _ = v24
    } else {
        v2 = -2147483648
    }
    var v25 string = "\n\t"
    var v26 string = "foo bar"
    var v27 float64 = 0.1
    v1 = -2147483648
    v27 = (float64(0.1) * 1e10)
    switch 4294967295 {
    case -58:
        v2 = (^(v2))
    case 1:
        v28 := []int{(v1 + -680), (v1 | v1), (^(v1))}
        _ = v28
    case 9223372036854775807:
        fmt.Println(v3)
    default:
        var v29 int64 = int64(v1)
        _ = v29
    }
    var v30 uint = 255
    var v31 float64 = 0.1
    v32 := []int{(v1 + v1), (v1 % (((v1 - v1)) | 1)), (v1 & -9223372036854775808), v1}
    var v33 int64 = -1
    fmt.Println(v25, v27, v30)
    for fi34 := 0; fi34 < 19; fi34++ {
        if ((fi34 - (v1 + (-(fi34)))) == -571) {
            var v35 int32 = v2
            switch (-(fi34)) {
            case 667:
                if len(v32) > 0 {
                    v32[0] = (v1 % ((v1) | 1))
                }
            case -422:
                v36 := GenS1{F0: 306}
                _ = v36
            }
            _ = v35
        } else {
            var v37 uint32 = 1
            _ = v37
        }
        _ = fi34
    }
    if true {
        v32 = append(v32, (v1 + 241))
    } else {
        fmt.Println(v27)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
}
