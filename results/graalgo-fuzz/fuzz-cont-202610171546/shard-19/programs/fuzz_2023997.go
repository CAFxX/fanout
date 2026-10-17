package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 uint32 }
func genFn0(p0 uint, p1 string) int64 {
    fmt.Println(p0, p0)
    p0 = (p0 << ((18446744073709551615) & 63))
    _ = p0
    _ = p1
    return -9223372036854775808
}

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    var v4 uint64 = 18446744073709551615
    var v5 float32 = (-2.25 + (3.14159 + 3.14159))
    v6 := []uint32{42, 65535, 4294967295}
    fmt.Println(v5, v1, v1)
    v7 := GenS1{F0: 65535}
    v8 := GenS1{F0: 404}
    if false {
        switch 2147483647 {
        case -2147483648:
            var v9 int32 = 27
            _ = v9
        case 128:
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            v10 := GenS2{F0: 0, F1: 42}
            _ = v10
        }
        v6 = append(v6, 134)
    }
    var v11 int32 = -2147483648
    var v12 uint64 = (^(v4))
    v1 = "\n\t"
    for fi13 := 0; fi13 < 10; fi13++ {
        _ = len(v2)
        if true {
            v4 = 18446744073709551615
        }
        _ = fi13
    }
    v6 = append(v6, 42)
    v6 = append(v6, 42)
    v1 = ""
    v5 += (float32(-2.25) / v5)
    var v14 uint32 = 1
    var v15 uint = 2
    fmt.Println(v14, v12)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
