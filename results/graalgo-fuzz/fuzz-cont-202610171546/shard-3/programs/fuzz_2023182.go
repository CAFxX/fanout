package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 int32, p1 int64, p2 uint16) int {
    if true {
        v5 := []int{32767}
        var v6 int64 = 994
        _ = v5
        _ = v6
    } else {
        p2 += (p2 / ((p2) | 1))
        switch "\n\t" {
        case "\n\t":
            var v7 bool = (65535 >= (p2 ^ (p2 | 1)))
            _ = v7
        default:
            p0 = 42
        }
    }
    if (0.1 <= 1.5) {
        p1 = (p1 ^ (p1 & 0))
    }
    v8 := []int{72}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return -1
}
func genFn1(p0 uint8, p1 uint64) int {
    switch 128 {
    case 2147483647:
        p0 = 0
    case 917:
        v9 := GenS1{F0: 1000}
        _ = v9
    case 42:
        var v10 string = "hello"
        _ = v10
    }
    p0 += (p0 + 173)
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 42
    var v3 int64 = -9223372036854775808
    v4 := make(map[uint32]int)
    fmt.Println(v2, v3)
    if false {
        var v11 int16 = -32768
        v3 = 20
        v3 += v3
        _ = v11
    }
    v12 := []int{195, 465, genFn1(uint8(v2), 42), genFn1(0, 18446744073709551615)}
    v13 := GenS1{F0: 539}
    var v14 int8 = 117
    if false {
        v14 = (v14 >> ((912) & 63))
        v2 = v2
    } else {
        for fi15 := 0; fi15 < 3; fi15++ {
            fmt.Println(v14)
            _ = fi15
        }
        if (true && true) {
            for fi16 := 0; fi16 < 10; fi16++ {
                v14 = v14
                var v17 string = "\n\t"
                v2 = (v2 - 255)
                _ = fi16
                _ = v17
            }
        }
    }
    switch ("\n\t" + ((("hello" + "foo bar") + "") + "foo bar")) {
    case "\n\t":
        _ = len(v12)
    default:
        _ = len(v12)
    }
    v18 := GenS1{F0: -32768}
    switch genFn1(166, 415) {
    case 488:
        var v19 uint8 = 2
        _ = v19
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
