package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
func genFn0(p0 string, p1 uint16) uint32 {
    if true {
        v7 := []int{2147483647, 128}
        v7 = append(v7, 0)
        _ = v7
    }
    var v8 int8 = 120
    p0 = p0
    v9 := GenS1{F0: 255, F1: -941}
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 2147483648
}

func main() {
    var v1 bool = false
    var v2 string = ""
    var v3 int64 = 1
    var v4 float64 = 1.5
    v5 := []string{}
    var v6 bool = (0.1 > (-2.25 / 1e10))
    for fi10 := 0; fi10 < 14; fi10++ {
        v4 = v4
        fmt.Println(v4, v3, v6)
        _ = fi10
    }
    v2 = v2
    for fi11 := 0; fi11 < 7; fi11++ {
        v2 = v2
        _ = fi11
    }
    v6 = false
    fmt.Println(v3, v1)
    v12 := GenS1{F0: 42, F1: 4294967295}
    switch "a" {
    case "\n\t":
        var v13 int8 = 127
        _ = v13
    }
    v2 = v2
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 int = (fi14 / ((fi14) | 1))
        v2 = v2
        _ = fi14
        _ = v15
    }
    if (-1 > -2147483648) {
        v4 -= (v4 * (1.5 - (-2.25 / ((0.0 / (v4 * (3.14159 * (((v4 * (v4 * float64(1e10))) - 0.0) + v4)))) * -2.25))))
    } else {
        if false {
            if false {
                v16 := GenS1{F0: 65535, F1: -387}
                var v17 uint8 = 2
                var v18 uint32 = 928
                _ = v16
                _ = v17
                _ = v18
            }
            if (18446744073709551615 >= 2) {
                if (v6 || (v6 && v1)) {
                    if false {
                        v4 = 0.1
                    }
                    if len(v5) > 0 {
                        v5[0] = v2
                    }
                }
                v19 := GenS1{F0: 42, F1: 127}
                _ = v19
            }
        }
        v5 = append(v5, v2)
        var v20 int64 = (v3 * v3)
        _ = v20
    }
    var v21 float32 = 0.1
    switch v2 {
    case "a":
        fmt.Println(v4, v4)
    case "foo bar":
        var v22 uint16 = 65535
        _ = v22
    case "\n\t":
        v2 = (v2 + "\n\t")
    }
    v23 := GenS1{F0: 0, F1: 796}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
