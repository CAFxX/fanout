package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64) int {
    if (840 == 65535) {
        var v4 float64 = 3.14159
        v4 = float64((float32(0.0) / 3.14159))
        _ = v4
    }
    if (255 == 402) {
        var v5 float32 = (3.14159 - (0.1 / (1e10 / float32(3.14159))))
        v6 := []int{643, 0, 128}
        var v7 uint32 = 65535
        _ = v5
        _ = v6
        _ = v7
    }
    switch -2147483648 {
    case 1:
        p0 -= (p0 >> ((416) & 63))
    case -235:
        p0 = (p0 & p0)
    case 4294967295:
        p0 += 1
    }
    if true {
        v8 := []int{32767, -9223372036854775808, 0}
        _ = v8
    } else {
        p0 += (-(p0))
        var v9 uint = 416
        v10 := GenS1{F0: 1000}
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        fmt.Println(p0)
        _ = fi11
    }
    _ = p0
    return -9223372036854775808
}

func main() {
    var v1 uint32 = 415
    var v2 int16 = 42
    v3 := make(map[string]bool)
    var v12 uint16 = 452
    v13 := GenS1{F0: v2}
    v14 := GenS1{F0: v2}
    if (!false) {
        var v15 int64 = 596
        switch 128 {
        case 9223372036854775807:
            var v16 uint8 = 2
            _ = v16
        case 4294967295:
            var v17 uint32 = 0
            _ = v17
        default:
            v15 += v15
        }
        _ = v15
    } else {
        var v18 int32 = 721
        _ = v18
    }
    fmt.Println(v2, v12, v1)
    for fi19 := 0; fi19 < 19; fi19++ {
        var v20 uint8 = 42
        fmt.Println(v20)
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 1; fi21++ {
        var v22 uint8 = 0
        _ = fi21
        _ = v22
    }
    var v23 int = 42
    var v24 float64 = float64(0.0)
    v12 = uint16(v1)
    if false {
        delete(v3, "a")
    } else {
        v3[""] = ((255 <= (v12 % ((v12) | 1))) || ((3.14159 + 1e10) == 0.1))
        v25 := []int{}
        v23 -= -2147483648
        _ = v25
    }
    fmt.Println(v1, v23, v24)
    v24 += v24
    fmt.Println(v24, v1, v24)
    v26 := []int{(v23 | v23), -654, genFn0(42), 1}
    v27 := GenS2{F0: 819}
    var v28 bool = (-2147483648 > 196)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
