package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int16) string {
    fmt.Println(p0)
    var v4 int64 = 1
    _ = p0
    _ = v4
    return ("foo bar" + "")
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint64 = 254
    var v3 int64 = 1
    fmt.Println(v2)
    v1 = v1
    fmt.Println(v1)
    switch 1 {
    case 0:
        var v5 int64 = (^(v3))
        _ = v5
    case -9223372036854775808:
        v6 := GenS1{F0: 0}
        _ = v6
    }
    v7 := []int{32767, 127, 9223372036854775807, -1}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    v3 += v3
    fmt.Println(v2)
    switch "hello" {
    case "":
        if true {
            v8 := GenS1{F0: 813}
            for fi9 := 0; fi9 < 12; fi9++ {
                if (180 >= 65535) {
                    v10 := GenS1{F0: -1}
                    var v11 uint16 = 65535
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            v3 = (v3 + (v3 - v3))
            _ = v8
        }
    default:
        v2 += v2
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        switch "" {
        case "":
            v2 = uint64(v1)
        case "foo bar":
            if len(v7) >= 2 {
                v7 = v7[1:len(v7)]
            }
        default:
            fmt.Println(v3, v3)
        }
        var v13 int8 = 0
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            v7 = append(v7, (fi15 * fi15))
            v3 += (v3 & -9223372036854775808)
            _ = fi15
        }
        var v16 uint = 65535
        var v17 float64 = 1e10
        _ = fi14
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v7
}
