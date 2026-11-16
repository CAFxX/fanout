package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    var v1 uint = 2
    v2 := make(map[string]uint32)
    var v3 uint = 42
    for fi4 := 0; fi4 < 5; fi4++ {
        v3 = (v3 & v1)
        v1 = (^(v3))
        _ = fi4
    }
    if (42 <= 2147483647) {
        for fi5 := 0; fi5 < 16; fi5++ {
            v3 -= 0
            _ = fi5
        }
        v6 := GenS1{F0: 875}
        v1 -= v1
        _ = v6
    }
    var v7 uint8 = 250
    v7 = v7
    v8 := GenS1{F0: uint32(v7)}
    var v9 int16 = 1000
    var v10 int32 = 271
    var v11 uint16 = 2
    v12 := GenS1{F0: 889}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
