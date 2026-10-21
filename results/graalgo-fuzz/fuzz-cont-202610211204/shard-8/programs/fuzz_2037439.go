package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string; F1 int64; F2 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 int; F1 uint16; F2 string }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 uint32; F1 int8 }
func genFn0() int32 {
    for fi5 := 0; fi5 < 6; fi5++ {
        fmt.Println(fi5)
        var v6 int8 = 127
        v6 -= -44
        _ = fi5
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 uint64 = 2
        for fi9 := 0; fi9 < 14; fi9++ {
            if true {
                var v10 float32 = 3.14159
                _ = v10
            }
            var v11 int8 = -104
            var v12 uint8 = 42
            if ((!false) && (18446744073709551615 <= 340)) {
                if ((false && ((v11 > v11) && false)) && (381 >= 2147483647)) {
                    v8 = (v8 >> ((uint(v8)) & 63))
                    if (141 < 255) {
                        v13 := GenS1{F0: (!true), F1: -927, F2: 32767}
                        _ = v13
                    }
                    switch (fi9 << ((uint(v8)) & 63)) {
                    case -1:
                        var v14 bool = false
                        _ = v14
                    }
                    switch (("" + "hello") + "a") {
                    case "foo bar":
                        v15 := []int{}
                        _ = v15
                    case "hello":
                        v16 := GenS2{F0: v8, F1: 1}
                        _ = v16
                    case "\n\t":
                        var v17 int8 = (v11 - v11)
                        _ = v17
                    }
                }
                switch "foo bar" {
                case "\n\t":
                    var v18 bool = true
                    _ = v18
                }
            }
            switch ("hello" + (("a" + "hello") + ("a" + ("hello" + "foo bar")))) {
            case "foo bar":
                v19 := GenS1{F0: true, F1: -2147483648, F2: 32767}
                _ = v19
            case "hello":
                v20 := GenS1{F0: false, F1: 0, F2: -32768}
                _ = v20
            }
            v21 := []int{525, -897}
            var v22 int64 = -9223372036854775808
            if true {
                var v23 bool = false
                v22 = v22
                v11 = 1
                var v24 float32 = -2.25
                var v25 uint = 0
                fmt.Println(v12, v8)
                if true {
                    v12 += (v12 >> (((v25 & (v25 << (((v25 & v25)) & 63)))) & 63))
                } else {
                    var v26 int32 = -2147483648
                    if len(v21) > 0 {
                        v21[0] = 42
                    }
                    _ = v26
                }
                _ = v23
                _ = v24
                _ = v25
            }
            fmt.Println(fi9, fi7, v22)
            _ = fi9
            _ = v11
            _ = v12
            _ = v21
            _ = v22
        }
        _ = fi7
        _ = v8
    }
    var v27 uint8 = 0
    _ = v27
    return 42
}
func genFn1(p0 int, p1 int32) int {
    v28 := []int{-1, p0, (-(p0))}
    p0 = p0
    _ = p0
    _ = p1
    _ = v28
    return -500
}

