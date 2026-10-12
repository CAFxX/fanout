package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    var v1 uint8 = 2
    v2 := []int{}
    var v3 int64 = 321
    v4 := []uint32{2147483647, 0, 1}
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: 18446744073709551615, F1: "foo bar"}
        _ = fi5
        _ = v6
    }
    v1 = 66
    v3 += (v3 + v3)
    fmt.Println(v3)
    var v7 int32 = 0
    var v8 int8 = 0
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println(v7, v3)
    v7 = v7
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 uint16 = 936
            var v12 int32 = 2147483647
            switch fi9 {
            case 128:
                switch fi10 {
                case -9223372036854775808:
                    var v13 string = (("foo bar" + "") + "")
                    _ = v13
                case 760:
                    switch "foo bar" {
                    case "\n\t":
                        v14 := GenS1{F0: 18446744073709551615, F1: "\n\t"}
                        _ = v14
                    case "":
                        var v15 uint32 = 73
                        _ = v15
                    }
                case 214:
                    v3 += v3
                }
            case 9223372036854775807:
                if true {
                    v12 += 2147483647
                } else {
                    v1 += v1
                    v12 -= 42
                }
            default:
                v16 := GenS2{F0: (^(v8)), F1: v8}
                v11 = 65535
                v3 -= 9223372036854775807
                _ = v16
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
