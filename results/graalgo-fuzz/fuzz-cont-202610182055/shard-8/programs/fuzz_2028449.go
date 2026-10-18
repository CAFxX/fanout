package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 string }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 bool; F2 int32 }
type GenS6 struct { F0 uint; F1 int32 }
func genFn0(p0 int8, p1 uint8) uint8 {
    for fi4 := 0; fi4 < 1; fi4++ {
        switch (fi4 >> ((339) & 63)) {
        case -779:
            switch "" {
            case "\n\t":
                v5 := []int{}
                _ = v5
            default:
                switch (fi4 ^ (fi4 << ((0) & 63))) {
                case -409:
                    var v6 int8 = (-(p0))
                    _ = v6
                case -196:
                    switch fi4 {
                    case 955:
                        if (-1 > -1) {
                            v7 := []int{32767}
                            _ = v7
                        }
                    case 4294967295:
                        v8 := []int{(fi4 >> ((18446744073709551615) & 63))}
                        _ = v8
                    default:
                        var v9 int16 = -32768
                        _ = v9
                    }
                default:
                    for fi10 := 0; fi10 < 13; fi10++ {
                        switch 128 {
                        case -9223372036854775808:
                            v11 := []int{(fi10 << ((748) & 63))}
                            _ = v11
                        default:
                            v12 := GenS1{F0: 838}
                            var v13 int32 = -1
                            _ = v12
                            _ = v13
                        }
                        v14 := GenS1{F0: 0}
                        var v15 float32 = (((float32(-2.25) / -2.25) / ((0.0 + (0.1 / -2.25)) - float32(1e10))) - 3.14159)
                        _ = fi10
                        _ = v14
                        _ = v15
                    }
                }
            }
        case -2147483648:
            for fi16 := 0; fi16 < 9; fi16++ {
                if (65535 >= 255) {
                    switch (("foo bar" + ("a" + "a")) + "\n\t") {
                    case "\n\t":
                        var v17 uint32 = 2147483648
                        _ = v17
                    case "a":
                        var v18 uint = 255
                        _ = v18
                    }
                } else {
                    var v19 int16 = -277
                    var v20 int = fi4
                    _ = v19
                    _ = v20
                }
                if ((p1 & p1) < p1) {
                    v21 := []int{int(p0), 9223372036854775807, 829, fi16}
                    fmt.Println(fi16, fi16)
                    p1 += 1
                    _ = v21
                } else {
                    var v22 int8 = (p0 ^ (p0 | p0))
                    _ = v22
                }
                _ = fi16
            }
        default:
            p0 = p0
        }
        var v23 uint16 = 0
        var v24 uint = 1
        _ = fi4
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: 32767}
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := []int{42}
        _ = fi26
        _ = v27
    }
    _ = p0
    _ = p1
    _ = v25
    return p1
}

