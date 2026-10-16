package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 string }
type GenS2 struct { F0 int; F1 int }

func main() {
    var v1 uint16 = 264
    v2 := []uint32{42, 65535}
    var v3 float32 = float32((-2.25 + 1e10))
    var v4 uint8 = 42
    if len(v2) > 0 {
        v2[0] = 547
    }
    var v5 int = -9223372036854775808
    fmt.Println(v3, v3)
    v6 := GenS1{F0: 2, F1: 200, F2: "hello"}
    switch "a" {
    case "foo bar":
        var v7 float64 = 0.0
        _ = v7
    case "":
        switch "\n\t" {
        case "":
            v2 = append(v2, 65535)
        default:
            if true {
                var v8 int16 = -598
                for fi9 := 0; fi9 < 13; fi9++ {
                    v10 := GenS1{F0: 18446744073709551615, F1: 129, F2: ("foo bar" + "\n\t")}
                    v3 = 3.14159
                    v5 += v5
                    _ = fi9
                    _ = v10
                }
                _ = v8
            }
        }
    case "hello":
        v5 += (v5 * 128)
    }
    v11 := GenS2{F0: (v5 - 32768), F1: 128}
    var v12 string = ("\n\t" + (("a" + ("foo bar" + "")) + ("\n\t" + "foo bar")))
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    switch 32768 {
    case -9223372036854775808:
        v2 = append(v2, 901)
    case 36:
        if len(v2) > 0 {
            v2[0] = 753
        }
    }
    v4 += v4
    var v13 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
