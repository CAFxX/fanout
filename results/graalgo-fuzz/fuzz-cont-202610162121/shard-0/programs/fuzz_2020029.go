package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 int32 }

func main() {
    var v1 uint = 65535
    v2 := make(map[string]string)
    var v3 int = 128
    v4 := make(map[string]string)
    switch -949 {
    case 32767:
        for fi5 := 0; fi5 < 14; fi5++ {
            v3 = (fi5 >> (((v1 / ((v1) | 1))) & 63))
            _ = fi5
        }
    }
    if false {
        switch (-(v3)) {
        case -2147483648:
            delete(v4, (("a" + "a") + "a"))
        default:
            v1 = 0
        }
        v6 := []int{}
        _ = v6
    } else {
        if (4294967295 != 580) {
            var v7 uint = v1
            v3 = v3
            _ = v7
        }
        switch "\n\t" {
        case "hello":
            v1 -= (v1 >> ((65535) & 63))
        case "foo bar":
            v1 = v1
        case "":
            var v8 uint32 = 2147483647
            _ = v8
        }
    }
    v9 := GenS1{F0: (-2.25 >= 1e10)}
    v10 := []int{v3, (^(v3)), v3}
    _ = len(v4)
    var v11 uint32 = 0
    v12 := GenS2{F0: 1}
    var v13 float32 = -2.25
    v14 := GenS3{F0: 100, F1: -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
