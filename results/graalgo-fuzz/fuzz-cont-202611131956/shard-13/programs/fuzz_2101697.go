package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint32; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 int8; F1 int32; F2 int16 }
type GenS8 struct { F0 string; F1 int64; F2 uint32 }

func main() {
    var v1 uint8 = 127
    var v2 uint = 781
    var v3 uint8 = v1
    v4 := []int{-174, -833, 4294967295}
    if len(v4) > 0 {
        v4[0] = 32767
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        if len(v4) > 0 {
            v4[0] = 0
        }
        v6 := GenS1{F0: 2147483647}
        _ = fi5
        _ = v6
    }
    switch "" {
    case "a":
        if false {
            _ = len(v4)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v2, v2, v2)
        }
    }
    if len(v4) > 0 {
        v4[0] = 32768
    }
    var v7 int16 = 42
    switch "" {
    case "foo bar":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        v8 := GenS2{F0: -9223372036854775808}
        _ = v8
    case "":
        v2 += (v2 ^ 1)
    default:
        v9 := GenS2{F0: 816}
        _ = v9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 12; fi11++ {
            var v12 int = fi10
            _ = fi11
            _ = v12
        }
        switch (fi10 % ((fi10) | 1)) {
        case 4294967295:
            fmt.Println(v7)
        case 402:
            v7 = v7
        }
        _ = fi10
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        v14 := GenS3{F0: 42}
        for fi15 := 0; fi15 < 19; fi15++ {
            switch "a" {
            case "hello":
                v2 -= v2
            case "\n\t":
                v3 -= (v1 | (v3 / (((v1 | (v1 / ((v1) | 1)))) | 1)))
            default:
                v16 := GenS4{F0: (v2 * v2)}
                v17 := GenS5{F0: 1, F1: 25, F2: v2}
                _ = len(v4)
                v3 = (v1 >> (((v2 | (^(v2)))) & 63))
                v2 = (v2 ^ v2)
                if (1.5 >= 0.1) {
                    var v18 int16 = (v7 * (v7 << ((430) & 63)))
                    v19 := GenS6{F0: (v2 >> ((4294967295) & 63))}
                    _ = v18
                    _ = v19
                }
                v7 = v7
                v20 := GenS1{F0: -1}
                switch fi15 {
                case 1:
                    v7 -= 0
                case 127:
                    if true {
                        if true {
                            fmt.Println(fi15, fi15)
                            var v21 bool = true
                            switch -584 {
                            case -9223372036854775808:
                                var v22 bool = true
                                _ = v22
                            }
                            _ = v21
                        } else {
                            fmt.Println(v7, v7)
                            v3 = v1
                            var v23 uint16 = 1
                            _ = v23
                        }
                        var v24 int32 = 1
                        _ = v24
                    }
                case 991:
                    v25 := GenS7{F0: -128, F1: 166, F2: v7}
                    _ = v25
                default:
                    switch -1 {
                    case 965:
                        v26 := GenS8{F0: "", F1: -1, F2: 945}
                        _ = v26
                    case 9223372036854775807:
                        var v27 uint64 = 884
                        _ = v27
                    case 32767:
                        fmt.Println(v3, v7, v2)
                    }
                }
                v3 = (v1 << ((4294967295) & 63))
                _ = v16
                _ = v17
                _ = v20
            }
            var v28 uint32 = 713
            _ = fi15
            _ = v28
        }
        _ = fi13
        _ = v14
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
