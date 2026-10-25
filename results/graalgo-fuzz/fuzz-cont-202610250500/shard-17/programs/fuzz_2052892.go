package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
func genFn0(p0 int16, p1 uint64) uint8 {
    fmt.Println(p1, p0)
    switch "\n\t" {
    case "a":
        var v7 float64 = 1e10
        _ = v7
    case "":
        for fi8 := 0; fi8 < 20; fi8++ {
            p1 += p1
            v9 := []int{}
            _ = fi8
            _ = v9
        }
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 int8 = -42
    v2 := []int{2147483647, 0}
    var v3 uint32 = 65535
    var v4 int = 127
    v5 := make(map[int]bool)
    v6 := make(map[int]uint32)
    var v10 int8 = 42
    switch "hello" {
    case "foo bar":
        v11 := v6[-1]
        _ = v11
    }
    var v12 int32 = 2147483647
    if len(v2) > 0 {
        v2[0] = v4
    }
    if (!(!true)) {
        delete(v6, (v4 << ((404) & 63)))
        v2 = append(v2, v4)
    } else {
        var v13 int16 = 0
        v5[(v4 & -192)] = true
        var v14 bool = ((v4 == v4) && true)
        _ = v13
        _ = v14
    }
    var v15 string = "\n\t"
    var v16 int32 = 42
    var v17 int16 = 449
    v15 = "\n\t"
    var v18 int64 = -402
    fmt.Println(v1, v1, v12)
    if ((v3 | (v3 | 183)) != (v3 << ((255) & 63))) {
        switch v4 {
        case -1:
            v19 := GenS1{F0: (^(v4)), F1: 255}
            _ = v19
        case 32767:
            if true {
                switch v4 {
                case 9223372036854775807:
                    for fi20 := 0; fi20 < 14; fi20++ {
                        v4 += 1
                        switch v15 {
                        case "hello":
                            if true {
                                var v21 int16 = -32768
                                _ = v21
                            }
                        case "":
                            v1 = v10
                        default:
                            fmt.Println(v18)
                        }
                        v4 = v4
                        _ = fi20
                    }
                default:
                    for fi22 := 0; fi22 < 18; fi22++ {
                        var v23 float64 = 3.14159
                        _ = fi22
                        _ = v23
                    }
                }
            }
        }
    }
    var v24 int = 128
    var v25 int8 = (v10 >> ((0) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
