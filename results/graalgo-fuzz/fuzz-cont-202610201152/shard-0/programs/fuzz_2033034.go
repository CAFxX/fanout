package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint8; F1 int64 }
func genFn0(p0 int16, p1 int32, p2 uint) uint16 {
    p0 += p0
    if true {
        v5 := GenS1{F0: 0, F1: -296, F2: 0}
        v6 := GenS1{F0: (p1 | 1), F1: 32767, F2: int(p1)}
        switch 0 {
        case 671:
            var v7 uint64 = 1
            _ = v7
        case 646:
            if false {
                v8 := []int{32767, 361}
                _ = v8
            } else {
                if true {
                    switch -2147483648 {
                    case 128:
                        var v9 int64 = 9223372036854775807
                        _ = v9
                    }
                    fmt.Println(p2)
                } else {
                    var v10 int64 = -1
                    if true {
                        v11 := GenS1{F0: -786, F1: (p0 ^ p0), F2: 9223372036854775807}
                        _ = v11
                    } else {
                        var v12 float32 = -2.25
                        _ = v12
                    }
                    v10 = (v10 ^ v10)
                    _ = v10
                }
                fmt.Println(p2)
            }
        case 127:
            v13 := GenS1{F0: p1, F1: p0, F2: 42}
            _ = v13
        default:
            v14 := []int{}
            _ = v14
        }
        _ = v5
        _ = v6
    }
    v15 := []int{0, 646, -1, 128}
    p0 = (p0 + p0)
    v16 := GenS1{F0: p1, F1: (p0 % ((1000) | 1)), F2: int(p2)}
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v16
    return 65535
}
func genFn1(p0 int8) int64 {
    v17 := GenS1{F0: -1, F1: -32768, F2: 2147483647}
    var v18 uint64 = 1
    v19 := GenS2{F0: 42, F1: 1, F2: v18}
    fmt.Println(p0, v18)
    for fi20 := 0; fi20 < 1; fi20++ {
        var v21 uint16 = 708
        _ = fi20
        _ = v21
    }
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return 42
}

func main() {
    var v1 int = 4294967295
    var v2 string = ""
    var v3 uint8 = 2
    var v4 bool = false
    v22 := GenS2{F0: 1, F1: 32767, F2: 915}
    v2 = ""
    if ((v1 | v1) != (-(v1))) {
        v1 = v1
        fmt.Println(v1, v4, v3)
    }
    switch v2 {
    case "foo bar":
        if v4 {
            fmt.Println(v1)
            var v23 uint8 = (v3 ^ 2)
            _ = v23
        }
    case "a":
        v24 := []int{(v1 | 128)}
        _ = v24
    }
    var v25 float64 = float64(-2.25)
    switch v2 {
    case "":
        v26 := GenS3{F0: (v3 >> ((387) & 63)), F1: 9223372036854775807}
        _ = v26
    case "\n\t":
        v1 = int(v3)
    default:
        v2 = v2
    }
    v25 -= v25
    var v27 int16 = 1
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(v27, v25)
        _ = fi28
    }
    if false {
        v29 := []int{}
        _ = v29
    }
    var v30 string = ""
    if v4 {
        v31 := GenS2{F0: 42, F1: 404, F2: 1}
        var v32 uint16 = 255
        _ = v31
        _ = v32
    } else {
        switch v2 {
        case "":
            v25 = v25
        }
        fmt.Println(v25)
        for fi33 := 0; fi33 < 20; fi33++ {
            if (int32(fi33) < 452) {
                for fi34 := 0; fi34 < 12; fi34++ {
                    v35 := GenS3{F0: 0, F1: -63}
                    v30 = ("" + v30)
                    var v36 uint64 = 42
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            }
            var v37 uint8 = v3
            _ = fi33
            _ = v37
        }
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int8 = -128
        var v40 uint = 0
        _ = fi38
        _ = v39
        _ = v40
    }
    var v41 uint64 = 4294967295
    var v42 uint16 = 1
    var v43 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v22
    _ = v25
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v41
    _ = v42
    _ = v43
}
