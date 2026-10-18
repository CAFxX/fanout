package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 42
    v3 := make(map[string]uint32)
    var v4 uint64 = uint64(v1)
    v5 := make(map[string]int)
    v1 = (v1 & v1)
    if ((1e10 >= (-2.25 + 3.14159)) && (42 > 0)) {
        v4 = 2
        v1 -= v1
    }
    var v6 float64 = (0.0 * 3.14159)
    for fi7 := 0; fi7 < 13; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 float64 = 1.5
            var v10 uint8 = v2
            if (-822 != -102) {
                if (2147483647 != 0) {
                    fmt.Println(v2, v6)
                }
                v4 -= 0
                if (!(false && (0.1 != 1.5))) {
                    fmt.Println(v9, v9)
                }
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS1{F0: 0, F1: 0}
        _ = fi7
        _ = v11
    }
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 5; fi13++ {
            v14 := []int{-1, (fi12 & fi12), fi12}
            _ = fi13
            _ = v14
        }
        fmt.Println(v1)
        v15 := GenS1{F0: -1, F1: 1}
        _ = fi12
        _ = v15
    }
    v16 := GenS2{F0: 0}
    fmt.Println(v2, v6, v4)
    v1 = v1
    var v17 uint8 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
