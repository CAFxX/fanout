package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 uint32; F2 uint16 }

func main() {
    var v1 int8 = 116
    v2 := make(map[string]int)
    var v3 string = (("" + "hello") + ("\n\t" + "a"))
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    v6 := make(map[string]int)
    fmt.Println(v3)
    var v7 uint64 = 18446744073709551615
    v8 := GenS1{F0: 1}
    _ = len(v4)
    v9 := GenS2{F0: 42, F1: 1, F2: 255}
    v7 -= v7
    var v10 int = 9223372036854775807
    v1 += 42
    v10 = (v10 >> ((4294967295) & 63))
    for fi11 := 0; fi11 < 8; fi11++ {
        if true {
            if false {
                v12 := GenS1{F0: 353}
                _ = v12
            } else {
                if (((^(v7)) >= v7) && (v3 < v3)) {
                    fmt.Println(v1, v7)
                }
                fmt.Println(v10, v7, v3)
                v1 = -21
            }
        }
        v3 = v3
        var v13 uint64 = (v7 + (^(v7)))
        _ = fi11
        _ = v13
    }
    switch v3 {
    case "hello":
        v1 += -128
    case "a":
        if (float64(1.5) >= 0.1) {
            if (!false) {
                v10 += (v10 ^ v10)
            } else {
                if len(v5) > 0 {
                    v5[0] = "a"
                }
            }
            v14 := GenS1{F0: 2}
            _ = v14
        } else {
            v15 := GenS1{F0: 594}
            var v16 float32 = 1e10
            _ = v15
            _ = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
