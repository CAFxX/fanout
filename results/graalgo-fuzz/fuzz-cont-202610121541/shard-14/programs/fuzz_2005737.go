package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }

func main() {
    var v1 uint32 = 42
    var v2 int16 = -32768
    v3 := []string{"hello", ((("a" + "a") + "a") + "")}
    var v4 uint = 782
    var v5 int16 = (-(v2))
    var v6 int32 = 754
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v4 = (v4 * (v4 * (v4 % ((v4) | 1))))
    v2 += 0
    if false {
        switch "\n\t" {
        case "":
            for fi7 := 0; fi7 < 9; fi7++ {
                var v8 uint32 = v1
                _ = fi7
                _ = v8
            }
        case "foo bar":
            var v9 uint16 = uint16(v4)
            _ = v9
        case "hello":
            if len(v3) > 0 {
                v3[0] = "foo bar"
            }
        default:
            for fi10 := 0; fi10 < 16; fi10++ {
                if false {
                    if false {
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v1 = v1
                            var v12 uint8 = 42
                            v4 = (^(v4))
                            _ = len(v3)
                            _ = fi11
                            _ = v12
                        }
                    }
                } else {
                    var v13 float32 = 0.0
                    v3 = append(v3, "\n\t")
                    _ = v13
                }
                _ = fi10
            }
        }
    }
    for fi14 := 0; fi14 < 17; fi14++ {
        v15 := GenS1{F0: 1, F1: 2}
        _ = fi14
        _ = v15
    }
    var v16 bool = ("\n\t" <= "\n\t")
    v4 -= (v4 * 861)
    fmt.Println(v2, v5)
    v2 = v2
    var v17 uint64 = 18446744073709551615
    switch "foo bar" {
    case "\n\t":
        var v18 float64 = 1e10
        _ = v18
    case "hello":
        v3 = append(v3, (("hello" + "foo bar") + ""))
    }
    var v19 int = int(v6)
    v3 = append(v3, "")
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
