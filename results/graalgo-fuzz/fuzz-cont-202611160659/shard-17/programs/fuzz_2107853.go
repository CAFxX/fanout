package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int; F2 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 uint8; F2 uint64 }

func main() {
    v1 := []int{4294967295, 733, -1, -9223372036854775808}
    var v2 uint32 = 65535
    var v3 int64 = 583
    v4 := make(map[int]bool)
    v5 := GenS1{F0: 37, F1: -440, F2: (v2 % ((42) | 1))}
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := GenS2{F0: 127}
        var v8 uint16 = 255
        if len(v1) > 0 {
            v1[0] = (-(fi6))
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS2{F0: 128}
    switch 0 {
    case 127:
        fmt.Println(v2, v3)
    case 9223372036854775807:
        var v10 int64 = (v3 + (v3 | (-(v3))))
        _ = v10
    case 2147483647:
        var v11 float32 = 0.1
        _ = v11
    }
    v12 := GenS3{F0: 89, F1: 65535}
    switch 4294967295 {
    case -104:
        var v13 int64 = v3
        _ = v13
    case 9223372036854775807:
        if false {
            v3 += (^(v3))
            var v14 uint = 206
            _ = v14
        } else {
            fmt.Println(v3, v3)
        }
    case 4294967295:
        if true {
            var v15 uint = 18446744073709551615
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi16 := 0; fi16 < 3; fi16++ {
                v3 += 42
                _ = fi16
            }
            _ = v15
        } else {
            v3 -= 9223372036854775807
            v17 := GenS3{F0: 0, F1: 925}
            v3 -= 1
            _ = v17
        }
    default:
        for fi18 := 0; fi18 < 1; fi18++ {
            _ = len(v4)
            if false {
                v3 += v3
                var v19 int8 = 114
                if (128 > 1) {
                    v20 := GenS2{F0: 253}
                    fmt.Println(v2, v19, v2)
                    for fi21 := 0; fi21 < 5; fi21++ {
                        if false {
                            switch ("foo bar" + ("\n\t" + "\n\t")) {
                            case "\n\t":
                                fmt.Println(v2, v2, v2)
                            case "foo bar":
                                _ = len(v4)
                            }
                        }
                        _ = fi21
                    }
                    _ = v20
                } else {
                    fmt.Println(fi18, v3, v3)
                    if (("" != "\n\t") && false) {
                        fmt.Println(v2, v3, v3)
                        switch "" {
                        case "a":
                            v3 -= 9223372036854775807
                        case "\n\t":
                            v22 := v4[(-(fi18))]
                            _ = v22
                        default:
                            var v23 int64 = v3
                            _ = v23
                        }
                        if true {
                            switch fi18 {
                            case -522:
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            case -101:
                                for fi24 := 0; fi24 < 7; fi24++ {
                                    var v25 int32 = 42
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = fi18
                                        }
                                        v26 := GenS1{F0: 255, F1: (v3 / ((-9223372036854775808) | 1)), F2: v2}
                                        v19 += v19
                                        _ = v26
                                    }
                                    _ = fi24
                                    _ = v25
                                }
                            default:
                                v2 = (v2 & (v2 ^ v2))
                            }
                            if (2 > 255) {
                                for fi27 := 0; fi27 < 11; fi27++ {
                                    v28 := GenS3{F0: 37, F1: 448}
                                    _ = fi27
                                    _ = v28
                                }
                            } else {
                                delete(v4, -2147483648)
                            }
                            if false {
                                v1 = append(v1, (fi18 & (^(fi18))))
                                v29 := v4[int(v3)]
                                switch "foo bar" {
                                case "hello":
                                    var v30 uint8 = 151
                                    _ = v30
                                case "a":
                                    fmt.Println(v2, v29)
                                }
                                _ = v29
                            }
                        } else {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            for fi31 := 0; fi31 < 14; fi31++ {
                                v19 = v19
                                delete(v4, fi18)
                                _ = fi31
                            }
                            v32 := GenS4{F0: 0, F1: (fi18 >> ((0) & 63)), F2: 99}
                            _ = v32
                        }
                    }
                }
                _ = v19
            }
            _ = fi18
        }
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        if false {
            if (80 > 42) {
                var v34 int8 = -128
                _ = v34
            }
            fmt.Println(fi33, v3, v2)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            switch fi33 {
            case -799:
                v3 = v3
            default:
                var v35 int64 = 9223372036854775807
                _ = v35
            }
            v36 := GenS5{F0: 1}
            var v37 int32 = -1
            _ = v36
            _ = v37
        }
        v38 := GenS6{F0: 127, F1: 2, F2: 2}
        _ = fi33
        _ = v38
    }
    v3 -= v3
    fmt.Println(v2, v2)
    v1 = append(v1, 32768)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
