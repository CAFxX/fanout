package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 uint32; F2 uint }
func genFn0(p0 int, p1 int64, p2 string) int32 {
    v4 := []int{}
    var v5 uint64 = 0
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := GenS1{F0: 101}
        v8 := GenS1{F0: 2}
        switch p2 {
        case "a":
            var v9 string = "hello"
            _ = v9
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v10 := GenS2{F0: 2}
    v11 := GenS2{F0: 371}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v4
    _ = v5
    return -684
}

func main() {
    v1 := make(map[int]bool)
    v2 := []string{("\n\t" + "\n\t"), "foo bar", "\n\t", "hello"}
    var v3 int64 = 42
    var v12 string = "foo bar"
    v12 = v12
    var v13 int32 = 42
    v13 = v13
    for fi14 := 0; fi14 < 2; fi14++ {
        for fi15 := 0; fi15 < 4; fi15++ {
            switch v12 {
            case "a":
                var v16 uint = 65535
                _ = v16
            default:
                switch "foo bar" {
                case "foo bar":
                    if (((-2.25 / ((1e10 / (3.14159 - float64(-2.25))) + -2.25)) > 1.5) || false) {
                        v17 := GenS2{F0: 351}
                        if false {
                            switch fi15 {
                            case 32768:
                                v13 = (v13 * -2147483648)
                            default:
                                v2 = append(v2, v12)
                            }
                            var v18 uint64 = 0
                            _ = v18
                        }
                        if (2 > 791) {
                            var v19 uint32 = 65535
                            var v20 int = (fi14 >> ((18446744073709551615) & 63))
                            _ = v19
                            _ = v20
                        } else {
                            var v21 float32 = 3.14159
                            v22 := GenS1{F0: 18446744073709551615}
                            v23 := GenS2{F0: 42}
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                        v24 := GenS3{F0: 255, F1: 476, F2: 266}
                        _ = v17
                        _ = v24
                    } else {
                        if true {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(fi15, v12, v3)
                        } else {
                            fmt.Println(fi15, v3)
                            v25 := GenS2{F0: 353}
                            var v26 string = v12
                            v27 := GenS1{F0: 666}
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v25
                            _ = v26
                            _ = v27
                        }
                        v13 = v13
                        if len(v2) > 0 {
                            v2[0] = ("a" + "foo bar")
                        }
                        if (1.5 > 0.1) {
                            if true {
                                var v28 int64 = v3
                                if ("foo bar" < v12) {
                                    v29 := GenS2{F0: 32767}
                                    _ = v29
                                }
                                v28 += (v28 * (v28 | (^(v28))))
                                _ = v28
                            } else {
                                fmt.Println(v12)
                                v3 = (v3 ^ v3)
                                if len(v2) > 0 {
                                    v2[0] = v12
                                }
                            }
                        }
                    }
                case "":
                    switch v12 {
                    case "a":
                        switch fi15 {
                        case 29:
                            _ = len(v1)
                        default:
                            v1[128] = (!(1e10 < ((((((float32(float64(((3.14159 * 1e10) / (1.5 + (float32(-2.25) + float32(0.0)))))) * float32(0.1)) + ((3.14159 / -2.25) - (((1.5 + 0.1) - float32((1.5 * 0.0))) * 3.14159))) + -2.25) / (0.1 / ((0.1 + 0.0) / 1e10))) + 0.1) - -2.25)))
                        }
                    case "":
                        v30 := GenS3{F0: 255, F1: 4294967295, F2: 923}
                        _ = v30
                    }
                }
            }
            _ = fi15
        }
        switch fi14 {
        case -9223372036854775808:
            for fi31 := 0; fi31 < 11; fi31++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                v12 = ((v12 + (((v12 + (v12 + "foo bar")) + "a") + "hello")) + "\n\t")
                _ = fi31
            }
        case 2147483647:
            v13 -= 551
        }
        _ = fi14
    }
    var v32 int64 = v3
    var v33 uint16 = 405
    if false {
        switch 1 {
        case 32768:
            v34 := GenS3{F0: 2, F1: 628, F2: 976}
            _ = v34
        case 9223372036854775807:
            v12 = (v12 + v12)
        case 127:
            var v35 uint8 = 128
            _ = v35
        }
    } else {
        v12 = v12
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v32
    _ = v33
}
