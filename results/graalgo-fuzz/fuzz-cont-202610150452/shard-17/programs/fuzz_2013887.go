package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 uint64, p1 int8) uint8 {
    var v4 int64 = -1
    var v5 int32 = -434
    fmt.Println(p0, p1, p1)
    v6 := GenS1{F0: 994, F1: 168}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return 128
}

func main() {
    var v1 uint8 = 127
    var v2 int32 = -898
    var v3 int16 = -813
    v1 = 42
    var v7 string = "hello"
    var v8 float32 = (0.0 * 0.0)
    var v9 int8 = 1
    switch 32768 {
    case 4294967295:
        if true {
            v8 += v8
            var v10 int64 = 0
            var v11 uint64 = 18446744073709551615
            _ = v10
            _ = v11
        } else {
            var v12 int = 32767
            var v13 uint16 = 544
            _ = v12
            _ = v13
        }
    case 32767:
        v8 += 1.5
    }
    var v14 uint32 = 2147483648
    v2 -= v2
    v15 := GenS2{F0: 4294967295, F1: v2}
    var v16 int32 = (v2 >> ((42) & 63))
    if false {
        var v17 int32 = int32(v3)
        _ = v17
    }
    v3 += v3
    fmt.Println(v9, v7)
    v18 := GenS2{F0: 0, F1: (v2 + v2)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
