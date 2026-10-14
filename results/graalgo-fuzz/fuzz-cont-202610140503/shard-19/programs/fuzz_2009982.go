package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
func genFn0(p0 bool, p1 int32) bool {
    p1 = p1
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return (!false)
}
func genFn1(p0 uint32) uint32 {
    var v4 float64 = 3.14159
    v4 += v4
    fmt.Println(v4, v4)
    for fi5 := 0; fi5 < 14; fi5++ {
        var v6 uint = 1
        _ = fi5
        _ = v6
    }
    v4 = 1e10
    _ = p0
    _ = v4
    return p0
}

func main() {
    var v1 uint8 = 40
    var v2 uint32 = 65535
    v3 := make(map[int]int)
    v7 := GenS1{F0: 1, F1: 32767}
    v2 += (v2 >> ((255) & 63))
    switch 32767 {
    case 127:
        v1 = v1
    case 692:
        delete(v3, 9223372036854775807)
    case 780:
        if false {
            var v8 int8 = 42
            _ = v8
        }
    }
    var v9 int64 = 0
    v2 -= 65535
    v9 += v9
    v2 = 1
    for fi10 := 0; fi10 < 5; fi10++ {
        var v11 uint64 = 4294967295
        if true {
            v3[fi10] = (-(fi10))
        }
        _ = fi10
        _ = v11
    }
    var v12 uint16 = 425
    if (!false) {
        delete(v3, 32767)
        var v13 float32 = float32(float64((1.5 + (-2.25 - 0.0))))
        _ = v13
    } else {
        v14 := GenS1{F0: 624, F1: 919}
        v15 := GenS1{F0: 646, F1: -252}
        if false {
            var v16 int16 = 0
            v16 += v16
            _ = v16
        }
        _ = v14
        _ = v15
    }
    v3[4294967295] = 127
    var v17 int = 0
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
