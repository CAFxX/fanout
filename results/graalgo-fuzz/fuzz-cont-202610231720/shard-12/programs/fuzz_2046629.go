package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint16 }
func genFn0(p0 string, p1 int16) int {
    var v5 float32 = 3.14159
    v6 := []int{}
    var v7 bool = (127 != 0)
    var v8 uint32 = 566
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 4294967295
}
func genFn1(p0 string) int32 {
    if false {
        fmt.Println(p0)
    }
    var v9 int = -2147483648
    v9 = (v9 ^ v9)
    _ = p0
    _ = v9
    return 0
}

func main() {
    v1 := []uint32{2147483648, 2147483648, 857, 1}
    var v2 bool = true
    v3 := []uint32{4294967295}
    var v4 uint32 = 4294967295
    v4 = 42
    if true {
        switch ((("" + "a") + ("\n\t" + "a")) + ("\n\t" + "a")) {
        case "foo bar":
            var v10 float64 = (0.1 - (1.5 / 1.5))
            _ = v10
        case "":
            for fi11 := 0; fi11 < 1; fi11++ {
                v12 := GenS1{F0: 4294967295}
                fmt.Println(v4, v4)
                v13 := GenS1{F0: 1}
                _ = fi11
                _ = v12
                _ = v13
            }
        case "\n\t":
            v4 = v4
        }
        v14 := GenS1{F0: 4294967295}
        switch "\n\t" {
        case "hello":
            fmt.Println(v2, v4)
        case "\n\t":
            v15 := GenS1{F0: 436}
            _ = v15
        case "foo bar":
            if ((!(v2 || v2)) && v2) {
                if v2 {
                    v4 -= (v4 & (v4 | v4))
                    fmt.Println(v4)
                    v16 := GenS1{F0: 1}
                    _ = v16
                } else {
                    v3 = append(v3, v4)
                }
            }
        default:
            v17 := GenS1{F0: 379}
            _ = v17
        }
        _ = v14
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        v1 = append(v1, v4)
        _ = fi18
    }
    var v19 uint32 = (v4 << ((2) & 63))
    for fi20 := 0; fi20 < 1; fi20++ {
        v4 -= (v4 - (v4 / (((v19 >> ((255) & 63))) | 1)))
        for fi21 := 0; fi21 < 6; fi21++ {
            v22 := GenS2{F0: 0, F1: 687}
            if len(v3) > 0 {
                v3[0] = (v4 % (((^(v19))) | 1))
            }
            _ = fi21
            _ = v22
        }
        switch "\n\t" {
        case "\n\t":
            for fi23 := 0; fi23 < 2; fi23++ {
                v19 = 0
                fmt.Println(v19, v4)
                _ = fi23
            }
        case "":
            var v24 float32 = (1.5 * ((0.0 / 3.14159) * ((3.14159 + 0.1) / 1.0)))
            _ = v24
        default:
            v19 = 0
        }
        _ = fi20
    }
    for fi25 := 0; fi25 < 20; fi25++ {
        v19 = (v4 & (v4 << ((2) & 63)))
        _ = fi25
    }
    if (v19 >= (^(v4))) {
        _ = len(v1)
        fmt.Println(v4)
    }
    v3 = append(v3, (v4 & 877))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
