package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    v1 := []uint32{42, 42}
    var v2 int16 = 884
    v3 := make(map[string]bool)
    var v4 uint64 = 2
    if (796 < -1) {
        v2 -= (v2 % (((v2 & int16(v4))) | 1))
        var v5 int8 = -21
        var v6 int8 = v5
        _ = v5
        _ = v6
    } else {
        var v7 float32 = (1.5 + (-2.25 * ((-2.25 + (1.5 / 1e10)) - 0.0)))
        var v8 int = -963
        _ = v7
        _ = v8
    }
    fmt.Println(v4, v2)
    for fi9 := 0; fi9 < 1; fi9++ {
        switch "hello" {
        case "hello":
            fmt.Println(fi9, fi9, fi9)
        default:
            fmt.Println(v2)
        }
        v4 += v4
        for fi10 := 0; fi10 < 19; fi10++ {
            if false {
                var v11 uint64 = 939
                var v12 int64 = 1
                _ = v11
                _ = v12
            }
            _ = fi10
        }
        _ = fi9
    }
    if false {
        if false {
            v4 = v4
        }
    }
    v13 := v3["\n\t"]
    _ = len(v1)
    var v14 int = -2147483648
    v15 := GenS1{F0: 912}
    var v16 bool = v13
    if (!false) {
        v13 = (false && false)
    } else {
        switch (-(v14)) {
        case 32767:
            fmt.Println(v16, v2)
        }
        fmt.Println(v2)
        for fi17 := 0; fi17 < 9; fi17++ {
            fmt.Println(v4, v4)
            _ = fi17
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
