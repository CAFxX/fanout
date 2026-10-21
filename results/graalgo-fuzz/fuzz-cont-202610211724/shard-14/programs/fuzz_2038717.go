package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int8; F1 int; F2 int8 }
type GenS4 struct { F0 uint; F1 uint; F2 uint8 }
func genFn0() bool {
    var v5 uint16 = 32767
    if (((298 >= 1) && (!(!true))) || (!true)) {
        var v6 int32 = 2147483647
        fmt.Println(v5)
        v6 = v6
        _ = v6
    } else {
        var v7 float64 = 3.14159
        _ = v7
    }
    var v8 int16 = 280
    for fi9 := 0; fi9 < 20; fi9++ {
        v10 := GenS1{F0: 2147483647, F1: 648, F2: (-(v8))}
        v8 += v8
        _ = fi9
        _ = v10
    }
    _ = v5
    _ = v8
    return (!false)
}
func genFn1(p0 uint, p1 int, p2 int16) uint16 {
    var v11 uint32 = 2
    fmt.Println(p2)
    var v12 int = (p1 - p1)
    if (((p2 | (p2 << ((p0) & 63))) < (p2 | 0)) && true) {
        switch (("\n\t" + "") + ("foo bar" + "\n\t")) {
        case "hello":
            p2 = -32768
        }
        if (true || ((v11 * v11) == (^(v11)))) {
            v13 := GenS1{F0: 1, F1: 18446744073709551615, F2: p2}
            if true {
                p1 = v12
                p2 = 32767
            }
            if false {
                if (128 != 42) {
                    p0 = p0
                } else {
                    if ((0.0 * 1e10) > ((1e10 - 0.0) / 1.5)) {
                        if (4294967295 != 291) {
                            v14 := []int{int(v11)}
                            fmt.Println(v12)
                            _ = v14
                        }
                    }
                    switch "a" {
                    case "\n\t":
                        for fi15 := 0; fi15 < 3; fi15++ {
                            var v16 int16 = (p2 | -295)
                            v17 := []int{int(v11), p1, -231}
                            var v18 uint32 = 1
                            _ = fi15
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                    default:
                        if true {
                            p0 = uint(v11)
                            var v19 uint16 = 0
                            var v20 int = 32768
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v12, p2, p1)
                            var v21 int8 = -128
                            var v22 int8 = -29
                            _ = v21
                            _ = v22
                        }
                    }
                }
            } else {
                var v23 float32 = float32(-2.25)
                _ = v23
            }
            _ = v13
        } else {
            for fi24 := 0; fi24 < 6; fi24++ {
                if (true && true) {
                    if (42 <= 2) {
                        p1 = 9
                        switch v12 {
                        case 877:
                            if ((1.5 / (0.0 + 1e10)) > (-2.25 + (0.1 - (-2.25 / (((1.5 + (1e10 * 1.5)) - 3.14159) + (-2.25 + 0.1)))))) {
                                v25 := GenS1{F0: -319, F1: 42, F2: (p2 - (p2 | (p2 >> ((1) & 63))))}
                                _ = v25
                            }
                        case 0:
                            var v26 uint64 = 0
                            _ = v26
                        default:
                            for fi27 := 0; fi27 < 3; fi27++ {
                                v11 = v11
                                v11 += v11
                                _ = fi27
                            }
                        }
                        if true {
                            v28 := []int{(p1 % (((v12 << ((4294967295) & 63))) | 1)), (^(v12))}
                            if (true || (127 != 79)) {
                                fmt.Println(p1)
                                v29 := GenS1{F0: -674, F1: 318, F2: (-(p2))}
                                switch "" {
                                case "a":
                                    for fi30 := 0; fi30 < 4; fi30++ {
                                        if true {
                                            var v31 int16 = (p2 - p2)
                                            _ = v31
                                        }
                                        var v32 int64 = 0
                                        _ = fi30
                                        _ = v32
                                    }
                                }
                                _ = v29
                            } else {
                                p0 = 824
                                var v33 int64 = -534
                                _ = v33
                            }
                            _ = v28
                        }
                    } else {
                        var v34 string = "\n\t"
                        _ = v34
                    }
                } else {
                    var v35 uint = 2
                    for fi36 := 0; fi36 < 3; fi36++ {
                        p2 = 32767
                        _ = fi36
                    }
                    v12 -= fi24
                    _ = v35
                }
                var v37 uint = p0
                v38 := []int{v12, (p1 ^ 1), (p1 & fi24), v12}
                _ = fi24
                _ = v37
                _ = v38
            }
        }
        v39 := []int{855, p1, (p1 << (((p0 & (p0 * 255))) & 63)), (p1 * (^(v12)))}
        _ = v39
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 65535
}

func main() {
    var v1 int64 = 42
    var v2 int8 = 1
    v3 := make(map[string]string)
    v4 := []int{}
    var v40 uint16 = 42
    delete(v3, "hello")
    v1 -= v1
    if (65535 > 901) {
        fmt.Println(v40, v2)
        var v41 int16 = -1
        if ((^(v2)) < v2) {
            v40 += 255
        } else {
            var v42 uint16 = (v40 - (v40 >> ((383) & 63)))
            var v43 string = ("a" + "foo bar")
            _ = v42
            _ = v43
        }
        _ = v41
    } else {
        v44 := GenS2{F0: -716}
        for fi45 := 0; fi45 < 7; fi45++ {
            var v46 int64 = 790
            _ = fi45
            _ = v46
        }
        v40 += (v40 * 105)
        _ = v44
    }
    v40 = 227
    if (!true) {
        v47 := GenS2{F0: 466}
        for fi48 := 0; fi48 < 16; fi48++ {
            fmt.Println(fi48, v40)
            v49 := GenS1{F0: -801, F1: 42, F2: 42}
            _ = fi48
            _ = v49
        }
        _ = v47
    }
    var v50 uint64 = 889
    var v51 int8 = 42
    v52 := GenS3{F0: 42, F1: 0, F2: 13}
    v53 := GenS4{F0: 42, F1: 1, F2: 255}
    for fi54 := 0; fi54 < 11; fi54++ {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        var v55 bool = true
        _ = fi54
        _ = v55
    }
    switch "a" {
    case "foo bar":
        fmt.Println(v50, v51, v40)
    case "\n\t":
        if true {
            _ = len(v4)
        } else {
            switch "" {
            case "a":
                switch 128 {
                case -1:
                    var v56 int64 = v1
                    _ = v56
                default:
                    var v57 int32 = int32(v40)
                    _ = v57
                }
            case "":
                var v58 uint16 = 255
                _ = v58
            default:
                var v59 uint = 459
                _ = v59
            }
            v51 += 1
            for fi60 := 0; fi60 < 11; fi60++ {
                var v61 uint64 = v50
                _ = fi60
                _ = v61
            }
        }
    case "":
        v50 = v50
    default:
        v3[((("a" + ((("a" + ("" + ("hello" + ((("a" + "foo bar") + ("foo bar" + "")) + "\n\t")))) + "hello") + (("" + "foo bar") + "\n\t"))) + "") + "\n\t")] = "foo bar"
    }
    fmt.Println(v2, v40, v2)
    _ = len(v3)
    v4 = append(v4, 410)
    v2 = (v51 << ((385) & 63))
    v3[("\n\t" + "hello")] = "a"
    var v62 float64 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v62
}
