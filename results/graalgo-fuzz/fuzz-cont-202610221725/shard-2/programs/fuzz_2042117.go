package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint8 }

func main() {
    var v1 string = (("hello" + ("a" + "hello")) + "a")
    v2 := []string{}
    var v3 float64 = -2.25
    v4 := []string{(((v1 + "foo bar") + v1) + "foo bar"), v1}
    var v5 uint = 18446744073709551615
    var v6 uint8 = 0
    v7 := GenS1{F0: (v6 & v6), F1: int(v6), F2: v6}
    for fi8 := 0; fi8 < 10; fi8++ {
        v3 -= v3
        var v9 float32 = float32(float64(1.5))
        _ = fi8
        _ = v9
    }
    var v10 uint8 = (v6 >> (((v5 << (((v5 + (v5 | v5))) & 63))) & 63))
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v11 string = v1
    switch -113 {
    case 32767:
        for fi12 := 0; fi12 < 7; fi12++ {
            fmt.Println(v10, v3, v10)
            _ = fi12
        }
    case -9223372036854775808:
        if (((float64(3.14159) - v3) + ((3.14159 - v3) / 1.0)) > float64(1e10)) {
            fmt.Println(v1, v3)
        }
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if (455 != -1) {
            var v14 uint64 = 4294967295
            _ = v14
        } else {
            fmt.Println(v6, v10, fi13)
            if true {
                fmt.Println(v10, v1)
                _ = len(v4)
            }
        }
        v6 += v10
        _ = len(v4)
        _ = fi13
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
