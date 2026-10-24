package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string; F1 int16; F2 int8 }
func genFn0(p0 bool, p1 int16) bool {
    p0 = (((-2.25 >= (0.0 * 0.1)) || (p0 || (1.5 != 0.1))) && (-537 > -352))
    p0 = p0
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return p0
}

func main() {
    var v1 bool = true
    var v2 int32 = -1
    var v3 bool = v1
    v4 := GenS1{F0: 1000}
    if ((!(0 == 2147483647)) && ((2 <= 1) || (!((("foo bar" + "\n\t") < ("foo bar" + (("\n\t" + "hello") + ""))) && v1)))) {
        for fi5 := 0; fi5 < 7; fi5++ {
            fmt.Println(v3)
            _ = fi5
        }
        for fi6 := 0; fi6 < 5; fi6++ {
            if v3 {
                v7 := []int{32767, fi6}
                _ = v7
            }
            _ = fi6
        }
    }
    v8 := GenS1{F0: -915}
    for fi9 := 0; fi9 < 11; fi9++ {
        v10 := GenS1{F0: 1}
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: -32768}
    fmt.Println(v2)
    var v12 float64 = float64(3.14159)
    v12 += 0.0
    v3 = (1 == 84)
    var v13 bool = v3
    if v1 {
        var v14 uint64 = 192
        if (0 <= -816) {
            switch 32768 {
            case 1:
                fmt.Println(v2)
            case -202:
                v14 = v14
            }
            switch "" {
            case "":
                var v15 uint8 = 0
                _ = v15
            case "foo bar":
                v16 := GenS2{F0: (("hello" + "foo bar") + "foo bar"), F1: 1, F2: 41}
                _ = v16
            }
        }
        _ = v14
    }
    v17 := []int{0}
    switch "foo bar" {
    case "\n\t":
        if v13 {
            if v13 {
                v2 = (v2 * v2)
            }
        }
    case "a":
        fmt.Println(v2, v3, v1)
    }
    var v18 int16 = 32767
    v1 = false
    var v19 int64 = 0
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    v12 += -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
