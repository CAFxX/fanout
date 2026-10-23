package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint }
func genFn0() uint32 {
    v5 := []int{-2147483648}
    switch "foo bar" {
    case "hello":
        var v6 uint32 = 2147483648
        _ = v6
    }
    _ = v5
    return 0
}
func genFn1(p0 uint16) uint8 {
    var v7 float64 = -2.25
    p0 -= 1
    _ = p0
    _ = v7
    return 13
}

func main() {
    var v1 int32 = -1
    var v2 bool = true
    var v3 uint64 = 42
    var v4 uint8 = 73
    v3 += v3
    v8 := []int{271, 32768, 9223372036854775807, 4294967295}
    if (!false) {
        switch "" {
        case "\n\t":
            if len(v8) > 0 {
                v8[0] = 32768
            }
        case "hello":
            switch "hello" {
            case "a":
                if v2 {
                    for fi9 := 0; fi9 < 1; fi9++ {
                        if len(v8) > 0 {
                            v8[0] = (fi9 ^ (fi9 - 819))
                        }
                        v4 = (v4 & v4)
                        if (474 != -1) {
                            switch 2147483647 {
                            case -1:
                                fmt.Println(v1, v4, v2)
                            case 159:
                                for fi10 := 0; fi10 < 19; fi10++ {
                                    if true {
                                        v11 := GenS1{F0: 0, F1: 9223372036854775807, F2: v4}
                                        var v12 int16 = 734
                                        v13 := GenS1{F0: -9223372036854775808, F1: 9223372036854775807, F2: genFn1(2)}
                                        _ = v11
                                        _ = v12
                                        _ = v13
                                    }
                                    v14 := GenS2{F0: int8(v3), F1: -1}
                                    _ = fi10
                                    _ = v14
                                }
                            default:
                                var v15 uint = 4294967295
                                _ = v15
                            }
                        }
                        _ = fi9
                    }
                    var v16 bool = ("a" >= "")
                    _ = v16
                } else {
                    fmt.Println(v3, v1, v2)
                }
            case "":
                var v17 int16 = 1000
                _ = v17
            case "\n\t":
                v4 += v4
            default:
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
            }
        default:
            v1 += v1
        }
    }
    if true {
        v1 = v1
        if len(v8) > 0 {
            v8[0] = 32767
        }
    } else {
        for fi18 := 0; fi18 < 16; fi18++ {
            switch fi18 {
            case -164:
                var v19 int32 = 362
                _ = v19
            case 363:
                var v20 int32 = (v1 % (((-(v1))) | 1))
                _ = v20
            case 9223372036854775807:
                if ((v1 / (((v1 & (v1 + (v1 << ((255) & 63))))) | 1)) < 1) {
                    if (42 >= 32767) {
                        fmt.Println(fi18)
                        switch "a" {
                        case "":
                            if len(v8) > 0 {
                                v8[0] = fi18
                            }
                        case "hello":
                            v21 := GenS1{F0: 114, F1: 9223372036854775807, F2: genFn1(uint16(v3))}
                            _ = v21
                        default:
                            v22 := GenS2{F0: 42, F1: -1}
                            _ = v22
                        }
                        v23 := GenS1{F0: -1, F1: 9223372036854775807, F2: v4}
                        _ = v23
                    } else {
                        var v24 bool = (true && v2)
                        fmt.Println(v4, v1)
                        _ = v24
                    }
                    v25 := GenS3{F0: 4294967295}
                    _ = v25
                }
            }
            _ = fi18
        }
        var v26 int8 = 100
        v8 = append(v8, -9223372036854775808)
        _ = v26
    }
    for fi27 := 0; fi27 < 15; fi27++ {
        v2 = v2
        _ = fi27
    }
    v4 -= 255
    var v28 uint64 = v3
    if false {
        var v29 float32 = 0.1
        _ = v29
    }
    v30 := GenS2{F0: 0, F1: 127}
    fmt.Println(v28, v4)
    var v31 int16 = 213
    for fi32 := 0; fi32 < 11; fi32++ {
        if (false || v2) {
            var v33 float32 = float32(float64(1.5))
            var v34 uint = 2
            _ = v33
            _ = v34
        } else {
            if (1e10 == ((3.14159 - ((3.14159 + -2.25) - 1.5)) / 1e10)) {
                fmt.Println(fi32)
                if len(v8) > 0 {
                    v8[0] = (fi32 << ((42) & 63))
                }
            }
            v2 = ((v31 << ((178) & 63)) == v31)
        }
        _ = fi32
    }
    var v35 int16 = v31
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v8
}
