package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 bool; F2 bool }
type GenS5 struct { F0 int8; F1 uint32; F2 int16 }
type GenS6 struct { F0 int32; F1 string }
type GenS7 struct { F0 int16; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -124
    v3 := make(map[string]int)
    if false {
        if ((("foo bar" + ("\n\t" + "foo bar")) + ("a" + ("foo bar" + ("\n\t" + ("" + "hello"))))) <= ("\n\t" + "a")) {
            if (v2 >= 42) {
                if (65535 == 58) {
                    var v4 int64 = 817
                    fmt.Println(v2)
                    _ = v4
                }
                v5 := v3["foo bar"]
                _ = v5
            } else {
                v2 = 1
                if (991 >= 987) {
                    fmt.Println(v2)
                }
            }
        }
        switch 461 {
        case 0:
            switch -82 {
            case -845:
                var v6 uint64 = 436
                _ = v6
            case 128:
                delete(v3, "hello")
            default:
                for fi7 := 0; fi7 < 9; fi7++ {
                    for fi8 := 0; fi8 < 18; fi8++ {
                        _ = len(v3)
                        v2 -= (v2 | v2)
                        _ = fi8
                    }
                    _ = fi7
                }
            }
        default:
            var v9 int16 = 0
            _ = v9
        }
    } else {
        for fi10 := 0; fi10 < 8; fi10++ {
            v2 = v2
            _ = fi10
        }
        v2 = v2
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: 1, F1: "hello"}
        var v13 int32 = 0
        var v14 int64 = 42
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v2 = v2
    v2 += -1
    if false {
        v15 := []int{-1, 42}
        v16 := v3[("hello" + ("foo bar" + ("foo bar" + "a")))]
        switch v16 {
        case -754:
            v17 := GenS1{F0: 17, F1: "hello"}
            _ = v17
        case 32768:
            if true {
                v16 = v16
                _ = len(v15)
            }
        default:
            v16 += v16
        }
        _ = v15
        _ = v16
    } else {
        v2 = 125
    }
    v2 = 13
    v18 := v1["foo bar"]
    switch (v18 % (((-(v18))) | 1)) {
    case 738:
        v19 := GenS1{F0: 9223372036854775807, F1: "hello"}
        _ = v19
    }
    for fi20 := 0; fi20 < 4; fi20++ {
        switch 4294967295 {
        case 9223372036854775807:
            v2 = (v2 % (((v2 / (((-(v2))) | 1))) | 1))
        case 380:
            v21 := []int{v18}
            _ = v21
        }
        v2 = (^(v2))
        _ = fi20
    }
    var v22 int32 = 42
    v22 = v22
    for fi23 := 0; fi23 < 14; fi23++ {
        var v24 uint8 = 128
        if true {
            v2 = v2
            v25 := GenS1{F0: -504, F1: ""}
            if ((v24 >> ((2) & 63)) >= (^(v24))) {
                var v26 float32 = -2.25
                _ = v26
            }
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    if true {
        var v27 bool = true
        _ = v27
    } else {
        if (0 >= 0) {
            if false {
                v28 := GenS2{F0: -1, F1: 2, F2: -626}
                _ = v28
            } else {
                var v29 float32 = (1e10 / -2.25)
                fmt.Println(v2, v18)
                _ = v29
            }
            v30 := GenS3{F0: v18}
            _ = v30
        }
        v22 = v22
        if false {
            v31 := GenS3{F0: v18}
            _ = v31
        }
    }
    v32 := GenS4{F0: (v18 % ((127) | 1)), F1: true, F2: (("\n\t" + "a") == ("\n\t" + "a"))}
    for fi33 := 0; fi33 < 3; fi33++ {
        for fi34 := 0; fi34 < 6; fi34++ {
            delete(v1, "\n\t")
            fmt.Println(fi33)
            v22 -= (-(v22))
            if true {
                var v35 int64 = 1
                _ = v35
            }
            _ = fi34
        }
        for fi36 := 0; fi36 < 16; fi36++ {
            switch "" {
            case "hello":
                v2 -= v2
                v37 := GenS5{F0: v2, F1: 4294967295, F2: 996}
                _ = v37
            case "\n\t":
                var v38 float32 = 1e10
                _ = v38
            default:
                if (false || (2 != 2147483648)) {
                    var v39 uint8 = 200
                    _ = v39
                } else {
                    v22 = 0
                    if false {
                        v40 := GenS6{F0: v22, F1: "\n\t"}
                        _ = v40
                    }
                    if (("a" + "a") <= "") {
                        v41 := []int{}
                        var v42 int64 = 9223372036854775807
                        if (!false) {
                            var v43 uint8 = 250
                            var v44 uint8 = v43
                            if true {
                                v22 = v22
                            }
                            fmt.Println(v2, v2, v42)
                            if true {
                                if false {
                                    v2 = (v2 * v2)
                                    var v45 string = "a"
                                    fmt.Println(v45, v2, fi33)
                                    _ = v45
                                }
                                fmt.Println(v42)
                                fmt.Println(v42, v22, v43)
                            }
                            _ = v43
                            _ = v44
                        } else {
                            if len(v41) >= 2 {
                                v41 = v41[1:len(v41)]
                            }
                        }
                        _ = v41
                        _ = v42
                    }
                    var v46 int8 = (v2 % ((v2) | 1))
                    _ = v46
                }
            }
            switch "a" {
            case "a":
                switch (fi33 % (((v18 * (fi33 << ((908) & 63)))) | 1)) {
                case 356:
                    if (!(297 > 1)) {
                        var v47 float32 = float32(float64(0.1))
                        switch (fi33 ^ (v18 | v18)) {
                        case 32768:
                            var v48 uint16 = 65535
                            _ = v48
                        case -123:
                            v49 := v3["a"]
                            _ = v49
                        }
                        var v50 uint32 = 797
                        switch v18 {
                        case -692:
                            if false {
                                if true {
                                    v51 := GenS7{F0: 0, F1: 4294967295, F2: 603}
                                    v3["a"] = -9223372036854775808
                                    var v52 uint8 = 1
                                    v53 := GenS1{F0: 9223372036854775807, F1: "a"}
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                }
                                v22 = 0
                                v54 := GenS3{F0: (v18 * (v18 << ((128) & 63)))}
                                _ = v54
                            } else {
                                if true {
                                    v2 = (v2 >> ((0) & 63))
                                    fmt.Println(v50, v22, fi36)
                                } else {
                                    v18 -= v18
                                    v2 = (v2 >> ((694) & 63))
                                }
                            }
                        case -612:
                            v22 -= 42
                        case 0:
                            v55 := GenS5{F0: v2, F1: (v50 ^ (v50 | 65535)), F2: 1}
                            _ = v55
                        default:
                            v56 := GenS7{F0: 42, F1: 255, F2: 32767}
                            _ = v56
                        }
                        switch (^(fi36)) {
                        case -9223372036854775808:
                            fmt.Println(v2)
                        default:
                            switch (v18 ^ fi36) {
                            case 32768:
                                var v57 int64 = 0
                                _ = v57
                            case 128:
                                fmt.Println(fi36, v2, v47)
                            }
                        }
                        _ = v47
                        _ = v50
                    }
                case 127:
                    v58 := []int{9223372036854775807}
                    _ = v58
                case 636:
                    switch 32767 {
                    case 907:
                        _ = len(v1)
                    case 400:
                        if true {
                            switch -1 {
                            case 876:
                                v22 -= (v22 * v22)
                            case -9223372036854775808:
                                var v59 uint16 = 65535
                                _ = v59
                            }
                        } else {
                            v60 := GenS1{F0: 0, F1: "\n\t"}
                            switch fi33 {
                            case 0:
                                delete(v3, "\n\t")
                            case 593:
                                fmt.Println(v18, v22)
                            default:
                                v61 := GenS1{F0: 42, F1: ("foo bar" + "")}
                                _ = v61
                            }
                            _ = v60
                        }
                    }
                default:
                    v62 := []int{fi36, 571}
                    _ = v62
                }
                v22 += (v22 + (v22 / ((v22) | 1)))
            }
            v22 -= v22
            _ = fi36
        }
        if (46 == 65535) {
            switch "" {
            case "foo bar":
                for fi63 := 0; fi63 < 9; fi63++ {
                    fmt.Println(v18)
                    _ = fi63
                }
            }
        } else {
            var v64 bool = ((953 > 485) || true)
            for fi65 := 0; fi65 < 6; fi65++ {
                var v66 string = ""
                _ = fi65
                _ = v66
            }
            _ = v64
        }
        _ = fi33
    }
    var v67 int64 = -1
    var v68 int = v18
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v18
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v67
    _ = v68
}
