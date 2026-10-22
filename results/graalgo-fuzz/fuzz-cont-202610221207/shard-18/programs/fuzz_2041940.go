package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int32; F1 string; F2 int64 }
type GenS3 struct { F0 int64; F1 int16; F2 int }
type GenS4 struct { F0 uint64; F1 bool; F2 int8 }
func genFn0() uint32 {
    var v4 uint32 = 397
    for fi5 := 0; fi5 < 8; fi5++ {
        if (1 > 255) {
            v4 -= 524
            v4 -= v4
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := GenS1{F0: 956, F1: (fi6 << ((4294967295) & 63))}
                v4 += (v4 % ((v4) | 1))
                v8 := GenS2{F0: -838, F1: (("foo bar" + "foo bar") + "a"), F2: -1}
                _ = fi6
                _ = v7
                _ = v8
            }
        }
        var v9 uint32 = v4
        _ = fi5
        _ = v9
    }
    _ = v4
    return 2147483647
}

func main() {
    var v1 int16 = -1
    var v2 uint = 65535
    v3 := []uint32{}
    var v10 uint64 = 18446744073709551615
    for fi11 := 0; fi11 < 10; fi11++ {
        if (0 > -93) {
            var v12 string = ("hello" + "\n\t")
            _ = v12
        }
        fmt.Println(fi11, v2)
        v13 := GenS2{F0: 602, F1: ("foo bar" + ("a" + ("a" + "foo bar"))), F2: 9223372036854775807}
        _ = fi11
        _ = v13
    }
    v14 := GenS3{F0: 0, F1: v1, F2: 42}
    v1 -= v1
    if len(v3) > 0 {
        v3[0] = 559
    }
    switch 21 {
    case 2147483647:
        v2 += (v2 + (v2 | v2))
    default:
        v1 += -356
    }
    fmt.Println(v10, v1)
    var v15 string = ("a" + (("a" + "foo bar") + "a"))
    if false {
        v15 = v15
        var v16 uint8 = 42
        if false {
            if false {
                v1 -= 1000
                v1 = -1
            } else {
                v2 = (v2 & (v2 / ((247) | 1)))
            }
            _ = len(v3)
        } else {
            v17 := GenS1{F0: 1, F1: 127}
            _ = v17
        }
        _ = v16
    }
    var v18 uint = (v2 >> (((v2 & 255)) & 63))
    if true {
        _ = len(v3)
        v3 = append(v3, 0)
    }
    if len(v3) > 0 {
        v3[0] = 65535
    }
    if true {
        if false {
            v3 = append(v3, 2147483648)
        } else {
            var v19 uint16 = 200
            var v20 string = v15
            var v21 uint = (^(v18))
            _ = v19
            _ = v20
            _ = v21
        }
        fmt.Println(v2, v10, v1)
    }
    fmt.Println(v10, v1)
    v15 = ("foo bar" + v15)
    for fi22 := 0; fi22 < 19; fi22++ {
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 bool = false
            v25 := GenS4{F0: v10, F1: ((v24 && v24) && v24), F2: -128}
            _ = fi23
            _ = v24
            _ = v25
        }
        v26 := GenS4{F0: (v10 % ((v10) | 1)), F1: false, F2: -128}
        for fi27 := 0; fi27 < 19; fi27++ {
            var v28 uint32 = 1
            _ = fi27
            _ = v28
        }
        _ = fi22
        _ = v26
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
}
