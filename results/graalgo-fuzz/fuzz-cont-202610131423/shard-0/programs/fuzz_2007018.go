package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 string }

func main() {
    var v1 uint = 4294967295
    var v2 int8 = 22
    var v3 int16 = 668
    var v4 bool = (1 > 123)
    var v5 uint64 = 344
    v4 = v4
    v6 := []int{}
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 6; fi8++ {
            switch fi7 {
            case 1:
                fmt.Println(v2, fi8)
            case 922:
                fmt.Println(fi7, v3, v3)
            case 127:
                v2 = 71
            default:
                v5 += v5
            }
            _ = fi8
        }
        _ = fi7
    }
    v9 := GenS1{F0: 2147483648, F1: 9223372036854775807, F2: ((("foo bar" + ("a" + "")) + "\n\t") + ("" + (("a" + ("hello" + "")) + "a")))}
    var v10 uint = (v1 >> ((v1) & 63))
    switch 9223372036854775807 {
    case 786:
        if len(v6) > 0 {
            v6[0] = 9223372036854775807
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
