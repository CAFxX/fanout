package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8 }

func main() {
    var v1 int64 = 0
    v2 := make(map[int]bool)
    v3 := []uint32{2, 477, 572, 65535}
    var v4 int64 = -1
    fmt.Println(v4)
    v4 -= 0
    var v5 int16 = 0
    var v6 uint32 = 165
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    _ = len(v3)
    v1 = (v4 + 42)
    v2[-990] = ((!(float64(float32(float64(-2.25))) > ((1e10 * 0.0) / (float64(((3.14159 - 0.1) / -2.25)) * -2.25)))) && true)
    var v7 float32 = 0.1
    v8 := GenS1{F0: -1}
    v9 := GenS2{F0: 111}
    for fi10 := 0; fi10 < 2; fi10++ {
        v11 := GenS1{F0: 1}
        var v12 float64 = (1e10 * ((-2.25 + (1.5 - 3.14159)) * 3.14159))
        v6 += (v6 % ((2147483648) | 1))
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 uint8 = 0
    fmt.Println(v4, v5)
    var v14 uint32 = v6
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
