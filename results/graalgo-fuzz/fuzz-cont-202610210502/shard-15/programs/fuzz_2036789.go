package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 int64 }
type GenS2 struct { F0 uint32; F1 string; F2 uint32 }
type GenS3 struct { F0 int32; F1 int64; F2 uint8 }
type GenS4 struct { F0 uint16; F1 uint }

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{4294967295, 65535}
    v3 := make(map[int]uint32)
    v4 := []int{-1, -38, -736, 4294967295}
    v5 := make(map[uint32]uint32)
    var v6 uint32 = 65535
    v6 -= v6
    v7 := v3[42]
    switch ("foo bar" + ("hello" + "a")) {
    case "hello":
        for fi8 := 0; fi8 < 5; fi8++ {
            if false {
                var v9 int8 = 100
                _ = v9
            }
            v6 -= v6
            _ = fi8
        }
    case "a":
        v6 = v6
    case "":
        var v10 uint16 = 58
        _ = v10
    default:
        if len(v2) > 0 {
            v2[0] = (v7 / (((v6 & 732)) | 1))
        }
    }
    fmt.Println(v7, v6)
    fmt.Println(v7, v7)
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(v7, v6, fi11)
            if (!(-117 >= 100)) {
                fmt.Println(fi12)
                var v13 int8 = 11
                switch "" {
                case "foo bar":
                    fmt.Println(v13, v6, v7)
                case "":
                    v14 := GenS1{F0: (v7 % (((v7 ^ (v7 - 680))) | 1)), F1: 128, F2: 1}
                    _ = v14
                case "hello":
                    var v15 uint32 = 497
                    _ = v15
                }
                if true {
                    v16 := GenS2{F0: (v6 % ((302) | 1)), F1: "hello", F2: 2147483647}
                    _ = v16
                }
                if (42 <= -2147483648) {
                    if len(v2) > 0 {
                        v2[0] = (v7 - (v6 / ((v6) | 1)))
                    }
                } else {
                    fmt.Println(v7, v6, fi11)
                    v17 := GenS3{F0: -2147483648, F1: 0, F2: 42}
                    fmt.Println(fi12)
                    _ = v17
                }
                _ = v13
            } else {
                v6 += 4294967295
                v6 = (v7 - v6)
            }
            v7 += v7
            if (111 < int16(v6)) {
                switch fi11 {
                case 4294967295:
                    var v18 int32 = -1
                    _ = v18
                case 128:
                    switch fi12 {
                    case 2147483647:
                        v19 := GenS3{F0: 42, F1: 1, F2: 1}
                        _ = v19
                    case 4294967295:
                        v20 := GenS4{F0: 255, F1: 4294967295}
                        _ = v20
                    }
                default:
                    if false {
                        var v21 uint = 0
                        v7 -= v6
                        _ = v21
                    }
                }
                var v22 uint8 = 1
                _ = v22
            }
            _ = fi12
        }
        v1[fi11] = (("a" + "\n\t") + "a")
        v7 = v6
        _ = fi11
    }
    v6 += (v7 - (v6 << ((177) & 63)))
    if (false || false) {
        for fi23 := 0; fi23 < 3; fi23++ {
            fmt.Println(v6, fi23)
            var v24 uint64 = 0
            _ = fi23
            _ = v24
        }
        var v25 string = "\n\t"
        if true {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 int16 = 42
                v28 := GenS1{F0: v7, F1: fi26, F2: 9223372036854775807}
                fmt.Println(v27, v27)
                _ = fi26
                _ = v27
                _ = v28
            }
            if (2 == 0) {
                v7 += v6
            }
            v7 -= (v6 ^ (v6 << ((783) & 63)))
        }
        _ = v25
    }
    v7 += 208
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
