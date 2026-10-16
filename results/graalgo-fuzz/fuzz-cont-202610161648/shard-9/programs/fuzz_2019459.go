package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }

func main() {
    var v1 string = ("foo bar" + "hello")
    v2 := []string{"a", (v1 + ((((v1 + "a") + ((v1 + "a") + "")) + "a") + "a"))}
    var v3 uint16 = 310
    if (2 >= 18446744073709551615) {
        v3 -= v3
        var v4 uint16 = (v3 % ((v3) | 1))
        _ = v4
    }
    var v5 uint8 = 107
    fmt.Println(v5, v1, v1)
    v1 = ("a" + (v1 + "hello"))
    switch v1 {
    case "foo bar":
        switch -1 {
        case 32768:
            for fi6 := 0; fi6 < 13; fi6++ {
                var v7 float64 = float64((1.5 / -2.25))
                _ = fi6
                _ = v7
            }
        case 645:
            if (true && true) {
                v5 += v5
                v8 := GenS1{F0: 1, F1: 388}
                _ = v8
            }
        default:
            v3 += v3
        }
    case "hello":
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 uint32 = 2147483647
            _ = fi9
            _ = v10
        }
    }
    if ((uint32(v5) > 2147483647) || (true || (((v1 + (v1 + ("a" + ""))) > (v1 + "hello")) && ((true && true) && ((true && true) && false))))) {
        v1 = v1
    } else {
        fmt.Println(v5)
        var v11 uint32 = 123
        var v12 int64 = -1
        _ = v11
        _ = v12
    }
    if (-469 != -9223372036854775808) {
        if (18446744073709551615 != 21) {
            if true {
                var v13 uint16 = 585
                var v14 uint32 = 857
                _ = v13
                _ = v14
            }
            for fi15 := 0; fi15 < 19; fi15++ {
                fmt.Println(v1, v3, v1)
                v5 = (v5 + v5)
                _ = fi15
            }
            switch -9223372036854775808 {
            case 42:
                v5 += v5
            case -9223372036854775808:
                var v16 int16 = 32767
                _ = v16
            }
        }
    }
    fmt.Println(v1, v3, v3)
    v2 = append(v2, "hello")
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
