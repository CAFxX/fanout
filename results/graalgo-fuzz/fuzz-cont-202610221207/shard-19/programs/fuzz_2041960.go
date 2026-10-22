package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint64 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
func genFn0(p0 int64, p1 int64) uint {
    for fi7 := 0; fi7 < 18; fi7++ {
        switch (fi7 | (fi7 | fi7)) {
        case 2147483647:
            v8 := GenS1{F0: 42}
            _ = v8
        }
        _ = fi7
    }
    v9 := GenS1{F0: -2147483648}
    v10 := GenS1{F0: 2147483647}
    var v11 int = -780
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 1
}
func genFn1() uint16 {
    switch 223 {
    case 42:
        var v12 int64 = 0
        _ = v12
    case -295:
        var v13 int8 = 52
        _ = v13
    case -9223372036854775808:
        v14 := []int{128, 979, 127, 128}
        _ = v14
    }
    v15 := GenS2{F0: 127, F1: false, F2: 4294967295}
    switch "foo bar" {
    case "a":
        var v16 float64 = (1e10 * 1e10)
        _ = v16
    case "\n\t":
        var v17 int8 = 42
        _ = v17
    case "":
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 bool = (255 != 15)
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := []int{fi18, fi20, fi20, (fi18 - (fi20 * fi18))}
                fmt.Println(fi18, v19, fi20)
                _ = fi20
                _ = v21
            }
            _ = fi18
            _ = v19
        }
    }
    v22 := []int{42}
    var v23 uint32 = 4294967295
    _ = v15
    _ = v22
    _ = v23
    return 608
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]uint32)
    var v3 uint = 255
    v4 := make(map[string]bool)
    var v5 int64 = 1
    v6 := make(map[string]uint32)
    for fi24 := 0; fi24 < 10; fi24++ {
        var v25 float64 = (0.0 - 0.1)
        var v26 int64 = (v5 - v5)
        v25 = (v25 + v25)
        _ = fi24
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: 1}
    for fi28 := 0; fi28 < 6; fi28++ {
        delete(v2, "\n\t")
        _ = fi28
    }
    v29 := GenS2{F0: 0, F1: true, F2: 571}
    for fi30 := 0; fi30 < 13; fi30++ {
        if (2147483648 <= 863) {
            v31 := GenS3{F0: 42, F1: 65535, F2: (v3 & v3)}
            v32 := GenS2{F0: uint8(v3), F1: true, F2: 884}
            _ = v31
            _ = v32
        } else {
            var v33 int16 = -32768
            v3 = v3
            v34 := v1[fi30]
            _ = v33
            _ = v34
        }
        v35 := []int{}
        for fi36 := 0; fi36 < 18; fi36++ {
            v37 := GenS1{F0: int32(fi36)}
            _ = fi36
            _ = v37
        }
        _ = fi30
        _ = v35
    }
    v3 = 155
    v5 = v5
    fmt.Println(v5)
    var v38 bool = true
    v39 := []int{}
    switch 715 {
    case 653:
        if len(v39) >= 2 {
            v39 = v39[1:len(v39)]
        }
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        v39 = append(v39, (fi40 - (fi40 & fi40)))
        v3 -= genFn0((-(v5)), (v5 | -9223372036854775808))
        if ((v3 >> ((v3) & 63)) != v3) {
            var v41 int = fi40
            for fi42 := 0; fi42 < 7; fi42++ {
                v38 = (!false)
                _ = fi42
            }
            _ = v41
        } else {
            v43 := GenS4{F0: -32768}
            if v38 {
                for fi44 := 0; fi44 < 8; fi44++ {
                    if false {
                        v45 := GenS3{F0: (^(v3)), F1: 65535, F2: v3}
                        if len(v39) >= 2 {
                            v39 = v39[1:len(v39)]
                        }
                        if (false && (0.0 > 3.14159)) {
                            v46 := GenS1{F0: 609}
                            v5 = (^(v5))
                            switch (fi44 / ((-2147483648) | 1)) {
                            case 127:
                                v47 := GenS2{F0: 167, F1: (!true), F2: 2}
                                _ = v47
                            default:
                                var v48 int8 = -128
                                _ = v48
                            }
                            _ = v46
                        } else {
                            v3 = v3
                            var v49 uint16 = 1
                            var v50 uint = (v3 % (((v3 + (^(v3)))) | 1))
                            _ = v49
                            _ = v50
                        }
                        v38 = (1.5 != 0.0)
                        _ = v45
                    }
                    if len(v39) > 0 {
                        v39[0] = (-(fi44))
                    }
                    _ = fi44
                }
                v51 := GenS5{F0: v38}
                _ = v51
            }
            _ = v43
        }
        _ = fi40
    }
    if true {
        var v52 bool = (!(!(true || (v38 || false))))
        _ = v52
    } else {
        if v38 {
            var v53 uint8 = 2
            _ = v53
        }
    }
    fmt.Println(v3, v38)
    if false {
        v38 = (0 >= 0)
    } else {
        fmt.Println(v38)
    }
    var v54 bool = (!false)
    v55 := GenS2{F0: 154, F1: true, F2: 1}
    v56 := GenS2{F0: 255, F1: v38, F2: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v6
}
