package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int }
func genFn0() int64 {
    v4 := []int{-1, -39, 127}
    _ = len(v4)
    var v5 int = 781
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int8 = 42
    var v3 int32 = -160
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 string = (("foo bar" + "foo bar") + (("foo bar" + "foo bar") + ("hello" + "")))
        for fi8 := 0; fi8 < 2; fi8++ {
            var v9 int16 = 0
            _ = fi8
            _ = v9
        }
        _ = fi6
        _ = v7
    }
    var v10 uint16 = 427
    fmt.Println(v10, v1)
    var v11 int = 128
    v12 := GenS1{F0: 333}
    v10 = v10
    v13 := GenS2{F0: v11}
    v14 := []int{-2147483648}
    if (!false) {
        v3 -= v3
    } else {
        v3 -= v3
        var v15 float32 = (1e10 / 1e10)
        _ = v15
    }
    fmt.Println(v10, v1, v10)
    v16 := GenS1{F0: 4294967295}
    var v17 float32 = float32(-2.25)
    v1 = (v1 + (-(v1)))
    if false {
        if false {
            v10 -= (v10 * v10)
            if len(v14) > 0 {
                v14[0] = 9223372036854775807
            }
        }
        var v18 int32 = (^(v3))
        _ = v18
    }
    v2 = 0
    var v19 int16 = -701
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
}
