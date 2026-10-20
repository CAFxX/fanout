package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32 }
func genFn0(p0 int16) int64 {
    var v6 float64 = float64(float32(0.1))
    p0 -= (p0 << ((18446744073709551615) & 63))
    v7 := []int{128}
    v6 = (v6 - 3.14159)
    _ = p0
    _ = v6
    _ = v7
    return 774
}
func genFn1(p0 string, p1 int32, p2 int32) string {
    fmt.Println(p0, p1, p0)
    p2 += p1
    _ = p0
    _ = p1
    _ = p2
    return p0
}

func main() {
    var v1 uint16 = 179
    v2 := []uint32{2, 4294967295, 2, 65535}
    var v3 uint32 = 1
    v4 := []uint32{}
    var v5 float64 = -2.25
    v3 = v3
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 int8 = 42
        switch (fi8 & fi8) {
        case -820:
            var v10 uint64 = 18446744073709551615
            _ = v10
        }
        _ = fi8
        _ = v9
    }
    v1 = (v1 << ((42) & 63))
    switch 1 {
    case 2147483647:
        v11 := GenS1{F0: 127, F1: 4294967295}
        _ = v11
    case 128:
        if false {
            v12 := GenS2{F0: 2, F1: ("" + "a")}
            v1 = v1
            for fi13 := 0; fi13 < 6; fi13++ {
                if (-110 >= -110) {
                    var v14 int16 = -1
                    _ = v14
                }
                v15 := GenS3{F0: false}
                if ((v5 / ((0.0 - v5) * 0.1)) == v5) {
                    if len(v4) > 0 {
                        v4[0] = (^(v3))
                    }
                    fmt.Println(fi13)
                    v5 = (1.5 * 3.14159)
                } else {
                    var v16 int16 = 42
                    var v17 uint64 = 258
                    v16 = (v16 << ((504) & 63))
                    _ = v16
                    _ = v17
                }
                _ = fi13
                _ = v15
            }
            _ = v12
        }
    }
    v18 := GenS1{F0: 144, F1: 2}
    fmt.Println(v1, v5, v3)
    if false {
        if true {
            v5 = v5
            v3 -= (v3 ^ 2147483648)
            v19 := GenS4{F0: v3}
            _ = v19
        } else {
            fmt.Println(v5, v3, v3)
        }
    }
    v3 = v3
    for fi20 := 0; fi20 < 12; fi20++ {
        v3 = v3
        _ = fi20
    }
    var v21 uint8 = 127
    if false {
        if (((0.1 / 0.1) * (((1e10 + 1.5) - 1.5) * (3.14159 / (0.1 / ((1.5 / 1.0) - 1e10))))) < 1e10) {
            for fi22 := 0; fi22 < 3; fi22++ {
                fmt.Println(fi22)
                v4 = append(v4, (v3 - (v3 + v3)))
                _ = fi22
            }
        }
        var v23 int8 = -1
        var v24 bool = (472 > 42)
        _ = v23
        _ = v24
    } else {
        var v25 uint = 2
        fmt.Println(v1, v21, v3)
        v3 = (v3 & (v3 + (v3 + 2147483647)))
        _ = v25
    }
    var v26 int64 = -696
    var v27 int8 = -128
    v5 = 3.14159
    _ = len(v4)
    if true {
        for fi28 := 0; fi28 < 10; fi28++ {
            for fi29 := 0; fi29 < 7; fi29++ {
                if true {
                    if (true && true) {
                        v21 += (v21 * (v21 + v21))
                    }
                    var v30 float64 = v5
                    v2 = append(v2, v3)
                    _ = v30
                }
                _ = fi29
            }
            v3 = v3
            _ = fi28
        }
        if false {
            if false {
                var v31 bool = true
                fmt.Println(v1, v21)
                _ = v31
            } else {
                var v32 int = 0
                _ = v32
            }
            fmt.Println(v3)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        for fi33 := 0; fi33 < 3; fi33++ {
            var v34 float64 = v5
            _ = len(v2)
            var v35 float64 = v5
            _ = fi33
            _ = v34
            _ = v35
        }
    } else {
        var v36 uint8 = v21
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
