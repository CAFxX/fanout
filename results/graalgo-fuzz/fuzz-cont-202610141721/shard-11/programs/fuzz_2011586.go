package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint8 = 195
    var v2 int8 = -1
    var v3 int32 = -692
    if true {
        if true {
            v1 -= v1
            v4 := []int{32768, 2147483647, 42, -1}
            _ = v4
        } else {
            var v5 string = ("\n\t" + "hello")
            v2 += (v2 ^ (-(v2)))
            _ = v5
        }
        switch "a" {
        case "a":
            for fi6 := 0; fi6 < 9; fi6++ {
                v2 = v2
                v2 = v2
                v7 := GenS1{F0: (!((v1 | v1) >= v1))}
                _ = fi6
                _ = v7
            }
        }
        v3 += 2147483647
    }
    switch 32767 {
    case -1:
        if (!true) {
            fmt.Println(v3, v2)
            switch "" {
            case "foo bar":
                v8 := []int{4294967295, 9223372036854775807}
                _ = v8
            case "a":
                v3 = 2147483647
            }
        } else {
            fmt.Println(v1, v3)
            var v9 int16 = 42
            _ = v9
        }
    case 4294967295:
        v10 := GenS1{F0: (v3 < int32(v2))}
        _ = v10
    }
    var v11 uint8 = v1
    var v12 float32 = (3.14159 * float32((-2.25 - (float64(1e10) + 0.0))))
    if true {
        v2 += (v2 ^ v2)
    } else {
        v3 -= v3
        for fi13 := 0; fi13 < 6; fi13++ {
            v1 -= (v1 + v11)
            _ = fi13
        }
        v1 -= 2
    }
    v1 = (v11 / ((2) | 1))
    v14 := GenS1{F0: (((!((v1 ^ 0) != v1)) && (!true)) && false)}
    fmt.Println(v3, v11)
    fmt.Println(v3)
    var v15 uint8 = (v1 % ((127) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
