package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
func genFn0(p0 int8) int {
    fmt.Println(p0)
    p0 = 127
    v4 := []int{-9223372036854775808}
    fmt.Println(p0)
    p0 = (p0 / ((p0) | 1))
    _ = p0
    _ = v4
    return 323
}

func main() {
    v1 := []int{-828, -869}
    v2 := make(map[uint32]bool)
    var v3 int32 = 42
    v3 -= 250
    if (197 < 195) {
        v5 := v2[65535]
        _ = v5
    }
    v6 := v2[4294967295]
    delete(v2, 4294967295)
    for fi7 := 0; fi7 < 7; fi7++ {
        v3 += (^(v3))
        for fi8 := 0; fi8 < 6; fi8++ {
            v6 = v6
            var v9 uint = 602
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    switch 0 {
    case 4294967295:
        if (100 < -2) {
            for fi10 := 0; fi10 < 20; fi10++ {
                delete(v2, 2)
                _ = fi10
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := GenS1{F0: 697, F1: "\n\t"}
    fmt.Println(v6, v3)
    if (-779 >= -1) {
        _ = len(v2)
    }
    v3 = 346
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v6
}
