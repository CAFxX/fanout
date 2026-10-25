package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int8; F2 int8 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 uint32; F1 string; F2 string }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int8 }
func genFn0(p0 int32) int64 {
    for fi6 := 0; fi6 < 14; fi6++ {
        var v7 int16 = 1
        v8 := []int{fi6, fi6, 42}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 255, F1: 886}
    _ = p0
    _ = v9
    return 42
}
func genFn1(p0 bool, p1 int64) uint64 {
    if p0 {
        if false {
            v10 := []int{35}
            _ = v10
        }
        p1 += (^(p1))
    }
    p1 = p1
    v11 := GenS2{F0: (!p0), F1: 1}
    v12 := []int{-9223372036854775808, 2147483647}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 4294967295
}

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    var v4 int16 = 42
    v5 := make(map[uint32]uint32)
    v13 := []int{-2147483648, 0}
    if false {
        if false {
            switch "\n\t" {
            case "hello":
                var v14 int64 = -839
                _ = v14
            default:
                v15 := GenS3{F0: v1}
                _ = v15
            }
        }
    } else {
        switch "" {
        case "":
            var v16 string = ("foo bar" + "\n\t")
            _ = v16
        case "a":
            v17 := GenS4{F0: 906, F1: 100, F2: 42}
            _ = v17
        case "foo bar":
            if (1 == 128) {
                fmt.Println(v1)
                v18 := GenS3{F0: (v1 + v1)}
                _ = v18
            }
        }
        for fi19 := 0; fi19 < 20; fi19++ {
            if (((!false) || true) || (-1 <= genFn0(1))) {
                v1 += (v1 % ((42) | 1))
                var v20 uint8 = 105
                for fi21 := 0; fi21 < 19; fi21++ {
                    if len(v13) > 0 {
                        v13[0] = (fi21 >> ((42) & 63))
                    }
                    var v22 uint64 = 151
                    _ = fi21
                    _ = v22
                }
                _ = v20
            } else {
                v13 = append(v13, int(v4))
            }
            _ = fi19
        }
    }
    for fi23 := 0; fi23 < 13; fi23++ {
        switch (fi23 & 42) {
        case 428:
            v4 = v4
        }
        fmt.Println(v4, v4)
        _ = fi23
    }
    v1 = 249
    v24 := GenS5{F0: (v1 >> ((42) & 63)), F1: -128}
    for fi25 := 0; fi25 < 8; fi25++ {
        v26 := GenS1{F0: 32767, F1: 355}
        var v27 float64 = float64(-2.25)
        v4 -= (v4 | 42)
        _ = fi25
        _ = v26
        _ = v27
    }
    v28 := GenS2{F0: (genFn1(true, 0) > 18446744073709551615), F1: (v1 + v1)}
    v29 := GenS4{F0: 1, F1: -96, F2: 117}
    switch 395 {
    case 1:
        var v30 bool = false
        _ = v30
    default:
        if false {
            switch -185 {
            case 477:
                if ((!(true || true)) || false) {
                    v4 = (v4 + (^(v4)))
                    if false {
                        v1 -= (v1 % (((^(v1))) | 1))
                    } else {
                        v4 -= (v4 + (v4 % ((v4) | 1)))
                        v1 -= v1
                    }
                    if (!((0 <= 602) && true)) {
                        v1 -= v1
                        var v31 string = ("foo bar" + "\n\t")
                        _ = v31
                    }
                } else {
                    var v32 float32 = float32(0.1)
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    _ = v32
                }
            case 4294967295:
                v33 := GenS4{F0: -2147483648, F1: -1, F2: -18}
                _ = v33
            case -488:
                switch "" {
                case "a":
                    if ((1.5 + 1.5) <= ((1e10 - 3.14159) * 3.14159)) {
                        var v34 int16 = v4
                        _ = v34
                    }
                case "hello":
                    for fi35 := 0; fi35 < 14; fi35++ {
                        v4 = (-(v4))
                        for fi36 := 0; fi36 < 9; fi36++ {
                            var v37 float64 = -2.25
                            fmt.Println(fi36, v1)
                            fmt.Println(v1, v1)
                            _ = fi36
                            _ = v37
                        }
                        var v38 uint32 = v1
                        _ = fi35
                        _ = v38
                    }
                default:
                    v1 = v1
                }
            }
            var v39 int = 4294967295
            if true {
                var v40 float64 = 1.5
                switch "hello" {
                case "":
                    v41 := GenS6{F0: 1, F1: "a", F2: (("a" + "hello") + "a")}
                    _ = v41
                }
                for fi42 := 0; fi42 < 13; fi42++ {
                    fmt.Println(v39, v1)
                    v1 += 4294967295
                    _ = fi42
                }
                _ = v40
            }
            _ = v39
        } else {
            var v43 int = -505
            v44 := v5[1]
            _ = v43
            _ = v44
        }
    }
    var v45 int32 = 0
    if false {
        for fi46 := 0; fi46 < 5; fi46++ {
            v47 := GenS1{F0: 990, F1: 568}
            v4 = (v4 & (^(v4)))
            var v48 float32 = (3.14159 * 0.0)
            _ = fi46
            _ = v47
            _ = v48
        }
        for fi49 := 0; fi49 < 13; fi49++ {
            delete(v5, (v1 * (^(v1))))
            fmt.Println(fi49, v1)
            var v50 int64 = 1
            _ = fi49
            _ = v50
        }
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    switch 370 {
    case -671:
        fmt.Println(v1, v1)
    case -9223372036854775808:
        v51 := GenS6{F0: v1, F1: "hello", F2: "\n\t"}
        _ = v51
    default:
        fmt.Println(v1, v4)
    }
    fmt.Println(v4, v1)
    for fi52 := 0; fi52 < 6; fi52++ {
        if true {
            v53 := v5[v1]
            _ = v53
        } else {
            if (((((0.0 - 0.0) - 3.14159) + -2.25) - 0.1) <= 0.0) {
                v54 := GenS7{F0: 81, F1: genFn1(((^(v1)) <= 65535), 651)}
                switch "\n\t" {
                case "hello":
                    v55 := GenS7{F0: int8(v4), F1: 42}
                    _ = v55
                case "\n\t":
                    v56 := GenS8{F0: "\n\t"}
                    _ = v56
                }
                _ = v54
            } else {
                var v57 int32 = v45
                _ = v57
            }
        }
        if (false && false) {
            delete(v2, "hello")
            v58 := GenS9{F0: 0}
            _ = v58
        } else {
            if (false || (!(18446744073709551615 >= 0))) {
                if false {
                    for fi59 := 0; fi59 < 10; fi59++ {
                        v60 := GenS3{F0: v1}
                        if len(v13) > 0 {
                            v13[0] = (fi52 >> ((255) & 63))
                        }
                        switch -737 {
                        case 32767:
                            delete(v2, "")
                        case 127:
                            var v61 string = "\n\t"
                            _ = v61
                        }
                        _ = fi59
                        _ = v60
                    }
                } else {
                    v4 = v4
                }
            }
            if (false || false) {
                v4 -= 42
                var v62 uint = 0
                _ = v62
            } else {
                var v63 int64 = 0
                v64 := GenS4{F0: (v45 - (v45 - (v45 ^ (v45 * (^(v45)))))), F1: 27, F2: 1}
                _ = v63
                _ = v64
            }
        }
        v1 += v1
        _ = fi52
    }
    for fi65 := 0; fi65 < 20; fi65++ {
        fmt.Println(v1)
        v45 = (v45 << ((2) & 63))
        var v66 uint64 = 1
        _ = fi65
        _ = v66
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v45
    _ = v5
}
