package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 uint; F1 bool; F2 uint8 }
type GenS3 struct { F0 bool; F1 uint64; F2 int8 }
func genFn0(p0 uint16) bool {
    if false {
        fmt.Println(p0)
        var v5 int64 = int64(p0)
        fmt.Println(v5)
        _ = v5
    } else {
        var v6 float32 = (-2.25 - (-2.25 / 1.5))
        v7 := []int{}
        _ = v6
        _ = v7
    }
    p0 = (p0 | 65535)
    if false {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int = (fi8 << ((255) & 63))
            fmt.Println(fi8, v9, p0)
            _ = fi8
            _ = v9
        }
        p0 = (^(p0))
        v10 := []int{-254}
        _ = v10
    } else {
        p0 = p0
        fmt.Println(p0)
        var v11 uint32 = 2147483647
        _ = v11
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 8; fi13++ {
            p0 -= p0
            fmt.Println(p0, p0)
            _ = fi13
        }
        var v14 int64 = 42
        _ = fi12
        _ = v14
    }
    _ = p0
    return true
}

func main() {
    var v1 int = 1
    var v2 float32 = 0.1
    var v3 float64 = ((0.0 - 1.5) - (-2.25 + float64((v2 + ((v2 * float32((0.1 - ((((1e10 * 0.0) - -2.25) - 0.1) * 3.14159)))) + v2)))))
    v4 := make(map[int]string)
    var v15 float64 = 1e10
    v16 := []int{v1, v1}
    v17 := GenS1{F0: false, F1: (false && ((((137 > 871) || true) && false) || (-1 >= -9223372036854775808)))}
    for fi18 := 0; fi18 < 2; fi18++ {
        if (-32768 == -32768) {
            for fi19 := 0; fi19 < 19; fi19++ {
                var v20 int = fi18
                _ = fi19
                _ = v20
            }
            v3 = float64(v2)
        }
        var v21 int8 = 42
        v1 += 9223372036854775807
        _ = fi18
        _ = v21
    }
    v1 += v1
    _ = len(v4)
    v2 = v2
    switch -894 {
    case 127:
        v3 += float64(0.1)
    case 2147483647:
        v2 = v2
    case 294:
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
    default:
        for fi22 := 0; fi22 < 16; fi22++ {
            v23 := GenS2{F0: 0, F1: (654 >= 1), F2: 255}
            _ = fi22
            _ = v23
        }
    }
    fmt.Println(v3, v3)
    v24 := GenS3{F0: true, F1: 42, F2: -128}
    if (false && (!((true || false) && true))) {
        var v25 int64 = 9223372036854775807
        _ = v25
    }
    switch (("" + "") + (("" + "\n\t") + ("\n\t" + ("a" + "")))) {
    case "a":
        if false {
            fmt.Println(v15, v15)
            v1 = (-(v1))
        } else {
            var v26 uint16 = 42
            v16 = append(v16, 128)
            v3 = v15
            _ = v26
        }
    case "":
        switch ("" + (("\n\t" + "a") + "hello")) {
        case "foo bar":
            for fi27 := 0; fi27 < 5; fi27++ {
                v3 = ((((v15 + v3) - (float64(0.1) - -2.25)) - v3) / v3)
                _ = fi27
            }
        }
    default:
        for fi28 := 0; fi28 < 13; fi28++ {
            v29 := GenS1{F0: false, F1: true}
            v1 = (^(v1))
            _ = fi28
            _ = v29
        }
    }
    var v30 float32 = 1.5
    if false {
        v15 += ((v15 * v3) * v3)
        v30 += (v2 * v30)
    } else {
        v15 = (float64(v2) * 3.14159)
        v31 := GenS1{F0: (1 < 1), F1: true}
        v32 := GenS3{F0: (0 <= 2), F1: 1, F2: 100}
        _ = v31
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        switch "hello" {
        case "a":
            if true {
                v3 += v15
            }
        }
        fmt.Println(v30, fi33, v1)
        _ = fi33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v4
}
