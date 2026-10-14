package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 bool; F2 uint }
type GenS3 struct { F0 int; F1 int8 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint32; F1 int64 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint16; F1 int16 }
type GenS8 struct { F0 uint16; F1 int8 }
type GenS9 struct { F0 int8 }
type GenS10 struct { F0 int32; F1 uint8; F2 uint16 }
type GenS11 struct { F0 string; F1 int16; F2 int8 }
type GenS12 struct { F0 uint64 }
func genFn0() uint64 {
    var v5 uint16 = 65535
    if (4294967295 >= 18446744073709551615) {
        if false {
            for fi6 := 0; fi6 < 1; fi6++ {
                switch fi6 {
                case 0:
                    v5 = (v5 << ((2) & 63))
                }
                for fi7 := 0; fi7 < 19; fi7++ {
                    v5 -= (v5 ^ (v5 + (v5 | 42)))
                    if true {
                        switch "\n\t" {
                        case "a":
                            v8 := GenS1{F0: -698}
                            _ = v8
                        case "\n\t":
                            var v9 float64 = (((float64(0.1) * (0.1 / ((3.14159 - (0.0 - float64(0.1))) - (1.5 * (0.1 * (float64(1.5) + 3.14159)))))) + ((1.5 * 3.14159) - 1e10)) / float64((0.1 / ((((1e10 + 1e10) + -2.25) * ((float32(0.1) / (float32(((3.14159 * 1e10) * (0.0 + 0.1))) + float32((0.1 - 1.5)))) - 0.0)) / (0.1 * 1e10)))))
                            _ = v9
                        }
                        v5 = 65535
                        v10 := []int{}
                        _ = v10
                    }
                    _ = fi7
                }
                v11 := GenS2{F0: 0, F1: ((fi6 * (fi6 + (fi6 % ((fi6) | 1)))) != (fi6 * (fi6 - -2147483648))), F2: 42}
                _ = fi6
                _ = v11
            }
        } else {
            for fi12 := 0; fi12 < 19; fi12++ {
                v5 = 255
                _ = fi12
            }
            if false {
                v13 := GenS3{F0: 127, F1: 42}
                if true {
                    v14 := GenS4{F0: "foo bar", F1: 0}
                    switch "a" {
                    case "\n\t":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            var v16 int32 = -2147483648
                            switch (fi15 + 0) {
                            case -637:
                                if true {
                                    if true {
                                        v17 := GenS3{F0: 2147483647, F1: 39}
                                        _ = v17
                                    } else {
                                        if (0 > 100) {
                                            var v18 int8 = 42
                                            var v19 int16 = -32768
                                            v20 := GenS4{F0: ("a" + "foo bar"), F1: 2147483648}
                                            _ = v18
                                            _ = v19
                                            _ = v20
                                        } else {
                                            v21 := []int{(fi15 >> ((26) & 63))}
                                            v21 = append(v21, fi15)
                                            _ = v21
                                        }
                                        if (!((fi15 << ((42) & 63)) < fi15)) {
                                            fmt.Println(v5)
                                        }
                                    }
                                    fmt.Println(v16)
                                }
                            }
                            v5 = v5
                            _ = fi15
                            _ = v16
                        }
                    case "foo bar":
                        var v22 int16 = 0
                        _ = v22
                    case "hello":
                        fmt.Println(v5)
                    default:
                        fmt.Println(v5)
                    }
                    _ = v14
                } else {
                    v5 -= v5
                }
                _ = v13
            } else {
                v5 += 1
            }
            v5 += (v5 % ((v5) | 1))
        }
        v23 := GenS5{F0: 239, F1: -333}
        _ = v23
    } else {
        for fi24 := 0; fi24 < 6; fi24++ {
            v5 -= 158
            if false {
                v25 := GenS6{F0: 9223372036854775807}
                _ = v25
            }
            _ = fi24
        }
        var v26 int32 = int32(v5)
        _ = v26
    }
    switch ("" + ("hello" + ("hello" + "hello"))) {
    case "foo bar":
        v5 -= (v5 / ((v5) | 1))
    default:
        v27 := []int{0}
        _ = v27
    }
    _ = v5
    return 386
}

