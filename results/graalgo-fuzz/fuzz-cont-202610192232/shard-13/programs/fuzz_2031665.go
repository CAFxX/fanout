package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint32; F1 int16; F2 int32 }

func main() {
    var v1 uint16 = 2
    v2 := make(map[uint32]uint32)
    v3 := []int{32767, 9223372036854775807, -1}
    var v4 int8 = 42
    var v5 uint64 = uint64(v1)
    switch "foo bar" {
    case "":
        var v6 float32 = 0.1
        _ = v6
    }
    delete(v2, 715)
    v7 := GenS1{F0: -32768}
    v5 += (v5 ^ v5)
    switch 664 {
    case 185:
        for fi8 := 0; fi8 < 17; fi8++ {
            switch fi8 {
            case 1:
                var v9 int32 = 2147483647
                _ = v9
            case 9223372036854775807:
                var v10 int = (fi8 | fi8)
                _ = v10
            case -60:
                for fi11 := 0; fi11 < 10; fi11++ {
                    fmt.Println(v1)
                    _ = fi11
                }
            default:
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            }
            if (false && (-682 < -1)) {
                v1 = (v1 % ((v1) | 1))
                var v12 uint32 = 2
                v5 = (^(v5))
                _ = v12
            }
            v13 := GenS1{F0: 1000}
            _ = fi8
            _ = v13
        }
    case 32767:
        v14 := GenS2{F0: ((v1 * (v1 >> ((4294967295) & 63))) > (v1 ^ (v1 - 65535))), F1: -9223372036854775808}
        _ = v14
    default:
        switch 9223372036854775807 {
        case 0:
            var v15 bool = true
            _ = v15
        case 59:
            v5 = v5
        }
    }
    var v16 float64 = (((float64((-2.25 + (3.14159 * float32(0.0)))) + ((3.14159 * -2.25) / 1.5)) * (((1e10 * ((0.1 - (1.5 + -2.25)) * 0.0)) / 3.14159) - 0.0)) * (0.0 - 0.0))
    v17 := GenS3{F0: 539, F1: -32768, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
