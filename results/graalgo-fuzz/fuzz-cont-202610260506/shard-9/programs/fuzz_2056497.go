package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16; F1 string }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }

func main() {
    v1 := []uint32{1}
    v2 := []int{-799, 2147483647, 32767}
    v3 := []int{32767, -9223372036854775808, 1, 42}
    v4 := make(map[uint32]bool)
    v5 := make(map[string]int)
    v6 := make(map[uint32]string)
    var v7 int32 = -114
    var v8 uint8 = 128
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi9 := 0; fi9 < 8; fi9++ {
        v8 += v8
        var v10 int = 2147483647
        var v11 int = (v10 ^ (fi9 << ((0) & 63)))
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 bool = ((!(146 == -1)) && true)
    var v13 int8 = 100
    v14 := GenS1{F0: (^(v8)), F1: 1000, F2: -1}
    fmt.Println(v13, v13, v12)
    switch -111 {
    case 42:
        for fi15 := 0; fi15 < 17; fi15++ {
            v16 := GenS2{F0: 32, F1: ""}
            if v12 {
                v8 = 127
                for fi17 := 0; fi17 < 13; fi17++ {
                    v12 = v12
                    v8 -= v8
                    v12 = (uint32(v8) > 4294967295)
                    _ = fi17
                }
            } else {
                for fi18 := 0; fi18 < 8; fi18++ {
                    var v19 int32 = (v7 / ((v7) | 1))
                    if v12 {
                        v20 := GenS1{F0: v8, F1: 0, F2: 298}
                        _ = v20
                    }
                    var v21 uint64 = 654
                    switch "" {
                    case "a":
                        var v22 int32 = v7
                        _ = v22
                    case "foo bar":
                        var v23 uint = uint(v8)
                        _ = v23
                    }
                    v24 := GenS2{F0: 92, F1: ""}
                    _ = fi18
                    _ = v19
                    _ = v21
                    _ = v24
                }
                v7 = v7
            }
            _ = fi15
            _ = v16
        }
    case 961:
        v25 := GenS1{F0: (v8 + (v8 / ((2) | 1))), F1: 623, F2: -1}
        _ = v25
    case 1:
        for fi26 := 0; fi26 < 9; fi26++ {
            fmt.Println(v8)
            if (687 == -934) {
                var v27 uint16 = 65535
                switch fi26 {
                case 32767:
                    for fi28 := 0; fi28 < 6; fi28++ {
                        if len(v1) > 0 {
                            v1[0] = 4294967295
                        }
                        _ = fi28
                    }
                }
                _ = v27
            }
            v29 := v4[42]
            _ = fi26
            _ = v29
        }
    default:
        var v30 int16 = -32768
        _ = v30
    }
    v31 := GenS1{F0: (v8 * (^(v8))), F1: 610, F2: 42}
    if false {
        var v32 int64 = -833
        _ = v32
    } else {
        fmt.Println(v12, v7, v13)
        v33 := GenS3{F0: false, F1: -2147483648, F2: 552}
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
