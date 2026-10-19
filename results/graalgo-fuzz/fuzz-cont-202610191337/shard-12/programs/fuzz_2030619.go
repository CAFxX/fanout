package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int64; F2 string }
func genFn0() int16 {
    var v7 bool = ((((0.0 >= 0.0) && (466 <= 2147483647)) && (42 <= 1)) && (1e10 < ((float32(3.14159) * 1.5) + 0.1)))
    fmt.Println(v7)
    var v8 int32 = -743
    v7 = true
    var v9 float64 = 0.0
    _ = v7
    _ = v8
    _ = v9
    return int16(v8)
}

func main() {
    var v1 int32 = -972
    var v2 int16 = -1
    var v3 string = "a"
    v4 := make(map[int]bool)
    var v5 uint64 = 42
    v6 := []string{"foo bar", "a", ("foo bar" + v3), "a"}
    if false {
        fmt.Println(v1)
        v2 = v2
    }
    if false {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = len(v4)
    }
    var v10 float32 = float32(float64(float32((float64(float32((3.14159 / 0.1))) - ((1e10 + ((1.5 * -2.25) + 3.14159)) + 0.1)))))
    var v11 string = ""
    v12 := GenS1{F0: 199}
    var v13 int64 = 9223372036854775807
    v4[924] = false
    var v14 uint32 = 65535
    if (v2 != v2) {
        v14 = v14
        if len(v6) > 0 {
            v6[0] = v3
        }
        if (uint(v14) < 0) {
            for fi15 := 0; fi15 < 10; fi15++ {
                v1 = (v1 + 77)
                _ = len(v4)
                _ = fi15
            }
            var v16 uint8 = 128
            _ = v16
        } else {
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int32 = v1
                switch (fi17 - fi17) {
                case 875:
                    var v19 uint64 = v5
                    _ = v19
                default:
                    for fi20 := 0; fi20 < 2; fi20++ {
                        fmt.Println(v2, v10, v11)
                        _ = fi20
                    }
                }
                if true {
                    v21 := GenS2{F0: 267}
                    var v22 int64 = (v13 << ((0) & 63))
                    _ = v21
                    _ = v22
                }
                _ = fi17
                _ = v18
            }
            switch v11 {
            case "a":
                v5 -= v5
            default:
                v23 := v4[4294967295]
                _ = v23
            }
            v24 := GenS2{F0: 931}
            _ = v24
        }
    } else {
        switch 32768 {
        case -764:
            var v25 int16 = (v2 - -32768)
            _ = v25
        case 257:
            v26 := GenS3{F0: 0, F1: v13, F2: ("" + v3)}
            _ = v26
        case 853:
            v11 = ""
        }
    }
    var v27 float32 = (0.1 / (v10 - 1e10))
    var v28 bool = (!(v10 >= float32(-2.25)))
    for fi29 := 0; fi29 < 10; fi29++ {
        if v28 {
            var v30 uint32 = v14
            var v31 uint32 = (v30 % (((v30 >> ((930) & 63))) | 1))
            switch v11 {
            case "a":
                fmt.Println(v5)
            case "\n\t":
                v14 = v31
            case "":
                v30 -= (v30 & (v30 - (v31 >> ((65535) & 63))))
            default:
                v32 := GenS2{F0: 279}
                _ = v32
            }
            _ = v30
            _ = v31
        } else {
            var v33 float64 = float64(float32((-2.25 - (float64(-2.25) + (float64(0.0) - 1.5)))))
            v34 := v4[(fi29 * (fi29 | int(v14)))]
            _ = v33
            _ = v34
        }
        _ = fi29
    }
    switch -1 {
    case 32767:
        var v35 uint32 = 2147483647
        _ = v35
    case -661:
        var v36 float64 = float64(0.0)
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
