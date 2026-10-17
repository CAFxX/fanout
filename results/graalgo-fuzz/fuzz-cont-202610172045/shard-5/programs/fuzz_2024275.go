package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 uint64, p1 uint32, p2 int16) uint64 {
    if true {
        fmt.Println(p1)
        v6 := GenS1{F0: 127}
        p0 -= p0
        _ = v6
    } else {
        var v7 float32 = 3.14159
        _ = v7
    }
    p0 += (p0 << ((4294967295) & 63))
    if (true && false) {
        var v8 float64 = (((1.5 / 1.5) * 0.1) - 3.14159)
        _ = v8
    } else {
        if (false && false) {
            p2 = p2
            fmt.Println(p2, p1)
            v9 := GenS1{F0: 2147483647}
            _ = v9
        } else {
            p0 = 18446744073709551615
        }
        fmt.Println(p1, p1)
        p0 += p0
    }
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 bool = (float32((1.5 / float64(1e10))) >= (-2.25 * ((1e10 + 0.0) * 1e10)))
    var v2 uint64 = 0
    var v3 int = 469
    var v4 bool = v1
    var v5 float64 = 3.14159
    var v10 uint8 = 42
    v2 = v2
    var v11 string = "a"
    v12 := GenS1{F0: (v3 >> ((uint(v2)) & 63))}
    fmt.Println(v2, v4, v3)
    var v13 uint64 = 1
    for fi14 := 0; fi14 < 8; fi14++ {
        var v15 uint64 = (^(v13))
        v15 = v2
        v11 = v11
        _ = fi14
        _ = v15
    }
    var v16 uint8 = 1
    v1 = ((v4 && v4) && (genFn0((^(v13)), uint32(v16), 32767) < 18446744073709551615))
    var v17 uint16 = 624
    for fi18 := 0; fi18 < 8; fi18++ {
        var v19 float32 = 0.0
        for fi20 := 0; fi20 < 9; fi20++ {
            if (!(false || v1)) {
                v16 = v10
                v21 := GenS2{F0: -1}
                v22 := GenS3{F0: uint(v17), F1: fi20, F2: 91}
                _ = v21
                _ = v22
            }
            _ = fi20
        }
        fmt.Println(v16)
        _ = fi18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
