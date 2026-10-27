package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint16; F2 uint16 }

func main() {
    var v1 bool = (!false)
    v2 := []string{"\n\t", "hello", "hello"}
    var v3 int16 = -648
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 32767
    v1 = (0 != -845)
    v5 = v5
    for fi6 := 0; fi6 < 10; fi6++ {
        v3 = v3
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 100, F1: (v5 * 65535)}
            _ = fi7
            _ = v8
        }
        for fi9 := 0; fi9 < 8; fi9++ {
            if true {
                if (!(("a" + "") > "a")) {
                    if (v1 || (!true)) {
                        v3 += v3
                        if v1 {
                            v1 = v1
                            if v1 {
                                var v10 int64 = 1
                                v1 = v1
                                _ = v10
                            } else {
                                fmt.Println(fi9, v5)
                            }
                            var v11 uint8 = 1
                            _ = v11
                        } else {
                            var v12 bool = (((!(v1 || (0.0 > 0.1))) && v1) || false)
                            _ = v12
                        }
                    } else {
                        v1 = v1
                        var v13 int64 = 1
                        v14 := GenS2{F0: v13}
                        if len(v2) > 0 {
                            v2[0] = ""
                        }
                        if v1 {
                            fmt.Println(fi9, fi9, fi9)
                        } else {
                            switch fi6 {
                            case 4294967295:
                                v13 = v13
                            default:
                                var v15 uint64 = 4294967295
                                _ = v15
                            }
                            fmt.Println(v5)
                            if (v1 || ((v5 ^ (v5 % (((v5 >> ((814) & 63))) | 1))) > (v5 % ((v5) | 1)))) {
                                switch "\n\t" {
                                case "":
                                    v13 = (v13 & (v13 & 9223372036854775807))
                                case "a":
                                    if v1 {
                                        var v16 uint = 204
                                        v17 := v4[2147483647]
                                        v18 := GenS1{F0: 59, F1: 500}
                                        _ = v16
                                        _ = v17
                                        _ = v18
                                    } else {
                                        switch "foo bar" {
                                        case "a":
                                            fmt.Println(fi6, v3)
                                        case "":
                                            var v19 int8 = 42
                                            _ = v19
                                        default:
                                            v20 := GenS2{F0: (v13 / ((1) | 1))}
                                            _ = v20
                                        }
                                        switch (("\n\t" + ("\n\t" + "")) + "\n\t") {
                                        case "":
                                            if (uint32(v5) <= 2) {
                                                var v21 bool = v1
                                                _ = v21
                                            }
                                        case "\n\t":
                                            switch ("hello" + "hello") {
                                            case "hello":
                                                _ = len(v4)
                                            default:
                                                v5 += (v5 & (v5 + v5))
                                            }
                                        }
                                    }
                                case "\n\t":
                                    if (!(-32768 <= (v3 % ((715) | 1)))) {
                                        switch int(v5) {
                                        case 32768:
                                            if (280 <= 2) {
                                                v22 := GenS1{F0: 127, F1: v5}
                                                _ = v22
                                            }
                                        case 496:
                                            v1 = (!(1 <= 2))
                                        default:
                                            v2 = append(v2, "hello")
                                        }
                                    } else {
                                        var v23 uint64 = 0
                                        var v24 uint64 = (^(v23))
                                        _ = v23
                                        _ = v24
                                    }
                                default:
                                    if false {
                                        var v25 int8 = 127
                                        v26 := GenS2{F0: (v13 << ((633) & 63))}
                                        v27 := GenS1{F0: v25, F1: (v5 * v5)}
                                        _ = v25
                                        _ = v26
                                        _ = v27
                                    } else {
                                        fmt.Println(v13)
                                        v5 = 2
                                    }
                                    fmt.Println(fi9, fi9, fi6)
                                    var v28 int8 = 38
                                    _ = v28
                                }
                                delete(v4, 641)
                                var v29 float64 = 0.0
                                _ = v29
                            } else {
                                var v30 uint32 = 0
                                var v31 int8 = -77
                                var v32 float64 = (((0.1 * 0.0) - 3.14159) - 3.14159)
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            switch ("a" + ("" + (("hello" + "a") + "foo bar"))) {
                            case "hello":
                                v33 := GenS2{F0: v13}
                                _ = v33
                            default:
                                var v34 float64 = -2.25
                                _ = v34
                            }
                            v35 := GenS3{F0: fi9}
                            v4[1] = 840
                            var v36 uint16 = v5
                            _ = v35
                            _ = v36
                        }
                        _ = v13
                        _ = v14
                    }
                    var v37 float32 = 1e10
                    v37 = 1.5
                    _ = v37
                }
                if (!v1) {
                    var v38 int64 = 53
                    _ = v38
                }
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            } else {
                _ = len(v4)
                switch "" {
                case "foo bar":
                    v3 = -407
                case "":
                    var v39 int16 = (v3 >> ((0) & 63))
                    _ = v39
                }
                if v1 {
                    v3 = -1
                    if v1 {
                        v40 := GenS4{F0: 0, F1: v5, F2: (v5 + v5)}
                        _ = len(v2)
                        var v41 int8 = 127
                        _ = v40
                        _ = v41
                    } else {
                        if v1 {
                            var v42 float32 = (1.5 + 1e10)
                            var v43 uint16 = (v5 + 1)
                            _ = v42
                            _ = v43
                        } else {
                            var v44 uint = 2
                            _ = v44
                        }
                        v45 := GenS2{F0: 661}
                        _ = v45
                    }
                    v46 := GenS3{F0: 217}
                    _ = v46
                }
                v5 = 32767
            }
            var v47 int8 = -128
            _ = fi9
            _ = v47
        }
        _ = fi6
    }
    v48 := GenS1{F0: 1, F1: v5}
    var v49 int64 = -659
    var v50 float32 = (1e10 / (0.1 + float32(1.5)))
    v50 -= ((v50 + -2.25) / 1.0)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
