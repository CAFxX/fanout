package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint32 }
func genFn0() int8 {
    var v7 uint16 = 2
    if false {
        v7 -= (^(v7))
        var v8 int = 0
        v9 := []int{(v8 * 32767), (v8 << ((666) & 63)), -674}
        _ = v8
        _ = v9
    }
    _ = v7
    return 42
}

func main() {
    var v1 float32 = float32(0.0)
    var v2 float64 = ((1e10 / 1e10) * float64((v1 - v1)))
    var v3 string = (((("" + "hello") + "foo bar") + "a") + "hello")
    var v4 uint64 = 42
    v5 := make(map[string]uint32)
    v6 := []string{}
    v4 -= 845
    var v10 uint16 = 59
    var v11 uint8 = 8
    for fi12 := 0; fi12 < 12; fi12++ {
        var v13 string = v3
        v14 := GenS1{F0: (-163 > fi12), F1: ((530 == -295) || true)}
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 14; fi15++ {
        v2 -= float64((v1 / v1))
        var v16 string = "\n\t"
        _ = fi15
        _ = v16
    }
    var v17 int64 = 0
    v18 := GenS1{F0: true, F1: (!true)}
    if len(v6) > 0 {
        v6[0] = v3
    }
    v19 := GenS2{F0: 747}
    if ((2147483647 >= 649) || true) {
        fmt.Println(v17, v2)
        switch int(v11) {
        case -311:
            var v20 int32 = int32(v17)
            _ = v20
        case -214:
            var v21 int8 = -128
            _ = v21
        case 291:
            for fi22 := 0; fi22 < 1; fi22++ {
                fmt.Println(v17, v10)
                var v23 uint8 = v11
                _ = fi22
                _ = v23
            }
        }
        v6 = append(v6, v3)
    }
    fmt.Println(v1, v3, v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
