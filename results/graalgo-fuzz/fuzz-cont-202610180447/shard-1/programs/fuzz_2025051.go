package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int }
type GenS2 struct { F0 uint64; F1 bool; F2 string }

func main() {
    var v1 uint = 227
    var v2 int64 = int64(v1)
    v3 := make(map[int]bool)
    var v4 int64 = v2
    var v5 string = ("" + ("" + (("hello" + "\n\t") + "")))
    switch int(v1) {
    case 127:
        v6 := GenS1{F0: v5, F1: int8(v1), F2: 127}
        _ = v6
    case -724:
        for fi7 := 0; fi7 < 12; fi7++ {
            v5 = ("hello" + ("foo bar" + "\n\t"))
            if (((1 == 42) || false) || false) {
                var v8 int = (^(fi7))
                var v9 float64 = (float64(1e10) / 0.1)
                _ = v8
                _ = v9
            } else {
                v10 := v3[(fi7 | 1)]
                _ = v10
            }
            v5 = ("hello" + (v5 + (((v5 + v5) + v5) + "a")))
            _ = fi7
        }
    case -217:
        if (1 >= 42) {
            delete(v3, 32767)
            v2 -= (-(v2))
        }
    }
    var v11 int16 = -1
    var v12 uint8 = 127
    v13 := GenS1{F0: v5, F1: 0, F2: -9223372036854775808}
    v1 -= 439
    if (964 != 103) {
        v14 := []int{128, 127, 1, 0}
        _ = v14
    }
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 17; fi16++ {
            var v17 int16 = 32767
            var v18 int = fi15
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            var v19 int64 = 0
            var v20 float64 = 1e10
            v21 := GenS2{F0: 18446744073709551615, F1: true, F2: "\n\t"}
            _ = v19
            _ = v20
            _ = v21
        } else {
            if true {
                var v22 int8 = 42
                delete(v3, 178)
                _ = v22
            }
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
