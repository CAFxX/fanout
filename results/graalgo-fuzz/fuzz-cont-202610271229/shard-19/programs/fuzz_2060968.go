package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16; F1 int16; F2 uint32 }
func genFn0(p0 uint32) uint16 {
    if (537 <= 60) {
        var v7 int32 = -727
        if true {
            if true {
                var v8 uint8 = 127
                p0 = (p0 / ((p0) | 1))
                switch 176 {
                case 703:
                    if true {
                        v8 = v8
                        v8 += v8
                    }
                case -9223372036854775808:
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 uint16 = uint16(v8)
                        _ = fi9
                        _ = v10
                    }
                case 9223372036854775807:
                    fmt.Println(p0, v7, v8)
                }
                _ = v8
            } else {
                v11 := []int{32768, -574, 4294967295}
                _ = v11
            }
            v12 := GenS1{F0: 0, F1: 65535}
            _ = v12
        }
        v13 := GenS1{F0: 32768, F1: 65535}
        _ = v13
        _ = v7
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := []int{(-(fi14)), (fi14 ^ 2147483647)}
        _ = fi14
        _ = v15
    }
    v16 := []int{32767, 32768, -9223372036854775808}
    for fi17 := 0; fi17 < 9; fi17++ {
        if false {
            v18 := GenS2{F0: 42}
            _ = v18
        }
        var v19 string = ("" + "\n\t")
        _ = fi17
        _ = v19
    }
    p0 = 491
    _ = p0
    _ = v16
    return 32767
}
func genFn1(p0 uint64) uint {
    p0 = p0
    fmt.Println(p0)
    fmt.Println(p0)
    v20 := GenS2{F0: -91}
    switch 615 {
    case -20:
        fmt.Println(p0)
    case 186:
        v21 := []int{-2147483648}
        _ = v21
    }
    _ = p0
    _ = v20
    return uint(p0)
}

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    v4 := []int{320, 0, 9223372036854775807}
    v5 := make(map[uint32]string)
    var v6 uint = uint(v1)
    fmt.Println(v1)
    if len(v4) > 0 {
        v4[0] = 32768
    }
    v1 = v1
    var v22 uint32 = 65535
    var v23 int = 9223372036854775807
    for fi24 := 0; fi24 < 14; fi24++ {
        v23 += v23
        _ = fi24
    }
    for fi25 := 0; fi25 < 17; fi25++ {
        switch fi25 {
        case 1:
            delete(v5, v22)
        case 32768:
            if ((((1e10 / float64(-2.25)) / -2.25) + float64(-2.25)) != -2.25) {
                switch (v23 + 20) {
                case 997:
                    v26 := GenS2{F0: 1}
                    _ = v26
                default:
                    var v27 float32 = (1e10 * ((1e10 + 1.5) - 0.1))
                    _ = v27
                }
            } else {
                fmt.Println(v1)
            }
        default:
            v22 = v22
        }
        _ = fi25
    }
    v1 += v1
    if (154 >= -1) {
        switch v23 {
        case 474:
            v28 := v2[(v23 % (((v23 >> (((v6 - genFn1(18446744073709551615))) & 63))) | 1))]
            _ = v28
        default:
            v6 -= genFn1(1)
        }
        v6 = v6
    } else {
        v1 += (v1 ^ v1)
        v29 := GenS2{F0: 100}
        for fi30 := 0; fi30 < 17; fi30++ {
            fmt.Println(fi30, v22)
            var v31 uint64 = 2
            _ = fi30
            _ = v31
        }
        _ = v29
    }
    v23 = v23
    switch "\n\t" {
    case "foo bar":
        if false {
            var v32 uint8 = v1
            _ = v32
        }
    default:
        if true {
            var v33 int32 = -1
            _ = v33
        } else {
            var v34 int16 = -32768
            fmt.Println(v6)
            _ = v34
        }
    }
    v22 += 2147483648
    v23 = (v23 / (((v23 | 32767)) | 1))
    v6 -= genFn1(uint64(v6))
    v6 -= v6
    switch v23 {
    case 32767:
        var v35 uint16 = 42
        _ = v35
    case 32768:
        switch ("\n\t" + ("\n\t" + "")) {
        case "a":
            v6 = 2
        case "\n\t":
            v6 = v6
        default:
            var v36 int16 = -32768
            _ = v36
        }
    default:
        for fi37 := 0; fi37 < 6; fi37++ {
            if false {
                for fi38 := 0; fi38 < 16; fi38++ {
                    var v39 int32 = 2147483647
                    _ = fi38
                    _ = v39
                }
                if false {
                    _ = len(v4)
                    fmt.Println(v23, v6)
                    v22 = v22
                } else {
                    fmt.Println(v6, v1, v1)
                    v40 := GenS3{F0: 1000, F1: 42, F2: v22}
                    _ = v40
                }
                var v41 int32 = 1
                _ = v41
            }
            var v42 int64 = int64(v6)
            var v43 bool = (uint8(v22) > (v1 + (^(v1))))
            _ = fi37
            _ = v42
            _ = v43
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
