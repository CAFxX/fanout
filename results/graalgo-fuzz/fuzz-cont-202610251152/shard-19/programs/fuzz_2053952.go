package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32; F1 int16; F2 int8 }
type GenS5 struct { F0 int32; F1 int; F2 int16 }
type GenS6 struct { F0 string; F1 int; F2 string }
type GenS7 struct { F0 bool }

func main() {
    v1 := []int{-2147483648, 1}
    var v2 uint64 = 0
    var v3 uint = 789
    v4 := []string{("" + (("\n\t" + "a") + "a")), "hello", ((("foo bar" + "\n\t") + "\n\t") + "a"), "a"}
    var v5 int8 = 42
    v6 := []uint32{4294967295, 65535, 890}
    v5 += (-(v5))
    var v7 uint64 = (^(v2))
    if true {
        fmt.Println(v3, v7)
        fmt.Println(v5, v7, v5)
        var v8 bool = (v3 < (v3 >> (((^(v3))) & 63)))
        _ = v8
    }
    _ = len(v6)
    v6 = append(v6, 2)
    switch -2147483648 {
    case 1:
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 bool = false
            if (-2.25 < 1.5) {
                v3 += (v3 ^ (v3 / ((253) | 1)))
                for fi11 := 0; fi11 < 5; fi11++ {
                    if (!false) {
                        v12 := GenS1{F0: 2147483648}
                        v13 := GenS1{F0: 65535}
                        _ = v12
                        _ = v13
                    } else {
                        if (!true) {
                            switch "foo bar" {
                            case "\n\t":
                                v7 = v7
                            case "hello":
                                if false {
                                    fmt.Println(v7, v5)
                                }
                            case "foo bar":
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            }
                        }
                    }
                    _ = fi11
                }
            } else {
                var v14 int16 = 945
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v15 int32 = 2147483647
                _ = v14
                _ = v15
            }
            v5 = v5
            _ = fi9
            _ = v10
        }
    case 4294967295:
        var v16 uint64 = (v2 | (v2 >> ((v3) & 63)))
        _ = v16
    case 241:
        if true {
            var v17 int32 = 1
            v5 += (v5 & v5)
            for fi18 := 0; fi18 < 19; fi18++ {
                switch ("foo bar" + "foo bar") {
                case "a":
                    v19 := GenS1{F0: 2147483648}
                    _ = v19
                }
                var v20 int8 = int8(v3)
                _ = fi18
                _ = v20
            }
            _ = v17
        } else {
            for fi21 := 0; fi21 < 13; fi21++ {
                if false {
                    v1 = append(v1, (fi21 % (((fi21 ^ (fi21 << (((v3 ^ v3)) & 63)))) | 1)))
                    v22 := GenS1{F0: 2}
                    if (!false) {
                        v2 = uint64(v3)
                        v3 -= (^(v3))
                    } else {
                        v6 = append(v6, 748)
                    }
                    _ = v22
                } else {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                    if true {
                        var v23 int32 = -296
                        _ = v23
                    } else {
                        fmt.Println(v2)
                        if (42 > 1) {
                            v5 -= v5
                            var v24 uint16 = 0
                            _ = v24
                        } else {
                            for fi25 := 0; fi25 < 10; fi25++ {
                                var v26 uint16 = 42
                                _ = fi25
                                _ = v26
                            }
                            for fi27 := 0; fi27 < 7; fi27++ {
                                v5 = v5
                                _ = fi27
                            }
                        }
                    }
                }
                _ = fi21
            }
            switch 1 {
            case -2147483648:
                if true {
                    v28 := GenS1{F0: 1}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v28
                } else {
                    if ((!(!(v5 > (^(v5))))) && true) {
                        v29 := GenS2{F0: ("" + ("\n\t" + ("" + "foo bar")))}
                        if (!true) {
                            for fi30 := 0; fi30 < 14; fi30++ {
                                v31 := GenS3{F0: 2}
                                _ = fi30
                                _ = v31
                            }
                            fmt.Println(v5, v2)
                            switch "hello" {
                            case "a":
                                var v32 int16 = -1
                                _ = v32
                            default:
                                v7 = v7
                            }
                        }
                        _ = v29
                    } else {
                        v3 = 589
                        var v33 int = 127
                        v33 -= v33
                        _ = v33
                    }
                    if true {
                        if len(v1) > 0 {
                            v1[0] = -2147483648
                        }
                        if false {
                            _ = len(v6)
                        } else {
                            var v34 uint32 = 2147483648
                            if len(v1) > 0 {
                                v1[0] = -1
                            }
                            _ = v34
                        }
                        if false {
                            v35 := GenS3{F0: 124}
                            _ = v35
                        } else {
                            if (!(-1 <= 614)) {
                                for fi36 := 0; fi36 < 14; fi36++ {
                                    var v37 uint64 = 18446744073709551615
                                    fmt.Println(v2, v5, v7)
                                    _ = fi36
                                    _ = v37
                                }
                                fmt.Println(v5, v5)
                            }
                        }
                    } else {
                        switch "\n\t" {
                        case "\n\t":
                            v38 := GenS3{F0: 2}
                            _ = v38
                        case "a":
                            if (2 <= 255) {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                            } else {
                                v5 = (v5 % ((100) | 1))
                                v3 -= (v3 * 209)
                                v2 = uint64(v3)
                            }
                        case "":
                            v5 = int8(v7)
                        default:
                            v39 := GenS4{F0: 978, F1: 0, F2: v5}
                            _ = v39
                        }
                    }
                    var v40 int64 = -1
                    _ = v40
                }
            case 9223372036854775807:
                fmt.Println(v5)
            case 710:
                var v41 float64 = 0.1
                _ = v41
            default:
                v2 += v2
            }
        }
    default:
        for fi42 := 0; fi42 < 6; fi42++ {
            v7 -= 2
            if true {
                for fi43 := 0; fi43 < 13; fi43++ {
                    if (false && (100 < v5)) {
                        var v44 float32 = 3.14159
                        v45 := GenS3{F0: 998}
                        _ = v44
                        _ = v45
                    } else {
                        var v46 uint16 = 468
                        var v47 int16 = 0
                        fmt.Println(fi43)
                        _ = v46
                        _ = v47
                    }
                    _ = fi43
                }
            }
            var v48 int = 127
            _ = fi42
            _ = v48
        }
    }
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    v49 := GenS5{F0: 0, F1: 4294967295, F2: -1}
    if true {
        _ = len(v6)
        for fi50 := 0; fi50 < 1; fi50++ {
            for fi51 := 0; fi51 < 3; fi51++ {
                switch 128 {
                case 862:
                    v1 = append(v1, (fi51 % ((699) | 1)))
                case 9223372036854775807:
                    v5 += v5
                default:
                    v52 := GenS6{F0: "", F1: 4294967295, F2: "\n\t"}
                    _ = v52
                }
                v5 -= (v5 & (^(v5)))
                _ = fi51
            }
            var v53 float32 = (float32(0.0) - 0.0)
            _ = fi50
            _ = v53
        }
        if true {
            for fi54 := 0; fi54 < 17; fi54++ {
                var v55 uint16 = 42
                var v56 int32 = int32(v5)
                for fi57 := 0; fi57 < 13; fi57++ {
                    var v58 uint = (v3 >> ((v3) & 63))
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    var v59 int8 = (v5 + (v5 - (v5 ^ 120)))
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                _ = fi54
                _ = v55
                _ = v56
            }
            switch 128 {
            case -76:
                var v60 int = 127
                _ = v60
            case -439:
                var v61 int64 = -1
                _ = v61
            case 32767:
                switch 32767 {
                case 32767:
                    if true {
                        switch -1 {
                        case 9223372036854775807:
                            if (!true) {
                                v4 = append(v4, "hello")
                                fmt.Println(v2)
                            }
                        default:
                            fmt.Println(v7, v2, v5)
                        }
                    } else {
                        fmt.Println(v2, v3, v5)
                        var v62 int16 = -1
                        _ = v62
                    }
                case -644:
                    var v63 uint64 = v2
                    _ = v63
                default:
                    fmt.Println(v5, v5, v7)
                }
            default:
                if false {
                    var v64 bool = false
                    _ = v64
                }
            }
            if (!true) {
                for fi65 := 0; fi65 < 2; fi65++ {
                    v66 := GenS7{F0: (!true)}
                    var v67 uint32 = 2
                    _ = fi65
                    _ = v66
                    _ = v67
                }
            } else {
                var v68 int = 2147483647
                v5 = -128
                _ = v68
            }
        }
    } else {
        fmt.Println(v3)
    }
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
    _ = v7
}
