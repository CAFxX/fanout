package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int; F2 uint }
type GenS4 struct { F0 int64; F1 int16; F2 int32 }

func main() {
    var v1 uint32 = 2147483647
    var v2 uint32 = v1
    var v3 int16 = -842
    var v4 string = "\n\t"
    v3 -= (-(v3))
    v1 = v2
    v5 := GenS1{F0: v4, F1: 707}
    var v6 float64 = (0.0 * float64((((((-2.25 * ((0.1 + 1e10) - -2.25)) - ((1.5 * 0.0) - -2.25)) / 1.5) * 0.0) * 0.0)))
    var v7 string = v4
    v6 += v6
    fmt.Println(v7)
    v8 := GenS1{F0: (v7 + (v4 + ((((v7 + v7) + v4) + (("a" + v4) + v7)) + (v4 + ((v4 + "a") + (v4 + v7)))))), F1: 65535}
    v4 = v4
    if (false && (!true)) {
        switch -1 {
        case -903:
            v9 := GenS2{F0: (-169 == 0)}
            _ = v9
        default:
            v10 := GenS3{F0: 2, F1: -2147483648, F2: 1}
            _ = v10
        }
        if true {
            switch 1 {
            case 411:
                for fi11 := 0; fi11 < 3; fi11++ {
                    v7 = "foo bar"
                    v12 := GenS3{F0: 32767, F1: (fi11 - fi11), F2: 2}
                    var v13 float64 = 0.0
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            default:
                v2 += 2147483648
            }
            v3 = (v3 + (-(v3)))
        }
        for fi14 := 0; fi14 < 4; fi14++ {
            v15 := []int{}
            _ = fi14
            _ = v15
        }
    }
    fmt.Println(v2)
    v16 := GenS4{F0: -972, F1: v3, F2: -989}
    var v17 uint32 = 428
    var v18 uint32 = v1
    v19 := GenS3{F0: 483, F1: -918, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
