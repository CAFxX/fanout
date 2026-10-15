package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 int }
func genFn0(p0 int8, p1 uint) int {
    fmt.Println(p0, p1)
    p0 += -31
    p0 += (p0 & p0)
    for fi4 := 0; fi4 < 7; fi4++ {
        fmt.Println(p1, fi4)
        for fi5 := 0; fi5 < 9; fi5++ {
            fmt.Println(fi5, fi5, p1)
            switch "hello" {
            case "\n\t":
                switch (fi5 / ((fi4) | 1)) {
                case -1:
                    var v6 int8 = p0
                    _ = v6
                case 32767:
                    fmt.Println(fi5)
                }
            }
            fmt.Println(p0)
            p1 -= p1
            var v7 float64 = 3.14159
            _ = fi5
            _ = v7
        }
        v8 := GenS1{F0: (-(p0)), F1: 127, F2: fi4}
        _ = fi4
        _ = v8
    }
    _ = p0
    _ = p1
    return 0
}

func main() {
    var v1 bool = (11 <= 42)
    var v2 uint64 = 62
    var v3 uint8 = 127
    v2 -= v2
    var v9 uint16 = 32767
    var v10 uint32 = 65535
    var v11 uint32 = uint32(v2)
    for fi12 := 0; fi12 < 17; fi12++ {
        var v13 int16 = -32768
        v14 := []int{(fi12 / ((128) | 1))}
        v10 -= v10
        _ = fi12
        _ = v13
        _ = v14
    }
    v15 := GenS1{F0: 100, F1: (v3 >> ((uint(v11)) & 63)), F2: -1}
    var v16 float32 = (float32(-2.25) * 1e10)
    var v17 bool = (v1 && (0 <= 0))
    var v18 int16 = -842
    v19 := []int{32767, -9223372036854775808, 1, 32767}
    if len(v19) > 0 {
        v19[0] = 128
    }
    switch ("\n\t" + ("\n\t" + "hello")) {
    case "\n\t":
        var v20 uint8 = 128
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v9
}
