package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64; F1 bool }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []uint32{2147483648, 59}
    v2 := []uint32{2147483648, 218, 42, 2147483647}
    var v3 int64 = 1
    if true {
        var v4 uint16 = 65535
        v2 = append(v2, 2)
        switch "" {
        case "":
            var v5 uint16 = 2
            _ = v5
        case "\n\t":
            if true {
                v6 := GenS1{F0: -28}
                v7 := GenS2{F0: 2, F1: 42, F2: -1}
                var v8 uint8 = 42
                _ = v6
                _ = v7
                _ = v8
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v9 uint64 = 1
                v3 += (v3 / (((v3 & v3)) | 1))
                _ = v9
            }
        case "hello":
            if (-2.25 > ((1e10 * (0.0 / -2.25)) + -2.25)) {
                _ = len(v2)
                for fi10 := 0; fi10 < 4; fi10++ {
                    _ = len(v2)
                    v3 = -9223372036854775808
                    if true {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            v12 := GenS1{F0: -128}
                            var v13 float64 = 1e10
                            _ = fi11
                            _ = v12
                            _ = v13
                        }
                        v3 -= v3
                    }
                    _ = fi10
                }
                for fi14 := 0; fi14 < 7; fi14++ {
                    if false {
                        var v15 int = -1
                        if false {
                            var v16 uint32 = 148
                            _ = v16
                        } else {
                            var v17 int64 = (v3 + 9223372036854775807)
                            fmt.Println(v17, v3, v3)
                            var v18 uint16 = v4
                            _ = v17
                            _ = v18
                        }
                        _ = v15
                    } else {
                        var v19 int16 = 42
                        _ = v19
                    }
                    _ = fi14
                }
            }
        default:
            v3 = v3
        }
        _ = v4
    }
    if false {
        if true {
            var v20 uint8 = 0
            for fi21 := 0; fi21 < 4; fi21++ {
                v3 = (v3 & -9223372036854775808)
                fmt.Println(v20)
                _ = fi21
            }
            var v22 uint8 = 1
            _ = v20
            _ = v22
        }
        v3 = (v3 * v3)
        var v23 int16 = 32767
        _ = v23
    }
    fmt.Println(v3)
    var v24 int64 = 428
    fmt.Println(v3, v24)
    v24 += (^(v24))
    if (!true) {
        v3 -= (v3 & (v24 - (-(v24))))
        v25 := GenS1{F0: -128}
        fmt.Println(v24, v3)
        _ = v25
    }
    var v26 int64 = v24
    v27 := GenS3{F0: 0, F1: (4294967295 != 42)}
    var v28 string = ("" + "")
    v29 := GenS4{F0: 729}
    fmt.Println(v3)
    v3 = v26
    v28 = ("hello" + (v28 + v28))
    _ = len(v2)
    for fi30 := 0; fi30 < 18; fi30++ {
        _ = len(v1)
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if ((fi30 < (fi30 % ((fi30) | 1))) || false) {
            fmt.Println(v24, fi30, v26)
        }
        _ = fi30
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
