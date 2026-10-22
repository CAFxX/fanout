package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint }

func main() {
    var v1 uint32 = 915
    var v2 uint8 = 127
    var v3 int8 = 0
    v4 := GenS1{F0: 457, F1: (v1 * (^(v1)))}
    v5 := []int{32767}
    v1 = v1
    switch 509 {
    case 33:
        if len(v5) > 0 {
            v5[0] = 492
        }
    }
    for fi6 := 0; fi6 < 10; fi6++ {
        v1 -= (v1 / ((2147483648) | 1))
        _ = fi6
    }
    v1 = (v1 << ((345) & 63))
    switch "\n\t" {
    case "foo bar":
        if (!(795 == 65535)) {
            v3 -= (v3 - (v3 >> ((18446744073709551615) & 63)))
        } else {
            var v7 bool = (3.14159 > 1.5)
            if (0.0 < 0.1) {
                for fi8 := 0; fi8 < 12; fi8++ {
                    var v9 int8 = (v3 % (((v3 << ((65535) & 63))) | 1))
                    if v7 {
                        var v10 uint64 = 1
                        v1 = (v1 << ((670) & 63))
                        v5 = append(v5, (fi8 | fi8))
                        _ = v10
                    } else {
                        v11 := GenS2{F0: false, F1: 0}
                        _ = v11
                    }
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    _ = fi8
                    _ = v9
                }
            }
            switch ("\n\t" + "\n\t") {
            case "\n\t":
                for fi12 := 0; fi12 < 9; fi12++ {
                    v7 = (!(!((!(((781 == 2) || false) || (("foo bar" + "\n\t") <= ("a" + (("" + (((((("" + ("foo bar" + (("foo bar" + "") + ""))) + "") + "hello") + "foo bar") + "hello") + ("foo bar" + ""))) + "a"))))) && true)))
                    _ = fi12
                }
            case "":
                var v13 int64 = -138
                _ = v13
            }
            _ = v7
        }
    case "":
        v2 = (v2 * (v2 ^ v2))
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        switch "foo bar" {
        case "hello":
            v1 = (v1 & (^(v1)))
        default:
            v15 := GenS2{F0: true, F1: 42}
            _ = v15
        }
        fmt.Println(v1, fi14, v2)
        v16 := GenS1{F0: (-(fi14)), F1: 694}
        _ = fi14
        _ = v16
    }
    v1 += 846
    fmt.Println(v3, v2, v2)
    v3 += 42
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v17 uint8 = (v2 % (((v2 / (((^(v2))) | 1))) | 1))
    v17 -= v17
    switch int(v1) {
    case -730:
        _ = len(v5)
    case -1:
        var v18 float32 = float32(-2.25)
        _ = v18
    default:
        for fi19 := 0; fi19 < 18; fi19++ {
            var v20 bool = (fi19 < (fi19 % ((fi19) | 1)))
            _ = fi19
            _ = v20
        }
    }
    fmt.Println(v3, v2)
    if (-1 <= 42) {
        var v21 int32 = -1
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
