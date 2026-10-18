package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8; F2 int32 }
type GenS2 struct { F0 uint8; F1 uint16 }

func main() {
    var v1 int64 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[string]string)
    v4 := make(map[string]int)
    v5 := []uint32{808, 42, 1}
    var v6 float64 = 0.0
    for fi7 := 0; fi7 < 10; fi7++ {
        switch 141 {
        case 4294967295:
            v8 := GenS1{F0: "foo bar", F1: 78, F2: 322}
            _ = v8
        case 128:
            v1 += v1
        default:
            fmt.Println(v6)
        }
        fmt.Println(v6, fi7, v6)
        _ = fi7
    }
    for fi9 := 0; fi9 < 14; fi9++ {
        for fi10 := 0; fi10 < 14; fi10++ {
            _ = len(v5)
            _ = fi10
        }
        _ = fi9
    }
    var v11 int = -9223372036854775808
    if false {
        var v12 uint = 18446744073709551615
        v11 = v11
        _ = v12
    } else {
        var v13 float32 = (1e10 / 3.14159)
        v11 += v11
        if (true && (v1 < v1)) {
            if true {
                var v14 uint = 0
                _ = v14
            }
        } else {
            v1 -= (v1 * 877)
            if true {
                delete(v4, ("a" + ""))
            } else {
                _ = len(v3)
                if (18446744073709551615 <= 4294967295) {
                    var v15 int = 2147483647
                    _ = v15
                }
            }
            switch v11 {
            case 0:
                var v16 uint8 = 161
                _ = v16
            case 32767:
                var v17 float64 = v6
                _ = v17
            case -15:
                _ = len(v5)
            default:
                var v18 uint64 = 18446744073709551615
                _ = v18
            }
        }
        _ = v13
    }
    v5 = append(v5, 4294967295)
    v6 += v6
    v1 = (v1 | (v1 % (((v1 * v1)) | 1)))
    for fi19 := 0; fi19 < 15; fi19++ {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        var v20 uint16 = 65535
        v21 := GenS2{F0: 35, F1: v20}
        _ = fi19
        _ = v20
        _ = v21
    }
    var v22 string = ("foo bar" + "a")
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v5 = append(v5, 2)
    switch v22 {
    case "":
        v23 := GenS2{F0: 0, F1: 2}
        _ = v23
    case "\n\t":
        delete(v3, (v22 + ((("hello" + ("" + v22)) + "") + "a")))
    case "a":
        fmt.Println(v22, v1, v11)
    default:
        fmt.Println(v22, v1)
    }
    switch (v11 + (-(v11))) {
    case -1:
        var v24 uint32 = 42
        _ = v24
    }
    v25 := GenS1{F0: "", F1: 34, F2: -2147483648}
    var v26 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
