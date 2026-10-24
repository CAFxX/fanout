package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 string }
type GenS4 struct { F0 uint8; F1 uint32; F2 int32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string }
func genFn0(p0 uint8, p1 uint8) uint16 {
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := GenS1{F0: 1}
        fmt.Println(p0)
        v9 := GenS1{F0: -303}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        if ((1 >= 2) && true) {
            p0 += p1
            p1 = (p0 * 1)
            if true {
                p0 = (p0 % (((p1 | p1)) | 1))
                p0 = p1
            }
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 float64 = (1.5 * 0.0)
            if true {
                switch "a" {
                case "hello":
                    var v13 string = "hello"
                    _ = v13
                case "a":
                    if false {
                        if true {
                            var v14 int = fi11
                            p1 = (p1 & p1)
                            _ = v14
                        }
                        v15 := []int{}
                        var v16 uint = 482
                        _ = v15
                        _ = v16
                    } else {
                        var v17 float32 = 0.1
                        _ = v17
                    }
                case "":
                    var v18 int64 = 42
                    _ = v18
                default:
                    var v19 int64 = 0
                    _ = v19
                }
                fmt.Println(fi11, fi11, p1)
                p0 += p1
            } else {
                v20 := GenS2{F0: 32767, F1: 2147483647}
                v21 := GenS1{F0: -220}
                _ = v20
                _ = v21
            }
            v22 := GenS1{F0: 173}
            if true {
                v23 := []int{fi10, fi11}
                if (!(!(531 > 42))) {
                    p1 = (p1 & (^(p1)))
                }
                _ = v23
            }
            if false {
                v24 := []int{}
                if true {
                    var v25 uint16 = 467
                    v12 = v12
                    _ = v25
                }
                _ = v24
            } else {
                p0 = (p1 / ((p1) | 1))
            }
            _ = fi11
            _ = v12
            _ = v22
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    return 446
}
func genFn1() uint16 {
    for fi26 := 0; fi26 < 11; fi26++ {
        var v27 int64 = -9223372036854775808
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS1{F0: 42}
            fmt.Println(fi28, fi29)
            v31 := GenS3{F0: 2147483647, F1: 1, F2: ("foo bar" + "\n\t")}
            switch -2147483648 {
            case 4294967295:
                var v32 uint8 = 101
                _ = v32
            }
            var v33 uint16 = 1
            _ = fi29
            _ = v30
            _ = v31
            _ = v33
        }
        _ = fi28
    }
    var v34 int = -1
    v35 := []int{}
    if len(v35) > 0 {
        v35[0] = -2147483648
    }
    _ = v34
    _ = v35
    return genFn0(1, 255)
}

