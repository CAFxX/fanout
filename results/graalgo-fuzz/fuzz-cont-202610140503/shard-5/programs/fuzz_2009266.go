package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }
func genFn0(p0 uint64, p1 int32) string {
    var v5 uint64 = p0
    fmt.Println(p0, v5)
    _ = p0
    _ = p1
    _ = v5
    return "foo bar"
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{("hello" + "a"), "\n\t", "", "\n\t"}
    var v3 uint = 18446744073709551615
    var v4 string = ""
    fmt.Println(v1, v1, v4)
    fmt.Println(v3, v1)
    switch -9223372036854775808 {
    case -968:
        if true {
            if (2 > 42) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v6 bool = false
                _ = v6
            }
            v2 = append(v2, v4)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    }
    var v7 int16 = 0
    switch "\n\t" {
    case "a":
        v8 := GenS1{F0: 0, F1: 804, F2: 1}
        _ = v8
    case "hello":
        v9 := GenS1{F0: v1, F1: 4294967295, F2: (-(v1))}
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 16; fi10++ {
            v2 = append(v2, v4)
            v1 = (-(v1))
            _ = fi10
        }
    }
    v11 := GenS1{F0: v1, F1: 2, F2: v1}
    v12 := GenS2{F0: 42, F1: (v7 / (((v7 + 32767)) | 1)), F2: (v3 >> (((^(v3))) & 63))}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 string = v4
        _ = fi13
        _ = v14
    }
    v1 = -80
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
