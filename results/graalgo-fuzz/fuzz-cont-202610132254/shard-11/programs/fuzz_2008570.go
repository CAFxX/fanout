package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 int64 }
func genFn0(p0 bool, p1 uint64, p2 uint8) int8 {
    v4 := GenS1{F0: p0, F1: (-32768 < 32767)}
    fmt.Println(p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 127
}

func main() {
    var v1 string = "foo bar"
    var v2 int32 = 825
    var v3 uint64 = 2
    v5 := GenS1{F0: (("" >= v1) && (714 == 42)), F1: (18446744073709551615 != 4294967295)}
    for fi6 := 0; fi6 < 3; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            v3 -= (v3 >> ((666) & 63))
            var v8 int32 = -132
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    v9 := []int{-545, 42, 0}
    switch -9223372036854775808 {
    case 1:
        if ((v1 + v1) != "foo bar") {
            v10 := GenS1{F0: false, F1: ((0 >= 255) && false)}
            v3 = (v3 * (v3 << ((255) & 63)))
            _ = v10
        }
    default:
        if (704 > 788) {
            v11 := GenS2{F0: 1}
            _ = v11
        }
    }
    var v12 uint64 = (^(v3))
    v1 = v1
    var v13 string = ("\n\t" + "foo bar")
    fmt.Println(v13, v13, v2)
    for fi14 := 0; fi14 < 15; fi14++ {
        if len(v9) > 0 {
            v9[0] = (fi14 >> ((2) & 63))
        }
        v9 = append(v9, (fi14 * fi14))
        v1 = ("hello" + v1)
        _ = fi14
    }
    _ = len(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v9
}
