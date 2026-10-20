package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 uint16; F2 uint }
func genFn0(p0 uint8) uint16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 18; fi5++ {
            switch -2147483648 {
            case -51:
                if true {
                    if (true && true) {
                        v6 := GenS1{F0: "foo bar"}
                        _ = v6
                    }
                    var v7 float64 = 0.1
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            v8 := GenS1{F0: "a"}
                            _ = v8
                        }
                    }
                    _ = v7
                } else {
                    var v9 float64 = ((float64((1e10 - float32(float64(0.1)))) * 1e10) + 1e10)
                    _ = v9
                }
            case 32768:
                p0 += (p0 * (^(p0)))
            case 124:
                p0 = 253
            }
            p0 = (p0 ^ (p0 << ((100) & 63)))
            fmt.Println(fi5, p0, p0)
            var v10 uint16 = 0
            var v11 uint8 = 255
            _ = fi5
            _ = v10
            _ = v11
        }
        v12 := []int{fi4, 31, fi4}
        p0 = 1
        _ = fi4
        _ = v12
    }
    p0 -= p0
    var v13 uint = 4294967295
    var v14 int = 4294967295
    _ = p0
    _ = v13
    _ = v14
    return 275
}

func main() {
    var v1 uint = 42
    var v2 string = "hello"
    v3 := make(map[uint32]int)
    for fi15 := 0; fi15 < 11; fi15++ {
        switch v2 {
        case "\n\t":
            var v16 uint64 = 42
            _ = v16
        default:
            var v17 int8 = 127
            _ = v17
        }
        _ = fi15
    }
    v1 = 714
    v18 := []int{4294967295}
    var v19 int = 725
    var v20 uint16 = 2
    var v21 float32 = 1e10
    v20 = 476
    if (42 >= 779) {
        var v22 int32 = 2147483647
        _ = v22
    } else {
        if true {
            v23 := GenS2{F0: (v2 + v2), F1: 42, F2: v1}
            _ = v23
        }
        var v24 uint16 = (^(v20))
        _ = v24
    }
    var v25 int64 = 0
    var v26 uint64 = 46
    v3[0] = (v19 >> ((v1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v25
    _ = v26
    _ = v3
}
