package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool; F1 uint16; F2 bool }

func main() {
    var v1 uint = 255
    var v2 float64 = (-2.25 / -2.25)
    var v3 uint64 = 2
    if (((1.5 >= -2.25) && (1 > 47)) || (814 < 32768)) {
        var v4 int16 = -378
        _ = v4
    } else {
        v1 = v1
        for fi5 := 0; fi5 < 13; fi5++ {
            fmt.Println(v3, v1, v3)
            fmt.Println(v2)
            var v6 uint32 = 2
            _ = fi5
            _ = v6
        }
        v3 = 2
    }
    var v7 uint32 = 841
    v8 := GenS1{F0: 9223372036854775807}
    fmt.Println(v2)
    v9 := GenS2{F0: (true && (v3 == 0)), F1: 65535, F2: (!false)}
    switch "a" {
    case "":
        var v10 bool = ((true && false) && (true || true))
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := []int{2147483647, (fi11 + fi11)}
            v7 += (v7 - (v7 << ((v1) & 63)))
            _ = fi11
            _ = v12
        }
    default:
        v13 := []int{}
        _ = v13
    }
    var v14 uint32 = (v7 << ((uint(v7)) & 63))
    v7 = 65535
    var v15 float32 = -2.25
    v16 := []int{32767, 32768, 2147483647, -768}
    var v17 float64 = ((v2 / v2) / v2)
    var v18 float64 = -2.25
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
