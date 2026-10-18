package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
func genFn0(p0 string, p1 uint16) int64 {
    v5 := []int{-2147483648, -1}
    p0 = ((p0 + (p0 + p0)) + "")
    p0 = p0
    switch -538 {
    case 0:
        var v6 uint32 = 2
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}

func main() {
    v1 := []string{"", ((("\n\t" + "a") + "hello") + ("a" + "\n\t"))}
    v2 := make(map[int]bool)
    v3 := []int{0, -664, 127, 128}
    var v4 uint = 65535
    v4 -= v4
    v4 += (v4 | (v4 & v4))
    v4 = (v4 | v4)
    var v7 int = -783
    fmt.Println(v4, v7)
    v7 = (v7 >> (((v4 & (v4 - v4))) & 63))
    var v8 uint16 = 255
    v9 := GenS1{F0: 4294967295, F1: 1}
    switch v7 {
    case -403:
        switch "\n\t" {
        case "":
            v10 := GenS1{F0: 2147483648, F1: 2}
            _ = v10
        case "a":
            switch ("a" + ("foo bar" + "foo bar")) {
            case "":
                fmt.Println(v4, v8, v8)
            }
        default:
            v8 += v8
        }
    case 389:
        v4 += (v4 + (^(v4)))
    case -9223372036854775808:
        switch "hello" {
        case "a":
            if ((float64(1.5) + 3.14159) <= 0.1) {
                if len(v1) > 0 {
                    v1[0] = ""
                }
                _ = len(v1)
                fmt.Println(v8, v8, v7)
            }
        default:
            fmt.Println(v7, v4)
        }
    }
    v7 -= v7
    v11 := GenS1{F0: 65535, F1: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
