package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"\n\t", ("a" + ("\n\t" + ("a" + "a")))}
    var v3 int = 9223372036854775807
    v4 := []uint32{2147483647, 846, 1, 273}
    var v5 uint32 = 2147483647
    v5 += v5
    delete(v1, "hello")
    for fi6 := 0; fi6 < 12; fi6++ {
        if false {
            var v7 uint32 = v5
            for fi8 := 0; fi8 < 7; fi8++ {
                v9 := GenS1{F0: 42, F1: 42}
                v5 = 65535
                v10 := GenS1{F0: 351, F1: 1}
                switch fi6 {
                case 4294967295:
                    v5 = v5
                default:
                    v3 = (v3 / (((fi6 * (-(v3)))) | 1))
                }
                _ = fi8
                _ = v10
                _ = v9
            }
            var v11 uint32 = (v5 ^ v5)
            _ = v11
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(v5, v3)
    var v12 int = v3
    fmt.Println(v12)
    switch "hello" {
    case "a":
        v2 = append(v2, "\n\t")
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        v14 := GenS1{F0: 1, F1: -1}
        v12 -= v12
        _ = fi13
        _ = v14
    }
    var v15 float64 = (1e10 - 3.14159)
    var v16 int32 = 42
    v5 = (^(v5))
    if true {
        if (!false) {
            v15 = (v15 + (1e10 * v15))
            switch "" {
            case "hello":
                v16 = (v16 ^ v16)
            default:
                v5 = v5
            }
            _ = len(v4)
        } else {
            fmt.Println(v16, v12, v5)
            for fi17 := 0; fi17 < 10; fi17++ {
                switch "hello" {
                case "foo bar":
                    v5 += v5
                case "a":
                    v12 = (v3 - 128)
                default:
                    for fi18 := 0; fi18 < 11; fi18++ {
                        var v19 uint = 649
                        _ = fi18
                        _ = v19
                    }
                }
                _ = fi17
            }
            v5 -= 2147483647
        }
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            fmt.Println(v5)
            _ = fi20
        }
        v21 := v1["foo bar"]
        var v22 int8 = 102
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
