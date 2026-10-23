package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8; F2 int16 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS3 struct { F0 string; F1 uint8; F2 bool }
type GenS4 struct { F0 int64; F1 uint16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int64; F1 int8 }
type GenS7 struct { F0 uint32 }
func genFn0() int16 {
    if true {
        var v4 float64 = 0.1
        _ = v4
    }
    v5 := []int{-9223372036854775808, 127, 32768, 481}
    var v6 float64 = 1e10
    _ = v5
    _ = v6
    return -32768
}

func main() {
    v1 := []int{427, -150, -2147483648, 32767}
    var v2 bool = false
    var v3 uint32 = 1
    switch 510 {
    case 127:
        for fi7 := 0; fi7 < 10; fi7++ {
            for fi8 := 0; fi8 < 7; fi8++ {
                if (uint64(v3) <= 562) {
                    if (!(-1 >= 42)) {
                        if ((0.0 <= (1.5 - 1.5)) && v2) {
                            v2 = false
                            switch "foo bar" {
                            case "hello":
                                v3 += (^(v3))
                            case "\n\t":
                                fmt.Println(fi8)
                            case "a":
                                v2 = (!v2)
                            default:
                                if len(v1) > 0 {
                                    v1[0] = fi8
                                }
                            }
                        } else {
                            var v9 int16 = 0
                            switch "\n\t" {
                            case "\n\t":
                                if (1e10 != 3.14159) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    v9 = v9
                                    _ = len(v1)
                                    fmt.Println(v3, fi7, v2)
                                    var v10 uint = 18446744073709551615
                                    var v11 uint64 = 1
                                    v12 := GenS1{F0: 42, F1: -1, F2: (v9 ^ v9)}
                                    var v13 float64 = -2.25
                                    v11 += v11
                                    _ = v10
                                    _ = v11
                                    _ = v12
                                    _ = v13
                                }
                            case "a":
                                switch 127 {
                                case 9223372036854775807:
                                    _ = len(v1)
                                }
                            }
                            v14 := GenS1{F0: 1, F1: 89, F2: v9}
                            _ = v14
                            _ = v9
                        }
                        var v15 uint16 = 2
                        _ = v15
                    }
                }
                v16 := GenS2{F0: (!(-1 > -9223372036854775808)), F1: v3, F2: 3}
                fmt.Println(fi8)
                _ = fi8
                _ = v16
            }
            var v17 uint = 2
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi7
            _ = v17
        }
    case 813:
        if len(v1) > 0 {
            v1[0] = 32768
        }
    case 32767:
        v18 := GenS1{F0: 72, F1: 42, F2: -150}
        _ = v18
    default:
        v19 := GenS1{F0: 941, F1: -82, F2: genFn0()}
        _ = v19
    }
    var v20 float32 = 1.5
    v3 = (v3 & v3)
    v21 := GenS2{F0: true, F1: v3, F2: 2}
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    if (v3 > (v3 | (^(v3)))) {
        switch 1 {
        case 9223372036854775807:
            for fi22 := 0; fi22 < 6; fi22++ {
                v23 := GenS3{F0: "hello", F1: uint8(v3), F2: (false && false)}
                fmt.Println(v3, fi22)
                if len(v1) > 0 {
                    v1[0] = fi22
                }
                _ = fi22
                _ = v23
            }
        case 4294967295:
            for fi24 := 0; fi24 < 9; fi24++ {
                var v25 float32 = 1.5
                for fi26 := 0; fi26 < 7; fi26++ {
                    v27 := GenS4{F0: 316, F1: 467}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi26
                    _ = v27
                }
                _ = fi24
                _ = v25
            }
        }
    }
    v28 := GenS5{F0: 32767}
    v29 := GenS6{F0: 208, F1: 100}
    v30 := GenS2{F0: (!v2), F1: (v3 % (((^(v3))) | 1)), F2: 128}
    switch 32767 {
    case 2147483647:
        v3 = v3
    case 73:
        v3 += v3
    case -1:
        v31 := GenS7{F0: v3}
        _ = v31
    }
    var v32 int32 = 42
    for fi33 := 0; fi33 < 8; fi33++ {
        for fi34 := 0; fi34 < 11; fi34++ {
            switch "" {
            case "hello":
                switch fi33 {
                case 2147483647:
                    var v35 uint16 = 216
                    _ = v35
                default:
                    var v36 uint8 = 2
                    _ = v36
                }
            case "\n\t":
                var v37 uint32 = (v3 ^ (v3 % (((v3 + v3)) | 1)))
                _ = v37
            case "foo bar":
                var v38 float32 = float32(3.14159)
                v39 := GenS4{F0: 9223372036854775807, F1: 65535}
                v3 += (^(v3))
                _ = v38
                _ = v39
            default:
                var v40 uint32 = (^(v3))
                if false {
                    v41 := GenS7{F0: v40}
                    fmt.Println(v2, v3)
                    switch "foo bar" {
                    case "":
                        v3 += v3
                    case "\n\t":
                        switch int(v40) {
                        case -836:
                            var v42 int16 = -32768
                            _ = v42
                        case 683:
                            v43 := GenS2{F0: (1 > uint64(v40)), F1: (v3 / (((v3 - (v3 << ((65535) & 63)))) | 1)), F2: 127}
                            v44 := GenS5{F0: (fi33 * (fi33 >> ((4294967295) & 63)))}
                            _ = v43
                            _ = v44
                        default:
                            var v45 int8 = -89
                            _ = v45
                        }
                    }
                    _ = v41
                }
                _ = v40
            }
            _ = fi34
        }
        _ = fi33
    }
    fmt.Println(v2, v32, v20)
    switch "\n\t" {
    case "hello":
        v20 = (v20 + (v20 + (float32(0.0) + 1.5)))
    }
    for fi46 := 0; fi46 < 16; fi46++ {
        for fi47 := 0; fi47 < 5; fi47++ {
            if (v32 < v32) {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v1 = append(v1, 42)
            } else {
                v2 = false
            }
            v2 = v2
            var v48 uint = 2
            if (-32768 >= 42) {
                var v49 bool = true
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                fmt.Println(fi47, fi47)
                _ = v49
            } else {
                if (1 < int64(v32)) {
                    if true {
                        v20 = v20
                    } else {
                        v20 += 3.14159
                    }
                } else {
                    var v50 int8 = -1
                    _ = v50
                }
                var v51 float32 = (v20 + v20)
                v52 := GenS1{F0: 2, F1: 0, F2: 42}
                _ = v51
                _ = v52
            }
            v48 = (v48 - v48)
            _ = fi47
            _ = v48
        }
        _ = fi46
    }
    if ((-128 == -1) || (0.0 != (3.14159 / ((-2.25 + float64(3.14159)) - (1e10 * 3.14159))))) {
        v53 := GenS1{F0: 2, F1: 127, F2: 1000}
        _ = v53
    }
    fmt.Println(v2)
    fmt.Println(v32)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v32
}
