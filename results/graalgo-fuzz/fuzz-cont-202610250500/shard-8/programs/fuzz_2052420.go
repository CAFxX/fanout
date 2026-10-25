package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint16; F1 int8 }
func genFn0(p0 uint, p1 uint32) int64 {
    var v7 uint32 = p1
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(p0, p1, p0)
        fmt.Println(p1)
        var v9 bool = false
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 9223372036854775807
}
func genFn1(p0 string, p1 int8) uint64 {
    var v10 int16 = -809
    var v11 string = ((p0 + "") + (p0 + p0))
    var v12 float64 = 3.14159
    v13 := GenS1{F0: 2147483647, F1: v10}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    var v3 uint64 = 42
    v4 := make(map[string]int)
    var v5 uint32 = 2147483647
    v6 := make(map[string]string)
    fmt.Println(v5)
    var v14 float32 = (0.0 * -2.25)
    v3 = (v3 >> ((827) & 63))
    v15 := GenS2{F0: v3}
    for fi16 := 0; fi16 < 11; fi16++ {
        switch ("hello" + "\n\t") {
        case "\n\t":
            v17 := []int{}
            _ = v17
        default:
            if false {
                var v18 uint64 = (v3 >> ((0) & 63))
                switch (-(fi16)) {
                case 32767:
                    v19 := v6[""]
                    _ = v19
                default:
                    fmt.Println(v3, v14)
                }
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := GenS1{F0: (fi20 - 32767), F1: -936}
                    v5 += v5
                    v22 := GenS3{F0: uint16(v5), F1: 100}
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v18
            } else {
                v23 := []int{1, (^(fi16)), fi16, fi16}
                v14 = float32(1e10)
                _ = v23
            }
        }
        fmt.Println(v14, v2, v3)
        _ = fi16
    }
    fmt.Println(v2, v2, v14)
    v3 += v3
    v2 = (((-1 <= 0) && (v2 && v2)) && (v5 >= 926))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
