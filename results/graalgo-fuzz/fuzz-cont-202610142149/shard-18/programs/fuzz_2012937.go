package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS3 struct { F0 int }
func genFn0() int16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint8 = 1
        switch ("\n\t" + ("" + "a")) {
        case "hello":
            v5 -= v5
        default:
            v6 := GenS1{F0: 65535}
            _ = v6
        }
        v5 += v5
        _ = fi4
        _ = v5
    }
    var v7 uint32 = 2147483648
    if ((0 < 42) || (true && false)) {
        if true {
            v8 := []int{-1, 128}
            _ = v8
        } else {
            v7 = v7
            var v9 int8 = -128
            v10 := []int{9223372036854775807, 128, 9223372036854775807, 42}
            _ = v10
            _ = v9
        }
        if ((true || (v7 < 2)) && (255 > 1)) {
            v11 := GenS1{F0: 42}
            for fi12 := 0; fi12 < 8; fi12++ {
                v13 := []int{9223372036854775807, 42, fi12}
                var v14 uint32 = v7
                for fi15 := 0; fi15 < 10; fi15++ {
                    v16 := GenS2{F0: 911, F1: v7, F2: int64(v7)}
                    switch "\n\t" {
                    case "foo bar":
                        var v17 uint8 = 42
                        _ = v17
                    case "hello":
                        var v18 bool = (!true)
                        _ = v18
                    default:
                        if ((false || (!(true && true))) && ((-1 == -1) && true)) {
                            v13 = append(v13, fi15)
                            var v19 uint64 = 2
                            _ = v19
                        }
                    }
                    _ = fi15
                    _ = v16
                }
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        }
        var v20 uint16 = 65535
        _ = v20
    }
    v21 := GenS3{F0: -969}
    _ = v21
    _ = v7
    return -1
}

func main() {
    var v1 uint64 = 191
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    v2 = v2
    fmt.Println(v1, v2)
    fmt.Println(v1, v1)
    fmt.Println(v1)
    for fi22 := 0; fi22 < 4; fi22++ {
        if (true || true) {
            if ((false && ((v1 + (v1 + v1)) != 1)) && false) {
                fmt.Println(v1, v2)
                var v23 int16 = 1000
                _ = v23
            }
        } else {
            var v24 uint = 0
            var v25 uint16 = 42
            for fi26 := 0; fi26 < 13; fi26++ {
                fmt.Println(v25, v2, fi26)
                _ = fi26
            }
            _ = v24
            _ = v25
        }
        var v27 float32 = (-2.25 * 1.5)
        if (-585 <= -783) {
            _ = len(v3)
            fmt.Println(v27)
            var v28 uint64 = v1
            _ = v28
        }
        _ = fi22
        _ = v27
    }
    v1 += v1
    v29 := GenS1{F0: 443}
    switch 4294967295 {
    case -467:
        _ = len(v3)
    case 4294967295:
        v30 := GenS1{F0: 679}
        _ = v30
    default:
        fmt.Println(v2, v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
}
