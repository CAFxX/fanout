package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 int64 = 0
    var v3 uint16 = 210
    var v4 int16 = 1
    if false {
        switch -1 {
        case 42:
            v3 = v3
        case -9223372036854775808:
            switch 128 {
            case 32767:
                v5 := GenS1{F0: -1, F1: 65535}
                _ = v5
            }
        case -1:
            for fi6 := 0; fi6 < 14; fi6++ {
                v7 := []int{fi6}
                fmt.Println(fi6)
                _ = fi6
                _ = v7
            }
        }
    } else {
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 int = fi8
            _ = fi8
            _ = v9
        }
        v2 += (v2 & v2)
        v3 = v3
    }
    var v10 uint16 = v3
    if (1.5 != 0.0) {
        v10 = v10
    }
    var v11 int8 = 127
    var v12 uint8 = 127
    var v13 float32 = 3.14159
    v2 = v2
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := []int{(fi14 << ((167) & 63))}
        var v16 uint32 = 65535
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
