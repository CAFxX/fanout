package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }

func main() {
    var v1 int64 = -1
    v2 := make(map[int]uint32)
    var v3 uint32 = 77
    switch 42 {
    case 1:
        for fi4 := 0; fi4 < 17; fi4++ {
            v2[fi4] = (v3 & 1)
            v1 = v1
            v3 = v3
            _ = fi4
        }
    case -753:
        if true {
            v1 = (v1 ^ v1)
            var v5 uint8 = 1
            _ = v5
        }
    }
    var v6 int16 = 1
    v7 := GenS1{F0: 127, F1: (!((!true) && false)), F2: v1}
    for fi8 := 0; fi8 < 10; fi8++ {
        for fi9 := 0; fi9 < 19; fi9++ {
            if false {
                v10 := GenS1{F0: (fi9 | 127), F1: (("a" + "foo bar") > ((("foo bar" + "hello") + "") + ((("foo bar" + "") + "foo bar") + "a"))), F2: (-(v1))}
                v11 := GenS1{F0: (fi8 - 0), F1: true, F2: (v1 + v1)}
                v2[-40] = (v3 / ((887) | 1))
                _ = v10
                _ = v11
            }
            _ = fi9
        }
        v1 += int64(v3)
        _ = fi8
    }
    v12 := GenS2{F0: 0, F1: 2, F2: 18446744073709551615}
    var v13 uint32 = v3
    v14 := GenS1{F0: 143, F1: true, F2: 9223372036854775807}
    for fi15 := 0; fi15 < 7; fi15++ {
        v3 += v13
        v16 := GenS2{F0: -101, F1: 127, F2: 2}
        _ = fi15
        _ = v16
    }
    var v17 uint32 = v13
    switch -2147483648 {
    case 32768:
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 int = fi18
            v1 += (v1 * v1)
            var v20 uint16 = 644
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := v2[246]
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v6
    _ = v7
}
