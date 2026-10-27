package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }

func main() {
    v1 := make(map[int]bool)
    var v2 int = 0
    v3 := []uint32{2147483648, 666, 0}
    v4 := []uint32{2, 829, 454, 4294967295}
    v5 := []int{}
    var v6 int32 = -2147483648
    v6 -= (v6 >> ((445) & 63))
    fmt.Println(v6, v6)
    v6 += (v6 % ((v6) | 1))
    v7 := v1[v2]
    v8 := GenS1{F0: "\n\t", F1: 794}
    switch (((("a" + ("" + "a")) + "hello") + "hello") + "\n\t") {
    case "":
        fmt.Println(v6, v2, v7)
    case "\n\t":
        if v7 {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if v7 {
                fmt.Println(v7, v7)
                v9 := GenS1{F0: "hello", F1: 65535}
                switch "hello" {
                case "":
                    v5 = append(v5, (^(v2)))
                case "a":
                    var v10 bool = v7
                    _ = v10
                }
                _ = v9
            } else {
                var v11 float64 = -2.25
                _ = v11
            }
            var v12 string = "\n\t"
            _ = v12
        }
    default:
        var v13 int64 = 9223372036854775807
        _ = v13
    }
    var v14 string = (((((("" + "foo bar") + ("foo bar" + "a")) + "a") + "hello") + "foo bar") + ("foo bar" + ("hello" + "a")))
    v14 = ""
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
