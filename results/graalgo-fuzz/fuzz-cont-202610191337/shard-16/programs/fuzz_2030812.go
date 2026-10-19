package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 uint64; F1 uint8 }
type GenS5 struct { F0 uint16; F1 string }
func genFn0(p0 int16) int16 {
    switch "\n\t" {
    case "foo bar":
        p0 = p0
    case "hello":
        v7 := GenS1{F0: -9223372036854775808, F1: ((p0 + 227) <= (p0 + (p0 | 42))), F2: 234}
        _ = v7
    }
    v8 := GenS2{F0: 255, F1: 594, F2: 2}
    v9 := []int{}
    _ = p0
    _ = v8
    _ = v9
    return p0
}
func genFn1(p0 bool, p1 uint64) int8 {
    var v10 bool = (!p0)
    var v11 uint16 = 1
    var v12 int16 = 1000
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 100
}

func main() {
    v1 := make(map[uint32]int)
    var v2 bool = (!((!false) || true))
    v3 := []uint32{}
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{0, 2}
    var v6 float32 = (1e10 + 0.0)
    v2 = (0 <= 104)
    var v13 int64 = 812
    switch "" {
    case "foo bar":
        var v14 int64 = v13
        _ = v14
    default:
        fmt.Println(v6, v6)
    }
    var v15 int16 = -32768
    fmt.Println(v2)
    delete(v1, 2147483647)
    var v16 uint = 255
    var v17 uint16 = 80
    var v18 uint64 = uint64(v17)
    if (int32(v16) <= 42) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v19 bool = (!false)
        _ = v19
    } else {
        for fi20 := 0; fi20 < 17; fi20++ {
            _ = len(v5)
            var v21 int = fi20
            for fi22 := 0; fi22 < 5; fi22++ {
                _ = len(v1)
                v21 -= int(v17)
                _ = fi22
            }
            _ = fi20
            _ = v21
        }
        if false {
            var v23 int32 = 2147483647
            _ = len(v5)
            v2 = (!v2)
            _ = v23
        }
        fmt.Println(v16, v4)
    }
    v24 := GenS2{F0: v17, F1: 2, F2: (v4 & 0)}
    switch "" {
    case "hello":
        switch (("hello" + "a") + "\n\t") {
        case "a":
            if true {
                v25 := GenS3{F0: (^(v16)), F1: "a"}
                _ = v25
            } else {
                var v26 bool = v2
                delete(v1, 297)
                v16 = v16
                _ = v26
            }
        case "foo bar":
            v27 := GenS4{F0: (^(v4)), F1: 42}
            _ = v27
        case "\n\t":
            _ = len(v3)
        }
    case "foo bar":
        v28 := GenS5{F0: v17, F1: ("foo bar" + "")}
        _ = v28
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    var v29 uint16 = v17
    if v2 {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        switch 4294967295 {
        case -566:
            v5 = append(v5, 998)
        case 32767:
            if ((int32(v15) >= -2147483648) || false) {
                fmt.Println(v15)
                fmt.Println(v13)
                var v30 uint64 = 42
                _ = v30
            }
        default:
            var v31 string = ((("hello" + ("" + "a")) + ("" + "")) + ((("" + "") + ("" + "a")) + ""))
            _ = v31
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v24
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
