package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint64; F1 int; F2 int16 }
type GenS3 struct { F0 uint; F1 int8 }
func genFn0(p0 uint64, p1 int8, p2 uint) int16 {
    var v7 int16 = 0
    p1 -= p1
    if false {
        if (false && true) {
            var v8 uint16 = 0
            v9 := []int{-226}
            v10 := GenS1{F0: ("" + ""), F1: 2, F2: 597}
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 uint16 = 535
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 string = "\n\t"
            var v14 int16 = v7
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (v7 | 1)
}
func genFn1(p0 int16, p1 uint8, p2 uint64) string {
    switch -1 {
    case 4294967295:
        var v15 float32 = 0.1
        _ = v15
    case -393:
        v16 := GenS2{F0: (p2 & (^(p2))), F1: 32768, F2: p0}
        _ = v16
    case 32767:
        v17 := []int{}
        _ = v17
    default:
        p0 = genFn0(p2, -128, 4294967295)
    }
    v18 := []int{-1}
    var v19 int64 = -1
    switch "\n\t" {
    case "foo bar":
        v20 := GenS1{F0: "\n\t", F1: 2147483648, F2: 42}
        _ = v20
    case "":
        p1 -= p1
    default:
        v18 = append(v18, -9223372036854775808)
    }
    p1 += (p1 - (p1 >> ((65535) & 63)))
    _ = p0
    _ = p1
    _ = p2
    _ = v18
    _ = v19
    return ("" + ("foo bar" + "hello"))
}

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{1, 2147483648, 2147483647, 60}
    var v3 uint16 = 1
    var v4 uint64 = 18446744073709551615
    v5 := []uint32{uint32(v4)}
    var v6 int = 32767
    v21 := GenS2{F0: v4, F1: 32767, F2: -32768}
    fmt.Println(v3, v6)
    switch "" {
    case "hello":
        var v22 int = v6
        _ = v22
    case "a":
        v6 = v6
    case "\n\t":
        v6 += (v6 ^ (v6 / ((v6) | 1)))
    }
    if true {
        var v23 bool = false
        fmt.Println(v23, v1, v6)
        _ = v23
    } else {
        var v24 uint16 = v3
        _ = v24
    }
    switch "a" {
    case "\n\t":
        var v25 float32 = (v1 * float32((0.0 / 1e10)))
        _ = v25
    case "hello":
        for fi26 := 0; fi26 < 4; fi26++ {
            v4 = v4
            _ = fi26
        }
    default:
        var v27 float64 = 1.5
        _ = v27
    }
    for fi28 := 0; fi28 < 8; fi28++ {
        v4 = 18446744073709551615
        if false {
            v5 = append(v5, 2147483647)
            v29 := GenS2{F0: v4, F1: fi28, F2: 0}
            _ = v29
        } else {
            for fi30 := 0; fi30 < 4; fi30++ {
                if ((v4 - (v4 * v4)) == 42) {
                    v31 := GenS2{F0: v4, F1: v6, F2: genFn0((v4 >> ((4294967295) & 63)), -128, 255)}
                    var v32 float64 = 0.1
                    _ = v31
                    _ = v32
                }
                _ = fi30
            }
            var v33 int32 = -2147483648
            _ = v33
        }
        var v34 int64 = -1
        _ = fi28
        _ = v34
    }
    if ("hello" <= ("foo bar" + "foo bar")) {
        v3 = (v3 >> ((2) & 63))
        v35 := GenS1{F0: "a", F1: 1, F2: -1}
        v3 -= 682
        _ = v35
    }
    v36 := GenS3{F0: 13, F1: -1}
    if true {
        if (true && true) {
            v37 := GenS1{F0: "\n\t", F1: uint32(v4), F2: -863}
            _ = v37
        } else {
            fmt.Println(v6, v4, v3)
            var v38 uint8 = 1
            v3 -= v3
            _ = v38
        }
        switch "foo bar" {
        case "hello":
            _ = len(v2)
        }
    } else {
        v39 := GenS1{F0: "foo bar", F1: 0, F2: 0}
        _ = v39
    }
    for fi40 := 0; fi40 < 13; fi40++ {
        var v41 float32 = float32(0.0)
        fmt.Println(fi40, fi40)
        _ = fi40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
