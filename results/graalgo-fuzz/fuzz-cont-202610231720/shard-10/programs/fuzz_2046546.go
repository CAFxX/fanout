package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
func genFn0() int {
    var v5 int8 = 0
    var v6 int64 = 9223372036854775807
    _ = v5
    _ = v6
    return 32768
}

func main() {
    var v1 uint8 = 1
    var v2 float64 = (1e10 / 1.5)
    var v3 float32 = (1e10 - 1.5)
    v4 := make(map[int]bool)
    v3 = 1e10
    for fi7 := 0; fi7 < 12; fi7++ {
        _ = len(v4)
        var v8 string = ""
        _ = fi7
        _ = v8
    }
    var v9 int64 = 0
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 16; fi11++ {
            v4[(-(fi11))] = false
            v3 = (v3 + (v3 + 1e10))
            var v12 uint32 = 4294967295
            _ = fi11
            _ = v12
        }
        _ = fi10
    }
    if true {
        v13 := GenS1{F0: 255}
        for fi14 := 0; fi14 < 3; fi14++ {
            v3 = v3
            v2 = (float64((((1e10 + 1e10) + (v3 / float32(v2))) * v3)) * v2)
            var v15 string = "\n\t"
            _ = fi14
            _ = v15
        }
        _ = v13
    } else {
        v9 -= (v9 % ((v9) | 1))
        var v16 uint64 = 338
        if false {
            v1 -= v1
        }
        _ = v16
    }
    var v17 bool = (((v3 / float32((v2 - (1e10 * v2)))) - 1.5) == ((v3 - 0.1) - v3))
    var v18 float32 = float32(v2)
    var v19 float32 = float32(-2.25)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
