package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 uint }

func main() {
    v1 := make(map[int]bool)
    var v2 uint = 1
    v3 := []string{"foo bar", "hello", "hello", "hello"}
    fmt.Println(v2)
    fmt.Println(v2)
    v2 += (v2 / ((65535) | 1))
    switch "a" {
    case "foo bar":
        v4 := GenS1{F0: int(v2), F1: -1, F2: ""}
        _ = v4
    case "hello":
        v5 := GenS1{F0: 4294967295, F1: 1, F2: "hello"}
        _ = v5
    case "":
        v3 = append(v3, (("\n\t" + "hello") + (("foo bar" + "a") + ((("hello" + "") + "\n\t") + "hello"))))
    }
    switch "" {
    case "":
        v6 := GenS2{F0: (v2 & (v2 ^ 42))}
        _ = v6
    case "a":
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
    }
    v2 = (v2 >> (((^(v2))) & 63))
    _ = len(v1)
    v1[42] = ((false && (!(false && true))) && (false && (!((0.0 / ((0.0 + ((0.0 * (3.14159 + 1.5)) - 3.14159)) / float64(0.1))) >= float64((0.0 + ((float32(((3.14159 / 1.0) - -2.25)) * 1e10) / (-2.25 * -2.25))))))))
    fmt.Println(v2)
    v2 -= (v2 ^ v2)
    var v7 uint = (v2 - v2)
    for fi8 := 0; fi8 < 12; fi8++ {
        fmt.Println(v2, v2)
        fmt.Println(v7, v2)
        _ = fi8
    }
    var v9 int8 = 100
    if true {
        fmt.Println(v7)
        v7 = (v2 & v2)
        if (2 <= 0) {
            v10 := GenS1{F0: 2147483647, F1: 897, F2: ("foo bar" + "a")}
            _ = v10
        }
    }
    var v11 int = -9223372036854775808
    v2 = (v7 / (((^(v2))) | 1))
    fmt.Println(v9, v2)
    if (-1 == v9) {
        v9 = int8(v11)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v9
}
