package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint8 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 int8; F2 int16 }
type GenS6 struct { F0 bool; F1 bool }
type GenS7 struct { F0 string }
func genFn0(p0 int32, p1 int, p2 string) int16 {
    v6 := GenS1{F0: (p2 + p2), F1: -911}
    v7 := GenS1{F0: p2, F1: 489}
    p0 = p0
    if false {
        v8 := GenS2{F0: "hello"}
        _ = v8
    } else {
        var v9 int64 = 0
        var v10 int = (^(p1))
        var v11 uint64 = 0
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := []int{(p1 / (((-(p1))) | 1)), (p1 / (((p1 - (p1 | 32768))) | 1)), (p1 | 2147483647), p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v6
    _ = v7
    return 42
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    v3 := []string{}
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 147
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 uint64 = 716
        _ = fi13
        _ = v14
    }
    switch "" {
    case "foo bar":
        var v15 float64 = 3.14159
        _ = v15
    }
    if (42 != 128) {
        if true {
            v5 = 228
        }
    }
    fmt.Println(v5)
    _ = len(v3)
    var v16 bool = (824 >= 896)
    for fi17 := 0; fi17 < 7; fi17++ {
        v16 = (v16 || true)
        v16 = (65535 >= 620)
        v16 = (v16 || (fi17 != (fi17 % ((-212) | 1))))
        _ = fi17
    }
    var v18 uint = 1
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            if v16 {
                var v21 float32 = 1.5
                _ = v21
            } else {
                v16 = false
                fmt.Println(fi20, v5)
                var v22 uint8 = (v5 ^ uint8(v18))
                _ = v22
            }
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            if ((!(-1 <= -9223372036854775808)) && (!(42 < 1))) {
                var v23 int32 = -947
                if ((122 == 1) || v16) {
                    v24 := GenS1{F0: "hello", F1: 42}
                    if v16 {
                        var v25 int64 = 0
                        v26 := v4[283]
                        _ = v25
                        _ = v26
                    } else {
                        var v27 uint = 1
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        fmt.Println(fi20)
                        _ = v27
                    }
                    _ = v24
                } else {
                    if v16 {
                        var v28 int = fi19
                        switch "foo bar" {
                        case "\n\t":
                            v18 = 4294967295
                        case "":
                            var v29 int32 = v23
                            _ = v29
                        }
                        v30 := GenS2{F0: "\n\t"}
                        v5 += v5
                        _ = v28
                        _ = v30
                    }
                    switch (fi19 * fi20) {
                    case 32768:
                        v31 := GenS3{F0: (v18 | v18), F1: (v5 | 212)}
                        _ = v31
                    case -9223372036854775808:
                        if v16 {
                            v23 -= v23
                            v32 := GenS3{F0: (v18 - v18), F1: (v5 + 2)}
                            _ = v32
                        }
                    default:
                        if ((false || v16) && v16) {
                            if (false && v16) {
                                v23 -= v23
                                v23 = v23
                                if v16 {
                                    if (v16 || v16) {
                                        v16 = false
                                    } else {
                                        var v33 int8 = 98
                                        if true {
                                            v18 -= v18
                                        }
                                        _ = v33
                                    }
                                } else {
                                    var v34 int32 = -2147483648
                                    v18 = (v18 << ((v18) & 63))
                                    v23 = v34
                                    _ = v34
                                }
                                if len(v3) > 0 {
                                    v3[0] = ("\n\t" + ("a" + (("a" + ("hello" + ("foo bar" + "foo bar"))) + "a")))
                                }
                            }
                            fmt.Println(v5, fi19)
                        } else {
                            var v35 float64 = ((0.0 - (3.14159 - 1e10)) * 3.14159)
                            var v36 uint16 = 255
                            _ = v35
                            _ = v36
                        }
                    }
                    v37 := GenS4{F0: v5}
                    _ = v37
                }
                _ = v23
            } else {
                delete(v4, 897)
                fmt.Println(v16, v16)
            }
            _ = fi20
        }
        for fi38 := 0; fi38 < 11; fi38++ {
            v18 = (v18 % (((v18 << ((v18) & 63))) | 1))
            _ = fi38
        }
        _ = fi19
    }
    v16 = ((((false && (65535 < 597)) && v16) && false) && true)
    for fi39 := 0; fi39 < 20; fi39++ {
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 uint32 = 409
            _ = fi40
            _ = v41
        }
        var v42 float32 = float32(0.1)
        var v43 uint32 = 4294967295
        _ = fi39
        _ = v42
        _ = v43
    }
    for fi44 := 0; fi44 < 19; fi44++ {
        switch ("foo bar" + "") {
        case "hello":
            if v16 {
                if (-1 == -1) {
                    var v45 uint = (v18 * (v18 >> ((v18) & 63)))
                    _ = v45
                } else {
                    var v46 int = fi44
                    _ = v46
                }
                var v47 uint64 = 0
                _ = v47
            } else {
                for fi48 := 0; fi48 < 16; fi48++ {
                    v49 := GenS3{F0: (v18 + v18), F1: (v5 / (((v5 >> ((466) & 63))) | 1))}
                    v16 = v16
                    var v50 float32 = 0.1
                    _ = fi48
                    _ = v49
                    _ = v50
                }
                for fi51 := 0; fi51 < 4; fi51++ {
                    var v52 int16 = -32768
                    if ((fi51 + 2147483647) <= fi44) {
                        switch 32767 {
                        case 128:
                            v53 := GenS5{F0: 546, F1: -128, F2: v52}
                            _ = v53
                        }
                        var v54 uint64 = 1
                        v55 := GenS5{F0: 36, F1: -28, F2: v52}
                        var v56 int8 = 100
                        if (755 <= 1) {
                            _ = len(v3)
                            switch "hello" {
                            case "\n\t":
                                if (!v16) {
                                    fmt.Println(v5, v54)
                                    v16 = (v16 && ((42 == 42) || (v52 < genFn0(1, (fi51 | fi44), ("\n\t" + ((("a" + "foo bar") + "a") + ""))))))
                                } else {
                                    v18 = (^(v18))
                                    _ = len(v1)
                                    v56 -= (v56 << ((v18) & 63))
                                    _ = len(v3)
                                }
                            case "a":
                                var v57 int8 = v56
                                _ = v57
                            case "foo bar":
                                v58 := GenS5{F0: 1, F1: v56, F2: -1}
                                _ = v58
                            default:
                                v59 := GenS6{F0: v16, F1: v16}
                                _ = v59
                            }
                            if ((!v16) || v16) {
                                v60 := GenS7{F0: ("hello" + ("\n\t" + "a"))}
                                var v61 int32 = 1
                                _ = v60
                                _ = v61
                            } else {
                                var v62 int8 = int8(v18)
                                _ = v62
                            }
                            v16 = (true && true)
                        } else {
                            v63 := GenS5{F0: 0, F1: -67, F2: v52}
                            var v64 int32 = -1
                            v65 := GenS1{F0: ("\n\t" + ""), F1: 96}
                            _ = v63
                            _ = v64
                            _ = v65
                        }
                        _ = v54
                        _ = v55
                        _ = v56
                    } else {
                        if (!(!v16)) {
                            var v66 uint16 = 65535
                            var v67 int32 = 141
                            v68 := v4[4294967295]
                            _ = v66
                            _ = v67
                            _ = v68
                        } else {
                            var v69 float32 = 0.1
                            if false {
                                var v70 int64 = int64(v18)
                                _ = v70
                            } else {
                                v69 += (float32(1e10) - v69)
                            }
                            if (-9223372036854775808 <= 272) {
                                if (!((v69 - v69) <= ((v69 / (1.5 * float32(-2.25))) * float32(3.14159)))) {
                                    fmt.Println(v18, v18, fi44)
                                    switch "" {
                                    case "a":
                                        _ = len(v2)
                                    case "\n\t":
                                        v69 -= v69
                                    case "hello":
                                        fmt.Println(v52, fi51)
                                    default:
                                        var v71 float64 = (float64(((0.1 / 0.1) / (3.14159 / -2.25))) + 1e10)
                                        _ = v71
                                    }
                                    v72 := GenS7{F0: ""}
                                    _ = v72
                                } else {
                                    var v73 uint32 = 604
                                    _ = v73
                                }
                            } else {
                                var v74 int16 = (v52 ^ v52)
                                _ = v74
                            }
                            _ = v69
                        }
                    }
                    _ = fi51
                    _ = v52
                }
                _ = len(v4)
            }
        case "a":
            v18 = (v18 % (((v18 << ((204) & 63))) | 1))
        case "foo bar":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        default:
            switch "a" {
            case "\n\t":
                for fi75 := 0; fi75 < 20; fi75++ {
                    _ = len(v3)
                    var v76 float64 = -2.25
                    v76 = (((v76 - v76) - v76) * 3.14159)
                    _ = fi75
                    _ = v76
                }
            default:
                var v77 int16 = -1
                _ = v77
            }
        }
        _ = fi44
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
