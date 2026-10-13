package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint; F2 uint32 }
type GenS3 struct { F0 bool; F1 int; F2 int }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 int64; F1 bool; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 int16; F1 uint32; F2 int64 }
func genFn0(p0 int16, p1 uint) bool {
    v4 := []int{128, -2147483648}
    v5 := GenS1{F0: 42, F1: 215, F2: 87}
    var v6 int8 = 51
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return ((false && true) || true)
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int = 1
    var v3 float64 = 0.0
    var v7 uint = 35
    var v8 string = ("\n\t" + "\n\t")
    switch "hello" {
    case "a":
        v8 = ("" + v8)
    default:
        if ((!true) && false) {
            var v9 int = (-(v2))
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        v11 := []int{(fi10 | 32768), v2}
        v8 = (("foo bar" + "a") + "hello")
        var v12 int32 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    switch (^(v2)) {
    case 249:
        fmt.Println(v8)
    case -178:
        var v13 int32 = 669
        _ = v13
    }
    switch "" {
    case "hello":
        v14 := GenS2{F0: (v7 % ((v7) | 1)), F1: (v7 % ((v7) | 1)), F2: 65535}
        _ = v14
    case "\n\t":
        v15 := GenS3{F0: true, F1: v2, F2: (v2 % ((v2) | 1))}
        _ = v15
    default:
        var v16 float64 = v3
        _ = v16
    }
    if true {
        for fi17 := 0; fi17 < 9; fi17++ {
            v18 := []int{fi17, fi17, 2147483647}
            fmt.Println(v7, v2)
            var v19 int8 = 0
            _ = fi17
            _ = v18
            _ = v19
        }
        switch v8 {
        case "foo bar":
            if (1 > 100) {
                if (42 >= 42) {
                    v3 += (((1.5 / v3) + v3) / (v3 + v3))
                }
                if false {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v3 = 0.1
                        if true {
                            var v21 float64 = v3
                            var v22 float64 = v3
                            _ = v21
                            _ = v22
                        }
                        _ = fi20
                    }
                    fmt.Println(v7, v3, v3)
                    fmt.Println(v3)
                }
                v23 := GenS4{F0: 42, F1: 2147483647}
                _ = v23
            } else {
                v24 := GenS5{F0: 7, F1: ("a" != (v8 + v8)), F2: 1}
                var v25 bool = false
                _ = v24
                _ = v25
            }
        case "hello":
            fmt.Println(v2, v7)
        case "a":
            fmt.Println(v7, v2, v2)
        }
        v26 := GenS6{F0: false, F1: 1}
        _ = v26
    } else {
        fmt.Println(v8, v2)
    }
    switch (v2 * v2) {
    case 1:
        switch v8 {
        case "\n\t":
            fmt.Println(v2, v8)
        case "hello":
            v7 -= v7
        default:
            if true {
                v27 := []int{}
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v8)
                case "":
                    v8 = v8
                case "hello":
                    v7 = v7
                }
                _ = v27
            } else {
                for fi28 := 0; fi28 < 3; fi28++ {
                    for fi29 := 0; fi29 < 3; fi29++ {
                        if (!((!false) || false)) {
                            var v30 bool = false
                            var v31 int32 = -1
                            if v30 {
                                fmt.Println(v2)
                            }
                            _ = v30
                            _ = v31
                        } else {
                            v32 := GenS4{F0: 127, F1: -1}
                            v33 := []int{}
                            _ = v32
                            _ = v33
                        }
                        _ = fi29
                    }
                    _ = fi28
                }
                v3 = (((((-2.25 - v3) - v3) / ((float64(1.5) + ((v3 + float64((1.5 * 0.1))) - (v3 - float64(0.0)))) / float64(-2.25))) * v3) / v3)
                for fi34 := 0; fi34 < 6; fi34++ {
                    switch "\n\t" {
                    case "\n\t":
                        for fi35 := 0; fi35 < 4; fi35++ {
                            v3 = (v3 - float64(float32((-2.25 + 0.1))))
                            var v36 uint = v7
                            if true {
                                switch "a" {
                                case "\n\t":
                                    if (1 >= 526) {
                                        var v37 string = ("hello" + v8)
                                        _ = v37
                                    }
                                case "hello":
                                    if false {
                                        var v38 uint8 = 127
                                        _ = v38
                                    } else {
                                        fmt.Println(v3, v8, v7)
                                        fmt.Println(v7, fi35)
                                        if false {
                                            fmt.Println(fi35)
                                        }
                                    }
                                case "foo bar":
                                    v3 = v3
                                }
                                var v39 int64 = 9223372036854775807
                                v40 := v1[319]
                                v3 -= v3
                                v40 = (!v40)
                                if (!true) {
                                    if true {
                                        v41 := []int{32768, (v2 & (fi35 + v2)), (v2 % ((v2) | 1)), (fi34 << ((255) & 63))}
                                        _ = v41
                                    } else {
                                        if (176 >= 527) {
                                            var v42 uint = 1
                                            _ = v42
                                        }
                                    }
                                    v39 = (^(v39))
                                }
                                var v43 bool = false
                                v44 := GenS4{F0: 0, F1: 227}
                                v45 := GenS6{F0: false, F1: 0}
                                _ = v39
                                _ = v40
                                _ = v43
                                _ = v44
                                _ = v45
                            } else {
                                fmt.Println(fi35)
                                v7 = (v36 | 255)
                                if (-1 != fi35) {
                                    switch (v8 + v8) {
                                    case "a":
                                        v46 := []int{(v2 * -260), v2, (fi34 << ((v36) & 63))}
                                        _ = v46
                                    case "hello":
                                        v47 := GenS4{F0: 127, F1: 2147483647}
                                        _ = v47
                                    default:
                                        v48 := GenS4{F0: -56, F1: -163}
                                        _ = v48
                                    }
                                    v49 := GenS7{F0: 0, F1: 152, F2: 9223372036854775807}
                                    _ = v49
                                }
                            }
                            _ = fi35
                            _ = v36
                        }
                    case "hello":
                        var v50 uint8 = 42
                        _ = v50
                    default:
                        v7 += v7
                    }
                    for fi51 := 0; fi51 < 20; fi51++ {
                        if false {
                            v8 = v8
                            if false {
                                v8 = v8
                                v8 = ("\n\t" + "foo bar")
                            }
                        }
                        v2 -= fi34
                        if (2 == 65535) {
                            switch (-(v2)) {
                            case -9223372036854775808:
                                v7 += v7
                            case 32767:
                                fmt.Println(v2)
                            default:
                                v8 = v8
                            }
                            var v52 string = v8
                            v53 := v1[42]
                            v54 := GenS4{F0: -128, F1: -2147483648}
                            _ = v52
                            _ = v53
                            _ = v54
                        }
                        _ = fi51
                    }
                    fmt.Println(v7)
                    _ = fi34
                }
            }
        }
    case 2147483647:
        v55 := GenS5{F0: 284, F1: ((v7 > v7) || ("\n\t" >= ("hello" + "hello"))), F2: 0}
        _ = v55
    }
    var v56 uint64 = 221
    var v57 int32 = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v56
    _ = v57
    _ = v7
    _ = v8
}
