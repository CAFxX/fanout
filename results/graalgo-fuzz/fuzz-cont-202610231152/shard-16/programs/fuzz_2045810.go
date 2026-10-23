package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 bool }
type GenS2 struct { F0 int8; F1 int32 }
type GenS3 struct { F0 uint; F1 uint; F2 int16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 int }

func main() {
    v1 := []uint32{0, 860}
    var v2 uint32 = 2
    v3 := []uint32{(v2 & (v2 + (v2 / ((385) | 1))))}
    v4 := GenS1{F0: 2147483647, F1: 253, F2: true}
    if (2147483648 == v2) {
        if (!(-32768 <= -32768)) {
            for fi5 := 0; fi5 < 10; fi5++ {
                if false {
                    v6 := GenS2{F0: 42, F1: 2147483647}
                    v7 := GenS3{F0: 1, F1: 2, F2: 778}
                    _ = v6
                    _ = v7
                }
                if ((^(v2)) != 1) {
                    if true {
                        _ = len(v3)
                        fmt.Println(fi5)
                    } else {
                        v2 = v2
                        switch fi5 {
                        case 128:
                            v8 := GenS1{F0: 42, F1: 127, F2: true}
                            _ = v8
                        case 32767:
                            for fi9 := 0; fi9 < 18; fi9++ {
                                switch (((("a" + "a") + "a") + (("\n\t" + "\n\t") + "a")) + "\n\t") {
                                case "foo bar":
                                    v2 -= 1
                                case "hello":
                                    fmt.Println(fi9, v2)
                                }
                                fmt.Println(fi5)
                                v2 = (^(v2))
                                _ = fi9
                            }
                        default:
                            switch -9223372036854775808 {
                            case 1:
                                var v10 int64 = -9223372036854775808
                                _ = v10
                            case -2147483648:
                                fmt.Println(v2, v2)
                            }
                        }
                        fmt.Println(fi5, v2)
                    }
                }
                _ = fi5
            }
            v2 = v2
            v2 -= 1
        }
        if (true && (float32(1e10) == 1e10)) {
            v2 = 124
        } else {
            switch 0 {
            case -719:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if (-2147483648 < 0) {
                    v11 := GenS4{F0: 0, F1: 2, F2: 42}
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = v11
                }
            }
        }
    } else {
        var v12 int = 32768
        var v13 uint = 1
        _ = v12
        _ = v13
    }
    v14 := GenS4{F0: (v2 ^ 65535), F1: 42, F2: 42}
    if len(v3) > 0 {
        v3[0] = v2
    }
    v2 = 1
    v2 += v2
    var v15 uint16 = 0
    if false {
        if len(v3) > 0 {
            v3[0] = v2
        }
    }
    v2 += v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
