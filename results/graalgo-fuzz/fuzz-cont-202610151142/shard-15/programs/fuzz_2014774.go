package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 int64, p1 string, p2 uint8) uint64 {
    if (3.14159 < ((float64(0.0) * (((3.14159 + (-2.25 - float64(1.5))) + -2.25) + (((float64(1e10) - 1.5) + 0.1) * 0.0))) / 0.1)) {
        if false {
            if true {
                var v6 uint8 = p2
                if (!false) {
                    var v7 uint32 = 42
                    for fi8 := 0; fi8 < 16; fi8++ {
                        var v9 int32 = 636
                        _ = fi8
                        _ = v9
                    }
                    v10 := GenS1{F0: -179}
                    _ = v10
                    _ = v7
                } else {
                    var v11 uint16 = 859
                    p1 = (p1 + "hello")
                    p2 += 162
                    _ = v11
                }
                _ = v6
            }
        } else {
            fmt.Println(p2, p2, p2)
        }
        v12 := []int{}
        _ = v12
    }
    v13 := []int{42, 42, -745, 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 4294967295
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = 1.5
    var v3 int32 = 60
    var v4 int64 = 0
    v5 := make(map[uint32]uint32)
    if ((!true) || ("" > ("foo bar" + ("\n\t" + "")))) {
        v14 := []int{32768}
        v15 := GenS2{F0: 1000, F1: 644}
        _ = v14
        _ = v15
    } else {
        v2 += v2
        v16 := []int{127}
        _ = v16
    }
    var v17 int64 = v4
    for fi18 := 0; fi18 < 1; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            fmt.Println(v4, v4, v3)
            _ = fi19
        }
        v4 -= v17
        fmt.Println(v1, v2)
        _ = fi18
    }
    v20 := GenS2{F0: 222, F1: 42}
    v21 := []int{812}
    v4 += v17
    v4 += v17
    v3 -= 0
    v1 -= (v1 | (v1 / (((v1 + (^(v1)))) | 1)))
    fmt.Println(v1, v3)
    switch -2147483648 {
    case -16:
        v3 = v3
    }
    for fi22 := 0; fi22 < 16; fi22++ {
        fmt.Println(v1, v4, v17)
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
