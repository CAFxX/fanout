package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
func genFn0(p0 int16, p1 bool) uint16 {
    var v6 int16 = (p0 ^ p0)
    switch ("\n\t" + "foo bar") {
    case "":
        var v7 uint32 = 938
        _ = v7
    case "\n\t":
        v6 = -764
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p1, p0)
        }
        var v9 int8 = 0
        var v10 uint16 = 0
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        p0 = p0
        fmt.Println(p1, v6, fi11)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v6
    return 1
}

func main() {
    var v1 float64 = float64(float32(1.5))
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    v4 := []uint32{0, 2, 2, 42}
    v5 := []uint32{0, 2}
    switch "foo bar" {
    case "foo bar":
        var v12 float32 = 0.0
        _ = v12
    default:
        for fi13 := 0; fi13 < 13; fi13++ {
            for fi14 := 0; fi14 < 9; fi14++ {
                v15 := GenS1{F0: 42, F1: 4}
                var v16 int64 = 42
                _ = fi14
                _ = v15
                _ = v16
            }
            var v17 float64 = float64(1e10)
            _ = fi13
            _ = v17
        }
    }
    v1 = v1
    for fi18 := 0; fi18 < 16; fi18++ {
        if ((float64((3.14159 * (0.0 - 1e10))) + v1) != float64(3.14159)) {
            v1 = v1
        }
        _ = fi18
    }
    if (!(!true)) {
        v1 += v1
        v1 = ((-2.25 / v1) * (0.0 / (v1 - 1.5)))
        var v19 int = 127
        _ = v19
    } else {
        for fi20 := 0; fi20 < 10; fi20++ {
            v21 := GenS1{F0: 4294967295, F1: 222}
            _ = fi20
            _ = v21
        }
    }
    var v22 bool = true
    for fi23 := 0; fi23 < 5; fi23++ {
        if (4294967295 > 225) {
            if len(v5) > 0 {
                v5[0] = 387
            }
        }
        _ = fi23
    }
    delete(v3, 42)
    var v24 float64 = v1
    if (v22 || v22) {
        var v25 uint8 = 127
        _ = v25
    }
    v26 := GenS1{F0: 872, F1: 31}
    var v27 uint16 = 615
    var v28 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v24
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
