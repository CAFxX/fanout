package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int64 }
type GenS4 struct { F0 string; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int8; F1 uint; F2 int32 }
type GenS8 struct { F0 int16; F1 uint }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 uint }
type GenS11 struct { F0 bool; F1 int8; F2 int32 }
func genFn0(p0 string, p1 uint64, p2 uint8) bool {
    fmt.Println(p2, p2, p0)
    var v4 int8 = 100
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return true
}
func genFn1() int {
    var v5 float32 = (((1.5 + ((((-2.25 * 0.1) + -2.25) - (1e10 - 3.14159)) + (3.14159 * 0.0))) - 1e10) / 1e10)
    v6 := []int{42, 32768}
    if (1 != 2147483647) {
        v7 := GenS1{F0: 2147483647, F1: false}
        _ = v7
    }
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 int = 32767
    var v2 uint16 = 0
    var v3 int8 = 55
    for fi8 := 0; fi8 < 19; fi8++ {
        v2 += (^(v2))
        v9 := GenS2{F0: v3}
        if true {
            v10 := []int{(fi8 >> ((2) & 63))}
            switch "\n\t" {
            case "hello":
                var v11 uint = 991
                _ = v11
            case "a":
                var v12 float32 = 1.5
                _ = v12
            case "":
                v3 = v3
            default:
                switch "a" {
                case "":
                    v13 := GenS3{F0: "a", F1: 0}
                    _ = v13
                case "hello":
                    var v14 int32 = -1
                    _ = v14
                case "\n\t":
                    var v15 bool = false
                    _ = v15
                default:
                    var v16 uint = 1
                    _ = v16
                }
            }
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    var v17 int = v1
    var v18 uint = 330
    v2 = (^(v2))
    switch "hello" {
    case "hello":
        if (!false) {
            var v19 int8 = (v3 % ((-81) | 1))
            switch "a" {
            case "":
                var v20 uint64 = 157
                _ = v20
            }
            fmt.Println(v1, v3, v1)
            _ = v19
        }
    case "a":
        for fi21 := 0; fi21 < 1; fi21++ {
            switch (v1 - (fi21 << (((v18 + v18)) & 63))) {
            case -785:
                v2 -= 2
            }
            _ = fi21
        }
    case "foo bar":
        switch "foo bar" {
        case "a":
            v2 = 2
        default:
            v17 = (v17 << ((v18) & 63))
        }
    default:
        v22 := GenS4{F0: "\n\t", F1: 4294967295, F2: (v3 ^ v3)}
        _ = v22
    }
    var v23 float32 = 1.5
    v2 += (v2 / (((^(v2))) | 1))
    v17 -= genFn1()
    v17 = (v17 / ((v1) | 1))
    fmt.Println(v17, v23)
    var v24 int8 = (v3 - (v3 / (((v3 >> (((v18 / ((v18) | 1))) & 63))) | 1)))
    v25 := []int{(-(v17)), (-(v1))}
    v24 += 125
    switch genFn1() {
    case 4294967295:
        v23 = (3.14159 / 1.0)
    case 709:
        for fi26 := 0; fi26 < 2; fi26++ {
            v18 -= (v18 / ((v18) | 1))
            if (!false) {
                var v27 uint16 = (v2 / ((677) | 1))
                if false {
                    v27 += v2
                    for fi28 := 0; fi28 < 17; fi28++ {
                        if (1 != int64(v18)) {
                            if false {
                                if len(v25) >= 2 {
                                    v25 = v25[1:len(v25)]
                                }
                                if (630 == 0) {
                                    v29 := GenS3{F0: "\n\t", F1: -1}
                                    var v30 bool = false
                                    fmt.Println(v18, v23)
                                    v17 = fi28
                                    if ("hello" > ("\n\t" + "hello")) {
                                        if (2 < 1) {
                                            fmt.Println(v18, v18)
                                            var v31 int64 = 42
                                            v32 := GenS3{F0: "\n\t", F1: v31}
                                            _ = v31
                                            _ = v32
                                        }
                                    }
                                    _ = v29
                                    _ = v30
                                } else {
                                    if (((42 > 1) && (!(!false))) || (v18 < (v18 * 0))) {
                                        fmt.Println(fi26, v24)
                                        v18 -= (v18 ^ v18)
                                    }
                                }
                                fmt.Println(fi28)
                                switch ("\n\t" + "a") {
                                case "a":
                                    if (((v24 < (v24 % (((v24 << ((v18) & 63))) | 1))) || ((^(v18)) >= v18)) || true) {
                                        var v33 string = "a"
                                        var v34 uint16 = 2
                                        _ = v33
                                        _ = v34
                                    }
                                case "":
                                    v35 := GenS2{F0: -11}
                                    _ = v35
                                default:
                                    v18 = v18
                                }
                                v27 = v2
                            } else {
                                switch ("" + "a") {
                                case "":
                                    v18 += (^(v18))
                                case "a":
                                    if len(v25) >= 2 {
                                        v25 = v25[1:len(v25)]
                                    }
                                default:
                                    switch ("foo bar" + "a") {
                                    case "hello":
                                        switch (v17 | (fi28 + (^(fi28)))) {
                                        case 1:
                                            v23 = float32((3.14159 - 3.14159))
                                        default:
                                            _ = len(v25)
                                            v23 = float32(3.14159)
                                        }
                                        _ = len(v25)
                                        if true {
                                            if (1 >= 0) {
                                                switch 4294967295 {
                                                case 32767:
                                                    var v36 uint32 = 25
                                                    _ = v36
                                                }
                                            } else {
                                                if true {
                                                    if len(v25) > 0 {
                                                        v25[0] = -215
                                                    }
                                                    fmt.Println(v18)
                                                }
                                                fmt.Println(v27, fi28)
                                            }
                                            v25 = append(v25, (fi28 - 2147483647))
                                            if false {
                                                switch "a" {
                                                case "hello":
                                                    if true {
                                                        switch genFn1() {
                                                        case -2147483648:
                                                            v1 -= genFn1()
                                                        case 32767:
                                                            v24 -= (v3 % ((int8(fi26)) | 1))
                                                        case 42:
                                                            v23 += 1e10
                                                            v23 += ((v23 + v23) - -2.25)
                                                        }
                                                    }
                                                case "\n\t":
                                                    var v37 uint = v18
                                                    _ = v37
                                                case "a":
                                                    v2 = 2
                                                }
                                                if false {
                                                    switch ("a" + (("\n\t" + ("hello" + "hello")) + "a")) {
                                                    case "\n\t":
                                                        v25 = append(v25, genFn1())
                                                    }
                                                    var v38 uint64 = 42
                                                    var v39 string = ""
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    v40 := GenS5{F0: v27}
                                                    var v41 float32 = float32((-2.25 * 3.14159))
                                                    _ = v40
                                                    _ = v41
                                                }
                                                v24 = (v3 / (((v24 - (v24 * -41))) | 1))
                                            } else {
                                                if (255 <= 127) {
                                                    if ((v18 + (v18 >> ((v18) & 63))) != v18) {
                                                        v25 = append(v25, 127)
                                                        v2 = v2
                                                        v42 := GenS6{F0: v18, F1: v2}
                                                        _ = v42
                                                    }
                                                } else {
                                                    var v43 int64 = -9223372036854775808
                                                    v27 += v27
                                                    v18 -= (v18 | (v18 >> (((v18 << ((65535) & 63))) & 63)))
                                                    _ = v43
                                                }
                                                var v44 string = "\n\t"
                                                v45 := GenS2{F0: v3}
                                                _ = v44
                                                _ = v45
                                            }
                                        }
                                    case "foo bar":
                                        var v46 float64 = (1.5 + 1e10)
                                        _ = v46
                                    case "a":
                                        v23 += v23
                                    default:
                                        v27 -= 522
                                    }
                                }
                                var v47 uint8 = 245
                                _ = v47
                            }
                            if (9223372036854775807 < (^(v1))) {
                                v17 = genFn1()
                                if (!true) {
                                    if false {
                                        var v48 int = (fi28 | -691)
                                        if len(v25) > 0 {
                                            v25[0] = fi26
                                        }
                                        switch 42 {
                                        case 278:
                                            if (("" < "foo bar") && (!true)) {
                                                var v49 bool = false
                                                var v50 int16 = 42
                                                _ = v49
                                                _ = v50
                                            } else {
                                                var v51 int32 = -1
                                                v52 := GenS2{F0: 0}
                                                if false {
                                                    var v53 uint8 = uint8(v18)
                                                    _ = v53
                                                } else {
                                                    if true {
                                                        v23 -= (0.1 * (v23 + -2.25))
                                                        v27 = v2
                                                        v54 := GenS3{F0: ("a" + ""), F1: -1}
                                                        _ = v54
                                                    }
                                                    var v55 uint32 = 2147483648
                                                    _ = v55
                                                }
                                                _ = v51
                                                _ = v52
                                            }
                                        case 127:
                                            v56 := GenS7{F0: v24, F1: v18, F2: 2147483647}
                                            _ = v56
                                        default:
                                            v2 -= v27
                                        }
                                        _ = v48
                                    }
                                }
                                if (786 >= (v18 & (v18 << ((uint(v27)) & 63)))) {
                                    v57 := GenS8{F0: 1000, F1: v18}
                                    v27 = 151
                                    var v58 uint16 = v27
                                    var v59 uint64 = 154
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                } else {
                                    fmt.Println(v18)
                                    if true {
                                        var v60 int16 = 94
                                        fmt.Println(v3, v27, v24)
                                        if ((v18 ^ v18) != (^(v18))) {
                                            v60 -= (-(v60))
                                            fmt.Println(v1)
                                            fmt.Println(v18)
                                            v61 := GenS9{F0: v27}
                                            _ = v61
                                        }
                                        if len(v25) > 0 {
                                            v25[0] = genFn1()
                                        }
                                        _ = v60
                                    }
                                    if (247 > 2147483647) {
                                        var v62 uint16 = (v27 * 65535)
                                        if len(v25) > 0 {
                                            v25[0] = genFn1()
                                        }
                                        if (0.1 >= 1e10) {
                                            _ = len(v25)
                                            var v63 float32 = 1.5
                                            _ = v63
                                        }
                                        _ = v62
                                    } else {
                                        v64 := GenS10{F0: -1, F1: (v18 / ((v18) | 1))}
                                        if true {
                                            var v65 int16 = 1000
                                            v65 -= (v65 - v65)
                                            _ = v65
                                        }
                                        var v66 int8 = v24
                                        _ = v64
                                        _ = v66
                                    }
                                    var v67 uint32 = 795
                                    v68 := GenS11{F0: (((float64(v23) / 1.0) + 0.0) <= 0.1), F1: v3, F2: 981}
                                    _ = v67
                                    _ = v68
                                }
                            } else {
                                var v69 int32 = int32(v24)
                                v17 = v1
                                _ = v69
                            }
                            var v70 bool = false
                            fmt.Println(v3, v1, v1)
                            _ = v70
                        } else {
                            var v71 int8 = v3
                            v2 -= (^(v27))
                            var v72 int = fi26
                            v27 = 255
                            _ = v71
                            _ = v72
                        }
                        v17 -= 42
                        fmt.Println(v1)
                        var v73 int16 = int16(v1)
                        _ = fi28
                        _ = v73
                    }
                    var v74 float32 = v23
                    _ = v74
                }
                var v75 uint32 = 799
                _ = v27
                _ = v75
            } else {
                fmt.Println(v18, v18, v24)
                fmt.Println(v18)
                _ = len(v25)
            }
            _ = fi26
        }
    case -183:
        if len(v25) > 0 {
            v25[0] = v1
        }
    }
    var v76 bool = (1000 <= 42)
    for fi77 := 0; fi77 < 6; fi77++ {
        v78 := GenS5{F0: 2}
        _ = fi77
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v76:", v76)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v76
}
