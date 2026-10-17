package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 1e10
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    v6 := make(map[string]string)
    if false {
        if (2 != 86) {
            if false {
                v7 := []int{4294967295, 32768}
                v2 = v2
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
                _ = v7
            }
            var v8 int32 = -1
            var v9 float32 = v2
            _ = v8
            _ = v9
        } else {
            fmt.Println(v2)
            _ = len(v5)
            switch -9223372036854775808 {
            case 0:
                v10 := GenS1{F0: 42, F1: -9223372036854775808, F2: 537}
                _ = v10
            case 4294967295:
                fmt.Println(v2)
            }
        }
    }
    fmt.Println(v2)
    v2 += ((0.0 - ((1.5 * 1.5) / v2)) - v2)
    var v11 int8 = 113
    if (!(119 >= 2)) {
        fmt.Println(v11, v11)
    }
    fmt.Println(v2, v11)
    if (!false) {
        for fi12 := 0; fi12 < 4; fi12++ {
            switch "foo bar" {
            case "\n\t":
                if true {
                    v13 := []int{(fi12 % ((495) | 1))}
                    v11 = (v11 % ((v11) | 1))
                    _ = v13
                }
            case "":
                var v14 uint16 = 0
                _ = v14
            }
            v15 := GenS1{F0: -9223372036854775808, F1: fi12, F2: 1}
            _ = fi12
            _ = v15
        }
        if false {
            var v16 uint8 = 1
            fmt.Println(v11, v16, v2)
            _ = v16
        }
        v2 = 0.0
    }
    v2 += (v2 * v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
