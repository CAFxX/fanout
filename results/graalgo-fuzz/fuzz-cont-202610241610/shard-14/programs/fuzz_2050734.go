package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint8 }
type GenS2 struct { F0 uint64; F1 int16; F2 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0() int8 {
    v7 := []int{152, -570}
    if true {
        var v8 string = ("hello" + ("" + "\n\t"))
        _ = v8
    }
    _ = v7
    return 97
}
func genFn1(p0 string, p1 uint32, p2 int8) uint {
    var v9 int = -250
    for fi10 := 0; fi10 < 8; fi10++ {
        p0 = p0
        if (p1 >= p1) {
            for fi11 := 0; fi11 < 11; fi11++ {
                var v12 int64 = 9223372036854775807
                v13 := GenS1{F0: p1, F1: "hello", F2: 1}
                var v14 int64 = v12
                p1 = p1
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            var v15 int16 = 0
            _ = v15
        } else {
            p1 += 545
            for fi16 := 0; fi16 < 10; fi16++ {
                p0 = p0
                p2 -= (p2 >> ((0) & 63))
                _ = fi16
            }
            v17 := GenS1{F0: (^(p1)), F1: p0, F2: 127}
            _ = v17
        }
        p1 = 65535
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 189
}

func main() {
    var v1 float32 = 0.1
    v2 := []uint32{2147483648, 398, 517}
    v3 := make(map[string]string)
    v4 := make(map[uint32]string)
    v5 := []int{-342, 32767, 2147483647, 42}
    var v6 int16 = 42
    _ = len(v3)
    fmt.Println(v6, v6)
    for fi18 := 0; fi18 < 18; fi18++ {
        if false {
            var v19 int16 = v6
            _ = v19
        }
        _ = len(v2)
        _ = fi18
    }
    if true {
        fmt.Println(v1, v6)
        v1 = v1
        v6 = -546
    } else {
        var v20 uint32 = 42
        v21 := v3[("foo bar" + ("\n\t" + "a"))]
        _ = v20
        _ = v21
    }
    var v22 int32 = -1
    var v23 int = -91
    v24 := GenS2{F0: 454, F1: -1, F2: -9223372036854775808}
    switch ("\n\t" + "\n\t") {
    case "":
        v25 := GenS2{F0: 1, F1: v6, F2: 0}
        _ = v25
    case "foo bar":
        v1 = (3.14159 + 0.1)
    }
    switch "" {
    case "foo bar":
        fmt.Println(v6)
    default:
        var v26 float32 = (v1 / ((v1 - v1) / (v1 * 0.1)))
        _ = v26
    }
    switch "foo bar" {
    case "a":
        v27 := GenS1{F0: 532, F1: "\n\t", F2: 127}
        _ = v27
    case "hello":
        v28 := GenS3{F0: -1, F1: 0}
        _ = v28
    case "foo bar":
        var v29 int = (v23 / (((v23 % ((1) | 1))) | 1))
        _ = v29
    }
    var v30 int64 = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
