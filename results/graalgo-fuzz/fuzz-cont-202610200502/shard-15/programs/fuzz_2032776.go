package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]int)
    v3 := make(map[string]string)
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    var v6 int = 32768
    var v7 uint8 = 128
    var v8 int16 = -740
    v7 = (v7 + v7)
    v9 := []int{-352, v6, (v6 >> ((65535) & 63)), (v6 >> ((65535) & 63))}
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v1, v6)
        _ = fi10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            var v13 uint32 = 2147483647
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 16; fi14++ {
            v1 = v1
            var v15 int16 = (v8 >> ((42) & 63))
            v7 -= v7
            _ = fi14
            _ = v15
        }
        _ = fi11
    }
    switch v6 {
    case 2147483647:
        if (0.1 > (0.0 - float64((v1 + v1)))) {
            v6 -= 4294967295
        } else {
            v1 = 1e10
        }
    case 32767:
        v16 := GenS1{F0: "hello"}
        _ = v16
    case 128:
        fmt.Println(v6)
    }
    v7 = (v7 << ((255) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
