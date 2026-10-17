package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32; F1 uint; F2 int32 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16; F1 uint16; F2 int64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint64; F1 uint }
type GenS10 struct { F0 int8; F1 string }
type GenS11 struct { F0 uint32 }

func main() {
    var v1 uint64 = 42
    var v2 float64 = -2.25
    var v3 int = 1
    var v4 float64 = 1e10
    v5 := make(map[uint32]int)
    v6 := GenS1{F0: 42, F1: 59}
    fmt.Println(v1)
    v7 := GenS1{F0: 0, F1: 255}
    for fi8 := 0; fi8 < 16; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            v10 := []int{(v3 - (^(v3))), fi9, 42, (fi8 & 0)}
            v1 -= v1
            var v11 bool = (v4 <= ((v2 - (v2 / 1.0)) + float64(float32(1e10))))
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 int16 = -481
        _ = fi8
        _ = v12
    }
    var v13 int64 = 1
    if true {
        var v14 int16 = 1
        v14 += v14
        v5[2147483648] = v3
        _ = v14
    } else {
        for fi15 := 0; fi15 < 16; fi15++ {
            for fi16 := 0; fi16 < 8; fi16++ {
                fmt.Println(v1, v3)
                v17 := GenS1{F0: 127, F1: 1}
                if (1 <= 318) {
                    v2 += v2
                    fmt.Println(v13, v3)
                    switch "" {
                    case "":
                        switch "\n\t" {
                        case "\n\t":
                            v1 = 1
                        default:
                            var v18 int64 = (v13 | 9223372036854775807)
                            switch (fi16 / ((fi15) | 1)) {
                            case -1:
                                fmt.Println(v3)
                            default:
                                v19 := GenS1{F0: -120, F1: 1}
                                v20 := GenS2{F0: 32767, F1: 72}
                                _ = v19
                                _ = v20
                            }
                            v21 := GenS3{F0: -452, F1: 723, F2: 0}
                            if (4294967295 != 132) {
                                switch fi16 {
                                case -1:
                                    var v22 bool = false
                                    _ = v22
                                case 32768:
                                    v23 := GenS1{F0: 0, F1: 109}
                                    _ = v23
                                case 4294967295:
                                    var v24 uint8 = 128
                                    _ = v24
                                default:
                                    var v25 uint16 = 32767
                                    _ = v25
                                }
                                v26 := GenS1{F0: -99, F1: 2}
                                switch (fi16 - v3) {
                                case -483:
                                    switch "\n\t" {
                                    case "":
                                        fmt.Println(fi15, v2, fi16)
                                    case "\n\t":
                                        v1 += (^(v1))
                                    case "foo bar":
                                        v5[954] = 4294967295
                                        var v27 uint8 = 163
                                        _ = v27
                                    }
                                case -702:
                                    v28 := GenS2{F0: 504, F1: -5}
                                    _ = v28
                                case 128:
                                    v13 = v18
                                default:
                                    var v29 uint64 = 4294967295
                                    var v30 string = "foo bar"
                                    v31 := GenS4{F0: (v1 << ((931) & 63))}
                                    _ = v29
                                    _ = v30
                                    _ = v31
                                }
                                _ = v26
                            } else {
                                fmt.Println(v4, v2)
                                var v32 uint32 = 2147483648
                                switch (v3 & fi15) {
                                case 42:
                                    v33 := GenS5{F0: (!((true && (!false)) && (32767 == 65535)))}
                                    _ = v33
                                case -2147483648:
                                    v18 += v13
                                case -287:
                                    var v34 int32 = 413
                                    _ = v34
                                default:
                                    v35 := GenS6{F0: 594, F1: 799, F2: v18}
                                    _ = v35
                                }
                                _ = v32
                            }
                            _ = v18
                            _ = v21
                        }
                    case "a":
                        v36 := GenS7{F0: ""}
                        v13 -= -223
                        _ = v36
                    case "\n\t":
                        if (!(-1 != 0)) {
                            v1 = v1
                            fmt.Println(fi15, v13, fi15)
                            var v37 float64 = ((float64(1.5) + float64(0.0)) * v2)
                            v3 += (fi15 & (fi16 >> ((1) & 63)))
                            var v38 int = (fi16 ^ fi16)
                            _ = v37
                            _ = v38
                        } else {
                            v39 := []int{(fi15 & (fi15 / ((4294967295) | 1))), 32767, (v3 & (fi16 * (^(fi15))))}
                            var v40 string = ""
                            _ = v39
                            _ = v40
                        }
                    }
                    var v41 uint16 = 2
                    _ = v41
                }
                _ = fi16
                _ = v17
            }
            _ = fi15
        }
        switch "" {
        case "hello":
            var v42 float32 = 1e10
            _ = v42
        default:
            var v43 float64 = 0.0
            _ = v43
        }
        for fi44 := 0; fi44 < 8; fi44++ {
            if (!(!false)) {
                if false {
                    fmt.Println(v4)
                    v45 := []int{v3}
                    _ = v45
                } else {
                    var v46 int8 = 100
                    v1 -= v1
                    _ = v46
                }
                for fi47 := 0; fi47 < 17; fi47++ {
                    var v48 float32 = 0.1
                    if true {
                        switch ("hello" + ("foo bar" + "")) {
                        case "foo bar":
                            if true {
                                v49 := v5[760]
                                _ = v49
                            }
                        default:
                            if false {
                                v4 = float64(v48)
                                var v50 int64 = v13
                                var v51 int8 = -128
                                _ = v50
                                _ = v51
                            } else {
                                fmt.Println(v3, v3, v13)
                            }
                        }
                        v52 := GenS8{F0: 0}
                        _ = v52
                    }
                    _ = fi47
                    _ = v48
                }
            }
            v53 := GenS9{F0: 2, F1: 42}
            v1 -= v1
            _ = fi44
            _ = v53
        }
    }
    v4 = v2
    for fi54 := 0; fi54 < 8; fi54++ {
        switch fi54 {
        case -969:
            v4 = v4
        case 2147483647:
            v55 := v5[1]
            _ = v55
        }
        var v56 uint64 = v1
        var v57 bool = true
        _ = fi54
        _ = v56
        _ = v57
    }
    if true {
        v58 := GenS9{F0: v1, F1: 96}
        _ = v58
    } else {
        var v59 uint = 18446744073709551615
        _ = v59
    }
    v60 := []int{-2147483648}
    v61 := GenS10{F0: 0, F1: ((("" + ("" + (((("hello" + ("hello" + "a")) + "\n\t") + "hello") + (("" + (((("foo bar" + "a") + ((("" + "hello") + "\n\t") + "foo bar")) + "foo bar") + ((("a" + ("hello" + (("a" + (("foo bar" + "foo bar") + "hello")) + "hello"))) + "hello") + ("a" + "hello")))) + "a")))) + ("\n\t" + (("" + "") + ""))) + "foo bar")}
    for fi62 := 0; fi62 < 13; fi62++ {
        v2 = v4
        _ = len(v60)
        v4 = v2
        _ = fi62
    }
    fmt.Println(v4, v4, v2)
    v63 := GenS11{F0: 0}
    v64 := GenS9{F0: 0, F1: 704}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v60:", len(v60), v60)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v60
    _ = v61
    _ = v63
    _ = v64
    _ = v7
}
