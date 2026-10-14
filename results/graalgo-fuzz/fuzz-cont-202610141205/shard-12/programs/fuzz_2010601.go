package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16; F2 bool }
type GenS3 struct { F0 uint8; F1 int32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint }
type GenS7 struct { F0 int32; F1 int32; F2 int64 }
type GenS8 struct { F0 int8; F1 int64 }
func genFn0() int8 {
    if (0 > 1) {
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint8 = 1
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 string = "foo bar"
                _ = fi9
                _ = v10
            }
            _ = fi7
            _ = v8
        }
    }
    var v11 int32 = 42
    for fi12 := 0; fi12 < 4; fi12++ {
        var v13 float32 = 3.14159
        _ = fi12
        _ = v13
    }
    v11 -= v11
    v14 := GenS1{F0: 9223372036854775807, F1: 32767}
    _ = v11
    _ = v14
    return 0
}

func main() {
    var v1 int16 = 32767
    var v2 int32 = 627
    var v3 int = -997
    v4 := make(map[string]int)
    var v5 int8 = -1
    v6 := []string{((("\n\t" + "a") + "hello") + "a"), "foo bar", ""}
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 8; fi16++ {
            v17 := GenS2{F0: 2147483647, F1: 42, F2: true}
            switch v3 {
            case 4294967295:
                fmt.Println(fi15, fi15)
            }
            var v18 uint16 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = len(v4)
        _ = fi15
    }
    v5 = (v5 & v5)
    for fi19 := 0; fi19 < 1; fi19++ {
        v3 += fi19
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi19
    }
    v6 = append(v6, ("\n\t" + "\n\t"))
    if (float64((1e10 * 1e10)) == (0.0 / 1.5)) {
        var v20 float32 = ((float32(((1.5 * 1.5) - -2.25)) / 3.14159) * (-2.25 + 3.14159))
        var v21 int32 = (v2 ^ v2)
        for fi22 := 0; fi22 < 15; fi22++ {
            if true {
                fmt.Println(v1)
            }
            var v23 int16 = (v1 - (v1 & -511))
            var v24 int8 = v5
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = v20
        _ = v21
    } else {
        if true {
            if true {
                v25 := GenS1{F0: 1, F1: (v1 % ((v1) | 1))}
                _ = v25
            } else {
                var v26 int8 = (v5 << ((42) & 63))
                if false {
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 int8 = v26
                        v3 = v3
                        _ = fi27
                        _ = v28
                    }
                    switch (v3 + 42) {
                    case -9223372036854775808:
                        if (2 <= 9) {
                            v29 := GenS3{F0: 255, F1: 0}
                            var v30 float64 = 3.14159
                            v31 := GenS4{F0: 128, F1: (v5 % (((v26 + (v26 - (v26 ^ 100)))) | 1)), F2: (v2 >> ((1) & 63))}
                            _ = v29
                            _ = v30
                            _ = v31
                        }
                    case 127:
                        var v32 uint8 = 1
                        _ = v32
                    case 9223372036854775807:
                        for fi33 := 0; fi33 < 19; fi33++ {
                            for fi34 := 0; fi34 < 13; fi34++ {
                                switch fi33 {
                                case -1:
                                    v35 := v4[("\n\t" + "\n\t")]
                                    _ = v35
                                case 728:
                                    if (v1 > (v1 >> ((255) & 63))) {
                                        _ = len(v6)
                                        v36 := GenS1{F0: 9223372036854775807, F1: v1}
                                        _ = v36
                                    } else {
                                        var v37 bool = (!false)
                                        v37 = (5 < 0)
                                        if (!true) {
                                            switch v3 {
                                            case -2147483648:
                                                v26 -= (v5 >> ((0) & 63))
                                            case 128:
                                                v6 = append(v6, "")
                                            case -9223372036854775808:
                                                v3 += (v3 & fi34)
                                            }
                                        }
                                        _ = v37
                                    }
                                    fmt.Println(v2, v1)
                                    var v38 int32 = v2
                                    _ = v38
                                case 4294967295:
                                    if true {
                                        var v39 uint64 = 1
                                        if (-1 != v2) {
                                            v26 += v26
                                            var v40 int32 = (v2 | v2)
                                            _ = v40
                                        }
                                        _ = v39
                                    } else {
                                        v41 := GenS1{F0: 9223372036854775807, F1: v1}
                                        _ = v41
                                    }
                                default:
                                    fmt.Println(v2, v2)
                                }
                                _ = len(v4)
                                _ = fi34
                            }
                            _ = fi33
                        }
                    default:
                        if len(v6) > 0 {
                            v6[0] = "foo bar"
                        }
                    }
                    fmt.Println(v3, v26, v26)
                } else {
                    var v42 uint8 = 53
                    _ = v42
                }
                _ = v26
            }
            v1 = 32767
            var v43 uint32 = 435
            _ = v43
        } else {
            switch v3 {
            case -554:
                if (201 < 0) {
                    v3 += 4294967295
                    var v44 int64 = -300
                    _ = v44
                }
            }
            v45 := GenS1{F0: 976, F1: (v1 >> ((249) & 63))}
            _ = len(v4)
            _ = v45
        }
        v2 += v2
        if (((!(517 == -1)) || true) && false) {
            var v46 int = v3
            var v47 int = 9223372036854775807
            fmt.Println(v3)
            _ = v46
            _ = v47
        }
    }
    var v48 uint64 = 334
    for fi49 := 0; fi49 < 2; fi49++ {
        switch (v3 | v3) {
        case -9223372036854775808:
            if (!false) {
                if ((^(v5)) <= v5) {
                    v50 := GenS5{F0: (v1 << ((18446744073709551615) & 63))}
                    v5 = (v5 >> ((65535) & 63))
                    var v51 uint8 = 127
                    _ = v50
                    _ = v51
                } else {
                    fmt.Println(v1)
                    v6 = append(v6, "\n\t")
                }
                v48 += (v48 - 4294967295)
            } else {
                var v52 int64 = 9223372036854775807
                _ = v52
            }
        case 128:
            switch "" {
            case "foo bar":
                var v53 int16 = (v1 & v1)
                _ = v53
            case "a":
                for fi54 := 0; fi54 < 9; fi54++ {
                    v55 := GenS5{F0: (v1 >> ((42) & 63))}
                    v56 := GenS4{F0: 225, F1: genFn0(), F2: v2}
                    _ = fi54
                    _ = v55
                    _ = v56
                }
            case "hello":
                for fi57 := 0; fi57 < 10; fi57++ {
                    v2 = v2
                    _ = fi57
                }
            default:
                switch ("\n\t" + (("hello" + "") + "a")) {
                case "":
                    v58 := GenS1{F0: 1, F1: (v1 % ((42) | 1))}
                    _ = v58
                case "\n\t":
                    var v59 string = "a"
                    _ = v59
                case "foo bar":
                    if (v5 <= v5) {
                        if false {
                            for fi60 := 0; fi60 < 8; fi60++ {
                                if (475 != 65535) {
                                    v61 := GenS3{F0: 0, F1: (-(v2))}
                                    _ = v61
                                }
                                var v62 uint8 = 128
                                switch "a" {
                                case "":
                                    v63 := GenS5{F0: (^(v1))}
                                    _ = v63
                                case "\n\t":
                                    v1 = 534
                                case "a":
                                    v1 += (v1 % ((42) | 1))
                                default:
                                    fmt.Println(v62)
                                }
                                _ = fi60
                                _ = v62
                            }
                        } else {
                            if true {
                                var v64 int32 = v2
                                _ = v64
                            }
                            v1 -= int16(fi49)
                            switch 118 {
                            case -1:
                                v65 := GenS6{F0: -302, F1: 65535, F2: 250}
                                _ = v65
                            }
                        }
                    }
                default:
                    if (v2 == v2) {
                        fmt.Println(v3)
                        v5 -= v5
                    } else {
                        v3 -= 9223372036854775807
                        v66 := GenS7{F0: (v2 << ((633) & 63)), F1: (v2 << ((42) & 63)), F2: 42}
                        _ = v66
                    }
                }
            }
        }
        _ = fi49
    }
    v67 := GenS8{F0: v5, F1: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
    _ = v67
}
