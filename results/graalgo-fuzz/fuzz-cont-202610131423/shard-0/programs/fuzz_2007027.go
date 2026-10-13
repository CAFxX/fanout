package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int; F2 uint16 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS5 struct { F0 uint8; F1 string }
type GenS6 struct { F0 int8; F1 bool }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 int32; F1 int8; F2 uint32 }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 int16 }
type GenS12 struct { F0 string; F1 bool }
type GenS13 struct { F0 int64 }
type GenS14 struct { F0 int32; F1 string }
func genFn0(p0 int32) uint16 {
    for fi5 := 0; fi5 < 2; fi5++ {
        if true {
            var v6 int = (fi5 + (fi5 + fi5))
            v7 := GenS1{F0: 42}
            _ = v6
            _ = v7
        }
        v8 := GenS2{F0: -874}
        _ = fi5
        _ = v8
    }
    v9 := []int{}
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    _ = p0
    _ = v9
    return 359
}
func genFn1() uint64 {
    for fi10 := 0; fi10 < 20; fi10++ {
        fmt.Println(fi10)
        var v11 uint64 = 895
        v11 = (v11 - 948)
        _ = fi10
        _ = v11
    }
    switch "a" {
    case "":
        var v12 uint16 = 42
        _ = v12
    case "hello":
        var v13 string = ("foo bar" + "hello")
        _ = v13
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        if true {
            var v15 uint = 255
            if (!(false && (!(1 <= -1)))) {
                var v16 string = "foo bar"
                switch (fi14 - 127) {
                case -491:
                    for fi17 := 0; fi17 < 12; fi17++ {
                        var v18 uint16 = 216
                        _ = fi17
                        _ = v18
                    }
                case 127:
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v16 = v16
                        _ = fi19
                    }
                case 2147483647:
                    if false {
                        v20 := []int{}
                        _ = v20
                    } else {
                        for fi21 := 0; fi21 < 14; fi21++ {
                            v22 := GenS3{F0: (v15 & v15), F1: (^(fi14)), F2: 1}
                            v15 = 65535
                            _ = fi21
                            _ = v22
                        }
                        fmt.Println(v16, v16)
                        for fi23 := 0; fi23 < 11; fi23++ {
                            fmt.Println(v16)
                            var v24 int32 = 2147483647
                            _ = fi23
                            _ = v24
                        }
                    }
                default:
                    if (v16 <= "hello") {
                        var v25 uint16 = genFn0(2147483647)
                        _ = v25
                    }
                }
                v16 = (v16 + "foo bar")
                _ = v16
            }
            _ = v15
        }
        var v26 uint64 = 18446744073709551615
        if false {
            for fi27 := 0; fi27 < 17; fi27++ {
                if (!false) {
                    if true {
                        fmt.Println(v26)
                        v26 = 42
                        switch "\n\t" {
                        case "":
                            if ((906 != 0) && false) {
                                v28 := GenS4{F0: 255, F1: (^(v26)), F2: 0}
                                _ = v28
                            } else {
                                var v29 uint = 0
                                _ = v29
                            }
                        case "a":
                            v26 -= (v26 / ((4294967295) | 1))
                        case "\n\t":
                            var v30 int64 = 42
                            _ = v30
                        default:
                            var v31 uint8 = 2
                            _ = v31
                        }
                    } else {
                        var v32 int = -903
                        switch int(v26) {
                        case 9223372036854775807:
                            if false {
                                var v33 uint32 = 4294967295
                                if false {
                                    v34 := []int{(fi27 - fi27)}
                                    switch fi14 {
                                    case -650:
                                        if true {
                                            if len(v34) >= 2 {
                                                v34 = v34[1:len(v34)]
                                            }
                                            v32 -= fi27
                                            v26 = (^(v26))
                                        }
                                    case -2147483648:
                                        if (true && false) {
                                            if true {
                                                var v35 uint32 = 662
                                                _ = v35
                                            }
                                        } else {
                                            v32 += 9223372036854775807
                                            if false {
                                                var v36 uint = 18446744073709551615
                                                _ = v36
                                            } else {
                                                if (false && (32767 > 0)) {
                                                    switch ("hello" + ((("foo bar" + "a") + ("foo bar" + "\n\t")) + "hello")) {
                                                    case "\n\t":
                                                        if len(v34) >= 2 {
                                                            v34 = v34[1:len(v34)]
                                                        }
                                                    default:
                                                        var v37 int16 = 1
                                                        _ = v37
                                                    }
                                                } else {
                                                    _ = len(v34)
                                                }
                                                v33 = v33
                                                if true {
                                                    v32 += (^(fi27))
                                                    switch "foo bar" {
                                                    case "hello":
                                                        fmt.Println(v32, v26, v32)
                                                    case "":
                                                        if true {
                                                            if false {
                                                                v38 := GenS2{F0: 9223372036854775807}
                                                                v39 := GenS3{F0: 65535, F1: (fi27 >> ((2) & 63)), F2: 2}
                                                                var v40 string = "a"
                                                                switch v40 {
                                                                case "foo bar":
                                                                    _ = len(v34)
                                                                case "\n\t":
                                                                    if true {
                                                                        v40 = (v40 + v40)
                                                                        v34 = append(v34, 32768)
                                                                    }
                                                                case "":
                                                                    v41 := GenS2{F0: 345}
                                                                    _ = v41
                                                                }
                                                                if true {
                                                                    fmt.Println(v33, fi14)
                                                                }
                                                                _ = v38
                                                                _ = v39
                                                                _ = v40
                                                            }
                                                        }
                                                    }
                                                    v42 := GenS3{F0: 65535, F1: fi14, F2: 255}
                                                    _ = v42
                                                } else {
                                                    v43 := GenS5{F0: 128, F1: "foo bar"}
                                                    if (384 >= (^(v33))) {
                                                        v26 += (v26 + v26)
                                                        var v44 string = ("a" + "")
                                                        _ = v44
                                                    } else {
                                                        switch (^(fi27)) {
                                                        case 4294967295:
                                                            if true {
                                                                if len(v34) > 0 {
                                                                    v34[0] = (v32 * (-(v32)))
                                                                }
                                                                var v45 float64 = 3.14159
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    v45 += v45
                                                                case "hello":
                                                                    var v46 bool = true
                                                                    _ = v46
                                                                }
                                                                _ = v45
                                                            } else {
                                                                v32 -= (fi27 * (fi14 ^ fi14))
                                                                switch 32768 {
                                                                case 32768:
                                                                    v47 := GenS6{F0: -70, F1: ("a" > "\n\t")}
                                                                    _ = v47
                                                                default:
                                                                    v32 += (fi14 >> ((278) & 63))
                                                                }
                                                                if len(v34) >= 2 {
                                                                    v34 = v34[1:len(v34)]
                                                                }
                                                            }
                                                        case 128:
                                                            var v48 uint = 563
                                                            _ = v48
                                                        default:
                                                            _ = len(v34)
                                                        }
                                                        v26 += (v26 + (v26 / ((v26) | 1)))
                                                    }
                                                    _ = v43
                                                }
                                            }
                                            var v49 int32 = 2147483647
                                            _ = v49
                                        }
                                    }
                                    _ = v34
                                }
                                var v50 int32 = 990
                                _ = v33
                                _ = v50
                            } else {
                                var v51 float64 = ((1e10 / 0.1) / (1.5 - 0.1))
                                v51 = 3.14159
                                v52 := GenS5{F0: 127, F1: "hello"}
                                var v53 float64 = ((v51 - (v51 * (-2.25 / (v51 - (float64(0.1) * 3.14159))))) - v51)
                                fmt.Println(v32)
                                switch (("" + "foo bar") + "") {
                                case "a":
                                    if (true || (!(1.5 <= (v51 / (0.0 / ((float64((3.14159 - (1e10 / 1.0))) * (v53 / v51)) * v51)))))) {
                                        switch "" {
                                        case "foo bar":
                                            var v54 uint16 = 0
                                            v55 := GenS2{F0: 9223372036854775807}
                                            v54 -= (v54 ^ (v54 / ((v54) | 1)))
                                            _ = v54
                                            _ = v55
                                        case "\n\t":
                                            v56 := GenS7{F0: ("foo bar" + "\n\t")}
                                            _ = v56
                                        }
                                        v51 = v51
                                    } else {
                                        v57 := GenS2{F0: -283}
                                        _ = v57
                                    }
                                default:
                                    v53 += float64(0.1)
                                }
                                v58 := []int{(fi27 - fi14), v32}
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v58
                            }
                        }
                        v59 := GenS7{F0: "\n\t"}
                        _ = v32
                        _ = v59
                    }
                    switch fi27 {
                    case 32768:
                        if (((float64(((0.0 * (1.5 / 1e10)) - ((1.5 / 3.14159) / 1e10))) / 1.0) - -2.25) != (1.5 / 0.1)) {
                            if (!(0 >= 732)) {
                                v60 := []int{fi27, 32768}
                                v26 = 42
                                _ = v60
                            }
                            var v61 float64 = (0.0 / 1.0)
                            _ = v61
                        }
                    case -901:
                        v62 := GenS8{F0: 42}
                        _ = v62
                    case 2147483647:
                        if (false && (((v26 << ((364) & 63)) != v26) || true)) {
                            var v63 uint64 = 854
                            if (true && false) {
                                var v64 uint16 = 1
                                v65 := GenS9{F0: 0, F1: 100, F2: 2}
                                v63 = v26
                                switch fi27 {
                                case -795:
                                    var v66 string = "hello"
                                    _ = v66
                                default:
                                    v67 := GenS10{F0: 1}
                                    _ = v67
                                }
                                var v68 int = (fi27 / ((596) | 1))
                                _ = v64
                                _ = v65
                                _ = v68
                            }
                            _ = v63
                        } else {
                            fmt.Println(fi27)
                            var v69 int32 = -2147483648
                            _ = v69
                        }
                    default:
                        var v70 float32 = ((1e10 - 0.0) / 1e10)
                        _ = v70
                    }
                    v71 := GenS11{F0: -32768}
                    _ = v71
                } else {
                    v72 := []int{}
                    v73 := GenS12{F0: "hello", F1: true}
                    var v74 int16 = 42
                    _ = v72
                    _ = v73
                    _ = v74
                }
                _ = fi27
            }
            var v75 bool = false
            v76 := []int{fi14, (fi14 | (fi14 / (((fi14 - fi14)) | 1)))}
            _ = v75
            _ = v76
        }
        _ = fi14
        _ = v26
    }
    v77 := []int{-604}
    var v78 float32 = (((float32(1e10) / float32(0.1)) * float32(((0.0 - float64(3.14159)) / float64(((0.1 + -2.25) - float32(1e10)))))) + -2.25)
    _ = v77
    _ = v78
    return 0
}

