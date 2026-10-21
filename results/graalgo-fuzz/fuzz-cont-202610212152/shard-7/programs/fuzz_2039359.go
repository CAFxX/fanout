package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 uint32; F2 int8 }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint16; F1 int8 }
type GenS6 struct { F0 int; F1 int }
func genFn0(p0 uint) uint64 {
    if true {
        v7 := []int{2147483647, 42, 127, 293}
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                p0 -= (p0 - p0)
                p0 -= (^(p0))
                switch "\n\t" {
                case "foo bar":
                    if true {
                        var v10 uint64 = 91
                        switch "" {
                        case "a":
                            if len(v7) > 0 {
                                v7[0] = 224
                            }
                        case "\n\t":
                            if len(v7) > 0 {
                                v7[0] = (fi9 % ((127) | 1))
                            }
                        case "foo bar":
                            var v11 uint16 = uint16(v10)
                            _ = v11
                        default:
                            fmt.Println(fi9, v10)
                        }
                        var v12 uint64 = (^(v10))
                        _ = v10
                        _ = v12
                    } else {
                        fmt.Println(fi8, p0)
                        if false {
                            var v13 uint32 = 2147483647
                            v14 := GenS1{F0: 42}
                            var v15 float32 = -2.25
                            _ = v13
                            _ = v14
                            _ = v15
                        } else {
                            if false {
                                if (820 > -397) {
                                    p0 += 255
                                    v16 := GenS2{F0: -538, F1: 564, F2: 42}
                                    _ = v16
                                }
                                v17 := GenS3{F0: "hello", F1: "hello"}
                                var v18 string = "hello"
                                _ = v17
                                _ = v18
                            }
                            var v19 int32 = 1
                            _ = v19
                        }
                    }
                case "\n\t":
                    if true {
                        if (0.0 != ((3.14159 / ((((float32(1.5) - ((1.5 * 1.5) + ((-2.25 + 0.1) - -2.25))) / 1.0) - (-2.25 * -2.25)) * (((0.0 * 0.1) - 0.1) + float32(((1.5 + float64((1e10 - 1.5))) / float64(1.5)))))) * -2.25)) {
                            var v20 int16 = 1000
                            _ = v20
                        } else {
                            switch 2147483647 {
                            case 32767:
                                var v21 bool = true
                                _ = v21
                            default:
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                            }
                        }
                    }
                case "a":
                    fmt.Println(fi8, fi9, p0)
                default:
                    if false {
                        v22 := GenS2{F0: 965, F1: 33, F2: 0}
                        switch fi9 {
                        case -230:
                            v23 := GenS2{F0: 42, F1: 42, F2: int8(p0)}
                            _ = v23
                        case 32767:
                            if true {
                                var v24 uint16 = 1
                                _ = v24
                            }
                        case 32768:
                            var v25 int64 = -1
                            _ = v25
                        }
                        _ = v22
                    } else {
                        v26 := GenS1{F0: -32768}
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        p0 -= 2
                        _ = v26
                    }
                    if len(v7) > 0 {
                        v7[0] = (-(fi9))
                    }
                    var v27 int = (fi8 | (-(fi9)))
                    _ = v27
                }
                var v28 bool = false
                p0 -= p0
                _ = fi9
                _ = v28
            }
            if (!(2 == 860)) {
                v29 := GenS3{F0: ("" + ""), F1: "hello"}
                p0 += p0
                _ = v29
            }
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            _ = fi8
        }
        p0 = (^(p0))
        _ = v7
    }
    if (44 == 108) {
        switch (("foo bar" + ("" + "foo bar")) + (("" + ("a" + "")) + ("\n\t" + "foo bar"))) {
        case "a":
            p0 = p0
        default:
            switch 9223372036854775807 {
            case 9223372036854775807:
                v30 := GenS4{F0: true, F1: "hello"}
                _ = v30
            case -699:
                switch -843 {
                case 128:
                    if (2147483647 != -2147483648) {
                        for fi31 := 0; fi31 < 16; fi31++ {
                            fmt.Println(p0)
                            for fi32 := 0; fi32 < 19; fi32++ {
                                if (((((fi31 % (((fi31 >> (((p0 << ((p0) & 63))) & 63))) | 1)) > fi31) && true) || (2 <= 32767)) || ((((2 == 2) && (690 <= 42)) || true) || ((-2.25 >= 0.1) && (206 >= 248)))) {
                                    p0 -= 653
                                }
                                _ = fi32
                            }
                            p0 += p0
                            _ = fi31
                        }
                    }
                case 42:
                    for fi33 := 0; fi33 < 19; fi33++ {
                        var v34 string = "\n\t"
                        v35 := GenS4{F0: (-2.25 > -2.25), F1: (((("" + v34) + (v34 + ((v34 + "hello") + v34))) + "hello") + v34)}
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            }
        }
    }
    v36 := []int{-541, -2147483648}
    var v37 float64 = (1.5 * (3.14159 - -2.25))
    if true {
        v37 -= ((v37 + -2.25) - v37)
    }
    _ = p0
    _ = v36
    _ = v37
    return 42
}

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]bool)
    var v3 uint = 1
    var v4 uint16 = 147
    var v5 float64 = float64(1.5)
    var v6 uint64 = 1
    var v38 uint8 = 42
    v38 = (v38 / (((v38 & (^(v38)))) | 1))
    for fi39 := 0; fi39 < 14; fi39++ {
        if ((fi39 & fi39) == -1) {
            v38 = uint8(v1)
        }
        var v40 uint8 = v38
        var v41 uint8 = (v40 << ((v3) & 63))
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 uint32 = 2
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 int = (fi43 & (fi43 - fi43))
        v45 := GenS5{F0: 1, F1: 21}
        _ = fi43
        _ = v44
        _ = v45
    }
    v1 = v1
    fmt.Println(v4, v42)
    fmt.Println(v5, v1, v3)
    fmt.Println(v38, v38)
    var v46 uint64 = (v6 + (v6 + v6))
    for fi47 := 0; fi47 < 3; fi47++ {
        for fi48 := 0; fi48 < 7; fi48++ {
            v3 = 0
            fmt.Println(v38)
            _ = fi48
        }
        if (1e10 <= float32(v5)) {
            for fi49 := 0; fi49 < 7; fi49++ {
                if (!true) {
                    fmt.Println(v6)
                }
                v1 = (v42 * v42)
                v50 := GenS6{F0: 32768, F1: (fi49 << ((uint(v4)) & 63))}
                _ = fi49
                _ = v50
            }
            var v51 int64 = -1
            _ = v51
        } else {
            for fi52 := 0; fi52 < 4; fi52++ {
                fmt.Println(v1, v38)
                _ = fi52
            }
            if (!false) {
                if true {
                    if (-1 >= 42) {
                        var v53 float64 = 1.5
                        delete(v2, (^(v42)))
                        _ = v53
                    }
                } else {
                    v6 = (v46 + v6)
                }
            }
        }
        v54 := []int{fi47, (fi47 & (fi47 / (((fi47 & 32767)) | 1))), fi47, 0}
        _ = fi47
        _ = v54
    }
    v55 := v2[(v42 - v1)]
    var v56 int = 525
    switch "a" {
    case "a":
        for fi57 := 0; fi57 < 1; fi57++ {
            for fi58 := 0; fi58 < 3; fi58++ {
                v59 := []int{}
                _ = fi58
                _ = v59
            }
            _ = fi57
        }
    default:
        v6 = (^(v6))
    }
    var v60 int32 = 42
    var v61 float64 = (v5 / (float64(3.14159) * v5))
    var v62 uint = (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v6:", v6)
    fmt.Println("v60:", v60)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v46
    _ = v5
    _ = v55
    _ = v56
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
