package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 int32; F1 uint32; F2 uint64 }
type GenS6 struct { F0 uint; F1 int16 }
type GenS7 struct { F0 uint; F1 uint }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint32, p1 uint32) int {
    v7 := []int{127}
    var v8 uint8 = 127
    for fi9 := 0; fi9 < 3; fi9++ {
        v10 := GenS1{F0: (v8 ^ 127)}
        v11 := GenS2{F0: 42, F1: -399}
        var v12 int64 = 42
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 32768
}

func main() {
    var v1 float32 = (-2.25 * -2.25)
    v2 := []string{"foo bar", "\n\t", ("a" + "")}
    var v3 string = "a"
    v4 := []uint32{923, 24, 4294967295}
    var v5 string = ""
    var v6 int8 = -1
    for fi13 := 0; fi13 < 15; fi13++ {
        switch ("foo bar" + v5) {
        case "hello":
            if len(v4) > 0 {
                v4[0] = 0
            }
        case "foo bar":
            switch -785 {
            case 1:
                fmt.Println(v3)
            case 42:
                v3 = v5
            case 0:
                if (!false) {
                    if false {
                        if (("a" + (v3 + v3)) > ("a" + "")) {
                            v3 = v5
                            for fi14 := 0; fi14 < 15; fi14++ {
                                if false {
                                    v6 = v6
                                }
                                _ = len(v2)
                                v15 := GenS1{F0: 2}
                                _ = fi14
                                _ = v15
                            }
                        }
                        var v16 uint16 = 527
                        var v17 float64 = 1e10
                        _ = v16
                        _ = v17
                    }
                }
            default:
                var v18 int32 = -2147483648
                _ = v18
            }
        }
        var v19 float32 = (v1 - (3.14159 + (v1 * 0.1)))
        _ = fi13
        _ = v19
    }
    v20 := GenS2{F0: -74, F1: -1}
    v21 := GenS2{F0: 42, F1: 1000}
    switch 32768 {
    case 4294967295:
        var v22 uint64 = 2
        _ = v22
    default:
        v3 = ("a" + "a")
    }
    if false {
        v23 := GenS2{F0: 42, F1: -538}
        v6 = v6
        var v24 int64 = -1
        _ = v23
        _ = v24
    } else {
        if (!true) {
            var v25 int8 = (-(v6))
            fmt.Println(v5)
            _ = v25
        }
        var v26 int8 = v6
        var v27 bool = (166 < 2)
        _ = v26
        _ = v27
    }
    v2 = append(v2, ("" + v5))
    v5 = (v5 + ("a" + (v5 + "foo bar")))
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 int16 = 1000
        for fi30 := 0; fi30 < 1; fi30++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v31 float32 = v1
            fmt.Println(fi28)
            var v32 int32 = 42
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    if false {
        v5 = v5
        for fi33 := 0; fi33 < 10; fi33++ {
            for fi34 := 0; fi34 < 6; fi34++ {
                _ = len(v4)
                _ = len(v4)
                _ = fi34
            }
            if true {
                v4 = append(v4, 968)
                v6 -= (v6 - v6)
            }
            var v35 int16 = 42
            _ = fi33
            _ = v35
        }
        if true {
            for fi36 := 0; fi36 < 20; fi36++ {
                if true {
                    if len(v4) > 0 {
                        v4[0] = 2147483647
                    }
                    v37 := GenS3{F0: -32768}
                    fmt.Println(v3, v3, fi36)
                    _ = v37
                }
                _ = fi36
            }
        } else {
            for fi38 := 0; fi38 < 2; fi38++ {
                switch fi38 {
                case 0:
                    v39 := GenS4{F0: -32768, F1: genFn0(303, 2147483648)}
                    _ = v39
                }
                fmt.Println(fi38, fi38, v5)
                if (255 < 747) {
                    switch fi38 {
                    case 32767:
                        if false {
                            var v40 float64 = 0.0
                            for fi41 := 0; fi41 < 15; fi41++ {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                var v42 float64 = v40
                                _ = fi41
                                _ = v42
                            }
                            var v43 uint64 = 103
                            _ = v40
                            _ = v43
                        }
                    case 4294967295:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case -32:
                        for fi44 := 0; fi44 < 4; fi44++ {
                            v45 := GenS5{F0: -677, F1: 42, F2: 42}
                            if true {
                                v3 = v5
                                var v46 uint32 = 766
                                if (!(fi38 > (fi38 ^ -9223372036854775808))) {
                                    fmt.Println(v1)
                                    if (2 >= 18446744073709551615) {
                                        var v47 int8 = (v6 % ((v6) | 1))
                                        v47 -= v47
                                        if false {
                                            _ = len(v2)
                                            fmt.Println(v6)
                                            fmt.Println(v6, v5)
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        }
                                        _ = v47
                                    }
                                } else {
                                    switch (("hello" + v5) + ((v5 + v3) + "")) {
                                    case "\n\t":
                                        fmt.Println(fi38)
                                    case "a":
                                        fmt.Println(fi38, v3)
                                    default:
                                        if false {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            v1 -= (v1 - v1)
                                        }
                                    }
                                    if ((v3 < ("a" + v5)) || false) {
                                        if (!(-229 > 32767)) {
                                            switch fi44 {
                                            case 394:
                                                var v48 bool = false
                                                _ = v48
                                            case -707:
                                                switch v5 {
                                                case "":
                                                    fmt.Println(v5)
                                                default:
                                                    if ((fi38 << ((318) & 63)) <= int(v6)) {
                                                        var v49 string = (v3 + ("foo bar" + "a"))
                                                        v46 -= (v46 | v46)
                                                        if true {
                                                            switch (fi38 << ((48) & 63)) {
                                                            case 385:
                                                                v50 := GenS6{F0: 2, F1: 42}
                                                                _ = v50
                                                            default:
                                                                var v51 float64 = ((float64(v1) / 3.14159) * 3.14159)
                                                                _ = v51
                                                            }
                                                            v46 += 317
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        _ = v49
                                                    } else {
                                                        var v52 float64 = -2.25
                                                        if false {
                                                            var v53 int64 = 9223372036854775807
                                                            _ = v53
                                                        } else {
                                                            v54 := GenS1{F0: 127}
                                                            v52 -= -2.25
                                                            switch 32767 {
                                                            case 32767:
                                                                var v55 int16 = 581
                                                                _ = v55
                                                            default:
                                                                fmt.Println(v52, v46)
                                                            }
                                                            _ = v54
                                                        }
                                                        v1 += v1
                                                        _ = v52
                                                    }
                                                }
                                                var v56 uint = 2
                                                _ = v56
                                            }
                                            fmt.Println(fi38, v46, v6)
                                            if true {
                                                v3 = "\n\t"
                                                fmt.Println(v3, v5)
                                                v46 += (v46 >> ((4294967295) & 63))
                                            }
                                        } else {
                                            var v57 uint = 55
                                            v58 := GenS1{F0: 42}
                                            _ = v57
                                            _ = v58
                                        }
                                        v3 = v3
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = v46
                            }
                            v3 = ("a" + v5)
                            _ = fi44
                            _ = v45
                        }
                    default:
                        switch 32767 {
                        case 9223372036854775807:
                            if false {
                                for fi59 := 0; fi59 < 17; fi59++ {
                                    var v60 int32 = -2147483648
                                    v61 := GenS5{F0: v60, F1: 65535, F2: 9}
                                    _ = fi59
                                    _ = v60
                                    _ = v61
                                }
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(fi38, fi38, v6)
                            } else {
                                if false {
                                    v62 := GenS6{F0: 456, F1: -452}
                                    if true {
                                        var v63 uint8 = 1
                                        for fi64 := 0; fi64 < 7; fi64++ {
                                            switch 0 {
                                            case 202:
                                                var v65 uint8 = (v63 % ((v63) | 1))
                                                _ = v65
                                            default:
                                                v66 := GenS3{F0: -293}
                                                _ = v66
                                            }
                                            _ = fi64
                                        }
                                        for fi67 := 0; fi67 < 3; fi67++ {
                                            if (-32768 >= 951) {
                                                var v68 uint16 = 2
                                                v68 += v68
                                                _ = v68
                                            } else {
                                                var v69 int16 = 597
                                                var v70 uint32 = 466
                                                switch ((v3 + v3) + v3) {
                                                case "hello":
                                                    if ((v1 < (v1 / (v1 * ((v1 * ((v1 * ((((v1 + ((v1 / v1) + float32(1e10))) * (v1 + (v1 - (v1 - (v1 / (float32(3.14159) * (v1 - v1))))))) + float32(float64(0.1))) / (float32(3.14159) + v1))) / 1.0)) + float32(float64(-2.25)))))) && ((v69 == 912) || ((18446744073709551615 <= 18446744073709551615) || true))) {
                                                        v71 := GenS6{F0: 197, F1: (v69 | (v69 * (v69 & -32768)))}
                                                        _ = v71
                                                    }
                                                case "foo bar":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                case "":
                                                    var v72 bool = (!((3.14159 >= 1e10) || (2147483647 != -1)))
                                                    _ = v72
                                                }
                                                _ = v69
                                                _ = v70
                                            }
                                            v1 = (v1 + -2.25)
                                            switch v5 {
                                            case "a":
                                                var v73 float32 = (v1 / v1)
                                                _ = v73
                                            case "\n\t":
                                                fmt.Println(v1, v6, v63)
                                            default:
                                                var v74 int64 = -1
                                                _ = v74
                                            }
                                            _ = fi67
                                        }
                                        _ = v63
                                    }
                                    _ = v62
                                } else {
                                    switch (((v3 + v3) + v5) + (v3 + "hello")) {
                                    case "a":
                                        for fi75 := 0; fi75 < 8; fi75++ {
                                            var v76 bool = true
                                            var v77 int16 = 0
                                            _ = fi75
                                            _ = v76
                                            _ = v77
                                        }
                                    }
                                    if (false && true) {
                                        if (!(!(fi38 >= (^(fi38))))) {
                                            if true {
                                                v78 := GenS7{F0: 1, F1: 2}
                                                v79 := GenS4{F0: -707, F1: fi38}
                                                v80 := GenS5{F0: 2147483647, F1: 2, F2: 159}
                                                _ = v78
                                                _ = v79
                                                _ = v80
                                            }
                                            fmt.Println(v6, v5)
                                            var v81 uint32 = 65535
                                            _ = v81
                                        } else {
                                            var v82 int = fi38
                                            v3 = "foo bar"
                                            _ = v82
                                        }
                                        if false {
                                            for fi83 := 0; fi83 < 6; fi83++ {
                                                v2 = append(v2, (v5 + v3))
                                                var v84 uint = 4294967295
                                                _ = fi83
                                                _ = v84
                                            }
                                            _ = len(v2)
                                        }
                                    }
                                }
                                var v85 uint16 = 42
                                _ = v85
                            }
                        case -9223372036854775808:
                            switch v5 {
                            case "hello":
                                if (1 <= 255) {
                                    v86 := GenS8{F0: 1}
                                    switch (fi38 >> ((42) & 63)) {
                                    case -636:
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                    case -2147483648:
                                        if false {
                                            for fi87 := 0; fi87 < 18; fi87++ {
                                                v6 -= v6
                                                _ = fi87
                                            }
                                            var v88 string = ("hello" + "foo bar")
                                            _ = v88
                                        } else {
                                            v2 = append(v2, v5)
                                        }
                                    case 127:
                                        switch (-(fi38)) {
                                        case 4294967295:
                                            if ((v3 + ((v5 + (("hello" + ("\n\t" + v5)) + (v3 + v5))) + ((v5 + v3) + ""))) == v5) {
                                                var v89 float64 = (float64((v1 * v1)) - (float64(v1) / (3.14159 / 0.1)))
                                                v1 = 1e10
                                                _ = v89
                                            }
                                        case 2147483647:
                                            for fi90 := 0; fi90 < 18; fi90++ {
                                                v6 = -128
                                                fmt.Println(v6, v3, fi90)
                                                v5 = (v3 + (("a" + (((((((v3 + "a") + "a") + ("" + v5)) + "foo bar") + ((v3 + (((v5 + v3) + v3) + v5)) + ("\n\t" + v5))) + (v5 + (v3 + v3))) + "")) + ("hello" + v3)))
                                                v1 += 1.5
                                                _ = fi90
                                            }
                                        default:
                                            v91 := GenS3{F0: -1}
                                            _ = v91
                                        }
                                    default:
                                        v6 -= v6
                                    }
                                    _ = v86
                                } else {
                                    for fi92 := 0; fi92 < 11; fi92++ {
                                        _ = len(v4)
                                        switch (^(fi38)) {
                                        case -169:
                                            v1 += float32((1.5 / 3.14159))
                                        }
                                        _ = fi92
                                    }
                                }
                            case "foo bar":
                                for fi93 := 0; fi93 < 14; fi93++ {
                                    if true {
                                        var v94 string = "a"
                                        v3 = "hello"
                                        v4 = append(v4, 0)
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v6 = v6
                                        _ = v94
                                    }
                                    switch "a" {
                                    case "":
                                        v95 := GenS2{F0: 469, F1: 392}
                                        _ = v95
                                    default:
                                        v3 = (v5 + v3)
                                    }
                                    _ = fi93
                                }
                            }
                        }
                    }
                    var v96 uint32 = 0
                    for fi97 := 0; fi97 < 20; fi97++ {
                        v98 := GenS4{F0: -770, F1: (fi97 % ((-9223372036854775808) | 1))}
                        var v99 uint32 = 221
                        v5 = ("\n\t" + v5)
                        _ = fi97
                        _ = v98
                        _ = v99
                    }
                    _ = v96
                }
                _ = fi38
            }
        }
    }
    v2 = append(v2, ("foo bar" + (v5 + ("foo bar" + "hello"))))
    _ = len(v2)
    if (42 != -1) {
        v100 := GenS1{F0: 1}
        _ = len(v4)
        var v101 uint8 = 255
        _ = v100
        _ = v101
    } else {
        _ = len(v4)
        fmt.Println(v6, v5)
        fmt.Println(v5, v6, v3)
    }
    if ((-32768 >= 1) && true) {
        v5 = (v5 + ((("" + v3) + ((("a" + "a") + ((((v5 + v5) + "") + ((v3 + v5) + v3)) + v3)) + v5)) + v5))
    } else {
        if len(v2) > 0 {
            v2[0] = v3
        }
    }
    var v102 uint8 = 0
    var v103 int16 = 1
    if (((v3 + "") + v3) <= v3) {
        v6 -= int8(v103)
    } else {
        for fi104 := 0; fi104 < 4; fi104++ {
            v105 := GenS9{F0: 2147483647, F1: 0}
            _ = fi104
            _ = v105
        }
        fmt.Println(v5, v3)
    }
    var v106 uint16 = 0
    v5 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v102:", v102)
    fmt.Println("v103:", v103)
    fmt.Println("v106:", v106)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v102
    _ = v103
    _ = v106
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
