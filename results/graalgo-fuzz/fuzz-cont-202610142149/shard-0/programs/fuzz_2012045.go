package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint8; F2 string }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 bool }
type GenS5 struct { F0 int8; F1 int }
type GenS6 struct { F0 int; F1 int64; F2 int64 }
type GenS7 struct { F0 uint; F1 uint32 }
type GenS8 struct { F0 int64; F1 uint32 }
func genFn0(p0 uint64) int {
    p0 = (p0 % ((42) | 1))
    v4 := GenS1{F0: true}
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v4
    return 128
}

func main() {
    v1 := make(map[int]bool)
    var v2 int = -9223372036854775808
    var v3 uint32 = 2147483648
    if (!false) {
        var v5 int16 = -1
        _ = v5
    } else {
        v6 := GenS1{F0: true}
        _ = v6
    }
    for fi7 := 0; fi7 < 7; fi7++ {
        switch ("a" + ("hello" + ("a" + "a"))) {
        case "\n\t":
            var v8 bool = (!false)
            _ = v8
        case "hello":
            v3 = v3
        case "foo bar":
            fmt.Println(fi7, v3)
        }
        var v9 int8 = -128
        var v10 uint16 = 32767
        _ = fi7
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 127, F1: 128, F2: ("a" + "a")}
    var v12 int64 = 0
    switch "" {
    case "hello":
        var v13 bool = false
        _ = v13
    case "":
        v3 = v3
    case "\n\t":
        var v14 bool = false
        _ = v14
    }
    var v15 uint = 255
    if false {
        if true {
            v16 := GenS1{F0: (10 >= v15)}
            if true {
                if true {
                    v15 = v15
                    var v17 float32 = 3.14159
                    _ = v17
                }
                v18 := GenS1{F0: (true || true)}
                _ = v18
            }
            v19 := v1[v2]
            _ = v16
            _ = v19
        } else {
            if (42 >= 42) {
                switch "foo bar" {
                case "\n\t":
                    switch "\n\t" {
                    case "\n\t":
                        v20 := GenS3{F0: 1}
                        _ = v20
                    default:
                        switch 1 {
                        case 128:
                            if true {
                                v21 := GenS1{F0: ((v2 | int(v3)) > v2)}
                                v3 += 2147483648
                                _ = v21
                            }
                        case 4294967295:
                            for fi22 := 0; fi22 < 19; fi22++ {
                                v23 := GenS1{F0: (!false)}
                                v15 = v15
                                var v24 bool = ((v3 / (((^(v3))) | 1)) == (v3 + (^(v3))))
                                _ = fi22
                                _ = v23
                                _ = v24
                            }
                        case 265:
                            v2 -= -908
                        default:
                            switch "" {
                            case "foo bar":
                                v12 -= v12
                            case "\n\t":
                                for fi25 := 0; fi25 < 14; fi25++ {
                                    if (!true) {
                                        v26 := GenS4{F0: false, F1: ((("hello" + ("" + "\n\t")) + "") <= "foo bar")}
                                        _ = v26
                                    } else {
                                        var v27 uint16 = 0
                                        var v28 int64 = v12
                                        _ = v27
                                        _ = v28
                                    }
                                    _ = fi25
                                }
                            default:
                                for fi29 := 0; fi29 < 16; fi29++ {
                                    v2 = (fi29 + (fi29 / (((fi29 * v2)) | 1)))
                                    var v30 int = fi29
                                    _ = fi29
                                    _ = v30
                                }
                            }
                        }
                    }
                case "":
                    v31 := GenS4{F0: false, F1: false}
                    _ = v31
                case "foo bar":
                    fmt.Println(v3)
                }
                for fi32 := 0; fi32 < 1; fi32++ {
                    if false {
                        fmt.Println(fi32)
                        var v33 int8 = 1
                        _ = v33
                    }
                    var v34 int64 = (v12 & (^(v12)))
                    v34 = 1
                    _ = fi32
                    _ = v34
                }
            }
            if (242 <= 255) {
                fmt.Println(v2)
                var v35 float32 = 1.5
                _ = v35
            } else {
                for fi36 := 0; fi36 < 18; fi36++ {
                    v2 = 32767
                    _ = fi36
                }
                for fi37 := 0; fi37 < 13; fi37++ {
                    var v38 int16 = 0
                    v39 := GenS2{F0: 210, F1: 199, F2: "a"}
                    switch genFn0(1) {
                    case 128:
                        if true {
                            if false {
                                switch "" {
                                case "hello":
                                    for fi40 := 0; fi40 < 8; fi40++ {
                                        if false {
                                            v41 := []int{(v2 * (fi40 | (fi40 / ((v2) | 1))))}
                                            _ = v41
                                        } else {
                                            v38 = v38
                                        }
                                        _ = fi40
                                    }
                                case "\n\t":
                                    fmt.Println(v12, v15)
                                case "foo bar":
                                    var v42 uint16 = 0
                                    _ = v42
                                }
                                delete(v1, 9223372036854775807)
                                for fi43 := 0; fi43 < 6; fi43++ {
                                    v15 -= (v15 - v15)
                                    var v44 uint8 = 46
                                    fmt.Println(fi37, v3)
                                    _ = fi43
                                    _ = v44
                                }
                            }
                        } else {
                            fmt.Println(v3, v15, v2)
                            v12 += (v12 >> ((v15) & 63))
                        }
                    default:
                        var v45 uint16 = 201
                        _ = v45
                    }
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
        }
        v2 -= (^(v2))
    } else {
        v46 := GenS5{F0: -128, F1: 32768}
        switch genFn0(839) {
        case 446:
            if (-9223372036854775808 == int64(v2)) {
                v47 := []int{-773, (v2 & int(v15)), (v2 << (((v15 << (((v15 + 0)) & 63))) & 63))}
                switch ("foo bar" + "") {
                case "foo bar":
                    var v48 uint8 = 127
                    _ = v48
                case "\n\t":
                    if ((273 > 1000) || (40 > 1)) {
                        v3 -= 336
                    } else {
                        fmt.Println(v3)
                        if false {
                            v2 = -596
                        }
                        v49 := GenS3{F0: 2147483647}
                        _ = v49
                    }
                case "":
                    var v50 float32 = 0.1
                    _ = v50
                default:
                    delete(v1, int(v3))
                }
                _ = v47
            }
        default:
            v1[(v2 / ((9223372036854775807) | 1))] = (-128 < 127)
        }
        _ = v46
    }
    var v51 int = v2
    v12 = v12
    for fi52 := 0; fi52 < 2; fi52++ {
        v53 := GenS6{F0: 586, F1: -1, F2: (v12 | (v12 | (v12 * (-(v12)))))}
        var v54 float64 = ((1e10 * 1.5) / 3.14159)
        for fi55 := 0; fi55 < 14; fi55++ {
            if true {
                var v56 uint32 = (v3 + v3)
                switch fi55 {
                case -9223372036854775808:
                    v57 := GenS5{F0: int8(v2), F1: genFn0(18446744073709551615)}
                    _ = v57
                case 1:
                    fmt.Println(v2, v15, fi55)
                case -164:
                    delete(v1, -2147483648)
                }
                _ = v56
            } else {
                fmt.Println(v15, v15, fi52)
            }
            v58 := []int{-483, 128}
            switch "a" {
            case "hello":
                var v59 float64 = (v54 + v54)
                _ = v59
            default:
                v54 += ((float64(1e10) / v54) / v54)
            }
            _ = fi55
            _ = v58
        }
        _ = fi52
        _ = v53
        _ = v54
    }
    fmt.Println(v51)
    v1[(v2 >> (((^(v15))) & 63))] = (false && false)
    var v60 uint = (v15 << ((uint(v3)) & 63))
    v3 = (v3 * v3)
    if false {
        if (42 != 36) {
            fmt.Println(v15, v60)
        }
        v60 = (^(v60))
    } else {
        for fi61 := 0; fi61 < 12; fi61++ {
            if (1.5 >= 0.1) {
                v12 += (v12 + v12)
            } else {
                fmt.Println(v2, v15)
                switch fi61 {
                case 32767:
                    switch (fi61 % ((genFn0(0)) | 1)) {
                    case 128:
                        v62 := GenS2{F0: 128, F1: 1, F2: (("" + "hello") + "\n\t")}
                        _ = v62
                    case -761:
                        v63 := GenS6{F0: v2, F1: -752, F2: v12}
                        _ = v63
                    case -2147483648:
                        delete(v1, (fi61 ^ -9223372036854775808))
                    default:
                        v60 += (v60 ^ 255)
                    }
                case -963:
                    if (uint16(v3) > 189) {
                        v64 := GenS5{F0: 100, F1: (fi61 & v2)}
                        _ = v64
                    } else {
                        if (v12 > (^(v12))) {
                            v12 = v12
                        } else {
                            switch genFn0(0) {
                            case 9223372036854775807:
                                var v65 float32 = 0.1
                                _ = v65
                            }
                        }
                        fmt.Println(v2)
                        v51 += int(v3)
                    }
                }
            }
            if false {
                v66 := []int{}
                v12 = v12
                for fi67 := 0; fi67 < 11; fi67++ {
                    if true {
                        if true {
                            var v68 float64 = ((1e10 - 1e10) / 3.14159)
                            _ = v68
                        }
                        _ = len(v66)
                    } else {
                        var v69 int64 = v12
                        _ = v69
                    }
                    v70 := v1[genFn0(1)]
                    if (v70 || true) {
                        v51 = (v51 - (v51 % (((-(v51))) | 1)))
                        v71 := GenS4{F0: ((308 >= 0) || true), F1: true}
                        _ = v71
                    } else {
                        switch (fi67 + 1) {
                        case 9223372036854775807:
                            var v72 int8 = 100
                            _ = v72
                        default:
                            v73 := GenS7{F0: 18446744073709551615, F1: v3}
                            _ = v73
                        }
                    }
                    v74 := GenS7{F0: v15, F1: v3}
                    _ = fi67
                    _ = v70
                    _ = v74
                }
                _ = v66
            }
            _ = fi61
        }
    }
    if (-1 >= genFn0(2)) {
        v75 := GenS2{F0: 2, F1: 0, F2: ""}
        v76 := GenS8{F0: 42, F1: v3}
        _ = v75
        _ = v76
    } else {
        v51 = genFn0(2)
        v15 = v60
        v77 := GenS6{F0: (v51 ^ (v51 << ((1) & 63))), F1: -1, F2: v12}
        _ = v77
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v51:", v51)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v51
    _ = v60
}
