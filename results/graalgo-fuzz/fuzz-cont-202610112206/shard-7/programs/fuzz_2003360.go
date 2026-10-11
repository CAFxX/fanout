package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 uint32) uint {
    var v4 int64 = -9223372036854775808
    var v5 uint = 42
    switch 32768 {
    case 32767:
        fmt.Println(v5, v5, p0)
    case 2147483647:
        v6 := GenS1{F0: (v4 == v4)}
        _ = v6
    case -349:
        fmt.Println(p0)
    default:
        v7 := GenS2{F0: -350}
        _ = v7
    }
    p0 -= 2147483647
    _ = p0
    _ = v4
    _ = v5
    return (^(v5))
}
func genFn1(p0 int32, p1 int16, p2 uint64) int64 {
    v8 := GenS2{F0: (p1 / (((p1 % ((p1) | 1))) | 1))}
    fmt.Println(p1)
    v9 := []int{-643, 9223372036854775807}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -59
}

func main() {
    v1 := []uint32{4294967295, 2147483648}
    var v2 float32 = 0.0
    var v3 uint16 = 309
    v10 := GenS3{F0: 18446744073709551615}
    fmt.Println(v2, v2)
    fmt.Println(v2)
    switch ("a" + "foo bar") {
    case "a":
        v2 -= 1e10
    default:
        if false {
            v3 = v3
            v3 = v3
            for fi11 := 0; fi11 < 20; fi11++ {
                v3 = (v3 >> ((125) & 63))
                switch 77 {
                case -1:
                    var v12 string = "foo bar"
                    _ = v12
                default:
                    fmt.Println(fi11, fi11)
                }
                if true {
                    var v13 int64 = 42
                    _ = v13
                } else {
                    v14 := GenS2{F0: 42}
                    var v15 uint16 = (v3 + 0)
                    _ = v14
                    _ = v15
                }
                _ = fi11
            }
        }
    }
    v16 := GenS1{F0: (995 != genFn0(2147483647))}
    if true {
        var v17 uint32 = 4294967295
        if true {
            v18 := GenS3{F0: 0}
            if len(v1) > 0 {
                v1[0] = v17
            }
            _ = v18
        } else {
            var v19 int8 = -125
            _ = v19
        }
        fmt.Println(v3, v2, v3)
        _ = v17
    } else {
        var v20 uint64 = 4294967295
        _ = v20
    }
    for fi21 := 0; fi21 < 17; fi21++ {
        v3 = (v3 >> ((704) & 63))
        _ = fi21
    }
    fmt.Println(v3, v2)
    fmt.Println(v2, v2)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 uint8 = 82
        _ = fi22
        _ = v23
    }
    fmt.Println(v3, v3)
    for fi24 := 0; fi24 < 8; fi24++ {
        fmt.Println(v3)
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
}
