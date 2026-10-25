package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 bool }
type GenS2 struct { F0 int64; F1 bool; F2 int32 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []uint32{822, 2, 2}
    var v3 int16 = 0
    var v4 int8 = 127
    var v5 uint32 = 553
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v6 := GenS1{F0: ((-2147483648 >= 0) && ((v3 * (v3 / ((1000) | 1))) != v3)), F1: 1, F2: (!true)}
    v2 = append(v2, (v5 >> ((255) & 63)))
    _ = len(v1)
    v7 := GenS1{F0: (((0.1 / 1e10) == 0.0) || (-70 == 127)), F1: 557, F2: (-71 == -2147483648)}
    v8 := GenS2{F0: 9223372036854775807, F1: (("\n\t" + ("a" + "\n\t")) < "hello"), F2: 0}
    v9 := GenS2{F0: 580, F1: (1 != 0), F2: 1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v10 := v1[(v5 + v5)]
    delete(v1, (v5 << ((4294967295) & 63)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
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
