package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 uint32 }
type GenS3 struct { F0 int8; F1 uint64 }

func main() {
    var v1 float32 = float32((0.1 - 0.1))
    v2 := []string{"\n\t", "a", ("foo bar" + ("" + "a"))}
    var v3 int32 = -354
    var v4 float64 = 1e10
    var v5 uint32 = 2147483647
    v6 := make(map[string]int)
    fmt.Println(v5)
    var v7 int32 = (v3 * v3)
    v8 := v6["\n\t"]
    if false {
        switch v8 {
        case 0:
            switch v8 {
            case 127:
                if ((((float64(float32((v4 * v4))) * ((v4 * -2.25) + v4)) / (1e10 + float64((v1 - v1)))) - (v4 * 3.14159)) == v4) {
                    for fi9 := 0; fi9 < 9; fi9++ {
                        var v10 uint = 539
                        v1 -= (v1 + v1)
                        fmt.Println(v5)
                        _ = fi9
                        _ = v10
                    }
                    var v11 uint64 = 0
                    v5 = v5
                    _ = v11
                }
            case 673:
                fmt.Println(v3, v8, v4)
            case 42:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                v4 = (float64(3.14159) * (0.1 + (v4 * 0.0)))
            }
        case 128:
            for fi12 := 0; fi12 < 16; fi12++ {
                v1 = -2.25
                for fi13 := 0; fi13 < 11; fi13++ {
                    var v14 uint64 = 755
                    _ = fi13
                    _ = v14
                }
                _ = fi12
            }
        case 1:
            var v15 uint64 = 18446744073709551615
            _ = v15
        }
        var v16 uint64 = 42
        v17 := GenS1{F0: (true && true), F1: 4294967295}
        _ = v16
        _ = v17
    }
    switch ((("foo bar" + "hello") + ((("hello" + "") + "foo bar") + (("\n\t" + "") + "a"))) + "") {
    case "":
        v18 := GenS1{F0: false, F1: 255}
        _ = v18
    case "foo bar":
        v5 = v5
    }
    v19 := GenS2{F0: -1, F1: (v5 | v5)}
    if (-1 != 42) {
        switch (v8 & int(v3)) {
        case -932:
            switch "foo bar" {
            case "\n\t":
                v20 := GenS3{F0: -1, F1: 42}
                _ = v20
            default:
                fmt.Println(v5)
            }
        default:
            var v21 int8 = 89
            _ = v21
        }
        switch "\n\t" {
        case "\n\t":
            for fi22 := 0; fi22 < 12; fi22++ {
                fmt.Println(fi22, v8, v7)
                _ = fi22
            }
        default:
            if (0 == uint64(v5)) {
                if false {
                    v1 = v1
                }
            } else {
                v5 = (v5 % (((v5 ^ (v5 + (v5 | 358)))) | 1))
                v5 = v5
            }
        }
    }
    var v23 uint = 0
    delete(v6, (("hello" + "\n\t") + ("a" + (("foo bar" + "") + ("\n\t" + "a")))))
    if false {
        v3 += (-(v7))
    } else {
        var v24 int64 = -9223372036854775808
        fmt.Println(v23, v23, v8)
        _ = v24
    }
    v25 := GenS3{F0: 127, F1: 2}
    fmt.Println(v23)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
