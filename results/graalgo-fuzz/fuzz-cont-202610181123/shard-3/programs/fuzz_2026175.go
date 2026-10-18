package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0(p0 int64, p1 uint32) int32 {
    var v7 uint32 = (p1 / ((2147483647) | 1))
    var v8 int = 4294967295
    v9 := []int{v8, 32767, v8, -9223372036854775808}
    var v10 uint8 = 255
    p0 -= (p0 >> ((2) & 63))
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 846
}

func main() {
    var v1 uint32 = 2
    var v2 float64 = 3.14159
    var v3 uint = 690
    v4 := []uint32{uint32(v3), v1, 114, v1}
    v5 := make(map[int]string)
    var v6 float32 = 1.5
    v5[-714] = "hello"
    v6 = (v6 * float32(1e10))
    var v11 uint32 = v1
    switch 32768 {
    case 546:
        var v12 int64 = 42
        _ = v12
    case 361:
        v6 = float32(v2)
    default:
        var v13 int16 = -842
        _ = v13
    }
    v1 += v11
    for fi14 := 0; fi14 < 17; fi14++ {
        fmt.Println(v3)
        var v15 float32 = v6
        _ = fi14
        _ = v15
    }
    var v16 float32 = v6
    var v17 uint8 = 128
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 uint32 = (v1 % (((v1 >> (((v3 + (v3 - 65535))) & 63))) | 1))
        for fi20 := 0; fi20 < 10; fi20++ {
            v16 -= v16
            v1 -= v1
            fmt.Println(v16, fi20, v2)
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    switch 32767 {
    case -889:
        fmt.Println(v17, v11)
    case 128:
        v11 = v1
    case -2147483648:
        for fi21 := 0; fi21 < 5; fi21++ {
            v4 = append(v4, 422)
            v22 := GenS1{F0: 42}
            _ = fi21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
