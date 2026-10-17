package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 string; F2 int16 }
type GenS3 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS4 struct { F0 int8; F1 int64; F2 uint64 }
type GenS5 struct { F0 int; F1 int16; F2 bool }
type GenS6 struct { F0 uint8; F1 int16 }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint64 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 uint8 }
type GenS13 struct { F0 uint64; F1 string }
func genFn0(p0 bool, p1 string, p2 string) bool {
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            p2 = p1
            var v7 int = (fi6 % (((fi6 >> ((1) & 63))) | 1))
            _ = fi6
            _ = v7
        }
        var v8 uint32 = 859
        _ = fi5
        _ = v8
    }
    p2 = p2
    _ = p0
    _ = p1
    _ = p2
    return (!(!(((0.1 / ((float64(0.1) * 3.14159) * 1e10)) == (3.14159 + 3.14159)) && p0)))
}
func genFn1(p0 uint16) uint16 {
    if (947 >= 4294967295) {
        v9 := []int{0, 127, 46, -359}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 4; fi10++ {
            v11 := []int{(fi10 * fi10), 485}
            switch "a" {
            case "foo bar":
                p0 -= 65535
            case "a":
                switch "a" {
                case "foo bar":
                    var v12 int8 = 127
                    _ = v12
                case "a":
                    var v13 uint32 = 1
                    _ = v13
                default:
                    for fi14 := 0; fi14 < 6; fi14++ {
                        switch fi14 {
                        case 1:
                            var v15 int = fi10
                            _ = v15
                        case 184:
                            fmt.Println(p0)
                        case -187:
                            var v16 uint = 716
                            _ = v16
                        }
                        var v17 bool = true
                        _ = fi14
                        _ = v17
                    }
                }
            case "\n\t":
                for fi18 := 0; fi18 < 8; fi18++ {
                    switch ("" + "") {
                    case "":
                        p0 -= 1
                        v19 := GenS1{F0: 42}
                        v11 = append(v11, (fi18 - 32767))
                        _ = v19
                    case "hello":
                        v20 := GenS2{F0: 180, F1: "foo bar", F2: -1}
                        _ = v20
                    case "\n\t":
                        fmt.Println(fi18, fi10)
                        v21 := GenS1{F0: -2147483648}
                        if len(v11) > 0 {
                            v11[0] = (fi18 % ((4294967295) | 1))
                        }
                        _ = v21
                    }
                    var v22 uint64 = uint64(p0)
                    _ = fi18
                    _ = v22
                }
            }
            _ = fi10
            _ = v11
        }
        var v23 bool = (-537 <= -9223372036854775808)
        fmt.Println(p0, v23)
        _ = v23
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        var v25 uint = 42
        v25 -= (v25 ^ 2)
        v25 = (v25 & 4294967295)
        _ = fi24
        _ = v25
    }
    _ = p0
    return 2
}

