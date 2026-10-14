package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8) int16 {
    var v5 uint32 = 2147483648
    v6 := []int{32767, 42, 798}
    var v7 uint64 = 2
    _ = len(v6)
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return -271
}
func genFn1(p0 uint, p1 int16) uint8 {
    var v8 uint64 = 24
    for fi9 := 0; fi9 < 2; fi9++ {
        p0 = 65535
        v10 := []int{(^(fi9)), fi9}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 42}
    var v12 int = 164
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v8
    return 42
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = ("a" + ("foo bar" + "foo bar"))
    var v3 int32 = -64
    v4 := make(map[int]string)
    var v13 uint32 = 2147483648
    v14 := GenS1{F0: 1}
    v15 := GenS2{F0: 1, F1: 128}
    _ = len(v1)
    v2 = v2
    switch "foo bar" {
    case "\n\t":
        switch 114 {
        case 4294967295:
            fmt.Println(v13, v2, v2)
        case 128:
            v16 := v1[2147483647]
            _ = v16
        default:
            fmt.Println(v2)
        }
    }
    v17 := GenS2{F0: 1, F1: 127}
    var v18 int32 = v3
    v2 = "a"
    v19 := GenS2{F0: 4294967295, F1: 2}
    v20 := []int{0, 415}
    fmt.Println(v18, v2, v2)
    v21 := GenS3{F0: 405}
    var v22 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