func main() {
    var v1 float64 = 1.5
    v2 := []uint32{2147483648, 42, 5}
    v3 := make(map[int]bool)
    v4 := []uint32{}
    if true {
        if false {
            if false {
                v29 := GenS1{F0: true, F1: 1, F2: 1000}
                for fi30 := 0; fi30 < 14; fi30++ {
                    for fi31 := 0; fi31 < 18; fi31++ {
                        switch genFn1(genFn1((fi30 ^ genFn1(genFn1((fi30 * fi30), 438), 924)), 18), -2147483648) {
                        case 32768:
                            if (1 >= 128) {
                                _ = len(v4)
                                if true {
                                    if (-32 != -62) {
                                        var v32 int32 = -1
                                        _ = v32
                                    }
                                    var v33 int64 = -9223372036854775808
                                    _ = v33
                                } else {
                                    if (42 > 0) {
                                        v34 := GenS2{F0: 42, F1: 688}
                                        _ = v34
                                    }
                                    v1 = float64((-2.25 + 0.0))
                                    v35 := GenS3{F0: ("foo bar" + "\n\t"), F1: -266, F2: 42}
                                    _ = v35
                                }
                                v2 = append(v2, 42)
                            } else {
                                v1 += v1
                            }
                        default:
                            v1 += 3.14159
                            v1 = float64((-2.25 - (1e10 * -2.25)))
                            if ((fi30 | fi31) >= fi30) {
                                var v36 int64 = -1
                                _ = v36
                            } else {
                                switch genFn1((fi31 - genFn1((fi31 & 128), 2147483647)), 42) {
                                case 127:
                                    v1 = (float64((-2.25 + 0.0)) * 0.1)
                                case 42:
                                    v1 -= (v1 + (3.14159 - (float64((1.5 * (1e10 + (0.0 / (-2.25 / (float32(v1) - float32(v1))))))) * (0.0 + float64(0.0)))))
                                }
                            }
                            v37 := GenS2{F0: 2, F1: 2}
                            _ = v37
                        }
                        v3[(fi31 * fi30)] = (1000 > -1)
                        _ = fi31
                    }
                    v38 := GenS4{F0: 1, F1: 418}
                    _ = fi30
                    _ = v38
                }
                delete(v3, genFn1(128, -785))
                _ = v29
            } else {
                var v39 int32 = genFn0()
                var v40 int8 = 42
                v41 := GenS4{F0: -417, F1: 833}
                _ = v39
                _ = v40
                _ = v41
            }
            v42 := GenS5{F0: 524}
            _ = v42
        } else {
            for fi43 := 0; fi43 < 6; fi43++ {
                switch 0 {
                case 32768:
                    v1 += 1e10
                case 128:
                    var v44 int32 = 2147483647
                    _ = v44
                case 363:
                    delete(v3, 128)
                }
                var v45 uint8 = 2
                v46 := GenS6{F0: 128}
                _ = fi43
                _ = v45
                _ = v46
            }
            v4 = append(v4, 42)
        }
    } else {
        for fi47 := 0; fi47 < 11; fi47++ {
            switch ("a" + "") {
            case "foo bar":
                for fi48 := 0; fi48 < 15; fi48++ {
                    if false {
                        v49 := GenS7{F0: false, F1: 42}
                        v1 += ((((1e10 * 0.0) - v1) - float64(0.1)) * 0.1)
                        _ = v49
                    } else {
                        v50 := GenS2{F0: 42, F1: 2}
                        _ = v50
                    }
                    _ = fi48
                }
            case "\n\t":
                for fi51 := 0; fi51 < 5; fi51++ {
                    if (-2147483648 == -2147483648) {
                        var v52 int = (fi51 + (fi47 ^ (fi51 % (((fi47 - (fi51 % ((fi47) | 1)))) | 1))))
                        _ = v52
                    } else {
                        fmt.Println(fi47, fi47)
                        v53 := GenS8{F0: (-(fi51)), F1: 42, F2: "a"}
                        v4 = append(v4, 4294967295)
                        _ = v53
                    }
                    var v54 int32 = -1
                    _ = fi51
                    _ = v54
                }
            }
            _ = fi47
        }
        v55 := GenS9{F0: 1}
        if (v1 != 0.1) {
            var v56 int = -1
            var v57 uint16 = 2
            v57 = v57
            _ = v56
            _ = v57
        } else {
            var v58 float64 = (-2.25 - ((v1 + float64(0.1)) * 0.1))
            _ = len(v3)
            _ = v58
        }
        _ = v55
    }
    var v59 float64 = (((v1 + v1) + v1) - 1.5)
    v60 := v3[1]
    var v61 uint16 = 65535
    if (106 != 2) {
        fmt.Println(v61, v61)
    } else {
        switch "\n\t" {
        case "":
            v61 += v61
        case "hello":
            v59 = ((v59 + 3.14159) / float64(3.14159))
        case "a":
            _ = len(v2)
        default:
            fmt.Println(v60, v60, v60)
        }
        v62 := GenS10{F0: 42, F1: 127}
        v63 := GenS1{F0: v60, F1: 1, F2: -32768}
        _ = v62
        _ = v63
    }
    var v64 int = 749
    _ = len(v3)
    switch (v64 << ((18446744073709551615) & 63)) {
    case 32767:
        v1 = float64((float32((v59 * (((v59 / (1.5 - (v59 + (float64(float32(float64(-2.25))) * 1e10)))) * ((0.0 / float64(float32(((v1 * v59) - v59)))) - float64(3.14159))) + v59))) - (0.1 + 1.5)))
    }
    var v65 int16 = -32768
    for fi66 := 0; fi66 < 5; fi66++ {
        v60 = (false || false)
        v67 := GenS5{F0: 687}
        v2 = append(v2, 1)
        _ = fi66
        _ = v67
    }
    var v68 uint64 = 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
    _ = v61
    _ = v64
    _ = v65
    _ = v68
}
