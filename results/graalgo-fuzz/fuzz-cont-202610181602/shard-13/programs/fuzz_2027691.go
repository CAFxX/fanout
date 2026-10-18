package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 bool }
type GenS2 struct { F0 string; F1 string; F2 uint8 }

func main() {
    var v1 int16 = 1000
    var v2 bool = true
    v3 := []int{}
    var v4 bool = (!true)
    var v5 uint16 = 0
    v6 := make(map[int]bool)
    var v7 float64 = ((1e10 * float64(((3.14159 - 0.0) + 0.0))) - 3.14159)
    var v8 int8 = 100
    v9 := GenS1{F0: v5, F1: v5, F2: (0 <= 42)}
    switch (("" + (((("a" + "a") + "foo bar") + ("\n\t" + ("" + "a"))) + "\n\t")) + "\n\t") {
    case "a":
        v10 := GenS2{F0: ("" + "\n\t"), F1: "hello", F2: 86}
        _ = v10
    case "foo bar":
        v11 := GenS2{F0: "a", F1: (("\n\t" + "\n\t") + ""), F2: 42}
        _ = v11
    case "\n\t":
        switch -1 {
        case 42:
            _ = len(v3)
        }
    default:
        var v12 int8 = v8
        _ = v12
    }
    switch 32768 {
    case 932:
        v13 := GenS2{F0: "", F1: "foo bar", F2: 255}
        _ = v13
    case -2147483648:
        fmt.Println(v2, v8)
    }
    v2 = v4
    v8 += (v8 - (v8 * (v8 % ((v8) | 1))))
    switch 967 {
    case 9223372036854775807:
        for fi14 := 0; fi14 < 17; fi14++ {
            v15 := GenS1{F0: (v5 & v5), F1: (v5 - (v5 - v5)), F2: true}
            switch "" {
            case "a":
                v6[(fi14 & fi14)] = false
            default:
                for fi16 := 0; fi16 < 3; fi16++ {
                    v17 := GenS1{F0: v5, F1: (v5 ^ 255), F2: true}
                    if false {
                        var v18 int16 = v1
                        _ = v18
                    }
                    v2 = (true || v4)
                    _ = fi16
                    _ = v17
                }
            }
            var v19 uint = 255
            _ = fi14
            _ = v15
            _ = v19
        }
    case 32768:
        v2 = v4
    case 2147483647:
        v3 = append(v3, 42)
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
