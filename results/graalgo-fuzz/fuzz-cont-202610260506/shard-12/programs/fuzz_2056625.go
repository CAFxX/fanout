package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64 }

func main() {
    var v1 int = 611
    v2 := []uint32{}
    var v3 uint32 = 544
    v4 := []string{}
    var v5 uint16 = 32767
    _ = len(v4)
    for fi6 := 0; fi6 < 20; fi6++ {
        var v7 int16 = 32767
        _ = len(v4)
        var v8 uint64 = 536
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 int8 = 0
    if false {
        v10 := GenS1{F0: 487, F1: -351}
        var v11 bool = true
        _ = v10
        _ = v11
    } else {
        if false {
            v12 := GenS1{F0: 323, F1: int32(v1)}
            var v13 bool = (-2147483648 != 2147483647)
            _ = v12
            _ = v13
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            if (0 <= (v5 * (v5 / ((v5) | 1)))) {
                fmt.Println(v5, v5, v9)
            } else {
                switch "foo bar" {
                case "a":
                    v5 += 65535
                case "hello":
                    var v14 float64 = (1.5 / float64(float32(float64(float32(((1.5 * (float64(0.0) * 3.14159)) + 3.14159))))))
                    _ = v14
                case "foo bar":
                    for fi15 := 0; fi15 < 5; fi15++ {
                        v3 += 543
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if true {
                            v16 := GenS1{F0: 18446744073709551615, F1: 2147483647}
                            v5 = 1
                            _ = v16
                        }
                        _ = fi15
                    }
                }
            }
        }
        v17 := GenS1{F0: 18446744073709551615, F1: 1}
        _ = v17
    }
    v18 := GenS2{F0: 825}
    switch "foo bar" {
    case "\n\t":
        var v19 int = (v1 + v1)
        _ = v19
    case "foo bar":
        v9 = 0
    }
    if true {
        if false {
            v20 := GenS1{F0: 1, F1: 42}
            _ = v20
        }
        fmt.Println(v9, v9)
    }
    var v21 float64 = 0.1
    var v22 int8 = (v9 ^ (v9 % ((42) | 1)))
    var v23 uint = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
