package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint8, p1 uint8) int8 {
    p0 = 127
    v4 := []int{440, 197}
    _ = p0
    _ = p1
    _ = v4
    return -128
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 2
    v3 := make(map[int]string)
    v2 += (v2 * (v2 ^ (^(v2))))
    for fi5 := 0; fi5 < 8; fi5++ {
        switch (fi5 | (fi5 - (fi5 + 9223372036854775807))) {
        case 630:
            fmt.Println(fi5)
        case -179:
            if true {
                var v6 uint64 = v2
                v2 -= 0
                fmt.Println(fi5, v2, v6)
                _ = v6
            } else {
                var v7 uint = 897
                _ = v7
            }
        default:
            var v8 bool = false
            _ = v8
        }
        _ = fi5
    }
    if false {
        v9 := GenS1{F0: "foo bar"}
        if (2 > 2) {
            fmt.Println(v2)
        }
        _ = v9
    }
    _ = len(v3)
    v2 -= v2
    switch "hello" {
    case "foo bar":
        var v10 int = 824
        _ = v10
    default:
        delete(v1, 197)
    }
    if true {
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 bool = false
            v12 = (-128 < -128)
            _ = fi11
            _ = v12
        }
    } else {
        for fi13 := 0; fi13 < 16; fi13++ {
            var v14 int16 = 89
            v3[fi13] = "foo bar"
            _ = fi13
            _ = v14
        }
        fmt.Println(v2)
        v2 = 4294967295
    }
    v15 := GenS2{F0: false, F1: 32767}
    var v16 int = -989
    fmt.Println(v2)
    var v17 uint = 42
    var v18 int16 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
