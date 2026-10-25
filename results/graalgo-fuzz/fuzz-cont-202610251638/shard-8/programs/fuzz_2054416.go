package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 bool }
func genFn0(p0 int32, p1 string) uint {
    var v7 float64 = 1.5
    fmt.Println(p0)
    fmt.Println(p0, p1, p0)
    _ = p0
    _ = p1
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 uint16 = 339
    v2 := make(map[uint32]int)
    var v3 uint64 = 287
    var v4 uint64 = v3
    v5 := make(map[string]bool)
    v6 := []string{"hello", "hello", ("\n\t" + "a"), "hello"}
    if ((-2.25 / 0.1) != float32(3.14159)) {
        v8 := GenS1{F0: 746, F1: 42, F2: false}
        _ = v8
    }
    fmt.Println(v3, v1)
    var v9 float32 = (float32(float64(1e10)) + 3.14159)
    var v10 float64 = ((-2.25 - 0.0) - 3.14159)
    switch -1 {
    case -2147483648:
        fmt.Println(v4, v10)
    case 4294967295:
        _ = len(v2)
    default:
        var v11 string = ""
        _ = v11
    }
    if (!((721 > 514) || false)) {
        switch -1 {
        case -310:
            fmt.Println(v9, v4)
        case -9223372036854775808:
            v12 := GenS1{F0: 790, F1: -1, F2: true}
            _ = v12
        }
        v10 += 0.0
    }
    if (418 > 766) {
        fmt.Println(v10)
        delete(v5, "a")
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 uint = 18446744073709551615
            v15 := GenS1{F0: 517, F1: 127, F2: false}
            v4 = 4294967295
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    v9 = float32(1e10)
    v16 := v2[0]
    var v17 int32 = 2147483647
    if (!(!true)) {
        var v18 int16 = -252
        _ = v18
    }
    v2[uint32(v3)] = v16
    var v19 int64 = 657
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
