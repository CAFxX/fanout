package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int; F2 int8 }
type GenS2 struct { F0 bool }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 65535
    var v3 int8 = 65
    v4 := []int{}
    var v5 int16 = 42
    v6 := GenS1{F0: v3, F1: 425, F2: v3}
    v3 = (-(v3))
    var v7 int32 = 42
    v3 = v3
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int64 = 0
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v11 uint16 = 524
            _ = fi9
            _ = v10
            _ = v11
        }
        _ = fi8
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        v3 = 1
        _ = fi12
    }
    v5 = 586
    var v13 uint = (v2 + 985)
    switch "hello" {
    case "foo bar":
        for fi14 := 0; fi14 < 8; fi14++ {
            if false {
                var v15 int8 = (v3 >> ((42) & 63))
                var v16 float32 = 0.0
                _ = v15
                _ = v16
            } else {
                if ((v5 - (-(v5))) > v5) {
                    if len(v4) > 0 {
                        v4[0] = (fi14 ^ -853)
                    }
                    var v17 bool = true
                    var v18 float64 = 3.14159
                    _ = v17
                    _ = v18
                }
                v13 = (^(v2))
                var v19 int = fi14
                _ = v19
            }
            v20 := GenS2{F0: false}
            v13 += 42
            _ = fi14
            _ = v20
        }
    case "":
        fmt.Println(v2, v7, v5)
    }
    switch "hello" {
    case "\n\t":
        var v21 float32 = 1.5
        _ = v21
    default:
        v2 = v2
    }
    v4 = append(v4, 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
