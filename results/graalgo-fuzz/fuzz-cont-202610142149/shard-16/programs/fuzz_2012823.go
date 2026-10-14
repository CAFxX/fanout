package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint64 = 1
    var v2 int = 127
    var v3 int8 = -64
    v4 := []uint32{1}
    var v5 string = "foo bar"
    v3 -= 0
    for fi6 := 0; fi6 < 7; fi6++ {
        var v7 uint32 = 605
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi6
        _ = v7
    }
    var v8 float64 = ((-2.25 - (3.14159 + 1.5)) / -2.25)
    v3 -= 0
    v9 := GenS1{F0: 175, F1: 0, F2: 2147483648}
    v2 = v2
    var v10 uint32 = 2
    v8 = (v8 + ((-2.25 * (float64(0.0) - ((v8 * 1e10) - float64(1.5)))) + (float64((((1.5 + 3.14159) - -2.25) - (-2.25 * 3.14159))) * 1e10)))
    for fi11 := 0; fi11 < 14; fi11++ {
        v12 := GenS2{F0: v1}
        if (232 < 82) {
            v3 = (v3 | v3)
            switch v5 {
            case "":
                v13 := GenS1{F0: 255, F1: 0, F2: v10}
                _ = v13
            case "foo bar":
                v8 += v8
            }
            var v14 uint32 = (v10 ^ 0)
            _ = v14
        }
        var v15 int8 = (v3 ^ 100)
        _ = fi11
        _ = v12
        _ = v15
    }
    var v16 float64 = v8
    v17 := GenS1{F0: 2, F1: 146, F2: 42}
    fmt.Println(v8, v2)
    var v18 uint64 = uint64(v10)
    for fi19 := 0; fi19 < 7; fi19++ {
        for fi20 := 0; fi20 < 4; fi20++ {
            fmt.Println(v10, v3, v8)
            var v21 uint = 0
            _ = len(v4)
            _ = fi20
            _ = v21
        }
        var v22 int = (fi19 / (((fi19 - (^(fi19)))) | 1))
        _ = fi19
        _ = v22
    }
    v23 := GenS1{F0: 128, F1: 255, F2: v10}
    for fi24 := 0; fi24 < 12; fi24++ {
        v5 = ""
        _ = fi24
    }
    v25 := GenS2{F0: (v18 & v1)}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
