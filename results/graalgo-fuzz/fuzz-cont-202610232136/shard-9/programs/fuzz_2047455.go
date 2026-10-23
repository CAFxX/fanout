package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 int8; F2 uint }

func main() {
    v1 := []int{}
    var v2 float64 = -2.25
    var v3 uint = 75
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 uint64 = 4294967295
        if (v5 < v5) {
            var v6 bool = (127 == 42)
            v5 = 991
            if len(v1) > 0 {
                v1[0] = fi4
            }
            _ = v6
        }
        switch "foo bar" {
        case "foo bar":
            var v7 int32 = 607
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v2 = v2
    fmt.Println(v2)
    var v8 uint = (v3 * v3)
    var v9 uint8 = 127
    var v10 int64 = -1
    v10 = v10
    if (!true) {
        v11 := GenS1{F0: 0, F1: 1}
        v8 += v8
        if (v2 != (0.0 + v2)) {
            for fi12 := 0; fi12 < 11; fi12++ {
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 int = fi13
                    var v15 uint32 = 0
                    v16 := GenS2{F0: fi12, F1: 32767}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
                _ = fi12
            }
        } else {
            switch "hello" {
            case "\n\t":
                if false {
                    v1 = append(v1, -525)
                    for fi17 := 0; fi17 < 8; fi17++ {
                        var v18 uint32 = 65535
                        _ = fi17
                        _ = v18
                    }
                }
            case "a":
                fmt.Println(v8, v3)
            }
            if false {
                fmt.Println(v8)
                for fi19 := 0; fi19 < 15; fi19++ {
                    v10 -= 9223372036854775807
                    v20 := GenS3{F0: -9223372036854775808, F1: -126, F2: (v8 * (v3 ^ (v8 % ((42) | 1))))}
                    _ = fi19
                    _ = v20
                }
            }
            v21 := GenS2{F0: 9223372036854775807, F1: -1}
            _ = v21
        }
        _ = v11
    }
    var v22 float64 = v2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
