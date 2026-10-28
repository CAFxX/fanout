package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 uint; F1 uint8 }

func main() {
    var v1 uint32 = 2147483647
    var v2 int = 9223372036854775807
    var v3 uint16 = 32767
    var v4 int16 = 1000
    v5 := GenS1{F0: v2, F1: true, F2: 9223372036854775807}
    for fi6 := 0; fi6 < 1; fi6++ {
        v3 = v3
        v7 := GenS2{F0: 2, F1: uint64(v1)}
        _ = fi6
        _ = v7
    }
    if true {
        v8 := []int{(^(v2)), 9223372036854775807, (v2 - (v2 + (v2 * -783))), v2}
        v2 -= (v2 & (v2 + -326))
        switch "a" {
        case "hello":
            v3 -= (v3 + 585)
        default:
            switch "" {
            case "":
                v4 = (v4 & 0)
            case "a":
                for fi9 := 0; fi9 < 20; fi9++ {
                    if true {
                        v3 = (v3 - v3)
                        fmt.Println(v3, v4, v2)
                        v10 := GenS3{F0: 651, F1: 165}
                        _ = v10
                    } else {
                        v8 = append(v8, int(v3))
                        v2 = v2
                    }
                    if true {
                        var v11 int64 = 9223372036854775807
                        var v12 uint64 = 505
                        _ = v11
                        _ = v12
                    } else {
                        switch 993 {
                        case 177:
                            if true {
                                switch ("hello" + ("a" + "\n\t")) {
                                case "\n\t":
                                    var v13 uint8 = 42
                                    _ = v13
                                case "a":
                                    _ = len(v8)
                                default:
                                    if true {
                                        var v14 uint = 663
                                        var v15 uint = v14
                                        var v16 uint = (v14 / ((v14) | 1))
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    }
                                }
                                for fi17 := 0; fi17 < 7; fi17++ {
                                    var v18 int8 = 100
                                    v19 := GenS2{F0: 255, F1: 18446744073709551615}
                                    _ = fi17
                                    _ = v18
                                    _ = v19
                                }
                            }
                        }
                    }
                    _ = fi9
                }
            }
        }
        _ = v8
    }
    v3 -= uint16(v1)
    var v20 uint16 = v3
    var v21 uint16 = (v20 / ((32767) | 1))
    v22 := GenS1{F0: v2, F1: true, F2: v2}
    for fi23 := 0; fi23 < 5; fi23++ {
        v24 := GenS1{F0: 2147483647, F1: true, F2: -2147483648}
        var v25 bool = false
        v25 = (0 <= 0)
        _ = fi23
        _ = v24
        _ = v25
    }
    switch "a" {
    case "\n\t":
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := []int{-2147483648, 699, (^(fi26)), (v2 % ((4294967295) | 1))}
            _ = fi26
            _ = v27
        }
    case "":
        v1 -= (v1 - v1)
    default:
        var v28 float32 = -2.25
        _ = v28
    }
    fmt.Println(v3, v4, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
