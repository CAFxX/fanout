package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint16; F2 string }

func main() {
    v1 := []int{2147483647, 1, -924, -86}
    var v2 float32 = 1.5
    var v3 string = "a"
    var v4 uint8 = 128
    var v5 uint = uint(v4)
    v2 -= (v2 / (((v2 / v2) - v2) + (((0.1 - 0.1) * v2) * (v2 * (v2 + (-2.25 / v2))))))
    v6 := GenS1{F0: -1, F1: 2}
    var v7 int8 = -61
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v8 int16 = -378
    if (!(true && true)) {
        if false {
            v5 = v5
        } else {
            switch v3 {
            case "":
                for fi9 := 0; fi9 < 6; fi9++ {
                    for fi10 := 0; fi10 < 15; fi10++ {
                        if (32767 < 1) {
                            v2 += v2
                            var v11 uint32 = 4294967295
                            _ = v11
                        }
                        if (!(-1 <= 42)) {
                            v12 := GenS1{F0: 42, F1: 38}
                            _ = v12
                        } else {
                            var v13 bool = true
                            var v14 uint64 = 42
                            _ = len(v1)
                            v8 += v8
                            _ = v13
                            _ = v14
                        }
                        v15 := GenS1{F0: (v7 & (v7 ^ v7)), F1: 2}
                        _ = fi10
                        _ = v15
                    }
                    fmt.Println(v7)
                    fmt.Println(v2, v7)
                    _ = fi9
                }
            case "hello":
                if (32767 != 42) {
                    v16 := GenS1{F0: (v7 >> (((v5 << (((^(v5))) & 63))) & 63)), F1: 42}
                    if true {
                        for fi17 := 0; fi17 < 12; fi17++ {
                            v7 = (v7 << ((255) & 63))
                            _ = fi17
                        }
                    }
                    _ = v16
                }
            case "\n\t":
                switch 9223372036854775807 {
                case 4294967295:
                    fmt.Println(v5)
                case 431:
                    v5 = v5
                case 127:
                    v18 := GenS1{F0: 42, F1: 2}
                    _ = v18
                default:
                    if len(v1) > 0 {
                        v1[0] = 2147483647
                    }
                }
            default:
                switch -114 {
                case 4294967295:
                    var v19 int16 = v8
                    _ = v19
                case 1:
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v4 = v4
                        _ = fi20
                    }
                case -1:
                    v21 := GenS2{F0: v4, F1: 255, F2: -9223372036854775808}
                    _ = v21
                }
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v22 := GenS3{F0: 0, F1: (v4 & (v4 >> ((18446744073709551615) & 63)))}
    var v23 float64 = 3.14159
    v3 = v3
    var v24 float32 = v2
    var v25 uint16 = 65535
    if true {
        v23 = (float64(v24) / v23)
        if len(v1) > 0 {
            v1[0] = 710
        }
    } else {
        fmt.Println(v23, v2)
        v2 = float32(float64((0.1 * float32(0.1))))
    }
    var v26 string = v3
    if (!false) {
        if true {
            var v27 int64 = 0
            if (899 <= 2147483647) {
                v28 := GenS2{F0: (v4 + 168), F1: (v5 | (v5 % ((v5) | 1))), F2: int64(v5)}
                v5 += 0
                v26 = v26
                _ = v28
            } else {
                var v29 string = "hello"
                v29 = v29
                var v30 bool = true
                _ = v29
                _ = v30
            }
            _ = v27
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if false {
                var v31 uint64 = 459
                fmt.Println(v24, v7, v4)
                v32 := GenS3{F0: v4, F1: v4}
                _ = v31
                _ = v32
            }
            v33 := GenS4{F0: 0, F1: (v25 << (((v5 | 1)) & 63)), F2: "\n\t"}
            _ = v33
        }
        for fi34 := 0; fi34 < 8; fi34++ {
            switch v26 {
            case "a":
                if len(v1) > 0 {
                    v1[0] = fi34
                }
            case "\n\t":
                var v35 uint8 = v4
                _ = v35
            default:
                for fi36 := 0; fi36 < 4; fi36++ {
                    var v37 int8 = (-(v7))
                    _ = fi36
                    _ = v37
                }
            }
            for fi38 := 0; fi38 < 19; fi38++ {
                switch 4294967295 {
                case -2147483648:
                    v1 = append(v1, fi34)
                case 128:
                    if (!(65535 <= 2)) {
                        v8 = 1
                    }
                case 357:
                    v26 = (v26 + ((("a" + (("\n\t" + (("\n\t" + v3) + ("\n\t" + v26))) + v3)) + "") + (("\n\t" + "foo bar") + v3)))
                default:
                    var v39 int = fi34
                    var v40 int32 = -549
                    _ = v39
                    _ = v40
                }
                v2 += (v24 * 3.14159)
                v41 := GenS1{F0: v7, F1: 2}
                _ = fi38
                _ = v41
            }
            _ = fi34
        }
    } else {
        var v42 bool = (v5 > (v5 / (((v5 ^ (^(v5)))) | 1)))
        for fi43 := 0; fi43 < 4; fi43++ {
            _ = len(v1)
            switch v26 {
            case "":
                v8 = -57
            default:
                v44 := GenS1{F0: 127, F1: 1}
                _ = v44
            }
            _ = fi43
        }
        fmt.Println(v25, v23)
        _ = v42
    }
    v45 := GenS3{F0: (v4 + 42), F1: (^(v4))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v45
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
