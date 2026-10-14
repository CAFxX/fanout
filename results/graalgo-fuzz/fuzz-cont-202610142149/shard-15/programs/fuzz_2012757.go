package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint8 }
func genFn0(p0 int8, p1 uint32, p2 int) int16 {
    var v7 int32 = -627
    switch "\n\t" {
    case "foo bar":
        var v8 uint64 = 42
        _ = v8
    }
    var v9 int32 = v7
    if true {
        switch ("hello" + ("foo bar" + ("\n\t" + "\n\t"))) {
        case "":
            for fi10 := 0; fi10 < 5; fi10++ {
                if false {
                    for fi11 := 0; fi11 < 20; fi11++ {
                        var v12 float32 = 1.5
                        var v13 bool = false
                        v7 += (-(v7))
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    p1 = p1
                }
                _ = fi10
            }
        default:
            p1 -= (p1 - p1)
        }
    }
    p0 -= p0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v9
    return -1
}

func main() {
    var v1 float64 = (0.1 + float64((1e10 - ((1.5 + 1e10) + 3.14159))))
    var v2 uint8 = 1
    v3 := make(map[int]uint32)
    v4 := make(map[uint32]string)
    var v5 float64 = v1
    var v6 int64 = 42
    v5 = v5
    v1 -= v5
    v2 += v2
    v14 := GenS1{F0: v2}
    v5 += 0.1
    v15 := []int{}
    fmt.Println(v5)
    var v16 int8 = 42
    _ = len(v15)
    v17 := GenS1{F0: v2}
    v18 := GenS2{F0: 989, F1: v2}
    if true {
        for fi19 := 0; fi19 < 6; fi19++ {
            for fi20 := 0; fi20 < 17; fi20++ {
                v6 += (v6 + (v6 >> ((666) & 63)))
                _ = fi20
            }
            v16 -= 100
            fmt.Println(v2, v2)
            _ = fi19
        }
        v6 -= (v6 ^ -523)
        var v21 int32 = 187
        _ = v21
    } else {
        for fi22 := 0; fi22 < 6; fi22++ {
            v6 += v6
            _ = fi22
        }
        var v23 int8 = -128
        if ((1 < 387) || (!true)) {
            fmt.Println(v5, v23, v5)
            for fi24 := 0; fi24 < 12; fi24++ {
                v1 = (v1 / (v1 + v5))
                _ = fi24
            }
        }
        _ = v23
    }
    var v25 uint = 898
    var v26 uint16 = 764
    v3[-2147483648] = 2147483648
    for fi27 := 0; fi27 < 15; fi27++ {
        v28 := GenS1{F0: uint8(v26)}
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
