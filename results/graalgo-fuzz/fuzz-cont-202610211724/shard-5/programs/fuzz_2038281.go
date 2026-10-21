package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int32; F1 uint8; F2 int32 }

func main() {
    v1 := []string{"foo bar", "foo bar", "foo bar"}
    var v2 int = 4294967295
    v3 := []string{"a", ""}
    var v4 uint = 0
    var v5 int = v2
    v2 = (v2 >> ((18446744073709551615) & 63))
    var v6 int64 = 1
    var v7 float32 = 1.5
    for fi8 := 0; fi8 < 3; fi8++ {
        if true {
            v1 = append(v1, "a")
            v6 -= (v6 << (((v4 << ((v4) & 63))) & 63))
            v1 = append(v1, "\n\t")
        }
        _ = fi8
    }
    for fi9 := 0; fi9 < 15; fi9++ {
        switch "\n\t" {
        case "":
            if false {
                var v10 uint8 = 127
                for fi11 := 0; fi11 < 12; fi11++ {
                    var v12 string = "a"
                    _ = fi11
                    _ = v12
                }
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := GenS1{F0: 18446744073709551615}
                    v6 = 42
                    _ = fi13
                    _ = v14
                }
                _ = v10
            }
        default:
            v5 -= int(v6)
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        var v15 int8 = 100
        _ = fi9
        _ = v15
    }
    v7 -= v7
    v2 -= -9223372036854775808
    fmt.Println(v4, v6)
    if len(v3) > 0 {
        v3[0] = "a"
    }
    var v16 int16 = 645
    var v17 uint16 = 65535
    var v18 float64 = 3.14159
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v19 := GenS2{F0: -1, F1: 2, F2: 42}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
