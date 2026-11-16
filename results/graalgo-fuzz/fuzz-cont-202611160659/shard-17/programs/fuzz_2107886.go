package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() bool {
    var v4 float32 = 0.1
    if false {
        if false {
            v5 := GenS1{F0: 1}
            _ = v5
        }
    }
    v4 += (((-2.25 / v4) / 1.0) * v4)
    if false {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint32 = 2147483648
            v8 := GenS1{F0: 2}
            _ = fi6
            _ = v7
            _ = v8
        }
        if (0 <= 65535) {
            v9 := GenS1{F0: 941}
            _ = v9
        }
        v10 := GenS1{F0: 1}
        _ = v10
    }
    _ = v4
    return (true && false)
}

func main() {
    var v1 uint64 = 905
    var v2 uint32 = 4294967295
    var v3 uint16 = 228
    if true {
        v11 := GenS1{F0: (^(v1))}
        _ = v11
    }
    var v12 float32 = float32(0.1)
    var v13 int32 = 1
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int = 2147483647
        v1 = v1
        _ = fi14
        _ = v15
    }
    var v16 float32 = 3.14159
    v3 -= v3
    var v17 uint = 644
    v16 = v16
    v18 := []int{-2147483648, int(v3)}
    var v19 int8 = -128
    v20 := GenS1{F0: uint64(v2)}
    v18 = append(v18, 127)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    v12 = v16
    v3 += (v3 << (((^(v17))) & 63))
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    var v21 uint8 = 2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
}
