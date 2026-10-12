package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 uint, p1 uint, p2 int64) int {
    p2 = p2
    if true {
        v5 := GenS1{F0: 128, F1: -1}
        fmt.Println(p2, p2)
        if false {
            var v6 int64 = (^(p2))
            v6 = 1
            if (18446744073709551615 == 0) {
                fmt.Println(v6, p2, p0)
                v7 := GenS1{F0: 57, F1: 1}
                var v8 bool = (!true)
                _ = v7
                _ = v8
            } else {
                v9 := GenS2{F0: v6}
                _ = v9
            }
            _ = v6
        }
        _ = v5
    } else {
        v10 := GenS1{F0: 42, F1: 42}
        var v11 int16 = -840
        _ = v10
        _ = v11
    }
    if false {
        v12 := []int{}
        if true {
            for fi13 := 0; fi13 < 20; fi13++ {
                p0 = p1
                fmt.Println(fi13, p1)
                v14 := GenS3{F0: 0}
                _ = fi13
                _ = v14
            }
        } else {
            fmt.Println(p1, p1, p0)
            var v15 float64 = 1e10
            if true {
                fmt.Println(p2, p0)
                v15 -= float64(1.5)
                v12 = append(v12, -993)
            }
            _ = v15
        }
        _ = v12
    }
    if ((false && (-128 == -1)) && true) {
        v16 := GenS2{F0: 9223372036854775807}
        fmt.Println(p2, p1)
        _ = v16
    } else {
        switch "foo bar" {
        case "\n\t":
            p0 += (p1 >> ((p1) & 63))
        case "":
            p1 = (p0 + (p0 / ((p0) | 1)))
        case "a":
            var v17 int64 = p2
            _ = v17
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 643
}
func genFn1() string {
    switch 127 {
    case -9223372036854775808:
        var v18 string = ""
        _ = v18
    }
    for fi19 := 0; fi19 < 7; fi19++ {
        if false {
            if false {
                var v20 int16 = 0
                _ = v20
            }
        }
        var v21 int64 = 651
        v22 := GenS3{F0: 4294967295}
        _ = fi19
        _ = v21
        _ = v22
    }
    if true {
        var v23 float32 = 1.5
        _ = v23
    }
    return "hello"
}

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]int)
    var v3 int64 = -458
    var v4 uint16 = 1
    var v24 int64 = -144
    delete(v2, 438)
    fmt.Println(v4, v3, v4)
    for fi25 := 0; fi25 < 12; fi25++ {
        v26 := GenS1{F0: 127, F1: -63}
        if ((true || ((-2.25 <= (0.0 - 1e10)) && false)) || (false && false)) {
            v27 := []int{(fi25 + 9223372036854775807), (fi25 + 1), (fi25 * genFn0(1, 0, (v24 + v24))), (fi25 ^ (fi25 / ((127) | 1)))}
            fmt.Println(v3, v24)
            _ = v27
        }
        _ = fi25
        _ = v26
    }
    var v28 string = ("" + ("a" + ""))
    v3 -= v3
    v24 += (v3 >> ((42) & 63))
    v1 = v1
    v29 := v2[42]
    var v30 int64 = v24
    delete(v2, 4294967295)
    v31 := []int{-9223372036854775808, (v29 << ((280) & 63)), v29, 340}
    var v32 bool = (65535 != 2)
    v28 = ((("foo bar" + v28) + v28) + v28)
    v33 := GenS1{F0: 255, F1: -1}
    fmt.Println(v3)
    var v34 uint8 = 128
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
