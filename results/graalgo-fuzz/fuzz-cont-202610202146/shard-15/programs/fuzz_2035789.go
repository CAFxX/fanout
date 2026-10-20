package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }

func main() {
    var v1 uint16 = 32767
    v2 := []uint32{2, 2, 65535}
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]string)
    v5 := []int{2147483647, -394, -1, 0}
    var v6 uint32 = 42
    var v7 uint16 = 2
    switch ("\n\t" + "hello") {
    case "":
        var v8 uint = 1
        _ = v8
    default:
        var v9 uint32 = (v6 * v6)
        _ = v9
    }
    if (true && false) {
        var v10 float64 = -2.25
        fmt.Println(v6)
        _ = v10
    } else {
        var v11 uint64 = 0
        v11 = v11
        _ = v11
    }
    fmt.Println(v6, v6, v6)
    var v12 int = -1
    var v13 uint = 4294967295
    for fi14 := 0; fi14 < 14; fi14++ {
        var v15 int = v12
        for fi16 := 0; fi16 < 1; fi16++ {
            switch 127 {
            case 4294967295:
                fmt.Println(v6, v12, fi16)
            case 42:
                v3[v6] = v6
            case -856:
                var v17 uint16 = v7
                _ = v17
            }
            switch v12 {
            case -9223372036854775808:
                v18 := GenS1{F0: 1000, F1: v13}
                _ = v18
            }
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
