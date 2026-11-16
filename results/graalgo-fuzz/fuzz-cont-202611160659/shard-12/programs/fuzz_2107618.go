package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 bool }
type GenS2 struct { F0 int }
func genFn0(p0 uint64, p1 uint) uint {
    p0 = (p0 >> ((208) & 63))
    v6 := GenS1{F0: (false && (false && ("" <= "a"))), F1: 127, F2: false}
    fmt.Println(p1, p1)
    var v7 int64 = -1
    if true {
        fmt.Println(p0, v7, v7)
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return p1
}

func main() {
    var v1 uint = 65535
    var v2 uint16 = 0
    var v3 uint = (v1 + 18446744073709551615)
    var v4 uint32 = 2147483647
    var v5 uint64 = 254
    v4 += (^(v4))
    v4 -= (v4 % ((65535) | 1))
    var v8 int = 2147483647
    var v9 bool = false
    v9 = (((v8 ^ v8) != v8) && ((true || (v4 >= 65535)) && v9))
    switch v8 {
    case -776:
        var v10 int32 = -349
        _ = v10
    default:
        var v11 float64 = 3.14159
        _ = v11
    }
    fmt.Println(v1, v2)
    for fi12 := 0; fi12 < 15; fi12++ {
        v13 := []int{(fi12 % ((v8) | 1)), (fi12 & 9223372036854775807)}
        v14 := GenS1{F0: (!false), F1: -1, F2: v9}
        switch fi12 {
        case -309:
            var v15 float64 = 1.5
            _ = v15
        case 9223372036854775807:
            v3 = uint(v4)
        case -417:
            v9 = true
        }
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi16 := 0; fi16 < 7; fi16++ {
        v9 = v9
        v3 += (v1 - (^(v1)))
        _ = fi16
    }
    v17 := []int{}
    if false {
        v17 = append(v17, (v8 | v8))
        fmt.Println(v2, v1)
    }
    if (-36 <= -128) {
        var v18 uint32 = v4
        var v19 string = ""
        var v20 int8 = -128
        _ = v18
        _ = v19
        _ = v20
    } else {
        v9 = true
        fmt.Println(v3)
        for fi21 := 0; fi21 < 19; fi21++ {
            v22 := GenS2{F0: v8}
            fmt.Println(v5)
            fmt.Println(v4, v4, v5)
            _ = fi21
            _ = v22
        }
    }
    v4 = v4
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 uint32 = uint32(v1)
        switch int(v2) {
        case -424:
            var v25 float64 = float64(float32(0.1))
            _ = v25
        case 578:
            var v26 uint64 = (v5 + uint64(v24))
            _ = v26
        case 4294967295:
            var v27 int16 = 702
            _ = v27
        }
        v28 := GenS2{F0: -672}
        _ = fi23
        _ = v24
        _ = v28
    }
    v17 = append(v17, v8)
    v4 -= v4
    v5 = 18446744073709551615
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
