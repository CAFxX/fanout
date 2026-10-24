package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 string; F1 int }
type GenS7 struct { F0 int64; F1 uint8 }
func genFn0(p0 uint, p1 int8, p2 uint8) int16 {
    var v5 uint = p0
    v6 := GenS1{F0: v5, F1: 0, F2: 32767}
    v7 := GenS2{F0: 127, F1: -32768, F2: ("\n\t" + "foo bar")}
    v8 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 217
}

func main() {
    var v1 int32 = -105
    var v2 uint32 = 0
    var v3 uint64 = 4294967295
    var v4 int8 = 1
    for fi9 := 0; fi9 < 19; fi9++ {
        switch ("a" + ("" + "hello")) {
        case "\n\t":
            v10 := []int{int(v3), (fi9 | (-(fi9))), fi9}
            _ = v10
        case "foo bar":
            var v11 int = (fi9 & -9223372036854775808)
            _ = v11
        case "a":
            var v12 bool = false
            _ = v12
        default:
            if ("foo bar" <= ("a" + ("a" + "\n\t"))) {
                fmt.Println(v2, v4, v3)
                v13 := []int{-1, 32768, (fi9 % (((fi9 / ((-603) | 1))) | 1))}
                _ = v13
            } else {
                v14 := GenS3{F0: 42}
                switch fi9 {
                case 0:
                    for fi15 := 0; fi15 < 11; fi15++ {
                        v16 := GenS4{F0: (v2 >> ((1) & 63)), F1: v4, F2: 165}
                        fmt.Println(v3)
                        var v17 uint8 = 127
                        _ = fi15
                        _ = v16
                        _ = v17
                    }
                case 32768:
                    var v18 bool = true
                    _ = v18
                case 241:
                    v19 := GenS3{F0: (-(v4))}
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        if false {
                            v4 -= v4
                        }
                        if true {
                            var v21 uint64 = v3
                            _ = v21
                        }
                        v22 := []int{-497}
                        v3 = v3
                        _ = fi20
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 6; fi23++ {
                    v24 := GenS4{F0: 2, F1: v4, F2: 255}
                    var v25 string = ("a" + (("a" + "a") + "hello"))
                    v26 := GenS3{F0: (v4 * v4)}
                    v27 := []int{fi23}
                    _ = fi23
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                _ = v14
            }
        }
        var v28 uint64 = 18446744073709551615
        _ = fi9
        _ = v28
    }
    v29 := GenS2{F0: 0, F1: int16(v3), F2: ("hello" + "a")}
    fmt.Println(v4, v3)
    v30 := GenS1{F0: 255, F1: 42, F2: 2}
    var v31 int8 = -1
    v32 := []int{1, 32768, -9223372036854775808}
    v3 = (v3 * (v3 << ((1) & 63)))
    v3 -= (v3 + v3)
    var v33 int = -2147483648
    if false {
        if (!false) {
            var v34 string = "hello"
            _ = v34
        } else {
            v35 := GenS2{F0: v33, F1: -32768, F2: ("" + "")}
            if false {
                v36 := GenS5{F0: 127}
                _ = v36
            } else {
                for fi37 := 0; fi37 < 2; fi37++ {
                    for fi38 := 0; fi38 < 14; fi38++ {
                        v39 := GenS6{F0: "a", F1: (v33 % ((-250) | 1))}
                        _ = fi38
                        _ = v39
                    }
                    v2 = v2
                    v3 = v3
                    _ = fi37
                }
                _ = len(v32)
                if len(v32) >= 2 {
                    v32 = v32[1:len(v32)]
                }
            }
            _ = v35
        }
        for fi40 := 0; fi40 < 8; fi40++ {
            _ = len(v32)
            _ = fi40
        }
    }
    switch ("" + "\n\t") {
    case "":
        _ = len(v32)
    case "foo bar":
        v1 = (v1 | 2147483647)
    case "\n\t":
        var v41 uint16 = 255
        _ = v41
    }
    v42 := GenS7{F0: 9223372036854775807, F1: 1}
    if (18446744073709551615 <= 1) {
        switch v33 {
        case -228:
            v4 = 127
        case 127:
            v1 = int32(v33)
        case -1:
            var v43 uint = 11
            _ = v43
        }
    } else {
        if len(v32) > 0 {
            v32[0] = (^(v33))
        }
        v2 = 310
        switch 846 {
        case 32768:
            v1 -= (v1 >> ((455) & 63))
        case -548:
            for fi44 := 0; fi44 < 14; fi44++ {
                if len(v32) > 0 {
                    v32[0] = (v33 - (^(fi44)))
                }
                var v45 uint16 = 65535
                switch 9223372036854775807 {
                case -832:
                    var v46 float64 = 1.5
                    _ = v46
                case 1:
                    var v47 int8 = (v31 | -77)
                    _ = v47
                default:
                    fmt.Println(v3, v4, fi44)
                }
                _ = fi44
                _ = v45
            }
        case 1:
            for fi48 := 0; fi48 < 8; fi48++ {
                switch "foo bar" {
                case "\n\t":
                    var v49 int8 = v4
                    _ = v49
                case "a":
                    v50 := GenS2{F0: -816, F1: genFn0(501, 29, 255), F2: "a"}
                    _ = v50
                case "foo bar":
                    if len(v32) >= 2 {
                        v32 = v32[1:len(v32)]
                    }
                }
                fmt.Println(v31, fi48)
                if (int16(v33) <= 1000) {
                    for fi51 := 0; fi51 < 15; fi51++ {
                        var v52 int16 = -266
                        if (0 > (fi51 % ((2147483647) | 1))) {
                            if len(v32) >= 2 {
                                v32 = v32[1:len(v32)]
                            }
                        }
                        v53 := GenS6{F0: "foo bar", F1: (^(v33))}
                        v54 := GenS5{F0: 204}
                        if false {
                            fmt.Println(fi51)
                        } else {
                            v2 += (v2 | 207)
                            fmt.Println(v1)
                            v32 = append(v32, fi51)
                        }
                        if len(v32) > 0 {
                            v32[0] = 4294967295
                        }
                        _ = fi51
                        _ = v52
                        _ = v53
                        _ = v54
                    }
                    var v55 int64 = int64(v2)
                    var v56 int64 = -884
                    _ = v55
                    _ = v56
                }
                _ = fi48
            }
        default:
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
        }
    }
    if len(v32) >= 2 {
        v32 = v32[1:len(v32)]
    }
    var v57 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v42
    _ = v57
}
