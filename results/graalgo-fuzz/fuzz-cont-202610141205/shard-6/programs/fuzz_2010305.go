package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 float32 = (0.0 * 0.0)
    var v2 bool = false
    var v3 float64 = (1e10 * (1.5 / float64((float32(float64(1e10)) / (float32(0.0) + v1)))))
    v4 := []string{"\n\t", "", ("\n\t" + ("foo bar" + "hello"))}
    var v5 uint16 = 357
    var v6 int16 = 42
    var v7 float64 = float64(float32(-2.25))
    switch 127 {
    case 896:
        for fi8 := 0; fi8 < 5; fi8++ {
            v9 := GenS1{F0: (-(v6))}
            _ = fi8
            _ = v9
        }
    case -186:
        v7 += (float64(v1) - ((v7 + (v3 + v7)) - v7))
    case -938:
        var v10 int16 = -32768
        _ = v10
    }
    fmt.Println(v6, v6)
    if ((float32((float64((v1 / v1)) - v7)) == v1) || ((289 != 146) && v2)) {
        fmt.Println(v2)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 = (^(v5))
    } else {
        v5 += v5
        v11 := GenS1{F0: (v6 - v6)}
        _ = v11
    }
    var v12 uint = 257
    var v13 string = "hello"
    switch v13 {
    case "foo bar":
        var v14 int16 = 0
        _ = v14
    case "":
        var v15 int8 = 0
        _ = v15
    }
    for fi16 := 0; fi16 < 3; fi16++ {
        fmt.Println(v6)
        var v17 int32 = -1
        v18 := GenS1{F0: v6}
        _ = fi16
        _ = v17
        _ = v18
    }
    v3 = (((v3 - ((v3 - 0.1) - v7)) - 0.1) / v3)
    v19 := GenS1{F0: (v6 + (v6 / ((v6) | 1)))}
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
