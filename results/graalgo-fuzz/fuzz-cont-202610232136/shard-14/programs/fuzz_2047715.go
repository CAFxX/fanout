package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint16 = 492
    v2 := []string{"\n\t", "\n\t"}
    v3 := make(map[string]bool)
    var v4 bool = (false && (18446744073709551615 > 42))
    v5 := GenS1{F0: 0, F1: "", F2: ""}
    switch 128 {
    case 601:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        fmt.Println(v1)
    }
    switch (("hello" + "") + "\n\t") {
    case "foo bar":
        switch "a" {
        case "":
            var v6 int16 = 1
            _ = v6
        case "foo bar":
            if true {
                v7 := GenS2{F0: 42, F1: 824, F2: 1}
                var v8 bool = v4
                _ = v7
                _ = v8
            }
        default:
            var v9 string = ("a" + ("a" + "hello"))
            _ = v9
        }
    case "a":
        v10 := GenS1{F0: 42, F1: "a", F2: ("\n\t" + ("a" + ""))}
        _ = v10
    case "\n\t":
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(v4, v1, v1)
            v12 := GenS3{F0: (^(v1))}
            _ = fi11
            _ = v12
        }
    }
    var v13 uint32 = 2147483647
    switch "foo bar" {
    case "hello":
        v13 += (v13 + (^(v13)))
    }
    v4 = (!(((0.1 > -2.25) || false) || v4))
    v1 += (^(v1))
    v1 = (v1 ^ (v1 % ((v1) | 1)))
    var v14 string = ("" + "foo bar")
    v13 = (^(v13))
    if (int8(v1) != 82) {
        _ = len(v2)
    }
    var v15 float32 = 1e10
    for fi16 := 0; fi16 < 10; fi16++ {
        if len(v2) > 0 {
            v2[0] = v14
        }
        _ = fi16
    }
    v17 := GenS2{F0: 1, F1: (v13 + v13), F2: 606}
    v13 = 53
    if (v4 && (-733 != 0)) {
        fmt.Println(v1)
        var v18 float32 = v15
        _ = v18
    } else {
        if (-32768 <= 0) {
            var v19 int = 906
            fmt.Println(v1, v1)
            _ = v19
        }
        var v20 int = -9223372036854775808
        _ = v20
    }
    fmt.Println(v1, v13)
    v4 = ((v15 == float32((float64(v15) - -2.25))) || true)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
