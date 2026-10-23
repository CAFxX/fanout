package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint8; F1 uint64; F2 uint8 }
type GenS5 struct { F0 int16; F1 int8; F2 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 42
    var v3 int32 = 1
    v4 := make(map[string]uint32)
    v5 := []uint32{}
    v5 = append(v5, 586)
    switch (-(v2)) {
    case -9223372036854775808:
        v6 := GenS1{F0: 42, F1: 228}
        _ = v6
    }
    delete(v1, 491)
    v3 += v3
    for fi7 := 0; fi7 < 3; fi7++ {
        if (!false) {
            if len(v5) > 0 {
                v5[0] = 1
            }
            if len(v5) > 0 {
                v5[0] = 42
            }
            v8 := GenS1{F0: 67, F1: 0}
            _ = v8
        } else {
            var v9 int8 = 100
            switch (("" + "hello") + ("\n\t" + "foo bar")) {
            case "foo bar":
                v9 = v9
            case "\n\t":
                var v10 uint32 = 2147483647
                _ = v10
            default:
                var v11 bool = (!(!(!(110 != 0))))
                _ = v11
            }
            _ = v9
        }
        if ((v3 >= (v3 << ((1) & 63))) || false) {
            for fi12 := 0; fi12 < 3; fi12++ {
                var v13 int64 = 42
                _ = fi12
                _ = v13
            }
        } else {
            v14 := GenS2{F0: 2, F1: -1}
            _ = v14
        }
        v15 := GenS2{F0: 2147483648, F1: 42}
        _ = fi7
        _ = v15
    }
    v3 = v3
    v16 := GenS3{F0: -1}
    fmt.Println(v2, v2)
    fmt.Println(v3, v2)
    v3 += 0
    if len(v5) > 0 {
        v5[0] = 32
    }
    v3 = (v3 * (v3 / ((v3) | 1)))
    if true {
        v17 := GenS4{F0: 128, F1: 42, F2: 127}
        _ = v17
    }
    switch ("hello" + "\n\t") {
    case "a":
        var v18 int8 = 100
        _ = v18
    }
    v4[("hello" + "\n\t")] = 2147483647
    var v19 int8 = 42
    if false {
        if (870 != 18446744073709551615) {
            if false {
                for fi20 := 0; fi20 < 16; fi20++ {
                    v3 += v3
                    if len(v5) > 0 {
                        v5[0] = 378
                    }
                    v19 = v19
                    _ = fi20
                }
            } else {
                v19 -= -1
                v19 += v19
            }
            var v21 float64 = float64(1.5)
            var v22 uint64 = 18446744073709551615
            _ = v21
            _ = v22
        }
        v23 := GenS5{F0: 1000, F1: -23, F2: v3}
        _ = v23
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
