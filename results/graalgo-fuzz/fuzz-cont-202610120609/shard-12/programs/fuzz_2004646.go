package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int16 }
func genFn0(p0 int16, p1 int32, p2 int) uint32 {
    var v7 bool = ("foo bar" < "hello")
    if v7 {
        p0 += p0
        v8 := []int{p2}
        _ = v8
    } else {
        if ("foo bar" < "\n\t") {
            v7 = v7
            fmt.Println(p2)
        }
        switch "hello" {
        case "a":
            fmt.Println(v7, p2)
        case "\n\t":
            v9 := []int{(p2 + (p2 << ((255) & 63)))}
            _ = v9
        default:
            fmt.Println(p1)
        }
        switch p2 {
        case -1:
            v10 := []int{(p2 << ((872) & 63)), -56, p2, p2}
            _ = v10
        case 9223372036854775807:
            p1 += p1
        }
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v7 = (v7 || v7)
        var v12 uint = 584
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}
func genFn1() bool {
    v13 := []int{}
    for fi14 := 0; fi14 < 4; fi14++ {
        var v15 int64 = -17
        fmt.Println(fi14)
        _ = fi14
        _ = v15
    }
    _ = v13
    return true
}

func main() {
    var v1 int16 = 32767
    var v2 bool = (!false)
    var v3 int8 = 100
    var v4 float64 = 1e10
    v5 := []string{"hello", "\n\t", "foo bar"}
    v6 := []uint32{}
    v2 = (-2.25 < (float64(((1.5 - 0.0) + 1e10)) - -2.25))
    switch "" {
    case "\n\t":
        v16 := GenS1{F0: 0, F1: 254, F2: (v1 >> ((826) & 63))}
        _ = v16
    case "hello":
        fmt.Println(v1, v1, v3)
    case "":
        v6 = append(v6, 2147483647)
    }
    v6 = append(v6, genFn0(v1, 1, -981))
    v5 = append(v5, "foo bar")
    fmt.Println(v2, v4, v4)
    fmt.Println(v1, v4, v2)
    var v17 string = ""
    v2 = true
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
