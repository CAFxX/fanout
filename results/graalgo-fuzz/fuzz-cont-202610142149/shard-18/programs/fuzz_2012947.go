package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 3.14159
    var v3 int = 0
    var v4 bool = true
    var v5 int16 = -32768
    v6 := GenS1{F0: 238, F1: -1}
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 int = (v3 << ((69) & 63))
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int32 = -2147483648
        var v11 uint16 = 2
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 uint = 34
    v4 = false
    v3 = v3
    v5 = int16(v12)
    v13 := GenS2{F0: (-968 >= 42), F1: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
