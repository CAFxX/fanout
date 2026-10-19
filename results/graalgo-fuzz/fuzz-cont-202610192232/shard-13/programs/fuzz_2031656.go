package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 uint64 }

func main() {
    var v1 int32 = 1
    v2 := make(map[uint32]int)
    v3 := []uint32{65535, 0}
    v4 := make(map[uint32]bool)
    v1 = -2147483648
    if (false && (true && true)) {
        _ = len(v2)
        var v5 float32 = 3.14159
        _ = v5
    }
    fmt.Println(v1)
    v6 := v2[42]
    v6 = v6
    v1 += v1
    v7 := GenS1{F0: 0, F1: 1000}
    v8 := GenS1{F0: 2, F1: 1}
    switch v6 {
    case 32767:
        v9 := GenS1{F0: 0, F1: 0}
        _ = v9
    case 177:
        switch "a" {
        case "":
            if true {
                switch v6 {
                case 341:
                    for fi10 := 0; fi10 < 5; fi10++ {
                        if true {
                            var v11 float32 = ((3.14159 * 3.14159) / (-2.25 / 1e10))
                            v6 += v6
                            var v12 uint64 = 1
                            _ = v11
                            _ = v12
                        }
                        _ = fi10
                    }
                case 4294967295:
                    v13 := GenS1{F0: 2, F1: int16(v1)}
                    _ = v13
                default:
                    var v14 int32 = 42
                    _ = v14
                }
                switch (("a" + "") + "a") {
                case "\n\t":
                    v15 := GenS2{F0: true, F1: 18446744073709551615}
                    _ = v15
                case "a":
                    v1 -= (v1 & (v1 | v1))
                case "":
                    for fi16 := 0; fi16 < 6; fi16++ {
                        v3 = append(v3, 2147483647)
                        v17 := GenS2{F0: (4294967295 <= 18446744073709551615), F1: 0}
                        _ = fi16
                        _ = v17
                    }
                }
            }
        case "foo bar":
            fmt.Println(v1, v6)
        default:
            v6 = 9223372036854775807
        }
    }
    fmt.Println(v1, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
