package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int8; F1 int64; F2 uint64 }
type GenS3 struct { F0 int; F1 bool }
type GenS4 struct { F0 uint64; F1 int32; F2 uint32 }
func genFn0(p0 string, p1 int64) bool {
    if (0 <= 1) {
        switch 2147483647 {
        case 32767:
            v6 := GenS1{F0: 1, F1: p0}
            _ = v6
        case 4294967295:
            var v7 int32 = -1
            _ = v7
        default:
            p0 = "a"
        }
    }
    v8 := GenS2{F0: 72, F1: (p1 * (p1 % ((p1) | 1))), F2: 2}
    _ = p0
    _ = p1
    _ = v8
    return (!(!((true || false) || (!true))))
}

func main() {
    v1 := make(map[int]string)
    var v2 int = 1
    var v3 uint8 = 42
    var v4 float64 = ((3.14159 - -2.25) + float64(3.14159))
    var v5 uint32 = 0
    if true {
        for fi9 := 0; fi9 < 3; fi9++ {
            var v10 float32 = float32(v4)
            _ = fi9
            _ = v10
        }
    }
    fmt.Println(v5)
    v3 -= 42
    for fi11 := 0; fi11 < 7; fi11++ {
        if (!(293 < 365)) {
            v12 := GenS2{F0: 0, F1: 0, F2: 4294967295}
            _ = v12
        } else {
            v13 := GenS1{F0: 1, F1: (("" + "a") + "hello")}
            v5 = v5
            v3 += v3
            _ = v13
        }
        fmt.Println(v3, v3)
        var v14 int16 = 754
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 16; fi15++ {
        if false {
            v16 := GenS3{F0: (v2 % ((v2) | 1)), F1: (uint32(v3) != (v5 << ((65535) & 63)))}
            _ = v16
        } else {
            v17 := GenS4{F0: 830, F1: 978, F2: v5}
            _ = v17
        }
        var v18 uint64 = 18446744073709551615
        _ = fi15
        _ = v18
    }
    v19 := []int{(^(v2)), v2}
    var v20 bool = (888 >= -2147483648)
    for fi21 := 0; fi21 < 20; fi21++ {
        v22 := GenS2{F0: 0, F1: -1, F2: 18446744073709551615}
        v23 := GenS2{F0: -128, F1: 0, F2: 0}
        v24 := GenS3{F0: -814, F1: (!true)}
        _ = fi21
        _ = v22
        _ = v23
        _ = v24
    }
    switch v2 {
    case 2147483647:
        fmt.Println(v5)
    case -2147483648:
        switch "\n\t" {
        case "foo bar":
            fmt.Println(v5)
        }
    case 1:
        var v25 int64 = 1
        _ = v25
    default:
        var v26 uint8 = 128
        _ = v26
    }
    v1[1] = ""
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
