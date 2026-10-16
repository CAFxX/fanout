package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 bool }

func main() {
    var v1 uint16 = 2
    var v2 string = "hello"
    var v3 uint32 = 0
    var v4 string = (v2 + ((((v2 + v2) + ((v2 + v2) + "a")) + ((v2 + "\n\t") + "\n\t")) + v2))
    var v5 int8 = 5
    v6 := make(map[int]uint32)
    v5 += v5
    for fi7 := 0; fi7 < 17; fi7++ {
        switch ("\n\t" + v2) {
        case "hello":
            v8 := GenS1{F0: v1, F1: 42}
            _ = v8
        case "\n\t":
            v9 := []int{}
            _ = v9
        case "foo bar":
            _ = len(v6)
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            v3 = v3
            if true {
                var v11 int = (fi10 ^ (fi7 * fi10))
                var v12 int = (fi7 % (((fi10 >> ((692) & 63))) | 1))
                switch (^(v12)) {
                case -2147483648:
                    if (128 > fi7) {
                        var v13 int32 = -362
                        v14 := []int{}
                        switch (fi10 + v12) {
                        case -128:
                            var v15 uint8 = 255
                            _ = v15
                        }
                        var v16 uint64 = 18446744073709551615
                        switch v2 {
                        case "foo bar":
                            v11 = -9223372036854775808
                        }
                        if ((v1 + 42) >= 0) {
                            v4 = ((v4 + ("foo bar" + v2)) + ("" + v4))
                        }
                        _ = v13
                        _ = v14
                        _ = v16
                    }
                case -1:
                    v4 = (v2 + "hello")
                default:
                    v5 -= v5
                }
                var v17 uint16 = 42
                _ = v11
                _ = v12
                _ = v17
            } else {
                var v18 bool = true
                var v19 float64 = 0.0
                _ = v18
                _ = v19
            }
            _ = fi10
        }
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 uint = 1
            var v22 float64 = 3.14159
            switch v4 {
            case "hello":
                fmt.Println(fi20, v2, v22)
            case "foo bar":
                switch "hello" {
                case "foo bar":
                    v2 = v4
                case "hello":
                    if false {
                        v23 := GenS1{F0: (v1 / (((v1 << ((v21) & 63))) | 1)), F1: 1000}
                        var v24 uint32 = v3
                        var v25 float64 = ((v22 / ((-2.25 - v22) / float64((float32(v22) - (1.5 / float32(v22)))))) - (1.5 - v22))
                        v5 += 100
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                case "\n\t":
                    v5 = (v5 >> (((v21 / (((v21 / ((v21) | 1))) | 1))) & 63))
                }
            default:
                var v26 float64 = 0.1
                _ = v26
            }
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi7
    }
    v27 := GenS1{F0: (v1 | (v1 | 32767)), F1: -233}
    var v28 uint64 = 4294967295
    var v29 uint8 = 172
    for fi30 := 0; fi30 < 13; fi30++ {
        if true {
            if false {
                if true {
                    v31 := []int{}
                    var v32 bool = (-947 <= 0)
                    _ = v31
                    _ = v32
                } else {
                    v29 = (v29 + 255)
                    var v33 uint64 = 0
                    _ = v33
                }
                var v34 float32 = -2.25
                _ = v34
            }
        }
        v3 -= (^(v3))
        for fi35 := 0; fi35 < 7; fi35++ {
            v5 -= (v5 ^ v5)
            _ = fi35
        }
        _ = fi30
    }
    v36 := GenS2{F0: v4, F1: 42}
    for fi37 := 0; fi37 < 2; fi37++ {
        v2 = "foo bar"
        for fi38 := 0; fi38 < 5; fi38++ {
            v39 := []int{32767, (fi37 << ((2) & 63)), fi38}
            if (false && false) {
                v28 = v28
                var v40 int64 = 0
                if false {
                    v39 = append(v39, (fi37 >> ((4294967295) & 63)))
                    if len(v39) >= 2 {
                        v39 = v39[1:len(v39)]
                    }
                    v1 -= (v1 * uint16(v3))
                } else {
                    if false {
                        v41 := GenS3{F0: (42 > -2147483648), F1: v29}
                        switch 128 {
                        case -535:
                            if false {
                                v29 += 127
                            }
                        case -9223372036854775808:
                            switch "a" {
                            case "a":
                                v39 = append(v39, (fi37 % (((fi38 ^ (-(fi38)))) | 1)))
                            case "\n\t":
                                var v42 uint32 = 42
                                if ((true && (!true)) && ((0.0 + float64(((float32(0.1) * (3.14159 / (0.0 - 1e10))) - 0.1))) == (0.0 * ((1.5 / 3.14159) - 3.14159)))) {
                                    v43 := GenS2{F0: (("foo bar" + v4) + (v4 + v2)), F1: 465}
                                    _ = v43
                                }
                                v44 := v6[-983]
                                _ = v42
                                _ = v44
                            case "":
                                fmt.Println(v40)
                            }
                            fmt.Println(v3, v1)
                            if false {
                                var v45 int32 = 663
                                if (491 <= 1000) {
                                    v29 -= (^(v29))
                                    var v46 int16 = 0
                                    v45 = v45
                                    _ = v46
                                } else {
                                    var v47 uint16 = v1
                                    var v48 float64 = -2.25
                                    _ = v47
                                    _ = v48
                                }
                                _ = v45
                            }
                            var v49 uint32 = v3
                            _ = v49
                        case 238:
                            v50 := GenS4{F0: (-32768 < 1)}
                            _ = v50
                        default:
                            v40 += (v40 | v40)
                            if len(v39) > 0 {
                                v39[0] = (-(fi37))
                            }
                        }
                        v51 := v6[976]
                        _ = v41
                        _ = v51
                    }
                }
                _ = v40
            }
            _ = fi38
            _ = v39
        }
        switch ((((v2 + v2) + "foo bar") + v4) + "a") {
        case "\n\t":
            for fi52 := 0; fi52 < 1; fi52++ {
                fmt.Println(v5, v2)
                v28 -= v28
                _ = fi52
            }
        case "":
            fmt.Println(v3, v2)
        default:
            var v53 uint32 = (v3 << ((255) & 63))
            _ = v53
        }
        _ = fi37
    }
    fmt.Println(v4, v3, v28)
    var v54 uint32 = 4294967295
    v54 = v54
    if false {
        var v55 float64 = 1e10
        _ = v55
    } else {
        switch "\n\t" {
        case "foo bar":
            if true {
                var v56 uint8 = 0
                for fi57 := 0; fi57 < 7; fi57++ {
                    fmt.Println(v28)
                    _ = fi57
                }
                var v58 int = -1
                _ = v56
                _ = v58
            } else {
                for fi59 := 0; fi59 < 6; fi59++ {
                    v29 -= 67
                    if true {
                        for fi60 := 0; fi60 < 20; fi60++ {
                            if (-1 < 0) {
                                v61 := GenS1{F0: (v1 | v1), F1: 994}
                                fmt.Println(v3, v29, fi60)
                                fmt.Println(v5, v1)
                                _ = v61
                            }
                            v62 := []int{int(v1)}
                            v63 := GenS3{F0: (int(v1) == (fi60 % ((1) | 1))), F1: 129}
                            _ = fi60
                            _ = v62
                            _ = v63
                        }
                        fmt.Println(v28)
                        if true {
                            switch (fi59 + (-(fi59))) {
                            case -370:
                                var v64 float64 = -2.25
                                _ = v64
                            case -9223372036854775808:
                                for fi65 := 0; fi65 < 13; fi65++ {
                                    fmt.Println(v4, v54, v1)
                                    v3 -= (v3 | v3)
                                    var v66 int = (fi59 + fi59)
                                    _ = fi65
                                    _ = v66
                                }
                            default:
                                fmt.Println(v29, fi59)
                            }
                        } else {
                            v67 := []int{32768}
                            v5 += (v5 >> ((65535) & 63))
                            _ = v67
                        }
                    }
                    _ = fi59
                }
            }
        default:
            _ = len(v6)
        }
    }
    if false {
        fmt.Println(v4, v54, v54)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v54
    _ = v6
}
