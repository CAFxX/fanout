package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS6 struct { F0 int; F1 int8 }
type GenS7 struct { F0 uint8; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS9 struct { F0 int; F1 int8 }
type GenS10 struct { F0 uint8; F1 int64; F2 uint }
type GenS11 struct { F0 string; F1 int32 }
type GenS12 struct { F0 bool; F1 int8 }
type GenS13 struct { F0 int }
func genFn0() bool {
    for fi5 := 0; fi5 < 18; fi5++ {
        for fi6 := 0; fi6 < 3; fi6++ {
            v7 := []int{}
            var v8 uint32 = 2147483647
            _ = fi6
            _ = v7
            _ = v8
        }
        if (false || false) {
            if true {
                var v9 string = (("" + "") + ("\n\t" + "\n\t"))
                v9 = ((v9 + (v9 + (v9 + v9))) + v9)
                for fi10 := 0; fi10 < 16; fi10++ {
                    switch (v9 + v9) {
                    case "hello":
                        v11 := GenS1{F0: (fi5 & (fi5 << ((42) & 63)))}
                        _ = v11
                    case "\n\t":
                        v12 := GenS2{F0: 310}
                        _ = v12
                    }
                    v9 = v9
                    _ = fi10
                }
                _ = v9
            }
            var v13 uint = 716
            v13 += v13
            _ = v13
        } else {
            var v14 uint16 = 32767
            var v15 bool = ((988 == 0) && (!false))
            v16 := GenS2{F0: 2}
            _ = v14
            _ = v15
            _ = v16
        }
        var v17 int16 = 32767
        _ = fi5
        _ = v17
    }
    var v18 uint32 = 0
    v19 := GenS3{F0: (v18 + v18), F1: 173}
    v20 := GenS4{F0: 717}
    _ = v18
    _ = v19
    _ = v20
    return true
}
func genFn1(p0 uint8, p1 int16, p2 string) uint {
    p2 = ((p2 + p2) + "foo bar")
    switch p2 {
    case "a":
        v21 := []int{-431, 224}
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return 255
}

func main() {
    var v1 int16 = 1
    var v2 int16 = 1000
    var v3 int8 = -99
    v4 := []int{4294967295}
    v3 -= v3
    for fi22 := 0; fi22 < 10; fi22++ {
        v2 = 0
        v1 = (v2 << ((1) & 63))
        for fi23 := 0; fi23 < 2; fi23++ {
            v1 += -32768
            v2 += v1
            v1 = (v1 - v2)
            _ = fi23
        }
        _ = fi22
    }
    v1 -= 651
    var v24 bool = (("\n\t" + ("a" + ("\n\t" + "\n\t"))) >= (("a" + "") + "a"))
    var v25 string = ""
    fmt.Println(v1)
    v26 := GenS5{F0: 788, F1: 2, F2: 1}
    if v24 {
        var v27 float32 = ((((-2.25 - float32((((((0.0 * ((float64(1.5) - (1e10 / float64(1e10))) * ((0.1 / (float64(3.14159) + 0.0)) + 3.14159))) / (3.14159 / (1.5 / 3.14159))) - 3.14159) / 1.5) * 0.0))) + (0.0 / float32((((-2.25 - 0.1) - 0.1) + 3.14159)))) * 1e10) + ((float32(0.1) - ((3.14159 / 1.0) / 3.14159)) * -2.25))
        v25 = ("" + "\n\t")
        _ = v27
    } else {
        v28 := GenS6{F0: -238, F1: v3}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v28
    }
    for fi29 := 0; fi29 < 11; fi29++ {
        if ((v1 != -1) && (18446744073709551615 != 42)) {
            switch v25 {
            case "foo bar":
                switch v25 {
                case "foo bar":
                    v30 := GenS6{F0: (fi29 / ((-533) | 1)), F1: (v3 / ((85) | 1))}
                    _ = v30
                case "":
                    fmt.Println(v25, v3, v25)
                default:
                    switch 680 {
                    case 32767:
                        var v31 float64 = 3.14159
                        _ = v31
                    case -49:
                        if (((v3 | (v3 ^ (v3 & (v3 * 0)))) <= 1) && false) {
                            if v24 {
                                v2 = -1
                                if (!v24) {
                                    v32 := GenS7{F0: 127, F1: (v25 < v25), F2: 520}
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    for fi33 := 0; fi33 < 6; fi33++ {
                                        v25 = v25
                                        _ = fi33
                                    }
                                    _ = v32
                                }
                            } else {
                                switch fi29 {
                                case 32767:
                                    var v34 uint64 = 733
                                    _ = v34
                                case 127:
                                    switch (fi29 % (((^(fi29))) | 1)) {
                                    case 127:
                                        for fi35 := 0; fi35 < 12; fi35++ {
                                            var v36 string = v25
                                            v37 := GenS8{F0: 2147483648, F1: -1, F2: 255}
                                            var v38 float64 = (0.1 + 3.14159)
                                            switch fi35 {
                                            case -2147483648:
                                                v36 = (v25 + v36)
                                            case -633:
                                                v36 = "hello"
                                            case 32767:
                                                switch v36 {
                                                case "foo bar":
                                                    fmt.Println(v2, v24, v3)
                                                }
                                            }
                                            var v39 uint = 548
                                            _ = fi35
                                            _ = v36
                                            _ = v37
                                            _ = v38
                                            _ = v39
                                        }
                                    }
                                case 42:
                                    _ = len(v4)
                                default:
                                    if v24 {
                                        v40 := GenS5{F0: 65535, F1: 255, F2: 2}
                                        switch (fi29 * (fi29 / (((-(fi29))) | 1))) {
                                        case -2147483648:
                                            v41 := GenS9{F0: 1, F1: (v3 - 1)}
                                            _ = v41
                                        case 32768:
                                            v42 := GenS10{F0: 127, F1: 0, F2: 125}
                                            _ = v42
                                        case -1:
                                            if true {
                                                var v43 int32 = 1
                                                _ = v43
                                            }
                                        }
                                        var v44 uint8 = 0
                                        _ = v40
                                        _ = v44
                                    }
                                }
                                v24 = (65535 != 173)
                                for fi45 := 0; fi45 < 12; fi45++ {
                                    v3 += v3
                                    v4 = append(v4, 9223372036854775807)
                                    var v46 int16 = v1
                                    _ = fi45
                                    _ = v46
                                }
                            }
                            v47 := GenS11{F0: (v25 + "a"), F1: 1}
                            _ = v47
                        } else {
                            var v48 float64 = (((((float64(-2.25) - (0.1 / ((1e10 / (0.1 * float64(1.5))) - 0.0))) + 0.1) * (float64(1.5) - 3.14159)) - (0.1 - (0.1 + -2.25))) * (((float64(float32(((1.5 * (-2.25 - 1e10)) + (1e10 - (1.5 - 1e10))))) + (1.5 * (((3.14159 + -2.25) / 1.0) - -2.25))) + float64(3.14159)) * 0.1))
                            _ = v48
                        }
                    default:
                        for fi49 := 0; fi49 < 17; fi49++ {
                            v50 := GenS5{F0: 32767, F1: 643, F2: 0}
                            var v51 string = v25
                            v25 = (v51 + (v51 + v25))
                            if false {
                                v1 = (-(v1))
                            } else {
                                if (v24 && v24) {
                                    if len(v4) > 0 {
                                        v4[0] = (fi49 + (fi49 % ((-569) | 1)))
                                    }
                                    fmt.Println(v1, fi29)
                                    v52 := GenS12{F0: true, F1: v3}
                                    _ = v52
                                }
                            }
                            _ = fi49
                            _ = v50
                            _ = v51
                        }
                    }
                }
            case "a":
                var v53 string = "hello"
                _ = v53
            }
        } else {
            for fi54 := 0; fi54 < 15; fi54++ {
                switch (fi54 | (fi54 >> ((991) & 63))) {
                case 32767:
                    if len(v4) > 0 {
                        v4[0] = (fi29 ^ int(v2))
                    }
                case 9223372036854775807:
                    v3 -= 0
                case 130:
                    var v55 float32 = (1e10 - (0.1 - (((1.5 * -2.25) * 0.1) / 1e10)))
                    v56 := GenS8{F0: 65535, F1: (-(v2)), F2: 65535}
                    v24 = (v24 || (((-666 == -927) || true) && (408 >= 930)))
                    fmt.Println(v24, v1)
                    _ = v55
                    _ = v56
                }
                _ = fi54
            }
        }
        v3 = v3
        v1 += 0
        _ = fi29
    }
    switch "foo bar" {
    case "a":
        switch v25 {
        case "\n\t":
            _ = len(v4)
        case "":
            fmt.Println(v25)
        case "foo bar":
            v2 = (v1 % ((42) | 1))
        default:
            var v57 uint32 = 941
            _ = v57
        }
    }
    v58 := GenS1{F0: 32767}
    for fi59 := 0; fi59 < 3; fi59++ {
        v60 := GenS4{F0: 881}
        if len(v4) > 0 {
            v4[0] = (-(fi59))
        }
        if len(v4) > 0 {
            v4[0] = fi59
        }
        _ = fi59
        _ = v60
    }
    v61 := GenS13{F0: 4294967295}
    var v62 uint8 = 126
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v58
    _ = v61
    _ = v62
}
