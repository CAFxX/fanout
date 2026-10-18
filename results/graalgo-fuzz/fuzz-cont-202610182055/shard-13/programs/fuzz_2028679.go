package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]uint32)
    var v3 int16 = 1000
    var v4 int32 = 1
    for fi5 := 0; fi5 < 9; fi5++ {
        for fi6 := 0; fi6 < 12; fi6++ {
            v3 = v3
            v1 = v1
            _ = fi6
        }
        for fi7 := 0; fi7 < 14; fi7++ {
            fmt.Println(fi7)
            _ = fi7
        }
        _ = fi5
    }
    v4 -= 532
    v8 := GenS1{F0: (127 <= -1), F1: 18446744073709551615}
    var v9 int16 = (v3 - (v3 ^ v3))
    switch 9223372036854775807 {
    case 4294967295:
        v1 = (^(v1))
    }
    v4 -= (v4 ^ (v4 << ((42) & 63)))
    if true {
        fmt.Println(v9)
    }
    v4 -= (v4 - int32(v9))
    v4 += (v4 % ((138) | 1))
    v4 += (v4 ^ 1)
    fmt.Println(v3)
    var v10 uint = 869
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int = (fi12 & fi12)
            var v14 uint16 = 0
            var v15 int16 = (v9 * v9)
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi11
    }
    var v16 bool = false
    for fi17 := 0; fi17 < 9; fi17++ {
        v4 = (^(v4))
        var v18 int16 = (v9 >> ((255) & 63))
        _ = fi17
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
