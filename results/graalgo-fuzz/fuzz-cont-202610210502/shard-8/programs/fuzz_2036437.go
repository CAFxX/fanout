package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int8; F2 int32 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 int32 = 1
    v2 := make(map[string]bool)
    var v3 bool = ((310 != 2) && (!(!((!false) || ((0.0 - -2.25) == (1.5 + -2.25))))))
    v4 := make(map[int]int)
    fmt.Println(v3)
    var v5 bool = true
    if v3 {
        v5 = true
        var v6 uint = 65535
        _ = v6
    } else {
        switch 2147483647 {
        case -9223372036854775808:
            switch "hello" {
            case "a":
                v7 := GenS1{F0: 2147483647, F1: 2}
                _ = v7
            }
        case 32767:
            v8 := GenS2{F0: 1, F1: 0, F2: (v1 ^ v1)}
            _ = v8
        case 9223372036854775807:
            var v9 float64 = -2.25
            _ = v9
        default:
            v10 := GenS3{F0: 1}
            _ = v10
        }
        var v11 float64 = ((0.1 * float64(1.5)) - 0.1)
        var v12 uint32 = 4294967295
        _ = v11
        _ = v12
    }
    v5 = (!(!(true && ((0 < 2) || ((-2.25 / 0.1) <= 3.14159)))))
    switch 2147483647 {
    case 32767:
        var v13 uint64 = 18446744073709551615
        _ = v13
    case -757:
        fmt.Println(v1)
    }
    fmt.Println(v1, v3, v5)
    switch "foo bar" {
    case "hello":
        v1 += (v1 ^ v1)
    case "a":
        var v14 uint64 = 94
        _ = v14
    case "\n\t":
        v3 = (((v1 & v1) == (v1 / (((-(v1))) | 1))) && true)
    }
    if (!(!((v3 || (2 > 65535)) || ((((!(81 > 127)) && ((!(float32((1.5 / 3.14159)) >= ((0.1 * 3.14159) / (1e10 * ((3.14159 + 0.0) * 3.14159))))) && (false && (0 > 1)))) || true) || v5)))) {
        var v15 float64 = 0.1
        _ = v15
    }
    delete(v2, ("" + ""))
    _ = len(v2)
    if (0 != 127) {
        var v16 float32 = (float32(0.1) + float32(0.1))
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
