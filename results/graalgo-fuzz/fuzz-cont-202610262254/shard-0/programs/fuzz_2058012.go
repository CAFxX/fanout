package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 int32; F1 uint64; F2 int }
type GenS4 struct { F0 int8; F1 int64; F2 int64 }
type GenS5 struct { F0 int8; F1 uint8 }
type GenS6 struct { F0 bool }
func genFn0() uint {
    var v6 float32 = -2.25
    v7 := GenS1{F0: 65535}
    _ = v6
    _ = v7
    return 2
}
func genFn1(p0 int64, p1 uint8) int {
    if false {
        fmt.Println(p1, p1)
    } else {
        var v8 uint64 = 4294967295
        for fi9 := 0; fi9 < 17; fi9++ {
            if true {
                if ((-867 < -1) && false) {
                    v10 := GenS1{F0: 1}
                    for fi11 := 0; fi11 < 20; fi11++ {
                        if false {
                            v12 := GenS1{F0: 4294967295}
                            v13 := []int{(fi9 % (((fi9 >> ((1) & 63))) | 1)), fi9}
                            _ = v12
                            _ = v13
                        } else {
                            var v14 int16 = 853
                            switch (((("hello" + ("hello" + ("hello" + "hello"))) + ("\n\t" + (("" + "foo bar") + "\n\t"))) + "hello") + "hello") {
                            case "\n\t":
                                fmt.Println(fi9, fi9)
                            case "hello":
                                var v15 bool = (254 != 1)
                                _ = v15
                            default:
                                if false {
                                    v16 := GenS2{F0: 2, F1: 0}
                                    _ = v16
                                } else {
                                    v17 := []int{}
                                    v8 += v8
                                    _ = v17
                                }
                                if (true || (23 >= 0)) {
                                    v18 := []int{(fi9 & (fi9 % ((128) | 1)))}
                                    switch ("foo bar" + "\n\t") {
                                    case "\n\t":
                                        v19 := GenS3{F0: 0, F1: (v8 ^ v8), F2: (fi11 - fi9)}
                                        if (!false) {
                                            p1 = (^(p1))
                                            fmt.Println(p1, v14)
                                            v18 = append(v18, fi9)
                                        }
                                        var v20 int16 = v14
                                        _ = v19
                                        _ = v20
                                    default:
                                        if (-2.25 == 0.0) {
                                            p0 = -1
                                        }
                                    }
                                    _ = v18
                                } else {
                                    v21 := GenS2{F0: 0, F1: 663}
                                    v14 = -32768
                                    _ = v21
                                }
                            }
                            var v22 uint16 = 0
                            _ = v14
                            _ = v22
                        }
                        _ = fi11
                    }
                    _ = v10
                }
                v23 := GenS3{F0: 0, F1: v8, F2: fi9}
                v24 := []int{}
                _ = v23
                _ = v24
            } else {
                var v25 string = "a"
                v26 := GenS4{F0: 37, F1: (p0 % ((p0) | 1)), F2: p0}
                _ = v25
                _ = v26
            }
            _ = fi9
        }
        v27 := GenS2{F0: 42, F1: 556}
        _ = v27
        _ = v8
    }
    v28 := GenS1{F0: 0}
    for fi29 := 0; fi29 < 13; fi29++ {
        var v30 uint8 = p1
        p1 = 0
        _ = fi29
        _ = v30
    }
    _ = p0
    _ = p1
    _ = v28
    return -89
}

