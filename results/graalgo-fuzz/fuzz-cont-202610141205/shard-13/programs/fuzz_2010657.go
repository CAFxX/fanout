package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int; F1 string; F2 uint8 }

func main() {
    var v1 float32 = 0.1
    var v2 float64 = 1e10
    v3 := []uint32{4294967295, 2147483647, 2147483647, 386}
    var v4 uint16 = 32767
    var v5 int16 = -1
    if false {
        fmt.Println(v2, v4)
        v3 = append(v3, 941)
        _ = len(v3)
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        switch 0 {
        case 218:
            v7 := GenS1{F0: 1, F1: (fi6 | fi6)}
            _ = v7
        case 2147483647:
            v2 -= v2
        default:
            switch (^(fi6)) {
            case 0:
                for fi8 := 0; fi8 < 11; fi8++ {
                    fmt.Println(fi8)
                    _ = fi8
                }
            default:
                var v9 float64 = v2
                _ = v9
            }
        }
        v10 := GenS1{F0: 1, F1: 9223372036854775807}
        _ = fi6
        _ = v10
    }
    switch 32768 {
    case 750:
        if true {
            v11 := GenS1{F0: 2, F1: 2147483647}
            v5 = v5
            _ = v11
        } else {
            v1 = (float32(0.0) / v1)
            switch ("a" + "\n\t") {
            case "hello":
                v1 -= v1
            case "foo bar":
                switch "\n\t" {
                case "\n\t":
                    v4 = (v4 / ((v4) | 1))
                }
            }
        }
    case -426:
        fmt.Println(v1)
    case -775:
        var v12 bool = false
        _ = v12
    default:
        var v13 int8 = 0
        _ = v13
    }
    var v14 string = ("\n\t" + ("" + ""))
    v15 := GenS2{F0: -977, F1: "", F2: 244}
    v16 := GenS1{F0: 0, F1: 0}
    var v17 uint32 = 2147483648
    for fi18 := 0; fi18 < 12; fi18++ {
        if false {
            v5 += int16(fi18)
            var v19 int32 = 1
            if false {
                v3 = append(v3, 1)
            } else {
                if len(v3) > 0 {
                    v3[0] = (v17 * (v17 | v17))
                }
            }
            _ = v19
        }
        _ = fi18
    }
    v3 = append(v3, v17)
    fmt.Println(v1, v5)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
