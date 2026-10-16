package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 bool; F1 bool }
type GenS3 struct { F0 int32; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int16; F1 uint8; F2 uint }
type GenS5 struct { F0 uint8; F1 int16; F2 int64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 int16; F1 int64; F2 uint8 }
type GenS8 struct { F0 uint8; F1 bool; F2 int16 }
func genFn0(p0 uint32) int32 {
    if (2147483647 > 1) {
        if true {
            if ((true && false) && true) {
                fmt.Println(p0)
            } else {
                var v6 int64 = 9223372036854775807
                _ = v6
            }
        } else {
            v7 := GenS1{F0: 980, F1: "foo bar"}
            v8 := GenS2{F0: (42 > 1000), F1: true}
            var v9 int16 = -819
            _ = v7
            _ = v8
            _ = v9
        }
        var v10 uint16 = 0
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if true {
            for fi12 := 0; fi12 < 15; fi12++ {
                var v13 int64 = -1
                if (-2.25 <= ((((0.0 / (1e10 - ((-2.25 + float32(1e10)) + (((3.14159 + 1.5) - 3.14159) + ((-2.25 / (1e10 * 0.1)) - -2.25))))) + 3.14159) * 0.0) - 1.5)) {
                    if false {
                        var v14 uint = 4294967295
                        _ = v14
                    } else {
                        v13 = (v13 | v13)
                        switch ("foo bar" + "\n\t") {
                        case "\n\t":
                            p0 -= (p0 + p0)
                        default:
                            v15 := GenS1{F0: 4294967295, F1: "\n\t"}
                            v16 := []int{fi11, fi11, (fi12 & (^(fi11)))}
                            if len(v16) > 0 {
                                v16[0] = (fi12 - 32767)
                            }
                            _ = v15
                            _ = v16
                        }
                    }
                    p0 += (p0 - p0)
                }
                switch "a" {
                case "":
                    var v17 float32 = 3.14159
                    _ = v17
                case "foo bar":
                    if (false || false) {
                        switch "" {
                        case "foo bar":
                            var v18 int8 = 0
                            _ = v18
                        case "a":
                            v19 := GenS3{F0: 2147483647, F1: 2, F2: 373}
                            _ = v19
                        case "hello":
                            v13 = (v13 & v13)
                        default:
                            fmt.Println(fi11, p0, fi12)
                        }
                        p0 = 0
                    } else {
                        v20 := GenS4{F0: 1000, F1: 2, F2: 42}
                        var v21 int64 = (^(v13))
                        p0 = p0
                        _ = v20
                        _ = v21
                    }
                case "a":
                    v22 := []int{fi11, (fi12 | fi12), fi12, -9223372036854775808}
                    _ = v22
                }
                _ = fi12
                _ = v13
            }
            fmt.Println(fi11, fi11)
        }
        _ = fi11
    }
    var v23 uint64 = 2
    _ = p0
    _ = v23
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]uint32)
    var v3 uint8 = 171
    v4 := []string{"", "", "", (("hello" + (("" + "") + "hello")) + ("foo bar" + ("\n\t" + "foo bar")))}
    var v5 int16 = 534
    var v24 int32 = 2147483647
    var v25 string = "\n\t"
    v26 := GenS5{F0: v3, F1: (v5 / ((v5) | 1)), F2: -9223372036854775808}
    var v27 int16 = (v5 - 204)
    if (943 == 0) {
        switch 319 {
        case 127:
            var v28 int = -286
            _ = v28
        case -1:
            var v29 float64 = -2.25
            _ = v29
        case 1:
            var v30 uint = 0
            _ = v30
        }
        v31 := GenS5{F0: (v3 << ((18446744073709551615) & 63)), F1: v5, F2: -1}
        v32 := GenS2{F0: (true || (true && true)), F1: true}
        _ = v31
        _ = v32
    } else {
        switch "a" {
        case "\n\t":
            v27 = v27
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            fmt.Println(v27, v25, v25)
            var v34 float64 = 3.14159
            v24 += genFn0(4294967295)
            _ = fi33
            _ = v34
        }
        v35 := GenS6{F0: -504, F1: v24}
        _ = v35
    }
    switch 128 {
    case -613:
        fmt.Println(v27, v5, v5)
    case 0:
        var v36 int32 = (-(v24))
        _ = v36
    }
    fmt.Println(v25, v3, v3)
    v27 = -32768
    if (!(32768 == -486)) {
        v5 -= (v27 << ((18446744073709551615) & 63))
        var v37 int16 = (v5 ^ (v5 | v27))
        _ = v37
    } else {
        fmt.Println(v25)
    }
    v27 = -818
    for fi38 := 0; fi38 < 15; fi38++ {
        v4 = append(v4, v25)
        switch fi38 {
        case 128:
            v39 := GenS2{F0: true, F1: (fi38 > fi38)}
            _ = v39
        case 122:
            v24 = (v24 / ((genFn0(2147483648)) | 1))
        case 2147483647:
            if (9223372036854775807 >= 42) {
                var v40 int = -1
                var v41 float32 = (3.14159 + float32(-2.25))
                var v42 bool = true
                _ = v40
                _ = v41
                _ = v42
            } else {
                fmt.Println(v5, fi38, v25)
                v43 := GenS7{F0: -32768, F1: 42, F2: (^(v3))}
                var v44 int64 = 9223372036854775807
                _ = v43
                _ = v44
            }
        }
        v5 = v27
        _ = fi38
    }
    _ = len(v1)
    v24 = (v24 + 865)
    for fi45 := 0; fi45 < 15; fi45++ {
        fmt.Println(v27, v24, v25)
        var v46 uint = 65535
        for fi47 := 0; fi47 < 14; fi47++ {
            var v48 uint64 = 42
            if false {
                switch fi47 {
                case 329:
                    v27 += v27
                case 32768:
                    v49 := GenS6{F0: -9223372036854775808, F1: (v24 >> (((^(v46))) & 63))}
                    _ = v49
                case 0:
                    v3 = v3
                default:
                    v50 := GenS8{F0: uint8(v46), F1: (v25 == ("foo bar" + ((v25 + "\n\t") + (v25 + v25)))), F2: (-(v5))}
                    _ = v50
                }
                _ = len(v4)
            } else {
                fmt.Println(fi45)
                var v51 uint8 = v3
                _ = v51
            }
            _ = fi47
            _ = v48
        }
        _ = fi45
        _ = v46
    }
    if len(v4) > 0 {
        v4[0] = v25
    }
    var v52 uint64 = 4294967295
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v52
}
