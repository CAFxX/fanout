package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    v1 := make(map[int]uint32)
    var v2 string = "foo bar"
    var v3 int32 = 0
    v4 := []string{"", ""}
    v5 := make(map[int]uint32)
    var v6 int64 = 9223372036854775807
    v7 := GenS1{F0: 42, F1: 917}
    v8 := GenS1{F0: 255, F1: 811}
    var v9 uint32 = 2147483647
    switch v2 {
    case "":
        v10 := GenS1{F0: 255, F1: v9}
        _ = v10
    case "hello":
        for fi11 := 0; fi11 < 7; fi11++ {
            if ((!(0.0 != 0.1)) && (true && false)) {
                if len(v4) > 0 {
                    v4[0] = (v2 + (v2 + v2))
                }
            } else {
                var v12 float64 = 1e10
                v13 := GenS1{F0: 255, F1: v9}
                _ = v12
                _ = v13
            }
            fmt.Println(v6, fi11, v9)
            _ = fi11
        }
    default:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v14 bool = false
    v6 = 1
    for fi15 := 0; fi15 < 2; fi15++ {
        switch (fi15 << ((18446744073709551615) & 63)) {
        case 32768:
            fmt.Println(v2, v3)
        case 0:
            v14 = v14
        default:
            var v16 uint64 = 1
            _ = v16
        }
        _ = fi15
    }
    if (true && ((v2 + "foo bar") < ((("\n\t" + v2) + v2) + "foo bar"))) {
        _ = len(v4)
    } else {
        _ = len(v1)
    }
    if v14 {
        v3 -= (v3 & (-(v3)))
    }
    switch -327 {
    case 9223372036854775807:
        fmt.Println(v14)
    case 4294967295:
        v17 := GenS2{F0: -1, F1: 297}
        _ = v17
    case 146:
        if ((float32(1.5) * float32(1.5)) < (3.14159 + 1.5)) {
            v3 += v3
            v2 = (("hello" + v2) + v2)
            for fi18 := 0; fi18 < 5; fi18++ {
                v2 = ((v2 + v2) + ("a" + v2))
                fmt.Println(v3, v14)
                if (1 == 42) {
                    var v19 int64 = (v6 - (-(v6)))
                    for fi20 := 0; fi20 < 19; fi20++ {
                        if (42 == 0) {
                            v21 := GenS2{F0: 32767, F1: 2}
                            _ = v21
                        }
                        _ = fi20
                    }
                    fmt.Println(fi18, v9)
                    _ = v19
                }
                _ = fi18
            }
        } else {
            v22 := GenS3{F0: true, F1: 750}
            _ = v22
        }
    default:
        v9 += 42
    }
    var v23 string = "hello"
    v24 := GenS1{F0: 18446744073709551615, F1: (v9 & v9)}
    v25 := GenS3{F0: v14, F1: (^(v9))}
    var v26 int16 = 32767
    var v27 float32 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
