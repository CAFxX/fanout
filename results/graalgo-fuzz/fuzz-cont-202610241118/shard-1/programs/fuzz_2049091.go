package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 int8 }

func main() {
    v1 := []string{}
    var v2 int = 202
    v3 := []int{9223372036854775807, (v2 >> ((65535) & 63)), (v2 + v2)}
    var v4 float64 = 1e10
    var v5 float64 = v4
    var v6 uint = 116
    for fi7 := 0; fi7 < 13; fi7++ {
        var v8 uint16 = 32767
        _ = fi7
        _ = v8
    }
    fmt.Println(v6, v6)
    var v9 uint8 = 1
    v10 := GenS1{F0: 871, F1: v2, F2: 42}
    if (730 < 618) {
        var v11 string = ("" + "")
        for fi12 := 0; fi12 < 11; fi12++ {
            switch v2 {
            case 0:
                for fi13 := 0; fi13 < 7; fi13++ {
                    switch ("" + v11) {
                    case "hello":
                        v9 = v9
                    case "\n\t":
                        if true {
                            v2 = (^(v2))
                            var v14 int32 = -1
                            fmt.Println(fi12, v6)
                            _ = v14
                        }
                    default:
                        v15 := GenS1{F0: uint32(v9), F1: 4294967295, F2: 873}
                        _ = v15
                    }
                    v4 += ((v5 + (0.1 + float64(1e10))) * v4)
                    _ = fi13
                }
            case 42:
                var v16 uint8 = (^(v9))
                _ = v16
            case 32767:
                var v17 int16 = -1
                _ = v17
            default:
                var v18 uint64 = 4294967295
                _ = v18
            }
            v19 := GenS1{F0: 2, F1: 128, F2: 599}
            _ = fi12
            _ = v19
        }
        v20 := GenS2{F0: 42, F1: 127}
        _ = v11
        _ = v20
    }
    v2 = (-(v2))
    if false {
        var v21 uint = (v6 * v6)
        v3 = append(v3, (^(v2)))
        _ = v21
    } else {
        var v22 int64 = 64
        v6 = (^(v6))
        var v23 int8 = int8(v22)
        _ = v22
        _ = v23
    }
    for fi24 := 0; fi24 < 5; fi24++ {
        v2 -= -377
        var v25 int32 = 2147483647
        _ = fi24
        _ = v25
    }
    var v26 int64 = 9223372036854775807
    fmt.Println(v9)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
