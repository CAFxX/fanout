package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint; F1 int32; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 int64 }
type GenS5 struct { F0 uint; F1 uint32 }
type GenS6 struct { F0 uint8 }

func main() {
    var v1 int8 = 42
    var v2 uint16 = 255
    v3 := []int{-43, 1, 32768}
    var v4 int16 = 32767
    var v5 uint64 = 42
    v6 := GenS1{F0: 2, F1: 2}
    fmt.Println(v5)
    v7 := GenS2{F0: 255, F1: -227, F2: -9223372036854775808}
    var v8 float32 = 1.5
    var v9 float32 = (v8 + (v8 + v8))
    if true {
        fmt.Println(v2)
        v3 = append(v3, 820)
        for fi10 := 0; fi10 < 20; fi10++ {
            if false {
                v11 := GenS1{F0: 106, F1: 90}
                fmt.Println(v9, v5)
                _ = v11
            } else {
                var v12 uint8 = 0
                _ = v12
            }
            if len(v3) > 0 {
                v3[0] = -2147483648
            }
            _ = fi10
        }
    } else {
        for fi13 := 0; fi13 < 1; fi13++ {
            if (-51 >= (v1 << ((42) & 63))) {
                v9 = 0.1
                var v14 int32 = 0
                switch fi13 {
                case -695:
                    v15 := GenS1{F0: 4294967295, F1: 2}
                    _ = v15
                case -441:
                    v16 := GenS3{F0: v5}
                    _ = v16
                default:
                    v14 = (v14 / ((v14) | 1))
                }
                _ = v14
            } else {
                var v17 bool = (("a" + ("\n\t" + "hello")) == (("\n\t" + "hello") + "a"))
                v1 = v1
                var v18 uint64 = 169
                _ = v17
                _ = v18
            }
            v19 := GenS3{F0: (^(v5))}
            for fi20 := 0; fi20 < 7; fi20++ {
                var v21 uint8 = 213
                v22 := GenS4{F0: (v4 + v4), F1: 525, F2: 1}
                _ = fi20
                _ = v21
                _ = v22
            }
            _ = fi13
            _ = v19
        }
        var v23 float32 = float32(float64(0.0))
        _ = v23
    }
    for fi24 := 0; fi24 < 10; fi24++ {
        if true {
            var v25 uint32 = 2147483648
            var v26 float64 = (-2.25 + (3.14159 * ((((0.1 / -2.25) * 1.5) / 1.0) / float64(v8))))
            _ = v25
            _ = v26
        }
        var v27 int16 = (v4 >> ((1) & 63))
        v28 := GenS5{F0: 65535, F1: 65535}
        _ = fi24
        _ = v27
        _ = v28
    }
    v29 := GenS6{F0: 42}
    var v30 uint16 = (^(v2))
    v2 = v30
    for fi31 := 0; fi31 < 9; fi31++ {
        v3 = append(v3, (-(fi31)))
        v8 = v9
        for fi32 := 0; fi32 < 11; fi32++ {
            var v33 int16 = int16(v2)
            var v34 uint16 = (v30 % (((v30 ^ (v2 * v30))) | 1))
            _ = fi32
            _ = v33
            _ = v34
        }
        _ = fi31
    }
    if (v5 != v5) {
        if false {
            v8 = (float32(1e10) / v9)
            v35 := GenS3{F0: v5}
            _ = v35
        } else {
            v36 := GenS2{F0: 42, F1: 888, F2: -1}
            fmt.Println(v4, v8)
            _ = v36
        }
        var v37 int16 = -32768
        _ = len(v3)
        _ = v37
    }
    if false {
        var v38 bool = true
        if ("" == "foo bar") {
            v39 := GenS5{F0: 0, F1: 4294967295}
            var v40 uint64 = v5
            _ = v39
            _ = v40
        }
        var v41 uint16 = (v2 / ((v2) | 1))
        _ = v38
        _ = v41
    }
    fmt.Println(v5)
    v3 = append(v3, 32768)
    v4 -= (v4 - (v4 >> ((481) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
