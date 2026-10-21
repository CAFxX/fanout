package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 int8; F1 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int16; F1 int64 }
type GenS7 struct { F0 bool; F1 int8; F2 string }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint64 = 2
    var v2 string = "hello"
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    for fi6 := 0; fi6 < 11; fi6++ {
        switch v2 {
        case "hello":
            for fi7 := 0; fi7 < 16; fi7++ {
                v2 = v2
                v8 := GenS1{F0: (v1 % (((v1 % (((v1 >> ((65535) & 63))) | 1))) | 1)), F1: 1000, F2: 128}
                _ = fi7
                _ = v8
            }
        case "":
            switch ("a" + v2) {
            case "hello":
                for fi9 := 0; fi9 < 1; fi9++ {
                    v1 += v1
                    var v10 int16 = -1
                    v10 = v10
                    _ = fi9
                    _ = v10
                }
            case "\n\t":
                if ((42 >= -1) && (!true)) {
                    var v11 uint64 = (v1 >> ((18446744073709551615) & 63))
                    _ = v11
                } else {
                    v12 := GenS1{F0: 217, F1: -484, F2: 42}
                    _ = v12
                }
            case "a":
                var v13 int32 = -2147483648
                _ = v13
            default:
                switch v2 {
                case "foo bar":
                    fmt.Println(v1, v1)
                case "\n\t":
                    var v14 int32 = 42
                    _ = v14
                default:
                    if false {
                        var v15 uint = 18446744073709551615
                        fmt.Println(v15)
                        _ = v15
                    }
                }
            }
        case "a":
            v16 := GenS2{F0: (v2 + ("a" + (v2 + "foo bar")))}
            _ = v16
        }
        var v17 int64 = 710
        _ = fi6
        _ = v17
    }
    fmt.Println(v1)
    if ((v1 | (v1 - 4294967295)) <= (v1 << ((65535) & 63))) {
        var v18 float32 = -2.25
        v19 := []int{9223372036854775807, 42}
        _ = v18
        _ = v19
    } else {
        var v20 bool = true
        if true {
            if (!v20) {
                v20 = ((-9223372036854775808 > -1) || v20)
                switch -721 {
                case -649:
                    v20 = v20
                case 4294967295:
                    if false {
                        if v20 {
                            v2 = v2
                            fmt.Println(v2)
                        }
                        if v20 {
                            switch 4294967295 {
                            case 127:
                                v20 = v20
                            }
                            v21 := v3[996]
                            _ = v21
                        }
                    } else {
                        var v22 bool = (false || true)
                        fmt.Println(v20)
                        _ = v22
                    }
                case 128:
                    var v23 int32 = 2147483647
                    _ = v23
                }
                fmt.Println(v1, v20)
            } else {
                switch 127 {
                case 9223372036854775807:
                    for fi24 := 0; fi24 < 17; fi24++ {
                        if v20 {
                            for fi25 := 0; fi25 < 3; fi25++ {
                                v1 += 2
                                _ = fi25
                            }
                            v2 = (v2 + (v2 + v2))
                            var v26 int16 = 1
                            _ = v26
                        } else {
                            v27 := GenS2{F0: v2}
                            switch ("" + v2) {
                            case "a":
                                v28 := GenS2{F0: ""}
                                _ = v28
                            }
                            _ = len(v5)
                            _ = v27
                        }
                        fmt.Println(fi24, v2, v1)
                        if false {
                            var v29 int16 = 722
                            var v30 int32 = 1
                            _ = v29
                            _ = v30
                        }
                        _ = fi24
                    }
                case 217:
                    if (2 != 42) {
                        v1 = (v1 / ((v1) | 1))
                        for fi31 := 0; fi31 < 4; fi31++ {
                            fmt.Println(fi31, fi31)
                            v32 := []int{0}
                            _ = fi31
                            _ = v32
                        }
                    }
                default:
                    for fi33 := 0; fi33 < 5; fi33++ {
                        if v20 {
                            var v34 int = -2147483648
                            _ = v34
                        } else {
                            var v35 int16 = 202
                            v36 := []int{(fi33 % ((fi33) | 1)), (fi33 * fi33), 1, fi33}
                            _ = v35
                            _ = v36
                        }
                        v20 = (0 >= 649)
                        _ = fi33
                    }
                }
                v37 := []int{2147483647}
                if len(v37) > 0 {
                    v37[0] = 0
                }
                _ = v37
            }
            v38 := GenS3{F0: v1, F1: 86}
            _ = v38
        } else {
            if false {
                v39 := GenS2{F0: ("\n\t" + "hello")}
                var v40 float32 = ((1e10 + (1e10 + (-2.25 - 0.1))) + (0.1 / (-2.25 - 0.0)))
                _ = v39
                _ = v40
            } else {
                for fi41 := 0; fi41 < 11; fi41++ {
                    v42 := GenS2{F0: v2}
                    _ = fi41
                    _ = v42
                }
            }
        }
        for fi43 := 0; fi43 < 4; fi43++ {
            v20 = (v20 || v20)
            _ = fi43
        }
        _ = v20
    }
    v44 := GenS2{F0: (v2 + (v2 + v2))}
    var v45 float32 = (3.14159 + -2.25)
    switch 806 {
    case -456:
        v46 := GenS4{F0: 1, F1: (v1 % ((v1) | 1))}
        _ = v46
    case -617:
        var v47 int8 = 42
        _ = v47
    case 536:
        v48 := GenS4{F0: -128, F1: 363}
        _ = v48
    default:
        fmt.Println(v2)
    }
    for fi49 := 0; fi49 < 15; fi49++ {
        if false {
            v50 := GenS5{F0: 1}
            var v51 float64 = (float64(v45) / 1.5)
            fmt.Println(v45)
            _ = v50
            _ = v51
        } else {
            switch (("hello" + (("hello" + ("a" + "a")) + v2)) + v2) {
            case "foo bar":
                v4[(fi49 | 0)] = 1
            case "a":
                v1 -= (^(v1))
            case "hello":
                v52 := GenS6{F0: -1, F1: 0}
                _ = v52
            default:
                var v53 uint = 42
                _ = v53
            }
        }
        _ = fi49
    }
    var v54 uint64 = v1
    var v55 uint8 = 250
    var v56 uint32 = 2147483647
    v57 := GenS7{F0: (245 >= (v56 << ((155) & 63))), F1: -13, F2: v2}
    switch 127 {
    case 1:
        v1 = 644
    case 449:
        _ = len(v4)
    default:
        v56 += 42
    }
    v56 -= v56
    v58 := GenS8{F0: v55}
    var v59 int = 127
    if (380 <= 42) {
        v2 = v2
        var v60 int16 = int16(v1)
        _ = v60
    } else {
        var v61 float64 = ((0.1 * -2.25) / 1e10)
        if true {
            switch v2 {
            case "foo bar":
                fmt.Println(v59, v1)
            case "a":
                switch v59 {
                case 1:
                    fmt.Println(v56, v59)
                case -9223372036854775808:
                    v62 := GenS6{F0: -519, F1: 42}
                    _ = v62
                default:
                    v63 := GenS4{F0: 0, F1: 1}
                    _ = v63
                }
            case "\n\t":
                v64 := GenS9{F0: 1, F1: 42}
                _ = v64
            default:
                for fi65 := 0; fi65 < 6; fi65++ {
                    var v66 int = (fi65 | (fi65 / ((975) | 1)))
                    fmt.Println(v54, v55, v59)
                    _ = fi65
                    _ = v66
                }
            }
            fmt.Println(v54, v54, v55)
            if (!true) {
                fmt.Println(v56)
            } else {
                for fi67 := 0; fi67 < 14; fi67++ {
                    v5["\n\t"] = false
                    v68 := GenS6{F0: 42, F1: 391}
                    _ = fi67
                    _ = v68
                }
                for fi69 := 0; fi69 < 12; fi69++ {
                    var v70 uint64 = (v1 ^ (v1 * (v1 ^ (^(v1)))))
                    v3[v56] = uint32(v1)
                    _ = fi69
                    _ = v70
                }
            }
        }
        _ = v61
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
