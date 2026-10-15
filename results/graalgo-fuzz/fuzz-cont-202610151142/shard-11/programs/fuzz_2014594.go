package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool; F1 bool; F2 uint8 }

func main() {
    var v1 int = 32767
    v2 := make(map[string]int)
    var v3 bool = true
    v4 := []int{v1, -9223372036854775808, (v1 ^ (v1 ^ (v1 ^ v1))), v1}
    for fi5 := 0; fi5 < 20; fi5++ {
        if v3 {
            _ = len(v2)
            if (v3 && (32767 >= 2)) {
                v1 -= (v1 << ((26) & 63))
            }
            v1 = 127
        }
        v6 := GenS1{F0: 0}
        _ = fi5
        _ = v6
    }
    v3 = (!(v3 && true))
    if v3 {
        v7 := GenS1{F0: 606}
        v3 = v3
        switch "" {
        case "a":
            fmt.Println(v1, v1)
        case "hello":
            v1 -= v1
        case "":
            var v8 int = (v1 % (((v1 >> ((1) & 63))) | 1))
            _ = v8
        default:
            if len(v4) > 0 {
                v4[0] = (^(v1))
            }
        }
        _ = v7
    }
    var v9 float64 = 0.0
    switch (v1 ^ v1) {
    case 0:
        var v10 int = v1
        _ = v10
    case 465:
        if false {
            _ = len(v2)
            v3 = false
        } else {
            for fi11 := 0; fi11 < 5; fi11++ {
                for fi12 := 0; fi12 < 14; fi12++ {
                    fmt.Println(v3, v1, fi12)
                    var v13 string = "\n\t"
                    _ = fi12
                    _ = v13
                }
                v14 := GenS1{F0: 444}
                v3 = false
                _ = fi11
                _ = v14
            }
            fmt.Println(v3)
        }
    case -962:
        var v15 bool = v3
        _ = v15
    default:
        var v16 int8 = 40
        _ = v16
    }
    v9 -= ((v9 - (3.14159 - v9)) / ((0.0 - 1e10) * (((v9 * ((float64(0.0) - 1e10) + 0.0)) * 1.5) - (float64((float32(v9) * 3.14159)) / (v9 / 1.0)))))
    v17 := GenS2{F0: (!((!v3) && v3)), F1: true, F2: 184}
    switch "foo bar" {
    case "a":
        v4 = append(v4, (v1 / ((2147483647) | 1)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
