package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int8; F2 bool }
type GenS3 struct { F0 uint32; F1 uint16; F2 int32 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int64; F1 uint16; F2 int32 }
func genFn0(p0 int64, p1 string) uint8 {
    v6 := GenS1{F0: 127, F1: 42}
    var v7 string = p1
    fmt.Println(v7, v7)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 bool) bool {
    v8 := []int{-951}
    switch 1 {
    case 4294967295:
        if (460 == 42) {
            p0 = true
            var v9 bool = false
            switch -800 {
            case 127:
                var v10 int = -1
                _ = v10
            case 9223372036854775807:
                if p0 {
                    v11 := GenS1{F0: -35, F1: 2147483648}
                    fmt.Println(p0, p0)
                    _ = v11
                }
            case 42:
                v12 := GenS2{F0: -9223372036854775808, F1: 1, F2: p0}
                _ = v12
            }
            _ = v9
        }
    case -9223372036854775808:
        var v13 int64 = -9223372036854775808
        _ = v13
    }
    var v14 uint64 = 809
    _ = p0
    _ = v14
    _ = v8
    return p0
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint64 = 4294967295
    var v3 bool = false
    v4 := make(map[int]string)
    v5 := []string{"\n\t", ("" + ("a" + "foo bar"))}
    v15 := GenS3{F0: 4294967295, F1: 422, F2: 1}
    v16 := GenS1{F0: -72, F1: 150}
    v17 := GenS4{F0: -555, F1: 88}
    if (-109 <= 1) {
        v3 = true
        var v18 int = 128
        v19 := GenS5{F0: (^(v2)), F1: -1}
        _ = v18
        _ = v19
    } else {
        var v20 uint64 = v2
        v21 := GenS6{F0: -1, F1: 65535, F2: 283}
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        _ = v20
        _ = v21
    }
    v3 = false
    fmt.Println(v3, v2)
    v3 = (623 <= 42)
    for fi22 := 0; fi22 < 3; fi22++ {
        switch "hello" {
        case "foo bar":
            var v23 int8 = -1
            _ = v23
        case "hello":
            if (-371 < fi22) {
                v2 -= (^(v2))
            }
        }
        var v24 float32 = ((1.5 * 1e10) * (1.5 + 0.1))
        _ = fi22
        _ = v24
    }
    fmt.Println(v3, v3)
    switch 128 {
    case 177:
        v25 := GenS5{F0: (^(v2)), F1: 42}
        _ = v25
    }
    var v26 float32 = -2.25
    switch "\n\t" {
    case "":
        v3 = ((false || v3) && v3)
    default:
        for fi27 := 0; fi27 < 2; fi27++ {
            if (v3 || v3) {
                v26 -= v26
                fmt.Println(v3, v3)
                if ("\n\t" == "hello") {
                    var v28 uint32 = 294
                    var v29 uint16 = 123
                    for fi30 := 0; fi30 < 9; fi30++ {
                        v2 += v2
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = fi30
                    }
                    _ = v28
                    _ = v29
                }
            }
            switch 0 {
            case -1:
                if v3 {
                    switch (fi27 << ((18446744073709551615) & 63)) {
                    case 641:
                        switch "foo bar" {
                        case "foo bar":
                            var v31 int32 = 0
                            _ = v31
                        case "":
                            v3 = (!v3)
                        }
                    default:
                        v3 = v3
                    }
                    switch fi27 {
                    case -224:
                        for fi32 := 0; fi32 < 20; fi32++ {
                            v33 := v4[fi27]
                            if false {
                                v34 := v4[fi32]
                                _ = v34
                            } else {
                                var v35 uint = 210
                                _ = v35
                            }
                            v26 = (3.14159 - v26)
                            _ = fi32
                            _ = v33
                        }
                    default:
                        var v36 uint32 = 291
                        _ = v36
                    }
                    v37 := v1[4294967295]
                    _ = v37
                }
            case 128:
                v2 = (v2 * (v2 * 42))
            case -717:
                var v38 float64 = ((((-2.25 * 0.1) + 1.5) + (((1.5 / 1.0) + 1e10) / ((float64(v26) - float64(float32(0.0))) + (3.14159 / 1e10)))) * ((1.5 * 3.14159) / (1.5 - -2.25)))
                _ = v38
            default:
                fmt.Println(v2, fi27, fi27)
            }
            switch ("foo bar" + ("a" + ("foo bar" + "a"))) {
            case "a":
                v39 := GenS1{F0: 127, F1: 103}
                _ = v39
            default:
                var v40 float64 = (-2.25 - 0.0)
                _ = v40
            }
            _ = fi27
        }
    }
    var v41 float32 = ((v26 / ((1.5 - (v26 / v26)) * v26)) - (v26 - 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v41
    _ = v5
}
