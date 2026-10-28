package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string) uint32 {
    v7 := []int{0, 127, 42, 127}
    v8 := GenS1{F0: "hello", F1: 42}
    _ = p0
    _ = v7
    _ = v8
    return 2147483648
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]int)
    var v3 uint = 65535
    var v4 int64 = 0
    var v5 int8 = 1
    v6 := []uint32{970}
    switch (("a" + "a") + ("a" + "foo bar")) {
    case "\n\t":
        v3 = v3
    case "a":
        for fi9 := 0; fi9 < 3; fi9++ {
            fmt.Println(v3)
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 6; fi10++ {
        var v11 uint16 = 932
        var v12 int32 = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    v4 -= 1
    if (!false) {
        delete(v2, "foo bar")
        _ = len(v2)
    } else {
        if true {
            v5 += 42
            var v13 int8 = -128
            _ = v13
        } else {
            for fi14 := 0; fi14 < 14; fi14++ {
                _ = len(v6)
                _ = fi14
            }
        }
        if false {
            v4 = (^(v1))
            v3 = (v3 + (v3 | (v3 ^ v3)))
        }
    }
    v15 := v2["\n\t"]
    v15 -= 13
    v16 := GenS2{F0: 0}
    var v17 int64 = int64(v5)
    if true {
        v1 += -9223372036854775808
        v1 += (v17 + 512)
    } else {
        var v18 float64 = float64(3.14159)
        switch "foo bar" {
        case "a":
            var v19 uint64 = 178
            _ = v19
        case "hello":
            for fi20 := 0; fi20 < 9; fi20++ {
                var v21 uint16 = 32767
                var v22 uint8 = 255
                v23 := GenS2{F0: 43}
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        default:
            fmt.Println(v5)
        }
        _ = v18
    }
    for fi24 := 0; fi24 < 6; fi24++ {
        _ = len(v6)
        v4 += (v4 - 9223372036854775807)
        var v25 int16 = 42
        _ = fi24
        _ = v25
    }
    var v26 int8 = v5
    v3 -= (v3 ^ v3)
    _ = len(v2)
    switch "foo bar" {
    case "foo bar":
        var v27 bool = true
        _ = v27
    case "\n\t":
        if len(v6) > 0 {
            v6[0] = 42
        }
    default:
        switch ("a" + (("a" + "a") + "")) {
        case "foo bar":
            fmt.Println(v17, v5)
        default:
            v28 := GenS3{F0: (v5 % ((v5) | 1))}
            _ = v28
        }
    }
    switch "a" {
    case "a":
        if false {
            v29 := GenS4{F0: 65535}
            _ = v29
        } else {
            var v30 float64 = 0.1
            var v31 int64 = 1
            _ = v30
            _ = v31
        }
    }
    var v32 uint16 = uint16(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
