package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint16; F1 int64 }
func genFn0(p0 int32) bool {
    p0 = (p0 * p0)
    switch "a" {
    case "\n\t":
        v5 := GenS1{F0: 979, F1: 0}
        _ = v5
    }
    _ = p0
    return (-9223372036854775808 == 984)
}
func genFn1(p0 bool) uint32 {
    p0 = (true && false)
    var v6 uint16 = 843
    var v7 uint = 1
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := []int{fi9, fi8, fi8, fi8}
            _ = fi9
            _ = v10
        }
        v6 = v6
        if ((v6 | (v6 | 16)) > (v6 >> ((v7) & 63))) {
            if p0 {
                v11 := GenS1{F0: 42, F1: 0}
                v12 := []int{fi8, fi8}
                v6 += (^(v6))
                _ = v11
                _ = v12
            } else {
                v7 -= 42
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 uint8 = 75
                    _ = fi13
                    _ = v14
                }
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(v6, v6, p0)
                    switch fi15 {
                    case -774:
                        switch ("hello" + "a") {
                        case "a":
                            v16 := []int{fi8, 278}
                            _ = v16
                        default:
                            switch "hello" {
                            case "\n\t":
                                v7 -= (v7 | (v7 >> (((v7 << (((v7 % ((1) | 1))) & 63))) & 63)))
                            case "":
                                switch ("" + "") {
                                case "a":
                                    switch fi8 {
                                    case 9223372036854775807:
                                        v17 := []int{(fi8 - 42), (fi8 | fi15), (fi15 % ((fi8) | 1)), int(v7)}
                                        _ = v17
                                    case -9223372036854775808:
                                        if (p0 || true) {
                                            v18 := []int{}
                                            _ = v18
                                        }
                                    default:
                                        v19 := []int{(fi8 / ((fi8) | 1)), fi8}
                                        _ = v19
                                    }
                                case "foo bar":
                                    var v20 bool = (-1 < 1)
                                    _ = v20
                                case "\n\t":
                                    v6 = (^(v6))
                                }
                            case "hello":
                                v6 = v6
                            default:
                                v21 := GenS2{F0: 85}
                                _ = v21
                            }
                        }
                    case 128:
                        p0 = true
                        fmt.Println(v6, fi15)
                        var v22 uint = 2
                        _ = v22
                    }
                    _ = fi15
                }
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                if p0 {
                    v24 := GenS3{F0: (fi8 / ((9223372036854775807) | 1))}
                    _ = v24
                } else {
                    v25 := GenS4{F0: v6, F1: -485}
                    _ = v25
                }
                v7 = 2
                _ = fi23
            }
        }
        _ = fi8
    }
    var v26 int8 = 100
    _ = p0
    _ = v26
    _ = v6
    _ = v7
    return 243
}

func main() {
    var v1 uint64 = 192
    var v2 float32 = ((float32((-2.25 * 0.0)) - -2.25) + 1.5)
    var v3 uint = 65535
    var v4 bool = true
    v27 := []int{-9223372036854775808, 32768, 127, -2147483648}
    switch 2147483647 {
    case 4294967295:
        fmt.Println(v2, v3, v4)
    case 128:
        var v28 uint16 = 546
        _ = v28
    }
    _ = len(v27)
    v1 += (^(v1))
    v3 -= v3
    var v29 uint = (v3 + v3)
    var v30 uint64 = v1
    if len(v27) > 0 {
        v27[0] = 889
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
