package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 string }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 uint16; F1 uint8; F2 string }
func genFn0(p0 uint, p1 bool) uint {
    if (65535 < p0) {
        var v5 int64 = 591
        var v6 int8 = -128
        switch 2147483647 {
        case -1:
            if (2 >= 18) {
                v6 = v6
                v7 := GenS1{F0: v6, F1: 0, F2: "foo bar"}
                _ = v7
            } else {
                var v8 uint8 = 1
                var v9 float64 = ((float64(-2.25) / 1.0) - 3.14159)
                _ = v8
                _ = v9
            }
        default:
            var v10 uint32 = 2
            _ = v10
        }
        _ = v5
        _ = v6
    } else {
        var v11 uint64 = 383
        _ = v11
    }
    var v12 int16 = 881
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 13; fi14++ {
            fmt.Println(fi14)
            _ = fi14
        }
        for fi15 := 0; fi15 < 20; fi15++ {
            v16 := []int{fi15}
            _ = fi15
            _ = v16
        }
        v17 := GenS2{F0: fi13, F1: 9223372036854775807}
        _ = fi13
        _ = v17
    }
    for fi18 := 0; fi18 < 17; fi18++ {
        v19 := []int{1, (fi18 / ((fi18) | 1)), (fi18 & fi18)}
        v20 := GenS1{F0: 127, F1: 127, F2: ("" + "\n\t")}
        _ = fi18
        _ = v19
        _ = v20
    }
    switch ("hello" + "a") {
    case "hello":
        p0 = (p0 * p0)
    default:
        v12 += v12
    }
    _ = p0
    _ = p1
    _ = v12
    return (p0 >> ((18446744073709551615) & 63))
}

func main() {
    var v1 float64 = float64(1e10)
    v2 := make(map[int]bool)
    v3 := []uint32{2147483647, 42}
    v4 := []uint32{2, 2, 501}
    _ = len(v2)
    var v21 int8 = -82
    if true {
        v21 = v21
        var v22 uint16 = 2
        v1 += v1
        _ = v22
    } else {
        var v23 uint32 = 114
        _ = v23
    }
    v24 := v2[42]
    _ = len(v2)
    if (1000 >= 1000) {
        v3 = append(v3, 2147483648)
        v25 := v2[128]
        _ = v25
    }
    switch "\n\t" {
    case "a":
        var v26 int16 = 1000
        _ = v26
    case "foo bar":
        if (!false) {
            fmt.Println(v1, v21, v21)
            if v24 {
                if true {
                    var v27 uint = 2
                    _ = v27
                }
                var v28 string = ("hello" + ("\n\t" + "hello"))
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                v1 += v1
                if (42 != 93) {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    var v30 int16 = -299
                    for fi31 := 0; fi31 < 1; fi31++ {
                        if v24 {
                            _ = len(v3)
                            var v32 int64 = 42
                            v21 -= v21
                            _ = v32
                        }
                        _ = fi31
                    }
                    _ = v30
                }
                var v33 bool = true
                _ = fi29
                _ = v33
            }
        } else {
            v24 = ((!(("foo bar" + "hello") < (((((("foo bar" + "hello") + "a") + "hello") + "\n\t") + "foo bar") + "hello"))) || (("foo bar" + "hello") < "a"))
        }
    case "":
        v24 = (2 <= 42)
    default:
        v21 = v21
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS3{F0: 961, F1: 2, F2: ("" + "a")}
        var v36 uint = 4294967295
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v24
    _ = v3
    _ = v4
}
