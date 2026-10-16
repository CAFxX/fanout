package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint16, p1 int16, p2 uint) uint8 {
    var v7 uint64 = 18446744073709551615
    p1 -= (p1 - p1)
    v8 := GenS1{F0: 804}
    v9 := GenS1{F0: 2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 128
}
func genFn1() int64 {
    switch 9223372036854775807 {
    case 1:
        var v10 bool = (!(!(42 > -1)))
        _ = v10
    case 32767:
        var v11 int32 = 1
        _ = v11
    case -2147483648:
        var v12 uint32 = 1
        _ = v12
    default:
        v13 := GenS1{F0: 50}
        _ = v13
    }
    var v14 string = ("\n\t" + ("a" + "foo bar"))
    _ = v14
    return 42
}

func main() {
    v1 := []uint32{932, 0, 2147483647}
    var v2 int32 = -1
    v3 := []string{"a", "hello", "\n\t"}
    var v4 uint = 529
    var v5 float32 = 0.1
    var v6 string = ("hello" + ((("foo bar" + "hello") + "foo bar") + (("" + "a") + "")))
    var v15 int32 = v2
    v5 = v5
    _ = len(v1)
    for fi16 := 0; fi16 < 5; fi16++ {
        switch (fi16 << (((v4 ^ v4)) & 63)) {
        case 42:
            v17 := GenS1{F0: 42}
            _ = v17
        case 32767:
            switch (((v6 + "") + v6) + v6) {
            case "foo bar":
                for fi18 := 0; fi18 < 2; fi18++ {
                    if (0 > 18446744073709551615) {
                        var v19 uint16 = 425
                        v2 -= (v15 ^ (v2 | (v2 ^ -133)))
                        v19 = v19
                        _ = v19
                    }
                    _ = fi18
                }
            case "":
                var v20 string = v6
                _ = v20
            case "\n\t":
                for fi21 := 0; fi21 < 13; fi21++ {
                    var v22 int16 = -620
                    if (4294967295 >= 0) {
                        if (1 > -644) {
                            v23 := GenS2{F0: 127, F1: 2}
                            v22 -= v22
                            fmt.Println(v22)
                            _ = v23
                        }
                    }
                    _ = fi21
                    _ = v22
                }
            default:
                v24 := GenS3{F0: fi16}
                _ = v24
            }
        case -356:
            switch v6 {
            case "foo bar":
                var v25 uint32 = 2147483647
                _ = v25
            case "\n\t":
                var v26 uint8 = 42
                _ = v26
            default:
                switch v6 {
                case "\n\t":
                    v2 = v15
                default:
                    var v27 int16 = 32767
                    _ = v27
                }
            }
        }
        _ = fi16
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 15; fi29++ {
            var v30 int = fi28
            if len(v1) > 0 {
                v1[0] = 136
            }
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    var v31 float64 = (float64(v5) - ((1e10 + (3.14159 * ((0.0 * (0.0 * 1e10)) / (float64((0.0 - 1.5)) * 3.14159)))) + ((1e10 * 1e10) / 1.5)))
    var v32 uint16 = 566
    fmt.Println(v15, v5, v31)
    switch v6 {
    case "hello":
        var v33 uint8 = 255
        _ = v33
    }
    var v34 uint16 = (^(v32))
    var v35 uint64 = 18446744073709551615
    v36 := GenS3{F0: 0}
    if len(v1) > 0 {
        v1[0] = 401
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
