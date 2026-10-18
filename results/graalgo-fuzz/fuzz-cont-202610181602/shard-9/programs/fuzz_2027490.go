package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64; F1 int64; F2 uint8 }
type GenS5 struct { F0 int8; F1 uint }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint16; F1 uint; F2 bool }
type GenS8 struct { F0 string }
type GenS9 struct { F0 int }
type GenS10 struct { F0 int16 }
func genFn0(p0 uint, p1 uint16, p2 uint16) bool {
    v4 := []int{127, 871}
    p0 = p0
    var v5 uint8 = 1
    p2 = uint16(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return (("hello" + ((("hello" + ("hello" + ("a" + "hello"))) + "\n\t") + "foo bar")) >= "foo bar")
}
func genFn1(p0 uint32, p1 int32) uint64 {
    if false {
        for fi6 := 0; fi6 < 9; fi6++ {
            v7 := GenS1{F0: 42, F1: 255}
            var v8 int64 = 42
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        fmt.Println(p0, p1)
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(fi9)
        _ = fi9
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        p1 = p1
        if (785 < 65535) {
            v11 := GenS1{F0: 9223372036854775807, F1: 1}
            for fi12 := 0; fi12 < 10; fi12++ {
                p1 += 164
                var v13 uint64 = 1
                _ = fi12
                _ = v13
            }
            _ = v11
        }
        _ = fi10
    }
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := []int{fi14, fi14}
        _ = len(v15)
        var v16 string = "a"
        _ = fi14
        _ = v15
        _ = v16
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    var v1 float64 = (1.5 - (1.5 - -2.25))
    v2 := make(map[uint32]uint32)
    v3 := []uint32{270, 2147483647, 4294967295}
    v17 := v2[1]
    fmt.Println(v17, v1)
    switch ("foo bar" + ("\n\t" + (("" + "\n\t") + "a"))) {
    case "foo bar":
        for fi18 := 0; fi18 < 2; fi18++ {
            switch 42 {
            case -1:
                v19 := GenS2{F0: "\n\t"}
                _ = v19
            case 127:
                v20 := GenS1{F0: -554, F1: 933}
                _ = v20
            default:
                for fi21 := 0; fi21 < 16; fi21++ {
                    switch (fi18 + (fi21 * (fi18 | fi18))) {
                    case -339:
                        fmt.Println(v1, fi21, fi21)
                    case 42:
                        if (-23 <= 833) {
                            fmt.Println(v17)
                            fmt.Println(v17, v1)
                            v17 = (v17 << ((4294967295) & 63))
                            fmt.Println(fi21, fi18, v17)
                            v17 += (v17 & 209)
                            fmt.Println(fi21, v17)
                        } else {
                            v22 := GenS3{F0: 1}
                            _ = v22
                        }
                    default:
                        v23 := GenS4{F0: -9223372036854775808, F1: -9223372036854775808, F2: 127}
                        _ = v23
                    }
                    v17 = (v17 * (^(v17)))
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    _ = fi21
                }
            }
            for fi24 := 0; fi24 < 18; fi24++ {
                v1 = 3.14159
                fmt.Println(fi24)
                var v25 uint64 = 4294967295
                var v26 bool = true
                if true {
                    var v27 int64 = 355
                    _ = v27
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            v28 := GenS5{F0: 1, F1: 770}
            _ = fi18
            _ = v28
        }
    case "hello":
        if true {
            v29 := GenS4{F0: 851, F1: -1, F2: 1}
            v30 := GenS6{F0: "foo bar"}
            if (((1 > 0) || true) || false) {
                for fi31 := 0; fi31 < 1; fi31++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v32 string = "hello"
                    _ = fi31
                    _ = v32
                }
                switch "foo bar" {
                case "a":
                    var v33 string = "hello"
                    _ = v33
                default:
                    v1 = v1
                }
                var v34 string = ("hello" + "a")
                _ = v34
            } else {
                if ((!(19 > 42)) || (-128 > 1)) {
                    fmt.Println(v1, v17)
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v1 -= v1
                    }
                } else {
                    v35 := GenS5{F0: -17, F1: 18446744073709551615}
                    _ = v35
                }
                var v36 bool = (138 != 1)
                var v37 uint16 = 2
                _ = v36
                _ = v37
            }
            _ = v29
            _ = v30
        }
    }
    switch "a" {
    case "foo bar":
        for fi38 := 0; fi38 < 2; fi38++ {
            v17 += (v17 | (v17 + v17))
            v1 = (v1 * (0.0 - v1))
            v1 += float64(0.0)
            _ = fi38
        }
    case "hello":
        if (!(137 != 18446744073709551615)) {
            var v39 uint8 = 84
            v40 := GenS2{F0: "hello"}
            _ = v39
            _ = v40
        }
    case "a":
        switch -9223372036854775808 {
        case -522:
            v41 := GenS6{F0: (("foo bar" + "hello") + ("\n\t" + "hello"))}
            _ = v41
        default:
            if false {
                var v42 uint32 = (v17 / ((0) | 1))
                if len(v3) > 0 {
                    v3[0] = v42
                }
                _ = v42
            } else {
                var v43 string = "\n\t"
                _ = v43
            }
        }
    default:
        fmt.Println(v17, v17)
    }
    for fi44 := 0; fi44 < 4; fi44++ {
        fmt.Println(v1, fi44, fi44)
        v1 = v1
        v17 = v17
        _ = fi44
    }
    for fi45 := 0; fi45 < 11; fi45++ {
        for fi46 := 0; fi46 < 6; fi46++ {
            if (0 == 58) {
                switch "foo bar" {
                case "\n\t":
                    if (!(!(("a" <= "hello") && false))) {
                        if (100 <= 1) {
                            v47 := GenS3{F0: 42}
                            fmt.Println(fi45, fi46, v17)
                            _ = v47
                        }
                        v1 = (v1 * (v1 * (v1 + (v1 - v1))))
                    } else {
                        v1 -= (v1 / ((v1 - (float64(3.14159) + (((v1 / 1e10) - v1) / v1))) + -2.25))
                        v48 := GenS6{F0: "foo bar"}
                        _ = v48
                    }
                default:
                    if (!true) {
                        switch -2147483648 {
                        case 0:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            if (!(-680 < 698)) {
                                if (120 == 128) {
                                    switch fi46 {
                                    case -994:
                                        if (!true) {
                                            switch 128 {
                                            case 127:
                                                switch "a" {
                                                case "foo bar":
                                                    v1 = 0.1
                                                    var v49 float64 = v1
                                                    _ = v49
                                                case "hello":
                                                    var v50 uint32 = (v17 + 42)
                                                    _ = v50
                                                }
                                            }
                                            var v51 int32 = -349
                                            v52 := GenS7{F0: 32767, F1: 0, F2: (!((0.1 - (0.0 * 0.1)) != (0.1 / 1.5)))}
                                            _ = v51
                                            _ = v52
                                        }
                                    case 42:
                                        v1 -= v1
                                    case 4294967295:
                                        switch ("hello" + "") {
                                        case "foo bar":
                                            fmt.Println(v17, fi45, fi45)
                                        case "":
                                            if (!(2147483647 > 42)) {
                                                var v53 int = (fi45 * fi46)
                                                var v54 uint8 = 127
                                                _ = v53
                                                _ = v54
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = 220
                                                }
                                                v55 := v2[v17]
                                                v56 := GenS8{F0: ""}
                                                _ = v55
                                                _ = v56
                                            }
                                        default:
                                            v17 = (v17 >> ((18446744073709551615) & 63))
                                        }
                                    default:
                                        v1 += float64(0.0)
                                    }
                                    v57 := GenS9{F0: (fi45 % ((800) | 1))}
                                    if (-1 != 2147483647) {
                                        v2[266] = (^(v17))
                                        switch (fi46 & (fi46 & (fi46 / (((fi46 / ((32767) | 1))) | 1)))) {
                                        case -833:
                                            _ = len(v2)
                                        default:
                                            fmt.Println(fi46)
                                        }
                                    }
                                    _ = v57
                                }
                                v3 = append(v3, (v17 & v17))
                            }
                            if ((true || (829 != 65535)) || (255 > 0)) {
                                v1 += ((1e10 - v1) * v1)
                            }
                        }
                        v3 = append(v3, v17)
                    }
                }
                var v58 uint64 = 280
                _ = v58
            } else {
                var v59 bool = (1 > -615)
                fmt.Println(v1)
                var v60 int = -2147483648
                _ = v59
                _ = v60
            }
            fmt.Println(v1, fi45, v17)
            fmt.Println(v17, fi45, v1)
            _ = fi46
        }
        v61 := GenS10{F0: 0}
        if (!false) {
            switch "\n\t" {
            case "hello":
                for fi62 := 0; fi62 < 2; fi62++ {
                    var v63 int8 = 1
                    fmt.Println(fi45)
                    _ = fi62
                    _ = v63
                }
            case "":
                var v64 uint64 = 42
                _ = v64
            default:
                v17 = (v17 * (v17 - v17))
            }
        } else {
            for fi65 := 0; fi65 < 3; fi65++ {
                switch fi65 {
                case -1:
                    switch (fi65 / ((-2147483648) | 1)) {
                    case 632:
                        v17 = v17
                    case 9223372036854775807:
                        v17 = (v17 - 2147483647)
                    case 404:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    default:
                        var v66 uint32 = v17
                        _ = v66
                    }
                }
                _ = fi65
            }
            switch "a" {
            case "a":
                var v67 uint64 = 0
                _ = v67
            default:
                if true {
                    var v68 uint = 65535
                    for fi69 := 0; fi69 < 8; fi69++ {
                        if len(v3) > 0 {
                            v3[0] = (v17 | 0)
                        }
                        _ = fi69
                    }
                    _ = v68
                } else {
                    var v70 int64 = 9223372036854775807
                    fmt.Println(v70, v17, fi45)
                    var v71 int64 = (v70 % ((v70) | 1))
                    _ = v70
                    _ = v71
                }
            }
            fmt.Println(fi45, fi45, v17)
        }
        _ = fi45
        _ = v61
    }
    switch "" {
    case "foo bar":
        v72 := v2[v17]
        _ = v72
    case "hello":
        var v73 uint = 255
        _ = v73
    case "a":
        v17 = v17
    default:
        var v74 int16 = -32768
        _ = v74
    }
    var v75 float32 = (float32((-2.25 * (((-2.25 - v1) / v1) / v1))) + 1.5)
    v76 := GenS1{F0: 9223372036854775807, F1: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v75:", v75)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v75
    _ = v76
}
