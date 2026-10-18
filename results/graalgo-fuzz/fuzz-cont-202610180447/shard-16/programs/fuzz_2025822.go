package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 string }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 int }
type GenS7 struct { F0 int32; F1 uint64 }
type GenS8 struct { F0 string; F1 uint64 }
func genFn0(p0 int32, p1 int64) int64 {
    var v6 int16 = 1000
    var v7 int8 = 22
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", "hello", ("\n\t" + "\n\t"), ""}
    v3 := make(map[string]string)
    var v4 uint16 = 42
    v5 := make(map[string]int)
    fmt.Println(v4)
    if ((2 > 719) || false) {
        var v8 uint8 = 53
        fmt.Println(v8, v4)
        _ = v8
    } else {
        fmt.Println(v4)
    }
    if (!(-1 <= 1)) {
        _ = len(v2)
    }
    switch 449 {
    case 9223372036854775807:
        if (false && false) {
            v4 += 815
            fmt.Println(v4)
            var v9 uint = 18446744073709551615
            _ = v9
        } else {
            for fi10 := 0; fi10 < 10; fi10++ {
                v4 = v4
                v4 = (v4 & v4)
                _ = fi10
            }
            switch 32768 {
            case -2147483648:
                v11 := GenS1{F0: genFn0(315, 1)}
                _ = v11
            case 32767:
                var v12 int = -1
                _ = v12
            }
        }
    case 32767:
        var v13 uint64 = 18446744073709551615
        _ = v13
    case -109:
        v4 = v4
    default:
        v4 = 558
    }
    delete(v1, 42)
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v4, fi14)
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        v2 = append(v2, "a")
        _ = fi14
    }
    v4 = (v4 + 42)
    v4 = (v4 & 42)
    switch "foo bar" {
    case "\n\t":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    case "":
        if (!(v4 != 32767)) {
            v4 = (v4 - (^(v4)))
            if false {
                var v15 uint = 2
                if (!true) {
                    v15 += v15
                    v16 := GenS1{F0: -1}
                    _ = v16
                } else {
                    var v17 uint16 = 276
                    for fi18 := 0; fi18 < 3; fi18++ {
                        switch 128 {
                        case 42:
                            if ((false || false) && true) {
                                for fi19 := 0; fi19 < 4; fi19++ {
                                    v15 += (v15 ^ 2)
                                    delete(v1, fi19)
                                    _ = fi19
                                }
                            } else {
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if false {
                                        if false {
                                            if (!(true && (((false || true) && (0 == -619)) && ((742 > 65535) || false)))) {
                                                fmt.Println(v4, v4)
                                            }
                                        }
                                        if (!true) {
                                            fmt.Println(fi18, v17, v17)
                                            if (465 != uint32(v4)) {
                                                v17 -= v17
                                            }
                                        }
                                        var v21 uint64 = 42
                                        _ = v21
                                    }
                                    _ = fi20
                                }
                            }
                        case 0:
                            v22 := GenS2{F0: 2, F1: ("hello" + ("hello" + "\n\t"))}
                            _ = v22
                        default:
                            v17 = (v17 & v17)
                        }
                        if ((false || false) && false) {
                            var v23 uint16 = 277
                            for fi24 := 0; fi24 < 5; fi24++ {
                                if false {
                                    fmt.Println(v4, v23)
                                } else {
                                    if (42 == 42) {
                                        if (100 > -54) {
                                            if true {
                                                var v25 int16 = 32767
                                                var v26 uint64 = 1
                                                v26 = (v26 ^ (^(v26)))
                                                _ = v25
                                                _ = v26
                                            } else {
                                                switch fi24 {
                                                case 127:
                                                    switch (fi18 & fi24) {
                                                    case 4294967295:
                                                        var v27 int32 = -893
                                                        _ = v27
                                                    case -2147483648:
                                                        if (fi24 <= -1) {
                                                            switch 42 {
                                                            case 42:
                                                                switch ("foo bar" + "foo bar") {
                                                                case "":
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "a"
                                                                    }
                                                                default:
                                                                    var v28 bool = (0.0 >= 1e10)
                                                                    _ = v28
                                                                }
                                                                switch (("hello" + "foo bar") + "hello") {
                                                                case "hello":
                                                                    if true {
                                                                        var v29 int16 = 42
                                                                        v30 := GenS2{F0: 64, F1: (("" + "hello") + ((("hello" + (((("a" + ("\n\t" + "")) + "hello") + "a") + (("foo bar" + "\n\t") + ""))) + "hello") + "\n\t"))}
                                                                        v31 := v5[(((("foo bar" + (((("" + "foo bar") + "") + ("\n\t" + "foo bar")) + ("foo bar" + ("hello" + ("\n\t" + ("a" + ("a" + "\n\t"))))))) + "a") + ("hello" + ("foo bar" + "hello"))) + "")]
                                                                        var v32 string = "a"
                                                                        _ = v29
                                                                        _ = v30
                                                                        _ = v31
                                                                        _ = v32
                                                                    } else {
                                                                        if (30 > 42) {
                                                                            v17 -= (^(v23))
                                                                        }
                                                                        v15 = v15
                                                                    }
                                                                case "\n\t":
                                                                    var v33 uint = v15
                                                                    _ = v33
                                                                case "a":
                                                                    v3["a"] = "foo bar"
                                                                }
                                                            case 32768:
                                                                switch (("foo bar" + "a") + (("\n\t" + ("foo bar" + ("a" + ("hello" + "hello")))) + "a")) {
                                                                case "hello":
                                                                    switch fi18 {
                                                                    case 127:
                                                                        var v34 int64 = 0
                                                                        _ = v34
                                                                    case 270:
                                                                        v2 = append(v2, "a")
                                                                    case 9223372036854775807:
                                                                        if (true && (fi18 < fi24)) {
                                                                            v35 := v3["hello"]
                                                                            _ = v35
                                                                        } else {
                                                                            var v36 int32 = 1
                                                                            if true {
                                                                                v17 -= 42
                                                                                v37 := GenS1{F0: genFn0(v36, 813)}
                                                                                v38 := GenS3{F0: genFn0(v36, 569), F1: false}
                                                                                _ = v37
                                                                                _ = v38
                                                                            } else {
                                                                                fmt.Println(v17)
                                                                                v4 += 42
                                                                                fmt.Println(v36)
                                                                            }
                                                                            _ = v36
                                                                        }
                                                                    default:
                                                                        v39 := GenS3{F0: -157, F1: false}
                                                                        _ = v39
                                                                    }
                                                                case "\n\t":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case "":
                                                                    v40 := v1[(fi18 ^ (fi18 * 1))]
                                                                    _ = v40
                                                                default:
                                                                    if true {
                                                                        v41 := GenS4{F0: 113}
                                                                        var v42 int32 = 1
                                                                        _ = v41
                                                                        _ = v42
                                                                    } else {
                                                                        var v43 uint32 = 2147483648
                                                                        var v44 int64 = 9223372036854775807
                                                                        _ = v43
                                                                        _ = v44
                                                                    }
                                                                }
                                                            case -1:
                                                                v15 = 0
                                                            default:
                                                                v45 := GenS3{F0: 0, F1: (-1 != 42)}
                                                                delete(v1, fi24)
                                                                _ = v45
                                                            }
                                                            if (v15 > (v15 & v15)) {
                                                                _ = len(v1)
                                                            } else {
                                                                fmt.Println(v4, v4, v17)
                                                                var v46 int16 = -1
                                                                _ = v46
                                                            }
                                                        } else {
                                                            var v47 uint16 = v4
                                                            v17 += v4
                                                            _ = v47
                                                        }
                                                    }
                                                }
                                            }
                                            var v48 uint16 = (^(v4))
                                            var v49 float64 = float64(float32(((-2.25 - 0.1) - 3.14159)))
                                            _ = v48
                                            _ = v49
                                        }
                                        fmt.Println(fi24)
                                    } else {
                                        fmt.Println(v15, fi24, v4)
                                        fmt.Println(v23, v4, fi24)
                                        v50 := v1[32768]
                                        v15 = (v15 % (((^(v15))) | 1))
                                        _ = v50
                                    }
                                    switch "\n\t" {
                                    case "a":
                                        var v51 int16 = 42
                                        _ = v51
                                    case "":
                                        switch -9223372036854775808 {
                                        case 264:
                                            switch "a" {
                                            case "hello":
                                                var v52 int64 = 0
                                                _ = v52
                                            case "foo bar":
                                                var v53 int16 = 1
                                                _ = v53
                                            case "\n\t":
                                                fmt.Println(fi18, v23, v15)
                                            default:
                                                v54 := GenS5{F0: 0}
                                                _ = v54
                                            }
                                        case 805:
                                            if false {
                                                v1[fi24] = false
                                            }
                                        }
                                    }
                                    v55 := GenS4{F0: 2}
                                    _ = v55
                                }
                                _ = fi24
                            }
                            _ = v23
                        } else {
                            v56 := GenS6{F0: 1000, F1: fi18}
                            var v57 int = (fi18 * (-(fi18)))
                            _ = v56
                            _ = v57
                        }
                        _ = fi18
                    }
                    _ = v17
                }
                if true {
                    fmt.Println(v4, v15)
                } else {
                    if false {
                        var v58 uint = (v15 << (((v15 + v15)) & 63))
                        var v59 string = ""
                        var v60 float64 = (1e10 * 3.14159)
                        _ = v58
                        _ = v59
                        _ = v60
                    } else {
                        fmt.Println(v4)
                        fmt.Println(v4)
                    }
                    for fi61 := 0; fi61 < 16; fi61++ {
                        v15 += v15
                        if true {
                            switch "\n\t" {
                            case "":
                                if (true || ((127 == -16) || (false && (1 == 1)))) {
                                    v62 := v1[987]
                                    v62 = (465 != 711)
                                    _ = v62
                                } else {
                                    for fi63 := 0; fi63 < 1; fi63++ {
                                        fmt.Println(v4, v15)
                                        v64 := GenS7{F0: 2147483647, F1: 18446744073709551615}
                                        _ = fi63
                                        _ = v64
                                    }
                                    if (((3.14159 + (0.1 / 1e10)) >= (float64(3.14159) - 3.14159)) || true) {
                                        v4 -= 42
                                        var v65 float64 = 3.14159
                                        fmt.Println(v4, v65)
                                        _ = v65
                                    } else {
                                        var v66 uint32 = 529
                                        if (42 != 277) {
                                            _ = len(v1)
                                            v1[0] = false
                                        } else {
                                            v66 = (^(v66))
                                            fmt.Println(v15, v66, v66)
                                            fmt.Println(fi61)
                                        }
                                        v4 -= (v4 % ((v4) | 1))
                                        _ = v66
                                    }
                                }
                            }
                            v4 = v4
                            for fi67 := 0; fi67 < 4; fi67++ {
                                var v68 string = "hello"
                                _ = fi67
                                _ = v68
                            }
                        } else {
                            switch (fi61 | 992) {
                            case 560:
                                var v69 float32 = float32(-2.25)
                                _ = v69
                            default:
                                var v70 float64 = ((0.0 + (1.5 / 3.14159)) * (1e10 / 1e10))
                                _ = v70
                            }
                            v71 := GenS6{F0: 303, F1: 32767}
                            _ = len(v2)
                            _ = v71
                        }
                        _ = fi61
                    }
                    if (0 >= -2147483648) {
                        v72 := GenS8{F0: (("\n\t" + ("a" + ((("\n\t" + "hello") + "foo bar") + "a"))) + "hello"), F1: 938}
                        var v73 uint16 = (v4 / ((v4) | 1))
                        v74 := GenS4{F0: 42}
                        _ = v72
                        _ = v73
                        _ = v74
                    }
                }
                _ = v15
            }
        }
    case "foo bar":
        fmt.Println(v4)
    }
    var v75 int64 = -41
    v76 := GenS7{F0: 1, F1: 2}
    v77 := GenS1{F0: 0}
    var v78 float64 = float64(((float32(3.14159) + 0.0) * ((-2.25 - float32(float64((1.5 * 1.5)))) * -2.25)))
    v78 = ((0.0 + (float64(1.5) * 1e10)) / (v78 / 0.1))
    v4 = (v4 * 255)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v75:", v75)
    fmt.Println("v78:", v78)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v75
    _ = v76
    _ = v77
    _ = v78
}
