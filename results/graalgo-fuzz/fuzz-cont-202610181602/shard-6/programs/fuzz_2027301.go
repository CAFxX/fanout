package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 uint8 }
type GenS2 struct { F0 uint; F1 int8; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int8; F1 string; F2 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = "\n\t"
    var v3 int64 = 42
    var v4 uint32 = 2147483648
    switch "" {
    case "hello":
        fmt.Println(v2)
    default:
        var v5 float64 = (0.0 + 0.0)
        _ = v5
    }
    fmt.Println(v3, v4, v4)
    if true {
        v6 := GenS1{F0: (((true || false) || (996 == 128)) || ((!true) && (-32768 > -1))), F1: true, F2: 255}
        v7 := GenS2{F0: 18446744073709551615, F1: 42, F2: (1 <= 1)}
        _ = v6
        _ = v7
    }
    v8 := GenS3{F0: 4294967295}
    v9 := GenS4{F0: 42}
    v3 = (v3 + v3)
    for fi10 := 0; fi10 < 9; fi10++ {
        v4 = 2147483647
        _ = fi10
    }
    v3 += (v3 % ((42) | 1))
    switch -270 {
    case -2147483648:
        if (0.1 >= ((3.14159 / float64(1.5)) * 0.1)) {
            v4 = (v4 & v4)
            v3 -= v3
            v3 -= v3
        } else {
            var v11 int32 = 2147483647
            _ = v11
        }
    default:
        fmt.Println(v3, v4)
    }
    v3 += v3
    if false {
        var v12 int64 = (v3 & v3)
        fmt.Println(v2, v3)
        _ = v12
    } else {
        if (!(127 < -88)) {
            v4 += 921
            var v13 string = v2
            v14 := GenS5{F0: -85, F1: (v2 + (v2 + v13)), F2: (v4 + v4)}
            _ = v13
            _ = v14
        }
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 int = fi15
            fmt.Println(v16, v3, v3)
            _ = fi15
            _ = v16
        }
    }
    v17 := GenS3{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
