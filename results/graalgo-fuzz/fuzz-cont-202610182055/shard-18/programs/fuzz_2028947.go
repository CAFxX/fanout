package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 string }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64; F1 int8; F2 bool }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 uint32; F2 int64 }

func main() {
    v1 := []int{32768, 128, -9223372036854775808}
    v2 := make(map[string]uint32)
    var v3 uint16 = 51
    for fi4 := 0; fi4 < 6; fi4++ {
        for fi5 := 0; fi5 < 13; fi5++ {
            v3 += (v3 ^ 1)
            v3 -= v3
            fmt.Println(v3, v3, fi4)
            v3 = v3
            var v6 int = fi4
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    if true {
        v7 := v2[("" + "\n\t")]
        if (0 == 42) {
            v7 = 1
            fmt.Println(v3, v7)
        } else {
            v8 := GenS1{F0: 833, F1: false, F2: ("hello" + ("a" + "hello"))}
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := GenS2{F0: 42, F1: 2}
                var v11 uint8 = 127
                _ = fi9
                _ = v10
                _ = v11
            }
            fmt.Println(v3, v3)
            _ = v8
        }
        if false {
            _ = len(v2)
            _ = len(v2)
            switch ("hello" + ("" + ("" + "foo bar"))) {
            case "\n\t":
                var v12 uint64 = 361
                _ = v12
            }
        }
        _ = v7
    } else {
        v3 = v3
        var v13 uint64 = 1
        _ = v13
    }
    v3 = (v3 | v3)
    v1 = append(v1, -9223372036854775808)
    var v14 uint16 = 0
    var v15 float64 = 1.5
    switch -2147483648 {
    case 32768:
        v16 := GenS2{F0: 1, F1: 2}
        _ = v16
    default:
        switch 127 {
        case 484:
            v2["foo bar"] = 727
        case -1:
            v15 += 0.0
        }
    }
    v3 = (v14 / ((1) | 1))
    v17 := GenS2{F0: 394, F1: 18446744073709551615}
    for fi18 := 0; fi18 < 16; fi18++ {
        v15 = 0.1
        if false {
            var v19 bool = false
            switch "" {
            case "foo bar":
                var v20 bool = v19
                _ = v20
            default:
                v19 = (v14 > v3)
            }
            _ = v19
        } else {
            if (false && (42 < -538)) {
                switch ("foo bar" + ("a" + ("\n\t" + "\n\t"))) {
                case "hello":
                    var v21 int64 = 42
                    _ = v21
                case "a":
                    v22 := GenS3{F0: -2147483648}
                    _ = v22
                case "":
                    fmt.Println(fi18)
                }
                switch ("\n\t" + ("\n\t" + "a")) {
                case "foo bar":
                    for fi23 := 0; fi23 < 15; fi23++ {
                        switch fi23 {
                        case -2147483648:
                            v3 = 1
                        default:
                            v3 = v3
                        }
                        _ = fi23
                    }
                case "":
                    if (false && (2 <= 676)) {
                        _ = len(v1)
                        _ = len(v2)
                    }
                }
            } else {
                if (v3 > (v14 / (((v14 * (v14 / (((v14 / (((^(v14))) | 1))) | 1)))) | 1))) {
                    v3 -= v14
                    fmt.Println(v3, v3)
                } else {
                    _ = len(v2)
                }
            }
            var v24 uint32 = 2147483647
            if true {
                v3 = v14
            } else {
                if (v3 <= (v14 << ((607) & 63))) {
                    if (("a" + ("" + (("foo bar" + "\n\t") + "hello"))) <= "\n\t") {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            if true {
                                v3 = v14
                                fmt.Println(fi25, v14, fi18)
                            } else {
                                var v26 int64 = 0
                                var v27 uint8 = 128
                                v28 := GenS4{F0: 18446744073709551615, F1: -128, F2: ((!(42 > -532)) && (true && false))}
                                _ = v26
                                _ = v27
                                _ = v28
                            }
                            _ = fi25
                        }
                        if (1 < uint8(v3)) {
                            var v29 int32 = -538
                            _ = v29
                        } else {
                            for fi30 := 0; fi30 < 3; fi30++ {
                                var v31 int = (fi18 | fi18)
                                var v32 uint = uint(v14)
                                _ = fi30
                                _ = v31
                                _ = v32
                            }
                            var v33 bool = true
                            _ = v33
                        }
                        for fi34 := 0; fi34 < 6; fi34++ {
                            if (uint64(v24) > 873) {
                                var v35 float32 = 3.14159
                                var v36 int16 = 436
                                _ = v35
                                _ = v36
                            }
                            if false {
                                var v37 int = (fi18 >> ((2) & 63))
                                v3 = (v14 / ((0) | 1))
                                _ = v37
                            }
                            var v38 int8 = -128
                            _ = fi34
                            _ = v38
                        }
                    } else {
                        for fi39 := 0; fi39 < 16; fi39++ {
                            if len(v1) > 0 {
                                v1[0] = fi18
                            }
                            var v40 uint16 = v14
                            if (((v24 % ((2147483647) | 1)) < (v24 % ((872) | 1))) || false) {
                                fmt.Println(v3, fi39, v40)
                                v15 += (((1.5 / v15) - 0.0) / v15)
                                fmt.Println(fi39, fi39)
                                fmt.Println(v24)
                                fmt.Println(fi18, v40)
                                v40 = 1
                                switch "" {
                                case "\n\t":
                                    var v41 uint64 = 4294967295
                                    _ = v41
                                default:
                                    v42 := GenS5{F0: 2}
                                    _ = v42
                                }
                            } else {
                                _ = len(v1)
                                delete(v2, "a")
                            }
                            _ = fi39
                            _ = v40
                        }
                        v24 += v24
                    }
                    v1 = append(v1, fi18)
                }
                v15 = v15
            }
            _ = v24
        }
        v15 = float64(float32(3.14159))
        _ = fi18
    }
    var v43 uint32 = 42
    var v44 string = "\n\t"
    if (!false) {
        v15 = 1e10
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v45 := GenS6{F0: -930, F1: (v43 >> ((0) & 63)), F2: 42}
    v43 -= (^(v43))
    if true {
        for fi46 := 0; fi46 < 16; fi46++ {
            fmt.Println(v43, v14, v3)
            _ = fi46
        }
    } else {
        if false {
            v1 = append(v1, 42)
            if false {
                var v47 float32 = 1e10
                _ = v47
            } else {
                fmt.Println(v43, v3, v3)
                v44 = (((v44 + (v44 + v44)) + (v44 + (v44 + v44))) + "\n\t")
                v3 -= (v3 | 0)
            }
            fmt.Println(v44)
        } else {
            v43 -= v43
            v43 += (v43 * v43)
        }
        v2["a"] = v43
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v43
    _ = v44
    _ = v45
}
