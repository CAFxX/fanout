package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int16; F2 uint32 }
type GenS3 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 uint64; F1 string; F2 string }
type GenS6 struct { F0 int8; F1 uint8; F2 bool }
func genFn0() uint8 {
    var v4 bool = ((429 != -819) && true)
    if (237 < 65535) {
        var v5 uint32 = 746
        var v6 uint8 = 2
        _ = v5
        _ = v6
    }
    _ = v4
    return 11
}
func genFn1(p0 int8) int8 {
    if (1000 == 52) {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint = 2
            v9 := GenS1{F0: 42}
            v8 += (v8 >> ((v8) & 63))
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        v10 := []int{-9223372036854775808, -9223372036854775808, -855, -471}
        _ = v10
    }
    if false {
        v11 := GenS2{F0: 71, F1: -1, F2: 1}
        v12 := GenS2{F0: 986, F1: -1, F2: 2}
        p0 -= p0
        _ = v11
        _ = v12
    }
    _ = p0
    return p0
}

func main() {
    var v1 uint = 4294967295
    var v2 uint = 921
    v3 := []int{-411, 42, -263}
    fmt.Println(v1, v2)
    if ("\n\t" > "foo bar") {
        v2 = v2
        switch "\n\t" {
        case "foo bar":
            v2 -= 18446744073709551615
        }
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            if ((!false) && (true || (!(1 == -2147483648)))) {
                if len(v3) > 0 {
                    v3[0] = (-(fi13))
                }
                var v14 float64 = 1.5
                switch "\n\t" {
                case "\n\t":
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v14 = 0.1
                        v3 = append(v3, (fi15 & (fi15 ^ fi15)))
                        _ = fi15
                    }
                case "hello":
                    var v16 uint8 = 1
                    _ = v16
                case "foo bar":
                    var v17 bool = true
                    _ = v17
                default:
                    fmt.Println(v1, v1)
                }
                _ = v14
            }
            for fi18 := 0; fi18 < 17; fi18++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if true {
                    switch "foo bar" {
                    case "a":
                        var v19 int16 = 1000
                        _ = v19
                    }
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    if (((fi13 != fi13) || (((1e10 + float32(0.1)) <= 1.5) || (42 <= 2))) && false) {
                        v20 := GenS3{F0: 804, F1: 0, F2: 851}
                        switch (fi18 & (fi18 & -9223372036854775808)) {
                        case 2147483647:
                            v1 -= (v2 << ((2) & 63))
                        }
                        v3 = append(v3, (^(fi18)))
                        _ = v20
                    } else {
                        var v21 int32 = 2147483647
                        fmt.Println(v1, v1)
                        if true {
                            fmt.Println(fi13, v2)
                            v22 := GenS4{F0: 128, F1: 4294967295}
                            v2 = (v2 >> ((4294967295) & 63))
                            _ = v22
                        }
                        v23 := GenS4{F0: 2, F1: 18446744073709551615}
                        _ = v21
                        _ = v23
                    }
                    if (42 == -1) {
                        switch fi13 {
                        case 2147483647:
                            switch (fi13 & 4294967295) {
                            case 128:
                                var v24 int8 = 103
                                _ = v24
                            case 116:
                                v3 = append(v3, fi13)
                            case 2147483647:
                                var v25 int = fi18
                                _ = v25
                            default:
                                if false {
                                    v1 += (v2 | (v2 + v2))
                                    var v26 int8 = int8(fi13)
                                    v27 := GenS2{F0: -400, F1: 118, F2: 65535}
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case 1:
                            var v28 uint16 = 42
                            _ = v28
                        default:
                            var v29 uint16 = 2
                            _ = v29
                        }
                        fmt.Println(fi18, fi13, fi18)
                    } else {
                        fmt.Println(fi13, v1, fi13)
                    }
                } else {
                    v30 := GenS5{F0: 1, F1: "foo bar", F2: "hello"}
                    _ = v30
                }
                v31 := GenS1{F0: 18446744073709551615}
                _ = fi18
                _ = v31
            }
            v2 -= v1
            _ = fi13
        }
    }
    var v32 uint64 = 0
    var v33 string = (("hello" + "foo bar") + "foo bar")
    var v34 uint32 = 722
    v32 -= (v32 % ((18446744073709551615) | 1))
    if (((3.14159 * 1.5) / ((1e10 + 3.14159) + 0.1)) == -2.25) {
        if true {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        var v35 int16 = 860
        _ = v35
    } else {
        v36 := GenS2{F0: 32767, F1: 949, F2: v34}
        _ = v36
    }
    v37 := GenS1{F0: (v2 % (((v1 & 539)) | 1))}
    if (genFn1(-107) == -40) {
        var v38 int64 = 0
        if ((v34 != v34) && ((false && false) || true)) {
            var v39 float32 = ((0.0 + float32(1e10)) * (-2.25 + -2.25))
            _ = v39
        } else {
            for fi40 := 0; fi40 < 16; fi40++ {
                v41 := GenS5{F0: v32, F1: v33, F2: (v33 + (v33 + v33))}
                _ = fi40
                _ = v41
            }
        }
        _ = v38
    } else {
        v42 := GenS2{F0: -32768, F1: 42, F2: v34}
        var v43 uint8 = 192
        v43 = v43
        _ = v42
        _ = v43
    }
    v34 -= (v34 * v34)
    switch ((v33 + v33) + ("a" + "hello")) {
    case "":
        fmt.Println(v2, v1)
    case "foo bar":
        switch -1 {
        case -2147483648:
            var v44 uint8 = 42
            _ = v44
        }
    default:
        if (128 < 102) {
            for fi45 := 0; fi45 < 5; fi45++ {
                fmt.Println(fi45)
                _ = fi45
            }
            fmt.Println(v33)
            v1 -= (v1 ^ 42)
        }
    }
    for fi46 := 0; fi46 < 9; fi46++ {
        v32 = (v32 & 42)
        _ = fi46
    }
    v3 = append(v3, 951)
    for fi47 := 0; fi47 < 15; fi47++ {
        switch 229 {
        case 127:
            fmt.Println(v34, v34)
        case -499:
            if false {
                switch "\n\t" {
                case "foo bar":
                    if true {
                        v48 := GenS2{F0: 0, F1: 0, F2: (^(v34))}
                        _ = v48
                    }
                default:
                    v3 = append(v3, fi47)
                }
            } else {
                switch -241 {
                case 42:
                    for fi49 := 0; fi49 < 6; fi49++ {
                        v50 := GenS6{F0: -128, F1: genFn0(), F2: (false || false)}
                        var v51 string = v33
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi49
                        _ = v50
                        _ = v51
                    }
                case 127:
                    if true {
                        var v52 int8 = 127
                        for fi53 := 0; fi53 < 4; fi53++ {
                            if true {
                                fmt.Println(v1, fi47, v2)
                                v54 := GenS4{F0: 127, F1: (v32 - (v32 * v32))}
                                _ = v54
                            } else {
                                v33 = v33
                                switch (^(fi47)) {
                                case 694:
                                    v55 := GenS4{F0: 128, F1: (^(v32))}
                                    _ = v55
                                }
                                v2 -= (^(v1))
                            }
                            _ = fi53
                        }
                        _ = v52
                    } else {
                        v2 -= v1
                        for fi56 := 0; fi56 < 16; fi56++ {
                            v57 := GenS2{F0: int16(v1), F1: 1, F2: 2147483647}
                            fmt.Println(fi56, fi47, fi47)
                            _ = fi56
                            _ = v57
                        }
                    }
                }
                var v58 int = (fi47 & (fi47 & fi47))
                fmt.Println(v34, fi47, v1)
                _ = v58
            }
        }
        if len(v3) > 0 {
            v3[0] = (fi47 ^ -67)
        }
        _ = fi47
    }
    _ = len(v3)
    if true {
        v32 = (v32 >> ((v2) & 63))
        v33 = v33
        for fi59 := 0; fi59 < 1; fi59++ {
            var v60 uint64 = v32
            _ = fi59
            _ = v60
        }
    }
    var v61 uint32 = v34
    var v62 uint = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v37
    _ = v61
    _ = v62
}
