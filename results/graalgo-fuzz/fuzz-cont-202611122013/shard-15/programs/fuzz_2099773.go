package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int16 }
type GenS2 struct { F0 uint8; F1 int }

func main() {
    var v1 int8 = 79
    v2 := []string{"a", "foo bar", "\n\t", "\n\t"}
    var v3 uint = 2
    var v4 int = -156
    var v5 int8 = (v1 | 42)
    if (!false) {
        switch "foo bar" {
        case "":
            v3 = v3
        case "a":
            var v6 float64 = 1.5
            _ = v6
        default:
            var v7 string = "\n\t"
            _ = v7
        }
    }
    fmt.Println(v5)
    v1 = 0
    var v8 string = "foo bar"
    _ = len(v2)
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 float64 = 0.0
        v11 := GenS1{F0: (42 >= fi9), F1: 9223372036854775807, F2: 42}
        if ((fi9 != (v4 + fi9)) || (!false)) {
            if false {
                switch (v4 / ((127) | 1)) {
                case 2147483647:
                    var v12 string = ("hello" + ("\n\t" + v8))
                    _ = v12
                case -728:
                    fmt.Println(v8)
                }
                var v13 int64 = 9223372036854775807
                _ = v13
            } else {
                var v14 uint16 = 984
                v8 = (v8 + "")
                switch (fi9 | (^(v4))) {
                case -940:
                    v3 += (v3 & (v3 >> ((v3) & 63)))
                case 60:
                    fmt.Println(v10)
                case -261:
                    if false {
                        for fi15 := 0; fi15 < 13; fi15++ {
                            v5 += 100
                            var v16 float32 = (((1.5 / 3.14159) + ((-2.25 - -2.25) / 1.0)) / 0.1)
                            _ = fi15
                            _ = v16
                        }
                        if false {
                            v17 := GenS2{F0: 243, F1: (v4 / (((-(v4))) | 1))}
                            var v18 string = "hello"
                            _ = v17
                            _ = v18
                        }
                        var v19 uint64 = 113
                        _ = v19
                    } else {
                        if false {
                            var v20 uint16 = (v14 - v14)
                            _ = v20
                        }
                        v21 := GenS1{F0: true, F1: 491, F2: 0}
                        if len(v2) > 0 {
                            v2[0] = v8
                        }
                        _ = v21
                    }
                }
                _ = v14
            }
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    v1 -= v5
    v1 = v5
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
