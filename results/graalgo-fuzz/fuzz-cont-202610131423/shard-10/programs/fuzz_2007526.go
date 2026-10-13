package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 0
    v2 := make(map[uint32]string)
    var v3 float64 = 1e10
    var v4 string = "\n\t"
    var v5 uint32 = 2147483648
    fmt.Println(v1, v3, v5)
    v4 = ("hello" + v4)
    var v6 int = 128
    if (!true) {
        var v7 int32 = 596
        _ = v7
    } else {
        switch "a" {
        case "\n\t":
            v1 += (v1 / ((v1) | 1))
        default:
            for fi8 := 0; fi8 < 5; fi8++ {
                delete(v2, (^(v5)))
                v9 := v2[v5]
                _ = fi8
                _ = v9
            }
        }
        var v10 float32 = float32(0.1)
        v10 = (1.5 + ((v10 * float32(v3)) + v10))
        _ = v10
    }
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := []int{(-(v6)), fi11, int(v5), (fi11 >> ((340) & 63))}
        var v13 int64 = 1
        var v14 float64 = 0.0
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    v3 += (3.14159 - (v3 * v3))
    v15 := []int{}
    fmt.Println(v5, v5, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
