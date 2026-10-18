package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 bool; F1 uint64; F2 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 uint64 }
func genFn0(p0 int8) uint16 {
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := []int{127}
        var v6 int32 = -1
        p0 = p0
        _ = fi4
        _ = v5
        _ = v6
    }
    if ("a" >= ("\n\t" + "")) {
        fmt.Println(p0)
        var v7 int = -33
        v7 = (v7 | -545)
        _ = v7
    }
    _ = p0
    return 42
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint8 = 128
    v3 := make(map[int]int)
    v8 := GenS1{F0: 42, F1: 4294967295}
    v9 := []int{-2147483648, -9223372036854775808}
    if true {
        if (v2 <= 127) {
            var v10 int8 = -1
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        v2 = (^(v2))
        if false {
            var v12 uint16 = genFn0(-39)
            _ = v12
        }
        if (-2.25 > (0.0 * float64((0.0 - 0.1)))) {
            if (502 < 0) {
                var v13 int32 = 42
                switch "hello" {
                case "a":
                    switch (fi11 & (fi11 & 32768)) {
                    case 42:
                        var v14 int = (-(fi11))
                        _ = v14
                    default:
                        var v15 int16 = -32768
                        _ = v15
                    }
                }
                fmt.Println(v13, v1, fi11)
                _ = v13
            }
            v1 += (v1 % ((v1) | 1))
        } else {
            v2 -= (v2 + 1)
            v2 += (^(v2))
            var v16 uint64 = 42
            _ = v16
        }
        _ = fi11
    }
    v17 := GenS2{F0: true, F1: 2, F2: 538}
    fmt.Println(v2, v1)
    var v18 uint32 = 0
    fmt.Println(v18, v1, v2)
    v19 := GenS3{F0: (v18 >> ((2) & 63)), F1: v1, F2: 42}
    fmt.Println(v1, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
