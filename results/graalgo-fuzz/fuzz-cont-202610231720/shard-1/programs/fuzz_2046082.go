package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint; F2 int32 }
func genFn0(p0 uint, p1 int16, p2 int64) int32 {
    for fi5 := 0; fi5 < 8; fi5++ {
        p2 -= int64(fi5)
        p2 = 1
        _ = fi5
    }
    v6 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -2147483648
}
func genFn1(p0 int8, p1 int64, p2 int16) uint16 {
    var v7 uint8 = 110
    switch "a" {
    case "hello":
        if ((p0 + (p0 | (p0 ^ (p0 & (p0 & 0))))) <= p0) {
            fmt.Println(p2, p0)
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                p0 = (^(p0))
                _ = fi8
            }
            var v9 int8 = (p0 % ((p0) | 1))
            p1 = (p1 % (((p1 + p1)) | 1))
            _ = v9
        }
    case "\n\t":
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(fi10, fi10)
            if true {
                p1 -= (p1 + p1)
            } else {
                fmt.Println(p2)
            }
            _ = fi10
        }
    case "":
        v11 := []int{-1, 32767, 42}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 float32 = (((((float32(0.0) - ((0.0 + 0.1) + (0.0 + 1e10))) / 3.14159) * -2.25) / 3.14159) - 0.1)
    v3 := []uint32{}
    var v4 bool = (-183 < 1000)
    if v4 {
        v12 := GenS1{F0: -591}
        for fi13 := 0; fi13 < 12; fi13++ {
            var v14 int16 = 32767
            _ = fi13
            _ = v14
        }
        if true {
            if v4 {
                v4 = v4
                fmt.Println(v4, v2)
                v1[("hello" + "")] = "foo bar"
            }
            v15 := GenS1{F0: -9223372036854775808}
            for fi16 := 0; fi16 < 4; fi16++ {
                var v17 uint8 = 127
                _ = fi16
                _ = v17
            }
            _ = v15
        }
        _ = v12
    } else {
        v4 = (!v4)
    }
    v18 := GenS2{F0: 1, F1: 1, F2: (v4 || ("\n\t" == ""))}
    var v19 float32 = v2
    var v20 uint64 = 1
    switch "hello" {
    case "hello":
        v21 := GenS1{F0: -773}
        _ = v21
    case "\n\t":
        var v22 float32 = (v19 / ((((1e10 - v19) * ((1.5 * (0.0 + 0.0)) * float32(0.0))) / v19) - v19))
        _ = v22
    default:
        if (((1.5 * 3.14159) * float64((float32((0.1 + -2.25)) * (v19 - 1.5)))) == ((3.14159 + float64(0.1)) + float64((0.0 / float32((3.14159 - float64(0.1))))))) {
            v23 := GenS3{F0: -959}
            for fi24 := 0; fi24 < 14; fi24++ {
                var v25 uint64 = v20
                for fi26 := 0; fi26 < 4; fi26++ {
                    if (!v4) {
                        if (2147483648 >= 65535) {
                            switch (fi24 / (((fi24 - (fi24 * fi26))) | 1)) {
                            case -2147483648:
                                if false {
                                    v2 = (v19 / v19)
                                    if v4 {
                                        var v27 int32 = genFn0(351, 0, int64(fi24))
                                        _ = v27
                                    } else {
                                        _ = len(v1)
                                        var v28 uint64 = v20
                                        _ = v28
                                    }
                                    v25 = v25
                                } else {
                                    if v4 {
                                        var v29 int8 = 0
                                        var v30 float32 = v2
                                        _ = v29
                                        _ = v30
                                    }
                                    v19 = v2
                                }
                            case -1:
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case 2147483647:
                                v3 = append(v3, 65535)
                            default:
                                v4 = false
                            }
                            var v31 uint64 = v25
                            fmt.Println(v25, v20)
                            _ = v31
                        } else {
                            fmt.Println(v25)
                            var v32 uint32 = 992
                            _ = v32
                        }
                        var v33 uint64 = 4294967295
                        _ = v33
                    } else {
                        v3 = append(v3, 42)
                        fmt.Println(fi26, fi26)
                    }
                    fmt.Println(v19, fi24, v2)
                    _ = len(v3)
                    v2 -= float32(((-2.25 + 0.0) / 1.5))
                    _ = fi26
                }
                v2 = (((v2 - (v19 - (1.5 + 0.0))) + (v2 * 0.0)) * v19)
                _ = fi24
                _ = v25
            }
            _ = v23
        } else {
            v4 = (!(v4 || ((66 <= 0) && (true || true))))
            if ((-2.25 + 3.14159) > (3.14159 / ((3.14159 + (1e10 - (-2.25 + -2.25))) + ((0.0 + 1.5) * (((-2.25 - 0.1) / 0.1) / 1.5))))) {
                if (42 <= -9223372036854775808) {
                    fmt.Println(v4, v2)
                } else {
                    if ((!v4) && (32767 != 0)) {
                        for fi34 := 0; fi34 < 17; fi34++ {
                            switch "foo bar" {
                            case "\n\t":
                                v35 := GenS4{F0: 699}
                                _ = v35
                            case "":
                                v36 := v1[("hello" + "a")]
                                _ = v36
                            case "foo bar":
                                v19 = v19
                            default:
                                delete(v1, "hello")
                            }
                            fmt.Println(v4, v20)
                            if (v4 && v4) {
                                v37 := GenS1{F0: 42}
                                switch 127 {
                                case -2147483648:
                                    v3 = append(v3, uint32(v20))
                                case 128:
                                    v20 = 470
                                case 42:
                                    v4 = ((float32(float64(v2)) + (v19 - 0.0)) < 1.5)
                                default:
                                    var v38 int8 = 42
                                    _ = v38
                                }
                                _ = v37
                            } else {
                                switch ("hello" + "") {
                                case "":
                                    for fi39 := 0; fi39 < 16; fi39++ {
                                        v2 -= 1e10
                                        v1["\n\t"] = ""
                                        _ = fi39
                                    }
                                case "a":
                                    switch "" {
                                    case "foo bar":
                                        v3 = append(v3, 0)
                                    }
                                }
                            }
                            _ = fi34
                        }
                        if (false && (-659 != 674)) {
                            v40 := GenS1{F0: 42}
                            if v4 {
                                var v41 int8 = 42
                                if (-2147483648 <= -2147483648) {
                                    var v42 float32 = v2
                                    var v43 float64 = float64((v19 / (((0.1 / v42) - 1e10) * v19)))
                                    _ = v42
                                    _ = v43
                                } else {
                                    v41 -= v41
                                    if ((false || ((2 > 42) || (!(1 > 128)))) && v4) {
                                        if v4 {
                                            for fi44 := 0; fi44 < 13; fi44++ {
                                                var v45 int64 = 42
                                                delete(v1, "a")
                                                _ = fi44
                                                _ = v45
                                            }
                                        } else {
                                            if (((150 < 2) || (42 >= 0)) && true) {
                                                v46 := GenS1{F0: 0}
                                                _ = v46
                                            } else {
                                                fmt.Println(v19, v4)
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v47 uint64 = v20
                                                _ = v47
                                            }
                                            var v48 int16 = 32767
                                            _ = v48
                                        }
                                    }
                                }
                                _ = len(v3)
                                _ = v41
                            }
                            _ = v40
                        } else {
                            v49 := v1[(((("hello" + (("" + ("foo bar" + "")) + (("hello" + ("hello" + "")) + "\n\t"))) + ("a" + "foo bar")) + "foo bar") + "\n\t")]
                            _ = v49
                        }
                        var v50 uint64 = (^(v20))
                        _ = v50
                    } else {
                        var v51 int32 = 42
                        var v52 int16 = 464
                        _ = v51
                        _ = v52
                    }
                    fmt.Println(v4, v4, v2)
                }
            }
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
        }
    }
    v19 -= float32(1e10)
    for fi53 := 0; fi53 < 8; fi53++ {
        v54 := GenS3{F0: -1}
        var v55 uint16 = genFn1(0, 42, 234)
        _ = fi53
        _ = v54
        _ = v55
    }
    var v56 uint = 4294967295
    var v57 uint16 = 65535
    v58 := GenS5{F0: -61}
    switch 0 {
    case -528:
        var v59 int16 = 0
        _ = v59
    default:
        switch "hello" {
        case "\n\t":
            switch -290 {
            case -294:
                var v60 int8 = 0
                _ = v60
            case 128:
                fmt.Println(v20, v56, v56)
            case 402:
                var v61 int = -809
                _ = v61
            }
        default:
            for fi62 := 0; fi62 < 7; fi62++ {
                if len(v3) > 0 {
                    v3[0] = 829
                }
                _ = fi62
            }
        }
    }
    v4 = (255 <= 0)
    switch "\n\t" {
    case "\n\t":
        v56 += (v56 << (((v56 - (v56 / ((v56) | 1)))) & 63))
    case "":
        if ("foo bar" < "\n\t") {
            switch "a" {
            case "\n\t":
                fmt.Println(v20, v19, v57)
            case "foo bar":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            default:
                fmt.Println(v20)
            }
            if v4 {
                if (((!(190 != 79)) && (1 < 4294967295)) && (984 > -1)) {
                    var v63 uint = v56
                    var v64 uint32 = 1
                    if v4 {
                        v57 = genFn1(127, -1, 32767)
                    } else {
                        if len(v3) > 0 {
                            v3[0] = v64
                        }
                    }
                    _ = v63
                    _ = v64
                } else {
                    v65 := GenS2{F0: -216, F1: -1, F2: true}
                    var v66 uint64 = v20
                    var v67 int8 = -128
                    _ = v65
                    _ = v66
                    _ = v67
                }
                v4 = v4
            } else {
                v68 := GenS6{F0: 1, F1: (v56 ^ 255), F2: -2147483648}
                v57 = v57
                fmt.Println(v4, v56)
                _ = v68
            }
            v4 = v4
        } else {
            var v69 int16 = 32767
            if len(v3) > 0 {
                v3[0] = 4294967295
            }
            _ = len(v1)
            _ = v69
        }
    }
    fmt.Println(v2, v57, v4)
    if v4 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi70 := 0; fi70 < 13; fi70++ {
            switch 1 {
            case 1:
                v56 = v56
            case -893:
                for fi71 := 0; fi71 < 7; fi71++ {
                    v20 -= v20
                    var v72 uint8 = 255
                    _ = fi71
                    _ = v72
                }
            }
            var v73 uint = v56
            _ = fi70
            _ = v73
        }
        delete(v1, "foo bar")
    } else {
        switch 9223372036854775807 {
        case 1:
            v56 = v56
        case -730:
            fmt.Println(v56, v19, v56)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v56
    _ = v57
    _ = v58
}
