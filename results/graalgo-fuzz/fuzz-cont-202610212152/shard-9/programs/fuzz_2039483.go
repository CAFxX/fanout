package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint8 = 127
    var v2 int = 828
    var v3 uint8 = (v1 >> ((4294967295) & 63))
    v4 := []int{203, (v2 * (^(v2)))}
    var v5 int64 = 42
    var v6 int16 = 849
    for fi7 := 0; fi7 < 10; fi7++ {
        _ = len(v4)
        switch v2 {
        case 128:
            if (-1 < 100) {
                var v8 int = (v2 >> ((65535) & 63))
                for fi9 := 0; fi9 < 14; fi9++ {
                    v10 := GenS1{F0: true}
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
        case 42:
            if len(v4) > 0 {
                v4[0] = (fi7 * (v2 ^ (v2 / (((v2 & fi7)) | 1))))
            }
        default:
            fmt.Println(v2, v3, v6)
        }
        _ = fi7
    }
    switch "a" {
    case "hello":
        var v11 uint64 = 2
        _ = v11
    default:
        if (false || true) {
            var v12 int = v2
            v1 = (v3 & v1)
            _ = v12
        } else {
            switch v2 {
            case 128:
                v13 := GenS1{F0: true}
                _ = v13
            default:
                v1 = v3
            }
            v14 := GenS2{F0: "\n\t", F1: 42}
            v6 += 1000
            _ = v14
        }
    }
    var v15 int32 = 1
    if (v5 >= (-(v5))) {
        var v16 int16 = v6
        _ = v16
    } else {
        for fi17 := 0; fi17 < 14; fi17++ {
            var v18 float64 = 0.1
            v19 := GenS3{F0: 65}
            v2 -= fi17
            _ = fi17
            _ = v18
            _ = v19
        }
        v6 = 0
        v2 = v2
    }
    var v20 int8 = 1
    var v21 float64 = (1.5 - (float64(0.1) / 0.1))
    v5 = 0
    fmt.Println(v20)
    v22 := GenS3{F0: (v20 & -25)}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
