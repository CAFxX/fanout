package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint; F1 uint32; F2 bool }
type GenS4 struct { F0 uint }
func genFn0(p0 int16) uint32 {
    v6 := GenS1{F0: -67, F1: 484}
    fmt.Println(p0)
    p0 -= (p0 >> ((2) & 63))
    p0 = p0
    fmt.Println(p0)
    _ = p0
    _ = v6
    return 157
}

func main() {
    var v1 int64 = -1
    v2 := []uint32{778, 1, 2}
    var v3 int16 = -1
    v4 := make(map[uint32]int)
    var v5 uint16 = 709
    var v7 uint8 = 2
    switch "a" {
    case "a":
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 float32 = (0.0 * 0.1)
            var v10 uint64 = 2
            _ = fi8
            _ = v10
            _ = v9
        }
    case "foo bar":
        v11 := GenS2{F0: "\n\t", F1: (v1 * v1), F2: v5}
        _ = v11
    case "hello":
        var v12 uint16 = (v5 << ((0) & 63))
        _ = v12
    }
    var v13 uint16 = (v5 + v5)
    if true {
        if (uint8(v13) != v7) {
            for fi14 := 0; fi14 < 2; fi14++ {
                for fi15 := 0; fi15 < 7; fi15++ {
                    switch "a" {
                    case "foo bar":
                        v7 += 128
                    }
                    if (42 != v7) {
                        switch ("foo bar" + ("foo bar" + "foo bar")) {
                        case "\n\t":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        case "":
                            v5 += v5
                            v5 -= (^(v13))
                            v16 := GenS1{F0: 1, F1: 2}
                            _ = v16
                        }
                        if len(v2) > 0 {
                            v2[0] = genFn0(42)
                        }
                        switch "\n\t" {
                        case "foo bar":
                            var v17 uint32 = 294
                            _ = v17
                        case "hello":
                            if true {
                                delete(v4, 1)
                                if (fi15 != (fi14 + (-(fi14)))) {
                                    if (true && (0 < 2)) {
                                        fmt.Println(v5, v13, v3)
                                    } else {
                                        var v18 string = ""
                                        switch "a" {
                                        case "a":
                                            v13 += 129
                                        case "\n\t":
                                            var v19 int32 = 42
                                            _ = v19
                                        default:
                                            v13 += v13
                                            if (1.5 < 1e10) {
                                                _ = len(v2)
                                                v3 = (v3 + (v3 + v3))
                                            }
                                            if (!true) {
                                                v2 = append(v2, 786)
                                                if (!(!true)) {
                                                    switch ((v18 + "") + "a") {
                                                    case "\n\t":
                                                        v20 := GenS3{F0: 2, F1: 0, F2: (fi14 >= fi14)}
                                                        _ = v20
                                                    case "":
                                                        v2 = append(v2, 4294967295)
                                                    case "hello":
                                                        v13 += 255
                                                    }
                                                    fmt.Println(v5)
                                                } else {
                                                    v13 = (v13 + 675)
                                                    fmt.Println(v5, v5, v7)
                                                    var v21 float32 = 0.1
                                                    v22 := GenS3{F0: 255, F1: 438, F2: true}
                                                    var v23 int8 = 100
                                                    _ = v21
                                                    _ = v22
                                                    _ = v23
                                                }
                                                v24 := GenS2{F0: v18, F1: v1, F2: (v5 << ((825) & 63))}
                                                _ = v24
                                            }
                                        }
                                        _ = v18
                                    }
                                    var v25 uint = 42
                                    var v26 int8 = 1
                                    _ = v25
                                    _ = v26
                                }
                                var v27 uint8 = (^(v7))
                                _ = v27
                            } else {
                                switch ("foo bar" + ("foo bar" + "a")) {
                                case "a":
                                    if (17 != -1) {
                                        v28 := v4[65535]
                                        v5 += 563
                                        _ = v28
                                    } else {
                                        v7 = (v7 >> ((0) & 63))
                                        var v29 bool = true
                                        _ = v29
                                    }
                                case "hello":
                                    v30 := GenS2{F0: "foo bar", F1: v1, F2: 42}
                                    _ = v30
                                }
                                if len(v2) > 0 {
                                    v2[0] = 4294967295
                                }
                                fmt.Println(fi15)
                            }
                        case "a":
                            v31 := GenS1{F0: -128, F1: 4294967295}
                            _ = v31
                        }
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi15
                }
                _ = fi14
            }
            var v32 int = 42
            _ = v32
        }
    } else {
        var v33 float32 = 1e10
        _ = v33
    }
    v2 = append(v2, 2147483647)
    if ((v7 + (v7 | 1)) >= v7) {
        v7 += v7
        v13 -= v5
        for fi34 := 0; fi34 < 10; fi34++ {
            fmt.Println(v7, v13)
            v35 := v4[genFn0(v3)]
            var v36 int16 = (v3 ^ (v3 - (v3 + (v3 + v3))))
            _ = fi34
            _ = v35
            _ = v36
        }
    } else {
        v37 := GenS2{F0: "", F1: (v1 & (v1 & v1)), F2: (v13 + v13)}
        var v38 uint = 1
        if true {
            var v39 uint64 = 4294967295
            var v40 int64 = (-(v1))
            fmt.Println(v39)
            _ = v39
            _ = v40
        }
        _ = v37
        _ = v38
    }
    v7 -= v7
    var v41 string = "foo bar"
    for fi42 := 0; fi42 < 7; fi42++ {
        switch (fi42 / ((fi42) | 1)) {
        case 127:
            for fi43 := 0; fi43 < 20; fi43++ {
                v1 -= (v1 / ((v1) | 1))
                v5 += v5
                var v44 int16 = 42
                var v45 int32 = 0
                v46 := GenS3{F0: 4294967295, F1: 140, F2: (!(0 < 65535))}
                _ = fi43
                _ = v44
                _ = v45
                _ = v46
            }
        case 32767:
            v5 -= 180
        case -9223372036854775808:
            switch (("a" + v41) + (v41 + (v41 + ((((v41 + ("" + (("" + "foo bar") + v41))) + v41) + (("\n\t" + v41) + (v41 + "foo bar"))) + ("foo bar" + v41))))) {
            case "\n\t":
                if true {
                    fmt.Println(v1, v7, v13)
                    var v47 uint32 = 84
                    switch (fi42 ^ fi42) {
                    case 623:
                        for fi48 := 0; fi48 < 19; fi48++ {
                            v13 = v13
                            var v49 uint8 = (v7 * v7)
                            _ = fi48
                            _ = v49
                        }
                    case 127:
                        if false {
                            v1 += (v1 % ((0) | 1))
                        } else {
                            v41 = (v41 + ((((v41 + "foo bar") + v41) + (v41 + v41)) + (v41 + ("a" + (("foo bar" + v41) + "foo bar")))))
                            var v50 int64 = (v1 % (((v1 % (((v1 | -266)) | 1))) | 1))
                            switch fi42 {
                            case 9223372036854775807:
                                v1 = v50
                            case 127:
                                if (v41 >= v41) {
                                    v3 = 32767
                                    v1 -= (v50 << ((1) & 63))
                                } else {
                                    v3 = 0
                                    for fi51 := 0; fi51 < 3; fi51++ {
                                        if (-2.25 <= (0.0 * 1.5)) {
                                            v52 := GenS2{F0: v41, F1: int64(fi42), F2: (v5 << ((991) & 63))}
                                            _ = v52
                                        } else {
                                            v3 = v3
                                            v53 := GenS1{F0: -128, F1: 42}
                                            v1 = 696
                                            if ((v1 << ((801) & 63)) > (v1 * (-(v1)))) {
                                                v13 = 42
                                                _ = len(v2)
                                                v2 = append(v2, 2)
                                            }
                                            v54 := GenS4{F0: uint(v5)}
                                            var v55 bool = false
                                            switch 128 {
                                            case -61:
                                                v41 = v41
                                            default:
                                                var v56 float64 = 0.0
                                                _ = v56
                                            }
                                            _ = v53
                                            _ = v54
                                            _ = v55
                                        }
                                        if false {
                                            if false {
                                                _ = len(v4)
                                                v13 -= v13
                                            }
                                            if (-102 == 7) {
                                                switch 859 {
                                                case 0:
                                                    var v57 bool = true
                                                    v57 = ((v13 * v5) == (v5 & (v5 + v13)))
                                                    v7 -= (v7 & v7)
                                                    _ = v57
                                                }
                                            } else {
                                                v4[4294967295] = 127
                                            }
                                        }
                                        _ = fi51
                                    }
                                    var v58 int64 = v1
                                    _ = v58
                                }
                            case -781:
                                v59 := GenS2{F0: v41, F1: v50, F2: uint16(v7)}
                                _ = v59
                            default:
                                var v60 float64 = (1.5 + -2.25)
                                _ = v60
                            }
                            _ = v50
                        }
                    case 781:
                        var v61 int = 0
                        _ = v61
                    }
                    _ = v47
                } else {
                    if true {
                        fmt.Println(v1, v3, v5)
                    } else {
                        for fi62 := 0; fi62 < 2; fi62++ {
                            var v63 uint8 = (^(v7))
                            v41 = (v41 + ((v41 + "") + ""))
                            v64 := GenS4{F0: 110}
                            _ = fi62
                            _ = v63
                            _ = v64
                        }
                    }
                    v3 -= v3
                }
            case "a":
                switch v41 {
                case "a":
                    var v65 int64 = v1
                    _ = v65
                default:
                    var v66 int64 = -9223372036854775808
                    _ = v66
                }
            }
        }
        _ = fi42
    }
    v1 += (v1 * -1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v5
    _ = v7
}
