package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint64 }

func main() {
    var v1 uint8 = 172
    var v2 int32 = -356
    var v3 int = 42
    v4 := make(map[string]string)
    var v5 uint16 = 32767
    var v6 uint64 = 0
    switch "\n\t" {
    case "\n\t":
        if true {
            v5 += 255
            switch v3 {
            case 1:
                for fi7 := 0; fi7 < 6; fi7++ {
                    if true {
                        var v8 uint64 = v6
                        _ = v8
                    }
                    _ = fi7
                }
            case 42:
                fmt.Println(v3)
            }
        }
    case "hello":
        v3 = -102
    case "a":
        v9 := GenS1{F0: -970, F1: ("\n\t" + ""), F2: v6}
        _ = v9
    default:
        v3 += (v3 << ((461) & 63))
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        fmt.Println(v2, v6, v6)
        _ = fi10
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        v3 -= (v3 ^ -474)
        var v12 float64 = -2.25
        _ = fi11
        _ = v12
    }
    var v13 float32 = 1e10
    switch -2147483648 {
    case 127:
        v1 += v1
    }
    var v14 uint8 = v1
    v3 = 32767
    v15 := GenS1{F0: (v2 + v2), F1: ("\n\t" + ("a" + "foo bar")), F2: (v6 / ((18446744073709551615) | 1))}
    v14 -= v14
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
