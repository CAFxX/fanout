package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 uint8; F1 uint16 }
func genFn0(p0 string, p1 bool, p2 bool) int16 {
    switch p0 {
    case "hello":
        switch -2147483648 {
        case 762:
            var v7 string = "hello"
            _ = v7
        default:
            fmt.Println(p2, p2, p0)
        }
    case "":
        var v8 float32 = float32(1.5)
        _ = v8
    default:
        v9 := []int{0, 0}
        _ = v9
    }
    if ((((1.5 - 0.1) / -2.25) < 0.1) && (p0 == (p0 + p0))) {
        fmt.Println(p0, p1)
    }
    if p2 {
        v10 := GenS1{F0: (0 >= 208)}
        fmt.Println(p0, p1)
        switch -337 {
        case -1:
            var v11 float32 = 0.1
            _ = v11
        }
        _ = v10
    }
    fmt.Println(p1)
    if true {
        var v12 float32 = 0.0
        v13 := GenS1{F0: p1}
        fmt.Println(p0, p2, v12)
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}
func genFn1() uint8 {
    switch 32768 {
    case 1:
        var v14 uint16 = 39
        _ = v14
    case 42:
        v15 := GenS1{F0: (false && true)}
        _ = v15
    case 2147483647:
        var v16 uint16 = 0
        _ = v16
    }
    var v17 int8 = -30
    v17 -= (v17 >> ((0) & 63))
    if (42 <= -32768) {
        v17 += v17
    }
    _ = v17
    return 42
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]int)
    v3 := make(map[int]uint32)
    v4 := make(map[string]string)
    var v5 int64 = 9223372036854775807
    var v6 float32 = float32(1.5)
    for fi18 := 0; fi18 < 20; fi18++ {
        for fi19 := 0; fi19 < 9; fi19++ {
            v20 := GenS1{F0: (true || false)}
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    switch 9223372036854775807 {
    case -775:
        v21 := GenS1{F0: (("a" + "a") != "hello")}
        _ = v21
    case 280:
        var v22 int = -9223372036854775808
        _ = v22
    default:
        var v23 int64 = (v5 & (^(v5)))
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 uint32 = 1
        _ = fi24
        _ = v25
    }
    for fi26 := 0; fi26 < 5; fi26++ {
        v1 -= (v1 + (^(v1)))
        if (-909 >= (fi26 / (((^(fi26))) | 1))) {
            for fi27 := 0; fi27 < 20; fi27++ {
                var v28 uint8 = 215
                v1 = (v1 << ((248) & 63))
                _ = fi27
                _ = v28
            }
            for fi29 := 0; fi29 < 14; fi29++ {
                v6 -= v6
                _ = fi29
            }
            fmt.Println(v6)
        }
        _ = fi26
    }
    var v30 uint16 = 63
    v31 := GenS2{F0: (v5 * v5)}
    v32 := []int{0, 128, -764}
    v4["hello"] = "a"
    v6 = v6
    switch 32767 {
    case 2147483647:
        v33 := GenS3{F0: -592, F1: v1}
        _ = v33
    default:
        var v34 int32 = -972
        _ = v34
    }
    for fi35 := 0; fi35 < 17; fi35++ {
        v36 := GenS4{F0: 128, F1: (v30 | 252)}
        if true {
            if false {
                var v37 string = "a"
                _ = v37
            }
        } else {
            v30 = (v30 % (((v30 | 46)) | 1))
            var v38 uint8 = 128
            var v39 uint16 = 255
            _ = v38
            _ = v39
        }
        _ = fi35
        _ = v36
    }
    fmt.Println(v30, v5)
    switch ("\n\t" + "") {
    case "\n\t":
        var v40 uint32 = (v1 + (v1 + v1))
        _ = v40
    case "hello":
        fmt.Println(v1, v5)
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
