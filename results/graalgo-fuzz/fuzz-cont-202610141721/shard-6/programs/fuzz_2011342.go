package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 uint16; F2 bool }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int; F1 uint8; F2 int }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 uint; F1 uint8; F2 bool }
type GenS8 struct { F0 bool; F1 uint32; F2 bool }
func genFn0(p0 bool, p1 bool) uint8 {
    p1 = p0
    fmt.Println(p1, p1)
    fmt.Println(p1, p0)
    v7 := []int{-306}
    _ = p0
    _ = p1
    _ = v7
    return 0
}
func genFn1() bool {
    switch -1 {
    case -629:
        var v8 int64 = 42
        _ = v8
    }
    if true {
        if (true || (842 <= -778)) {
            var v9 string = ""
            var v10 uint16 = 42
            v11 := []int{128, int(v10)}
            _ = v10
            _ = v11
            _ = v9
        }
        var v12 int16 = -32768
        _ = v12
    }
    if false {
        v13 := []int{2147483647}
        for fi14 := 0; fi14 < 6; fi14++ {
            for fi15 := 0; fi15 < 9; fi15++ {
                if len(v13) > 0 {
                    v13[0] = (fi14 % ((fi15) | 1))
                }
                _ = fi15
            }
            _ = fi14
        }
        for fi16 := 0; fi16 < 8; fi16++ {
            var v17 float32 = 3.14159
            if false {
                if false {
                    v17 = v17
                } else {
                    v13 = append(v13, (fi16 / ((1) | 1)))
                }
                var v18 uint64 = 0
                var v19 int16 = -311
                _ = v18
                _ = v19
            } else {
                for fi20 := 0; fi20 < 10; fi20++ {
                    var v21 uint8 = 255
                    var v22 float64 = float64(v17)
                    var v23 int = fi16
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v23
                }
                v24 := GenS1{F0: (936 <= -416)}
                for fi25 := 0; fi25 < 7; fi25++ {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = fi25
                }
                _ = v24
            }
            _ = fi16
            _ = v17
        }
        _ = v13
    } else {
        v26 := []int{-275}
        var v27 float32 = (float32((0.1 + 1.5)) - float32((((0.0 + (0.0 / 0.1)) * 1.5) / -2.25)))
        v28 := GenS2{F0: 2, F1: 65535, F2: true}
        _ = v26
        _ = v27
        _ = v28
    }
    return false
}

