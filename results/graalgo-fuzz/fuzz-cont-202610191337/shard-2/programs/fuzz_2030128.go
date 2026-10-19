package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint32; F1 int; F2 int64 }

func main() {
    v1 := []int{798, -959, 42, -211}
    var v2 uint32 = 2147483647
    v3 := []int{-360, 42}
    var v4 uint8 = 1
    if true {
        switch (("hello" + ("hello" + "")) + (("foo bar" + "a") + ("hello" + ("\n\t" + "a")))) {
        case "a":
            var v5 bool = (208 <= 0)
            _ = v5
        default:
            v6 := GenS1{F0: 86, F1: int64(v4)}
            _ = v6
        }
        v2 = v2
        var v7 int16 = 1000
        _ = v7
    }
    if false {
        _ = len(v1)
    }
    var v8 int = 643
    var v9 float32 = 0.0
    var v10 float64 = 0.0
    for fi11 := 0; fi11 < 3; fi11++ {
        if true {
            var v12 string = "hello"
            _ = v12
        } else {
            var v13 uint32 = (^(v2))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if true {
                _ = len(v1)
                switch (-(v8)) {
                case 4294967295:
                    switch (^(fi11)) {
                    case 0:
                        v10 = (3.14159 + v10)
                    case 358:
                        v8 += (fi11 % (((fi11 << ((255) & 63))) | 1))
                    case 1:
                        v13 = (v13 | v13)
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                default:
                    v14 := GenS1{F0: (^(v8)), F1: 239}
                    _ = v14
                }
            }
            _ = v13
        }
        _ = fi11
    }
    v15 := GenS1{F0: (v8 + 32768), F1: -9223372036854775808}
    v10 += 1.5
    v9 -= v9
    var v16 uint8 = 1
    v9 -= v9
    v16 = v16
    var v17 float32 = 1e10
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(v16, v9, v17)
        _ = fi18
    }
    if (v10 <= v10) {
        v19 := GenS2{F0: v2, F1: (v8 & (v8 * -2147483648)), F2: 9223372036854775807}
        v16 = v16
        _ = v19
    }
    v2 = (v2 * (v2 << ((2) & 63)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
