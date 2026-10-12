package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 string; F1 bool; F2 int }
func genFn0(p0 int16) int64 {
    if (!(!(!false))) {
        p0 = p0
    }
    switch "a" {
    case "\n\t":
        fmt.Println(p0)
    case "hello":
        v6 := []int{}
        _ = v6
    default:
        p0 -= p0
    }
    var v7 bool = false
    v8 := []int{42}
    _ = p0
    _ = v7
    _ = v8
    return -695
}

func main() {
    var v1 float64 = float64((-2.25 / 3.14159))
    var v2 int32 = 1
    var v3 uint16 = 255
    var v4 float32 = 1e10
    var v5 uint16 = (v3 >> ((0) & 63))
    var v9 uint8 = 245
    for fi10 := 0; fi10 < 1; fi10++ {
        fmt.Println(v2)
        _ = fi10
    }
    v11 := GenS1{F0: 205, F1: -332}
    v3 = v3
    v4 -= (((float32(v1) - (v4 - (float32(3.14159) + ((v4 + v4) - ((v4 - (((v4 - v4) - 0.1) * 0.1)) / 3.14159))))) * -2.25) - (v4 + (float32((0.1 / (float64(v4) * (v1 * v1)))) + v4)))
    for fi12 := 0; fi12 < 8; fi12++ {
        var v13 int64 = 1
        v9 += (v9 << ((255) & 63))
        _ = fi12
        _ = v13
    }
    v2 = v2
    if (true || false) {
        v14 := GenS1{F0: v9, F1: -9223372036854775808}
        var v15 float32 = (3.14159 / (v4 + 0.1))
        _ = v14
        _ = v15
    }
    v16 := []int{561, 53, 723}
    var v17 uint = 4294967295
    if (v17 <= v17) {
        switch "hello" {
        case "\n\t":
            v9 += v9
        case "a":
            _ = len(v16)
        default:
            for fi18 := 0; fi18 < 8; fi18++ {
                if len(v16) >= 2 {
                    v16 = v16[1:len(v16)]
                }
                _ = fi18
            }
        }
    } else {
        var v19 int = 421
        v20 := GenS2{F0: ("\n\t" + "hello"), F1: (true && true), F2: v19}
        v21 := GenS1{F0: v9, F1: v19}
        _ = v19
        _ = v20
        _ = v21
    }
    v9 = (v9 + (^(v9)))
    v22 := GenS1{F0: (v9 + v9), F1: int(v2)}
    v23 := GenS2{F0: "foo bar", F1: false, F2: -136}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
