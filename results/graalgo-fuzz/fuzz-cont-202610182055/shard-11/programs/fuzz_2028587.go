package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }

func main() {
    v1 := []string{"a", "a", "foo bar"}
    v2 := make(map[uint32]string)
    v3 := make(map[string]string)
    var v4 float64 = 1e10
    v5 := []string{"hello"}
    v6 := make(map[string]bool)
    v4 = (-2.25 - ((((float64(0.1) * v4) - -2.25) / v4) * -2.25))
    if (!true) {
        var v7 int = 185
        for fi8 := 0; fi8 < 1; fi8++ {
            fmt.Println(v4, fi8, v7)
            for fi9 := 0; fi9 < 18; fi9++ {
                fmt.Println(fi8, v4, v7)
                if false {
                    var v10 int16 = 42
                    var v11 int32 = -624
                    v10 = v10
                    _ = v10
                    _ = v11
                }
                _ = fi9
            }
            _ = fi8
        }
        v12 := GenS1{F0: 9223372036854775807, F1: 1}
        _ = v12
        _ = v7
    }
    var v13 uint64 = 356
    v13 = 2
    if (-32768 <= 626) {
        v4 = (v4 * ((3.14159 + 1e10) * 1e10))
        var v14 uint16 = 433
        for fi15 := 0; fi15 < 2; fi15++ {
            _ = len(v3)
            var v16 int32 = 1
            if (true || true) {
                delete(v2, 2147483648)
                v16 += (v16 & (v16 % ((-1) | 1)))
            } else {
                if true {
                    delete(v2, 48)
                    v14 = (v14 | (v14 >> ((2) & 63)))
                    fmt.Println(v13, v4)
                }
                var v17 int32 = v16
                fmt.Println(v16, v13, v14)
                _ = v17
            }
            _ = fi15
            _ = v16
        }
        _ = v14
    }
    var v18 uint = 42
    _ = len(v6)
    delete(v2, 417)
    fmt.Println(v13, v4, v13)
    v19 := GenS1{F0: 699, F1: 623}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
