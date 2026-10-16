package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
func genFn0(p0 int32) int8 {
    var v7 uint32 = 2
    var v8 int16 = 126
    v9 := []int{4294967295, 0, -9223372036854775808}
    switch ("" + ((("" + "") + "a") + "a")) {
    case "":
        switch "foo bar" {
        case "a":
            v8 += (v8 / ((v8) | 1))
        default:
            var v10 uint64 = 577
            _ = v10
        }
    default:
        v9 = append(v9, 128)
    }
    var v11 string = "hello"
    _ = p0
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 1
}
func genFn1(p0 uint16, p1 uint8, p2 int32) int16 {
    switch (((("foo bar" + "") + "a") + ("hello" + "\n\t")) + ("" + "hello")) {
    case "":
        p1 = (p1 - (^(p1)))
    case "hello":
        if true {
            if (65535 <= 0) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := GenS1{F0: 4294967295, F1: (0.1 != 1e10), F2: 255}
                    p0 += p0
                    _ = fi12
                    _ = v13
                }
                fmt.Println(p0, p2, p2)
                v14 := GenS1{F0: 640, F1: true, F2: (p1 - 42)}
                _ = v14
            } else {
                v15 := GenS2{F0: 42, F1: 1}
                v16 := []int{676}
                p1 = p1
                _ = v15
                _ = v16
            }
            var v17 int = 127
            switch "hello" {
            case "\n\t":
                var v18 uint64 = 18446744073709551615
                _ = v18
            default:
                var v19 uint = 4294967295
                _ = v19
            }
            _ = v17
        }
    case "\n\t":
        v20 := GenS1{F0: 388, F1: (p1 <= (p1 ^ p1)), F2: (p1 / ((p1) | 1))}
        _ = v20
    }
    var v21 uint64 = 42
    p1 = p1
    if (-81 < genFn0((p2 >> ((764) & 63)))) {
        fmt.Println(p0, p2, p1)
    } else {
        v22 := []int{127, 735}
        _ = v22
    }
    if true {
        fmt.Println(p2, p1)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return -32768
}

func main() {
    v1 := []uint32{2147483647}
    v2 := []uint32{1, 432}
    var v3 int = 127
    v4 := []int{v3}
    var v5 float32 = float32(3.14159)
    v6 := []uint32{339, 1}
    v3 = 2147483647
    for fi23 := 0; fi23 < 4; fi23++ {
        _ = len(v6)
        _ = fi23
    }
    switch (("" + "a") + ((("foo bar" + "") + "foo bar") + "foo bar")) {
    case "\n\t":
        fmt.Println(v5)
    }
    v5 = v5
    v3 -= (v3 << ((150) & 63))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    v4 = append(v4, (^(v3)))
    var v24 int = (v3 * 493)
    switch 714 {
    case -71:
        var v25 int = 2147483647
        _ = v25
    case 127:
        if (true && false) {
            fmt.Println(v3, v24, v24)
            var v26 string = "\n\t"
            _ = v26
        }
    case 2147483647:
        v5 -= v5
    }
    var v27 int16 = -209
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
