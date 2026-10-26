package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 uint8 }
func genFn0(p0 uint64, p1 uint64, p2 uint16) string {
    fmt.Println(p0, p2, p0)
    var v6 bool = (255 > 184)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return "hello"
}

func main() {
    var v1 bool = (-42 == 0)
    var v2 bool = (true && false)
    var v3 string = "a"
    var v4 uint16 = 32767
    v5 := make(map[int]bool)
    fmt.Println(v1, v4)
    if (-2.25 <= 1e10) {
        var v7 bool = false
        for fi8 := 0; fi8 < 17; fi8++ {
            v9 := GenS1{F0: -560, F1: 42, F2: 120}
            v4 -= v4
            _ = fi8
            _ = v9
        }
        var v10 float32 = float32(3.14159)
        _ = v10
        _ = v7
    } else {
        v3 = ""
        var v11 uint64 = 0
        _ = v11
    }
    var v12 uint32 = uint32(v4)
    var v13 int32 = 1
    if ((v1 || (!v2)) || ((!((v12 << ((42) & 63)) != (^(v12)))) && true)) {
        v12 -= (v12 - v12)
        for fi14 := 0; fi14 < 3; fi14++ {
            for fi15 := 0; fi15 < 7; fi15++ {
                fmt.Println(v12, v4, v1)
                _ = fi15
            }
            for fi16 := 0; fi16 < 2; fi16++ {
                fmt.Println(fi16)
                v2 = v2
                _ = fi16
            }
            _ = fi14
        }
    }
    v13 = v13
    v12 -= (v12 ^ 0)
    if v1 {
        v17 := GenS1{F0: 0, F1: 42, F2: 42}
        var v18 bool = v1
        _ = v17
        _ = v18
    }
    var v19 float32 = 1e10
    for fi20 := 0; fi20 < 18; fi20++ {
        v12 = (v12 << ((1) & 63))
        _ = fi20
    }
    v21 := GenS2{F0: 172, F1: 1, F2: 0}
    var v22 bool = true
    v1 = (v22 || (0.1 != 1.5))
    delete(v5, 32768)
    v22 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
