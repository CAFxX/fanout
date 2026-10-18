package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 int8 }
type GenS4 struct { F0 uint8; F1 uint }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int64; F1 uint64 }

func main() {
    var v1 float32 = 1.5
    var v2 bool = (!((false && (!false)) || (!true)))
    v3 := make(map[int]int)
    var v4 int8 = 98
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: v2}
        for fi7 := 0; fi7 < 11; fi7++ {
            if v2 {
                fmt.Println(v2, v4, v4)
                if v2 {
                    v8 := GenS1{F0: ((264 > -6) && v2)}
                    var v9 uint8 = 42
                    v2 = v2
                    _ = v8
                    _ = v9
                } else {
                    v2 = (v2 && (32767 == 93))
                }
                fmt.Println(fi5, fi5)
            } else {
                fmt.Println(v1, fi7, fi7)
                v10 := GenS2{F0: 0, F1: 18446744073709551615}
                _ = v10
            }
            _ = fi7
        }
        _ = fi5
        _ = v6
    }
    if ("" <= "hello") {
        v1 = 0.1
        var v11 int8 = v4
        _ = v11
    }
    if true {
        var v12 int32 = -1
        var v13 float32 = ((1e10 - v1) + v1)
        fmt.Println(v13)
        _ = v12
        _ = v13
    } else {
        fmt.Println(v4)
        switch "a" {
        case "":
            switch -52 {
            case -1:
                fmt.Println(v2)
            }
        }
        fmt.Println(v2, v2)
    }
    var v14 uint64 = 76
    v1 += v1
    if (!v2) {
        v15 := GenS3{F0: 61, F1: 436, F2: v4}
        if (0 >= 38) {
            var v16 float32 = v1
            fmt.Println(v2, v16)
            v17 := GenS1{F0: false}
            _ = v16
            _ = v17
        }
        _ = v15
    } else {
        var v18 uint16 = 511
        var v19 uint16 = 255
        var v20 float32 = 0.0
        _ = v18
        _ = v19
        _ = v20
    }
    if (0.0 < -2.25) {
        var v21 uint = 896
        delete(v3, 828)
        if false {
            v22 := GenS4{F0: 1, F1: 253}
            _ = v22
        } else {
            var v23 uint32 = 1
            _ = v23
        }
        _ = v21
    }
    v24 := GenS4{F0: 2, F1: 887}
    if (v4 > v4) {
        fmt.Println(v1)
        v1 += ((0.1 / float32(float64(v1))) - 1.5)
        var v25 uint = 0
        _ = v25
    } else {
        var v26 uint8 = 0
        for fi27 := 0; fi27 < 6; fi27++ {
            switch "a" {
            case "hello":
                fmt.Println(v4, v14)
            case "a":
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 float32 = v1
                    _ = fi28
                    _ = v29
                }
            case "foo bar":
                fmt.Println(v14)
            default:
                fmt.Println(v4)
            }
            var v30 uint = 42
            _ = fi27
            _ = v30
        }
        _ = v26
    }
    v31 := GenS4{F0: 127, F1: 65535}
    fmt.Println(v4, v14, v4)
    v32 := GenS4{F0: 53, F1: 18446744073709551615}
    v14 += (v14 * v14)
    v33 := GenS5{F0: 42}
    for fi34 := 0; fi34 < 14; fi34++ {
        fmt.Println(v4, v14)
        v14 = 354
        v35 := GenS6{F0: -1, F1: 2}
        _ = fi34
        _ = v35
    }
    var v36 float32 = v1
    var v37 bool = ((1 <= 709) && false)
    v3[930] = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v4
}
