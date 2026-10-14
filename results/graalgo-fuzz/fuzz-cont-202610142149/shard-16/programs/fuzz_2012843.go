package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int32; F2 bool }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 int8; F2 uint }
type GenS6 struct { F0 int; F1 uint64; F2 uint }
type GenS7 struct { F0 uint32; F1 int64; F2 int64 }

func main() {
    var v1 int8 = -128
    v2 := make(map[int]uint32)
    v3 := make(map[string]uint32)
    var v4 string = "\n\t"
    v4 = v4
    for fi5 := 0; fi5 < 18; fi5++ {
        var v6 uint = 255
        fmt.Println(v6)
        _ = fi5
        _ = v6
    }
    var v7 uint8 = 42
    if true {
        if false {
            _ = len(v3)
            switch -445 {
            case -784:
                switch -1 {
                case 1:
                    if true {
                        v8 := GenS1{F0: v1, F1: 42, F2: v7}
                        _ = v8
                    }
                case 432:
                    v9 := v2[1]
                    _ = v9
                case 228:
                    v10 := GenS2{F0: (v7 * v7), F1: 113, F2: true}
                    _ = v10
                }
            case 4294967295:
                v11 := GenS1{F0: 1, F1: 1, F2: (^(v7))}
                _ = v11
            default:
                v12 := GenS3{F0: v1}
                _ = v12
            }
        } else {
            v4 = ""
            for fi13 := 0; fi13 < 15; fi13++ {
                v1 -= (v1 << ((4294967295) & 63))
                v14 := GenS4{F0: -787}
                var v15 uint32 = 2
                _ = fi13
                _ = v14
                _ = v15
            }
            if (!true) {
                for fi16 := 0; fi16 < 8; fi16++ {
                    if (42 <= 4294967295) {
                        v17 := []int{32767, -9223372036854775808}
                        for fi18 := 0; fi18 < 18; fi18++ {
                            v19 := v3["a"]
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    } else {
                        if true {
                            if (((0.0 + 1.5) / 1e10) != float32(1.5)) {
                                if false {
                                    var v20 uint8 = (v7 >> ((255) & 63))
                                    _ = v20
                                }
                                for fi21 := 0; fi21 < 3; fi21++ {
                                    var v22 float32 = (0.0 / 1.0)
                                    _ = fi21
                                    _ = v22
                                }
                            } else {
                                var v23 uint16 = 698
                                if (false || (255 == v7)) {
                                    for fi24 := 0; fi24 < 1; fi24++ {
                                        if false {
                                            fmt.Println(v1, v1)
                                            if (((5 != int8(v23)) || (!(-1 >= 42))) || true) {
                                                var v25 uint = 2
                                                v23 = 2
                                                _ = len(v2)
                                                _ = v25
                                            } else {
                                                var v26 uint64 = 18446744073709551615
                                                v26 = (v26 - (^(v26)))
                                                v1 = 42
                                                v27 := GenS3{F0: (v1 - v1)}
                                                _ = v26
                                                _ = v27
                                            }
                                            var v28 int32 = -2147483648
                                            _ = v28
                                        } else {
                                            if ((false && true) || false) {
                                                switch fi24 {
                                                case 42:
                                                    switch fi16 {
                                                    case -484:
                                                        v29 := GenS2{F0: 2, F1: -707, F2: false}
                                                        _ = v29
                                                    default:
                                                        switch v4 {
                                                        case "a":
                                                            switch "\n\t" {
                                                            case "hello":
                                                                fmt.Println(v7, fi16, v1)
                                                            case "\n\t":
                                                                if false {
                                                                    v30 := GenS5{F0: 2, F1: (v1 >> ((2) & 63)), F2: 0}
                                                                    _ = len(v3)
                                                                    var v31 float64 = 1.5
                                                                    _ = v30
                                                                    _ = v31
                                                                } else {
                                                                    switch -249 {
                                                                    case 127:
                                                                        if false {
                                                                            delete(v3, ("hello" + v4))
                                                                            v4 = "\n\t"
                                                                        } else {
                                                                            v32 := []int{(fi16 - fi16), (fi16 * 111), (fi24 >> ((331) & 63)), 0}
                                                                            v33 := GenS6{F0: fi24, F1: 905, F2: 893}
                                                                            switch int(v7) {
                                                                            case 32767:
                                                                                v23 = (v23 | v23)
                                                                            case 127:
                                                                                fmt.Println(v23)
                                                                            case 4294967295:
                                                                                var v34 float32 = 1e10
                                                                                _ = v34
                                                                            default:
                                                                                var v35 float64 = -2.25
                                                                                _ = v35
                                                                            }
                                                                            _ = v32
                                                                            _ = v33
                                                                        }
                                                                        v36 := []int{9223372036854775807}
                                                                        _ = v36
                                                                    case 32767:
                                                                        if true {
                                                                            v37 := GenS5{F0: uint(v23), F1: v1, F2: 698}
                                                                            _ = len(v3)
                                                                            delete(v2, (^(fi24)))
                                                                            _ = v37
                                                                        }
                                                                    case -516:
                                                                        v1 = -1
                                                                    default:
                                                                        if false {
                                                                            v38 := GenS2{F0: (v7 + v7), F1: 2147483647, F2: false}
                                                                            v23 += v23
                                                                            var v39 bool = true
                                                                            _ = v38
                                                                            _ = v39
                                                                        } else {
                                                                            if ((int8(v7) == v1) && false) {
                                                                                v40 := []int{}
                                                                                switch (("hello" + v4) + ((v4 + "") + v4)) {
                                                                                case "foo bar":
                                                                                    var v41 int = (fi16 % ((fi24) | 1))
                                                                                    _ = v41
                                                                                case "hello":
                                                                                    fmt.Println(v4)
                                                                                default:
                                                                                    var v42 string = v4
                                                                                    _ = v42
                                                                                }
                                                                                _ = v40
                                                                            } else {
                                                                                v1 -= (^(v1))
                                                                                fmt.Println(v4)
                                                                                v4 = "foo bar"
                                                                            }
                                                                            v23 -= (v23 - v23)
                                                                        }
                                                                    }
                                                                    delete(v2, -1)
                                                                    if false {
                                                                        v43 := []int{(fi24 ^ 339), (fi16 - 857), 128, fi24}
                                                                        v1 = v1
                                                                        fmt.Println(v1, fi24)
                                                                        _ = v43
                                                                    }
                                                                }
                                                            default:
                                                                v4 = "hello"
                                                                v44 := GenS5{F0: 2, F1: v1, F2: 468}
                                                                var v45 int8 = v1
                                                                _ = v44
                                                                _ = v45
                                                            }
                                                        case "hello":
                                                            fmt.Println(fi16)
                                                        case "foo bar":
                                                            v1 = v1
                                                            var v46 float64 = 1e10
                                                            v4 = v4
                                                            v23 = 255
                                                            var v47 int = 501
                                                            _ = v46
                                                            _ = v47
                                                        default:
                                                            fmt.Println(fi24, v23)
                                                        }
                                                    }
                                                case -827:
                                                    var v48 int32 = 345
                                                    _ = v48
                                                case -1:
                                                    v7 += v7
                                                default:
                                                    var v49 uint32 = uint32(v7)
                                                    _ = v49
                                                }
                                            }
                                            v1 = 100
                                        }
                                        v23 = (v23 / ((v23) | 1))
                                        _ = fi24
                                    }
                                } else {
                                    v3["foo bar"] = 805
                                }
                                _ = v23
                            }
                            for fi50 := 0; fi50 < 10; fi50++ {
                                switch (fi50 | (fi16 % ((fi50) | 1))) {
                                case 4294967295:
                                    v51 := GenS4{F0: -2147483648}
                                    _ = v51
                                }
                                fmt.Println(v1, fi16)
                                _ = fi50
                            }
                        }
                    }
                    v1 -= 100
                    _ = fi16
                }
                var v52 int16 = 42
                v4 = "foo bar"
                _ = v52
            }
        }
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        v54 := GenS1{F0: 104, F1: 0, F2: (v7 >> ((4294967295) & 63))}
        for fi55 := 0; fi55 < 12; fi55++ {
            v56 := GenS1{F0: (v1 >> ((255) & 63)), F1: 1, F2: (v7 / (((v7 * 0)) | 1))}
            var v57 uint16 = 42
            if true {
                var v58 int32 = -308
                _ = v58
            }
            _ = fi55
            _ = v56
            _ = v57
        }
        var v59 int32 = -1
        _ = fi53
        _ = v54
        _ = v59
    }
    for fi60 := 0; fi60 < 14; fi60++ {
        for fi61 := 0; fi61 < 17; fi61++ {
            v7 += (v7 * v7)
            var v62 uint64 = 831
            if (((-2.25 - (((1.5 - 0.1) / 1.5) + -2.25)) * 1e10) != (3.14159 + (0.1 / 1.0))) {
                v7 -= v7
                var v63 uint32 = 2147483647
                _ = v63
            } else {
                var v64 uint16 = 65535
                _ = v64
            }
            _ = fi61
            _ = v62
        }
        _ = fi60
    }
    v7 = (v7 >> ((uint(v7)) & 63))
    v1 = 100
    v65 := GenS1{F0: -120, F1: 1, F2: v7}
    if (!true) {
        for fi66 := 0; fi66 < 11; fi66++ {
            var v67 uint8 = (^(v7))
            v2[fi66] = 65535
            _ = fi66
            _ = v67
        }
        for fi68 := 0; fi68 < 2; fi68++ {
            fmt.Println(v4, fi68)
            _ = fi68
        }
    } else {
        fmt.Println(v7)
        for fi69 := 0; fi69 < 12; fi69++ {
            var v70 int = fi69
            var v71 string = "foo bar"
            if false {
                for fi72 := 0; fi72 < 16; fi72++ {
                    v7 += 127
                    v7 += (v7 | (^(v7)))
                    _ = fi72
                }
            } else {
                if ((v1 / ((42) | 1)) == v1) {
                    var v73 uint16 = 2
                    v73 = 234
                    switch fi69 {
                    case 127:
                        if (16 <= 156) {
                            v74 := GenS4{F0: -555}
                            v75 := GenS7{F0: 2147483648, F1: 54, F2: 9223372036854775807}
                            _ = v74
                            _ = v75
                        } else {
                            v76 := GenS2{F0: v7, F1: -1, F2: true}
                            _ = v76
                        }
                    case 128:
                        if false {
                            var v77 int64 = -1
                            _ = v77
                        }
                    }
                    _ = v73
                } else {
                    var v78 uint = 18446744073709551615
                    fmt.Println(fi69)
                    _ = v78
                }
            }
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = len(v2)
    }
    fmt.Println(v4)
    v79 := []int{-206, 305}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v79:", len(v79), v79)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v65
    _ = v7
    _ = v79
}
