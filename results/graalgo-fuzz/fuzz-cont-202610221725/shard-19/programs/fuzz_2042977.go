package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 bool }
func genFn0() uint64 {
    v7 := GenS1{F0: -1, F1: 4294967295}
    v8 := []int{1, 127, 0, -2147483648}
    if len(v8) > 0 {
        v8[0] = 2147483647
    }
    _ = v7
    _ = v8
    return 284
}

func main() {
    v1 := []int{42, 5, 32768, -450}
    var v2 int16 = 55
    var v3 uint8 = 255
    var v4 uint8 = (v3 & (v3 - (v3 >> ((255) & 63))))
    var v5 int = -2147483648
    var v6 int32 = 757
    var v9 float64 = (((((float64(1.5) * float64(-2.25)) + -2.25) * (((1e10 / 0.1) + -2.25) + 3.14159)) + 0.0) / (3.14159 - 0.0))
    var v10 uint = 112
    var v11 bool = false
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v9 = v9
    switch (v5 & v5) {
    case 0:
        var v12 uint64 = 1
        _ = v12
    case -9223372036854775808:
        if true {
            v6 -= -418
            var v13 bool = (!v11)
            if (42 != v5) {
                if len(v1) > 0 {
                    v1[0] = int(v10)
                }
            } else {
                if true {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        v15 := GenS2{F0: false}
                        _ = fi14
                        _ = v15
                    }
                } else {
                    switch "a" {
                    case "hello":
                        var v16 uint16 = 357
                        _ = v16
                    }
                    v1 = append(v1, 2147483647)
                    if true {
                        v6 = (v6 ^ v6)
                        switch "foo bar" {
                        case "foo bar":
                            v17 := GenS2{F0: true}
                            _ = v17
                        case "":
                            fmt.Println(v6)
                        case "\n\t":
                            v9 -= (0.1 * -2.25)
                        }
                        var v18 float32 = ((0.0 + 0.1) * (float32(-2.25) / 1.5))
                        _ = v18
                    }
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = len(v1)
            }
            _ = v13
        } else {
            if ((v2 ^ v2) != (v2 / ((v2) | 1))) {
                switch ("\n\t" + ("\n\t" + ("\n\t" + "foo bar"))) {
                case "\n\t":
                    v2 = v2
                case "foo bar":
                    v4 -= uint8(v10)
                default:
                    var v19 int8 = 42
                    _ = v19
                }
                fmt.Println(v4, v3)
                switch "hello" {
                case "\n\t":
                    v20 := GenS2{F0: (("a" + (("\n\t" + ("\n\t" + "\n\t")) + "a")) > "a")}
                    _ = v20
                case "a":
                    v21 := GenS1{F0: (v5 * -447), F1: 2}
                    _ = v21
                case "hello":
                    for fi22 := 0; fi22 < 13; fi22++ {
                        v3 = v4
                        _ = fi22
                    }
                }
            }
        }
    }
    if ((("" + "\n\t") <= "") || (true || (v11 || ((1.5 < v9) || (1e10 <= 1e10))))) {
        var v23 int64 = 92
        v1 = append(v1, (v5 << (((v10 << ((42) & 63))) & 63)))
        v24 := GenS1{F0: 188, F1: 18446744073709551615}
        _ = v23
        _ = v24
    } else {
        var v25 int32 = (^(v6))
        _ = v25
    }
    fmt.Println(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
