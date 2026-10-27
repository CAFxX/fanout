package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint64; F1 bool }
type GenS7 struct { F0 uint64 }
func genFn0(p0 int8, p1 int8, p2 uint16) int8 {
    var v5 int16 = -32768
    if false {
        for fi6 := 0; fi6 < 16; fi6++ {
            v7 := GenS1{F0: -221}
            if true {
                var v8 uint8 = 32
                var v9 int64 = 1
                _ = v8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
        v10 := GenS2{F0: (v5 / ((v5) | 1))}
        _ = v10
    } else {
        var v11 int = 4294967295
        fmt.Println(v11)
        _ = v11
    }
    p0 = p1
    if (2 < p2) {
        switch ("hello" + "hello") {
        case "foo bar":
            for fi12 := 0; fi12 < 19; fi12++ {
                p1 = p0
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(p1)
                    _ = fi13
                }
                _ = fi12
            }
        default:
            v5 -= v5
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 ^ (p1 % ((p1) | 1)))
}

func main() {
    v1 := []string{"a", ("\n\t" + "foo bar")}
    v2 := []int{-1}
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]bool)
    for fi14 := 0; fi14 < 11; fi14++ {
        var v15 int = fi14
        switch fi14 {
        case -2147483648:
            var v16 uint64 = 42
            _ = v16
        case 4294967295:
            for fi17 := 0; fi17 < 1; fi17++ {
                v15 = 2147483647
                if false {
                    switch fi17 {
                    case 367:
                        if ((3.14159 != 1.5) || true) {
                            switch ("hello" + "hello") {
                            case "hello":
                                var v18 int32 = 0
                                _ = v18
                            case "":
                                v15 += 695
                            case "foo bar":
                                v19 := GenS3{F0: 405, F1: 32768}
                                _ = v19
                            }
                            v15 += -1
                            var v20 int8 = -128
                            var v21 uint64 = 42
                            if true {
                                v15 = (^(v15))
                                switch v15 {
                                case 954:
                                    fmt.Println(fi14, fi17, v20)
                                default:
                                    var v22 int16 = 13
                                    _ = v22
                                }
                            } else {
                                if (-175 > -1) {
                                    v15 -= 9223372036854775807
                                    if (0 > 175) {
                                        v15 -= (v15 % ((-2147483648) | 1))
                                        v23 := GenS4{F0: 1}
                                        _ = v23
                                    }
                                    var v24 float32 = 0.1
                                    _ = v24
                                } else {
                                    if false {
                                        var v25 uint = 535
                                        if true {
                                            v26 := GenS3{F0: (fi17 / ((-258) | 1)), F1: 32767}
                                            fmt.Println(v20, v15)
                                            _ = v26
                                        } else {
                                            var v27 uint8 = 92
                                            if (77 >= (v27 + v27)) {
                                                var v28 uint16 = 844
                                                v25 += 255
                                                var v29 float64 = float64(0.0)
                                                var v30 float64 = ((v29 + float64(float32(((float64(0.0) - v29) - (v29 * ((v29 - 1.5) / (1e10 - v29))))))) / v29)
                                                _ = v28
                                                _ = v29
                                                _ = v30
                                            } else {
                                                fmt.Println(v20, v20, fi17)
                                            }
                                            v31 := GenS4{F0: 2147483647}
                                            _ = v27
                                            _ = v31
                                        }
                                        v32 := GenS2{F0: 42}
                                        _ = v25
                                        _ = v32
                                    }
                                }
                                if false {
                                    fmt.Println(v21)
                                }
                                if (!(true && (false || (200 != 930)))) {
                                    v21 = (v21 << ((4) & 63))
                                } else {
                                    switch v15 {
                                    case 42:
                                        switch "hello" {
                                        case "hello":
                                            if true {
                                                if false {
                                                    var v33 int = 955
                                                    if (2 == 287) {
                                                        var v34 uint32 = 294
                                                        _ = v34
                                                    } else {
                                                        var v35 uint64 = 4294967295
                                                        _ = v35
                                                    }
                                                    _ = v33
                                                } else {
                                                    var v36 float64 = 3.14159
                                                    v15 = -576
                                                    fmt.Println(v21)
                                                    _ = v36
                                                }
                                            }
                                        case "a":
                                            if false {
                                                fmt.Println(v21, v20, v20)
                                            } else {
                                                v21 -= (v21 / (((v21 & v21)) | 1))
                                                var v37 string = "a"
                                                _ = v37
                                            }
                                        case "":
                                            v21 -= 42
                                        default:
                                            if ((fi14 << ((255) & 63)) == fi17) {
                                                if true {
                                                    switch (^(fi14)) {
                                                    case 1:
                                                        v38 := GenS3{F0: 128, F1: v15}
                                                        _ = v38
                                                    default:
                                                        switch fi14 {
                                                        case 9223372036854775807:
                                                            if true {
                                                                v20 = genFn0(genFn0(v20, v20, 42), 111, 1)
                                                            } else {
                                                                var v39 uint8 = 141
                                                                _ = v39
                                                            }
                                                        default:
                                                            v21 += v21
                                                        }
                                                    }
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v40 int32 = -2147483648
                                                    _ = v40
                                                }
                                            } else {
                                                v21 = (v21 * v21)
                                                v3[(fi17 / ((fi14) | 1))] = 774
                                            }
                                        }
                                        switch (fi14 | (fi14 % (((fi17 & 0)) | 1))) {
                                        case 42:
                                            var v41 int32 = -660
                                            _ = v41
                                        case 32767:
                                            v21 += v21
                                        case -820:
                                            v42 := GenS5{F0: -183}
                                            _ = v42
                                        }
                                    case 9223372036854775807:
                                        var v43 float32 = (0.1 / -2.25)
                                        _ = v43
                                    case 725:
                                        var v44 float64 = 0.0
                                        _ = v44
                                    default:
                                        var v45 bool = (!(-2.25 > ((3.14159 * 0.0) - 3.14159)))
                                        _ = v45
                                    }
                                }
                                var v46 uint8 = 2
                                _ = v46
                            }
                            _ = v20
                            _ = v21
                        } else {
                            v15 = 0
                            v3[fi17] = 65535
                            if (false && false) {
                                switch (((((("foo bar" + "a") + "a") + "") + "") + "hello") + "hello") {
                                case "":
                                    v47 := GenS2{F0: 42}
                                    _ = v47
                                }
                            }
                        }
                    case -1:
                        v15 += (-(v15))
                    }
                    v48 := GenS6{F0: 4294967295, F1: (!((float32(0.0) - 1e10) <= 1.5))}
                    var v49 int = v15
                    _ = v48
                    _ = v49
                } else {
                    v15 -= 9223372036854775807
                    if true {
                        var v50 string = ""
                        var v51 int = (fi14 & v15)
                        _ = v50
                        _ = v51
                    } else {
                        switch -9223372036854775808 {
                        case 2147483647:
                            var v52 int16 = 1000
                            v52 -= v52
                            var v53 string = ""
                            _ = v52
                            _ = v53
                        case 9223372036854775807:
                            v15 += (fi14 / (((fi14 << ((2) & 63))) | 1))
                        }
                        var v54 int32 = 0
                        var v55 uint16 = 966
                        _ = v54
                        _ = v55
                    }
                    if (-1 > 100) {
                        v15 = -83
                        if false {
                            var v56 int64 = -537
                            var v57 int8 = 30
                            _ = v56
                            _ = v57
                        }
                        v15 = (fi14 + (fi14 / (((fi14 << ((4294967295) & 63))) | 1)))
                    } else {
                        v15 -= fi17
                        v2 = append(v2, fi14)
                        if true {
                            if len(v2) > 0 {
                                v2[0] = 42
                            }
                            var v58 uint32 = 2
                            _ = v58
                        } else {
                            switch 9223372036854775807 {
                            case 42:
                                v15 = v15
                            default:
                                var v59 uint64 = 2
                                _ = v59
                            }
                        }
                    }
                    if true {
                        var v60 float64 = ((1e10 / 1e10) + 1.5)
                        v61 := GenS7{F0: 769}
                        var v62 uint = 566
                        v63 := GenS4{F0: 0}
                        _ = v60
                        _ = v61
                        _ = v62
                        _ = v63
                    } else {
                        v64 := GenS1{F0: 594}
                        var v65 uint64 = 18446744073709551615
                        v3[(-(fi17))] = 691
                        if (251 > -394) {
                            var v66 int64 = 42
                            var v67 int32 = -2147483648
                            _ = v66
                            _ = v67
                        }
                        _ = v64
                        _ = v65
                    }
                }
                var v68 int16 = 1000
                _ = fi17
                _ = v68
            }
        }
        var v69 string = ("a" + "")
        _ = fi14
        _ = v15
        _ = v69
    }
    var v70 float64 = 1e10
    if true {
        for fi71 := 0; fi71 < 16; fi71++ {
            fmt.Println(v70, fi71)
            _ = fi71
        }
    }
    delete(v4, 2147483648)
    if len(v1) > 0 {
        v1[0] = ""
    }
    var v72 uint8 = 127
    if (1 < 279) {
        fmt.Println(v72, v72)
        v73 := v4[2147483647]
        var v74 float64 = (1.5 * (v70 + 3.14159))
        _ = v73
        _ = v74
    }
    if true {
        var v75 uint32 = 1
        _ = v75
    }
    var v76 int32 = 42
    var v77 uint32 = 0
    var v78 int64 = -583
    if (true || (128 <= v72)) {
        v79 := GenS5{F0: 42}
        _ = v79
    } else {
        var v80 int32 = (v76 * (v76 | 0))
        v77 += (v77 / ((v77) | 1))
        for fi81 := 0; fi81 < 1; fi81++ {
            for fi82 := 0; fi82 < 12; fi82++ {
                var v83 string = ""
                if false {
                    v84 := GenS6{F0: 0, F1: false}
                    _ = v84
                }
                _ = fi82
                _ = v83
            }
            _ = fi81
        }
        _ = v80
    }
    if true {
        if ("hello" == ("a" + "")) {
            switch 32768 {
            case 9223372036854775807:
                var v85 string = (("" + "a") + ("foo bar" + "foo bar"))
                _ = v85
            case -350:
                var v86 float64 = v70
                _ = v86
            case 2147483647:
                v77 += v77
            default:
                _ = len(v2)
            }
        } else {
            if false {
                var v87 string = "a"
                v70 += v70
                _ = v87
            }
            v78 -= -9223372036854775808
        }
    }
    v77 = v77
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v70:", v70)
    fmt.Println("v72:", v72)
    fmt.Println("v76:", v76)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v70
    _ = v72
    _ = v76
    _ = v77
    _ = v78
}
