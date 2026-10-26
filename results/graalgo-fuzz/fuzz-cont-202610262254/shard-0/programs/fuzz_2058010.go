package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint64; F1 int; F2 uint32 }
type GenS3 struct { F0 uint16; F1 bool; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int32; F1 uint8; F2 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 int; F2 string }
func genFn0(p0 int16) uint8 {
    if (!false) {
        var v6 uint16 = 1
        fmt.Println(p0, p0)
        _ = v6
    } else {
        v7 := GenS1{F0: 42, F1: 95}
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 uint8 = 42
            v10 := []int{fi8, 473, int(v9)}
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        switch 128 {
        case 756:
            if true {
                v11 := []int{1, -993}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 5; fi12++ {
                    p0 = p0
                    _ = fi12
                }
                v13 := []int{}
                switch 42 {
                case -2147483648:
                    var v14 uint64 = 18446744073709551615
                    _ = v14
                case 2147483647:
                    fmt.Println(p0)
                case 32768:
                    var v15 float64 = 0.0
                    _ = v15
                }
                _ = v13
            }
        default:
            switch ((("a" + (("a" + ("\n\t" + "a")) + (((("a" + ("foo bar" + "")) + "\n\t") + ("hello" + ("a" + "a"))) + "a"))) + "foo bar") + "\n\t") {
            case "\n\t":
                var v16 float32 = (1e10 - (float32(0.1) / 1.0))
                _ = v16
            case "":
                fmt.Println(p0)
            case "foo bar":
                fmt.Println(p0)
            default:
                v17 := GenS1{F0: 1, F1: 42}
                _ = v17
            }
        }
        _ = v7
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        for fi19 := 0; fi19 < 8; fi19++ {
            v20 := GenS1{F0: 255, F1: -1}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    switch 32768 {
    case 128:
        v21 := GenS1{F0: 2, F1: 127}
        _ = v21
    default:
        v22 := GenS2{F0: 1, F1: -2147483648, F2: 2}
        _ = v22
    }
    _ = p0
    return 128
}
func genFn1(p0 int32, p1 uint32) uint {
    p0 = (p0 + p0)
    switch 9223372036854775807 {
    case 127:
        v23 := []int{9223372036854775807}
        _ = v23
    case -841:
        if (true && (725 < -1)) {
            p0 = -2147483648
            var v24 uint16 = 65535
            _ = v24
        } else {
            switch -989 {
            case 824:
                var v25 int16 = 42
                _ = v25
            case 2147483647:
                v26 := []int{-713}
                _ = v26
            default:
                p0 -= (p0 % ((0) | 1))
            }
            var v27 int = 2147483647
            _ = v27
        }
    }
    var v28 int16 = 367
    _ = p0
    _ = p1
    _ = v28
    return 0
}

func main() {
    var v1 uint64 = 0
    var v2 string = ("a" + ("a" + ("a" + "")))
    v3 := []uint32{2147483647, 4294967295, 65535, 65535}
    var v4 int16 = -405
    var v5 int64 = -9223372036854775808
    for fi29 := 0; fi29 < 16; fi29++ {
        v1 = (v1 & (v1 % ((2) | 1)))
        if false {
            v1 = v1
        }
        fmt.Println(v5, fi29)
        _ = fi29
    }
    v2 = v2
    fmt.Println(v2, v1)
    var v30 int64 = v5
    var v31 int32 = 661
    if len(v3) > 0 {
        v3[0] = 65535
    }
    if true {
        for fi32 := 0; fi32 < 14; fi32++ {
            v33 := GenS3{F0: 2, F1: true, F2: true}
            _ = fi32
            _ = v33
        }
        if (((v31 - (v31 + 2147483647)) != (v31 >> ((2) & 63))) && (230 == 128)) {
            if false {
                for fi34 := 0; fi34 < 3; fi34++ {
                    switch 0 {
                    case 1:
                        for fi35 := 0; fi35 < 5; fi35++ {
                            v1 -= v1
                            _ = len(v3)
                            _ = fi35
                        }
                    case 32767:
                        var v36 uint32 = 4294967295
                        _ = v36
                    case 9223372036854775807:
                        fmt.Println(v2)
                    }
                    var v37 int = fi34
                    switch "\n\t" {
                    case "foo bar":
                        var v38 uint32 = 65535
                        _ = v38
                    case "hello":
                        switch (((("foo bar" + v2) + (v2 + "")) + v2) + v2) {
                        case "hello":
                            v39 := GenS4{F0: true}
                            _ = v39
                        case "":
                            if (!true) {
                                v5 = (v5 / (((^(v5))) | 1))
                            }
                        case "a":
                            for fi40 := 0; fi40 < 11; fi40++ {
                                v41 := GenS5{F0: v31, F1: 191, F2: v37}
                                v37 += 32767
                                _ = fi40
                                _ = v41
                            }
                        default:
                            if (!true) {
                                switch v2 {
                                case "foo bar":
                                    for fi42 := 0; fi42 < 18; fi42++ {
                                        fmt.Println(v2, v5, v5)
                                        fmt.Println(v31, v30)
                                        _ = fi42
                                    }
                                case "":
                                    v43 := GenS2{F0: 1, F1: v37, F2: 358}
                                    _ = v43
                                }
                                _ = len(v3)
                            }
                        }
                    case "":
                        var v44 uint64 = v1
                        _ = v44
                    }
                    _ = fi34
                    _ = v37
                }
                switch v2 {
                case "":
                    var v45 uint64 = v1
                    _ = v45
                case "hello":
                    fmt.Println(v1)
                default:
                    var v46 uint32 = 65535
                    _ = v46
                }
                for fi47 := 0; fi47 < 13; fi47++ {
                    var v48 bool = true
                    fmt.Println(v48, v4, v1)
                    var v49 int8 = -1
                    _ = fi47
                    _ = v48
                    _ = v49
                }
            } else {
                _ = len(v3)
                var v50 float32 = 1.5
                _ = v50
            }
            if false {
                var v51 uint64 = 4294967295
                _ = v51
            } else {
                fmt.Println(v2, v1)
                for fi52 := 0; fi52 < 19; fi52++ {
                    v53 := GenS3{F0: 548, F1: (!((4294967295 <= 65535) || false)), F2: true}
                    _ = fi52
                    _ = v53
                }
            }
            fmt.Println(v1, v1, v30)
        } else {
            var v54 int8 = -1
            for fi55 := 0; fi55 < 13; fi55++ {
                v56 := GenS6{F0: v31}
                v3 = append(v3, 2147483648)
                v57 := GenS6{F0: 0}
                _ = fi55
                _ = v56
                _ = v57
            }
            if (-781 <= -2147483648) {
                if (4294967295 < 41) {
                    var v58 uint16 = 0
                    v5 -= 9223372036854775807
                    _ = v58
                }
                var v59 int32 = (-(v31))
                _ = v59
            } else {
                fmt.Println(v4, v1)
                v60 := GenS5{F0: -2147483648, F1: 42, F2: 2147483647}
                _ = v60
            }
            _ = v54
        }
        v4 = (v4 << ((65535) & 63))
    } else {
        v30 = (v5 | (v30 * int64(v1)))
        v31 += (v31 | v31)
    }
    v61 := GenS6{F0: (v31 << ((genFn1(-2147483648, 65535)) & 63))}
    var v62 uint = 2
    v63 := GenS4{F0: false}
    v5 -= 9223372036854775807
    if len(v3) > 0 {
        v3[0] = 42
    }
    if (!false) {
        fmt.Println(v4)
        if true {
            if ((-108 <= 127) || (32767 > 42)) {
                v5 -= (v5 - 1)
            }
        } else {
            for fi64 := 0; fi64 < 7; fi64++ {
                _ = len(v3)
                for fi65 := 0; fi65 < 4; fi65++ {
                    v66 := GenS7{F0: v5, F1: 4294967295, F2: "foo bar"}
                    fmt.Println(v4, v62, v4)
                    v1 = (v1 * (v1 | 4294967295))
                    _ = fi65
                    _ = v66
                }
                v62 -= (v62 << ((uint(v1)) & 63))
                _ = fi64
            }
        }
        var v67 int32 = (v31 << (((v62 & 29)) & 63))
        _ = v67
    } else {
        if (1.5 != 1e10) {
            switch v2 {
            case "\n\t":
                v4 -= v4
            }
            v68 := GenS4{F0: false}
            var v69 int64 = v30
            _ = v68
            _ = v69
        } else {
            v5 += 1
            if ((!(0.0 >= 1.5)) && false) {
                switch "" {
                case "a":
                    if true {
                        var v70 int32 = (-(v31))
                        v31 = 1
                        _ = len(v3)
                        _ = v70
                    } else {
                        v71 := GenS1{F0: 186, F1: int8(v31)}
                        _ = v71
                    }
                case "foo bar":
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case "\n\t":
                    _ = len(v3)
                }
            }
        }
    }
    var v72 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v62:", v62)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v61
    _ = v62
    _ = v63
    _ = v72
}
