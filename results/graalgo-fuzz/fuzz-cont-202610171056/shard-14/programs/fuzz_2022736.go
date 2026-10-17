package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 int16 }

func main() {
    var v1 uint32 = 759
    var v2 float64 = -2.25
    v3 := make(map[string]string)
    v4 := GenS1{F0: 276, F1: -31, F2: 8}
    var v5 uint32 = (v1 + 0)
    var v6 int16 = 42
    v7 := []int{565}
    v1 -= 686
    v7 = append(v7, 590)
    var v8 float32 = 1e10
    v9 := GenS1{F0: 1, F1: -154, F2: (^(v6))}
    v10 := GenS1{F0: 65535, F1: -1, F2: (v6 + -1)}
    var v11 uint16 = 2
    var v12 uint64 = 627
    var v13 int32 = 2147483647
    if false {
        for fi14 := 0; fi14 < 19; fi14++ {
            for fi15 := 0; fi15 < 17; fi15++ {
                if true {
                    v16 := GenS1{F0: 32767, F1: -9223372036854775808, F2: (v6 & v6)}
                    _ = v16
                }
                var v17 float64 = (((float64(v8) * 0.1) * (float64(v8) / 0.1)) + 1.5)
                v17 += 3.14159
                _ = fi15
                _ = v17
            }
            v7 = append(v7, (fi14 << ((65535) & 63)))
            _ = fi14
        }
        switch ("hello" + "\n\t") {
        case "\n\t":
            switch 32767 {
            case 4294967295:
                var v18 int64 = -1
                _ = v18
            case 1:
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            default:
                v11 -= v11
            }
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    } else {
        var v19 int = 32767
        var v20 bool = (v11 <= 68)
        _ = v19
        _ = v20
    }
    fmt.Println(v13, v11, v2)
    v13 += (-(v13))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
