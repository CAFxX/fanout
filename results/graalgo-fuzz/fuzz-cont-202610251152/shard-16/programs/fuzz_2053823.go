package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
func genFn0(p0 int64) uint16 {
    var v6 int32 = -879
    v7 := []int{}
    _ = p0
    _ = v6
    _ = v7
    return 0
}

func main() {
    var v1 float64 = (-2.25 + (-2.25 / -2.25))
    var v2 float64 = v1
    var v3 int32 = 42
    var v4 int16 = 0
    var v5 float32 = 1.5
    for fi8 := 0; fi8 < 16; fi8++ {
        v4 = (v4 % ((int16(fi8)) | 1))
        _ = fi8
    }
    var v9 uint8 = 0
    v10 := GenS1{F0: (v4 * (-(v4)))}
    v11 := GenS1{F0: v4}
    if ((v9 - 1) >= v9) {
        var v12 bool = false
        v4 = (^(v4))
        _ = v12
    } else {
        v5 = v5
        var v13 uint = 1
        v14 := []int{127}
        _ = v13
        _ = v14
    }
    v15 := GenS2{F0: (("hello" + "a") + "a")}
    v1 += 0.1
    for fi16 := 0; fi16 < 18; fi16++ {
        fmt.Println(v9, v4)
        _ = fi16
    }
    v4 = -898
    v5 = float32(((float64(v5) * (v1 * (1e10 + 1e10))) - (v2 / 1.5)))
    v3 = (v3 & v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
