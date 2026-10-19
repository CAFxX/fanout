package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint32; F1 uint; F2 string }
type GenS3 struct { F0 int16; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint16; F1 bool; F2 int64 }
type GenS5 struct { F0 int; F1 int16 }
func genFn0(p0 uint8, p1 int32, p2 int32) uint32 {
    fmt.Println(p1, p0, p2)
    var v5 int16 = -313
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 65535
}
func genFn1(p0 int, p1 int16) int16 {
    p0 -= p0
    for fi6 := 0; fi6 < 10; fi6++ {
        v7 := []int{(p0 | 127), (p0 | -746), -2147483648}
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := GenS1{F0: 2147483647, F1: 959}
            var v10 uint64 = 358
            p0 = p0
            _ = fi8
            _ = v10
            _ = v9
        }
        v11 := GenS2{F0: 442, F1: 255, F2: "foo bar"}
        _ = fi6
        _ = v11
        _ = v7
    }
    _ = p0
    _ = p1
    return p1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ""
    var v3 uint16 = 413
    var v4 uint32 = 468
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 12; fi13++ {
            v1[genFn0(128, 74, 0)] = (!(865 <= 863))
            var v14 int16 = 24
            _ = fi13
            _ = v14
        }
        _ = fi12
    }
    v2 = v2
    var v15 int64 = -1
    var v16 string = "\n\t"
    fmt.Println(v2, v4)
    for fi17 := 0; fi17 < 11; fi17++ {
        fmt.Println(v4)
        for fi18 := 0; fi18 < 19; fi18++ {
            v19 := GenS3{F0: -1, F1: (v3 - v3), F2: 0}
            var v20 bool = (v4 > 4294967295)
            if (fi17 > 650) {
                v21 := []int{}
                fmt.Println(v15, v15, v2)
                var v22 int64 = (-(v15))
                _ = v21
                _ = v22
            } else {
                v3 -= 2
            }
            _ = fi18
            _ = v19
            _ = v20
        }
        _ = fi17
    }
    v23 := GenS4{F0: 935, F1: true, F2: (v15 << ((2) & 63))}
    switch v2 {
    case "\n\t":
        v24 := GenS2{F0: (v4 / ((genFn0(127, 2147483647, 2147483647)) | 1)), F1: 65535, F2: (("a" + (v2 + v2)) + v2)}
        _ = v24
    case "":
        var v25 float32 = 0.0
        _ = v25
    }
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v3)
        v2 = ((((v2 + "hello") + (v2 + v16)) + "") + "foo bar")
        if false {
            v27 := []int{}
            _ = v27
        } else {
            v28 := GenS5{F0: (fi26 << ((2) & 63)), F1: 42}
            _ = v28
        }
        _ = fi26
    }
    if true {
        var v29 int16 = 18
        if true {
            if true {
                v30 := v1[80]
                _ = v30
            } else {
                var v31 float32 = -2.25
                _ = v31
            }
        }
        var v32 uint8 = 128
        _ = v29
        _ = v32
    }
    for fi33 := 0; fi33 < 2; fi33++ {
        if false {
            fmt.Println(v4)
        } else {
            fmt.Println(v15, v4, v16)
            fmt.Println(fi33, v2, fi33)
        }
        _ = fi33
    }
    var v34 uint = 832
    _ = len(v1)
    var v35 uint32 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v34
    _ = v35
    _ = v4
}
