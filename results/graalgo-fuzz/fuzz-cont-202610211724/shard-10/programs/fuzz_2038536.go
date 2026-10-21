package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16; F1 uint16; F2 int }
type GenS5 struct { F0 string; F1 string }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = -176
    v3 := []uint32{}
    var v4 string = ("a" + "foo bar")
    v5 := []uint32{2147483647, 1, 2147483648}
    fmt.Println(v4, v4)
    for fi6 := 0; fi6 < 7; fi6++ {
        fmt.Println(v2, fi6)
        if (!(true && ((!false) && true))) {
            for fi7 := 0; fi7 < 11; fi7++ {
                switch (fi6 - fi7) {
                case 2147483647:
                    fmt.Println(fi6, fi7, v2)
                case 128:
                    var v8 int8 = 100
                    _ = v8
                default:
                    v4 = (("" + v4) + "a")
                }
                _ = fi7
            }
        }
        delete(v1, fi6)
        _ = fi6
    }
    v2 = (v2 ^ -1)
    for fi9 := 0; fi9 < 2; fi9++ {
        fmt.Println(v4, v2)
        switch 128 {
        case 127:
            v1[fi9] = ("\n\t" + v4)
        default:
            for fi10 := 0; fi10 < 4; fi10++ {
                fmt.Println(fi9, fi10)
                v2 += v2
                _ = fi10
            }
        }
        var v11 uint32 = 995
        _ = fi9
        _ = v11
    }
    fmt.Println(v2, v4)
    var v12 string = (v4 + "\n\t")
    for fi13 := 0; fi13 < 1; fi13++ {
        fmt.Println(v12, fi13, fi13)
        var v14 uint64 = 599
        if false {
            v4 = ""
            var v15 uint8 = 128
            _ = v15
        }
        _ = fi13
        _ = v14
    }
    if true {
        if len(v5) > 0 {
            v5[0] = 2
        }
    }
    if (1 != 255) {
        for fi16 := 0; fi16 < 2; fi16++ {
            fmt.Println(fi16, v12, v2)
            _ = fi16
        }
    }
    v17 := GenS1{F0: 39, F1: 2147483647}
    var v18 int16 = 401
    fmt.Println(v12, v4)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v19 := GenS2{F0: "a", F1: (-1 <= 7)}
    var v20 uint32 = 2147483647
    v21 := v1[-883]
    if ((-1 != -2147483648) && true) {
        fmt.Println(v20, v12, v4)
        v22 := GenS3{F0: v21, F1: 0}
        if ((("a" + "") + v12) <= (v4 + v12)) {
            var v23 uint16 = 65535
            var v24 uint8 = 2
            if false {
                if (!false) {
                    if (!(0 > (v23 * v23))) {
                        for fi25 := 0; fi25 < 17; fi25++ {
                            var v26 uint = 747
                            for fi27 := 0; fi27 < 8; fi27++ {
                                if (892 > 2) {
                                    v12 = "\n\t"
                                } else {
                                    var v28 int = 4294967295
                                    v29 := GenS2{F0: "\n\t", F1: true}
                                    _ = v28
                                    _ = v29
                                }
                                var v30 float64 = 0.1
                                if ((false || (v2 < (v2 >> (((v26 & (v26 - (v26 % (((v26 * (v26 * (v26 | v26)))) | 1))))) & 63)))) || true) {
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                } else {
                                    v21 = "hello"
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v31 uint8 = (^(v24))
                                    _ = v31
                                }
                                var v32 uint64 = 0
                                switch (fi27 + fi25) {
                                case 0:
                                    v21 = "foo bar"
                                case 304:
                                    v26 -= (v26 % ((v26) | 1))
                                }
                                _ = fi27
                                _ = v30
                                _ = v32
                            }
                            _ = fi25
                            _ = v26
                        }
                        switch v4 {
                        case "a":
                            for fi33 := 0; fi33 < 18; fi33++ {
                                fmt.Println(v4, fi33, v4)
                                fmt.Println(v2, v12, v24)
                                if false {
                                    for fi34 := 0; fi34 < 1; fi34++ {
                                        switch 42 {
                                        case 2147483647:
                                            v35 := GenS4{F0: v18, F1: (^(v23)), F2: int(v2)}
                                            _ = v35
                                        case 32768:
                                            switch v21 {
                                            case "foo bar":
                                                var v36 float64 = 0.1
                                                _ = v36
                                            case "\n\t":
                                                fmt.Println(v18, fi33, fi34)
                                            case "hello":
                                                var v37 uint16 = (^(v23))
                                                _ = v37
                                            default:
                                                var v38 float64 = (-2.25 * ((1.5 / ((0.1 - 0.0) / (float64(-2.25) - 0.0))) + (1.5 - (1e10 / 1.5))))
                                                _ = v38
                                            }
                                        }
                                        v24 = (v24 | (v24 + 1))
                                        _ = fi34
                                    }
                                    var v39 uint32 = v20
                                    if (-902 != int32(v18)) {
                                        var v40 float32 = float32(1e10)
                                        _ = v40
                                    } else {
                                        v41 := GenS4{F0: (v18 & (v18 * (v18 - 1))), F1: v23, F2: fi33}
                                        fmt.Println(v20, v2)
                                        for fi42 := 0; fi42 < 10; fi42++ {
                                            v4 = v12
                                            var v43 uint16 = 0
                                            switch 2147483647 {
                                            case -268:
                                                v24 += (v24 % ((v24) | 1))
                                            case 2147483647:
                                                var v44 float32 = float32((float64(((float32(0.0) + (0.0 - 0.0)) + float32(0.0))) * 1.5))
                                                _ = v44
                                            default:
                                                v45 := v1[0]
                                                _ = v45
                                            }
                                            var v46 int16 = (v18 >> ((255) & 63))
                                            _ = fi42
                                            _ = v43
                                            _ = v46
                                        }
                                        _ = v41
                                    }
                                    _ = v39
                                }
                                _ = fi33
                            }
                        case "hello":
                            v4 = ("a" + v21)
                        default:
                            var v47 float64 = 3.14159
                            _ = v47
                        }
                        for fi48 := 0; fi48 < 14; fi48++ {
                            fmt.Println(v23, v20)
                            v49 := GenS4{F0: 0, F1: 65535, F2: fi48}
                            v2 = v2
                            _ = fi48
                            _ = v49
                        }
                    }
                    _ = len(v3)
                    v50 := GenS5{F0: ("\n\t" + v4), F1: ""}
                    _ = v50
                }
                v20 += (^(v20))
            }
            _ = v23
            _ = v24
        }
        _ = v22
    }
    v4 = (v4 + (v21 + ""))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
