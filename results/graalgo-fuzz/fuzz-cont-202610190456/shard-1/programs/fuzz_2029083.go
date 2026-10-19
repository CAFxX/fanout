package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int64 }

func main() {
    var v1 uint = 0
    var v2 int32 = -776
    var v3 int16 = 1000
    switch -9223372036854775808 {
    case -945:
        v4 := []int{}
        _ = v4
    case 128:
        v5 := []int{42}
        _ = v5
    case -9223372036854775808:
        var v6 uint = v1
        _ = v6
    default:
        v3 = (v3 | v3)
    }
    var v7 uint16 = 912
    var v8 uint8 = 1
    for fi9 := 0; fi9 < 12; fi9++ {
        v3 -= (^(v3))
        if (2147483647 > 774) {
            for fi10 := 0; fi10 < 6; fi10++ {
                if (false || (18446744073709551615 < 4294967295)) {
                    var v11 uint8 = 255
                    _ = v11
                } else {
                    var v12 uint8 = v8
                    _ = v12
                }
                _ = fi10
            }
        } else {
            for fi13 := 0; fi13 < 7; fi13++ {
                if false {
                    v7 += (v7 - 255)
                }
                _ = fi13
            }
            v14 := GenS1{F0: ("foo bar" + "foo bar"), F1: fi9}
            _ = v14
        }
        _ = fi9
    }
    fmt.Println(v8, v7, v8)
    v15 := []int{}
    fmt.Println(v2)
    v16 := GenS2{F0: -9223372036854775808}
    switch "hello" {
    case "foo bar":
        fmt.Println(v3, v1, v8)
    case "":
        v2 -= 683
    case "hello":
        if true {
            fmt.Println(v8, v8)
            fmt.Println(v8, v7, v1)
        } else {
            _ = len(v15)
        }
    default:
        v17 := GenS2{F0: -826}
        _ = v17
    }
    switch 42 {
    case 626:
        v7 += 255
    default:
        v1 -= (v1 % ((v1) | 1))
    }
    switch 42 {
    case -340:
        var v18 int64 = 42
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v7
    _ = v8
}
