package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16; F2 uint }
type GenS3 struct { F0 int; F1 string; F2 bool }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int64 }
func genFn0(p0 uint, p1 bool, p2 int64) uint8 {
    v6 := GenS1{F0: 18446744073709551615}
    for fi7 := 0; fi7 < 4; fi7++ {
        p1 = (42 < 2)
        _ = fi7
    }
    var v8 float64 = -2.25
    var v9 bool = (v8 <= 3.14159)
    var v10 float32 = (((1.5 * ((1e10 * -2.25) + float32(v8))) - 0.1) - 1.5)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    _ = v8
    _ = v9
    return 255
}

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := []string{"a", "a"}
    v4 := []uint32{2147483648}
    v5 := make(map[uint32]bool)
    var v11 float64 = ((0.1 - ((1.5 * 0.0) / -2.25)) + 0.1)
    var v12 int64 = 265
    v11 -= ((3.14159 - -2.25) * v11)
    for fi13 := 0; fi13 < 15; fi13++ {
        v14 := GenS1{F0: 1}
        v15 := GenS1{F0: 42}
        var v16 int16 = -32768
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    var v17 uint64 = 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v12 += 9223372036854775807
    if true {
        v12 = v12
        switch 128 {
        case -2147483648:
            fmt.Println(v17, v12, v17)
        case 268:
            var v18 int8 = -128
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 float32 = 0.0
            v21 := GenS2{F0: v17, F1: 32767, F2: 2}
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    if true {
        if (726 <= 668) {
            v22 := GenS3{F0: 9223372036854775807, F1: ("hello" + ("hello" + ("hello" + "foo bar"))), F2: true}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        v4 = append(v4, 2147483648)
        _ = fi23
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        delete(v2, (fi24 + int(v17)))
        fmt.Println(v12, v17)
        delete(v2, (fi24 ^ (fi24 / (((fi24 / ((fi24) | 1))) | 1))))
        _ = fi24
    }
    if true {
        for fi25 := 0; fi25 < 14; fi25++ {
            switch "\n\t" {
            case "":
                v12 = v12
            case "a":
                v26 := GenS4{F0: 65535, F1: 42, F2: -128}
                _ = v26
            case "hello":
                v27 := GenS4{F0: 65535, F1: 179, F2: 9}
                _ = v27
            default:
                v11 = 1.5
            }
            _ = fi25
        }
        if (!(!(65535 > 909))) {
            v12 += (^(v12))
        }
    } else {
        if ((false && false) && (1 >= 0)) {
            var v28 uint64 = (v17 ^ 18446744073709551615)
            _ = v28
        }
    }
    if true {
        switch "a" {
        case "foo bar":
            v29 := v1["a"]
            _ = v29
        case "":
            v11 -= (3.14159 * 0.0)
        case "a":
            v30 := GenS1{F0: 2}
            _ = v30
        }
        v31 := GenS5{F0: (v17 - 1), F1: 42, F2: 51}
        fmt.Println(v12, v12, v17)
        _ = v31
    } else {
        var v32 uint = 0
        if false {
            var v33 int16 = -1
            _ = len(v5)
            v34 := GenS5{F0: (v17 >> ((778) & 63)), F1: 204, F2: v12}
            _ = v33
            _ = v34
        } else {
            var v35 bool = ((v32 / ((v32) | 1)) != v32)
            var v36 bool = false
            var v37 uint16 = 0
            _ = v35
            _ = v36
            _ = v37
        }
        v12 -= -9223372036854775808
        _ = v32
    }
    for fi38 := 0; fi38 < 14; fi38++ {
        for fi39 := 0; fi39 < 6; fi39++ {
            v11 += float64(3.14159)
            fmt.Println(v11)
            var v40 float64 = v11
            fmt.Println(fi39, v12)
            _ = fi39
            _ = v40
        }
        v12 = v12
        delete(v5, 434)
        _ = fi38
    }
    _ = len(v3)
    fmt.Println(v11)
    switch "hello" {
    case "a":
        for fi41 := 0; fi41 < 20; fi41++ {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if (true || false) {
                var v42 uint = 380
                v17 = (^(v17))
                _ = v42
            }
            _ = fi41
        }
    case "":
        var v43 uint = 18446744073709551615
        _ = v43
    default:
        fmt.Println(v12, v17)
    }
    v17 -= (^(v17))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
