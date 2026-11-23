package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    v1 := []uint32{}
    var v2 int8 = -32
    var v3 int32 = 831
    v4 := make(map[uint32]bool)
    var v5 int = 128
    delete(v4, 2147483648)
    v6 := GenS1{F0: 1, F1: 634}
    delete(v4, 42)
    _ = len(v4)
    v3 = (v3 << ((0) & 63))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 += -47
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            var v9 uint8 = 120
            var v10 float64 = 1.5
            _ = len(v4)
            _ = fi8
            _ = v10
            _ = v9
        }
        if (523 <= 1) {
            if true {
                var v11 int = 846
                for fi12 := 0; fi12 < 7; fi12++ {
                    if false {
                        v13 := GenS1{F0: 0, F1: 4294967295}
                        v5 = v11
                        _ = v13
                    }
                    _ = fi12
                }
                fmt.Println(v2)
                _ = v11
            }
        } else {
            v2 -= 100
            var v14 uint16 = 1
            _ = v14
        }
        var v15 int16 = 42
        _ = fi7
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
