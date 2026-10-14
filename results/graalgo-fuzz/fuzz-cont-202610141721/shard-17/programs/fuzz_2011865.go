package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int }
func genFn0(p0 int8, p1 uint64) uint32 {
    v5 := []int{}
    fmt.Println(p1)
    _ = len(v5)
    v5 = append(v5, 9223372036854775807)
    _ = p0
    _ = p1
    _ = v5
    return 65535
}

func main() {
    v1 := []uint32{2147483648, 1}
    v2 := []int{0}
    var v3 uint64 = 469
    v4 := []string{"hello", ("\n\t" + ("" + "\n\t"))}
    var v6 int16 = 32767
    if ("hello" > "") {
        switch "hello" {
        case "":
            var v7 int = 9223372036854775807
            _ = v7
        case "hello":
            if ((v3 * (v3 << ((4294967295) & 63))) > (v3 | 18446744073709551615)) {
                if true {
                    fmt.Println(v3)
                } else {
                    v8 := GenS1{F0: 0, F1: 4294967295}
                    v3 -= v3
                    var v9 int64 = 286
                    _ = v8
                    _ = v9
                }
            }
        case "foo bar":
            var v10 int8 = -21
            _ = v10
        default:
            v3 = (v3 - v3)
        }
        v3 -= (^(v3))
    } else {
        v3 = v3
        v11 := GenS1{F0: -907, F1: -1}
        var v12 string = "a"
        _ = v11
        _ = v12
    }
    var v13 int16 = 32767
    switch 127 {
    case 2147483647:
        fmt.Println(v13)
    case 9223372036854775807:
        v14 := GenS1{F0: 9223372036854775807, F1: -22}
        _ = v14
    default:
        v2 = append(v2, -9223372036854775808)
    }
    var v15 uint64 = (v3 - (v3 + (v3 * 0)))
    var v16 uint = 65535
    var v17 float32 = 0.1
    switch 2147483647 {
    case -2147483648:
        v13 = v13
    case -579:
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    }
    for fi18 := 0; fi18 < 10; fi18++ {
        switch "foo bar" {
        case "hello":
            fmt.Println(v16)
        default:
            v19 := GenS1{F0: 42, F1: fi18}
            _ = v19
        }
        _ = fi18
    }
    var v20 string = "foo bar"
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v6
}
