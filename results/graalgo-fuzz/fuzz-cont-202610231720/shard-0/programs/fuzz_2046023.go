package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 int16 }
func genFn0(p0 bool, p1 int32, p2 uint64) uint64 {
    if ((false || false) || p0) {
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(fi6, p1)
            p1 += -2147483648
            fmt.Println(p2)
            _ = fi6
        }
    }
    v7 := GenS1{F0: 42, F1: (p1 % ((p1) | 1)), F2: 71}
    v8 := GenS1{F0: 110, F1: p1, F2: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int16 = -239
    v2 := make(map[string]int)
    var v3 string = "foo bar"
    var v4 int = -95
    v5 := []uint32{0}
    if len(v5) > 0 {
        v5[0] = 0
    }
    v3 = ""
    var v9 int16 = (-(v1))
    var v10 float64 = (0.1 * 1e10)
    switch ((v3 + "hello") + "") {
    case "hello":
        v11 := GenS1{F0: 1, F1: 42, F2: 127}
        _ = v11
    case "\n\t":
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 uint32 = 2147483647
            v9 = 42
            v4 = (fi12 | (fi12 << ((4294967295) & 63)))
            _ = fi12
            _ = v13
        }
    }
    var v14 uint8 = 225
    if (v1 == (v9 << ((255) & 63))) {
        v15 := v2[v3]
        _ = v15
    } else {
        v3 = "\n\t"
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        switch (v4 % ((-1) | 1)) {
        case 0:
            for fi17 := 0; fi17 < 13; fi17++ {
                fmt.Println(v9, v10)
                fmt.Println(v3, v3)
                v9 += 32767
                _ = fi17
            }
        case 4294967295:
            switch (v3 + ("\n\t" + v3)) {
            case "foo bar":
                v18 := GenS1{F0: 65535, F1: 2147483647, F2: 255}
                _ = v18
            default:
                v4 += -9223372036854775808
            }
        }
        v19 := GenS2{F0: 367, F1: v14, F2: v9}
        var v20 uint32 = 4294967295
        _ = fi16
        _ = v19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
