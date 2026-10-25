package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint32; F1 bool; F2 string }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 uint32; F1 int32 }
type GenS7 struct { F0 uint32; F1 int16; F2 int64 }
func genFn0(p0 int16) uint16 {
    v6 := GenS1{F0: "\n\t", F1: 42, F2: 884}
    if false {
        fmt.Println(p0)
        v7 := []int{2147483647}
        _ = v7
    } else {
        for fi8 := 0; fi8 < 1; fi8++ {
            v9 := GenS2{F0: 42, F1: 892, F2: 11}
            fmt.Println(fi8)
            _ = fi8
            _ = v9
        }
    }
    if ((!false) && (2147483648 <= 65535)) {
        switch "hello" {
        case "":
            if false {
                var v10 uint = 255
                v11 := GenS1{F0: "hello", F1: -941, F2: 255}
                v10 = v10
                _ = v10
                _ = v11
            } else {
                p0 = (p0 | p0)
                p0 = -1
            }
        case "a":
            var v12 bool = (!true)
            _ = v12
        default:
            v13 := GenS1{F0: (("" + "\n\t") + "a"), F1: 42, F2: 2}
            _ = v13
        }
        var v14 uint16 = 278
        switch "" {
        case "foo bar":
            switch "a" {
            case "\n\t":
                v15 := []int{881, 42, 2147483647, -9223372036854775808}
                _ = v15
            default:
                p0 += p0
            }
        case "\n\t":
            for fi16 := 0; fi16 < 12; fi16++ {
                for fi17 := 0; fi17 < 16; fi17++ {
                    if false {
                        v18 := GenS2{F0: 0, F1: 42, F2: (p0 / (((p0 - 32767)) | 1))}
                        var v19 uint32 = 497
                        v20 := GenS3{F0: v19, F1: ((fi16 ^ fi17) != (fi16 >> ((490) & 63))), F2: "\n\t"}
                        v19 += v19
                        v21 := []int{(fi16 ^ 32767)}
                        switch -1 {
                        case 2147483647:
                            switch "foo bar" {
                            case "\n\t":
                                v19 -= 94
                                if false {
                                    if true {
                                        var v22 int8 = -40
                                        _ = v22
                                    }
                                } else {
                                    if (v14 >= (^(v14))) {
                                        var v23 int = (fi17 / ((2147483647) | 1))
                                        v19 -= v19
                                        fmt.Println(v14, v14)
                                        _ = v23
                                    }
                                }
                                v21 = append(v21, fi17)
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                if len(v21) >= 2 {
                                    v21 = v21[1:len(v21)]
                                }
                                var v24 int64 = 9223372036854775807
                                var v25 uint = 4294967295
                                v21 = append(v21, (fi16 | (^(fi17))))
                                switch "\n\t" {
                                case "\n\t":
                                    v26 := GenS4{F0: (v19 >> (((v25 - v25)) & 63)), F1: 0}
                                    _ = v26
                                case "foo bar":
                                    if true {
                                        v19 += (v19 * v19)
                                    }
                                default:
                                    var v27 string = "\n\t"
                                    _ = v27
                                }
                                _ = v24
                                _ = v25
                            case "foo bar":
                                var v28 uint16 = v14
                                _ = v28
                            case "hello":
                                var v29 int32 = 42
                                _ = v29
                            default:
                                var v30 uint32 = 1
                                _ = v30
                            }
                            if true {
                                if (-1 < 0) {
                                    v19 += (v19 >> ((42) & 63))
                                    var v31 bool = true
                                    _ = v31
                                } else {
                                    fmt.Println(v19, fi16)
                                    v32 := GenS5{F0: 131, F1: int32(fi16)}
                                    _ = v32
                                }
                                var v33 float64 = (float64(float32((1e10 + (0.1 * 0.1)))) * 1e10)
                                var v34 int16 = 947
                                _ = v33
                                _ = v34
                            }
                        }
                        var v35 int64 = 1
                        _ = v18
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v35
                    }
                    v36 := GenS6{F0: 2147483648, F1: 1}
                    _ = fi17
                    _ = v36
                }
                _ = fi16
            }
        case "hello":
            if true {
                var v37 int8 = 100
                var v38 float32 = (float32(0.1) / ((float32(float64(1.5)) * 0.0) + (3.14159 / ((3.14159 - (1e10 + (((0.1 + 3.14159) - (3.14159 - ((-2.25 + -2.25) - -2.25))) * (float32((float64(1.5) + 0.0)) + ((3.14159 / float32(((1.5 * 0.1) + 3.14159))) + (((float32(float64(1e10)) * 0.1) / 3.14159) / 1e10)))))) / (1e10 * 1.5)))))
                _ = v37
                _ = v38
            } else {
                v39 := []int{}
                var v40 uint8 = 126
                v41 := GenS4{F0: 801, F1: 0}
                _ = v39
                _ = v40
                _ = v41
            }
        }
        _ = v14
    } else {
        p0 -= p0
        var v42 uint16 = 0
        v42 += (v42 ^ v42)
        _ = v42
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (true && (444 > 1))
    var v3 int8 = 1
    v4 := []int{127, 2147483647, 32768}
    var v5 string = "a"
    for fi43 := 0; fi43 < 9; fi43++ {
        fmt.Println(fi43)
        if (true || v2) {
            var v44 float32 = float32(-2.25)
            for fi45 := 0; fi45 < 13; fi45++ {
                v2 = v2
                v3 -= (v3 & (v3 << ((18446744073709551615) & 63)))
                _ = fi45
            }
            for fi46 := 0; fi46 < 17; fi46++ {
                fmt.Println(v44, v2, v2)
                if len(v4) > 0 {
                    v4[0] = 32767
                }
                v4 = append(v4, (^(fi43)))
                _ = fi46
            }
            _ = v44
        } else {
            if (("a" != v5) || true) {
                for fi47 := 0; fi47 < 7; fi47++ {
                    if v2 {
                        v48 := GenS2{F0: 428, F1: 42, F2: -32768}
                        _ = v48
                    }
                    _ = fi47
                }
                for fi49 := 0; fi49 < 3; fi49++ {
                    switch fi43 {
                    case 32768:
                        v2 = (4294967295 == 42)
                    default:
                        var v50 float64 = 3.14159
                        _ = v50
                    }
                    v2 = (18446744073709551615 < 18446744073709551615)
                    v51 := GenS4{F0: 65535, F1: 42}
                    _ = fi49
                    _ = v51
                }
                if (461 > 4294967295) {
                    v5 = (("a" + v5) + (v5 + "\n\t"))
                }
            } else {
                var v52 int64 = -1
                _ = v52
            }
        }
        _ = fi43
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        v2 = v2
        var v54 string = ""
        var v55 uint32 = 2
        _ = fi53
        _ = v54
        _ = v55
    }
    var v56 uint16 = 0
    for fi57 := 0; fi57 < 17; fi57++ {
        var v58 int8 = v3
        v56 = (v56 * v56)
        _ = fi57
        _ = v58
    }
    var v59 int32 = 42
    fmt.Println(v2, v5)
    _ = len(v1)
    var v60 int = 32767
    _ = len(v4)
    for fi61 := 0; fi61 < 19; fi61++ {
        var v62 uint = 0
        v63 := GenS2{F0: (v59 % ((-1) | 1)), F1: 296, F2: 32767}
        _ = fi61
        _ = v62
        _ = v63
    }
    var v64 int64 = 202
    if false {
        v59 += v59
        v56 = (v56 ^ genFn0(42))
    } else {
        v65 := GenS7{F0: 0, F1: int16(v64), F2: (^(v64))}
        fmt.Println(v56, v64, v59)
        _ = v65
    }
    v59 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v59
    _ = v60
    _ = v64
}
