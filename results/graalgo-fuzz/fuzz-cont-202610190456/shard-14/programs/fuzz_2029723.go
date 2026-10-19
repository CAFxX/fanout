package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 bool }

func main() {
    v1 := make(map[string]int)
    var v2 int = 127
    v3 := make(map[int]bool)
    v4 := []string{("hello" + "\n\t"), ("\n\t" + "hello")}
    v2 -= (v2 % ((v2) | 1))
    if (1.5 == -2.25) {
        v2 -= (v2 & 128)
    } else {
        v2 -= (v2 | -29)
    }
    switch v2 {
    case 4294967295:
        switch -247 {
        case 2147483647:
            var v5 float64 = -2.25
            _ = v5
        case 42:
            v6 := GenS1{F0: -1, F1: 1, F2: false}
            _ = v6
        default:
            v2 = 1
        }
    case 128:
        var v7 uint8 = 127
        _ = v7
    case 135:
        v2 -= (v2 - (v2 + (v2 >> ((972) & 63))))
    default:
        fmt.Println(v2)
    }
    fmt.Println(v2)
    var v8 uint16 = 947
    var v9 float64 = 0.1
    switch v2 {
    case 127:
        if false {
            v2 = v2
        }
    default:
        var v10 uint16 = (v8 >> ((42) & 63))
        _ = v10
    }
    if len(v4) > 0 {
        v4[0] = "foo bar"
    }
    if (42 <= 42) {
        switch "\n\t" {
        case "foo bar":
            var v11 uint8 = 1
            _ = v11
        case "hello":
            var v12 int = 0
            _ = v12
        case "\n\t":
            if true {
                var v13 int32 = 2147483647
                for fi14 := 0; fi14 < 3; fi14++ {
                    var v15 float64 = (-2.25 / 0.1)
                    if len(v4) > 0 {
                        v4[0] = (((("a" + "hello") + "a") + ("" + "\n\t")) + "a")
                    }
                    var v16 bool = (2147483648 == 4294967295)
                    _ = fi14
                    _ = v15
                    _ = v16
                }
                _ = v13
            }
        }
    } else {
        fmt.Println(v9, v8, v9)
        var v17 uint64 = 4294967295
        _ = v17
    }
    if false {
        switch (v2 - v2) {
        case -2147483648:
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(v9, fi18)
                _ = len(v3)
                for fi19 := 0; fi19 < 8; fi19++ {
                    v2 -= int(v8)
                    v9 += (3.14159 - -2.25)
                    var v20 uint8 = 42
                    v9 -= v9
                    _ = fi19
                    _ = v20
                }
                _ = fi18
            }
        case 42:
            var v21 int32 = 1
            _ = v21
        }
    }
    var v22 float64 = (v9 * v9)
    fmt.Println(v22, v2, v22)
    v23 := GenS1{F0: 1, F1: 65535, F2: ((^(v8)) > (v8 + (v8 / ((v8) | 1))))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
