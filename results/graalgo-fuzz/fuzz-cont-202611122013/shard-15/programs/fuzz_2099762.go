package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64; F1 int64; F2 int }
type GenS4 struct { F0 uint32; F1 uint16; F2 int16 }
type GenS5 struct { F0 uint8; F1 uint; F2 string }
type GenS6 struct { F0 int32; F1 int64; F2 int }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 int16 }
type GenS9 struct { F0 int; F1 int32 }
func genFn0(p0 int32) uint16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        switch "\n\t" {
        case "hello":
            for fi7 := 0; fi7 < 3; fi7++ {
                fmt.Println(fi7, fi6, fi6)
                var v8 int = 4294967295
                _ = fi7
                _ = v8
            }
        case "foo bar":
            var v9 int64 = int64(p0)
            _ = v9
        case "a":
            var v10 string = "a"
            _ = v10
        }
        var v11 uint8 = 232
        v12 := []int{-209}
        _ = fi6
        _ = v11
        _ = v12
    }
    if (!("" < ("\n\t" + "a"))) {
        switch -42 {
        case -2147483648:
            for fi13 := 0; fi13 < 16; fi13++ {
                p0 -= p0
                v14 := []int{}
                var v15 int = fi13
                _ = fi13
                _ = v14
                _ = v15
            }
        case 32767:
            var v16 int64 = 676
            _ = v16
        case 128:
            v17 := []int{-512, 42, 520, -619}
            _ = v17
        default:
            switch "a" {
            case "hello":
                p0 += p0
            case "\n\t":
                var v18 string = ""
                _ = v18
            default:
                v19 := []int{}
                _ = v19
            }
        }
        fmt.Println(p0)
    } else {
        p0 += (p0 & p0)
        var v20 int16 = -1
        _ = v20
    }
    fmt.Println(p0)
    var v21 int64 = 469
    _ = p0
    _ = v21
    return 0
}

