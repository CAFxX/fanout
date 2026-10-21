package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 int32, p1 string, p2 int64) uint64 {
    p2 = 1
    v6 := GenS1{F0: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float64 = (0.1 / (1.5 - (3.14159 / 1.0)))
    var v3 int16 = -1
    var v4 int32 = -2147483648
    var v5 int32 = (v4 | (v4 / (((v4 << ((311) & 63))) | 1)))
    for fi7 := 0; fi7 < 20; fi7++ {
        var v8 float32 = (0.1 + 0.1)
        switch (fi7 % ((fi7) | 1)) {
        case -139:
            if true {
                v3 = (v3 >> ((42) & 63))
            } else {
                if false {
                    if false {
                        _ = len(v1)
                    } else {
                        var v9 int = 0
                        _ = v9
                    }
                    var v10 uint = 255
                    _ = v10
                }
            }
        case 4294967295:
            v8 -= v8
        default:
            var v11 int = (^(fi7))
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    v12 := GenS1{F0: 42}
    v13 := GenS1{F0: 42}
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 string = (("hello" + "hello") + "")
        _ = fi14
        _ = v15
    }
    v16 := []int{42, 427, 999}
    v17 := v1[0]
    for fi18 := 0; fi18 < 1; fi18++ {
        _ = len(v16)
        _ = fi18
    }
    var v19 int = 2147483647
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
