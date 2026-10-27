package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 int16 }
type GenS2 struct { F0 bool; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 int16; F1 int; F2 bool }
func genFn0(p0 uint8, p1 int32, p2 int) bool {
    if (("hello" + "foo bar") == "") {
        switch (p2 >> ((18446744073709551615) & 63)) {
        case -1:
            v6 := GenS1{F0: p0, F1: (2147483647 == 2147483647), F2: -1}
            _ = v6
        case 128:
            for fi7 := 0; fi7 < 18; fi7++ {
                fmt.Println(p0, p1)
                _ = fi7
            }
        }
        var v8 uint64 = 4294967295
        _ = v8
    } else {
        v9 := GenS1{F0: p0, F1: true, F2: 759}
        _ = v9
    }
    var v10 bool = ((p2 ^ int(p1)) <= p2)
    p1 += p1
    var v11 float32 = 1.5
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (4294967295 < 18446744073709551615)
}
func genFn1(p0 int8, p1 uint32) int8 {
    v12 := GenS1{F0: 74, F1: false, F2: 1}
    for fi13 := 0; fi13 < 9; fi13++ {
        for fi14 := 0; fi14 < 14; fi14++ {
            v15 := []int{4294967295}
            if (false && false) {
                switch "" {
                case "\n\t":
                    switch ((("" + ("a" + "foo bar")) + "foo bar") + ("a" + "")) {
                    case "foo bar":
                        v16 := GenS2{F0: (!false), F1: -9223372036854775808, F2: 0}
                        _ = v16
                    case "\n\t":
                        if len(v15) > 0 {
                            v15[0] = -435
                        }
                    default:
                        fmt.Println(p1)
                        v15 = append(v15, 653)
                    }
                case "hello":
                    var v17 uint = 0
                    if ((false && false) || (1.5 <= 1e10)) {
                        var v18 float64 = -2.25
                        _ = v18
                    }
                    if true {
                        switch (-(fi14)) {
                        case 434:
                            fmt.Println(fi14)
                        case 128:
                            v17 -= v17
                        case 32767:
                            p0 = (^(p0))
                        default:
                            var v19 float32 = 1e10
                            _ = v19
                        }
                    }
                    _ = v17
                case "a":
                    if len(v15) >= 2 {
                        v15 = v15[1:len(v15)]
                    }
                }
                p0 -= (p0 >> ((4294967295) & 63))
                var v20 int = fi13
                _ = v20
            } else {
                var v21 int8 = (p0 & -43)
                v21 -= (p0 << ((1) & 63))
                _ = v21
            }
            _ = fi14
            _ = v15
        }
        for fi22 := 0; fi22 < 2; fi22++ {
            v23 := []int{}
            var v24 int32 = -1
            _ = fi22
            _ = v23
            _ = v24
        }
        _ = fi13
    }
    if (255 != 255) {
        var v25 uint8 = 2
        _ = v25
    }
    v26 := []int{-1, -9223372036854775808, 2147483647}
    _ = p0
    _ = p1
    _ = v12
    _ = v26
    return p0
}

func main() {
    v1 := []uint32{2147483648}
    v2 := make(map[string]int)
    var v3 bool = (!(!false))
    var v4 uint64 = 802
    var v5 int32 = int32(v4)
    var v27 int64 = 0
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    if ((!v3) || (39 > 138)) {
        v28 := GenS3{F0: 2, F1: (v4 ^ v4)}
        v29 := GenS4{F0: -32768, F1: 128, F2: v3}
        _ = v28
        _ = v29
    } else {
        for fi30 := 0; fi30 < 12; fi30++ {
            v3 = (691 >= 0)
            _ = fi30
        }
        fmt.Println(v5, v4, v27)
    }
    var v31 int8 = 127
    v31 -= (v31 << ((65535) & 63))
    if (127 > 32767) {
        v3 = v3
        var v32 float32 = 0.1
        _ = v32
    }
    v4 = 237
    for fi33 := 0; fi33 < 10; fi33++ {
        var v34 int8 = v31
        _ = fi33
        _ = v34
    }
    var v35 int = 0
    if v3 {
        var v36 uint = 0
        _ = v36
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
