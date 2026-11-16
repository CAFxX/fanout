package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int32; F2 int8 }
type GenS3 struct { F0 string; F1 uint8; F2 string }

func main() {
    var v1 string = "hello"
    var v2 uint32 = 1
    var v3 uint16 = 2
    var v4 uint64 = uint64(v2)
    var v5 int = 128
    switch 371 {
    case 472:
        v5 += v5
    case 22:
        var v6 int64 = 42
        _ = v6
    case 32767:
        v7 := []int{(-(v5)), 128, v5, v5}
        _ = v7
    default:
        var v8 uint = 59
        _ = v8
    }
    v2 += (v2 % (((v2 / (((v2 & v2)) | 1))) | 1))
    v5 -= v5
    v4 = 863
    v2 -= (v2 ^ (v2 - 1))
    fmt.Println(v5, v1)
    v1 = ("a" + v1)
    if (0 <= 504) {
        if true {
            for fi9 := 0; fi9 < 10; fi9++ {
                var v10 float64 = (0.0 / (float64(-2.25) * 1e10))
                v3 += (v3 | v3)
                for fi11 := 0; fi11 < 13; fi11++ {
                    var v12 int64 = -9223372036854775808
                    v13 := []int{2147483647, (fi11 - (v5 << ((20) & 63)))}
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                _ = fi9
                _ = v10
            }
            fmt.Println(v2, v2)
            v5 = -1
        }
        v14 := GenS1{F0: v1}
        if false {
            v15 := GenS2{F0: 1, F1: 731, F2: 0}
            _ = v15
        }
        _ = v14
    }
    v16 := GenS2{F0: -418, F1: -1, F2: -128}
    v4 += (v4 | 1)
    v17 := GenS2{F0: 42, F1: -977, F2: 42}
    var v18 uint32 = (v2 * (v2 | v2))
    v19 := []int{v5, v5, (v5 * 32767), (^(v5))}
    var v20 string = v1
    switch (v5 | v5) {
    case 32768:
        if (317 <= 467) {
            if (true || ((v4 - (v4 - 18446744073709551615)) != (v4 & v4))) {
                if (true || false) {
                    switch v5 {
                    case 738:
                        if false {
                            v21 := GenS3{F0: "hello", F1: uint8(v4), F2: "hello"}
                            v22 := GenS2{F0: 505, F1: 1, F2: 1}
                            _ = v21
                            _ = v22
                        }
                    }
                    v4 = (v4 >> ((255) & 63))
                }
                v1 = v1
                fmt.Println(v20, v20)
            } else {
                for fi23 := 0; fi23 < 16; fi23++ {
                    switch 162 {
                    case 0:
                        v3 += 65535
                    default:
                        for fi24 := 0; fi24 < 3; fi24++ {
                            var v25 float64 = 1e10
                            var v26 uint8 = 128
                            _ = fi24
                            _ = v25
                            _ = v26
                        }
                    }
                    _ = fi23
                }
                v1 = (("\n\t" + v1) + v20)
                switch "" {
                case "":
                    var v27 int64 = -15
                    _ = v27
                case "foo bar":
                    v1 = "\n\t"
                }
            }
        } else {
            var v28 uint = 4294967295
            v19 = append(v19, (-(v5)))
            _ = v28
        }
    case -69:
        var v29 int8 = -56
        _ = v29
    case 127:
        v30 := GenS2{F0: 42, F1: int32(v18), F2: 0}
        _ = v30
    }
    v4 -= 2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
