package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint64 }
type GenS4 struct { F0 string; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int8; F1 string; F2 uint64 }
type GenS9 struct { F0 bool; F1 string; F2 uint16 }
type GenS10 struct { F0 uint; F1 bool; F2 uint16 }
func genFn0() string {
    var v7 bool = false
    if ((2 <= 255) && (!(!false))) {
        if v7 {
            var v8 float32 = 3.14159
            v9 := GenS1{F0: 1}
            _ = v8
            _ = v9
        } else {
            v10 := GenS1{F0: 1}
            var v11 int16 = -32768
            var v12 uint64 = 42
            _ = v10
            _ = v11
            _ = v12
        }
    } else {
        fmt.Println(v7)
    }
    var v13 int = -518
    v14 := []int{v13, v13, (v13 & v13), v13}
    _ = v13
    _ = v14
    _ = v7
    return (("a" + "\n\t") + "")
}

func main() {
    var v1 int64 = -1
    var v2 int16 = -32768
    var v3 uint16 = 32767
    v4 := []string{"a", "", "a", "a"}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]bool)
    fmt.Println(v1)
    switch -9223372036854775808 {
    case -2147483648:
        switch "" {
        case "hello":
            if true {
                if false {
                    v15 := GenS1{F0: int16(v1)}
                    switch 656 {
                    case 32767:
                        for fi16 := 0; fi16 < 15; fi16++ {
                            if ((false || false) && (!((134 <= v2) || ("" != ("\n\t" + (("foo bar" + "hello") + "\n\t")))))) {
                                for fi17 := 0; fi17 < 3; fi17++ {
                                    v2 -= (v2 & 1)
                                    _ = fi17
                                }
                                var v18 int16 = (v2 % (((v2 | v2)) | 1))
                                v19 := GenS2{F0: (v3 * (v3 / ((v3) | 1)))}
                                _ = v18
                                _ = v19
                            }
                            _ = fi16
                        }
                    case 42:
                        _ = len(v4)
                    case 0:
                        v1 -= (v1 + 405)
                    }
                    _ = v15
                } else {
                    v20 := GenS2{F0: v3}
                    v3 += v3
                    _ = v20
                }
                switch "" {
                case "a":
                    v1 -= 858
                case "hello":
                    for fi21 := 0; fi21 < 12; fi21++ {
                        fmt.Println(v3, fi21)
                        var v22 string = ""
                        _ = fi21
                        _ = v22
                    }
                }
            } else {
                var v23 uint8 = 1
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = v23
            }
        case "\n\t":
            v24 := GenS3{F0: 290, F1: -2147483648, F2: 1}
            _ = v24
        }
    case 127:
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            switch (("" + "hello") + ("a" + "a")) {
            case "a":
                v3 = v3
            }
            var v25 int8 = -84
            _ = v25
        }
    default:
        for fi26 := 0; fi26 < 14; fi26++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            for fi27 := 0; fi27 < 12; fi27++ {
                var v28 int16 = (v2 - (v2 + 274))
                v2 = (-(v28))
                _ = fi27
                _ = v28
            }
            _ = fi26
        }
    }
    for fi29 := 0; fi29 < 18; fi29++ {
        _ = len(v6)
        for fi30 := 0; fi30 < 6; fi30++ {
            v2 -= v2
            fmt.Println(v1, fi30, fi29)
            if false {
                v2 = 1
                v5[(fi30 + fi29)] = 2
            } else {
                v3 = v3
                delete(v5, (fi30 + (fi30 >> ((2) & 63))))
            }
            var v31 uint16 = 1
            var v32 bool = ((("" + (("a" + "hello") + "a")) + "hello") != "hello")
            _ = fi30
            _ = v31
            _ = v32
        }
        _ = fi29
    }
    if true {
        v33 := GenS1{F0: (v2 / (((^(v2))) | 1))}
        v34 := GenS3{F0: 1, F1: -2147483648, F2: 0}
        _ = v33
        _ = v34
    } else {
        var v35 uint = 751
        _ = v35
    }
    for fi36 := 0; fi36 < 14; fi36++ {
        v3 = v3
        var v37 float32 = 3.14159
        if true {
            var v38 int64 = v1
            _ = v38
        }
        _ = fi36
        _ = v37
    }
    v3 += (v3 ^ v3)
    for fi39 := 0; fi39 < 19; fi39++ {
        for fi40 := 0; fi40 < 16; fi40++ {
            if ((!(4294967295 > 1)) || (0.1 == 1e10)) {
                _ = len(v4)
                if (!(915 < 32767)) {
                    v41 := GenS4{F0: ("\n\t" + ""), F1: 1, F2: 1}
                    v1 = (v1 * v1)
                    if (!((!("a" == "a")) || (false || (true || false)))) {
                        if (1e10 >= (3.14159 / 0.1)) {
                            var v42 int64 = v1
                            _ = v42
                        } else {
                            switch "foo bar" {
                            case "\n\t":
                                switch 9223372036854775807 {
                                case -744:
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                case 0:
                                    switch ((("hello" + "a") + "") + ("" + "")) {
                                    case "a":
                                        var v43 int16 = v2
                                        _ = v43
                                    case "":
                                        if (-1 >= 1) {
                                            v2 += v2
                                        }
                                    }
                                default:
                                    var v44 bool = ((1 >= 42) && true)
                                    _ = v44
                                }
                            case "a":
                                v2 += (^(v2))
                            default:
                                _ = len(v6)
                            }
                            v45 := GenS3{F0: 4294967295, F1: -2147483648, F2: 18446744073709551615}
                            switch fi39 {
                            case 32767:
                                v46 := GenS5{F0: v3}
                                _ = v46
                            case 981:
                                if ((1e10 - 0.0) != (-2.25 + 1.5)) {
                                    v47 := v5[(-(fi40))]
                                    _ = v47
                                }
                            case 743:
                                if len(v4) > 0 {
                                    v4[0] = ("hello" + "a")
                                }
                            default:
                                if (-1 != 42) {
                                    switch "foo bar" {
                                    case "hello":
                                        v2 = v2
                                    case "\n\t":
                                        var v48 int = fi40
                                        var v49 float64 = 1e10
                                        _ = v48
                                        _ = v49
                                    case "":
                                        fmt.Println(fi39)
                                    default:
                                        if (true && (!false)) {
                                            _ = len(v6)
                                            v2 = (v2 & 1000)
                                        } else {
                                            var v50 string = ("hello" + "\n\t")
                                            _ = v50
                                        }
                                    }
                                    _ = len(v5)
                                    fmt.Println(v1)
                                } else {
                                    if true {
                                        fmt.Println(v1, v2, v2)
                                        v5[fi40] = 2
                                        var v51 int8 = 100
                                        _ = v51
                                    }
                                    var v52 uint = 255
                                    v53 := GenS2{F0: uint16(v52)}
                                    _ = v52
                                    _ = v53
                                }
                            }
                            _ = v45
                        }
                    }
                    _ = v41
                } else {
                    var v54 string = "a"
                    switch "\n\t" {
                    case "":
                        v4 = append(v4, ("foo bar" + v54))
                    }
                    v2 = 42
                    v3 = 255
                    _ = v54
                }
                var v55 int64 = v1
                _ = v55
            } else {
                var v56 int16 = 1000
                v57 := GenS3{F0: 2147483647, F1: -2147483648, F2: 42}
                _ = v56
                _ = v57
            }
            v58 := GenS6{F0: 1, F1: 65535}
            v3 = 0
            _ = fi40
            _ = v58
        }
        delete(v6, 4294967295)
        for fi59 := 0; fi59 < 8; fi59++ {
            switch 32768 {
            case 32768:
                if (-128 == 67) {
                    if len(v4) > 0 {
                        v4[0] = "foo bar"
                    }
                }
            case 128:
                var v60 uint32 = 977
                _ = v60
            }
            var v61 bool = (!true)
            fmt.Println(fi59, v3, v61)
            _ = fi59
            _ = v61
        }
        _ = fi39
    }
    for fi62 := 0; fi62 < 1; fi62++ {
        switch "\n\t" {
        case "hello":
            fmt.Println(v3)
        case "\n\t":
            if true {
                var v63 uint = 230
                for fi64 := 0; fi64 < 14; fi64++ {
                    if true {
                        var v65 uint32 = 1
                        var v66 uint64 = 0
                        if (191 > v63) {
                            var v67 uint64 = (v66 << (((v63 - v63)) & 63))
                            v65 -= (v65 & v65)
                            _ = v67
                        }
                        _ = v65
                        _ = v66
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        v68 := v6[1]
                        v63 += (v63 >> ((v63) & 63))
                        _ = v68
                    }
                    _ = fi64
                }
                switch ("a" + "") {
                case "foo bar":
                    if true {
                        var v69 uint8 = 127
                        v63 += (v63 - v63)
                        fmt.Println(v1, v2)
                        _ = v69
                    }
                case "":
                    switch fi62 {
                    case -607:
                        v70 := GenS7{F0: true}
                        _ = v70
                    case 553:
                        v71 := GenS5{F0: 0}
                        _ = v71
                    case 2147483647:
                        var v72 uint16 = v3
                        _ = v72
                    }
                default:
                    var v73 int32 = 2147483647
                    _ = v73
                }
                _ = v63
            } else {
                for fi74 := 0; fi74 < 14; fi74++ {
                    fmt.Println(v3, v3)
                    _ = fi74
                }
                v75 := GenS8{F0: 83, F1: ((("a" + "hello") + "") + "a"), F2: 0}
                _ = v75
            }
        }
        var v76 string = "foo bar"
        if (-1 < -128) {
            v1 = v1
            if false {
                v2 += v2
                v77 := GenS9{F0: true, F1: (v76 + v76), F2: (v3 - 200)}
                _ = v77
            } else {
                v3 = (v3 + 883)
                v78 := v5[(fi62 | (fi62 << ((688) & 63)))]
                fmt.Println(v76)
                _ = v78
            }
            var v79 int16 = (v2 & v2)
            _ = v79
        }
        _ = fi62
        _ = v76
    }
    if false {
        var v80 string = "hello"
        _ = v80
    }
    var v81 float32 = float32(((((((((((3.14159 + 3.14159) + float64(3.14159)) / 1.5) + ((((1.5 - -2.25) - 0.0) - (3.14159 * 0.0)) * 0.0)) + 0.1) + ((float64((1.5 / 3.14159)) / 1e10) - (0.1 * 0.0))) - 1.5) * ((0.0 / 0.1) / 0.1)) - (3.14159 * ((3.14159 / (1.5 / 1.5)) / ((1e10 * 0.1) + 1.5)))) - 3.14159))
    fmt.Println(v2)
    v82 := GenS10{F0: 816, F1: (884 != 32768), F2: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v81
    _ = v82
}
