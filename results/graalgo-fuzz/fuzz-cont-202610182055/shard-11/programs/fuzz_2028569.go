package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }

func main() {
    var v1 int16 = 1000
    v2 := []string{"\n\t", "\n\t", "foo bar"}
    var v3 uint32 = 661
    var v4 string = "a"
    fmt.Println(v1)
    if false {
        fmt.Println(v3)
    } else {
        fmt.Println(v3, v4)
        var v5 float32 = 1e10
        _ = v5
    }
    v4 = "hello"
    for fi6 := 0; fi6 < 19; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            if len(v2) > 0 {
                v2[0] = v4
            }
            var v8 int = (fi6 / (((fi7 >> ((4294967295) & 63))) | 1))
            var v9 uint = 1
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi6
    }
    var v10 float32 = 3.14159
    v11 := GenS1{F0: ((v10 != -2.25) && (2 >= 96)), F1: false}
    fmt.Println(v4)
    v3 = v3
    v12 := GenS1{F0: (false || (false && true)), F1: (false || (-1 <= 127))}
    v4 = ("\n\t" + (v4 + ((v4 + v4) + "foo bar")))
    var v13 int16 = 32767
    fmt.Println(v4)
    _ = len(v2)
    if (!((46 < v3) && (!(42 > 2)))) {
        var v14 int16 = (-(v13))
        v15 := GenS1{F0: false, F1: false}
        v13 -= 723
        _ = v14
        _ = v15
    }
    v3 = v3
    for fi16 := 0; fi16 < 14; fi16++ {
        var v17 int64 = -1
        var v18 bool = false
        var v19 uint16 = 993
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    var v20 int64 = -194
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
