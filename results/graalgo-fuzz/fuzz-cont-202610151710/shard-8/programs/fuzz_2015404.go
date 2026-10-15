package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }

func main() {
    v1 := []string{("" + "\n\t"), "\n\t", (("\n\t" + ("foo bar" + "\n\t")) + ("hello" + "a")), ("foo bar" + ("hello" + "foo bar"))}
    v2 := make(map[int]uint32)
    var v3 float64 = 0.0
    v4 := []int{-901, 9223372036854775807, 336}
    v5 := []string{"", "hello", (("" + "foo bar") + "")}
    var v6 string = "hello"
    var v7 int = 2147483647
    for fi8 := 0; fi8 < 15; fi8++ {
        switch (v7 * -688) {
        case 0:
            fmt.Println(v3, v6)
        case -850:
            fmt.Println(v6, v7, v6)
        case -264:
            fmt.Println(fi8, v6, v6)
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi8
    }
    fmt.Println(v7)
    var v9 int8 = 83
    switch (v7 % ((4294967295) | 1)) {
    case 111:
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 int8 = v9
            _ = fi10
            _ = v11
        }
    default:
        var v12 uint = 4294967295
        _ = v12
    }
    var v13 int16 = int16(v7)
    fmt.Println(v9, v6)
    if (2 == 852) {
        v5 = append(v5, v6)
        v13 = (v13 % ((v13) | 1))
    }
    v14 := GenS1{F0: 112, F1: (v7 ^ (v7 >> ((0) & 63)))}
    var v15 int32 = 968
    if (127 > (v7 & v7)) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v16 int16 = v13
    for fi17 := 0; fi17 < 14; fi17++ {
        v9 += v9
        var v18 float64 = v3
        _ = fi17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