func main() {
    var v1 int = 128
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    var v4 int32 = 42
    var v5 string = "foo bar"
    switch "a" {
    case "":
        v22 := []int{}
        _ = v22
    case "hello":
        v4 += -1
    case "foo bar":
        switch v5 {
        case "hello":
            v1 = v1
        case "a":
            v23 := GenS1{F0: 127}
            _ = v23
        case "foo bar":
            if true {
                var v24 string = (v5 + v5)
                for fi25 := 0; fi25 < 4; fi25++ {
                    if false {
                        for fi26 := 0; fi26 < 17; fi26++ {
                            var v27 uint = 1
                            _ = fi26
                            _ = v27
                        }
                        v2[v5] = v1
                    } else {
                        fmt.Println(v24, fi25)
                        for fi28 := 0; fi28 < 2; fi28++ {
                            switch ("foo bar" + "") {
                            case "":
                                v4 += (v4 ^ -2147483648)
                            case "hello":
                                switch v5 {
                                case "foo bar":
                                    fmt.Println(v5, fi28)
                                    var v29 string = "\n\t"
                                    _ = v29
                                default:
                                    if (42 < 487) {
                                        var v30 uint8 = 42
                                        fmt.Println(v5, fi25, v30)
                                        var v31 int32 = -573
                                        _ = len(v2)
                                        _ = v30
                                        _ = v31
                                    } else {
                                        delete(v2, v5)
                                    }
                                }
                                var v32 int32 = (v4 % ((v4) | 1))
                                v4 = v4
                                _ = v32
                            case "a":
                                v33 := GenS2{F0: v4, F1: 1}
                                _ = v33
                            }
                            _ = fi28
                        }
                    }
                    _ = fi25
                }
                if false {
                    v4 -= v4
                }
                _ = v24
            }
        default:
            v5 = "foo bar"
        }
    default:
        var v34 uint16 = 2
        _ = v34
    }
    switch "" {
    case "hello":
        var v35 bool = (-106 >= 1)
        _ = v35
    case "":
        v1 += 42
    case "foo bar":
        fmt.Println(v1)
    default:
        for fi36 := 0; fi36 < 16; fi36++ {
            v37 := GenS3{F0: 42, F1: 265, F2: (fi36 - fi36)}
            if false {
                var v38 int64 = 0
                v5 = v5
                for fi39 := 0; fi39 < 11; fi39++ {
                    if true {
                        if false {
                            fmt.Println(fi39, v5)
                            switch (fi39 / ((-1) | 1)) {
                            case 128:
                                var v40 uint64 = 18446744073709551615
                                _ = v40
                            case 9223372036854775807:
                                switch (fi36 / ((127) | 1)) {
                                case 42:
                                    v4 += v4
                                }
                            default:
                                if false {
                                    if (true && (v38 > 456)) {
                                        v38 = v38
                                    }
                                } else {
                                    v4 += (^(v4))
                                    v1 = (v1 | (^(v1)))
                                }
                            }
                        }
                        var v41 float64 = 0.1
                        v38 = (v38 << ((65535) & 63))
                        _ = v41
                    } else {
                        if (255 != 32767) {
                            if false {
                                var v42 uint16 = 2
                                v43 := GenS4{F0: 1, F1: genFn0(v4), F2: -1}
                                fmt.Println(v1, fi39)
                                _ = v42
                                _ = v43
                            }
                            v44 := GenS5{F0: 1, F1: 18446744073709551615, F2: v5}
                            _ = v44
                        }
                        v1 = (v1 ^ fi39)
                    }
                    v5 = v5
                    _ = fi39
                }
                _ = v38
            }
            _ = fi36
            _ = v37
        }
    }
    var v45 uint = 255
    if (-441 < (v4 / ((v4) | 1))) {
        _ = len(v2)
    } else {
        v46 := []int{(^(v1))}
        var v47 uint8 = 127
        _ = v46
        _ = v47
    }
    v45 -= 65535
    for fi48 := 0; fi48 < 14; fi48++ {
        if false {
            var v49 uint64 = 18446744073709551615
            _ = v49
        } else {
            var v50 int16 = 0
            var v51 int16 = -449
            v1 = v1
            _ = v50
            _ = v51
        }
        _ = fi48
    }
    if false {
        v52 := []int{32768}
        if len(v52) >= 2 {
            v52 = v52[1:len(v52)]
        }
        _ = v52
    } else {
        if (!true) {
            for fi53 := 0; fi53 < 2; fi53++ {
                if false {
                    v54 := GenS2{F0: v4, F1: -1}
                    var v55 uint64 = 2
                    _ = v54
                    _ = v55
                }
                v56 := GenS6{F0: v4, F1: 0, F2: fi53}
                switch (v1 * v1) {
                case 9223372036854775807:
                    for fi57 := 0; fi57 < 15; fi57++ {
                        v1 += fi53
                        fmt.Println(v5, v45)
                        _ = fi57
                    }
                case -1:
                    v58 := []int{v1, -2147483648, -1, (v1 + v1)}
                    _ = v58
                }
                _ = fi53
                _ = v56
            }
        } else {
            delete(v3, (v5 + "\n\t"))
            switch (v1 / (((v1 - (v1 - (v1 / ((v1) | 1))))) | 1)) {
            case 0:
                for fi59 := 0; fi59 < 20; fi59++ {
                    v60 := GenS5{F0: 167, F1: 906, F2: "hello"}
                    delete(v3, "\n\t")
                    switch (v5 + (((v5 + (v5 + v5)) + ("\n\t" + ((((v5 + "\n\t") + v5) + "\n\t") + "\n\t"))) + (v5 + "hello"))) {
                    case "a":
                        v61 := []int{(v1 >> ((v45) & 63))}
                        _ = v61
                    case "hello":
                        if (4294967295 == 42) {
                            v62 := GenS6{F0: v4, F1: -1, F2: (fi59 >> ((v45) & 63))}
                            _ = v62
                        }
                    case "":
                        v45 -= 65535
                    }
                    _ = fi59
                    _ = v60
                }
            default:
                v63 := GenS7{F0: 255}
                _ = v63
            }
            var v64 uint16 = 651
            _ = v64
        }
    }
    if true {
        switch (v1 & 2147483647) {
        case 0:
            var v65 int16 = 1
            _ = v65
        case -1:
            v5 = v5
        case -828:
            switch v1 {
            case -929:
                v66 := GenS5{F0: 127, F1: 42, F2: ((v5 + ("foo bar" + v5)) + ((v5 + "") + "\n\t"))}
                _ = v66
            case 1:
                v67 := GenS8{F0: 217, F1: -32768}
                _ = v67
            case -792:
                v1 -= -1
            default:
                switch 1 {
                case -611:
                    v4 -= v4
                case 0:
                    var v68 float32 = 1e10
                    _ = v68
                }
            }
        default:
            v69 := GenS5{F0: 128, F1: 18446744073709551615, F2: v5}
            _ = v69
        }
        var v70 int = v1
        var v71 int16 = -875
        _ = v70
        _ = v71
    }
    v72 := []int{(v1 / ((v1) | 1)), 128, (-(v1))}
    var v73 int64 = 1
    fmt.Println(v4, v45)
    if ((839 != 1) && ((float32(1.5) / (0.1 + ((1.5 + 1.5) - 0.0))) == 1.5)) {
        var v74 string = (("hello" + ("a" + "foo bar")) + ("foo bar" + ((v5 + (((("" + v5) + ("" + ("" + v5))) + (v5 + (((v5 + v5) + v5) + "hello"))) + v5)) + (v5 + (v5 + "a")))))
        v75 := GenS9{F0: 9223372036854775807, F1: int32(v1)}
        _ = v74
        _ = v75
    }
    if (false && ((0 != -128) && ((((((0.1 + 0.1) * 1.5) / 1.5) / 1e10) < 3.14159) || false))) {
        switch v5 {
        case "":
            if false {
                v76 := GenS6{F0: (v4 & (v4 >> ((v45) & 63))), F1: 0, F2: int(v45)}
                fmt.Println(v1, v45, v4)
                switch ("a" + ((v5 + ("a" + v5)) + v5)) {
                case "":
                    if true {
                        for fi77 := 0; fi77 < 4; fi77++ {
                            v1 = -9223372036854775808
                            fmt.Println(fi77, v1)
                            delete(v3, v5)
                            _ = fi77
                        }
                        v2["foo bar"] = v1
                        fmt.Println(v5)
                    }
                case "\n\t":
                    v73 = v73
                }
                _ = v76
            } else {
                for fi78 := 0; fi78 < 5; fi78++ {
                    for fi79 := 0; fi79 < 16; fi79++ {
                        delete(v2, "foo bar")
                        _ = fi79
                    }
                    _ = fi78
                }
            }
        }
        v73 = (v73 + 42)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v72:", len(v72), v72)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v72
    _ = v73
}
