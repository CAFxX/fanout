package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    var v3 int = 32768
    fmt.Println(v3)
    var v4 uint32 = 652
    var v5 float64 = 1.5
    v6 := []int{(v3 << ((4294967295) & 63)), 9223372036854775807, v3}
    var v7 float32 = (1.5 / 3.14159)
    for fi8 := 0; fi8 < 2; fi8++ {
        v5 += v5
        if (478 >= 42) {
            switch 590 {
            case 316:
                fmt.Println(v3, fi8)
            case 1:
                var v9 bool = (!true)
                _ = v9
            case -1:
                v5 -= v5
            }
            _ = len(v2)
            v6 = append(v6, 0)
        } else {
            if false {
                for fi10 := 0; fi10 < 16; fi10++ {
                    _ = len(v6)
                    v11 := GenS1{F0: 27, F1: 42, F2: 1}
                    _ = fi10
                    _ = v11
                }
            } else {
                for fi12 := 0; fi12 < 8; fi12++ {
                    _ = len(v6)
                    v13 := GenS1{F0: 42, F1: fi12, F2: 65535}
                    var v14 uint = 2
                    _ = fi12
                    _ = v13
                    _ = v14
                }
            }
            v4 += 2147483648
        }
        var v15 int64 = -9223372036854775808
        _ = fi8
        _ = v15
    }
    v4 -= (v4 | v4)
    var v16 int32 = -2147483648
    v7 -= 1.5
    for fi17 := 0; fi17 < 9; fi17++ {
        for fi18 := 0; fi18 < 18; fi18++ {
            v6 = append(v6, 128)
            _ = fi18
        }
        switch (v3 - 42) {
        case 0:
            v3 += int(v4)
        case -681:
            fmt.Println(v7, v3)
        case -1:
            fmt.Println(fi17)
        default:
            if ((9223372036854775807 >= -525) || false) {
                v16 = v16
            } else {
                var v19 bool = true
                v16 = (-(v16))
                _ = v19
            }
        }
        _ = fi17
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
