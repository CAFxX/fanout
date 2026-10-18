package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 uint16 }
type GenS2 struct { F0 bool; F1 int16; F2 int64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16 }
type GenS5 struct { F0 uint32; F1 int64; F2 uint }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 65535
    v3 := make(map[uint32]bool)
    v4 := []uint32{0, 586, 556}
    v5 := []uint32{}
    var v6 int32 = -1
    if ((-(v6)) != 42) {
        v6 -= v6
    } else {
        if (4294967295 != 536) {
            v5 = append(v5, 4294967295)
            v7 := GenS1{F0: 65535, F1: -1, F2: (v2 / (((^(v2))) | 1))}
            _ = v7
        }
        var v8 int = -9223372036854775808
        _ = len(v1)
        _ = v8
    }
    v9 := v1[560]
    v10 := v3[651]
    var v11 int64 = 155
    fmt.Println(v9, v11, v11)
    _ = len(v3)
    if (false && (-131 > 1)) {
        fmt.Println(v9, v2, v10)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi12 := 0; fi12 < 5; fi12++ {
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(fi13, fi12)
                v6 -= 2147483647
                var v14 uint32 = 42
                _ = fi13
                _ = v14
            }
            var v15 bool = v10
            var v16 int8 = 100
            _ = fi12
            _ = v15
            _ = v16
        }
    } else {
        v17 := GenS1{F0: 2, F1: v11, F2: 0}
        v18 := GenS1{F0: 0, F1: v11, F2: v2}
        v19 := GenS2{F0: (("foo bar" + ("hello" + "foo bar")) >= "foo bar"), F1: -770, F2: (v11 & v11)}
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS2{F0: v10, F1: 0, F2: 1}
        v22 := GenS1{F0: 1, F1: (v11 + 843), F2: v2}
        v23 := GenS3{F0: ("a" + "hello")}
        _ = fi20
        _ = v21
        _ = v22
        _ = v23
    }
    fmt.Println(v9, v11, v9)
    var v24 int16 = 42
    var v25 uint = 65535
    fmt.Println(v9, v11, v24)
    if (v10 || v10) {
        var v26 float32 = 0.0
        v27 := GenS3{F0: ""}
        _ = v26
        _ = v27
    } else {
        v11 = (^(v11))
    }
    fmt.Println(v2, v9, v6)
    if (100 < -75) {
        v3[65535] = (!v10)
        v2 = v2
    }
    for fi28 := 0; fi28 < 10; fi28++ {
        delete(v3, 65535)
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := v1[2147483648]
            switch "foo bar" {
            case "hello":
                fmt.Println(v11, v2)
            case "a":
                var v31 uint64 = 4294967295
                _ = v31
            case "foo bar":
                if v10 {
                    var v32 int = (v30 * 58)
                    if true {
                        v33 := v1[764]
                        _ = v33
                    }
                    _ = v32
                } else {
                    v34 := GenS1{F0: v25, F1: 9223372036854775807, F2: 42}
                    if v10 {
                        if (-2.25 >= -2.25) {
                            if v10 {
                                v30 += (fi29 + (fi29 - -681))
                                delete(v3, 41)
                            } else {
                                v2 += (v2 % ((v2) | 1))
                                var v35 float64 = (-2.25 + -2.25)
                                v6 -= v6
                                _ = v35
                            }
                            v36 := GenS4{F0: 306, F1: v2}
                            _ = v36
                        } else {
                            fmt.Println(v30)
                        }
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = v34
                }
            }
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v37 := GenS5{F0: 592, F1: (v11 | v11), F2: (^(v25))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
