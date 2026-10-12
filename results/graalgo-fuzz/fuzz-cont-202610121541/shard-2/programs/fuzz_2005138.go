package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 bool; F1 uint32; F2 uint16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 uint64; F2 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 string; F1 uint32; F2 int32 }
type GenS12 struct { F0 uint16; F1 int64 }
func genFn0() int32 {
    var v6 uint64 = 4294967295
    var v7 int = -1
    v7 += (v7 << ((65535) & 63))
    var v8 float64 = 3.14159
    v9 := []int{}
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint16) int {
    for fi10 := 0; fi10 < 19; fi10++ {
        if (!false) {
            for fi11 := 0; fi11 < 3; fi11++ {
                fmt.Println(fi11)
                p0 = p0
                if false {
                    fmt.Println(fi10)
                    if true {
                        v12 := []int{fi10, fi10, 1, (fi10 >> ((788) & 63))}
                        var v13 uint8 = 127
                        _ = v12
                        _ = v13
                    }
                }
                _ = fi11
            }
            v14 := GenS1{F0: -2147483648, F1: p0}
            p0 -= p0
            _ = v14
        } else {
            p0 = p0
        }
        for fi15 := 0; fi15 < 19; fi15++ {
            p0 = p0
            switch (("foo bar" + "") + "\n\t") {
            case "foo bar":
                fmt.Println(fi10)
            case "a":
                switch ("\n\t" + "") {
                case "":
                    var v16 bool = (genFn0() > 0)
                    _ = v16
                }
                p0 -= p0
            case "":
                p0 = (p0 & p0)
            }
            if (1.5 <= (0.0 - 1.5)) {
                var v17 uint32 = 0
                p0 -= (p0 - p0)
                _ = v17
            } else {
                var v18 int32 = -1
                v19 := GenS1{F0: v18, F1: p0}
                var v20 int64 = -9223372036854775808
                if true {
                    if true {
                        if true {
                            v21 := GenS1{F0: genFn0(), F1: p0}
                            if true {
                                if ((v18 % (((v18 << ((65535) & 63))) | 1)) >= v18) {
                                    v22 := []int{fi10, fi15, (fi15 ^ (fi10 << ((42) & 63))), (fi10 & 9223372036854775807)}
                                    switch fi10 {
                                    case 32768:
                                        p0 += (^(p0))
                                    }
                                    _ = v22
                                }
                            }
                            v23 := []int{}
                            switch fi10 {
                            case 32768:
                                var v24 int32 = genFn0()
                                _ = v24
                            }
                            _ = v21
                            _ = v23
                        }
                    }
                }
                var v25 int32 = (v18 >> ((446) & 63))
                _ = v18
                _ = v19
                _ = v20
                _ = v25
            }
            _ = fi15
        }
        for fi26 := 0; fi26 < 20; fi26++ {
            fmt.Println(fi26)
            fmt.Println(p0)
            var v27 int16 = 1000
            _ = fi26
            _ = v27
        }
        _ = fi10
    }
    fmt.Println(p0)
    _ = p0
    return 372
}

