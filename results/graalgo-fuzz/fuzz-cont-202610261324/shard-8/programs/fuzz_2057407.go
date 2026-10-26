package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 bool }
type GenS2 struct { F0 bool }

func main() {
    var v1 uint32 = 285
    v2 := make(map[int]string)
    var v3 float32 = 3.14159
    v4 := make(map[string]string)
    var v5 uint = 178
    v6 := []string{"hello", "a"}
    for fi7 := 0; fi7 < 4; fi7++ {
        switch (^(fi7)) {
        case 2147483647:
            v3 = (1.5 * v3)
        case -610:
            for fi8 := 0; fi8 < 7; fi8++ {
                fmt.Println(v3)
                v9 := GenS1{F0: (fi8 / ((127) | 1)), F1: 42, F2: (1 != 65535)}
                _ = fi8
                _ = v9
            }
        case 0:
            v3 += (v3 + v3)
        }
        fmt.Println(v3, v3)
        for fi10 := 0; fi10 < 7; fi10++ {
            v6 = append(v6, "a")
            _ = fi10
        }
        _ = fi7
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        v2[(fi11 >> ((18446744073709551615) & 63))] = "hello"
        for fi12 := 0; fi12 < 12; fi12++ {
            v13 := GenS2{F0: true}
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    switch ("hello" + ("" + "foo bar")) {
    case "\n\t":
        var v14 uint64 = 651
        _ = v14
    case "hello":
        if len(v6) > 0 {
            v6[0] = ("" + "\n\t")
        }
    default:
        var v15 uint32 = v1
        _ = v15
    }
    var v16 uint = (v5 | (v5 + 1))
    var v17 uint = (v5 % (((v16 - (v16 / ((65535) | 1)))) | 1))
    var v18 uint16 = 261
    delete(v4, ("foo bar" + ""))
    _ = len(v6)
    v19 := GenS2{F0: true}
    var v20 int = 378
    fmt.Println(v3, v16, v20)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
