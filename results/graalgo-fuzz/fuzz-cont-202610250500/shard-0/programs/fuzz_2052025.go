package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 int }

func main() {
    var v1 uint16 = 42
    v2 := make(map[int]uint32)
    var v3 int = 32767
    var v4 uint64 = 4294967295
    var v5 bool = false
    v6 := GenS1{F0: -1, F1: 216, F2: -9223372036854775808}
    if v5 {
        var v7 int8 = -9
        _ = v7
    }
    delete(v2, v3)
    v2[128] = uint32(v4)
    v4 -= v4
    if v5 {
        v5 = false
        v8 := GenS2{F0: 583, F1: -1, F2: -2147483648}
        switch 716 {
        case -10:
            v9 := v2[int(v1)]
            _ = v9
        case -1:
            var v10 int8 = -128
            _ = v10
        }
        _ = v8
    } else {
        var v11 float32 = float32(((0.0 * 3.14159) - 1e10))
        if true {
            fmt.Println(v4)
        }
        switch (v3 % (((v3 / (((v3 % ((-281) | 1))) | 1))) | 1)) {
        case 32768:
            switch "a" {
            case "hello":
                var v12 float32 = float32(((1e10 + (float64(1e10) + ((1.5 * 0.0) / 1.5))) + (-2.25 * -2.25)))
                _ = v12
            }
        case -9223372036854775808:
            var v13 uint16 = 808
            _ = v13
        case 9223372036854775807:
            var v14 uint32 = 4294967295
            _ = v14
        }
        _ = v11
    }
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
