package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }

func main() {
    var v1 string = "\n\t"
    var v2 uint64 = 994
    var v3 uint64 = (v2 & 2)
    v4 := []string{v1, (v1 + (v1 + v1)), v1, "\n\t"}
    _ = len(v4)
    v3 += v3
    if true {
        v2 = v2
    } else {
        for fi5 := 0; fi5 < 15; fi5++ {
            if (255 == 1) {
                var v6 uint8 = 1
                var v7 int = 2147483647
                if (0 != 597) {
                    v6 = (v6 & 211)
                    var v8 string = "\n\t"
                    _ = v8
                }
                _ = v6
                _ = v7
            }
            _ = fi5
        }
    }
    v3 -= (v3 << ((0) & 63))
    _ = len(v4)
    var v9 uint16 = 2
    for fi10 := 0; fi10 < 6; fi10++ {
        if (!((^(fi10)) < fi10)) {
            v1 = ("a" + ((v1 + v1) + (v1 + v1)))
            for fi11 := 0; fi11 < 20; fi11++ {
                var v12 string = ("hello" + v1)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi11
                _ = v12
            }
        } else {
            v4 = append(v4, (((v1 + "\n\t") + v1) + v1))
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi10
    }
    var v13 float32 = 1e10
    fmt.Println(v9, v13, v1)
    var v14 int64 = 9223372036854775807
    fmt.Println(v2, v2)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (42 == -1) {
        var v15 uint = uint(v9)
        _ = v15
    } else {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v13 -= ((float32((3.14159 * -2.25)) / ((v13 + v13) * (((v13 / 1.5) / 3.14159) * v13))) / v13)
        v14 -= -582
    }
    switch -9223372036854775808 {
    case -173:
        v16 := GenS1{F0: 54, F1: 127}
        _ = v16
    case 119:
        v1 = "a"
    }
    fmt.Println(v9, v2)
    v13 = float32((1e10 / 1.5))
    var v17 uint8 = 127
    for fi18 := 0; fi18 < 2; fi18++ {
        fmt.Println(v2, v17)
        v2 -= (v3 * (v3 | 602))
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
