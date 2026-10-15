package main

import (
    "fmt"
)

func genFn0(p0 uint) uint8 {
    fmt.Println(p0)
    var v4 uint64 = 18446744073709551615
    for fi5 := 0; fi5 < 9; fi5++ {
        v4 += v4
        _ = fi5
    }
    switch 32768 {
    case 127:
        p0 = 582
    case 482:
        p0 += (p0 * 18446744073709551615)
    case 125:
        v4 = 18446744073709551615
    }
    _ = p0
    _ = v4
    return 255
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]int)
    var v3 float64 = 3.14159
    var v6 int64 = 1
    for fi7 := 0; fi7 < 16; fi7++ {
        switch fi7 {
        case 4294967295:
            var v8 uint16 = 1
            _ = v8
        default:
            v6 += v6
        }
        var v9 uint16 = 2
        v10 := []int{(fi7 % ((1) | 1)), (fi7 | fi7)}
        _ = fi7
        _ = v10
        _ = v9
    }
    v3 -= 0.1
    var v11 bool = v1
    v6 -= 1
    fmt.Println(v6, v11)
    var v12 int64 = v6
    v3 += 0.1
    v13 := []int{-125}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v6
}
