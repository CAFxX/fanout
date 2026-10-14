package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int16; F1 int }
type GenS6 struct { F0 uint8; F1 string }
type GenS7 struct { F0 int16; F1 uint32; F2 int }
type GenS8 struct { F0 uint; F1 bool }
func genFn0(p0 int32) int8 {
    v5 := []int{9223372036854775807, -9223372036854775808}
    if ((p0 ^ (p0 >> ((42) & 63))) != p0) {
        var v6 int = 318
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v6
    }
    _ = len(v5)
    for fi7 := 0; fi7 < 5; fi7++ {
        if true {
            switch "a" {
            case "hello":
                var v8 bool = false
                _ = v8
            default:
                switch "" {
                case "foo bar":
                    for fi9 := 0; fi9 < 20; fi9++ {
                        var v10 bool = true
                        p0 = (^(p0))
                        if false {
                            switch (("" + ("hello" + "hello")) + "a") {
                            case "\n\t":
                                switch "\n\t" {
                                case "foo bar":
                                    fmt.Println(v10, p0)
                                case "a":
                                    var v11 uint16 = 0
                                    _ = v11
                                case "\n\t":
                                    var v12 string = "hello"
                                    _ = v12
                                default:
                                    p0 = p0
                                }
                            }
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        _ = fi9
                        _ = v10
                    }
                default:
                    switch "foo bar" {
                    case "a":
                        var v13 float32 = 0.0
                        _ = v13
                    case "hello":
                        var v14 string = "\n\t"
                        _ = v14
                    }
                }
            }
            if (false || (-1 < 1)) {
                fmt.Println(p0)
            } else {
                var v15 uint32 = 2147483647
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS1{F0: 4294967295}
                    v15 += v15
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            switch 9223372036854775807 {
            case 32767:
                if ("\n\t" <= ("" + "\n\t")) {
                    if false {
                        if (32767 == -32768) {
                            p0 = -1
                            for fi18 := 0; fi18 < 13; fi18++ {
                                if len(v5) > 0 {
                                    v5[0] = fi18
                                }
                                _ = fi18
                            }
                        } else {
                            for fi19 := 0; fi19 < 1; fi19++ {
                                p0 = (p0 & (p0 | (p0 & p0)))
                                fmt.Println(fi19, fi7)
                                var v20 float32 = 3.14159
                                _ = fi19
                                _ = v20
                            }
                        }
                        var v21 string = ("a" + "a")
                        _ = v21
                    } else {
                        p0 = p0
                        v22 := GenS2{F0: "foo bar", F1: 95, F2: 160}
                        _ = v22
                    }
                    var v23 int16 = 1000
                    _ = v23
                } else {
                    v24 := GenS2{F0: ("foo bar" + ("\n\t" + "foo bar")), F1: 1, F2: 909}
                    _ = v24
                }
            case 2147483647:
                p0 = p0
            }
        } else {
            var v25 uint32 = 220
            fmt.Println(v25)
            _ = v25
        }
        _ = fi7
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 int64 = -9223372036854775808
    var v3 uint16 = 32767
    v4 := []int{42}
    _ = len(v1)
    var v26 uint64 = 45
    var v27 float64 = (float64(1.5) + 1e10)
    v4 = append(v4, -2147483648)
    var v28 int16 = 1000
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v3, v28)
    default:
        for fi29 := 0; fi29 < 19; fi29++ {
            fmt.Println(v26)
            v30 := GenS3{F0: (v26 + (v26 >> ((18446744073709551615) & 63)))}
            _ = fi29
            _ = v30
        }
    }
    if (32767 > v3) {
        v2 += (v2 - (^(v2)))
        for fi31 := 0; fi31 < 18; fi31++ {
            var v32 uint32 = 42
            _ = fi31
            _ = v32
        }
        for fi33 := 0; fi33 < 8; fi33++ {
            for fi34 := 0; fi34 < 8; fi34++ {
                var v35 int = fi34
                v36 := GenS1{F0: 943}
                _ = fi34
                _ = v35
                _ = v36
            }
            switch (fi33 + fi33) {
            case -2147483648:
                var v37 float32 = ((0.0 + 1.5) * (1e10 + 1.5))
                _ = v37
            case 127:
                fmt.Println(v2, v26, v3)
            case -475:
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            v4 = append(v4, fi33)
            _ = fi33
        }
    }
    if ((v2 % ((int64(v3)) | 1)) >= (v2 << ((uint(v26)) & 63))) {
        switch 1 {
        case -2147483648:
            switch "" {
            case "hello":
                delete(v1, "hello")
            case "\n\t":
                for fi38 := 0; fi38 < 6; fi38++ {
                    if false {
                        delete(v1, ((("\n\t" + "\n\t") + "") + "hello"))
                        var v39 int32 = 0
                        if true {
                            var v40 int16 = (^(v28))
                            delete(v1, "\n\t")
                            _ = len(v1)
                            _ = v40
                        } else {
                            fmt.Println(fi38)
                        }
                        _ = v39
                    }
                    for fi41 := 0; fi41 < 18; fi41++ {
                        v42 := GenS4{F0: v2, F1: -44}
                        fmt.Println(fi41, v3, v2)
                        v2 = v2
                        if false {
                            v43 := GenS5{F0: (v28 & v28), F1: fi41}
                            _ = v43
                        } else {
                            v1[("foo bar" + ("\n\t" + "hello"))] = 42
                        }
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v28 += (v28 & v28)
                        default:
                            v26 += 4294967295
                            v44 := GenS2{F0: ("hello" + "\n\t"), F1: v3, F2: 297}
                            _ = v44
                        }
                        _ = fi41
                        _ = v42
                    }
                    var v45 int16 = v28
                    _ = fi38
                    _ = v45
                }
            case "":
                v2 += (^(v2))
            }
        case 588:
            var v46 uint8 = 255
            _ = v46
        case 1:
            fmt.Println(v2, v27)
        default:
            v47 := GenS1{F0: 4294967295}
            _ = v47
        }
        if false {
            if true {
                var v48 uint64 = (^(v26))
                if ("" != ("foo bar" + "")) {
                    var v49 int8 = 42
                    _ = v49
                } else {
                    if true {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        for fi50 := 0; fi50 < 2; fi50++ {
                            v51 := GenS6{F0: 1, F1: ("foo bar" + "hello")}
                            if (2147483647 < fi50) {
                                v52 := GenS7{F0: -920, F1: 738, F2: fi50}
                                var v53 int16 = (-(v28))
                                _ = v52
                                _ = v53
                            } else {
                                var v54 float32 = 0.0
                                for fi55 := 0; fi55 < 12; fi55++ {
                                    var v56 uint8 = 40
                                    switch "" {
                                    case "":
                                        var v57 float64 = (float64(float32(v27)) * v27)
                                        _ = v57
                                    case "a":
                                        v58 := GenS5{F0: (v28 % ((0) | 1)), F1: fi50}
                                        _ = v58
                                    default:
                                        var v59 bool = true
                                        _ = v59
                                    }
                                    _ = fi55
                                    _ = v56
                                }
                                _ = v54
                            }
                            _ = fi50
                            _ = v51
                        }
                    }
                    fmt.Println(v2, v2, v48)
                    v3 = (v3 * (v3 + 722))
                }
                v60 := v1["a"]
                _ = v48
                _ = v60
            }
            for fi61 := 0; fi61 < 9; fi61++ {
                var v62 int16 = (v28 * v28)
                if ((false && false) && (false || false)) {
                    switch fi61 {
                    case -182:
                        var v63 uint8 = 255
                        _ = v63
                    case -66:
                        fmt.Println(fi61)
                    }
                    v2 = v2
                    fmt.Println(v26, v2)
                } else {
                    v28 = 32767
                    var v64 int = -9223372036854775808
                    _ = v64
                }
                v65 := GenS8{F0: 65535, F1: (-9223372036854775808 >= fi61)}
                _ = fi61
                _ = v62
                _ = v65
            }
            switch ("\n\t" + "hello") {
            case "hello":
                fmt.Println(v28)
            }
        } else {
            v66 := GenS6{F0: 0, F1: "hello"}
            var v67 float32 = (float32(((((3.14159 - v27) + 3.14159) * (1.5 / float64((float32(((v27 - (v27 - -2.25)) + 1e10)) * 0.0)))) + 0.0)) - 3.14159)
            _ = v66
            _ = v67
        }
    } else {
        var v68 uint16 = v3
        for fi69 := 0; fi69 < 6; fi69++ {
            var v70 uint16 = v3
            var v71 int64 = 42
            v68 -= v70
            _ = fi69
            _ = v70
            _ = v71
        }
        _ = v68
    }
    _ = len(v4)
    v3 += v3
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