func main() {
    var v1 bool = true
    var v2 bool = v1
    v3 := []uint32{0, 2147483648, 761}
    var v4 uint16 = 255
    if false {
        v79 := GenS5{F0: 128, F1: ("a" + "foo bar")}
        _ = v79
    } else {
        for fi80 := 0; fi80 < 20; fi80++ {
            var v81 int64 = 1
            _ = fi80
            _ = v81
        }
        switch ("a" + "\n\t") {
        case "\n\t":
            v82 := GenS11{F0: 1}
            _ = v82
        case "foo bar":
            v1 = false
        default:
            for fi83 := 0; fi83 < 20; fi83++ {
                fmt.Println(v1)
                _ = len(v3)
                var v84 float32 = float32((1.5 - 1.5))
                _ = fi83
                _ = v84
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    if (0 >= 608) {
        for fi85 := 0; fi85 < 3; fi85++ {
            v4 -= 492
            v2 = false
            var v86 uint32 = 214
            _ = fi85
            _ = v86
        }
        v4 += (^(v4))
        v2 = false
    } else {
        var v87 string = "foo bar"
        _ = v87
    }
    v1 = (748 > 2)
    for fi88 := 0; fi88 < 3; fi88++ {
        var v89 float32 = 1e10
        v4 += v4
        _ = fi88
        _ = v89
    }
    var v90 uint16 = genFn0(202)
    switch -9223372036854775808 {
    case -40:
        v91 := GenS13{F0: -9223372036854775808}
        _ = v91
    }
    var v92 float32 = 0.1
    switch -847 {
    case 2147483647:
        if len(v3) > 0 {
            v3[0] = 4294967295
        }
    case 127:
        _ = len(v3)
    default:
        fmt.Println(v1, v2, v4)
    }
    var v93 int = 1
    v94 := GenS1{F0: (v4 & v4)}
    var v95 uint64 = 315
    if v2 {
        for fi96 := 0; fi96 < 18; fi96++ {
            fmt.Println(v95, v92)
            var v97 int64 = 0
            _ = fi96
            _ = v97
        }
        if (42 <= -128) {
            fmt.Println(v92, v92, v4)
            var v98 uint32 = 42
            _ = v98
        } else {
            var v99 uint = 2
            switch "foo bar" {
            case "hello":
                if false {
                    for fi100 := 0; fi100 < 8; fi100++ {
                        var v101 int16 = 917
                        for fi102 := 0; fi102 < 10; fi102++ {
                            var v103 float64 = 3.14159
                            if len(v3) > 0 {
                                v3[0] = 65535
                            }
                            fmt.Println(v103, v99, v90)
                            _ = fi102
                            _ = v103
                        }
                        for fi104 := 0; fi104 < 12; fi104++ {
                            if (!true) {
                                fmt.Println(fi104, v1, v2)
                            } else {
                                switch "\n\t" {
                                case "\n\t":
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                }
                                v105 := GenS14{F0: -1, F1: ("foo bar" + (("\n\t" + "hello") + "\n\t"))}
                                _ = v105
                            }
                            _ = fi104
                        }
                        _ = fi100
                        _ = v101
                    }
                } else {
                    _ = len(v3)
                }
            case "foo bar":
                if (42 >= 0) {
                    v90 -= (v90 & (v90 & (^(v4))))
                }
            case "a":
                if (!("" < "\n\t")) {
                    var v106 uint64 = v95
                    _ = v106
                } else {
                    if v1 {
                        var v107 uint = uint(v95)
                        _ = v107
                    } else {
                        if (!("foo bar" <= "\n\t")) {
                            v3 = append(v3, 684)
                            fmt.Println(v99, v99, v95)
                            v108 := GenS3{F0: (v99 | (v99 - (v99 ^ 438))), F1: (v93 | 0), F2: (v4 | (v4 - genFn0(2147483647)))}
                            _ = v108
                        }
                    }
                }
            }
            _ = v99
        }
        var v109 string = "foo bar"
        _ = v109
    } else {
        v95 -= (v95 | 4294967295)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v90:", v90)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v90
    _ = v92
    _ = v93
    _ = v94
    _ = v95
}
