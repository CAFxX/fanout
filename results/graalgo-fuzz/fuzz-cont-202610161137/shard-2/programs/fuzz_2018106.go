package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int }
type GenS2 struct { F0 uint64; F1 int64 }
func genFn0() int16 {
    v4 := []int{2147483647, 0}
    v5 := GenS1{F0: -90, F1: "foo bar", F2: 128}
    var v6 int8 = 42
    _ = v4
    _ = v5
    _ = v6
    return -927
}

func main() {
    var v1 int32 = 1
    var v2 uint8 = 128
    var v3 string = "\n\t"
    var v7 int8 = 0
    v8 := []int{-448, -684}
    var v9 int = 260
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v10 uint64 = 568
    var v11 uint = 0
    v11 = 4294967295
    for fi12 := 0; fi12 < 3; fi12++ {
        var v13 int = 9223372036854775807
        _ = fi12
        _ = v13
    }
    if false {
        fmt.Println(v1, v10)
        if (362 != -291) {
            v1 -= -917
            for fi14 := 0; fi14 < 9; fi14++ {
                switch v3 {
                case "\n\t":
                    var v15 int32 = (^(v1))
                    _ = v15
                case "a":
                    v7 += -1
                }
                switch v3 {
                case "foo bar":
                    var v16 int64 = 9223372036854775807
                    _ = v16
                case "a":
                    var v17 string = ""
                    _ = v17
                case "":
                    v10 += (v10 / ((2) | 1))
                default:
                    var v18 string = "hello"
                    _ = v18
                }
                v19 := GenS2{F0: (^(v10)), F1: 0}
                _ = fi14
                _ = v19
            }
        }
    } else {
        if ((1.5 < 0.0) && (!false)) {
            v9 = (-(v9))
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
