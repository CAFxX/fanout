package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }

func main() {
    var v1 int = 42
    v2 := []int{(v1 >> ((255) & 63))}
    var v3 uint64 = 2
    var v4 int64 = int64(v1)
    v5 := GenS1{F0: 18446744073709551615, F1: 255}
    v4 = v4
    for fi6 := 0; fi6 < 17; fi6++ {
        v1 += (-(v1))
        _ = fi6
    }
    if ("hello" >= ("hello" + "hello")) {
        if false {
            var v7 int16 = 977
            _ = v7
        } else {
            switch "hello" {
            case "a":
                v1 = (v1 / ((v1) | 1))
            }
        }
    }
    v3 -= (^(v3))
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 13; fi9++ {
            v1 += fi8
            _ = fi9
        }
        v3 += (v3 ^ (v3 & (^(v3))))
        _ = len(v2)
        _ = fi8
    }
    fmt.Println(v4)
    v1 = v1
    var v10 int16 = 1
    switch "" {
    case "foo bar":
        switch ("foo bar" + (("\n\t" + ("a" + "\n\t")) + "")) {
        case "\n\t":
            v10 -= int16(v1)
        case "":
            switch "hello" {
            case "":
                for fi11 := 0; fi11 < 8; fi11++ {
                    v10 = (v10 << ((42) & 63))
                    for fi12 := 0; fi12 < 6; fi12++ {
                        var v13 bool = (!(v3 != 18446744073709551615))
                        v14 := GenS1{F0: 110, F1: 42}
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 uint64 = v3
                    _ = fi11
                    _ = v15
                }
            default:
                var v16 uint16 = 255
                _ = v16
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
