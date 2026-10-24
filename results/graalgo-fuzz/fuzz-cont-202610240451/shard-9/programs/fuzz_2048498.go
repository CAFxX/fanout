package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int; F1 uint16 }
func genFn0() int32 {
    switch -65 {
    case 258:
        var v7 uint64 = 2
        _ = v7
    default:
        var v8 int64 = -9223372036854775808
        _ = v8
    }
    var v9 uint8 = 1
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 string = "foo bar"
        _ = fi10
        _ = v11
    }
    _ = v9
    return 2147483647
}
func genFn1(p0 uint64) int8 {
    v12 := GenS1{F0: 0, F1: 757}
    v13 := []int{4294967295, 32768}
    if (!("\n\t" > "")) {
        v14 := GenS1{F0: 1, F1: -2147483648}
        var v15 uint16 = 783
        _ = v14
        _ = v15
    } else {
        var v16 string = ""
        for fi17 := 0; fi17 < 15; fi17++ {
            for fi18 := 0; fi18 < 1; fi18++ {
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                p0 -= 18446744073709551615
                _ = fi18
            }
            _ = fi17
        }
        _ = v16
    }
    _ = p0
    _ = v12
    _ = v13
    return -63
}

func main() {
    v1 := []string{"a"}
    var v2 int = 9223372036854775807
    var v3 int16 = 42
    v4 := []string{""}
    v5 := []int{(v2 >> ((65535) & 63))}
    v6 := []string{}
    fmt.Println(v3, v2)
    fmt.Println(v3)
    var v19 int8 = 127
    fmt.Println(v19, v19)
    v20 := GenS2{F0: (v2 >> ((255) & 63)), F1: 0}
    var v21 int = v2
    fmt.Println(v3, v19)
    v19 += (v19 / ((genFn1(2)) | 1))
    v21 = (-(v2))
    v3 -= v3
    v19 = (v19 | (^(v19)))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
