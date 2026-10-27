package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
func genFn0() bool {
    var v4 float32 = 1.5
    v5 := GenS1{F0: 2, F1: 255}
    for fi6 := 0; fi6 < 15; fi6++ {
        v4 += (((v4 - v4) + v4) - v4)
        for fi7 := 0; fi7 < 5; fi7++ {
            v4 = 3.14159
            v4 -= (v4 / v4)
            v8 := []int{376, -495, 198, (^(fi7))}
            var v9 uint64 = 2
            v9 += v9
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    _ = v4
    _ = v5
    return ((!true) || false)
}

func main() {
    var v1 float32 = ((-2.25 / 1.0) - -2.25)
    var v2 int64 = 42
    v3 := []uint32{42, 1}
    var v10 bool = (!((true || (!false)) && false))
    if len(v3) > 0 {
        v3[0] = 652
    }
    switch ("hello" + "\n\t") {
    case "":
        var v11 uint32 = 2147483648
        _ = v11
    case "\n\t":
        if v10 {
            var v12 int = -1
            _ = len(v3)
            _ = v12
        }
    default:
        v10 = ((!v10) && v10)
    }
    var v13 int64 = 42
    var v14 uint64 = 42
    v2 = v2
    var v15 int16 = -1
    var v16 int32 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
}
