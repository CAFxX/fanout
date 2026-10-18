package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string; F1 uint16 }
type GenS4 struct { F0 uint16; F1 int32 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 uint64 }
func genFn0() uint {
    v4 := GenS1{F0: 447, F1: 209, F2: 100}
    switch 42 {
    case -2147483648:
        v5 := GenS1{F0: 1, F1: 255, F2: -31}
        _ = v5
    case -80:
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 string = "hello"
            v8 := []int{fi6, (fi6 / ((fi6) | 1)), (fi6 | (fi6 * (fi6 + (^(fi6)))))}
            var v9 uint32 = 903
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    case 32767:
        var v10 int16 = 0
        _ = v10
    }
    var v11 uint8 = 255
    fmt.Println(v11)
    _ = v11
    _ = v4
    return 42
}
func genFn1() bool {
    v12 := []int{-794, -9223372036854775808, 955}
    var v13 uint8 = 25
    _ = v12
    _ = v13
    return false
}

func main() {
    var v1 uint8 = 127
    var v2 string = "foo bar"
    var v3 string = ((v2 + (v2 + v2)) + (("hello" + ((v2 + "") + "hello")) + ("\n\t" + v2)))
    v1 += (v1 - v1)
    v14 := GenS1{F0: 4294967295, F1: v1, F2: 100}
    var v15 string = (((("a" + "") + v3) + (v2 + (("hello" + "a") + v2))) + ("foo bar" + v2))
    v2 = v15
    if (true || (0 != 65535)) {
        var v16 uint8 = 127
        var v17 int8 = 0
        if false {
            v15 = v3
            if (-1 < (v17 + -128)) {
                var v18 uint16 = 255
                _ = v18
            }
        }
        _ = v16
        _ = v17
    } else {
        v19 := GenS2{F0: 42, F1: (v1 % ((v1) | 1))}
        for fi20 := 0; fi20 < 7; fi20++ {
            v21 := GenS1{F0: 1, F1: v1, F2: -1}
            for fi22 := 0; fi22 < 2; fi22++ {
                v23 := GenS1{F0: 224, F1: 127, F2: 30}
                if (64 == 42) {
                    v15 = v3
                    if false {
                        v15 = v15
                        if true {
                            var v24 uint32 = 914
                            _ = v24
                        }
                    }
                    var v25 int64 = -305
                    _ = v25
                } else {
                    var v26 uint8 = v1
                    if (true && (v2 != "")) {
                        v26 = v26
                    } else {
                        v27 := GenS3{F0: "\n\t", F1: 42}
                        _ = v27
                    }
                    _ = v26
                }
                fmt.Println(v1)
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    for fi28 := 0; fi28 < 2; fi28++ {
        fmt.Println(v2, v1, v1)
        if true {
            var v29 int64 = 9223372036854775807
            var v30 float64 = float64((1.5 * float32((((1e10 * 1.5) + (1.5 - float64(0.1))) * 0.0))))
            _ = v29
            _ = v30
        }
        _ = fi28
    }
    fmt.Println(v15, v2, v2)
    fmt.Println(v15)
    if (3.14159 == 1e10) {
        switch (("foo bar" + "") + "\n\t") {
        case "hello":
            if (65535 < 2) {
                for fi31 := 0; fi31 < 3; fi31++ {
                    v15 = "hello"
                    v2 = v3
                    var v32 string = "a"
                    _ = fi31
                    _ = v32
                }
                v33 := []int{42, -849}
                _ = v33
            } else {
                for fi34 := 0; fi34 < 4; fi34++ {
                    v35 := []int{(fi34 / ((fi34) | 1))}
                    if (((-9223372036854775808 != 0) && (!(!false))) || (707 < 42)) {
                        switch 9223372036854775807 {
                        case 1:
                            v36 := GenS4{F0: 873, F1: 1}
                            _ = v36
                        case 4294967295:
                            v1 = v1
                        case 32768:
                            v37 := GenS2{F0: genFn0(), F1: 0}
                            _ = v37
                        }
                    }
                    v38 := GenS4{F0: 255, F1: 2147483647}
                    _ = fi34
                    _ = v35
                    _ = v38
                }
            }
        case "\n\t":
            var v39 int8 = -128
            _ = v39
        default:
            v40 := GenS3{F0: (v2 + "a"), F1: 42}
            _ = v40
        }
        for fi41 := 0; fi41 < 10; fi41++ {
            if false {
                v42 := GenS1{F0: 42, F1: 52, F2: 42}
                fmt.Println(fi41, v2, v1)
                _ = v42
            }
            _ = fi41
        }
    }
    v1 += 2
    var v43 int16 = 0
    switch 762 {
    case -2147483648:
        v1 = v1
    case -494:
        var v44 float32 = 3.14159
        _ = v44
    case 128:
        v45 := GenS2{F0: 42, F1: v1}
        _ = v45
    default:
        v43 -= v43
    }
    switch 128 {
    case 32767:
        fmt.Println(v15, v3, v2)
    case 489:
        var v46 uint = 65535
        _ = v46
    case 0:
        switch "\n\t" {
        case "\n\t":
            if (-1 >= -1) {
                v47 := GenS5{F0: 9223372036854775807, F1: -45, F2: 2}
                var v48 int32 = -1
                _ = v47
                _ = v48
            } else {
                v43 -= (v43 - (-(v43)))
                v49 := GenS6{F0: 18446744073709551615}
                var v50 int16 = v43
                _ = v49
                _ = v50
            }
        case "hello":
            var v51 int32 = 42
            _ = v51
        case "":
            v52 := []int{4294967295}
            _ = v52
        }
    default:
        if (100 <= 2) {
            switch 127 {
            case -225:
                v43 += (v43 ^ 1000)
            case 4294967295:
                v53 := []int{0, 127, -1}
                _ = v53
            case -223:
                if (2147483647 > 983) {
                    for fi54 := 0; fi54 < 7; fi54++ {
                        switch fi54 {
                        case 383:
                            for fi55 := 0; fi55 < 13; fi55++ {
                                v56 := []int{fi55, fi54, fi54}
                                var v57 int = -2147483648
                                var v58 uint16 = 42
                                _ = fi55
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                        case 2147483647:
                            v59 := []int{-2147483648, 42, fi54, fi54}
                            _ = v59
                        default:
                            fmt.Println(v1)
                        }
                        v1 = 111
                        for fi60 := 0; fi60 < 1; fi60++ {
                            fmt.Println(fi54)
                            v3 = ""
                            switch (fi54 ^ (-(fi54))) {
                            case 9223372036854775807:
                                var v61 float32 = 0.0
                                _ = v61
                            case -9223372036854775808:
                                if false {
                                    v15 = v15
                                }
                                v1 = v1
                            default:
                                var v62 int32 = -1
                                _ = v62
                            }
                            _ = fi60
                        }
                        _ = fi54
                    }
                    v43 = -1
                    for fi63 := 0; fi63 < 15; fi63++ {
                        v3 = v15
                        _ = fi63
                    }
                } else {
                    if ((!((!((1.5 + 1.5) >= (-2.25 - 3.14159))) && true)) && (1e10 != 1e10)) {
                        var v64 bool = true
                        _ = v64
                    } else {
                        for fi65 := 0; fi65 < 16; fi65++ {
                            if true {
                                for fi66 := 0; fi66 < 3; fi66++ {
                                    var v67 uint8 = 42
                                    _ = fi66
                                    _ = v67
                                }
                                v68 := []int{(fi65 - (fi65 ^ -1))}
                                v1 += (v1 % (((v1 * (v1 ^ 42))) | 1))
                                _ = v68
                            } else {
                                v69 := []int{fi65, fi65}
                                _ = v69
                            }
                            _ = fi65
                        }
                    }
                }
            default:
                fmt.Println(v2, v3)
            }
            v43 += (v43 / ((v43) | 1))
            if (896 <= 270) {
                v70 := GenS6{F0: 18446744073709551615}
                fmt.Println(v1, v43)
                for fi71 := 0; fi71 < 18; fi71++ {
                    if ((fi71 >> ((593) & 63)) <= fi71) {
                        v72 := []int{4294967295}
                        var v73 string = (v3 + ((("foo bar" + (v2 + v15)) + ("a" + v15)) + (v3 + (v3 + v3))))
                        _ = v72
                        _ = v73
                    }
                    _ = fi71
                }
                _ = v70
            }
        }
    }
    v74 := []int{-23}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v74:", len(v74), v74)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v43
    _ = v74
}
