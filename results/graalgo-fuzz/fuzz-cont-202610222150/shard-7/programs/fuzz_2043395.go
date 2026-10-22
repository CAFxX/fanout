package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 string }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 int8 }
func genFn0(p0 uint64, p1 uint8) int16 {
    var v7 uint8 = (p1 * p1)
    if false {
        v8 := []int{int(p0)}
        var v9 float32 = 3.14159
        _ = v8
        _ = v9
    }
    if (-294 > 42) {
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int64 = 0
            _ = fi10
            _ = v11
        }
        v12 := []int{583, -378, int(v7)}
        _ = v12
    } else {
        p0 += p0
    }
    var v13 float32 = 0.1
    _ = p0
    _ = p1
    _ = v13
    _ = v7
    return -32768
}

func main() {
    var v1 int64 = 367
    var v2 uint32 = 8
    v3 := []int{32767}
    v4 := make(map[uint32]uint32)
    v5 := []int{}
    var v6 uint32 = v2
    switch "foo bar" {
    case "a":
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint32 = v2
            _ = len(v3)
            _ = fi14
            _ = v15
        }
    case "\n\t":
        v6 = v2
    default:
        v16 := GenS1{F0: 2147483647, F1: 1, F2: "a"}
        _ = v16
    }
    fmt.Println(v1, v2)
    switch 142 {
    case 0:
        fmt.Println(v6)
    case -858:
        v17 := GenS2{F0: 948}
        _ = v17
    case 4294967295:
        v6 = 42
    default:
        if false {
            for fi18 := 0; fi18 < 14; fi18++ {
                delete(v4, 666)
                _ = fi18
            }
        }
    }
    var v19 int16 = -32768
    var v20 int = 0
    v2 -= v6
    for fi21 := 0; fi21 < 1; fi21++ {
        v22 := GenS3{F0: uint(v6), F1: 311, F2: -109}
        var v23 int8 = 1
        v20 -= 198
        _ = fi21
        _ = v22
        _ = v23
    }
    if false {
        var v24 float64 = 1.5
        var v25 uint8 = 2
        _ = v24
        _ = v25
    }
    var v26 uint8 = 1
    switch "foo bar" {
    case "":
        fmt.Println(v2, v26)
    case "foo bar":
        v2 = v2
    }
    if (((true || (false && (!false))) && (true || (true || ((!true) && ((uint64(v26) <= 685) && false))))) && (!(true && ((true && true) || ((-(v20)) == -9223372036854775808))))) {
        switch 32767 {
        case -1:
            if ((v19 + (v19 * v19)) > genFn0(42, (v26 / ((v26) | 1)))) {
                var v27 bool = false
                var v28 int = (v20 << ((65535) & 63))
                _ = v27
                _ = v28
            } else {
                var v29 float64 = (-2.25 - -2.25)
                _ = v29
            }
        }
        v6 = v2
        fmt.Println(v20)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
