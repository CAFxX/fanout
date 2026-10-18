package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 int64 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 int64; F1 uint64; F2 uint32 }
type GenS9 struct { F0 uint32; F1 uint64; F2 uint }
func genFn0() int16 {
    switch -1 {
    case 127:
        v7 := GenS1{F0: 520, F1: 0, F2: 42}
        _ = v7
    case 0:
        var v8 float32 = (((0.0 + 0.0) / 1.0) / 3.14159)
        _ = v8
    case 999:
        switch -321 {
        case 0:
            var v9 int8 = 42
            _ = v9
        case 127:
            v10 := GenS1{F0: 760, F1: 2147483647, F2: 255}
            _ = v10
        }
    }
    if (!(!(-363 != 47))) {
        var v11 int16 = 693
        if (42 <= 0) {
            var v12 int16 = 1
            var v13 uint64 = 2
            _ = v12
            _ = v13
        }
        v11 -= (^(v11))
        _ = v11
    } else {
        if true {
            for fi14 := 0; fi14 < 5; fi14++ {
                fmt.Println(fi14)
                for fi15 := 0; fi15 < 17; fi15++ {
                    var v16 int = fi14
                    v16 = 9223372036854775807
                    fmt.Println(fi14, v16)
                    v17 := []int{}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                _ = fi14
            }
            var v18 float64 = 3.14159
            if (true && (true && true)) {
                v19 := []int{-9223372036854775808, 720}
                if len(v19) > 0 {
                    v19[0] = 32768
                }
                var v20 uint64 = 42
                _ = v19
                _ = v20
            }
            _ = v18
        }
    }
    return -32768
}
func genFn1() uint16 {
    var v21 float64 = ((float64((0.1 * float32(0.0))) * ((((1.5 - 1e10) / 1.5) - 3.14159) + -2.25)) + 0.1)
    v22 := GenS2{F0: (1 < 4294967295), F1: 9223372036854775807, F2: 1}
    _ = v21
    _ = v22
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = true
    var v3 uint8 = 129
    var v4 float64 = -2.25
    v5 := []uint32{4294967295, 4294967295, 305}
    var v6 bool = false
    v23 := GenS1{F0: 4294967295, F1: 1, F2: 42}
    fmt.Println(v2)
    v6 = (18446744073709551615 < 4294967295)
    v24 := GenS2{F0: (4294967295 >= 102), F1: -9223372036854775808, F2: -9223372036854775808}
    v25 := v1["foo bar"]
    fmt.Println(v25)
    if true {
        v26 := v1[("foo bar" + ("\n\t" + ("\n\t" + (("foo bar" + "a") + "hello"))))]
        _ = v26
    } else {
        var v27 uint64 = 546
        v4 -= float64(1e10)
        v6 = v6
        _ = v27
    }
    var v28 uint16 = 2
    v29 := v1["hello"]
    v28 = v28
    v2 = v2
    switch "foo bar" {
    case "a":
        var v30 uint64 = uint64(v3)
        _ = v30
    default:
        var v31 int64 = 1
        _ = v31
    }
    if v6 {
        v6 = v2
        v28 = v28
        switch v29 {
        case -118:
            var v32 float64 = 1.5
            _ = v32
        case 42:
            for fi33 := 0; fi33 < 18; fi33++ {
                var v34 uint32 = 4294967295
                _ = fi33
                _ = v34
            }
        default:
            v25 = v29
        }
    }
    delete(v1, "\n\t")
    var v35 int = v25
    for fi36 := 0; fi36 < 2; fi36++ {
        v25 -= fi36
        if (435 <= (-(v35))) {
            if false {
                if v6 {
                    if (127 != 100) {
                        v37 := GenS1{F0: 511, F1: 42, F2: (^(v28))}
                        var v38 uint8 = v3
                        if true {
                            var v39 int16 = 1
                            _ = len(v1)
                            _ = v39
                        }
                        _ = v37
                        _ = v38
                    } else {
                        for fi40 := 0; fi40 < 9; fi40++ {
                            v35 = (v25 % (((v25 >> ((255) & 63))) | 1))
                            if true {
                                if true {
                                    var v41 int16 = genFn0()
                                    if v2 {
                                        switch "a" {
                                        case "":
                                            switch ("foo bar" + ("hello" + "\n\t")) {
                                            case "\n\t":
                                                _ = len(v1)
                                            case "":
                                                switch (-(v29)) {
                                                case 2147483647:
                                                    switch ("\n\t" + "hello") {
                                                    case "":
                                                        v42 := GenS1{F0: 18446744073709551615, F1: -774, F2: (v28 / ((genFn1()) | 1))}
                                                        _ = v42
                                                    default:
                                                        if len(v5) > 0 {
                                                            v5[0] = 970
                                                        }
                                                    }
                                                case 1:
                                                    v28 += 0
                                                case 4294967295:
                                                    switch -1 {
                                                    case -2147483648:
                                                        v43 := v1["foo bar"]
                                                        if len(v5) > 0 {
                                                            v5[0] = 0
                                                        }
                                                        v44 := GenS1{F0: 0, F1: 0, F2: v28}
                                                        v3 = (v3 / (((^(v3))) | 1))
                                                        _ = v43
                                                        _ = v44
                                                    case 772:
                                                        v35 += (v35 / ((9223372036854775807) | 1))
                                                    }
                                                }
                                            }
                                        case "foo bar":
                                            v29 = fi40
                                        case "\n\t":
                                            v4 = v4
                                        default:
                                            fmt.Println(v29, v29)
                                        }
                                        v6 = true
                                    }
                                    if false {
                                        switch "hello" {
                                        case "hello":
                                            var v45 uint16 = v28
                                            _ = v45
                                        case "foo bar":
                                            v29 += v35
                                        }
                                        fmt.Println(v4)
                                    } else {
                                        v46 := GenS3{F0: (v3 << ((4294967295) & 63)), F1: (-(v41)), F2: 42}
                                        if (42 <= 0) {
                                            var v47 uint32 = 2
                                            v48 := v1[("foo bar" + ((("\n\t" + (("a" + "foo bar") + "")) + "\n\t") + ""))]
                                            _ = v47
                                            _ = v48
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            delete(v1, ("\n\t" + "hello"))
                                            v4 = v4
                                        }
                                        _ = v46
                                    }
                                    _ = v41
                                } else {
                                    _ = len(v5)
                                    var v49 uint = 255
                                    v50 := GenS3{F0: 42, F1: 32767, F2: 253}
                                    _ = v49
                                    _ = v50
                                }
                                if false {
                                    if false {
                                        var v51 uint32 = 2
                                        _ = v51
                                    }
                                    if false {
                                        if (v2 && true) {
                                            fmt.Println(v6)
                                        }
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = 2
                                    }
                                    if (true || v6) {
                                        v52 := GenS4{F0: true, F1: 1}
                                        delete(v1, ("\n\t" + ("foo bar" + ("hello" + ""))))
                                        _ = v52
                                    }
                                    v4 = 1e10
                                    if v2 {
                                        switch "\n\t" {
                                        case "a":
                                            var v53 bool = (v2 && (v2 || v2))
                                            _ = v53
                                        case "":
                                            fmt.Println(fi36)
                                        }
                                    } else {
                                        v54 := GenS4{F0: v6, F1: 42}
                                        var v55 int64 = -944
                                        v55 = v55
                                        var v56 string = "foo bar"
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    var v57 uint8 = (v3 & v3)
                                    v58 := GenS3{F0: v3, F1: genFn0(), F2: -965}
                                    var v59 bool = v6
                                    v60 := GenS2{F0: (0 >= -2147483648), F1: 936, F2: 42}
                                    v61 := GenS3{F0: (v57 * (v57 & (^(v3)))), F1: 1000, F2: 84}
                                    v29 = (fi40 | 32768)
                                    _ = v57
                                    _ = v58
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                            } else {
                                switch (("" + "hello") + (("" + "a") + "hello")) {
                                case "\n\t":
                                    var v62 bool = (-375 == 42)
                                    _ = v62
                                case "hello":
                                    if len(v5) > 0 {
                                        v5[0] = 2
                                    }
                                case "a":
                                    var v63 uint8 = 0
                                    _ = v63
                                }
                                _ = len(v5)
                                if v2 {
                                    v3 -= v3
                                    var v64 int8 = -14
                                    var v65 int64 = 42
                                    _ = v64
                                    _ = v65
                                } else {
                                    delete(v1, "")
                                    if (9223372036854775807 <= 717) {
                                        v35 += (-(fi40))
                                    } else {
                                        if (!v6) {
                                            _ = len(v1)
                                            switch "hello" {
                                            case "a":
                                                v35 -= 832
                                            case "\n\t":
                                                v66 := v1["hello"]
                                                _ = v66
                                            case "foo bar":
                                                var v67 bool = (42 < 1)
                                                _ = v67
                                            default:
                                                var v68 uint32 = 2147483647
                                                _ = v68
                                            }
                                            var v69 uint32 = 4294967295
                                            _ = v69
                                        } else {
                                            var v70 uint = 0
                                            if (fi36 > -9223372036854775808) {
                                                switch "hello" {
                                                case "a":
                                                    var v71 uint8 = 255
                                                    _ = v71
                                                case "hello":
                                                    if (-1 > 2147483647) {
                                                        v1["\n\t"] = v25
                                                    } else {
                                                        switch (fi40 << ((v70) & 63)) {
                                                        case -2147483648:
                                                            v29 = 32767
                                                        case 9223372036854775807:
                                                            var v72 uint8 = 128
                                                            _ = v72
                                                        case 797:
                                                            v73 := GenS5{F0: (v70 + 65535), F1: 568}
                                                            _ = v73
                                                        default:
                                                            var v74 uint8 = 42
                                                            v75 := GenS6{F0: (!(!(406 < 42)))}
                                                            if ((^(v74)) == 1) {
                                                                v3 = 82
                                                            }
                                                            _ = v74
                                                            _ = v75
                                                        }
                                                        if (!((!(v2 && (true && false))) || (!false))) {
                                                            var v76 int16 = 42
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(v6)
                                                            case "foo bar":
                                                                var v77 int64 = -732
                                                                _ = v77
                                                            case "a":
                                                                v5 = append(v5, 2147483647)
                                                            }
                                                            _ = v76
                                                        } else {
                                                            var v78 float32 = 3.14159
                                                            _ = v78
                                                        }
                                                    }
                                                }
                                                v25 = 722
                                            } else {
                                                switch ("\n\t" + "") {
                                                case "hello":
                                                    v25 = (^(fi40))
                                                case "a":
                                                    if v6 {
                                                        v70 -= (^(v70))
                                                    } else {
                                                        var v79 uint32 = 65535
                                                        _ = v79
                                                    }
                                                }
                                                v80 := GenS6{F0: v6}
                                                if (!false) {
                                                    var v81 uint64 = 0
                                                    fmt.Println(fi40, v29)
                                                    var v82 int8 = -128
                                                    _ = v81
                                                    _ = v82
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = 1
                                                }
                                                v6 = v2
                                                v6 = false
                                                _ = v80
                                            }
                                            _ = v70
                                        }
                                        v6 = (967 > 18446744073709551615)
                                        v4 -= 1e10
                                    }
                                }
                                v29 = v35
                                fmt.Println(v25)
                            }
                            _ = fi40
                        }
                        v1["\n\t"] = 454
                    }
                    v25 -= 4294967295
                    if (!((1e10 + ((0.0 / -2.25) * (0.1 / float32(v4)))) != float32(float64(((float32(1e10) / 1.5) / -2.25))))) {
                        if v6 {
                            v29 -= -264
                            fmt.Println(v4, fi36, v25)
                        }
                        v28 += genFn1()
                        v83 := GenS7{F0: -21}
                        _ = v83
                    } else {
                        var v84 int32 = int32(v35)
                        if (("\n\t" + "foo bar") == ("\n\t" + "hello")) {
                            switch ("" + "foo bar") {
                            case "a":
                                var v85 int64 = 9223372036854775807
                                _ = v85
                            default:
                                for fi86 := 0; fi86 < 20; fi86++ {
                                    v35 -= 1
                                    v87 := GenS8{F0: 826, F1: 4294967295, F2: 2147483648}
                                    var v88 uint16 = v28
                                    _ = fi86
                                    _ = v87
                                    _ = v88
                                }
                            }
                            var v89 int32 = (v84 + (v84 ^ -1))
                            if (uint8(v28) < v3) {
                                for fi90 := 0; fi90 < 3; fi90++ {
                                    if (32767 <= 1) {
                                        _ = len(v5)
                                    }
                                    _ = fi90
                                }
                                fmt.Println(fi36)
                            }
                            _ = v89
                        }
                        fmt.Println(v35, v6)
                        _ = v84
                    }
                }
                fmt.Println(v35)
            } else {
                switch -3 {
                case 128:
                    v91 := GenS6{F0: ((!v2) || v2)}
                    _ = v91
                case -9223372036854775808:
                    for fi92 := 0; fi92 < 19; fi92++ {
                        var v93 float32 = 3.14159
                        _ = fi92
                        _ = v93
                    }
                case 0:
                    v94 := v1["\n\t"]
                    _ = v94
                default:
                    for fi95 := 0; fi95 < 13; fi95++ {
                        v96 := GenS4{F0: (42 != 2147483647), F1: -2147483648}
                        var v97 int = 128
                        var v98 int32 = -2147483648
                        var v99 float32 = -2.25
                        _ = fi95
                        _ = v96
                        _ = v97
                        _ = v98
                        _ = v99
                    }
                }
            }
        }
        _ = fi36
    }
    for fi100 := 0; fi100 < 16; fi100++ {
        v25 = (^(v29))
        for fi101 := 0; fi101 < 2; fi101++ {
            if v6 {
                var v102 uint32 = 170
                var v103 uint16 = v28
                var v104 uint16 = (v103 << ((42) & 63))
                _ = v102
                _ = v103
                _ = v104
            }
            v4 -= 1.5
            v105 := GenS2{F0: false, F1: (v29 / (((fi101 / ((0) | 1))) | 1)), F2: 417}
            v106 := GenS9{F0: 2147483647, F1: 18446744073709551615, F2: 896}
            _ = fi101
            _ = v105
            _ = v106
        }
        _ = fi100
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
