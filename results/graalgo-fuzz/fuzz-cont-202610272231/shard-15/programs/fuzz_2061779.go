package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 int8; F1 int16 }
type GenS4 struct { F0 string; F1 int64; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint64; F2 int32 }
func genFn0(p0 int8, p1 uint64) string {
    v4 := GenS1{F0: 961, F1: (^(p0))}
    v5 := GenS1{F0: 42, F1: p0}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return "hello"
}

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -584
    v3 := make(map[int]bool)
    if (!(!(true || false))) {
        var v6 int64 = 1
        _ = v6
    }
    var v7 uint = 255
    var v8 int32 = v2
    fmt.Println(v2, v2)
    fmt.Println(v8, v7, v7)
    var v9 int32 = (-(v8))
    if (127 == 139) {
        if (32767 < 42) {
            fmt.Println(v7, v8, v9)
            var v10 int16 = -954
            v11 := GenS2{F0: 127, F1: 0}
            _ = v10
            _ = v11
        } else {
            var v12 bool = true
            fmt.Println(v12, v8)
            _ = v12
        }
    }
    if true {
        if ((2147483647 >= 4294967295) || (-1 != 42)) {
            v9 = (v8 + v2)
            for fi13 := 0; fi13 < 15; fi13++ {
                var v14 uint64 = 263
                var v15 int8 = -1
                _ = fi13
                _ = v14
                _ = v15
            }
            if (!(-42 >= -47)) {
                for fi16 := 0; fi16 < 18; fi16++ {
                    var v17 uint64 = 474
                    _ = fi16
                    _ = v17
                }
            } else {
                _ = len(v1)
                if (42 != 152) {
                    fmt.Println(v8, v2)
                    if true {
                        if (!((v7 / ((1) | 1)) >= v7)) {
                            v2 -= (v9 - (v8 - 42))
                            v18 := GenS2{F0: -2147483648, F1: 0}
                            _ = v18
                        }
                    } else {
                        delete(v1, "a")
                        var v19 int32 = v9
                        for fi20 := 0; fi20 < 15; fi20++ {
                            v21 := GenS3{F0: 42, F1: -648}
                            fmt.Println(fi20, v2)
                            switch "a" {
                            case "foo bar":
                                fmt.Println(v8)
                            case "\n\t":
                                v2 = 1
                            case "hello":
                                if false {
                                    v19 = (v19 / ((567) | 1))
                                    if true {
                                        var v22 int8 = 59
                                        switch ("foo bar" + "") {
                                        case "foo bar":
                                            v23 := GenS4{F0: ("hello" + ("\n\t" + ("a" + (("hello" + "foo bar") + "\n\t")))), F1: 9223372036854775807, F2: 4294967295}
                                            _ = v23
                                        default:
                                            if false {
                                                if (!(false || true)) {
                                                    v2 = v2
                                                    v9 -= (v9 / (((v2 | 731)) | 1))
                                                    fmt.Println(fi20)
                                                }
                                                switch fi20 {
                                                case 32767:
                                                    fmt.Println(v7)
                                                default:
                                                    for fi24 := 0; fi24 < 19; fi24++ {
                                                        fmt.Println(fi20, v9)
                                                        var v25 bool = false
                                                        switch 128 {
                                                        case 42:
                                                            switch (fi20 | fi24) {
                                                            case -9223372036854775808:
                                                                if v25 {
                                                                    if true {
                                                                        switch fi24 {
                                                                        case 443:
                                                                            v7 -= (v7 * 854)
                                                                        }
                                                                        var v26 uint64 = 1
                                                                        v27 := GenS4{F0: "\n\t", F1: int64(v2), F2: uint32(v7)}
                                                                        _ = v26
                                                                        _ = v27
                                                                    } else {
                                                                        var v28 int = fi20
                                                                        v29 := []int{-9223372036854775808, (^(fi20)), fi24}
                                                                        if true {
                                                                            var v30 int16 = 846
                                                                            if v25 {
                                                                                var v31 int = fi24
                                                                                fmt.Println(v30)
                                                                                _ = v31
                                                                            } else {
                                                                                var v32 string = ("foo bar" + ("foo bar" + ""))
                                                                                v33 := GenS5{F0: 42}
                                                                                _ = v32
                                                                                _ = v33
                                                                            }
                                                                            _ = v30
                                                                        }
                                                                        _ = v28
                                                                        _ = v29
                                                                    }
                                                                    v34 := []int{-9223372036854775808, fi20, (fi20 % ((fi24) | 1)), fi24}
                                                                    _ = v34
                                                                } else {
                                                                    v35 := []int{fi20, 2147483647}
                                                                    _ = v35
                                                                }
                                                            }
                                                        case 1:
                                                            var v36 int32 = (v19 % ((int32(v22)) | 1))
                                                            _ = v36
                                                        }
                                                        _ = fi24
                                                        _ = v25
                                                    }
                                                }
                                            }
                                        }
                                        _ = v22
                                    } else {
                                        if true {
                                            fmt.Println(fi20)
                                        }
                                    }
                                    fmt.Println(v9)
                                }
                            default:
                                v37 := GenS6{F0: false, F1: 4294967295, F2: (v8 + 2147483647)}
                                _ = v37
                            }
                            _ = fi20
                            _ = v21
                        }
                        _ = v19
                    }
                } else {
                    v38 := GenS4{F0: ("" + ("foo bar" + "\n\t")), F1: 0, F2: 42}
                    v39 := []int{}
                    _ = v38
                    _ = v39
                }
            }
        }
        fmt.Println(v9, v8)
    } else {
        for fi40 := 0; fi40 < 9; fi40++ {
            fmt.Println(v7)
            for fi41 := 0; fi41 < 13; fi41++ {
                var v42 int8 = 100
                fmt.Println(v42, v8)
                var v43 float64 = float64(0.1)
                _ = fi41
                _ = v42
                _ = v43
            }
            _ = fi40
        }
        v44 := []int{32767}
        _ = v44
    }
    for fi45 := 0; fi45 < 9; fi45++ {
        if ((!(1.5 != 0.1)) || false) {
            var v46 uint32 = 4294967295
            v47 := GenS1{F0: (v2 << (((v7 % ((v7) | 1))) & 63)), F1: 127}
            _ = v46
            _ = v47
        }
        var v48 uint64 = 18446744073709551615
        v49 := []int{(fi45 / ((fi45) | 1)), 2147483647, fi45}
        _ = fi45
        _ = v48
        _ = v49
    }
    for fi50 := 0; fi50 < 17; fi50++ {
        v51 := []int{fi50}
        _ = len(v3)
        var v52 uint16 = 255
        _ = fi50
        _ = v51
        _ = v52
    }
    for fi53 := 0; fi53 < 4; fi53++ {
        v54 := []int{(fi53 * fi53), (-(fi53))}
        var v55 string = "foo bar"
        var v56 uint = 4294967295
        _ = fi53
        _ = v54
        _ = v55
        _ = v56
    }
    if (!(732 >= v7)) {
        var v57 float32 = ((0.1 * (float32(((0.0 + 1.5) * ((((float64(((1.5 + 0.1) - -2.25)) - (3.14159 - ((0.0 - 0.1) + 1e10))) + 0.1) / 1.0) - (1e10 + -2.25)))) * 3.14159)) + -2.25)
        _ = v57
    } else {
        v7 += v7
    }
    v2 = (v9 - (v8 ^ (v9 * (^(v8)))))
    if true {
        v58 := GenS3{F0: 1, F1: int16(v7)}
        _ = v58
    }
    for fi59 := 0; fi59 < 2; fi59++ {
        v60 := []int{}
        v9 -= (^(v8))
        delete(v3, fi59)
        _ = fi59
        _ = v60
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        fmt.Println(fi61, v8)
        if true {
            var v62 int = (-(fi61))
            _ = v62
        } else {
            var v63 int64 = 0
            _ = v63
        }
        var v64 int16 = 1000
        _ = fi61
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
