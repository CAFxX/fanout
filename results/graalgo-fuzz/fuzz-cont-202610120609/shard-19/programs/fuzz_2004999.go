package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
func genFn0() int {
    if false {
        v5 := []int{95, -145, -2147483648}
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v5
    } else {
        var v6 uint32 = 42
        var v7 uint8 = 127
        v6 = 823
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 int = (fi8 % (((^(fi8))) | 1))
        _ = fi8
        _ = v9
    }
    return 128
}
func genFn1(p0 int, p1 int64, p2 uint8) uint8 {
    if (p0 != (p0 % ((p0) | 1))) {
        fmt.Println(p0)
        v10 := GenS1{F0: 42, F1: "a"}
        _ = v10
    }
    p2 = (p2 | 208)
    for fi11 := 0; fi11 < 9; fi11++ {
        p2 += (p2 + 1)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = p2
    return 128
}

func main() {
    v1 := []string{}
    var v2 bool = false
    var v3 int64 = 301
    var v4 uint32 = 2
    var v12 int8 = -109
    v4 += 885
    fmt.Println(v2, v2)
    _ = len(v1)
    v3 += (v3 ^ v3)
    if (255 == 1) {
        switch "a" {
        case "a":
            _ = len(v1)
        }
    } else {
        if len(v1) > 0 {
            v1[0] = ""
        }
        fmt.Println(v12)
        var v13 int16 = -983
        _ = v13
    }
    v12 -= (v12 * 121)
    for fi14 := 0; fi14 < 2; fi14++ {
        if v2 {
            _ = len(v1)
            v15 := GenS1{F0: 42, F1: "\n\t"}
            _ = v15
        }
        for fi16 := 0; fi16 < 9; fi16++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi16
        }
        if false {
            v3 = v3
            if v2 {
                v17 := GenS1{F0: 32767, F1: "a"}
                var v18 float64 = 0.1
                v18 += v18
                _ = v17
                _ = v18
            }
        }
        _ = fi14
    }
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 int64 = 42
        v1 = append(v1, "")
        _ = fi19
        _ = v20
    }
    var v21 uint8 = 229
    v12 = (v12 << ((925) & 63))
    var v22 uint8 = 255
    v2 = v2
    var v23 int8 = v12
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
