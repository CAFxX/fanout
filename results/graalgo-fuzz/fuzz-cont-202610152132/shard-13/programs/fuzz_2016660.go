package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 string }
type GenS5 struct { F0 string; F1 uint16; F2 string }
func genFn0(p0 uint8) uint {
    p0 -= (p0 ^ p0)
    switch 128 {
    case 127:
        v4 := []int{255, 9223372036854775807, 202}
        _ = v4
    case 128:
        var v5 int16 = 32767
        _ = v5
    case 455:
        var v6 int8 = -54
        _ = v6
    }
    _ = p0
    return 18446744073709551615
}
func genFn1(p0 uint16, p1 uint64, p2 uint32) uint64 {
    var v7 int8 = -1
    var v8 uint = 42
    switch "foo bar" {
    case "":
        var v9 float64 = float64(((3.14159 / 3.14159) - float32(1e10)))
        _ = v9
    case "\n\t":
        for fi10 := 0; fi10 < 8; fi10++ {
            if (true || ((fi10 < (fi10 | fi10)) || true)) {
                if (219 >= p1) {
                    v11 := []int{fi10, int(p0)}
                    for fi12 := 0; fi12 < 10; fi12++ {
                        v13 := GenS1{F0: true}
                        var v14 string = (("\n\t" + "\n\t") + "\n\t")
                        fmt.Println(p1, p1, v7)
                        p2 += 65535
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    _ = v11
                } else {
                    v15 := GenS1{F0: (42 <= 32767)}
                    var v16 uint = (v8 % (((v8 * (^(v8)))) | 1))
                    _ = v15
                    _ = v16
                }
                fmt.Println(p1)
            } else {
                v17 := []int{}
                var v18 int32 = 1
                v18 += (^(v18))
                _ = v17
                _ = v18
            }
            if false {
                if (int8(p0) == (v7 - (^(v7)))) {
                    v19 := []int{2147483647, (^(fi10)), (fi10 >> (((^(v8))) & 63))}
                    fmt.Println(p2)
                    _ = v19
                } else {
                    p1 = p1
                    v20 := GenS1{F0: false}
                    _ = v20
                }
                v21 := []int{(fi10 | (fi10 >> ((v8) & 63))), int(p2), 127}
                switch ("" + "foo bar") {
                case "hello":
                    v22 := GenS1{F0: (!(42 == 1000))}
                    _ = v22
                case "a":
                    if len(v21) >= 2 {
                        v21 = v21[1:len(v21)]
                    }
                default:
                    fmt.Println(v8)
                }
                _ = v21
            }
            p2 -= 423
            _ = fi10
        }
    case "hello":
        fmt.Println(p2, p0)
    }
    switch 127 {
    case 1:
        fmt.Println(p0, v7)
    case 128:
        var v23 string = "\n\t"
        _ = v23
    case 32768:
        for fi24 := 0; fi24 < 7; fi24++ {
            for fi25 := 0; fi25 < 3; fi25++ {
                fmt.Println(p0, fi25)
                v26 := GenS2{F0: v7}
                if (((fi24 >= fi25) || true) || true) {
                    v27 := GenS3{F0: (p1 - 868)}
                    _ = v27
                }
                v28 := GenS4{F0: int64(fi25), F1: (("a" + ("" + "foo bar")) + "")}
                switch "\n\t" {
                case "\n\t":
                    v8 -= v8
                case "":
                    p1 += 456
                }
                var v29 int8 = 100
                switch (fi25 & 758) {
                case -192:
                    var v30 int64 = -9223372036854775808
                    _ = v30
                case -231:
                    v31 := []int{1}
                    _ = v31
                case 861:
                    v8 = 4294967295
                }
                _ = fi25
                _ = v26
                _ = v28
                _ = v29
            }
            var v32 int16 = 42
            _ = fi24
            _ = v32
        }
    default:
        for fi33 := 0; fi33 < 8; fi33++ {
            for fi34 := 0; fi34 < 1; fi34++ {
                v35 := GenS3{F0: (p1 & 241)}
                p2 -= 42
                _ = fi34
                _ = v35
            }
            var v36 uint32 = 2147483648
            _ = fi33
            _ = v36
        }
    }
    switch "\n\t" {
    case "foo bar":
        for fi37 := 0; fi37 < 8; fi37++ {
            if (!(v7 != v7)) {
                v38 := []int{(fi37 << (((v8 & (^(v8)))) & 63))}
                v7 = 127
                if (-38 < 57) {
                    if ((p2 % ((uint32(p0)) | 1)) < (p2 % ((p2) | 1))) {
                        v39 := GenS2{F0: v7}
                        var v40 uint64 = uint64(v8)
                        _ = v39
                        _ = v40
                    }
                }
                _ = v38
            } else {
                p2 = 356
                fmt.Println(fi37, p1, v7)
            }
            for fi41 := 0; fi41 < 7; fi41++ {
                var v42 int8 = (v7 ^ (v7 - v7))
                fmt.Println(p1)
                _ = fi41
                _ = v42
            }
            if (424 <= p1) {
                v8 = (v8 ^ v8)
            }
            _ = fi37
        }
    default:
        p0 = (p0 % (((^(p0))) | 1))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    v1 := []int{0, -9223372036854775808, 32767, -891}
    v2 := make(map[int]string)
    var v3 float64 = 1.5
    v43 := GenS2{F0: 100}
    for fi44 := 0; fi44 < 17; fi44++ {
        for fi45 := 0; fi45 < 19; fi45++ {
            if (!false) {
                v46 := GenS1{F0: (!false)}
                _ = v46
            }
            _ = fi45
        }
        fmt.Println(fi44, v3)
        _ = fi44
    }
    fmt.Println(v3)
    for fi47 := 0; fi47 < 7; fi47++ {
        var v48 float32 = 1e10
        _ = fi47
        _ = v48
    }
    if false {
        v3 += 1.5
        var v49 uint = 255
        _ = v49
    }
    fmt.Println(v3)
    v3 -= (0.0 - (v3 - 0.0))
    if (255 != 433) {
        delete(v2, -2147483648)
        v50 := v2[32768]
        _ = v50
    } else {
        if (!(false || (1e10 >= float64(0.1)))) {
            if true {
                v51 := GenS1{F0: ((1.5 > (1.5 / float32((1.5 / 1.0)))) || true)}
                _ = v51
            }
            if true {
                v3 += 1e10
            }
        }
        v3 = v3
        _ = len(v2)
    }
    switch 128 {
    case 32767:
        for fi52 := 0; fi52 < 10; fi52++ {
            var v53 int32 = -2147483648
            _ = fi52
            _ = v53
        }
    case 111:
        fmt.Println(v3)
    case 1:
        v54 := v2[-7]
        _ = v54
    }
    for fi55 := 0; fi55 < 18; fi55++ {
        for fi56 := 0; fi56 < 8; fi56++ {
            var v57 uint32 = 619
            v3 = (v3 + 1.5)
            _ = fi56
            _ = v57
        }
        var v58 float64 = v3
        if len(v1) > 0 {
            v1[0] = fi55
        }
        _ = fi55
        _ = v58
    }
    var v59 float64 = (float64(3.14159) + 0.0)
    switch -1 {
    case 127:
        switch (("" + "foo bar") + "\n\t") {
        case "":
            var v60 string = "foo bar"
            _ = v60
        case "a":
            if (!(177 < 128)) {
                var v61 float32 = 1e10
                fmt.Println(v3)
                _ = v61
            } else {
                var v62 int64 = 42
                switch -9223372036854775808 {
                case -9223372036854775808:
                    var v63 string = (("a" + "foo bar") + (("" + ("foo bar" + (("foo bar" + (((("\n\t" + "a") + "") + "a") + "a")) + "\n\t"))) + ((("" + "") + "\n\t") + ("\n\t" + "a"))))
                    _ = v63
                case 42:
                    if (2147483647 == 42) {
                        if (false || (80 != 1)) {
                            if true {
                                v64 := GenS3{F0: 2}
                                v62 -= (v62 | (v62 ^ v62))
                                _ = v64
                            } else {
                                delete(v2, -764)
                                v62 -= (-(v62))
                            }
                            v62 = (^(v62))
                            _ = len(v1)
                        } else {
                            var v65 uint8 = 138
                            _ = v65
                        }
                        v2[0] = (("a" + "") + "\n\t")
                        for fi66 := 0; fi66 < 9; fi66++ {
                            fmt.Println(v3)
                            for fi67 := 0; fi67 < 14; fi67++ {
                                _ = len(v1)
                                _ = fi67
                            }
                            if false {
                                if (((1.5 / (0.1 * 3.14159)) < (((0.1 * 1e10) + (1.5 - (0.0 - 1e10))) - -2.25)) && ((!false) && (!true))) {
                                    v68 := GenS5{F0: "hello", F1: 471, F2: ("\n\t" + "hello")}
                                    _ = v68
                                } else {
                                    v59 -= 1e10
                                    v59 += 0.0
                                    v3 = v3
                                }
                                var v69 uint16 = 228
                                var v70 uint32 = 564
                                _ = v69
                                _ = v70
                            }
                            _ = fi66
                        }
                    }
                case -1:
                    v1 = append(v1, 184)
                default:
                    for fi71 := 0; fi71 < 20; fi71++ {
                        v62 -= v62
                        _ = fi71
                    }
                }
                fmt.Println(v62, v3, v3)
                _ = v62
            }
        default:
            if (18446744073709551615 >= 955) {
                if ("hello" == ("foo bar" + "hello")) {
                    if (128 <= 127) {
                        if (1.5 != (0.0 + float64(0.0))) {
                            v2[127] = ""
                        } else {
                            var v72 int32 = 193
                            _ = v72
                        }
                        if false {
                            v3 -= float64((((1e10 - 3.14159) / 1.5) / 1.0))
                            delete(v2, 127)
                        } else {
                            _ = len(v1)
                        }
                        v3 -= ((v59 - v59) / ((-2.25 - -2.25) * ((float64(((1e10 * 0.0) / 1.5)) + v3) - 0.1)))
                    }
                    var v73 uint = 661
                    _ = v73
                }
                var v74 int8 = 100
                _ = v74
            } else {
                switch "a" {
                case "a":
                    switch "\n\t" {
                    case "hello":
                        fmt.Println(v3, v3)
                    case "":
                        var v75 int64 = -1
                        _ = v75
                    default:
                        var v76 float32 = 0.1
                        _ = v76
                    }
                default:
                    var v77 uint64 = 0
                    _ = v77
                }
                fmt.Println(v3, v59)
            }
        }
    case -159:
        var v78 string = ""
        _ = v78
    default:
        var v79 uint = 255
        _ = v79
    }
    fmt.Println(v3)
    var v80 float64 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v59:", v59)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
    _ = v59
    _ = v80
}
