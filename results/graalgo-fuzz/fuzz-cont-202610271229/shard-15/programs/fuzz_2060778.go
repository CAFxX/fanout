package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 bool; F1 int16; F2 int32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint32; F1 int8; F2 uint }
type GenS7 struct { F0 int8; F1 uint16 }
type GenS8 struct { F0 uint64; F1 bool }
func genFn0(p0 int) uint64 {
    v4 := GenS1{F0: "foo bar", F1: 100}
    var v5 string = ""
    v6 := GenS2{F0: 65535, F1: ""}
    v7 := []int{(p0 & p0), 0}
    for fi8 := 0; fi8 < 2; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
            if (128 <= 127) {
                var v10 int64 = 0
                p0 = -9223372036854775808
                var v11 float32 = (1e10 / float32(0.1))
                _ = v10
                _ = v11
            } else {
                if true {
                    v5 = "a"
                    if ((v5 + v5) <= v5) {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v12 int64 = -1
                        var v13 uint64 = 42
                        _ = v12
                        _ = v13
                    } else {
                        p0 -= fi8
                    }
                    if (false && true) {
                        var v14 uint64 = 971
                        v15 := GenS1{F0: v5, F1: 42}
                        var v16 string = v5
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                }
                var v17 string = v5
                p0 = (fi8 % (((fi9 | (fi9 >> ((65535) & 63)))) | 1))
                _ = v17
            }
            _ = fi9
        }
        var v18 float32 = 3.14159
        v19 := GenS3{F0: true, F1: -32768, F2: 2147483647}
        _ = fi8
        _ = v18
        _ = v19
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    return 2
}

