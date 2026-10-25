package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 uint32, p1 int8, p2 int8) uint32 {
    p1 -= (p1 >> ((95) & 63))
    p1 += int8(p0)
    var v4 uint32 = (^(p0))
    var v5 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return p0
}
func genFn1(p0 int8) int16 {
    v6 := GenS1{F0: -218, F1: 938}
    fmt.Println(p0)
    v7 := GenS2{F0: -7, F1: false}
    _ = p0
    _ = v6
    _ = v7
    return 32767
}

func main() {
    var v1 uint16 = 308
    v2 := make(map[int]string)
    var v3 bool = (!(33 != 42))
    v3 = ((("foo bar" + "\n\t") + "foo bar") < "hello")
    delete(v2, -1)
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 uint8 = 0
            v3 = true
            fmt.Println(v10, v1)
            _ = fi9
            _ = v10
        }
        v3 = (3.14159 != (-2.25 / ((((1.5 - ((0.1 + (1e10 - -2.25)) * 0.1)) / ((0.1 / (3.14159 + 0.1)) / 3.14159)) * 0.1) / 0.1)))
        _ = fi8
    }
    v1 = (v1 >> ((254) & 63))
    var v11 uint64 = 687
    v1 = (v1 + 1)
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := GenS2{F0: 48, F1: false}
        fmt.Println(v3)
        v14 := GenS2{F0: -1, F1: v3}
        _ = fi12
        _ = v13
        _ = v14
    }
    v1 = (v1 ^ (v1 % (((v1 & 0)) | 1)))
    if false {
        v11 += 18446744073709551615
        var v15 int32 = -33
        v16 := GenS1{F0: 0, F1: 128}
        _ = v15
        _ = v16
    } else {
        var v17 uint = 65535
        _ = v17
    }
    var v18 bool = ((!v3) || v3)
    var v19 int8 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
