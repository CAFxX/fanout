package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    var v1 uint = 183
    var v2 int8 = 42
    var v3 string = "\n\t"
    v4 := make(map[string]uint32)
    v5 := []int{128, 2147483647}
    var v6 float32 = 0.0
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 float64 = float64(v6)
        if (!(!((false || (false && false)) || true))) {
            var v9 string = "\n\t"
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    var v10 bool = false
    var v11 int8 = (v2 * 127)
    var v12 float64 = 0.1
    var v13 uint = (v1 - (^(v1)))
    var v14 float32 = v6
    for fi15 := 0; fi15 < 19; fi15++ {
        v6 += (float32(float64(v14)) - 0.1)
        _ = fi15
    }
    v11 = v2
    for fi16 := 0; fi16 < 9; fi16++ {
        fmt.Println(v2, v10)
        _ = fi16
    }
    _ = len(v4)
    var v17 int32 = 1
    v18 := GenS1{F0: 1}
    for fi19 := 0; fi19 < 19; fi19++ {
        v3 = "hello"
        v14 -= v14
        _ = fi19
    }
    for fi20 := 0; fi20 < 6; fi20++ {
        v13 += 42
        v21 := GenS1{F0: 100}
        if v10 {
            for fi22 := 0; fi22 < 2; fi22++ {
                v2 -= (v11 >> ((18446744073709551615) & 63))
                v23 := GenS1{F0: v11}
                _ = fi22
                _ = v23
            }
            var v24 int16 = 1
            fmt.Println(v13, v24, v2)
            _ = v24
        }
        _ = fi20
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
