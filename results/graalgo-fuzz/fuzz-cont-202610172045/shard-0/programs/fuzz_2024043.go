package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 int }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint; F1 int32; F2 int32 }
type GenS4 struct { F0 string; F1 int16; F2 int16 }
func genFn0(p0 int, p1 int16, p2 uint16) string {
    if (!(4294967295 <= 2)) {
        v5 := []int{p0, (p0 * p0), 717}
        switch ("\n\t" + "") {
        case "":
            for fi6 := 0; fi6 < 17; fi6++ {
                for fi7 := 0; fi7 < 20; fi7++ {
                    p1 += p1
                    p2 -= p2
                    _ = fi7
                }
                p1 += (p1 % ((1) | 1))
                for fi8 := 0; fi8 < 3; fi8++ {
                    v9 := GenS1{F0: (p2 - (^(p2))), F1: 702, F2: p0}
                    v10 := GenS1{F0: (p2 & p2), F1: 2, F2: 42}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = fi6
            }
        case "\n\t":
            p2 -= p2
        }
        _ = v5
    } else {
        var v11 uint = 18446744073709551615
        _ = v11
    }
    switch "a" {
    case "foo bar":
        if false {
            for fi12 := 0; fi12 < 20; fi12++ {
                fmt.Println(p2)
                switch 0 {
                case -1:
                    var v13 int64 = -1
                    _ = v13
                case 4294967295:
                    var v14 int = (fi12 >> ((1) & 63))
                    _ = v14
                }
                _ = fi12
            }
        } else {
            var v15 float32 = (((-2.25 * (0.0 * 1.5)) - 0.0) / (3.14159 + 1.5))
            var v16 uint32 = 2
            _ = v15
            _ = v16
        }
    case "":
        v17 := GenS2{F0: 2, F1: p0}
        _ = v17
    default:
        fmt.Println(p1, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return ("a" + (((((("hello" + "foo bar") + "") + "foo bar") + "\n\t") + ((((("foo bar" + "") + ("foo bar" + ("\n\t" + "a"))) + "foo bar") + ("hello" + ("a" + (((("a" + (("\n\t" + "foo bar") + "hello")) + "hello") + "hello") + "a")))) + "a")) + "hello"))
}
func genFn1(p0 uint64, p1 int16, p2 bool) uint32 {
    var v18 int = 461
    v19 := GenS2{F0: 42, F1: 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return 2147483648
}

func main() {
    var v1 uint8 = 128
    v2 := []string{"\n\t", "hello", "a"}
    var v3 int32 = -349
    v4 := make(map[string]int)
    v3 = (v3 << ((65535) & 63))
    if (!false) {
        var v20 float64 = 0.0
        v20 += (float64(((float32(-2.25) + (-2.25 + 1e10)) - (3.14159 - 0.1))) - v20)
        _ = v20
    } else {
        fmt.Println(v1, v1)
        var v21 int32 = (v3 & v3)
        _ = v21
    }
    v22 := GenS1{F0: 337, F1: 2, F2: 127}
    for fi23 := 0; fi23 < 14; fi23++ {
        if true {
            v1 = v1
            v2 = append(v2, ("a" + "hello"))
            switch "\n\t" {
            case "":
                var v24 int64 = 9223372036854775807
                _ = v24
            }
        } else {
            var v25 int64 = 9223372036854775807
            _ = v25
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            fmt.Println(v1)
            _ = fi26
        }
        _ = fi23
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        v1 -= v1
        _ = fi27
    }
    v28 := GenS2{F0: v1, F1: 988}
    for fi29 := 0; fi29 < 5; fi29++ {
        fmt.Println(v1, v1)
        v1 -= v1
        _ = len(v4)
        _ = fi29
    }
    fmt.Println(v1)
    v30 := GenS1{F0: 65535, F1: 65535, F2: -780}
    v1 += 21
    if true {
        for fi31 := 0; fi31 < 17; fi31++ {
            switch "hello" {
            case "foo bar":
                fmt.Println(fi31, v3, fi31)
            case "a":
                fmt.Println(v3, v1)
            default:
                v1 += 128
            }
            _ = fi31
        }
        for fi32 := 0; fi32 < 8; fi32++ {
            if false {
                var v33 bool = true
                _ = v33
            } else {
                v3 = v3
                v34 := GenS2{F0: v1, F1: (fi32 ^ -1)}
                _ = v34
            }
            fmt.Println(v3, fi32)
            _ = fi32
        }
        var v35 uint16 = 42
        _ = v35
    }
    v3 = v3
    v36 := GenS3{F0: 4294967295, F1: 443, F2: v3}
    v37 := GenS4{F0: "", F1: int16(v1), F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v22
    _ = v28
    _ = v3
    _ = v30
    _ = v36
    _ = v37
    _ = v4
}
