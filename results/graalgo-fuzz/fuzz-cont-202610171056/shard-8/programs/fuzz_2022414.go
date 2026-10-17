package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 int8 }
type GenS2 struct { F0 int; F1 int16; F2 string }
type GenS3 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 string = ""
    var v2 int64 = 1
    var v3 int16 = 1000
    v4 := make(map[uint32]bool)
    if (!false) {
        v5 := GenS1{F0: 2, F1: 32767, F2: -1}
        fmt.Println(v1, v1)
        v2 += v2
        _ = v5
    }
    var v6 uint32 = 4294967295
    v7 := []int{688, 1}
    var v8 uint32 = (v6 * v6)
    fmt.Println(v1)
    var v9 uint16 = 0
    switch v1 {
    case "":
        v1 = (("a" + ((v1 + v1) + v1)) + v1)
    case "a":
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(v6, v9, v3)
            if true {
                v11 := GenS2{F0: fi10, F1: (v3 << ((42) & 63)), F2: (v1 + v1)}
                for fi12 := 0; fi12 < 15; fi12++ {
                    v6 -= v8
                    _ = fi12
                }
                _ = v11
            }
            _ = fi10
        }
    default:
        for fi13 := 0; fi13 < 11; fi13++ {
            if (((0.1 == -2.25) || false) && true) {
                for fi14 := 0; fi14 < 15; fi14++ {
                    v2 = (v2 ^ v2)
                    v7 = append(v7, 446)
                    var v15 float32 = 1.5
                    _ = fi14
                    _ = v15
                }
            } else {
                v16 := GenS1{F0: 127, F1: (v9 + v9), F2: 78}
                _ = v16
            }
            switch fi13 {
            case -9223372036854775808:
                v9 = 65535
            case 9223372036854775807:
                v17 := GenS1{F0: 128, F1: v9, F2: -1}
                _ = v17
            default:
                fmt.Println(v8)
            }
            _ = fi13
        }
    }
    if true {
        if (2 >= (^(v6))) {
            if false {
                v4[v8] = false
            }
        }
        var v18 uint32 = v8
        _ = len(v4)
        _ = v18
    } else {
        for fi19 := 0; fi19 < 1; fi19++ {
            switch v1 {
            case "a":
                switch fi19 {
                case 9223372036854775807:
                    if ((((3.14159 + -2.25) - 3.14159) != (0.1 + 3.14159)) && true) {
                        var v20 bool = true
                        for fi21 := 0; fi21 < 9; fi21++ {
                            if (("\n\t" + v1) == (v1 + v1)) {
                                v22 := GenS3{F0: 2, F1: v3, F2: 18446744073709551615}
                                v23 := GenS3{F0: 2, F1: 32767, F2: 349}
                                _ = v22
                                _ = v23
                            } else {
                                _ = len(v4)
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                var v24 float32 = -2.25
                                _ = v24
                            }
                            v3 -= 0
                            _ = fi21
                        }
                        _ = v20
                    } else {
                        v25 := GenS1{F0: 238, F1: 0, F2: -128}
                        var v26 uint8 = 255
                        var v27 uint = 2
                        _ = v25
                        _ = v26
                        _ = v27
                    }
                case -9223372036854775808:
                    v28 := GenS4{F0: 0}
                    _ = v28
                default:
                    v29 := GenS1{F0: 1, F1: (v9 << ((1) & 63)), F2: 0}
                    _ = v29
                }
            case "hello":
                for fi30 := 0; fi30 < 10; fi30++ {
                    v31 := GenS1{F0: 43, F1: (v9 / ((v9) | 1)), F2: 1}
                    v32 := GenS1{F0: 1, F1: v9, F2: int8(fi30)}
                    _ = fi30
                    _ = v31
                    _ = v32
                }
            case "":
                switch "a" {
                case "hello":
                    for fi33 := 0; fi33 < 15; fi33++ {
                        var v34 uint32 = 65535
                        v2 = -9223372036854775808
                        _ = fi33
                        _ = v34
                    }
                case "\n\t":
                    if (-2.25 != 3.14159) {
                        if false {
                            v6 = (v6 + v8)
                        }
                        v1 = v1
                        var v35 float32 = ((3.14159 * (0.1 - 1.5)) * 0.1)
                        _ = v35
                    }
                default:
                    switch (v1 + ("\n\t" + (v1 + v1))) {
                    case "\n\t":
                        v3 -= 0
                    }
                }
            }
            v3 = (v3 | v3)
            _ = fi19
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
