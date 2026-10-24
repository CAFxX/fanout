package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }

func main() {
    var v1 bool = true
    var v2 int32 = 0
    var v3 float64 = (1e10 * (1.5 + -2.25))
    var v4 int16 = 1000
    var v5 uint32 = 2147483648
    v6 := []int{2147483647, 9223372036854775807}
    fmt.Println(v4)
    v5 -= (v5 * (v5 ^ (v5 - (v5 / (((v5 - (^(v5)))) | 1)))))
    switch "\n\t" {
    case "hello":
        if v1 {
            var v7 uint16 = 42
            v2 = 42
            for fi8 := 0; fi8 < 20; fi8++ {
                switch ("hello" + "a") {
                case "hello":
                    v7 = v7
                default:
                    var v9 int16 = (^(v4))
                    _ = v9
                }
                _ = fi8
            }
            _ = v7
        }
    }
    _ = len(v6)
    v3 = (1e10 - v3)
    v10 := GenS1{F0: 255}
    v2 = (v2 << ((42) & 63))
    if true {
        fmt.Println(v4)
        if (false && true) {
            switch "" {
            case "":
                var v11 uint64 = 4294967295
                _ = v11
            case "\n\t":
                for fi12 := 0; fi12 < 20; fi12++ {
                    var v13 int16 = (^(v4))
                    fmt.Println(v5)
                    if len(v6) > 0 {
                        v6[0] = (fi12 & (fi12 | (fi12 >> ((255) & 63))))
                    }
                    _ = fi12
                    _ = v13
                }
            case "foo bar":
                v5 = (^(v5))
            }
        } else {
            if len(v6) > 0 {
                v6[0] = 0
            }
        }
        var v14 int = 0
        _ = v14
    }
    fmt.Println(v3)
    v1 = v1
    v2 = 0
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
