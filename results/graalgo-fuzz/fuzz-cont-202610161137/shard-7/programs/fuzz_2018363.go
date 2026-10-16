package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int32; F1 int16; F2 bool }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32, p1 uint, p2 uint8) bool {
    fmt.Println(p0, p0)
    if false {
        switch "hello" {
        case "hello":
            v6 := GenS1{F0: 4294967295, F1: 43}
            _ = v6
        case "\n\t":
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := []int{505, (fi7 ^ fi7), fi7, (fi7 >> ((p1) & 63))}
                v9 := GenS2{F0: -2147483648, F1: -593, F2: (18446744073709551615 > 2)}
                _ = fi7
                _ = v8
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 7; fi10++ {
                v11 := []int{(fi10 % (((fi10 % ((int(p1)) | 1))) | 1)), 127, (fi10 & fi10), (fi10 | 1)}
                var v12 uint16 = 65535
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        v13 := []int{9223372036854775807, 42, 2147483647, -1}
        _ = v13
    } else {
        v14 := []int{419, 42}
        for fi15 := 0; fi15 < 9; fi15++ {
            if len(v14) > 0 {
                v14[0] = (fi15 / (((-(fi15))) | 1))
            }
            _ = fi15
        }
        _ = v14
    }
    p2 -= 42
    switch 0 {
    case -205:
        var v16 int64 = 313
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    return (2 >= 698)
}

func main() {
    v1 := []string{"foo bar", ""}
    var v2 uint32 = 607
    var v3 uint = 4294967295
    v4 := make(map[uint32]bool)
    var v5 uint16 = 0
    _ = len(v4)
    v5 -= (v5 - v5)
    if len(v1) > 0 {
        v1[0] = ("foo bar" + "a")
    }
    v17 := GenS3{F0: -32768}
    if false {
        if (635 != 1000) {
            var v18 int8 = -1
            v3 = v3
            var v19 uint8 = 42
            _ = v18
            _ = v19
        }
    }
    v3 -= v3
    fmt.Println(v2, v2, v5)
    fmt.Println(v5)
    v2 += 42
    v20 := GenS3{F0: -561}
    var v21 bool = true
    v2 = uint32(v5)
    var v22 uint8 = 127
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
