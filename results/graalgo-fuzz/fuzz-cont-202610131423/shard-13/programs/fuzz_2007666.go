package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 bool; F1 string; F2 uint16 }

func main() {
    var v1 bool = false
    v2 := []int{-9223372036854775808, -9223372036854775808, -9223372036854775808}
    var v3 int64 = 42
    v4 := []int{-473, 0, -2147483648, 127}
    v5 := []uint32{42, 1, 2147483648}
    if v1 {
        for fi6 := 0; fi6 < 14; fi6++ {
            v7 := GenS1{F0: 2, F1: v3}
            _ = fi6
            _ = v7
        }
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS1{F0: 2, F1: v3}
            if v1 {
                if (!false) {
                    v1 = false
                    v3 -= v3
                } else {
                    var v10 float64 = (1.5 + 0.0)
                    fmt.Println(v3, fi8, fi8)
                    _ = v10
                }
                _ = len(v4)
            } else {
                v3 = (v3 >> ((255) & 63))
            }
            var v11 uint8 = 255
            _ = fi8
            _ = v11
            _ = v9
        }
        v3 += (v3 - v3)
    } else {
        if true {
            var v12 int = -2147483648
            _ = v12
        }
        v13 := GenS2{F0: v1, F1: ("a" + "\n\t"), F2: 1}
        _ = v13
    }
    var v14 bool = (91 == 42)
    var v15 int = -9223372036854775808
    v1 = (v3 >= 42)
    var v16 float64 = 1.5
    var v17 uint16 = 504
    v18 := GenS2{F0: v1, F1: ("hello" + ((("a" + ("\n\t" + "a")) + ("foo bar" + ("hello" + "a"))) + ("\n\t" + "foo bar"))), F2: (v17 - 0)}
    switch 127 {
    case -2147483648:
        fmt.Println(v1, v14, v14)
    }
    _ = len(v2)
    v15 -= -2147483648
    var v19 string = "\n\t"
    v20 := GenS1{F0: 4294967295, F1: (v3 | (^(v3)))}
    fmt.Println(v17, v1, v19)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
