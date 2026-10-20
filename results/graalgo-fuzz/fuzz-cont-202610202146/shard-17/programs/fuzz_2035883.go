package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 uint16; F2 int }
type GenS3 struct { F0 uint8; F1 int32; F2 bool }
type GenS4 struct { F0 uint; F1 uint64 }
type GenS5 struct { F0 int64; F1 int; F2 int8 }
func genFn0(p0 uint32, p1 bool) int8 {
    fmt.Println(p0, p1)
    var v6 int16 = 0
    var v7 string = "foo bar"
    fmt.Println(p0, v7, p1)
    switch 32767 {
    case -2147483648:
        if (!(!false)) {
            v8 := GenS1{F0: -9223372036854775808}
            switch v7 {
            case "hello":
                v9 := []int{32767, 940}
                _ = v9
            case "a":
                v6 = v6
            default:
                v10 := GenS2{F0: ("\n\t" + ((v7 + v7) + v7)), F1: 1, F2: 0}
                _ = v10
            }
            v6 -= v6
            _ = v8
        }
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 100
}
func genFn1(p0 int16, p1 int64) uint32 {
    v11 := GenS3{F0: 2, F1: 42, F2: true}
    v12 := []int{709, 2147483647, 127, 128}
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 931
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]bool)
    var v3 float32 = (3.14159 - 1e10)
    var v4 int32 = 1
    var v5 bool = false
    v1 = (v1 & 255)
    v13 := GenS4{F0: 0, F1: 1}
    v14 := GenS5{F0: 105, F1: 32768, F2: 42}
    for fi15 := 0; fi15 < 8; fi15++ {
        fmt.Println(v3, v5, fi15)
        _ = fi15
    }
    var v16 uint8 = 105
    switch "\n\t" {
    case "hello":
        var v17 bool = false
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v1 -= v1
        var v19 uint16 = v1
        _ = fi18
        _ = v19
    }
    delete(v2, 2147483648)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
