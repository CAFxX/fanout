package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int32; F2 int16 }
func genFn0(p0 uint64, p1 uint16, p2 uint8) int16 {
    if (307 < 9223372036854775807) {
        v5 := GenS1{F0: 9223372036854775807}
        v6 := []int{-2147483648}
        _ = v5
        _ = v6
    }
    v7 := GenS1{F0: 32}
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return -32768
}
func genFn1() bool {
    var v9 uint64 = 198
    v10 := []int{-1}
    for fi11 := 0; fi11 < 14; fi11++ {
        v9 -= 2
        _ = fi11
    }
    var v12 int16 = -32768
    v9 -= 18446744073709551615
    _ = v10
    _ = v12
    _ = v9
    return (2147483647 < 32767)
}

func main() {
    var v1 bool = (0.0 < 1e10)
    var v2 bool = ((100 <= 42) || (0 >= 2147483648))
    var v3 uint64 = 1
    var v4 uint = 0
    for fi13 := 0; fi13 < 4; fi13++ {
        v4 = (^(v4))
        var v14 int = (-(fi13))
        _ = fi13
        _ = v14
    }
    v4 += v4
    v3 += v3
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint16 = 616
        for fi17 := 0; fi17 < 12; fi17++ {
            var v18 int16 = 604
            fmt.Println(v16, v4)
            _ = fi17
            _ = v18
        }
        _ = fi15
        _ = v16
    }
    if true {
        var v19 uint = 42
        _ = v19
    } else {
        for fi20 := 0; fi20 < 1; fi20++ {
            switch "a" {
            case "":
                v21 := GenS2{F0: 0, F1: 42, F2: 1000}
                _ = v21
            default:
                v1 = v2
            }
            var v22 uint64 = (v3 | (v3 ^ v3))
            _ = fi20
            _ = v22
        }
        var v23 int = 0
        _ = v23
    }
    v2 = (635 <= 0)
    v24 := GenS2{F0: 42, F1: -1, F2: 0}
    if (v3 > 146) {
        v25 := []int{9223372036854775807}
        _ = v25
    } else {
        v1 = true
    }
    v2 = (!((32767 > 32767) && (false || (v1 && v1))))
    var v26 int32 = -1
    v27 := []int{32767, 795, -331, -623}
    switch "hello" {
    case "a":
        if v2 {
            if (1 < 127) {
                for fi28 := 0; fi28 < 3; fi28++ {
                    var v29 int = (fi28 | -2147483648)
                    _ = fi28
                    _ = v29
                }
            } else {
                v30 := GenS1{F0: -336}
                _ = v30
            }
            v26 -= v26
        } else {
            v2 = v1
        }
    case "hello":
        if v1 {
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 int8 = 127
                _ = fi31
                _ = v32
            }
        }
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        for fi34 := 0; fi34 < 7; fi34++ {
            var v35 string = ""
            var v36 int8 = 127
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
