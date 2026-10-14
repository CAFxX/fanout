package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int; F1 int8; F2 int64 }
type GenS4 struct { F0 int16; F1 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint; F1 int64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 int64; F1 int64; F2 uint16 }
func genFn0() int32 {
    v5 := []int{-424, 32767, -28, -9223372036854775808}
    var v6 uint8 = 255
    var v7 float64 = float64(1e10)
    if (4294967295 > 327) {
        fmt.Println(v7, v6)
    }
    if true {
        v5 = append(v5, 42)
        if false {
            if len(v5) > 0 {
                v5[0] = 127
            }
        } else {
            if (4294967295 < 18446744073709551615) {
                for fi8 := 0; fi8 < 9; fi8++ {
                    var v9 uint = 65535
                    _ = fi8
                    _ = v9
                }
            } else {
                fmt.Println(v6)
                var v10 int64 = 0
                _ = v10
            }
        }
        v11 := GenS1{F0: 641, F1: 42}
        _ = v11
    }
    _ = v5
    _ = v6
    _ = v7
    return -2147483648
}
func genFn1(p0 bool) uint16 {
    var v12 uint = 1
    var v13 int16 = -32768
    var v14 uint = 255
    v15 := []int{-383, 4294967295}
    if p0 {
        var v16 string = (("" + ("\n\t" + (("\n\t" + "") + (("\n\t" + (("a" + ("foo bar" + "\n\t")) + "a")) + ("\n\t" + "\n\t"))))) + "")
        var v17 int16 = v13
        if false {
            v18 := GenS1{F0: (v12 % ((0) | 1)), F1: 42}
            _ = v18
        } else {
            var v19 int = 0
            var v20 int64 = 9223372036854775807
            _ = v19
            _ = v20
        }
        _ = v16
        _ = v17
    } else {
        switch -1 {
        case -732:
            var v21 uint8 = 2
            _ = v21
        }
        p0 = ((255 == 227) || ((v13 + v13) != (v13 | (v13 % ((v13) | 1)))))
        var v22 int = 32767
        _ = v22
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    return 2
}

func main() {
    var v1 bool = (((4294967295 > 1) || (false && false)) && (true || false))
    var v2 uint16 = 65535
    var v3 uint64 = uint64(v2)
    var v4 uint = 255
    v4 = (v4 | v4)
    if (0.1 == 0.0) {
        v3 -= v3
    } else {
        var v23 int32 = 785
        v2 = (v2 % (((v2 % (((v2 << ((0) & 63))) | 1))) | 1))
        _ = v23
    }
    fmt.Println(v4, v4)
    if true {
        v24 := GenS1{F0: (v4 | 4294967295), F1: (v3 ^ (^(v3)))}
        v25 := GenS1{F0: (v4 + v4), F1: 0}
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: 1, F1: -9223372036854775808}
    for fi27 := 0; fi27 < 11; fi27++ {
        v28 := GenS3{F0: 73, F1: 100, F2: -1}
        var v29 float64 = (1e10 * (-2.25 + 0.1))
        switch "" {
        case "a":
            v3 += (v3 % ((v3) | 1))
        case "hello":
            if (!(1e10 < 1.5)) {
                v29 += v29
                for fi30 := 0; fi30 < 11; fi30++ {
                    var v31 int = 7
                    if (!v1) {
                        v3 = (v3 * 512)
                    } else {
                        v4 = (v4 >> ((v4) & 63))
                        v32 := []int{-9223372036854775808, 32768, 127}
                        switch 9223372036854775807 {
                        case 518:
                            if ((0.1 / 1.5) == 0.1) {
                                var v33 int32 = -428
                                _ = v33
                            }
                        case 32767:
                            if (fi30 < (fi27 * fi30)) {
                                v2 += 0
                                switch "a" {
                                case "hello":
                                    v4 = 953
                                }
                                if ((^(v3)) >= v3) {
                                    v34 := GenS4{F0: -32768, F1: (42 == 0)}
                                    switch fi27 {
                                    case 2147483647:
                                        switch (fi30 + 128) {
                                        case 2147483647:
                                            switch (fi27 & (fi30 * v31)) {
                                            case -351:
                                                fmt.Println(v1)
                                            case 32768:
                                                var v35 int32 = 553
                                                _ = v35
                                            case -460:
                                                if v1 {
                                                    switch ("a" + "a") {
                                                    case "\n\t":
                                                        v2 += (v2 & v2)
                                                        switch "foo bar" {
                                                        case "\n\t":
                                                            v36 := GenS5{F0: 42}
                                                            _ = v36
                                                        default:
                                                            var v37 uint = v4
                                                            switch ("a" + "\n\t") {
                                                            case "":
                                                                v2 += (v2 / ((genFn1(true)) | 1))
                                                            case "foo bar":
                                                                v3 += v3
                                                            }
                                                            v2 += (v2 ^ v2)
                                                            _ = v37
                                                        }
                                                    case "foo bar":
                                                        if (v1 && v1) {
                                                            v38 := GenS3{F0: (-(fi30)), F1: 127, F2: 691}
                                                            v39 := GenS6{F0: -40}
                                                            _ = v38
                                                            _ = v39
                                                        } else {
                                                            v1 = (-9223372036854775808 >= 1)
                                                            v40 := GenS4{F0: -1, F1: (556 <= 4294967295)}
                                                            _ = v40
                                                        }
                                                    }
                                                    switch 4294967295 {
                                                    case 32768:
                                                        fmt.Println(v4)
                                                    case -680:
                                                        v32 = append(v32, fi30)
                                                    default:
                                                        fmt.Println(fi27)
                                                    }
                                                }
                                            default:
                                                var v41 uint8 = 2
                                                _ = v41
                                            }
                                        default:
                                            switch "a" {
                                            case "a":
                                                if len(v32) >= 2 {
                                                    v32 = v32[1:len(v32)]
                                                }
                                            case "hello":
                                                switch (fi27 % ((-799) | 1)) {
                                                case 4294967295:
                                                    var v42 uint16 = 65535
                                                    _ = v42
                                                case 32768:
                                                    var v43 int8 = -42
                                                    _ = v43
                                                case 0:
                                                    switch (fi27 - (v31 * fi30)) {
                                                    case 35:
                                                        v44 := GenS5{F0: 32767}
                                                        _ = v44
                                                    case 128:
                                                        v3 = 0
                                                    }
                                                }
                                            default:
                                                v31 = 32768
                                            }
                                        }
                                    default:
                                        var v45 int8 = 32
                                        _ = v45
                                    }
                                    v31 = fi27
                                    _ = v34
                                } else {
                                    v1 = (-1 > -726)
                                    if v1 {
                                        var v46 int = -1
                                        if true {
                                            var v47 uint = (v4 + v4)
                                            var v48 int8 = 1
                                            _ = v47
                                            _ = v48
                                        } else {
                                            var v49 float64 = v29
                                            _ = v49
                                        }
                                        _ = v46
                                    } else {
                                        switch "\n\t" {
                                        case "a":
                                            v29 += (v29 + v29)
                                        }
                                    }
                                    v2 -= (^(v2))
                                }
                            } else {
                                switch (fi27 ^ 32768) {
                                case 83:
                                    if len(v32) > 0 {
                                        v32[0] = -9223372036854775808
                                    }
                                case 1:
                                    v3 = 2
                                    v29 = (v29 / v29)
                                default:
                                    fmt.Println(v2, fi30)
                                }
                            }
                        case 4294967295:
                            var v50 uint8 = 1
                            _ = v50
                        default:
                            _ = len(v32)
                        }
                        var v51 uint = v4
                        _ = v32
                        _ = v51
                    }
                    _ = fi30
                    _ = v31
                }
            } else {
                var v52 uint8 = 42
                _ = v52
            }
        default:
            v53 := GenS7{F0: (v4 & v4), F1: 598}
            _ = v53
        }
        _ = fi27
        _ = v28
        _ = v29
    }
    var v54 uint8 = 0
    if ((4294967295 > uint(v3)) || (!v1)) {
        v4 -= 0
        var v55 int8 = 127
        _ = v55
    }
    var v56 int8 = 100
    for fi57 := 0; fi57 < 14; fi57++ {
        for fi58 := 0; fi58 < 3; fi58++ {
            if true {
                v56 = v56
            }
            _ = fi58
        }
        switch 128 {
        case -1:
            for fi59 := 0; fi59 < 3; fi59++ {
                var v60 uint64 = (v3 - (v3 * v3))
                v61 := []int{}
                v62 := GenS4{F0: 831, F1: v1}
                _ = fi59
                _ = v60
                _ = v61
                _ = v62
            }
        case 528:
            if true {
                switch 32767 {
                case 42:
                    v63 := []int{42, (fi57 >> (((v4 | v4)) & 63)), fi57}
                    _ = v63
                default:
                    for fi64 := 0; fi64 < 2; fi64++ {
                        switch (fi64 * (fi57 << ((v4) & 63))) {
                        case 127:
                            v56 -= int8(v54)
                        case 764:
                            v1 = v1
                        }
                        if v1 {
                            switch (fi64 - fi57) {
                            case -439:
                                fmt.Println(v1)
                            case -747:
                                var v65 int = int(v4)
                                _ = v65
                            case 4294967295:
                                v54 = (^(v54))
                            default:
                                fmt.Println(fi57)
                            }
                            fmt.Println(v54, v1, v4)
                            fmt.Println(v3, v1)
                        } else {
                            v66 := []int{1, fi64}
                            fmt.Println(v3)
                            if false {
                                v67 := GenS8{F0: "foo bar"}
                                _ = v67
                            }
                            _ = v66
                        }
                        v68 := []int{(fi64 ^ (fi64 / ((fi64) | 1)))}
                        v2 -= (^(v2))
                        _ = fi64
                        _ = v68
                    }
                }
            }
        case 42:
            v69 := []int{fi57, (fi57 * -305), fi57, 631}
            _ = v69
        default:
            var v70 float64 = 3.14159
            _ = v70
        }
        _ = fi57
    }
    fmt.Println(v3, v1)
    v1 = v1
    if (true || (!(((1.5 / 1.0) / ((1e10 - 0.0) / 3.14159)) > 0.1))) {
        v71 := GenS8{F0: "a"}
        var v72 uint16 = 42
        _ = v71
        _ = v72
    } else {
        v73 := []int{-411, 9223372036854775807, -1}
        if true {
            _ = len(v73)
        } else {
            if (!false) {
                v74 := GenS7{F0: (v4 % (((v4 << ((2) & 63))) | 1)), F1: 0}
                _ = len(v73)
                _ = v74
            }
            var v75 int32 = genFn0()
            _ = v75
        }
        _ = v73
    }
    if true {
        if v1 {
            fmt.Println(v56, v54, v56)
            for fi76 := 0; fi76 < 17; fi76++ {
                var v77 int32 = 470
                switch int(v77) {
                case 128:
                    for fi78 := 0; fi78 < 15; fi78++ {
                        fmt.Println(v1, v56, v56)
                        var v79 uint32 = 1
                        switch fi78 {
                        case 2147483647:
                            v80 := GenS6{F0: 96}
                            _ = v80
                        case 9223372036854775807:
                            var v81 bool = ((v54 > uint8(v4)) && ((^(v54)) < (v54 + (^(v54)))))
                            _ = v81
                        case 128:
                            var v82 int64 = 9223372036854775807
                            _ = v82
                        default:
                            var v83 float32 = (0.1 - (((1e10 - 3.14159) - 1e10) / 0.1))
                            _ = v83
                        }
                        _ = fi78
                        _ = v79
                    }
                case -20:
                    switch (^(fi76)) {
                    case -1:
                        if v1 {
                            v84 := []int{fi76}
                            _ = v84
                        } else {
                            if (845 != 42) {
                                var v85 string = "foo bar"
                                fmt.Println(v1, v3)
                                v1 = (genFn1(v1) > (v2 << ((v4) & 63)))
                                _ = v85
                            } else {
                                v54 = v54
                            }
                            var v86 bool = (false && v1)
                            if v86 {
                                var v87 int64 = int64(v4)
                                v88 := []int{fi76}
                                _ = v87
                                _ = v88
                            } else {
                                var v89 int8 = (v56 % ((v56) | 1))
                                if (v1 && v86) {
                                    v90 := []int{fi76}
                                    v1 = v1
                                    _ = v90
                                } else {
                                    if v86 {
                                        v3 += (^(v3))
                                        for fi91 := 0; fi91 < 7; fi91++ {
                                            var v92 int16 = int16(v77)
                                            _ = fi91
                                            _ = v92
                                        }
                                        for fi93 := 0; fi93 < 15; fi93++ {
                                            v94 := GenS9{F0: v3}
                                            _ = fi93
                                            _ = v94
                                        }
                                    } else {
                                        if (v86 && true) {
                                            v89 = v56
                                            fmt.Println(v86, v3, fi76)
                                        } else {
                                            var v95 bool = v86
                                            _ = v95
                                        }
                                        v86 = (!false)
                                    }
                                }
                                for fi96 := 0; fi96 < 19; fi96++ {
                                    if (2147483648 == uint32(v54)) {
                                        var v97 int8 = (v89 ^ (v89 % (((^(v89))) | 1)))
                                        var v98 int16 = 1
                                        _ = v97
                                        _ = v98
                                    }
                                    v3 = (v3 * 42)
                                    var v99 float32 = 0.0
                                    fmt.Println(v3)
                                    _ = fi96
                                    _ = v99
                                }
                                _ = v89
                            }
                            _ = v86
                        }
                    case -2147483648:
                        var v100 uint16 = (v2 | 255)
                        _ = v100
                    }
                case 32767:
                    if (v1 && v1) {
                        var v101 uint = v4
                        _ = v101
                    }
                }
                _ = fi76
                _ = v77
            }
        } else {
            switch 2147483647 {
            case -9223372036854775808:
                v102 := GenS8{F0: "\n\t"}
                _ = v102
            }
            fmt.Println(v54, v1)
            v103 := GenS10{F0: -512, F1: 1, F2: genFn1((((733 <= 42) && v1) || (false || v1)))}
            _ = v103
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v54
    _ = v56
}
