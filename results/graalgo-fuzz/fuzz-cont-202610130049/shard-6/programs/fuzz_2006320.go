package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int }
type GenS2 struct { F0 int16; F1 uint8; F2 int8 }
type GenS3 struct { F0 int; F1 int8 }

func main() {
    v1 := []string{("hello" + ("a" + ""))}
    v2 := make(map[uint32]string)
    var v3 int8 = 55
    var v4 uint64 = 42
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: (^(v4)), F1: (-(fi5))}
        switch "" {
        case "\n\t":
            v4 = (^(v4))
        case "a":
            fmt.Println(v4, v3, v4)
        case "hello":
            v7 := GenS1{F0: v4, F1: fi5}
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    switch -851 {
    case 127:
        var v8 uint32 = 65535
        _ = v8
    case -2147483648:
        if (1000 > -32768) {
            if (-251 == -1) {
                switch -34 {
                case -9223372036854775808:
                    v3 += v3
                }
                var v9 uint64 = (v4 * v4)
                _ = v9
            }
            var v10 int64 = -9223372036854775808
            _ = v10
        } else {
            v3 -= v3
        }
    case -538:
        var v11 int16 = -1
        _ = v11
    default:
        v1 = append(v1, "foo bar")
    }
    v4 -= (^(v4))
    if (4294967295 >= 2) {
        var v12 uint8 = 1
        _ = v12
    }
    v3 += -85
    if true {
        if ((9223372036854775807 < -72) && (false && (("\n\t" + "foo bar") > "\n\t"))) {
            v3 = 1
            var v13 string = ""
            var v14 int32 = 0
            _ = v13
            _ = v14
        }
    } else {
        v15 := GenS2{F0: 42, F1: 255, F2: (v3 ^ 100)}
        _ = v15
    }
    v4 += 251
    var v16 int64 = -9223372036854775808
    for fi17 := 0; fi17 < 5; fi17++ {
        v16 -= -1
        _ = fi17
    }
    var v18 float32 = 1e10
    v19 := GenS3{F0: -75, F1: 17}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
