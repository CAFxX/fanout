package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint8 }

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = true
    var v3 float32 = 0.0
    var v4 uint64 = 1
    v5 := make(map[uint32]int)
    v6 := []int{4294967295, 0, 32768}
    v2 = v2
    fmt.Println(v4)
    var v7 int64 = -9223372036854775808
    switch ("" + "\n\t") {
    case "":
        v2 = (((-80 < -128) && (72 == 255)) && true)
    }
    v5[512] = 0
    if (("a" != "") || false) {
        v4 += (v4 & (v4 ^ (v4 % ((914) | 1))))
    } else {
        delete(v5, 2)
        v2 = (0.0 >= (1.5 - 0.1))
    }
    v1[1] = -9223372036854775808
    for fi8 := 0; fi8 < 5; fi8++ {
        v7 += (v7 | v7)
        switch "foo bar" {
        case "a":
            v9 := GenS1{F0: 824, F1: 944}
            _ = v9
        case "hello":
            var v10 int32 = -2147483648
            _ = v10
        }
        _ = fi8
    }
    v7 -= -1
    v2 = v2
    var v11 int = 9223372036854775807
    _ = len(v1)
    switch "" {
    case "foo bar":
        v12 := GenS1{F0: 663, F1: 0}
        _ = v12
    case "a":
        delete(v1, 1)
    case "hello":
        for fi13 := 0; fi13 < 2; fi13++ {
            if (false && (false && v2)) {
                fmt.Println(v4, v11, v11)
            }
            v14 := GenS2{F0: 0}
            switch 908 {
            case 438:
                v3 = 1e10
            case 32767:
                v15 := GenS3{F0: 1}
                _ = v15
            case -26:
                var v16 uint = 4294967295
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    }
    if false {
        v17 := GenS4{F0: 0, F1: 1000, F2: 128}
        _ = v17
    } else {
        var v18 float32 = (v3 * float32(float64(float32(3.14159))))
        _ = v18
    }
    v2 = (65535 >= 65535)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
