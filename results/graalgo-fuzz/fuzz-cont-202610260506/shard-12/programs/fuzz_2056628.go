package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int64; F1 uint64; F2 int32 }

func main() {
    var v1 int = -2147483648
    var v2 string = "a"
    v3 := make(map[uint32]bool)
    if true {
        var v4 float32 = (0.0 / 1e10)
        for fi5 := 0; fi5 < 11; fi5++ {
            var v6 float32 = v4
            for fi7 := 0; fi7 < 8; fi7++ {
                v1 -= fi7
                if false {
                    fmt.Println(v6)
                    v1 += (v1 << ((796) & 63))
                    var v8 uint8 = 121
                    _ = v8
                }
                var v9 float64 = ((1e10 + 1e10) - 0.1)
                _ = fi7
                _ = v9
            }
            _ = fi5
            _ = v6
        }
        _ = v4
    }
    var v10 int8 = 100
    switch -886 {
    case -2147483648:
        var v11 int = (v1 - (^(v1)))
        _ = v11
    case -1:
        v3[2] = false
    default:
        v12 := GenS1{F0: 4294967295, F1: (^(v10)), F2: 2147483647}
        _ = v12
    }
    v13 := GenS2{F0: 2147483647, F1: 18446744073709551615}
    v1 = (v1 & 32768)
    switch "\n\t" {
    case "a":
        var v14 string = (((v2 + "") + v2) + v2)
        _ = v14
    case "foo bar":
        switch v2 {
        case "hello":
            v1 -= v1
        case "\n\t":
            v10 -= 7
        case "":
            switch v1 {
            case 32768:
                for fi15 := 0; fi15 < 13; fi15++ {
                    var v16 uint32 = 159
                    _ = fi15
                    _ = v16
                }
            case 838:
                if ((255 < 2) && true) {
                    v17 := GenS3{F0: v1}
                    _ = v17
                }
            default:
                v10 = v10
            }
        default:
            v3[2147483648] = true
        }
    }
    v10 = v10
    var v18 uint32 = 0
    fmt.Println(v2, v2)
    v19 := GenS4{F0: 124, F1: 355, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v10
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
