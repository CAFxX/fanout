package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 string; F1 int32; F2 int8 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 224
    var v3 int = 281
    v4 := []int{(^(v3))}
    var v5 int8 = 100
    var v6 int16 = (v2 ^ 42)
    if true {
        v7 := v1[477]
        var v8 uint = 255
        _ = v7
        _ = v8
    } else {
        if (v3 == v3) {
            switch "\n\t" {
            case "foo bar":
                var v9 float64 = 1.5
                _ = v9
            case "a":
                _ = len(v4)
            }
        } else {
            if true {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v2, v6, v6)
                default:
                    var v10 string = "foo bar"
                    _ = v10
                }
            }
        }
        if false {
            v4 = append(v4, 32768)
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v2 = v6
        } else {
            v11 := GenS1{F0: v6, F1: 65535}
            _ = v11
        }
        switch v3 {
        case 127:
            v12 := GenS2{F0: "a", F1: -309, F2: (^(v5))}
            _ = v12
        case -801:
            v6 = v2
        case 0:
            fmt.Println(v5, v5)
        default:
            switch "" {
            case "hello":
                v13 := GenS2{F0: ("foo bar" + "hello"), F1: 850, F2: v5}
                _ = v13
            case "a":
                fmt.Println(v3)
            default:
                switch "\n\t" {
                case "\n\t":
                    v5 = v5
                case "hello":
                    fmt.Println(v6, v5, v2)
                case "a":
                    v4 = append(v4, v3)
                }
            }
        }
    }
    var v14 string = (("hello" + "a") + ("a" + "\n\t"))
    fmt.Println(v5, v14)
    v15 := GenS2{F0: v14, F1: 44, F2: (v5 << ((255) & 63))}
    v6 -= v2
    var v16 float64 = (1.5 / float64(1e10))
    v2 = v6
    v5 -= (^(v5))
    var v17 uint8 = 0
    if false {
        v16 += float64(float32((float64(-2.25) - 1.5)))
        v6 = (v2 % (((v6 | v2)) | 1))
    } else {
        for fi18 := 0; fi18 < 9; fi18++ {
            _ = len(v4)
            _ = fi18
        }
        var v19 int32 = 42
        _ = v19
    }
    v20 := GenS3{F0: v6}
    if (2 != 124) {
        v21 := GenS1{F0: v2, F1: 65535}
        var v22 int64 = 0
        if ((v17 >> ((1) & 63)) < v17) {
            v17 -= (v17 | (v17 / ((63) | 1)))
            v6 = (v6 / ((1) | 1))
        }
        _ = v21
        _ = v22
    } else {
        switch v14 {
        case "a":
            var v23 int32 = 393
            _ = v23
        case "hello":
            if ((v5 + (v5 / (((v5 - (v5 ^ v5))) | 1))) <= v5) {
                var v24 bool = (!(0.0 > 3.14159))
                var v25 float32 = ((1e10 - (float32(((3.14159 - 1.5) * 1.5)) / (float32(float64(1.5)) - 1e10))) / 1.0)
                _ = v24
                _ = v25
            }
        default:
            for fi26 := 0; fi26 < 14; fi26++ {
                v27 := GenS2{F0: v14, F1: 2147483647, F2: v5}
                _ = fi26
                _ = v27
            }
        }
    }
    if false {
        if (65535 == 0) {
            fmt.Println(v2, v2)
            v16 = v16
        }
    }
    var v28 int16 = (v2 | v6)
    for fi29 := 0; fi29 < 4; fi29++ {
        fmt.Println(v16, v5, v16)
        v5 += (v5 * (v5 << ((997) & 63)))
        _ = fi29
    }
    v28 -= (v2 + (v2 << ((1) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
