package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint }
func genFn0(p0 int64, p1 bool) uint {
    v5 := []int{2147483647, 32768, 1, 128}
    _ = len(v5)
    p0 = p0
    p1 = false
    _ = p0
    _ = p1
    _ = v5
    return 255
}

func main() {
    var v1 uint8 = 207
    var v2 uint32 = 905
    var v3 uint32 = v2
    var v4 int = 685
    v2 = v2
    if true {
        var v6 uint32 = 4294967295
        var v7 bool = false
        for fi8 := 0; fi8 < 19; fi8++ {
            v9 := GenS1{F0: -128, F1: (v4 >> ((42) & 63))}
            fmt.Println(v3, v7)
            v10 := GenS1{F0: -11, F1: v4}
            _ = fi8
            _ = v10
            _ = v9
        }
        _ = v6
        _ = v7
    } else {
        if ((false || (42 == 99)) || false) {
            var v11 int16 = -1
            v4 += 128
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := []int{v4, v4, (v4 % ((128) | 1)), fi12}
                fmt.Println(v1, v1, v3)
                v3 -= 2147483647
                _ = fi12
                _ = v13
            }
            _ = v11
        } else {
            v1 = (v1 >> ((1) & 63))
            for fi14 := 0; fi14 < 13; fi14++ {
                fmt.Println(v4, v2)
                v15 := []int{(v4 * (-(v4))), (v4 % ((v4) | 1)), fi14, 32767}
                _ = fi14
                _ = v15
            }
        }
        var v16 uint32 = v3
        _ = v16
    }
    for fi17 := 0; fi17 < 1; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            if (!(!true)) {
                v19 := []int{fi18, (-(fi18)), fi18}
                v3 = v3
                _ = v19
            } else {
                var v20 float64 = (-2.25 * 1e10)
                fmt.Println(fi17, v4, fi17)
                _ = v20
            }
            _ = fi18
        }
        var v21 uint16 = 255
        _ = fi17
        _ = v21
    }
    v3 = (^(v2))
    var v22 float64 = (1.5 * 0.0)
    v22 = ((float64(3.14159) / ((v22 - v22) - v22)) / v22)
    v23 := GenS2{F0: 926}
    v1 = v1
    v24 := GenS1{F0: -128, F1: -1}
    for fi25 := 0; fi25 < 3; fi25++ {
        v22 = v22
        var v26 int = (^(v4))
        _ = fi25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
