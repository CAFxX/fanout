package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }

func main() {
    var v1 float64 = (float64(1e10) * 0.0)
    v2 := []uint32{0, 2}
    v3 := []string{"\n\t", "\n\t", "foo bar"}
    var v4 uint = 251
    fmt.Println(v4, v1)
    var v5 uint32 = 322
    v6 := GenS1{F0: 2}
    if false {
        var v7 uint = (^(v4))
        _ = v7
    }
    var v8 float64 = float64((0.1 + 1.5))
    v9 := GenS1{F0: 931}
    for fi10 := 0; fi10 < 13; fi10++ {
        for fi11 := 0; fi11 < 10; fi11++ {
            v8 += v1
            switch (^(fi11)) {
            case 128:
                if len(v2) > 0 {
                    v2[0] = v5
                }
            case 0:
                var v12 string = "\n\t"
                _ = v12
            case 9223372036854775807:
                v1 -= ((((v8 + (float64(-2.25) + 1.5)) / (v1 + 1e10)) / (v8 - 0.1)) - 0.1)
            default:
                var v13 uint = 714
                _ = v13
            }
            var v14 int64 = 0
            _ = fi11
            _ = v14
        }
        _ = fi10
    }
    var v15 float64 = v8
    switch "hello" {
    case "hello":
        if ((318 < 754) && false) {
            var v16 float32 = 3.14159
            fmt.Println(v15)
            v17 := GenS1{F0: 0}
            _ = v16
            _ = v17
        }
    }
    if len(v2) > 0 {
        v2[0] = v5
    }
    if (!false) {
        if false {
            v1 -= v1
            v15 = (0.1 - 1e10)
            if (!((v4 * v4) == uint(v5))) {
                v5 = 691
                var v18 uint16 = 65535
                var v19 int8 = 1
                _ = v18
                _ = v19
            } else {
                v1 = (v15 + v15)
                var v20 uint8 = 2
                _ = v20
            }
        }
        v5 = 2147483647
    }
    for fi21 := 0; fi21 < 16; fi21++ {
        var v22 float64 = v8
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 19; fi23++ {
        fmt.Println(v15, v5, v4)
        v24 := GenS2{F0: 32767}
        var v25 int64 = -421
        _ = fi23
        _ = v24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
