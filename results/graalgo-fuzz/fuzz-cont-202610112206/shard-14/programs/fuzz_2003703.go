package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int16; F1 uint; F2 int8 }
type GenS6 struct { F0 int8; F1 int16 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 int64; F1 int64 }
func genFn0(p0 uint64, p1 uint16, p2 bool) uint64 {
    p2 = p2
    if (((1e10 > 0.0) || (((!p2) && (p2 || p2)) || ((-2.25 * ((((1e10 * 0.0) - (0.0 / (3.14159 - -2.25))) - (3.14159 * -2.25)) * 1e10)) < (1e10 * (float32(0.0) / 1e10))))) && p2) {
        for fi7 := 0; fi7 < 10; fi7++ {
            if (!("" > "\n\t")) {
                v8 := []int{128, (-(fi7))}
                v9 := GenS1{F0: 89, F1: 1}
                _ = v8
                _ = v9
            }
            _ = fi7
        }
        var v10 uint = 18446744073709551615
        v11 := GenS2{F0: 0}
        _ = v10
        _ = v11
    } else {
        fmt.Println(p1, p2)
    }
    var v12 uint64 = (p0 + 18446744073709551615)
    var v13 float32 = (float32(((1.5 + (0.1 - 3.14159)) - 1e10)) - 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return p0
}
func genFn1(p0 int8, p1 uint32) uint64 {
    p1 -= 65535
    v14 := []int{4294967295, 4294967295, 497, -2147483648}
    _ = p0
    _ = p1
    _ = v14
    return 2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[string]string)
    var v3 uint64 = 913
    v4 := []int{100}
    var v5 int64 = -1
    v6 := make(map[string]uint32)
    fmt.Println(v3, v3)
    switch (("foo bar" + (("a" + "hello") + "hello")) + (("\n\t" + ("a" + "")) + ("hello" + "hello"))) {
    case "hello":
        var v15 int16 = -964
        _ = v15
    case "\n\t":
        _ = len(v6)
    default:
        for fi16 := 0; fi16 < 3; fi16++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi16
        }
    }
    v17 := GenS2{F0: 32767}
    for fi18 := 0; fi18 < 1; fi18++ {
        v3 = genFn1(42, 1)
        if true {
            fmt.Println(v5, fi18, fi18)
            switch (fi18 | (^(fi18))) {
            case -482:
                v19 := GenS3{F0: (fi18 + 127), F1: false}
                _ = v19
            case 4294967295:
                v3 = 509
            case -431:
                switch "foo bar" {
                case "\n\t":
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case "a":
                    var v20 float32 = -2.25
                    _ = v20
                default:
                    v3 = v3
                }
            default:
                if true {
                    if true {
                        delete(v1, 1)
                        switch (("" + "\n\t") + ("" + "foo bar")) {
                        case "foo bar":
                            v5 -= v5
                        case "a":
                            v5 -= v5
                        default:
                            v21 := v6["a"]
                            _ = v21
                        }
                        var v22 int32 = 42
                        _ = v22
                    }
                    var v23 int8 = 0
                    _ = v23
                }
            }
            v24 := GenS2{F0: -695}
            _ = v24
        }
        if true {
            v25 := GenS4{F0: (^(v3))}
            for fi26 := 0; fi26 < 17; fi26++ {
                v3 = 0
                fmt.Println(v5)
                _ = fi26
            }
            _ = v25
        } else {
            v27 := GenS2{F0: -850}
            v3 = v3
            v28 := GenS5{F0: 1, F1: 0, F2: -123}
            _ = v27
            _ = v28
        }
        _ = fi18
    }
    if true {
        _ = len(v6)
        _ = len(v4)
    } else {
        for fi29 := 0; fi29 < 16; fi29++ {
            var v30 int16 = 42
            _ = fi29
            _ = v30
        }
        v31 := GenS6{F0: 73, F1: 688}
        v32 := GenS7{F0: 1, F1: 9223372036854775807}
        _ = v31
        _ = v32
    }
    v6["\n\t"] = 1
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if (!true) {
        v33 := GenS3{F0: 9223372036854775807, F1: (!false)}
        _ = v33
    }
    v34 := GenS8{F0: v5, F1: (v5 - (v5 & v5))}
    var v35 uint = 479
    v35 = 204
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
