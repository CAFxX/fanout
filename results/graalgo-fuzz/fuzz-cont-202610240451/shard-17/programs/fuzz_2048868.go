package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    var v1 float32 = (0.1 + (3.14159 * 1e10))
    var v2 uint8 = 128
    var v3 int = -9223372036854775808
    v4 := make(map[uint32]bool)
    v5 := []int{v3, -2147483648, 4294967295}
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: 0, F1: 37, F2: 32767}
        switch "a" {
        case "foo bar":
            v2 += (v2 ^ (v2 << ((1) & 63)))
        case "":
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 uint64 = 0
                fmt.Println(fi6, v1)
                _ = fi8
                _ = v9
            }
        }
        var v10 int32 = 42
        _ = fi6
        _ = v10
        _ = v7
    }
    if false {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    v1 = v1
    for fi11 := 0; fi11 < 14; fi11++ {
        if (42 == 42) {
            var v12 int8 = 100
            _ = v12
        }
        v13 := GenS1{F0: 42, F1: v2, F2: 0}
        var v14 uint = 53
        _ = fi11
        _ = v13
        _ = v14
    }
    fmt.Println(v3)
    if true {
        for fi15 := 0; fi15 < 7; fi15++ {
            if (!true) {
                v1 = v1
                fmt.Println(v1, v3, fi15)
            }
            if (1 <= -1) {
                if (644 == 4294967295) {
                    var v16 bool = true
                    var v17 int32 = 1
                    _ = v16
                    _ = v17
                }
            } else {
                switch "\n\t" {
                case "hello":
                    var v18 int64 = 42
                    _ = v18
                case "foo bar":
                    for fi19 := 0; fi19 < 10; fi19++ {
                        if (true || true) {
                            switch "" {
                            case "hello":
                                v3 += (fi15 / (((^(fi15))) | 1))
                            }
                            v2 -= 200
                            v4[2] = (18446744073709551615 == 42)
                        } else {
                            v20 := GenS2{F0: 18446744073709551615, F1: 614}
                            var v21 float32 = ((v1 / (float32(1.5) * (0.0 + v1))) + v1)
                            v1 = (v21 / 1.5)
                            _ = v20
                            _ = v21
                        }
                        v22 := GenS3{F0: ("" + "hello"), F1: -1}
                        var v23 uint8 = (v2 % ((v2) | 1))
                        v1 -= v1
                        var v24 int32 = -229
                        _ = fi19
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                case "a":
                    v2 -= 20
                }
            }
            _ = fi15
        }
        var v25 float64 = -2.25
        _ = v25
    }
    var v26 float64 = 3.14159
    for fi27 := 0; fi27 < 16; fi27++ {
        v2 = (v2 & 82)
        _ = fi27
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
