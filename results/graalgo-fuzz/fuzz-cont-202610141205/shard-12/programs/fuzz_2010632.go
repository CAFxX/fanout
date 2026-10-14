package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 uint = 42
    var v4 int = -809
    if (18 == -121) {
        switch "" {
        case "a":
            var v5 float64 = 1e10
            _ = v5
        case "":
            for fi6 := 0; fi6 < 17; fi6++ {
                v4 += -565
                switch "\n\t" {
                case "hello":
                    v3 += 1
                default:
                    v3 -= 769
                }
                _ = fi6
            }
        }
        v3 -= v3
    } else {
        v4 += v4
        var v7 int8 = -128
        if (65535 != 793) {
            v8 := []int{}
            v4 = v4
            _ = v8
        }
        _ = v7
    }
    var v9 uint32 = 4294967295
    v10 := []int{}
    for fi11 := 0; fi11 < 8; fi11++ {
        v9 += (v9 >> ((v3) & 63))
        _ = fi11
    }
    var v12 uint = (v3 ^ v3)
    for fi13 := 0; fi13 < 2; fi13++ {
        if true {
            v14 := GenS1{F0: v3}
            v15 := GenS1{F0: uint(v9)}
            _ = v14
            _ = v15
        }
        _ = fi13
    }
    _ = len(v10)
    delete(v1, (^(v4)))
    if (("a" + ("hello" + "hello")) <= ("\n\t" + ("foo bar" + "a"))) {
        for fi16 := 0; fi16 < 10; fi16++ {
            if false {
                v17 := GenS2{F0: "\n\t"}
                var v18 bool = (v9 != (v9 % ((2) | 1)))
                _ = len(v10)
                _ = v17
                _ = v18
            }
            _ = fi16
        }
    }
    v1[-63] = ((v12 >> ((v12) & 63)) == (v12 << ((v3) & 63)))
    v19 := GenS3{F0: 32767, F1: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
