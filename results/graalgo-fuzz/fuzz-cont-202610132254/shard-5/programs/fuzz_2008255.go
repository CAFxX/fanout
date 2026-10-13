package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 string; F1 int64; F2 uint64 }
type GenS5 struct { F0 uint; F1 bool; F2 uint }
func genFn0(p0 int, p1 string) bool {
    v7 := GenS1{F0: 0, F1: true}
    p0 = -9223372036854775808
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 float32 = 0.0
    var v4 int32 = 1
    v5 := []string{"\n\t", "\n\t", "foo bar"}
    var v6 uint = 2
    for fi8 := 0; fi8 < 19; fi8++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 uint64 = 0
            var v11 bool = v1
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    if v1 {
        delete(v2, 4294967295)
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 bool = (938 > 38)
            _ = fi12
            _ = v13
        }
    } else {
        v14 := GenS2{F0: 42}
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(v6, v6)
            fmt.Println(v6, v4, fi15)
            _ = fi15
        }
        _ = v14
    }
    if ((226 == 128) || (v1 || (0 <= -1))) {
        _ = len(v2)
        v16 := GenS1{F0: 897, F1: (!(1 >= 42))}
        _ = v16
    }
    v17 := GenS2{F0: 1}
    v1 = (!v1)
    v18 := GenS2{F0: 65535}
    v6 += v6
    v1 = false
    switch "" {
    case "foo bar":
        switch -1 {
        case -1:
            var v19 string = "a"
            _ = v19
        default:
            if true {
                v2[2147483647] = -1
                switch 32768 {
                case -751:
                    if (v4 == (^(v4))) {
                        for fi20 := 0; fi20 < 17; fi20++ {
                            v5 = append(v5, "")
                            if false {
                                switch fi20 {
                                case -9223372036854775808:
                                    v21 := GenS3{F0: "foo bar", F1: 127}
                                    _ = v21
                                case -515:
                                    v22 := GenS4{F0: ("" + "\n\t"), F1: 42, F2: 4294967295}
                                    _ = v22
                                }
                                fmt.Println(v1, v6, v1)
                            } else {
                                v3 += v3
                                v23 := v2[2]
                                _ = v23
                            }
                            _ = fi20
                        }
                        fmt.Println(v1)
                        fmt.Println(v3, v1, v1)
                    } else {
                        for fi24 := 0; fi24 < 20; fi24++ {
                            switch (fi24 >> ((v6) & 63)) {
                            case 9223372036854775807:
                                if (2 >= 62) {
                                    var v25 float32 = ((1e10 - -2.25) + v3)
                                    switch (((("foo bar" + "") + ("" + "\n\t")) + "a") + "foo bar") {
                                    case "a":
                                        for fi26 := 0; fi26 < 9; fi26++ {
                                            v25 -= (1e10 + 1.5)
                                            _ = fi26
                                        }
                                    }
                                    switch (fi24 | (fi24 * (^(fi24)))) {
                                    case -9223372036854775808:
                                        switch ((("hello" + "hello") + "\n\t") + (("foo bar" + "") + "")) {
                                        case "hello":
                                            v27 := GenS5{F0: (v6 ^ (v6 << ((42) & 63))), F1: v1, F2: v6}
                                            _ = v27
                                        case "a":
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        case "foo bar":
                                            v3 = v25
                                        }
                                    case 42:
                                        var v28 int16 = -18
                                        _ = v28
                                    }
                                    _ = v25
                                } else {
                                    _ = len(v2)
                                    v1 = v1
                                }
                            case 32767:
                                switch fi24 {
                                case 510:
                                    for fi29 := 0; fi29 < 5; fi29++ {
                                        if (-52 <= -1) {
                                            if len(v5) > 0 {
                                                v5[0] = ""
                                            }
                                            v6 -= 880
                                            if v1 {
                                                if (((float64(1e10) - float64(v3)) / 1.5) != 1.5) {
                                                    v1 = (0.1 >= (-2.25 + (1.5 - 1e10)))
                                                    var v30 int64 = -257
                                                    if v1 {
                                                        var v31 uint32 = 42
                                                        _ = v31
                                                    } else {
                                                        v4 += 1
                                                        v3 = float32(((0.1 + 1e10) * -2.25))
                                                    }
                                                    _ = v30
                                                }
                                                v6 += v6
                                            }
                                            v6 += v6
                                            v1 = v1
                                        }
                                        _ = fi29
                                    }
                                case -2147483648:
                                    for fi32 := 0; fi32 < 11; fi32++ {
                                        var v33 uint = v6
                                        _ = fi32
                                        _ = v33
                                    }
                                case -63:
                                    switch (fi24 >> ((v6) & 63)) {
                                    case 4294967295:
                                        v6 = v6
                                    }
                                }
                            default:
                                var v34 int16 = 42
                                _ = v34
                            }
                            _ = fi24
                        }
                        for fi35 := 0; fi35 < 10; fi35++ {
                            var v36 int8 = -28
                            var v37 int8 = (v36 | (v36 / ((-128) | 1)))
                            _ = fi35
                            _ = v36
                            _ = v37
                        }
                        for fi38 := 0; fi38 < 20; fi38++ {
                            var v39 bool = v1
                            var v40 uint32 = 221
                            var v41 float64 = 3.14159
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    }
                case 128:
                    var v42 int8 = 127
                    _ = v42
                case 0:
                    var v43 uint16 = 32767
                    _ = v43
                default:
                    v44 := GenS2{F0: 0}
                    _ = v44
                }
                if v1 {
                    v45 := GenS3{F0: "a", F1: -104}
                    _ = v45
                } else {
                    switch "hello" {
                    case "\n\t":
                        var v46 float32 = (float32(1.5) - v3)
                        _ = v46
                    case "":
                        var v47 int32 = v4
                        _ = v47
                    }
                }
            } else {
                v6 = (v6 % ((v6) | 1))
                var v48 string = "\n\t"
                var v49 string = "a"
                _ = v48
                _ = v49
            }
        }
    case "":
        for fi50 := 0; fi50 < 19; fi50++ {
            for fi51 := 0; fi51 < 13; fi51++ {
                var v52 int32 = v4
                if v1 {
                    var v53 string = ("\n\t" + "foo bar")
                    var v54 float64 = 1.5
                    _ = v53
                    _ = v54
                }
                fmt.Println(fi51, v6)
                _ = fi51
                _ = v52
            }
            v55 := GenS1{F0: 4294967295, F1: (-9223372036854775808 == -252)}
            _ = fi50
            _ = v55
        }
    case "a":
        var v56 float64 = float64((v3 + ((((1.5 + (((v3 / (v3 / -2.25)) * v3) - float32(3.14159))) / 1.5) - 3.14159) * ((1.5 + ((3.14159 + (((v3 * v3) * 0.1) * v3)) + v3)) * ((float32(-2.25) / v3) - float32(3.14159))))))
        _ = v56
    default:
        _ = len(v5)
    }
    for fi57 := 0; fi57 < 20; fi57++ {
        fmt.Println(v1, v3, fi57)
        _ = fi57
    }
    var v58 string = "a"
    var v59 uint32 = 65535
    _ = len(v2)
    v60 := GenS4{F0: v58, F1: 9223372036854775807, F2: 972}
    var v61 int8 = 100
    var v62 uint = (^(v6))
    switch "a" {
    case "":
        var v63 int16 = 678
        _ = v63
    case "hello":
        delete(v2, 1)
    }
    for fi64 := 0; fi64 < 12; fi64++ {
        v4 += v4
        var v65 float64 = -2.25
        switch v58 {
        case "hello":
            switch "foo bar" {
            case "a":
                v59 += v59
            case "foo bar":
                v66 := GenS1{F0: 0, F1: v1}
                _ = v66
            case "":
                for fi67 := 0; fi67 < 2; fi67++ {
                    if len(v5) > 0 {
                        v5[0] = ("\n\t" + ("\n\t" + (v58 + "hello")))
                    }
                    switch v58 {
                    case "":
                        if (false || (!true)) {
                            var v68 int8 = -1
                            if v1 {
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                fmt.Println(v68)
                            } else {
                                var v69 string = ((v58 + "foo bar") + v58)
                                if len(v5) > 0 {
                                    v5[0] = "a"
                                }
                                var v70 int16 = -199
                                var v71 string = v69
                                _ = v69
                                _ = v70
                                _ = v71
                            }
                            var v72 uint8 = 55
                            switch "" {
                            case "a":
                                fmt.Println(v1, v1)
                                if false {
                                    v59 = v59
                                    fmt.Println(v4, fi64)
                                    v73 := GenS3{F0: v58, F1: (^(v68))}
                                    _ = v73
                                } else {
                                    var v74 float32 = (v3 + (1.5 + -2.25))
                                    if v1 {
                                        v75 := GenS2{F0: 872}
                                        var v76 float64 = (0.0 - v65)
                                        _ = v75
                                        _ = v76
                                    } else {
                                        v74 += v3
                                        switch fi67 {
                                        case -472:
                                            _ = len(v5)
                                        case -172:
                                            v62 -= v62
                                        case 1:
                                            switch 32768 {
                                            case 852:
                                                var v77 int8 = int8(v4)
                                                _ = v77
                                            case 9223372036854775807:
                                                var v78 uint16 = 255
                                                _ = v78
                                            case 2147483647:
                                                v1 = (v1 || v1)
                                            }
                                        }
                                        v79 := v2[2147483647]
                                        _ = v79
                                    }
                                    _ = v74
                                }
                                v80 := GenS2{F0: 32767}
                                _ = v80
                            case "":
                                var v81 int32 = (v4 | -1)
                                _ = v81
                            case "\n\t":
                                _ = len(v2)
                            default:
                                var v82 uint8 = (v72 >> ((v62) & 63))
                                _ = v82
                            }
                            fmt.Println(v6)
                            _ = v68
                            _ = v72
                        }
                    case "\n\t":
                        v83 := GenS2{F0: 460}
                        _ = v83
                    }
                    switch v58 {
                    case "hello":
                        v84 := GenS4{F0: ("a" + ("hello" + v58)), F1: int64(fi64), F2: 0}
                        _ = v84
                    }
                    var v85 string = v58
                    _ = fi67
                    _ = v85
                }
            default:
                v5 = append(v5, v58)
            }
        }
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v60
    _ = v61
    _ = v62
}
