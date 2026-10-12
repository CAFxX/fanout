package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 int8 }
type GenS2 struct { F0 int8; F1 string; F2 uint64 }
func genFn0() int32 {
    var v6 int = 422
    var v7 int8 = 100
    for fi8 := 0; fi8 < 2; fi8++ {
        v6 -= (v6 | (v6 | fi8))
        for fi9 := 0; fi9 < 17; fi9++ {
            var v10 int8 = -128
            var v11 int32 = 1
            v10 = v7
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    var v12 uint32 = 2
    switch (^(v6)) {
    case 0:
        var v13 int8 = (v7 >> ((4294967295) & 63))
        _ = v13
    case 32767:
        v12 = (v12 * v12)
    }
    _ = v12
    _ = v6
    _ = v7
    return -625
}
func genFn1(p0 int16, p1 int32, p2 uint64) uint64 {
    for fi14 := 0; fi14 < 18; fi14++ {
        var v15 int8 = 1
        v16 := GenS1{F0: p1, F1: 182, F2: v15}
        var v17 int = (fi14 ^ 128)
        _ = fi14
        _ = v15
        _ = v16
        _ = v17
    }
    fmt.Println(p2, p1)
    for fi18 := 0; fi18 < 17; fi18++ {
        fmt.Println(p2, p1)
        p2 = 4294967295
        var v19 uint32 = 1
        _ = fi18
        _ = v19
    }
    if false {
        fmt.Println(p0, p2, p0)
        fmt.Println(p0, p2, p1)
    } else {
        p1 = 2147483647
        fmt.Println(p0, p0, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    return 18446744073709551615
}

func main() {
    v1 := []string{"a", (("hello" + "a") + "a"), "a"}
    var v2 float32 = (-2.25 + 1e10)
    v3 := make(map[string]bool)
    var v4 float32 = ((v2 / ((v2 + (((v2 + 0.1) / v2) - float32(float64(v2)))) - v2)) - 1e10)
    v5 := []int{32767, 2147483647, -615, 489}
    v2 -= v2
    for fi20 := 0; fi20 < 2; fi20++ {
        var v21 uint8 = 144
        if (0 != genFn1(670, 0, 18446744073709551615)) {
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 int = 32767
                _ = fi22
                _ = v23
            }
        }
        if false {
            if ((-12 <= int8(fi20)) || (0 > 2147483647)) {
                var v24 int16 = 32767
                v24 = (-(v24))
                for fi25 := 0; fi25 < 20; fi25++ {
                    v4 = 1e10
                    _ = fi25
                }
                _ = v24
            }
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v21 = (v21 & 0)
        }
        _ = fi20
        _ = v21
    }
    v4 = 0.0
    _ = len(v3)
    if false {
        if true {
            var v26 float64 = (1e10 - 3.14159)
            _ = v26
        }
        v4 += (1.5 / (float32((float64(1.5) - (-2.25 + 3.14159))) * (float32((0.1 - (0.0 + (0.0 - ((((1e10 / 1e10) - (((1.5 - 0.1) + -2.25) * (-2.25 / (1e10 + (0.1 - -2.25))))) / 0.1) - 1e10))))) - v4)))
        switch "\n\t" {
        case "a":
            var v27 int8 = 42
            _ = v27
        }
    } else {
        if true {
            switch "foo bar" {
            case "hello":
                var v28 int64 = -216
                _ = v28
            default:
                switch (("hello" + "") + "foo bar") {
                case "":
                    if (!(false && true)) {
                        for fi29 := 0; fi29 < 2; fi29++ {
                            for fi30 := 0; fi30 < 15; fi30++ {
                                fmt.Println(fi30)
                                fmt.Println(fi29, v4, v2)
                                v2 += ((v2 / (v4 + 0.1)) + v2)
                                _ = fi30
                            }
                            if (!(!(-2147483648 > 2147483647))) {
                                var v31 uint32 = 2147483647
                                v32 := v3[("\n\t" + "\n\t")]
                                _ = v31
                                _ = v32
                            } else {
                                switch fi29 {
                                case 390:
                                    for fi33 := 0; fi33 < 13; fi33++ {
                                        v34 := GenS2{F0: -128, F1: "hello", F2: 1}
                                        var v35 int64 = 1
                                        delete(v3, (("\n\t" + ("" + "foo bar")) + ("\n\t" + "hello")))
                                        _ = fi33
                                        _ = v34
                                        _ = v35
                                    }
                                case 32767:
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case 32768:
                                    for fi36 := 0; fi36 < 15; fi36++ {
                                        switch "" {
                                        case "foo bar":
                                            v2 += (v4 / v2)
                                        }
                                        v2 -= v4
                                        var v37 int32 = 1
                                        var v38 int8 = 1
                                        v37 = (v37 + v37)
                                        _ = fi36
                                        _ = v37
                                        _ = v38
                                    }
                                }
                            }
                            for fi39 := 0; fi39 < 9; fi39++ {
                                fmt.Println(fi29, v4)
                                var v40 uint16 = 1
                                _ = fi39
                                _ = v40
                            }
                            _ = fi29
                        }
                        for fi41 := 0; fi41 < 3; fi41++ {
                            var v42 uint16 = 1
                            fmt.Println(v4, v4)
                            _ = fi41
                            _ = v42
                        }
                        switch 42 {
                        case 32768:
                            var v43 int = 2147483647
                            _ = v43
                        default:
                            if (!false) {
                                for fi44 := 0; fi44 < 5; fi44++ {
                                    v45 := GenS2{F0: -128, F1: "foo bar", F2: 4294967295}
                                    var v46 float64 = 1e10
                                    _ = fi44
                                    _ = v45
                                    _ = v46
                                }
                                v47 := GenS1{F0: -891, F1: 0, F2: 0}
                                var v48 uint64 = 769
                                _ = v47
                                _ = v48
                            }
                        }
                    }
                case "\n\t":
                    v2 = float32((float64(v4) * (-2.25 - 0.0)))
                case "hello":
                    var v49 uint = 4294967295
                    _ = v49
                }
            }
            if ((17 == 42) || ("" >= ("hello" + ((("foo bar" + ("foo bar" + "\n\t")) + "\n\t") + "a")))) {
                var v50 int = 42
                _ = v50
            }
        }
        fmt.Println(v2, v2)
    }
    if (-128 < 42) {
        var v51 int8 = -1
        _ = v51
    } else {
        var v52 int32 = -1
        if (true && true) {
            switch "" {
            case "hello":
                if false {
                    v53 := GenS1{F0: v52, F1: 258, F2: -1}
                    var v54 float64 = 3.14159
                    var v55 bool = (!(((1e10 > 3.14159) && false) && (!false)))
                    _ = v53
                    _ = v54
                    _ = v55
                }
            default:
                v56 := GenS1{F0: genFn0(), F1: 1, F2: -1}
                _ = v56
            }
        } else {
            v57 := v3[("a" + (("\n\t" + "a") + "\n\t"))]
            v58 := GenS1{F0: (v52 | v52), F1: 1, F2: -1}
            _ = v57
            _ = v58
        }
        _ = v52
    }
    var v59 int64 = 9223372036854775807
    fmt.Println(v2, v2)
    fmt.Println(v4, v4, v2)
    if false {
        switch (("" + "hello") + "") {
        case "foo bar":
            if false {
                v59 = 1
            } else {
                v5 = append(v5, 9223372036854775807)
                var v60 float64 = (((0.1 + (float64((v2 * (v4 + 0.0))) + float64((0.1 - float32(1e10))))) + -2.25) / ((0.0 - (1e10 / 1.0)) - (0.0 + 0.1)))
                v2 = (((-2.25 + v4) + v4) / v4)
                _ = v60
            }
        }
    }
    v4 -= v4
    if (1 != 1) {
        v5 = append(v5, 0)
        v2 = 0.0
    }
    var v61 uint64 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v59:", v59)
    fmt.Println("v61:", v61)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v59
    _ = v61
}
