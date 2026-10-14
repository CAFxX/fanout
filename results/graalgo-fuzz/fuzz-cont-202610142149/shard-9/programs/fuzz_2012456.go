package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 int64 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint16 }
type GenS3 struct { F0 int; F1 uint64; F2 int8 }
type GenS4 struct { F0 int32; F1 int8 }
type GenS5 struct { F0 uint16; F1 uint8 }
func genFn0() uint {
    v4 := []int{32768, 32767, 4294967295}
    v5 := GenS1{F0: 961, F1: -32768, F2: -9223372036854775808}
    var v6 int16 = -32768
    for fi7 := 0; fi7 < 5; fi7++ {
        _ = len(v4)
        for fi8 := 0; fi8 < 17; fi8++ {
            fmt.Println(v6)
            _ = fi8
        }
        switch (((("a" + "foo bar") + "") + "foo bar") + ("" + "hello")) {
        case "a":
            var v9 int32 = 2147483647
            _ = v9
        case "":
            var v10 int32 = 0
            _ = v10
        default:
            fmt.Println(v6, v6)
        }
        _ = fi7
    }
    _ = v4
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 uint32, p2 bool) uint {
    v11 := []int{-9223372036854775808, -564, 2147483647, 128}
    switch 128 {
    case 2147483647:
        v12 := GenS1{F0: p0, F1: -714, F2: -1}
        _ = v12
    case 42:
        if (42 == -1) {
            for fi13 := 0; fi13 < 10; fi13++ {
                if (float64(0.0) <= (float64(1e10) + 0.1)) {
                    p2 = false
                } else {
                    if len(v11) > 0 {
                        v11[0] = (fi13 + fi13)
                    }
                    fmt.Println(p2, p1)
                }
                p2 = (!(3.14159 < 0.1))
                v11 = append(v11, 0)
                _ = fi13
            }
            var v14 uint64 = 0
            _ = v14
        } else {
            _ = len(v11)
            fmt.Println(p1, p1, p2)
            p2 = p2
        }
    default:
        for fi15 := 0; fi15 < 9; fi15++ {
            if len(v11) > 0 {
                v11[0] = (fi15 << ((265) & 63))
            }
            v16 := GenS1{F0: (^(p0)), F1: 32767, F2: 9223372036854775807}
            fmt.Println(p1, p1)
            _ = fi15
            _ = v16
        }
    }
    switch (("foo bar" + "hello") + "foo bar") {
    case "hello":
        var v17 int64 = -9223372036854775808
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return genFn0()
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int64 = -1
    var v3 string = ""
    v2 = 1
    v18 := GenS1{F0: 834, F1: 32767, F2: v2}
    v2 -= -9223372036854775808
    v19 := GenS2{F0: 4294967295, F1: 127, F2: 65535}
    v3 = "a"
    v1[4294967295] = 2
    if (!true) {
        _ = len(v1)
    } else {
        v20 := GenS3{F0: -2147483648, F1: 42, F2: 27}
        v21 := v1[1]
        _ = v20
        _ = v21
    }
    switch (("a" + v3) + ("" + v3)) {
    case "a":
        switch "a" {
        case "foo bar":
            v2 = v2
        }
    case "\n\t":
        v2 = (v2 - (^(v2)))
    }
    fmt.Println(v3)
    v22 := []int{-471, -179, 127, -2147483648}
    for fi23 := 0; fi23 < 4; fi23++ {
        switch 127 {
        case -506:
            if (((-2.25 - 0.1) / 0.1) >= (0.0 - -2.25)) {
                fmt.Println(fi23, v3)
                if (false || false) {
                    v24 := GenS4{F0: 1, F1: 100}
                    for fi25 := 0; fi25 < 20; fi25++ {
                        if true {
                            fmt.Println(fi25, v2, v3)
                            var v26 string = (v3 + v3)
                            _ = v26
                        } else {
                            var v27 int8 = -1
                            var v28 uint32 = 42
                            v28 += (v28 ^ (^(v28)))
                            _ = v27
                            _ = v28
                        }
                        v29 := GenS5{F0: 32767, F1: 127}
                        _ = fi25
                        _ = v29
                    }
                    _ = v24
                }
                var v30 int8 = 0
                _ = v30
            } else {
                v2 = v2
            }
        case -2147483648:
            for fi31 := 0; fi31 < 2; fi31++ {
                v32 := GenS5{F0: 0, F1: 128}
                v33 := v1[1]
                _ = len(v1)
                _ = fi31
                _ = v32
                _ = v33
            }
        case -873:
            v2 = -408
        }
        _ = fi23
    }
    v2 -= (v2 % ((v2) | 1))
    fmt.Println(v3)
    for fi34 := 0; fi34 < 7; fi34++ {
        fmt.Println(fi34, fi34)
        _ = fi34
    }
    var v35 int32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v35
}
