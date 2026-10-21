package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int32 }

func main() {
    v1 := []int{42, 127}
    var v2 int8 = 0
    var v3 uint64 = 4294967295
    v4 := []string{}
    var v5 int8 = (v2 / ((v2) | 1))
    var v6 float64 = 0.0
    v3 = (v3 + (v3 / ((v3) | 1)))
    v2 -= -63
    v2 = v2
    for fi7 := 0; fi7 < 20; fi7++ {
        v5 -= -75
        _ = fi7
    }
    if true {
        for fi8 := 0; fi8 < 18; fi8++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v9 string = ""
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            v3 = v3
            fmt.Println(fi10)
            var v11 uint = 255
            _ = fi10
            _ = v11
        }
    } else {
        v2 -= v5
    }
    var v12 uint16 = 65
    switch "hello" {
    case "foo bar":
        v3 = v3
    default:
        v13 := GenS1{F0: v12, F1: -466, F2: -1}
        _ = v13
    }
    if (526 < 4294967295) {
        _ = len(v4)
    } else {
        switch -700 {
        case -2147483648:
            var v14 int = 32768
            _ = v14
        default:
            v12 = v12
        }
        v15 := GenS1{F0: (v12 | 1), F1: 953, F2: 1}
        v5 += int8(v12)
        _ = v15
    }
    fmt.Println(v6)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
