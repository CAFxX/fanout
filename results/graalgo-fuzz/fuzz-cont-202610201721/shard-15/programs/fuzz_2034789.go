package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }

func main() {
    var v1 uint8 = 255
    var v2 uint = 255
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    v5 := make(map[string]uint32)
    var v6 uint64 = 42
    for fi7 := 0; fi7 < 5; fi7++ {
        switch "" {
        case "a":
            for fi8 := 0; fi8 < 15; fi8++ {
                switch -9223372036854775808 {
                case 0:
                    var v9 uint8 = (v1 % ((255) | 1))
                    _ = v9
                default:
                    var v10 float64 = float64(float32(float64(3.14159)))
                    _ = v10
                }
                v6 = (^(v6))
                v11 := []int{32767, (fi8 - 9223372036854775807), 32767}
                _ = fi8
                _ = v11
            }
        case "hello":
            var v12 float32 = -2.25
            _ = v12
        default:
            if (0 > 434) {
                v13 := GenS1{F0: (fi7 << ((962) & 63)), F1: 752}
                fmt.Println(v1, v1, v1)
                v2 -= (v2 >> ((255) & 63))
                _ = v13
            } else {
                fmt.Println(v2)
            }
        }
        if true {
            v1 += (v1 ^ (^(v1)))
            var v14 uint = 4294967295
            v6 = 18446744073709551615
            _ = v14
        }
        _ = fi7
    }
    var v15 uint8 = v1
    var v16 string = ("foo bar" + "\n\t")
    for fi17 := 0; fi17 < 5; fi17++ {
        v18 := GenS1{F0: int(v2), F1: 42}
        fmt.Println(v2)
        _ = fi17
        _ = v18
    }
    var v19 uint = (v2 | (v2 / ((v2) | 1)))
    fmt.Println(v16, v6, v15)
    switch "hello" {
    case "\n\t":
        switch -9223372036854775808 {
        case -397:
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 bool = (!(v1 == (^(v1))))
                _ = fi20
                _ = v21
            }
        case 1:
            var v22 uint32 = 65535
            _ = v22
        }
    case "a":
        v16 = v16
    }
    var v23 int = 42
    v16 = "hello"
    _ = len(v3)
    v24 := GenS1{F0: 0, F1: 32767}
    var v25 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
