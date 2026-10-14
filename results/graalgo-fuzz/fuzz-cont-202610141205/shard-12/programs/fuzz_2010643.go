package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 int; F1 uint32 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 int = 128
    v2 := make(map[int]int)
    var v3 uint32 = 259
    v4 := []uint32{v3, v3, 1}
    v1 -= 1
    switch ("\n\t" + "a") {
    case "a":
        v5 := GenS1{F0: 0, F1: v3, F2: 42}
        _ = v5
    }
    v3 -= 0
    fmt.Println(v3, v1)
    switch v1 {
    case -146:
        v6 := GenS2{F0: (-(v1)), F1: v3}
        _ = v6
    default:
        var v7 float32 = 1.5
        _ = v7
    }
    var v8 int16 = 1
    var v9 int64 = -9223372036854775808
    v1 = 2147483647
    _ = len(v4)
    var v10 uint64 = 4294967295
    var v11 int64 = 9223372036854775807
    if (2 < (v10 - v10)) {
        if false {
            v4 = append(v4, (v3 << ((18446744073709551615) & 63)))
            v10 += v10
        } else {
            fmt.Println(v3, v9, v8)
        }
        fmt.Println(v11, v9)
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            v4 = append(v4, uint32(v10))
            var v13 uint8 = 127
            _ = fi12
            _ = v13
        }
    }
    var v14 bool = true
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v4 = append(v4, (^(v3)))
        fmt.Println(v9, v8)
        v10 -= (v10 + v10)
    } else {
        v15 := GenS3{F0: -9223372036854775808}
        for fi16 := 0; fi16 < 10; fi16++ {
            v8 -= v8
            v17 := GenS1{F0: 991, F1: (v3 << ((65535) & 63)), F2: (v9 ^ 42)}
            v18 := GenS2{F0: fi16, F1: (^(v3))}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    if v14 {
        for fi19 := 0; fi19 < 14; fi19++ {
            var v20 int = (v1 - (v1 % ((fi19) | 1)))
            _ = fi19
            _ = v20
        }
        var v21 int32 = 738
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
