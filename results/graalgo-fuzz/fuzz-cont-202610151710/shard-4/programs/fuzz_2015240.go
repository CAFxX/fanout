package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int8; F1 string }

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 2
    var v3 uint = 2
    v4 := make(map[string]bool)
    var v5 int = 9223372036854775807
    var v6 bool = (-2147483648 == -527)
    if v6 {
        v2 = (^(v2))
    } else {
        v7 := []int{(v5 << ((v3) & 63))}
        v2 = v2
        _ = v7
    }
    var v8 int8 = -128
    v9 := []int{v5, 401}
    v10 := GenS1{F0: ("" + ""), F1: -183}
    if (!false) {
        var v11 bool = v6
        _ = v11
    }
    fmt.Println(v8, v2)
    if len(v9) > 0 {
        v9[0] = (v5 ^ (v5 >> (((v3 | (v3 << ((2) & 63)))) & 63)))
    }
    fmt.Println(v5, v2, v6)
    if ((v6 || ((1 != v1) || v6)) && v6) {
        v3 = uint(v2)
        switch "foo bar" {
        case "a":
            if v6 {
                if false {
                    if (42 >= 586) {
                        v12 := GenS2{F0: (v5 + -390)}
                        for fi13 := 0; fi13 < 20; fi13++ {
                            for fi14 := 0; fi14 < 10; fi14++ {
                                v6 = v6
                                switch 927 {
                                case -1:
                                    if true {
                                        fmt.Println(v3, v3)
                                    } else {
                                        if v6 {
                                            v15 := GenS2{F0: (v5 % (((v5 >> (((^(v3))) & 63))) | 1))}
                                            v16 := GenS1{F0: "foo bar", F1: int16(fi13)}
                                            var v17 uint = (v3 | v3)
                                            v18 := GenS2{F0: fi13}
                                            v3 -= (^(v3))
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                    }
                                }
                                if (("a" + (("foo bar" + (("\n\t" + "a") + "")) + (("hello" + "hello") + "hello"))) >= ("\n\t" + "hello")) {
                                    v19 := GenS3{F0: 4294967295, F1: -251, F2: 213}
                                    fmt.Println(v6, v2)
                                    if (float64(-2.25) != 3.14159) {
                                        var v20 uint8 = v2
                                        v4["\n\t"] = (v6 || ((v20 != 181) && (-382 > int64(v20))))
                                        _ = v20
                                    }
                                    _ = v19
                                } else {
                                    fmt.Println(v8)
                                    fmt.Println(v8)
                                    if true {
                                        fmt.Println(fi14, v1)
                                    } else {
                                        if (-47 >= v8) {
                                            switch "" {
                                            case "foo bar":
                                                if v6 {
                                                    switch fi14 {
                                                    case 0:
                                                        v2 += (v2 & (v2 >> ((1) & 63)))
                                                    case 1:
                                                        fmt.Println(fi13, v8)
                                                    case 158:
                                                        v2 += 0
                                                    default:
                                                        var v21 uint8 = 195
                                                        _ = v21
                                                    }
                                                } else {
                                                    var v22 uint8 = (^(v2))
                                                    v4["hello"] = (4294967295 == 2)
                                                    _ = v22
                                                }
                                            default:
                                                fmt.Println(v6)
                                            }
                                        }
                                    }
                                }
                                var v23 int8 = (v8 + (v8 * (v8 - v8)))
                                _ = fi14
                                _ = v23
                            }
                            v24 := v4["hello"]
                            _ = fi13
                            _ = v24
                        }
                        var v25 int = 4294967295
                        _ = v12
                        _ = v25
                    } else {
                        if v6 {
                            if v6 {
                                switch -9223372036854775808 {
                                case 9223372036854775807:
                                    v3 = (v3 << ((18446744073709551615) & 63))
                                case -2147483648:
                                    _ = len(v4)
                                }
                                v1 = (v1 % ((v1) | 1))
                                if false {
                                    if ((v3 / ((v3) | 1)) > 18446744073709551615) {
                                        fmt.Println(v2)
                                    }
                                }
                            }
                            v26 := GenS3{F0: 4294967295, F1: -9223372036854775808, F2: v1}
                            v1 = v1
                            _ = v26
                        }
                        v9 = append(v9, 2147483647)
                    }
                    v27 := GenS4{F0: 127}
                    var v28 int = v5
                    _ = v27
                    _ = v28
                } else {
                    v6 = v6
                    var v29 int8 = (v8 >> (((v3 % (((v3 & v3)) | 1))) & 63))
                    _ = v29
                }
            } else {
                var v30 int32 = 2147483647
                _ = v30
            }
        }
        var v31 float64 = -2.25
        _ = v31
    }
    if (int64(v2) != 9223372036854775807) {
        v6 = true
        for fi32 := 0; fi32 < 18; fi32++ {
            var v33 int16 = 1000
            var v34 uint16 = v1
            var v35 string = "foo bar"
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
    } else {
        delete(v4, ("hello" + ("\n\t" + ("foo bar" + ("foo bar" + "a")))))
        if v6 {
            _ = len(v9)
            if true {
                if len(v9) > 0 {
                    v9[0] = (v5 % ((v5) | 1))
                }
                var v36 uint = 2
                v37 := GenS2{F0: (v5 - v5)}
                _ = v36
                _ = v37
            }
            for fi38 := 0; fi38 < 16; fi38++ {
                var v39 uint64 = 2
                if len(v9) >= 2 {
                    v9 = v9[1:len(v9)]
                }
                _ = fi38
                _ = v39
            }
        } else {
            if len(v9) > 0 {
                v9[0] = 32767
            }
            v9 = append(v9, (-(v5)))
        }
    }
    switch "\n\t" {
    case "hello":
        if v6 {
            v2 = (v2 % (((v2 & (v2 & uint8(v3)))) | 1))
            v1 = v1
        } else {
            switch ((("\n\t" + "hello") + "hello") + (("" + ((("a" + "foo bar") + "foo bar") + ("\n\t" + "foo bar"))) + "\n\t")) {
            case "hello":
                var v40 float32 = 1.5
                _ = v40
            case "\n\t":
                var v41 float32 = -2.25
                _ = v41
            case "":
                if ((v3 & v3) < v3) {
                    if false {
                        var v42 int16 = 0
                        v1 -= 90
                        _ = v42
                    }
                    for fi43 := 0; fi43 < 12; fi43++ {
                        v44 := GenS4{F0: v8}
                        fmt.Println(v5, v1, v8)
                        _ = fi43
                        _ = v44
                    }
                } else {
                    switch (v5 & v5) {
                    case 42:
                        var v45 uint64 = 18446744073709551615
                        _ = v45
                    }
                    v6 = v6
                }
            default:
                for fi46 := 0; fi46 < 7; fi46++ {
                    switch fi46 {
                    case -891:
                        var v47 string = ((((("a" + "\n\t") + "foo bar") + ("a" + "foo bar")) + ("hello" + "a")) + "\n\t")
                        _ = v47
                    default:
                        switch 0 {
                        case 0:
                            switch "" {
                            case "hello":
                                delete(v4, ("" + ("hello" + "")))
                            case "a":
                                if v6 {
                                    var v48 float64 = float64(0.0)
                                    v3 = uint(v2)
                                    v1 = v1
                                    _ = v48
                                }
                            default:
                                for fi49 := 0; fi49 < 10; fi49++ {
                                    if v6 {
                                        var v50 int16 = 32767
                                        v51 := v4[("\n\t" + "")]
                                        fmt.Println(v50, v6)
                                        _ = v50
                                        _ = v51
                                    } else {
                                        v52 := GenS4{F0: 0}
                                        if v6 {
                                            fmt.Println(v3, v6, v1)
                                            delete(v4, "a")
                                        } else {
                                            v53 := GenS5{F0: 1}
                                            if len(v9) > 0 {
                                                v9[0] = fi46
                                            }
                                            v8 += -1
                                            _ = v53
                                        }
                                        if v6 {
                                            v54 := GenS2{F0: (fi46 & (fi49 | 0))}
                                            var v55 float32 = 1e10
                                            _ = v54
                                            _ = v55
                                        }
                                        _ = v52
                                    }
                                    var v56 uint8 = v2
                                    _ = len(v4)
                                    v6 = v6
                                    v6 = true
                                    var v57 string = ("" + "a")
                                    if v6 {
                                        v58 := v4[v57]
                                        switch (fi49 & (fi46 << (((v3 * v3)) & 63))) {
                                        case -1:
                                            var v59 int16 = 0
                                            var v60 uint = 826
                                            _ = v59
                                            _ = v60
                                        case 4294967295:
                                            var v61 int = v5
                                            _ = v61
                                        case 374:
                                            v6 = v6
                                        default:
                                            var v62 int64 = 923
                                            _ = v62
                                        }
                                        var v63 uint = 255
                                        var v64 int32 = 927
                                        switch v57 {
                                        case "":
                                            v2 += 42
                                        case "hello":
                                            if len(v9) >= 2 {
                                                v9 = v9[1:len(v9)]
                                            }
                                        case "\n\t":
                                            v3 = v63
                                        default:
                                            if ((v58 && (v6 && (v64 >= (v64 % ((-666) | 1))))) || ((0 < 1000) && v58)) {
                                                v4[v57] = false
                                                v58 = false
                                                switch (v5 % ((v5) | 1)) {
                                                case 191:
                                                    v65 := v4["hello"]
                                                    _ = v65
                                                case 42:
                                                    if (!v58) {
                                                        v57 = (((v57 + ("hello" + v57)) + v57) + v57)
                                                        v2 = v56
                                                        var v66 uint32 = 863
                                                        _ = v66
                                                    } else {
                                                        _ = len(v9)
                                                        v64 += (v64 ^ 2147483647)
                                                        var v67 uint64 = 308
                                                        _ = v67
                                                    }
                                                case 4294967295:
                                                    if (true && v58) {
                                                        v9 = append(v9, (^(fi49)))
                                                    }
                                                default:
                                                    v63 += v63
                                                }
                                            } else {
                                                v3 -= (v3 - v3)
                                                v64 += int32(v8)
                                            }
                                        }
                                        _ = v58
                                        _ = v63
                                        _ = v64
                                    } else {
                                        var v68 float64 = 0.0
                                        _ = v68
                                    }
                                    if len(v9) > 0 {
                                        v9[0] = v5
                                    }
                                    v6 = v6
                                    _ = fi49
                                    _ = v56
                                    _ = v57
                                }
                            }
                        case 9223372036854775807:
                            fmt.Println(v1, v1, v2)
                        case 719:
                            switch -912 {
                            case 127:
                                var v69 uint64 = 18446744073709551615
                                _ = v69
                            case 9223372036854775807:
                                v3 = (^(v3))
                            case 146:
                                for fi70 := 0; fi70 < 13; fi70++ {
                                    _ = len(v9)
                                    v71 := GenS6{F0: -959}
                                    v6 = v6
                                    _ = fi70
                                    _ = v71
                                }
                            }
                        }
                    }
                    for fi72 := 0; fi72 < 19; fi72++ {
                        v73 := GenS4{F0: (v8 << ((v3) & 63))}
                        v9 = append(v9, 32767)
                        _ = fi72
                        _ = v73
                    }
                    _ = fi46
                }
            }
            if v6 {
                switch ("\n\t" + "a") {
                case "\n\t":
                    if false {
                        var v74 float64 = (3.14159 * 0.0)
                        var v75 int64 = 763
                        _ = v74
                        _ = v75
                    }
                case "":
                    for fi76 := 0; fi76 < 5; fi76++ {
                        if v6 {
                            v2 = (v2 / ((1) | 1))
                            var v77 string = "\n\t"
                            _ = v77
                        }
                        for fi78 := 0; fi78 < 14; fi78++ {
                            fmt.Println(v3, fi78, fi78)
                            _ = fi78
                        }
                        v3 += (v3 | v3)
                        _ = fi76
                    }
                default:
                    var v79 uint16 = (v1 | 65535)
                    _ = v79
                }
                v3 += uint(v1)
                if v6 {
                    for fi80 := 0; fi80 < 20; fi80++ {
                        v8 += (v8 / (((v8 & -54)) | 1))
                        _ = fi80
                    }
                    if (v3 >= 427) {
                        switch "hello" {
                        case "\n\t":
                            v6 = v6
                        default:
                            var v81 int16 = 364
                            _ = v81
                        }
                        v8 = -78
                        v2 += uint8(v1)
                    }
                } else {
                    v82 := GenS2{F0: 42}
                    fmt.Println(v2, v6, v6)
                    _ = v82
                }
            }
            v83 := GenS7{F0: v8, F1: "hello"}
            _ = v83
        }
    case "foo bar":
        var v84 bool = ((true || (!true)) && false)
        _ = v84
    case "\n\t":
        v8 -= (v8 | v8)
    }
    var v85 float32 = 3.14159
    fmt.Println(v8, v2)
    v9 = append(v9, -367)
    v5 = (v5 ^ (v5 / ((v5) | 1)))
    if (!true) {
        fmt.Println(v6, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v85:", v85)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v85
    _ = v9
}
