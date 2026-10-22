package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 float32 = float32(1e10)
    var v2 int16 = 42
    var v3 uint16 = 1
    v4 := []uint32{42, 1, 2147483648, 65535}
    var v5 int32 = 1
    v1 = v1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v6 int8 = 100
    if (1 >= 93) {
        var v7 uint8 = 30
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = len(v4)
        _ = v7
    }
    fmt.Println(v1, v1, v3)
    v8 := GenS1{F0: false}
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 uint64 = 42
        fmt.Println(v10, v6, v6)
        v1 = (v1 + v1)
        _ = fi9
        _ = v10
    }
    fmt.Println(v2, v2, v1)
    v4 = append(v4, 0)
    for fi11 := 0; fi11 < 3; fi11++ {
        if false {
            if len(v4) > 0 {
                v4[0] = 2
            }
        }
        _ = fi11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        v13 := GenS1{F0: false}
        _ = fi12
        _ = v13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
