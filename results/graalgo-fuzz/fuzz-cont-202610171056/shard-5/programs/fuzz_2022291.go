package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint; F2 uint16 }
type GenS3 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint64, p1 uint) int64 {
    var v6 uint64 = p0
    p0 += 1
    var v7 string = "hello"
    if ((!false) && (-9223372036854775808 < -1)) {
        v8 := []int{-2147483648, -854, 128, 127}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32) int16 {
    p0 = p0
    var v9 uint = 65535
    _ = p0
    _ = v9
    return 0
}

func main() {
    var v1 int = 42
    v2 := make(map[int]string)
    var v3 int32 = 730
    var v4 int8 = 127
    var v5 int16 = 1
    v10 := GenS1{F0: 349, F1: 4294967295}
    v11 := GenS1{F0: 0, F1: 42}
    _ = len(v2)
    switch "a" {
    case "":
        for fi12 := 0; fi12 < 18; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                v5 -= v5
                v14 := []int{9223372036854775807, v1, -461}
                _ = fi13
                _ = v14
            }
            v3 += v3
            var v15 int8 = v4
            _ = fi12
            _ = v15
        }
    case "hello":
        v3 = (v3 / (((v3 >> ((4294967295) & 63))) | 1))
    default:
        for fi16 := 0; fi16 < 15; fi16++ {
            fmt.Println(v4)
            v17 := GenS1{F0: 2, F1: 186}
            _ = fi16
            _ = v17
        }
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        v19 := GenS2{F0: 42, F1: 4294967295, F2: 255}
        _ = fi18
        _ = v19
    }
    v3 = 1
    v4 -= (v4 * int8(v5))
    v20 := GenS3{F0: 127, F1: 2}
    _ = len(v2)
    var v21 int64 = -1
    var v22 uint64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
