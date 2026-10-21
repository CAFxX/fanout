package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint64; F1 int64 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 uint32; F1 uint32; F2 int8 }
type GenS5 struct { F0 int16; F1 uint16; F2 string }
type GenS6 struct { F0 int16; F1 int16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 uint }
type GenS9 struct { F0 int; F1 uint8 }

func main() {
    v1 := make(map[string]string)
    var v2 uint8 = 29
    v3 := []uint32{2}
    var v4 int16 = -1
    var v5 uint8 = v2
    v6 := []string{((("hello" + ("a" + "")) + "\n\t") + "foo bar"), ("" + "hello"), (("a" + "a") + ""), ""}
    fmt.Println(v5)
    for fi7 := 0; fi7 < 3; fi7++ {
        if (!false) {
            v8 := v1["hello"]
            v9 := GenS1{F0: -9223372036854775808, F1: v8}
            var v10 int64 = 9223372036854775807
            _ = v10
            _ = v8
            _ = v9
        } else {
            v4 -= v4
            v1["hello"] = (("foo bar" + (("hello" + ("a" + "a")) + (("hello" + ((("hello" + "foo bar") + "\n\t") + ("hello" + ""))) + ("foo bar" + (("\n\t" + "\n\t") + ("hello" + "a")))))) + ("hello" + "hello"))
        }
        _ = fi7
    }
    v11 := GenS1{F0: 42, F1: ("a" + "\n\t")}
    var v12 float64 = 1.5
    v13 := v1[("foo bar" + "\n\t")]
    var v14 uint16 = 65535
    if (!false) {
        v15 := GenS1{F0: -9223372036854775808, F1: ("" + ("a" + (v13 + v13)))}
        switch 1 {
        case -2147483648:
            switch "foo bar" {
            case "foo bar":
                fmt.Println(v14, v2, v14)
            }
        default:
            if true {
                if true {
                    if false {
                        v16 := GenS2{F0: 42, F1: 42}
                        fmt.Println(v2, v5)
                        _ = v16
                    }
                    v12 = 3.14159
                    v17 := GenS3{F0: v4, F1: (v14 | v14)}
                    _ = v17
                }
            }
        }
        v14 = (v14 ^ (v14 >> ((357) & 63)))
        _ = v15
    } else {
        var v18 uint = 65535
        fmt.Println(v13, v13, v12)
        _ = v18
    }
    var v19 int8 = 0
    var v20 uint = 4294967295
    if (false && false) {
        var v21 string = v13
        v19 -= v19
        _ = v21
    } else {
        var v22 uint8 = (^(v2))
        var v23 int16 = (v4 / (((v4 % ((v4) | 1))) | 1))
        _ = v22
        _ = v23
    }
    var v24 float32 = 3.14159
    switch -1 {
    case 32768:
        switch ((v13 + "\n\t") + "") {
        case "foo bar":
            var v25 bool = true
            _ = v25
        case "\n\t":
            for fi26 := 0; fi26 < 18; fi26++ {
                v20 -= 907
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 uint16 = (v14 << ((752) & 63))
                    switch fi27 {
                    case 42:
                        v29 := GenS4{F0: 590, F1: 1, F2: 26}
                        _ = v29
                    case -815:
                        v30 := GenS5{F0: (v4 * v4), F1: v14, F2: v13}
                        var v31 int64 = -1
                        v31 += v31
                        switch int(v19) {
                        case 341:
                            switch 2147483647 {
                            case 4294967295:
                                v32 := GenS3{F0: (v4 * v4), F1: (v28 & 255)}
                                _ = v32
                            default:
                                v5 = 91
                            }
                        }
                        v33 := GenS6{F0: 1000, F1: (-(v4))}
                        _ = v30
                        _ = v31
                        _ = v33
                    case 128:
                        if false {
                            switch "" {
                            case "foo bar":
                                switch int(v4) {
                                case -1:
                                    v13 = v13
                                case -9223372036854775808:
                                    if false {
                                        if true {
                                            v19 = 127
                                            if (0 <= -724) {
                                                fmt.Println(v20, v28)
                                            }
                                            if (!false) {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                fmt.Println(v28, v19, v4)
                                            }
                                        } else {
                                            var v34 float64 = (1.5 / ((v12 - v12) / (-2.25 - v12)))
                                            var v35 int16 = (v4 | v4)
                                            _ = v34
                                            _ = v35
                                        }
                                    }
                                case 128:
                                    if true {
                                        if (781 >= 252) {
                                            v19 = v19
                                            var v36 bool = (!((v24 > (v24 + 1e10)) && false))
                                            var v37 bool = v36
                                            _ = v36
                                            _ = v37
                                        } else {
                                            var v38 int64 = 457
                                            var v39 uint = (v20 / ((v20) | 1))
                                            _ = v38
                                            _ = v39
                                        }
                                        v1[((("foo bar" + v13) + ((v13 + "\n\t") + v13)) + (v13 + (v13 + v13)))] = v13
                                        if false {
                                            v40 := GenS6{F0: 996, F1: v4}
                                            _ = v40
                                        }
                                        if (2 < 18446744073709551615) {
                                            if (v19 != (-(v19))) {
                                                fmt.Println(fi26)
                                                var v41 uint8 = 1
                                                var v42 bool = (true && false)
                                                if len(v6) > 0 {
                                                    v6[0] = v13
                                                }
                                                var v43 bool = ((v24 * v24) > -2.25)
                                                _ = v41
                                                _ = v42
                                                _ = v43
                                            } else {
                                                if len(v6) > 0 {
                                                    v6[0] = (v13 + (v13 + ("hello" + v13)))
                                                }
                                                delete(v1, v13)
                                            }
                                            if true {
                                                if (true || true) {
                                                    if (0 < 128) {
                                                        switch (v13 + (v13 + "\n\t")) {
                                                        case "foo bar":
                                                            delete(v1, v13)
                                                        case "a":
                                                            _ = len(v3)
                                                        case "\n\t":
                                                            v24 -= float32((v12 + v12))
                                                        }
                                                        _ = len(v3)
                                                    }
                                                    fmt.Println(fi27, v4, fi27)
                                                    var v44 string = v13
                                                    _ = v44
                                                } else {
                                                    var v45 uint16 = v14
                                                    _ = v45
                                                }
                                            }
                                            v12 = v12
                                            if (!false) {
                                                fmt.Println(v2)
                                            }
                                            v46 := GenS6{F0: (v4 + v4), F1: (v4 & v4)}
                                            _ = v46
                                        } else {
                                            if true {
                                                v2 = 128
                                                v47 := v1[v13]
                                                switch (fi27 | (fi26 / (((fi26 / ((fi26) | 1))) | 1))) {
                                                case 823:
                                                    v28 -= 912
                                                default:
                                                    delete(v1, v47)
                                                }
                                                var v48 uint16 = v14
                                                _ = v47
                                                _ = v48
                                            } else {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v49 := GenS7{F0: v4}
                                                _ = v49
                                            }
                                        }
                                        fmt.Println(fi26, v2, v20)
                                    }
                                default:
                                    v28 -= 429
                                }
                            case "":
                                var v50 uint16 = (^(v14))
                                _ = v50
                            }
                            fmt.Println(v2)
                            v51 := GenS8{F0: (v28 >> ((65535) & 63)), F1: (v20 >> ((v20) & 63))}
                            v6 = append(v6, v13)
                            if (-9223372036854775808 == -351) {
                                var v52 uint8 = (^(v2))
                                v53 := GenS9{F0: (fi27 | fi27), F1: (v5 / ((v2) | 1))}
                                v5 += v5
                                v54 := GenS7{F0: 42}
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                            _ = v51
                        }
                    }
                    _ = fi27
                    _ = v28
                }
                switch (v13 + (v13 + "hello")) {
                case "foo bar":
                    var v55 uint8 = (v5 & (v5 / (((v5 << ((uint(v2)) & 63))) | 1)))
                    _ = v55
                case "":
                    v19 = 1
                case "hello":
                    switch "" {
                    case "":
                        for fi56 := 0; fi56 < 20; fi56++ {
                            _ = len(v6)
                            switch v13 {
                            case "":
                                if ((v14 * 65535) >= v14) {
                                    v5 = 122
                                } else {
                                    var v57 uint = (v20 / ((0) | 1))
                                    _ = v57
                                }
                                if (!true) {
                                    if (false || (v19 > -128)) {
                                        v2 = (v2 & 0)
                                    }
                                    if false {
                                        v2 = v5
                                        var v58 uint16 = v14
                                        fmt.Println(v58, v24, v2)
                                        _ = v58
                                    } else {
                                        fmt.Println(v19, fi26, v14)
                                        fmt.Println(v12)
                                    }
                                    v59 := GenS6{F0: v4, F1: -732}
                                    _ = v59
                                } else {
                                    fmt.Println(v12, v14)
                                }
                                var v60 int16 = (-(v4))
                                _ = v60
                            case "hello":
                                fmt.Println(v4, v12, v2)
                            case "\n\t":
                                var v61 int = (fi56 | fi56)
                                _ = v61
                            }
                            _ = fi56
                        }
                    default:
                        for fi62 := 0; fi62 < 12; fi62++ {
                            v3 = append(v3, 1)
                            if false {
                                var v63 bool = false
                                v13 = "\n\t"
                                var v64 int8 = -1
                                _ = v63
                                _ = v64
                            } else {
                                switch "foo bar" {
                                case "\n\t":
                                    if len(v3) > 0 {
                                        v3[0] = 2147483647
                                    }
                                case "hello":
                                    _ = len(v1)
                                case "foo bar":
                                    v24 = v24
                                default:
                                    v65 := v1[v13]
                                    _ = v65
                                }
                            }
                            _ = fi62
                        }
                    }
                }
                _ = fi26
            }
        case "a":
            var v66 uint8 = (v5 + (^(v5)))
            _ = v66
        default:
            fmt.Println(v12)
        }
    case 42:
        if len(v6) > 0 {
            v6[0] = "hello"
        }
    }
    var v67 int16 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v67
}
