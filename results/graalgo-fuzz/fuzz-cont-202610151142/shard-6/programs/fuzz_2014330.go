package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS2 struct { F0 int; F1 string; F2 int64 }

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{}
    v3 := []uint32{}
    var v4 bool = true
    _ = len(v2)
    if false {
        v4 = (!(!v4))
    } else {
        if v4 {
            var v5 int8 = 100
            _ = v5
        } else {
            var v6 uint64 = 1
            if (-2.25 <= 1.5) {
                switch "hello" {
                case "hello":
                    switch "\n\t" {
                    case "a":
                        v4 = (65535 > 4294967295)
                    case "foo bar":
                        v2 = append(v2, "hello")
                    }
                case "a":
                    fmt.Println(v6)
                }
                for fi7 := 0; fi7 < 12; fi7++ {
                    var v8 int16 = -1
                    var v9 int8 = 42
                    for fi10 := 0; fi10 < 15; fi10++ {
                        v8 -= (v8 / ((14) | 1))
                        var v11 uint16 = 255
                        v12 := GenS1{F0: 928, F1: 628, F2: (v6 % ((42) | 1))}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                v4 = false
            } else {
                v4 = v4
                if (2147483648 <= 799) {
                    v13 := GenS1{F0: -718, F1: 741, F2: v6}
                    var v14 int32 = 42
                    _ = v13
                    _ = v14
                }
                v4 = false
            }
            _ = v6
        }
    }
    if v4 {
        v15 := GenS2{F0: 127, F1: "", F2: 354}
        _ = v15
    }
    v16 := GenS2{F0: -605, F1: "", F2: 1}
    var v17 uint16 = 2
    _ = len(v1)
    var v18 uint64 = 0
    switch 32767 {
    case 128:
        var v19 uint = 1
        _ = v19
    default:
        for fi20 := 0; fi20 < 1; fi20++ {
            var v21 int8 = 1
            switch fi20 {
            case -2147483648:
                v22 := GenS2{F0: (fi20 / ((32767) | 1)), F1: (("\n\t" + (("a" + "foo bar") + "hello")) + "hello"), F2: 9223372036854775807}
                _ = v22
            case -9223372036854775808:
                _ = len(v3)
            case 0:
                fmt.Println(v17)
            }
            _ = fi20
            _ = v21
        }
    }
    var v23 float64 = 0.1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
