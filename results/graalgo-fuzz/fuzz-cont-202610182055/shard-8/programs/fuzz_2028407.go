package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 string }

func main() {
    var v1 int8 = 127
    v2 := []int{-679, 4294967295, -1}
    var v3 float64 = -2.25
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 := GenS1{F0: -363, F1: 399, F2: 923}
    v5 := GenS2{F0: 32767, F1: 719, F2: "hello"}
    if true {
        v6 := GenS2{F0: 2, F1: 0, F2: ""}
        switch "a" {
        case "foo bar":
            v1 = v1
        }
        _ = v6
    } else {
        v2 = append(v2, -981)
        for fi7 := 0; fi7 < 14; fi7++ {
            var v8 uint64 = 182
            if len(v2) > 0 {
                v2[0] = fi7
            }
            var v9 uint8 = 255
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 uint32 = 4294967295
    fmt.Println(v10)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v10)
    v2 = append(v2, -2147483648)
    v3 = 0.0
    v11 := GenS2{F0: 65535, F1: v10, F2: "foo bar"}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
