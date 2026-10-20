package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 int8; F1 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[string]bool)
    var v3 int16 = 0
    var v4 uint16 = 0
    var v5 int64 = -1
    v4 = v4
    var v6 bool = false
    v7 := GenS1{F0: 925, F1: 42}
    v1[(("" + "hello") + "a")] = (!(v4 <= 255))
    var v8 bool = false
    for fi9 := 0; fi9 < 19; fi9++ {
        fmt.Println(v5, v6, v3)
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int32 = 890
            var v12 float32 = ((0.0 * float32((1e10 / 0.1))) - 0.1)
            v5 += (v5 ^ int64(fi10))
            v12 = (v12 + float32(0.1))
            var v13 float32 = 3.14159
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            fmt.Println(v4, v5)
            switch ("a" + (("" + "") + "foo bar")) {
            case "":
                var v16 uint16 = v4
                _ = v16
            default:
                if v6 {
                    v17 := GenS2{F0: v8, F1: 0, F2: v3}
                    var v18 string = ""
                    v19 := []int{(fi14 - fi15), (^(fi14)), (-(fi15))}
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    if false {
                        fmt.Println(fi14, v6, v5)
                        v8 = (v4 != (^(v4)))
                    } else {
                        fmt.Println(fi14, fi15)
                        v20 := GenS2{F0: false, F1: 0, F2: (v3 << ((14) & 63))}
                        v3 = v3
                        v21 := GenS2{F0: (0 > 127), F1: uint64(v4), F2: (v3 - (v3 + v3))}
                        v22 := GenS1{F0: 0, F1: 32767}
                        if false {
                            v23 := GenS1{F0: -69, F1: (fi15 % ((fi14) | 1))}
                            _ = v23
                        }
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                }
            }
            _ = fi15
        }
        var v24 string = "foo bar"
        switch v24 {
        case "hello":
            var v25 float32 = (1e10 + (1.5 - (((0.0 + 3.14159) - 3.14159) - 1e10)))
            _ = v25
        default:
            v8 = true
        }
        _ = fi14
        _ = v24
    }
    for fi26 := 0; fi26 < 8; fi26++ {
        v4 += v4
        switch "hello" {
        case "hello":
            switch fi26 {
            case 9223372036854775807:
                var v27 int32 = -1
                _ = v27
            case -9223372036854775808:
                delete(v1, ("" + "hello"))
            default:
                fmt.Println(v3, v6)
            }
        case "\n\t":
            var v28 uint64 = 1
            _ = v28
        }
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 int8 = 2
            v31 := GenS3{F0: (v30 & (v30 * 1)), F1: 18446744073709551615}
            _ = fi29
            _ = v30
            _ = v31
        }
        _ = fi26
    }
    fmt.Println(v8, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
