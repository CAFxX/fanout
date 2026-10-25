package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string; F1 bool }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[int]int)
    v3 := []int{127, 128}
    var v4 float32 = (1.5 / 1.5)
    var v5 uint16 = 1
    var v6 uint16 = v5
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v6, fi7, v6)
        _ = fi7
    }
    var v8 float32 = v4
    fmt.Println(v4)
    for fi9 := 0; fi9 < 3; fi9++ {
        var v10 int16 = 1
        for fi11 := 0; fi11 < 11; fi11++ {
            v10 = v10
            var v12 uint = 265
            var v13 string = ((("" + ("foo bar" + "a")) + "a") + "\n\t")
            _ = fi11
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
    }
    if (false || false) {
        v8 -= ((v8 - v8) + v8)
    }
    if false {
        if false {
            switch "a" {
            case "foo bar":
                if len(v3) > 0 {
                    v3[0] = 819
                }
            case "\n\t":
                v6 = (^(v5))
            }
        } else {
            fmt.Println(v5)
            var v14 int32 = -720
            _ = v14
        }
    } else {
        var v15 uint16 = (v6 & (v5 & (v6 | (v5 % ((285) | 1)))))
        _ = v15
    }
    _ = len(v2)
    v8 = v8
    var v16 uint32 = (v1 ^ (v1 >> ((255) & 63)))
    _ = len(v3)
    switch 682 {
    case -632:
        var v17 float32 = (float32(1.5) * 3.14159)
        _ = v17
    }
    if true {
        v16 -= v1
        v18 := GenS1{F0: ("\n\t" + "a")}
        _ = v18
    }
    switch -768 {
    case -864:
        v19 := GenS2{F0: ("\n\t" + "foo bar"), F1: (2 <= (^(v5)))}
        _ = v19
    default:
        v20 := GenS2{F0: "a", F1: (false || false)}
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
