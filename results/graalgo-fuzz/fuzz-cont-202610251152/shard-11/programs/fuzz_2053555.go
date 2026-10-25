package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 uint }
type GenS7 struct { F0 bool; F1 int64; F2 string }
type GenS8 struct { F0 uint32; F1 bool; F2 bool }
func genFn0(p0 string, p1 int16) int8 {
    switch 139 {
    case -1:
        for fi6 := 0; fi6 < 16; fi6++ {
            var v7 uint32 = 842
            if (v7 != v7) {
                for fi8 := 0; fi8 < 11; fi8++ {
                    fmt.Println(p1)
                    var v9 uint64 = 0
                    _ = fi8
                    _ = v9
                }
                for fi10 := 0; fi10 < 18; fi10++ {
                    v7 = v7
                    v11 := []int{(fi10 << ((4294967295) & 63)), fi10, fi6, fi6}
                    _ = fi10
                    _ = v11
                }
            } else {
                switch p0 {
                case "\n\t":
                    var v12 int16 = p1
                    _ = v12
                default:
                    v7 += v7
                }
            }
            fmt.Println(p1, fi6)
            _ = fi6
            _ = v7
        }
    case 32768:
        if (4294967295 != 42) {
            switch -1 {
            case 32767:
                for fi13 := 0; fi13 < 2; fi13++ {
                    fmt.Println(p1)
                    v14 := GenS1{F0: p1, F1: 852, F2: 127}
                    _ = fi13
                    _ = v14
                }
            case -1:
                v15 := GenS1{F0: p1, F1: 65535, F2: 210}
                _ = v15
            case 2147483647:
                fmt.Println(p1, p1)
            default:
                fmt.Println(p0, p1)
            }
            switch p0 {
            case "\n\t":
                p0 = ("\n\t" + p0)
            case "a":
                var v16 int8 = -64
                _ = v16
            case "foo bar":
                v17 := []int{1, 656, 2147483647}
                _ = v17
            }
            v18 := []int{-2147483648, 119, 32767, -880}
            _ = v18
        } else {
            p0 = ((p0 + p0) + "\n\t")
            fmt.Println(p0, p1)
        }
    }
    v19 := GenS1{F0: p1, F1: 1, F2: 2}
    _ = p0
    _ = p1
    _ = v19
    return 127
}
func genFn1(p0 string) string {
    if true {
        switch 128 {
        case 4294967295:
            p0 = "hello"
        case 259:
            var v20 uint8 = 1
            _ = v20
        case 42:
            if false {
                v21 := GenS1{F0: -32768, F1: 2147483648, F2: 1}
                v22 := []int{-116, -9223372036854775808}
                for fi23 := 0; fi23 < 3; fi23++ {
                    p0 = "\n\t"
                    fmt.Println(fi23, fi23)
                    _ = fi23
                }
                _ = v21
                _ = v22
            }
        }
        v24 := []int{-847}
        _ = v24
    }
    switch 0 {
    case -1:
        var v25 uint32 = 565
        _ = v25
    case -761:
        v26 := []int{2147483647}
        _ = v26
    case -578:
        switch p0 {
        case "":
            if false {
                p0 = ""
                for fi27 := 0; fi27 < 7; fi27++ {
                    fmt.Println(p0)
                    _ = fi27
                }
                var v28 int64 = 0
                _ = v28
            } else {
                p0 = (p0 + "hello")
            }
        default:
            v29 := GenS1{F0: 0, F1: 42, F2: 255}
            _ = v29
        }
    }
    var v30 string = (p0 + p0)
    fmt.Println(v30, v30)
    switch p0 {
    case "\n\t":
        p0 = "\n\t"
    }
    _ = p0
    _ = v30
    return p0
}

