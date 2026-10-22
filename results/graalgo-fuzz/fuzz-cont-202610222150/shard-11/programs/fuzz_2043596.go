package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    var v1 string = ""
    var v2 int8 = -33
    var v3 int64 = 532
    v4 := GenS1{F0: 42, F1: 538}
    v5 := []int{1}
    fmt.Println(v3)
    v1 = (("foo bar" + (v1 + v1)) + ((v1 + v1) + ("hello" + ((v1 + v1) + v1))))
    fmt.Println(v3, v1, v2)
    for fi6 := 0; fi6 < 7; fi6++ {
        if (!(32767 > 489)) {
            switch 4294967295 {
            case 0:
                var v7 uint8 = 255
                _ = v7
            case 32768:
                if len(v5) > 0 {
                    v5[0] = fi6
                }
            }
            if len(v5) > 0 {
                v5[0] = fi6
            }
        }
        _ = fi6
    }
    switch "" {
    case "\n\t":
        var v8 uint32 = 0
        _ = v8
    case "foo bar":
        switch v1 {
        case "\n\t":
            v2 += -23
        case "a":
            v9 := GenS2{F0: ("foo bar" + (v1 + ("foo bar" + "\n\t"))), F1: 18446744073709551615}
            _ = v9
        default:
            fmt.Println(v3, v2)
        }
    }
    fmt.Println(v2, v2, v2)
    var v10 bool = (42 == 271)
    v1 = v1
    v3 = (v3 % ((v3) | 1))
    v1 = v1
    var v11 uint64 = 1
    _ = len(v5)
    fmt.Println(v2, v1)
    switch -1 {
    case 9223372036854775807:
        _ = len(v5)
    default:
        v12 := GenS2{F0: (((v1 + v1) + v1) + (("foo bar" + v1) + v1)), F1: (v11 << ((18446744073709551615) & 63))}
        _ = v12
    }
    switch v1 {
    case "a":
        _ = len(v5)
    case "foo bar":
        var v13 uint = 2
        _ = v13
    default:
        for fi14 := 0; fi14 < 2; fi14++ {
            _ = len(v5)
            var v15 uint8 = 127
            _ = fi14
            _ = v15
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
