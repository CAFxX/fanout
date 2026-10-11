package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int8 }
type GenS4 struct { F0 int; F1 string; F2 uint8 }
func genFn0(p0 int32) uint8 {
    switch -908 {
    case 127:
        for fi5 := 0; fi5 < 13; fi5++ {
            p0 = 2147483647
            _ = fi5
        }
    }
    for fi6 := 0; fi6 < 17; fi6++ {
        v7 := []int{9223372036854775807}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    p0 -= (p0 - p0)
    _ = p0
    return 42
}
func genFn1(p0 int, p1 string) uint16 {
    v8 := GenS1{F0: 1, F1: 0, F2: 0}
    for fi9 := 0; fi9 < 5; fi9++ {
        p0 = fi9
        p0 = p0
        var v10 uint = 2
        _ = fi9
        _ = v10
    }
    var v11 float64 = 1e10
    for fi12 := 0; fi12 < 11; fi12++ {
        var v13 uint8 = 42
        p1 = "a"
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return 32767
}

func main() {
    var v1 int64 = -411
    var v2 uint = 0
    v3 := []uint32{2147483647, 0}
    v4 := make(map[string]string)
    if true {
        v2 = (v2 << (((v2 + 2)) & 63))
    }
    if ((v1 * v1) == -9223372036854775808) {
        fmt.Println(v2, v2)
        var v14 uint8 = 177
        _ = v14
    }
    fmt.Println(v2)
    for fi15 := 0; fi15 < 7; fi15++ {
        if (0 <= 398) {
            delete(v4, ("" + ("foo bar" + ("\n\t" + ((((("\n\t" + "a") + "foo bar") + (("hello" + "foo bar") + "hello")) + "") + "")))))
            var v16 bool = true
            var v17 uint16 = 32767
            _ = v16
            _ = v17
        }
        if (v1 > -296) {
            v18 := GenS2{F0: 1000, F1: 140}
            var v19 int = fi15
            var v20 int16 = 32767
            _ = v18
            _ = v19
            _ = v20
        }
        var v21 int16 = 32767
        _ = fi15
        _ = v21
    }
    if len(v3) > 0 {
        v3[0] = 42
    }
    switch -1 {
    case 128:
        v22 := GenS1{F0: 0, F1: 0, F2: 644}
        _ = v22
    case 0:
        switch -9223372036854775808 {
        case -580:
            fmt.Println(v1)
        case 32768:
            v23 := GenS3{F0: 2, F1: 100}
            _ = v23
        case -946:
            for fi24 := 0; fi24 < 10; fi24++ {
                var v25 int8 = 1
                if false {
                    v26 := GenS3{F0: 0, F1: 100}
                    _ = v26
                }
                _ = fi24
                _ = v25
            }
        default:
            fmt.Println(v2, v2)
        }
    }
    if (206 >= 9223372036854775807) {
        fmt.Println(v1, v2)
    } else {
        v27 := GenS4{F0: 32767, F1: "hello", F2: 127}
        _ = v27
    }
    for fi28 := 0; fi28 < 13; fi28++ {
        var v29 uint8 = 128
        _ = fi28
        _ = v29
    }
    if (!true) {
        var v30 uint8 = 47
        _ = v30
    } else {
        fmt.Println(v1, v1)
        v2 -= (v2 & 65535)
        var v31 uint = 18446744073709551615
        _ = v31
    }
    var v32 int = int(v1)
    if len(v3) > 0 {
        v3[0] = 67
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
}
