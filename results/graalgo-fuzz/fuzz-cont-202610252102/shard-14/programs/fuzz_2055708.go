package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 int }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 int; F1 uint8 }
type GenS4 struct { F0 int64; F1 int32; F2 uint }
type GenS5 struct { F0 string; F1 uint64; F2 uint8 }
func genFn0(p0 int64, p1 int16) uint32 {
    v4 := GenS1{F0: 2, F1: (p1 << ((868) & 63)), F2: -9223372036854775808}
    p0 -= p0
    for fi5 := 0; fi5 < 16; fi5++ {
        switch (("foo bar" + "") + "\n\t") {
        case "a":
            fmt.Println(p1)
        case "":
            for fi6 := 0; fi6 < 19; fi6++ {
                fmt.Println(p0, p1)
                _ = fi6
            }
        default:
            v7 := GenS2{F0: 127, F1: 27}
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        switch "\n\t" {
        case "a":
            for fi9 := 0; fi9 < 11; fi9++ {
                p1 += (p1 ^ p1)
                _ = fi9
            }
        case "":
            v10 := GenS3{F0: fi8, F1: 0}
            _ = v10
        case "\n\t":
            if (!false) {
                v11 := GenS3{F0: fi8, F1: 128}
                var v12 uint = 18446744073709551615
                _ = v11
                _ = v12
            }
        default:
            var v13 int32 = 42
            _ = v13
        }
        _ = fi8
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 int = -240
    var v2 uint64 = 4294967295
    v3 := []string{"a", "hello", "\n\t"}
    if (("\n\t" != ("a" + ("hello" + ("\n\t" + ("" + ""))))) || (1000 < -943)) {
        fmt.Println(v1)
        for fi14 := 0; fi14 < 1; fi14++ {
            v15 := GenS4{F0: -630, F1: 42, F2: 342}
            _ = fi14
            _ = v15
        }
    } else {
        v3 = append(v3, ("hello" + (("a" + ("foo bar" + "")) + "\n\t")))
    }
    var v16 float64 = 0.1
    v1 = (^(v1))
    v2 -= (v2 | 42)
    _ = len(v3)
    if false {
        var v17 int = 127
        for fi18 := 0; fi18 < 2; fi18++ {
            if false {
                var v19 uint64 = v2
                _ = v19
            }
            var v20 uint64 = 399
            _ = fi18
            _ = v20
        }
        var v21 float64 = (float64(1e10) * 1.5)
        _ = v17
        _ = v21
    }
    if (-1 == (v1 + (v1 * (v1 % ((32767) | 1))))) {
        var v22 int64 = -1
        _ = v22
    } else {
        var v23 int = 0
        v24 := GenS5{F0: "hello", F1: (^(v2)), F2: 1}
        _ = v23
        _ = v24
    }
    v25 := GenS5{F0: "hello", F1: v2, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v3
}
