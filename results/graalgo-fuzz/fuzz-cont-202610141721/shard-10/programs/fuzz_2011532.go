package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = -2.25
    var v3 int32 = -1
    var v4 float32 = (float32((v2 + (((v2 - (v2 - (v2 / ((v2 + (3.14159 - v2)) + (v2 * v2))))) * v2) + v2))) * 0.0)
    for fi5 := 0; fi5 < 7; fi5++ {
        if (!(-1 != -1)) {
            for fi6 := 0; fi6 < 9; fi6++ {
                v7 := GenS1{F0: 128}
                _ = fi6
                _ = v7
            }
            switch (fi5 | fi5) {
            case 0:
                delete(v1, fi5)
            default:
                var v8 uint = 42
                _ = v8
            }
        } else {
            v9 := GenS2{F0: 0}
            _ = v9
        }
        fmt.Println(v4, v2)
        v4 = v4
        _ = fi5
    }
    fmt.Println(v3, v2)
    fmt.Println(v4)
    for fi10 := 0; fi10 < 1; fi10++ {
        if (-1 >= 32767) {
            v3 += v3
            v11 := []int{fi10, 42, fi10, fi10}
            _ = v11
        } else {
            v3 = 0
            v4 = float32(0.1)
        }
        v12 := GenS1{F0: 181}
        _ = fi10
        _ = v12
    }
    v3 = 0
    delete(v1, -1)
    for fi13 := 0; fi13 < 4; fi13++ {
        for fi14 := 0; fi14 < 5; fi14++ {
            var v15 uint = 1
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    v16 := []int{128, -1, 9223372036854775807, 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
