package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int, p1 uint8, p2 int8) uint64 {
    var v7 uint = 255
    var v8 int64 = 9223372036854775807
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint16 = 954
    v2 := []string{}
    v3 := make(map[int]uint32)
    v4 := []string{"hello", "foo bar"}
    var v5 string = ("" + "\n\t")
    var v6 uint = 255
    for fi9 := 0; fi9 < 13; fi9++ {
        v10 := GenS1{F0: (v6 ^ (v6 + v6))}
        if true {
            v1 -= 1
            _ = len(v4)
        }
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint32 = 2147483647
            if len(v2) > 0 {
                v2[0] = v5
            }
            _ = fi11
            _ = v12
        }
        _ = fi9
        _ = v10
    }
    v13 := GenS2{F0: 197}
    var v14 uint32 = 418
    var v15 uint = v6
    var v16 bool = ((-2.25 < float64((-2.25 + (((0.1 * (float32((1e10 / 1e10)) + ((0.0 + (0.0 - -2.25)) * 3.14159))) - 0.0) + -2.25)))) || false)
    var v17 int16 = 394
    var v18 uint32 = (v14 << (((v15 + v15)) & 63))
    var v19 int32 = 42
    var v20 string = v5
    v21 := GenS2{F0: 1}
    var v22 bool = (v16 || ((v1 ^ v1) > 0))
    v23 := GenS1{F0: v6}
    var v24 int32 = v19
    fmt.Println(v16, v24, v6)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
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
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
