package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    var v1 bool = (0.0 > 0.1)
    var v2 int64 = 0
    var v3 int64 = v2
    var v4 bool = (171 == 42)
    fmt.Println(v2, v2, v1)
    v3 = (v2 + (-(v3)))
    switch 599 {
    case -805:
        v1 = v1
    default:
        v5 := GenS1{F0: 0, F1: "foo bar"}
        _ = v5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v3 += v3
        _ = fi6
    }
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: 4294967295, F1: ""}
        v9 := []int{}
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 19; fi10++ {
        if ((65535 <= 42) || (42 == -1)) {
            if ((254 != -1) && (3.14159 <= ((0.1 * 1e10) / (3.14159 - -2.25)))) {
                var v11 bool = (((v3 / (((v2 >> ((244) & 63))) | 1)) < (v2 * (-(v2)))) || false)
                _ = v11
            } else {
                var v12 int16 = -32768
                v13 := GenS2{F0: 42, F1: v3}
                v14 := []int{(-(fi10)), -2147483648, (-(fi10)), 32767}
                _ = v12
                _ = v13
                _ = v14
            }
        }
        _ = fi10
    }
    fmt.Println(v3, v1, v1)
    var v15 float32 = -2.25
    var v16 uint16 = 923
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
