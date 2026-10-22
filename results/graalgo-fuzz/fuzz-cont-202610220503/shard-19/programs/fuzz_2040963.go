package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 int16; F1 uint8 }

func main() {
    var v1 int8 = -39
    var v2 uint16 = 32767
    var v3 uint8 = 0
    v4 := GenS1{F0: v2, F1: int16(v1), F2: false}
    var v5 float64 = -2.25
    var v6 int16 = 1
    fmt.Println(v6, v3, v1)
    v7 := GenS1{F0: 690, F1: (v6 | (v6 + 1)), F2: false}
    var v8 uint64 = 1
    switch ("hello" + ("" + "\n\t")) {
    case "a":
        for fi9 := 0; fi9 < 9; fi9++ {
            var v10 int8 = (v1 + v1)
            switch fi9 {
            case -575:
                if false {
                    for fi11 := 0; fi11 < 7; fi11++ {
                        v3 -= v3
                        _ = fi11
                    }
                    v1 += v1
                }
            case 42:
                if true {
                    fmt.Println(v2, v3, v10)
                    var v12 int = fi9
                    _ = v12
                } else {
                    fmt.Println(v8)
                    v13 := GenS1{F0: v2, F1: -273, F2: ((-(fi9)) < 219)}
                    fmt.Println(v1, v2)
                    _ = v13
                }
            }
            var v14 float64 = v5
            _ = fi9
            _ = v10
            _ = v14
        }
    default:
        var v15 int32 = 0
        _ = v15
    }
    var v16 int8 = (v1 * v1)
    v17 := GenS1{F0: (v2 & v2), F1: v6, F2: false}
    if ((float64(((1e10 - 3.14159) - (1.5 * ((-2.25 - 0.1) - float32(float64(-2.25)))))) * 0.0) == ((0.1 * v5) + float64(((float32(1e10) / 1.0) - (0.1 / 3.14159))))) {
        v2 += (v2 * 1)
        v18 := GenS2{F0: (^(v6)), F1: v3}
        _ = v18
    } else {
        switch 4294967295 {
        case -378:
            fmt.Println(v1, v16, v2)
        case 32768:
            v8 = 4294967295
        }
        v19 := GenS2{F0: v6, F1: (^(v3))}
        _ = v19
    }
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
