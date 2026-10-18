package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int8, p1 uint32) uint8 {
    fmt.Println(p1, p1)
    p0 = (p0 >> ((1) & 63))
    p1 -= 23
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := make(map[int]int)
    var v2 uint16 = 255
    v3 := make(map[int]string)
    var v4 float64 = ((-2.25 - -2.25) - -2.25)
    var v5 bool = false
    v6 := make(map[uint32]int)
    var v7 uint16 = 1
    var v8 int32 = -1
    var v9 int = 0
    if v5 {
        var v10 int8 = 42
        _ = v10
    } else {
        for fi11 := 0; fi11 < 19; fi11++ {
            for fi12 := 0; fi12 < 5; fi12++ {
                if (0 != 0) {
                    v4 = ((float64((0.0 / 1e10)) / v4) + 1e10)
                    var v13 string = "\n\t"
                    var v14 uint64 = 637
                    _ = v13
                    _ = v14
                }
                _ = fi12
            }
            _ = fi11
        }
        var v15 float64 = v4
        switch "\n\t" {
        case "a":
            var v16 int = (v9 ^ (^(v9)))
            _ = v16
        default:
            v7 += (^(v7))
        }
        _ = v15
    }
    v17 := []int{(v9 & (-(v9)))}
    v18 := GenS1{F0: v5}
    for fi19 := 0; fi19 < 1; fi19++ {
        var v20 float64 = v4
        _ = fi19
        _ = v20
    }
    fmt.Println(v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
