package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 bool }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 uint16 }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint32 }
func genFn0(p0 int16) bool {
    if true {
        if (1 == 0) {
            if false {
                var v4 uint64 = 2
                var v5 uint8 = 128
                v4 = 1
                _ = v4
                _ = v5
            } else {
                switch "\n\t" {
                case "foo bar":
                    for fi6 := 0; fi6 < 5; fi6++ {
                        if (!false) {
                            v7 := GenS1{F0: -9223372036854775808, F1: 42, F2: (true || true)}
                            var v8 uint32 = 536
                            _ = v7
                            _ = v8
                        } else {
                            v9 := []int{fi6, (fi6 * 4294967295), fi6, -9223372036854775808}
                            _ = v9
                        }
                        fmt.Println(fi6, fi6)
                        for fi10 := 0; fi10 < 20; fi10++ {
                            if false {
                                p0 = (-(p0))
                                var v11 int = 99
                                _ = v11
                            } else {
                                fmt.Println(fi6, p0)
                                var v12 int8 = -96
                                _ = v12
                            }
                            v13 := GenS1{F0: 42, F1: 42, F2: false}
                            _ = fi10
                            _ = v13
                        }
                        _ = fi6
                    }
                }
            }
            var v14 int16 = p0
            fmt.Println(v14, v14)
            _ = v14
        } else {
            var v15 float32 = -2.25
            _ = v15
        }
        fmt.Println(p0)
        p0 = (p0 << ((4294967295) & 63))
    } else {
        var v16 uint8 = 117
        _ = v16
    }
    p0 = (p0 * 1000)
    fmt.Println(p0)
    var v17 float32 = 1.5
    var v18 uint8 = 128
    _ = p0
    _ = v17
    _ = v18
    return false
}
func genFn1(p0 int32) uint64 {
    if (!(true || (733 != 558))) {
        v19 := GenS2{F0: "hello"}
        _ = v19
    } else {
        if false {
            p0 = p0
            p0 += p0
        } else {
            if false {
                var v20 bool = (-721 > -676)
                var v21 int64 = 9223372036854775807
                var v22 int8 = 1
                _ = v20
                _ = v21
                _ = v22
            }
        }
        if false {
            v23 := []int{42}
            _ = v23
        }
    }
    fmt.Println(p0)
    v24 := []int{}
    p0 -= (p0 | -183)
    p0 = (p0 + 42)
    _ = p0
    _ = v24
    return 4294967295
}

func main() {
    v1 := []int{}
    v2 := []int{-850, -2147483648, 1, -2147483648}
    var v3 int64 = 812
    v25 := GenS3{F0: 42, F1: 2}
    fmt.Println(v3)
    fmt.Println(v3)
    var v26 int64 = -724
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v27 uint16 = 369
    v26 += (-(v26))
    v28 := GenS4{F0: genFn1(521)}
    for fi29 := 0; fi29 < 16; fi29++ {
        var v30 float32 = 0.1
        if (float64(v30) == 0.0) {
            var v31 int = (fi29 >> ((0) & 63))
            _ = v31
        }
        fmt.Println(v27, v26, v30)
        _ = fi29
        _ = v30
    }
    var v32 float32 = float32(0.1)
    v3 = (v3 & v26)
    for fi33 := 0; fi33 < 12; fi33++ {
        fmt.Println(v3, v32)
        _ = fi33
    }
    var v34 uint16 = 255
    if false {
        var v35 float32 = (v32 * float32(1.5))
        var v36 int16 = -32768
        v37 := GenS5{F0: (v26 << ((550) & 63)), F1: v27}
        _ = v35
        _ = v36
        _ = v37
    }
    v38 := GenS6{F0: 898}
    v1 = append(v1, 9223372036854775807)
    fmt.Println(v34, v3)
    if false {
        var v39 uint8 = 127
        var v40 float32 = (1e10 * ((v32 - (0.1 * 1.5)) * (((float32(1e10) - v32) / v32) / 1.0)))
        _ = v39
        _ = v40
    } else {
        v32 -= v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v34
    _ = v38
}
