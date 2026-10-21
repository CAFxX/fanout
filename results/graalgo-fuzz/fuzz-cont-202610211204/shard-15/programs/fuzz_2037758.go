package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 string }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 bool; F1 int16 }
type GenS7 struct { F0 uint8; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint; F1 int16; F2 int }
type GenS9 struct { F0 uint8; F1 uint16; F2 int16 }
func genFn0(p0 int16, p1 uint) int8 {
    var v7 float32 = (1.5 * (0.0 + 0.0))
    switch 922 {
    case 599:
        v8 := GenS1{F0: 0, F1: 1000, F2: "a"}
        _ = v8
    case 902:
        if (42 >= 42) {
            var v9 int8 = -128
            v10 := GenS1{F0: 32, F1: (p0 | p0), F2: ("a" + "")}
            _ = v10
            _ = v9
        }
    }
    v11 := []int{2147483647, 9223372036854775807}
    _ = p0
    _ = p1
    _ = v11
    _ = v7
    return -1
}

func main() {
    var v1 string = "foo bar"
    var v2 uint8 = 251
    v3 := make(map[uint32]bool)
    var v4 float32 = 1.5
    v5 := []string{}
    v6 := []int{-9223372036854775808}
    v1 = (v1 + v1)
    v12 := GenS1{F0: 727, F1: 42, F2: v1}
    v13 := GenS1{F0: uint16(v2), F1: 174, F2: ((v1 + ("a" + ("foo bar" + "foo bar"))) + "")}
    for fi14 := 0; fi14 < 1; fi14++ {
        fmt.Println(v1)
        v15 := GenS1{F0: 2, F1: -1, F2: ((v1 + "\n\t") + ("" + "a"))}
        switch (fi14 + fi14) {
        case 700:
            if (-1 == 42) {
                v4 += (1.5 + v4)
                v16 := GenS1{F0: 821, F1: 0, F2: ((v1 + "hello") + v1)}
                _ = v16
            } else {
                v17 := GenS2{F0: 2, F1: 1}
                var v18 uint = 209
                for fi19 := 0; fi19 < 4; fi19++ {
                    v20 := GenS1{F0: 65535, F1: -815, F2: "\n\t"}
                    if (-9223372036854775808 < 42) {
                        delete(v3, 65535)
                        v3[0] = false
                    }
                    _ = fi19
                    _ = v20
                }
                _ = v17
                _ = v18
            }
        case 4294967295:
            v21 := GenS3{F0: "foo bar", F1: v2, F2: 255}
            _ = v21
        default:
            var v22 uint = 168
            _ = v22
        }
        _ = fi14
        _ = v15
    }
    v23 := GenS2{F0: 358, F1: 0}
    switch "foo bar" {
    case "foo bar":
        v4 = (v4 + 3.14159)
    case "a":
        v2 += 82
    case "":
        v2 = v2
    default:
        var v24 string = (v1 + v1)
        _ = v24
    }
    if (0 > -39) {
        switch 32767 {
        case 42:
            if true {
                if true {
                    fmt.Println(v4)
                    switch 128 {
                    case -2147483648:
                        for fi25 := 0; fi25 < 16; fi25++ {
                            v1 = (("\n\t" + (("a" + v1) + v1)) + v1)
                            var v26 int = fi25
                            var v27 int64 = 385
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    case 128:
                        v1 = ""
                    case 32767:
                        if true {
                            if true {
                                switch 1 {
                                case 420:
                                    v28 := GenS3{F0: v1, F1: 53, F2: 255}
                                    _ = v28
                                }
                            }
                            fmt.Println(v4)
                        }
                    default:
                        if len(v6) > 0 {
                            v6[0] = 32767
                        }
                    }
                    var v29 float64 = 3.14159
                    _ = v29
                } else {
                    v30 := GenS4{F0: 42}
                    _ = v30
                }
                v2 = v2
                var v31 int = -1
                _ = v31
            }
        case 114:
            v32 := GenS5{F0: -14}
            _ = v32
        case 128:
            if (0 > 2) {
                var v33 int32 = 42
                v34 := GenS3{F0: "", F1: 240, F2: 462}
                _ = v33
                _ = v34
            }
        default:
            var v35 uint32 = 207
            _ = v35
        }
        fmt.Println(v1, v2)
        fmt.Println(v2, v2)
    }
    fmt.Println(v1, v4)
    v36 := GenS6{F0: ((!true) && ((1 <= -188) && true)), F1: 1000}
    if false {
        v37 := GenS1{F0: 0, F1: -524, F2: (v1 + v1)}
        for fi38 := 0; fi38 < 10; fi38++ {
            var v39 int32 = 42
            _ = fi38
            _ = v39
        }
        _ = v37
    } else {
        v4 = (v4 * (float32(0.0) * 1e10))
        fmt.Println(v2)
        if true {
            delete(v3, 2147483647)
        }
    }
    v4 += ((-2.25 / v4) - ((float32(-2.25) - v4) * float32(1.5)))
    v4 = v4
    fmt.Println(v2)
    for fi40 := 0; fi40 < 7; fi40++ {
        for fi41 := 0; fi41 < 15; fi41++ {
            v42 := GenS7{F0: (v2 * v2), F1: 2, F2: -9223372036854775808}
            if (!false) {
                if true {
                    v4 -= (float32(float64(1e10)) * v4)
                    var v43 uint32 = 218
                    var v44 uint16 = 697
                    _ = v43
                    _ = v44
                }
            }
            switch (v1 + v1) {
            case "\n\t":
                v2 = (v2 ^ 120)
            case "a":
                v45 := GenS6{F0: (0 == 2), F1: -387}
                _ = v45
            case "foo bar":
                v46 := GenS8{F0: 53, F1: 32767, F2: (fi40 << ((2) & 63))}
                _ = v46
            default:
                v47 := GenS9{F0: 42, F1: uint16(v2), F2: -695}
                if false {
                    v48 := v3[728]
                    _ = v48
                }
                _ = v47
            }
            _ = fi41
            _ = v42
        }
        _ = fi40
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        if true {
            v50 := GenS9{F0: (v2 << ((255) & 63)), F1: 255, F2: 32767}
            var v51 string = (v1 + (("" + (v1 + v1)) + v1))
            _ = v50
            _ = v51
        } else {
            v52 := GenS6{F0: false, F1: -32768}
            _ = v52
        }
        for fi53 := 0; fi53 < 2; fi53++ {
            v54 := GenS7{F0: (v2 >> ((880) & 63)), F1: uint16(v2), F2: -9223372036854775808}
            if (-1 == 66) {
                var v55 float32 = 1.5
                v4 = v55
                _ = v55
            }
            v56 := GenS3{F0: "hello", F1: (^(v2)), F2: 186}
            var v57 int32 = 638
            v58 := GenS4{F0: 0}
            _ = fi53
            _ = v54
            _ = v56
            _ = v57
            _ = v58
        }
        switch (fi49 / ((fi49) | 1)) {
        case 9223372036854775807:
            for fi59 := 0; fi59 < 15; fi59++ {
                var v60 float32 = (1.5 * ((v4 / 1.0) + -2.25))
                _ = fi59
                _ = v60
            }
        default:
            v2 = (v2 - (v2 - v2))
        }
        _ = fi49
    }
    fmt.Println(v2)
    v3[4294967295] = false
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
