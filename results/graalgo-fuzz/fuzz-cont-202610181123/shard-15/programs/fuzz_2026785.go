package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS2 struct { F0 bool; F1 int64; F2 int8 }
type GenS3 struct { F0 bool }

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []uint32{1, 659, 0, 1}
    var v3 int64 = 1
    var v4 int64 = (^(v3))
    var v5 float64 = (0.0 / (float64(-2.25) * 1e10))
    var v6 uint8 = 212
    var v7 uint64 = 642
    for fi8 := 0; fi8 < 11; fi8++ {
        if len(v2) > 0 {
            v2[0] = 2147483648
        }
        var v9 int32 = 2147483647
        v10 := GenS1{F0: 77, F1: v6, F2: 2147483648}
        _ = fi8
        _ = v10
        _ = v9
    }
    if true {
        if false {
            v6 += v6
        }
        v11 := GenS2{F0: false, F1: (v3 | (^(v1))), F2: -1}
        _ = v11
    } else {
        v12 := GenS1{F0: uint(v6), F1: v6, F2: 562}
        v13 := GenS3{F0: true}
        switch -9223372036854775808 {
        case 128:
            var v14 uint8 = (v6 >> ((32) & 63))
            _ = v14
        }
        _ = v12
        _ = v13
    }
    v4 += v4
    switch ("hello" + "a") {
    case "hello":
        for fi15 := 0; fi15 < 4; fi15++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi15
        }
    }
    fmt.Println(v4, v1)
    var v16 int64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
