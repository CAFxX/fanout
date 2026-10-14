package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 int }
type GenS2 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint64 = 456
    var v2 int = 0
    v3 := []int{(v2 ^ (v2 & 32768)), v2, (^(v2))}
    var v4 int = (v2 & -919)
    var v5 uint8 = 127
    var v6 float64 = 1.5
    switch "foo bar" {
    case "foo bar":
        for fi7 := 0; fi7 < 16; fi7++ {
            v1 -= v1
            _ = fi7
        }
    case "\n\t":
        for fi8 := 0; fi8 < 5; fi8++ {
            v2 += v2
            if (42 > 1000) {
                var v9 int8 = 127
                for fi10 := 0; fi10 < 13; fi10++ {
                    var v11 float64 = v6
                    _ = fi10
                    _ = v11
                }
                _ = v9
            } else {
                v4 = 1
                v2 -= (fi8 + (v2 * 32768))
            }
            _ = fi8
        }
    default:
        switch "a" {
        case "hello":
            switch ((("a" + "") + "foo bar") + ("foo bar" + ((((("hello" + "a") + "\n\t") + "a") + "a") + (("\n\t" + "\n\t") + "")))) {
            case "hello":
                var v12 uint32 = 526
                _ = v12
            case "":
                for fi13 := 0; fi13 < 19; fi13++ {
                    fmt.Println(v2, fi13, v5)
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v14 := GenS1{F0: -1, F1: 265, F2: v4}
                    _ = fi13
                    _ = v14
                }
            }
        }
    }
    v15 := GenS1{F0: int8(v5), F1: 605, F2: (v4 ^ (v4 * -1))}
    if false {
        var v16 uint = 18446744073709551615
        fmt.Println(v16, v5, v16)
        _ = v16
    }
    fmt.Println(v4, v1)
    v5 = uint8(v1)
    if true {
        v5 = v5
        fmt.Println(v5)
        fmt.Println(v1, v2)
    }
    v4 = v2
    fmt.Println(v4)
    v17 := GenS2{F0: v1, F1: 554}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
