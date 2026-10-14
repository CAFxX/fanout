package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 bool; F1 int64; F2 int16 }
type GenS5 struct { F0 int64; F1 int8; F2 uint }
type GenS6 struct { F0 int16; F1 uint16; F2 int64 }
type GenS7 struct { F0 uint16; F1 int }
type GenS8 struct { F0 uint32; F1 uint; F2 uint64 }

func main() {
    v1 := []uint32{2147483647}
    var v2 uint = 583
    var v3 uint8 = 0
    v4 := []string{"foo bar"}
    v5 := make(map[string]int)
    var v6 string = ("foo bar" + ("\n\t" + "\n\t"))
    v2 = 0
    v7 := GenS1{F0: 4294967295}
    fmt.Println(v3)
    var v8 uint64 = 2
    for fi9 := 0; fi9 < 9; fi9++ {
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v2, fi10)
            if true {
                if (((!false) && false) || false) {
                    v6 = "hello"
                    switch fi9 {
                    case 2147483647:
                        v6 = v6
                    default:
                        if ((v8 == 4294967295) || (true || (2 == 906))) {
                            var v11 uint32 = 4294967295
                            var v12 int32 = -156
                            switch v6 {
                            case "foo bar":
                                if false {
                                    if false {
                                        v13 := GenS1{F0: (v8 - v8)}
                                        if (((65535 >= 255) && true) || false) {
                                            var v14 uint8 = 128
                                            _ = v14
                                        } else {
                                            v8 = v8
                                            var v15 uint64 = (v8 % (((^(v8))) | 1))
                                            _ = v15
                                        }
                                        if false {
                                            v16 := GenS1{F0: (v8 << ((v2) & 63))}
                                            v2 -= (v2 / ((v2) | 1))
                                            _ = v16
                                        }
                                        _ = v13
                                    }
                                    v17 := GenS2{F0: 42, F1: -1}
                                    switch fi9 {
                                    case -9223372036854775808:
                                        v18 := GenS3{F0: 127, F1: 12}
                                        _ = v18
                                    case 127:
                                        _ = len(v5)
                                    case 42:
                                        v11 = v11
                                    }
                                    _ = v17
                                }
                            }
                            _ = v11
                            _ = v12
                        } else {
                            switch (fi10 - fi9) {
                            case -1:
                                v3 += 42
                            case 42:
                                var v19 int16 = 846
                                _ = v19
                            case 2147483647:
                                var v20 bool = true
                                _ = v20
                            }
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            v2 += v2
                        }
                    }
                    v3 = v3
                    switch 9223372036854775807 {
                    case 128:
                        var v21 uint16 = 94
                        _ = v21
                    case -2147483648:
                        var v22 uint = (v2 % ((v2) | 1))
                        v23 := GenS3{F0: 255, F1: -473}
                        _ = v22
                        _ = v23
                    default:
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                } else {
                    if (255 >= 1) {
                        if (2 >= v3) {
                            if ((true || false) && (127 != 3)) {
                                v24 := GenS2{F0: 780, F1: 127}
                                switch v6 {
                                case "\n\t":
                                    v5[v6] = fi10
                                default:
                                    v2 -= 44
                                }
                                _ = v24
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v2 += (v2 >> ((uint(v8)) & 63))
                            }
                            fmt.Println(v8)
                            var v25 uint8 = v3
                            _ = v25
                        }
                        switch v6 {
                        case "hello":
                            var v26 uint64 = (^(v8))
                            _ = v26
                        case "foo bar":
                            switch "" {
                            case "hello":
                                v27 := GenS4{F0: false, F1: -901, F2: -114}
                                _ = v27
                            case "foo bar":
                                v2 += v2
                            case "":
                                if len(v1) > 0 {
                                    v1[0] = 48
                                }
                            default:
                                switch (fi10 / (((fi10 + fi10)) | 1)) {
                                case -20:
                                    v8 = v8
                                case 892:
                                    _ = len(v5)
                                case 4294967295:
                                    fmt.Println(v6, v2, v6)
                                default:
                                    var v28 int = (-(fi9))
                                    _ = v28
                                }
                            }
                        }
                        v29 := v5[v6]
                        _ = v29
                    }
                    if false {
                        v4 = append(v4, (v6 + v6))
                        v8 += 0
                    } else {
                        switch 42 {
                        case 127:
                            v5["hello"] = (fi9 << (((^(v2))) & 63))
                            if true {
                                var v30 int32 = 2147483647
                                v2 = (^(v2))
                                _ = v30
                            }
                        case -9223372036854775808:
                            _ = len(v5)
                        case 0:
                            if true {
                                if ((42 < 42) && false) {
                                    switch "a" {
                                    case "\n\t":
                                        if false {
                                            v31 := v5[v6]
                                            var v32 string = v6
                                            if ("\n\t" != "foo bar") {
                                                v33 := GenS3{F0: (^(v3)), F1: -514}
                                                var v34 uint8 = (^(v3))
                                                _ = v33
                                                _ = v34
                                            }
                                            v32 = v32
                                            var v35 uint = (v2 * (v2 & (v2 >> (((v2 >> ((65535) & 63))) & 63))))
                                            _ = v31
                                            _ = v32
                                            _ = v35
                                        } else {
                                            switch (^(fi10)) {
                                            case 2147483647:
                                                var v36 int64 = -9223372036854775808
                                                _ = v36
                                            case 32768:
                                                var v37 uint16 = 700
                                                _ = v37
                                            case 337:
                                                if ((fi9 == fi10) || true) {
                                                    v5["a"] = -9223372036854775808
                                                }
                                            default:
                                                switch (fi10 - fi10) {
                                                case 127:
                                                    v38 := GenS5{F0: -1, F1: int8(v2), F2: 255}
                                                    if true {
                                                        var v39 float32 = 1.5
                                                        _ = v39
                                                    }
                                                    fmt.Println(v6)
                                                    _ = v38
                                                case 32767:
                                                    v2 = 422
                                                case 4294967295:
                                                    var v40 uint8 = v3
                                                    _ = v40
                                                default:
                                                    if (!false) {
                                                        v3 = v3
                                                        v3 = 97
                                                    } else {
                                                        var v41 bool = (!((v6 < v6) || (-1 == 0)))
                                                        v42 := GenS2{F0: 1, F1: -1}
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                }
                                                var v43 int32 = 642
                                                switch v6 {
                                                case "":
                                                    v4 = append(v4, (v6 + v6))
                                                    v8 -= 1
                                                case "a":
                                                    switch v6 {
                                                    case "a":
                                                        var v44 uint8 = (v3 | 2)
                                                        _ = v44
                                                    }
                                                default:
                                                    var v45 uint64 = (v8 % ((v8) | 1))
                                                    _ = v45
                                                }
                                                v3 -= (v3 % ((v3) | 1))
                                                v6 = (v6 + (v6 + v6))
                                                _ = v43
                                            }
                                        }
                                    }
                                    if ((!true) && (!true)) {
                                        v8 += 0
                                        v8 -= v8
                                        fmt.Println(v2)
                                    }
                                } else {
                                    v1 = append(v1, 2)
                                    var v46 int8 = 0
                                    v47 := GenS6{F0: 0, F1: 525, F2: 9223372036854775807}
                                    _ = v46
                                    _ = v47
                                }
                            } else {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        default:
                            fmt.Println(fi9, fi9)
                        }
                        switch fi9 {
                        case -165:
                            var v48 int8 = -128
                            _ = v48
                        case 128:
                            v49 := GenS7{F0: 2, F1: (fi9 | (fi9 + fi9))}
                            _ = v49
                        case 921:
                            _ = len(v5)
                        }
                        if true {
                            var v50 uint16 = 495
                            v51 := v5[(v6 + v6)]
                            _ = v50
                            _ = v51
                        }
                    }
                    var v52 bool = true
                    v1 = append(v1, uint32(v3))
                    var v53 uint = v2
                    _ = v52
                    _ = v53
                }
                switch "" {
                case "\n\t":
                    v3 = (^(v3))
                }
            }
            _ = fi10
        }
        _ = fi9
    }
    if (-256 >= -1) {
        _ = len(v5)
        switch -315 {
        case -9223372036854775808:
            v5["foo bar"] = 1
        case 1:
            v3 = 255
        case 32768:
            if false {
                var v54 uint16 = 1
                _ = v54
            } else {
                v55 := GenS8{F0: 2147483648, F1: (v2 / ((420) | 1)), F2: 2}
                var v56 uint = (v2 ^ 790)
                v56 += 255
                _ = v55
                _ = v56
            }
        }
    }
    for fi57 := 0; fi57 < 5; fi57++ {
        var v58 string = "hello"
        if false {
            v59 := GenS8{F0: 2147483647, F1: (v2 ^ 563), F2: v8}
            _ = v59
        } else {
            v8 = v8
        }
        _ = fi57
        _ = v58
    }
    v8 = 4294967295
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
