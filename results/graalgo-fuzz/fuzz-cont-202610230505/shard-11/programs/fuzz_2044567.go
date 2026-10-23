package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 bool }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0(p0 bool, p1 bool, p2 uint32) int {
    fmt.Println(p0)
    var v6 uint16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 2147483647
}

func main() {
    var v1 int = -29
    v2 := []string{"\n\t", ("foo bar" + "foo bar"), ""}
    var v3 uint32 = 2
    v4 := make(map[string]bool)
    var v5 uint16 = 178
    var v7 int8 = 100
    var v8 uint = 947
    v2 = append(v2, "")
    v1 = -994
    var v9 uint16 = v5
    switch "\n\t" {
    case "":
        switch ("foo bar" + "a") {
        case "foo bar":
            switch (^(v1)) {
            case 32767:
                v7 += 100
            case 2147483647:
                v7 -= (v7 * int8(v9))
            default:
                fmt.Println(v7, v3)
            }
        }
    case "foo bar":
        v8 = (v8 % ((2) | 1))
    default:
        var v10 int32 = 42
        _ = v10
    }
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    fmt.Println(v9, v7)
    var v11 uint32 = (v3 % ((2) | 1))
    v12 := GenS1{F0: -2147483648, F1: true, F2: false}
    v13 := GenS2{F0: 0, F1: (v7 | int8(v1))}
    var v14 int = v1
    _ = len(v2)
    switch (v1 | (v14 - (v14 ^ 0))) {
    case 32768:
        v15 := GenS1{F0: -513, F1: false, F2: (42 < int64(v8))}
        _ = v15
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        if (!((!false) && (uint64(v5) <= uint64(v5)))) {
            for fi17 := 0; fi17 < 6; fi17++ {
                var v18 float32 = 1e10
                v18 = v18
                v11 = v3
                _ = fi17
                _ = v18
            }
            v11 += v11
            for fi19 := 0; fi19 < 8; fi19++ {
                fmt.Println(v14, v9, v11)
                v20 := v4["foo bar"]
                v3 += v3
                _ = fi19
                _ = v20
            }
        }
        var v21 int64 = -9223372036854775808
        v22 := GenS3{F0: (v7 | 127), F1: 42}
        _ = fi16
        _ = v21
        _ = v22
    }
    v23 := GenS3{F0: v7, F1: v5}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
