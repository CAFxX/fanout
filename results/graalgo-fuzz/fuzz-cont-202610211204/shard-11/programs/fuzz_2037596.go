package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int8 }

func main() {
    var v1 uint32 = 0
    v2 := []string{((("foo bar" + "\n\t") + ("foo bar" + "\n\t")) + ""), "hello", ("a" + "\n\t")}
    v3 := make(map[int]string)
    var v4 uint8 = 0
    v4 = v4
    var v5 uint64 = uint64(v4)
    if len(v2) > 0 {
        v2[0] = (("\n\t" + "foo bar") + "foo bar")
    }
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float32 = ((1e10 - float32(float64((1e10 + 1e10)))) * -2.25)
        _ = fi6
        _ = v7
    }
    _ = len(v3)
    var v8 float32 = (1e10 * 1e10)
    v9 := v3[-926]
    v10 := GenS1{F0: -22, F1: 42, F2: 42}
    switch "hello" {
    case "a":
        for fi11 := 0; fi11 < 1; fi11++ {
            v12 := GenS1{F0: fi11, F1: -945, F2: 10}
            var v13 uint8 = 0
            if len(v2) > 0 {
                v2[0] = v9
            }
            _ = fi11
            _ = v12
            _ = v13
        }
    default:
        fmt.Println(v8, v1)
    }
    if (!false) {
        var v14 int16 = -1
        delete(v3, -477)
        if len(v2) > 0 {
            v2[0] = v9
        }
        _ = v14
    } else {
        if true {
            if len(v2) > 0 {
                v2[0] = ((v9 + (v9 + "a")) + ((v9 + ((v9 + "\n\t") + v9)) + v9))
            }
            v3[42] = v9
            var v15 int32 = 0
            _ = v15
        } else {
            v9 = v9
        }
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 16; fi17++ {
                var v18 uint = 2
                v5 = v5
                switch v9 {
                case "":
                    var v19 int16 = 1
                    _ = v19
                default:
                    switch 2147483647 {
                    case 32767:
                        var v20 int8 = -128
                        _ = v20
                    }
                }
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
    }
    v5 -= (v5 | (v5 + (v5 % ((41) | 1))))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
