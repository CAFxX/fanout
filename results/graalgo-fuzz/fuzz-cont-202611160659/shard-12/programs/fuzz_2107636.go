package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int32 = -413
    var v2 uint64 = 4294967295
    v3 := []int{-9223372036854775808, 32768, 219, -429}
    var v4 uint = 2
    v5 := []string{("" + "\n\t"), "", ("hello" + "foo bar")}
    var v6 int32 = 1
    if len(v5) > 0 {
        v5[0] = "\n\t"
    }
    for fi7 := 0; fi7 < 20; fi7++ {
        switch "" {
        case "foo bar":
            switch "a" {
            case "":
                var v8 int64 = -9223372036854775808
                _ = v8
            case "a":
                v1 = (v1 / ((-336) | 1))
            }
        case "\n\t":
            for fi9 := 0; fi9 < 2; fi9++ {
                fmt.Println(fi7, v4, v6)
                var v10 uint8 = 49
                v6 = (v1 ^ (v1 * v6))
                _ = fi9
                _ = v10
            }
        case "a":
            v4 = 235
        default:
            fmt.Println(v6, v2, v2)
        }
        _ = fi7
    }
    var v11 bool = false
    for fi12 := 0; fi12 < 14; fi12++ {
        for fi13 := 0; fi13 < 10; fi13++ {
            switch (^(fi13)) {
            case 32768:
                var v14 uint16 = 2
                _ = v14
            case -369:
                fmt.Println(v6, v11, fi12)
            }
            _ = fi13
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 int32 = v6
            v17 := GenS1{F0: 153, F1: -1}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi12
    }
    var v18 int16 = int16(v4)
    v1 = (v1 * (v1 % ((620) | 1)))
    fmt.Println(v2)
    var v19 uint16 = 100
    v20 := GenS1{F0: 128, F1: 71}
    if (v11 && (((v6 ^ (v1 & 990)) <= (^(v1))) && v11)) {
        var v21 float32 = ((1.5 + ((float32(0.1) - (3.14159 - ((1e10 * 1.5) / 3.14159))) / 0.1)) * 0.1)
        fmt.Println(v2)
        _ = v21
    }
    v19 -= 0
    switch ("foo bar" + "foo bar") {
    case "a":
        if false {
            v2 += (v2 % (((v2 - v2)) | 1))
            if (2 == 255) {
                if v11 {
                    if true {
                        switch -920 {
                        case 32768:
                            switch -1 {
                            case 896:
                                _ = len(v5)
                            default:
                                if len(v5) > 0 {
                                    v5[0] = ("foo bar" + "\n\t")
                                }
                            }
                        default:
                            v6 += (^(v6))
                        }
                    } else {
                        if (float32(3.14159) > 1.5) {
                            v22 := GenS2{F0: 1}
                            v6 = (-(v1))
                            _ = v22
                        }
                        fmt.Println(v19, v2)
                        var v23 string = "\n\t"
                        _ = v23
                    }
                    v18 += (v18 ^ v18)
                    var v24 int64 = 1
                    _ = v24
                }
                v4 = v4
            }
        }
    case "hello":
        v1 += v1
    case "\n\t":
        v25 := GenS2{F0: uint8(v4)}
        _ = v25
    }
    v3 = append(v3, -61)
    var v26 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
