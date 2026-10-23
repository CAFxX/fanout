package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 string; F2 int16 }
type GenS3 struct { F0 uint; F1 int16; F2 int64 }
type GenS4 struct { F0 int }
func genFn0(p0 uint64, p1 uint) uint {
    v6 := GenS1{F0: -124}
    p1 -= 691
    v7 := []int{}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 814
}
func genFn1(p0 uint8, p1 int32, p2 int) uint16 {
    var v8 uint64 = 42
    v9 := GenS1{F0: int8(p1)}
    v10 := []int{p2, -2147483648, 10}
    switch "foo bar" {
    case "hello":
        for fi11 := 0; fi11 < 8; fi11++ {
            if ((false && (!true)) && false) {
                v10 = append(v10, (^(p2)))
                fmt.Println(p2, v8)
            }
            _ = len(v10)
            fmt.Println(fi11, fi11, fi11)
            _ = fi11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return uint16(p0)
}

func main() {
    var v1 uint8 = 203
    var v2 float32 = 1e10
    var v3 string = "hello"
    v4 := []int{-9223372036854775808, 32767}
    var v5 string = v3
    v12 := GenS1{F0: 0}
    v5 = v3
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 uint32 = 2147483647
        _ = fi13
        _ = v14
    }
    var v15 uint16 = 462
    var v16 uint16 = v15
    if false {
        if true {
            fmt.Println(v2)
        } else {
            v17 := GenS2{F0: -32768, F1: "a", F2: -32768}
            _ = v17
        }
        v18 := GenS2{F0: 895, F1: v5, F2: 32767}
        _ = v18
    }
    v2 = float32(3.14159)
    for fi19 := 0; fi19 < 17; fi19++ {
        v20 := GenS3{F0: 18446744073709551615, F1: -1, F2: -207}
        if true {
            switch 2147483647 {
            case 121:
                if (!(genFn1(v1, -866, fi19) <= 255)) {
                    switch 832 {
                    case -2147483648:
                        v5 = v3
                    case -717:
                        v21 := GenS4{F0: (fi19 | (fi19 << ((genFn0(4294967295, 18446744073709551615)) & 63)))}
                        _ = v21
                    case 42:
                        var v22 uint32 = uint32(v16)
                        _ = v22
                    default:
                        var v23 bool = false
                        _ = v23
                    }
                    var v24 bool = ((0 != 949) || false)
                    _ = v24
                } else {
                    v1 -= 127
                    if (-275 != (fi19 / (((fi19 & (fi19 / ((812) | 1)))) | 1))) {
                        v1 = v1
                        for fi25 := 0; fi25 < 2; fi25++ {
                            var v26 uint16 = v16
                            _ = fi25
                            _ = v26
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case 749:
                fmt.Println(v3)
            case 42:
                var v27 uint16 = v15
                _ = v27
            }
            v28 := GenS2{F0: 0, F1: "foo bar", F2: 42}
            _ = v28
        } else {
            var v29 int32 = -1
            _ = v29
        }
        v2 -= (3.14159 + v2)
        _ = fi19
        _ = v20
    }
    var v30 int16 = -32768
    for fi31 := 0; fi31 < 9; fi31++ {
        var v32 uint8 = (v1 + (v1 | (^(v1))))
        _ = fi31
        _ = v32
    }
    v16 += genFn1(127, 280, 127)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
