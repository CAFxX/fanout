package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 int64; F1 int8; F2 int16 }
type GenS3 struct { F0 uint; F1 uint32; F2 string }
type GenS4 struct { F0 uint32; F1 bool; F2 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint64; F1 string; F2 int8 }
type GenS7 struct { F0 int32; F1 int16; F2 uint }
func genFn0(p0 uint64, p1 uint16) uint64 {
    var v6 uint64 = p0
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 string = "hello"
        var v9 int8 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 15; fi10++ {
        switch (fi10 & fi10) {
        case 32767:
            fmt.Println(v6, fi10)
        }
        v11 := []int{fi10, (fi10 & -9223372036854775808), (^(fi10))}
        _ = fi10
        _ = v11
    }
    if (1 >= 42) {
        v12 := GenS1{F0: -296, F1: ("hello" + (("hello" + "foo bar") + "\n\t"))}
        _ = v12
    } else {
        v13 := []int{32767}
        if len(v13) >= 2 {
            v13 = v13[1:len(v13)]
        }
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v6
    return 2
}
func genFn1() uint {
    switch "" {
    case "\n\t":
        v14 := GenS1{F0: 42, F1: "\n\t"}
        _ = v14
    case "hello":
        var v15 uint = 255
        _ = v15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: 743, F1: "hello"}
        _ = fi16
        _ = v17
    }
    v18 := GenS1{F0: -9223372036854775808, F1: (("a" + "hello") + ("a" + "hello"))}
    var v19 float32 = 0.1
    _ = v18
    _ = v19
    return 65535
}