func main() {
    v1 := []int{32768}
    var v2 uint64 = 200
    var v3 float32 = ((1.5 - float32(((((float64(1.5) + 0.0) + -2.25) / 1.0) - 1e10))) - 0.0)
    v4 := []string{("foo bar" + ""), "foo bar", "", "\n\t"}
    v5 := make(map[string]int)
    var v31 int32 = 2147483647
    v32 := GenS2{F0: 4294967295}
    v3 = float32(0.0)
    for fi33 := 0; fi33 < 11; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            var v35 float64 = 1.5
            var v36 uint16 = 2
            v36 += (^(v36))
            var v37 float64 = (float64((float32(0.0) - v3)) + (v35 + v35))
            var v38 int = (fi33 << ((1) & 63))
            _ = fi34
            _ = v35
            _ = v36
            _ = v37
            _ = v38
        }
        for fi39 := 0; fi39 < 14; fi39++ {
            var v40 float64 = float64(float32((((3.14159 + (((1.5 / 1e10) - 1.5) + 1.5)) * 1e10) - (-2.25 / float64((3.14159 + (-2.25 - v3)))))))
            _ = fi39
            _ = v40
        }
        v41 := GenS3{F0: 0, F1: -9223372036854775808}
        _ = fi33
        _ = v41
    }
    switch "\n\t" {
    case "\n\t":
        for fi42 := 0; fi42 < 13; fi42++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v43 := GenS4{F0: true}
            _ = fi42
            _ = v43
        }
    }
    v2 -= (v2 << ((2) & 63))
    v2 -= v2
    var v44 uint8 = 128
    switch "" {
    case "a":
        v31 = v31
    case "\n\t":
        v31 -= (v31 | -554)
    }
    var v45 float64 = 3.14159
    v46 := GenS3{F0: 127, F1: -1}
    if false {
        if true {
            if len(v4) > 0 {
                v4[0] = ("" + (("hello" + "\n\t") + (("a" + ((("a" + "foo bar") + "hello") + "foo bar")) + "\n\t")))
            }
            v47 := GenS5{F0: 0}
            _ = v47
        } else {
            fmt.Println(v2, v3, v45)
            for fi48 := 0; fi48 < 13; fi48++ {
                v45 = -2.25
                v49 := GenS6{F0: 800, F1: 255}
                _ = fi48
                _ = v49
            }
        }
    }
    v44 = (v44 - 127)
    var v50 uint16 = uint16(v2)
    fmt.Println(v2, v50, v44)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch ("a" + ("\n\t" + "a")) {
    case "foo bar":
        switch "" {
        case "foo bar":
            switch ("" + "foo bar") {
            case "foo bar":
                var v51 bool = (v31 != -2147483648)
                _ = v51
            case "\n\t":
                var v52 bool = true
                _ = v52
            case "hello":
                v53 := GenS7{F0: false, F1: 0, F2: ("\n\t" + "hello")}
                _ = v53
            default:
                switch 42 {
                case -33:
                    fmt.Println(v44, v50, v50)
                case 32767:
                    if ((0 == -1) && (true && (true || true))) {
                        var v54 int16 = 42
                        if len(v1) > 0 {
                            v1[0] = 128
                        }
                        var v55 int32 = v31
                        _ = v54
                        _ = v55
                    } else {
                        switch "foo bar" {
                        case "a":
                            v44 = (v44 % ((v44) | 1))
                        case "":
                            v44 -= 42
                        }
                        var v56 uint16 = v50
                        for fi57 := 0; fi57 < 16; fi57++ {
                            for fi58 := 0; fi58 < 20; fi58++ {
                                v56 = v50
                                if false {
                                    v2 = v2
                                } else {
                                    v2 += v2
                                    v3 -= -2.25
                                    if false {
                                        v2 -= (^(v2))
                                        v2 -= 244
                                        var v59 float32 = (v3 - v3)
                                        delete(v5, "a")
                                        var v60 int32 = (^(v31))
                                        _ = v59
                                        _ = v60
                                    } else {
                                        var v61 bool = ((^(fi57)) > -1)
                                        _ = v61
                                    }
                                }
                                v45 = 1.5
                                _ = fi58
                            }
                            v56 -= 32767
                            _ = fi57
                        }
                        _ = v56
                    }
                default:
                    for fi62 := 0; fi62 < 3; fi62++ {
                        if true {
                            var v63 int = fi62
                            _ = v63
                        }
                        _ = fi62
                    }
                }
            }
        case "":
            switch 32767 {
            case 32768:
                v2 += (v2 * (^(v2)))
            case 1:
                var v64 float64 = (v45 / v45)
                _ = v64
            default:
                v65 := GenS8{F0: 861, F1: (true && (false || (!false))), F2: (43 <= 341)}
                _ = v65
            }
        }
    case "\n\t":
        var v66 uint64 = uint64(v50)
        _ = v66
    default:
        if ((!false) && (v2 == uint64(v44))) {
            fmt.Println(v45, v45)
        }
    }
    if (!true) {
        v2 = (v2 << ((18446744073709551615) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v50
}
