package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint; F1 uint16 }
func genFn0(p0 bool) uint16 {
    p0 = (255 >= 2)
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 42}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    v8 := []int{32768, 9223372036854775807}
    _ = p0
    _ = v8
    return 916
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 128
    var v5 uint16 = 289
    if true {
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint8 = (v4 << ((65535) & 63))
            _ = fi9
            _ = v10
        }
        v5 = (v5 % ((genFn0((-1 != -1))) | 1))
    } else {
        fmt.Println(v1)
        v11 := GenS1{F0: 42}
        v5 = 65535
        _ = v11
    }
    v12 := GenS2{F0: -2147483648}
    var v13 float64 = (1.5 * -2.25)
    v13 = 1.5
    var v14 string = v1
    v15 := []int{}
    v16 := GenS1{F0: 42}
    v4 -= (v4 << ((18446744073709551615) & 63))
    v17 := GenS2{F0: 0}
    var v18 uint = 18446744073709551615
    v4 += 15
    var v19 uint16 = (^(v5))
    v2[2147483647] = 959
    var v20 uint32 = 342
    var v21 int8 = 100
    v22 := GenS3{F0: (v18 << ((4294967295) & 63)), F1: genFn0((((float32(3.14159) + (1.5 + float32(3.14159))) - float32((3.14159 / (1.5 / v13)))) > 0.0))}
    var v23 uint = v18
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
