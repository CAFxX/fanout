package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 int8 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]bool)
    var v3 float32 = (1e10 + 1e10)
    var v4 uint8 = 1
    var v5 int64 = -9223372036854775808
    v4 += (v4 / (((v4 + v4)) | 1))
    v6 := []int{4294967295, 9223372036854775807, -2147483648}
    var v7 uint16 = 277
    switch 133 {
    case 9223372036854775807:
        v5 -= v5
    }
    var v8 int = 4294967295
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v9 uint = uint(v7)
    v5 = (v5 * v5)
    v10 := GenS1{F0: 0, F1: false}
    var v11 int16 = 32767
    v8 += v8
    switch "a" {
    case "foo bar":
        for fi12 := 0; fi12 < 5; fi12++ {
            if (!(!(!true))) {
                var v13 string = "foo bar"
                _ = v13
            }
            v14 := GenS1{F0: 656, F1: true}
            v15 := GenS1{F0: 108, F1: false}
            _ = fi12
            _ = v14
            _ = v15
        }
    case "":
        for fi16 := 0; fi16 < 4; fi16++ {
            v5 += 9223372036854775807
            fmt.Println(v11)
            for fi17 := 0; fi17 < 5; fi17++ {
                v4 += v4
                _ = fi17
            }
            _ = fi16
        }
    case "hello":
        switch int(v4) {
        case -727:
            for fi18 := 0; fi18 < 6; fi18++ {
                if (float32(0.1) < v3) {
                    fmt.Println(v5, fi18)
                    _ = len(v1)
                } else {
                    var v19 uint32 = 0
                    v20 := GenS2{F0: int8(v8)}
                    _ = v19
                    _ = v20
                }
                _ = fi18
            }
        default:
            fmt.Println(v11, v7, v3)
        }
    default:
        fmt.Println(v3)
    }
    v5 += v5
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
