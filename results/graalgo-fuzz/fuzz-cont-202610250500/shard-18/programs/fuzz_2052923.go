package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint; F2 int16 }

func main() {
    var v1 uint = 757
    v2 := make(map[int]uint32)
    var v3 uint64 = 2
    v4 := make(map[uint32]string)
    v5 := make(map[int]int)
    v6 := []int{0, 1, -289, 4294967295}
    var v7 uint8 = 128
    v5[42] = 127
    fmt.Println(v3, v7)
    v7 -= 128
    switch 127 {
    case 0:
        v1 -= v1
    case 4294967295:
        var v8 float64 = (0.0 - 3.14159)
        _ = v8
    }
    v9 := GenS1{F0: false, F1: 0}
    if (0 >= 981) {
        switch 0 {
        case -2147483648:
            fmt.Println(v3, v7, v3)
        case 4294967295:
            v7 = v7
        case 42:
            v10 := GenS2{F0: 0, F1: false}
            _ = v10
        default:
            var v11 int8 = -128
            _ = v11
        }
        v12 := GenS2{F0: 2, F1: true}
        _ = v12
    } else {
        if true {
            for fi13 := 0; fi13 < 6; fi13++ {
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS2{F0: 42, F1: false}
                    fmt.Println(v1)
                    v7 -= uint8(v3)
                    fmt.Println(fi14, v1, v1)
                    var v16 uint8 = v7
                    if (0.1 == 1.5) {
                        var v17 float32 = 0.1
                        _ = v17
                    } else {
                        var v18 uint = (v1 - (v1 - (v1 * 65535)))
                        switch "\n\t" {
                        case "foo bar":
                            if (false && ("hello" <= "\n\t")) {
                                v19 := v4[2147483647]
                                v2[(fi13 - (fi14 / ((fi14) | 1)))] = 707
                                v7 = 127
                                _ = v19
                            } else {
                                var v20 float64 = 0.0
                                v16 = (v16 - (^(v7)))
                                var v21 string = "foo bar"
                                if true {
                                    var v22 string = v21
                                    v23 := GenS3{F0: true, F1: 349}
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    _ = v22
                                    _ = v23
                                } else {
                                    if (uint32(v1) == 2147483648) {
                                        var v24 float32 = 1.5
                                        var v25 int32 = 748
                                        if true {
                                            v26 := GenS1{F0: false, F1: 1}
                                            var v27 int16 = 32767
                                            v20 += ((v20 / (v20 - 3.14159)) - (v20 * (1e10 / v20)))
                                            _ = v26
                                            _ = v27
                                        }
                                        _ = v24
                                        _ = v25
                                    } else {
                                        if (false || (uint32(v3) != 65535)) {
                                            var v28 float64 = (v20 / -2.25)
                                            _ = v28
                                        } else {
                                            v1 = 255
                                            switch ((((v21 + v21) + ((v21 + (v21 + v21)) + v21)) + v21) + "") {
                                            case "hello":
                                                v3 = v3
                                            default:
                                                switch fi13 {
                                                case 4294967295:
                                                    _ = len(v6)
                                                case 74:
                                                    v20 += 0.0
                                                default:
                                                    var v29 bool = true
                                                    _ = v29
                                                }
                                            }
                                        }
                                        v30 := GenS4{F0: v7}
                                        if (v16 > 42) {
                                            fmt.Println(v18, fi13, v18)
                                            v21 = ("\n\t" + v21)
                                        } else {
                                            v1 = v18
                                            if false {
                                                v6 = append(v6, (fi14 | fi14))
                                                v21 = v21
                                                if (uint(v7) != v1) {
                                                    v31 := GenS1{F0: ((false || (((1.5 / (1.5 / 0.1)) * 0.0) == ((3.14159 / (1e10 - (float32(((float64((0.0 - (-2.25 * 3.14159))) * v20) - v20)) / -2.25))) / ((-2.25 + (1e10 + 3.14159)) - ((((0.0 + (0.0 * 1.5)) + 3.14159) * float32(v20)) * float32(1e10)))))) && (920 == 9223372036854775807)), F1: 1000}
                                                    _ = v31
                                                }
                                            }
                                            v5[(fi13 - (fi13 - (fi13 ^ fi14)))] = (fi14 / ((fi14) | 1))
                                        }
                                        _ = v30
                                    }
                                    delete(v2, 9223372036854775807)
                                }
                                _ = v20
                                _ = v21
                            }
                            if false {
                                v7 -= (v16 << ((v18) & 63))
                            }
                        case "hello":
                            var v32 float64 = 1e10
                            _ = v32
                        }
                        _ = v18
                    }
                    if (int64(fi14) < 1) {
                        if (!(4294967295 != 2)) {
                            var v33 int16 = 0
                            fmt.Println(v1, v33, v3)
                            fmt.Println(fi13)
                            _ = v33
                        } else {
                            if true {
                                fmt.Println(v7)
                                if false {
                                    v34 := GenS1{F0: (v16 >= v16), F1: 42}
                                    var v35 int16 = 0
                                    switch "" {
                                    case "foo bar":
                                        fmt.Println(v7, v16)
                                    default:
                                        v36 := GenS1{F0: (65535 < 85), F1: v35}
                                        _ = v36
                                    }
                                    var v37 string = ((("a" + ("foo bar" + "a")) + "a") + (("foo bar" + "") + ("" + ("" + ("foo bar" + ("foo bar" + "\n\t"))))))
                                    _ = v34
                                    _ = v35
                                    _ = v37
                                }
                            } else {
                                var v38 float64 = 1.5
                                v16 = (v16 ^ 42)
                                var v39 int16 = -32768
                                _ = v38
                                _ = v39
                            }
                            delete(v4, 2)
                            v1 = 632
                        }
                        v16 = v7
                        var v40 uint8 = (v16 * (v16 & v16))
                        _ = v40
                    } else {
                        fmt.Println(fi14)
                        v41 := GenS3{F0: (!false), F1: -853}
                        if ((true && ("hello" < ("a" + "hello"))) || (((v3 + v3) == 4294967295) || false)) {
                            v42 := GenS2{F0: 0, F1: false}
                            _ = v42
                        } else {
                            v1 -= (v1 % ((656) | 1))
                            if (false || true) {
                                var v43 int32 = 42
                                if (947 <= 992) {
                                    var v44 float32 = (0.1 + (1.5 * (float32((-2.25 - 0.0)) - float32(3.14159))))
                                    v44 = (v44 * 0.0)
                                    _ = v44
                                } else {
                                    var v45 float32 = (-2.25 - 0.0)
                                    v43 += (v43 + -2147483648)
                                    _ = v45
                                }
                                if (!false) {
                                    if ((0.1 - 1e10) <= -2.25) {
                                        if len(v6) > 0 {
                                            v6[0] = (-(fi13))
                                        }
                                        switch 42 {
                                        case 128:
                                            if false {
                                                v3 = (v3 / (((v3 | 42)) | 1))
                                                fmt.Println(v3, v7, v3)
                                                fmt.Println(fi14, v43)
                                            }
                                        default:
                                            var v46 uint64 = 2
                                            _ = v46
                                        }
                                    } else {
                                        v47 := GenS5{F0: v43}
                                        fmt.Println(v16, v3, v43)
                                        _ = v47
                                    }
                                }
                                _ = v43
                            } else {
                                switch fi14 {
                                case 42:
                                    switch ((("a" + "a") + "\n\t") + ((((("\n\t" + "\n\t") + ("hello" + "a")) + "a") + "\n\t") + "\n\t")) {
                                    case "":
                                        v7 = (v7 + v16)
                                    case "foo bar":
                                        if (v3 > (v3 << (((^(v1))) & 63))) {
                                            v7 = (v16 | (v7 - (v16 % (((v16 >> (((v1 % (((^(v1))) | 1))) & 63))) | 1))))
                                            var v48 bool = false
                                            _ = v48
                                        }
                                    case "a":
                                        fmt.Println(fi14)
                                    }
                                case -705:
                                    v1 += v1
                                    var v49 int = 959
                                    _ = v49
                                case -9223372036854775808:
                                    var v50 uint32 = 42
                                    _ = v50
                                default:
                                    v51 := GenS4{F0: (v7 >> ((v1) & 63))}
                                    _ = v51
                                }
                                var v52 uint = v1
                                _ = v52
                            }
                            var v53 int64 = 42
                            _ = v53
                        }
                        _ = v41
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                var v54 int64 = 9223372036854775807
                _ = fi13
                _ = v54
            }
            v7 -= (^(v7))
            for fi55 := 0; fi55 < 11; fi55++ {
                if true {
                    switch ("hello" + ("\n\t" + "")) {
                    case "":
                        switch fi55 {
                        case 2147483647:
                            v56 := v4[42]
                            _ = v56
                        case 32768:
                            for fi57 := 0; fi57 < 11; fi57++ {
                                v3 += (v3 >> (((v1 % ((v1) | 1))) & 63))
                                v7 += v7
                                var v58 float64 = float64(((0.0 + 1.5) / 1e10))
                                _ = len(v6)
                                _ = fi57
                                _ = v58
                            }
                        }
                    case "hello":
                        var v59 int64 = -9223372036854775808
                        _ = v59
                    case "a":
                        fmt.Println(v3)
                    default:
                        for fi60 := 0; fi60 < 14; fi60++ {
                            fmt.Println(fi60, fi55, v3)
                            _ = fi60
                        }
                    }
                    v3 = (v3 >> ((v1) & 63))
                } else {
                    fmt.Println(v1)
                    v61 := GenS6{F0: -853, F1: 708, F2: int16(v1)}
                    _ = v61
                }
                fmt.Println(fi55, v1, v3)
                if true {
                    v7 -= v7
                }
                _ = fi55
            }
        }
        var v62 int16 = 32767
        v1 -= 18446744073709551615
        _ = v62
    }
    v4[uint32(v7)] = ((("\n\t" + "\n\t") + "\n\t") + "\n\t")
    v63 := GenS3{F0: ((v7 << (((^(v1))) & 63)) >= v7), F1: 1}
    var v64 uint32 = 1
    var v65 int8 = 1
    var v66 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v7
    _ = v9
}
