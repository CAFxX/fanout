package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint16 }
func genFn0() int {
    var v6 int = -2147483648
    v6 = 0
    switch ("\n\t" + ("foo bar" + "hello")) {
    case "":
        v7 := GenS1{F0: 185}
        _ = v7
    case "\n\t":
        var v8 int32 = 1
        _ = v8
    case "a":
        var v9 int = (v6 % ((v6) | 1))
        _ = v9
    default:
        v10 := GenS1{F0: 2}
        _ = v10
    }
    _ = v6
    return (v6 ^ (v6 >> ((255) & 63)))
}

func main() {
    var v1 string = ("\n\t" + "\n\t")
    var v2 uint8 = 127
    var v3 int64 = -99
    var v4 uint64 = 456
    var v5 float64 = 0.0
    switch "hello" {
    case "a":
        v11 := []int{-1, 638}
        _ = v11
    case "hello":
        fmt.Println(v4, v4, v3)
    }
    switch 32767 {
    case 2147483647:
        v12 := GenS1{F0: (v2 + v2)}
        _ = v12
    case 0:
        v13 := []int{4294967295, 0, 9223372036854775807}
        _ = v13
    }
    v14 := GenS1{F0: v2}
    v5 += v5
    v15 := []int{}
    v16 := GenS2{F0: 124, F1: 65535, F2: 42}
    v3 -= -680
    v5 = 3.14159
    var v17 int16 = 256
    v17 = (v17 % ((v17) | 1))
    _ = len(v15)
    v15 = append(v15, -1)
    for fi18 := 0; fi18 < 11; fi18++ {
        switch int(v17) {
        case 801:
            var v19 uint16 = 366
            _ = v19
        default:
            var v20 int64 = (v3 / (((v3 / ((-1) | 1))) | 1))
            _ = v20
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
