package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int8 }
type GenS2 struct { F0 int64; F1 uint32 }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]bool)
    var v3 string = ""
    var v4 float32 = -2.25
    var v5 int = 42
    var v6 uint8 = 226
    v1 -= v1
    var v7 uint = 281
    v7 -= (v7 << ((249) & 63))
    fmt.Println(v5)
    var v8 uint = v7
    fmt.Println(v6, v7)
    v9 := []int{-68, (v5 / (((v5 * (v5 << (((v7 << ((18446744073709551615) & 63))) & 63)))) | 1)), v5, -730}
    fmt.Println(v5, v4)
    v6 = (^(v6))
    var v10 uint32 = 42
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 11; fi12++ {
            if true {
                v9 = append(v9, fi11)
                v13 := GenS1{F0: -12, F1: "a", F2: 0}
                _ = v13
            } else {
                switch v5 {
                case 9223372036854775807:
                    v14 := GenS1{F0: -56, F1: "foo bar", F2: 111}
                    _ = v14
                case 2147483647:
                    v1 = v1
                }
                switch "" {
                case "hello":
                    v15 := v2[("\n\t" + ("" + v3))]
                    _ = v15
                case "\n\t":
                    v16 := GenS1{F0: 127, F1: (v3 + v3), F2: -1}
                    _ = v16
                }
            }
            _ = fi12
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        v2[((v3 + v3) + (v3 + ""))] = (v3 == ("hello" + v3))
        if true {
            switch "hello" {
            case "":
                v18 := GenS2{F0: -9223372036854775808, F1: v10}
                _ = v18
            case "foo bar":
                var v19 int = (v5 + v5)
                _ = v19
            }
            if false {
                v10 = v10
            }
        }
        v7 -= (v8 >> ((v7) & 63))
        _ = fi17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
