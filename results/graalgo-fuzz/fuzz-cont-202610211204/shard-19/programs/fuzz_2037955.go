package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 bool }
func genFn0() bool {
    var v6 uint8 = 2
    v6 += (v6 >> ((108) & 63))
    if false {
        v7 := []int{}
        v8 := GenS1{F0: 2, F1: v6}
        v6 = (v6 % ((v6) | 1))
        _ = v7
        _ = v8
    }
    v6 += v6
    _ = v6
    return false
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int16 = -32768
    v3 := make(map[string]int)
    var v4 int = 0
    v5 := []string{}
    if ((v4 != (v4 << ((65535) & 63))) || true) {
        switch "a" {
        case "foo bar":
            v2 -= (-(v2))
        case "hello":
            switch (v4 % (((v4 << ((2) & 63))) | 1)) {
            case 4294967295:
                v1 += (v1 >> ((65535) & 63))
            default:
                v1 -= (v1 ^ 1)
            }
        default:
            switch "foo bar" {
            case "hello":
                v4 = (v4 ^ v4)
            }
        }
        var v9 int = v4
        if true {
            switch (v4 * 0) {
            case -9223372036854775808:
                v10 := GenS2{F0: (!true)}
                _ = v10
            default:
                _ = len(v5)
            }
        }
        _ = v9
    }
    fmt.Println(v2, v4, v4)
    for fi11 := 0; fi11 < 15; fi11++ {
        for fi12 := 0; fi12 < 7; fi12++ {
            var v13 int32 = 42
            var v14 uint16 = 532
            _ = fi12
            _ = v13
            _ = v14
        }
        if (true && false) {
            var v15 bool = false
            _ = v15
        } else {
            _ = len(v5)
            var v16 int64 = int64(fi11)
            _ = v16
        }
        _ = fi11
    }
    fmt.Println(v4, v4, v2)
    for fi17 := 0; fi17 < 7; fi17++ {
        fmt.Println(fi17)
        switch fi17 {
        case 4294967295:
            _ = len(v5)
        case 9223372036854775807:
            v18 := GenS2{F0: true}
            _ = v18
        }
        _ = fi17
    }
    fmt.Println(v1, v1, v1)
    v1 = (v1 << ((42) & 63))
    var v19 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
