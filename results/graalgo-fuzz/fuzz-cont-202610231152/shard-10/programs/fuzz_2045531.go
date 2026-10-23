package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string }
type GenS2 struct { F0 int32; F1 int64; F2 int64 }

func main() {
    v1 := []uint32{65535}
    v2 := []string{"", ("a" + "hello"), ("foo bar" + (("hello" + ("" + "hello")) + ""))}
    var v3 float64 = 0.1
    fmt.Println(v3)
    if true {
        switch "a" {
        case "foo bar":
            v4 := GenS1{F0: -3, F1: "hello"}
            _ = v4
        default:
            v3 = 0.1
        }
    } else {
        v5 := GenS2{F0: 0, F1: -285, F2: 0}
        v6 := GenS2{F0: -2147483648, F1: 470, F2: 9223372036854775807}
        _ = v5
        _ = v6
    }
    fmt.Println(v3)
    v3 -= v3
    if ((128 >= 59) || (896 > 255)) {
        var v7 uint8 = 1
        if (437 == 2147483647) {
            var v8 uint8 = v7
            if false {
                v1 = append(v1, 42)
                v3 -= v3
            }
            for fi9 := 0; fi9 < 20; fi9++ {
                v7 += 171
                _ = fi9
            }
            _ = v8
        }
        _ = v7
    } else {
        v10 := GenS1{F0: -1, F1: "\n\t"}
        _ = v10
    }
    v3 = float64((1.5 - float32((float64(0.0) - v3))))
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += float64(1e10)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v3)
    for fi11 := 0; fi11 < 2; fi11++ {
        if ((!true) && false) {
            var v12 int64 = -9223372036854775808
            _ = v12
        } else {
            v13 := GenS2{F0: 2147483647, F1: 0, F2: 0}
            _ = v13
        }
        v14 := GenS2{F0: 42, F1: -713, F2: -805}
        _ = fi11
        _ = v14
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v3
}