func main() {
    var v1 int64 = 1
    var v2 bool = false
    v3 := []string{"\n\t", ("a" + ("a" + (("\n\t" + "") + "foo bar"))), "\n\t"}
    if (!false) {
        v20 := GenS2{F0: 18446744073709551615, F1: "\n\t"}
        _ = v20
    } else {
        var v21 uint8 = 127
        v21 = (v21 << ((65535) & 63))
        _ = v21
    }
    var v22 uint64 = 2
    switch 1 {
    case 0:
        v2 = true
    case 32767:
        v22 = (^(v22))
    default:
        if (!true) {
            for fi23 := 0; fi23 < 16; fi23++ {
                if true {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        var v25 int = fi23
                        _ = fi24
                        _ = v25
                    }
                }
                if len(v3) > 0 {
                    v3[0] = "hello"
                }
                switch "foo bar" {
                case "a":
                    v1 += v1
                case "\n\t":
                    var v26 uint64 = genFn0(fi23)
                    _ = v26
                case "":
                    var v27 bool = (42 < -108)
                    _ = v27
                default:
                    if false {
                        for fi28 := 0; fi28 < 17; fi28++ {
                            switch "foo bar" {
                            case "hello":
                                v2 = (-1 >= 9223372036854775807)
                            case "foo bar":
                                v2 = v2
                            }
                            var v29 int16 = 0
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi28
                            _ = v29
                        }
                    } else {
                        if v2 {
                            for fi30 := 0; fi30 < 8; fi30++ {
                                if (("foo bar" + (("\n\t" + "\n\t") + "")) <= ("foo bar" + (("a" + ("hello" + "hello")) + ("a" + "")))) {
                                    switch "" {
                                    case "hello":
                                        v31 := GenS3{F0: ((127 <= 100) || true), F1: -739, F2: 0}
                                        _ = v31
                                    case "":
                                        var v32 int32 = -682
                                        _ = v32
                                    }
                                    v2 = (2 >= 368)
                                    v1 = v1
                                } else {
                                    v33 := GenS3{F0: (42 <= 1000), F1: 0, F2: 2147483647}
                                    switch ("\n\t" + "") {
                                    case "":
                                        v34 := GenS4{F0: 4294967295}
                                        _ = v34
                                    case "\n\t":
                                        v35 := GenS3{F0: v2, F1: 1, F2: -746}
                                        _ = v35
                                    }
                                    switch fi30 {
                                    case -1:
                                        v2 = v2
                                    }
                                    _ = v33
                                }
                                _ = fi30
                            }
                            fmt.Println(v2)
                            v22 = (v22 - v22)
                        } else {
                            v36 := GenS2{F0: 42, F1: ("hello" + "")}
                            _ = v36
                        }
                        if v2 {
                            v37 := GenS3{F0: v2, F1: 1000, F2: -355}
                            _ = v37
                        }
                        if true {
                            fmt.Println(v1, v2)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    }
                }
                _ = fi23
            }
            v22 += v22
            v3 = append(v3, "hello")
        }
    }
    for fi38 := 0; fi38 < 7; fi38++ {
        v39 := GenS5{F0: 0}
        if (!(int(v22) < fi38)) {
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                var v40 uint = 0
                v2 = v2
                var v41 uint8 = 42
                _ = v40
                _ = v41
            }
            switch "" {
            case "":
                fmt.Println(fi38, v2)
            case "foo bar":
                for fi42 := 0; fi42 < 15; fi42++ {
                    if (2 >= 639) {
                        if (136 == 2) {
                            if len(v3) > 0 {
                                v3[0] = "a"
                            }
                            var v43 int = -2147483648
                            _ = v43
                        }
                        switch 2147483647 {
                        case 9223372036854775807:
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                        case -141:
                            switch "foo bar" {
                            case "\n\t":
                                var v44 bool = (1000 < 1)
                                _ = v44
                            default:
                                switch (fi38 * fi38) {
                                case -116:
                                    v2 = v2
                                case 2147483647:
                                    var v45 bool = v2
                                    _ = v45
                                default:
                                    var v46 float32 = -2.25
                                    _ = v46
                                }
                            }
                            v47 := GenS6{F0: 2147483647, F1: -115, F2: 4294967295}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v22 -= (v22 - v22)
                            var v48 int32 = -946
                            _ = v47
                            _ = v48
                        }
                    }
                    switch ("a" + (("foo bar" + "") + "foo bar")) {
                    case "\n\t":
                        v49 := GenS2{F0: 476, F1: ("a" + ("a" + "foo bar"))}
                        _ = v49
                    case "a":
                        fmt.Println(v22, v22, v1)
                    case "":
                        fmt.Println(fi38)
                    default:
                        if len(v3) > 0 {
                            v3[0] = (("hello" + "a") + ("a" + "hello"))
                        }
                    }
                    switch (fi42 + (fi38 / ((-1) | 1))) {
                    case 4294967295:
                        if false {
                            var v50 uint64 = 0
                            if v2 {
                                var v51 uint64 = (v50 ^ v50)
                                var v52 int16 = 132
                                v51 -= 730
                                var v53 uint64 = (v51 >> ((4294967295) & 63))
                                v52 -= 1
                                var v54 int16 = (v52 << ((uint(v53)) & 63))
                                _ = v51
                                _ = v52
                                _ = v53
                                _ = v54
                            }
                            v22 += v50
                            _ = v50
                        }
                        switch (^(fi42)) {
                        case 127:
                            if false {
                                if false {
                                    fmt.Println(v2, v2)
                                    if (v2 && v2) {
                                        switch ("hello" + ("hello" + "")) {
                                        case "hello":
                                            switch "hello" {
                                            case "":
                                                var v55 int = (^(fi38))
                                                _ = v55
                                            default:
                                                if len(v3) > 0 {
                                                    v3[0] = ("foo bar" + "\n\t")
                                                }
                                            }
                                        }
                                        v2 = ((83 < 115) && (3.14159 > (0.0 + -2.25)))
                                    } else {
                                        var v56 int8 = 127
                                        v57 := GenS2{F0: 42, F1: ""}
                                        var v58 uint = 885
                                        switch (fi42 >> ((v58) & 63)) {
                                        case -105:
                                            if v2 {
                                                v59 := GenS7{F0: 100, F1: 65535}
                                                _ = v59
                                            }
                                        }
                                        var v60 int16 = 0
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                        _ = v60
                                    }
                                }
                                var v61 uint16 = 2
                                _ = v61
                            }
                        }
                    default:
                        fmt.Println(v2, fi42, fi42)
                    }
                    _ = fi42
                }
            case "\n\t":
                var v62 int = (fi38 - fi38)
                _ = v62
            }
        }
        for fi63 := 0; fi63 < 16; fi63++ {
            if v2 {
                v64 := GenS1{F0: "", F1: -62}
                _ = v64
            }
            _ = fi63
        }
        _ = fi38
        _ = v39
    }
    v2 = true
    for fi65 := 0; fi65 < 20; fi65++ {
        v66 := GenS8{F0: (v22 % (((^(v22))) | 1)), F1: true}
        _ = fi65
        _ = v66
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v67 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v67
}
