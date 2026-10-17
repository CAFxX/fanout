package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int64; F1 uint; F2 int8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 uint8; F1 uint8; F2 int16 }

func main() {
    var v1 string = ("\n\t" + (("hello" + "\n\t") + "foo bar"))
    v2 := []uint32{}
    var v3 int16 = 1
    v4 := make(map[string]bool)
    v5 := []int{}
    var v6 int16 = v3
    v7 := GenS1{F0: 2, F1: 54}
    for fi8 := 0; fi8 < 14; fi8++ {
        v6 = 42
        switch fi8 {
        case 231:
            var v9 uint = 0
            _ = v9
        case 32767:
            v10 := GenS2{F0: -1, F1: 65535, F2: 0}
            _ = v10
        default:
            var v11 int64 = -9223372036854775808
            _ = v11
        }
        switch fi8 {
        case 2147483647:
            if false {
                for fi12 := 0; fi12 < 13; fi12++ {
                    var v13 uint8 = 128
                    if true {
                        v1 = v1
                    }
                    _ = len(v5)
                    _ = fi12
                    _ = v13
                }
                v14 := GenS1{F0: 18446744073709551615, F1: v3}
                _ = v14
            } else {
                v15 := GenS3{F0: 779}
                _ = v15
            }
        case 1:
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := GenS4{F0: (42 != 1)}
                switch (v1 + v1) {
                case "foo bar":
                    if true {
                        switch v1 {
                        case "hello":
                            fmt.Println(fi8)
                        }
                        v1 = v1
                        var v18 uint32 = 2
                        switch v1 {
                        case "":
                            var v19 uint64 = 18446744073709551615
                            var v20 int64 = 42
                            _ = v19
                            _ = v20
                        case "hello":
                            var v21 float64 = 0.1
                            _ = v21
                        default:
                            fmt.Println(fi16, v18)
                        }
                        v3 = (v6 ^ v6)
                        var v22 int64 = 9223372036854775807
                        _ = v18
                        _ = v22
                    } else {
                        if false {
                            fmt.Println(fi8, v3)
                            if len(v2) > 0 {
                                v2[0] = 678
                            }
                        }
                        var v23 uint16 = 2
                        v24 := GenS4{F0: true}
                        _ = v23
                        _ = v24
                    }
                case "\n\t":
                    switch v1 {
                    case "a":
                        v25 := GenS1{F0: 4294967295, F1: (v3 + (v3 % ((v6) | 1)))}
                        _ = v25
                    }
                case "hello":
                    v1 = v1
                }
                _ = fi16
                _ = v17
            }
        }
        _ = fi8
    }
    v26 := GenS5{F0: 0, F1: (!true)}
    switch -457 {
    case -2:
        fmt.Println(v1, v6)
    }
    var v27 int8 = 127
    for fi28 := 0; fi28 < 11; fi28++ {
        v6 = int16(fi28)
        _ = fi28
    }
    if ((((((-2.25 / float64(0.1)) / 1.0) / (1.5 + ((1e10 / 0.1) + float64(0.1)))) / 1e10) <= (((1e10 * -2.25) / 0.1) * -2.25)) || false) {
        var v29 int8 = (v27 & v27)
        if (!false) {
            fmt.Println(v6, v6, v6)
            var v30 int8 = v27
            _ = len(v2)
            _ = v30
        } else {
            var v31 int64 = -291
            _ = v31
        }
        _ = v29
    }
    delete(v4, v1)
    fmt.Println(v27)
    v1 = v1
    var v32 int32 = 2147483647
    if (-9223372036854775808 <= 1) {
        var v33 float32 = 0.1
        v34 := GenS1{F0: 1, F1: v6}
        _ = v33
        _ = v34
    }
    v4[v1] = ((v27 % (((v27 % (((v27 ^ 127)) | 1))) | 1)) == (v27 | (v27 - (v27 - 127))))
    v3 = v6
    v35 := GenS6{F0: 128, F1: 42, F2: -311}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
