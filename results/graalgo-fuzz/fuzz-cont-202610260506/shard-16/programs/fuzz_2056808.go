package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint }

func main() {
    var v1 uint32 = 653
    v2 := []int{202, -537, -1, 32767}
    var v3 uint = 284
    var v4 uint64 = 4294967295
    v5 := GenS1{F0: v4}
    var v6 uint32 = (v1 * 0)
    v3 -= (v3 << ((v3) & 63))
    var v7 string = "foo bar"
    v8 := GenS1{F0: v4}
    var v9 int = 127
    var v10 float32 = -2.25
    v11 := GenS1{F0: v4}
    v12 := GenS2{F0: -41, F1: v1, F2: (v3 ^ v3)}
    v13 := GenS1{F0: v4}
    if true {
        fmt.Println(v7, v10)
        v9 = v9
        for fi14 := 0; fi14 < 1; fi14++ {
            var v15 int = fi14
            if (!(0 != (^(v3)))) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                fmt.Println(fi14)
            }
            var v16 uint16 = 2
            _ = fi14
            _ = v15
            _ = v16
        }
    } else {
        if (0 >= 42) {
            _ = len(v2)
            var v17 int64 = 42
            switch -387 {
            case 498:
                v3 -= 1
            default:
                var v18 bool = false
                _ = v18
            }
            _ = v17
        } else {
            v1 = 0
            for fi19 := 0; fi19 < 14; fi19++ {
                v3 += (v3 & 0)
                _ = fi19
            }
        }
        fmt.Println(v7, v10, v3)
        var v20 int64 = -210
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
