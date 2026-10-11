package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 uint8 = 242
    switch 759 {
    case 9223372036854775807:
        fmt.Println(v3)
    case 128:
        if true {
            v3 -= v3
        }
    default:
        var v4 int16 = 1
        _ = v4
    }
    if false {
        if (1 < 127) {
            fmt.Println(v3)
        }
        for fi5 := 0; fi5 < 8; fi5++ {
            var v6 bool = true
            var v7 int64 = 481
            _ = fi5
            _ = v6
            _ = v7
        }
        var v8 int32 = -249
        _ = v8
    }
    var v9 bool = false
    var v10 int = 9223372036854775807
    for fi11 := 0; fi11 < 13; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            v13 := GenS1{F0: int16(v3)}
            if (32767 <= 0) {
                var v14 bool = v9
                if (((fi12 <= (fi12 >> ((65535) & 63))) && (0 > 42)) || (65535 <= 186)) {
                    v15 := GenS1{F0: -32768}
                    v16 := GenS2{F0: 480, F1: 697, F2: (2 < 42)}
                    v17 := v1[4294967295]
                    _ = v15
                    _ = v16
                    _ = v17
                }
                _ = v14
            } else {
                fmt.Println(fi12, v9, v10)
                v18 := GenS1{F0: 32767}
                _ = v18
            }
            _ = fi12
            _ = v13
        }
        if ((v3 < 191) || v9) {
            v19 := GenS2{F0: 698, F1: 2147483647, F2: ((3.14159 / (3.14159 + (3.14159 - (3.14159 / (1.5 / (((3.14159 / 3.14159) + -2.25) + 0.0)))))) == 1.5)}
            _ = v19
        }
        switch ("" + "hello") {
        case "a":
            v20 := v1[719]
            _ = v20
        case "\n\t":
            fmt.Println(v9, v9)
        }
        _ = fi11
    }
    v3 -= (v3 + 128)
    var v21 string = "foo bar"
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 int8 = -50
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        var v25 float32 = -2.25
        _ = fi24
        _ = v25
    }
    v3 += 28
    fmt.Println(v21, v21, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v9
}
