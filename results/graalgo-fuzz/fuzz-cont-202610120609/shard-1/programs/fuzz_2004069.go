package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint8; F1 string; F2 uint }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint32; F1 int16; F2 uint }
type GenS7 struct { F0 uint8; F1 uint64 }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 int16 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 uint; F1 uint16 }
type GenS15 struct { F0 uint64 }
type GenS16 struct { F0 int }
type GenS17 struct { F0 uint16; F1 uint16 }
type GenS18 struct { F0 int; F1 uint16; F2 int16 }
type GenS19 struct { F0 uint16; F1 bool; F2 int16 }
type GenS20 struct { F0 uint; F1 uint64 }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 int16; F1 string; F2 uint32 }
type GenS23 struct { F0 string; F1 int32; F2 int16 }
func genFn0(p0 int16, p1 int, p2 uint16) string {
    v6 := []int{}
    p2 += 501
    v7 := GenS1{F0: (p2 - p2), F1: 32767}
    for fi8 := 0; fi8 < 5; fi8++ {
        if true {
            p2 += p2
            if true {
                if (("" + "\n\t") != "foo bar") {
                    p0 += -1
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v9 := GenS2{F0: 255, F1: ("hello" + "hello"), F2: 18446744073709551615}
                    _ = v9
                }
            } else {
                var v10 uint64 = 1
                v11 := GenS1{F0: 32767, F1: 60}
                if false {
                    var v12 uint = 42
                    if (127 >= 1) {
                        v13 := GenS2{F0: 0, F1: "", F2: (v12 - v12)}
                        _ = v13
                    } else {
                        fmt.Println(v12)
                    }
                    if len(v6) > 0 {
                        v6[0] = fi8
                    }
                    _ = v12
                }
                _ = v10
                _ = v11
            }
        }
        var v14 int8 = -1
        _ = fi8
        _ = v14
    }
    fmt.Println(p2, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return (("\n\t" + "") + "a")
}
func genFn1() int8 {
    var v15 int8 = -128
    if false {
        for fi16 := 0; fi16 < 6; fi16++ {
            if (0.0 != (1.5 / 1e10)) {
                for fi17 := 0; fi17 < 6; fi17++ {
                    switch "hello" {
                    case "a":
                        switch 42 {
                        case 1:
                            var v18 int8 = v15
                            _ = v18
                        case -163:
                            var v19 bool = false
                            _ = v19
                        case 32767:
                            switch ("a" + ("a" + "")) {
                            case "\n\t":
                                fmt.Println(v15, fi17)
                            default:
                                v20 := []int{(fi17 + 32767), fi16}
                                _ = v20
                            }
                        default:
                            v21 := GenS2{F0: 0, F1: ((("foo bar" + ((("hello" + "foo bar") + ("" + ("\n\t" + ("a" + (("foo bar" + "foo bar") + "\n\t"))))) + "foo bar")) + ("foo bar" + "foo bar")) + "hello"), F2: 0}
                            v22 := []int{(fi16 / (((-(fi16))) | 1)), fi17, fi17, (fi17 % ((fi17) | 1))}
                            var v23 int32 = 0
                            var v24 uint32 = 65535
                            if ("\n\t" == (("foo bar" + "\n\t") + ("" + ("" + ("" + "foo bar"))))) {
                                v24 -= v24
                            } else {
                                v23 -= -661
                            }
                            _ = v21
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                    }
                    var v25 float32 = 0.0
                    _ = fi17
                    _ = v25
                }
                v15 -= (v15 & (v15 & (v15 >> ((931) & 63))))
            } else {
                fmt.Println(fi16, fi16)
            }
            _ = fi16
        }
        var v26 bool = true
        _ = v26
    }
    _ = v15
    return (v15 - (^(v15)))
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]string)
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    v5 := []string{((("\n\t" + "") + ("" + ("hello" + (("hello" + "\n\t") + "hello")))) + ("\n\t" + ("" + "foo bar")))}
    var v27 float32 = float32((0.0 * 1.5))
    var v28 bool = true
    v27 = v27
    for fi29 := 0; fi29 < 8; fi29++ {
        switch (("" + "\n\t") + "") {
        case "":
            v28 = false
        case "a":
            v28 = v28
        }
        _ = fi29
    }
    if v28 {
        v30 := GenS1{F0: 888, F1: 65535}
        _ = v30
    }
    if (42 <= 0) {
        v28 = v28
        v28 = v28
        switch 0 {
        case 1:
            for fi31 := 0; fi31 < 1; fi31++ {
                var v32 int = (fi31 & (fi31 | fi31))
                _ = len(v3)
                if (-978 >= 0) {
                    v28 = true
                } else {
                    fmt.Println(v27)
                }
                _ = fi31
                _ = v32
            }
        default:
            var v33 uint32 = 408
            _ = v33
        }
    } else {
        for fi34 := 0; fi34 < 13; fi34++ {
            fmt.Println(v27)
            _ = fi34
        }
    }
    if v28 {
        v27 = (v27 / ((v27 - v27) * v27))
        v35 := GenS2{F0: 202, F1: "\n\t", F2: 65535}
        if v28 {
            var v36 int32 = 835
            _ = v36
        } else {
            for fi37 := 0; fi37 < 10; fi37++ {
                for fi38 := 0; fi38 < 16; fi38++ {
                    if (0 > 366) {
                        v28 = (true || false)
                        var v39 int64 = 9223372036854775807
                        fmt.Println(v27)
                        _ = v39
                    } else {
                        if (!v28) {
                            v27 = 1e10
                            v40 := GenS2{F0: 0, F1: "", F2: 255}
                            v2[-2147483648] = "hello"
                            _ = v40
                        }
                        switch "" {
                        case "foo bar":
                            v27 += (float32(0.0) / (1.5 * v27))
                            v41 := GenS1{F0: 555, F1: 2}
                            if ((v28 && (1e10 < ((0.0 - 1.5) - float64(v27)))) && (!v28)) {
                                v28 = true
                                v42 := GenS3{F0: -938, F1: (v28 || (4294967295 < fi37)), F2: 0}
                                _ = v42
                            } else {
                                v1 = append(v1, 0)
                                if true {
                                    v27 = float32(0.1)
                                } else {
                                    if v28 {
                                        v27 = (float32((0.0 - 0.0)) / (v27 + (float32(1.5) - (v27 + (float32(0.1) / v27)))))
                                        if (v28 || v28) {
                                            var v43 int8 = -128
                                            v43 = (v43 | genFn1())
                                            _ = v43
                                        } else {
                                            if ("foo bar" > ("hello" + ("foo bar" + "a"))) {
                                                v44 := GenS3{F0: 1000, F1: (42 > 42), F2: 255}
                                                var v45 int32 = 1
                                                var v46 int8 = genFn1()
                                                _ = v44
                                                _ = v45
                                                _ = v46
                                            } else {
                                                var v47 int32 = -142
                                                switch "foo bar" {
                                                case "\n\t":
                                                    v28 = (65535 < 535)
                                                case "a":
                                                    var v48 bool = (42 < v47)
                                                    _ = v48
                                                default:
                                                    if (!((!v28) || false)) {
                                                        var v49 string = "hello"
                                                        _ = v49
                                                    }
                                                }
                                                v27 += -2.25
                                                _ = v47
                                            }
                                            v27 = 3.14159
                                        }
                                        if (true && v28) {
                                            var v50 int8 = 100
                                            if ((v27 / v27) != v27) {
                                                v28 = false
                                                v51 := GenS4{F0: 42}
                                                if (127 <= 0) {
                                                    if false {
                                                        var v52 uint16 = 25
                                                        var v53 int64 = 0
                                                        var v54 uint64 = 42
                                                        _ = v52
                                                        _ = v53
                                                        _ = v54
                                                    } else {
                                                        delete(v4, 42)
                                                        var v55 uint8 = 0
                                                        _ = v55
                                                    }
                                                    if (0 >= 621) {
                                                        fmt.Println(v27, v50, v50)
                                                    } else {
                                                        var v56 string = ""
                                                        switch -970 {
                                                        case 131:
                                                            var v57 int = 1
                                                            _ = v57
                                                        case 70:
                                                            if v28 {
                                                                if ((!v28) && (!(!((v50 / (((v50 >> ((18446744073709551615) & 63))) | 1)) < v50)))) {
                                                                    if false {
                                                                        var v58 uint = 18446744073709551615
                                                                        fmt.Println(v28)
                                                                        _ = v58
                                                                    } else {
                                                                        switch ((v56 + v56) + "hello") {
                                                                        case "hello":
                                                                            fmt.Println(v50)
                                                                        }
                                                                        v59 := GenS2{F0: 62, F1: "a", F2: 0}
                                                                        var v60 uint = 1
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                    switch v56 {
                                                                    case "a":
                                                                        fmt.Println(v28)
                                                                    case "\n\t":
                                                                        switch fi38 {
                                                                        case 9223372036854775807:
                                                                            fmt.Println(v50)
                                                                        case 1:
                                                                            v61 := v4[2147483647]
                                                                            _ = v61
                                                                        default:
                                                                            switch v56 {
                                                                            case "hello":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            case "":
                                                                                if ((69 != 127) || ((v56 + v56) > "hello")) {
                                                                                    v62 := GenS1{F0: 1, F1: 0}
                                                                                    if true {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        fmt.Println(fi38, fi38, fi37)
                                                                                    }
                                                                                    var v63 bool = (v50 != (v50 ^ genFn1()))
                                                                                    _ = v62
                                                                                    _ = v63
                                                                                }
                                                                            }
                                                                        }
                                                                    case "hello":
                                                                        v64 := v2[(^(fi37))]
                                                                        _ = v64
                                                                    default:
                                                                        switch 976 {
                                                                        case -2147483648:
                                                                            if v28 {
                                                                                v65 := GenS2{F0: 0, F1: v56, F2: 681}
                                                                                v66 := v4[753]
                                                                                _ = v65
                                                                                _ = v66
                                                                            } else {
                                                                                v27 -= 0.0
                                                                            }
                                                                        default:
                                                                            switch (("a" + ("\n\t" + (v56 + ((v56 + (v56 + v56)) + v56)))) + v56) {
                                                                            case "":
                                                                                if (!(!true)) {
                                                                                    v50 -= (^(v50))
                                                                                    if v28 {
                                                                                        var v67 uint32 = 0
                                                                                        if (-9223372036854775808 > 53) {
                                                                                            if v28 {
                                                                                                if len(v5) >= 2 {
                                                                                                    v5 = v5[1:len(v5)]
                                                                                                }
                                                                                                _ = len(v2)
                                                                                            } else {
                                                                                                switch fi38 {
                                                                                                case -889:
                                                                                                    var v68 uint32 = (v67 - 0)
                                                                                                    _ = v68
                                                                                                case 128:
                                                                                                    fmt.Println(v28, v56, fi37)
                                                                                                case 127:
                                                                                                    v27 += v27
                                                                                                }
                                                                                            }
                                                                                            var v69 int = (fi38 / (((fi37 | (fi37 / ((fi37) | 1)))) | 1))
                                                                                            v27 = (v27 - ((0.0 - float32((1.5 - 1.5))) - (v27 * 1.5)))
                                                                                            var v70 float64 = 1e10
                                                                                            v56 = "hello"
                                                                                            v3[192] = (v67 | (v67 % (((v67 + (v67 << ((0) & 63)))) | 1)))
                                                                                            v2[v69] = v56
                                                                                            _ = v69
                                                                                            _ = v70
                                                                                        } else {
                                                                                            v56 = (((v56 + "a") + v56) + v56)
                                                                                            v27 = v27
                                                                                        }
                                                                                        fmt.Println(v27, v50, v56)
                                                                                        _ = v67
                                                                                    } else {
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            var v71 int8 = (-(v50))
                                                                                            _ = v71
                                                                                        case "":
                                                                                            v27 = v27
                                                                                        case "a":
                                                                                            var v72 uint8 = 0
                                                                                            _ = v72
                                                                                        default:
                                                                                            v73 := GenS5{F0: 2}
                                                                                            if v28 {
                                                                                                switch (v56 + v56) {
                                                                                                case "":
                                                                                                    if true {
                                                                                                        var v74 uint16 = 152
                                                                                                        if v28 {
                                                                                                            fmt.Println(v56)
                                                                                                        } else {
                                                                                                            if (18446744073709551615 > 65535) {
                                                                                                                var v75 int = 9223372036854775807
                                                                                                                switch v56 {
                                                                                                                case "":
                                                                                                                    fmt.Println(v56, v75, v28)
                                                                                                                case "\n\t":
                                                                                                                    if false {
                                                                                                                        v50 += int8(fi37)
                                                                                                                        switch "\n\t" {
                                                                                                                        case "hello":
                                                                                                                            v76 := GenS5{F0: (v74 / (((v74 % ((565) | 1))) | 1))}
                                                                                                                            _ = v76
                                                                                                                        case "a":
                                                                                                                            fmt.Println(v75, v74, v75)
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v77 := GenS5{F0: v74}
                                                                                                                    _ = v77
                                                                                                                default:
                                                                                                                    fmt.Println(v75)
                                                                                                                }
                                                                                                                _ = v75
                                                                                                            } else {
                                                                                                                switch (v56 + "hello") {
                                                                                                                case "":
                                                                                                                    var v78 string = v56
                                                                                                                    _ = v78
                                                                                                                case "hello":
                                                                                                                    v79 := GenS6{F0: 3, F1: 0, F2: 42}
                                                                                                                    var v80 bool = (-1 < int(v50))
                                                                                                                    v74 -= (v74 / (((v74 / (((v74 & v74)) | 1))) | 1))
                                                                                                                    v4[922] = (!v28)
                                                                                                                    v74 += v74
                                                                                                                    var v81 int8 = (^(v50))
                                                                                                                    _ = v79
                                                                                                                    _ = v80
                                                                                                                    _ = v81
                                                                                                                case "a":
                                                                                                                    if true {
                                                                                                                        v28 = v28
                                                                                                                        delete(v4, 110)
                                                                                                                        var v82 string = ("a" + (v56 + ((v56 + v56) + v56)))
                                                                                                                        _ = v82
                                                                                                                    }
                                                                                                                default:
                                                                                                                    v56 = ""
                                                                                                                }
                                                                                                                v4[842] = (((float32((((float64((v27 - v27)) + 0.0) - 1.5) - float64((v27 + v27)))) + v27) - (0.1 + v27)) == float32((-2.25 + (0.1 + 0.0))))
                                                                                                                v74 = (v74 / ((589) | 1))
                                                                                                            }
                                                                                                            var v83 float64 = float64((((v27 / (v27 * (v27 - v27))) + 0.1) / (v27 / (float32(1e10) + v27))))
                                                                                                            v84 := GenS6{F0: 562, F1: 1, F2: 681}
                                                                                                            v56 = v56
                                                                                                            _ = v83
                                                                                                            _ = v84
                                                                                                        }
                                                                                                        if (!false) {
                                                                                                            var v85 int64 = -337
                                                                                                            v50 = (v50 * v50)
                                                                                                            if v28 {
                                                                                                                v74 += v74
                                                                                                                v86 := GenS2{F0: 148, F1: v56, F2: 0}
                                                                                                                var v87 string = (v56 + ("" + v56))
                                                                                                                _ = v86
                                                                                                                _ = v87
                                                                                                            }
                                                                                                            _ = v85
                                                                                                        }
                                                                                                        _ = v74
                                                                                                    } else {
                                                                                                        if ((769 >= 4294967295) || (0.1 == 1.5)) {
                                                                                                            var v88 uint32 = 701
                                                                                                            v28 = (122 != 4294967295)
                                                                                                            var v89 uint = 0
                                                                                                            switch fi37 {
                                                                                                            case -380:
                                                                                                                switch fi38 {
                                                                                                                case 42:
                                                                                                                    fmt.Println(v28, v89, v88)
                                                                                                                case 763:
                                                                                                                    v89 -= 1
                                                                                                                    v56 = v56
                                                                                                                default:
                                                                                                                    switch v56 {
                                                                                                                    case "foo bar":
                                                                                                                        v90 := GenS7{F0: 138, F1: 42}
                                                                                                                        v28 = (511 > 42)
                                                                                                                        fmt.Println(v27, v89, fi37)
                                                                                                                        _ = v90
                                                                                                                    case "":
                                                                                                                        v27 -= 1e10
                                                                                                                    case "hello":
                                                                                                                        delete(v3, fi37)
                                                                                                                    default:
                                                                                                                        if (!v28) {
                                                                                                                            v89 += (v89 % (((v89 - (^(v89)))) | 1))
                                                                                                                        } else {
                                                                                                                            var v91 uint16 = 42
                                                                                                                            switch (fi38 % (((fi38 / ((32767) | 1))) | 1)) {
                                                                                                                            case 32768:
                                                                                                                                if (fi37 == fi37) {
                                                                                                                                    _ = len(v5)
                                                                                                                                    _ = len(v2)
                                                                                                                                    v92 := GenS3{F0: 1000, F1: true, F2: 4}
                                                                                                                                    var v93 uint32 = v88
                                                                                                                                    var v94 string = ""
                                                                                                                                    if (v28 && ((!v28) && v28)) {
                                                                                                                                        var v95 int16 = -1
                                                                                                                                        if v28 {
                                                                                                                                            switch (fi38 / (((-(fi38))) | 1)) {
                                                                                                                                            case -1:
                                                                                                                                                switch ("\n\t" + "a") {
                                                                                                                                                case "hello":
                                                                                                                                                    fmt.Println(v89)
                                                                                                                                                case "":
                                                                                                                                                    v96 := GenS8{F0: (v91 - 0)}
                                                                                                                                                    _ = v96
                                                                                                                                                }
                                                                                                                                            case 1:
                                                                                                                                                v97 := v4[v88]
                                                                                                                                                _ = v97
                                                                                                                                            case 9223372036854775807:
                                                                                                                                                v2[2147483647] = ("a" + (("hello" + "") + ((((v56 + ((v56 + "hello") + "foo bar")) + v56) + ("a" + v94)) + ("hello" + ("hello" + v94)))))
                                                                                                                                            default:
                                                                                                                                                v5 = append(v5, "\n\t")
                                                                                                                                            }
                                                                                                                                            if v28 {
                                                                                                                                                v56 = (v56 + "\n\t")
                                                                                                                                                if true {
                                                                                                                                                    v98 := GenS3{F0: v95, F1: true, F2: 255}
                                                                                                                                                    _ = v98
                                                                                                                                                }
                                                                                                                                                v27 = v27
                                                                                                                                            } else {
                                                                                                                                                if len(v1) >= 2 {
                                                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                                                }
                                                                                                                                                v28 = false
                                                                                                                                                v50 = v50
                                                                                                                                            }
                                                                                                                                            v95 = (v95 / ((v95) | 1))
                                                                                                                                            var v99 int16 = (v95 ^ -183)
                                                                                                                                            _ = v99
                                                                                                                                        }
                                                                                                                                        v89 = v89
                                                                                                                                        fmt.Println(v56)
                                                                                                                                        if true {
                                                                                                                                            switch v56 {
                                                                                                                                            case "hello":
                                                                                                                                                v50 = genFn1()
                                                                                                                                            case "a":
                                                                                                                                                _ = len(v3)
                                                                                                                                            case "\n\t":
                                                                                                                                                var v100 uint = 1
                                                                                                                                                _ = v100
                                                                                                                                            }
                                                                                                                                            var v101 bool = (false || ((0.1 + -2.25) == float64(v27)))
                                                                                                                                            _ = v101
                                                                                                                                        }
                                                                                                                                        _ = v95
                                                                                                                                    } else {
                                                                                                                                        v102 := GenS9{F0: 1000}
                                                                                                                                        v3[fi37] = uint32(v91)
                                                                                                                                        switch 2147483647 {
                                                                                                                                        case 1:
                                                                                                                                            if (!v28) {
                                                                                                                                                switch (v56 + v94) {
                                                                                                                                                case "":
                                                                                                                                                    v103 := GenS10{F0: 0}
                                                                                                                                                    var v104 int16 = 509
                                                                                                                                                    v56 = ((("" + v56) + v94) + (v94 + v94))
                                                                                                                                                    _ = v103
                                                                                                                                                    _ = v104
                                                                                                                                                default:
                                                                                                                                                    v89 = v89
                                                                                                                                                }
                                                                                                                                                fmt.Println(v56, v89)
                                                                                                                                            } else {
                                                                                                                                                if v28 {
                                                                                                                                                    v105 := GenS2{F0: 1, F1: v56, F2: uint(v91)}
                                                                                                                                                    _ = v105
                                                                                                                                                } else {
                                                                                                                                                    var v106 int64 = 1
                                                                                                                                                    v107 := GenS4{F0: v89}
                                                                                                                                                    _ = v106
                                                                                                                                                    _ = v107
                                                                                                                                                }
                                                                                                                                                v108 := v3[int(v88)]
                                                                                                                                                v93 = (v108 ^ v88)
                                                                                                                                                if (float32(0.0) == v27) {
                                                                                                                                                    if (!(((v27 / float32(-2.25)) - ((v27 * 3.14159) * float32(3.14159))) > 1e10)) {
                                                                                                                                                        v5 = append(v5, (v94 + ("hello" + "foo bar")))
                                                                                                                                                        if (!(0 == 0)) {
                                                                                                                                                            switch -151 {
                                                                                                                                                            case 127:
                                                                                                                                                                v1 = append(v1, 2147483647)
                                                                                                                                                            default:
                                                                                                                                                                var v109 int8 = 100
                                                                                                                                                                _ = v109
                                                                                                                                                            }
                                                                                                                                                            var v110 string = v56
                                                                                                                                                            _ = v110
                                                                                                                                                        } else {
                                                                                                                                                            if true {
                                                                                                                                                                var v111 int8 = (v50 - (v50 | 100))
                                                                                                                                                                v112 := GenS3{F0: -32768, F1: v28, F2: 127}
                                                                                                                                                                _ = v111
                                                                                                                                                                _ = v112
                                                                                                                                                            }
                                                                                                                                                            switch v94 {
                                                                                                                                                            case "\n\t":
                                                                                                                                                                v3[-1] = 2147483648
                                                                                                                                                            }
                                                                                                                                                            var v113 int8 = genFn1()
                                                                                                                                                            if ((1.5 - (1e10 - 0.0)) <= 0.1) {
                                                                                                                                                                switch fi37 {
                                                                                                                                                                case 220:
                                                                                                                                                                    v113 = (v113 >> ((18446744073709551615) & 63))
                                                                                                                                                                case 128:
                                                                                                                                                                    if len(v1) > 0 {
                                                                                                                                                                        v1[0] = 2147483648
                                                                                                                                                                    }
                                                                                                                                                                case -1:
                                                                                                                                                                    _ = len(v5)
                                                                                                                                                                default:
                                                                                                                                                                    v114 := GenS1{F0: (v91 * v91), F1: 8}
                                                                                                                                                                    switch ("\n\t" + (v56 + "\n\t")) {
                                                                                                                                                                    case "":
                                                                                                                                                                        v115 := GenS11{F0: (!v28)}
                                                                                                                                                                        _ = v115
                                                                                                                                                                    case "\n\t":
                                                                                                                                                                        switch fi38 {
                                                                                                                                                                        case 128:
                                                                                                                                                                            if len(v5) > 0 {
                                                                                                                                                                                v5[0] = (v94 + (v94 + v94))
                                                                                                                                                                            }
                                                                                                                                                                            _ = len(v2)
                                                                                                                                                                        case -9223372036854775808:
                                                                                                                                                                            var v116 uint32 = (v93 % (((v108 / (((v93 * v93)) | 1))) | 1))
                                                                                                                                                                            _ = v116
                                                                                                                                                                        default:
                                                                                                                                                                            if len(v1) >= 2 {
                                                                                                                                                                                v1 = v1[1:len(v1)]
                                                                                                                                                                            }
                                                                                                                                                                        }
                                                                                                                                                                    case "hello":
                                                                                                                                                                        if v28 {
                                                                                                                                                                            if (!(!v28)) {
                                                                                                                                                                                var v117 int16 = 756
                                                                                                                                                                                if false {
                                                                                                                                                                                    v27 = (v27 / float32(3.14159))
                                                                                                                                                                                    if (v28 && true) {
                                                                                                                                                                                        delete(v2, 2147483647)
                                                                                                                                                                                    }
                                                                                                                                                                                    var v118 bool = (float64(-2.25) <= 0.0)
                                                                                                                                                                                    _ = v118
                                                                                                                                                                                } else {
                                                                                                                                                                                    if (210 != 128) {
                                                                                                                                                                                        var v119 bool = (!false)
                                                                                                                                                                                        var v120 int = (-(fi38))
                                                                                                                                                                                        _ = v119
                                                                                                                                                                                        _ = v120
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v108 += (v88 ^ v108)
                                                                                                                                                                                        switch fi38 {
                                                                                                                                                                                        case 633:
                                                                                                                                                                                            v121 := GenS4{F0: 95}
                                                                                                                                                                                            _ = v121
                                                                                                                                                                                        case 2147483647:
                                                                                                                                                                                            _ = len(v2)
                                                                                                                                                                                        case 32768:
                                                                                                                                                                                            fmt.Println(v108, v93)
                                                                                                                                                                                        }
                                                                                                                                                                                        delete(v4, v93)
                                                                                                                                                                                        v122 := v2[985]
                                                                                                                                                                                        _ = v122
                                                                                                                                                                                    }
                                                                                                                                                                                    if v28 {
                                                                                                                                                                                        delete(v4, v88)
                                                                                                                                                                                        fmt.Println(v89)
                                                                                                                                                                                        v123 := GenS12{F0: 255}
                                                                                                                                                                                        var v124 bool = true
                                                                                                                                                                                        v93 = (v88 * (v93 >> ((680) & 63)))
                                                                                                                                                                                        _ = v123
                                                                                                                                                                                        _ = v124
                                                                                                                                                                                    } else {
                                                                                                                                                                                        if true {
                                                                                                                                                                                            var v125 float32 = v27
                                                                                                                                                                                            v126 := GenS13{F0: -2147483648}
                                                                                                                                                                                            delete(v4, v88)
                                                                                                                                                                                            var v127 uint64 = 158
                                                                                                                                                                                            fmt.Println(v93, v50)
                                                                                                                                                                                            v128 := GenS14{F0: (v89 | 42), F1: (v91 - 171)}
                                                                                                                                                                                            delete(v4, (v88 - uint32(v91)))
                                                                                                                                                                                            _ = v125
                                                                                                                                                                                            _ = v126
                                                                                                                                                                                            _ = v127
                                                                                                                                                                                            _ = v128
                                                                                                                                                                                        }
                                                                                                                                                                                        fmt.Println(v50, v28)
                                                                                                                                                                                    }
                                                                                                                                                                                    if (42 != uint(v88)) {
                                                                                                                                                                                        v28 = v28
                                                                                                                                                                                        if len(v1) > 0 {
                                                                                                                                                                                            v1[0] = 2147483648
                                                                                                                                                                                        }
                                                                                                                                                                                    } else {
                                                                                                                                                                                        v27 = (0.1 * v27)
                                                                                                                                                                                        var v129 int16 = (-(v117))
                                                                                                                                                                                        _ = v129
                                                                                                                                                                                    }
                                                                                                                                                                                }
                                                                                                                                                                                _ = v117
                                                                                                                                                                            } else {
                                                                                                                                                                                if v28 {
                                                                                                                                                                                    v88 += v108
                                                                                                                                                                                }
                                                                                                                                                                                v56 = v56
                                                                                                                                                                                v88 += (v93 >> (((v89 & 65535)) & 63))
                                                                                                                                                                            }
                                                                                                                                                                            switch ((((v94 + v56) + v56) + "hello") + (((v94 + (((v56 + v56) + ((v94 + (v94 + "a")) + ("a" + ("" + v94)))) + "foo bar")) + (v56 + ("hello" + ("hello" + "\n\t")))) + "foo bar")) {
                                                                                                                                                                            case "a":
                                                                                                                                                                                if (!(v28 && v28)) {
                                                                                                                                                                                    var v130 uint16 = v91
                                                                                                                                                                                    _ = v130
                                                                                                                                                                                } else {
                                                                                                                                                                                    v131 := GenS7{F0: 0, F1: 18446744073709551615}
                                                                                                                                                                                    v108 = uint32(v91)
                                                                                                                                                                                    var v132 int8 = v113
                                                                                                                                                                                    fmt.Println(v28)
                                                                                                                                                                                    _ = v131
                                                                                                                                                                                    _ = v132
                                                                                                                                                                                }
                                                                                                                                                                            }
                                                                                                                                                                            v27 = -2.25
                                                                                                                                                                        } else {
                                                                                                                                                                            fmt.Println(v88, fi38, v108)
                                                                                                                                                                            var v133 float32 = (v27 - v27)
                                                                                                                                                                            var v134 int32 = -1
                                                                                                                                                                            var v135 int32 = (v134 - v134)
                                                                                                                                                                            _ = v133
                                                                                                                                                                            _ = v134
                                                                                                                                                                            _ = v135
                                                                                                                                                                        }
                                                                                                                                                                        var v136 uint8 = 186
                                                                                                                                                                        fmt.Println(v28, v50)
                                                                                                                                                                        _ = v136
                                                                                                                                                                    }
                                                                                                                                                                    _ = v114
                                                                                                                                                                }
                                                                                                                                                                fmt.Println(v91)
                                                                                                                                                                v5 = append(v5, v94)
                                                                                                                                                                v88 = 42
                                                                                                                                                            }
                                                                                                                                                            var v137 float64 = (-2.25 - ((-2.25 - 0.0) / 3.14159))
                                                                                                                                                            v89 -= v89
                                                                                                                                                            _ = v113
                                                                                                                                                            _ = v137
                                                                                                                                                        }
                                                                                                                                                        var v138 string = "a"
                                                                                                                                                        var v139 uint = uint(v108)
                                                                                                                                                        var v140 int = (fi38 - (fi38 | int(v91)))
                                                                                                                                                        var v141 string = (v138 + v138)
                                                                                                                                                        if (v28 || v28) {
                                                                                                                                                            var v142 uint = uint(v108)
                                                                                                                                                            var v143 int64 = 1
                                                                                                                                                            _ = v142
                                                                                                                                                            _ = v143
                                                                                                                                                        } else {
                                                                                                                                                            v50 = v50
                                                                                                                                                        }
                                                                                                                                                        _ = v138
                                                                                                                                                        _ = v139
                                                                                                                                                        _ = v140
                                                                                                                                                        _ = v141
                                                                                                                                                    } else {
                                                                                                                                                        if v28 {
                                                                                                                                                            switch v56 {
                                                                                                                                                            case "\n\t":
                                                                                                                                                                v94 = (v56 + "hello")
                                                                                                                                                            case "a":
                                                                                                                                                                v144 := GenS1{F0: (v91 * (v91 - 255)), F1: (v91 | (v91 >> ((1) & 63)))}
                                                                                                                                                                _ = v144
                                                                                                                                                            case "hello":
                                                                                                                                                                v145 := GenS9{F0: -32768}
                                                                                                                                                                _ = v145
                                                                                                                                                            default:
                                                                                                                                                                var v146 string = (v56 + v94)
                                                                                                                                                                _ = v146
                                                                                                                                                            }
                                                                                                                                                            var v147 int8 = v50
                                                                                                                                                            fmt.Println(v89, fi38, fi37)
                                                                                                                                                            v91 = v91
                                                                                                                                                            _ = v147
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    v108 = 2
                                                                                                                                                } else {
                                                                                                                                                    if len(v5) >= 2 {
                                                                                                                                                        v5 = v5[1:len(v5)]
                                                                                                                                                    }
                                                                                                                                                    v27 = v27
                                                                                                                                                }
                                                                                                                                                _ = v108
                                                                                                                                            }
                                                                                                                                        default:
                                                                                                                                            var v148 bool = ((true && v28) && false)
                                                                                                                                            _ = v148
                                                                                                                                        }
                                                                                                                                        _ = v102
                                                                                                                                    }
                                                                                                                                    switch (fi37 >> ((139) & 63)) {
                                                                                                                                    case 337:
                                                                                                                                        var v149 string = v94
                                                                                                                                        v150 := GenS11{F0: v28}
                                                                                                                                        if v28 {
                                                                                                                                            fmt.Println(v149)
                                                                                                                                        } else {
                                                                                                                                            var v151 string = ""
                                                                                                                                            v152 := GenS3{F0: 1, F1: (v28 || (false || v28)), F2: 42}
                                                                                                                                            v28 = false
                                                                                                                                            _ = v151
                                                                                                                                            _ = v152
                                                                                                                                        }
                                                                                                                                        _ = v149
                                                                                                                                        _ = v150
                                                                                                                                    case -2147483648:
                                                                                                                                        if (42 >= -977) {
                                                                                                                                            if (v28 && v28) {
                                                                                                                                                v28 = (v89 < 255)
                                                                                                                                                v2[(fi38 << ((v89) & 63))] = v56
                                                                                                                                            } else {
                                                                                                                                                v153 := GenS6{F0: 0, F1: 727, F2: 507}
                                                                                                                                                delete(v4, v93)
                                                                                                                                                _ = v153
                                                                                                                                            }
                                                                                                                                        }
                                                                                                                                    case 2147483647:
                                                                                                                                        var v154 int64 = 925
                                                                                                                                        _ = v154
                                                                                                                                    default:
                                                                                                                                        var v155 int32 = -1
                                                                                                                                        _ = v155
                                                                                                                                    }
                                                                                                                                    _ = v92
                                                                                                                                    _ = v93
                                                                                                                                    _ = v94
                                                                                                                                }
                                                                                                                                switch int(v88) {
                                                                                                                                case 1:
                                                                                                                                    fmt.Println(v88)
                                                                                                                                }
                                                                                                                            default:
                                                                                                                                switch fi37 {
                                                                                                                                case 705:
                                                                                                                                    fmt.Println(v27, v91, v91)
                                                                                                                                case 342:
                                                                                                                                    if v28 {
                                                                                                                                        var v156 int = -259
                                                                                                                                        if v28 {
                                                                                                                                            var v157 int64 = 188
                                                                                                                                            var v158 uint32 = 854
                                                                                                                                            var v159 uint = (v89 - (v89 * 62))
                                                                                                                                            v160 := GenS15{F0: 42}
                                                                                                                                            if v28 {
                                                                                                                                                v50 = (v50 >> ((v89) & 63))
                                                                                                                                            }
                                                                                                                                            _ = v157
                                                                                                                                            _ = v158
                                                                                                                                            _ = v159
                                                                                                                                            _ = v160
                                                                                                                                        }
                                                                                                                                        _ = v156
                                                                                                                                    }
                                                                                                                                case 951:
                                                                                                                                    v161 := GenS6{F0: (v88 & (v88 * v88)), F1: 42, F2: (v89 * (v89 - 941))}
                                                                                                                                    _ = v161
                                                                                                                                }
                                                                                                                            }
                                                                                                                            _ = v91
                                                                                                                        }
                                                                                                                        var v162 uint = (v89 << ((v89) & 63))
                                                                                                                        _ = v162
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                var v163 float64 = (0.1 / 3.14159)
                                                                                                                _ = v163
                                                                                                            }
                                                                                                            fmt.Println(v50)
                                                                                                            v50 -= (v50 & (v50 / ((v50) | 1)))
                                                                                                            var v164 bool = (65535 > v89)
                                                                                                            v164 = false
                                                                                                            _ = len(v5)
                                                                                                            _ = v164
                                                                                                            _ = v88
                                                                                                            _ = v89
                                                                                                        } else {
                                                                                                            var v165 float32 = v27
                                                                                                            v166 := GenS10{F0: 610}
                                                                                                            fmt.Println(fi38, fi37, v56)
                                                                                                            _ = v165
                                                                                                            _ = v166
                                                                                                        }
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v167 float32 = -2.25
                                                                                                _ = v167
                                                                                            }
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                    }
                                                                    v50 = (^(v50))
                                                                }
                                                                if false {
                                                                    v28 = (5 > 255)
                                                                    v168 := v3[(fi38 << ((131) & 63))]
                                                                    _ = v168
                                                                } else {
                                                                    if v28 {
                                                                        v56 = v56
                                                                        v169 := GenS1{F0: 65535, F1: 32767}
                                                                        v1 = append(v1, 2)
                                                                        _ = v169
                                                                    }
                                                                }
                                                            } else {
                                                                var v170 int = fi38
                                                                if (128 >= 255) {
                                                                    switch fi37 {
                                                                    case -9223372036854775808:
                                                                        delete(v4, 1)
                                                                    case 128:
                                                                        _ = len(v5)
                                                                    }
                                                                    if v28 {
                                                                        if false {
                                                                            v56 = (v56 + ("hello" + v56))
                                                                            if true {
                                                                                if false {
                                                                                    var v171 uint32 = 114
                                                                                    _ = v171
                                                                                } else {
                                                                                    if v28 {
                                                                                        v5 = append(v5, "a")
                                                                                        if true {
                                                                                            var v172 uint64 = 632
                                                                                            if v28 {
                                                                                                if v28 {
                                                                                                    if true {
                                                                                                        var v173 float64 = (float64(0.0) - -2.25)
                                                                                                        v28 = (181 >= 801)
                                                                                                        _ = v173
                                                                                                    } else {
                                                                                                        if ((v27 - (float32((3.14159 / 1.5)) / 1.5)) != 0.0) {
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = 4294967295
                                                                                                            }
                                                                                                        }
                                                                                                        var v174 uint64 = (^(v172))
                                                                                                        if (101 > 2) {
                                                                                                            var v175 bool = true
                                                                                                            _ = v175
                                                                                                        } else {
                                                                                                            v176 := GenS16{F0: fi38}
                                                                                                            v1 = append(v1, 381)
                                                                                                            if false {
                                                                                                                switch v56 {
                                                                                                                case "\n\t":
                                                                                                                    v170 += (v170 * -671)
                                                                                                                case "":
                                                                                                                    v56 = v56
                                                                                                                default:
                                                                                                                    switch fi37 {
                                                                                                                    case -587:
                                                                                                                        v27 = v27
                                                                                                                    case 42:
                                                                                                                        var v177 int = 4294967295
                                                                                                                        _ = v177
                                                                                                                    }
                                                                                                                }
                                                                                                                if len(v5) >= 2 {
                                                                                                                    v5 = v5[1:len(v5)]
                                                                                                                }
                                                                                                                var v178 float32 = (v27 * float32(1e10))
                                                                                                                _ = v178
                                                                                                            }
                                                                                                            _ = v176
                                                                                                        }
                                                                                                        _ = len(v4)
                                                                                                        v179 := GenS17{F0: 32767, F1: 453}
                                                                                                        _ = v174
                                                                                                        _ = v179
                                                                                                    }
                                                                                                    switch fi37 {
                                                                                                    case 974:
                                                                                                        var v180 string = v56
                                                                                                        _ = v180
                                                                                                    case 32767:
                                                                                                        var v181 int16 = 42
                                                                                                        _ = v181
                                                                                                    case 2147483647:
                                                                                                        fmt.Println(v172, v50)
                                                                                                    default:
                                                                                                        var v182 uint16 = 523
                                                                                                        _ = v182
                                                                                                    }
                                                                                                }
                                                                                                var v183 uint = 255
                                                                                                var v184 int = (fi37 - (-(fi37)))
                                                                                                _ = v183
                                                                                                _ = v184
                                                                                            }
                                                                                            v170 = -9223372036854775808
                                                                                            _ = v172
                                                                                        } else {
                                                                                            var v185 uint32 = 915
                                                                                            v186 := GenS5{F0: 2}
                                                                                            switch 127 {
                                                                                            case 1:
                                                                                                if true {
                                                                                                    v56 = ("" + "foo bar")
                                                                                                    if (!true) {
                                                                                                        fmt.Println(v28)
                                                                                                        if (2 < 255) {
                                                                                                            v187 := GenS18{F0: 9223372036854775807, F1: 2, F2: 1}
                                                                                                            switch ("hello" + v56) {
                                                                                                            case "\n\t":
                                                                                                                v188 := GenS15{F0: 2}
                                                                                                                _ = v188
                                                                                                            case "foo bar":
                                                                                                                if (0.1 == -2.25) {
                                                                                                                    if len(v1) >= 2 {
                                                                                                                        v1 = v1[1:len(v1)]
                                                                                                                    }
                                                                                                                    v50 -= -1
                                                                                                                    v170 = (v170 >> ((0) & 63))
                                                                                                                }
                                                                                                            case "":
                                                                                                                v189 := v2[(v170 - (fi38 * (v170 << ((255) & 63))))]
                                                                                                                _ = len(v5)
                                                                                                                _ = v189
                                                                                                            default:
                                                                                                                var v190 bool = v28
                                                                                                                _ = v190
                                                                                                            }
                                                                                                            var v191 uint64 = 366
                                                                                                            _ = v187
                                                                                                            _ = v191
                                                                                                        }
                                                                                                        var v192 int = (v170 - fi38)
                                                                                                        var v193 string = "a"
                                                                                                        _ = v192
                                                                                                        _ = v193
                                                                                                    } else {
                                                                                                        v170 += -743
                                                                                                        var v194 float32 = v27
                                                                                                        v194 += v27
                                                                                                        _ = v194
                                                                                                    }
                                                                                                    var v195 int64 = -1
                                                                                                    var v196 uint64 = 1
                                                                                                    _ = v195
                                                                                                    _ = v196
                                                                                                }
                                                                                            }
                                                                                            _ = v185
                                                                                            _ = v186
                                                                                        }
                                                                                    } else {
                                                                                        _ = len(v3)
                                                                                        v197 := GenS19{F0: 292, F1: true, F2: -716}
                                                                                        if len(v5) > 0 {
                                                                                            v5[0] = "\n\t"
                                                                                        }
                                                                                        _ = v197
                                                                                    }
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                                var v198 uint8 = 255
                                                                                fmt.Println(fi37, v50, v198)
                                                                                var v199 int32 = -201
                                                                                switch (fi38 % (((fi37 | v170)) | 1)) {
                                                                                case 1:
                                                                                    _ = len(v3)
                                                                                }
                                                                                switch v56 {
                                                                                case "":
                                                                                    var v200 int16 = -1
                                                                                    _ = v200
                                                                                case "foo bar":
                                                                                    v28 = (!true)
                                                                                default:
                                                                                    v201 := GenS2{F0: (v198 & v198), F1: "", F2: 0}
                                                                                    _ = v201
                                                                                }
                                                                                _ = v198
                                                                                _ = v199
                                                                            } else {
                                                                                switch "a" {
                                                                                case "":
                                                                                    var v202 int32 = 2147483647
                                                                                    _ = v202
                                                                                case "foo bar":
                                                                                    v50 = (v50 | v50)
                                                                                case "a":
                                                                                    v50 = v50
                                                                                }
                                                                            }
                                                                        } else {
                                                                            v50 = v50
                                                                            var v203 int16 = 683
                                                                            _ = v203
                                                                        }
                                                                        var v204 int = v170
                                                                        v204 = v204
                                                                        _ = v204
                                                                    } else {
                                                                        fmt.Println(v50)
                                                                        v205 := GenS4{F0: 0}
                                                                        if (42 <= 531) {
                                                                            v56 = "\n\t"
                                                                        } else {
                                                                            if v28 {
                                                                                fmt.Println(v28)
                                                                                v56 = v56
                                                                            } else {
                                                                                if false {
                                                                                    if len(v5) >= 2 {
                                                                                        v5 = v5[1:len(v5)]
                                                                                    }
                                                                                }
                                                                                var v206 uint64 = 4294967295
                                                                                _ = v206
                                                                            }
                                                                        }
                                                                        _ = v205
                                                                    }
                                                                    if v28 {
                                                                        fmt.Println(v50, v27)
                                                                        var v207 float32 = 0.1
                                                                        v208 := v3[(v170 / ((fi37) | 1))]
                                                                        _ = v207
                                                                        _ = v208
                                                                    }
                                                                } else {
                                                                    if false {
                                                                        if false {
                                                                            v209 := GenS20{F0: 42, F1: 982}
                                                                            if (-1 <= 32767) {
                                                                                if v28 {
                                                                                    v210 := GenS21{F0: 2147483647}
                                                                                    switch "hello" {
                                                                                    case "\n\t":
                                                                                        v211 := GenS17{F0: 0, F1: 741}
                                                                                        _ = v211
                                                                                    case "foo bar":
                                                                                        var v212 int64 = 610
                                                                                        _ = v212
                                                                                    case "a":
                                                                                        v170 = (fi37 % ((v170) | 1))
                                                                                    }
                                                                                    _ = len(v1)
                                                                                    _ = v210
                                                                                }
                                                                                v213 := GenS21{F0: 2147483647}
                                                                                v56 = "hello"
                                                                                _ = v213
                                                                            } else {
                                                                                fmt.Println(v50)
                                                                                v214 := v3[(v170 << ((2) & 63))]
                                                                                var v215 int64 = 0
                                                                                _ = v214
                                                                                _ = v215
                                                                            }
                                                                            v216 := GenS10{F0: -920}
                                                                            v5 = append(v5, ((v56 + (v56 + (("hello" + (("" + (v56 + v56)) + (v56 + "hello"))) + v56))) + ("foo bar" + "\n\t")))
                                                                            _ = v209
                                                                            _ = v216
                                                                        } else {
                                                                            fmt.Println(fi37, v28)
                                                                        }
                                                                    } else {
                                                                        var v217 uint64 = 648
                                                                        _ = v217
                                                                    }
                                                                    switch v56 {
                                                                    case "\n\t":
                                                                        var v218 int8 = (v50 / (((v50 * v50)) | 1))
                                                                        var v219 int = (v170 ^ (^(fi38)))
                                                                        _ = v218
                                                                        _ = v219
                                                                    case "":
                                                                        _ = len(v4)
                                                                    case "hello":
                                                                        v56 = (v56 + v56)
                                                                    default:
                                                                        var v220 uint = 336
                                                                        _ = v220
                                                                    }
                                                                    var v221 int = v170
                                                                    _ = v221
                                                                }
                                                                var v222 int64 = 0
                                                                v56 = "\n\t"
                                                                _ = v170
                                                                _ = v222
                                                            }
                                                            v223 := v3[(fi38 << ((0) & 63))]
                                                            v224 := GenS22{F0: 0, F1: "hello", F2: 2147483647}
                                                            _ = v223
                                                            _ = v224
                                                        default:
                                                            _ = len(v4)
                                                        }
                                                        _ = v56
                                                    }
                                                } else {
                                                    var v225 bool = v28
                                                    _ = v225
                                                }
                                                fmt.Println(v27, v27)
                                                var v226 bool = v28
                                                _ = v226
                                                _ = v51
                                            }
                                            _ = v50
                                        }
                                    }
                                    fmt.Println(v27, fi38, v28)
                                    v227 := GenS23{F0: "foo bar", F1: 597, F2: 32767}
                                    _ = v227
                                }
                                v228 := GenS1{F0: 42, F1: 32767}
                                _ = v228
                            }
                            _ = v41
                        }
                        fmt.Println(v28, v28, v28)
                    }
                    _ = fi38
                }
                _ = fi37
            }
            v229 := GenS21{F0: 42}
            _ = v229
        }
        _ = v35
    } else {
        v230 := GenS5{F0: 2}
        _ = v230
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if v28 {
        var v231 uint32 = 1
        _ = v231
    } else {
        v232 := GenS3{F0: 1000, F1: (1 > 219), F2: 173}
        _ = v232
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
