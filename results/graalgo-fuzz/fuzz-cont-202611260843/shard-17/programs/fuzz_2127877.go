package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int32; F1 uint; F2 uint32 }
type GenS3 struct { F0 int64; F1 bool }

func main() {
    var v1 int16 = -32768
    var v2 uint8 = 0
    var v3 uint = 1
    var v4 int8 = -1
    v5 := make(map[string]uint32)
    var v6 uint64 = 1
    var v7 int16 = (v1 * 1)
    if (-1 < 2147483647) {
        v8 := GenS1{F0: 255}
        if (!true) {
            if true {
                v9 := GenS2{F0: -569, F1: v3, F2: 1}
                v10 := v5["foo bar"]
                _ = v10
                _ = v9
            }
            delete(v5, "hello")
        } else {
            fmt.Println(v6, v3)
        }
        _ = v8
    }
    v6 = (v6 ^ v6)
    switch 564 {
    case -1:
        var v11 uint64 = 4294967295
        _ = v11
    case 9223372036854775807:
        v12 := []int{32768, 9223372036854775807, -2147483648}
        _ = v12
    case 32768:
        fmt.Println(v2, v4)
    default:
        for fi13 := 0; fi13 < 20; fi13++ {
            v6 = v6
            fmt.Println(v4)
            v14 := []int{fi13, fi13}
            _ = fi13
            _ = v14
        }
    }
    v15 := []int{-9223372036854775808, -9223372036854775808, 128, -9223372036854775808}
    v7 -= (v1 * (v1 & v1))
    var v16 uint16 = 847
    v5["a"] = 988
    var v17 float32 = 1e10
    var v18 int16 = v1
    v7 -= (v1 | (v18 & (v18 >> ((1) & 63))))
    v15 = append(v15, 161)
    for fi19 := 0; fi19 < 5; fi19++ {
        v2 = 128
        fmt.Println(v16, v1)
        _ = fi19
    }
    fmt.Println(v6, v16, v1)
    v5[((("a" + ("hello" + "foo bar")) + "") + ("hello" + "foo bar"))] = 157
    v15 = append(v15, 1)
    var v20 uint64 = 18446744073709551615
    v21 := GenS3{F0: -1, F1: (-994 <= -2147483648)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
