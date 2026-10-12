package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[string]string)
    v3 := make(map[uint32]int)
    v4 := make(map[uint32]uint32)
    v5 := make(map[uint32]uint32)
    v1 = v1
    if (((!false) || false) || (955 <= 1)) {
        fmt.Println(v1)
        v6 := GenS1{F0: 175, F1: 32767}
        _ = v6
    } else {
        if false {
            var v7 bool = (3.14159 <= (1.5 + 1.5))
            _ = v7
        } else {
            switch -2147483648 {
            case -9223372036854775808:
                var v8 string = ""
                _ = v8
            case -512:
                for fi9 := 0; fi9 < 19; fi9++ {
                    if true {
                        v1 = -1
                        fmt.Println(fi9, v1)
                    }
                    _ = fi9
                }
            }
        }
    }
    var v10 int8 = -80
    v11 := GenS1{F0: 42, F1: 42}
    v12 := []int{-9223372036854775808, 4294967295, 408, 127}
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 127, F1: 41}
        if (0 != 1) {
            _ = len(v5)
            var v15 int32 = 1
            var v16 int32 = v15
            _ = v15
            _ = v16
        }
        var v17 uint16 = 970
        _ = fi13
        _ = v14
        _ = v17
    }
    fmt.Println(v10)
    switch "hello" {
    case "hello":
        v18 := v5[42]
        _ = v18
    default:
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := GenS1{F0: 31, F1: 819}
            var v21 float32 = -2.25
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v10 += -124
    for fi22 := 0; fi22 < 16; fi22++ {
        if false {
            var v23 uint16 = 65535
            _ = len(v4)
            _ = v23
        }
        _ = fi22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
