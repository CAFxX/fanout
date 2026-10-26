package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 int8; F2 int }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 string; F1 uint64; F2 uint16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 uint8 }

func main() {
    v1 := []uint32{2147483647, 65535}
    v2 := make(map[string]int)
    var v3 int16 = 544
    var v4 bool = true
    var v5 float32 = (0.1 + 3.14159)
    var v6 int8 = 100
    if false {
        v5 = float32(((1e10 + ((0.0 + float64(v5)) / (1e10 - 1.5))) * -2.25))
        switch -9223372036854775808 {
        case -23:
            var v7 uint16 = 0
            _ = v7
        case 2147483647:
            if (9223372036854775807 <= 491) {
                for fi8 := 0; fi8 < 13; fi8++ {
                    switch "foo bar" {
                    case "":
                        v6 = v6
                    case "a":
                        if false {
                            delete(v2, "")
                        }
                    case "hello":
                        v4 = false
                    default:
                        fmt.Println(v4)
                    }
                    if (128 >= 255) {
                        fmt.Println(v5, v4)
                        switch fi8 {
                        case 6:
                            var v9 bool = v4
                            _ = v9
                        }
                    } else {
                        switch ("foo bar" + "") {
                        case "":
                            v10 := GenS1{F0: "a", F1: v6}
                            _ = v10
                        case "hello":
                            v1 = append(v1, 4294967295)
                        default:
                            v3 += v3
                        }
                    }
                    if (!((v6 & v6) >= (v6 >> ((752) & 63)))) {
                        var v11 int = fi8
                        var v12 string = "foo bar"
                        _ = v11
                        _ = v12
                    }
                    _ = fi8
                }
                var v13 int32 = -565
                for fi14 := 0; fi14 < 2; fi14++ {
                    v15 := GenS1{F0: "a", F1: (v6 - v6)}
                    switch "foo bar" {
                    case "foo bar":
                        if (0.1 != (1e10 + (1e10 * -2.25))) {
                            switch (fi14 & 812) {
                            case 32767:
                                var v16 int8 = (v6 ^ (v6 & 1))
                                _ = v16
                            }
                            if len(v1) > 0 {
                                v1[0] = 332
                            }
                        } else {
                            v17 := GenS1{F0: "\n\t", F1: (v6 * v6)}
                            switch (fi14 | 127) {
                            case 2147483647:
                                for fi18 := 0; fi18 < 13; fi18++ {
                                    v3 += v3
                                    if (2 == 152) {
                                        v13 -= (v13 ^ v13)
                                    }
                                    _ = fi18
                                }
                            case 32767:
                                var v19 uint32 = 65535
                                _ = v19
                            }
                            v20 := GenS1{F0: ("" + (("a" + "foo bar") + "\n\t")), F1: (v6 % ((v6) | 1))}
                            _ = v17
                            _ = v20
                        }
                    case "a":
                        var v21 uint64 = 1
                        _ = v21
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v13
            } else {
                if v4 {
                    var v22 int = -9223372036854775808
                    _ = v22
                } else {
                    v23 := v2[(("hello" + "") + ("a" + "foo bar"))]
                    switch "foo bar" {
                    case "foo bar":
                        var v24 int16 = v3
                        _ = v24
                    default:
                        if (v4 || false) {
                            fmt.Println(v23, v5)
                            if v4 {
                                v6 = 42
                                fmt.Println(v5, v23)
                            }
                            for fi25 := 0; fi25 < 15; fi25++ {
                                v26 := GenS2{F0: 2147483647, F1: 127, F2: ("hello" + ("hello" + "hello"))}
                                delete(v2, "")
                                var v27 uint8 = 134
                                _ = fi25
                                _ = v26
                                _ = v27
                            }
                        }
                    }
                    _ = v23
                }
                var v28 uint = 2
                _ = v28
            }
        case 4294967295:
            if true {
                var v29 float32 = (1.5 - v5)
                _ = v29
            }
        default:
            v30 := GenS2{F0: 0, F1: 2, F2: ("foo bar" + "hello")}
            _ = v30
        }
    }
    v31 := GenS3{F0: 65535}
    var v32 int = 9223372036854775807
    switch ("hello" + "\n\t") {
    case "a":
        v33 := GenS4{F0: "a", F1: v6, F2: v32}
        _ = v33
    default:
        v2["a"] = (v32 & v32)
    }
    v6 = (v6 + v6)
    v34 := GenS3{F0: 65535}
    var v35 int64 = 1
    if v4 {
        v6 += (-(v6))
    } else {
        v35 = v35
        for fi36 := 0; fi36 < 1; fi36++ {
            for fi37 := 0; fi37 < 18; fi37++ {
                if len(v1) > 0 {
                    v1[0] = 4294967295
                }
                _ = fi37
            }
            v5 -= (v5 / -2.25)
            if true {
                v38 := GenS2{F0: 835, F1: 245, F2: "foo bar"}
                var v39 uint8 = 80
                fmt.Println(fi36, v4, v3)
                _ = v38
                _ = v39
            }
            _ = fi36
        }
        v40 := GenS3{F0: 1}
        _ = v40
    }
    v35 += v35
    var v41 string = "a"
    v42 := GenS5{F0: (v3 & v3), F1: 2}
    switch (v32 ^ 9223372036854775807) {
    case 554:
        var v43 string = (v41 + (v41 + v41))
        _ = v43
    case 42:
        v6 += (v6 % (((v6 - (v6 & (v6 ^ 42)))) | 1))
    default:
        if v4 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            var v44 uint16 = 32767
            for fi45 := 0; fi45 < 14; fi45++ {
                v46 := v2[v41]
                var v47 uint = 0
                _ = fi45
                _ = v46
                _ = v47
            }
            v48 := GenS5{F0: 32767, F1: uint32(v44)}
            _ = v44
            _ = v48
        }
    }
    switch 32768 {
    case -462:
        var v49 uint8 = 59
        _ = v49
    case -1:
        v5 -= (((v5 * -2.25) + (3.14159 * 1e10)) * (v5 / 1e10))
    }
    var v50 uint64 = 578
    v5 = (1.5 - v5)
    for fi51 := 0; fi51 < 14; fi51++ {
        if true {
            v52 := GenS3{F0: 2147483647}
            switch "foo bar" {
            case "foo bar":
                if len(v1) > 0 {
                    v1[0] = 2
                }
            case "":
                _ = len(v1)
            case "a":
                v53 := GenS6{F0: (4294967295 != 255), F1: 255}
                _ = v53
            }
            fmt.Println(v50, v4)
            _ = v52
        } else {
            _ = len(v2)
        }
        for fi54 := 0; fi54 < 18; fi54++ {
            switch v41 {
            case "a":
                v55 := GenS7{F0: "", F1: v50, F2: 614}
                _ = v55
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            if v4 {
                if v4 {
                    fmt.Println(v32)
                    if ((235 == 32767) || ((1e10 / 1e10) == float64(float32((float64(((0.1 * v5) + 1e10)) * 3.14159))))) {
                        _ = len(v2)
                        if (2 <= 2) {
                            switch fi54 {
                            case -483:
                                if (false || (false || ((32768 > -625) && true))) {
                                    _ = len(v2)
                                }
                            default:
                                switch (v32 >> ((4294967295) & 63)) {
                                case 1:
                                    v2[v41] = v32
                                }
                            }
                        }
                    }
                }
            } else {
                v41 = ("\n\t" + v41)
                if (v35 < (v35 >> ((914) & 63))) {
                    v56 := GenS2{F0: 100, F1: 128, F2: "\n\t"}
                    v57 := GenS8{F0: v3}
                    _ = v56
                    _ = v57
                } else {
                    var v58 int = (fi51 + fi51)
                    v58 = fi54
                    if v4 {
                        var v59 uint = 95
                        _ = v59
                    }
                    _ = v58
                }
                if false {
                    v60 := GenS9{F0: 0}
                    _ = v60
                } else {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    var v61 int64 = (v35 >> ((1) & 63))
                    v3 -= -978
                    _ = v61
                }
            }
            _ = fi54
        }
        _ = fi51
    }
    if (((1e10 / float64(v5)) / 1.0) >= -2.25) {
        var v62 uint64 = (v50 * v50)
        _ = v62
    } else {
        var v63 uint16 = 65535
        v35 -= 0
        fmt.Println(v5, v35)
        _ = v63
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v4
    _ = v41
    _ = v42
    _ = v5
    _ = v50
    _ = v6
}
