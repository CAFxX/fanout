package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
type GenS2 struct { F0 int32; F1 uint64 }
func genFn0(p0 int64, p1 int, p2 uint32) int64 {
    var v7 string = ("" + "\n\t")
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p0 | -9223372036854775808)
}

func main() {
    v1 := make(map[string]int)
    var v2 float64 = 3.14159
    v3 := make(map[int]uint32)
    var v4 string = "\n\t"
    v5 := make(map[uint32]int)
    var v6 uint = 1
    v8 := GenS1{F0: (v6 - v6), F1: (^(v6)), F2: v4}
    var v9 int16 = 264
    v9 += (v9 % ((-32768) | 1))
    switch 42 {
    case 2147483647:
        if true {
            var v10 uint32 = 65535
            for fi11 := 0; fi11 < 20; fi11++ {
                v12 := GenS1{F0: v6, F1: 2, F2: v4}
                _ = fi11
                _ = v12
            }
            var v13 uint = (v6 * 4294967295)
            _ = v10
            _ = v13
        } else {
            if false {
                var v14 uint64 = 42
                for fi15 := 0; fi15 < 16; fi15++ {
                    v16 := GenS1{F0: 2, F1: (v6 - (v6 ^ 0)), F2: "foo bar"}
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            delete(v3, 762)
        }
    default:
        var v17 float64 = v2
        _ = v17
    }
    for fi18 := 0; fi18 < 12; fi18++ {
        var v19 int8 = 118
        var v20 int8 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    fmt.Println(v2, v9)
    if ((true || ((!true) && false)) && true) {
        v21 := []int{-341, -9223372036854775808, -337}
        _ = v21
    }
    if true {
        v6 += (v6 << (((v6 * (v6 << ((v6) & 63)))) & 63))
    }
    if (!false) {
        _ = len(v1)
        switch v4 {
        case "foo bar":
            for fi22 := 0; fi22 < 7; fi22++ {
                var v23 float32 = 1e10
                _ = fi22
                _ = v23
            }
        default:
            var v24 uint32 = 2147483648
            _ = v24
        }
        v25 := GenS1{F0: 42, F1: (v6 | v6), F2: ((v4 + v4) + (v4 + v4))}
        _ = v25
    } else {
        v26 := GenS2{F0: 224, F1: 939}
        _ = v26
    }
    fmt.Println(v2, v9)
    v27 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