func main() {
    var v1 float64 = float64(1e10)
    var v2 int64 = 0
    v3 := make(map[uint32]uint32)
    v4 := []uint32{0, 111, 4294967295}
    v5 := make(map[uint32]string)
    var v6 int8 = 35
    v1 = v1
    if (true || true) {
        for fi36 := 0; fi36 < 10; fi36++ {
            switch fi36 {
            case 128:
                if len(v4) > 0 {
                    v4[0] = 2147483648
                }
            case 42:
                switch "foo bar" {
                case "\n\t":
                    v6 = v6
                case "hello":
                    v6 = (v6 | (v6 / ((-83) | 1)))
                case "":
                    v1 += v1
                default:
                    if true {
                        _ = len(v3)
                    }
                }
            default:
                if len(v4) > 0 {
                    v4[0] = 2147483647
                }
            }
            for fi37 := 0; fi37 < 5; fi37++ {
                var v38 bool = (!false)
                switch "a" {
                case "a":
                    if (!(18446744073709551615 == 1)) {
                        if v38 {
                            if (-2.25 != float32(float64(1e10))) {
                                v1 -= v1
                            } else {
                                v2 = v2
                                switch (fi37 & fi36) {
                                case 978:
                                    switch fi37 {
                                    case 682:
                                        v38 = (!("" != ""))
                                    default:
                                        fmt.Println(v6)
                                    }
                                case 418:
                                    v6 = (v6 / ((-82) | 1))
                                }
                                v39 := GenS3{F0: 42, F1: 32767, F2: "a"}
                                if v38 {
                                    if len(v4) > 0 {
                                        v4[0] = 2147483648
                                    }
                                } else {
                                    v6 += 42
                                }
                                _ = v39
                            }
                            var v40 bool = true
                            _ = v40
                        }
                        var v41 int64 = -1
                        _ = v41
                    } else {
                        v1 = v1
                    }
                case "":
                    if false {
                        v38 = v38
                    }
                case "hello":
                    v42 := GenS4{F0: 128, F1: 4294967295, F2: -629}
                    _ = v42
                }
                _ = fi37
                _ = v38
            }
            v6 = v6
            _ = fi36
        }
    } else {
        v1 += v1
        if ((!true) && ((9223372036854775807 >= -797) && false)) {
            fmt.Println(v1, v2)
        }
    }
    switch 128 {
    case 800:
        for fi43 := 0; fi43 < 4; fi43++ {
            v44 := GenS1{F0: 0}
            v45 := GenS5{F0: "\n\t"}
            var v46 uint64 = 29
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
    default:
        if (!true) {
            if false {
                _ = len(v5)
                var v47 float32 = (1.5 + float32((v1 * float64((0.0 + 1e10)))))
                for fi48 := 0; fi48 < 13; fi48++ {
                    v2 += -9223372036854775808
                    if ((("" + "a") + (("foo bar" + "") + "\n\t")) == "\n\t") {
                        for fi49 := 0; fi49 < 6; fi49++ {
                            switch ("" + ("foo bar" + "\n\t")) {
                            case "hello":
                                if false {
                                    switch "a" {
                                    case "":
                                        var v50 int64 = -1
                                        _ = v50
                                    }
                                    if (genFn0(1, 210) > 363) {
                                        fmt.Println(fi49, v1)
                                        fmt.Println(v2, v6)
                                    }
                                    switch "foo bar" {
                                    case "hello":
                                        fmt.Println(fi49, v1)
                                    case "\n\t":
                                        var v51 int64 = (v2 & -1)
                                        _ = v51
                                    case "foo bar":
                                        var v52 float64 = float64(1e10)
                                        _ = v52
                                    }
                                }
                            case "a":
                                v6 = 96
                            default:
                                switch ("foo bar" + "foo bar") {
                                case "\n\t":
                                    v53 := GenS6{F0: 4294967295}
                                    _ = v53
                                default:
                                    v1 = -2.25
                                }
                            }
                            fmt.Println(v1)
                            switch -385 {
                            case 2147483647:
                                if false {
                                    var v54 int8 = v6
                                    _ = v54
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v55 int64 = v2
                                    v55 = (v55 % ((v2) | 1))
                                    fmt.Println(v47)
                                    _ = v55
                                }
                            case 146:
                                v1 += v1
                            default:
                                switch ("" + "\n\t") {
                                case "foo bar":
                                    if false {
                                        fmt.Println(v2, v2, fi49)
                                    } else {
                                        switch (fi48 % ((int(v2)) | 1)) {
                                        case 42:
                                            var v56 uint64 = 521
                                            _ = v56
                                        case 2147483647:
                                            v57 := GenS1{F0: -401}
                                            var v58 float32 = 3.14159
                                            _ = v57
                                            _ = v58
                                        default:
                                            v59 := v3[2]
                                            _ = v59
                                        }
                                        v2 = v2
                                        v1 = v1
                                    }
                                case "a":
                                    switch fi48 {
                                    case 127:
                                        var v60 uint32 = 639
                                        _ = v60
                                    case -9223372036854775808:
                                        fmt.Println(v1, fi48, v47)
                                    default:
                                        if (!((!true) || (!(1 < 984)))) {
                                            switch "" {
                                            case "hello":
                                                fmt.Println(fi49, fi48)
                                            case "a":
                                                if (v6 != v6) {
                                                    fmt.Println(v47, fi48)
                                                }
                                            default:
                                                fmt.Println(v6, v1, fi49)
                                            }
                                            v2 -= v2
                                            var v61 int32 = 42
                                            v62 := GenS7{F0: 2147483647}
                                            _ = v61
                                            _ = v62
                                        }
                                    }
                                case "hello":
                                    var v63 string = ("a" + (("" + "a") + "a"))
                                    _ = v63
                                default:
                                    var v64 uint = 0
                                    _ = v64
                                }
                            }
                            _ = fi49
                        }
                    } else {
                        v5[0] = "foo bar"
                        for fi65 := 0; fi65 < 16; fi65++ {
                            v66 := GenS1{F0: 518}
                            if false {
                                delete(v5, 198)
                                fmt.Println(v47, v47, v2)
                                var v67 float32 = (v47 * 1e10)
                                v1 = ((0.0 - float64(v67)) * 0.1)
                                if true {
                                    var v68 uint = 400
                                    v68 = 240
                                    var v69 int64 = (v2 << ((v68) & 63))
                                    _ = v68
                                    _ = v69
                                }
                                var v70 uint32 = 28
                                v6 -= (-(v6))
                                _ = v67
                                _ = v70
                            }
                            _ = fi65
                            _ = v66
                        }
                        if false {
                            if false {
                                fmt.Println(v2, v6, v1)
                                fmt.Println(v2, fi48)
                            }
                            switch fi48 {
                            case 1:
                                fmt.Println(v1)
                            case 2147483647:
                                var v71 uint64 = 0
                                _ = v71
                            case 32768:
                                _ = len(v4)
                            }
                        } else {
                            var v72 uint8 = 41
                            switch "" {
                            case "":
                                v72 -= (v72 - v72)
                            case "\n\t":
                                v73 := GenS5{F0: "hello"}
                                _ = v73
                            case "foo bar":
                                v5[459] = "foo bar"
                            }
                            switch 727 {
                            case 127:
                                if true {
                                    for fi74 := 0; fi74 < 11; fi74++ {
                                        v2 -= v2
                                        v75 := GenS2{F0: 42, F1: 1}
                                        _ = fi74
                                        _ = v75
                                    }
                                } else {
                                    var v76 uint = 42
                                    var v77 int32 = -58
                                    v3[577] = 1
                                    _ = v76
                                    _ = v77
                                }
                            case 47:
                                if (v1 != ((v1 / (0.1 * v1)) + ((v1 + v1) / (v1 / float64((-2.25 / float32(((float64((((float32(((float64(v47) + v1) - float64(1e10))) + (float32(3.14159) + ((3.14159 - v47) / float32(v1)))) * 1e10) + (v47 + v47))) * (1e10 - float64((v47 / v47)))) * float64(v47))))))))) {
                                    v2 += v2
                                    switch fi48 {
                                    case 32767:
                                        var v78 uint = 419
                                        _ = v78
                                    case 9223372036854775807:
                                        var v79 int64 = (v2 << ((18446744073709551615) & 63))
                                        _ = v79
                                    default:
                                        if true {
                                            v47 = ((v47 * (v47 - float32(((v1 / float64(((1e10 * (v47 + v47)) - v47))) - (v1 + (v1 - 0.1)))))) / (3.14159 / (((v47 - -2.25) * float32(-2.25)) * (float32(v1) * v47))))
                                        } else {
                                            switch (^(fi48)) {
                                            case 32768:
                                                if (!(!(112 > 104))) {
                                                    var v80 bool = false
                                                    _ = v80
                                                } else {
                                                    v81 := GenS1{F0: 2147483647}
                                                    _ = v81
                                                }
                                            }
                                            if false {
                                                v82 := GenS2{F0: 42, F1: 2147483647}
                                                _ = v82
                                            } else {
                                                v83 := GenS2{F0: 32767, F1: -2147483648}
                                                v47 = v47
                                                delete(v3, 2147483648)
                                                _ = v83
                                            }
                                            for fi84 := 0; fi84 < 13; fi84++ {
                                                if (0 != -1) {
                                                    v2 = v2
                                                    var v85 uint = 4294967295
                                                    v47 = 3.14159
                                                    _ = v85
                                                }
                                                switch "" {
                                                case "a":
                                                    delete(v5, 738)
                                                case "hello":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                var v86 int8 = (v6 % ((0) | 1))
                                                _ = fi84
                                                _ = v86
                                            }
                                        }
                                    }
                                }
                            default:
                                v1 -= v1
                            }
                            _ = v72
                        }
                    }
                    _ = fi48
                }
                _ = v47
            } else {
                v6 += v6
                _ = len(v4)
            }
            if len(v4) > 0 {
                v4[0] = 0
            }
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (("" + ("a" + "a")) + "hello") {
            case "hello":
                v3[1] = 1
            case "":
                for fi87 := 0; fi87 < 13; fi87++ {
                    fmt.Println(v6, v1, fi87)
                    var v88 string = ("foo bar" + "foo bar")
                    _ = fi87
                    _ = v88
                }
            case "\n\t":
                var v89 float64 = float64(1.5)
                _ = v89
            default:
                v2 -= -9223372036854775808
            }
        }
    }
    switch 4294967295 {
    case -1:
        v90 := GenS3{F0: 1, F1: 473, F2: ("foo bar" + ("hello" + ((("\n\t" + "a") + "a") + (("" + "hello") + ("\n\t" + ("\n\t" + ("\n\t" + "\n\t")))))))}
        _ = v90
    default:
        var v91 uint = 901
        _ = v91
    }
    for fi92 := 0; fi92 < 9; fi92++ {
        v1 = ((v1 - v1) + v1)
        _ = fi92
    }
    v1 -= (1.5 + v1)
    v93 := GenS1{F0: int32(v6)}
    if false {
        var v94 int8 = 100
        for fi95 := 0; fi95 < 7; fi95++ {
            for fi96 := 0; fi96 < 8; fi96++ {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v97 int = -9223372036854775808
                _ = fi96
                _ = v97
            }
            for fi98 := 0; fi98 < 3; fi98++ {
                v6 = v94
                var v99 bool = true
                v99 = (1 > -2147483648)
                _ = fi98
                _ = v99
            }
            v100 := GenS8{F0: "hello"}
            _ = fi95
            _ = v100
        }
        if false {
            if (42 <= 4294967295) {
                var v101 uint = 4294967295
                v102 := v3[4294967295]
                v101 = v101
                _ = v101
                _ = v102
            }
        } else {
            _ = len(v3)
        }
        _ = v94
    }
    for fi103 := 0; fi103 < 7; fi103++ {
        v2 -= v2
        for fi104 := 0; fi104 < 13; fi104++ {
            var v105 float32 = 1e10
            var v106 string = ("\n\t" + "a")
            var v107 int64 = (v2 + v2)
            _ = fi104
            _ = v105
            _ = v106
            _ = v107
        }
        v6 += (v6 >> ((737) & 63))
        _ = fi103
    }
    _ = len(v4)
    v108 := v5[2147483648]
    v4 = append(v4, 2147483647)
    var v109 uint = 927
    fmt.Println("v1:", v1)
    fmt.Println("v108:", v108)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v108
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v93
}
