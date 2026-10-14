package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int64 }

func main() {
    v1 := []int{42, 128}
    v2 := []string{"foo bar", "\n\t", "", ("foo bar" + "foo bar")}
    var v3 float64 = 0.0
    var v4 int8 = 127
    v1 = append(v1, 4294967295)
    v3 = v3
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v5 := GenS1{F0: "\n\t", F1: v4}
    fmt.Println(v3)
    if true {
        if false {
            switch 128 {
            case 32767:
                v6 := GenS1{F0: "hello", F1: (v4 % (((v4 | 65)) | 1))}
                _ = v6
            case 26:
                v7 := GenS2{F0: 236}
                _ = v7
            case 2147483647:
                var v8 uint64 = 0
                _ = v8
            }
        } else {
            v4 = 1
            var v9 uint = 927
            _ = v9
        }
    }
    v10 := GenS2{F0: 0}
    _ = len(v2)
    if ((!(411 < 686)) || true) {
        var v11 int = -859
        v3 -= 1.5
        v11 -= (-(v11))
        _ = v11
    }
    v4 = v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
