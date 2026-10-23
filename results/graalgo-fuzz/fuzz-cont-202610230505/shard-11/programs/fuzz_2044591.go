package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint8 }

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 0.1
    v3 := []uint32{}
    var v4 string = "\n\t"
    v5 := GenS1{F0: 873, F1: 0, F2: 128}
    if true {
        if true {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                fmt.Println(v1, v1, v4)
                v2 += (float64(1.5) / 0.1)
            } else {
                fmt.Println(v1, v4)
                v4 = v4
                for fi6 := 0; fi6 < 7; fi6++ {
                    v4 = ((v4 + v4) + v4)
                    v7 := GenS1{F0: 65535, F1: 4294967295, F2: 1}
                    for fi8 := 0; fi8 < 6; fi8++ {
                        v9 := GenS1{F0: 824, F1: 42, F2: 42}
                        _ = fi8
                        _ = v9
                    }
                    _ = fi6
                    _ = v7
                }
            }
            fmt.Println(v2)
        }
        _ = len(v3)
        v2 -= float64(((0.0 + 0.1) - (0.1 / 0.1)))
    } else {
        for fi10 := 0; fi10 < 1; fi10++ {
            if (42 > 1000) {
                var v11 bool = true
                _ = v11
            } else {
                fmt.Println(v1)
                v4 = v4
            }
            var v12 int64 = 1
            _ = fi10
            _ = v12
        }
        fmt.Println(v2, v1, v4)
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: 255, F1: 485, F2: 1}
            fmt.Println(v1, fi13)
            v2 += 0.1
            _ = fi13
            _ = v14
        }
    }
    v1 = (-(v1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v3 = append(v3, 1)
    var v15 uint = 18446744073709551615
    _ = len(v3)
    var v16 uint = (v15 + 406)
    v17 := GenS1{F0: 2, F1: 4294967295, F2: 0}
    v4 = "hello"
    var v18 int = 4294967295
    v19 := GenS1{F0: 0, F1: 0, F2: 153}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
