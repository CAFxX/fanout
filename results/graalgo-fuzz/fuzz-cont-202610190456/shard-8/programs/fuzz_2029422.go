package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 65535
    var v3 string = "\n\t"
    v2 -= v2
    v4 := v1[(((("foo bar" + (v3 + (v3 + v3))) + (("foo bar" + v3) + "hello")) + v3) + ((v3 + v3) + "\n\t"))]
    if v4 {
        v5 := []int{696, -441, 4294967295, 32768}
        v2 += v2
        fmt.Println(v4, v2)
        _ = v5
    } else {
        if v4 {
            var v6 uint64 = 1
            if true {
                fmt.Println(v2, v2)
                if v4 {
                    fmt.Println(v4)
                    var v7 float32 = 0.1
                    _ = v7
                }
                fmt.Println(v4, v3, v4)
            } else {
                v1[v3] = v4
            }
            v8 := GenS1{F0: (v2 < v2), F1: (v2 << ((4294967295) & 63))}
            _ = v6
            _ = v8
        } else {
            var v9 string = v3
            _ = v9
        }
    }
    v2 += v2
    v10 := []int{785, -9223372036854775808, -2147483648}
    for fi11 := 0; fi11 < 15; fi11++ {
        var v12 uint64 = 1
        var v13 uint8 = 1
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := GenS2{F0: v3}
    v1[v3] = true
    for fi15 := 0; fi15 < 9; fi15++ {
        v3 = v3
        delete(v1, "\n\t")
        switch "foo bar" {
        case "":
            v16 := GenS1{F0: v4, F1: (^(v2))}
            _ = v16
        case "\n\t":
            v17 := GenS3{F0: ("a" + v3), F1: 42, F2: 255}
            _ = v17
        case "a":
            v4 = (923 == 1)
        }
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
