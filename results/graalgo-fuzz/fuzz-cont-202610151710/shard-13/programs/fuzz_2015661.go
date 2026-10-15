package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int64 }
type GenS2 struct { F0 int16; F1 string; F2 uint }
func genFn0(p0 uint8, p1 int8, p2 uint16) uint64 {
    if true {
        v6 := []int{}
        _ = v6
    } else {
        v7 := GenS1{F0: false, F1: -9223372036854775808, F2: -546}
        _ = v7
    }
    var v8 uint64 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 42
    var v2 bool = (false && (true || (!true)))
    v3 := []int{-750, 128, -506}
    var v4 int8 = 127
    var v5 int8 = v4
    var v9 int16 = 536
    v9 = (v9 / ((v9) | 1))
    fmt.Println(v2)
    var v10 int8 = (^(v4))
    var v11 uint = 212
    v2 = v2
    v12 := GenS1{F0: v2, F1: -128, F2: -1}
    if true {
        var v13 string = ((("hello" + "foo bar") + "hello") + "hello")
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: (v1 > (v1 >> (((v11 + 42)) & 63))), F1: 617, F2: -1}
        _ = fi14
        _ = v15
    }
    switch 1 {
    case 707:
        v10 -= v5
    case 127:
        v5 = v5
    default:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v5, v5, v5)
        case "":
            v3 = append(v3, -1)
        case "a":
            var v16 float64 = ((float64((0.0 + (((-2.25 - -2.25) + 0.1) + 1.5))) - ((-2.25 + ((float64(1.5) + (1.5 + 1e10)) * 1.5)) / float64(-2.25))) * 1e10)
            _ = v16
        }
    }
    if v2 {
        v3 = append(v3, -103)
        for fi17 := 0; fi17 < 20; fi17++ {
            v3 = append(v3, 32768)
            fmt.Println(v10, v10)
            v3 = append(v3, fi17)
            _ = fi17
        }
    } else {
        fmt.Println(v10)
    }
    v18 := GenS1{F0: (974 > 1), F1: 9223372036854775807, F2: -1}
    v19 := GenS2{F0: (v9 << ((v11) & 63)), F1: ("a" + ("foo bar" + "hello")), F2: (v11 - (^(v11)))}
    var v20 int = 602
    v10 -= v10
    v2 = v2
    var v21 uint8 = 163
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
