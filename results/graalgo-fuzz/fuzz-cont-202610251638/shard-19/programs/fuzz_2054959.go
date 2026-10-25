package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint; F1 uint64; F2 bool }

func main() {
    var v1 uint64 = 79
    v2 := []uint32{}
    v3 := []uint32{}
    var v4 uint16 = 32767
    v5 := make(map[uint32]string)
    v6 := GenS1{F0: 42, F1: true, F2: -9223372036854775808}
    var v7 bool = (!((false && true) && false))
    switch "a" {
    case "\n\t":
        var v8 uint32 = 449
        _ = v8
    case "":
        var v9 uint16 = v4
        _ = v9
    default:
        v1 += 928
    }
    fmt.Println(v4)
    v4 = v4
    fmt.Println(v7)
    fmt.Println(v7)
    v7 = true
    for fi10 := 0; fi10 < 10; fi10++ {
        if ((^(fi10)) != (-(fi10))) {
            v11 := GenS2{F0: 18446744073709551615, F1: v1, F2: (!false)}
            switch fi10 {
            case -470:
                switch "" {
                case "foo bar":
                    var v12 uint32 = 2
                    _ = v12
                }
            case 2147483647:
                _ = len(v2)
            case -463:
                v1 = v1
            default:
                var v13 uint = 4294967295
                _ = v13
            }
            _ = v11
        } else {
            if true {
                if true {
                    if true {
                        var v14 uint8 = 1
                        v15 := GenS1{F0: -1, F1: v7, F2: 0}
                        v3 = append(v3, 65535)
                        _ = v14
                        _ = v15
                    }
                }
                for fi16 := 0; fi16 < 16; fi16++ {
                    v4 = 1
                    _ = fi16
                }
                if v7 {
                    v7 = false
                    fmt.Println(v7, v7, fi10)
                    var v17 uint16 = v4
                    _ = v17
                }
            }
        }
        var v18 float64 = 1e10
        var v19 int = (fi10 * 4294967295)
        _ = fi10
        _ = v18
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
