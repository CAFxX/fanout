package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 uint; F2 string }

func main() {
    var v1 uint = 817
    v2 := make(map[uint32]int)
    var v3 uint32 = 0
    var v4 uint = (v1 - (v1 * 255))
    var v5 int32 = 0
    fmt.Println(v4, v5)
    v6 := GenS1{F0: 106}
    if (!(1e10 == ((float32((1e10 * 1e10)) - 1e10) * (1.5 / -2.25)))) {
        v7 := v2[65535]
        var v8 int32 = (v5 % ((-1) | 1))
        _ = v7
        _ = v8
    } else {
        v1 -= 723
        var v9 bool = (((v3 * (v3 + v3)) == v3) && (v5 == (v5 << (((v1 & v1)) & 63))))
        _ = v9
    }
    for fi10 := 0; fi10 < 18; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS1{F0: -1}
            var v13 uint = v1
            v5 += (v5 - -1)
            v1 += v13
            _ = fi11
            _ = v12
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            v15 := GenS1{F0: 101}
            v16 := GenS1{F0: 100}
            _ = fi14
            _ = v15
            _ = v16
        }
        v17 := []int{}
        _ = fi10
        _ = v17
    }
    var v18 int8 = 0
    v19 := GenS2{F0: true, F1: v1, F2: ("" + ("a" + "hello"))}
    var v20 uint64 = 122
    for fi21 := 0; fi21 < 5; fi21++ {
        v18 += (v18 + v18)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
