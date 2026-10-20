package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32, p1 uint32, p2 string) int {
    if false {
        var v7 uint8 = 0
        _ = v7
    } else {
        var v8 string = p2
        if false {
            fmt.Println(v8, p0)
            var v9 int = 32768
            _ = v9
        }
        p2 = "hello"
        _ = v8
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint8 = 127
    var v2 int8 = 100
    var v3 uint64 = 726
    v4 := []uint32{65535, 42, 42, 258}
    v5 := make(map[string]bool)
    v6 := []uint32{1}
    v3 -= (v3 << ((42) & 63))
    fmt.Println(v2)
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    var v10 uint8 = (v1 << ((4294967295) & 63))
    switch "hello" {
    case "hello":
        for fi11 := 0; fi11 < 4; fi11++ {
            v2 = v2
            for fi12 := 0; fi12 < 2; fi12++ {
                _ = len(v6)
                v5["hello"] = ((9223372036854775807 < (fi11 | (fi12 ^ (-(fi12))))) && false)
                _ = fi12
            }
            _ = fi11
        }
    default:
        _ = len(v5)
    }
    v13 := GenS1{F0: 934}
    if true {
        fmt.Println(v3)
        var v14 uint32 = 42
        _ = v14
    } else {
        fmt.Println(v2, v3)
        var v15 string = "\n\t"
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int64 = 219
            var v18 float32 = 0.0
            v2 -= v2
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    fmt.Println(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