func main() {
    var v1 string = ((("foo bar" + "\n\t") + "a") + "hello")
    var v2 float32 = 0.1
    v3 := []uint32{2147483648, 2147483647}
    v4 := make(map[int]uint32)
    v2 = (v2 * ((v2 / (float32(float64(v2)) / float32(-2.25))) * 3.14159))
    if (true || false) {
        v2 = v2
    } else {
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 18; fi27++ {
                if (1 <= 0) {
                    if true {
                        v1 = v1
                        fmt.Println(fi26, fi26)
                        if true {
                            switch fi27 {
                            case -623:
                                var v28 bool = (float64(((float32(3.14159) * (v2 * v2)) / v2)) == ((float64(v2) + 1.5) / 0.1))
                                _ = v28
                            case 127:
                                switch "hello" {
                                case "foo bar":
                                    switch 42 {
                                    case 128:
                                        _ = len(v3)
                                    default:
                                        _ = len(v4)
                                    }
                                case "a":
                                    v1 = v1
                                    v2 += (v2 - float32(((1.5 / 1.5) / (0.0 * (3.14159 / float64(v2))))))
                                default:
                                    if true {
                                        v4[fi26] = 0
                                        v2 -= (v2 + ((v2 * 1e10) + (v2 + v2)))
                                    } else {
                                        if (!(!false)) {
                                            v4[2147483647] = 42
                                            switch v1 {
                                            case "\n\t":
                                                v1 = v1
                                            case "a":
                                                fmt.Println(v1, v2)
                                            default:
                                                var v29 string = (v1 + (v1 + v1))
                                                _ = v29
                                            }
                                        }
                                        if (false && (127 == (fi26 ^ (fi27 - (fi27 << ((132) & 63)))))) {
                                            var v30 uint64 = 42
                                            _ = len(v3)
                                            _ = v30
                                        } else {
                                            var v31 uint16 = 251
                                            if (-714 < 1) {
                                                var v32 bool = false
                                                _ = v32
                                            } else {
                                                v33 := GenS2{F0: 65535, F1: v1, F2: 736}
                                                fmt.Println(v2, v1, v31)
                                                fmt.Println(fi27, v2, fi26)
                                                _ = v33
                                            }
                                            var v34 int8 = 127
                                            _ = v31
                                            _ = v34
                                        }
                                    }
                                    v35 := v4[fi26]
                                    _ = v35
                                }
                            case 0:
                                if true {
                                    var v36 int32 = -1
                                    if len(v3) > 0 {
                                        v3[0] = 2
                                    }
                                    if (1 < 0) {
                                        v2 = 3.14159
                                    } else {
                                        switch 1 {
                                        case 182:
                                            v1 = (v1 + "\n\t")
                                        case 4294967295:
                                            switch (("a" + "a") + v1) {
                                            case "hello":
                                                if (2 != 255) {
                                                    _ = len(v4)
                                                    var v37 uint = 4294967295
                                                    _ = v37
                                                }
                                                v38 := GenS3{F0: 2, F1: (42 < 1), F2: 103}
                                                _ = v38
                                            case "a":
                                                var v39 float32 = ((((float32((1.5 / 3.14159)) - v2) / v2) - v2) + 1e10)
                                                _ = v39
                                            }
                                        case 32768:
                                            v40 := GenS4{F0: -68, F1: 0, F2: 104}
                                            _ = v40
                                        }
                                        if (326 != 65535) {
                                            v41 := GenS5{F0: 0, F1: 914, F2: true}
                                            var v42 bool = (-456 <= 42)
                                            _ = v41
                                            _ = v42
                                        }
                                    }
                                    v43 := GenS5{F0: (fi27 >> ((2) & 63)), F1: 42, F2: ((!(v2 >= (0.0 * 1.5))) && (false || (!true)))}
                                    _ = v36
                                    _ = v43
                                } else {
                                    switch "foo bar" {
                                    case "hello":
                                        v44 := GenS6{F0: 0, F1: 39}
                                        _ = v44
                                    case "foo bar":
                                        if ((("hello" + v1) < v1) && true) {
                                            v2 -= (-2.25 / -2.25)
                                            if true {
                                                fmt.Println(fi27, v2)
                                            }
                                            v1 = v1
                                        }
                                    }
                                }
                            default:
                                v45 := GenS7{F0: 2}
                                _ = v45
                            }
                            if (!(false && false)) {
                                v46 := GenS7{F0: 2147483648}
                                _ = v46
                            } else {
                                switch "a" {
                                case "\n\t":
                                    if (2147483647 < 0) {
                                        v2 = (float32(1e10) + v2)
                                        v4[(fi27 ^ -476)] = 1
                                    }
                                case "hello":
                                    if true {
                                        if false {
                                            v47 := GenS8{F0: -1, F1: 700, F2: 2147483648}
                                            if (!false) {
                                                var v48 float32 = 0.0
                                                v48 -= v48
                                                _ = v48
                                            } else {
                                                switch (^(fi26)) {
                                                case -509:
                                                    v1 = (("a" + "a") + v1)
                                                case 32768:
                                                    fmt.Println(v1, fi27, fi26)
                                                case 4294967295:
                                                    var v49 int64 = -9223372036854775808
                                                    _ = v49
                                                default:
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        var v50 uint16 = 255
                                                        _ = v50
                                                    case "":
                                                        v1 = "a"
                                                    }
                                                }
                                            }
                                            v1 = v1
                                            _ = v47
                                        }
                                        fmt.Println(fi27, fi27, v2)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                case "foo bar":
                                    v51 := GenS9{F0: 42}
                                    _ = v51
                                default:
                                    if (!false) {
                                        v2 = float32(1.5)
                                        var v52 bool = false
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        v53 := GenS6{F0: 189, F1: -32768}
                                        v1 = "hello"
                                        v4[215] = 65535
                                        _ = v52
                                        _ = v53
                                    } else {
                                        v2 = v2
                                        var v54 int32 = 0
                                        _ = v54
                                    }
                                }
                            }
                            var v55 bool = false
                            _ = v55
                        } else {
                            v56 := GenS10{F0: 1}
                            v57 := GenS4{F0: 0, F1: -1, F2: 42}
                            _ = v56
                            _ = v57
                        }
                    }
                    var v58 uint8 = 2
                    v59 := v4[9223372036854775807]
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                    _ = v58
                    _ = v59
                }
                v2 -= v2
                v2 = 0.1
                switch 2147483647 {
                case 334:
                    var v60 string = ("foo bar" + "\n\t")
                    _ = v60
                }
                _ = fi27
            }
            _ = fi26
        }
        for fi61 := 0; fi61 < 10; fi61++ {
            v2 += (v2 - 1e10)
            _ = fi61
        }
        switch 32768 {
        case 42:
            v62 := GenS11{F0: 481}
            _ = v62
        case -2147483648:
            if true {
                v63 := GenS6{F0: 36, F1: 0}
                _ = v63
            }
        case -1:
            v1 = ""
        }
    }
    var v64 uint64 = 90
    if false {
        switch "a" {
        case "foo bar":
            fmt.Println(v64, v64)
        }
    }
    fmt.Println(v64, v2)
    var v65 bool = false
    v66 := GenS12{F0: 208}
    if true {
        if ((v65 || true) || (3.14159 == ((1e10 / 1.0) / (0.0 + (-2.25 + (-2.25 - 0.1)))))) {
            v65 = (v65 || (v2 < ((v2 + 1e10) * (v2 / v2))))
        }
    } else {
        v67 := GenS10{F0: (v64 << ((1) & 63))}
        _ = v67
    }
    var v68 uint64 = v64
    v69 := GenS3{F0: (v64 << ((211) & 63)), F1: v65, F2: 1}
    v1 = "\n\t"
    v70 := GenS13{F0: 0, F1: ("hello" + v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v64
    _ = v65
    _ = v66
    _ = v68
    _ = v69
    _ = v70
}
