package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 uint16 = 2
    v2 := []string{"foo bar"}
    var v3 int32 = 42
    v4 := make(map[string]uint32)
    var v5 int = 673
    var v6 float32 = float32(-2.25)
    for fi7 := 0; fi7 < 20; fi7++ {
        v8 := GenS1{F0: v3}
        _ = fi7
        _ = v8
    }
    if false {
        v3 -= (v3 % (((^(v3))) | 1))
    } else {
        if (float32(0.0) > 0.1) {
            var v9 float64 = -2.25
            if ((false || true) || false) {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                var v10 int32 = (v3 % ((2147483647) | 1))
                var v11 uint16 = (v1 & v1)
                _ = v10
                _ = v11
            }
            _ = v9
        }
        v1 = v1
        v12 := GenS2{F0: 127, F1: v1}
        _ = v12
    }
    if false {
        fmt.Println(v3, v1, v3)
    }
    if (!(v5 < (-(v5)))) {
        if false {
            fmt.Println(v6, v1)
            var v13 uint = 255
            switch (v5 - v5) {
            case 9223372036854775807:
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 bool = true
                    var v16 uint32 = 2
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = v13
        } else {
            if (v5 > 937) {
                var v17 uint = 4294967295
                _ = v17
            }
        }
    }
    v18 := GenS1{F0: (v3 - (v3 | v3))}
    fmt.Println(v5)
    for fi19 := 0; fi19 < 3; fi19++ {
        var v20 uint64 = 994
        v6 -= v6
        _ = fi19
        _ = v20
    }
    var v21 int16 = 0
    var v22 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
