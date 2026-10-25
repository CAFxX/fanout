package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8; F2 bool }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32; F1 string }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int64; F1 uint16; F2 uint }

func main() {
    var v1 uint8 = 128
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 bool = (false && (45 == 42))
    var v5 int = -363
    var v6 uint16 = 393
    v1 += (v1 & v1)
    if (!(true || v4)) {
        v7 := GenS1{F0: 542}
        _ = v7
    } else {
        v5 -= v5
    }
    delete(v2, v5)
    v8 := GenS1{F0: 0}
    if (v6 < v6) {
        for fi9 := 0; fi9 < 3; fi9++ {
            for fi10 := 0; fi10 < 8; fi10++ {
                v4 = (!true)
                var v11 int8 = -1
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
        fmt.Println(v6, v5, v4)
    } else {
        v6 -= v6
        var v12 int64 = 9223372036854775807
        if true {
            v2[-750] = (((("hello" + "") + (("" + "hello") + "\n\t")) + "foo bar") + "\n\t")
            if false {
                for fi13 := 0; fi13 < 18; fi13++ {
                    fmt.Println(v6, fi13)
                    if v4 {
                        switch "" {
                        case "hello":
                            switch ((("hello" + "hello") + "foo bar") + "") {
                            case "hello":
                                var v14 int32 = 124
                                _ = v14
                            default:
                                fmt.Println(v12, v5)
                            }
                        case "\n\t":
                            v4 = (v6 < v6)
                        case "a":
                            v15 := []int{}
                            _ = v15
                        }
                    } else {
                        var v16 string = ("" + "\n\t")
                        _ = v16
                    }
                    _ = fi13
                }
                if v4 {
                    for fi17 := 0; fi17 < 4; fi17++ {
                        fmt.Println(v6, v1, v6)
                        _ = fi17
                    }
                    if ((v5 & 127) > (v5 >> ((686) & 63))) {
                        fmt.Println(v4, v6)
                        if ((((-2.25 >= 1e10) || true) || (!true)) || (v12 >= v12)) {
                            v18 := GenS2{F0: 42, F1: -1, F2: false}
                            if v4 {
                                switch v5 {
                                case -2147483648:
                                    for fi19 := 0; fi19 < 16; fi19++ {
                                        if ((!(0 == 42)) && (((float64(float32(1e10)) * 3.14159) / ((1e10 + ((float64(1e10) / 1.0) + 0.0)) / ((0.0 * ((((0.0 * 0.0) + 0.1) / 0.1) - 0.1)) + float64(-2.25)))) == 0.1)) {
                                            if v4 {
                                                if (((v6 % ((65535) | 1)) > 32767) && v4) {
                                                    _ = len(v2)
                                                    if (65535 != 42) {
                                                        for fi20 := 0; fi20 < 17; fi20++ {
                                                            v21 := GenS2{F0: 0, F1: 100, F2: true}
                                                            _ = fi20
                                                            _ = v21
                                                        }
                                                        v4 = v4
                                                        var v22 uint8 = (v1 * v1)
                                                        _ = v22
                                                    } else {
                                                        v2[-2147483648] = ("hello" + "")
                                                        for fi23 := 0; fi23 < 18; fi23++ {
                                                            v4 = false
                                                            if false {
                                                                var v24 uint32 = 0
                                                                _ = len(v2)
                                                                _ = v24
                                                            }
                                                            v6 += 32767
                                                            _ = fi23
                                                        }
                                                    }
                                                    if (v6 <= 2) {
                                                        v5 += 0
                                                        for fi25 := 0; fi25 < 6; fi25++ {
                                                            switch (v5 * -2147483648) {
                                                            case 128:
                                                                var v26 uint = 900
                                                                _ = v26
                                                            default:
                                                                v6 += v6
                                                            }
                                                            var v27 int16 = -602
                                                            _ = fi25
                                                            _ = v27
                                                        }
                                                        if (((v12 < (v12 << ((18446744073709551615) & 63))) && (int32(fi19) == -624)) && ((!v4) || (!v4))) {
                                                            if (-2.25 >= 0.1) {
                                                                v28 := []int{-9223372036854775808}
                                                                switch v5 {
                                                                case -975:
                                                                    v1 += 255
                                                                }
                                                                v29 := GenS2{F0: 42, F1: 77, F2: false}
                                                                _ = v28
                                                                _ = v29
                                                            }
                                                            v5 += 4294967295
                                                        } else {
                                                            var v30 float32 = 1.5
                                                            if (((-2.25 + 0.0) - ((3.14159 + float64((0.0 - float32(1.5)))) / (float64(float32(((0.0 - 3.14159) / 1e10))) + ((1e10 * (1e10 + 0.0)) / float64((0.0 + v30)))))) <= 1e10) {
                                                                var v31 float64 = float64(v30)
                                                                v31 = (1.5 - v31)
                                                                v32 := GenS2{F0: 42, F1: 34, F2: v4}
                                                                _ = v31
                                                                _ = v32
                                                            }
                                                            _ = v30
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        _ = fi19
                                    }
                                case -1:
                                    for fi33 := 0; fi33 < 9; fi33++ {
                                        for fi34 := 0; fi34 < 10; fi34++ {
                                            v35 := []int{689, (fi34 / ((fi34) | 1))}
                                            _ = len(v35)
                                            _ = fi34
                                            _ = v35
                                        }
                                        if (!false) {
                                            v36 := GenS2{F0: 42, F1: 1, F2: (1.5 >= (0.0 / (float32((1e10 + 1e10)) - 1.5)))}
                                            for fi37 := 0; fi37 < 10; fi37++ {
                                                v38 := GenS3{F0: (v1 | (v1 - 0)), F1: -43}
                                                v39 := []int{(fi33 * (fi37 << ((4294967295) & 63))), (fi37 % (((fi33 ^ int(v12))) | 1))}
                                                v40 := GenS4{F0: 32767, F1: 992}
                                                switch 2147483647 {
                                                case 128:
                                                    v5 += (fi33 | 4294967295)
                                                case -278:
                                                    fmt.Println(v4, v6)
                                                    v41 := v2[fi37]
                                                    _ = v41
                                                }
                                                var v42 uint8 = 137
                                                _ = fi37
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                                _ = v42
                                            }
                                            switch (fi33 - (^(fi33))) {
                                            case 952:
                                                var v43 bool = true
                                                _ = v43
                                            }
                                            _ = v36
                                        } else {
                                            var v44 int8 = 47
                                            for fi45 := 0; fi45 < 20; fi45++ {
                                                switch (fi33 << ((4294967295) & 63)) {
                                                case 127:
                                                    v46 := GenS1{F0: 4294967295}
                                                    _ = v46
                                                case 1:
                                                    v47 := GenS5{F0: v44, F1: "foo bar"}
                                                    switch "foo bar" {
                                                    case "hello":
                                                        switch "hello" {
                                                        case "hello":
                                                            if false {
                                                                var v48 float32 = 0.1
                                                                switch (fi45 & (-(fi33))) {
                                                                case 127:
                                                                    v49 := GenS6{F0: 42}
                                                                    _ = v49
                                                                case -957:
                                                                    v50 := GenS7{F0: 0, F1: ("a" + "a")}
                                                                    _ = v50
                                                                default:
                                                                    delete(v3, (-(v5)))
                                                                }
                                                                _ = v48
                                                            } else {
                                                                if (int(v12) <= (fi45 - 361)) {
                                                                    var v51 bool = false
                                                                    _ = v51
                                                                }
                                                            }
                                                        case "":
                                                            if (0.1 >= 1e10) {
                                                                if (v44 >= (-(v44))) {
                                                                    var v52 int16 = 32767
                                                                    if true {
                                                                        if false {
                                                                            fmt.Println(v52, v6)
                                                                            var v53 float32 = 3.14159
                                                                            _ = v53
                                                                        }
                                                                        var v54 int8 = (v44 ^ v44)
                                                                        _ = v54
                                                                    }
                                                                    var v55 int64 = (^(v12))
                                                                    _ = v52
                                                                    _ = v55
                                                                }
                                                                v56 := []int{32767}
                                                                _ = v56
                                                            }
                                                        case "a":
                                                            if (v4 || v4) {
                                                                v57 := GenS4{F0: 42, F1: 1}
                                                                _ = v57
                                                            }
                                                        }
                                                        v12 = (v12 | v12)
                                                        v58 := GenS4{F0: 0, F1: 2147483647}
                                                        _ = v58
                                                    case "\n\t":
                                                        v6 -= (v6 >> ((255) & 63))
                                                    case "foo bar":
                                                        v59 := v3[127]
                                                        _ = v59
                                                    }
                                                    _ = v47
                                                case -2147483648:
                                                    fmt.Println(v4, v6)
                                                }
                                                _ = fi45
                                            }
                                            for fi60 := 0; fi60 < 1; fi60++ {
                                                if v4 {
                                                    fmt.Println(v1, v12)
                                                    v61 := GenS8{F0: 128}
                                                    v62 := GenS9{F0: 30}
                                                    _ = v61
                                                    _ = v62
                                                } else {
                                                    var v63 float64 = (float64((3.14159 + float32((1.5 * 0.0)))) / 1e10)
                                                    switch ("a" + ("\n\t" + "\n\t")) {
                                                    case "a":
                                                        switch -9223372036854775808 {
                                                        case -650:
                                                            v4 = v4
                                                        default:
                                                            v63 = 3.14159
                                                        }
                                                    case "hello":
                                                        v12 = -173
                                                    default:
                                                        fmt.Println(v63, fi33, v4)
                                                    }
                                                    if v4 {
                                                        v64 := []int{(v5 / ((fi33) | 1)), 128}
                                                        switch v5 {
                                                        case 0:
                                                            fmt.Println(v5, v12)
                                                        }
                                                        v44 = int8(v1)
                                                        _ = v64
                                                    }
                                                    _ = v63
                                                }
                                                _ = fi60
                                            }
                                            _ = v44
                                        }
                                        _ = fi33
                                    }
                                default:
                                    var v65 uint64 = 2
                                    _ = v65
                                }
                                v6 = v6
                            } else {
                                switch "a" {
                                case "":
                                    v12 = (v12 ^ (v12 << ((4294967295) & 63)))
                                }
                                v66 := GenS4{F0: 1000, F1: 0}
                                if false {
                                    v12 = v12
                                    var v67 int16 = 148
                                    _ = v67
                                } else {
                                    v68 := []int{}
                                    _ = v68
                                }
                                _ = v66
                            }
                            _ = v18
                        }
                    }
                    switch ("foo bar" + (("a" + "\n\t") + "hello")) {
                    case "foo bar":
                        v69 := v3[v5]
                        _ = v69
                    case "":
                        v3[v5] = ((("\n\t" + ("foo bar" + ("a" + ("foo bar" + ((("foo bar" + ("\n\t" + "hello")) + ("foo bar" + "a")) + ((("hello" + "foo bar") + "foo bar") + "hello")))))) + "foo bar") + "foo bar")
                    default:
                        v70 := GenS5{F0: -112, F1: "foo bar"}
                        _ = v70
                    }
                }
            } else {
                if ((v5 << ((476) & 63)) != (v5 >> ((124) & 63))) {
                    v71 := []int{v5, v5}
                    _ = v71
                }
                v72 := GenS2{F0: 76, F1: -95, F2: (v4 && (("foo bar" < "a") || true))}
                _ = v72
            }
        }
        _ = v12
    }
    fmt.Println(v6, v6)
    if false {
        if (v4 && v4) {
            v4 = ("hello" > "a")
            var v73 int = v5
            _ = v73
        }
        v5 = (v5 - (v5 >> ((2) & 63)))
        for fi74 := 0; fi74 < 4; fi74++ {
            fmt.Println(v5, v5, v4)
            _ = fi74
        }
    }
    v75 := []int{v5, (v5 | v5), (v5 >> ((2) & 63))}
    v5 = v5
    var v76 uint16 = v6
    if v4 {
        switch "a" {
        case "foo bar":
            var v77 float64 = (1e10 / 1.0)
            _ = v77
        case "":
            v78 := GenS10{F0: -143, F1: 307, F2: 65535}
            _ = v78
        }
    } else {
        for fi79 := 0; fi79 < 16; fi79++ {
            for fi80 := 0; fi80 < 17; fi80++ {
                switch v5 {
                case 32768:
                    v81 := GenS9{F0: -128}
                    v5 += (fi79 % ((fi79) | 1))
                    _ = v81
                default:
                    if false {
                        fmt.Println(v4)
                    }
                }
                _ = fi80
            }
            _ = fi79
        }
        if ((1 <= 1) && v4) {
            switch v5 {
            case 42:
                v5 = 4294967295
            case 128:
                var v82 uint = 42
                _ = v82
            case 2147483647:
                v1 -= (v1 | v1)
            default:
                delete(v2, v5)
            }
            v1 += 128
            v3[v5] = "hello"
        }
        var v83 int64 = 0
        _ = v83
    }
    if v4 {
        for fi84 := 0; fi84 < 13; fi84++ {
            switch (("a" + "foo bar") + "\n\t") {
            case "":
                for fi85 := 0; fi85 < 12; fi85++ {
                    switch fi84 {
                    case 213:
                        fmt.Println(fi85, fi85, fi85)
                        var v86 uint = 65
                        _ = v86
                    case 2147483647:
                        if v4 {
                            if v4 {
                                v87 := GenS5{F0: 100, F1: ("a" + ("\n\t" + "hello"))}
                                var v88 bool = false
                                _ = v87
                                _ = v88
                            } else {
                                fmt.Println(v1)
                                v5 = v5
                            }
                        } else {
                            if len(v75) > 0 {
                                v75[0] = (fi85 ^ v5)
                            }
                        }
                    }
                    _ = fi85
                }
            case "foo bar":
                if v4 {
                    var v89 int32 = 248
                    _ = v89
                }
            case "a":
                delete(v3, (v5 & -519))
            default:
                for fi90 := 0; fi90 < 16; fi90++ {
                    var v91 uint = 962
                    var v92 int8 = -128
                    _ = fi90
                    _ = v91
                    _ = v92
                }
            }
            if (-1 > v5) {
                switch (v5 >> ((0) & 63)) {
                case 127:
                    v93 := GenS7{F0: 42, F1: ""}
                    _ = v93
                case 4294967295:
                    v94 := v2[127]
                    _ = v94
                }
                v1 = (^(v1))
                _ = len(v2)
            } else {
                var v95 bool = (true && (!(2 != 1)))
                v6 += v76
                _ = v95
            }
            var v96 uint32 = 42
            _ = fi84
            _ = v96
        }
        fmt.Println(v5, v5, v4)
    }
    var v97 float32 = 0.1
    v76 = 25
    if (42 == 4294967295) {
        v98 := GenS10{F0: 42, F1: (v76 + (v76 * 65535)), F2: 4294967295}
        switch ("hello" + "hello") {
        case "":
            _ = len(v3)
        default:
            var v99 float32 = v97
            _ = v99
        }
        _ = v98
    } else {
        if ((!(v4 && (v4 && true))) && (1 <= 986)) {
            v100 := v2[v5]
            _ = v100
        } else {
            v5 = 982
            v101 := GenS3{F0: v1, F1: -1}
            _ = v101
        }
        v102 := GenS7{F0: 2147483647, F1: "foo bar"}
        v2[(^(v5))] = ""
        _ = v102
    }
    for fi103 := 0; fi103 < 11; fi103++ {
        v5 -= (v5 * fi103)
        _ = fi103
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v75:", len(v75), v75)
    fmt.Println("v76:", v76)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v75
    _ = v76
    _ = v8
    _ = v97
}
