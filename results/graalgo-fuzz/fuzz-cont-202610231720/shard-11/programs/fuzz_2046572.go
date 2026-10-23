package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 uint32) int16 {
    var v6 string = ("" + "\n\t")
    var v7 uint16 = 788
    fmt.Println(v6, p0)
    var v8 float32 = float32(1.5)
    if true {
        v9 := GenS1{F0: 49, F1: 0, F2: 4294967295}
        v6 = "\n\t"
        v10 := GenS2{F0: 4294967295, F1: 0}
        _ = v10
        _ = v9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 635
}
func genFn1(p0 int16) int16 {
    v11 := []int{32767, -151, 1}
    var v12 int32 = -2147483648
    v13 := GenS3{F0: -1, F1: 531}
    _ = p0
    _ = v11
    _ = v12
    _ = v13
    return (p0 << ((42) & 63))
}

func main() {
    var v1 int = 893
    var v2 float32 = 3.14159
    v3 := make(map[string]uint32)
    var v4 int32 = 0
    var v5 string = ("foo bar" + ("a" + ("foo bar" + "hello")))
    switch ((v5 + "hello") + v5) {
    case "\n\t":
        v14 := GenS1{F0: 497, F1: 32767, F2: 898}
        _ = v14
    case "":
        fmt.Println(v5, v1, v2)
    case "hello":
        for fi15 := 0; fi15 < 16; fi15++ {
            var v16 uint8 = 221
            v3[v5] = 2147483647
            _ = fi15
            _ = v16
        }
    }
    var v17 int = -350
    v5 = v5
    v18 := GenS1{F0: 65535, F1: 512, F2: 42}
    v19 := []int{(-(v17)), -855, (v17 - v1)}
    fmt.Println(v17)
    switch "foo bar" {
    case "":
        v17 = v17
    default:
        fmt.Println(v1, v4)
    }
    var v20 int8 = 1
    v1 += v1
    v5 = ("foo bar" + "")
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
