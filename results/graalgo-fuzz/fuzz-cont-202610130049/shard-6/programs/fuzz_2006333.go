package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
func genFn0(p0 int8, p1 int) uint16 {
    var v6 int16 = 42
    var v7 bool = (65535 != 2)
    fmt.Println(p1, p1, p1)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1() uint8 {
    var v8 uint64 = 4294967295
    v9 := []int{-9223372036854775808, 127, 2147483647, 127}
    if true {
        var v10 int = 128
        var v11 string = ("a" + ((((("a" + ("\n\t" + ("hello" + ("hello" + "hello")))) + ("foo bar" + ("foo bar" + ("\n\t" + "hello")))) + ("a" + (("" + "") + ""))) + "\n\t") + ("hello" + "")))
        _ = v10
        _ = v11
    }
    var v12 uint64 = v8
    var v13 int = -1
    _ = v12
    _ = v13
    _ = v8
    _ = v9
    return 78
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = -1
    var v3 int32 = 2147483647
    v4 := make(map[string]string)
    var v5 float32 = v1
    var v14 float64 = (1.5 * (-2.25 / (1.5 / 1.0)))
    fmt.Println(v1, v3)
    switch "\n\t" {
    case "":
        for fi15 := 0; fi15 < 4; fi15++ {
            if (("" < "") || true) {
                fmt.Println(v1)
            } else {
                var v16 int32 = -53
                _ = v16
            }
            _ = len(v4)
            var v17 float64 = ((((v14 - v14) / v14) * v14) * v14)
            _ = fi15
            _ = v17
        }
    default:
        delete(v4, "foo bar")
    }
    v18 := []int{2147483647, -148, -323}
    var v19 int64 = 1
    v2 += v2
    v5 = v5
    for fi20 := 0; fi20 < 5; fi20++ {
        v14 = 1e10
        var v21 uint8 = 128
        _ = fi20
        _ = v21
    }
    switch "a" {
    case "hello":
        v22 := GenS1{F0: -128, F1: v19}
        _ = v22
    case "foo bar":
        v1 -= v5
    default:
        var v23 string = ""
        _ = v23
    }
    v24 := GenS1{F0: -128, F1: -322}
    v19 += 1
    v25 := GenS1{F0: 127, F1: -1}
    var v26 int64 = -9223372036854775808
    if (false && true) {
        fmt.Println(v5, v5, v26)
        delete(v4, "a")
        v3 += v3
    } else {
        for fi27 := 0; fi27 < 9; fi27++ {
            v3 += -1
            switch "" {
            case "\n\t":
                fmt.Println(v14, v5)
            case "foo bar":
                v3 -= (v3 & v3)
            case "a":
                switch (fi27 % (((fi27 >> ((255) & 63))) | 1)) {
                case -9223372036854775808:
                    var v28 int32 = v3
                    _ = v28
                case 1:
                    if (!false) {
                        var v29 int16 = -32768
                        for fi30 := 0; fi30 < 12; fi30++ {
                            var v31 int = fi27
                            if len(v18) >= 2 {
                                v18 = v18[1:len(v18)]
                            }
                            if (0 > 65535) {
                                v2 = v29
                            }
                            v32 := GenS1{F0: 0, F1: (v26 * v26)}
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                        for fi33 := 0; fi33 < 12; fi33++ {
                            var v34 uint8 = 179
                            var v35 uint8 = genFn1()
                            var v36 bool = true
                            switch fi27 {
                            case -9223372036854775808:
                                v18 = append(v18, fi33)
                            default:
                                v35 = v35
                            }
                            _ = fi33
                            _ = v34
                            _ = v35
                            _ = v36
                        }
                        _ = v29
                    } else {
                        fmt.Println(v3, v5)
                        switch (fi27 & fi27) {
                        case -23:
                            if (314 >= 2147483648) {
                                if len(v18) >= 2 {
                                    v18 = v18[1:len(v18)]
                                }
                                v37 := GenS1{F0: 127, F1: v26}
                                fmt.Println(v14)
                                _ = v37
                            } else {
                                v4[("a" + "foo bar")] = "hello"
                            }
                        case -496:
                            var v38 int16 = -471
                            _ = v38
                        default:
                            fmt.Println(v3)
                        }
                        var v39 int8 = 4
                        _ = v39
                    }
                case 32767:
                    var v40 int = fi27
                    _ = v40
                default:
                    switch fi27 {
                    case -258:
                        v41 := GenS1{F0: -15, F1: (^(v19))}
                        _ = v41
                    case 370:
                        switch fi27 {
                        case 32767:
                            var v42 float32 = 0.1
                            _ = v42
                        default:
                            fmt.Println(v2, v1, v2)
                        }
                    case -1:
                        switch "hello" {
                        case "\n\t":
                            v43 := GenS1{F0: 42, F1: (v26 << ((1) & 63))}
                            _ = v43
                        case "foo bar":
                            var v44 float32 = -2.25
                            _ = v44
                        default:
                            v19 = (v19 ^ (v26 + (v26 | (v19 << ((443) & 63)))))
                        }
                    }
                }
            }
            if false {
                var v45 bool = false
                var v46 int8 = -110
                v47 := GenS1{F0: -128, F1: (v26 << ((255) & 63))}
                _ = v45
                _ = v46
                _ = v47
            }
            _ = fi27
        }
    }
    fmt.Println(v3, v3, v5)
    v18 = append(v18, 1)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
