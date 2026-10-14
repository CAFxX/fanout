package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 uint64 }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint8; F1 int8 }
func genFn0(p0 uint32) uint16 {
    p0 -= (p0 & p0)
    if false {
        var v5 int64 = 9223372036854775807
        _ = v5
    }
    _ = p0
    return 933
}

func main() {
    var v1 uint16 = 454
    v2 := make(map[uint32]string)
    var v3 int = -1
    var v4 int32 = -763
    v6 := GenS1{F0: 1000, F1: 1, F2: 404}
    v7 := []int{(v3 >> ((4294967295) & 63)), v3, v3, v3}
    var v8 string = "foo bar"
    v7 = append(v7, (v3 + v3))
    if (128 < 0) {
        var v9 float32 = 3.14159
        _ = v9
    }
    switch (v3 % (((v3 + (-(v3)))) | 1)) {
    case 4294967295:
        var v10 float32 = 3.14159
        _ = v10
    case 42:
        if true {
            for fi11 := 0; fi11 < 2; fi11++ {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v13 := GenS2{F0: 918, F1: -1}
                    var v14 bool = (252 <= 2)
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if (!(true && (fi11 < int(v1)))) {
                    v15 := GenS3{F0: 128, F1: 0}
                    var v16 string = (v8 + v8)
                    _ = v15
                    _ = v16
                } else {
                    v1 -= genFn0(4294967295)
                }
                v17 := GenS3{F0: 1, F1: 1}
                _ = fi11
                _ = v17
            }
        } else {
            switch (v3 % ((-9223372036854775808) | 1)) {
            case 9223372036854775807:
                switch v8 {
                case "":
                    if false {
                        var v18 string = ((("" + v8) + v8) + ("hello" + v8))
                        var v19 int8 = 107
                        _ = v18
                        _ = v19
                    }
                default:
                    var v20 uint = 226
                    _ = v20
                }
            case 2147483647:
                if len(v7) > 0 {
                    v7[0] = v3
                }
            default:
                for fi21 := 0; fi21 < 7; fi21++ {
                    var v22 float64 = (float64(float32((-2.25 * 0.1))) * 1e10)
                    switch "a" {
                    case "foo bar":
                        v22 -= (v22 + 0.0)
                    default:
                        fmt.Println(fi21)
                    }
                    _ = fi21
                    _ = v22
                }
            }
            if (true && false) {
                for fi23 := 0; fi23 < 5; fi23++ {
                    var v24 int = v3
                    v25 := GenS1{F0: -32768, F1: 2, F2: 18446744073709551615}
                    fmt.Println(v3, v3)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = len(v7)
                for fi26 := 0; fi26 < 1; fi26++ {
                    fmt.Println(v4, v1)
                    v1 = (v1 / ((2) | 1))
                    var v27 uint8 = 128
                    _ = fi26
                    _ = v27
                }
            }
        }
    case 127:
        switch (v3 / (((v3 & 9223372036854775807)) | 1)) {
        case -2147483648:
            var v28 string = ("foo bar" + v8)
            _ = v28
        default:
            var v29 int16 = 391
            _ = v29
        }
    }
    _ = len(v2)
    var v30 uint64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
