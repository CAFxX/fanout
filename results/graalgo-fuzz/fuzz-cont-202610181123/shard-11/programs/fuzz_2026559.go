package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint64 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    var v3 int = 2147483647
    v4 := []int{-175, 951}
    var v5 uint16 = 79
    v6 := GenS1{F0: 42, F1: true, F2: 18446744073709551615}
    fmt.Println(v5, v5)
    v1[115] = 162
    switch "a" {
    case "\n\t":
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            for fi7 := 0; fi7 < 9; fi7++ {
                if false {
                    if true {
                        switch "hello" {
                        case "":
                            var v8 uint16 = 227
                            _ = v8
                        case "foo bar":
                            v5 = (^(v5))
                        case "\n\t":
                            v9 := GenS1{F0: 0, F1: (!("foo bar" == "")), F2: 1}
                            _ = v9
                        default:
                            var v10 uint32 = 4294967295
                            _ = v10
                        }
                        v5 = 2
                        var v11 int32 = 1
                        _ = v11
                    }
                }
                switch "" {
                case "":
                    var v12 int8 = 89
                    _ = v12
                }
                v5 = v5
                _ = fi7
            }
        case "":
            var v13 uint8 = 1
            _ = v13
        case "\n\t":
            v5 = (v5 ^ (v5 / (((v5 << ((1) & 63))) | 1)))
        }
    case "hello":
        v5 = (^(v5))
    case "":
        v3 += (v3 * -424)
    }
    for fi14 := 0; fi14 < 11; fi14++ {
        switch 127 {
        case 185:
            v15 := v2["foo bar"]
            _ = v15
        }
        for fi16 := 0; fi16 < 13; fi16++ {
            v3 = fi14
            var v17 bool = true
            v18 := v1[2147483647]
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS1{F0: 2147483647, F1: true, F2: 2}
        _ = fi14
        _ = v19
    }
    v20 := GenS1{F0: 932, F1: true, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
