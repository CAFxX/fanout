package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 uint8 }
type GenS7 struct { F0 uint8; F1 bool }
type GenS8 struct { F0 uint64; F1 int16 }
type GenS9 struct { F0 uint64; F1 int64; F2 uint8 }
func genFn0(p0 uint8, p1 int32) uint8 {
    fmt.Println(p1, p0)
    fmt.Println(p0, p1)
    v5 := []int{int(p0), 0, 847, -9223372036854775808}
    var v6 int = 32767
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v6, fi7)
        var v8 string = (((((("foo bar" + "") + "a") + "foo bar") + ("hello" + "\n\t")) + "foo bar") + "foo bar")
        if false {
            var v9 int8 = -128
            if len(v5) > 0 {
                v5[0] = int(p0)
            }
            if false {
                switch "" {
                case "":
                    v10 := GenS1{F0: (((v8 + v8) + "hello") + "")}
                    _ = v10
                case "hello":
                    v11 := GenS1{F0: v8}
                    _ = v11
                }
                p0 = (p0 >> ((1) & 63))
                if (255 > 255) {
                    v6 = v6
                    var v12 float64 = float64(0.0)
                    _ = v12
                }
            } else {
                for fi13 := 0; fi13 < 20; fi13++ {
                    p0 = (p0 % (((p0 + 127)) | 1))
                    _ = fi13
                }
                _ = len(v5)
            }
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 << ((709) & 63))
}

