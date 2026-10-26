package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }
func genFn0(p0 uint, p1 int16) uint {
    switch 127 {
    case 863:
        v5 := []int{}
        _ = v5
    }
    v6 := []int{32768, 0, 896, 924}
    var v7 int64 = 1
    var v8 int16 = p1
    p1 += p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint = 1
    var v2 int = 32767
    v3 := []int{(v2 | -745), v2}
    v4 := make(map[string]uint32)
    var v9 int16 = 42
    var v10 int16 = (v9 | 1)
    _ = len(v4)
    if false {
        delete(v4, "hello")
        v10 = (v10 * (v9 & v9))
        v11 := v4["hello"]
        _ = v11
    } else {
        var v12 int8 = -1
        v4["hello"] = 4294967295
        _ = v12
    }
    var v13 float32 = 0.1
    fmt.Println(v10)
    v14 := v4["\n\t"]
    for fi15 := 0; fi15 < 20; fi15++ {
        switch fi15 {
        case 32768:
            v16 := GenS1{F0: 2, F1: 4294967295, F2: 363}
            _ = v16
        case 180:
            var v17 uint32 = (v14 << ((v1) & 63))
            _ = v17
        case 32767:
            v13 = 3.14159
        }
        v1 -= 737
        _ = fi15
    }
    v13 += 0.1
    for fi18 := 0; fi18 < 3; fi18++ {
        if (823 > uint64(v14)) {
            v3 = append(v3, (v2 << ((1) & 63)))
            fmt.Println(v2)
            v19 := GenS1{F0: 1, F1: 2147483648, F2: v14}
            _ = v19
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
