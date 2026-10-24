package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32; F1 uint }

func main() {
    var v1 float32 = (-2.25 / 3.14159)
    v2 := make(map[int]uint32)
    var v3 int16 = -429
    v4 := make(map[int]int)
    v5 := make(map[string]bool)
    var v6 uint64 = 0
    v3 = v3
    v6 = (v6 - (^(v6)))
    v7 := v4[-122]
    v8 := GenS1{F0: 2147483647}
    for fi9 := 0; fi9 < 5; fi9++ {
        fmt.Println(v1, v3)
        if (true && true) {
            if true {
                switch v7 {
                case 9223372036854775807:
                    if false {
                        v3 -= -882
                        v2[-949] = 700
                    } else {
                        var v10 uint64 = (v6 - 18446744073709551615)
                        _ = v10
                    }
                case 4294967295:
                    v11 := GenS1{F0: -1}
                    _ = v11
                case -849:
                    var v12 uint = 336
                    _ = v12
                default:
                    v7 = (-(fi9))
                }
                v7 += (fi9 & 128)
                _ = len(v5)
            } else {
                var v13 uint = 81
                _ = len(v2)
                _ = v13
            }
        }
        v14 := GenS1{F0: -2147483648}
        _ = fi9
        _ = v14
    }
    var v15 uint16 = 65535
    v16 := []int{v7}
    for fi17 := 0; fi17 < 20; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 int = -541
            v19 += fi17
            _ = fi18
            _ = v19
        }
        switch (fi17 >> ((255) & 63)) {
        case 32768:
            fmt.Println(v1, v7)
        }
        _ = fi17
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 uint64 = (v6 * 2)
        v22 := GenS2{F0: 228, F1: 940}
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
