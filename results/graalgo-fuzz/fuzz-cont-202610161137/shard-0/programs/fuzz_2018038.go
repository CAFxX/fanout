package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 string; F1 uint16 }
type GenS7 struct { F0 uint8; F1 uint32; F2 int64 }
type GenS8 struct { F0 uint64; F1 bool }
func genFn0() uint16 {
    var v6 int = -223
    var v7 float32 = (1.5 + 0.1)
    _ = v6
    _ = v7
    return 783
}

func main() {
    v1 := []int{4294967295, 9223372036854775807, 4294967295}
    var v2 int32 = 2147483647
    v3 := make(map[int]int)
    var v4 int64 = -554
    var v5 uint64 = 242
    var v8 int16 = 1000
    var v9 uint = 4294967295
    v5 -= v5
    if (0 >= v9) {
        if true {
            v10 := GenS1{F0: false}
            for fi11 := 0; fi11 < 1; fi11++ {
                fmt.Println(v9, v4, v8)
                for fi12 := 0; fi12 < 4; fi12++ {
                    fmt.Println(fi11)
                    _ = fi12
                }
                _ = fi11
            }
            _ = v10
        }
    }
    fmt.Println(v8, v8)
    if true {
        if true {
            for fi13 := 0; fi13 < 18; fi13++ {
                v14 := GenS1{F0: (true && ("\n\t" < "a"))}
                if false {
                    v1 = append(v1, -9223372036854775808)
                    var v15 uint = v9
                    _ = v15
                } else {
                    fmt.Println(v9, fi13)
                    var v16 int = fi13
                    if len(v1) > 0 {
                        v1[0] = v16
                    }
                    _ = v16
                }
                _ = fi13
                _ = v14
            }
            delete(v3, 32767)
            var v17 float64 = ((-2.25 * ((0.0 / (0.0 + float64(1.5))) / 1e10)) * 1.5)
            _ = v17
        } else {
            v2 += (-(v2))
            v18 := GenS2{F0: (^(v2)), F1: (v8 + (v8 % (((v8 / ((v8) | 1))) | 1))), F2: (v4 | (v4 << (((v9 << ((v9) & 63))) & 63)))}
            _ = v18
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            v20 := GenS1{F0: (2147483648 >= 23)}
            if (!(!(!(v8 == 32767)))) {
                var v21 float32 = (-2.25 - ((1.5 / 3.14159) + float32((0.1 + (-2.25 + -2.25)))))
                v22 := GenS2{F0: int32(v8), F1: v8, F2: (v4 % (((v4 - v4)) | 1))}
                var v23 uint32 = 224
                _ = v21
                _ = v22
                _ = v23
            } else {
                v2 = v2
                if true {
                    switch fi19 {
                    case 4294967295:
                        switch fi19 {
                        case -865:
                            v2 = 1
                        case 128:
                            if (!true) {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v24 := GenS1{F0: (false || true)}
                                _ = v24
                            } else {
                                switch fi19 {
                                case 32767:
                                    var v25 int32 = (v2 + 2147483647)
                                    _ = v25
                                case 4294967295:
                                    if ((v8 << ((v9) & 63)) >= -32768) {
                                        for fi26 := 0; fi26 < 3; fi26++ {
                                            fmt.Println(v9, v9)
                                            switch fi26 {
                                            case -93:
                                                switch (fi19 / (((fi26 << ((v9) & 63))) | 1)) {
                                                case 0:
                                                    if ((3.14159 / 1.0) > -2.25) {
                                                        switch "" {
                                                        case "foo bar":
                                                            v27 := GenS3{F0: v4, F1: 0}
                                                            _ = v27
                                                        case "hello":
                                                            v9 = v9
                                                        default:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v28 uint32 = 742
                                                                _ = v28
                                                            case "a":
                                                                var v29 int64 = v4
                                                                _ = v29
                                                            case "\n\t":
                                                                if (!true) {
                                                                    var v30 bool = (236 <= 1)
                                                                    var v31 float64 = float64(float32(3.14159))
                                                                    _ = v30
                                                                    _ = v31
                                                                } else {
                                                                    switch fi19 {
                                                                    case -9223372036854775808:
                                                                        v32 := v3[fi26]
                                                                        _ = v32
                                                                    }
                                                                    v33 := GenS4{F0: false, F1: 0, F2: v5}
                                                                    v34 := GenS4{F0: false, F1: -2147483648, F2: (v5 << ((v9) & 63))}
                                                                    if false {
                                                                        var v35 uint32 = 65535
                                                                        fmt.Println(v5)
                                                                        _ = v35
                                                                    } else {
                                                                        _ = len(v1)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        var v36 int8 = -128
                                                                        fmt.Println(v9, v9)
                                                                        _ = v36
                                                                    }
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                            default:
                                                                v37 := GenS5{F0: "a"}
                                                                _ = v37
                                                            }
                                                        }
                                                        v38 := GenS1{F0: (!(0.1 > 3.14159))}
                                                        v2 -= (v2 % ((v2) | 1))
                                                        _ = v38
                                                    } else {
                                                        var v39 string = "\n\t"
                                                        _ = v39
                                                    }
                                                case 434:
                                                    fmt.Println(v9)
                                                case -919:
                                                    _ = len(v1)
                                                default:
                                                    var v40 int = fi26
                                                    _ = v40
                                                }
                                            case 455:
                                                if false {
                                                    v4 -= v4
                                                    if (206 > 893) {
                                                        switch fi26 {
                                                        case 552:
                                                            if (fi26 > (fi19 * 255)) {
                                                                v41 := GenS2{F0: (^(v2)), F1: int16(v2), F2: -1}
                                                                v4 = (^(v4))
                                                                _ = v41
                                                            }
                                                        case 4:
                                                            v8 = int16(v2)
                                                        default:
                                                            var v42 float32 = float32(3.14159)
                                                            _ = v42
                                                        }
                                                        var v43 uint32 = 645
                                                        _ = v43
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    if (255 >= 255) {
                                                        v5 = uint64(v9)
                                                        var v44 int8 = 100
                                                        _ = v44
                                                    }
                                                    var v45 uint = v9
                                                    v4 = (^(v4))
                                                    v2 -= v2
                                                    _ = v45
                                                } else {
                                                    var v46 int32 = (v2 ^ v2)
                                                    var v47 uint16 = 897
                                                    var v48 uint8 = 1
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                }
                                            case -9223372036854775808:
                                                v4 -= (v4 + 1)
                                            }
                                            _ = fi26
                                        }
                                    }
                                }
                                v9 = (v9 + (v9 ^ (v9 - v9)))
                                for fi49 := 0; fi49 < 20; fi49++ {
                                    var v50 float32 = 3.14159
                                    v51 := GenS2{F0: 2147483647, F1: v8, F2: 0}
                                    var v52 uint = 42
                                    v53 := v3[(^(fi19))]
                                    switch fi19 {
                                    case 2147483647:
                                        switch (((("hello" + ("hello" + "a")) + "a") + "\n\t") + "hello") {
                                        case "":
                                            var v54 float32 = v50
                                            _ = v54
                                        case "\n\t":
                                            v55 := GenS6{F0: ("\n\t" + "a"), F1: 488}
                                            _ = v55
                                        case "a":
                                            var v56 int16 = (^(v8))
                                            _ = v56
                                        default:
                                            var v57 uint8 = 112
                                            _ = v57
                                        }
                                    case -433:
                                        var v58 int32 = v2
                                        _ = v58
                                    case 995:
                                        var v59 uint32 = 2147483648
                                        _ = v59
                                    }
                                    fmt.Println(v8)
                                    _ = fi49
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                            }
                        case -559:
                            for fi60 := 0; fi60 < 7; fi60++ {
                                delete(v3, 1)
                                var v61 int = (fi19 / ((int(v5)) | 1))
                                v62 := GenS7{F0: 42, F1: 4294967295, F2: v4}
                                _ = fi60
                                _ = v61
                                _ = v62
                            }
                        default:
                            v2 += 0
                        }
                    default:
                        v1 = append(v1, int(v4))
                    }
                }
            }
            v2 += v2
            _ = fi19
            _ = v20
        }
        v2 += (-(v2))
    } else {
        var v63 string = "\n\t"
        v4 = (^(v4))
        var v64 string = (v63 + ("hello" + v63))
        _ = v63
        _ = v64
    }
    if true {
        for fi65 := 0; fi65 < 3; fi65++ {
            fmt.Println(v9, v2)
            v66 := v3[(-(fi65))]
            v5 = (v5 + v5)
            _ = fi65
            _ = v66
        }
        var v67 float64 = 1e10
        _ = v67
    } else {
        fmt.Println(v8, v4, v4)
        v8 += v8
        fmt.Println(v9, v9)
    }
    if (42 < -2) {
        v68 := GenS8{F0: 1, F1: ((v5 | 2) == (v5 - (v5 - 2)))}
        _ = v68
    }
    v69 := GenS4{F0: (!(true && true)), F1: v2, F2: 18446744073709551615}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v69
    _ = v8
    _ = v9
}
