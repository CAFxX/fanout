package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 9223372036854775807
    v3 := []int{}
    var v4 uint = 4294967295
    var v5 int = -855
    v4 = 18446744073709551615
    v6 := v1[-1]
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 float32 = float32(float64(3.14159))
        var v9 int32 = -2147483648
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        fmt.Println(v6, v4, v2)
        v11 := GenS1{F0: (v2 << ((v4) & 63)), F1: (v4 % ((728) | 1))}
        if (!(true && (0.1 < (-2.25 + (float64((0.0 + 3.14159)) + 1e10))))) {
            for fi12 := 0; fi12 < 19; fi12++ {
                v13 := GenS1{F0: (v2 | v2), F1: uint(v6)}
                if ((false || true) && (1 > 1)) {
                    v14 := GenS2{F0: int64(v6)}
                    var v15 float64 = (((-2.25 * 3.14159) - (0.0 / 0.1)) / (1e10 / float64(float32((-2.25 / 0.1)))))
                    fmt.Println(v2)
                    _ = v14
                    _ = v15
                }
                v2 += v2
                _ = fi12
                _ = v13
            }
            v16 := GenS2{F0: (v2 + 1)}
            _ = v16
        } else {
            if true {
                switch (fi10 << ((0) & 63)) {
                case 935:
                    var v17 uint16 = 15
                    _ = v17
                }
                v18 := GenS1{F0: (v2 | 302), F1: 2}
                _ = v18
            } else {
                fmt.Println(v5, v4)
                v5 = fi10
                var v19 uint = 2
                _ = v19
            }
            switch (("foo bar" + "") + "hello") {
            case "hello":
                if true {
                    v6 = v6
                } else {
                    v6 -= v6
                    v5 = (v5 >> ((v4) & 63))
                    v5 += 412
                }
            default:
                if (42 > 9) {
                    var v20 bool = (!false)
                    _ = v20
                }
            }
            for fi21 := 0; fi21 < 17; fi21++ {
                v3 = append(v3, (fi21 | 520))
                _ = len(v3)
                switch v5 {
                case 4294967295:
                    var v22 int32 = 1
                    _ = v22
                default:
                    v6 += 2147483647
                }
                _ = fi21
            }
        }
        _ = fi10
        _ = v11
    }
    v23 := GenS3{F0: v6}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v4)
    if true {
        var v24 int = v5
        v25 := GenS4{F0: 18446744073709551615}
        _ = v24
        _ = v25
    }
    if false {
        v2 = -418
        v26 := GenS4{F0: v4}
        _ = v26
    } else {
        var v27 float64 = (1e10 - 0.1)
        v28 := GenS2{F0: 296}
        v27 += v27
        _ = v27
        _ = v28
    }
    fmt.Println(v6, v4)
    var v29 uint32 = v6
    if true {
        var v30 int16 = 718
        switch (("\n\t" + "\n\t") + "hello") {
        case "foo bar":
            var v31 uint8 = 128
            _ = v31
        case "\n\t":
            v6 += (v6 | (v29 >> ((42) & 63)))
        default:
            for fi32 := 0; fi32 < 10; fi32++ {
                fmt.Println(v30)
                fmt.Println(v6, v6, v30)
                _ = fi32
            }
        }
        _ = v30
    }
    var v33 uint64 = 42
    v4 -= (v4 % (((v4 << (((v4 + (v4 | v4))) & 63))) | 1))
    var v34 int64 = (v2 & 42)
    v33 = v33
    var v35 int = (v5 << (((^(v4))) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
