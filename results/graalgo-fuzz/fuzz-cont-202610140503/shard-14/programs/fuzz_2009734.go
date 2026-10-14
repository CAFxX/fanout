package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 string }
type GenS2 struct { F0 bool; F1 int64; F2 int }
func genFn0(p0 bool, p1 string) int {
    v4 := GenS1{F0: 861, F1: 42, F2: p1}
    p1 = "hello"
    p1 = "hello"
    _ = p0
    _ = p1
    _ = v4
    return 32768
}

func main() {
    var v1 int64 = 606
    v2 := []uint32{0, 2147483648, 0}
    v3 := []int{32768, 0}
    var v5 int8 = 127
    switch "a" {
    case "\n\t":
        if ("" < "\n\t") {
            var v6 uint64 = 0
            fmt.Println(v5)
            v5 = 42
            _ = v6
        } else {
            _ = len(v2)
        }
    }
    var v7 string = "hello"
    v5 = (v5 ^ (v5 & (v5 % ((1) | 1))))
    var v8 float32 = 0.0
    if (true || (18446744073709551615 >= 1)) {
        switch v7 {
        case "":
            for fi9 := 0; fi9 < 18; fi9++ {
                for fi10 := 0; fi10 < 5; fi10++ {
                    v11 := GenS2{F0: (false && ((float32(1.5) + ((v8 / v8) / (v8 / v8))) <= v8)), F1: (-(v1)), F2: fi9}
                    _ = fi10
                    _ = v11
                }
                _ = fi9
            }
        case "foo bar":
            var v12 int16 = 0
            _ = v12
        case "\n\t":
            for fi13 := 0; fi13 < 7; fi13++ {
                fmt.Println(v8)
                switch (v7 + v7) {
                case "foo bar":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v14 := GenS1{F0: 0, F1: 18446744073709551615, F2: v7}
                _ = fi13
                _ = v14
            }
        default:
            if true {
                var v15 int = -9223372036854775808
                fmt.Println(v7, v1)
                var v16 int16 = 356
                _ = v15
                _ = v16
            } else {
                var v17 int32 = 0
                _ = v17
            }
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        v1 = (v1 ^ 1)
    } else {
        var v18 uint8 = 101
        _ = v18
    }
    v19 := GenS1{F0: 4294967295, F1: 0, F2: v7}
    var v20 uint16 = 1
    for fi21 := 0; fi21 < 7; fi21++ {
        if len(v3) > 0 {
            v3[0] = (fi21 * fi21)
        }
        _ = fi21
    }
    v1 = 186
    v20 += v20
    var v22 uint8 = 1
    for fi23 := 0; fi23 < 4; fi23++ {
        v2 = append(v2, 126)
        _ = fi23
    }
    for fi24 := 0; fi24 < 14; fi24++ {
        v20 = (v20 + uint16(v22))
        v5 = v5
        var v25 uint16 = (v20 >> ((42) & 63))
        _ = fi24
        _ = v25
    }
    if ((^(v1)) != v1) {
        fmt.Println(v1, v1, v8)
    }
    if false {
        switch genFn0((false || (false && (v8 <= v8))), "") {
        case 32768:
            var v26 uint16 = v20
            _ = v26
        case 9:
            if len(v2) > 0 {
                v2[0] = 4294967295
            }
        }
        _ = len(v2)
        fmt.Println(v22)
    } else {
        _ = len(v3)
    }
    var v27 uint16 = (v20 ^ v20)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v27
    _ = v3
    _ = v5
    _ = v7
    _ = v8
}
