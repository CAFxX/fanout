package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32; F1 bool }

func main() {
    var v1 uint = 0
    var v2 bool = false
    var v3 string = "a"
    var v4 uint = v1
    v5 := make(map[uint32]bool)
    v6 := GenS1{F0: 2, F1: 936}
    v3 = (v3 + ("a" + v3))
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 bool = v2
        var v9 float64 = 0.1
        v8 = (!(42 > -110))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        v11 := GenS2{F0: 65535, F1: (v1 == v4)}
        for fi12 := 0; fi12 < 3; fi12++ {
            v1 += 791
            var v13 uint = (v1 - v4)
            var v14 bool = (true && (v2 || v2))
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 int16 = 32767
        _ = fi10
        _ = v11
        _ = v15
    }
    var v16 float32 = 0.1
    v17 := GenS2{F0: 89, F1: true}
    var v18 int8 = 42
    fmt.Println(v2, v16)
    v19 := GenS2{F0: 101, F1: v2}
    v16 = v16
    v2 = true
    v2 = (!((!(1 <= 248)) && v2))
    for fi20 := 0; fi20 < 20; fi20++ {
        if ((v2 || false) || v2) {
            v2 = v2
        }
        var v21 uint8 = 108
        v22 := GenS1{F0: 0, F1: 25}
        _ = fi20
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 8; fi23++ {
        if (!false) {
            fmt.Println(v1, v1)
            for fi24 := 0; fi24 < 6; fi24++ {
                v25 := []int{(fi23 + 9223372036854775807), fi24}
                v26 := v5[0]
                _ = fi24
                _ = v25
                _ = v26
            }
            fmt.Println(fi23, v3, v18)
        }
        if ((0 < 42) && (!(fi23 < (fi23 << (((v4 << (((v1 & 4294967295)) & 63))) & 63))))) {
            v27 := []int{int(v4), 42}
            v2 = ((false && (!v2)) || (1000 <= 159))
            fmt.Println(v18)
            _ = v27
        }
        _ = fi23
    }
    fmt.Println(v1, v18)
    v16 -= ((float32((float64((v16 + float32(((1.5 + 3.14159) - (0.1 / 1e10))))) - -2.25)) * 1.5) / v16)
    switch v3 {
    case "":
        v1 = 42
    case "hello":
        var v28 uint32 = 386
        _ = v28
    }
    for fi29 := 0; fi29 < 19; fi29++ {
        fmt.Println(v16)
        fmt.Println(v3)
        v18 = 127
        _ = fi29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
