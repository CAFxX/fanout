package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }

func main() {
    var v1 int8 = -1
    var v2 int64 = 97
    v3 := make(map[uint32]string)
    v4 := make(map[string]uint32)
    var v5 int32 = -2147483648
    v6 := GenS1{F0: 69, F1: 32768}
    v7 := v4["a"]
    var v8 uint32 = v7
    if (false || true) {
        v5 += v5
        if (true && (v8 == 2147483647)) {
            var v9 int64 = v2
            var v10 uint8 = 1
            for fi11 := 0; fi11 < 16; fi11++ {
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 uint8 = (v10 % (((v10 + (^(v10)))) | 1))
                    switch "a" {
                    case "hello":
                        v7 = v8
                    case "\n\t":
                        v14 := GenS1{F0: (v8 << ((65535) & 63)), F1: fi12}
                        _ = v14
                    default:
                        v2 = v9
                    }
                    _ = fi12
                    _ = v13
                }
                for fi15 := 0; fi15 < 13; fi15++ {
                    v2 = (v9 - v9)
                    var v16 int64 = 1
                    var v17 uint8 = v10
                    _ = fi15
                    _ = v16
                    _ = v17
                }
                for fi18 := 0; fi18 < 12; fi18++ {
                    v10 = 240
                    var v19 int = (fi18 ^ 32768)
                    _ = fi18
                    _ = v19
                }
                _ = fi11
            }
            _ = v10
            _ = v9
        } else {
            v8 = 65535
        }
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v3[(v8 << ((104) & 63))] = "a"
        _ = fi20
    }
    fmt.Println(v7, v2)
    v2 -= (^(v2))
    fmt.Println(v1)
    v5 -= 42
    v21 := GenS1{F0: v7, F1: 4294967295}
    var v22 float32 = -2.25
    v22 = v22
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
