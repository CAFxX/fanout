package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 uint32 }
func genFn0(p0 string, p1 bool, p2 int) uint32 {
    p0 = ((p0 + p0) + ("hello" + p0))
    fmt.Println(p1, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    return 65535
}
func genFn1() uint8 {
    v6 := []int{1, 2147483647}
    var v7 uint = 834
    var v8 uint8 = 255
    v7 += 215
    for fi9 := 0; fi9 < 14; fi9++ {
        fmt.Println(v8, v8)
        var v10 uint64 = 0
        var v11 uint16 = 65535
        _ = fi9
        _ = v10
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    return (v8 % (((v8 << (((v7 / (((v7 + v7)) | 1))) & 63))) | 1))
}

func main() {
    var v1 float32 = (float32(0.0) * 1.5)
    v2 := []int{0}
    v3 := make(map[string]uint32)
    var v4 uint16 = 762
    v5 := make(map[string]string)
    _ = len(v2)
    v12 := v3[""]
    v13 := v5["hello"]
    v12 = v12
    var v14 int32 = -104
    v14 += v14
    var v15 int8 = -95
    switch 4294967295 {
    case -237:
        switch v13 {
        case "foo bar":
            var v16 int = -842
            _ = v16
        case "hello":
            v17 := v3[(("a" + v13) + "\n\t")]
            _ = v17
        default:
            v18 := GenS1{F0: (v14 & (v14 << ((619) & 63)))}
            _ = v18
        }
    case -1:
        var v19 string = v13
        _ = v19
    }
    v20 := GenS2{F0: -9223372036854775808, F1: (^(v12))}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
