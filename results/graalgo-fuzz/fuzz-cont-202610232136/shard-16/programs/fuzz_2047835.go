package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 bool; F1 int32; F2 uint64 }
type GenS3 struct { F0 int8; F1 int32 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 int32; F1 int16; F2 uint8 }
type GenS6 struct { F0 uint32; F1 uint8; F2 uint }
type GenS7 struct { F0 bool; F1 string; F2 int64 }
type GenS8 struct { F0 int64; F1 uint64 }
type GenS9 struct { F0 int64; F1 bool; F2 bool }
func genFn0(p0 bool, p1 int16) uint32 {
    var v4 uint64 = 42
    for fi5 := 0; fi5 < 11; fi5++ {
        var v6 int = fi5
        if p0 {
            p1 = (p1 - (p1 >> ((uint(v4)) & 63)))
            var v7 uint16 = 65535
            v8 := GenS1{F0: 861, F1: int64(fi5), F2: 949}
            _ = v7
            _ = v8
        } else {
            if ((239 > uint32(v4)) && true) {
                var v9 int16 = (^(p1))
                p0 = true
                v10 := GenS1{F0: -2147483648, F1: 1, F2: v4}
                _ = v10
                _ = v9
            } else {
                fmt.Println(v4, fi5)
            }
        }
        _ = fi5
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}
func genFn1(p0 bool, p1 int16, p2 int) int8 {
    var v11 uint64 = 1
    if p0 {
        fmt.Println(p2, p1)
        switch "" {
        case "":
            v12 := GenS1{F0: -1, F1: -658, F2: 512}
            _ = v12
        case "a":
            switch p2 {
            case 127:
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint = 0
                    v14 -= (v14 * v14)
                    _ = fi13
                    _ = v14
                }
            }
        }
    }
    v15 := []int{-1, p2, (p2 * p2), p2}
    p1 = (^(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v15
    return 0
}

func main() {
    var v1 uint16 = 301
    var v2 int64 = -1
    v3 := make(map[string]bool)
    v1 -= 212
    fmt.Println(v2, v2)
    switch "a" {
    case "":
        for fi16 := 0; fi16 < 1; fi16++ {
            fmt.Println(fi16)
            _ = fi16
        }
    case "hello":
        switch "hello" {
        case "\n\t":
            if false {
                if false {
                    fmt.Println(v1, v2)
                    for fi17 := 0; fi17 < 8; fi17++ {
                        if (((((((float64(0.1) - float64(3.14159)) + -2.25) * (1e10 + 1e10)) == 1e10) || true) && ((!(false || (806 == 0))) || false)) && false) {
                            if true {
                                if false {
                                    if (!true) {
                                        for fi18 := 0; fi18 < 15; fi18++ {
                                            switch (fi18 / (((fi18 - fi18)) | 1)) {
                                            case -121:
                                                v19 := []int{}
                                                _ = v19
                                            case -484:
                                                var v20 uint8 = 213
                                                _ = v20
                                            case 4294967295:
                                                v2 = v2
                                            default:
                                                v1 = v1
                                            }
                                            v21 := GenS1{F0: 42, F1: (v2 + v2), F2: 18446744073709551615}
                                            _ = fi18
                                            _ = v21
                                        }
                                        for fi22 := 0; fi22 < 19; fi22++ {
                                            if (!false) {
                                                var v23 int = (fi22 ^ 0)
                                                _ = v23
                                            } else {
                                                v24 := GenS2{F0: true, F1: 1, F2: 958}
                                                v2 -= (v2 - (^(v2)))
                                                v1 = v1
                                                var v25 uint8 = 123
                                                fmt.Println(fi22, fi17)
                                                _ = v24
                                                _ = v25
                                            }
                                            if false {
                                                v26 := []int{-1}
                                                _ = v26
                                            } else {
                                                v27 := []int{fi22, fi17, 32768, (fi22 % (((fi22 | (fi17 / ((4294967295) | 1)))) | 1))}
                                                v1 -= v1
                                                v27 = append(v27, fi22)
                                                var v28 float64 = 0.1
                                                var v29 float64 = (float64((0.0 + 1.5)) * float64(((3.14159 / (1.5 * 3.14159)) + 3.14159)))
                                                _ = v27
                                                _ = v28
                                                _ = v29
                                            }
                                            _ = fi22
                                        }
                                        v1 = (v1 | 1)
                                    } else {
                                        v30 := GenS2{F0: false, F1: 2147483647, F2: 522}
                                        if ((1000 <= 1000) && (122 < 42)) {
                                            var v31 int8 = 0
                                            for fi32 := 0; fi32 < 1; fi32++ {
                                                fmt.Println(v2, v1, fi32)
                                                var v33 uint8 = 255
                                                v34 := GenS3{F0: (v31 >> ((285) & 63)), F1: -858}
                                                _ = fi32
                                                _ = v33
                                                _ = v34
                                            }
                                            _ = v31
                                        }
                                        v35 := GenS2{F0: (0.0 != ((1e10 + 3.14159) * 3.14159)), F1: 42, F2: 42}
                                        _ = v30
                                        _ = v35
                                    }
                                    switch (fi17 % ((fi17) | 1)) {
                                    case -741:
                                        v36 := GenS3{F0: 127, F1: 0}
                                        _ = v36
                                    case 479:
                                        var v37 string = (("a" + "foo bar") + "foo bar")
                                        _ = v37
                                    case 128:
                                        fmt.Println(v2, fi17)
                                    }
                                    v38 := v3[((("a" + (("\n\t" + "\n\t") + ("foo bar" + "\n\t"))) + "foo bar") + ("foo bar" + "\n\t"))]
                                    _ = v38
                                } else {
                                    _ = len(v3)
                                }
                                var v39 int32 = 178
                                _ = v39
                            }
                            var v40 uint32 = 267
                            _ = v40
                        } else {
                            for fi41 := 0; fi41 < 7; fi41++ {
                                if (false || true) {
                                    switch fi41 {
                                    case 186:
                                        var v42 int32 = -2147483648
                                        _ = v42
                                    case -959:
                                        v43 := GenS2{F0: (!false), F1: int32(fi17), F2: 4294967295}
                                        _ = v43
                                    case 32768:
                                        v44 := []int{int(v1), fi41, (fi41 ^ 42)}
                                        _ = v44
                                    default:
                                        var v45 int32 = 1
                                        _ = v45
                                    }
                                }
                                if ((v2 ^ (^(v2))) >= v2) {
                                    if false {
                                        v46 := GenS1{F0: 1, F1: (^(v2)), F2: 10}
                                        if (3.14159 <= float64((0.1 / 3.14159))) {
                                            var v47 float64 = (1.5 / 1.0)
                                            var v48 int64 = v2
                                            _ = len(v3)
                                            _ = v47
                                            _ = v48
                                        } else {
                                            var v49 int16 = 42
                                            var v50 int64 = v2
                                            if (((float32(float64(0.1)) / 3.14159) / (3.14159 * ((1e10 / 1e10) + -2.25))) >= 0.1) {
                                                if ((((((3.14159 + 1.5) / 1.0) / 0.1) - 1e10) * 3.14159) > 0.1) {
                                                    v50 = (v2 % (((v2 * (v50 + 9223372036854775807))) | 1))
                                                    switch ("hello" + "foo bar") {
                                                    case "foo bar":
                                                        if true {
                                                            var v51 string = "a"
                                                            v50 = (v2 << ((65535) & 63))
                                                            _ = v51
                                                        } else {
                                                            var v52 uint8 = 255
                                                            var v53 int64 = v2
                                                            _ = v52
                                                            _ = v53
                                                        }
                                                    default:
                                                        v54 := []int{(^(fi41)), fi41, (fi41 - (fi41 | fi17)), (fi41 ^ (fi41 - (-(fi17))))}
                                                        _ = v54
                                                    }
                                                } else {
                                                    v55 := GenS4{F0: -556, F1: -1}
                                                    _ = v55
                                                }
                                                var v56 uint8 = 1
                                                if (!("" < ("\n\t" + ""))) {
                                                    if false {
                                                        fmt.Println(v56)
                                                    } else {
                                                        fmt.Println(v50)
                                                    }
                                                    v57 := GenS3{F0: -1, F1: 0}
                                                    _ = v57
                                                }
                                                _ = v56
                                            } else {
                                                v49 += (^(v49))
                                                v1 = (v1 * (v1 | v1))
                                            }
                                            _ = v49
                                            _ = v50
                                        }
                                        _ = v46
                                    } else {
                                        if true {
                                            var v58 int32 = int32(fi17)
                                            v59 := GenS3{F0: genFn1(true, 42, (^(fi17))), F1: v58}
                                            var v60 uint = 42
                                            _ = v58
                                            _ = v59
                                            _ = v60
                                        }
                                        v1 = 42
                                    }
                                    fmt.Println(v2)
                                    if true {
                                        switch "\n\t" {
                                        case "hello":
                                            var v61 float64 = (0.1 * (-2.25 + float64(3.14159)))
                                            _ = v61
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                v62 := GenS5{F0: -2147483648, F1: -32768, F2: 127}
                                                _ = v62
                                            }
                                        }
                                        fmt.Println(fi41, fi17)
                                        if ((!true) || ((4294967295 > 2) && true)) {
                                            if false {
                                                v63 := []int{(-(fi17)), (fi41 % ((32768) | 1)), fi17}
                                                var v64 bool = (true && true)
                                                if false {
                                                    fmt.Println(v64, v64)
                                                }
                                                _ = v63
                                                _ = v64
                                            } else {
                                                _ = len(v3)
                                                switch "a" {
                                                case "\n\t":
                                                    v65 := v3["\n\t"]
                                                    _ = v65
                                                case "hello":
                                                    v66 := []int{580, fi17, int(v1), 143}
                                                    _ = v66
                                                default:
                                                    if true {
                                                        v67 := GenS5{F0: 0, F1: 42, F2: 2}
                                                        _ = v67
                                                    } else {
                                                        _ = len(v3)
                                                    }
                                                }
                                            }
                                        }
                                    } else {
                                        switch ("hello" + "foo bar") {
                                        case "hello":
                                            v68 := GenS1{F0: 2147483647, F1: (v2 / ((9223372036854775807) | 1)), F2: 2}
                                            _ = v68
                                        }
                                        if false {
                                            fmt.Println(v2, v2)
                                        }
                                        v2 += 0
                                        v2 = (v2 ^ v2)
                                        v69 := GenS6{F0: 0, F1: 227, F2: 2}
                                        _ = v69
                                    }
                                } else {
                                    v70 := GenS2{F0: (-135 <= 42), F1: 888, F2: 257}
                                    v71 := GenS7{F0: (!false), F1: (("" + ("a" + "a")) + "a"), F2: v2}
                                    var v72 uint8 = 75
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                                _ = fi41
                            }
                        }
                        if true {
                            for fi73 := 0; fi73 < 20; fi73++ {
                                v1 = (v1 + v1)
                                v74 := v3["\n\t"]
                                _ = fi73
                                _ = v74
                            }
                            fmt.Println(fi17)
                            var v75 uint = 255
                            _ = v75
                        }
                        _ = fi17
                    }
                } else {
                    switch 2147483647 {
                    case 1:
                        switch ((("hello" + "hello") + (("foo bar" + "foo bar") + "a")) + ("hello" + "foo bar")) {
                        case "foo bar":
                            var v76 int16 = -659
                            _ = v76
                        case "":
                            v1 += v1
                        case "hello":
                            v2 -= v2
                        default:
                            for fi77 := 0; fi77 < 7; fi77++ {
                                for fi78 := 0; fi78 < 20; fi78++ {
                                    v79 := GenS5{F0: -2147483648, F1: int16(v1), F2: 42}
                                    _ = fi78
                                    _ = v79
                                }
                                v80 := GenS6{F0: 65535, F1: 255, F2: 255}
                                _ = fi77
                                _ = v80
                            }
                        }
                    }
                    v2 = (v2 & v2)
                }
                if true {
                    if ((((!((-1 == 42) || false)) || (!(((2 != 65535) && (348 == 65535)) && (false && (("a" + (("\n\t" + ("hello" + "foo bar")) + "")) == ("\n\t" + "\n\t")))))) && true) || (1 < -164)) {
                        v2 = (v2 >> ((466) & 63))
                        fmt.Println(v2, v2)
                        var v81 uint = 255
                        _ = v81
                    } else {
                        fmt.Println(v2, v2)
                    }
                }
                for fi82 := 0; fi82 < 11; fi82++ {
                    v83 := []int{(fi82 % (((fi82 - 2147483647)) | 1)), (fi82 ^ (fi82 >> ((42) & 63)))}
                    v2 = (v2 / ((v2) | 1))
                    _ = fi82
                    _ = v83
                }
            } else {
                var v84 bool = true
                v85 := []int{}
                _ = v84
                _ = v85
            }
        case "":
            var v86 int64 = (v2 & v2)
            _ = v86
        case "foo bar":
            var v87 uint8 = 128
            _ = v87
        }
    }
    for fi88 := 0; fi88 < 6; fi88++ {
        var v89 uint8 = 1
        var v90 bool = (255 >= 919)
        v91 := []int{(fi88 << ((2) & 63)), (-(fi88)), fi88}
        _ = fi88
        _ = v89
        _ = v90
        _ = v91
    }
    var v92 uint32 = 42
    for fi93 := 0; fi93 < 8; fi93++ {
        for fi94 := 0; fi94 < 17; fi94++ {
            if true {
                v95 := []int{}
                var v96 int16 = 1
                _ = v95
                _ = v96
            }
            var v97 uint = 2
            _ = fi94
            _ = v97
        }
        v3["hello"] = true
        for fi98 := 0; fi98 < 18; fi98++ {
            v3["\n\t"] = (0 >= 2)
            var v99 uint64 = 113
            _ = fi98
            _ = v99
        }
        _ = fi93
    }
    v100 := GenS5{F0: 639, F1: 0, F2: 255}
    v1 = v1
    v1 -= v1
    if true {
        switch "" {
        case "":
            v101 := GenS8{F0: (v2 ^ 466), F1: 2}
            _ = v101
        case "\n\t":
            fmt.Println(v2, v2, v1)
        case "a":
            var v102 uint32 = (v92 * (v92 << ((192) & 63)))
            _ = v102
        default:
            v1 += (v1 * v1)
        }
        for fi103 := 0; fi103 < 20; fi103++ {
            for fi104 := 0; fi104 < 1; fi104++ {
                v105 := GenS9{F0: (v2 + (v2 >> ((2) & 63))), F1: true, F2: true}
                delete(v3, ("foo bar" + "foo bar"))
                _ = fi104
                _ = v105
            }
            if (182 == 0) {
                if (4294967295 != 255) {
                    v2 = 1
                    switch (("a" + "hello") + "a") {
                    case "\n\t":
                        fmt.Println(v92, v92)
                    default:
                        v1 += (v1 >> ((1) & 63))
                    }
                    for fi106 := 0; fi106 < 18; fi106++ {
                        var v107 uint64 = 18446744073709551615
                        v108 := GenS1{F0: 42, F1: 0, F2: (v107 / ((v107) | 1))}
                        _ = fi106
                        _ = v107
                        _ = v108
                    }
                }
                v2 += (v2 & v2)
            }
            _ = fi103
        }
        var v109 int32 = 42
        _ = v109
    }
    v92 = genFn0(false, 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v92:", v92)
    _ = v1
    _ = v100
    _ = v2
    _ = v3
    _ = v92
}
