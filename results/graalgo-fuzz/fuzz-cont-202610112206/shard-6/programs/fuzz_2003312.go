package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS2 struct { F0 int }

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[string]string)
    v3 := make(map[string]bool)
    v4 := []uint32{}
    var v5 int16 = -492
    if ((18446744073709551615 <= 42) || false) {
        var v6 uint16 = 567
        switch "hello" {
        case "a":
            if ((0.1 - (1.5 * 3.14159)) <= -2.25) {
                fmt.Println(v6, v6)
                v7 := GenS1{F0: 2147483647, F1: 42, F2: -623}
                _ = v7
            } else {
                var v8 string = "a"
                v9 := GenS1{F0: 0, F1: 1, F2: -384}
                _ = v8
                _ = v9
            }
        }
        _ = v6
    }
    var v10 uint8 = 128
    v5 += v5
    var v11 float64 = 3.14159
    v12 := GenS1{F0: 2, F1: 4294967295, F2: (v5 % ((v5) | 1))}
    fmt.Println(v10)
    if false {
        v5 = (v5 ^ (^(v5)))
    } else {
        if (-9223372036854775808 < 2147483647) {
            v11 = v11
        } else {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v13 string = "hello"
            _ = v13
        }
        for fi14 := 0; fi14 < 15; fi14++ {
            for fi15 := 0; fi15 < 14; fi15++ {
                v10 = v10
                v16 := GenS2{F0: -9223372036854775808}
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
