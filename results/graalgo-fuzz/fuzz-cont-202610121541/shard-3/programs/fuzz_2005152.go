package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS6 struct { F0 int; F1 int64; F2 int8 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 int64; F1 uint16; F2 int64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint32; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int = -2147483648
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    var v5 int8 = -75
    var v6 int16 = 1000
    v5 = 1
    var v7 string = "hello"
    switch v7 {
    case "a":
        v8 := GenS1{F0: (v5 << ((0) & 63))}
        _ = v8
    case "":
        switch (((((v7 + (("hello" + "foo bar") + "hello")) + "a") + v7) + ("" + v7)) + v7) {
        case "foo bar":
            if (1.5 == 0.0) {
                for fi9 := 0; fi9 < 6; fi9++ {
                    v10 := []int{(^(fi9)), 1, 41}
                    _ = fi9
                    _ = v10
                }
                if true {
                    var v11 int16 = (v6 & (v6 & 32767))
                    var v12 string = (v7 + (v7 + (("hello" + "foo bar") + v7)))
                    if true {
                        var v13 int16 = v11
                        fmt.Println(v7, v6)
                        _ = v13
                    } else {
                        switch v12 {
                        case "hello":
                            if false {
                                if (true && ((-2.25 / 3.14159) >= -2.25)) {
                                    fmt.Println(v7, v5, v7)
                                    if (-2.25 < (0.1 - ((-2.25 + 1e10) - -2.25))) {
                                        v5 = v5
                                        v14 := GenS2{F0: v12, F1: (((v12 + v12) + ((("a" + ((v12 + ("hello" + "\n\t")) + v7)) + "") + ("foo bar" + v12))) != v7)}
                                        fmt.Println(v6, v5)
                                        _ = v14
                                    } else {
                                        v15 := []int{v2}
                                        v16 := GenS2{F0: v7, F1: true}
                                        _ = v15
                                        _ = v16
                                    }
                                } else {
                                    v17 := GenS2{F0: "\n\t", F1: false}
                                    _ = v17
                                }
                                switch "\n\t" {
                                case "\n\t":
                                    v18 := GenS2{F0: v7, F1: true}
                                    _ = v18
                                case "hello":
                                    v11 = 32767
                                case "foo bar":
                                    v7 = v12
                                }
                            }
                        case "":
                            var v19 float32 = 3.14159
                            _ = v19
                        default:
                            var v20 bool = true
                            _ = v20
                        }
                        var v21 float64 = (1e10 + 0.1)
                        _ = v21
                    }
                    _ = v11
                    _ = v12
                } else {
                    v7 = ((v7 + v7) + (v7 + "a"))
                    if ((v6 * 32767) == (v6 / (((v6 % (((v6 + 42)) | 1))) | 1))) {
                        if false {
                            for fi22 := 0; fi22 < 12; fi22++ {
                                if true {
                                    if true {
                                        fmt.Println(fi22, v5, v7)
                                        for fi23 := 0; fi23 < 7; fi23++ {
                                            if false {
                                                fmt.Println(v6, fi23, v7)
                                            } else {
                                                var v24 int = 366
                                                var v25 bool = (!(false || true))
                                                if (true || (757 == 1)) {
                                                    switch (v2 >> ((65535) & 63)) {
                                                    case 42:
                                                        v26 := GenS3{F0: v5}
                                                        _ = v26
                                                    case -407:
                                                        v2 = fi22
                                                    case 9223372036854775807:
                                                        var v27 float64 = 0.0
                                                        _ = v27
                                                    default:
                                                        fmt.Println(v7, v25)
                                                    }
                                                    v28 := GenS2{F0: ("hello" + v7), F1: v25}
                                                    v29 := GenS4{F0: 596}
                                                    _ = v28
                                                    _ = v29
                                                } else {
                                                    var v30 uint64 = 42
                                                    _ = v30
                                                }
                                                fmt.Println(v25, fi22)
                                                var v31 uint16 = 255
                                                v32 := []int{-584, -2147483648, fi23}
                                                _ = v24
                                                _ = v25
                                                _ = v31
                                                _ = v32
                                            }
                                            _ = fi23
                                        }
                                    }
                                    fmt.Println(v7, v7, v7)
                                    for fi33 := 0; fi33 < 4; fi33++ {
                                        v7 = "foo bar"
                                        _ = fi33
                                    }
                                } else {
                                    if true {
                                        if (!false) {
                                            for fi34 := 0; fi34 < 20; fi34++ {
                                                v6 = (v6 >> ((101) & 63))
                                                fmt.Println(fi22, v7, fi34)
                                                switch -93 {
                                                case 42:
                                                    if ((!true) && false) {
                                                        v35 := GenS1{F0: (^(v5))}
                                                        var v36 string = (v7 + v7)
                                                        _ = v35
                                                        _ = v36
                                                    }
                                                    fmt.Println(v6, v2)
                                                case 128:
                                                    switch ((v7 + "a") + (v7 + ("" + (v7 + v7)))) {
                                                    case "a":
                                                        v37 := GenS4{F0: 2}
                                                        _ = v37
                                                    case "":
                                                        if false {
                                                            v2 -= 1
                                                            v38 := []int{(v2 + (fi22 >> ((255) & 63))), (fi22 / ((fi22) | 1)), fi34, v2}
                                                            var v39 int = v2
                                                            _ = v38
                                                            _ = v39
                                                        }
                                                    default:
                                                        switch ("hello" + "\n\t") {
                                                        case "\n\t":
                                                            if (false && ((!(128 == 0)) && (!true))) {
                                                                v5 += (v5 >> ((65535) & 63))
                                                                if false {
                                                                    v5 = -128
                                                                    v40 := GenS5{F0: 2147483647, F1: -1, F2: 0}
                                                                    v7 = "a"
                                                                    var v41 string = "\n\t"
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                                v2 += (fi22 >> ((4294967295) & 63))
                                                            } else {
                                                                if (true && false) {
                                                                    v42 := []int{fi34}
                                                                    var v43 uint = 765
                                                                    _ = v42
                                                                    _ = v43
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if true {
                                                        v2 -= -533
                                                    } else {
                                                        fmt.Println(fi22, v6, v2)
                                                    }
                                                    v44 := []int{}
                                                    _ = v44
                                                case -9223372036854775808:
                                                    switch (v7 + v7) {
                                                    case "a":
                                                        if true {
                                                            v45 := GenS5{F0: 267, F1: 9223372036854775807, F2: 1}
                                                            switch "a" {
                                                            case "hello":
                                                                fmt.Println(v5, fi22)
                                                            default:
                                                                var v46 uint8 = 255
                                                                _ = v46
                                                            }
                                                            fmt.Println(fi34, v5)
                                                            _ = v45
                                                        }
                                                    }
                                                }
                                                _ = fi34
                                            }
                                            switch v7 {
                                            case "foo bar":
                                                v47 := GenS1{F0: 100}
                                                _ = v47
                                            case "\n\t":
                                                for fi48 := 0; fi48 < 10; fi48++ {
                                                    var v49 uint = 2
                                                    v50 := GenS2{F0: (v7 + "\n\t"), F1: (false || (!(35 > 42)))}
                                                    if ((((0.1 / 1.5) != 3.14159) || false) || true) {
                                                        var v51 int8 = 127
                                                        _ = v51
                                                    }
                                                    _ = fi48
                                                    _ = v49
                                                    _ = v50
                                                }
                                            default:
                                                if true {
                                                    v52 := GenS6{F0: (v2 | -1), F1: 0, F2: v5}
                                                    fmt.Println(v6, fi22, v2)
                                                    var v53 int64 = 1
                                                    _ = v52
                                                    _ = v53
                                                }
                                            }
                                            switch (v7 + v7) {
                                            case "hello":
                                                delete(v4, v7)
                                            case "":
                                                v7 = v7
                                            default:
                                                fmt.Println(v5)
                                            }
                                        }
                                    }
                                    fmt.Println(v2, v5)
                                }
                                switch -1 {
                                case 686:
                                    var v54 float64 = ((0.0 * ((0.0 * -2.25) + ((-2.25 + (float64((1.5 * 1e10)) + -2.25)) - float64(0.1)))) + ((1e10 - 0.0) + float64((float32(0.1) - 1.5))))
                                    _ = v54
                                case 9223372036854775807:
                                    for fi55 := 0; fi55 < 18; fi55++ {
                                        v56 := GenS4{F0: 19}
                                        var v57 string = v7
                                        fmt.Println(fi22, fi22, v2)
                                        _ = fi55
                                        _ = v56
                                        _ = v57
                                    }
                                case 32768:
                                    var v58 string = ("\n\t" + "\n\t")
                                    _ = v58
                                default:
                                    v59 := GenS4{F0: 4294967295}
                                    _ = v59
                                }
                                switch -1 {
                                case -371:
                                    if (true || true) {
                                        v60 := GenS7{F0: (!(!(1 != 1))), F1: 65535}
                                        v6 = (-(v6))
                                        if false {
                                            v61 := GenS3{F0: (v5 % ((v5) | 1))}
                                            v5 = v5
                                            _ = v61
                                        }
                                        _ = v60
                                    } else {
                                        v62 := GenS4{F0: 0}
                                        for fi63 := 0; fi63 < 18; fi63++ {
                                            var v64 int8 = v5
                                            _ = fi63
                                            _ = v64
                                        }
                                        v1[0] = v2
                                        _ = v62
                                    }
                                }
                                _ = fi22
                            }
                            v65 := GenS4{F0: 2}
                            var v66 uint8 = 2
                            _ = v65
                            _ = v66
                        }
                    } else {
                        if false {
                            fmt.Println(v2, v5)
                        } else {
                            for fi67 := 0; fi67 < 4; fi67++ {
                                fmt.Println(v5, v2)
                                _ = fi67
                            }
                            if true {
                                v5 = -55
                                if (!(!false)) {
                                    var v68 int16 = v6
                                    for fi69 := 0; fi69 < 6; fi69++ {
                                        fmt.Println(v7, v2)
                                        v70 := []int{247, 32768, (v2 & v2), 435}
                                        v71 := GenS3{F0: -128}
                                        _ = fi69
                                        _ = v70
                                        _ = v71
                                    }
                                    var v72 float64 = -2.25
                                    _ = v68
                                    _ = v72
                                } else {
                                    v7 = (v7 + (((("" + v7) + (v7 + (v7 + v7))) + v7) + "a"))
                                    v6 -= (v6 & (v6 ^ (v6 - (-(v6)))))
                                    v73 := GenS3{F0: int8(v2)}
                                    _ = v73
                                }
                            } else {
                                for fi74 := 0; fi74 < 17; fi74++ {
                                    v75 := GenS8{F0: false}
                                    _ = fi74
                                    _ = v75
                                }
                            }
                        }
                    }
                }
            }
        case "a":
            v6 = (v6 ^ 32767)
        case "hello":
            for fi76 := 0; fi76 < 7; fi76++ {
                v3[4294967295] = ((v7 + v7) + v7)
                var v77 bool = true
                _ = fi76
                _ = v77
            }
        }
    case "\n\t":
        v2 -= v2
    }
    for fi78 := 0; fi78 < 3; fi78++ {
        if (-791 < 0) {
            v2 = fi78
        } else {
            var v79 uint16 = 65535
            v80 := GenS9{F0: 9223372036854775807, F1: (v79 >> ((0) & 63)), F2: 1}
            _ = v79
            _ = v80
        }
        _ = fi78
    }
    delete(v1, 661)
    v81 := []int{v2, v2}
    var v82 uint32 = 65535
    v5 = -43
    if true {
        switch "hello" {
        case "foo bar":
            v83 := GenS9{F0: -9223372036854775808, F1: 709, F2: -291}
            _ = v83
        case "a":
            var v84 int32 = -2147483648
            _ = v84
        case "":
            if (502 >= (^(v82))) {
                for fi85 := 0; fi85 < 9; fi85++ {
                    var v86 string = (v7 + "hello")
                    _ = fi85
                    _ = v86
                }
                switch "\n\t" {
                case "foo bar":
                    v87 := GenS5{F0: v82, F1: 0, F2: 2}
                    _ = v87
                case "hello":
                    v82 = v82
                case "a":
                    if true {
                        for fi88 := 0; fi88 < 19; fi88++ {
                            var v89 string = v7
                            v90 := GenS10{F0: (true && true)}
                            var v91 bool = (1e10 < ((1e10 * 0.0) / (-2.25 - ((float32(0.0) + -2.25) * 0.1))))
                            _ = fi88
                            _ = v89
                            _ = v90
                            _ = v91
                        }
                    }
                default:
                    v82 += 2147483648
                }
                var v92 uint8 = 42
                _ = v92
            }
        }
        var v93 uint16 = uint16(v82)
        v5 -= v5
        _ = v93
    } else {
        v5 += (v5 * v5)
        v6 -= 1
        if true {
            fmt.Println(v82, v82, v6)
            switch (v2 << ((18446744073709551615) & 63)) {
            case 1:
                if (1e10 == 1.5) {
                    var v94 int64 = 0
                    _ = v94
                } else {
                    var v95 int = v2
                    var v96 bool = (9223372036854775807 == 0)
                    _ = v95
                    _ = v96
                }
            default:
                fmt.Println(v5)
            }
            switch 484 {
            case -2147483648:
                fmt.Println(v7)
            }
        }
    }
    fmt.Println(v7, v6, v7)
    fmt.Println(v2, v6, v7)
    v7 = ("a" + v7)
    var v97 float64 = 1.5
    _ = len(v3)
    for fi98 := 0; fi98 < 5; fi98++ {
        v2 = fi98
        switch (fi98 << ((42) & 63)) {
        case 608:
            _ = len(v81)
        }
        for fi99 := 0; fi99 < 19; fi99++ {
            var v100 int16 = v6
            if (true || (v7 <= v7)) {
                var v101 uint8 = 127
                _ = v101
            }
            _ = fi99
            _ = v100
        }
        _ = fi98
    }
    var v102 int64 = 0
    for fi103 := 0; fi103 < 19; fi103++ {
        v7 = (v7 + "foo bar")
        switch 32768 {
        case -988:
            _ = len(v3)
        case 127:
            for fi104 := 0; fi104 < 16; fi104++ {
                var v105 uint64 = 439
                var v106 uint16 = 736
                switch v2 {
                case 4294967295:
                    v82 = (^(v82))
                default:
                    switch (v7 + v7) {
                    case "hello":
                        v2 = 423
                    case "a":
                        var v107 int = (fi104 % ((-9223372036854775808) | 1))
                        _ = v107
                    }
                }
                _ = fi104
                _ = v105
                _ = v106
            }
        case -1:
            v108 := GenS11{F0: v82, F1: 757}
            _ = v108
        }
        var v109 int64 = v102
        _ = fi103
        _ = v109
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v102:", v102)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v81:", len(v81), v81)
    fmt.Println("v82:", v82)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v81
    _ = v82
    _ = v97
}
