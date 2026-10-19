package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64; F1 bool; F2 int32 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0() uint16 {
    var v4 float64 = float64(0.1)
    var v5 int32 = -53
    var v6 int32 = v5
    v4 -= v4
    fmt.Println(v4, v4, v6)
    _ = v4
    _ = v5
    _ = v6
    return 65535
}
func genFn1(p0 int64, p1 uint16, p2 int16) uint {
    v7 := []int{574, 32767, -1, 9223372036854775807}
    p1 += genFn0()
    v8 := GenS1{F0: 2147483647, F1: 674, F2: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 693
    v2 := []int{-1, 42}
    var v3 int64 = -351
    for fi9 := 0; fi9 < 16; fi9++ {
        switch "hello" {
        case "\n\t":
            for fi10 := 0; fi10 < 5; fi10++ {
                if true {
                    var v11 float32 = -2.25
                    _ = v11
                }
                var v12 int8 = -128
                var v13 int32 = 666
                _ = fi10
                _ = v12
                _ = v13
            }
        case "":
            var v14 int8 = -128
            _ = v14
        }
        v15 := GenS2{F0: 0, F1: ("\n\t" <= "hello"), F2: 465}
        _ = fi9
        _ = v15
    }
    var v16 float32 = ((((0.0 - 0.1) * 3.14159) / 1.0) - 1e10)
    v2 = append(v2, -2147483648)
    var v17 uint = 4294967295
    switch "hello" {
    case "hello":
        for fi18 := 0; fi18 < 9; fi18++ {
            fmt.Println(v3, v3, v16)
            if true {
                if true {
                    for fi19 := 0; fi19 < 14; fi19++ {
                        if false {
                            v20 := GenS2{F0: 759, F1: (true || false), F2: -742}
                            if ("\n\t" != "") {
                                var v21 float32 = v16
                                _ = v21
                            } else {
                                if true {
                                    if (genFn0() == 621) {
                                        fmt.Println(v1)
                                    }
                                }
                            }
                            v16 -= float32(1.5)
                            _ = v20
                        } else {
                            v2 = append(v2, 762)
                        }
                        var v22 int8 = -1
                        _ = fi19
                        _ = v22
                    }
                    v23 := GenS2{F0: 2, F1: true, F2: -2147483648}
                    v1 = 2
                    _ = v23
                }
            } else {
                var v24 uint8 = 1
                var v25 uint = (v17 & v17)
                _ = v24
                _ = v25
            }
            var v26 string = "\n\t"
            _ = fi18
            _ = v26
        }
    }
    fmt.Println(v16, v3)
    _ = len(v2)
    var v27 int32 = 257
    var v28 uint16 = genFn0()
    fmt.Println(v3, v28)
    switch 2147483647 {
    case 0:
        v29 := GenS3{F0: 1, F1: int16(v28)}
        _ = v29
    case 32768:
        v30 := GenS1{F0: (v1 % (((v1 | (^(v1)))) | 1)), F1: genFn1((v3 & -1), 2, -111), F2: 1}
        _ = v30
    }
    v1 = 1
    v28 = (v28 / ((uint16(v1)) | 1))
    fmt.Println(v3, v27, v16)
    switch "\n\t" {
    case "a":
        v3 = (v3 << (((v17 / (((v17 >> (((v17 % (((^(v17))) | 1))) & 63))) | 1))) & 63))
    case "foo bar":
        fmt.Println(v1, v3)
    default:
        for fi31 := 0; fi31 < 14; fi31++ {
            _ = len(v2)
            _ = fi31
        }
    }
    var v32 float32 = v16
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS3{F0: 18446744073709551615, F1: -1}
        v27 += v27
        _ = fi33
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v32
}
