package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32767, -745, -9223372036854775808}
    v3 := make(map[uint32]uint32)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 := GenS1{F0: 4294967295, F1: 1}
    if ("\n\t" >= (("\n\t" + ("" + ("hello" + ("\n\t" + "")))) + ("\n\t" + ""))) {
        var v5 uint16 = 1
        var v6 uint32 = 0
        _ = v5
        _ = v6
    }
    var v7 int64 = -9223372036854775808
    v7 = 9223372036854775807
    var v8 uint16 = 1
    switch 127 {
    case -329:
        var v9 bool = (18446744073709551615 > 2)
        _ = v9
    case 699:
        switch "foo bar" {
        case "":
            switch "hello" {
            case "\n\t":
                v10 := v1[-537]
                _ = v10
            default:
                v2 = append(v2, 638)
            }
        default:
            _ = len(v2)
        }
    default:
        var v11 bool = false
        _ = v11
    }
    v7 = 613
    v8 = 0
    v7 = v7
    var v12 int64 = (v7 - (-(v7)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
