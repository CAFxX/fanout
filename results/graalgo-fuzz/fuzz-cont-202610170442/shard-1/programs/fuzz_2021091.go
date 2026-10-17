package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int32; F1 bool; F2 uint16 }
func genFn0(p0 uint16) uint32 {
    for fi7 := 0; fi7 < 11; fi7++ {
        var v8 uint = 42
        p0 = (p0 % ((p0) | 1))
        if ((^(p0)) <= 834) {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := GenS1{F0: 27}
                var v11 float32 = 0.1
                _ = fi9
                _ = v10
                _ = v11
            }
            switch ("a" + "hello") {
            case "a":
                var v12 bool = (2 >= 2)
                _ = v12
            default:
                for fi13 := 0; fi13 < 14; fi13++ {
                    p0 -= (p0 + (p0 ^ p0))
                    var v14 int = fi13
                    _ = fi13
                    _ = v14
                }
            }
        }
        _ = fi7
        _ = v8
    }
    v15 := []int{}
    var v16 uint8 = 255
    if len(v15) > 0 {
        v15[0] = -1
    }
    _ = p0
    _ = v15
    _ = v16
    return 2
}

func main() {
    v1 := []uint32{0, 2}
    var v2 int16 = 1000
    var v3 uint64 = 42
    var v4 int64 = 9223372036854775807
    v5 := []string{"", "\n\t", "\n\t"}
    var v6 int32 = 0
    if (v3 <= (^(v3))) {
        var v17 uint32 = 565
        _ = v17
    } else {
        switch "foo bar" {
        case "a":
            fmt.Println(v2, v3, v4)
        case "\n\t":
            v1 = append(v1, 2147483647)
        }
        var v18 int = 42
        switch (v18 ^ -518) {
        case 690:
            if true {
                v3 = (v3 << ((65535) & 63))
            }
        case -463:
            fmt.Println(v2)
        case -2147483648:
            if false {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                switch 1 {
                case 268:
                    v3 = (v3 - v3)
                }
            }
        }
        _ = v18
    }
    fmt.Println(v3)
    v3 += v3
    v19 := GenS2{F0: (-(v6)), F1: (-78 <= 8), F2: 65535}
    if (128 != 127) {
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS2{F0: 733, F1: false, F2: 1}
            fmt.Println(v4, v3)
            var v22 int32 = (v6 | (v6 + v6))
            _ = fi20
            _ = v21
            _ = v22
        }
        switch (("\n\t" + "a") + "\n\t") {
        case "":
            var v23 bool = true
            _ = v23
        case "a":
            if ((!(1 != 42)) || (!(!false))) {
                v2 -= (v2 | 42)
            }
        case "\n\t":
            v6 += -540
        default:
            for fi24 := 0; fi24 < 20; fi24++ {
                fmt.Println(fi24)
                v3 -= v3
                _ = fi24
            }
        }
    } else {
        if (false || false) {
            if false {
                var v25 int32 = v6
                v1 = append(v1, 1)
                fmt.Println(v3, v2)
                _ = v25
            }
            if len(v1) > 0 {
                v1[0] = genFn0(1)
            }
            if true {
                v1 = append(v1, genFn0(2))
                var v26 uint8 = 127
                if (!(!(-2.25 != 1e10))) {
                    switch "" {
                    case "foo bar":
                        var v27 uint8 = (v26 % ((127) | 1))
                        _ = v27
                    default:
                        var v28 int = 32768
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 18; fi29++ {
                        var v30 int32 = (-(v6))
                        var v31 float32 = (float32(float64(3.14159)) * (0.1 + (((3.14159 + float32(-2.25)) * 1.5) * 0.1)))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                } else {
                    v2 = v2
                    v26 -= v26
                }
                _ = v26
            } else {
                fmt.Println(v3, v6, v3)
                _ = len(v1)
            }
        }
        v4 = (v4 / (((v4 + v4)) | 1))
    }
    var v32 int32 = v6
    v33 := GenS2{F0: (-(v32)), F1: (32767 > 433), F2: 442}
    var v34 int = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
