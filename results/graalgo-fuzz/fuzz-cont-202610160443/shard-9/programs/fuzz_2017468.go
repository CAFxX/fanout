package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 bool }
func genFn0() bool {
    var v7 float64 = 0.0
    v7 = float64(0.0)
    _ = v7
    return (1 == 2)
}

func main() {
    var v1 uint = 2
    var v2 uint8 = 127
    var v3 uint64 = 2
    v4 := make(map[string]bool)
    var v5 int16 = 1
    v6 := []string{"\n\t", "a", ""}
    if false {
        for fi8 := 0; fi8 < 19; fi8++ {
            var v9 int16 = (v5 / ((v5) | 1))
            _ = fi8
            _ = v9
        }
    }
    _ = len(v4)
    v5 = -32768
    var v10 int64 = 1
    var v11 uint16 = uint16(v2)
    switch "a" {
    case "\n\t":
        v1 = 2
    default:
        var v12 float32 = 0.0
        _ = v12
    }
    v3 = 42
    v13 := GenS1{F0: v2, F1: 32767, F2: true}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v14 int = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
