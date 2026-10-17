package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int; F2 int8 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint; F1 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int8 }
type GenS11 struct { F0 uint64; F1 uint8 }
func genFn0(p0 int8) int16 {
    p0 = (p0 << ((814) & 63))
    p0 = p0
    var v7 uint32 = 4294967295
    v8 := GenS1{F0: 152}
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 int64, p1 int64, p2 int64) string {
    for fi9 := 0; fi9 < 20; fi9++ {
        p2 = (p0 - p0)
        v10 := GenS1{F0: 295}
        if (true && ((fi9 ^ fi9) <= (fi9 - 127))) {
            if ((-2.25 / 1.0) >= 0.0) {
                if true {
                    v11 := []int{fi9, 2147483647, (^(fi9))}
                    for fi12 := 0; fi12 < 13; fi12++ {
                        switch (fi9 >> ((0) & 63)) {
                        case 128:
                            p1 -= (p2 / (((p1 ^ (p1 ^ p0))) | 1))
                        case -1:
                            switch fi9 {
                            case -2147483648:
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                var v13 float64 = 3.14159
                                _ = v13
                            case 223:
                                var v14 string = (("foo bar" + "a") + ("" + ((("" + "foo bar") + "hello") + ("hello" + ""))))
                                _ = v14
                            case 0:
                                p0 = 0
                            default:
                                v11 = append(v11, (fi12 * (-(fi9))))
                            }
                        }
                        fmt.Println(p0, p0, fi12)
                        var v15 int = 127
                        if len(v11) > 0 {
                            v11[0] = (-(fi12))
                        }
                        fmt.Println(p2, fi12)
                        v11 = append(v11, v15)
                        _ = fi12
                        _ = v15
                    }
                    var v16 int = -2147483648
                    _ = v11
                    _ = v16
                }
                for fi17 := 0; fi17 < 12; fi17++ {
                    fmt.Println(fi9)
                    _ = fi17
                }
                p1 = 1
            }
            v18 := GenS2{F0: -11}
            _ = v18
        } else {
            p1 += p1
        }
        _ = fi9
        _ = v10
    }
    var v19 string = ("hello" + "hello")
    var v20 int16 = 32767
    v20 = 0
    var v21 bool = ((v19 <= "a") && (1.5 < 0.0))
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v20
    _ = v21
    return v19
}

