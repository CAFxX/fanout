package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 uint64 }
func genFn0(p0 uint16) bool {
    fmt.Println(p0)
    var v4 uint64 = 18446744073709551615
    _ = p0
    _ = v4
    return (!(!(464 != 42)))
}

func main() {
    var v1 uint8 = 179
    v2 := []string{"hello", "a", "a", ""}
    v3 := []uint32{2147483648, 808, 2, 2}
    var v5 uint8 = 145
    if true {
        var v6 uint16 = 0
        _ = v6
    }
    fmt.Println(v5, v5)
    for fi7 := 0; fi7 < 3; fi7++ {
        v1 = 1
        _ = fi7
    }
    v5 += (v5 * 255)
    v8 := GenS1{F0: 818, F1: 172, F2: 18446744073709551615}
    var v9 uint8 = v1
    v9 = 190
    var v10 uint32 = 221
    var v11 float32 = (1e10 - 3.14159)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v12 uint32 = (v10 << ((65535) & 63))
    var v13 uint64 = 347
    var v14 float64 = ((float64(-2.25) / 1e10) - 0.0)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v5
    _ = v8
    _ = v9
}
