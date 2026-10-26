package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 string, p1 uint16, p2 int64) bool {
    if (279 > 18446744073709551615) {
        var v6 int = 118
        v7 := GenS1{F0: 0, F1: 42}
        _ = v6
        _ = v7
    } else {
        v8 := []int{217}
        var v9 float32 = float32((1e10 - ((1.5 / (0.1 + (1.5 + 0.1))) / (3.14159 / (float64(-2.25) * 3.14159)))))
        p0 = p0
        _ = v8
        _ = v9
    }
    if (0 > 0) {
        switch ((p0 + p0) + "hello") {
        case "a":
            if (42 < 241) {
                fmt.Println(p2, p0)
                var v10 bool = (p1 == (p1 << ((1) & 63)))
                _ = v10
            }
        case "hello":
            var v11 uint8 = 0
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            fmt.Println(p2, p1, fi12)
            var v13 int64 = 42
            _ = fi12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            if false {
                v16 := GenS1{F0: 1, F1: -32768}
                _ = v16
            }
            p1 = p1
            if ("foo bar" < "\n\t") {
                p0 = p0
                var v17 uint64 = 18446744073709551615
                var v18 uint64 = (v17 * (v17 * v17))
                _ = v17
                _ = v18
            } else {
                var v19 int32 = 794
                v20 := []int{(-(fi14)), -57, 4294967295, fi14}
                _ = v19
                _ = v20
            }
            _ = fi15
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    return (65535 >= 18446744073709551615)
}

func main() {
    var v1 uint8 = 99
    var v2 int16 = 417
    var v3 float32 = (0.0 * 3.14159)
    var v4 string = "\n\t"
    var v5 uint64 = 0
    fmt.Println(v4)
    v21 := GenS1{F0: v5, F1: (v2 * 32767)}
    var v22 uint8 = v1
    v2 += v2
    v23 := GenS2{F0: 245, F1: 9223372036854775807}
    for fi24 := 0; fi24 < 10; fi24++ {
        if true {
            for fi25 := 0; fi25 < 14; fi25++ {
                v26 := GenS2{F0: (v22 >> ((0) & 63)), F1: 9223372036854775807}
                _ = fi25
                _ = v26
            }
            var v27 bool = true
            _ = v27
        }
        v28 := []int{32767, (-(fi24)), int(v1)}
        _ = fi24
        _ = v28
    }
    fmt.Println(v2, v4, v2)
    for fi29 := 0; fi29 < 15; fi29++ {
        fmt.Println(v5, v5)
        _ = fi29
    }
    var v30 string = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
