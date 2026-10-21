package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[string]bool)
    v3 := make(map[int]int)
    for fi4 := 0; fi4 < 7; fi4++ {
        if true {
            var v5 string = ""
            _ = v5
        } else {
            _ = len(v3)
            switch fi4 {
            case 811:
                for fi6 := 0; fi6 < 6; fi6++ {
                    var v7 float32 = 1.5
                    _ = len(v3)
                    v8 := []int{(fi4 / ((fi6) | 1)), 1}
                    _ = fi6
                    _ = v7
                    _ = v8
                }
            default:
                v9 := GenS1{F0: 516}
                _ = v9
            }
        }
        _ = fi4
    }
    var v10 int = 2147483647
    _ = len(v3)
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint16 = 2
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS1{F0: v12}
            v10 += (-(v10))
            _ = fi13
            _ = v14
        }
        var v15 uint = 4294967295
        _ = fi11
        _ = v12
        _ = v15
    }
    var v16 int16 = -184
    v17 := GenS1{F0: 42}
    if (false && false) {
        v18 := []int{(v10 & (v10 + (v10 * v10))), v10, 522}
        _ = v18
    }
    fmt.Println(v16)
    v10 -= (v10 >> ((42) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
