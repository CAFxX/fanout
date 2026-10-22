package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]string)
    var v2 uint32 = 2147483647
    v3 := []uint32{}
    v2 = 763
    var v4 uint32 = v2
    switch 442 {
    case 4294967295:
        switch 336 {
        case -135:
            var v5 float64 = 1.5
            _ = v5
        case -93:
            var v6 float32 = (((1.5 / 3.14159) - -2.25) / 1.0)
            _ = v6
        case 32767:
            if ((true || (!false)) && true) {
                var v7 string = "\n\t"
                _ = v7
            }
        default:
            v4 = v4
        }
    case 127:
        v2 = v2
    }
    fmt.Println(v4)
    for fi8 := 0; fi8 < 18; fi8++ {
        if len(v3) > 0 {
            v3[0] = v2
        }
        _ = fi8
    }
    delete(v1, 128)
    v3 = append(v3, (v2 >> ((578) & 63)))
    v3 = append(v3, (^(v2)))
    var v9 int32 = int32(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
