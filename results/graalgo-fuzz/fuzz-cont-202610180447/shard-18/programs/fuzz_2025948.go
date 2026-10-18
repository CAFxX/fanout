package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 int8 }

func main() {
    var v1 uint = 215
    var v2 uint = v1
    v3 := make(map[string]string)
    v4 := []int{-956, 9223372036854775807, 1}
    var v5 string = ("a" + ("" + ("foo bar" + "foo bar")))
    var v6 int64 = 0
    fmt.Println(v2, v5, v2)
    switch v5 {
    case "\n\t":
        v7 := GenS1{F0: 71, F1: (v5 + v5), F2: 837}
        _ = v7
    case "a":
        var v8 bool = true
        _ = v8
    default:
        switch (v5 + v5) {
        case "hello":
            switch 675 {
            case -363:
                if false {
                    var v9 bool = (!false)
                    _ = v9
                }
            case -9223372036854775808:
                switch -197 {
                case 0:
                    for fi10 := 0; fi10 < 3; fi10++ {
                        var v11 float64 = ((float64(float32(1e10)) + 3.14159) / 1e10)
                        _ = fi10
                        _ = v11
                    }
                case 239:
                    var v12 uint32 = 508
                    _ = v12
                case 837:
                    var v13 uint16 = 224
                    _ = v13
                default:
                    v14 := v3[(v5 + (v5 + "foo bar"))]
                    _ = v14
                }
            default:
                v15 := GenS2{F0: false, F1: 426, F2: 100}
                _ = v15
            }
        case "a":
            var v16 int8 = -115
            _ = v16
        default:
            v1 += v1
        }
    }
    v4 = append(v4, 568)
    if false {
        v1 += (v2 & (v1 >> ((v1) & 63)))
        switch ((("foo bar" + "") + (v5 + v5)) + (v5 + (v5 + v5))) {
        case "":
            var v17 uint = v1
            _ = v17
        }
    }
    var v18 int32 = 42
    var v19 int16 = 553
    v20 := GenS1{F0: 160, F1: ("a" + "foo bar"), F2: 0}
    v19 = v19
    var v21 int64 = (-(v6))
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
