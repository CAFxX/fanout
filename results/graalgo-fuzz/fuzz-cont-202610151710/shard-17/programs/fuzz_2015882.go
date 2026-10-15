package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int64, p1 string) string {
    switch p1 {
    case "hello":
        p0 -= (p0 / (((p0 + 9223372036854775807)) | 1))
    case "\n\t":
        v4 := GenS1{F0: 650, F1: 128, F2: -2147483648}
        _ = v4
    default:
        p0 = (^(p0))
    }
    var v5 uint = 18446744073709551615
    _ = p0
    _ = p1
    _ = v5
    return ((("hello" + ((p1 + p1) + ((p1 + p1) + (p1 + p1)))) + ((p1 + (p1 + "")) + p1)) + p1)
}

func main() {
    v1 := []string{"hello"}
    var v2 int8 = 100
    var v3 int8 = v2
    _ = len(v1)
    switch 32768 {
    case 2147483647:
        v2 -= v3
    case 4294967295:
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    }
    v3 += (-(v2))
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int64 = 1
            fmt.Println(v2, fi7)
            _ = fi7
            _ = v8
        }
        _ = fi6
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3, v2)
    for fi9 := 0; fi9 < 9; fi9++ {
        v10 := GenS2{F0: 127}
        _ = fi9
        _ = v10
    }
    v3 += (v3 % ((-18) | 1))
    v2 = (^(v3))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
