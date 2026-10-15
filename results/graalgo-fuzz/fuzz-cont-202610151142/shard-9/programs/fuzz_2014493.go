package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8; F2 int16 }
type GenS2 struct { F0 int64; F1 string; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 int64 }
type GenS5 struct { F0 int64; F1 uint8; F2 bool }
func genFn0(p0 int) string {
    p0 -= p0
    for fi4 := 0; fi4 < 4; fi4++ {
        p0 = -192
        fmt.Println(fi4)
        _ = fi4
    }
    v5 := GenS1{F0: 42, F1: 143, F2: 676}
    _ = p0
    _ = v5
    return ("" + "foo bar")
}
func genFn1(p0 int64, p1 uint64, p2 bool) int16 {
    p0 -= p0
    p1 -= p1
    p2 = p2
    var v6 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -527
}

func main() {
    var v1 bool = true
    v2 := []string{"", "a"}
    v3 := make(map[string]uint32)
    fmt.Println(v1)
    fmt.Println(v1)
    for fi7 := 0; fi7 < 18; fi7++ {
        switch (fi7 << ((1) & 63)) {
        case -9223372036854775808:
            v1 = (-2.25 > 3.14159)
        case 4294967295:
            fmt.Println(v1, v1)
        case 9223372036854775807:
            v1 = (!true)
        default:
            if (("a" >= ("hello" + "")) && (float32((0.0 / (1e10 / -2.25))) > 1.5)) {
                for fi8 := 0; fi8 < 3; fi8++ {
                    fmt.Println(fi8, v1, fi7)
                    v1 = (!(-2147483648 <= 313))
                    _ = fi8
                }
                v1 = (1 >= -705)
            } else {
                v9 := GenS2{F0: 9223372036854775807, F1: "", F2: -5}
                _ = v9
            }
        }
        v10 := GenS3{F0: 0, F1: 0}
        v11 := v3[("" + ("hello" + (("\n\t" + (((("" + "\n\t") + "\n\t") + ("a" + "\n\t")) + ("\n\t" + ("" + "\n\t")))) + "\n\t")))]
        _ = fi7
        _ = v10
        _ = v11
    }
    var v12 uint = 530
    v13 := GenS2{F0: -9223372036854775808, F1: "\n\t", F2: -63}
    var v14 uint8 = 255
    var v15 float64 = ((1e10 / 1e10) + 1.5)
    for fi16 := 0; fi16 < 8; fi16++ {
        if (65535 >= uint32(v14)) {
            v12 -= (v12 >> ((v12) & 63))
        } else {
            switch "a" {
            case "\n\t":
                var v17 string = ("hello" + "\n\t")
                _ = v17
            }
            var v18 float32 = (0.1 + 3.14159)
            _ = v18
        }
        switch (fi16 >> ((v12) & 63)) {
        case 32767:
            delete(v3, (((("hello" + "hello") + "a") + "") + ("a" + "foo bar")))
        case 128:
            for fi19 := 0; fi19 < 9; fi19++ {
                fmt.Println(v1, v12, v12)
                var v20 float32 = 1.5
                var v21 int64 = 1
                var v22 uint16 = 627
                var v23 string = "\n\t"
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
                _ = v23
            }
        }
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(v14)
            v14 += v14
            _ = fi24
        }
        _ = fi16
    }
    v25 := GenS4{F0: 255, F1: 1}
    for fi26 := 0; fi26 < 13; fi26++ {
        v12 += v12
        var v27 int = (fi26 << ((65535) & 63))
        _ = fi26
        _ = v27
    }
    v1 = v1
    v28 := GenS5{F0: 1, F1: (v14 % ((v14) | 1)), F2: (!((32767 > 255) && true))}
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v28
    _ = v3
}
