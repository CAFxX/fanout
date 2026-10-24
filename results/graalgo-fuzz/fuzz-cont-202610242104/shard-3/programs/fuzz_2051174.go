package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64 }
type GenS2 struct { F0 uint8; F1 int }
func genFn0() uint {
    v5 := []int{9223372036854775807, -994}
    if len(v5) > 0 {
        v5[0] = 42
    }
    v6 := GenS1{F0: 42, F1: 1}
    var v7 string = ("" + "")
    v7 = v7
    _ = v5
    _ = v6
    _ = v7
    return 638
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = true
    var v3 int64 = -365
    v4 := []string{}
    var v8 bool = v2
    v3 -= 304
    var v9 int = 1
    v1["\n\t"] = ("a" + "hello")
    v10 := GenS2{F0: 1, F1: -9223372036854775808}
    if (-1 == 749) {
        v3 -= (v3 * (-(v3)))
        for fi11 := 0; fi11 < 9; fi11++ {
            var v12 string = "hello"
            _ = fi11
            _ = v12
        }
        fmt.Println(v3, v9)
    }
    switch "" {
    case "\n\t":
        _ = len(v1)
    case "foo bar":
        for fi13 := 0; fi13 < 6; fi13++ {
            for fi14 := 0; fi14 < 7; fi14++ {
                v8 = (((65535 < 32767) || ((-2.25 - 0.1) <= 3.14159)) && (4294967295 < 86))
                _ = fi14
            }
            fmt.Println(v2, fi13, fi13)
            _ = fi13
        }
    case "":
        v8 = (-2147483648 != 0)
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            v16 := GenS2{F0: 41, F1: 32768}
            var v17 uint8 = 0
            fmt.Println(v2, v3)
            _ = fi15
            _ = v16
            _ = v17
        }
    }
    v18 := GenS2{F0: 1, F1: (v9 / ((v9) | 1))}
    v19 := GenS1{F0: -862, F1: 606}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
