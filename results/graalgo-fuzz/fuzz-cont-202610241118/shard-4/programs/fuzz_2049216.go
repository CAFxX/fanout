package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint64 }
func genFn0(p0 bool, p1 uint8, p2 string) string {
    p2 = (p2 + p2)
    v5 := GenS1{F0: 1, F1: 900, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}

func main() {
    v1 := []string{}
    var v2 int64 = 9223372036854775807
    var v3 float64 = -2.25
    var v4 int32 = -1
    if false {
        switch ("\n\t" + "a") {
        case "a":
            var v6 uint64 = 1
            _ = v6
        case "":
            v3 -= v3
        }
    }
    v2 += v2
    for fi7 := 0; fi7 < 6; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            var v9 int8 = -125
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        for fi11 := 0; fi11 < 4; fi11++ {
            if ("foo bar" > "hello") {
                v12 := GenS1{F0: 170, F1: 1, F2: 1}
                v13 := GenS1{F0: 601, F1: 2, F2: 18446744073709551615}
                var v14 int64 = int64(v4)
                _ = v12
                _ = v13
                _ = v14
            } else {
                v15 := GenS1{F0: 2147483647, F1: 252, F2: 363}
                _ = v15
            }
            var v16 float32 = 1e10
            var v17 uint64 = 42
            _ = fi11
            _ = v16
            _ = v17
        }
        _ = fi10
    }
    v18 := GenS1{F0: 840, F1: 123, F2: 4294967295}
    var v19 uint = 158
    v4 += -2147483648
    if len(v1) > 0 {
        v1[0] = "foo bar"
    }
    var v20 float64 = ((v3 * 1.5) + 0.0)
    var v21 float32 = -2.25
    var v22 int16 = 1000
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
