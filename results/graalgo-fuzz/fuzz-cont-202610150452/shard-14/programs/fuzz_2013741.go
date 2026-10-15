package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint32; F1 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int8 }
func genFn0(p0 uint32, p1 string, p2 uint) int {
    for fi6 := 0; fi6 < 8; fi6++ {
        var v7 int64 = 525
        _ = fi6
        _ = v7
    }
    p0 += 2
    if (!false) {
        p1 = p1
        v8 := []int{-2147483648, 0, 936, -245}
        _ = v8
    } else {
        var v9 uint32 = (p0 % ((p0) | 1))
        _ = v9
    }
    var v10 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -9223372036854775808
}
func genFn1() string {
    switch 407 {
    case 901:
        var v11 uint64 = 2
        _ = v11
    case 42:
        var v12 int16 = -32768
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v14 := GenS1{F0: 117, F1: 9223372036854775807, F2: 0}
        _ = fi13
        _ = v14
    }
    return ""
}

func main() {
    v1 := []int{}
    var v2 float32 = (1e10 - 1.5)
    var v3 int64 = -1
    var v4 uint16 = 1
    var v5 uint = 0
    var v15 bool = false
    switch 4294967295 {
    case 786:
        if false {
            var v16 uint32 = 0
            for fi17 := 0; fi17 < 3; fi17++ {
                v5 -= (^(v5))
                v3 = (v3 - v3)
                fmt.Println(v5)
                _ = fi17
            }
            if v15 {
                var v18 float32 = (v2 - -2.25)
                _ = len(v1)
                switch int(v4) {
                case 32768:
                    if v15 {
                        var v19 bool = true
                        for fi20 := 0; fi20 < 3; fi20++ {
                            var v21 int = -290
                            for fi22 := 0; fi22 < 19; fi22++ {
                                v21 -= (v21 >> ((v5) & 63))
                                if len(v1) > 0 {
                                    v1[0] = (fi22 >> ((v5) & 63))
                                }
                                _ = fi22
                            }
                            var v23 uint32 = (^(v16))
                            _ = fi20
                            _ = v21
                            _ = v23
                        }
                        v24 := GenS2{F0: v16, F1: -102}
                        _ = v19
                        _ = v24
                    } else {
                        v25 := GenS2{F0: uint32(v4), F1: 1}
                        v4 = 803
                        v1 = append(v1, 42)
                        _ = v25
                    }
                case 32767:
                    var v26 string = "a"
                    _ = v26
                case 4294967295:
                    v2 = v18
                }
                _ = v18
            }
            _ = v16
        } else {
            var v27 float32 = float32((3.14159 / 3.14159))
            v3 += v3
            _ = v27
        }
    case -9223372036854775808:
        v28 := GenS3{F0: v3}
        _ = v28
    }
    var v29 uint8 = 42
    switch genFn0(1, ("hello" + "a"), 65535) {
    case 128:
        v30 := GenS4{F0: -1}
        _ = v30
    }
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if len(v1) > 0 {
        v1[0] = 9223372036854775807
    }
    var v31 float32 = v2
    var v32 float64 = 0.0
    for fi33 := 0; fi33 < 9; fi33++ {
        switch fi33 {
        case 1:
            v5 += (v5 % ((v5) | 1))
        case 4294967295:
            v31 += (v31 * (v2 - v2))
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = fi33
    }
    if (v2 != (float32(1e10) * (float32(v32) + v2))) {
        v34 := GenS4{F0: 100}
        _ = v34
    } else {
        v5 = (v5 + v5)
        var v35 uint8 = (v29 | 42)
        _ = v35
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
