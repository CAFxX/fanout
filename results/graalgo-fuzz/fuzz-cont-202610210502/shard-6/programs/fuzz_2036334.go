package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []string{}
    v2 := make(map[uint32]int)
    var v3 string = "a"
    var v4 float32 = ((3.14159 / 1e10) - 0.0)
    if (true && (true && (true && true))) {
        if false {
            v4 = ((v4 - 3.14159) + ((0.1 * ((v4 * (float32((float64(float32(0.0)) / (1.5 - float64(0.0)))) + ((-2.25 * (v4 / 1.0)) / 3.14159))) / v4)) - float32((1e10 + (0.1 + float64(1e10))))))
        } else {
            for fi5 := 0; fi5 < 19; fi5++ {
                fmt.Println(v3)
                _ = fi5
            }
            fmt.Println(v3, v4)
            _ = len(v2)
        }
        for fi6 := 0; fi6 < 4; fi6++ {
            v7 := v2[4294967295]
            delete(v2, 596)
            _ = fi6
            _ = v7
        }
    } else {
        v4 = v4
        fmt.Println(v3, v4)
        v4 = (float32(1.5) / v4)
    }
    v4 = v4
    v4 = v4
    v8 := GenS1{F0: 0}
    var v9 int16 = -1
    var v10 uint64 = 4294967295
    v11 := GenS1{F0: 255}
    v12 := GenS1{F0: 255}
    for fi13 := 0; fi13 < 15; fi13++ {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint64 = v10
            _ = fi14
            _ = v15
        }
        fmt.Println(fi13, fi13, fi13)
        _ = fi13
    }
    var v16 int = -668
    v17 := GenS1{F0: 1}
    var v18 int64 = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
