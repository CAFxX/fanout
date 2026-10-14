package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8; F2 uint8 }
type GenS2 struct { F0 int16; F1 int16; F2 int }
type GenS3 struct { F0 int32; F1 int }
type GenS4 struct { F0 int32; F1 int32; F2 uint32 }
func genFn0(p0 int64) int16 {
    var v7 float64 = (-2.25 / 1.0)
    v7 = ((v7 * float64(3.14159)) + (v7 + ((v7 / (float64(3.14159) * 3.14159)) * 1.5)))
    v7 = (v7 * 0.0)
    v8 := GenS1{F0: 4294967295, F1: 159, F2: 42}
    _ = p0
    _ = v7
    _ = v8
    return 1000
}

func main() {
    var v1 uint8 = 119
    var v2 int8 = 0
    var v3 int32 = 42
    v4 := []string{"", "a"}
    var v5 int32 = (v3 ^ v3)
    v6 := []uint32{657, 351, 1}
    v2 -= v2
    _ = len(v6)
    _ = len(v4)
    var v9 int8 = v2
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 uint64 = uint64(v1)
            switch -9223372036854775808 {
            case -9223372036854775808:
                v5 = v3
            case 396:
                v13 := GenS2{F0: int16(fi10), F1: genFn0(-227), F2: fi10}
                _ = v13
            case 128:
                v12 = v12
            }
            if (2147483647 < 166) {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            } else {
                _ = len(v4)
                fmt.Println(fi10, v1, v12)
                if ("hello" >= "\n\t") {
                    if true {
                        if (-32768 > genFn0(1)) {
                            var v14 uint = 18446744073709551615
                            if true {
                                if (623 >= 42) {
                                    var v15 uint8 = v1
                                    switch (fi11 - (fi10 + (-(fi10)))) {
                                    case -2147483648:
                                        v2 = (v2 * v2)
                                    }
                                    switch "" {
                                    case "hello":
                                        var v16 uint32 = 831
                                        _ = v16
                                    case "foo bar":
                                        v15 -= (v1 % (((^(v1))) | 1))
                                        switch "hello" {
                                        case "\n\t":
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        case "hello":
                                            v17 := GenS3{F0: (-(v5)), F1: (fi11 ^ fi10)}
                                            _ = v17
                                        }
                                        switch "hello" {
                                        case "a":
                                            if len(v6) > 0 {
                                                v6[0] = 396
                                            }
                                        default:
                                            if false {
                                                if true {
                                                    v15 -= (v15 * v1)
                                                    if false {
                                                        _ = len(v6)
                                                    }
                                                    fmt.Println(v14, v12, v15)
                                                    v3 = (v3 >> (((v14 ^ v14)) & 63))
                                                    var v18 int8 = 42
                                                    var v19 uint64 = v12
                                                    var v20 float32 = 1e10
                                                    _ = v18
                                                    _ = v19
                                                    _ = v20
                                                } else {
                                                    if false {
                                                        var v21 int8 = 100
                                                        _ = v21
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = "a"
                                                        }
                                                        switch fi10 {
                                                        case 2147483647:
                                                            var v22 uint = (v14 % ((398) | 1))
                                                            _ = v22
                                                        case 128:
                                                            if len(v6) > 0 {
                                                                v6[0] = 1
                                                            }
                                                        default:
                                                            v23 := GenS4{F0: v3, F1: v3, F2: 42}
                                                            _ = v23
                                                        }
                                                        fmt.Println(v12, v12)
                                                    }
                                                }
                                                if (true && (42 != (v9 + v9))) {
                                                    if false {
                                                        v24 := GenS1{F0: 2147483648, F1: 171, F2: (v15 % ((179) | 1))}
                                                        _ = v24
                                                    } else {
                                                        v6 = append(v6, 4294967295)
                                                    }
                                                } else {
                                                    v12 = (^(v12))
                                                    fmt.Println(fi11)
                                                }
                                                v5 = v5
                                                v25 := GenS4{F0: 852, F1: 1, F2: 2147483647}
                                                var v26 int64 = -306
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                    case "":
                                        var v27 uint16 = 0
                                        _ = v27
                                    }
                                    _ = v15
                                }
                            }
                            _ = v14
                        } else {
                            v3 = -2147483648
                        }
                    } else {
                        switch ("foo bar" + (("\n\t" + ("" + "")) + ("foo bar" + "foo bar"))) {
                        case "a":
                            v28 := GenS1{F0: 42, F1: v1, F2: 128}
                            _ = v28
                        default:
                            v12 -= 0
                        }
                        v1 = (v1 / ((v1) | 1))
                        v29 := GenS4{F0: (^(v5)), F1: v3, F2: 2147483647}
                        _ = v29
                    }
                    if len(v6) > 0 {
                        v6[0] = 65535
                    }
                    fmt.Println(v1)
                    v30 := GenS2{F0: -679, F1: -32768, F2: (fi11 & (fi11 >> ((65535) & 63)))}
                    _ = v30
                }
            }
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    v1 -= (v1 | (v1 % ((v1) | 1)))
    var v31 uint16 = 65535
    v6 = append(v6, 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
