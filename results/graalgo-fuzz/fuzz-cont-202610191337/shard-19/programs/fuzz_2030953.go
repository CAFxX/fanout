package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        v8 := []int{fi7}
        v8 = append(v8, fi7)
        _ = fi7
        _ = v8
    }
    if (!false) {
        v9 := []int{115, 279, 2147483647, -604}
        var v10 int8 = -1
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int8 = 42
        _ = fi11
        _ = v12
    }
    return 4294967295
}
func genFn1() int64 {
    var v13 int8 = -128
    for fi14 := 0; fi14 < 14; fi14++ {
        fmt.Println(v13, v13)
        fmt.Println(fi14)
        v15 := GenS1{F0: (v13 >> ((0) & 63))}
        _ = fi14
        _ = v15
    }
    v13 = v13
    _ = v13
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = 1
    v3 := make(map[uint32]bool)
    var v4 int8 = -1
    var v5 string = "hello"
    v6 := make(map[uint32]uint32)
    v5 = v5
    var v16 int = 32768
    v16 -= (v16 / ((0) | 1))
    v17 := GenS1{F0: v4}
    fmt.Println(v5)
    v3[773] = (!(4294967295 == 42))
    if false {
        var v18 uint64 = 42
        v19 := GenS1{F0: v4}
        fmt.Println(v4, v5)
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := []int{(-(v16)), (-(v16))}
        _ = fi20
        _ = v21
    }
    var v22 float32 = -2.25
    for fi23 := 0; fi23 < 20; fi23++ {
        var v24 float64 = -2.25
        _ = fi23
        _ = v24
    }
    if (298 > 1) {
        switch v5 {
        case "a":
            for fi25 := 0; fi25 < 15; fi25++ {
                v5 = v5
                _ = fi25
            }
        case "hello":
            v26 := v6[42]
            _ = v26
        case "\n\t":
            v2 = 42
        }
        for fi27 := 0; fi27 < 8; fi27++ {
            fmt.Println(v22, v22)
            _ = fi27
        }
    }
    var v28 uint = 42
    var v29 uint32 = 4294967295
    var v30 float32 = v22
    v29 += 2
    var v31 string = "\n\t"
    v32 := GenS1{F0: v4}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
