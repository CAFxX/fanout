package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 int) uint32 {
    fmt.Println(p0)
    switch 42 {
    case 127:
        p0 = (p0 ^ p0)
    case 42:
        p0 = (p0 / ((p0) | 1))
    case -2147483648:
        for fi6 := 0; fi6 < 17; fi6++ {
            switch "foo bar" {
            case "":
                for fi7 := 0; fi7 < 9; fi7++ {
                    v8 := []int{fi7, (p0 << ((202) & 63))}
                    fmt.Println(fi7, p0, fi7)
                    fmt.Println(fi6, p0, fi6)
                    _ = fi7
                    _ = v8
                }
            case "\n\t":
                switch p0 {
                case 32767:
                    var v9 int8 = 37
                    _ = v9
                case 42:
                    var v10 uint32 = 65535
                    _ = v10
                default:
                    for fi11 := 0; fi11 < 16; fi11++ {
                        v12 := []int{p0}
                        var v13 uint = 0
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            }
            var v14 int64 = 1
            if (42 == 2) {
                v15 := GenS1{F0: 4294967295}
                var v16 float64 = 0.1
                p0 = 940
                _ = v15
                _ = v16
            } else {
                p0 = p0
                var v17 float32 = 1e10
                fmt.Println(p0)
                _ = v17
            }
            _ = fi6
            _ = v14
        }
    default:
        var v18 int32 = 0
        _ = v18
    }
    if true {
        p0 += (p0 / (((^(p0))) | 1))
    }
    v19 := []int{-638, (p0 + -9223372036854775808), 1, -2147483648}
    for fi20 := 0; fi20 < 11; fi20++ {
        p0 -= fi20
        p0 = (fi20 | (^(fi20)))
        _ = fi20
    }
    _ = p0
    _ = v19
    return 42
}

func main() {
    var v1 uint16 = 563
    v2 := []uint32{538, 65535}
    var v3 string = "hello"
    var v4 uint = 748
    v5 := make(map[uint32]bool)
    var v21 int32 = 42
    switch -2147483648 {
    case 32767:
        fmt.Println(v3, v1, v4)
    default:
        v1 -= (v1 + 110)
    }
    v2 = append(v2, 2)
    fmt.Println(v4)
    if (-1 < 42) {
        var v22 int32 = (v21 << (((v4 & (v4 % ((v4) | 1)))) & 63))
        v4 -= v4
        _ = v22
    }
    if false {
        var v23 int64 = -326
        fmt.Println(v3)
        _ = v23
    } else {
        switch (v3 + (v3 + ("hello" + v3))) {
        case "foo bar":
            if false {
                var v24 float64 = 0.0
                v25 := GenS2{F0: 656}
                _ = v24
                _ = v25
            } else {
                var v26 float64 = (1e10 + float64(0.0))
                _ = v26
            }
        case "hello":
            var v27 string = (v3 + ((v3 + (v3 + "a")) + v3))
            _ = v27
        default:
            v3 = v3
        }
        v28 := GenS2{F0: (v1 | (v1 - v1))}
        var v29 int16 = 627
        _ = v28
        _ = v29
    }
    fmt.Println(v3)
    switch "a" {
    case "":
        v30 := GenS3{F0: 375}
        _ = v30
    case "foo bar":
        if true {
            var v31 uint16 = v1
            switch 2147483647 {
            case -398:
                var v32 int = 32767
                _ = v32
            default:
                if (uint(v1) > 42) {
                    v33 := GenS4{F0: false, F1: (v4 % ((255) | 1))}
                    _ = v33
                }
            }
            v3 = "foo bar"
            _ = v31
        } else {
            v34 := v5[0]
            if v34 {
                if ((!v34) && (-2.25 > -2.25)) {
                    if ((!(1 == -916)) || (!(false || (v1 == 156)))) {
                        if v34 {
                            fmt.Println(v1, v4)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                    }
                    if v34 {
                        fmt.Println(v34, v1, v4)
                        var v35 int32 = (-(v21))
                        v36 := GenS4{F0: false, F1: v4}
                        _ = v35
                        _ = v36
                    }
                    v34 = v34
                } else {
                    v37 := GenS1{F0: 2147483648}
                    var v38 int16 = -701
                    var v39 string = "a"
                    _ = v37
                    _ = v38
                    _ = v39
                }
                var v40 int32 = (v21 / ((int32(v1)) | 1))
                _ = v40
            } else {
                var v41 uint64 = 1
                v42 := GenS4{F0: v34, F1: v4}
                _ = v41
                _ = v42
            }
            _ = v34
        }
    }
    var v43 int16 = 42
    v4 = v4
    fmt.Println(v21)
    var v44 float32 = -2.25
    for fi45 := 0; fi45 < 14; fi45++ {
        fmt.Println(v43)
        v46 := GenS5{F0: 30, F1: (fi45 + fi45)}
        _ = fi45
        _ = v46
    }
    v47 := GenS1{F0: 4294967295}
    v3 = ""
    if (127 <= int(v43)) {
        switch ("\n\t" + (((((((("foo bar" + v3) + "foo bar") + ((v3 + v3) + v3)) + ((v3 + (v3 + v3)) + v3)) + ((v3 + "hello") + v3)) + v3) + ("\n\t" + v3)) + v3)) {
        case "":
            switch v3 {
            case "a":
                var v48 int16 = 32767
                _ = v48
            case "":
                var v49 uint8 = 255
                _ = v49
            }
        case "hello":
            var v50 uint16 = (^(v1))
            _ = v50
        case "\n\t":
            if true {
                v21 = (v21 - v21)
            }
        }
        v5[genFn0(-510)] = (42 >= 42)
        var v51 int16 = v43
        _ = v51
    }
    var v52 uint = (v4 + (v4 - (v4 * (v4 % ((v4) | 1)))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v47
    _ = v5
    _ = v52
}
