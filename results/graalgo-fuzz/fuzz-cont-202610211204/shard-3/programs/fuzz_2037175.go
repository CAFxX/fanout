package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 string; F1 int8 }

func main() {
    v1 := []int{1, 42, -2147483648}
    v2 := make(map[string]string)
    v3 := make(map[int]string)
    v4 := make(map[uint32]string)
    var v5 float32 = 0.1
    _ = len(v4)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 7; fi7++ {
            if false {
                _ = len(v1)
            }
            _ = fi7
        }
        var v8 uint64 = 808
        _ = fi6
        _ = v8
    }
    v5 = v5
    v5 = (v5 - 3.14159)
    var v9 int8 = 42
    fmt.Println(v5)
    for fi10 := 0; fi10 < 15; fi10++ {
        v9 -= v9
        fmt.Println(v5)
        if false {
            v11 := GenS1{F0: 1, F1: 128}
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 int = fi12
                v4[2147483647] = "foo bar"
                _ = fi12
                _ = v13
            }
            switch "hello" {
            case "\n\t":
                fmt.Println(v5, v9, v5)
            case "":
                switch 270 {
                case 0:
                    var v14 int16 = 1000
                    _ = v14
                default:
                    switch (fi10 + (fi10 % ((32767) | 1))) {
                    case 32767:
                        var v15 int64 = 684
                        _ = v15
                    }
                }
            }
            _ = v11
        }
        _ = fi10
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        for fi17 := 0; fi17 < 17; fi17++ {
            switch "\n\t" {
            case "a":
                switch "" {
                case "hello":
                    v18 := GenS1{F0: 18446744073709551615, F1: 128}
                    _ = v18
                case "":
                    if true {
                        var v19 int32 = 1
                        _ = v19
                    }
                case "foo bar":
                    v9 -= -1
                }
            default:
                switch "\n\t" {
                case "":
                    v5 = (v5 / ((0.1 / 3.14159) / (1.5 - 0.1)))
                case "a":
                    if len(v1) > 0 {
                        v1[0] = fi16
                    }
                case "hello":
                    v9 = v9
                default:
                    var v20 uint = 255
                    _ = v20
                }
            }
            if false {
                var v21 uint = 442
                var v22 int32 = -333
                fmt.Println(v5, v21, v22)
                _ = v21
                _ = v22
            }
            _ = fi17
        }
        v23 := GenS2{F0: (("a" + "a") + ""), F1: v9}
        _ = fi16
        _ = v23
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