func main() {
    v1 := []uint32{}
    var v2 float64 = 1e10
    var v3 int64 = 1
    var v4 string = "hello"
    var v5 int = -2147483648
    v1 = append(v1, 73)
    switch v4 {
    case "foo bar":
        if false {
            var v31 uint32 = 805
            _ = v31
        }
    case "\n\t":
        v32 := GenS5{F0: -1, F1: 176}
        _ = v32
    case "hello":
        v3 -= (v3 - v3)
    default:
        switch v4 {
        case "a":
            v4 = "a"
        case "hello":
            if true {
                v5 -= genFn1((v3 / (((-(v3))) | 1)), 42)
                v4 = ""
            }
        }
    }
    var v33 uint16 = 65535
    var v34 int8 = -4
    switch "\n\t" {
    case "\n\t":
        switch v4 {
        case "hello":
            v1 = append(v1, 2147483647)
        case "a":
            v5 -= 0
        }
    }
    for fi35 := 0; fi35 < 18; fi35++ {
        if (true && (v34 >= 100)) {
            var v36 int = fi35
            fmt.Println(v2)
            _ = v36
        }
        var v37 uint8 = 28
        _ = fi35
        _ = v37
    }
    for fi38 := 0; fi38 < 13; fi38++ {
        for fi39 := 0; fi39 < 3; fi39++ {
            v33 = (v33 * v33)
            var v40 uint8 = 128
            fmt.Println(v33)
            _ = fi39
            _ = v40
        }
        _ = fi38
    }
    for fi41 := 0; fi41 < 16; fi41++ {
        v42 := GenS5{F0: (v34 + (v34 >> ((1) & 63))), F1: 2}
        for fi43 := 0; fi43 < 2; fi43++ {
            if true {
                if ((0.1 * ((1.5 / -2.25) / ((0.1 - -2.25) + ((float32(3.14159) - 1e10) - 0.1)))) != float32(0.1)) {
                    var v44 uint = 42
                    _ = len(v1)
                    _ = v44
                }
                v45 := GenS6{F0: (0 != 2147483647)}
                v34 -= v34
                _ = v45
            } else {
                v2 += float64(-2.25)
                v4 = (("hello" + "\n\t") + v4)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            var v46 float64 = 1e10
            _ = fi43
            _ = v46
        }
        _ = fi41
        _ = v42
    }
    for fi47 := 0; fi47 < 13; fi47++ {
        v5 = 32768
        v48 := GenS2{F0: (v33 * (v33 - v33)), F1: 18446744073709551615}
        v3 = -9223372036854775808
        _ = fi47
        _ = v48
    }
    var v49 float64 = ((v2 / 3.14159) + float64(1.5))
    for fi50 := 0; fi50 < 13; fi50++ {
        v4 = v4
        fmt.Println(v5)
        v3 = -9223372036854775808
        _ = fi50
    }
    v51 := GenS2{F0: 42, F1: 414}
    for fi52 := 0; fi52 < 3; fi52++ {
        for fi53 := 0; fi53 < 20; fi53++ {
            if (!(1 >= 42)) {
                var v54 float64 = float64(0.1)
                _ = v54
            }
            var v55 int64 = v3
            if false {
                switch (fi52 | (v5 | fi53)) {
                case -1:
                    switch (v4 + "a") {
                    case "\n\t":
                        fmt.Println(v3)
                    case "a":
                        var v56 uint = uint(v33)
                        _ = v56
                    default:
                        v5 = (^(fi53))
                    }
                case 2147483647:
                    _ = len(v1)
                case 127:
                    var v57 uint8 = 127
                    _ = v57
                }
                v5 = (fi53 - fi52)
                switch "\n\t" {
                case "\n\t":
                    fmt.Println(v5, v3, v33)
                case "a":
                    if ((v5 * v5) != fi53) {
                        v4 = v4
                    } else {
                        v34 -= 26
                        var v58 uint16 = v33
                        _ = v58
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    var v59 int = (fi53 - (fi52 - fi53))
                    if len(v1) > 0 {
                        v1[0] = 99
                    }
                    v4 = "a"
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    v33 += v33
                    _ = v59
                case "":
                    switch ("\n\t" + v4) {
                    case "\n\t":
                        var v60 uint8 = 0
                        _ = v60
                    case "hello":
                        v33 -= v33
                    default:
                        var v61 int8 = v34
                        _ = v61
                    }
                }
            } else {
                var v62 float64 = ((v49 - ((v49 / v49) - 3.14159)) / -2.25)
                var v63 float64 = 0.1
                fmt.Println(v62, v5)
                _ = v62
                _ = v63
            }
            _ = fi53
            _ = v55
        }
        var v64 int8 = (v34 << ((893) & 63))
        _ = fi52
        _ = v64
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v49
    _ = v5
    _ = v51
}