func main() {
    v1 := []int{0}
    var v2 uint64 = 4294967295
    v3 := []string{"hello"}
    var v4 int16 = 42
    if false {
        fmt.Println(v2)
        var v28 float32 = 0.1
        _ = v28
    } else {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v29 uint32 = 729
    var v30 int64 = 9223372036854775807
    switch -606 {
    case 42:
        var v31 int64 = 1
        _ = v31
    case 4294967295:
        for fi32 := 0; fi32 < 10; fi32++ {
            v29 = 213
            if true {
                v33 := GenS7{F0: 335, F1: (^(v4))}
                if false {
                    for fi34 := 0; fi34 < 10; fi34++ {
                        var v35 int = fi32
                        v4 = (v4 | 32767)
                        _ = fi34
                        _ = v35
                    }
                } else {
                    switch fi32 {
                    case 32767:
                        for fi36 := 0; fi36 < 13; fi36++ {
                            v37 := GenS8{F0: 726, F1: -1}
                            v29 = 187
                            _ = fi36
                            _ = v37
                        }
                    case 42:
                        for fi38 := 0; fi38 < 13; fi38++ {
                            var v39 int32 = 1
                            var v40 int16 = (^(v4))
                            _ = len(v3)
                            fmt.Println(v40, v30)
                            if len(v3) > 0 {
                                v3[0] = ("hello" + ("a" + (("\n\t" + ("\n\t" + "a")) + "hello")))
                            }
                            var v41 float32 = 1e10
                            if true {
                                var v42 int32 = (v39 * v39)
                                if true {
                                    if false {
                                        v43 := GenS9{F0: 42}
                                        var v44 uint64 = v2
                                        _ = v43
                                        _ = v44
                                    }
                                }
                                _ = v42
                            } else {
                                var v45 uint16 = 652
                                if ((v4 * (v4 * (v4 * v4))) < 32767) {
                                    if len(v3) > 0 {
                                        v3[0] = "\n\t"
                                    }
                                }
                                _ = v45
                            }
                            if true {
                                v29 -= v29
                            }
                            _ = fi38
                            _ = v39
                            _ = v40
                            _ = v41
                        }
                    case -2147483648:
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v1 = append(v1, (fi32 % ((fi32) | 1)))
                            if true {
                                v47 := GenS10{F0: 653, F1: 0, F2: 0}
                                var v48 uint64 = v2
                                v48 -= v48
                                _ = v47
                                _ = v48
                            } else {
                                v4 = -965
                                var v49 uint8 = 127
                                _ = v49
                            }
                            _ = fi46
                        }
                    default:
                        v50 := GenS11{F0: ("hello" + "hello"), F1: v4, F2: -53}
                        _ = v50
                    }
                    if false {
                        var v51 string = ("a" + "a")
                        _ = v51
                    }
                }
                _ = v33
            } else {
                var v52 int8 = -49
                _ = v52
            }
            _ = fi32
        }
    }
    switch -645 {
    case 2147483647:
        v29 -= 2147483648
    default:
        v53 := GenS11{F0: "hello", F1: 1, F2: 54}
        _ = v53
    }
    switch 0 {
    case 42:
        _ = len(v1)
    case 0:
        if true {
            for fi54 := 0; fi54 < 13; fi54++ {
                if (v30 <= v30) {
                    v55 := GenS11{F0: "a", F1: v4, F2: 1}
                    fmt.Println(v2, v30)
                    _ = len(v1)
                    _ = v55
                }
                var v56 float64 = 0.0
                for fi57 := 0; fi57 < 9; fi57++ {
                    if false {
                        fmt.Println(v29)
                        fmt.Println(fi54)
                        fmt.Println(v2, v29, fi54)
                    } else {
                        switch (^(fi54)) {
                        case 2147483647:
                            v56 = (1e10 + 1.5)
                        default:
                            v29 = v29
                        }
                    }
                    v29 = 994
                    var v58 string = "hello"
                    var v59 bool = (int8(v4) != -128)
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                _ = fi54
                _ = v56
            }
            switch "\n\t" {
            case "a":
                if (!false) {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                } else {
                    if len(v3) > 0 {
                        v3[0] = "hello"
                    }
                    var v60 bool = false
                    _ = v60
                }
            case "\n\t":
                var v61 uint64 = 955
                _ = v61
            case "foo bar":
                _ = len(v1)
            }
            _ = len(v3)
        }
    case 340:
        for fi62 := 0; fi62 < 20; fi62++ {
            if ((v4 ^ v4) == v4) {
                for fi63 := 0; fi63 < 7; fi63++ {
                    v64 := GenS12{F0: genFn0()}
                    v29 = (v29 & 218)
                    _ = fi63
                    _ = v64
                }
                var v65 float64 = 3.14159
                _ = v65
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            switch fi62 {
            case -868:
                for fi66 := 0; fi66 < 14; fi66++ {
                    v29 = v29
                    _ = fi66
                }
            case 4294967295:
                var v67 int32 = -2147483648
                _ = v67
            }
            var v68 int = -22
            _ = fi62
            _ = v68
        }
    }
    switch (("a" + "") + "a") {
    case "hello":
        var v69 uint64 = v2
        _ = v69
    case "foo bar":
        v2 -= 2
    }
    var v70 uint16 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v70:", v70)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v70
}
