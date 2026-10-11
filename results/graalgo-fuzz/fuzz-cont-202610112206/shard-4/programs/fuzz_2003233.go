package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 int32 }
type GenS2 struct { F0 int; F1 int }
func genFn0(p0 int64) uint16 {
    p0 = -1
    var v7 float32 = 1.5
    var v8 int16 = -32768
    switch 127 {
    case 128:
        var v9 uint16 = 847
        _ = v9
    case 0:
        for fi10 := 0; fi10 < 3; fi10++ {
            switch "foo bar" {
            case "foo bar":
                for fi11 := 0; fi11 < 9; fi11++ {
                    var v12 string = "a"
                    v13 := GenS1{F0: 245, F1: -1, F2: 9}
                    fmt.Println(fi11, v8, p0)
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case "hello":
                v14 := GenS1{F0: 42, F1: 0, F2: 1}
                _ = v14
            }
            v15 := []int{}
            switch (-(fi10)) {
            case 9223372036854775807:
                p0 -= (p0 / (((p0 ^ p0)) | 1))
            case 0:
                _ = len(v15)
            }
            _ = fi10
            _ = v15
        }
    case 367:
        v16 := GenS2{F0: -9223372036854775808, F1: -866}
        _ = v16
    }
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := []uint32{4294967295, 1, 2, 523}
    v4 := []string{"hello", ("\n\t" + "a")}
    var v5 float64 = 3.14159
    var v6 uint8 = 245
    v17 := v2[1]
    v1[1] = 2147483647
    if len(v4) > 0 {
        v4[0] = v17
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        var v19 int16 = -869
        _ = fi18
        _ = v19
    }
    if true {
        var v20 float64 = (float64(1.5) * v5)
        _ = v20
    }
    for fi21 := 0; fi21 < 14; fi21++ {
        var v22 float64 = (-2.25 / float64(1e10))
        var v23 int32 = 42
        v22 -= 3.14159
        _ = fi21
        _ = v22
        _ = v23
    }
    if true {
        v2[-2147483648] = "\n\t"
    }
    if (0 <= 9223372036854775807) {
        switch 861 {
        case 0:
            v24 := GenS1{F0: 65535, F1: 127, F2: 0}
            _ = v24
        default:
            var v25 uint32 = 2147483647
            _ = v25
        }
        if true {
            var v26 bool = true
            var v27 bool = (78 != 1)
            _ = v26
            _ = v27
        }
    }
    v28 := GenS1{F0: 18446744073709551615, F1: 60, F2: 0}
    var v29 float64 = float64(1.5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
