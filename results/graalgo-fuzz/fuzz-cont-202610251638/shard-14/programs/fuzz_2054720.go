package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 int8 }
type GenS4 struct { F0 int; F1 uint64; F2 int16 }
func genFn0(p0 uint32, p1 uint32) uint8 {
    if true {
        p0 = (^(p0))
    } else {
        for fi4 := 0; fi4 < 12; fi4++ {
            if (!true) {
                switch (^(fi4)) {
                case 42:
                    var v5 int = fi4
                    _ = v5
                case 996:
                    p1 += 65535
                }
                p0 -= 251
            } else {
                v6 := []int{(fi4 >> ((2) & 63)), (-(fi4)), fi4}
                _ = v6
            }
            _ = fi4
        }
    }
    var v7 float32 = (((0.1 + 1.5) + 0.1) * (((0.1 / (-2.25 - float32(0.1))) * 0.0) + 0.0))
    switch (("foo bar" + "a") + (("foo bar" + "\n\t") + "foo bar")) {
    case "\n\t":
        var v8 int32 = 42
        _ = v8
    case "a":
        if (("\n\t" + "foo bar") <= (("hello" + ("foo bar" + "a")) + "hello")) {
            if (!false) {
                if true {
                    v9 := GenS1{F0: 433, F1: 1}
                    _ = v9
                }
            }
            v10 := GenS1{F0: 255, F1: 0}
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 uint32 = (p0 - p1)
        v13 := GenS1{F0: 419, F1: -937}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 int = -1
    _ = p0
    _ = p1
    _ = v14
    _ = v7
    return 42
}

func main() {
    var v1 uint32 = 42
    v2 := make(map[uint32]string)
    var v3 string = ""
    var v15 int8 = -128
    if true {
        v16 := GenS2{F0: 152}
        _ = v16
    }
    for fi17 := 0; fi17 < 6; fi17++ {
        var v18 int32 = int32(fi17)
        if false {
            v3 = ("" + (v3 + v3))
            v15 = (^(v15))
            v19 := GenS3{F0: 227, F1: 118, F2: v15}
            _ = v19
        }
        _ = fi17
        _ = v18
    }
    v20 := []int{-2147483648, 529, -345}
    if false {
        var v21 int64 = 1
        _ = v21
    }
    v20 = append(v20, -2147483648)
    if true {
        v22 := GenS1{F0: 855, F1: 925}
        var v23 int8 = 1
        if (18446744073709551615 >= 18446744073709551615) {
            v15 += (v15 % ((v15) | 1))
        } else {
            v2[v1] = (v3 + v3)
            v24 := GenS2{F0: 127}
            if (3.14159 >= 1e10) {
                var v25 float32 = ((0.1 / 1.0) / 3.14159)
                _ = v25
            }
            _ = v24
        }
        _ = v22
        _ = v23
    } else {
        v15 = 1
        v3 = (v3 + "a")
        fmt.Println(v1)
    }
    _ = len(v2)
    var v26 uint64 = 4294967295
    for fi27 := 0; fi27 < 11; fi27++ {
        v3 = v3
        _ = fi27
    }
    for fi28 := 0; fi28 < 18; fi28++ {
        var v29 uint16 = 65535
        v1 += 735
        v30 := GenS3{F0: 0, F1: 42, F2: (v15 * 116)}
        _ = fi28
        _ = v29
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        var v32 uint32 = v1
        v33 := GenS4{F0: 4294967295, F1: v26, F2: 1000}
        var v34 uint64 = 42
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    fmt.Println(v3, v26)
    var v35 string = v3
    switch 128 {
    case -336:
        v15 = 53
    case 625:
        var v36 uint = 4294967295
        _ = v36
    }
    v35 = v35
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v35
}
