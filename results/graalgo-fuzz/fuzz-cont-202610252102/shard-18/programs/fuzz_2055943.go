package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 bool }
type GenS4 struct { F0 string; F1 int64; F2 string }
type GenS5 struct { F0 int64; F1 int16 }
type GenS6 struct { F0 int64; F1 uint8 }
type GenS7 struct { F0 uint16; F1 uint8; F2 bool }
type GenS8 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = -128
    v3 := []string{"hello", "a", ""}
    var v4 int = 42
    v5 := []uint32{}
    switch "a" {
    case "a":
        v1[1] = true
    case "hello":
        switch v4 {
        case 127:
            for fi6 := 0; fi6 < 16; fi6++ {
                var v7 float32 = 0.0
                _ = fi6
                _ = v7
            }
        default:
            v8 := GenS1{F0: -9223372036854775808, F1: 2}
            _ = v8
        }
    case "":
        var v9 int8 = 0
        _ = v9
    default:
        v10 := GenS2{F0: 1}
        _ = v10
    }
    var v11 int16 = -948
    switch "" {
    case "\n\t":
        v11 -= 1
    case "":
        _ = len(v5)
    case "foo bar":
        v12 := GenS3{F0: -854, F1: ((3.14159 * (0.0 - 0.0)) <= -2.25)}
        _ = v12
    default:
        v3 = append(v3, ("\n\t" + "foo bar"))
    }
    var v13 float64 = (0.1 + (1.5 + -2.25))
    var v14 uint32 = 416
    if (!true) {
        switch "a" {
        case "hello":
            if (4294967295 >= 4294967295) {
                v14 += 468
            }
        case "foo bar":
            v15 := GenS3{F0: -9223372036854775808, F1: false}
            _ = v15
        }
        for fi16 := 0; fi16 < 15; fi16++ {
            v4 = (fi16 ^ v4)
            var v17 int16 = v11
            v18 := GenS1{F0: -736, F1: 4294967295}
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            v19 := v1[(v14 >> ((65535) & 63))]
            var v20 int8 = 3
            _ = v19
            _ = v20
        }
    } else {
        v1[(^(v14))] = true
        if false {
            v14 = 121
            v21 := GenS1{F0: int64(v2), F1: 42}
            _ = v21
        } else {
            if false {
                var v22 int32 = 2147483647
                v22 += v22
                _ = v22
            } else {
                var v23 int8 = (v2 | v2)
                _ = v23
            }
            if false {
                if ((0.1 != (1.5 - -2.25)) && false) {
                    v24 := GenS4{F0: "", F1: 42, F2: "hello"}
                    for fi25 := 0; fi25 < 6; fi25++ {
                        if (false || false) {
                            var v26 uint64 = 0
                            v14 = v14
                            for fi27 := 0; fi27 < 4; fi27++ {
                                switch fi27 {
                                case 0:
                                    v13 = float64((0.1 + (1e10 * (((1.5 * (-2.25 / float32(v13))) / 3.14159) * 0.0))))
                                case 4294967295:
                                    switch "a" {
                                    case "":
                                        if false {
                                            v14 = v14
                                        } else {
                                            v1[(v14 - v14)] = (false && (("a" + "a") >= "foo bar"))
                                            if false {
                                                v4 += 128
                                            }
                                        }
                                        v4 = v4
                                        v1[0] = false
                                    case "hello":
                                        switch ((("" + ("hello" + "a")) + ("" + "\n\t")) + ("a" + "\n\t")) {
                                        case "foo bar":
                                            v4 = fi27
                                        }
                                        var v28 uint8 = 0
                                        _ = v28
                                    }
                                }
                                v1[(v14 << ((18446744073709551615) & 63))] = (true && true)
                                _ = fi27
                            }
                            _ = v26
                        } else {
                            for fi29 := 0; fi29 < 10; fi29++ {
                                var v30 int32 = -2147483648
                                var v31 uint32 = (v14 - v14)
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                        }
                        var v32 float64 = ((v13 - (v13 * v13)) + v13)
                        fmt.Println(v14, fi25, fi25)
                        _ = fi25
                        _ = v32
                    }
                    v14 = v14
                    _ = v24
                }
                switch ("" + ("hello" + "\n\t")) {
                case "foo bar":
                    v33 := GenS3{F0: 615, F1: (((v4 | 32768) <= v4) && (1 == 0))}
                    _ = v33
                case "a":
                    v4 -= (v4 % (((v4 & (v4 + 32768))) | 1))
                case "":
                    switch (v4 << ((686) & 63)) {
                    case 1:
                        var v34 uint64 = 42
                        _ = v34
                    case 32767:
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case 357:
                        v35 := GenS3{F0: -114, F1: ((!true) || (!true))}
                        _ = v35
                    }
                default:
                    v36 := GenS5{F0: -514, F1: v11}
                    _ = v36
                }
            } else {
                v5 = append(v5, v14)
            }
            if false {
                v4 += (v4 - (^(v4)))
            } else {
                v2 = (-(v2))
                v37 := GenS6{F0: 9223372036854775807, F1: 128}
                _ = v37
            }
        }
    }
    if false {
        for fi38 := 0; fi38 < 6; fi38++ {
            for fi39 := 0; fi39 < 17; fi39++ {
                if (!true) {
                    if true {
                        v40 := GenS1{F0: -264, F1: 65535}
                        var v41 int64 = 0
                        v42 := GenS3{F0: (v41 - int64(v14)), F1: false}
                        _ = v40
                        _ = v41
                        _ = v42
                    } else {
                        v13 += (v13 + 0.0)
                        v2 += (v2 / ((v2) | 1))
                        switch fi38 {
                        case 298:
                            v43 := GenS7{F0: 725, F1: 1, F2: true}
                            _ = v43
                        default:
                            v44 := GenS8{F0: ""}
                            _ = v44
                        }
                    }
                    v45 := GenS8{F0: (("" + (("" + "\n\t") + "foo bar")) + (("hello" + "hello") + "a"))}
                    _ = v45
                }
                _ = fi39
            }
            v2 = -111
            var v46 int32 = -1
            _ = fi38
            _ = v46
        }
        fmt.Println(v4, v4, v4)
        var v47 uint32 = v14
        _ = v47
    } else {
        var v48 uint64 = 2
        _ = v48
    }
    v11 = (-(v11))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
