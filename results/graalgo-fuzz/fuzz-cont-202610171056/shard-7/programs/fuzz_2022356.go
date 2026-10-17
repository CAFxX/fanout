package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int; F1 uint32 }
type GenS5 struct { F0 int64; F1 int64; F2 int32 }
type GenS6 struct { F0 int16; F1 int32 }

func main() {
    var v1 uint16 = 411
    v2 := []int{-379, -146, -715, 1}
    var v3 int8 = 25
    var v4 uint16 = (v1 << ((42) & 63))
    var v5 bool = (1 > -963)
    var v6 uint8 = 127
    for fi7 := 0; fi7 < 15; fi7++ {
        v1 = (v1 / ((v1) | 1))
        _ = fi7
    }
    var v8 uint16 = v4
    v5 = v5
    var v9 uint = 42
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := GenS1{F0: v6, F1: uint(v6), F2: (v6 * (^(v6)))}
        switch ("\n\t" + "") {
        case "hello":
            v12 := GenS2{F0: (v3 ^ v3), F1: v1}
            _ = v12
        case "a":
            if false {
                switch fi10 {
                case -952:
                    for fi13 := 0; fi13 < 19; fi13++ {
                        fmt.Println(v5, v4, v5)
                        _ = fi13
                    }
                case 793:
                    fmt.Println(v5, v4)
                case 42:
                    switch (("\n\t" + "a") + "foo bar") {
                    case "":
                        switch "\n\t" {
                        case "":
                            var v14 uint32 = 65535
                            _ = v14
                        case "hello":
                            v5 = (42 == 1)
                        }
                    case "foo bar":
                        v15 := GenS3{F0: uint16(v6)}
                        _ = v15
                    case "hello":
                        v16 := GenS4{F0: (fi10 >> (((^(v9))) & 63)), F1: 0}
                        _ = v16
                    }
                default:
                    if (32767 != fi10) {
                        if false {
                            v17 := GenS1{F0: (v6 >> (((v9 | v9)) & 63)), F1: v9, F2: uint8(v4)}
                            fmt.Println(fi10)
                            fmt.Println(v1, v6, v5)
                            _ = v17
                        } else {
                            v5 = false
                            fmt.Println(v3, v5)
                        }
                        v18 := GenS2{F0: v3, F1: 0}
                        _ = v18
                    } else {
                        var v19 bool = v5
                        _ = v19
                    }
                }
            } else {
                v20 := GenS4{F0: fi10, F1: uint32(v9)}
                _ = len(v2)
                if v5 {
                    v21 := GenS5{F0: 555, F1: 437, F2: 42}
                    fmt.Println(v4)
                    _ = v21
                } else {
                    switch 529 {
                    case 642:
                        var v22 int32 = 318
                        _ = v22
                    case -971:
                        _ = len(v2)
                    case 128:
                        v23 := GenS6{F0: 0, F1: int32(v3)}
                        _ = v23
                    default:
                        var v24 bool = v5
                        _ = v24
                    }
                    v4 -= (v1 << ((215) & 63))
                }
                _ = v20
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        _ = fi10
        _ = v11
    }
    var v25 int32 = -2147483648
    switch "" {
    case "":
        v9 = v9
    case "foo bar":
        fmt.Println(v8)
    case "a":
        fmt.Println(v1)
    default:
        if (v5 || (v5 && v5)) {
            v9 -= v9
            if false {
                v4 += v8
                v4 += v4
                if ("" < ("" + ("" + ("hello" + "hello")))) {
                    v2 = append(v2, 1)
                    switch -9223372036854775808 {
                    case -851:
                        v2 = append(v2, -9223372036854775808)
                    case 826:
                        var v26 uint16 = 65535
                        _ = v26
                    case 2147483647:
                        fmt.Println(v9, v3, v6)
                    }
                    v8 = v8
                }
            }
        }
    }
    switch 32767 {
    case -2147483648:
        v1 -= (v8 | v8)
    default:
        var v27 uint64 = 1
        _ = v27
    }
    fmt.Println(v6)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
