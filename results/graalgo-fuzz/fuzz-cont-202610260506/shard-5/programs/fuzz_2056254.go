package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 int64; F1 uint32; F2 bool }
type GenS3 struct { F0 int32; F1 uint64; F2 bool }
type GenS4 struct { F0 uint8; F1 int }
type GenS5 struct { F0 int8; F1 uint; F2 uint16 }
type GenS6 struct { F0 uint32; F1 int32; F2 int16 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 uint64; F1 string }

func main() {
    var v1 int8 = 0
    v2 := []string{"\n\t", (("" + "a") + "hello"), "a", "\n\t"}
    v3 := make(map[string]uint32)
    var v4 int32 = 1
    var v5 uint8 = 127
    for fi6 := 0; fi6 < 11; fi6++ {
        if false {
            var v7 int32 = -1
            v8 := GenS1{F0: 0, F1: 42}
            if (0 == 0) {
                v9 := GenS1{F0: 415, F1: 42}
                _ = v9
            }
            _ = v7
            _ = v8
        } else {
            if false {
                var v10 uint = 390
                v4 -= (v4 * (v4 << ((2) & 63)))
                var v11 int64 = int64(v5)
                _ = v10
                _ = v11
            }
            v4 -= v4
            if true {
                for fi12 := 0; fi12 < 15; fi12++ {
                    v5 -= (^(v5))
                    v5 -= (^(v5))
                    _ = fi12
                }
                for fi13 := 0; fi13 < 13; fi13++ {
                    switch fi13 {
                    case 32767:
                        switch (fi6 % (((-(fi13))) | 1)) {
                        case -624:
                            v1 += v1
                        case -652:
                            if (1 > 0) {
                                v14 := GenS1{F0: 1, F1: 1}
                                _ = v14
                            } else {
                                if (false || (v5 <= v5)) {
                                    v15 := GenS2{F0: -807, F1: 2147483647, F2: true}
                                    if (0.1 >= 1.5) {
                                        v4 = (v4 & 1)
                                        v5 = (v5 ^ (v5 - (v5 / ((201) | 1))))
                                        var v16 uint16 = 255
                                        var v17 uint16 = (v16 | 0)
                                        _ = v16
                                        _ = v17
                                    }
                                    _ = v15
                                } else {
                                    v4 -= (v4 ^ (v4 / ((v4) | 1)))
                                }
                                switch fi6 {
                                case 32767:
                                    if (((v5 - v5) < (v5 | (v5 / ((18) | 1)))) || (0.0 == -2.25)) {
                                        if true {
                                            v18 := GenS3{F0: (^(v4)), F1: uint64(v5), F2: ((-(fi13)) < 0)}
                                            v19 := v3["\n\t"]
                                            _ = v18
                                            _ = v19
                                        }
                                    }
                                    var v20 uint = 0
                                    _ = v20
                                case 4294967295:
                                    fmt.Println(v4)
                                }
                            }
                        case -331:
                            v4 += v4
                        }
                    }
                    v4 = v4
                    _ = fi13
                }
                switch 944 {
                case 127:
                    var v21 uint64 = uint64(v5)
                    _ = v21
                default:
                    if true {
                        if false {
                            for fi22 := 0; fi22 < 15; fi22++ {
                                v23 := GenS4{F0: v5, F1: (fi6 - (fi22 >> ((18446744073709551615) & 63)))}
                                v5 = 224
                                _ = fi22
                                _ = v23
                            }
                            var v24 int = fi6
                            if true {
                                switch v24 {
                                case 42:
                                    var v25 uint8 = 255
                                    _ = v25
                                case 526:
                                    switch "" {
                                    case "a":
                                        v1 += (^(v1))
                                    default:
                                        if (float64((float32(-2.25) * float32(3.14159))) != 0.0) {
                                            v26 := GenS5{F0: v1, F1: 4294967295, F2: 42}
                                            fmt.Println(v4, v1)
                                            v27 := GenS3{F0: v4, F1: 4294967295, F2: (false && (!true))}
                                            _ = v26
                                            _ = v27
                                        }
                                    }
                                case 4294967295:
                                    fmt.Println(v5, v24, v4)
                                }
                                v28 := GenS2{F0: -537, F1: 1, F2: (((true && false) && (false && true)) || false)}
                                _ = len(v3)
                                _ = v28
                            }
                            _ = v24
                        }
                        _ = len(v3)
                    }
                }
            }
        }
        _ = fi6
    }
    v4 = -2147483648
    v2 = append(v2, "hello")
    if (false && false) {
        var v29 float64 = 1.5
        var v30 uint8 = v5
        v5 = (v5 * v5)
        _ = v29
        _ = v30
    }
    if true {
        v31 := v3["hello"]
        _ = v31
    } else {
        for fi32 := 0; fi32 < 7; fi32++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi32
        }
        for fi33 := 0; fi33 < 16; fi33++ {
            v1 -= 42
            var v34 uint = 998
            _ = fi33
            _ = v34
        }
        if false {
            var v35 uint = 462
            if true {
                switch "" {
                case "\n\t":
                    v4 += 0
                case "foo bar":
                    if (0 >= 1) {
                        fmt.Println(v35, v1)
                        for fi36 := 0; fi36 < 17; fi36++ {
                            switch fi36 {
                            case 0:
                                for fi37 := 0; fi37 < 9; fi37++ {
                                    var v38 float64 = float64(3.14159)
                                    _ = fi37
                                    _ = v38
                                }
                            default:
                                for fi39 := 0; fi39 < 20; fi39++ {
                                    if true {
                                        if false {
                                            var v40 uint8 = v5
                                            if false {
                                                v41 := GenS2{F0: 0, F1: 2147483648, F2: true}
                                                var v42 bool = false
                                                if ((v35 + 1) >= 0) {
                                                    var v43 int32 = 42
                                                    fmt.Println(v42, v4, fi39)
                                                    _ = v43
                                                } else {
                                                    if (v4 < 861) {
                                                        var v44 uint8 = v40
                                                        v1 = (^(v1))
                                                        v45 := GenS6{F0: 139, F1: (v4 | (v4 >> ((v35) & 63))), F2: -896}
                                                        v1 += 127
                                                        _ = v44
                                                        _ = v45
                                                    } else {
                                                        fmt.Println(v35, fi39, v4)
                                                    }
                                                    v46 := GenS7{F0: v42}
                                                    if v42 {
                                                        v2 = append(v2, ("" + "foo bar"))
                                                        v42 = (!false)
                                                        v47 := GenS8{F0: 9}
                                                        _ = v47
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = "\n\t"
                                                        }
                                                        v42 = ((v1 / ((int8(v35)) | 1)) < v1)
                                                        v40 += 1
                                                        switch fi39 {
                                                        case 127:
                                                            var v48 uint64 = 18446744073709551615
                                                            delete(v3, "\n\t")
                                                            fmt.Println(v40, v1)
                                                            _ = v48
                                                        default:
                                                            v42 = (1 <= 9223372036854775807)
                                                        }
                                                    }
                                                    _ = v46
                                                }
                                                _ = v41
                                                _ = v42
                                            }
                                            var v49 bool = (v4 <= 1)
                                            _ = v40
                                            _ = v49
                                        } else {
                                            var v50 bool = ((v4 >> ((0) & 63)) != int32(fi39))
                                            delete(v3, "a")
                                            v3["foo bar"] = 259
                                            _ = v50
                                        }
                                        v51 := GenS2{F0: 9223372036854775807, F1: 2147483647, F2: true}
                                        _ = v51
                                    } else {
                                        v52 := GenS9{F0: (true && (!(!false))), F1: 4294967295, F2: 218}
                                        _ = v52
                                    }
                                    v5 += 127
                                    _ = fi39
                                }
                            }
                            v1 = (v1 >> ((2) & 63))
                            _ = fi36
                        }
                    }
                case "":
                    var v53 float64 = (0.0 * 0.1)
                    _ = v53
                }
                for fi54 := 0; fi54 < 6; fi54++ {
                    v55 := GenS6{F0: 4294967295, F1: 2147483647, F2: 160}
                    if (!(-686 > -9223372036854775808)) {
                        _ = len(v3)
                        var v56 string = "a"
                        var v57 uint8 = v5
                        _ = v56
                        _ = v57
                    } else {
                        var v58 uint64 = 2
                        for fi59 := 0; fi59 < 2; fi59++ {
                            v2 = append(v2, "\n\t")
                            var v60 int32 = v4
                            _ = fi59
                            _ = v60
                        }
                        _ = v58
                    }
                    switch (fi54 & (^(fi54))) {
                    case 128:
                        var v61 string = "hello"
                        _ = v61
                    case 0:
                        fmt.Println(fi54, v1)
                    case 2147483647:
                        v1 -= (v1 >> ((v35) & 63))
                    }
                    _ = fi54
                    _ = v55
                }
                for fi62 := 0; fi62 < 16; fi62++ {
                    v63 := GenS10{F0: 1}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi62
                    _ = v63
                }
            } else {
                var v64 uint32 = 1
                _ = v64
            }
            _ = v35
        }
    }
    v4 = (v4 ^ v4)
    if false {
        v1 = (v1 * v1)
        if (42 < 0) {
            v65 := GenS11{F0: -1}
            _ = v65
        }
    }
    delete(v3, "")
    v4 = (v4 ^ (v4 >> ((310) & 63)))
    v5 = (^(v5))
    v66 := GenS5{F0: (v1 * (v1 ^ -46)), F1: 2, F2: uint16(v5)}
    for fi67 := 0; fi67 < 19; fi67++ {
        var v68 int16 = -32768
        v69 := GenS12{F0: 846, F1: "hello"}
        _ = fi67
        _ = v68
        _ = v69
    }
    var v70 uint32 = 2147483648
    var v71 int8 = (v1 | (v1 % (((v1 - v1)) | 1)))
    switch "\n\t" {
    case "\n\t":
        _ = len(v3)
    case "hello":
        if (480 < 42) {
            v5 = v5
            v71 += v71
        }
    case "foo bar":
        v1 = 100
    }
    if (4294967295 == 603) {
        v4 = (v4 % (((-(v4))) | 1))
        v70 = (^(v70))
        for fi72 := 0; fi72 < 1; fi72++ {
            v73 := GenS8{F0: -537}
            _ = fi72
            _ = v73
        }
    }
    var v74 uint8 = (^(v5))
    v70 += v70
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v66
    _ = v70
    _ = v71
    _ = v74
}
