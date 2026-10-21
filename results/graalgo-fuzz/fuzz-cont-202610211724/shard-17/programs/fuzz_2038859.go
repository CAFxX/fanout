package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
func genFn0(p0 uint16, p1 bool, p2 int16) uint32 {
    fmt.Println(p2)
    v4 := GenS1{F0: 169, F1: 138}
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 4294967295
}

func main() {
    var v1 bool = true
    v2 := make(map[string]uint32)
    v3 := []uint32{1, 574, 646}
    var v5 int16 = 0
    switch "foo bar" {
    case "\n\t":
        v6 := GenS1{F0: 1, F1: 42}
        _ = v6
    case "foo bar":
        switch 2147483647 {
        case -9223372036854775808:
            v5 += v5
        case 4294967295:
            if false {
                v7 := v2["hello"]
                _ = v7
            }
        case -248:
            v8 := v2["hello"]
            _ = v8
        default:
            fmt.Println(v5, v1)
        }
    case "hello":
        for fi9 := 0; fi9 < 8; fi9++ {
            for fi10 := 0; fi10 < 13; fi10++ {
                var v11 uint32 = 2147483648
                _ = fi10
                _ = v11
            }
            if (-2.25 != 1.5) {
                _ = len(v3)
            }
            _ = fi9
        }
    default:
        switch 127 {
        case 127:
            v12 := GenS1{F0: 0, F1: 276}
            _ = v12
        case 341:
            var v13 string = ("foo bar" + ("a" + (("\n\t" + "hello") + "a")))
            _ = v13
        }
    }
    v5 = 1
    if v1 {
        v5 += (^(v5))
    }
    v1 = false
    var v14 int = 399
    var v15 uint16 = 255
    v16 := GenS1{F0: 4294967295, F1: 289}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v5
}
