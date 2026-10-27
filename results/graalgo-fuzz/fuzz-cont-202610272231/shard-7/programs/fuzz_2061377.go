package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint16) uint8 {
    p0 = 453
    p0 += (p0 ^ (p0 * p0))
    var v7 uint = 42
    switch 0 {
    case -1:
        if (!false) {
            p0 = p0
        }
    case 32767:
        v7 = 664
    case 2147483647:
        p0 += p0
    }
    v7 = v7
    _ = p0
    _ = v7
    return 1
}

func main() {
    var v1 bool = (-9223372036854775808 == 1)
    var v2 float32 = (0.0 / -2.25)
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 255
    var v5 bool = (true && false)
    v6 := []int{128}
    for fi8 := 0; fi8 < 17; fi8++ {
        v3[42] = 2
        _ = fi8
    }
    v9 := GenS1{F0: 127}
    v10 := v3[0]
    v11 := GenS1{F0: -107}
    if false {
        v3[65535] = 42
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        fmt.Println(v10, v4)
    } else {
        fmt.Println(v2, v5)
    }
    fmt.Println(v4)
    var v12 uint16 = 1
    for fi13 := 0; fi13 < 10; fi13++ {
        v5 = (v5 || v5)
        var v14 int16 = int16(v10)
        _ = fi13
        _ = v14
    }
    if v5 {
        var v15 uint8 = (v4 * 2)
        v16 := GenS1{F0: 100}
        switch -666 {
        case -9223372036854775808:
            var v17 uint64 = 4294967295
            _ = v17
        }
        _ = v15
        _ = v16
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        v19 := v3[v10]
        var v20 uint32 = 669
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
