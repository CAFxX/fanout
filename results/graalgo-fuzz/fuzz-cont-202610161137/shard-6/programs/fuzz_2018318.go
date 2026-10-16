package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    v1 := []uint32{2147483648, 2}
    v2 := make(map[int]uint32)
    var v3 uint16 = 896
    for fi4 := 0; fi4 < 3; fi4++ {
        for fi5 := 0; fi5 < 3; fi5++ {
            var v6 string = "foo bar"
            _ = fi5
            _ = v6
        }
        v3 += v3
        _ = fi4
    }
    v7 := v2[127]
    _ = len(v1)
    v8 := GenS1{F0: 2147483647}
    var v9 float64 = float64(((float32(float64((3.14159 + 3.14159))) / -2.25) - float32(0.0)))
    v9 = v9
    v9 -= float64(-2.25)
    v7 = (v7 ^ 65535)
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint16 = 0
        v11 -= 255
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 2147483647}
    if (false && ((v3 > v3) && false)) {
        var v13 int = 32768
        _ = v13
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