func main() {
    var v1 int8 = 127
    var v2 uint = 889
    v3 := make(map[int]string)
    v4 := []int{-930, 32767, 2147483647, 42}
    var v14 float64 = 0.0
    if true {
        for fi15 := 0; fi15 < 16; fi15++ {
            if false {
                var v16 int = fi15
                v17 := GenS1{F0: "foo bar"}
                _ = v16
                _ = v17
            } else {
                switch -437 {
                case 891:
                    var v18 uint32 = 2147483648
                    _ = v18
                case 32768:
                    var v19 float32 = (3.14159 * (1e10 - 1e10))
                    _ = v19
                default:
                    v20 := GenS1{F0: "a"}
                    _ = v20
                }
            }
            fmt.Println(v1)
            _ = fi15
        }
    } else {
        var v21 string = ""
        var v22 uint8 = 128
        _ = v21
        _ = v22
    }
    v23 := GenS2{F0: 2}
    v1 += v1
    _ = len(v4)
    var v24 uint32 = 762
    if ("" <= ("hello" + ("" + "a"))) {
        fmt.Println(v24, v2, v2)
        v2 = v2
        for fi25 := 0; fi25 < 1; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                v24 += v24
                _ = fi26
            }
            _ = len(v3)
            if false {
                if true {
                    switch "a" {
                    case "foo bar":
                        v14 = v14
                    }
                    fmt.Println(v14)
                    var v27 uint64 = 0
                    _ = v27
                } else {
                    var v28 float64 = 1e10
                    _ = len(v4)
                    _ = v28
                }
                var v29 int8 = (v1 / (((v1 - (v1 & 20))) | 1))
                fmt.Println(v29)
                _ = v29
            } else {
                v3[(fi25 >> ((v2) & 63))] = "foo bar"
                switch "" {
                case "hello":
                    for fi30 := 0; fi30 < 3; fi30++ {
                        switch fi25 {
                        case -2147483648:
                            switch fi30 {
                            case 4294967295:
                                v31 := GenS2{F0: 18446744073709551615}
                                _ = v31
                            case -2147483648:
                                if false {
                                    v1 = v1
                                } else {
                                    v2 = (v2 | v2)
                                    var v32 uint64 = 194
                                    _ = v32
                                }
                            case 144:
                                if len(v4) > 0 {
                                    v4[0] = -287
                                }
                            default:
                                fmt.Println(fi25, fi30)
                            }
                        case 9223372036854775807:
                            v24 -= 65535
                        case 42:
                            v14 = v14
                        default:
                            fmt.Println(v14, fi30)
                        }
                        switch 607 {
                        case 42:
                            if true {
                                switch (fi25 >> (((v2 & v2)) & 63)) {
                                case 32767:
                                    var v33 int = (fi30 ^ fi25)
                                    _ = v33
                                case 9223372036854775807:
                                    var v34 uint64 = 0
                                    fmt.Println(fi30)
                                    _ = v34
                                default:
                                    if true {
                                        var v35 uint32 = (v24 / ((v24) | 1))
                                        var v36 uint8 = 2
                                        switch "" {
                                        case "\n\t":
                                            var v37 uint64 = 18446744073709551615
                                            v38 := GenS1{F0: ("hello" + "a")}
                                            var v39 float64 = ((((v14 * v14) * v14) * v14) - (3.14159 - (-2.25 / v14)))
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        case "hello":
                                            v40 := GenS2{F0: 613}
                                            _ = v40
                                        default:
                                            if true {
                                                delete(v3, fi30)
                                            }
                                        }
                                        _ = v35
                                        _ = v36
                                    } else {
                                        v41 := GenS1{F0: "hello"}
                                        var v42 uint32 = v24
                                        _ = v41
                                        _ = v42
                                    }
                                }
                                v43 := GenS2{F0: 38}
                                if ((fi25 / (((fi30 + (fi30 + -719))) | 1)) != (-(fi25))) {
                                    v44 := GenS3{F0: true, F1: 681}
                                    switch ((("a" + (("\n\t" + "foo bar") + ("foo bar" + ""))) + ("\n\t" + (("\n\t" + "foo bar") + ("hello" + "a")))) + "hello") {
                                    case "\n\t":
                                        if (65535 == 19) {
                                            var v45 uint = v2
                                            if (float32(v14) <= -2.25) {
                                                if (!true) {
                                                    v46 := GenS1{F0: "a"}
                                                    var v47 int64 = 9223372036854775807
                                                    var v48 uint = (v2 | v2)
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                } else {
                                                    v45 += (v45 << (((v2 % ((v2) | 1))) & 63))
                                                    if false {
                                                        v1 = 0
                                                    } else {
                                                        v2 -= (v2 ^ 2)
                                                        switch fi25 {
                                                        case 32767:
                                                            if false {
                                                                var v49 uint8 = 1
                                                                if (!(!(42 >= 2147483647))) {
                                                                    v45 -= (v2 - v2)
                                                                }
                                                                _ = v49
                                                            } else {
                                                                fmt.Println(v2)
                                                            }
                                                            var v50 uint8 = 255
                                                            var v51 uint32 = v24
                                                            if ((!false) || false) {
                                                                var v52 string = "a"
                                                                _ = v52
                                                            } else {
                                                                fmt.Println(fi25, fi30, v14)
                                                            }
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            _ = v50
                                                            _ = v51
                                                        case 1:
                                                            if (42 == 1) {
                                                                var v53 uint64 = 2
                                                                _ = v53
                                                            } else {
                                                                if (!true) {
                                                                    fmt.Println(fi25)
                                                                    if (-341 <= 537) {
                                                                        if true {
                                                                            if len(v4) >= 2 {
                                                                                v4 = v4[1:len(v4)]
                                                                            }
                                                                            if ((v1 - v1) >= v1) {
                                                                                switch "a" {
                                                                                case "\n\t":
                                                                                    var v54 int64 = 1
                                                                                    v55 := GenS3{F0: true, F1: 0}
                                                                                    _ = v54
                                                                                    _ = v55
                                                                                case "foo bar":
                                                                                    if (2 >= 42) {
                                                                                        if false {
                                                                                            var v56 uint8 = 42
                                                                                            if (!false) {
                                                                                                fmt.Println(fi30)
                                                                                                delete(v3, -1)
                                                                                            }
                                                                                            fmt.Println(fi30, v2, v56)
                                                                                            _ = v56
                                                                                        } else {
                                                                                            if true {
                                                                                                var v57 uint = v45
                                                                                                if true {
                                                                                                    fmt.Println(fi30, fi30)
                                                                                                } else {
                                                                                                    var v58 bool = (((v14 + v14) / 1.0) == (0.0 * v14))
                                                                                                    var v59 string = "a"
                                                                                                    v60 := GenS1{F0: "a"}
                                                                                                    _ = v58
                                                                                                    _ = v59
                                                                                                    _ = v60
                                                                                                }
                                                                                                v61 := GenS4{F0: -1}
                                                                                                v45 += v2
                                                                                                var v62 int16 = 1
                                                                                                _ = v57
                                                                                                _ = v61
                                                                                                _ = v62
                                                                                            } else {
                                                                                                v63 := v3[(fi30 >> ((874) & 63))]
                                                                                                if true {
                                                                                                    fmt.Println(v45, fi30)
                                                                                                    switch (fi25 - (fi30 | (fi25 << (((v2 + (v2 * (v45 + (v45 >> ((18446744073709551615) & 63)))))) & 63)))) {
                                                                                                    case 42:
                                                                                                        v24 = v24
                                                                                                    case 9223372036854775807:
                                                                                                        var v64 int64 = -1
                                                                                                        _ = v64
                                                                                                    default:
                                                                                                        var v65 int32 = 2147483647
                                                                                                        _ = v65
                                                                                                    }
                                                                                                } else {
                                                                                                    v24 -= (v24 ^ (v24 * 4294967295))
                                                                                                    v1 += (^(v1))
                                                                                                }
                                                                                                _ = v63
                                                                                            }
                                                                                            if (0.0 >= ((v14 * v14) * v14)) {
                                                                                                v2 += v2
                                                                                                v1 = -38
                                                                                            }
                                                                                        }
                                                                                        v14 = v14
                                                                                    } else {
                                                                                        v66 := GenS4{F0: -901}
                                                                                        _ = v66
                                                                                    }
                                                                                case "hello":
                                                                                    v67 := GenS1{F0: "foo bar"}
                                                                                    _ = v67
                                                                                }
                                                                                fmt.Println(v2, v2)
                                                                                var v68 int64 = 0
                                                                                _ = v68
                                                                            } else {
                                                                                var v69 int = 42
                                                                                v14 = v14
                                                                                if true {
                                                                                    v69 -= v69
                                                                                    v24 -= v24
                                                                                }
                                                                                _ = v69
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case 995:
                                                            v2 += v2
                                                        default:
                                                            var v70 int64 = 42
                                                            v45 = 42
                                                            var v71 float32 = ((float32(0.1) / float32(v14)) * (1e10 * (-2.25 / float32((((float64(1.5) / 3.14159) + v14) - ((((3.14159 / 1.5) + (v14 - ((v14 / v14) + (3.14159 / 1e10)))) * (float64((-2.25 - 1.5)) - v14)) / (v14 * float64(1e10))))))))
                                                            switch 130 {
                                                            case 128:
                                                                var v72 uint16 = uint16(v45)
                                                                _ = v72
                                                            case -1:
                                                                var v73 float64 = 1e10
                                                                _ = v73
                                                            case 2147483647:
                                                                fmt.Println(fi30)
                                                                v70 += (v70 & v70)
                                                                fmt.Println(v71)
                                                            default:
                                                                fmt.Println(fi30, fi25, v71)
                                                            }
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        switch "\n\t" {
                                                        case "":
                                                            var v74 uint8 = 128
                                                            _ = v74
                                                        default:
                                                            v75 := v3[fi25]
                                                            _ = v75
                                                        }
                                                        var v76 bool = true
                                                        if v76 {
                                                            v77 := GenS5{F0: ((v24 >> (((v2 % ((522) | 1))) & 63)) < 2147483648)}
                                                            v78 := GenS6{F0: 0, F1: 0}
                                                            v14 += (float64(float32(v14)) + (1e10 - 1.5))
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            v45 += uint(v24)
                                                        }
                                                        _ = v76
                                                    }
                                                }
                                            } else {
                                                var v79 int16 = -32768
                                                var v80 float32 = -2.25
                                                var v81 uint16 = 2
                                                _ = v79
                                                _ = v80
                                                _ = v81
                                            }
                                            fmt.Println(v14, v2)
                                            _ = v45
                                        } else {
                                            if ((fi30 << ((v2) & 63)) <= int(v1)) {
                                                var v82 int64 = -1
                                                var v83 int16 = 32767
                                                if false {
                                                    v2 -= (v2 - v2)
                                                    var v84 int = fi30
                                                    v14 += v14
                                                    _ = v84
                                                }
                                                _ = v82
                                                _ = v83
                                            }
                                        }
                                    case "hello":
                                        if len(v4) > 0 {
                                            v4[0] = (fi25 * fi30)
                                        }
                                    default:
                                        switch (^(fi25)) {
                                        case 9223372036854775807:
                                            v24 += v24
                                        case -98:
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch "" {
                                                case "a":
                                                    var v85 uint64 = 199
                                                    fmt.Println(v24)
                                                    v24 = v24
                                                    var v86 string = ("" + "")
                                                    _ = v85
                                                    _ = v86
                                                default:
                                                    _ = len(v3)
                                                }
                                            case "a":
                                                v1 = int8(v2)
                                            case "":
                                                v24 -= v24
                                            default:
                                                v87 := GenS3{F0: false, F1: 898}
                                                _ = v87
                                            }
                                        default:
                                            if true {
                                                v24 -= (v24 / ((2147483647) | 1))
                                                switch "a" {
                                                case "":
                                                    if len(v4) > 0 {
                                                        v4[0] = (-(fi25))
                                                    }
                                                }
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = -160
                                                }
                                                var v88 uint16 = 42
                                                v14 += v14
                                                _ = v88
                                            }
                                        }
                                    }
                                    _ = v44
                                }
                                _ = v43
                            }
                        }
                        _ = len(v3)
                        _ = fi30
                    }
                case "":
                    if ((v24 / (((v24 % (((v24 % ((2) | 1))) | 1))) | 1)) != 2) {
                        v2 = (v2 & (v2 << ((v2) & 63)))
                    }
                }
                v89 := GenS1{F0: (("a" + "foo bar") + "foo bar")}
                _ = v89
            }
            _ = fi25
        }
    } else {
        var v90 int32 = -644
        var v91 float32 = (1.5 + 3.14159)
        delete(v3, 128)
        _ = v90
        _ = v91
    }
    var v92 uint = (v2 * (^(v2)))
    for fi93 := 0; fi93 < 5; fi93++ {
        if (!false) {
            var v94 int16 = 1000
            v94 = v94
            _ = v94
        }
        var v95 float64 = v14
        _ = fi93
        _ = v95
    }
    for fi96 := 0; fi96 < 2; fi96++ {
        if (191 != 32767) {
            switch (fi96 >> (((v2 % ((v2) | 1))) & 63)) {
            case 89:
                v97 := GenS7{F0: 127, F1: true}
                _ = v97
            case 42:
                v98 := GenS8{F0: 18446744073709551615, F1: 1}
                _ = v98
            case 2147483647:
                v99 := GenS7{F0: 201, F1: false}
                _ = v99
            default:
                if (32767 >= 255) {
                    for fi100 := 0; fi100 < 10; fi100++ {
                        if true {
                            v2 -= (v92 >> ((42) & 63))
                            var v101 uint32 = 42
                            _ = v101
                        }
                        v102 := GenS9{F0: 4294967295, F1: -1, F2: 42}
                        fmt.Println(fi96)
                        _ = fi100
                        _ = v102
                    }
                } else {
                    fmt.Println(fi96, v14)
                    var v103 int16 = 42
                    v24 = (v24 * (v24 + 982))
                    _ = v103
                }
            }
        } else {
            var v104 uint8 = 255
            _ = v104
        }
        for fi105 := 0; fi105 < 4; fi105++ {
            var v106 int16 = 32767
            fmt.Println(v1)
            _ = fi105
            _ = v106
        }
        for fi107 := 0; fi107 < 19; fi107++ {
            v108 := v3[(^(fi107))]
            _ = fi107
            _ = v108
        }
        _ = fi96
    }
    fmt.Println(v92)
    v92 = 2
    v109 := GenS2{F0: 18446744073709551615}
    var v110 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v110:", v110)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v92:", v92)
    _ = v1
    _ = v109
    _ = v110
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v92
}
