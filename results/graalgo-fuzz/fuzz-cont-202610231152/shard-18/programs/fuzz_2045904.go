package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS6 struct { F0 bool; F1 int32; F2 uint32 }
type GenS7 struct { F0 uint8; F1 string }
type GenS8 struct { F0 uint }
func genFn0(p0 uint64, p1 int16) bool {
    v6 := GenS1{F0: 195}
    v7 := []int{968}
    if true {
        for fi8 := 0; fi8 < 15; fi8++ {
            if (1.5 < 1e10) {
                switch fi8 {
                case -338:
                    var v9 int64 = int64(p0)
                    _ = v9
                case -528:
                    if true {
                        switch "\n\t" {
                        case "":
                            for fi10 := 0; fi10 < 10; fi10++ {
                                var v11 string = (("foo bar" + "") + ("hello" + ("\n\t" + "foo bar")))
                                switch v11 {
                                case "a":
                                    var v12 int64 = 42
                                    _ = v12
                                case "hello":
                                    v13 := GenS2{F0: (p1 >> ((65535) & 63))}
                                    _ = v13
                                }
                                switch (fi10 * fi10) {
                                case 128:
                                    var v14 uint32 = 359
                                    _ = v14
                                default:
                                    switch v11 {
                                    case "\n\t":
                                        v15 := GenS1{F0: 128}
                                        _ = v15
                                    }
                                }
                                _ = fi10
                                _ = v11
                            }
                        case "\n\t":
                            for fi16 := 0; fi16 < 7; fi16++ {
                                v17 := GenS2{F0: p1}
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                _ = fi16
                                _ = v17
                            }
                        default:
                            var v18 bool = (1 == 1)
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 15; fi19++ {
                            var v20 uint16 = 0
                            if (false || false) {
                                var v21 uint64 = 4294967295
                                v22 := GenS1{F0: 255}
                                if (true || true) {
                                    var v23 uint64 = 4294967295
                                    var v24 int16 = (p1 * p1)
                                    var v25 int32 = 100
                                    switch "" {
                                    case "a":
                                        v26 := GenS3{F0: v20}
                                        _ = v26
                                    case "":
                                        v21 = (v23 >> ((18446744073709551615) & 63))
                                    default:
                                        fmt.Println(v21)
                                    }
                                    _ = v23
                                    _ = v24
                                    _ = v25
                                }
                                _ = v21
                                _ = v22
                            }
                            p0 += p0
                            _ = fi19
                            _ = v20
                        }
                        p1 += p1
                    } else {
                        v27 := GenS4{F0: 65535}
                        if ((float32((1e10 + (float64(1.5) / 3.14159))) < 0.1) || (true || false)) {
                            var v28 float64 = 3.14159
                            _ = v28
                        } else {
                            switch (^(fi8)) {
                            case 32768:
                                var v29 int64 = -9223372036854775808
                                _ = v29
                            case 0:
                                var v30 uint16 = 42
                                _ = v30
                            }
                            if (((1.5 + float64((1e10 * (1e10 * 0.1)))) >= -2.25) || false) {
                                var v31 uint = 65535
                                var v32 float64 = (-2.25 + ((1e10 / 1.5) + (-2.25 + 1e10)))
                                _ = v31
                                _ = v32
                            }
                            var v33 int16 = p1
                            _ = v33
                        }
                        var v34 uint8 = 1
                        _ = v27
                        _ = v34
                    }
                }
                if len(v7) > 0 {
                    v7[0] = (fi8 * fi8)
                }
                for fi35 := 0; fi35 < 18; fi35++ {
                    if true {
                        switch "" {
                        case "":
                            if false {
                                switch (^(fi35)) {
                                case 1:
                                    if true {
                                        var v36 uint64 = (p0 >> ((255) & 63))
                                        fmt.Println(fi8)
                                        _ = v36
                                    } else {
                                        if false {
                                            p1 = p1
                                        } else {
                                            var v37 uint8 = 162
                                            var v38 int64 = 1
                                            switch fi8 {
                                            case -1:
                                                p1 += 0
                                            case -615:
                                                switch 128 {
                                                case -1:
                                                    var v39 int64 = (v38 << ((65535) & 63))
                                                    _ = v39
                                                case 127:
                                                    var v40 int16 = -32768
                                                    v41 := GenS5{F0: 32767, F1: 0, F2: 127}
                                                    p0 -= (p0 ^ p0)
                                                    _ = v40
                                                    _ = v41
                                                default:
                                                    switch "a" {
                                                    case "a":
                                                        if ((65535 > 2147483648) || false) {
                                                            p1 = -48
                                                        } else {
                                                            if ((-2.25 / (1e10 / 3.14159)) != -2.25) {
                                                                _ = len(v7)
                                                            }
                                                            v7 = append(v7, fi8)
                                                        }
                                                    default:
                                                        v38 -= (v38 % ((v38) | 1))
                                                    }
                                                }
                                            default:
                                                if true {
                                                    switch "foo bar" {
                                                    case "hello":
                                                        if (1 == (p0 & (p0 - 249))) {
                                                            var v42 int64 = (v38 << ((137) & 63))
                                                            _ = v42
                                                        }
                                                    case "foo bar":
                                                        var v43 float32 = 1e10
                                                        _ = v43
                                                    default:
                                                        fmt.Println(p0, p0, fi35)
                                                    }
                                                    p0 = p0
                                                } else {
                                                    p0 = p0
                                                    if len(v7) >= 2 {
                                                        v7 = v7[1:len(v7)]
                                                    }
                                                    p1 = p1
                                                    fmt.Println(p0, fi35, fi8)
                                                }
                                            }
                                            v44 := GenS1{F0: (^(v37))}
                                            _ = v37
                                            _ = v38
                                            _ = v44
                                        }
                                        p0 = p0
                                        fmt.Println(fi8)
                                        if len(v7) >= 2 {
                                            v7 = v7[1:len(v7)]
                                        }
                                        var v45 uint32 = 2
                                        v7 = append(v7, int(p1))
                                        _ = len(v7)
                                        _ = v45
                                    }
                                }
                                switch fi8 {
                                case -387:
                                    fmt.Println(fi8, fi8)
                                case -2147483648:
                                    switch (fi35 + int(p1)) {
                                    case -9223372036854775808:
                                        p1 -= 42
                                    case 32767:
                                        p1 -= (p1 / (((p1 ^ -367)) | 1))
                                    default:
                                        p1 += (p1 / (((p1 - -32768)) | 1))
                                    }
                                default:
                                    _ = len(v7)
                                }
                                switch ("hello" + "hello") {
                                case "a":
                                    var v46 bool = ((true && false) && (-92 >= 0))
                                    if len(v7) > 0 {
                                        v7[0] = fi8
                                    }
                                    p1 = 0
                                    _ = v46
                                case "foo bar":
                                    var v47 int = 4294967295
                                    _ = v47
                                case "hello":
                                    var v48 float32 = (float32(((0.1 * 1.5) + (float64((((-2.25 + 1e10) - (1.5 + (1.5 / 3.14159))) * 1e10)) + 1.5))) / (0.0 - 0.1))
                                    _ = v48
                                default:
                                    p0 = (p0 - 0)
                                }
                                var v49 uint64 = (p0 * p0)
                                p1 -= p1
                                _ = v49
                            } else {
                                if true {
                                    if len(v7) >= 2 {
                                        v7 = v7[1:len(v7)]
                                    }
                                }
                            }
                        default:
                            if true {
                                v7 = append(v7, 32768)
                            } else {
                                var v50 uint8 = 150
                                _ = v50
                            }
                        }
                        _ = len(v7)
                        var v51 bool = true
                        _ = v51
                    }
                    if (!true) {
                        p0 += 424
                    } else {
                        p1 = (p1 - 42)
                        var v52 uint16 = 32767
                        var v53 uint32 = uint32(p0)
                        _ = v52
                        _ = v53
                    }
                    v54 := GenS6{F0: (255 < 904), F1: 0, F2: 0}
                    _ = fi35
                    _ = v54
                }
            } else {
                fmt.Println(p1, p1, p1)
                p1 = (-(p1))
                for fi55 := 0; fi55 < 16; fi55++ {
                    if len(v7) > 0 {
                        v7[0] = fi8
                    }
                    fmt.Println(p1, fi8, fi8)
                    var v56 uint64 = (p0 + (p0 | 1))
                    _ = fi55
                    _ = v56
                }
            }
            _ = fi8
        }
        if (!false) {
            v57 := GenS4{F0: 42}
            _ = v57
        }
    } else {
        p1 += -1
        fmt.Println(p1, p0)
        var v58 float32 = 0.1
        _ = v58
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return (!false)
}
func genFn1(p0 bool) uint64 {
    v59 := GenS6{F0: (p0 || p0), F1: -2147483648, F2: 98}
    for fi60 := 0; fi60 < 19; fi60++ {
        p0 = (p0 && p0)
        _ = fi60
    }
    var v61 uint8 = 115
    _ = p0
    _ = v59
    _ = v61
    return 18446744073709551615
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 bool = false
    var v5 string = "a"
    if v4 {
        if false {
            v62 := v1[("" + (v5 + v5))]
            _ = v62
        }
        _ = len(v3)
    } else {
        v5 = "\n\t"
        v5 = "a"
        for fi63 := 0; fi63 < 2; fi63++ {
            fmt.Println(v5)
            v64 := []int{}
            switch fi63 {
            case 2147483647:
                for fi65 := 0; fi65 < 10; fi65++ {
                    if v4 {
                        var v66 uint = 459
                        var v67 string = v5
                        var v68 uint64 = 18446744073709551615
                        _ = v66
                        _ = v67
                        _ = v68
                    }
                    _ = fi65
                }
            case 0:
                var v69 string = ((("a" + (v5 + v5)) + (v5 + "hello")) + ("hello" + "a"))
                _ = v69
            case 127:
                for fi70 := 0; fi70 < 10; fi70++ {
                    v5 = (v5 + "foo bar")
                    if (42 < 4294967295) {
                        v4 = v4
                        if (!v4) {
                            switch v5 {
                            case "a":
                                fmt.Println(fi70)
                            case "":
                                var v71 uint16 = 42
                                _ = v71
                            case "\n\t":
                                var v72 int16 = 42
                                _ = v72
                            default:
                                fmt.Println(fi70)
                            }
                            var v73 bool = (!v4)
                            if len(v64) >= 2 {
                                v64 = v64[1:len(v64)]
                            }
                            var v74 uint8 = 1
                            _ = v73
                            _ = v74
                        }
                    } else {
                        _ = len(v1)
                        var v75 uint64 = 1
                        if len(v64) > 0 {
                            v64[0] = 0
                        }
                        _ = v75
                    }
                    _ = fi70
                }
            }
            _ = fi63
            _ = v64
        }
    }
    if v4 {
        var v76 int8 = -1
        _ = v76
    } else {
        var v77 uint16 = 1
        fmt.Println(v4)
        _ = v77
    }
    var v78 string = (("" + v5) + (((v5 + (v5 + "foo bar")) + ("a" + v5)) + v5))
    if true {
        fmt.Println(v78)
        v79 := GenS7{F0: 42, F1: "\n\t"}
        if false {
            v78 = (v5 + v5)
            switch v78 {
            case "hello":
                fmt.Println(v78, v5, v4)
            case "foo bar":
                fmt.Println(v5)
            default:
                switch v78 {
                case "":
                    if v4 {
                        if v4 {
                            _ = len(v1)
                            v78 = "foo bar"
                        } else {
                            for fi80 := 0; fi80 < 2; fi80++ {
                                v78 = v78
                                if v4 {
                                    switch fi80 {
                                    case 9223372036854775807:
                                        v78 = v78
                                    case 0:
                                        var v81 float64 = 1.5
                                        _ = v81
                                    case 412:
                                        var v82 int16 = -212
                                        _ = v82
                                    default:
                                        fmt.Println(fi80, v4)
                                    }
                                    v5 = "\n\t"
                                }
                                fmt.Println(v78, fi80, v4)
                                _ = fi80
                            }
                        }
                        fmt.Println(v78)
                    }
                }
            }
            var v83 uint16 = 472
            _ = v83
        }
        _ = v79
    } else {
        var v84 uint64 = 339
        for fi85 := 0; fi85 < 10; fi85++ {
            var v86 float64 = (1e10 * 0.0)
            _ = fi85
            _ = v86
        }
        _ = v84
    }
    v87 := GenS2{F0: -32768}
    var v88 int = -9223372036854775808
    if (0 == 2) {
        switch v88 {
        case 1:
            if (!(968 != 249)) {
                for fi89 := 0; fi89 < 16; fi89++ {
                    for fi90 := 0; fi90 < 20; fi90++ {
                        var v91 int64 = 540
                        _ = fi90
                        _ = v91
                    }
                    var v92 int32 = 1
                    var v93 string = "foo bar"
                    _ = fi89
                    _ = v92
                    _ = v93
                }
                var v94 float64 = 0.0
                v78 = v5
                _ = v94
            }
        }
    }
    v88 -= (-(v88))
    var v95 int16 = -27
    v96 := v3["hello"]
    if v4 {
        var v97 int64 = -1
        var v98 float64 = 0.1
        if (489 > 1) {
            if v4 {
                var v99 int32 = 42
                var v100 uint16 = 0
                _ = v100
                _ = v99
            }
            v101 := GenS8{F0: 25}
            _ = v101
        }
        _ = v97
        _ = v98
    } else {
        var v102 float32 = 1.5
        fmt.Println(v96)
        _ = v102
    }
    if v4 {
        var v103 int8 = int8(v88)
        _ = v103
    }
    v78 = v78
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v78:", v78)
    fmt.Println("v88:", v88)
    fmt.Println("v95:", v95)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v78
    _ = v87
    _ = v88
    _ = v95
    _ = v96
}