func main() {
    v1 := []uint32{4294967295, 1, 972}
    var v2 uint32 = 2
    var v3 int8 = 37
    var v4 float32 = 0.0
    var v5 uint16 = 0
    v6 := []string{}
    var v29 uint32 = 240
    v30 := GenS2{F0: 1, F1: (^(v5)), F2: true}
    switch -104 {
    case 32768:
        var v31 int64 = 9223372036854775807
        _ = v31
    case 2147483647:
        var v32 int8 = (v3 / ((1) | 1))
        _ = v32
    default:
        var v33 string = ("a" + "hello")
        _ = v33
    }
    for fi34 := 0; fi34 < 19; fi34++ {
        v3 = (v3 / (((v3 + v3)) | 1))
        var v35 int32 = 655
        _ = fi34
        _ = v35
    }
    var v36 uint16 = v5
    v37 := GenS3{F0: 9223372036854775807, F1: 42}
    if false {
        if true {
            fmt.Println(v36)
            v38 := GenS4{F0: 940, F1: 124, F2: 42}
            _ = v38
        }
        var v39 uint = 555
        for fi40 := 0; fi40 < 13; fi40++ {
            switch fi40 {
            case 9223372036854775807:
                for fi41 := 0; fi41 < 17; fi41++ {
                    v29 += (^(v29))
                    if (!(false && false)) {
                        var v42 int8 = (v3 << ((v39) & 63))
                        fmt.Println(v39, fi40)
                        v42 += v42
                        _ = v42
                    } else {
                        if false {
                            var v43 uint = v39
                            _ = v43
                        } else {
                            var v44 int16 = 1000
                            if (!true) {
                                switch (fi41 % ((fi41) | 1)) {
                                case 71:
                                    var v45 int32 = -811
                                    _ = v45
                                case 127:
                                    v2 -= (v2 | 70)
                                }
                                v46 := GenS5{F0: 2, F1: v29}
                                var v47 float64 = (0.0 * -2.25)
                                var v48 int32 = 2147483647
                                v49 := GenS6{F0: 1000, F1: (^(v39)), F2: int16(v2)}
                                _ = v46
                                _ = v47
                                _ = v48
                                _ = v49
                            }
                            var v50 string = "foo bar"
                            _ = v44
                            _ = v50
                        }
                        v29 -= 2
                    }
                    fmt.Println(v4, v4, fi40)
                    v3 = int8(v5)
                    v51 := GenS5{F0: 128, F1: v29}
                    _ = fi41
                    _ = v51
                }
            default:
                for fi52 := 0; fi52 < 12; fi52++ {
                    if false {
                        if false {
                            switch -950 {
                            case 9223372036854775807:
                                v53 := GenS3{F0: 127, F1: 1}
                                _ = v53
                            default:
                                v39 -= 578
                            }
                            v54 := GenS3{F0: -910, F1: 1}
                            v5 += v36
                            _ = v54
                        }
                        if len(v6) > 0 {
                            v6[0] = ("\n\t" + ("\n\t" + "hello"))
                        }
                    }
                    fmt.Println(v2, v29)
                    v2 = v29
                    _ = fi52
                }
            }
            if false {
                if (!((int8(v5) >= (-(v3))) && false)) {
                    v5 -= uint16(v29)
                } else {
                    fmt.Println(v4)
                }
                var v55 uint16 = v36
                fmt.Println(v29)
                _ = v55
            } else {
                for fi56 := 0; fi56 < 12; fi56++ {
                    v4 = (1.5 - ((v4 - float32((0.0 - 1.5))) * ((3.14159 / v4) + v4)))
                    _ = fi56
                }
            }
            if (!((fi40 % ((fi40) | 1)) >= (fi40 % (((fi40 & 2147483647)) | 1)))) {
                var v57 int = (fi40 % ((9223372036854775807) | 1))
                switch -328 {
                case 829:
                    v58 := GenS5{F0: 73, F1: (^(v29))}
                    _ = v58
                case 128:
                    fmt.Println(v57, v57, v29)
                case 0:
                    v6 = append(v6, "\n\t")
                }
                var v59 float32 = v4
                _ = v57
                _ = v59
            } else {
                v2 -= (v29 >> ((0) & 63))
                fmt.Println(v36, v39)
            }
            _ = fi40
        }
        _ = v39
    }
    _ = len(v6)
    if true {
        if (42 < 128) {
            for fi60 := 0; fi60 < 19; fi60++ {
                fmt.Println(fi60, v5)
                var v61 string = "\n\t"
                _ = fi60
                _ = v61
            }
            v36 -= v5
            var v62 int32 = -335
            _ = v62
        } else {
            fmt.Println(v36, v2)
            fmt.Println(v5, v4, v4)
        }
    }
    if false {
        v6 = append(v6, "a")
        _ = len(v1)
        if false {
            v3 = v3
        }
    }
    if true {
        v63 := GenS7{F0: 725, F1: 1, F2: (!true)}
        if (true && (((0.0 + -2.25) - (0.0 / (-2.25 / 3.14159))) > (1e10 - -2.25))) {
            v5 += 0
            if ((false || (false || false)) || false) {
                v2 = v29
            }
            if true {
                for fi64 := 0; fi64 < 15; fi64++ {
                    var v65 int = (fi64 & (fi64 * (fi64 ^ -1)))
                    fmt.Println(v29, v3, v36)
                    _ = fi64
                    _ = v65
                }
                for fi66 := 0; fi66 < 10; fi66++ {
                    v2 = 65535
                    var v67 int64 = -9223372036854775808
                    var v68 uint16 = v5
                    _ = fi66
                    _ = v67
                    _ = v68
                }
            }
        }
        var v69 uint8 = 71
        _ = v63
        _ = v69
    } else {
        for fi70 := 0; fi70 < 14; fi70++ {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v36 = 780
            if (false || ((v3 | -8) != 108)) {
                for fi71 := 0; fi71 < 4; fi71++ {
                    var v72 string = (("\n\t" + "") + ("\n\t" + "hello"))
                    _ = len(v1)
                    _ = fi71
                    _ = v72
                }
            } else {
                var v73 uint32 = v2
                v74 := GenS4{F0: -2147483648, F1: 127, F2: (fi70 ^ fi70)}
                _ = v73
                _ = v74
            }
            _ = fi70
        }
        v75 := GenS8{F0: false, F1: 42, F2: (42 != -9223372036854775808)}
        _ = v75
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
