package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 uint64; F1 int8 }

func main() {
    var v1 float32 = (1e10 / 1.0)
    var v2 float64 = 1.5
    var v3 int16 = 311
    var v4 int64 = 1
    v5 := GenS1{F0: "a", F1: 720}
    var v6 uint = 65535
    switch int(v3) {
    case 4294967295:
        v2 -= (v2 - v2)
    default:
        for fi7 := 0; fi7 < 1; fi7++ {
            if (100 < -1) {
                if (42 > 127) {
                    var v8 int = fi7
                    _ = v8
                }
            }
            var v9 uint16 = 14
            _ = fi7
            _ = v9
        }
    }
    v10 := []int{170, 42, 128, 213}
    v2 = v2
    v1 += float32(float64(v1))
    v11 := GenS1{F0: "hello", F1: -1}
    var v12 uint8 = 0
    v13 := GenS1{F0: "foo bar", F1: -860}
    if ((v2 - float64((v1 / float32((v2 * ((v2 * v2) / ((v2 - v2) / float64(1.5)))))))) > ((v2 / v2) - v2)) {
        v1 = float32((v2 - (1e10 + v2)))
        var v14 int = 128
        _ = v14
    } else {
        if (100 <= int8(v4)) {
            v3 -= v3
            switch "a" {
            case "\n\t":
                v4 = -9223372036854775808
            case "a":
                if false {
                    v15 := GenS2{F0: 0, F1: 91}
                    _ = v15
                } else {
                    var v16 int = 4294967295
                    _ = v16
                }
            case "foo bar":
                if ((^(v4)) < (v4 << ((v6) & 63))) {
                    v12 -= (v12 % (((v12 % (((v12 ^ 1)) | 1))) | 1))
                    if false {
                        var v17 int32 = -2147483648
                        if false {
                            v12 = v12
                            v10 = append(v10, -2147483648)
                        }
                        for fi18 := 0; fi18 < 8; fi18++ {
                            var v19 int16 = (v3 / ((-417) | 1))
                            v2 = 1e10
                            _ = fi18
                            _ = v19
                        }
                        _ = v17
                    }
                    for fi20 := 0; fi20 < 2; fi20++ {
                        v10 = append(v10, (^(fi20)))
                        v6 -= 236
                        if true {
                            fmt.Println(v3, fi20)
                        }
                        _ = fi20
                    }
                } else {
                    v10 = append(v10, int(v4))
                    v21 := GenS1{F0: (("\n\t" + "a") + ("hello" + "foo bar")), F1: 42}
                    fmt.Println(v6)
                    _ = v21
                }
            }
        } else {
            v22 := GenS1{F0: "foo bar", F1: 1}
            for fi23 := 0; fi23 < 16; fi23++ {
                fmt.Println(v4)
                fmt.Println(v1, v4, v12)
                _ = fi23
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                var v25 uint16 = 65535
                var v26 int16 = v3
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = v22
        }
    }
    if ((65535 >= 708) && false) {
        var v27 uint16 = 42
        if true {
            v3 = 42
        }
        _ = v27
    } else {
        v3 += v3
        var v28 float32 = 1e10
        _ = v28
    }
    fmt.Println(v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
