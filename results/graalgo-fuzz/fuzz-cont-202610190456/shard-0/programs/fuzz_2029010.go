package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint }

func main() {
    v1 := []int{257, 0}
    var v2 int8 = 127
    v3 := []int{9223372036854775807}
    v4 := make(map[uint32]int)
    if len(v1) > 0 {
        v1[0] = 643
    }
    fmt.Println(v2)
    v2 -= (-(v2))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = 8
    var v5 string = "a"
    var v6 uint32 = 2
    v3 = append(v3, -1)
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 float32 = float32(float64((float32(1e10) / 1e10)))
        v5 = ""
        _ = fi7
        _ = v8
    }
    var v9 int32 = -340
    if true {
        v5 = ((("hello" + "a") + (("hello" + "foo bar") + v5)) + (v5 + "\n\t"))
        fmt.Println(v2, v9)
        for fi10 := 0; fi10 < 6; fi10++ {
            for fi11 := 0; fi11 < 3; fi11++ {
                v12 := GenS1{F0: 423}
                switch (fi10 | fi11) {
                case -31:
                    var v13 float64 = -2.25
                    _ = v13
                }
                _ = fi11
                _ = v12
            }
            _ = fi10
        }
    } else {
        v14 := GenS2{F0: 42, F1: 42}
        delete(v4, (v6 ^ v6))
        _ = v14
    }
    var v15 uint8 = 2
    switch -567 {
    case -1:
        switch v5 {
        case "foo bar":
            v2 -= (^(v2))
        }
    case 4294967295:
        v16 := GenS3{F0: 1, F1: 199}
        _ = v16
    default:
        v17 := GenS1{F0: -213}
        _ = v17
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
