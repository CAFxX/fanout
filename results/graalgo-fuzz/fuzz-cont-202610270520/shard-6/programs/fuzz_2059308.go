package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 int; F1 int8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int32 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 string; F1 uint8; F2 uint }
type GenS8 struct { F0 int64; F1 bool; F2 uint }
type GenS9 struct { F0 int8; F1 uint8; F2 string }
type GenS10 struct { F0 uint8; F1 string; F2 int }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 int64 = 429
        v7 -= 0
        _ = fi6
        _ = v7
    }
    var v8 uint64 = 951
    v9 := []int{127, int(v8)}
    if len(v9) > 0 {
        v9[0] = -1
    }
    _ = v8
    _ = v9
    return 42
}
func genFn1() int {
    var v10 uint = 65535
    if true {
        v11 := GenS1{F0: (!(-284 <= -1)), F1: (false && ((false || (128 != 1)) || (993 < -109))), F2: 2}
        v10 -= (^(v10))
        _ = v11
    }
    _ = v10
    return 2147483647
}

func main() {
    var v1 int = -2147483648
    var v2 string = ""
    v3 := make(map[string]string)
    v4 := make(map[uint32]bool)
    v5 := []int{}
    var v12 int16 = 0
    for fi13 := 0; fi13 < 4; fi13++ {
        v12 = -32768
        _ = fi13
    }
    if (127 != -116) {
        v14 := GenS2{F0: (v12 - (v12 << ((538) & 63))), F1: 497}
        v1 = (v1 >> ((772) & 63))
        _ = v14
    } else {
        v1 = (v1 | genFn1())
        fmt.Println(v12, v12)
        v2 = v2
    }
    v12 -= v12
    fmt.Println(v2)
    var v15 uint8 = 0
    for fi16 := 0; fi16 < 12; fi16++ {
        v15 -= v15
        _ = fi16
    }
    if (336 == v12) {
        for fi17 := 0; fi17 < 15; fi17++ {
            v18 := GenS2{F0: -397, F1: 4294967295}
            v5 = append(v5, genFn1())
            _ = fi17
            _ = v18
        }
        if true {
            v3[v2] = (v2 + (v2 + v2))
            v19 := GenS3{F0: (v1 + (v1 + v1)), F1: int8(v1), F2: 2147483647}
            _ = v19
        }
        var v20 uint64 = 620
        _ = v20
    }
    if false {
        v12 += v12
    }
    switch (v1 * 2147483647) {
    case 2147483647:
        if true {
            v21 := GenS1{F0: false, F1: true, F2: 65535}
            var v22 int64 = -88
            _ = v21
            _ = v22
        } else {
            for fi23 := 0; fi23 < 9; fi23++ {
                switch (v1 << ((1) & 63)) {
                case -884:
                    fmt.Println(v2, v1)
                case -507:
                    switch fi23 {
                    case 127:
                        v2 = v2
                    case -1:
                        switch v2 {
                        case "a":
                            for fi24 := 0; fi24 < 20; fi24++ {
                                if false {
                                    var v25 uint16 = 366
                                    _ = v25
                                } else {
                                    fmt.Println(fi23)
                                    if (!true) {
                                        v26 := GenS2{F0: v12, F1: 2}
                                        _ = v26
                                    }
                                    var v27 uint64 = 18446744073709551615
                                    v28 := GenS4{F0: 21, F1: 0}
                                    fmt.Println(v27, v2)
                                    _ = v27
                                    _ = v28
                                }
                                switch v2 {
                                case "\n\t":
                                    v29 := GenS5{F0: 42}
                                    _ = v29
                                }
                                v5 = append(v5, (fi23 + v1))
                                switch v1 {
                                case 127:
                                    v30 := GenS5{F0: v12}
                                    _ = v30
                                default:
                                    fmt.Println(v15)
                                }
                                if true {
                                    var v31 uint8 = (v15 - (v15 * 128))
                                    if (1e10 >= (0.0 * 0.0)) {
                                        switch (("a" + ("hello" + "\n\t")) + v2) {
                                        case "":
                                            v12 = v12
                                        default:
                                            if len(v5) > 0 {
                                                v5[0] = (fi24 - (fi24 % ((v1) | 1)))
                                            }
                                        }
                                    } else {
                                        var v32 int8 = 81
                                        if false {
                                            v33 := GenS6{F0: 1, F1: v2}
                                            var v34 bool = true
                                            _ = v33
                                            _ = v34
                                        } else {
                                            switch ((((v2 + "\n\t") + (("foo bar" + (v2 + ("\n\t" + v2))) + "a")) + v2) + "hello") {
                                            case "foo bar":
                                                switch v2 {
                                                case "":
                                                    switch (^(fi23)) {
                                                    case 232:
                                                        v3[(("foo bar" + v2) + (v2 + ("a" + (v2 + "hello"))))] = v2
                                                    case 1:
                                                        if (1 <= 9223372036854775807) {
                                                            fmt.Println(v32)
                                                        }
                                                    default:
                                                        if (false || true) {
                                                            v1 = v1
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        v35 := GenS7{F0: "a", F1: (v31 + 2), F2: 18446744073709551615}
                                                        _ = v35
                                                    }
                                                    fmt.Println(fi24)
                                                }
                                            case "":
                                                v36 := GenS8{F0: -9223372036854775808, F1: false, F2: 4294967295}
                                                _ = v36
                                            case "a":
                                                v31 = v15
                                            default:
                                                v37 := GenS8{F0: int64(v31), F1: (true && true), F2: 18446744073709551615}
                                                _ = v37
                                            }
                                            var v38 uint = 849
                                            var v39 int = 32767
                                            _ = v38
                                            _ = v39
                                        }
                                        v5 = append(v5, fi23)
                                        _ = v32
                                    }
                                    var v40 uint = 1
                                    v41 := GenS1{F0: (!(!(!false))), F1: false, F2: 255}
                                    _ = v31
                                    _ = v40
                                    _ = v41
                                }
                                _ = fi24
                            }
                        default:
                            switch v2 {
                            case "\n\t":
                                switch (v1 | genFn1()) {
                                case -2147483648:
                                    v15 = 127
                                case 0:
                                    var v42 int = genFn1()
                                    _ = v42
                                case 883:
                                    v15 = 125
                                }
                            case "a":
                                var v43 int8 = -128
                                _ = v43
                            default:
                                fmt.Println(v12, v12)
                            }
                        }
                    }
                case 0:
                    var v44 int64 = 810
                    _ = v44
                default:
                    if false {
                        v45 := GenS9{F0: -95, F1: (v15 | v15), F2: v2}
                        _ = len(v3)
                        _ = v45
                    }
                }
                _ = fi23
            }
            var v46 bool = (!true)
            _ = v46
        }
    case -805:
        if false {
            if false {
                v3[(v2 + v2)] = "foo bar"
                if true {
                    v5 = append(v5, v1)
                    if len(v5) > 0 {
                        v5[0] = v1
                    }
                }
            }
            var v47 int8 = -1
            _ = v47
        }
    default:
        v15 = (v15 + 1)
    }
    var v48 int = (-(v1))
    v49 := GenS10{F0: v15, F1: ((("\n\t" + "foo bar") + "\n\t") + v2), F2: v48}
    var v50 bool = (42 <= 127)
    v48 = (v48 - v48)
    v50 = (v15 <= 127)
    v50 = v50
    var v51 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
}
