package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint16; F1 uint8 }
func genFn0() uint {
    v4 := GenS1{F0: 0}
    v5 := GenS2{F0: true}
    var v6 uint32 = 65535
    switch "\n\t" {
    case "a":
        var v7 int8 = 0
        _ = v7
    default:
        v8 := GenS3{F0: 2, F1: 128}
        _ = v8
    }
    v9 := []int{127, 42, 117}
    _ = v4
    _ = v5
    _ = v6
    _ = v9
    return 704
}
func genFn1(p0 int16, p1 uint16, p2 uint8) uint8 {
    v10 := GenS4{F0: -176, F1: 0, F2: p2}
    for fi11 := 0; fi11 < 5; fi11++ {
        var v12 string = ""
        _ = fi11
        _ = v12
    }
    v13 := GenS5{F0: 65535, F1: (^(p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v13
    return p2
}

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    var v3 float32 = -2.25
    if (-1 < 42) {
        if true {
            switch 127 {
            case 42:
                v2 = v2
            }
        }
    } else {
        var v14 int32 = -1
        _ = v14
    }
    var v15 int32 = 710
    var v16 float64 = 1e10
    v2 = v2
    var v17 int = 4294967295
    var v18 int8 = -112
    var v19 uint16 = 208
    v19 -= v19
    fmt.Println(v3)
    var v20 uint = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
