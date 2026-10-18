package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 uint64; F2 uint }
func genFn0(p0 uint8, p1 uint16) uint64 {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := []int{-9223372036854775808, (fi7 + fi7), (^(fi7)), (-(fi7))}
        p1 -= (p1 ^ (p1 ^ p1))
        for fi9 := 0; fi9 < 20; fi9++ {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v10 bool = true
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    p1 -= p1
    p1 = p1
    if true {
        var v11 uint8 = (p0 * (p0 * p0))
        _ = v11
    }
    var v12 int16 = 805
    _ = p0
    _ = p1
    _ = v12
    return 4294967295
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint16 = 42
    v3 := []int{-2147483648, -2147483648}
    v4 := make(map[string]int)
    var v5 string = "a"
    v6 := []string{v5, "a", "\n\t"}
    v2 = 32767
    if true {
        fmt.Println(v2, v5)
        fmt.Println(v5)
    }
    var v13 float32 = (0.1 / -2.25)
    v14 := GenS1{F0: v5}
    v2 = v2
    v13 -= v13
    if true {
        switch "\n\t" {
        case "\n\t":
            switch v5 {
            case "":
                v15 := GenS1{F0: (v5 + "foo bar")}
                _ = v15
            case "hello":
                var v16 int16 = -148
                _ = v16
            }
        case "foo bar":
            v2 = 0
        }
    } else {
        v17 := GenS1{F0: v5}
        var v18 int32 = 0
        _ = v17
        _ = v18
    }
    v2 -= v2
    var v19 uint = 18446744073709551615
    var v20 int64 = -365
    v21 := GenS2{F0: -1, F1: uint64(v19), F2: 156}
    v2 = 1
    switch ((v5 + "foo bar") + v5) {
    case "hello":
        var v22 int16 = 1000
        _ = v22
    case "":
        switch 495 {
        case 2147483647:
            switch v5 {
            case "":
                v23 := GenS1{F0: v5}
                _ = v23
            case "foo bar":
                if (!false) {
                    if (true || (2 >= 42)) {
                        fmt.Println(v19, v20)
                        var v24 bool = (32767 > 2147483647)
                        _ = v24
                    }
                }
            case "\n\t":
                var v25 int8 = 100
                _ = v25
            default:
                v26 := v4[(v5 + ("a" + ""))]
                _ = v26
            }
        }
    case "\n\t":
        v27 := GenS2{F0: -1, F1: 500, F2: v19}
        _ = v27
    default:
        v28 := GenS2{F0: int32(v20), F1: 0, F2: 592}
        _ = v28
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        if (fi29 <= 1) {
            var v30 uint64 = genFn0(23, (v2 | (v2 & 1)))
            _ = v30
        } else {
            switch ("\n\t" + (v5 + v5)) {
            case "a":
                for fi31 := 0; fi31 < 16; fi31++ {
                    fmt.Println(fi31, v2, fi29)
                    _ = fi31
                }
            }
            v13 = -2.25
            v32 := GenS1{F0: ((("hello" + v5) + (v5 + "hello")) + ((((v5 + ((v5 + v5) + (v5 + "a"))) + v5) + v5) + v5))}
            _ = v32
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi29
    }
    var v33 uint8 = 231
    v3 = append(v3, -1)
    v34 := v1[1]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
