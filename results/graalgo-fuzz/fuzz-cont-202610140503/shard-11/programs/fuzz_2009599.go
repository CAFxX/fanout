package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint16; F1 uint8; F2 string }
type GenS7 struct { F0 int }
type GenS8 struct { F0 bool; F1 int; F2 uint }

func main() {
    var v1 int = -1
    v2 := make(map[int]string)
    var v3 uint32 = 0
    v4 := make(map[string]bool)
    v5 := GenS1{F0: 42, F1: -32768}
    fmt.Println(v3)
    fmt.Println(v3, v1)
    v6 := []int{v1}
    switch 127 {
    case 4294967295:
        v7 := GenS1{F0: 42, F1: 32767}
        _ = v7
    case 32767:
        v1 -= v1
    default:
        switch (("hello" + (("hello" + "") + "hello")) + "hello") {
        case "a":
            var v8 int16 = 32767
            _ = v8
        default:
            v3 += 2147483648
        }
    }
    v9 := GenS1{F0: 250, F1: 32767}
    switch (("" + ("\n\t" + ("\n\t" + "\n\t"))) + "hello") {
    case "foo bar":
        if true {
            if true {
                v1 -= 9223372036854775807
                switch ("a" + "a") {
                case "hello":
                    fmt.Println(v1, v3)
                case "":
                    if false {
                        v3 = v3
                        for fi10 := 0; fi10 < 7; fi10++ {
                            for fi11 := 0; fi11 < 3; fi11++ {
                                if (!(127 != -95)) {
                                    v3 -= 2147483648
                                    v1 = (fi11 % (((v1 & v1)) | 1))
                                    v12 := GenS2{F0: 2}
                                    _ = v12
                                }
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if true {
                                    fmt.Println(v3, fi10, fi11)
                                } else {
                                    var v13 int64 = 28
                                    fmt.Println(v3, v3, fi10)
                                    var v14 bool = ((527 == 32767) && (((3.14159 >= 1.5) && (false && (1 <= 108))) && (0.1 >= 0.0)))
                                    _ = v13
                                    _ = v14
                                }
                                _ = fi11
                            }
                            var v15 uint64 = 18446744073709551615
                            _ = fi10
                            _ = v15
                        }
                        switch (v1 * v1) {
                        case 32767:
                            fmt.Println(v3, v1)
                        case -1:
                            v3 += 42
                        case 0:
                            var v16 int8 = 107
                            _ = v16
                        }
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        v17 := GenS3{F0: 0}
                        var v18 int8 = 0
                        _ = v17
                        _ = v18
                    }
                default:
                    fmt.Println(v3, v1)
                }
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    var v20 string = ("" + ("a" + "hello"))
                    var v21 int16 = 1
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        } else {
            switch 128 {
            case -714:
                var v22 string = "hello"
                _ = v22
            case 42:
                if (!false) {
                    v23 := GenS1{F0: 2, F1: int16(v1)}
                    fmt.Println(v1)
                    _ = v23
                } else {
                    v1 += v1
                    var v24 int16 = -32768
                    _ = v24
                }
            }
            v3 += v3
            var v25 int = (v1 / ((-926) | 1))
            _ = v25
        }
    case "\n\t":
        v26 := GenS4{F0: 4294967295, F1: -1}
        _ = v26
    case "a":
        v27 := GenS3{F0: 2}
        _ = v27
    default:
        switch "\n\t" {
        case "":
            if ((-123 > 100) || true) {
                v1 = v1
                var v28 int64 = 1
                _ = v28
            } else {
                fmt.Println(v3, v3)
                v29 := GenS3{F0: 32767}
                _ = v29
            }
        case "\n\t":
            var v30 float64 = -2.25
            _ = v30
        case "a":
            for fi31 := 0; fi31 < 18; fi31++ {
                var v32 float64 = (0.0 / 3.14159)
                if (0 == 76) {
                    v33 := v4["hello"]
                    switch "a" {
                    case "foo bar":
                        for fi34 := 0; fi34 < 2; fi34++ {
                            v3 = (v3 >> ((710) & 63))
                            _ = fi34
                        }
                    case "":
                        var v35 int8 = 127
                        _ = v35
                    case "\n\t":
                        var v36 uint64 = 570
                        _ = v36
                    default:
                        if v33 {
                            switch fi31 {
                            case 127:
                                fmt.Println(v3, v33)
                            case -571:
                                if false {
                                    var v37 int32 = 394
                                    v1 += -325
                                    _ = v37
                                }
                            }
                            var v38 string = "a"
                            var v39 uint16 = 973
                            _ = v38
                            _ = v39
                        } else {
                            fmt.Println(fi31, fi31, fi31)
                        }
                    }
                    _ = v33
                } else {
                    if (false && true) {
                        var v40 uint32 = (v3 | 0)
                        var v41 uint = 351
                        delete(v4, ("" + "\n\t"))
                        _ = v40
                        _ = v41
                    } else {
                        v32 = (0.0 - 1e10)
                        if false {
                            var v42 uint8 = 106
                            if false {
                                v2[-298] = ("a" + "a")
                                var v43 uint64 = 18446744073709551615
                                _ = v43
                            }
                            var v44 uint32 = v3
                            _ = v42
                            _ = v44
                        } else {
                            var v45 uint16 = 255
                            for fi46 := 0; fi46 < 3; fi46++ {
                                _ = len(v4)
                                var v47 bool = false
                                _ = fi46
                                _ = v47
                            }
                            _ = v45
                        }
                    }
                    v48 := GenS5{F0: 42}
                    _ = v48
                }
                var v49 string = "foo bar"
                _ = fi31
                _ = v32
                _ = v49
            }
        default:
            v50 := GenS5{F0: -599}
            _ = v50
        }
    }
    v3 = (^(v3))
    fmt.Println(v1, v3)
    switch "" {
    case "a":
        switch "\n\t" {
        case "\n\t":
            if (1 < 42) {
                v1 -= 128
                if len(v6) > 0 {
                    v6[0] = (v1 & -234)
                }
                if (1.5 > 3.14159) {
                    v3 -= (^(v3))
                    var v51 int8 = 100
                    if true {
                        switch ("hello" + ("\n\t" + "hello")) {
                        case "\n\t":
                            delete(v2, v1)
                        }
                    }
                    _ = v51
                }
            }
        case "":
            var v52 int32 = -1
            _ = v52
        case "foo bar":
            for fi53 := 0; fi53 < 8; fi53++ {
                var v54 int64 = -1
                for fi55 := 0; fi55 < 4; fi55++ {
                    var v56 string = ((("\n\t" + "foo bar") + "") + (("a" + "hello") + "\n\t"))
                    if true {
                        if (32767 > 0) {
                            var v57 int8 = -1
                            _ = v57
                        }
                    }
                    v58 := GenS3{F0: 42}
                    _ = fi55
                    _ = v56
                    _ = v58
                }
                _ = fi53
                _ = v54
            }
        }
    case "":
        fmt.Println(v1)
    case "foo bar":
        var v59 bool = true
        _ = v59
    default:
        var v60 int16 = 42
        _ = v60
    }
    v1 -= 32767
    if false {
        var v61 uint16 = 223
        v62 := GenS6{F0: 42, F1: 166, F2: "\n\t"}
        var v63 string = ("a" + "hello")
        _ = v61
        _ = v62
        _ = v63
    } else {
        _ = len(v6)
        var v64 uint32 = v3
        for fi65 := 0; fi65 < 3; fi65++ {
            if (!true) {
                switch (v1 << ((255) & 63)) {
                case 32768:
                    v6 = append(v6, v1)
                case 0:
                    v66 := GenS5{F0: 1}
                    _ = v66
                case -29:
                    v3 -= 2147483648
                default:
                    var v67 uint16 = 835
                    _ = v67
                }
                var v68 int16 = -44
                var v69 uint8 = 127
                _ = v68
                _ = v69
            }
            switch v1 {
            case 127:
                switch v1 {
                case -2147483648:
                    if true {
                        if (2 >= 312) {
                            var v70 float32 = (-2.25 * 0.0)
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            _ = v70
                        } else {
                            switch v1 {
                            case 42:
                                fmt.Println(fi65, v1)
                            case -324:
                                var v71 bool = false
                                _ = v71
                            }
                            var v72 string = "foo bar"
                            _ = v72
                        }
                        v3 = (v64 >> ((4294967295) & 63))
                        v73 := GenS3{F0: 65535}
                        _ = v73
                    } else {
                        switch ((("\n\t" + "foo bar") + "\n\t") + "foo bar") {
                        case "\n\t":
                            fmt.Println(v1, v64)
                        }
                        var v74 float64 = (-2.25 + 3.14159)
                        switch v1 {
                        case -9223372036854775808:
                            for fi75 := 0; fi75 < 18; fi75++ {
                                v4["hello"] = true
                                if len(v6) > 0 {
                                    v6[0] = (fi65 ^ (v1 >> ((91) & 63)))
                                }
                                v6 = append(v6, v1)
                                _ = fi75
                            }
                        case 42:
                            if ((false || false) && (100 != 100)) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                var v76 float64 = 3.14159
                                switch (fi65 % ((127) | 1)) {
                                case 128:
                                    switch "hello" {
                                    case "":
                                        var v77 uint32 = (v3 ^ (v3 / (((v64 + 2147483648)) | 1)))
                                        _ = v77
                                    case "\n\t":
                                        switch "foo bar" {
                                        case "hello":
                                            v78 := GenS6{F0: 2, F1: 128, F2: "foo bar"}
                                            _ = v78
                                        case "foo bar":
                                            for fi79 := 0; fi79 < 6; fi79++ {
                                                v6 = append(v6, fi65)
                                                _ = fi79
                                            }
                                        }
                                    case "a":
                                        v76 = v76
                                    default:
                                        switch -2147483648 {
                                        case 32768:
                                            v80 := GenS7{F0: fi65}
                                            _ = v80
                                        case 624:
                                            v3 -= v3
                                        case -202:
                                            fmt.Println(v74, v74)
                                        default:
                                            fmt.Println(v1, v74)
                                        }
                                    }
                                }
                                _ = v76
                            } else {
                                fmt.Println(v3)
                            }
                        case 0:
                            fmt.Println(fi65, v74)
                        default:
                            fmt.Println(v64)
                        }
                        _ = v74
                    }
                case 128:
                    var v81 int8 = 42
                    _ = v81
                case -193:
                    switch fi65 {
                    case 621:
                        v82 := GenS8{F0: false, F1: (v1 & 128), F2: 255}
                        _ = v82
                    case 32768:
                        v83 := GenS7{F0: -9223372036854775808}
                        _ = v83
                    }
                }
            case 32768:
                delete(v4, "hello")
            default:
                v64 = 2147483647
            }
            _ = fi65
        }
        _ = v64
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
