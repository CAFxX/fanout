package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32; F2 int8 }

func main() {
    var v1 bool = false
    var v2 int32 = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 float64 = (1.5 - 1.5)
    v5 := make(map[int]int)
    v6 := []uint32{}
    var v7 bool = (!v1)
    v1 = (-1 >= -546)
    for fi8 := 0; fi8 < 9; fi8++ {
        fmt.Println(v1)
        if ((v7 && (164 >= -368)) || (184 == 2)) {
            if len(v6) > 0 {
                v6[0] = 42
            }
            if v1 {
                switch fi8 {
                case 0:
                    var v9 uint16 = 568
                    _ = v9
                case -916:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        v7 = v7
                        var v11 uint16 = 42
                        var v12 bool = (!false)
                        var v13 int = (fi10 % (((fi8 >> ((19) & 63))) | 1))
                        _ = fi10
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                case -203:
                    var v14 bool = v1
                    _ = v14
                default:
                    var v15 bool = v1
                    _ = v15
                }
            }
        } else {
            v7 = false
        }
        _ = fi8
    }
    for fi16 := 0; fi16 < 18; fi16++ {
        var v17 int8 = 42
        fmt.Println(v1)
        v18 := v5[fi16]
        _ = fi16
        _ = v17
        _ = v18
    }
    var v19 uint64 = 0
    if v1 {
        v7 = false
    }
    v2 = v2
    v2 = (v2 >> ((28) & 63))
    v7 = ((0 != 2) || (314 >= -32768))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 float64 = (v4 + (3.14159 + v4))
        v4 = v4
        _ = fi20
        _ = v21
    }
    v7 = (!(false && (1.5 > 1.5)))
    v22 := GenS1{F0: (v2 * v2), F1: 0, F2: -128}
    var v23 uint16 = 435
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
