package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 float64 = ((1.5 - 1.5) - float64((float32(0.1) + 0.1)))
    var v3 int32 = int32(v1)
    v4 := make(map[uint32]int)
    v5 := make(map[uint32]uint32)
    var v6 int64 = 42
    fmt.Println(v3)
    var v7 string = (("" + "hello") + "a")
    v2 = 3.14159
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            if (3.14159 == 0.0) {
                v4[1] = (fi9 / ((-948) | 1))
                v10 := GenS1{F0: 2}
                _ = v10
            }
            v2 += (v2 * v2)
            var v11 int8 = 100
            var v12 uint8 = uint8(v1)
            _ = fi9
            _ = v11
            _ = v12
        }
        var v13 int8 = -32
        v14 := GenS2{F0: 65535}
        _ = fi8
        _ = v13
        _ = v14
    }
    switch -2147483648 {
    case -2147483648:
        var v15 int8 = -106
        _ = v15
    }
    v5[2] = 1
    var v16 uint32 = 113
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
