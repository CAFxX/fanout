package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string }
func genFn0(p0 string, p1 uint32) int64 {
    v6 := []int{42, -2147483648}
    if true {
        if (42 <= uint64(p1)) {
            p1 += 1
        } else {
            p1 -= p1
            if len(v6) > 0 {
                v6[0] = -532
            }
            v7 := GenS1{F0: (!false), F1: 42}
            _ = v7
        }
        fmt.Println(p0, p0)
    }
    if len(v6) > 0 {
        v6[0] = 0
    }
    var v8 string = ("hello" + p0)
    for fi9 := 0; fi9 < 5; fi9++ {
        p1 += 671
        switch 2147483647 {
        case 2147483647:
            p1 = (p1 << ((0) & 63))
        case 128:
            var v10 uint64 = 2
            _ = v10
        case -9223372036854775808:
            for fi11 := 0; fi11 < 13; fi11++ {
                var v12 int64 = 1
                fmt.Println(v8, fi11, v12)
                p0 = (v8 + p0)
                _ = fi11
                _ = v12
            }
        }
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v8
    return 247
}
func genFn1(p0 uint8) int32 {
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: ("hello" != "hello"), F1: 1}
        _ = fi13
        _ = v14
    }
    var v15 string = "\n\t"
    p0 += (p0 % ((p0) | 1))
    _ = p0
    _ = v15
    return 499
}

func main() {
    var v1 uint8 = 42
    v2 := []uint32{}
    var v3 string = "a"
    var v4 int8 = -1
    var v5 float32 = (1.5 + 0.1)
    var v16 uint = 342
    v16 = (^(v16))
    v17 := GenS1{F0: true, F1: 9223372036854775807}
    fmt.Println(v5)
    v5 -= 1e10
    if false {
        if true {
            if ((!(false || false)) || ((v1 + v1) == (^(v1)))) {
                var v18 int16 = 32767
                _ = v18
            } else {
                if false {
                    var v19 uint64 = 18446744073709551615
                    _ = v19
                } else {
                    v4 = v4
                    v3 = v3
                }
                _ = len(v2)
                for fi20 := 0; fi20 < 8; fi20++ {
                    switch -12 {
                    case 9223372036854775807:
                        var v21 int8 = (-(v4))
                        _ = v21
                    }
                    var v22 uint8 = v1
                    var v23 float32 = (float32((0.1 * (3.14159 - ((3.14159 * -2.25) + 1.5)))) / 1e10)
                    _ = fi20
                    _ = v22
                    _ = v23
                }
            }
        }
        for fi24 := 0; fi24 < 13; fi24++ {
            var v25 bool = true
            for fi26 := 0; fi26 < 5; fi26++ {
                v1 -= v1
                v27 := GenS2{F0: 42}
                v5 = (v5 / 3.14159)
                _ = fi26
                _ = v27
            }
            v28 := GenS3{F0: -446}
            _ = fi24
            _ = v25
            _ = v28
        }
        v3 = ""
    }
    fmt.Println(v5)
    v29 := GenS4{F0: (v3 + v3)}
    v2 = append(v2, 1)
    if false {
        for fi30 := 0; fi30 < 12; fi30++ {
            v2 = append(v2, 42)
            var v31 float32 = ((v5 / (v5 / ((float32(-2.25) * (v5 * (0.0 * ((3.14159 * ((v5 - v5) * (v5 - float32(3.14159)))) / 0.1)))) * (v5 * 1.5)))) * v5)
            fmt.Println(v1, v4, v4)
            _ = fi30
            _ = v31
        }
        fmt.Println(v16)
        var v32 int16 = 1000
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
