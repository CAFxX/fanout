package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int8 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int16; F1 int; F2 bool }
type GenS8 struct { F0 int8; F1 int }
func genFn0(p0 uint) uint16 {
    if true {
        p0 += 948
        var v5 float32 = 3.14159
        fmt.Println(v5)
        _ = v5
    }
    var v6 string = "a"
    switch v6 {
    case "":
        var v7 int8 = 18
        _ = v7
    case "hello":
        v8 := []int{4294967295}
        _ = v8
    case "foo bar":
        fmt.Println(v6, v6)
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 uint32 = 768
    var v2 int32 = 42
    v3 := make(map[int]bool)
    var v4 int8 = int8(v2)
    if (354 > 746) {
        v9 := []int{}
        _ = v9
    }
    switch "\n\t" {
    case "\n\t":
        for fi10 := 0; fi10 < 12; fi10++ {
            if true {
                var v11 int8 = v4
                _ = v11
            }
            v3[(fi10 / ((fi10) | 1))] = true
            _ = fi10
        }
    case "foo bar":
        var v12 uint32 = (v1 - (v1 + v1))
        _ = v12
    default:
        fmt.Println(v4)
    }
    var v13 int = 32767
    switch (v13 / ((v13) | 1)) {
    case -9223372036854775808:
        var v14 int16 = int16(v13)
        _ = v14
    case -314:
        if (!true) {
            if (!false) {
                v15 := GenS1{F0: v1, F1: ("" + ("a" + "\n\t")), F2: (v4 & (v4 * v4))}
                var v16 bool = (18446744073709551615 > 42)
                _ = v15
                _ = v16
            } else {
                for fi17 := 0; fi17 < 15; fi17++ {
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v4 = v4
                        _ = fi18
                    }
                    v4 = (v4 ^ (v4 << ((337) & 63)))
                    v19 := GenS1{F0: 42, F1: "foo bar", F2: v4}
                    _ = fi17
                    _ = v19
                }
                v20 := GenS1{F0: (v1 + (v1 / ((v1) | 1))), F1: ("hello" + ("a" + ("" + ("\n\t" + ("hello" + "hello"))))), F2: 100}
                switch (v13 + v13) {
                case 295:
                    if false {
                        v21 := []int{}
                        v4 = (-(v4))
                        _ = v21
                    }
                }
                _ = v20
            }
            v2 = (^(v2))
            var v22 uint16 = 279
            _ = v22
        }
    }
    v1 = 2147483647
    var v23 int8 = -30
    var v24 uint64 = 32
    v25 := GenS2{F0: "\n\t", F1: (v24 - v24)}
    switch "a" {
    case "hello":
        v26 := GenS3{F0: 733}
        _ = v26
    case "\n\t":
        var v27 int = v13
        _ = v27
    case "":
        v28 := GenS4{F0: 0, F1: 1}
        _ = v28
    }
    if true {
        for fi29 := 0; fi29 < 17; fi29++ {
            for fi30 := 0; fi30 < 20; fi30++ {
                v31 := []int{(fi30 << ((73) & 63))}
                _ = fi30
                _ = v31
            }
            v32 := []int{(fi29 << ((uint(v24)) & 63)), (v13 ^ (fi29 * v13)), (v13 | (v13 / ((851) | 1)))}
            v24 += (v24 * uint64(v1))
            _ = fi29
            _ = v32
        }
        v33 := v3[v13]
        v34 := []int{(v13 - 79), (^(v13)), v13}
        _ = v33
        _ = v34
    }
    if (411 <= 1) {
        switch "hello" {
        case "foo bar":
            v35 := GenS2{F0: "foo bar", F1: (^(v24))}
            _ = v35
        case "a":
            delete(v3, (v13 / (((v13 << ((65535) & 63))) | 1)))
        case "\n\t":
            for fi36 := 0; fi36 < 14; fi36++ {
                v13 = 42
                v37 := GenS3{F0: 928}
                _ = fi36
                _ = v37
            }
        default:
            for fi38 := 0; fi38 < 2; fi38++ {
                var v39 int = (v13 / (((fi38 - -411)) | 1))
                var v40 int64 = -372
                var v41 uint8 = uint8(v24)
                _ = fi38
                _ = v39
                _ = v40
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 14; fi42++ {
            v2 += v2
            _ = fi42
        }
        _ = len(v3)
    }
    switch ("foo bar" + "\n\t") {
    case "\n\t":
        var v43 float64 = 0.1
        _ = v43
    case "a":
        if (32767 > 620) {
            if (false && true) {
                v23 = v23
                var v44 uint32 = v1
                fmt.Println(v23, v13)
                _ = v44
            }
        }
    case "foo bar":
        if (v4 < -97) {
            v24 = v24
            if ((-970 > -957) && true) {
                v45 := GenS5{F0: 1}
                if (genFn0(550) < genFn0(42)) {
                    var v46 uint16 = 283
                    _ = v46
                } else {
                    v47 := GenS5{F0: 2}
                    var v48 uint16 = genFn0(563)
                    fmt.Println(v4, v23, v48)
                    _ = v47
                    _ = v48
                }
                _ = v45
            } else {
                if false {
                    switch "hello" {
                    case "":
                        for fi49 := 0; fi49 < 18; fi49++ {
                            var v50 uint16 = 42
                            var v51 int8 = v4
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    case "\n\t":
                        if (2 == 960) {
                            v4 -= -1
                            for fi52 := 0; fi52 < 14; fi52++ {
                                var v53 int32 = (v2 * v2)
                                v54 := []int{(fi52 % ((int(v24)) | 1)), int(v53)}
                                var v55 uint = 42
                                _ = fi52
                                _ = v53
                                _ = v54
                                _ = v55
                            }
                        } else {
                            switch (v13 ^ 276) {
                            case 1:
                                var v56 uint16 = 255
                                _ = v56
                            default:
                                v57 := GenS2{F0: "foo bar", F1: v24}
                                _ = v57
                            }
                            v3[(-(v13))] = false
                        }
                    case "hello":
                        if (32767 != 480) {
                            var v58 uint64 = (v24 | 1)
                            _ = v58
                        } else {
                            var v59 uint64 = 2
                            switch (^(v13)) {
                            case 128:
                                if ((32767 == 722) && false) {
                                    switch "" {
                                    case "foo bar":
                                        var v60 string = "\n\t"
                                        _ = v60
                                    default:
                                        v61 := GenS6{F0: v4}
                                        _ = v61
                                    }
                                    v13 += (v13 * -9223372036854775808)
                                    fmt.Println(v24, v59)
                                }
                            default:
                                switch "foo bar" {
                                case "":
                                    var v62 float64 = float64(3.14159)
                                    _ = v62
                                case "\n\t":
                                    v24 += v24
                                case "hello":
                                    var v63 uint32 = v1
                                    _ = v63
                                }
                            }
                            _ = v59
                        }
                    default:
                        v4 = 100
                    }
                }
                var v64 int16 = -1
                _ = v64
            }
        } else {
            for fi65 := 0; fi65 < 10; fi65++ {
                var v66 bool = (-1 != (v2 ^ v2))
                var v67 uint32 = 2147483648
                fmt.Println(v13, v2)
                _ = fi65
                _ = v66
                _ = v67
            }
            if true {
                var v68 bool = (933 <= 65535)
                v69 := GenS5{F0: 1}
                _ = v68
                _ = v69
            } else {
                for fi70 := 0; fi70 < 13; fi70++ {
                    if (32767 == genFn0(uint(v24))) {
                        v71 := []int{42, (fi70 % ((1) | 1)), (-(v13)), (fi70 % ((fi70) | 1))}
                        var v72 uint32 = v1
                        v73 := GenS2{F0: "foo bar", F1: v24}
                        _ = v71
                        _ = v72
                        _ = v73
                    } else {
                        switch ("\n\t" + "a") {
                        case "a":
                            fmt.Println(v13, v2)
                        case "foo bar":
                            delete(v3, (v13 % (((^(v13))) | 1)))
                        default:
                            for fi74 := 0; fi74 < 12; fi74++ {
                                var v75 int64 = 1
                                switch 128 {
                                case 127:
                                    fmt.Println(fi70)
                                case 128:
                                    if false {
                                        var v76 float32 = (1e10 - float32(-2.25))
                                        _ = v76
                                    }
                                case 9223372036854775807:
                                    delete(v3, (fi70 ^ (-(fi74))))
                                default:
                                    if true {
                                        switch 42 {
                                        case 127:
                                            switch (("hello" + ((("\n\t" + "foo bar") + "\n\t") + "")) + ("foo bar" + "")) {
                                            case "foo bar":
                                                var v77 int32 = v2
                                                _ = v77
                                            default:
                                                fmt.Println(v1, v4, v1)
                                            }
                                        case -1:
                                            if (987 == 42) {
                                                fmt.Println(v13, v24, fi70)
                                                var v78 uint8 = 255
                                                var v79 int64 = 570
                                                _ = v78
                                                _ = v79
                                            }
                                        case -2147483648:
                                            fmt.Println(v4, v24, v23)
                                        }
                                    }
                                }
                                _ = fi74
                                _ = v75
                            }
                        }
                    }
                    for fi80 := 0; fi80 < 10; fi80++ {
                        if true {
                            switch ("" + ("hello" + "\n\t")) {
                            case "":
                                if (!((v24 + 42) <= v24)) {
                                    v81 := []int{}
                                    v82 := GenS7{F0: int16(v4), F1: fi80, F2: false}
                                    if (!true) {
                                        v23 += v23
                                    } else {
                                        if false {
                                            fmt.Println(v2)
                                        }
                                        if true {
                                            v4 -= v23
                                            v83 := GenS7{F0: 42, F1: (v13 & (v13 * v13)), F2: (42 == 1)}
                                            var v84 uint32 = v1
                                            _ = v83
                                            _ = v84
                                        }
                                        if len(v81) > 0 {
                                            v81[0] = (fi80 >> ((255) & 63))
                                        }
                                    }
                                    _ = v81
                                    _ = v82
                                }
                            case "foo bar":
                                var v85 bool = (2147483647 >= 0)
                                _ = v85
                            default:
                                v13 = v13
                            }
                            if true {
                                var v86 uint16 = 42
                                v24 += 4294967295
                                _ = v86
                            }
                            if true {
                                v13 = (fi80 << ((0) & 63))
                                v87 := []int{fi80, v13}
                                _ = v87
                            } else {
                                var v88 float64 = float64((1.5 / (-2.25 * 1e10)))
                                _ = v88
                            }
                        } else {
                            v2 = v2
                            var v89 uint = 4294967295
                            fmt.Println(v13)
                            _ = v89
                        }
                        v23 += (v4 & v23)
                        v2 -= (-(v2))
                        v90 := GenS8{F0: v4, F1: (v13 / ((128) | 1))}
                        v91 := GenS2{F0: "hello", F1: v24}
                        _ = fi80
                        _ = v90
                        _ = v91
                    }
                    fmt.Println(v23)
                    _ = fi70
                }
                v1 -= 171
            }
            v13 += v13
        }
    default:
        v92 := GenS8{F0: (v23 * 1), F1: v13}
        _ = v92
    }
    var v93 uint16 = 2
    v1 += (v1 << ((4294967295) & 63))
    switch (^(v13)) {
    case -926:
        v2 = v2
    default:
        v2 -= (v2 * 42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v13
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v93
}
