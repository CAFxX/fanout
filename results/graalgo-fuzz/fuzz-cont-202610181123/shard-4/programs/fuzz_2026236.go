package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint16 = 32767
    var v2 int8 = 0
    var v3 uint8 = 126
    var v4 uint32 = 458
    v5 := make(map[uint32]bool)
    v6 := make(map[string]string)
    switch "a" {
    case "\n\t":
        v7 := []int{127}
        _ = v7
    }
    switch -693 {
    case 127:
        var v8 int16 = 1
        _ = v8
    case 32767:
        switch 2147483647 {
        case -286:
            var v9 uint = 0
            _ = v9
        case 520:
            var v10 int64 = 42
            _ = v10
        }
    default:
        var v11 bool = false
        _ = v11
    }
    if false {
        if true {
            switch -264 {
            case -2147483648:
                v12 := GenS1{F0: 810, F1: true}
                _ = v12
            case 146:
                for fi13 := 0; fi13 < 14; fi13++ {
                    for fi14 := 0; fi14 < 15; fi14++ {
                        var v15 int64 = -616
                        fmt.Println(v4, v3, v4)
                        _ = fi14
                        _ = v15
                    }
                    _ = fi13
                }
            }
        } else {
            _ = len(v5)
            v1 += (v1 / (((v1 / (((v1 / ((255) | 1))) | 1))) | 1))
            v2 = v2
        }
    }
    v1 -= v1
    if (!(0.0 <= float64(3.14159))) {
        v4 = v4
    } else {
        if (1 != int32(v1)) {
            if true {
                v4 += (^(v4))
            }
        } else {
            fmt.Println(v4)
            v3 += 254
        }
    }
    v16 := GenS1{F0: -1, F1: true}
    v17 := []int{32768, 32767, 785, 149}
    if ((9223372036854775807 < 42) || (false && false)) {
        v4 = (^(v4))
        fmt.Println(v1, v1, v2)
    } else {
        v1 += v1
        var v18 bool = false
        v1 += v1
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            var v21 int = (fi19 | 0)
            _ = fi20
            _ = v21
        }
        if false {
            fmt.Println(v1, v1)
            if false {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = len(v17)
                v6["a"] = "hello"
            }
        }
        _ = fi19
    }
    if (!(false && false)) {
        v22 := GenS2{F0: 0}
        v23 := v5[v4]
        _ = v22
        _ = v23
    } else {
        if len(v17) > 0 {
            v17[0] = 0
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
