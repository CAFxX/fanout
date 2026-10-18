package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint; F2 int }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 int16 = 42
    v4 := []int{128}
    v5 := []string{"foo bar", ("\n\t" + (("hello" + (("foo bar" + "\n\t") + "hello")) + "\n\t")), ("foo bar" + ((("hello" + "") + "a") + ("" + "hello")))}
    v6 := GenS1{F0: 2, F1: "\n\t", F2: 127}
    v7 := v1[24]
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v3 = 1000
            if true {
                var v10 int = fi9
                _ = v10
            } else {
                if (!(((v7 >> ((65535) & 63)) == v7) && true)) {
                    var v11 float64 = ((-2.25 * 1.5) * (0.1 - 1.5))
                    switch "" {
                    case "hello":
                        if (127 < 42) {
                            if (((false && (false || false)) || (!true)) || (!true)) {
                                var v12 string = ("a" + "")
                                var v13 float32 = 1e10
                                var v14 bool = (false || ((0.1 / v13) == ((1e10 / v13) * (0.1 - v13))))
                                _ = v12
                                _ = v13
                                _ = v14
                            } else {
                                if true {
                                    if (((1e10 * ((0.0 - 3.14159) + 0.0)) < (0.1 + 1e10)) || (-1 >= 1)) {
                                        v11 = v11
                                        switch (fi8 * fi9) {
                                        case 32768:
                                            v11 = v11
                                        default:
                                            if false {
                                                v7 -= (v7 + 65535)
                                            } else {
                                                if false {
                                                    var v15 uint32 = 681
                                                    _ = v15
                                                }
                                            }
                                        }
                                    } else {
                                        if len(v4) > 0 {
                                            v4[0] = (-(fi9))
                                        }
                                        var v16 uint16 = 0
                                        v11 -= v11
                                        if true {
                                            v17 := GenS2{F0: 42, F1: 1, F2: 127}
                                            _ = v17
                                        }
                                        v18 := GenS1{F0: 127, F1: "", F2: int8(v7)}
                                        _ = v16
                                        _ = v18
                                    }
                                    v11 += ((0.1 * (0.1 + 3.14159)) / ((v11 * (v11 + 0.1)) * v11))
                                }
                                var v19 int64 = 0
                                fmt.Println(fi9, fi8, fi8)
                                _ = v19
                            }
                        } else {
                            var v20 uint32 = v7
                            _ = v20
                        }
                    case "a":
                        var v21 uint64 = 624
                        _ = v21
                    case "":
                        if (!false) {
                            var v22 int64 = -928
                            _ = v22
                        }
                    default:
                        var v23 uint8 = 0
                        _ = v23
                    }
                    v24 := GenS1{F0: 8, F1: "hello", F2: 0}
                    if (128 == 81) {
                        v25 := GenS3{F0: 255, F1: 4294967295, F2: fi9}
                        fmt.Println(v7)
                        _ = v25
                    }
                    var v26 float64 = v11
                    _ = v11
                    _ = v24
                    _ = v26
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                var v27 int32 = -1
                _ = v27
            }
            v28 := GenS1{F0: 127, F1: "\n\t", F2: 15}
            v3 = v3
            if (!(42 != 104)) {
                v29 := v1[1]
                var v30 string = ""
                switch ((v30 + v30) + (v30 + "foo bar")) {
                case "\n\t":
                    var v31 uint8 = 42
                    _ = v31
                case "foo bar":
                    if (!true) {
                        if (65535 >= 255) {
                            var v32 int = (^(fi9))
                            _ = v32
                        }
                        v30 = ""
                        fmt.Println(v29, v3, fi9)
                    } else {
                        var v33 int32 = 42
                        _ = v33
                    }
                default:
                    if true {
                        v34 := GenS1{F0: 42, F1: "hello", F2: 0}
                        fmt.Println(v29, v30)
                        _ = v34
                    }
                }
                _ = len(v1)
                v29 = (v29 % (((v29 + (^(v29)))) | 1))
                _ = v29
                _ = v30
            }
            v35 := v1[(^(v7))]
            if len(v4) > 0 {
                v4[0] = (^(fi8))
            }
            var v36 int = fi9
            _ = fi9
            _ = v28
            _ = v35
            _ = v36
        }
        v7 -= (v7 << ((2) & 63))
        v3 = v3
        _ = fi8
    }
    v3 -= v3
    var v37 uint16 = uint16(v7)
    var v38 float64 = -2.25
    if (float64(0.1) != 0.0) {
        fmt.Println(v37, v37)
        v7 = (v7 ^ (v7 >> ((813) & 63)))
        var v39 uint = 483
        _ = v39
    } else {
        for fi40 := 0; fi40 < 7; fi40++ {
            v37 += (v37 - (v37 >> ((532) & 63)))
            fmt.Println(v38)
            _ = fi40
        }
    }
    v3 = (v3 | v3)
    if (!(true && ((118 != 65535) && false))) {
        var v41 uint16 = v37
        var v42 uint64 = 4294967295
        var v43 uint8 = 167
        _ = v41
        _ = v42
        _ = v43
    }
    var v44 float32 = 1.5
    for fi45 := 0; fi45 < 3; fi45++ {
        v7 = 1
        var v46 uint32 = v7
        v5 = append(v5, "a")
        _ = fi45
        _ = v46
    }
    v37 = (v37 + v37)
    var v47 float32 = float32(3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v47
    _ = v5
    _ = v6
    _ = v7
}
