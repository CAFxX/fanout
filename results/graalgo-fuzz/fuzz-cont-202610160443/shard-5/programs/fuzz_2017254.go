package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint }

func main() {
    var v1 uint = 65535
    v2 := []string{}
    var v3 uint16 = 0
    if false {
        if false {
            var v4 uint16 = v3
            _ = v4
        }
    }
    v5 := GenS1{F0: 32768, F1: (v3 >> ((999) & 63)), F2: (v3 << ((v1) & 63))}
    if (("hello" == "hello") && (707 >= 4294967295)) {
        var v6 bool = (0.0 == ((1e10 - (3.14159 + (-2.25 * 0.1))) + 1e10))
        fmt.Println(v3, v1)
        _ = v6
    }
    v3 = v3
    var v7 uint8 = 130
    var v8 uint32 = 516
    if (uint32(v1) <= (v8 | (v8 % ((v8) | 1)))) {
        fmt.Println(v1, v3)
        v1 = (^(v1))
    }
    v9 := GenS2{F0: 0}
    fmt.Println(v3, v3, v3)
    if (!(!((^(v1)) <= v1))) {
        var v10 uint64 = 18446744073709551615
        v7 = v7
        for fi11 := 0; fi11 < 2; fi11++ {
            v12 := GenS3{F0: 42}
            v3 = 2
            _ = fi11
            _ = v12
        }
        _ = v10
    }
    var v13 uint64 = 2
    v7 = v7
    switch "hello" {
    case "hello":
        if ((^(v7)) > v7) {
            var v14 int64 = 42
            if (1e10 == float64(1e10)) {
                v14 = 0
                v2 = append(v2, "foo bar")
                v7 = (v7 - 70)
            }
            var v15 int8 = -107
            _ = v14
            _ = v15
        } else {
            var v16 int32 = -708
            _ = v16
        }
    case "foo bar":
        if false {
            var v17 uint32 = v8
            if (false && true) {
                v7 += (v7 >> (((v1 ^ v1)) & 63))
            } else {
                switch ("\n\t" + (("\n\t" + "a") + ("foo bar" + "a"))) {
                case "\n\t":
                    if (true || (((v1 / ((4294967295) | 1)) <= (v1 & (v1 & 248))) || true)) {
                        for fi18 := 0; fi18 < 8; fi18++ {
                            v13 = (^(v13))
                            v8 = (v8 * (v17 - 740))
                            _ = fi18
                        }
                    }
                case "hello":
                    for fi19 := 0; fi19 < 5; fi19++ {
                        v20 := GenS4{F0: 1, F1: (v1 << (((v1 ^ (v1 << ((v1) & 63)))) & 63))}
                        _ = fi19
                        _ = v20
                    }
                case "foo bar":
                    var v21 float32 = 1e10
                    _ = v21
                }
                var v22 uint = v1
                var v23 int16 = -1
                _ = v22
                _ = v23
            }
            _ = v17
        }
    default:
        if true {
            var v24 int16 = 407
            _ = v24
        } else {
            _ = len(v2)
            fmt.Println(v1, v13)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
