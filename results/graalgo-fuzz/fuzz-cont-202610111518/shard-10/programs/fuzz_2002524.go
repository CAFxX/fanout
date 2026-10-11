package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS2 struct { F0 uint16 }

func main() {
    var v1 bool = true
    v2 := []int{}
    v3 := []uint32{2147483647, 1, 2147483647, 793}
    var v4 int8 = 127
    v3 = append(v3, 188)
    v5 := GenS1{F0: 834, F1: 2147483648, F2: 1}
    v1 = v1
    v6 := GenS1{F0: 18446744073709551615, F1: 2, F2: 821}
    var v7 float32 = -2.25
    v1 = v1
    v1 = true
    fmt.Println(v7)
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 8; fi9++ {
            if len(v2) > 0 {
                v2[0] = fi9
            }
            var v10 uint8 = 95
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := GenS1{F0: 18446744073709551615, F1: 2, F2: 669}
    v12 := GenS2{F0: 0}
    switch (("" + "foo bar") + "") {
    case "hello":
        for fi13 := 0; fi13 < 3; fi13++ {
            if len(v2) > 0 {
                v2[0] = 309
            }
            switch "" {
            case "hello":
                if false {
                    switch "foo bar" {
                    case "":
                        fmt.Println(v7)
                    default:
                        var v14 bool = (v1 || v1)
                        _ = v14
                    }
                } else {
                    switch ("\n\t" + "") {
                    case "hello":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "\n\t":
                        var v15 float32 = 0.0
                        _ = v15
                    }
                    var v16 int32 = 418
                    _ = v16
                }
            case "\n\t":
                var v17 uint8 = 128
                _ = v17
            default:
                var v18 uint = 4294967295
                _ = v18
            }
            _ = fi13
        }
    case "\n\t":
        var v19 uint8 = 127
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
