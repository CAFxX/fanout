package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }

func main() {
    var v1 string = (("" + "hello") + "\n\t")
    v2 := make(map[uint32]bool)
    v3 := make(map[int]bool)
    v4 := make(map[uint32]uint32)
    v1 = v1
    for fi5 := 0; fi5 < 8; fi5++ {
        fmt.Println(v1, v1)
        _ = fi5
    }
    switch 586 {
    case -9223372036854775808:
        if false {
            v6 := GenS1{F0: 659}
            v7 := []int{41, 0, 4294967295, 2147483647}
            _ = v6
            _ = v7
        }
    default:
        v8 := GenS1{F0: 1}
        _ = v8
    }
    v9 := []int{127}
    for fi10 := 0; fi10 < 1; fi10++ {
        switch (v1 + v1) {
        case "a":
            var v11 float64 = (0.1 - 1.5)
            _ = v11
        default:
            var v12 int16 = 42
            _ = v12
        }
        for fi13 := 0; fi13 < 4; fi13++ {
            fmt.Println(fi10)
            var v14 float32 = (3.14159 + ((-2.25 * 1e10) - 0.0))
            var v15 int8 = 0
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi10
    }
    if (2 == 0) {
        v4[0] = 927
        fmt.Println(v1)
    }
    v1 = v1
    v16 := GenS2{F0: 2147483647}
    switch v1 {
    case "hello":
        v17 := v2[1]
        _ = v17
    }
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v1 = v1
        v1 = "\n\t"
        _ = fi18
    }
    v19 := v3[-9223372036854775808]
    var v20 uint32 = 2
    v20 = (v20 % ((689) | 1))
    var v21 bool = ((v20 / ((v20) | 1)) <= v20)
    if ((("a" + v1) != ("\n\t" + v1)) || v19) {
        var v22 uint32 = (v20 << ((2) & 63))
        _ = v22
    } else {
        var v23 string = v1
        v9 = append(v9, 32767)
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v9
}
