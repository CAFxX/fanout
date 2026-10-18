package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 int32; F1 uint16 }

func main() {
    var v1 float32 = (float32(1.5) / ((1.5 + (1.5 + 0.0)) * float32((1e10 / -2.25))))
    var v2 float32 = v1
    var v3 uint16 = 1
    var v4 uint16 = 2
    var v5 int = -9223372036854775808
    fmt.Println(v1, v2)
    if (0 != 0) {
        v5 -= (v5 ^ 359)
    }
    v1 = 0.0
    v2 += ((float32(1e10) - float32(-2.25)) + v1)
    v2 = v2
    v4 += v3
    switch (v5 - (v5 % (((v5 >> ((915) & 63))) | 1))) {
    case -973:
        for fi6 := 0; fi6 < 11; fi6++ {
            var v7 uint64 = 2
            _ = fi6
            _ = v7
        }
    case 685:
        switch "\n\t" {
        case "":
            v2 = 0.1
        case "a":
            for fi8 := 0; fi8 < 20; fi8++ {
                switch v5 {
                case 760:
                    var v9 uint32 = 2147483648
                    _ = v9
                case 4294967295:
                    v10 := []int{}
                    _ = v10
                case 9223372036854775807:
                    v1 += v1
                default:
                    v5 = (fi8 + (v5 * fi8))
                }
                v1 -= v1
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 string = "\n\t"
                    if true {
                        switch (v5 / ((v5) | 1)) {
                        case 128:
                            v12 = ("foo bar" + ("\n\t" + "a"))
                        case -1:
                            var v13 int = (^(v5))
                            _ = v13
                        case 665:
                            v14 := GenS1{F0: -1, F1: 4294967295}
                            _ = v14
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                _ = fi8
            }
        case "foo bar":
            if (!(18446744073709551615 >= 0)) {
                for fi15 := 0; fi15 < 18; fi15++ {
                    v1 += float32(((1.5 + 1e10) - 1.5))
                    fmt.Println(v4)
                    for fi16 := 0; fi16 < 13; fi16++ {
                        switch ((("\n\t" + "\n\t") + ("a" + ("a" + "a"))) + "a") {
                        case "":
                            v17 := GenS2{F0: 42, F1: v3}
                            _ = v17
                        case "foo bar":
                            if true {
                                fmt.Println(v4, v3, v2)
                                fmt.Println(v5, v5, fi15)
                                v18 := GenS2{F0: -1, F1: 42}
                                _ = v18
                            }
                        default:
                            var v19 int8 = 100
                            _ = v19
                        }
                        fmt.Println(v5, v1, v1)
                        _ = fi16
                    }
                    _ = fi15
                }
                var v20 float32 = v2
                _ = v20
            } else {
                v21 := []int{(v5 * v5)}
                _ = v21
            }
        default:
            var v22 uint64 = 407
            _ = v22
        }
    default:
        for fi23 := 0; fi23 < 17; fi23++ {
            var v24 bool = (0.1 >= (0.1 + 0.1))
            _ = fi23
            _ = v24
        }
    }
    var v25 int16 = -1
    fmt.Println(v3)
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS2{F0: 42, F1: v3}
        _ = fi26
        _ = v27
    }
    var v28 int64 = 1
    v29 := GenS1{F0: -128, F1: 1}
    var v30 int32 = 2147483647
    v28 += v28
    var v31 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