func main() {
    v1 := []string{}
    var v2 string = ("" + (("\n\t" + ("hello" + "\n\t")) + "hello"))
    v3 := make(map[uint32]bool)
    var v4 int = -9223372036854775808
    var v5 uint16 = 0
    if false {
        if true {
            fmt.Println(v5)
            v28 := GenS1{F0: 2147483647, F1: 804}
            if len(v1) > 0 {
                v1[0] = (v2 + ("\n\t" + (((v2 + ("\n\t" + v2)) + (((v2 + ((v2 + (v2 + v2)) + v2)) + "a") + v2)) + v2)))
            }
            _ = v28
        }
        _ = len(v1)
        v29 := GenS1{F0: 2147483647, F1: v5}
        _ = v29
    } else {
        v30 := GenS2{F0: (13 <= 0), F1: 42}
        switch "a" {
        case "foo bar":
            var v31 float32 = float32(3.14159)
            _ = v31
        case "a":
            if (!false) {
                _ = len(v1)
                for fi32 := 0; fi32 < 15; fi32++ {
                    v5 -= v5
                    var v33 float64 = (float64(1e10) / 1.0)
                    _ = fi32
                    _ = v33
                }
                v34 := GenS3{F0: (42 < -471), F1: 0, F2: v5}
                _ = v34
            } else {
                if (false || (v5 >= (v5 | (v5 & 0)))) {
                    for fi35 := 0; fi35 < 6; fi35++ {
                        v5 -= (v5 >> ((974) & 63))
                        fmt.Println(v4)
                        for fi36 := 0; fi36 < 6; fi36++ {
                            switch v2 {
                            case "a":
                                v1 = append(v1, ("a" + "\n\t"))
                            case "\n\t":
                                var v37 int8 = 127
                                _ = v37
                            case "foo bar":
                                v2 = (v2 + v2)
                            default:
                                if true {
                                    var v38 uint8 = 2
                                    _ = v38
                                }
                            }
                            _ = fi36
                        }
                        _ = fi35
                    }
                } else {
                    if false {
                        fmt.Println(v4, v5, v4)
                        fmt.Println(v2)
                    }
                    var v39 uint32 = 0
                    _ = v39
                }
                v2 = (v2 + ((("\n\t" + ("a" + v2)) + "hello") + v2))
                var v40 uint8 = 128
                _ = v40
            }
        case "\n\t":
            for fi41 := 0; fi41 < 17; fi41++ {
                fmt.Println(v5, v2, v5)
                if (((float64(1.5) / 1e10) * -2.25) <= ((0.1 / 0.1) * 0.0)) {
                    v42 := GenS4{F0: (v2 + v2)}
                    if (true && false) {
                        if false {
                            v43 := GenS5{F0: 298}
                            var v44 float32 = (1.5 + 1e10)
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                            _ = v43
                            _ = v44
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            v4 = v4
                        }
                        v4 = -465
                        v45 := GenS2{F0: true, F1: 0}
                        _ = v45
                    } else {
                        v1 = append(v1, v2)
                    }
                    _ = v42
                }
                if len(v1) > 0 {
                    v1[0] = ("" + "foo bar")
                }
                _ = fi41
            }
        }
        _ = v30
    }
    if false {
        v46 := GenS6{F0: (v5 >> ((2) & 63))}
        _ = v46
    }
    for fi47 := 0; fi47 < 1; fi47++ {
        for fi48 := 0; fi48 < 17; fi48++ {
            fmt.Println(v4)
            _ = fi48
        }
        if (!(true || (-1 < genFn0()))) {
            if false {
                var v49 int8 = -80
                _ = len(v1)
                if (255 != 519) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v49
            } else {
                v50 := v3[2147483648]
                _ = v50
            }
            v51 := GenS7{F0: 0, F1: 691, F2: 2}
            _ = v51
        } else {
            if true {
                v52 := GenS4{F0: "hello"}
                if false {
                    for fi53 := 0; fi53 < 2; fi53++ {
                        if true {
                            var v54 float64 = 3.14159
                            _ = v54
                        }
                        _ = fi53
                    }
                    var v55 float64 = 1.5
                    _ = len(v3)
                    _ = v55
                } else {
                    if true {
                        var v56 uint16 = (^(v5))
                        _ = v56
                    } else {
                        if ((1 == 42) || false) {
                            if (2 < 0) {
                                if (false && true) {
                                    v5 += v5
                                    var v57 uint = 255
                                    _ = v57
                                } else {
                                    var v58 uint64 = 1
                                    v3[42] = ((!true) && ((false || false) || (42 > 1)))
                                    if len(v1) > 0 {
                                        v1[0] = ("foo bar" + (v2 + ((("a" + (("hello" + (("" + "foo bar") + v2)) + v2)) + "foo bar") + "hello")))
                                    }
                                    _ = v58
                                }
                                var v59 uint16 = v5
                                v59 -= v59
                                _ = v59
                            } else {
                                v60 := v3[669]
                                var v61 uint64 = 1
                                _ = v60
                                _ = v61
                            }
                        } else {
                            var v62 bool = false
                            v63 := GenS5{F0: v5}
                            v62 = (793 > v4)
                            _ = v62
                            _ = v63
                        }
                    }
                    var v64 uint32 = 42
                    _ = v64
                }
                _ = v52
            } else {
                if (v4 != 2147483647) {
                    v4 = 335
                }
                v65 := GenS8{F0: 4294967295}
                _ = v65
            }
        }
        _ = fi47
    }
    v5 -= (v5 * (^(v5)))
    var v66 bool = false
    if (0 >= 100) {
        v67 := GenS2{F0: true, F1: 382}
        _ = len(v3)
        v68 := v3[615]
        _ = v67
        _ = v68
    }
    fmt.Println(v66, v66)
    for fi69 := 0; fi69 < 5; fi69++ {
        var v70 int64 = -1
        var v71 float64 = 0.0
        _ = fi69
        _ = v70
        _ = v71
    }
    v5 -= 693
    fmt.Println(v4, v2, v5)
    _ = len(v1)
    v72 := GenS9{F0: v5, F1: (42 <= -32768)}
    v73 := GenS9{F0: (v5 % ((v5) | 1)), F1: (!(-128 != 1))}
    for fi74 := 0; fi74 < 13; fi74++ {
        v75 := GenS10{F0: ("\n\t" + (("" + (v2 + ("a" + v2))) + (v2 + (v2 + "hello")))), F1: 42}
        _ = len(v1)
        _ = fi74
        _ = v75
    }
    if false {
        var v76 uint32 = 762
        v77 := GenS11{F0: (v2 + (v2 + v2)), F1: (v76 << ((42) & 63)), F2: 42}
        for fi78 := 0; fi78 < 15; fi78++ {
            if v66 {
                var v79 uint8 = 2
                var v80 uint64 = 2
                var v81 float64 = 0.0
                _ = v79
                _ = v80
                _ = v81
            } else {
                if false {
                    v82 := GenS9{F0: 0, F1: ((^(v5)) > v5)}
                    _ = v82
                } else {
                    for fi83 := 0; fi83 < 14; fi83++ {
                        var v84 uint64 = 149
                        _ = fi83
                        _ = v84
                    }
                    v5 = v5
                }
            }
            if (!true) {
                fmt.Println(v4, fi78)
            }
            var v85 uint64 = 2
            _ = fi78
            _ = v85
        }
        _ = v76
        _ = v77
    }
    switch (v4 ^ (v4 ^ genFn1(32767))) {
    case 549:
        var v86 int32 = -2147483648
        _ = v86
    case 751:
        v2 = "a"
    case 9223372036854775807:
        for fi87 := 0; fi87 < 11; fi87++ {
            fmt.Println(v5, v5)
            var v88 float32 = float32(1e10)
            for fi89 := 0; fi89 < 17; fi89++ {
                v90 := GenS12{F0: (v5 % (((v5 & v5)) | 1)), F1: 284}
                _ = fi89
                _ = v90
            }
            _ = fi87
            _ = v88
        }
    default:
        if (-32768 != int16(v4)) {
            v66 = (42 > 9223372036854775807)
            v66 = (uint(v5) == 689)
        }
    }
    for fi91 := 0; fi91 < 2; fi91++ {
        v5 -= (v5 - v5)
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
        _ = fi91
    }
    if true {
        switch (-(v4)) {
        case -9223372036854775808:
            v92 := GenS3{F0: false, F1: 65535, F2: (v5 + v5)}
            _ = v92
        case 4294967295:
            v66 = (678 >= 0)
        default:
            fmt.Println(v5)
        }
    } else {
        var v93 float32 = float32((((0.0 / 1.5) - 0.0) * float64((((0.1 * 0.0) * (((1e10 * -2.25) * ((1e10 - -2.25) + 1e10)) * 0.1)) / (0.1 - ((-2.25 + 0.0) - 3.14159))))))
        v94 := GenS9{F0: 32767, F1: false}
        var v95 int16 = 1000
        _ = v93
        _ = v94
        _ = v95
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v72
    _ = v73
}
