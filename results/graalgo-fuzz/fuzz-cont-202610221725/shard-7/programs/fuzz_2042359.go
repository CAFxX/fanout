package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int64 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[int]uint32)
    v3 := []int{-9223372036854775808, 812, 32767}
    var v4 float64 = 3.14159
    v5 := []int{}
    switch 4294967295 {
    case 2147483647:
        if true {
            switch "foo bar" {
            case "foo bar":
                if (-838 <= -9223372036854775808) {
                    v4 = v4
                } else {
                    if (144 == 42) {
                        for fi6 := 0; fi6 < 6; fi6++ {
                            if ((-1 == 407) || (("a" >= ("foo bar" + (("" + ("\n\t" + "")) + ("" + "hello")))) || false)) {
                                switch fi6 {
                                case -9223372036854775808:
                                    for fi7 := 0; fi7 < 6; fi7++ {
                                        if true {
                                            if (-1 != 0) {
                                                var v8 int32 = 0
                                                _ = v8
                                            }
                                            var v9 string = "foo bar"
                                            fmt.Println(v4, fi7)
                                            v10 := GenS1{F0: 4294967295}
                                            _ = v10
                                            _ = v9
                                        } else {
                                            v4 -= (v4 / 1.5)
                                            switch -9223372036854775808 {
                                            case -71:
                                                var v11 string = "hello"
                                                _ = v11
                                            case -163:
                                                var v12 float64 = v4
                                                _ = v12
                                            default:
                                                v4 -= v4
                                            }
                                            v1 = v1
                                            v4 = 1.5
                                        }
                                        _ = fi7
                                    }
                                default:
                                    switch fi6 {
                                    case 4294967295:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case 32767:
                                        v13 := GenS2{F0: 394, F1: 42, F2: 42}
                                        _ = v13
                                    case 174:
                                        v14 := v2[-9223372036854775808]
                                        _ = v14
                                    }
                                }
                                v2[(fi6 << ((42) & 63))] = 2147483648
                                for fi15 := 0; fi15 < 19; fi15++ {
                                    var v16 uint64 = 4294967295
                                    _ = fi15
                                    _ = v16
                                }
                            }
                            fmt.Println(fi6)
                            _ = fi6
                        }
                        v17 := GenS3{F0: v1}
                        switch 4294967295 {
                        case 0:
                            var v18 int = 804
                            _ = v18
                        case 127:
                            v4 += v4
                        case 780:
                            switch "" {
                            case "hello":
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case "":
                                if (!true) {
                                    for fi19 := 0; fi19 < 18; fi19++ {
                                        for fi20 := 0; fi20 < 13; fi20++ {
                                            var v21 float32 = 0.1
                                            v22 := GenS3{F0: (v1 % (((^(v1))) | 1))}
                                            if (42 > -9223372036854775808) {
                                                if false {
                                                    var v23 float64 = (float64((v21 * 1e10)) - 3.14159)
                                                    v23 = (0.0 * 0.1)
                                                    v21 = ((float32(3.14159) * 0.1) * 1e10)
                                                    _ = v23
                                                }
                                                var v24 float32 = v21
                                                _ = v24
                                            }
                                            _ = fi20
                                            _ = v21
                                            _ = v22
                                        }
                                        for fi25 := 0; fi25 < 13; fi25++ {
                                            v1 = (v1 ^ v1)
                                            fmt.Println(fi19, v4, fi25)
                                            _ = fi25
                                        }
                                        var v26 uint64 = 982
                                        _ = fi19
                                        _ = v26
                                    }
                                }
                            }
                        default:
                            var v27 uint8 = (v1 >> ((255) & 63))
                            _ = v27
                        }
                        _ = v17
                    } else {
                        fmt.Println(v4)
                        var v28 uint16 = 42
                        _ = v28
                    }
                }
            case "hello":
                fmt.Println(v1)
            default:
                v4 = (3.14159 - 1.5)
            }
        } else {
            switch "a" {
            case "":
                delete(v2, 1)
            case "\n\t":
                var v29 bool = false
                _ = v29
            }
            v30 := GenS4{F0: 2}
            v31 := GenS2{F0: 0, F1: 145, F2: 0}
            _ = v30
            _ = v31
        }
    case 4294967295:
        v4 -= (float64(1.5) / 3.14159)
    case 32767:
        v4 += (v4 + v4)
    default:
        v1 -= 193
    }
    v1 = v1
    fmt.Println(v4, v1)
    v1 -= 127
    var v32 int16 = -1
    v33 := GenS5{F0: 1}
    v32 = (v32 / ((v32) | 1))
    for fi34 := 0; fi34 < 8; fi34++ {
        var v35 uint8 = 2
        if false {
            var v36 string = "hello"
            _ = v36
        }
        _ = fi34
        _ = v35
    }
    v37 := GenS6{F0: 9223372036854775807}
    v32 -= (v32 | v32)
    switch "hello" {
    case "":
        v32 += (v32 | v32)
    default:
        v32 = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
