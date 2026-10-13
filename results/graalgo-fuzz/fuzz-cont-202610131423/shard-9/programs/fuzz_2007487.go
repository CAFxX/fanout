package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint64 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }

func main() {
    var v1 int = 4294967295
    var v2 string = "hello"
    v3 := []string{}
    var v4 int64 = -9223372036854775808
    v1 = (v1 | (v1 + (v1 & (v1 | 32768))))
    v3 = append(v3, (v2 + (v2 + "\n\t")))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch (v1 >> ((254) & 63)) {
    case 0:
        var v5 int = 0
        _ = v5
    case -2147483648:
        fmt.Println(v1, v4)
    case 9223372036854775807:
        if ((v4 - v4) == v4) {
            var v6 int = v1
            switch -624 {
            case 4294967295:
                var v7 uint64 = 42
                _ = v7
            default:
                fmt.Println(v6, v1, v4)
            }
            _ = v6
        }
    default:
        fmt.Println(v4)
    }
    var v8 int8 = 0
    v1 = v1
    for fi9 := 0; fi9 < 2; fi9++ {
        v2 = v2
        if true {
            var v10 float32 = 0.1
            _ = v10
        }
        for fi11 := 0; fi11 < 6; fi11++ {
            if false {
                if false {
                    v3 = append(v3, (v2 + "a"))
                } else {
                    var v12 int = fi9
                    v13 := GenS1{F0: (v2 + v2)}
                    if ((-32768 < 1) && ((false || true) || true)) {
                        switch ((v2 + (v2 + "\n\t")) + (((v2 + "foo bar") + v2) + v2)) {
                        case "":
                            v14 := GenS1{F0: v2}
                            _ = v14
                        case "hello":
                            var v15 float32 = -2.25
                            var v16 float64 = 0.1
                            var v17 int16 = 32767
                            var v18 float64 = float64(float32(v16))
                            _ = v15
                            _ = v16
                            _ = v17
                            _ = v18
                        default:
                            if (307 >= 255) {
                                var v19 bool = (1 == 42)
                                fmt.Println(v12)
                                v1 -= (^(v1))
                                _ = v19
                            }
                        }
                        var v20 int = (fi9 >> ((563) & 63))
                        switch v2 {
                        case "foo bar":
                            v2 = v2
                        }
                        v21 := GenS2{F0: 32767, F1: 669}
                        v22 := GenS2{F0: 0, F1: 1}
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    var v23 int32 = 0
                    var v24 float64 = -2.25
                    v1 -= (fi11 >> ((950) & 63))
                    _ = v12
                    _ = v13
                    _ = v23
                    _ = v24
                }
            }
            var v25 float32 = 0.1
            var v26 int16 = -1
            _ = fi11
            _ = v25
            _ = v26
        }
        _ = fi9
    }
    var v27 bool = (!((1.5 + ((0.0 + 1e10) / 1.0)) == 0.0))
    for fi28 := 0; fi28 < 11; fi28++ {
        v27 = true
        for fi29 := 0; fi29 < 10; fi29++ {
            var v30 float32 = 0.0
            v3 = append(v3, "")
            v31 := GenS2{F0: 32767, F1: 18446744073709551615}
            var v32 uint32 = 65535
            switch v2 {
            case "hello":
                fmt.Println(v32, v2)
            case "foo bar":
                switch fi29 {
                case 4294967295:
                    v33 := GenS1{F0: "foo bar"}
                    _ = v33
                case 127:
                    v1 = 710
                    _ = len(v3)
                    switch v2 {
                    case "hello":
                        switch v2 {
                        case "foo bar":
                            v34 := GenS2{F0: -32768, F1: 18446744073709551615}
                            _ = v34
                        case "":
                            if ((v27 || (v4 != v4)) && (13 <= 2)) {
                                var v35 int16 = -32768
                                var v36 int64 = 0
                                _ = v35
                                _ = v36
                            }
                        }
                    case "\n\t":
                        var v37 uint32 = (v32 + v32)
                        _ = v37
                    case "a":
                        v38 := GenS3{F0: 1, F1: (v1 * (v1 + (fi29 * 9223372036854775807)))}
                        _ = v38
                    }
                case -9223372036854775808:
                    _ = len(v3)
                default:
                    v39 := GenS3{F0: 467, F1: (fi28 / (((-(v1))) | 1))}
                    _ = v39
                }
            case "a":
                var v40 int32 = -2147483648
                _ = v40
            default:
                var v41 int64 = 0
                _ = v41
            }
            _ = fi29
            _ = v30
            _ = v31
            _ = v32
        }
        _ = fi28
    }
    if v27 {
        v42 := GenS1{F0: "\n\t"}
        if true {
            v1 = v1
            for fi43 := 0; fi43 < 3; fi43++ {
                for fi44 := 0; fi44 < 13; fi44++ {
                    v4 = (v4 | v4)
                    switch v2 {
                    case "hello":
                        v45 := GenS3{F0: 4294967295, F1: fi44}
                        _ = v45
                    case "foo bar":
                        var v46 int64 = 727
                        _ = v46
                    case "\n\t":
                        var v47 uint64 = 0
                        _ = v47
                    }
                    _ = fi44
                }
                for fi48 := 0; fi48 < 20; fi48++ {
                    if len(v3) > 0 {
                        v3[0] = v2
                    }
                    var v49 uint8 = 2
                    _ = fi48
                    _ = v49
                }
                v27 = (v27 || (v27 || (2 >= 1)))
                _ = fi43
            }
            _ = len(v3)
        }
        if (97 < 2) {
            var v50 int = (-(v1))
            for fi51 := 0; fi51 < 14; fi51++ {
                v2 = (v2 + v2)
                _ = fi51
            }
            var v52 uint64 = 0
            _ = v50
            _ = v52
        }
        _ = v42
    }
    v53 := GenS4{F0: 4294967295, F1: v4, F2: 255}
    fmt.Println(v8)
    switch v2 {
    case "":
        for fi54 := 0; fi54 < 8; fi54++ {
            fmt.Println(v8, v27, v1)
            _ = fi54
        }
    case "hello":
        if (255 >= 2) {
            var v55 float64 = 3.14159
            _ = v55
        } else {
            v2 = v2
            _ = len(v3)
        }
    }
    for fi56 := 0; fi56 < 12; fi56++ {
        v57 := GenS2{F0: -687, F1: 175}
        var v58 int32 = -2147483648
        var v59 float32 = 3.14159
        _ = fi56
        _ = v57
        _ = v58
        _ = v59
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi60 := 0; fi60 < 18; fi60++ {
        for fi61 := 0; fi61 < 6; fi61++ {
            var v62 float64 = (((0.1 / 0.1) * 3.14159) * 0.1)
            _ = fi61
            _ = v62
        }
        v8 -= -128
        var v63 int = 877
        _ = fi60
        _ = v63
    }
    var v64 float32 = 0.0
    v27 = (float32((3.14159 * 1.5)) <= ((0.1 * v64) + v64))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v64:", v64)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v53
    _ = v64
    _ = v8
}
