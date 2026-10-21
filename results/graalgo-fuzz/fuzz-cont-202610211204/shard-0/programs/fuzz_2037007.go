package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 int64 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint32 }
type GenS4 struct { F0 uint64; F1 uint }
type GenS5 struct { F0 int8; F1 int; F2 int8 }

func main() {
    var v1 uint = 302
    var v2 uint32 = 194
    var v3 int = 700
    var v4 uint8 = 207
    v5 := []int{v3, 32768}
    var v6 uint8 = v4
    if false {
        fmt.Println(v4, v6, v2)
        v3 -= (v3 & v3)
    } else {
        v6 = uint8(v2)
    }
    if true {
        var v7 string = (("\n\t" + "hello") + (("hello" + ("" + "a")) + ""))
        if false {
            fmt.Println(v3, v2, v2)
        } else {
            if false {
                v8 := GenS1{F0: uint64(v2), F1: uint32(v1), F2: 9223372036854775807}
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                v6 = (v6 << ((v1) & 63))
                _ = v8
            }
            var v9 int16 = int16(v2)
            _ = v9
        }
        _ = v7
    } else {
        if (v6 <= 1) {
            switch "foo bar" {
            case "foo bar":
                v6 += 0
            default:
                v3 = 2147483647
            }
        } else {
            var v10 float64 = 3.14159
            v10 += 3.14159
            for fi11 := 0; fi11 < 5; fi11++ {
                v12 := GenS1{F0: 18446744073709551615, F1: v2, F2: 642}
                for fi13 := 0; fi13 < 3; fi13++ {
                    if true {
                        var v14 int = -778
                        fmt.Println(v14, v3, v3)
                        v15 := GenS2{F0: "a"}
                        _ = v14
                        _ = v15
                    } else {
                        v10 -= ((v10 - 0.0) - float64(-2.25))
                        var v16 uint64 = 4294967295
                        var v17 uint64 = v16
                        _ = v16
                        _ = v17
                    }
                    if true {
                        switch fi11 {
                        case 699:
                            v5 = append(v5, fi11)
                        case -562:
                            v18 := GenS1{F0: 132, F1: v2, F2: -525}
                            _ = v18
                        case -9223372036854775808:
                            switch ((("\n\t" + "hello") + ("" + "\n\t")) + "foo bar") {
                            case "foo bar":
                                fmt.Println(v1, fi11, v6)
                            case "hello":
                                if (3.14159 >= 0.0) {
                                    v19 := GenS1{F0: 1, F1: 4294967295, F2: 0}
                                    _ = v19
                                } else {
                                    var v20 int32 = -816
                                    if len(v5) > 0 {
                                        v5[0] = fi13
                                    }
                                    fmt.Println(v20, fi13, fi11)
                                    _ = v20
                                }
                            case "":
                                if true {
                                    _ = len(v5)
                                    switch "\n\t" {
                                    case "foo bar":
                                        if (1 > 152) {
                                            v10 += v10
                                            v10 -= v10
                                        } else {
                                            fmt.Println(fi13, fi11, fi11)
                                            var v21 int8 = -1
                                            v22 := GenS3{F0: v1, F1: (v2 >> (((^(v1))) & 63))}
                                            _ = v21
                                            _ = v22
                                        }
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                    v10 += 1e10
                                } else {
                                    v2 += v2
                                }
                            }
                        }
                        var v23 float32 = 1.5
                        v4 = 162
                        _ = v23
                    }
                    fmt.Println(v2)
                    _ = fi13
                }
                _ = fi11
                _ = v12
            }
            _ = v10
        }
        switch -196 {
        case 9223372036854775807:
            for fi24 := 0; fi24 < 13; fi24++ {
                v25 := GenS4{F0: 42, F1: (v1 / (((v1 | v1)) | 1))}
                var v26 int8 = 1
                fmt.Println(v3, v6, v3)
                _ = fi24
                _ = v25
                _ = v26
            }
        case 0:
            if (((106 < 1) && true) || false) {
                var v27 int64 = -311
                v28 := GenS5{F0: -128, F1: int(v1), F2: 1}
                v29 := GenS3{F0: v1, F1: (v2 - 0)}
                _ = v27
                _ = v28
                _ = v29
            }
        case 127:
            var v30 int64 = 922
            _ = v30
        default:
            var v31 uint32 = v2
            _ = v31
        }
        var v32 int8 = 0
        _ = v32
    }
    v1 = uint(v2)
    v5 = append(v5, (v3 & 128))
    var v33 uint16 = 0
    fmt.Println(v33)
    v6 += v6
    var v34 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
