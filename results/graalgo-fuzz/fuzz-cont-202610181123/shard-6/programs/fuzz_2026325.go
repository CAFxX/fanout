package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }

func main() {
    var v1 uint = 612
    var v2 int8 = 1
    v3 := []string{"\n\t", ("\n\t" + ""), ("foo bar" + "a")}
    v2 = v2
    var v4 uint16 = 32767
    switch 127 {
    case 0:
        var v5 uint = (v1 - 0)
        _ = v5
    case -9223372036854775808:
        v2 += v2
    default:
        if len(v3) > 0 {
            v3[0] = ""
        }
    }
    if len(v3) > 0 {
        v3[0] = "foo bar"
    }
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            if true {
                v2 -= v2
                fmt.Println(v1, v1, v2)
                v4 += 469
            } else {
                v1 = 18446744073709551615
                var v8 uint16 = (v4 * (v4 ^ (v4 * 806)))
                _ = v8
            }
            _ = fi7
        }
        _ = fi6
    }
    v9 := GenS1{F0: -637, F1: 42}
    v4 -= (^(v4))
    v4 = (^(v4))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
