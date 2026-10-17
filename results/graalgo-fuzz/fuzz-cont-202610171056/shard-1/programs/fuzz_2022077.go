package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint16) string {
    p0 = (^(p0))
    var v7 float64 = (-2.25 / 1.0)
    _ = p0
    _ = v7
    return "a"
}
func genFn1(p0 uint16) int8 {
    p0 += 32767
    if (false && true) {
        var v8 int32 = 42
        p0 = p0
        _ = v8
    }
    var v9 int16 = -504
    _ = p0
    _ = v9
    return 42
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint = 42
    v3 := []int{}
    v4 := make(map[uint32]bool)
    v5 := make(map[uint32]bool)
    var v6 float64 = 3.14159
    v10 := GenS1{F0: 2147483647}
    v11 := GenS1{F0: 701}
    v6 = (v6 / (v6 + (float64((((0.1 + (v1 / 1e10)) - 1.5) * float32((0.1 * v6)))) / ((v6 - float64((v1 - 1.5))) - (float64((((((-2.25 * (float32(v6) - (((float32((v6 + float64(float32(v6)))) + 0.0) * 1e10) - v1))) - float32(v6)) - v1) * float32(0.0)) - ((v1 / (float32(1.5) + v1)) + v1))) / v6)))))
    v2 += v2
    v3 = append(v3, 868)
    for fi12 := 0; fi12 < 10; fi12++ {
        var v13 int32 = -2147483648
        _ = fi12
        _ = v13
    }
    if len(v3) > 0 {
        v3[0] = -593
    }
    v4[42] = (!false)
    switch 9223372036854775807 {
    case -450:
        for fi14 := 0; fi14 < 6; fi14++ {
            if (0 != -1) {
                _ = len(v4)
                v6 = 0.1
            }
            if true {
                var v15 int32 = 2147483647
                fmt.Println(v15)
                fmt.Println(v2)
                _ = v15
            }
            if true {
                var v16 uint16 = 65535
                _ = v16
            }
            _ = fi14
        }
    case 651:
        fmt.Println(v2)
    case -90:
        var v17 int64 = -1
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 int = fi18
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
