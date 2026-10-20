package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint; F1 int8 }

func main() {
    var v1 int32 = 0
    v2 := []string{}
    v3 := []int{-184, 528}
    var v4 uint64 = 18446744073709551615
    var v5 int64 = 738
    v6 := make(map[uint32]int)
    v7 := v6[1]
    v7 -= 32767
    var v8 uint16 = 180
    var v9 int32 = (v1 << ((1) & 63))
    switch v7 {
    case -2147483648:
        var v10 int = -1
        _ = v10
    case 944:
        if false {
            v1 += (v1 / ((v1) | 1))
            fmt.Println(v1)
            if true {
                if (!(!false)) {
                    v9 = (-(v1))
                    switch "hello" {
                    case "foo bar":
                        v8 += (^(v8))
                    case "\n\t":
                        var v11 uint64 = (v4 << ((2) & 63))
                        _ = v11
                    default:
                        v6[2147483648] = (^(v7))
                    }
                } else {
                    if false {
                        for fi12 := 0; fi12 < 15; fi12++ {
                            fmt.Println(v8)
                            switch 485 {
                            case -856:
                                v13 := GenS1{F0: 121, F1: 0, F2: 235}
                                _ = v13
                            case -9223372036854775808:
                                switch v7 {
                                case 42:
                                    for fi14 := 0; fi14 < 11; fi14++ {
                                        if ((v4 & 2) <= (v4 % ((154) | 1))) {
                                            var v15 uint32 = 553
                                            var v16 int64 = v5
                                            _ = v15
                                            _ = v16
                                        } else {
                                            v7 = (v7 % ((fi14) | 1))
                                        }
                                        if false {
                                            if true {
                                                v17 := GenS2{F0: 42, F1: 42}
                                                if (65535 < 0) {
                                                    v7 = (fi14 >> ((4294967295) & 63))
                                                    if ("\n\t" != "foo bar") {
                                                        v18 := GenS3{F0: 1}
                                                        v4 -= (^(v4))
                                                        _ = v18
                                                    }
                                                }
                                                _ = v17
                                            }
                                            v19 := GenS3{F0: 4294967295}
                                            _ = v19
                                        } else {
                                            if true {
                                                if len(v2) > 0 {
                                                    v2[0] = ((((((("" + "a") + ("hello" + (("\n\t" + "\n\t") + "a"))) + ("\n\t" + "a")) + "hello") + "foo bar") + "hello") + "\n\t")
                                                }
                                                fmt.Println(v4)
                                                v8 = (v8 >> ((501) & 63))
                                            } else {
                                                v20 := GenS1{F0: -128, F1: 100, F2: (v1 >> ((255) & 63))}
                                                _ = v20
                                            }
                                        }
                                        _ = fi14
                                    }
                                default:
                                    fmt.Println(v1, v4)
                                }
                            case 9223372036854775807:
                                if (false || true) {
                                    var v21 uint = 18446744073709551615
                                    _ = v21
                                }
                            default:
                                var v22 uint = 210
                                _ = v22
                            }
                            v23 := GenS4{F0: 811, F1: 42}
                            _ = fi12
                            _ = v23
                        }
                        for fi24 := 0; fi24 < 7; fi24++ {
                            v9 += v9
                            _ = fi24
                        }
                    }
                }
            }
        }
    }
    v25 := GenS3{F0: 675}
    if true {
        var v26 uint8 = 42
        v4 = 4294967295
        _ = v26
    } else {
        v7 -= (v7 + (v7 - v7))
    }
    delete(v6, 345)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v8, v9, v5)
    v27 := GenS3{F0: 65535}
    fmt.Println(v4)
    fmt.Println(v8, v9)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v9, v9)
    v4 += 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
