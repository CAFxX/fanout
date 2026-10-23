package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 bool }

func main() {
    var v1 uint8 = 131
    var v2 uint64 = 111
    var v3 uint64 = (v2 * 1)
    var v4 bool = true
    fmt.Println(v2, v1, v4)
    switch -1 {
    case 117:
        v2 += (^(v2))
    default:
        var v5 uint16 = 42
        _ = v5
    }
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: uint16(v1), F1: (v6 <= (((v6 / (((1.5 - v6) - ((v6 + (float64((1e10 + 0.0)) * 1.5)) - ((v6 * 3.14159) * v6))) * v6)) * (v6 + v6)) - (v6 + (v6 - 0.0)))), F2: (45 >= 42)}
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 float64 = v6
            _ = fi9
            _ = v10
        }
        var v11 int16 = 42
        _ = fi7
        _ = v11
        _ = v8
    }
    switch 127 {
    case 430:
        switch 0 {
        case 127:
            if v4 {
                fmt.Println(v2, v1, v4)
                for fi12 := 0; fi12 < 16; fi12++ {
                    v13 := []int{fi12, (fi12 ^ 758), (fi12 * -65)}
                    v14 := GenS1{F0: 255, F1: false, F2: v4}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
        case 42:
            switch 4294967295 {
            case -25:
                var v15 string = "a"
                _ = v15
            default:
                v16 := []int{}
                _ = v16
            }
        default:
            v4 = true
        }
    }
    var v17 int64 = -1
    fmt.Println(v2)
    v18 := []int{32767, 32768}
    v6 += float64(((0.0 * 3.14159) / 1e10))
    var v19 uint16 = 706
    var v20 uint16 = v19
    var v21 uint = 118
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v6
}