func main() {
    var v1 bool = (805 >= -823)
    v2 := []int{-640, 414, 128}
    var v3 int32 = -760
    v4 := make(map[uint32]bool)
    v5 := []uint32{0, 236, 779}
    v3 -= 383
    for fi20 := 0; fi20 < 12; fi20++ {
        v3 = (^(v3))
        var v21 float64 = -2.25
        _ = fi20
        _ = v21
    }
    var v22 uint64 = 4294967295
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint32 = 418
        _ = fi23
        _ = v24
    }
    _ = len(v4)
    if v1 {
        var v25 int64 = 604
        if false {
            for fi26 := 0; fi26 < 16; fi26++ {
                for fi27 := 0; fi27 < 8; fi27++ {
                    v22 = genFn0(v22, 191)
                    fmt.Println(v22, fi26)
                    var v28 uint = 980
                    _ = fi27
                    _ = v28
                }
                v29 := GenS1{F0: 0, F1: "\n\t"}
                var v30 uint = 4294967295
                _ = fi26
                _ = v29
                _ = v30
            }
            switch "" {
            case "\n\t":
                var v31 uint64 = 0
                _ = v31
            }
        } else {
            var v32 float64 = 3.14159
            _ = v32
        }
        delete(v4, 2147483648)
        _ = v25
    } else {
        for fi33 := 0; fi33 < 20; fi33++ {
            var v34 float64 = 0.1
            _ = fi33
            _ = v34
        }
        v35 := GenS1{F0: -9223372036854775808, F1: "hello"}
        switch -2147483648 {
        case 101:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case 4294967295:
            v36 := GenS2{F0: 1, F1: -107, F2: -646}
            _ = v36
        case 42:
            v37 := GenS2{F0: 1, F1: -128, F2: 777}
            _ = v37
        default:
            switch "a" {
            case "foo bar":
                if false {
                    v5 = append(v5, 2147483648)
                }
            case "a":
                if false {
                    fmt.Println(v22, v3)
                    if true {
                        v38 := GenS1{F0: 9223372036854775807, F1: "hello"}
                        v22 -= v22
                        for fi39 := 0; fi39 < 15; fi39++ {
                            switch fi39 {
                            case -2147483648:
                                var v40 uint64 = (v22 - (v22 % (((v22 ^ v22)) | 1)))
                                _ = v40
                            case 127:
                                if ((false || ((v22 - genFn0((v22 & v22), 30)) != genFn0(v22, 218))) || (v1 && (!(42 != 175)))) {
                                    for fi41 := 0; fi41 < 18; fi41++ {
                                        v22 += (v22 % (((^(v22))) | 1))
                                        v3 = v3
                                        _ = fi41
                                    }
                                    if v1 {
                                        v22 = v22
                                    } else {
                                        for fi42 := 0; fi42 < 7; fi42++ {
                                            if true {
                                                var v43 bool = ((0.0 <= -2.25) || (-777 >= -1))
                                                var v44 string = "a"
                                                _ = v43
                                                _ = v44
                                            } else {
                                                v4[1] = (3.14159 <= (3.14159 * 1.5))
                                                v5 = append(v5, 502)
                                            }
                                            _ = fi42
                                        }
                                        for fi45 := 0; fi45 < 19; fi45++ {
                                            fmt.Println(v22)
                                            switch fi45 {
                                            case 32767:
                                                v46 := GenS3{F0: 477, F1: 42, F2: ("foo bar" + "foo bar")}
                                                _ = v46
                                            case 2147483647:
                                                var v47 int32 = 42
                                                _ = v47
                                            }
                                            _ = fi45
                                        }
                                    }
                                } else {
                                    for fi48 := 0; fi48 < 7; fi48++ {
                                        _ = len(v2)
                                        switch 1 {
                                        case 127:
                                            fmt.Println(v1)
                                            if (!(!(9223372036854775807 == 42))) {
                                                v1 = v1
                                            }
                                            v49 := GenS4{F0: 2, F1: (true || v1), F2: -828}
                                            _ = v49
                                        case 1:
                                            switch 9223372036854775807 {
                                            case -9223372036854775808:
                                                var v50 int32 = (-(v3))
                                                _ = v50
                                            }
                                        default:
                                            v1 = (2 != genFn1())
                                            fmt.Println(v22, v1, v22)
                                            if (false || true) {
                                                if (!(68 != -98)) {
                                                    fmt.Println(v1)
                                                    if v1 {
                                                        var v51 uint16 = 32767
                                                        v3 += 2147483647
                                                        var v52 uint8 = 255
                                                        var v53 uint16 = (^(v51))
                                                        _ = v51
                                                        _ = v52
                                                        _ = v53
                                                    }
                                                    v54 := GenS5{F0: -32768}
                                                    _ = v54
                                                } else {
                                                    delete(v4, 42)
                                                    v55 := GenS1{F0: 42, F1: ""}
                                                    v56 := GenS2{F0: -9223372036854775808, F1: 127, F2: 32767}
                                                    _ = v55
                                                    _ = v56
                                                }
                                                v3 += (v3 | v3)
                                                _ = len(v4)
                                                v3 = (-(v3))
                                            } else {
                                                var v57 uint64 = (v22 % ((801) | 1))
                                                switch fi48 {
                                                case 1:
                                                    fmt.Println(v1, v57, fi39)
                                                }
                                                _ = v57
                                            }
                                        }
                                        v58 := GenS6{F0: (^(v22)), F1: "a", F2: -1}
                                        var v59 float32 = float32((3.14159 / (1.5 / (1.5 / 1.0))))
                                        var v60 uint32 = 2147483647
                                        _ = fi48
                                        _ = v58
                                        _ = v59
                                        _ = v60
                                    }
                                    var v61 uint8 = 0
                                    var v62 float64 = (1e10 * 0.1)
                                    _ = v61
                                    _ = v62
                                }
                            }
                            fmt.Println(v22, v22)
                            _ = fi39
                        }
                        _ = v38
                    }
                }
            default:
                v3 = v3
            }
        }
        _ = v35
    }
    if v1 {
        for fi63 := 0; fi63 < 8; fi63++ {
            v64 := GenS4{F0: 2147483648, F1: v1, F2: 1}
            var v65 bool = false
            v3 = (v3 + v3)
            _ = fi63
            _ = v64
            _ = v65
        }
    } else {
        var v66 int32 = 1
        switch 1 {
        case 650:
            _ = len(v2)
        case 2147483647:
            v22 = v22
        }
        _ = v66
    }
    v67 := GenS7{F0: v3, F1: -32768, F2: 286}
    for fi68 := 0; fi68 < 1; fi68++ {
        v1 = false
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v67
}
