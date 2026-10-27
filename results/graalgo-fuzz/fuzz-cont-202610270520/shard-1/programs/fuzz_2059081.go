package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 int; F2 string }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 string; F1 int16; F2 string }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint16) uint8 {
    v4 := []int{1, 1}
    p0 -= p0
    _ = p0
    _ = v4
    return 42
}

func main() {
    var v1 int16 = 42
    var v2 string = (((("" + "foo bar") + "hello") + "\n\t") + "")
    var v3 float32 = (0.1 + -2.25)
    v5 := GenS1{F0: 39}
    if (255 == genFn0(0)) {
        fmt.Println(v2, v2)
        fmt.Println(v1, v3)
        v2 = "foo bar"
    }
    for fi6 := 0; fi6 < 5; fi6++ {
        fmt.Println(v1, v3, v2)
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int16 = (v1 + (v1 >> ((255) & 63)))
            fmt.Println(fi6, v1)
            _ = fi7
            _ = v8
        }
        v1 = 1
        _ = fi6
    }
    switch v2 {
    case "\n\t":
        var v9 int = 9223372036854775807
        _ = v9
    default:
        var v10 uint64 = 42
        _ = v10
    }
    var v11 uint32 = 37
    v12 := []int{174}
    if false {
        for fi13 := 0; fi13 < 7; fi13++ {
            if false {
                var v14 float64 = (((0.1 / 1.5) + (1e10 + (0.1 / 1.0))) + float64(((v3 + (v3 + (float32(1.5) - v3))) / (((-2.25 / v3) - float32(1e10)) / float32(1.5)))))
                _ = v14
            }
            if len(v12) > 0 {
                v12[0] = (-(fi13))
            }
            v15 := GenS1{F0: 309}
            _ = fi13
            _ = v15
        }
        switch v2 {
        case "a":
            for fi16 := 0; fi16 < 7; fi16++ {
                v3 = v3
                _ = fi16
            }
        case "foo bar":
            fmt.Println(v1, v3)
        case "\n\t":
            for fi17 := 0; fi17 < 13; fi17++ {
                for fi18 := 0; fi18 < 3; fi18++ {
                    _ = len(v12)
                    if (v2 > v2) {
                        if ((v1 + 1) >= v1) {
                            v1 = (^(v1))
                            switch "a" {
                            case "\n\t":
                                switch fi17 {
                                case 32767:
                                    if false {
                                        fmt.Println(v1, fi18, v11)
                                    }
                                case 200:
                                    var v19 int16 = v1
                                    _ = v19
                                default:
                                    fmt.Println(fi17, v3, v11)
                                }
                            default:
                                switch -400 {
                                case 1:
                                    var v20 int = fi18
                                    _ = v20
                                case -215:
                                    v12 = append(v12, fi17)
                                case 128:
                                    if (!((127 >= 255) || (!(135 > 4294967295)))) {
                                        v21 := GenS1{F0: uint16(v11)}
                                        var v22 uint8 = 206
                                        var v23 int64 = 0
                                        _ = v21
                                        _ = v22
                                        _ = v23
                                    }
                                }
                            }
                            if false {
                                v24 := GenS1{F0: 65535}
                                var v25 int16 = (v1 - (^(v1)))
                                _ = v24
                                _ = v25
                            } else {
                                if true {
                                    switch "" {
                                    case "hello":
                                        fmt.Println(fi17, v1)
                                    default:
                                        fmt.Println(v1, v1)
                                    }
                                }
                                _ = len(v12)
                                v11 = 4294967295
                            }
                            v2 = v2
                            switch "\n\t" {
                            case "\n\t":
                                if len(v12) >= 2 {
                                    v12 = v12[1:len(v12)]
                                }
                            default:
                                v26 := GenS2{F0: 100, F1: 42}
                                _ = v26
                            }
                        }
                    }
                    _ = fi18
                }
                _ = len(v12)
                for fi27 := 0; fi27 < 1; fi27++ {
                    var v28 int32 = 0
                    if true {
                        var v29 string = (v2 + "\n\t")
                        v28 -= (v28 / ((v28) | 1))
                        _ = v29
                    } else {
                        fmt.Println(v11)
                        fmt.Println(v28, v3, v11)
                    }
                    _ = fi27
                    _ = v28
                }
                _ = fi17
            }
        }
    } else {
        v3 = (((v3 * v3) - 1e10) * (v3 + 1.5))
        for fi30 := 0; fi30 < 7; fi30++ {
            _ = len(v12)
            switch (fi30 | (fi30 % ((fi30) | 1))) {
            case 42:
                v11 -= v11
            case -9223372036854775808:
                v1 += v1
            default:
                v31 := GenS3{F0: 128, F1: fi30, F2: "a"}
                _ = v31
            }
            _ = fi30
        }
    }
    for fi32 := 0; fi32 < 15; fi32++ {
        v2 = (v2 + v2)
        _ = fi32
    }
    var v33 string = v2
    for fi34 := 0; fi34 < 18; fi34++ {
        if (((-115 < -128) || true) && true) {
            var v35 int64 = 42
            _ = v35
        } else {
            if false {
                v33 = "hello"
                v11 -= 0
                if true {
                    if (4294967295 >= 2) {
                        v3 += (float32(-2.25) + 1.5)
                        var v36 uint32 = v11
                        _ = v36
                    } else {
                        v12 = append(v12, fi34)
                        if (true && (!false)) {
                            switch ("" + v2) {
                            case "a":
                                var v37 int8 = -110
                                _ = v37
                            default:
                                if false {
                                    var v38 uint64 = 538
                                    _ = v38
                                } else {
                                    v11 = (^(v11))
                                    v39 := GenS4{F0: v11, F1: (v1 % ((v1) | 1))}
                                    _ = v39
                                }
                            }
                            v12 = append(v12, (-(fi34)))
                        } else {
                            var v40 uint32 = v11
                            v33 = ""
                            switch 32768 {
                            case 1:
                                v41 := GenS5{F0: "\n\t", F1: (v1 & 1000), F2: "\n\t"}
                                _ = v41
                            case 2147483647:
                                v1 = -1
                            default:
                                for fi42 := 0; fi42 < 3; fi42++ {
                                    v43 := GenS4{F0: 2147483648, F1: v1}
                                    _ = fi42
                                    _ = v43
                                }
                            }
                            _ = v40
                        }
                    }
                }
            } else {
                v44 := GenS6{F0: -62}
                switch 4294967295 {
                case -2147483648:
                    fmt.Println(v2, v3)
                case 2147483647:
                    for fi45 := 0; fi45 < 10; fi45++ {
                        var v46 int8 = 36
                        var v47 int64 = -1
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                }
                v3 -= v3
                _ = v44
            }
            if false {
                v12 = append(v12, fi34)
                var v48 uint32 = 1
                v49 := GenS7{F0: (true && (!((32767 <= 32767) || true))), F1: 42}
                _ = v48
                _ = v49
            }
        }
        _ = fi34
    }
    if true {
        fmt.Println(v11, v33, v1)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v33
    _ = v5
}
