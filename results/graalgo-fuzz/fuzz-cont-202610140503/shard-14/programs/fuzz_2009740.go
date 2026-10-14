package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int8 }

func main() {
    v1 := []string{}
    var v2 uint = 42
    var v3 bool = (float32(((3.14159 + 1.5) - 0.0)) == 0.0)
    v4 := []uint32{}
    var v5 uint8 = 127
    v3 = v3
    switch 153 {
    case -952:
        if false {
            v3 = v3
        }
    case 32768:
        if false {
            switch ("a" + "\n\t") {
            case "a":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            default:
                var v6 uint64 = 296
                _ = v6
            }
            _ = len(v1)
        } else {
            if true {
                var v7 uint64 = 42
                v7 -= (v7 / ((v7) | 1))
                _ = v7
            } else {
                fmt.Println(v2)
            }
        }
    }
    v2 -= 896
    v8 := GenS1{F0: "\n\t", F1: -401, F2: -128}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v9 int16 = -648
    v5 -= v5
    var v10 int8 = 0
    v11 := GenS1{F0: ("hello" + "a"), F1: 128, F2: v10}
    v3 = (-1 <= (v9 >> (((v2 / ((v2) | 1))) & 63)))
    var v12 int8 = (v10 % ((v10) | 1))
    var v13 float32 = (3.14159 - -2.25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
