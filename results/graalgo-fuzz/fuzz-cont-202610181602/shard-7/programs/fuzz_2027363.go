package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 int32; F2 uint32 }

func main() {
    var v1 uint16 = 169
    var v2 uint32 = 4294967295
    var v3 uint32 = v2
    v4 := []uint32{v3, v2, (v3 >> ((334) & 63)), 2147483647}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v5 := GenS1{F0: v1}
    var v6 int = 127
    if true {
        v7 := GenS2{F0: 1, F1: -2147483648, F2: (v2 % ((v2) | 1))}
        for fi8 := 0; fi8 < 1; fi8++ {
            for fi9 := 0; fi9 < 3; fi9++ {
                v10 := GenS1{F0: v1}
                if len(v4) > 0 {
                    v4[0] = v3
                }
                v3 = (v2 % (((v3 & (^(v2)))) | 1))
                _ = fi9
                _ = v10
            }
            if (64 >= 103) {
                var v11 string = ("" + ("hello" + "foo bar"))
                if false {
                    if true {
                        for fi12 := 0; fi12 < 11; fi12++ {
                            if true {
                                switch (fi12 * fi12) {
                                case 128:
                                    if (!(true && false)) {
                                        var v13 uint = 797
                                        v1 = (v1 - v1)
                                        v6 = (v6 ^ v6)
                                        _ = v13
                                    }
                                case -426:
                                    var v14 float32 = 0.0
                                    var v15 int = (v6 / (((fi12 ^ (v6 / ((fi12) | 1)))) | 1))
                                    _ = v14
                                    _ = v15
                                case 0:
                                    v16 := GenS2{F0: 407, F1: -413, F2: v3}
                                    _ = v16
                                default:
                                    var v17 float32 = 1.5
                                    _ = v17
                                }
                                fmt.Println(v6, v11)
                            } else {
                                var v18 int64 = 889
                                _ = v18
                            }
                            _ = fi12
                        }
                        v11 = "hello"
                        var v19 int16 = -32768
                        _ = v19
                    }
                    var v20 bool = true
                    v3 -= (v2 << ((954) & 63))
                    _ = v20
                } else {
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 int8 = -117
                        switch v11 {
                        case "\n\t":
                            var v23 uint8 = 196
                            _ = v23
                        case "foo bar":
                            v6 = -858
                            switch "hello" {
                            case "a":
                                fmt.Println(fi21)
                            }
                            fmt.Println(v1)
                        default:
                            if false {
                                v24 := GenS1{F0: v1}
                                _ = v24
                            } else {
                                fmt.Println(v2, v11, v2)
                                v22 += (v22 >> ((1) & 63))
                                switch fi8 {
                                case 4294967295:
                                    _ = len(v4)
                                case -698:
                                    var v25 int64 = -1
                                    _ = v25
                                default:
                                    v26 := GenS1{F0: (v1 - (v1 ^ v1))}
                                    _ = v26
                                }
                            }
                        }
                        _ = fi21
                        _ = v22
                    }
                }
                v27 := GenS2{F0: 42, F1: int32(v2), F2: uint32(v1)}
                _ = v11
                _ = v27
            }
            v28 := GenS1{F0: 65535}
            _ = fi8
            _ = v28
        }
        _ = v7
    }
    var v29 int16 = 0
    var v30 int = 9223372036854775807
    if (!false) {
        fmt.Println(v3, v29, v29)
        v3 -= v3
        switch -838 {
        case 0:
            v30 = (v30 >> ((65535) & 63))
        default:
            fmt.Println(v1, v2)
        }
    }
    for fi31 := 0; fi31 < 13; fi31++ {
        _ = len(v4)
        switch "\n\t" {
        case "\n\t":
            v3 += v3
        case "foo bar":
            var v32 int16 = (v29 % ((v29) | 1))
            _ = v32
        default:
            for fi33 := 0; fi33 < 8; fi33++ {
                var v34 uint = 266
                _ = fi33
                _ = v34
            }
        }
        v29 += (v29 << ((150) & 63))
        _ = fi31
    }
    v35 := GenS1{F0: (v1 % (((v1 + 673)) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
