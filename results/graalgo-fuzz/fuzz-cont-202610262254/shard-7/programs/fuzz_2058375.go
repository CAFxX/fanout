package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint8; F1 int64; F2 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int32; F1 uint32; F2 uint64 }
func genFn0(p0 string, p1 bool) int {
    fmt.Println(p1, p0)
    var v6 uint16 = 32767
    v7 := []int{42, -9223372036854775808, -788}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 436
}

func main() {
    var v1 float64 = (1e10 + (0.0 + 1e10))
    var v2 string = "\n\t"
    v3 := make(map[string]bool)
    var v4 bool = (!((2 <= 18446744073709551615) && false))
    v5 := make(map[string]string)
    var v8 int32 = -2147483648
    for fi9 := 0; fi9 < 16; fi9++ {
        for fi10 := 0; fi10 < 16; fi10++ {
            if (v4 || v4) {
                v11 := GenS1{F0: -116, F1: 42, F2: 64}
                var v12 string = v2
                if (1 != 170) {
                    if (262 >= 0) {
                        if v4 {
                            v13 := GenS1{F0: 55, F1: 42, F2: -9223372036854775808}
                            v14 := GenS2{F0: 4294967295, F1: 18446744073709551615}
                            _ = v13
                            _ = v14
                        }
                        var v15 int16 = -32768
                        if (42 <= 42) {
                            v16 := GenS2{F0: 42, F1: 42}
                            v4 = v4
                            switch fi10 {
                            case 128:
                                var v17 int64 = -1
                                _ = v17
                            case 42:
                                fmt.Println(fi9)
                            case 32768:
                                switch fi10 {
                                case -475:
                                    v18 := []int{fi9}
                                    _ = v18
                                }
                            default:
                                var v19 bool = (!v4)
                                _ = v19
                            }
                            _ = v16
                        } else {
                            if (139 < 0) {
                                var v20 uint64 = 48
                                v21 := GenS1{F0: int8(v8), F1: v15, F2: 42}
                                _ = v20
                                _ = v21
                            }
                            if ((523 >= 0) || (false || (845 != 4294967295))) {
                                v4 = ((1 > 32768) && (fi10 != fi9))
                                var v22 int64 = 9223372036854775807
                                v23 := []int{(fi9 | (fi10 - fi9)), (fi10 - genFn0(v12, false)), fi9}
                                delete(v3, v12)
                                _ = v22
                                _ = v23
                            } else {
                                v2 = v2
                                fmt.Println(fi10, v8, v15)
                                v24 := GenS1{F0: 42, F1: -262, F2: 42}
                                _ = v24
                            }
                            v1 += -2.25
                        }
                        _ = v15
                    } else {
                        switch (fi9 * genFn0("a", (1 != 1))) {
                        case 9223372036854775807:
                            _ = len(v3)
                        default:
                            var v25 uint8 = 1
                            var v26 float64 = v1
                            _ = v25
                            _ = v26
                        }
                    }
                }
                if (v4 || (1 != -32768)) {
                    if true {
                        switch (fi10 << ((453) & 63)) {
                        case -417:
                            v4 = (127 >= 226)
                        case -2147483648:
                            v27 := GenS3{F0: 2, F1: 0, F2: -124}
                            _ = v27
                        default:
                            v28 := GenS1{F0: 118, F1: -32768, F2: 0}
                            _ = v28
                        }
                        v29 := GenS4{F0: 1000}
                        _ = v29
                    } else {
                        v30 := []int{}
                        var v31 float32 = ((-2.25 - 0.0) / float32((1e10 + 1e10)))
                        v32 := GenS3{F0: 10, F1: 9223372036854775807, F2: 42}
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    fmt.Println(v12, fi10)
                    v33 := GenS5{F0: 900}
                    _ = v33
                }
                v34 := v3[v12]
                var v35 int16 = 32767
                if (995 != 0) {
                    var v36 int64 = 42
                    var v37 int32 = 928
                    _ = v36
                    _ = v37
                }
                _ = v11
                _ = v12
                _ = v34
                _ = v35
            } else {
                v8 -= v8
            }
            _ = fi10
        }
        _ = fi9
    }
    switch "foo bar" {
    case "":
        var v38 uint32 = 2147483648
        _ = v38
    case "hello":
        v39 := []int{genFn0(((v2 + (v2 + ("foo bar" + v2))) + v2), v4)}
        _ = v39
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        var v41 int16 = -32768
        _ = fi40
        _ = v41
    }
    fmt.Println(v4, v1)
    for fi42 := 0; fi42 < 5; fi42++ {
        v43 := GenS6{F0: v8, F1: 4294967295, F2: 18446744073709551615}
        fmt.Println(v8, v4)
        _ = fi42
        _ = v43
    }
    var v44 int32 = (v8 + (-(v8)))
    for fi45 := 0; fi45 < 20; fi45++ {
        switch fi45 {
        case 0:
            fmt.Println(v2, fi45, v4)
        case -344:
            if false {
                v46 := []int{-874, (fi45 - (fi45 * fi45))}
                _ = v46
            }
        case -311:
            var v47 int8 = -128
            _ = v47
        }
        v48 := []int{genFn0(v2, v4)}
        if v4 {
            if true {
                v49 := GenS5{F0: 9223372036854775807}
                v1 = (v1 + v1)
                v50 := v3[v2]
                _ = v49
                _ = v50
            }
        } else {
            v8 += (v44 * 1)
            fmt.Println(v44, v1)
            var v51 bool = v4
            _ = v51
        }
        _ = fi45
        _ = v48
    }
    var v52 int32 = (^(v8))
    if v4 {
        for fi53 := 0; fi53 < 16; fi53++ {
            v2 = v2
            _ = fi53
        }
    } else {
        v54 := []int{9223372036854775807, 4294967295, -2147483648, genFn0(v2, (1 <= 9223372036854775807))}
        _ = v54
    }
    switch -515 {
    case 128:
        var v55 uint64 = 42
        _ = v55
    case 127:
        for fi56 := 0; fi56 < 19; fi56++ {
            fmt.Println(v4)
            if ((42 > 0) && (42 > 4294967295)) {
                var v57 int8 = 20
                v44 -= v8
                v2 = v2
                _ = v57
            } else {
                v52 -= v44
                var v58 uint16 = 42
                var v59 int32 = (v52 << ((564) & 63))
                _ = v58
                _ = v59
            }
            switch (fi56 * genFn0(v2, true)) {
            case -566:
                var v60 int = (fi56 << ((255) & 63))
                _ = v60
            case 349:
                v8 = v44
            default:
                v2 = ("hello" + v2)
            }
            _ = fi56
        }
    }
    v4 = ((float64(0.0) / (float64(0.0) / ((v1 - v1) / 3.14159))) > v1)
    for fi61 := 0; fi61 < 13; fi61++ {
        var v62 uint = 812
        _ = fi61
        _ = v62
    }
    v8 += (v44 ^ (v44 % (((-(v44))) | 1)))
    v63 := GenS2{F0: 18446744073709551615, F1: 2}
    fmt.Println(v44)
    var v64 int64 = 141
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v52
    _ = v63
    _ = v64
    _ = v8
}
