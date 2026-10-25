package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }

func main() {
    var v1 int64 = 1
    v2 := make(map[string]bool)
    var v3 int32 = int32(v1)
    v4 := make(map[string]uint32)
    v5 := make(map[int]string)
    v6 := []uint32{}
    var v7 int8 = int8(v1)
    var v8 int16 = 32767
    var v9 uint16 = 821
    v7 -= (v7 * -85)
    if len(v6) > 0 {
        v6[0] = 2147483647
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint16 = (v9 | (v9 & v9))
        switch (fi10 * -575) {
        case 128:
            var v12 uint64 = 18446744073709551615
            _ = v12
        case 42:
            fmt.Println(v7, v11)
        case -89:
            delete(v2, ("hello" + "foo bar"))
        default:
            var v13 float64 = 1e10
            _ = v13
        }
        _ = fi10
        _ = v11
    }
    var v14 float32 = (((3.14159 * -2.25) + float32(3.14159)) - 0.0)
    for fi15 := 0; fi15 < 5; fi15++ {
        var v16 int32 = (v3 * v3)
        if true {
            var v17 int64 = int64(v7)
            v1 += v1
            _ = v17
        }
        if false {
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                v18 := GenS1{F0: v16, F1: 2}
                _ = v18
            case "a":
                switch (fi15 * fi15) {
                case 42:
                    if false {
                        v19 := GenS1{F0: (v3 | v16), F1: 1}
                        v6 = append(v6, 42)
                        fmt.Println(v3, v1)
                        _ = v19
                    }
                case 2147483647:
                    v1 -= v1
                }
            default:
                for fi20 := 0; fi20 < 11; fi20++ {
                    fmt.Println(v9, fi20, v3)
                    var v21 uint8 = 128
                    v22 := GenS1{F0: (v16 + v3), F1: (^(v21))}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            }
        } else {
            var v23 string = "a"
            var v24 int64 = (v1 % (((v1 + int64(v8))) | 1))
            _ = v23
            _ = v24
        }
        _ = fi15
        _ = v16
    }
    v25 := v4["\n\t"]
    var v26 string = "\n\t"
    delete(v2, v26)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
