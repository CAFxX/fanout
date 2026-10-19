package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 int16; F1 uint64; F2 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int16 }
type GenS6 struct { F0 uint; F1 int64; F2 int32 }
type GenS7 struct { F0 string; F1 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 string }
type GenS10 struct { F0 uint8; F1 uint; F2 int }
type GenS11 struct { F0 uint8; F1 uint }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 int32; F1 uint16 }
func genFn0() int16 {
    var v6 int = 0
    v6 = (v6 % ((42) | 1))
    var v7 int32 = 42
    for fi8 := 0; fi8 < 20; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 string = ""
            v11 := []int{(-(fi9))}
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    _ = v6
    _ = v7
    return 465
}

func main() {
    v1 := []uint32{593}
    var v2 uint16 = 32767
    var v3 uint32 = 2147483647
    v4 := []int{}
    v5 := []uint32{(v3 << ((0) & 63))}
    switch ("\n\t" + "\n\t") {
    case "hello":
        if len(v1) > 0 {
            v1[0] = v3
        }
    default:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v3 = 65535
    var v12 int32 = 1
    var v13 uint16 = (v2 >> ((0) & 63))
    fmt.Println(v2)
    var v14 float64 = ((1.5 * 3.14159) * 0.1)
    v3 -= v3
    if (4294967295 >= 382) {
        var v15 int8 = -128
        v16 := GenS1{F0: v3, F1: "hello", F2: v12}
        v17 := GenS2{F0: 32767, F1: 0, F2: -1}
        _ = v15
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        fmt.Println(v14)
        for fi19 := 0; fi19 < 12; fi19++ {
            var v20 uint8 = 1
            v21 := GenS2{F0: 42, F1: 589, F2: 42}
            v13 += (v2 / (((^(v2))) | 1))
            _ = fi19
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 12; fi22++ {
            if false {
                var v23 int64 = 9223372036854775807
                var v24 int32 = (v12 * 0)
                _ = v23
                _ = v24
            }
            _ = fi22
        }
        _ = fi18
    }
    switch ("a" + "a") {
    case "foo bar":
        for fi25 := 0; fi25 < 9; fi25++ {
            v26 := GenS1{F0: (v3 % ((354) | 1)), F1: "a", F2: (v12 % ((v12) | 1))}
            switch int(v13) {
            case -2147483648:
                var v27 int64 = 560
                _ = v27
            case 1:
                if false {
                    var v28 uint32 = (v3 & v3)
                    if true {
                        switch fi25 {
                        case -1:
                            if len(v4) > 0 {
                                v4[0] = fi25
                            }
                        case 4294967295:
                            var v29 float32 = (-2.25 + (1.5 - (1e10 / 3.14159)))
                            _ = v29
                        default:
                            var v30 int16 = -302
                            _ = v30
                        }
                        switch (fi25 & 32767) {
                        case 537:
                            for fi31 := 0; fi31 < 1; fi31++ {
                                if len(v1) > 0 {
                                    v1[0] = (v28 ^ 2147483648)
                                }
                                _ = fi31
                            }
                        case 9223372036854775807:
                            v28 -= v28
                        case 1:
                            for fi32 := 0; fi32 < 6; fi32++ {
                                v1 = append(v1, v28)
                                switch ("\n\t" + "a") {
                                case "\n\t":
                                    v3 -= 65535
                                }
                                if false {
                                    switch "\n\t" {
                                    case "hello":
                                        v28 += (v3 << ((4294967295) & 63))
                                    }
                                } else {
                                    var v33 int64 = 1
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if true {
                                        if false {
                                            var v34 int8 = 125
                                            v2 = 0
                                            v3 = v3
                                            if (1000 >= -1) {
                                                if len(v1) > 0 {
                                                    v1[0] = 42
                                                }
                                                v35 := GenS3{F0: -971}
                                                _ = v35
                                            } else {
                                                var v36 uint64 = 2
                                                if ((fi32 << ((1) & 63)) <= fi25) {
                                                    switch ("foo bar" + ("a" + "hello")) {
                                                    case "":
                                                        var v37 uint = 4294967295
                                                        _ = v37
                                                    default:
                                                        v3 += uint32(v36)
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                } else {
                                                    var v38 int8 = (-(v34))
                                                    var v39 bool = ((4294967295 == 893) && (!false))
                                                    _ = v38
                                                    _ = v39
                                                }
                                                if (false || true) {
                                                    switch ("foo bar" + ("foo bar" + "foo bar")) {
                                                    case "a":
                                                        v40 := GenS4{F0: (fi25 | (^(fi25))), F1: 301, F2: v13}
                                                        _ = v40
                                                    }
                                                    v2 = 394
                                                } else {
                                                    _ = len(v5)
                                                    v14 = (v14 - v14)
                                                    v41 := GenS5{F0: (v13 | v2), F1: -1}
                                                    _ = v41
                                                }
                                                _ = v36
                                            }
                                            if true {
                                                if (true || false) {
                                                    v2 -= 1
                                                    var v42 int32 = -90
                                                    if (false || (!false)) {
                                                        v43 := GenS2{F0: -32768, F1: 976, F2: (v33 + (v33 | v33))}
                                                        fmt.Println(v3, v13)
                                                        _ = v43
                                                    }
                                                    _ = v42
                                                } else {
                                                    v28 = 1
                                                    v44 := GenS6{F0: 18446744073709551615, F1: v33, F2: v12}
                                                    _ = v44
                                                }
                                                switch fi25 {
                                                case -2147483648:
                                                    switch fi25 {
                                                    case 42:
                                                        v45 := GenS2{F0: -889, F1: 1, F2: 0}
                                                        _ = v45
                                                    }
                                                default:
                                                    if false {
                                                        var v46 uint = 2
                                                        _ = v46
                                                    }
                                                }
                                                if (((v14 + float64(-2.25)) / 1.5) > v14) {
                                                    v47 := GenS7{F0: ((("" + ("\n\t" + "a")) + ((("\n\t" + ("" + (("foo bar" + ("foo bar" + (("hello" + "hello") + (("hello" + "hello") + "")))) + ("foo bar" + "\n\t")))) + "\n\t") + "\n\t")) + ""), F1: fi32}
                                                    v3 += v3
                                                    var v48 int = fi32
                                                    _ = v47
                                                    _ = v48
                                                }
                                            } else {
                                                var v49 int32 = v12
                                                _ = v49
                                            }
                                            _ = v34
                                        }
                                        if false {
                                            v12 -= v12
                                        }
                                    }
                                    _ = v33
                                }
                                v5 = append(v5, v28)
                                _ = fi32
                            }
                        default:
                            _ = len(v4)
                        }
                    } else {
                        switch (fi25 ^ (fi25 - (fi25 + fi25))) {
                        case 0:
                            v13 -= (v2 + (v2 + v2))
                        default:
                            for fi50 := 0; fi50 < 13; fi50++ {
                                var v51 uint32 = 2147483648
                                v2 = 901
                                var v52 int64 = 9223372036854775807
                                _ = fi50
                                _ = v51
                                _ = v52
                            }
                        }
                    }
                    var v53 uint64 = 293
                    _ = v28
                    _ = v53
                }
            case 4294967295:
                if len(v1) > 0 {
                    v1[0] = (v3 >> ((255) & 63))
                }
            default:
                for fi54 := 0; fi54 < 20; fi54++ {
                    var v55 int8 = 36
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    switch 0 {
                    case -786:
                        var v56 int = 2147483647
                        if len(v4) > 0 {
                            v4[0] = 32768
                        }
                        _ = v56
                    case 2147483647:
                        v57 := GenS8{F0: 0}
                        _ = v57
                    case -2147483648:
                        fmt.Println(v13)
                    }
                    v58 := GenS9{F0: "\n\t"}
                    switch fi25 {
                    case 127:
                        var v59 int8 = 42
                        _ = v59
                    case 32767:
                        v60 := GenS1{F0: v3, F1: "hello", F2: (v12 >> ((99) & 63))}
                        _ = v60
                    default:
                        v61 := GenS5{F0: v2, F1: genFn0()}
                        _ = v61
                    }
                    _ = fi54
                    _ = v55
                    _ = v58
                }
            }
            for fi62 := 0; fi62 < 12; fi62++ {
                _ = len(v4)
                v4 = append(v4, fi25)
                _ = fi62
            }
            _ = fi25
            _ = v26
        }
    case "hello":
        var v63 int16 = -1
        _ = v63
    case "\n\t":
        var v64 int = 9223372036854775807
        _ = v64
    default:
        var v65 float64 = float64(0.1)
        _ = v65
    }
    v13 += (v13 ^ v13)
    if (2 == uint64(v2)) {
        switch 9223372036854775807 {
        case 1:
            for fi66 := 0; fi66 < 13; fi66++ {
                v67 := GenS10{F0: 168, F1: 42, F2: fi66}
                var v68 uint8 = 42
                if true {
                    v69 := GenS1{F0: (v3 + 1), F1: "\n\t", F2: v12}
                    _ = v69
                }
                _ = fi66
                _ = v67
                _ = v68
            }
        case 2147483647:
            var v70 float32 = (-2.25 / 1.0)
            _ = v70
        case -2147483648:
            v14 += -2.25
        }
        var v71 int = 725
        _ = v71
    }
    v72 := GenS11{F0: 127, F1: 154}
    v73 := GenS12{F0: 42}
    v74 := GenS13{F0: -1, F1: 447}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v73
    _ = v74
}
