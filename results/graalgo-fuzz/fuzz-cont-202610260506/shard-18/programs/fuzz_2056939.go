package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int32; F1 int64 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint }
type GenS4 struct { F0 int; F1 int; F2 uint8 }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 bool; F1 bool; F2 int64 }
type GenS7 struct { F0 string; F1 int8; F2 uint }
type GenS8 struct { F0 int8 }
func genFn0(p0 int8) bool {
    var v4 int16 = 32767
    var v5 string = "foo bar"
    var v6 int16 = v4
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 -= 1000
        v8 := []int{9223372036854775807, 0}
        var v9 bool = (2 == 127)
        _ = fi7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return (9223372036854775807 > 9223372036854775807)
}
func genFn1(p0 bool, p1 int32) string {
    fmt.Println(p0, p1)
    var v10 int = 42
    _ = p0
    _ = p1
    _ = v10
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[int]int)
    v3 := make(map[int]int)
    v11 := []int{2147483647, -2147483648, 490}
    var v12 float64 = float64(1e10)
    v13 := GenS1{F0: 203, F1: 42}
    if true {
        v14 := GenS1{F0: 18446744073709551615, F1: 65535}
        if len(v11) >= 2 {
            v11 = v11[1:len(v11)]
        }
        var v15 uint16 = 1
        _ = v14
        _ = v15
    }
    switch -1 {
    case 127:
        _ = len(v1)
    case -2147483648:
        for fi16 := 0; fi16 < 10; fi16++ {
            for fi17 := 0; fi17 < 14; fi17++ {
                delete(v1, "foo bar")
                _ = fi17
            }
            v18 := GenS1{F0: 42, F1: 2147483648}
            _ = fi16
            _ = v18
        }
    default:
        v12 = (float64(3.14159) - v12)
    }
    v12 = (v12 / ((v12 + v12) / 0.1))
    v19 := GenS1{F0: 728, F1: 4294967295}
    var v20 float32 = (float32(3.14159) * -2.25)
    if false {
        fmt.Println(v12, v20)
        for fi21 := 0; fi21 < 1; fi21++ {
            v20 += 3.14159
            if true {
                v20 -= 0.1
                for fi22 := 0; fi22 < 12; fi22++ {
                    switch "\n\t" {
                    case "hello":
                        var v23 int64 = -1
                        _ = len(v3)
                        _ = v23
                    case "a":
                        switch "\n\t" {
                        case "\n\t":
                            switch fi22 {
                            case -253:
                                switch fi21 {
                                case 138:
                                    fmt.Println(v20, fi21)
                                case 127:
                                    if len(v11) >= 2 {
                                        v11 = v11[1:len(v11)]
                                    }
                                case -9223372036854775808:
                                    var v24 int8 = -1
                                    _ = v24
                                }
                            case 42:
                                v12 = float64((1.5 + v20))
                            case 114:
                                v12 = v12
                            }
                        case "":
                            v25 := GenS1{F0: 42, F1: 2}
                            _ = v25
                        }
                    default:
                        if (false || (2147483647 <= 4294967295)) {
                            v26 := GenS2{F0: -1, F1: 93}
                            _ = v26
                        } else {
                            fmt.Println(fi21)
                            var v27 float64 = v12
                            _ = v27
                        }
                    }
                    var v28 string = ""
                    if false {
                        fmt.Println(v28, fi22, v28)
                    } else {
                        if true {
                            v29 := GenS3{F0: 42, F1: 261, F2: 42}
                            _ = v29
                        }
                        var v30 int32 = 106
                        _ = len(v11)
                        var v31 int16 = 757
                        v32 := GenS4{F0: 9223372036854775807, F1: fi22, F2: 42}
                        v31 = (v31 * (v31 & (v31 + v31)))
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = fi22
                    _ = v28
                }
                var v33 string = ("\n\t" + ("foo bar" + "a"))
                _ = v33
            } else {
                if (4294967295 == 2147483647) {
                    v2[(fi21 | (fi21 % ((fi21) | 1)))] = (fi21 ^ (-(fi21)))
                    var v34 int8 = 127
                    for fi35 := 0; fi35 < 3; fi35++ {
                        if (406 < 1) {
                            var v36 int16 = 32767
                            _ = v36
                        } else {
                            fmt.Println(v12, v12, fi21)
                            var v37 uint64 = 761
                            v34 = -1
                            switch "" {
                            case "a":
                                v2[fi21] = (fi35 >> ((2) & 63))
                            }
                            switch (fi35 / ((fi21) | 1)) {
                            case 1:
                                v38 := GenS5{F0: 42, F1: v37}
                                _ = v38
                            case -2147483648:
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                            case 4294967295:
                                if (!((v20 / 0.1) <= (float32(((v12 / v12) + float64(v20))) / ((((1.5 / 3.14159) - 3.14159) - float32((v12 * v12))) - (v20 * 0.0))))) {
                                    v20 -= v20
                                    fmt.Println(v12, v37, v34)
                                }
                            }
                            _ = v37
                        }
                        v12 = v12
                        _ = fi35
                    }
                    _ = v34
                } else {
                    v12 -= v12
                    if len(v11) > 0 {
                        v11[0] = fi21
                    }
                }
                v12 -= (1.5 / 1.0)
                v20 += v20
            }
            _ = fi21
        }
    }
    v20 = 1.5
    switch ("\n\t" + ("a" + "\n\t")) {
    case "foo bar":
        v39 := GenS6{F0: (987 <= 1), F1: (((float64((v20 * (v20 / v20))) - v12) + 1.5) >= float64(((v20 * (0.0 / (3.14159 * v20))) + v20))), F2: 773}
        _ = v39
    case "\n\t":
        if false {
            v20 -= v20
            switch "" {
            case "foo bar":
                var v40 int8 = -51
                _ = v40
            case "":
                for fi41 := 0; fi41 < 18; fi41++ {
                    _ = len(v1)
                    _ = fi41
                }
            case "hello":
                var v42 int = -9223372036854775808
                _ = v42
            }
        }
    case "a":
        v20 = (v20 * 1e10)
    }
    if false {
        if false {
            if (4294967295 == 689) {
                for fi43 := 0; fi43 < 20; fi43++ {
                    for fi44 := 0; fi44 < 15; fi44++ {
                        v45 := GenS6{F0: true, F1: (false && (4294967295 == 0)), F2: 42}
                        v12 = v12
                        v12 = (v12 * v12)
                        _ = fi44
                        _ = v45
                    }
                    _ = fi43
                }
                v46 := GenS7{F0: "hello", F1: 1, F2: 42}
                if (1 >= -32768) {
                    var v47 float64 = (v12 + float64(float32(v12)))
                    v2[0] = 32768
                    _ = v47
                }
                _ = v46
            }
            fmt.Println(v20)
            var v48 int32 = -1
            _ = v48
        } else {
            fmt.Println(v20, v20)
        }
        v49 := GenS7{F0: "", F1: 42, F2: 18446744073709551615}
        v20 = 1.5
        _ = v49
    } else {
        _ = len(v11)
    }
    if len(v11) >= 2 {
        v11 = v11[1:len(v11)]
    }
    var v50 float32 = v20
    if false {
        v12 -= (v12 + 3.14159)
        var v51 uint8 = 2
        v50 = (v50 - ((v50 - 1e10) * v20))
        _ = v51
    } else {
        var v52 uint64 = 0
        switch 2147483647 {
        case -2147483648:
            switch 128 {
            case 0:
                switch 513 {
                case -9223372036854775808:
                    v20 = v50
                case 618:
                    switch -356 {
                    case 700:
                        v50 = v20
                    case 0:
                        v50 = 3.14159
                    default:
                        var v53 uint64 = (v52 / ((v52) | 1))
                        _ = v53
                    }
                case 9223372036854775807:
                    var v54 uint32 = 2147483647
                    _ = v54
                default:
                    var v55 float32 = (v50 - 0.1)
                    _ = v55
                }
            case 669:
                v56 := GenS5{F0: 825, F1: v52}
                _ = v56
            case 2147483647:
                switch 32767 {
                case 516:
                    if true {
                        if true {
                            v57 := GenS6{F0: false, F1: (127 > 127), F2: -9223372036854775808}
                            _ = v57
                        }
                        v50 += float32((0.1 - (((((v12 - 3.14159) / 1e10) / (v12 * 0.1)) - v12) + -2.25)))
                        v58 := GenS8{F0: -1}
                        _ = v58
                    }
                case 440:
                    v2[32767] = 4294967295
                }
            default:
                switch 9223372036854775807 {
                case 128:
                    v20 = v20
                }
            }
        default:
            for fi59 := 0; fi59 < 10; fi59++ {
                v20 -= 0.1
                var v60 int64 = -312
                _ = fi59
                _ = v60
            }
        }
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v50
}
