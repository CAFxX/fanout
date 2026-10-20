package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 int64 }

func main() {
    var v1 uint32 = 42
    v2 := []int{-299, 2147483647}
    var v3 int = -215
    var v4 int64 = -9223372036854775808
    var v5 int = (v3 ^ v3)
    fmt.Println(v1, v5, v4)
    fmt.Println(v5, v1, v3)
    var v6 int32 = 845
    fmt.Println(v1, v5, v6)
    v3 = v3
    for fi7 := 0; fi7 < 14; fi7++ {
        if (-1 <= int8(v4)) {
            fmt.Println(v5, v1, fi7)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v8 bool = false
            _ = v8
        }
        var v9 bool = (-431 <= -761)
        switch "hello" {
        case "\n\t":
            v9 = (462 == 338)
        default:
            switch -1 {
            case 4294967295:
                v10 := GenS1{F0: 32767}
                _ = v10
            default:
                fmt.Println(v3, v9, v1)
            }
        }
        _ = fi7
        _ = v9
    }
    if false {
        v11 := GenS2{F0: 104, F1: (-(v4))}
        _ = v11
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
