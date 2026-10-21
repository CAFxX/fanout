package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 int }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 bool) bool {
    v6 := GenS1{F0: p0, F1: 4294967295}
    p0 = p0
    v7 := []int{1, -94}
    v8 := GenS1{F0: (2 == 18446744073709551615), F1: 939}
    fmt.Println(p0)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return true
}

func main() {
    var v1 uint32 = 65535
    v2 := []string{((("\n\t" + "hello") + "a") + (("a" + "foo bar") + "")), "foo bar"}
    var v3 int16 = 978
    v4 := make(map[int]bool)
    var v5 uint = 521
    var v9 uint16 = uint16(v5)
    var v10 float64 = 0.0
    for fi11 := 0; fi11 < 2; fi11++ {
        fmt.Println(v3)
        if ((0.0 > v10) || (v1 == (v1 << (((v5 >> ((v5) & 63))) & 63)))) {
            var v12 uint8 = 128
            for fi13 := 0; fi13 < 19; fi13++ {
                var v14 uint16 = v9
                var v15 bool = false
                v9 = 352
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        _ = fi11
    }
    v9 += v9
    for fi16 := 0; fi16 < 20; fi16++ {
        for fi17 := 0; fi17 < 19; fi17++ {
            v3 = (v3 - int16(fi16))
            var v18 uint64 = 1
            _ = fi17
            _ = v18
        }
        v9 = v9
        _ = fi16
    }
    v19 := GenS2{F0: -2147483648, F1: 0}
    switch 0 {
    case 4294967295:
        for fi20 := 0; fi20 < 16; fi20++ {
            for fi21 := 0; fi21 < 7; fi21++ {
                delete(v4, 128)
                fmt.Println(fi20, v5)
                var v22 int32 = 42
                _ = fi21
                _ = v22
            }
            var v23 float32 = 1e10
            switch "\n\t" {
            case "a":
                var v24 int16 = 0
                _ = v24
            case "foo bar":
                v25 := v4[(fi20 & (fi20 - (^(fi20))))]
                _ = v25
            case "hello":
                switch fi20 {
                case 32768:
                    v26 := GenS1{F0: true, F1: (v1 << (((^(v5))) & 63))}
                    _ = v26
                case 505:
                    v9 -= 2
                default:
                    v10 -= v10
                }
            }
            _ = fi20
            _ = v23
        }
    case -375:
        v9 += 2
    default:
        fmt.Println(v1)
    }
    var v27 int32 = 0
    if ((v9 / ((v9) | 1)) == (^(v9))) {
        var v28 int16 = (-(v3))
        _ = v28
    } else {
        v29 := v4[2147483647]
        for fi30 := 0; fi30 < 18; fi30++ {
            for fi31 := 0; fi31 < 17; fi31++ {
                if v29 {
                    var v32 uint32 = 95
                    _ = v32
                }
                _ = fi31
            }
            _ = fi30
        }
        if (v29 && v29) {
            v5 -= 210
            v5 = v5
        }
        _ = v29
    }
    if true {
        v1 -= v1
        v5 += (v5 | 2)
    }
    for fi33 := 0; fi33 < 16; fi33++ {
        var v34 int64 = -1
        _ = fi33
        _ = v34
    }
    for fi35 := 0; fi35 < 9; fi35++ {
        var v36 uint8 = 0
        _ = fi35
        _ = v36
    }
    v37 := GenS3{F0: 2147483648, F1: (v27 & (v27 + int32(v9))), F2: 782}
    v10 = 0.1
    v9 -= (v9 % (((v9 - v9)) | 1))
    for fi38 := 0; fi38 < 9; fi38++ {
        if (fi38 >= fi38) {
            v2 = append(v2, "\n\t")
            var v39 uint16 = uint16(v1)
            _ = v39
        } else {
            delete(v4, 1)
        }
        for fi40 := 0; fi40 < 13; fi40++ {
            switch ("" + ((("\n\t" + "hello") + ((("\n\t" + "") + "a") + "hello")) + "\n\t")) {
            case "a":
                if false {
                    var v41 int8 = -29
                    if ((42 != 4294967295) || (!true)) {
                        if true {
                            fmt.Println(v41)
                            _ = len(v2)
                        }
                        v42 := GenS2{F0: v27, F1: v9}
                        _ = v42
                    } else {
                        v9 -= (v9 & 65535)
                        if (false && false) {
                            var v43 uint64 = 182
                            var v44 int8 = (-(v41))
                            _ = v43
                            _ = v44
                        }
                        var v45 uint64 = 2
                        _ = v45
                    }
                    v46 := GenS4{F0: (v9 + v9), F1: uint16(v1)}
                    v47 := GenS5{F0: 255, F1: 234}
                    _ = v41
                    _ = v46
                    _ = v47
                }
            case "\n\t":
                if true {
                    if (!false) {
                        var v48 float64 = (v10 + v10)
                        _ = v48
                    } else {
                        v49 := GenS6{F0: 1}
                        var v50 uint = v5
                        _ = v49
                        _ = v50
                    }
                } else {
                    v27 = (^(v27))
                    fmt.Println(fi40)
                }
            }
            if true {
                v5 = 65535
            } else {
                v51 := v4[2]
                var v52 float32 = (3.14159 / (float32((v10 - (v10 - 0.1))) + 1.5))
                _ = len(v4)
                _ = v51
                _ = v52
            }
            _ = fi40
        }
        _ = fi38
    }
    if ((!true) && true) {
        for fi53 := 0; fi53 < 13; fi53++ {
            v54 := GenS6{F0: 42}
            delete(v4, (fi53 % ((1) | 1)))
            _ = fi53
            _ = v54
        }
    } else {
        v3 += (v3 & 0)
        v3 = v3
        fmt.Println(v1, v9)
    }
    var v55 int32 = v27
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v55
    _ = v9
}
