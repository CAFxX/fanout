package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint8; F1 int; F2 int }
type GenS4 struct { F0 uint; F1 uint8 }
type GenS5 struct { F0 int; F1 int }
type GenS6 struct { F0 bool; F1 uint; F2 int16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 uint64 }
type GenS9 struct { F0 string; F1 int8; F2 int16 }
func genFn0(p0 uint16, p1 uint16) uint64 {
    p1 = (p0 << ((42) & 63))
    fmt.Println(p0)
    if false {
        var v6 float64 = 1.5
        _ = v6
    } else {
        if false {
            v7 := []int{4294967295, 127, 2147483647, 52}
            v8 := GenS1{F0: 1, F1: 42, F2: (32767 <= -1)}
            v9 := GenS1{F0: 68, F1: 1, F2: true}
            _ = v7
            _ = v8
            _ = v9
        } else {
            v10 := []int{685, 32768, 372}
            v11 := GenS2{F0: (-123 <= 230)}
            v12 := GenS3{F0: 1, F1: 42, F2: -413}
            _ = v10
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := GenS2{F0: true}
            _ = fi13
            _ = v14
        }
    }
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return uint64(p0)
}

func main() {
    var v1 uint16 = 1
    v2 := []int{42, 32767}
    v3 := make(map[int]uint32)
    v4 := []uint32{1, 615}
    var v5 float32 = -2.25
    delete(v3, 1)
    switch 10 {
    case -36:
        fmt.Println(v5, v1)
    case -247:
        fmt.Println(v1)
    default:
        if false {
            var v15 int = 0
            v16 := GenS4{F0: 415, F1: 128}
            v1 = (v1 / ((255) | 1))
            _ = v15
            _ = v16
        }
    }
    if false {
        if (true || (false && (!false))) {
            var v17 uint16 = (v1 & v1)
            for fi18 := 0; fi18 < 9; fi18++ {
                switch 1 {
                case -1:
                    var v19 float32 = ((v5 - float32(3.14159)) * -2.25)
                    _ = v19
                default:
                    switch "hello" {
                    case "hello":
                        fmt.Println(fi18)
                    case "a":
                        v20 := GenS5{F0: 128, F1: fi18}
                        _ = v20
                    case "\n\t":
                        for fi21 := 0; fi21 < 17; fi21++ {
                            v3[-2147483648] = 105
                            _ = fi21
                        }
                    default:
                        var v22 uint16 = 65535
                        _ = v22
                    }
                }
                fmt.Println(v17, v1, v17)
                _ = fi18
            }
            if false {
                v23 := GenS6{F0: false, F1: 920, F2: -32768}
                v24 := GenS7{F0: ""}
                v25 := GenS5{F0: 0, F1: -20}
                _ = v23
                _ = v24
                _ = v25
            }
            _ = v17
        }
    } else {
        fmt.Println(v5, v5)
    }
    var v26 float32 = v5
    if (float32(0.1) <= (v5 + 0.1)) {
        _ = len(v3)
        var v27 int16 = 1
        if (!false) {
            var v28 uint64 = 505
            v28 -= (v28 * 0)
            _ = v28
        }
        _ = v27
    }
    v26 -= (0.1 * 1e10)
    v1 += v1
    v1 = (v1 + v1)
    if (-2147483648 == -2147483648) {
        if (4294967295 == 4294967295) {
            if len(v2) > 0 {
                v2[0] = -2147483648
            }
            var v29 uint = 423
            var v30 int8 = 30
            _ = v29
            _ = v30
        } else {
            fmt.Println(v5, v26, v5)
            v31 := GenS8{F0: -128, F1: 0}
            v1 = (v1 * 120)
            _ = v31
        }
        fmt.Println(v5, v26, v1)
    }
    var v32 string = ("a" + "foo bar")
    fmt.Println(v26, v26)
    for fi33 := 0; fi33 < 12; fi33++ {
        v5 += 1e10
        v34 := GenS9{F0: "foo bar", F1: 42, F2: 1}
        fmt.Println(v1, v5)
        _ = fi33
        _ = v34
    }
    v1 += v1
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v5
}
