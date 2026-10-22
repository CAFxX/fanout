package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 int; F2 int8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = ((((0.1 - 0.1) / (0.0 - (1e10 + 1e10))) - (((1.5 - 0.1) + 1.5) / 1.0)) / 1.5)
    var v3 int = 921
    var v4 float64 = (0.1 * -2.25)
    if (42 == 4294967295) {
        fmt.Println(v4, v3)
        if (1000 != 1000) {
            var v5 float32 = (0.0 * v2)
            var v6 float64 = 3.14159
            var v7 float32 = ((v2 - (float32(0.1) * v2)) + v2)
            _ = v5
            _ = v6
            _ = v7
        } else {
            delete(v1, (v3 & (v3 % ((v3) | 1))))
            v3 -= (-(v3))
            var v8 uint8 = 2
            _ = v8
        }
        delete(v1, (-(v3)))
    } else {
        v9 := GenS1{F0: 4294967295}
        var v10 uint16 = 2
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 10; fi11++ {
        v3 = fi11
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 uint64 = 1
            if true {
                var v14 float64 = v4
                _ = v14
            }
            _ = fi12
            _ = v13
        }
        v15 := GenS1{F0: 2}
        _ = fi11
        _ = v15
    }
    var v16 float32 = v2
    if (false || false) {
        delete(v1, v3)
        for fi17 := 0; fi17 < 20; fi17++ {
            v18 := []int{(fi17 / ((32767) | 1)), 128, -1, 0}
            if len(v18) >= 2 {
                v18 = v18[1:len(v18)]
            }
            _ = fi17
            _ = v18
        }
    } else {
        v16 -= v16
        var v19 int8 = 87
        fmt.Println(v19)
        _ = v19
    }
    v2 -= (v16 / float32((v4 / ((v4 - v4) * v4))))
    var v20 int = v3
    v21 := GenS2{F0: 18446744073709551615, F1: 0, F2: 100}
    v20 += -1
    v22 := GenS2{F0: 286, F1: 127, F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
