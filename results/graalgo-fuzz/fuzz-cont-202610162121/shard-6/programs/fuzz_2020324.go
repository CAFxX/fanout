package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int }
type GenS2 struct { F0 int }

func main() {
    var v1 float32 = ((-2.25 / 3.14159) + float32(-2.25))
    var v2 uint32 = 0
    var v3 uint32 = 4294967295
    var v4 int64 = -9223372036854775808
    switch 32768 {
    case 269:
        var v5 uint64 = 2
        _ = v5
    case 735:
        var v6 string = "\n\t"
        _ = v6
    case 32768:
        v7 := []int{128, 32767, 0, -9223372036854775808}
        _ = v7
    default:
        v8 := GenS1{F0: ("hello" + ((("" + "hello") + "\n\t") + (("a" + "\n\t") + ("foo bar" + "")))), F1: 866, F2: int(v3)}
        _ = v8
    }
    v1 = v1
    v1 -= v1
    v9 := []int{-438, -398, 32768, -192}
    var v10 float32 = 1e10
    v10 -= v1
    fmt.Println(v1)
    if false {
        v11 := GenS1{F0: "\n\t", F1: 0, F2: 1}
        switch "hello" {
        case "a":
            for fi12 := 0; fi12 < 15; fi12++ {
                if false {
                    fmt.Println(v4, v2)
                    for fi13 := 0; fi13 < 14; fi13++ {
                        var v14 uint = 42
                        var v15 uint64 = 0
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                } else {
                    var v16 int64 = v4
                    _ = v16
                }
                switch 742 {
                case 127:
                    if true {
                        v3 = (v3 % (((^(v2))) | 1))
                        _ = len(v9)
                    } else {
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        v17 := GenS2{F0: fi12}
                        v18 := GenS1{F0: ("foo bar" + "a"), F1: 42, F2: fi12}
                        _ = v17
                        _ = v18
                    }
                }
                var v19 bool = true
                _ = fi12
                _ = v19
            }
        default:
            if (!true) {
                var v20 int = 32767
                _ = v20
            } else {
                fmt.Println(v1, v4, v1)
                v2 -= v2
            }
        }
        _ = v11
    } else {
        fmt.Println(v4, v3, v10)
    }
    v2 -= v2
    var v21 int = 806
    fmt.Println(v21, v21, v4)
    v2 -= 0
    var v22 int = (-(v21))
    v4 -= -299
    v10 += (v1 / -2.25)
    fmt.Println(v1, v10, v1)
    if ((4294967295 <= 0) && true) {
        v23 := GenS2{F0: v22}
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