func main() {
    var v1 uint64 = 699
    var v2 int = 9223372036854775807
    var v3 float64 = 1.5
    var v4 string = "foo bar"
    var v5 int = (v2 >> ((133) & 63))
    v6 := []uint32{2147483648, 411}
    var v22 uint32 = 4294967295
    v2 += (v5 % ((v2) | 1))
    switch "\n\t" {
    case "\n\t":
        for fi23 := 0; fi23 < 8; fi23++ {
            if (!true) {
                for fi24 := 0; fi24 < 8; fi24++ {
                    if len(v6) > 0 {
                        v6[0] = (^(v22))
                    }
                    var v25 int = fi24
                    _ = fi24
                    _ = v25
                }
            } else {
                var v26 int16 = -534
                _ = v26
            }
            _ = fi23
        }
    case "a":
        switch "hello" {
        case "hello":
            if true {
                var v27 uint8 = 0
                var v28 bool = false
                var v29 uint32 = (v22 << ((287) & 63))
                _ = v27
                _ = v28
                _ = v29
            } else {
                for fi30 := 0; fi30 < 9; fi30++ {
                    fmt.Println(v1)
                    v1 += uint64(v22)
                    _ = fi30
                }
                v31 := GenS3{F0: (v2 + (v5 * -9223372036854775808)), F1: (-(v5)), F2: 127}
                _ = v31
            }
        case "foo bar":
            var v32 string = "a"
            _ = v32
        case "a":
            v33 := GenS1{F0: 0}
            _ = v33
        default:
            switch "hello" {
            case "\n\t":
                for fi34 := 0; fi34 < 8; fi34++ {
                    v35 := GenS3{F0: 128, F1: (v5 % ((v2) | 1)), F2: -1}
                    fmt.Println(v5, v2)
                    v36 := GenS1{F0: 254}
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            case "hello":
                v37 := GenS2{F0: 0}
                _ = v37
            case "foo bar":
                v6 = append(v6, v22)
            }
        }
    }
    v22 = v22
    if false {
        var v38 uint = 65535
        if ((^(v38)) == v38) {
            switch (v2 | v5) {
            case -2147483648:
                if (255 < 1) {
                    var v39 int64 = 9223372036854775807
                    var v40 int8 = int8(v1)
                    var v41 uint64 = (v1 >> (((^(v38))) & 63))
                    _ = v39
                    _ = v40
                    _ = v41
                } else {
                    switch (v4 + ("hello" + ("foo bar" + "\n\t"))) {
                    case "a":
                        v4 = v4
                    case "hello":
                        var v42 bool = false
                        _ = v42
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
            }
            switch (v5 | (v5 & (v5 & v5))) {
            case -1:
                for fi43 := 0; fi43 < 9; fi43++ {
                    var v44 int16 = 1
                    if len(v6) > 0 {
                        v6[0] = v22
                    }
                    _ = fi43
                    _ = v44
                }
            case 2147483647:
                v5 -= (v2 * 32768)
            }
        }
        _ = v38
    } else {
        var v45 int64 = -1
        if (!true) {
            for fi46 := 0; fi46 < 8; fi46++ {
                if false {
                    v4 = v4
                    v2 += 32767
                } else {
                    switch (-(v5)) {
                    case -810:
                        v22 += 2
                    case 4:
                        switch (fi46 & v2) {
                        case -97:
                            v45 -= 0
                        case -192:
                            if (270 >= 18446744073709551615) {
                                v47 := GenS4{F0: 42, F1: ((4294967295 != 42) && (!(v22 >= 42)))}
                                _ = v47
                            }
                        }
                    case 4294967295:
                        v48 := GenS5{F0: 2, F1: 60}
                        _ = v48
                    default:
                        _ = len(v6)
                    }
                    var v49 uint32 = v22
                    _ = v49
                }
                _ = fi46
            }
            v22 = v22
            var v50 int32 = 453
            _ = v50
        } else {
            var v51 int32 = 42
            switch v5 {
            case 1:
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            case -615:
                v1 = v1
            case 32768:
                switch "hello" {
                case "a":
                    for fi52 := 0; fi52 < 10; fi52++ {
                        for fi53 := 0; fi53 < 13; fi53++ {
                            switch "hello" {
                            case "":
                                var v54 int = (fi53 & v2)
                                _ = v54
                            }
                            _ = fi53
                        }
                        _ = fi52
                    }
                case "foo bar":
                    _ = len(v6)
                case "hello":
                    v55 := GenS1{F0: 454}
                    _ = v55
                }
            default:
                switch v5 {
                case 295:
                    fmt.Println(v3, v5)
                case 1:
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                case 128:
                    for fi56 := 0; fi56 < 9; fi56++ {
                        if false {
                            if ((false || true) || false) {
                                var v57 int16 = 32767
                                _ = v57
                            } else {
                                v6 = append(v6, v22)
                                fmt.Println(v22, v51, v5)
                                for fi58 := 0; fi58 < 18; fi58++ {
                                    switch (v2 & 785) {
                                    case 0:
                                        _ = len(v6)
                                    case -272:
                                        v6 = append(v6, 65535)
                                    case 1:
                                        if false {
                                            v22 += v22
                                            var v59 int64 = v45
                                            if false {
                                                if true {
                                                    if len(v6) > 0 {
                                                        v6[0] = 284
                                                    }
                                                    var v60 uint64 = 18446744073709551615
                                                    _ = v60
                                                }
                                            } else {
                                                switch (fi58 & (fi58 << ((65535) & 63))) {
                                                case -860:
                                                    if (1000 >= 32767) {
                                                        v61 := GenS6{F0: v51}
                                                        if false {
                                                            if true {
                                                                fmt.Println(v22)
                                                                v51 = (^(v51))
                                                                var v62 uint64 = (v1 + v1)
                                                                _ = v62
                                                            }
                                                            if len(v6) >= 2 {
                                                                v6 = v6[1:len(v6)]
                                                            }
                                                            v51 = (v51 << ((1) & 63))
                                                            v22 += 69
                                                        }
                                                        switch (fi56 | (^(v5))) {
                                                        case 2147483647:
                                                            v6 = append(v6, 165)
                                                        case 20:
                                                            v3 = (v3 / float64((-2.25 * ((3.14159 - 1e10) + 1e10))))
                                                        case 42:
                                                            switch v4 {
                                                            case "\n\t":
                                                                var v63 uint64 = 4294967295
                                                                _ = v63
                                                            case "foo bar":
                                                                switch v4 {
                                                                case "foo bar":
                                                                    v64 := GenS7{F0: 127, F1: v1, F2: 32767}
                                                                    _ = v64
                                                                default:
                                                                    v65 := GenS3{F0: -427, F1: (^(v2)), F2: 81}
                                                                    _ = v65
                                                                }
                                                                var v66 int32 = int32(v45)
                                                                _ = v66
                                                            case "":
                                                                v67 := GenS3{F0: (^(fi56)), F1: v2, F2: -1}
                                                                _ = v67
                                                            default:
                                                                var v68 uint32 = (v22 | 2147483648)
                                                                _ = v68
                                                            }
                                                        }
                                                        _ = v61
                                                    } else {
                                                        if (((1.5 - v3) >= v3) || (!true)) {
                                                            if true {
                                                                v69 := GenS2{F0: 100}
                                                                switch v2 {
                                                                case 2147483647:
                                                                    v5 += (^(v5))
                                                                case -798:
                                                                    _ = len(v6)
                                                                default:
                                                                    v6 = append(v6, (^(v22)))
                                                                }
                                                                _ = v69
                                                            } else {
                                                                _ = len(v6)
                                                                var v70 uint32 = (v22 - 42)
                                                                _ = v70
                                                            }
                                                        }
                                                    }
                                                default:
                                                    v59 += -1
                                                }
                                                var v71 int = 128
                                                v72 := GenS8{F0: 411}
                                                _ = v71
                                                _ = v72
                                            }
                                            _ = v59
                                        } else {
                                            var v73 uint32 = 232
                                            v51 += (v51 * (v51 & int32(v45)))
                                            if (!(128 < 0)) {
                                                var v74 uint8 = 255
                                                var v75 uint16 = 1
                                                var v76 uint16 = (v75 | 32767)
                                                _ = v74
                                                _ = v75
                                                _ = v76
                                            }
                                            _ = v73
                                        }
                                    default:
                                        v77 := GenS3{F0: -914, F1: (-(v5)), F2: -1}
                                        _ = v77
                                    }
                                    if len(v6) > 0 {
                                        v6[0] = 2
                                    }
                                    _ = fi58
                                }
                            }
                            fmt.Println(v4)
                        } else {
                            if len(v6) > 0 {
                                v6[0] = (v22 % ((v22) | 1))
                            }
                            v78 := GenS9{F0: 18446744073709551615}
                            _ = v78
                        }
                        v79 := GenS10{F0: 127}
                        _ = fi56
                        _ = v79
                    }
                }
            }
            _ = v51
        }
        _ = v45
    }
    var v80 float32 = -2.25
    for fi81 := 0; fi81 < 13; fi81++ {
        if (((-2.25 * (v80 * v80)) - 1e10) >= 0.0) {
            switch v5 {
            case 2147483647:
                if true {
                    fmt.Println(v2, v80, v1)
                }
            }
            v82 := GenS11{F0: v1, F1: 1}
            _ = v82
        } else {
            fmt.Println(v2, v5, fi81)
            v83 := GenS8{F0: 45}
            _ = v83
        }
        switch ("a" + "") {
        case "hello":
            fmt.Println(v1, v1)
        case "":
            switch v4 {
            case "\n\t":
                for fi84 := 0; fi84 < 15; fi84++ {
                    _ = len(v6)
                    v1 = 1
                    v80 = v80
                    _ = fi84
                }
            case "":
                if (-1 != 42) {
                    var v85 uint16 = 673
                    _ = v85
                } else {
                    v80 -= (-2.25 * 1.5)
                    v6 = append(v6, (v22 / (((v22 + v22)) | 1)))
                }
            case "a":
                var v86 int8 = 100
                _ = v86
            default:
                if false {
                    var v87 float32 = (v80 - ((v80 - 0.1) + 1e10))
                    var v88 float64 = (v3 / 1.0)
                    fmt.Println(v87, v87, v22)
                    _ = v87
                    _ = v88
                } else {
                    var v89 int32 = 42
                    var v90 int64 = -369
                    _ = v89
                    _ = v90
                }
            }
        default:
            switch 128 {
            case 127:
                var v91 int = (fi81 << ((4294967295) & 63))
                _ = v91
            default:
                fmt.Println(v5, v4)
            }
        }
        _ = fi81
    }
    var v92 float64 = (((v3 - v3) / v3) - (float64(v80) - 1.5))
    var v93 uint = 255
    var v94 int16 = genFn0(0)
    v6 = append(v6, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v80:", v80)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v80
    _ = v92
    _ = v93
    _ = v94
}