func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]uint32)
    for fi28 := 0; fi28 < 13; fi28++ {
        fmt.Println(fi28)
        if true {
            for fi29 := 0; fi29 < 18; fi29++ {
                var v30 uint32 = 2
                switch (-(fi29)) {
                case -559:
                    v30 -= (v30 + (v30 % ((135) | 1)))
                case -348:
                    fmt.Println(fi29)
                }
                v1 = (v1 * v1)
                _ = fi29
                _ = v30
            }
            switch (fi28 ^ (fi28 * fi28)) {
            case -609:
                var v31 int32 = -1
                _ = v31
            default:
                v1 = (v1 >> ((433) & 63))
            }
            if true {
                if ((1 >= -44) && true) {
                    v1 += v1
                    switch (fi28 - fi28) {
                    case 127:
                        var v32 string = ("a" + "hello")
                        _ = v32
                    case 4294967295:
                        for fi33 := 0; fi33 < 9; fi33++ {
                            v34 := []int{9223372036854775807}
                            if false {
                                v35 := GenS2{F0: "a", F1: ("a" + ("foo bar" + "hello"))}
                                v36 := GenS1{F0: 65535}
                                _ = len(v3)
                                _ = v35
                                _ = v36
                            }
                            _ = fi33
                            _ = v34
                        }
                    case -868:
                        v1 = (v1 + (v1 - v1))
                    default:
                        var v37 float64 = (float64((3.14159 / (0.1 / 1.5))) / (1e10 / (-2.25 / float64(-2.25))))
                        _ = v37
                    }
                } else {
                    v1 -= (v1 / (((v1 / ((v1) | 1))) | 1))
                }
                v38 := v2[2147483648]
                _ = v38
            } else {
                v1 += (^(v1))
                v39 := GenS3{F0: -2147483648, F1: 0}
                switch "foo bar" {
                case "hello":
                    var v40 int8 = 127
                    _ = v40
                case "":
                    v3["\n\t"] = 4294967295
                case "foo bar":
                    for fi41 := 0; fi41 < 6; fi41++ {
                        switch ("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + (("" + "hello") + "\n\t")))) {
                        case "foo bar":
                            v42 := GenS4{F0: false}
                            _ = v42
                        case "hello":
                            v1 += v1
                        default:
                            v1 -= v1
                        }
                        if true {
                            if false {
                                fmt.Println(fi28, v1, fi28)
                            } else {
                                fmt.Println(v1)
                            }
                        }
                        _ = fi41
                    }
                default:
                    v3[(((("\n\t" + (("hello" + "") + ("a" + (("" + "\n\t") + "")))) + "foo bar") + "hello") + "hello")] = 65535
                }
                _ = v39
            }
        } else {
            v1 = 42
            v1 += v1
        }
        if false {
            v43 := []int{(fi28 ^ (fi28 % ((fi28) | 1)))}
            _ = v43
        }
        _ = fi28
    }
    if (100 == 42) {
        var v44 bool = (556 > 18446744073709551615)
        for fi45 := 0; fi45 < 17; fi45++ {
            var v46 uint32 = 2147483648
            v47 := []int{}
            if v44 {
                v48 := GenS2{F0: (((("" + "a") + (("a" + "foo bar") + "\n\t")) + "foo bar") + ("foo bar" + "hello")), F1: "a"}
                var v49 string = "\n\t"
                _ = v48
                _ = v49
            }
            _ = fi45
            _ = v46
            _ = v47
        }
        _ = v44
    }
    if (!(((("a" >= "a") && true) || false) && (!true))) {
        v50 := v2[2]
        for fi51 := 0; fi51 < 6; fi51++ {
            for fi52 := 0; fi52 < 12; fi52++ {
                v53 := v2[(v50 - v50)]
                _ = fi52
                _ = v53
            }
            v54 := GenS1{F0: 65535}
            var v55 float64 = -2.25
            _ = fi51
            _ = v54
            _ = v55
        }
        _ = v50
    } else {
        for fi56 := 0; fi56 < 16; fi56++ {
            var v57 int16 = -593
            if true {
                v58 := GenS4{F0: false}
                v59 := []int{}
                if (-32768 > v57) {
                    v1 = (v57 + v1)
                    var v60 int16 = v57
                    _ = v60
                } else {
                    v61 := GenS5{F0: 2, F1: false, F2: 2147483647}
                    for fi62 := 0; fi62 < 16; fi62++ {
                        var v63 uint8 = 255
                        _ = fi62
                        _ = v63
                    }
                    v64 := GenS1{F0: 0}
                    _ = v61
                    _ = v64
                }
                _ = v58
                _ = v59
            } else {
                if true {
                    switch (^(fi56)) {
                    case 2147483647:
                        if false {
                            fmt.Println(v57, v57, v57)
                            v57 -= v57
                            var v65 uint64 = 0
                            _ = v65
                        } else {
                            if false {
                                for fi66 := 0; fi66 < 20; fi66++ {
                                    var v67 int8 = -93
                                    _ = fi66
                                    _ = v67
                                }
                                v68 := GenS5{F0: 65535, F1: false, F2: -1}
                                v69 := []int{(fi56 / ((fi56) | 1))}
                                _ = v68
                                _ = v69
                            }
                            fmt.Println(fi56, v1)
                            var v70 bool = (255 >= 255)
                            _ = v70
                        }
                    case -2147483648:
                        var v71 int8 = -128
                        _ = v71
                    case -903:
                        for fi72 := 0; fi72 < 11; fi72++ {
                            var v73 float32 = 1.5
                            delete(v2, 2147483647)
                            _ = fi72
                            _ = v73
                        }
                    }
                }
                v74 := []int{fi56, (-(fi56)), -9223372036854775808}
                _ = v74
            }
            _ = fi56
            _ = v57
        }
        switch "hello" {
        case "foo bar":
            v1 = 776
        }
        var v75 int32 = 2147483647
        _ = v75
    }
    var v76 bool = true
    fmt.Println(v76, v76)
    v76 = (v76 || (v76 || v76))
    var v77 uint = 42
    _ = len(v2)
    v77 += v77
    switch ((("foo bar" + ((("" + ("hello" + "")) + "foo bar") + ("\n\t" + ((("a" + "hello") + ("foo bar" + ((((((("foo bar" + (("a" + "foo bar") + (("hello" + (((("foo bar" + "foo bar") + "foo bar") + "foo bar") + ("a" + "\n\t"))) + "\n\t"))) + "\n\t") + (("a" + "a") + ("foo bar" + "\n\t"))) + "hello") + "\n\t") + "foo bar") + "a"))) + "a")))) + "\n\t") + "a") {
    case "":
        if (42 == 260) {
            v76 = (float32(1.5) != 1e10)
        }
    case "hello":
        switch "\n\t" {
        case "foo bar":
            var v78 float32 = 1.5
            _ = v78
        default:
            var v79 uint32 = 4294967295
            _ = v79
        }
    case "a":
        if v76 {
            var v80 int64 = -537
            v76 = v76
            if v76 {
                v1 += (v1 * -447)
            }
            _ = v80
        }
    default:
        for fi81 := 0; fi81 < 9; fi81++ {
            switch "" {
            case "a":
                if false {
                    if ("a" < ("" + "\n\t")) {
                        fmt.Println(fi81)
                        var v82 bool = v76
                        _ = v82
                    } else {
                        switch (("foo bar" + "") + "foo bar") {
                        case "hello":
                            var v83 int8 = 42
                            _ = v83
                        }
                        if (!((1e10 - 0.1) > -2.25)) {
                            if (true && false) {
                                for fi84 := 0; fi84 < 6; fi84++ {
                                    v85 := []int{fi84, fi81}
                                    _ = fi84
                                    _ = v85
                                }
                            }
                            for fi86 := 0; fi86 < 19; fi86++ {
                                var v87 int = -9223372036854775808
                                if false {
                                    switch v87 {
                                    case -762:
                                        var v88 int32 = 0
                                        _ = v88
                                    }
                                    fmt.Println(v1, v76)
                                    var v89 int32 = 2147483647
                                    _ = v89
                                }
                                _ = fi86
                                _ = v87
                            }
                        } else {
                            fmt.Println(v1)
                        }
                        for fi90 := 0; fi90 < 9; fi90++ {
                            var v91 bool = (4294967295 > 0)
                            var v92 int16 = 0
                            _ = fi90
                            _ = v91
                            _ = v92
                        }
                    }
                }
            case "foo bar":
                fmt.Println(v1, v77)
            default:
                v93 := GenS6{F0: (v77 - 42), F1: 2147483647}
                _ = v93
            }
            v94 := GenS5{F0: 645, F1: v76, F2: 42}
            var v95 string = "\n\t"
            _ = fi81
            _ = v94
            _ = v95
        }
    }
    if (211 != 2147483648) {
        v96 := GenS2{F0: "foo bar", F1: ("hello" + "")}
        var v97 float32 = 0.1
        switch 128 {
        case 4294967295:
            v3["a"] = 0
        case 32767:
            if ("hello" < ("hello" + ("foo bar" + ("hello" + "\n\t")))) {
                if (!(835 <= -378)) {
                    var v98 float64 = (0.0 + (float64(v97) * 1.5))
                    _ = v98
                }
                v99 := []int{128, -186, -332, 2147483647}
                _ = v99
            }
        case 32768:
            if (-395 <= (v1 ^ (v1 << (((^(v77))) & 63)))) {
                v97 = (v97 + (1.5 * v97))
            }
        default:
            if (42 != -1) {
                if (true && v76) {
                    for fi100 := 0; fi100 < 20; fi100++ {
                        if true {
                            var v101 uint64 = 4294967295
                            for fi102 := 0; fi102 < 1; fi102++ {
                                var v103 int = (-(fi100))
                                _ = fi102
                                _ = v103
                            }
                            _ = v101
                        }
                        _ = fi100
                    }
                    v104 := []int{-1, 774, 1}
                    fmt.Println(v1)
                    _ = v104
                }
            }
        }
        _ = v96
        _ = v97
    } else {
        var v105 int8 = 42
        var v106 uint = v77
        _ = v105
        _ = v106
    }
    v76 = (true && v76)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    _ = v1
    _ = v2
    _ = v3
    _ = v76
    _ = v77
}
