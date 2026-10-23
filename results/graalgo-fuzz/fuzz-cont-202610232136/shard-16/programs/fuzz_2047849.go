package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int; F1 uint8; F2 string }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 uint64; F1 uint32 }
type GenS6 struct { F0 int16; F1 int32; F2 int }
func genFn0() string {
    if ((1000 == -32768) || (-477 > -9223372036854775808)) {
        v5 := []int{2147483647}
        _ = v5
    }
    var v6 bool = (106 == 42)
    switch -1 {
    case 127:
        if (-2.25 == 0.1) {
            v6 = (644 == 92)
            switch "" {
            case "":
                var v7 uint32 = 4294967295
                _ = v7
            case "a":
                v6 = (3.14159 >= 0.1)
            default:
                if v6 {
                    v6 = (612 != -1)
                }
            }
            for fi8 := 0; fi8 < 11; fi8++ {
                if false {
                    for fi9 := 0; fi9 < 15; fi9++ {
                        var v10 int32 = 0
                        var v11 uint = 4294967295
                        v6 = (true || v6)
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    var v12 float64 = 1e10
                    _ = v12
                }
                if true {
                    switch "\n\t" {
                    case "":
                        fmt.Println(fi8)
                    case "\n\t":
                        var v13 int16 = 32767
                        _ = v13
                    case "hello":
                        v14 := []int{}
                        _ = v14
                    default:
                        if v6 {
                            v15 := GenS1{F0: -788}
                            _ = v15
                        }
                    }
                } else {
                    var v16 bool = (true && (!true))
                    _ = v16
                }
                v6 = (v6 && ((!(4294967295 == 4294967295)) || false))
                _ = fi8
            }
        } else {
            var v17 uint32 = 2
            v18 := GenS1{F0: -9223372036854775808}
            _ = v17
            _ = v18
        }
    case -2147483648:
        var v19 uint16 = 2
        _ = v19
    case 32767:
        switch "foo bar" {
        case "":
            v20 := GenS1{F0: 9223372036854775807}
            _ = v20
        case "a":
            v21 := GenS1{F0: -1}
            _ = v21
        }
    }
    _ = v6
    return ("foo bar" + "hello")
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 uint8 = 1
    v4 := make(map[string]uint32)
    for fi22 := 0; fi22 < 10; fi22++ {
        for fi23 := 0; fi23 < 11; fi23++ {
            v3 -= v3
            _ = len(v1)
            var v24 int32 = 2147483647
            _ = fi23
            _ = v24
        }
        if true {
            v3 = (v3 % ((v3) | 1))
            if true {
                v3 = 0
                switch "a" {
                case "":
                    v3 -= v3
                }
            } else {
                if (!false) {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if true {
                            v3 += v3
                        }
                        var v26 string = "\n\t"
                        _ = fi25
                        _ = v26
                    }
                }
            }
            for fi27 := 0; fi27 < 6; fi27++ {
                var v28 int16 = 945
                v28 += v28
                _ = fi27
                _ = v28
            }
        } else {
            v29 := GenS1{F0: 0}
            v30 := GenS2{F0: 4294967295, F1: (v3 >> ((2) & 63)), F2: "hello"}
            _ = v29
            _ = v30
        }
        fmt.Println(v3, v3)
        _ = fi22
    }
    if (-1 > -9223372036854775808) {
        var v31 uint64 = 1
        if false {
            var v32 uint8 = 173
            var v33 int = 1
            _ = v32
            _ = v33
        }
        fmt.Println(v31, v31)
        _ = v31
    }
    for fi34 := 0; fi34 < 3; fi34++ {
        fmt.Println(v3)
        _ = fi34
    }
    var v35 uint32 = 1
    v3 += 42
    if (v3 >= (^(v3))) {
        v36 := []int{-271, -93, -480}
        _ = v36
    }
    v37 := GenS1{F0: 42}
    if false {
        for fi38 := 0; fi38 < 16; fi38++ {
            v39 := GenS3{F0: -1, F1: (v3 * v3)}
            _ = fi38
            _ = v39
        }
        v35 += v35
    } else {
        fmt.Println(v3)
        fmt.Println(v3)
    }
    v40 := GenS2{F0: -2147483648, F1: (v3 >> ((432) & 63)), F2: (("foo bar" + "\n\t") + ((("" + ("a" + ((("hello" + "") + "a") + (("\n\t" + "hello") + "\n\t")))) + (("\n\t" + "\n\t") + "foo bar")) + ""))}
    var v41 float32 = (0.0 * 1.5)
    for fi42 := 0; fi42 < 8; fi42++ {
        fmt.Println(fi42, v35, v3)
        if true {
            v43 := []int{-1}
            fmt.Println(fi42)
            _ = v43
        } else {
            switch 32768 {
            case 2147483647:
                v41 = ((float32((0.1 / (1e10 / (0.1 / 1.0)))) * v41) * 1.5)
            case 1:
                v41 = (3.14159 / -2.25)
            case 32767:
                delete(v4, "hello")
            default:
                v41 = (v41 / v41)
            }
            if (!true) {
                v3 = uint8(v35)
                var v44 int64 = 9223372036854775807
                _ = v44
            }
            fmt.Println(v41)
        }
        _ = fi42
    }
    switch 1 {
    case 42:
        var v45 uint64 = 2
        _ = v45
    case 4294967295:
        var v46 int64 = 9223372036854775807
        _ = v46
    case 128:
        delete(v2, 32768)
    default:
        if (false || (-131 != -937)) {
            var v47 string = ((("hello" + "a") + "a") + "")
            _ = v47
        }
    }
    if true {
        var v48 uint32 = 2147483647
        _ = v48
    } else {
        fmt.Println(v3, v3, v41)
        if false {
            var v49 int16 = 0
            _ = v49
        } else {
            if false {
                v50 := []int{-2147483648, -769}
                var v51 uint64 = 278
                v35 = (v35 - v35)
                _ = v50
                _ = v51
            } else {
                if true {
                    for fi52 := 0; fi52 < 5; fi52++ {
                        switch (-(fi52)) {
                        case 127:
                            v53 := GenS2{F0: -2147483648, F1: v3, F2: "a"}
                            _ = v53
                        default:
                            var v54 string = "hello"
                            _ = v54
                        }
                        _ = len(v1)
                        var v55 float64 = 0.0
                        _ = fi52
                        _ = v55
                    }
                    if false {
                        fmt.Println(v35, v3, v35)
                        switch 233 {
                        case 32767:
                            v56 := GenS3{F0: 32768, F1: (v3 * 65)}
                            _ = v56
                        }
                    } else {
                        v35 = v35
                        switch (("hello" + "\n\t") + ("a" + ("a" + ("hello" + "\n\t")))) {
                        case "\n\t":
                            if ((true && (0.0 < 0.0)) || false) {
                                if (false || (("\n\t" + ("hello" + "a")) <= "a")) {
                                    var v57 float64 = 3.14159
                                    _ = v57
                                }
                            }
                        case "":
                            if (true || false) {
                                if (255 == 65535) {
                                    for fi58 := 0; fi58 < 10; fi58++ {
                                        var v59 int64 = -1
                                        if (1 != 42) {
                                            if false {
                                                for fi60 := 0; fi60 < 9; fi60++ {
                                                    v61 := []int{fi58, (^(fi60))}
                                                    _ = fi60
                                                    _ = v61
                                                }
                                                v62 := GenS1{F0: -1}
                                                _ = v62
                                            }
                                            switch "hello" {
                                            case "":
                                                if true {
                                                    fmt.Println(v41)
                                                    v63 := GenS4{F0: -694, F1: "a"}
                                                    _ = v63
                                                }
                                            }
                                        } else {
                                            v64 := []int{32768, (^(fi58)), fi58}
                                            v59 = int64(fi58)
                                            _ = v64
                                        }
                                        _ = fi58
                                        _ = v59
                                    }
                                    v3 += (v3 >> ((290) & 63))
                                }
                                if (v3 != v3) {
                                    var v65 int64 = 0
                                    var v66 int16 = -32768
                                    _ = v65
                                    _ = v66
                                }
                                fmt.Println(v41, v3, v41)
                            }
                        }
                    }
                    v67 := GenS1{F0: -1}
                    _ = v67
                }
                for fi68 := 0; fi68 < 11; fi68++ {
                    delete(v2, fi68)
                    if (true || (!true)) {
                        v3 = (v3 * v3)
                        var v69 uint8 = 2
                        v35 = v35
                        _ = v69
                    } else {
                        v70 := GenS3{F0: (fi68 / (((fi68 ^ (fi68 >> ((756) & 63)))) | 1)), F1: v3}
                        delete(v4, ("foo bar" + "a"))
                        _ = v70
                    }
                    _ = fi68
                }
            }
            v71 := GenS5{F0: 0, F1: 0}
            _ = v71
        }
    }
    v72 := GenS5{F0: 2, F1: uint32(v3)}
    v3 = (v3 / ((v3) | 1))
    switch -507 {
    case -1:
        v35 = v35
    case -9223372036854775808:
        if (v35 != v35) {
            switch ("" + "\n\t") {
            case "foo bar":
                var v73 int16 = 704
                _ = v73
            case "a":
                if (!false) {
                    delete(v4, "a")
                    v74 := GenS6{F0: 430, F1: -2147483648, F2: 492}
                    _ = v74
                } else {
                    v35 += 4294967295
                    v75 := GenS2{F0: 0, F1: v3, F2: ""}
                    fmt.Println(v41)
                    _ = v75
                }
            }
            for fi76 := 0; fi76 < 5; fi76++ {
                fmt.Println(v35, fi76, fi76)
                v77 := []int{fi76, (fi76 >> ((1) & 63)), int(v35), (fi76 / ((fi76) | 1))}
                _ = fi76
                _ = v77
            }
            v78 := []int{-543, 1, -9223372036854775808, -2147483648}
            _ = v78
        }
    case 1:
        switch ("" + ("hello" + "")) {
        case "":
            if ((("foo bar" + ("hello" + "hello")) >= "") && true) {
                v35 = 4294967295
            }
        default:
            if false {
                switch 32767 {
                case 32767:
                    if (!false) {
                        v2[127] = (("hello" + "foo bar") + ((("a" + "foo bar") + ("foo bar" + ("a" + "a"))) + "a"))
                        var v79 uint16 = 398
                        _ = v79
                    } else {
                        v80 := GenS2{F0: -2147483648, F1: v3, F2: "foo bar"}
                        _ = v80
                    }
                case 9223372036854775807:
                    v3 = (v3 * uint8(v35))
                case -179:
                    var v81 float32 = v41
                    _ = v81
                default:
                    switch 32768 {
                    case -364:
                        var v82 uint8 = (v3 - v3)
                        _ = v82
                    }
                }
            }
        }
    }
    switch -504 {
    case -9223372036854775808:
        for fi83 := 0; fi83 < 11; fi83++ {
            v84 := GenS1{F0: 94}
            _ = fi83
            _ = v84
        }
    case -565:
        var v85 int32 = 0
        _ = v85
    case -972:
        if (v41 == 3.14159) {
            var v86 uint = 4294967295
            var v87 int16 = -1
            v41 = 0.0
            _ = v86
            _ = v87
        }
    default:
        switch "hello" {
        case "a":
            v88 := []int{-185, -1, 128}
            _ = v88
        default:
            v41 = v41
        }
    }
    var v89 int32 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v72
    _ = v89
}
