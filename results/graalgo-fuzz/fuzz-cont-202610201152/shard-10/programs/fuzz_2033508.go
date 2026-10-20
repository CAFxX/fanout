package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int; F2 uint64 }

func main() {
    var v1 uint8 = 170
    v2 := []string{"a", "\n\t", "", ("" + "a")}
    var v3 uint8 = v1
    v4 := make(map[uint32]int)
    var v5 bool = (!(false || false))
    v6 := v4[2147483648]
    v3 = (^(v1))
    v7 := GenS1{F0: (!((^(v3)) > (v3 % ((v3) | 1))))}
    v3 = 242
    switch (v6 * v6) {
    case 4294967295:
        if (((v6 >> ((uint(v3)) & 63)) < 2147483647) && v5) {
            if (((42 <= 42) && true) && true) {
                switch "" {
                case "\n\t":
                    for fi8 := 0; fi8 < 17; fi8++ {
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                        v9 := GenS2{F0: 777, F1: (fi8 - 779), F2: uint64(v1)}
                        _ = fi8
                        _ = v9
                    }
                case "hello":
                    v6 = (v6 ^ (v6 + v6))
                }
            }
            var v10 bool = ((v1 ^ 42) != v3)
            for fi11 := 0; fi11 < 9; fi11++ {
                v6 = 0
                switch "" {
                case "a":
                    _ = len(v2)
                default:
                    v2 = append(v2, ("a" + ("\n\t" + "foo bar")))
                }
                _ = fi11
            }
            _ = v10
        }
    default:
        fmt.Println(v5, v1)
    }
    var v12 uint32 = 42
    _ = len(v4)
    v3 -= (v3 % ((v3) | 1))
    v4[(v12 & uint32(v3))] = 128
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
