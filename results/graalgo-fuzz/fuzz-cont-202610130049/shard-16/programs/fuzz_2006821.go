package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
func genFn0(p0 int16) uint16 {
    v6 := []int{0}
    if true {
        var v7 int16 = 943
        var v8 int8 = 0
        _ = v7
        _ = v8
    } else {
        var v9 float64 = 1.5
        v10 := GenS1{F0: 0}
        if len(v6) > 0 {
            v6[0] = 523
        }
        _ = v10
        _ = v9
    }
    fmt.Println(p0)
    if len(v6) > 0 {
        v6[0] = 32768
    }
    _ = p0
    _ = v6
    return 0
}

func main() {
    var v1 uint = 824
    var v2 float64 = 0.0
    var v3 bool = (128 >= 2)
    v4 := make(map[int]int)
    v5 := []string{"\n\t", "\n\t", "\n\t", "a"}
    var v11 uint64 = 0
    _ = len(v5)
    var v12 int8 = -1
    v13 := GenS1{F0: 42}
    for fi14 := 0; fi14 < 15; fi14++ {
        var v15 float64 = (-2.25 / (v2 + (v2 * v2)))
        _ = fi14
        _ = v15
    }
    if ((-1 > 1) && (v2 > (1.5 + float64((3.14159 * (float32(3.14159) + (0.0 - float32((v2 / (float64(((1.5 - (-2.25 / ((-2.25 * (1e10 + 0.0)) - 0.0))) + float32(1.5))) - v2)))))))))) {
        v1 = 255
        var v16 uint16 = 0
        _ = v16
    }
    fmt.Println(v1, v12, v12)
    var v17 int64 = 1
    delete(v4, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
