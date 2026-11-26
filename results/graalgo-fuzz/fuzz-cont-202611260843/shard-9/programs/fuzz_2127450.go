package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }

func main() {
    var v1 float64 = 0.1
    var v2 float64 = v1
    v3 := []int{-539, -1, -9223372036854775808}
    var v4 int = -820
    if true {
        v2 -= (v2 + 1e10)
        v1 = 1e10
        var v5 int8 = 1
        _ = v5
    } else {
        switch "foo bar" {
        case "\n\t":
            v6 := GenS1{F0: -2147483648, F1: 2}
            _ = v6
        case "":
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 uint8 = 42
                _ = fi7
                _ = v8
            }
        default:
            v4 -= v4
        }
        var v9 uint = 2
        _ = v9
    }
    var v10 bool = (true || (false || (!(2 > 66))))
    if v10 {
        v10 = (true || true)
    }
    _ = len(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        if v10 {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            switch "foo bar" {
            case "":
                v10 = v10
            case "hello":
                fmt.Println(fi11, v10, fi11)
            }
        }
        _ = fi11
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    v3 = append(v3, 128)
    var v12 int32 = 217
    var v13 int64 = -9223372036854775808
    v3 = append(v3, (-(v4)))
    var v14 uint32 = 221
    var v15 bool = (v2 == ((1.5 / v1) * v1))
    v13 -= (^(v13))
    v10 = (0 != -1)
    switch ("\n\t" + "") {
    case "\n\t":
        if (!false) {
            var v16 uint32 = 42
            _ = v16
        } else {
            var v17 uint64 = 42
            var v18 int64 = v13
            for fi19 := 0; fi19 < 16; fi19++ {
                v4 += fi19
                _ = fi19
            }
            _ = v17
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
