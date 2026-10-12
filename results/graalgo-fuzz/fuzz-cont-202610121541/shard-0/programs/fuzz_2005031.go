package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := []uint32{2147483648, 263}
    var v2 float32 = (((0.0 + 1.5) / 1.0) * -2.25)
    v3 := []int{685, 2147483647, 127}
    var v4 uint32 = 468
    var v5 uint = 2
    var v6 int = 32767
    for fi7 := 0; fi7 < 18; fi7++ {
        var v8 float64 = 1e10
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        fmt.Println(v8)
        _ = fi7
        _ = v8
    }
    v6 -= v6
    v9 := GenS1{F0: 2}
    v4 = 71
    if (!false) {
        var v10 float64 = (0.0 * (1.5 / (1e10 / ((1.5 / 1.0) * ((0.0 / 1.5) + 3.14159)))))
        _ = v10
    }
    fmt.Println(v2, v4)
    fmt.Println(v4, v4)
    for fi11 := 0; fi11 < 14; fi11++ {
        var v12 int32 = -143
        _ = fi11
        _ = v12
    }
    var v13 uint8 = 155
    v3 = append(v3, (v6 * 128))
    for fi14 := 0; fi14 < 7; fi14++ {
        if (true && true) {
            v2 = v2
            for fi15 := 0; fi15 < 3; fi15++ {
                var v16 int32 = 0
                _ = fi15
                _ = v16
            }
        }
        v17 := GenS1{F0: 42}
        _ = fi14
        _ = v17
    }
    v2 += 1.5
    switch "hello" {
    case "a":
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 bool = false
            var v20 int32 = -1
            var v21 uint64 = 92
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    case "hello":
        v2 -= (1e10 + (v2 / v2))
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v2 = v2
        for fi23 := 0; fi23 < 13; fi23++ {
            fmt.Println(v4, v4, v13)
            _ = fi23
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi22
    }
    var v24 uint32 = v4
    if len(v3) > 0 {
        v3[0] = v6
    }
    var v25 uint64 = 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
