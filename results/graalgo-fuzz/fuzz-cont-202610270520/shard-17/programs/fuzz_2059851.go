package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int64; F1 int8; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []uint32{0, 2147483648}
    var v4 bool = false
    for fi5 := 0; fi5 < 1; fi5++ {
        v4 = false
        _ = fi5
    }
    fmt.Println(v4)
    if len(v3) > 0 {
        v3[0] = 0
    }
    _ = len(v1)
    v4 = (!(3.14159 < float64(1.5)))
    if v4 {
        v4 = (672 >= -9223372036854775808)
        var v6 bool = ((255 != 2) && (0 > 127))
        _ = v6
    } else {
        fmt.Println(v4)
        v7 := v1[("a" + "\n\t")]
        v8 := v2[42]
        _ = v7
        _ = v8
    }
    v4 = v4
    switch "hello" {
    case "":
        if (42 < 1) {
            v4 = (166 >= 42)
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 int8 = 100
                _ = fi9
                _ = v10
            }
            if v4 {
                v11 := GenS1{F0: -34, F1: 0}
                for fi12 := 0; fi12 < 19; fi12++ {
                    v13 := GenS1{F0: 100, F1: 0}
                    _ = fi12
                    _ = v13
                }
                _ = v11
            } else {
                v4 = true
            }
        } else {
            var v14 uint16 = 255
            if true {
                for fi15 := 0; fi15 < 17; fi15++ {
                    v16 := GenS2{F0: 42, F1: 42, F2: v14}
                    _ = fi15
                    _ = v16
                }
                delete(v1, "hello")
            }
            v4 = false
            _ = v14
        }
    case "a":
        fmt.Println(v4)
    case "hello":
        var v17 float32 = float32(0.0)
        _ = v17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
