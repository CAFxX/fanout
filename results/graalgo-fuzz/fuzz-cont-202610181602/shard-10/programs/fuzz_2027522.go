package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }

func main() {
    var v1 int64 = 0
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := []int{-629, 42, 32767, 32768}
    v1 = v1
    var v5 int32 = 1
    fmt.Println(v1, v5)
    if false {
        v6 := GenS1{F0: 18446744073709551615, F1: 0}
        _ = v6
    } else {
        v1 -= v1
        var v7 int = -702
        v1 = (v1 >> ((2) & 63))
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        v4 = append(v4, -1)
        _ = fi8
    }
    v9 := GenS1{F0: 18446744073709551615, F1: 2147483647}
    v1 = (v1 | 406)
    var v10 string = "\n\t"
    _ = len(v3)
    for fi11 := 0; fi11 < 11; fi11++ {
        v10 = (v10 + "a")
        v12 := GenS1{F0: 691, F1: 114}
        var v13 uint8 = 34
        _ = fi11
        _ = v12
        _ = v13
    }
    delete(v3, 127)
    if true {
        switch "" {
        case "a":
            if true {
                _ = len(v4)
                v4 = append(v4, -589)
                v5 += -1
            } else {
                var v14 uint32 = 684
                switch v10 {
                case "hello":
                    var v15 float32 = float32(3.14159)
                    _ = v15
                case "foo bar":
                    for fi16 := 0; fi16 < 14; fi16++ {
                        v17 := GenS1{F0: 4294967295, F1: v14}
                        _ = fi16
                        _ = v17
                    }
                default:
                    var v18 int16 = -301
                    _ = v18
                }
                _ = v14
            }
        }
        for fi19 := 0; fi19 < 17; fi19++ {
            var v20 uint = 18446744073709551615
            _ = fi19
            _ = v20
        }
        switch 32767 {
        case 0:
            v10 = v10
        case 32768:
            var v21 uint64 = 18446744073709551615
            _ = v21
        }
    } else {
        delete(v3, int(v1))
    }
    var v22 float64 = 1e10
    v23 := GenS1{F0: 18446744073709551615, F1: 995}
    switch 0 {
    case 350:
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 uint64 = 1
            _ = fi24
            _ = v25
        }
    default:
        v22 = 3.14159
    }
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 uint = 4294967295
        fmt.Println(v10, fi26, v22)
        v28 := v2[2147483647]
        _ = fi26
        _ = v27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
