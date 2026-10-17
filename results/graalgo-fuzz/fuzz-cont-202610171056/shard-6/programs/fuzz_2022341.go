package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint8; F2 uint32 }
type GenS4 struct { F0 uint8; F1 uint8 }
func genFn0() bool {
    var v4 int64 = 0
    var v5 int = -569
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := GenS1{F0: 255}
        var v8 uint64 = 2
        _ = fi6
        _ = v7
        _ = v8
    }
    switch "foo bar" {
    case "":
        for fi9 := 0; fi9 < 17; fi9++ {
            v4 -= 113
            if true {
                v5 -= v5
            }
            _ = fi9
        }
    }
    _ = v4
    _ = v5
    return true
}
func genFn1() int32 {
    v10 := []int{}
    if (3.14159 >= (1e10 + 1.5)) {
        var v11 uint32 = 571
        v10 = append(v10, 32768)
        switch 42 {
        case 127:
            v12 := GenS1{F0: 0}
            _ = v12
        case 9223372036854775807:
            v13 := GenS2{F0: "\n\t", F1: 251, F2: 1}
            _ = v13
        case 788:
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
        default:
            fmt.Println(v11)
        }
        _ = v11
    }
    var v14 float32 = ((3.14159 / 1e10) / 1.5)
    _ = len(v10)
    v10 = append(v10, -9223372036854775808)
    _ = v10
    _ = v14
    return -125
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint32 = 2
    var v3 uint64 = 192
    var v15 int16 = 1
    v3 = v3
    fmt.Println(v15, v2)
    v2 = (v2 + 65535)
    var v16 float32 = 3.14159
    v16 += v16
    switch (("hello" + "hello") + ("foo bar" + "\n\t")) {
    case "\n\t":
        v17 := []int{128, -705}
        _ = v17
    case "a":
        v15 = v15
    case "foo bar":
        if (100 == -1) {
            var v18 uint32 = (v2 << ((952) & 63))
            v15 -= v15
            _ = v18
        } else {
            var v19 uint64 = v3
            v20 := GenS1{F0: 42}
            var v21 int64 = 427
            _ = v19
            _ = v20
            _ = v21
        }
    default:
        if (-1 != genFn1()) {
            _ = len(v1)
            var v22 int64 = -493
            _ = v22
        }
    }
    v23 := []int{4294967295, 32768}
    v16 -= (0.0 * float32(((0.1 - (0.0 + 1e10)) + (1.5 / float64((1e10 * float32((0.0 + float64(v16)))))))))
    delete(v1, 32768)
    v2 = (v2 ^ (v2 % (((v2 >> ((65535) & 63))) | 1)))
    for fi24 := 0; fi24 < 4; fi24++ {
        v25 := GenS3{F0: (-(v15)), F1: 201, F2: 42}
        v15 += (v15 ^ -956)
        v3 = v3
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 16; fi26++ {
        v27 := GenS1{F0: 65535}
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 11; fi28++ {
        v29 := GenS1{F0: 32767}
        _ = fi28
        _ = v29
    }
    if (!false) {
        v30 := GenS1{F0: 42}
        var v31 int8 = -17
        v32 := GenS4{F0: 75, F1: 104}
        _ = v30
        _ = v31
        _ = v32
    } else {
        var v33 int32 = 1
        _ = v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
}
