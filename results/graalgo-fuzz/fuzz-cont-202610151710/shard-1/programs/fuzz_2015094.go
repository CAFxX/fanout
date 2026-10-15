package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 bool }

func main() {
    v1 := make(map[string]bool)
    v2 := []int{-1, -2147483648, 32767}
    var v3 uint = 458
    var v4 int8 = 108
    var v5 string = ("\n\t" + "\n\t")
    v1["foo bar"] = ((!false) && true)
    var v6 int64 = -9223372036854775808
    if (!(v6 <= (-(v6)))) {
        for fi7 := 0; fi7 < 14; fi7++ {
            if (1.5 <= 0.0) {
                _ = len(v1)
                _ = len(v1)
                fmt.Println(v4, fi7)
            }
            if false {
                var v8 int8 = (v4 << ((v3) & 63))
                _ = v8
            }
            _ = fi7
        }
    } else {
        var v9 int64 = (^(v6))
        if (!false) {
            for fi10 := 0; fi10 < 2; fi10++ {
                v9 -= (^(v9))
                v6 += v9
                _ = fi10
            }
            switch 9223372036854775807 {
            case -458:
                v11 := GenS1{F0: v9, F1: (^(v4)), F2: ((v3 ^ 4294967295) >= 4294967295)}
                _ = v11
            default:
                fmt.Println(v6, v5, v6)
            }
        } else {
            switch v5 {
            case "":
                switch 415 {
                case 9223372036854775807:
                    var v12 bool = (!(!true))
                    _ = v12
                default:
                    var v13 uint = 2
                    _ = v13
                }
            case "\n\t":
                _ = len(v1)
            case "a":
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 int32 = -2147483648
                    _ = fi14
                    _ = v15
                }
            }
        }
        _ = v9
    }
    v6 += v6
    v6 = (-(v6))
    v4 = (v4 % ((100) | 1))
    var v16 uint = 255
    v4 = v4
    var v17 int8 = v4
    v3 = v3
    var v18 float32 = 0.0
    var v19 string = v5
    for fi20 := 0; fi20 < 12; fi20++ {
        _ = len(v2)
        v2 = append(v2, fi20)
        _ = fi20
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
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
}
