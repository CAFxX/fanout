package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int8, p1 uint) uint16 {
    var v4 uint = p1
    v5 := []int{-926, -2147483648}
    fmt.Println(v4)
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 bool, p1 bool) int32 {
    v6 := []int{-1, 0, 1}
    var v7 uint64 = 0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return -535
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{1}
    var v3 bool = false
    v8 := GenS1{F0: 636}
    var v9 uint8 = 42
    v10 := GenS2{F0: 65535, F1: 4294967295}
    var v11 int = 2147483647
    v3 = (0.1 >= (float32(1e10) / 1.0))
    for fi12 := 0; fi12 < 5; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            var v14 uint8 = v9
            switch "foo bar" {
            case "\n\t":
                if v3 {
                    var v15 uint64 = 42
                    _ = v15
                }
            case "hello":
                if (!v3) {
                    var v16 int = 213
                    v14 = 42
                    v11 = (fi12 - v16)
                    v17 := GenS3{F0: 218}
                    var v18 float32 = 0.1
                    v19 := GenS3{F0: 0}
                    _ = v16
                    _ = v17
                    _ = v18
                    _ = v19
                } else {
                    v20 := GenS3{F0: 535}
                    var v21 float32 = float32(0.1)
                    _ = v20
                    _ = v21
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi13
            _ = v14
        }
        fmt.Println(v3, fi12)
        _ = fi12
    }
    v9 = v9
    fmt.Println(v3, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
