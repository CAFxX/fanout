package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8; F1 uint; F2 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int16 }
func genFn0(p0 int64, p1 int8, p2 string) int {
    for fi6 := 0; fi6 < 15; fi6++ {
        fmt.Println(p0, p0, p0)
        fmt.Println(fi6)
        v7 := GenS1{F0: 703}
        _ = fi6
        _ = v7
    }
    if (3.14159 > ((0.0 - float32(-2.25)) / 3.14159)) {
        p2 = p2
    } else {
        fmt.Println(p2, p0, p0)
        fmt.Println(p2, p0)
        if false {
            var v8 uint8 = 42
            var v9 int64 = p0
            _ = v8
            _ = v9
        }
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        switch 42 {
        case 127:
            var v11 uint8 = 127
            _ = v11
        }
        _ = fi10
    }
    if ((1e10 - 0.1) != 0.0) {
        var v12 int8 = (p1 % ((p1) | 1))
        _ = v12
    } else {
        var v13 string = p2
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return -95
}
func genFn1(p0 int32, p1 uint64) int32 {
    for fi14 := 0; fi14 < 11; fi14++ {
        fmt.Println(fi14)
        v15 := []int{fi14, (fi14 | int(p1)), 205}
        for fi16 := 0; fi16 < 17; fi16++ {
            p0 = p0
            var v17 uint64 = 926
            v18 := GenS2{F0: 214, F1: 2, F2: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
        _ = fi14
        _ = v15
    }
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    return (p0 - 2147483647)
}

func main() {
    var v1 uint8 = 2
    v2 := []string{"", "\n\t", "foo bar", "\n\t"}
    var v3 uint16 = 65535
    var v4 string = ""
    var v5 int = 741
    var v19 int64 = -9223372036854775808
    if (!(!true)) {
        switch (v4 + ((v4 + v4) + "a")) {
        case "foo bar":
            var v20 uint32 = 42
            _ = v20
        }
        var v21 int8 = 0
        _ = v21
    }
    if (!(-1 != 42)) {
        var v22 int = (v5 - 127)
        v23 := GenS2{F0: v1, F1: 2, F2: 42}
        _ = v22
        _ = v23
    }
    var v24 int = v5
    v25 := GenS3{F0: v4}
    var v26 uint16 = v3
    var v27 bool = ((1 != -492) && false)
    v28 := GenS4{F0: -32768}
    if len(v2) > 0 {
        v2[0] = v4
    }
    switch genFn0((v19 ^ (v19 % ((-1) | 1))), 1, "foo bar") {
    case 0:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case 9223372036854775807:
        for fi29 := 0; fi29 < 11; fi29++ {
            v24 = (^(fi29))
            _ = fi29
        }
    case -112:
        v2 = append(v2, "\n\t")
    default:
        v1 += (v1 << ((4294967295) & 63))
    }
    v19 = (v19 & (v19 & v19))
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
