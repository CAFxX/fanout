package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int32; F1 int32 }
type GenS3 struct { F0 int64; F1 int64; F2 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 string }
type GenS7 struct { F0 uint64; F1 int; F2 uint64 }
func genFn0(p0 string, p1 int32, p2 int8) int64 {
    var v7 uint16 = 0
    fmt.Println(p2)
    v7 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return -9223372036854775808
}

func main() {
    var v1 uint = 42
    v2 := []uint32{42}
    var v3 uint64 = 209
    var v4 uint16 = 933
    v5 := []string{"\n\t", ("hello" + "\n\t")}
    v6 := make(map[string]string)
    switch -253 {
    case 32767:
        if ((0.0 == 3.14159) || true) {
            v8 := GenS1{F0: (v4 & (v4 + 623)), F1: 374}
            for fi9 := 0; fi9 < 2; fi9++ {
                _ = len(v2)
                _ = fi9
            }
            _ = v8
        }
    default:
        v6["\n\t"] = "a"
    }
    fmt.Println(v1)
    v6["hello"] = "\n\t"
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 int = fi10
        _ = fi10
        _ = v11
    }
    if true {
        v12 := GenS2{F0: 2147483647, F1: -1}
        v3 -= (v3 & 4294967295)
        var v13 bool = ((1 != 42) || ("foo bar" == (("foo bar" + "foo bar") + "foo bar")))
        _ = v12
        _ = v13
    } else {
        v4 = 65535
        var v14 int8 = 26
        fmt.Println(v14)
        _ = v14
    }
    v3 = v3
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 float64 = 1.5
            switch ("hello" + (("a" + "foo bar") + "a")) {
            case "foo bar":
                v18 := GenS2{F0: -298, F1: -2147483648}
                _ = v18
            default:
                fmt.Println(fi16)
            }
            _ = fi16
            _ = v17
        }
        var v19 uint16 = v4
        if ((v1 - (v1 - (v1 - 65535))) != 336) {
            if true {
                for fi20 := 0; fi20 < 5; fi20++ {
                    if (((v1 != (v1 ^ (^(v1)))) && (!true)) || true) {
                        if (int16(fi15) != 260) {
                            switch fi15 {
                            case 32768:
                                if true {
                                    var v21 uint16 = (^(v19))
                                    fmt.Println(v19)
                                    v4 = (v4 + v19)
                                    _ = v21
                                }
                            case 0:
                                v22 := GenS2{F0: 1, F1: 40}
                                _ = v22
                            }
                        } else {
                            var v23 uint16 = v19
                            v4 = (^(v23))
                            _ = v23
                        }
                        switch "\n\t" {
                        case "\n\t":
                            switch ("hello" + ((("" + "") + "hello") + "a")) {
                            case "a":
                                v1 = v1
                            case "hello":
                                v19 = v4
                                var v24 string = "\n\t"
                                var v25 int8 = 100
                                _ = v24
                                _ = v25
                            case "foo bar":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            }
                        }
                        if (65535 != v1) {
                            v6[("\n\t" + "hello")] = "\n\t"
                        } else {
                            fmt.Println(fi20)
                            v19 -= 255
                            if (0 <= 109) {
                                _ = len(v5)
                            }
                        }
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi20
                }
                v26 := GenS3{F0: 9223372036854775807, F1: -1, F2: 376}
                fmt.Println(v4, v4, v4)
                _ = v26
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v1 = uint(v19)
        } else {
            if false {
                if (!(992 <= (v3 << ((v1) & 63)))) {
                    var v27 int64 = 206
                    var v28 int8 = 42
                    switch "\n\t" {
                    case "\n\t":
                        v1 += v1
                    case "foo bar":
                        if true {
                            v29 := GenS1{F0: 384, F1: -945}
                            _ = v29
                        } else {
                            fmt.Println(fi15, v27)
                            v19 = v4
                        }
                    }
                    _ = v27
                    _ = v28
                }
                var v30 string = ""
                _ = v30
            }
            v3 += v3
        }
        _ = fi15
        _ = v19
    }
    for fi31 := 0; fi31 < 10; fi31++ {
        v32 := GenS3{F0: genFn0(((("foo bar" + ("hello" + "hello")) + "\n\t") + "hello"), 2147483647, 127), F1: 42, F2: -525}
        v33 := GenS4{F0: 565}
        v34 := GenS4{F0: (v4 ^ (v4 % ((v4) | 1)))}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    v4 -= 32767
    if (!false) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if ((18446744073709551615 > (v3 << ((v1) & 63))) && (9223372036854775807 <= -1)) {
            var v35 uint16 = 42
            switch 4294967295 {
            case 9223372036854775807:
                if true {
                    v35 += 65535
                    switch "hello" {
                    case "foo bar":
                        switch -889 {
                        case -2147483648:
                            var v36 int32 = -2147483648
                            _ = v36
                        default:
                            v35 = 65535
                        }
                    default:
                        v2 = append(v2, 689)
                    }
                    for fi37 := 0; fi37 < 11; fi37++ {
                        switch (fi37 & (fi37 % (((^(fi37))) | 1))) {
                        case 0:
                            _ = len(v2)
                        case -9223372036854775808:
                            var v38 bool = ((-1 != 1000) && (-1 >= -1))
                            _ = v38
                        default:
                            if true {
                                for fi39 := 0; fi39 < 3; fi39++ {
                                    v40 := GenS5{F0: -1, F1: 42}
                                    switch "\n\t" {
                                    case "hello":
                                        v41 := v6["hello"]
                                        _ = v41
                                    case "foo bar":
                                        var v42 uint = v1
                                        _ = v42
                                    default:
                                        fmt.Println(fi37)
                                    }
                                    if (!false) {
                                        fmt.Println(fi37, fi39)
                                        var v43 int = (^(fi39))
                                        var v44 string = "a"
                                        _ = v43
                                        _ = v44
                                    }
                                    v35 = (v35 >> ((v1) & 63))
                                    if true {
                                        fmt.Println(v1, v35, v1)
                                        v45 := v6["hello"]
                                        _ = v45
                                    }
                                    _ = fi39
                                    _ = v40
                                }
                                fmt.Println(v35, v4, v4)
                                v5 = append(v5, "foo bar")
                            } else {
                                v1 = v1
                                for fi46 := 0; fi46 < 7; fi46++ {
                                    if len(v5) > 0 {
                                        v5[0] = ("" + "foo bar")
                                    }
                                    _ = fi46
                                }
                                for fi47 := 0; fi47 < 19; fi47++ {
                                    v35 += v4
                                    v48 := v6["hello"]
                                    var v49 bool = (-53 == 1)
                                    _ = fi47
                                    _ = v48
                                    _ = v49
                                }
                            }
                        }
                        var v50 float32 = ((3.14159 * 0.0) * 3.14159)
                        _ = fi37
                        _ = v50
                    }
                } else {
                    for fi51 := 0; fi51 < 11; fi51++ {
                        v3 += 0
                        v6["a"] = "a"
                        v4 = 32767
                        _ = fi51
                    }
                    v52 := GenS6{F0: "\n\t"}
                    for fi53 := 0; fi53 < 3; fi53++ {
                        if true {
                            switch ("\n\t" + ("hello" + (("hello" + ("a" + "a")) + "foo bar"))) {
                            case "":
                                v5 = append(v5, "hello")
                            case "\n\t":
                                var v54 float32 = (((float32((0.0 * (((((1.5 / 3.14159) + -2.25) * 1e10) + 1e10) - ((3.14159 * (3.14159 - float64((-2.25 + 1.5)))) * -2.25)))) * (3.14159 / 1e10)) - (0.0 * 0.0)) * 0.1)
                                _ = v54
                            case "foo bar":
                                for fi55 := 0; fi55 < 19; fi55++ {
                                    var v56 uint = v1
                                    _ = fi55
                                    _ = v56
                                }
                            }
                            v3 = v3
                            var v57 uint64 = 1
                            _ = v57
                        } else {
                            v58 := GenS7{F0: v3, F1: (^(fi53)), F2: (^(v3))}
                            if false {
                                v59 := GenS7{F0: 357, F1: fi53, F2: 758}
                                _ = v59
                            }
                            _ = len(v5)
                            _ = v58
                        }
                        v35 += 41
                        _ = fi53
                    }
                    _ = v52
                }
            case -265:
                if ((^(v3)) < 0) {
                    v60 := GenS3{F0: 69, F1: -1, F2: 1}
                    _ = v60
                }
            }
            var v61 int16 = int16(v1)
            _ = v35
            _ = v61
        }
        var v62 int16 = 42
        _ = v62
    } else {
        var v63 uint8 = 128
        _ = v63
    }
    fmt.Println(v4, v3, v4)
    switch 9223372036854775807 {
    case 9223372036854775807:
        var v64 uint64 = (v3 - 337)
        _ = v64
    case 127:
        v3 -= (v3 % ((939) | 1))
    }
    switch "hello" {
    case "hello":
        var v65 int64 = -418
        _ = v65
    default:
        var v66 uint = v1
        _ = v66
    }
    if ((!false) || ((48 <= 42) || false)) {
        v67 := v6[("" + "hello")]
        v4 = v4
        v68 := GenS5{F0: 122, F1: -226}
        _ = v67
        _ = v68
    } else {
        v4 += (v4 | (v4 * v4))
        v4 -= 255
        if true {
            v1 += v1
            var v69 uint16 = (v4 / ((255) | 1))
            if true {
                fmt.Println(v1, v69, v3)
                v69 = 123
                for fi70 := 0; fi70 < 3; fi70++ {
                    v3 = (v3 & (v3 & v3))
                    _ = fi70
                }
            }
            _ = v69
        }
    }
    v71 := GenS7{F0: 0, F1: -826, F2: (v3 - v3)}
    for fi72 := 0; fi72 < 5; fi72++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v73 := v6[((((("foo bar" + "a") + (("hello" + "foo bar") + "")) + "") + (("hello" + ("" + (("hello" + "a") + "foo bar"))) + ("foo bar" + ("" + "\n\t")))) + ("a" + "a"))]
        var v74 uint64 = v3
        _ = fi72
        _ = v73
        _ = v74
    }
    v2 = append(v2, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v71
}
