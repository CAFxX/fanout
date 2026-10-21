package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 int64; F1 uint8; F2 int8 }

func main() {
    var v1 string = ("a" + "hello")
    v2 := []int{-217}
    var v3 uint32 = 65535
    v4 := []string{}
    var v5 uint = 18446744073709551615
    var v6 int32 = -2147483648
    switch 9223372036854775807 {
    case 128:
        switch 9223372036854775807 {
        case 728:
            fmt.Println(v1, v5)
        case 9223372036854775807:
            for fi7 := 0; fi7 < 19; fi7++ {
                v8 := GenS1{F0: 42, F1: 2, F2: (v3 << (((v5 % ((305) | 1))) & 63))}
                var v9 int64 = -871
                _ = fi7
                _ = v8
                _ = v9
            }
        }
    case 0:
        v3 = v3
    default:
        v10 := GenS1{F0: 4294967295, F1: 255, F2: v3}
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        for fi12 := 0; fi12 < 1; fi12++ {
            fmt.Println(fi12, fi12)
            _ = fi12
        }
        v1 = v1
        v13 := GenS1{F0: 0, F1: 250, F2: 2147483648}
        _ = fi11
        _ = v13
    }
    var v14 float32 = 0.1
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := GenS2{F0: v5, F1: v3, F2: (0.0 == 1e10)}
            var v18 uint64 = 77
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = len(v2)
        var v19 string = "hello"
        _ = fi15
        _ = v19
    }
    v2 = append(v2, 4294967295)
    v5 -= 1
    v2 = append(v2, 4294967295)
    var v20 uint = (v5 % ((v5) | 1))
    for fi21 := 0; fi21 < 4; fi21++ {
        fmt.Println(v3)
        v22 := GenS3{F0: int64(v3), F1: 127, F2: -8}
        _ = fi21
        _ = v22
    }
    v14 = -2.25
    fmt.Println(v20, v14, v14)
    var v23 uint = (v5 % ((uint(v3)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
