package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }

func main() {
    var v1 uint = 1
    v2 := []uint32{uint32(v1), 4294967295, 2}
    var v3 int = 42
    var v4 uint8 = 255
    if len(v2) > 0 {
        v2[0] = 738
    }
    for fi5 := 0; fi5 < 12; fi5++ {
        var v6 uint = (^(v1))
        _ = fi5
        _ = v6
    }
    v3 += 9223372036854775807
    v7 := GenS1{F0: 1000, F1: ((("a" + "a") + "a") + ("\n\t" + ("\n\t" + "")))}
    v8 := GenS1{F0: 0, F1: ""}
    v2 = append(v2, 2147483647)
    switch v3 {
    case -597:
        var v9 bool = (32767 != uint16(v1))
        _ = v9
    case -2147483648:
        if len(v2) > 0 {
            v2[0] = 826
        }
    case 762:
        fmt.Println(v4)
    default:
        if false {
            v4 = 158
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    v4 = (^(v4))
    v3 += (v3 - -2147483648)
    var v10 int = -1
    var v11 int64 = 459
    var v12 uint = v1
    fmt.Println(v1, v11, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
