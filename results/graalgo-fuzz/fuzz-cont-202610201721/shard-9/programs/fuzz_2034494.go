package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
func genFn0() uint {
    var v7 float64 = 1.5
    for fi8 := 0; fi8 < 15; fi8++ {
        fmt.Println(v7, fi8)
        _ = fi8
    }
    fmt.Println(v7)
    v7 = v7
    _ = v7
    return 255
}

func main() {
    var v1 uint16 = 255
    var v2 int8 = 100
    var v3 float32 = 1.5
    var v4 uint8 = 0
    var v5 bool = (621 <= 1)
    v6 := make(map[uint32]int)
    var v9 int = 32767
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint16 = (v1 & (v1 - 32767))
        v11 = (v1 + v1)
        _ = fi10
        _ = v11
    }
    v12 := []int{v9}
    var v13 uint32 = 1
    var v14 uint32 = v13
    for fi15 := 0; fi15 < 18; fi15++ {
        v1 += (v1 | (v1 / (((^(v1))) | 1)))
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi15
    }
    delete(v6, (^(v13)))
    switch "" {
    case "hello":
        if true {
            var v16 float32 = float32(0.1)
            var v17 bool = v5
            _ = v16
            _ = v17
        }
    default:
        var v18 int = v9
        _ = v18
    }
    v2 -= v2
    switch "a" {
    case "\n\t":
        if (65535 > 2) {
            var v19 uint16 = (v1 + 0)
            v9 -= 9223372036854775807
            _ = v19
        } else {
            v20 := GenS1{F0: 1}
            var v21 float64 = (0.0 + (3.14159 - float64(v3)))
            _ = v20
            _ = v21
        }
    case "":
        for fi22 := 0; fi22 < 16; fi22++ {
            v5 = false
            var v23 int16 = -32768
            _ = fi22
            _ = v23
        }
    case "hello":
        var v24 string = "foo bar"
        _ = v24
    default:
        for fi25 := 0; fi25 < 18; fi25++ {
            v5 = true
            var v26 int16 = int16(v2)
            _ = fi25
            _ = v26
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
