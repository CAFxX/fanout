package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int; F1 int16 }

func main() {
    var v1 int8 = 42
    var v2 int8 = (^(v1))
    var v3 float32 = ((-2.25 + float32(((float64(1.5) + 1.5) * (1.5 / 1.0)))) - (1e10 * 0.0))
    var v4 uint32 = 1
    for fi5 := 0; fi5 < 6; fi5++ {
        switch fi5 {
        case 1:
            var v6 float32 = ((float32((3.14159 * (0.1 * 0.0))) / 0.1) / (float32(-2.25) - float32(1e10)))
            _ = v6
        case 655:
            for fi7 := 0; fi7 < 8; fi7++ {
                v2 -= (v2 & v2)
                _ = fi7
            }
        }
        v8 := GenS1{F0: fi5, F1: 42}
        var v9 int8 = (v2 / (((v2 >> ((2) & 63))) | 1))
        _ = fi5
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if false {
            v11 := []int{89, (fi10 % (((fi10 % ((fi10) | 1))) | 1)), (^(fi10))}
            v2 = (v1 << ((322) & 63))
            _ = v11
        } else {
            v4 = (v4 - v4)
            if true {
                v3 -= 1e10
                if false {
                    v12 := GenS2{F0: 42, F1: -32768}
                    _ = v12
                }
                v3 -= 1e10
            } else {
                v13 := []int{(^(fi10))}
                if (1 >= 1) {
                    v2 -= (v2 * (-(v2)))
                }
                v3 -= 1e10
                _ = v13
            }
        }
        _ = fi10
    }
    var v14 int = 9223372036854775807
    v15 := []int{(-(v14)), v14, v14}
    var v16 uint = 1
    v1 += (^(v2))
    switch "" {
    case "":
        v14 = (^(v14))
    case "foo bar":
        fmt.Println(v16, v3)
    case "a":
        switch (v14 * int(v1)) {
        case 32767:
            v17 := GenS3{F0: v14, F1: -32768}
            _ = v17
        case 129:
            if (482 <= v14) {
                v16 = 65535
                v16 -= v16
            }
        case 4294967295:
            var v18 uint32 = 172
            _ = v18
        default:
            v16 = v16
        }
    }
    _ = len(v15)
    var v19 uint32 = 65535
    v19 += (v19 * 922)
    var v20 uint8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
