package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16, p1 string) int64 {
    fmt.Println(p0)
    switch -416 {
    case 32768:
        v7 := GenS1{F0: 71}
        _ = v7
    }
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            var v10 string = p1
            _ = fi9
            _ = v10
        }
        p1 = (p1 + (p1 + p1))
        _ = fi8
    }
    var v11 float32 = 0.0
    _ = p0
    _ = p1
    _ = v11
    return 0
}

func main() {
    var v1 string = ""
    v2 := []string{}
    var v3 float32 = float32(0.1)
    var v4 int16 = 1
    v5 := make(map[int]string)
    v6 := []uint32{}
    var v12 uint32 = 1
    v13 := v5[582]
    v5[9223372036854775807] = (v1 + v13)
    v3 = v3
    delete(v5, 127)
    for fi14 := 0; fi14 < 10; fi14++ {
        v5[fi14] = v13
        var v15 int8 = 42
        _ = fi14
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        switch (-(fi16)) {
        case -1:
            v17 := GenS1{F0: 131}
            _ = v17
        case 1:
            var v18 bool = false
            _ = v18
        case -50:
            var v19 uint = 65535
            _ = v19
        default:
            var v20 uint32 = v12
            _ = v20
        }
        v12 += (v12 * v12)
        _ = fi16
    }
    v21 := GenS2{F0: -9223372036854775808}
    v12 = (v12 | 2147483647)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
