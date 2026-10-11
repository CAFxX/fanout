package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 int; F2 int }
type GenS8 struct { F0 int64 }
func genFn0(p0 uint16, p1 int16) string {
    if true {
        var v5 bool = false
        v6 := GenS1{F0: 117, F1: 2, F2: 42}
        _ = v5
        _ = v6
    } else {
        if (48 != 68) {
            v7 := GenS1{F0: -1, F1: 4294967295, F2: p0}
            var v8 uint = 213
            _ = v7
            _ = v8
        } else {
            p1 -= p1
        }
    }
    var v9 int64 = 739
    p1 -= (p1 * p1)
    for fi10 := 0; fi10 < 20; fi10++ {
        v9 -= (v9 & 9223372036854775807)
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v9
    return "a"
}
func genFn1(p0 uint64, p1 int8, p2 int) uint64 {
    switch "\n\t" {
    case "hello":
        for fi11 := 0; fi11 < 15; fi11++ {
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 bool = (1.5 != (3.14159 / 1.5))
                fmt.Println(p2, p0, fi11)
                var v14 int8 = (-(p1))
                fmt.Println(p1)
                if false {
                    switch (p2 | (^(fi12))) {
                    case -500:
                        switch "foo bar" {
                        case "foo bar":
                            var v15 int32 = 694
                            _ = v15
                        case "hello":
                            v14 = (v14 % ((v14) | 1))
                        default:
                            if (0.0 >= (0.1 * 1.5)) {
                                v13 = (!v13)
                                var v16 int16 = 1000
                                if ("hello" == "hello") {
                                    v17 := GenS2{F0: 255, F1: "hello"}
                                    v18 := []int{}
                                    fmt.Println(v13, p0)
                                    _ = v17
                                    _ = v18
                                }
                                _ = v16
                            }
                        }
                    case 42:
                        var v19 float64 = (-2.25 * 0.1)
                        _ = v19
                    case 424:
                        var v20 uint64 = (p0 >> ((255) & 63))
                        _ = v20
                    default:
                        fmt.Println(fi12)
                    }
                    v21 := GenS3{F0: (!true), F1: (^(p0))}
                    _ = v21
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            fmt.Println(p2, p2)
            _ = fi11
        }
    case "\n\t":
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 int32 = 42
            _ = fi22
            _ = v23
        }
    default:
        if true {
            p2 += (p2 << ((18446744073709551615) & 63))
            var v24 uint = 2
            fmt.Println(v24)
            _ = v24
        }
    }
    v25 := GenS4{F0: 32767, F1: -403}
    fmt.Println(p2)
    var v26 int16 = -1
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    _ = v26
    return p0
}

func main() {
    v1 := []string{"", "\n\t"}
    var v2 string = "foo bar"
    var v3 uint8 = 255
    var v4 uint16 = 547
    fmt.Println(v2, v3)
    if false {
        switch -9223372036854775808 {
        case 832:
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 string = ((v2 + v2) + v2)
                v1 = append(v1, ("foo bar" + (v2 + v2)))
                v4 = v4
                _ = fi27
                _ = v28
            }
        case 2147483647:
            var v29 float32 = 0.1
            _ = v29
        case 127:
            switch -844 {
            case 0:
                var v30 uint8 = 0
                _ = v30
            case 265:
                for fi31 := 0; fi31 < 10; fi31++ {
                    fmt.Println(v2, fi31, v2)
                    var v32 float64 = 0.0
                    _ = fi31
                    _ = v32
                }
            case 9223372036854775807:
                if ((^(v4)) <= 796) {
                    v4 += 878
                } else {
                    if false {
                        for fi33 := 0; fi33 < 3; fi33++ {
                            var v34 uint16 = (v4 | v4)
                            _ = fi33
                            _ = v34
                        }
                        v4 = (v4 % (((v4 ^ v4)) | 1))
                        v4 -= 1
                    } else {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    v2 = "a"
                }
            }
        }
        if (v3 > 46) {
            v35 := GenS5{F0: 42, F1: (!(65535 <= 18446744073709551615))}
            switch "hello" {
            case "foo bar":
                if (406 == -729) {
                    v2 = v2
                } else {
                    v1 = append(v1, v2)
                }
            case "\n\t":
                if (0 > -1) {
                    switch ((v2 + "\n\t") + v2) {
                    case "foo bar":
                        if false {
                            fmt.Println(v4, v4)
                        } else {
                            v4 -= 134
                        }
                    default:
                        v1 = append(v1, v2)
                    }
                    for fi36 := 0; fi36 < 19; fi36++ {
                        if ((0 <= 18446744073709551615) && ((v4 == (v4 + v4)) || false)) {
                            switch "foo bar" {
                            case "a":
                                v37 := GenS1{F0: 100, F1: 42, F2: v4}
                                _ = v37
                            case "":
                                fmt.Println(fi36)
                            case "\n\t":
                                fmt.Println(v4, fi36, v3)
                            default:
                                switch (fi36 & fi36) {
                                case -2147483648:
                                    var v38 uint64 = 4294967295
                                    _ = v38
                                case -859:
                                    var v39 bool = ((v4 << ((0) & 63)) < (v4 << ((4294967295) & 63)))
                                    _ = v39
                                case 21:
                                    v40 := GenS6{F0: false}
                                    _ = v40
                                default:
                                    for fi41 := 0; fi41 < 18; fi41++ {
                                        var v42 bool = true
                                        _ = fi41
                                        _ = v42
                                    }
                                }
                            }
                            if len(v1) > 0 {
                                v1[0] = ("" + ("" + v2))
                            }
                        } else {
                            var v43 float32 = (-2.25 - ((1e10 - 1.5) * 0.1))
                            v44 := GenS4{F0: (v4 ^ 42), F1: -817}
                            _ = v43
                            _ = v44
                        }
                        _ = fi36
                    }
                } else {
                    var v45 int32 = 42
                    for fi46 := 0; fi46 < 16; fi46++ {
                        var v47 uint32 = 65535
                        var v48 float64 = float64((1e10 + 0.1))
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    fmt.Println(v3)
                    _ = v45
                }
            case "a":
                if (9223372036854775807 > 42) {
                    if ((2147483647 > -347) || (2147483647 > 65535)) {
                        v49 := GenS4{F0: (v4 + (v4 ^ v4)), F1: 2147483647}
                        v50 := GenS7{F0: (v2 + (("hello" + ((("hello" + "\n\t") + (("\n\t" + ((v2 + "foo bar") + v2)) + "\n\t")) + v2)) + v2)), F1: 128, F2: -587}
                        _ = v49
                        _ = v50
                    } else {
                        for fi51 := 0; fi51 < 1; fi51++ {
                            v3 = (v3 & (v3 & (v3 | 128)))
                            v2 = "a"
                            v2 = ("" + v2)
                            _ = fi51
                        }
                        for fi52 := 0; fi52 < 20; fi52++ {
                            var v53 int16 = -707
                            if false {
                                fmt.Println(v2, fi52)
                            } else {
                                fmt.Println(v2, v3)
                                for fi54 := 0; fi54 < 14; fi54++ {
                                    v1 = append(v1, ("foo bar" + "a"))
                                    var v55 float32 = 3.14159
                                    _ = fi54
                                    _ = v55
                                }
                                if len(v1) > 0 {
                                    v1[0] = "\n\t"
                                }
                            }
                            _ = fi52
                            _ = v53
                        }
                        v1 = append(v1, v2)
                    }
                }
            default:
                v4 += v4
            }
            _ = v35
        }
    } else {
        var v56 float64 = 0.1
        var v57 uint8 = v3
        var v58 int16 = 1000
        _ = v56
        _ = v57
        _ = v58
    }
    var v59 uint32 = 2147483647
    fmt.Println(v3)
    v60 := GenS1{F0: -1, F1: genFn1(300, 127, -128), F2: 255}
    if false {
        var v61 uint16 = (^(v4))
        if true {
            for fi62 := 0; fi62 < 8; fi62++ {
                v63 := GenS2{F0: (^(v61)), F1: v2}
                _ = fi62
                _ = v63
            }
        }
        _ = v61
    }
    if false {
        switch "\n\t" {
        case "hello":
            v4 = 65535
        case "a":
            switch v2 {
            case "":
                v59 = v59
            case "hello":
                v64 := GenS8{F0: 0}
                _ = v64
            default:
                v2 = ""
            }
        case "foo bar":
            v3 += v3
        default:
            var v65 float32 = -2.25
            _ = v65
        }
        v4 += 2
    }
    for fi66 := 0; fi66 < 19; fi66++ {
        v2 = v2
        var v67 uint64 = 42
        switch (v2 + v2) {
        case "\n\t":
            var v68 float32 = 0.0
            _ = v68
        }
        _ = fi66
        _ = v67
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v60
}
