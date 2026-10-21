package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint32; F1 int; F2 uint64 }
func genFn0(p0 int, p1 int, p2 uint32) int64 {
    p1 = p0
    v5 := []int{}
    var v6 uint64 = 4294967295
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 248
}
func genFn1(p0 int8, p1 int32, p2 int64) bool {
    v7 := []int{42, 2147483647, 127}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]string)
    var v3 int8 = 0
    var v4 int64 = 9223372036854775807
    var v8 uint8 = 116
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int64 = v4
        if false {
            v4 = v10
            delete(v2, "a")
        } else {
            v3 = v3
            fmt.Println(fi9, v8, v10)
        }
        _ = fi9
        _ = v10
    }
    v4 -= (v4 % (((v4 | (v4 ^ 0))) | 1))
    switch "" {
    case "\n\t":
        for fi11 := 0; fi11 < 3; fi11++ {
            v1[("foo bar" + "\n\t")] = 1
            _ = fi11
        }
    }
    if true {
        v8 = (v8 % ((37) | 1))
        var v12 uint8 = (v8 * (v8 << ((0) & 63)))
        switch 32768 {
        case 1:
            v13 := []int{}
            _ = v13
        case 513:
            v8 = 128
        default:
            v14 := GenS1{F0: (^(v12))}
            _ = v14
        }
        _ = v12
    } else {
        v15 := []int{-335}
        if true {
            if ((!(true || ((false && true) || false))) && (false || true)) {
                var v16 uint32 = 42
                _ = v16
            }
            if (1.5 < float64(0.1)) {
                v17 := GenS1{F0: 2}
                v18 := GenS2{F0: 1}
                _ = v17
                _ = v18
            }
            var v19 int16 = 1000
            _ = v19
        }
        v20 := GenS3{F0: 1000}
        _ = v15
        _ = v20
    }
    var v21 float64 = (((1e10 - 1e10) - 1.5) - 3.14159)
    v22 := GenS1{F0: (v8 ^ v8)}
    for fi23 := 0; fi23 < 18; fi23++ {
        if (-1 > 1) {
            var v24 string = ""
            _ = v24
        } else {
            fmt.Println(v8, v21, v4)
            var v25 int64 = -9223372036854775808
            _ = v25
        }
        _ = fi23
    }
    for fi26 := 0; fi26 < 6; fi26++ {
        fmt.Println(v4, v21, v8)
        v4 -= (v4 % (((v4 * genFn0(788, fi26, 779))) | 1))
        _ = fi26
    }
    v27 := []int{-1, 1, -160}
    v28 := GenS1{F0: (v8 ^ v8)}
    fmt.Println(v3, v21)
    if false {
        fmt.Println(v8)
        var v29 int32 = 0
        switch "foo bar" {
        case "a":
            for fi30 := 0; fi30 < 9; fi30++ {
                var v31 int16 = 42
                v8 += (v8 ^ v8)
                if (!true) {
                    for fi32 := 0; fi32 < 10; fi32++ {
                        fmt.Println(fi32, v21, v4)
                        if (123 <= 566) {
                            v33 := GenS2{F0: v31}
                            switch "\n\t" {
                            case "hello":
                                var v34 float64 = (((float64(1e10) * v21) + (float64(1.5) - float64(0.1))) * ((0.0 / 0.1) - v21))
                                _ = v34
                            }
                            v4 = (v4 / (((v4 << ((42) & 63))) | 1))
                            _ = v33
                        } else {
                            v35 := GenS2{F0: v31}
                            switch (fi32 & (fi32 & 9223372036854775807)) {
                            case 1:
                                v36 := GenS2{F0: 159}
                                _ = v36
                            }
                            _ = v35
                        }
                        _ = fi32
                    }
                } else {
                    fmt.Println(v4, fi30, v8)
                }
                _ = fi30
                _ = v31
            }
        case "":
            if true {
                _ = len(v27)
            } else {
                if (!true) {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        fmt.Println(v4, fi37, v8)
                        _ = fi37
                    }
                    for fi38 := 0; fi38 < 11; fi38++ {
                        v8 = v8
                        fmt.Println(fi38, fi38, v4)
                        v8 = (v8 | (v8 >> ((0) & 63)))
                        _ = fi38
                    }
                    for fi39 := 0; fi39 < 19; fi39++ {
                        var v40 uint32 = 4294967295
                        switch (fi39 * (fi39 - (fi39 ^ (fi39 | (fi39 + (fi39 % ((fi39) | 1))))))) {
                        case -777:
                            if len(v27) >= 2 {
                                v27 = v27[1:len(v27)]
                            }
                        case -783:
                            for fi41 := 0; fi41 < 8; fi41++ {
                                v42 := GenS4{F0: 444, F1: 1000}
                                v43 := GenS1{F0: (v8 << ((996) & 63))}
                                delete(v2, "hello")
                                _ = fi41
                                _ = v42
                                _ = v43
                            }
                        case -636:
                            v29 = (^(v29))
                        }
                        _ = fi39
                        _ = v40
                    }
                }
                v44 := GenS5{F0: 2, F1: -2147483648, F2: 4294967295}
                _ = v44
            }
        case "\n\t":
            fmt.Println(v21, v21, v29)
        default:
            if true {
                var v45 string = "foo bar"
                if len(v27) > 0 {
                    v27[0] = 127
                }
                _ = v45
            } else {
                for fi46 := 0; fi46 < 16; fi46++ {
                    var v47 float64 = 1.5
                    for fi48 := 0; fi48 < 1; fi48++ {
                        if false {
                            var v49 int64 = (^(v4))
                            fmt.Println(v3)
                            _ = v49
                        } else {
                            var v50 uint = 42
                            if len(v27) > 0 {
                                v27[0] = (fi46 / ((32767) | 1))
                            }
                            switch "\n\t" {
                            case "hello":
                                v51 := GenS1{F0: v8}
                                _ = v51
                            default:
                                switch (fi48 << ((v50) & 63)) {
                                case 32768:
                                    _ = len(v27)
                                case -912:
                                    fmt.Println(fi48, fi46)
                                case 9223372036854775807:
                                    v2["foo bar"] = "a"
                                default:
                                    v4 = int64(v50)
                                }
                            }
                            _ = v50
                        }
                        if len(v27) > 0 {
                            v27[0] = (^(fi48))
                        }
                        _ = fi48
                    }
                    _ = fi46
                    _ = v47
                }
            }
        }
        _ = v29
    } else {
        for fi52 := 0; fi52 < 8; fi52++ {
            switch "a" {
            case "\n\t":
                v4 -= genFn0(fi52, 9223372036854775807, 675)
            case "hello":
                v21 -= (v21 * v21)
            case "a":
                var v53 float64 = 0.0
                _ = v53
            }
            var v54 int8 = v3
            var v55 int16 = int16(v3)
            _ = fi52
            _ = v54
            _ = v55
        }
    }
    var v56 uint32 = 2
    v56 += v56
    v57 := GenS4{F0: -867, F1: 590}
    v4 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v8
}
