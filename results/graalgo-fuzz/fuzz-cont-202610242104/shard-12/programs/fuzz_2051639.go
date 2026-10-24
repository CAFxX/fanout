package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }

func main() {
    var v1 string = "foo bar"
    v2 := make(map[string]uint32)
    v3 := []uint32{2147483647, 2147483648}
    v4 := make(map[string]uint32)
    var v5 int32 = 0
    v6 := make(map[uint32]string)
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: 2}
        v9 := GenS1{F0: 42}
        var v10 bool = false
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 uint8 = 217
    if (false || true) {
        for fi12 := 0; fi12 < 13; fi12++ {
            v1 = "foo bar"
            var v13 int32 = v5
            _ = fi12
            _ = v13
        }
    }
    fmt.Println(v11, v11, v11)
    var v14 int8 = -36
    v3 = append(v3, 4294967295)
    var v15 uint = 42
    fmt.Println(v11, v5, v11)
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 int16 = -528
        v11 = 2
        v18 := GenS1{F0: uint16(v15)}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
