package main

import (
    "fmt"
)

func genFn0(p0 string, p1 int64) string {
    var v5 int16 = 32767
    var v6 string = ("" + p0)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = (("" + "foo bar") + (("" + ("" + "a")) + (("hello" + "\n\t") + (("\n\t" + "a") + "\n\t"))))
    var v3 float64 = 1.5
    var v4 uint64 = 856
    var v7 uint16 = 255
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 14; fi9++ {
            var v10 float32 = 1e10
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := []int{708, 7, 50}
    _ = len(v11)
    fmt.Println(v3, v4)
    var v12 uint16 = (^(v7))
    var v13 int16 = -1
    for fi14 := 0; fi14 < 18; fi14++ {
        switch v2 {
        case "\n\t":
            fmt.Println(v13)
        default:
            switch fi14 {
            case -2147483648:
                v12 = uint16(v4)
            default:
                v12 = (v7 / ((315) | 1))
            }
        }
        _ = fi14
    }
    fmt.Println(v2, v2, v12)
    v12 += (^(v12))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
