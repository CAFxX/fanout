package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint; F1 int32; F2 bool }
type GenS4 struct { F0 int8; F1 uint16; F2 string }
type GenS5 struct { F0 int32; F1 int16; F2 int64 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint; F1 uint16 }
type GenS8 struct { F0 string; F1 uint }
type GenS9 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int64, p1 int16, p2 uint8) int32 {
    for fi5 := 0; fi5 < 5; fi5++ {
        if ("" != "hello") {
            switch fi5 {
            case -9223372036854775808:
                switch 9223372036854775807 {
                case 127:
                    if (false || false) {
                        v6 := GenS1{F0: false, F1: (p2 / ((p2) | 1)), F2: (fi5 * -879)}
                        p2 -= 2
                        p1 -= (p1 / (((p1 - -1)) | 1))
                        _ = v6
                    } else {
                        switch "foo bar" {
                        case "":
                            var v7 uint = 782
                            _ = v7
                        case "foo bar":
                            if false {
                                fmt.Println(p0, fi5)
                                p0 = 1
                            } else {
                                p0 += (^(p0))
                                v8 := []int{}
                                if len(v8) > 0 {
                                    v8[0] = 0
                                }
                                _ = v8
                            }
                        default:
                            if false {
                                for fi9 := 0; fi9 < 12; fi9++ {
                                    v10 := []int{fi9, -1, fi9, (fi5 + -9223372036854775808)}
                                    if (!false) {
                                        p1 -= 1
                                        if true {
                                            switch (fi9 ^ fi9) {
                                            case -1:
                                                fmt.Println(p2)
                                            }
                                        }
                                        p0 = 396
                                    } else {
                                        p1 += 1000
                                    }
                                    if len(v10) >= 2 {
                                        v10 = v10[1:len(v10)]
                                    }
                                    var v11 int32 = 42
                                    _ = fi9
                                    _ = v10
                                    _ = v11
                                }
                                v12 := []int{}
                                _ = v12
                            }
                        }
                        var v13 uint32 = 2
                        for fi14 := 0; fi14 < 11; fi14++ {
                            fmt.Println(v13, p2, p1)
                            v15 := GenS1{F0: false, F1: 0, F2: (-(fi14))}
                            _ = fi14
                            _ = v15
                        }
                        _ = v13
                    }
                case 0:
                    for fi16 := 0; fi16 < 18; fi16++ {
                        if ((2 != 549) || false) {
                            var v17 int32 = -1
                            switch (-(fi5)) {
                            case -1:
                                v18 := []int{(fi5 ^ (fi16 * (fi5 & (fi5 & fi16)))), (-(fi5))}
                                _ = v18
                            case -914:
                                var v19 float64 = ((0.1 / (1.5 / (float64((-2.25 - 0.1)) - 0.1))) / 0.1)
                                _ = v19
                            }
                            switch ("\n\t" + "\n\t") {
                            case "a":
                                var v20 int16 = p1
                                _ = v20
                            }
                            _ = v17
                        } else {
                            switch "hello" {
                            case "foo bar":
                                if true {
                                    v21 := GenS1{F0: true, F1: (p2 - (^(p2))), F2: 32768}
                                    _ = v21
                                } else {
                                    switch "" {
                                    case "foo bar":
                                        v22 := GenS2{F0: 42}
                                        _ = v22
                                    }
                                }
                            case "hello":
                                p2 = 2
                            }
                            var v23 bool = (p2 >= 53)
                            v23 = false
                            _ = v23
                        }
                        v24 := GenS3{F0: 255, F1: -1, F2: ("" < "\n\t")}
                        _ = fi16
                        _ = v24
                    }
                case 2147483647:
                    var v25 float64 = 1e10
                    _ = v25
                default:
                    var v26 float32 = 3.14159
                    _ = v26
                }
            case 212:
                if (true && ((p2 ^ 56) <= (^(p2)))) {
                    v27 := GenS4{F0: int8(p1), F1: 65535, F2: "hello"}
                    p2 -= 127
                    switch "a" {
                    case "\n\t":
                        p2 = p2
                    case "hello":
                        p1 = p1
                    case "a":
                        if ((true || false) || true) {
                            v28 := []int{}
                            for fi29 := 0; fi29 < 3; fi29++ {
                                var v30 uint64 = 933
                                _ = fi29
                                _ = v30
                            }
                            _ = v28
                        }
                    default:
                        fmt.Println(p2, p0, fi5)
                    }
                    _ = v27
                }
            }
        } else {
            v31 := []int{fi5, 127, (fi5 << ((4294967295) & 63))}
            var v32 uint16 = 255
            v31 = append(v31, (fi5 % ((fi5) | 1)))
            _ = v31
            _ = v32
        }
        _ = fi5
    }
    p2 = 128
    p1 = -32768
    if true {
        v33 := []int{5}
        var v34 int = -427
        v35 := GenS5{F0: -727, F1: p1, F2: (^(p0))}
        _ = v33
        _ = v34
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    var v1 int = 32768
    var v2 bool = (0 >= 42)
    v3 := []uint32{1}
    v4 := make(map[int]int)
    v36 := GenS3{F0: 1, F1: 0, F2: ((42 != 42) || v2)}
    var v37 uint32 = 2147483647
    for fi38 := 0; fi38 < 6; fi38++ {
        switch ("foo bar" + "foo bar") {
        case "hello":
            v2 = true
        }
        var v39 int8 = -75
        _ = fi38
        _ = v39
    }
    v1 += 42
    var v40 int = (v1 | (v1 % ((1) | 1)))
    if v2 {
        if len(v3) > 0 {
            v3[0] = (v37 >> ((4294967295) & 63))
        }
        for fi41 := 0; fi41 < 20; fi41++ {
            v42 := GenS6{F0: ("" + "hello")}
            _ = fi41
            _ = v42
        }
        v1 += (v40 & (^(v40)))
    } else {
        for fi43 := 0; fi43 < 18; fi43++ {
            for fi44 := 0; fi44 < 10; fi44++ {
                var v45 uint = 2
                _ = fi44
                _ = v45
            }
            v40 -= 719
            _ = fi43
        }
        var v46 uint8 = 76
        if (-349 > 495) {
            v1 -= v1
            v2 = (1 < 1)
        } else {
            if (v2 || (!(18446744073709551615 < 602))) {
                if (4294967295 == 18446744073709551615) {
                    v40 = (v40 + v1)
                    if false {
                        for fi47 := 0; fi47 < 20; fi47++ {
                            var v48 float32 = ((1e10 + float32((1.5 / 0.1))) / float32(1.5))
                            for fi49 := 0; fi49 < 6; fi49++ {
                                var v50 int = 2147483647
                                if (32767 >= 2) {
                                    switch "" {
                                    case "hello":
                                        switch v1 {
                                        case -16:
                                            fmt.Println(v40, v40)
                                        case 1:
                                            switch "a" {
                                            case "foo bar":
                                                if v2 {
                                                    if true {
                                                        fmt.Println(v40)
                                                        var v51 int = 199
                                                        switch "a" {
                                                        case "":
                                                            v52 := GenS2{F0: 1}
                                                            _ = v52
                                                        case "foo bar":
                                                            v46 += v46
                                                        default:
                                                            delete(v4, v40)
                                                        }
                                                        _ = v51
                                                    }
                                                }
                                            default:
                                                v50 -= 0
                                            }
                                        }
                                    case "foo bar":
                                        if ((v37 - (v37 * (v37 * (v37 / (((^(v37))) | 1))))) != 98) {
                                            var v53 string = ("\n\t" + ("hello" + "hello"))
                                            var v54 int32 = -896
                                            if (0 == 127) {
                                                if (100 <= int8(v37)) {
                                                    _ = len(v4)
                                                    fmt.Println(v37, v54)
                                                }
                                                if v2 {
                                                    if v2 {
                                                        if v2 {
                                                            fmt.Println(v54, v54)
                                                        } else {
                                                            var v55 uint16 = 416
                                                            v1 = int(v46)
                                                            var v56 uint8 = (v46 | 0)
                                                            v57 := GenS4{F0: -1, F1: (v55 * 146), F2: "a"}
                                                            v50 = -454
                                                            _ = v55
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                        v58 := GenS5{F0: 1, F1: 1, F2: 1}
                                                        _ = v58
                                                    }
                                                    if true {
                                                        var v59 float64 = -2.25
                                                        switch fi49 {
                                                        case -2147483648:
                                                            v60 := GenS7{F0: 4294967295, F1: 865}
                                                            _ = v60
                                                        case 32767:
                                                            if v2 {
                                                                v50 = v40
                                                            } else {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                var v61 uint8 = v46
                                                                _ = v61
                                                            }
                                                        }
                                                        v62 := GenS1{F0: (v46 != v46), F1: 31, F2: (fi47 << ((255) & 63))}
                                                        _ = v59
                                                        _ = v62
                                                    }
                                                } else {
                                                    switch ("a" + v53) {
                                                    case "":
                                                        var v63 int16 = 32767
                                                        _ = v63
                                                    case "\n\t":
                                                        var v64 uint8 = 128
                                                        _ = v64
                                                    default:
                                                        _ = len(v4)
                                                    }
                                                    var v65 uint8 = (v46 & v46)
                                                    _ = v65
                                                }
                                                v2 = true
                                            } else {
                                                v46 -= v46
                                                v1 -= fi49
                                                var v66 uint32 = (v37 >> ((0) & 63))
                                                _ = v66
                                            }
                                            _ = v53
                                            _ = v54
                                        }
                                    }
                                    var v67 int32 = 832
                                    _ = v67
                                } else {
                                    _ = len(v4)
                                }
                                v46 -= v46
                                _ = fi49
                                _ = v50
                            }
                            _ = fi47
                            _ = v48
                        }
                        delete(v4, v1)
                        v3 = append(v3, v37)
                    }
                    for fi68 := 0; fi68 < 12; fi68++ {
                        v40 -= (v40 >> ((65535) & 63))
                        v37 += (v37 << ((4294967295) & 63))
                        var v69 float32 = -2.25
                        _ = fi68
                        _ = v69
                    }
                } else {
                    if (true && ((("a" + "hello") + "a") < ("foo bar" + "foo bar"))) {
                        for fi70 := 0; fi70 < 8; fi70++ {
                            var v71 int = v1
                            _ = fi70
                            _ = v71
                        }
                        var v72 int64 = -1
                        _ = v72
                    } else {
                        fmt.Println(v37)
                        for fi73 := 0; fi73 < 5; fi73++ {
                            var v74 int16 = -733
                            fmt.Println(v1, v1, v37)
                            _ = fi73
                            _ = v74
                        }
                        if (("" + ("\n\t" + "hello")) == ("\n\t" + "")) {
                            switch -569 {
                            case 42:
                                for fi75 := 0; fi75 < 12; fi75++ {
                                    var v76 int8 = 120
                                    var v77 string = "\n\t"
                                    _ = len(v4)
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                }
                            case 0:
                                var v78 uint64 = 1
                                _ = v78
                            case -2147483648:
                                for fi79 := 0; fi79 < 6; fi79++ {
                                    fmt.Println(v1, v40)
                                    _ = fi79
                                }
                            default:
                                switch ("\n\t" + ("hello" + "hello")) {
                                case "a":
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                default:
                                    v80 := GenS8{F0: "", F1: 65535}
                                    _ = v80
                                }
                            }
                        }
                    }
                }
            } else {
                if true {
                    for fi81 := 0; fi81 < 11; fi81++ {
                        var v82 int16 = 0
                        _ = fi81
                        _ = v82
                    }
                }
                v83 := GenS4{F0: 1, F1: 32767, F2: "hello"}
                _ = v83
            }
        }
        _ = v46
    }
    v84 := GenS5{F0: genFn0(685, 155, 2), F1: -32768, F2: -323}
    v37 = 517
    if v2 {
        var v85 bool = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v85
    }
    v2 = (100 != 44)
    if v2 {
        var v86 int32 = int32(v1)
        v86 -= -655
        delete(v4, (v1 + v40))
        _ = v86
    } else {
        v87 := GenS9{F0: 695, F1: 369}
        _ = v87
    }
    if (v2 || v2) {
        fmt.Println(v2)
        var v88 float32 = (((0.1 + 0.1) * float32(0.0)) - float32(-2.25))
        _ = v88
    }
    switch ("\n\t" + "a") {
    case "hello":
        fmt.Println(v37, v1, v40)
    }
    var v89 float64 = 0.0
    v40 = (v1 ^ (v40 << ((1) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v84
    _ = v89
}
