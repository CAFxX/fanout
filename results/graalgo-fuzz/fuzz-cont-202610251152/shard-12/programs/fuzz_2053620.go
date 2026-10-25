package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }

func main() {
    var v1 int8 = 1
    var v2 uint = 1
    v3 := []string{"\n\t", "hello", "\n\t"}
    var v4 float32 = 0.1
    var v5 float32 = float32((0.0 + -2.25))
    v6 := make(map[int]uint32)
    v7 := GenS1{F0: 42, F1: -240, F2: v2}
    for fi8 := 0; fi8 < 4; fi8++ {
        if (((("\n\t" + "hello") + ((("\n\t" + "foo bar") + ("foo bar" + (("\n\t" + "\n\t") + "hello"))) + "")) + "hello") >= "") {
            if len(v3) > 0 {
                v3[0] = ""
            }
            _ = len(v6)
        }
        _ = fi8
    }
    v9 := GenS2{F0: false, F1: 1}
    v10 := v6[2147483647]
    for fi11 := 0; fi11 < 17; fi11++ {
        v1 -= (v1 << (((^(v2))) & 63))
        if (true || true) {
            var v12 bool = false
            v13 := GenS1{F0: 966, F1: 9223372036854775807, F2: 18446744073709551615}
            fmt.Println(v4, v2)
            _ = v12
            _ = v13
        }
        _ = fi11
    }
    fmt.Println(v10)
    v5 = (v5 * v4)
    v2 = (^(v2))
    v2 += (v2 ^ (v2 + (^(v2))))
    v1 = 0
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v14 := GenS3{F0: 1, F1: (v10 << ((42) & 63)), F2: v1}
    if len(v3) > 0 {
        v3[0] = (("a" + "foo bar") + ("" + (("\n\t" + ("foo bar" + ("\n\t" + "hello"))) + (("a" + ("\n\t" + "foo bar")) + ""))))
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 float64 = (0.1 * 0.1)
            var v18 uint64 = 42
            _ = fi16
            _ = v17
            _ = v18
        }
        fmt.Println(v10, v2, v5)
        v19 := GenS2{F0: (1.5 == -2.25), F1: -1}
        _ = fi15
        _ = v19
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
