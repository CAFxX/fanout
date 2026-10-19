package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 string; F1 int32 }
func genFn0(p0 uint32) int {
    v5 := []int{-2147483648, 0}
    var v6 int = 9223372036854775807
    if (true || (0 >= 96)) {
        for fi7 := 0; fi7 < 14; fi7++ {
            switch 317 {
            case -9223372036854775808:
                p0 = (p0 + 4294967295)
            default:
                for fi8 := 0; fi8 < 1; fi8++ {
                    if (42 < 46) {
                        var v9 float32 = (float32((1e10 / 1e10)) * float32(3.14159))
                        fmt.Println(p0, fi7)
                        if (!((p0 % ((p0) | 1)) < (p0 ^ (^(p0))))) {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            if (((((v9 / v9) * (0.0 - v9)) + v9) + v9) > float32(3.14159)) {
                                p0 += 2
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v10 := GenS1{F0: 2147483647}
                                v6 += (v6 << ((65535) & 63))
                                v11 := GenS2{F0: 133}
                                p0 = (p0 << ((uint(p0)) & 63))
                                _ = v10
                                _ = v11
                            } else {
                                var v12 int = (v6 / ((899) | 1))
                                _ = v12
                            }
                            fmt.Println(fi7, p0, fi7)
                        } else {
                            _ = len(v5)
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v9
                    }
                    p0 = p0
                    v6 += fi7
                    _ = fi8
                }
            }
            var v13 int8 = 42
            _ = fi7
            _ = v13
        }
        fmt.Println(p0)
    } else {
        if true {
            switch 0 {
            case -232:
                v6 -= 691
            case -1:
                switch (v6 & -62) {
                case 342:
                    if (!(2 < 255)) {
                        switch (v6 >> ((65535) & 63)) {
                        case 32767:
                            if len(v5) > 0 {
                                v5[0] = v6
                            }
                        case -2147483648:
                            var v14 uint64 = 2
                            _ = v14
                        case 1:
                            v6 = v6
                        default:
                            p0 += p0
                        }
                        var v15 string = ("" + "\n\t")
                        _ = v15
                    } else {
                        p0 += 2
                        fmt.Println(p0)
                        switch "" {
                        case "hello":
                            v16 := GenS3{F0: -14}
                            _ = v16
                        case "a":
                            p0 = p0
                        case "":
                            v17 := GenS4{F0: 1, F1: -1}
                            _ = v17
                        }
                    }
                }
            case 473:
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
            }
            p0 = 108
        }
        v18 := GenS1{F0: 394}
        _ = v18
    }
    v6 = v6
    _ = p0
    _ = v5
    _ = v6
    return (^(v6))
}
func genFn1(p0 string, p1 int16) int16 {
    v19 := []int{9223372036854775807, 0, -158, genFn0(2147483647)}
    switch 9223372036854775807 {
    case 427:
        p0 = ((p0 + p0) + (p0 + p0))
    }
    _ = p0
    _ = p1
    _ = v19
    return 42
}

func main() {
    var v1 uint32 = 0
    var v2 int64 = 1
    v3 := make(map[string]bool)
    var v4 int = -9223372036854775808
    v4 = v4
    v2 -= -9223372036854775808
    v2 += v2
    v20 := GenS3{F0: 42}
    v1 = v1
    v1 = v1
    var v21 uint8 = 1
    var v22 uint8 = 128
    var v23 float32 = 0.1
    v24 := GenS5{F0: "hello", F1: -2147483648}
    var v25 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
