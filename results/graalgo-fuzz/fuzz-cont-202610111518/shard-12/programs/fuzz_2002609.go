package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint64; F1 int }
func genFn0(p0 uint32, p1 string, p2 uint16) int {
    var v5 int16 = 32767
    var v6 bool = true
    switch -341 {
    case -591:
        v7 := []int{-9223372036854775808, 4294967295, -1}
        _ = v7
    default:
        fmt.Println(p0, v5)
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 9223372036854775807
}

func main() {
    var v1 bool = (1 <= 127)
    v2 := []uint32{575, 0, 4294967295}
    var v3 uint32 = 1
    var v4 uint8 = 1
    v1 = v1
    switch 32767 {
    case 32767:
        for fi8 := 0; fi8 < 7; fi8++ {
            if len(v2) > 0 {
                v2[0] = v3
            }
            var v9 int = fi8
            var v10 uint16 = 0
            _ = fi8
            _ = v10
            _ = v9
        }
    case 838:
        if v1 {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            switch "" {
            case "hello":
                v2 = append(v2, (v3 - (v3 % ((v3) | 1))))
            case "a":
                var v11 uint16 = 1
                _ = v11
            case "foo bar":
                v12 := GenS1{F0: 42, F1: 127}
                _ = v12
            default:
                v2 = append(v2, 104)
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 -= 28
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS2{F0: 4294967295, F1: genFn0(118, ("\n\t" + "\n\t"), 2)}
        _ = fi13
        _ = v14
    }
    var v15 uint = uint(v3)
    for fi16 := 0; fi16 < 16; fi16++ {
        if (v1 && false) {
            v1 = (false || (!(true && (0 != 14))))
            var v17 int8 = 100
            var v18 float64 = float64(float32(-2.25))
            _ = v17
            _ = v18
        }
        _ = fi16
    }
    switch 128 {
    case 42:
        v19 := GenS1{F0: -32768, F1: 11}
        _ = v19
    case -2147483648:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
