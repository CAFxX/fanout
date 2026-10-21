package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
func genFn0(p0 uint64, p1 int8, p2 int32) uint8 {
    for fi6 := 0; fi6 < 2; fi6++ {
        for fi7 := 0; fi7 < 18; fi7++ {
            p1 = (^(p1))
            var v8 uint32 = 0
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    var v9 string = ("a" + ("hello" + "a"))
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 44
}

func main() {
    var v1 int16 = -32768
    var v2 bool = false
    var v3 int16 = v1
    var v4 uint8 = 2
    var v5 float32 = 1e10
    var v10 int = 901
    switch "" {
    case "foo bar":
        if true {
            v3 = v1
            v11 := GenS1{F0: (^(v4)), F1: 42}
            var v12 int64 = -974
            _ = v11
            _ = v12
        } else {
            v13 := GenS1{F0: v4, F1: 985}
            _ = v13
        }
    default:
        v14 := []int{360, (v10 >> ((18446744073709551615) & 63)), (v10 & (-(v10))), v10}
        _ = v14
    }
    v5 = v5
    fmt.Println(v2, v1, v5)
    var v15 uint8 = v4
    v16 := GenS1{F0: v15, F1: 425}
    if true {
        v17 := GenS1{F0: v15, F1: 255}
        _ = v17
    } else {
        var v18 float32 = float32(0.1)
        _ = v18
    }
    var v19 float64 = (0.0 - (3.14159 * 3.14159))
    v4 = genFn0(2, 42, 0)
    var v20 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
