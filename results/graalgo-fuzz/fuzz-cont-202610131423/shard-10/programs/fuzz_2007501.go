package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int32; F1 uint; F2 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 bool; F2 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string; F1 uint32; F2 int }
func genFn0() uint64 {
    var v5 uint8 = 210
    v6 := GenS1{F0: 9223372036854775807, F1: 2}
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v5)
        v5 += 127
        _ = fi7
    }
    var v8 uint32 = 2147483647
    _ = v5
    _ = v6
    _ = v8
    return 1
}
func genFn1() uint {
    v9 := []int{}
    if false {
        if (65535 < 544) {
            var v10 string = ("hello" + "foo bar")
            _ = v10
        } else {
            var v11 uint8 = 237
            fmt.Println(v11)
            if false {
                var v12 uint16 = 65535
                _ = v12
            }
            _ = v11
        }
        var v13 int = 42
        _ = v13
    } else {
        for fi14 := 0; fi14 < 8; fi14++ {
            var v15 float32 = (1.5 - (0.0 + (0.0 / float32((((0.0 - 1.5) - (-2.25 * 0.1)) * float64(((0.1 * ((float32(0.0) + -2.25) * 3.14159)) / 3.14159)))))))
            _ = fi14
            _ = v15
        }
        var v16 uint64 = 4294967295
        v16 = v16
        _ = v16
    }
    v17 := GenS1{F0: 431, F1: 4294967295}
    v18 := GenS1{F0: 9223372036854775807, F1: 781}
    for fi19 := 0; fi19 < 18; fi19++ {
        var v20 uint8 = 255
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint32 = 706
            v23 := GenS1{F0: 4294967295, F1: 629}
            _ = fi21
            _ = v22
            _ = v23
        }
        _ = fi19
        _ = v20
    }
    _ = v17
    _ = v18
    _ = v9
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -128
    var v3 int64 = -9223372036854775808
    v4 := []uint32{65535}
    for fi24 := 0; fi24 < 19; fi24++ {
        v25 := GenS1{F0: fi24, F1: 760}
        var v26 int32 = -185
        var v27 uint16 = 187
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    v2 = v2
    for fi28 := 0; fi28 < 16; fi28++ {
        v2 -= v2
        fmt.Println(v2)
        if (3.14159 < -2.25) {
            v2 = (v2 + v2)
        }
        _ = fi28
    }
    v29 := GenS1{F0: 466, F1: 42}
    v3 = (v3 % ((v3) | 1))
    v2 -= (-(v2))
    for fi30 := 0; fi30 < 14; fi30++ {
        for fi31 := 0; fi31 < 15; fi31++ {
            fmt.Println(fi31, v3, v2)
            v3 -= (v3 << ((18446744073709551615) & 63))
            if (1 >= 2) {
                v32 := GenS1{F0: 0, F1: 417}
                _ = v32
            }
            _ = fi31
        }
        switch fi30 {
        case 568:
            v2 = (-(v2))
        case -2147483648:
            if false {
                v33 := GenS2{F0: 42, F1: 1, F2: (^(fi30))}
                v34 := v1[(fi30 + (fi30 | (fi30 & fi30)))]
                _ = v33
                _ = v34
            }
        case -848:
            v1[2147483647] = (fi30 + -322)
        }
        _ = fi30
    }
    switch "\n\t" {
    case "":
        for fi35 := 0; fi35 < 15; fi35++ {
            for fi36 := 0; fi36 < 15; fi36++ {
                if (false && (-1 >= 1000)) {
                    var v37 string = ("foo bar" + "\n\t")
                    v38 := GenS1{F0: fi35, F1: 42}
                    _ = v37
                    _ = v38
                } else {
                    fmt.Println(fi35)
                    if (("\n\t" != ((("a" + ("foo bar" + ("foo bar" + "hello"))) + "") + "")) && true) {
                        var v39 float32 = 3.14159
                        if false {
                            var v40 int64 = (^(v3))
                            switch "hello" {
                            case "a":
                                if len(v4) > 0 {
                                    v4[0] = 503
                                }
                            default:
                                var v41 bool = true
                                _ = v41
                            }
                            v42 := GenS3{F0: 0}
                            _ = v40
                            _ = v42
                        }
                        _ = v39
                    }
                    v43 := GenS2{F0: 42, F1: 784, F2: (fi35 | fi36)}
                    _ = v43
                }
                v44 := GenS4{F0: (v3 >> ((genFn1()) & 63)), F1: (42 == 1), F2: 776}
                if true {
                    fmt.Println(fi35, v3, v3)
                    v3 = v3
                } else {
                    var v45 bool = (!true)
                    _ = v45
                }
                _ = fi36
                _ = v44
            }
            _ = fi35
        }
    }
    v46 := GenS5{F0: (42 <= 1)}
    if true {
        switch "hello" {
        case "a":
            v2 -= v2
        case "\n\t":
            v47 := GenS5{F0: false}
            _ = v47
        }
    } else {
        v3 -= (v3 | (v3 << ((241) & 63)))
        for fi48 := 0; fi48 < 9; fi48++ {
            var v49 bool = true
            for fi50 := 0; fi50 < 16; fi50++ {
                var v51 int16 = 1
                v2 = (-(v2))
                var v52 uint64 = 45
                _ = fi50
                _ = v51
                _ = v52
            }
            _ = fi48
            _ = v49
        }
    }
    fmt.Println(v2)
    v53 := GenS6{F0: "hello", F1: 0, F2: -2147483648}
    for fi54 := 0; fi54 < 1; fi54++ {
        var v55 int64 = v3
        v2 += (-(v2))
        _ = fi54
        _ = v55
    }
    v2 = 42
    for fi56 := 0; fi56 < 14; fi56++ {
        v3 = (v3 - v3)
        var v57 float32 = float32((0.1 * -2.25))
        _ = fi56
        _ = v57
    }
    v4 = append(v4, 362)
    _ = len(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v46
    _ = v53
}
