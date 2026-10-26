package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 uint32, p1 uint16) uint64 {
    var v5 int64 = int64(p1)
    p0 = 0
    p1 = (p1 | p1)
    _ = p0
    _ = p1
    _ = v5
    return 940
}
func genFn1(p0 bool, p1 int64) uint8 {
    p0 = p0
    var v6 int = 127
    _ = p0
    _ = p1
    _ = v6
    return 42
}

func main() {
    var v1 uint16 = 42
    v2 := make(map[uint32]string)
    var v3 uint16 = 860
    var v4 uint32 = 4294967295
    if (!true) {
        v7 := GenS1{F0: -32768}
        _ = v7
    } else {
        var v8 float64 = (-2.25 + 1e10)
        _ = v8
    }
    if true {
        if true {
            v1 -= v3
        } else {
            v9 := GenS1{F0: 1000}
            for fi10 := 0; fi10 < 7; fi10++ {
                for fi11 := 0; fi11 < 12; fi11++ {
                    switch -428 {
                    case -9223372036854775808:
                        fmt.Println(fi11, fi11, v4)
                    case 9223372036854775807:
                        fmt.Println(fi10)
                        switch "\n\t" {
                        case "foo bar":
                            var v12 uint16 = v1
                            _ = v12
                        }
                        delete(v2, uint32(v1))
                        v13 := GenS1{F0: 42}
                        _ = v13
                    default:
                        if true {
                            var v14 int64 = -94
                            v4 = 42
                            if true {
                                v15 := v2[v4]
                                switch (fi11 >> ((uint(v4)) & 63)) {
                                case 0:
                                    var v16 int = fi10
                                    _ = v16
                                default:
                                    _ = len(v2)
                                }
                                v17 := []int{}
                                fmt.Println(fi11, v4)
                                _ = v15
                                _ = v17
                            }
                            _ = v14
                        } else {
                            switch fi11 {
                            case -798:
                                v3 -= (v3 - (v3 >> ((0) & 63)))
                            case 2147483647:
                                if false {
                                    v3 -= (v1 & (v3 / (((v1 / (((v3 / ((42) | 1))) | 1))) | 1)))
                                } else {
                                    switch -736 {
                                    case -370:
                                        v3 = v1
                                    case -2147483648:
                                        v4 = (v4 >> ((277) & 63))
                                    case -296:
                                        v3 += 199
                                    }
                                    v18 := GenS1{F0: 1000}
                                    _ = v18
                                }
                            }
                            v4 += v4
                        }
                    }
                    var v19 int16 = -1
                    if (9223372036854775807 >= -362) {
                        v20 := GenS2{F0: -1}
                        v21 := GenS3{F0: 18446744073709551615, F1: 0}
                        _ = v20
                        _ = v21
                    }
                    v22 := v2[v4]
                    _ = fi11
                    _ = v19
                    _ = v22
                }
                v3 -= (^(v3))
                _ = fi10
            }
            var v23 int8 = 100
            _ = v23
            _ = v9
        }
    }
    v4 = v4
    v24 := []int{}
    if (3.14159 != -2.25) {
        fmt.Println(v4, v1, v3)
        v25 := GenS1{F0: -32768}
        _ = v25
    } else {
        if ((("foo bar" + "hello") + "\n\t") > "foo bar") {
            v1 = v1
            if len(v24) >= 2 {
                v24 = v24[1:len(v24)]
            }
        }
        v4 += v4
        v26 := GenS4{F0: -32768}
        _ = v26
    }
    v3 = v3
    v2[v4] = "hello"
    v27 := GenS1{F0: 0}
    v28 := GenS5{F0: (v4 | (v4 ^ (^(v4)))), F1: 128}
    v3 = (^(v1))
    var v29 int8 = -1
    for fi30 := 0; fi30 < 15; fi30++ {
        switch "" {
        case "hello":
            var v31 int32 = 2147483647
            _ = v31
        }
        if (-2.25 > -2.25) {
            v24 = append(v24, 76)
            var v32 int32 = -496
            _ = v32
        } else {
            fmt.Println(v4, v1)
            v1 = v1
        }
        fmt.Println(v29)
        _ = fi30
    }
    v29 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
}
