package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint64 }
type GenS2 struct { F0 uint16; F1 int64; F2 int }

func main() {
    v1 := []int{9223372036854775807, 32768, 32767}
    v2 := make(map[uint32]int)
    var v3 uint16 = 0
    v4 := make(map[string]int)
    var v5 uint = 18446744073709551615
    var v6 uint = v5
    for fi7 := 0; fi7 < 9; fi7++ {
        for fi8 := 0; fi8 < 13; fi8++ {
            _ = len(v1)
            if true {
                var v9 uint64 = 18446744073709551615
                var v10 string = ((("hello" + ("a" + "a")) + ("foo bar" + "\n\t")) + "\n\t")
                if true {
                    v6 = v5
                    v11 := GenS1{F0: 1, F1: -3, F2: v9}
                    v6 += v6
                    _ = v11
                }
                fmt.Println(v10, v3)
                _ = v10
                _ = v9
            }
            v6 += 912
            _ = fi8
        }
        _ = fi7
    }
    v6 = 2
    switch -9223372036854775808 {
    case 96:
        var v12 uint = (v5 - v6)
        _ = v12
    case -898:
        for fi13 := 0; fi13 < 3; fi13++ {
            var v14 int = (fi13 + (fi13 % ((fi13) | 1)))
            _ = fi13
            _ = v14
        }
    }
    var v15 uint64 = 322
    for fi16 := 0; fi16 < 17; fi16++ {
        for fi17 := 0; fi17 < 6; fi17++ {
            var v18 uint64 = v15
            _ = fi17
            _ = v18
        }
        for fi19 := 0; fi19 < 15; fi19++ {
            var v20 int16 = -817
            v3 = (^(v3))
            _ = fi19
            _ = v20
        }
        _ = fi16
    }
    var v21 string = "hello"
    v21 = (v21 + ("" + "\n\t"))
    v2[60] = -2147483648
    v5 += (^(v5))
    var v22 uint32 = uint32(v15)
    switch 20 {
    case 993:
        for fi23 := 0; fi23 < 1; fi23++ {
            if true {
                var v24 uint8 = 226
                delete(v4, "foo bar")
                v22 -= v22
                _ = v24
            } else {
                v3 += uint16(v5)
                switch (v21 + (v21 + "")) {
                case "a":
                    var v25 string = "\n\t"
                    _ = v25
                case "":
                    _ = len(v1)
                default:
                    _ = len(v1)
                }
                v26 := v4[((v21 + v21) + v21)]
                _ = v26
            }
            _ = fi23
        }
    default:
        v27 := GenS2{F0: v3, F1: 837, F2: 802}
        _ = v27
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
