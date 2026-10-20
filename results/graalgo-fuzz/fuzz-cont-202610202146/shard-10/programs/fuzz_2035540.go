package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 string; F1 uint32 }

func main() {
    v1 := make(map[int]string)
    var v2 int16 = -1
    var v3 int64 = 638
    var v4 uint64 = 1
    var v5 float64 = 1.5
    v2 -= -1
    switch "" {
    case "foo bar":
        v3 -= -773
    case "":
        var v6 uint8 = 1
        _ = v6
    default:
        var v7 uint = 2
        _ = v7
    }
    var v8 bool = true
    for fi9 := 0; fi9 < 20; fi9++ {
        var v10 int64 = 330
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: 127, F1: 2147483647}
    for fi12 := 0; fi12 < 6; fi12++ {
        if (!v8) {
            delete(v1, (fi12 ^ (fi12 ^ -894)))
        }
        v13 := GenS1{F0: 42, F1: -884}
        v4 = (v4 - 50)
        _ = fi12
        _ = v13
    }
    switch "\n\t" {
    case "\n\t":
        v14 := GenS2{F0: ("hello" + ""), F1: 42}
        _ = v14
    case "foo bar":
        delete(v1, -9223372036854775808)
    case "hello":
        switch ("hello" + "") {
        case "\n\t":
            if v8 {
                for fi15 := 0; fi15 < 18; fi15++ {
                    var v16 float64 = (((v5 + (v5 / (0.1 + v5))) / (((v5 + v5) + v5) - 1.5)) - v5)
                    _ = fi15
                    _ = v16
                }
                v1[-312] = (("a" + ("hello" + "foo bar")) + "\n\t")
            }
        default:
            v1[363] = ("foo bar" + "hello")
        }
    }
    v17 := v1[9223372036854775807]
    v1[9223372036854775807] = "a"
    var v18 uint16 = uint16(v4)
    v4 = (v4 ^ v4)
    v19 := GenS1{F0: 2, F1: 822}
    for fi20 := 0; fi20 < 8; fi20++ {
        for fi21 := 0; fi21 < 8; fi21++ {
            v1[(fi21 | int(v2))] = v17
            _ = fi21
        }
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
