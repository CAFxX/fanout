package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }

func main() {
    var v1 int32 = -891
    v2 := []int{1, 2147483647, 32767}
    v3 := []string{}
    if (1.5 > (1.5 - 0.1)) {
        for fi4 := 0; fi4 < 4; fi4++ {
            switch fi4 {
            case -1:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 1:
                v1 = (v1 >> ((0) & 63))
            case -221:
                switch "foo bar" {
                case "a":
                    v5 := GenS1{F0: (true || false), F1: -1, F2: (fi4 >> ((18446744073709551615) & 63))}
                    _ = v5
                default:
                    v1 = v1
                }
            default:
                var v6 int32 = (v1 + (v1 << ((168) & 63)))
                _ = v6
            }
            fmt.Println(v1, fi4)
            _ = fi4
        }
    } else {
        for fi7 := 0; fi7 < 4; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 float64 = 1.5
                v1 = (-(v1))
                _ = fi8
                _ = v9
            }
            v1 += (v1 * v1)
            _ = fi7
        }
    }
    if true {
        v1 += v1
    }
    v1 += (v1 % (((^(v1))) | 1))
    if false {
        if true {
            v1 = (v1 >> ((54) & 63))
        }
        switch "a" {
        case "\n\t":
            var v10 int32 = 42
            _ = v10
        }
        v1 -= v1
    }
    v1 = v1
    var v11 uint16 = 32767
    v1 = v1
    v11 = (v11 * v11)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
}
