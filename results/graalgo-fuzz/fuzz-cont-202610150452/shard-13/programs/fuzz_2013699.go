package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint; F1 int64; F2 uint16 }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 uint16; F1 int16 }
func genFn0(p0 uint8, p1 uint8, p2 uint64) uint8 {
    v4 := []int{2147483647, 181, -9223372036854775808}
    if true {
        p2 += (p2 - p2)
    } else {
        var v5 uint64 = (p2 >> ((65535) & 63))
        _ = v5
    }
    v6 := GenS1{F0: -1, F1: -947}
    p2 += 18446744073709551615
    fmt.Println(p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v6
    return (p0 << ((65535) & 63))
}
func genFn1() bool {
    var v7 uint8 = 150
    fmt.Println(v7)
    fmt.Println(v7)
    switch -2147483648 {
    case 127:
        v8 := GenS2{F0: 2}
        _ = v8
    case -2147483648:
        var v9 int16 = -958
        _ = v9
    }
    _ = v7
    return true
}

func main() {
    var v1 int = 9223372036854775807
    var v2 bool = (1e10 <= (1.5 / 1.0))
    v3 := make(map[uint32]bool)
    v2 = (3.14159 > 1e10)
    var v10 int64 = -9
    var v11 int32 = -1
    var v12 bool = (v1 <= (v1 ^ v1))
    var v13 string = "hello"
    v1 = (v1 * v1)
    switch "hello" {
    case "\n\t":
        fmt.Println(v13, v1, v11)
    case "a":
        v2 = v12
    case "":
        var v14 int32 = 460
        _ = v14
    default:
        if (((3.14159 / 1.0) / (float64(float32(3.14159)) / -2.25)) != 1e10) {
            v2 = ((v12 || v12) && v12)
            v15 := []int{v1, v1}
            _ = v15
        } else {
            var v16 uint32 = 2
            for fi17 := 0; fi17 < 19; fi17++ {
                v2 = (!(!v12))
                if ((true || v2) && false) {
                    var v18 int32 = v11
                    _ = v18
                } else {
                    v19 := GenS2{F0: 677}
                    _ = v19
                }
                _ = fi17
            }
            v20 := GenS3{F0: 408, F1: (v11 / ((v11) | 1)), F2: v12}
            _ = v16
            _ = v20
        }
    }
    v21 := []int{}
    var v22 int32 = v11
    v23 := GenS3{F0: -2147483648, F1: v11, F2: (18446744073709551615 < 251)}
    if (-32768 < 32767) {
        var v24 uint8 = 128
        v10 -= v10
        var v25 uint8 = v24
        _ = v24
        _ = v25
    } else {
        var v26 int = v1
        if (!(2 < 0)) {
            for fi27 := 0; fi27 < 6; fi27++ {
                switch "\n\t" {
                case "\n\t":
                    v10 += v10
                case "hello":
                    var v28 string = "foo bar"
                    _ = v28
                default:
                    if v2 {
                        v22 -= (v22 >> ((797) & 63))
                        v13 = v13
                    } else {
                        if (2 > 4294967295) {
                            if v12 {
                                v29 := GenS3{F0: int32(v26), F1: v11, F2: v12}
                                fmt.Println(v13, v12)
                                _ = v29
                            }
                        } else {
                            var v30 uint16 = 1
                            fmt.Println(v30, v11, v12)
                            _ = v30
                        }
                        v21 = append(v21, v26)
                    }
                }
                _ = fi27
            }
        } else {
            if len(v21) > 0 {
                v21[0] = (v1 & v26)
            }
        }
        v22 = (v22 >> ((2) & 63))
        _ = v26
    }
    v31 := GenS4{F0: v12}
    if false {
        v32 := GenS2{F0: 60}
        _ = v32
    } else {
        v33 := GenS5{F0: 42, F1: int64(v11), F2: 255}
        v34 := v3[863]
        v35 := GenS6{F0: 88, F1: 289}
        _ = v33
        _ = v34
        _ = v35
    }
    v36 := GenS7{F0: 661, F1: 1}
    var v37 bool = (0 < 255)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v36
    _ = v37
}
