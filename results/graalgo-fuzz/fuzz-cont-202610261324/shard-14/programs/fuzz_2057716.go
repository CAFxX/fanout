package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8; F1 uint64 }

func main() {
    var v1 uint32 = 116
    var v2 bool = true
    var v3 int64 = 742
    v4 := make(map[int]uint32)
    v5 := make(map[string]bool)
    switch "" {
    case "\n\t":
        var v6 bool = (!((!(((((((((1e10 / 3.14159) / 0.1) + (((3.14159 + (1e10 + (1e10 - ((((3.14159 - ((-2.25 * ((-2.25 - (1e10 - (1.5 / 3.14159))) / 1.5)) / (0.1 / (-2.25 / 1e10)))) * 3.14159) + 1.5) / float64(-2.25))))) / (-2.25 - 3.14159)) - 0.0)) * (-2.25 - -2.25)) * (3.14159 + (0.0 / -2.25))) / 0.1) + 0.0) + 3.14159) >= 1e10)) && (829 == 2)))
        _ = v6
    default:
        v1 = v1
    }
    switch ("foo bar" + "hello") {
    case "hello":
        switch -268 {
        case 1:
            if v2 {
                var v7 int16 = 32767
                v8 := GenS1{F0: true}
                v7 += (v7 + v7)
                _ = v7
                _ = v8
            } else {
                var v9 string = (("\n\t" + "foo bar") + "")
                v10 := GenS1{F0: (1 < 465)}
                var v11 int64 = v3
                _ = v10
                _ = v11
                _ = v9
            }
        }
    case "a":
        v12 := GenS2{F0: 127, F1: 1}
        _ = v12
    }
    if v2 {
        fmt.Println(v3, v3)
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        v14 := v5["a"]
        var v15 int32 = 880
        _ = fi13
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 float64 = (0.0 + ((1e10 - (3.14159 - 0.0)) * 1e10))
        _ = fi16
        _ = v17
    }
    var v18 float64 = 1.5
    v18 = (v18 / float64(3.14159))
    if v2 {
        v2 = (672 == 32767)
        v19 := GenS2{F0: 255, F1: 903}
        _ = v19
    } else {
        v1 = (v1 << ((uint(v1)) & 63))
    }
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
