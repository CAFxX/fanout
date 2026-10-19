package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int64 }
type GenS2 struct { F0 int; F1 string }
type GenS3 struct { F0 int32; F1 bool }
type GenS4 struct { F0 int; F1 int8; F2 string }
type GenS5 struct { F0 uint32; F1 uint64 }
type GenS6 struct { F0 uint32; F1 int64 }
func genFn0(p0 bool, p1 uint64) bool {
    var v7 uint8 = 42
    v8 := GenS1{F0: -1, F1: "foo bar", F2: 516}
    switch "hello" {
    case "":
        fmt.Println(p1)
    default:
        v9 := GenS1{F0: 2147483647, F1: "\n\t", F2: 9223372036854775807}
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := []int{}
        switch fi10 {
        case 427:
            var v12 int = (fi10 % ((fi10) | 1))
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (387 == 261)
}
func genFn1(p0 bool, p1 int32) uint8 {
    v13 := GenS2{F0: 949, F1: (("" + "foo bar") + ("" + ("a" + "")))}
    for fi14 := 0; fi14 < 10; fi14++ {
        if true {
            var v15 uint32 = 4294967295
            v16 := []int{fi14, fi14, -1}
            p1 = p1
            _ = v15
            _ = v16
        }
        _ = fi14
    }
    if (p0 && ((!true) || false)) {
        var v17 int8 = 42
        v17 -= (v17 % ((v17) | 1))
        _ = v17
    } else {
        var v18 int64 = -1
        _ = v18
    }
    var v19 int64 = 42
    _ = p0
    _ = p1
    _ = v13
    _ = v19
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 255
    var v3 int32 = 165
    var v4 float64 = (0.1 * (0.1 + 1e10))
    var v5 float32 = ((1.5 + (1.5 + 0.1)) / 3.14159)
    var v6 uint64 = 4294967295
    v20 := GenS2{F0: 1, F1: "\n\t"}
    var v21 uint64 = v6
    v2 += v2
    var v22 int64 = int64(v21)
    var v23 string = ""
    v6 = v6
    switch -9223372036854775808 {
    case -9223372036854775808:
        v24 := GenS1{F0: v3, F1: "\n\t", F2: 9223372036854775807}
        _ = v24
    case 4294967295:
        if ((false && ((v3 >= (v3 % (((v3 << (((v2 ^ v2)) & 63))) | 1))) || false)) || (!true)) {
            if true {
                switch -9223372036854775808 {
                case 128:
                    v21 -= v6
                default:
                    for fi25 := 0; fi25 < 5; fi25++ {
                        v22 += v22
                        _ = fi25
                    }
                }
                switch 32768 {
                case 354:
                    fmt.Println(v5)
                case 126:
                    var v26 uint8 = 45
                    _ = v26
                case 9223372036854775807:
                    v22 += -138
                }
            }
            fmt.Println(v22)
            for fi27 := 0; fi27 < 9; fi27++ {
                var v28 int32 = v3
                switch v23 {
                case "\n\t":
                    for fi29 := 0; fi29 < 3; fi29++ {
                        if false {
                            v30 := GenS1{F0: 0, F1: ("a" + v23), F2: 1}
                            var v31 uint = v2
                            _ = v30
                            _ = v31
                        } else {
                            switch v23 {
                            case "foo bar":
                                v32 := GenS3{F0: (v28 & 382), F1: (v5 < 3.14159)}
                                _ = v32
                            case "\n\t":
                                v22 += (v22 % (((v22 ^ -9223372036854775808)) | 1))
                            case "hello":
                                v28 -= (v3 | v3)
                            default:
                                v4 += (v4 + float64(v5))
                            }
                            if (false || (127 <= genFn1((!true), (v3 / ((v3) | 1))))) {
                                var v33 float32 = ((v5 - ((3.14159 - float32(((((v4 - (-2.25 * -2.25)) * (v4 * (v4 - v4))) + float64(1.5)) + v4))) - v5)) + v5)
                                v34 := GenS4{F0: (^(fi27)), F1: 1, F2: "\n\t"}
                                v35 := GenS4{F0: 32768, F1: -1, F2: v23}
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                            var v36 float32 = v5
                            var v37 int32 = (v28 >> ((v2) & 63))
                            _ = v36
                            _ = v37
                        }
                        if true {
                            var v38 int8 = int8(fi29)
                            if false {
                                v39 := []int{(-(fi27))}
                                v28 += 42
                                v3 = -658
                                _ = v39
                            }
                            _ = v38
                        }
                        _ = fi29
                    }
                case "":
                    v40 := GenS5{F0: uint32(v2), F1: 153}
                    _ = v40
                }
                _ = fi27
                _ = v28
            }
        } else {
            v4 = (v4 + v4)
            v41 := v1[1]
            _ = v41
        }
    default:
        v42 := GenS3{F0: -400, F1: (804 < 32767)}
        _ = v42
    }
    fmt.Println(v23, v6, v2)
    v23 = "\n\t"
    var v43 int32 = -478
    switch "hello" {
    case "a":
        v1[2147483647] = -2147483648
    default:
        for fi44 := 0; fi44 < 15; fi44++ {
            if false {
                var v45 int = fi44
                if ((v45 << ((v2) & 63)) >= 4294967295) {
                    v3 = (v43 ^ 675)
                    if false {
                        v46 := GenS3{F0: v3, F1: (v23 != "hello")}
                        _ = v46
                    } else {
                        v47 := GenS6{F0: 764, F1: (v22 ^ (v22 & 0))}
                        if (!(v4 != (((((-2.25 / v4) / 1.5) - v4) + 1.5) * (v4 + (((((1.5 - (v4 * (v4 + v4))) + v4) * v4) * (((v4 - ((1.5 / (v4 + (-2.25 * 0.0))) - (1e10 + 1e10))) * (0.0 * 1e10)) / v4)) / (v4 / ((v4 / v4) + v4))))))) {
                            if ((v21 & (v6 * v21)) != (v21 % (((v21 % ((v21) | 1))) | 1))) {
                                v48 := []int{(fi44 | fi44)}
                                var v49 float32 = float32(v4)
                                _ = v48
                                _ = v49
                            }
                        } else {
                            v43 += v43
                            v1[0] = (-(fi44))
                        }
                        var v50 uint = 0
                        _ = v47
                        _ = v50
                    }
                } else {
                    var v51 float64 = ((v4 * (v4 - (float64(3.14159) / -2.25))) / 1e10)
                    var v52 uint8 = 42
                    _ = v51
                    _ = v52
                }
                _ = v45
            } else {
                for fi53 := 0; fi53 < 3; fi53++ {
                    v4 += 1e10
                    if false {
                        if ((1e10 <= v5) && true) {
                            v4 = (v4 * v4)
                        }
                    } else {
                        v4 += (((0.0 * 0.1) - 1e10) / 1.5)
                        v54 := GenS2{F0: fi44, F1: ((("" + (("a" + ("hello" + v23)) + ((((("hello" + v23) + "a") + (((((v23 + v23) + "") + ((v23 + (v23 + (((v23 + (v23 + ((v23 + v23) + v23))) + ("foo bar" + v23)) + v23))) + "foo bar")) + "hello") + "\n\t")) + v23) + (v23 + v23)))) + v23) + ("\n\t" + v23))}
                        fmt.Println(fi53)
                        _ = v54
                    }
                    _ = fi53
                }
            }
            fmt.Println(v22, v2)
            _ = fi44
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
