package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int8; F1 uint8; F2 int64 }
type GenS3 struct { F0 uint8; F1 uint16 }
type GenS4 struct { F0 int32; F1 uint16 }
type GenS5 struct { F0 string }
func genFn0(p0 int32, p1 int32) int {
    switch -435 {
    case 747:
        var v6 uint32 = 0
        _ = v6
    }
    switch "\n\t" {
    case "foo bar":
        v7 := GenS1{F0: -1, F1: 32767}
        _ = v7
    }
    _ = p0
    _ = p1
    return 32768
}

func main() {
    var v1 int64 = 1
    var v2 float32 = ((1.5 + 0.0) * 0.0)
    var v3 uint = 255
    var v4 uint16 = 42
    var v5 float32 = v2
    var v8 string = "foo bar"
    for fi9 := 0; fi9 < 9; fi9++ {
        v5 += v2
        fmt.Println(v3)
        var v10 int32 = 2147483647
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        switch (fi11 & fi11) {
        case -59:
            v12 := []int{(fi11 << ((v3) & 63))}
            _ = v12
        default:
            v8 = v8
        }
        for fi13 := 0; fi13 < 3; fi13++ {
            v3 = v3
            if true {
                v14 := GenS1{F0: 1, F1: v4}
                _ = v14
            }
            switch v8 {
            case "":
                var v15 uint = 4294967295
                _ = v15
            case "\n\t":
                v16 := GenS1{F0: 0, F1: (v4 + (^(v4)))}
                _ = v16
            }
            _ = fi13
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if false {
            for fi18 := 0; fi18 < 2; fi18++ {
                if (v1 == v1) {
                    switch 0 {
                    case 0:
                        fmt.Println(v5, v8)
                    default:
                        v19 := GenS2{F0: 100, F1: 127, F2: v1}
                        _ = v19
                    }
                }
                _ = fi18
            }
            v20 := GenS3{F0: 255, F1: (v4 + (v4 | v4))}
            var v21 int16 = 818
            _ = v20
            _ = v21
        }
        _ = fi17
    }
    if true {
        v3 = v3
        if (2147483647 < -1) {
            v22 := []int{127, genFn0(42, 1), 2147483647, 9223372036854775807}
            _ = v22
        }
    } else {
        v3 = v3
        v23 := GenS4{F0: 0, F1: (v4 % ((v4) | 1))}
        _ = v23
    }
    v24 := []int{int(v4), 127, 1, 9223372036854775807}
    fmt.Println(v8, v4, v5)
    v25 := GenS3{F0: 255, F1: 669}
    v26 := GenS1{F0: -22, F1: v4}
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    if len(v24) >= 2 {
        v24 = v24[1:len(v24)]
    }
    _ = len(v24)
    v8 = (v8 + "")
    var v27 uint64 = 42
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(v3)
        v29 := GenS5{F0: v8}
        _ = fi28
        _ = v29
    }
    if true {
        var v30 int8 = 42
        _ = v30
    } else {
        for fi31 := 0; fi31 < 11; fi31++ {
            fmt.Println(fi31)
            v27 = 1
            if false {
                if true {
                    if (v4 > v4) {
                        v1 = v1
                        v3 -= (v3 >> ((v3) & 63))
                    }
                } else {
                    v32 := GenS5{F0: v8}
                    _ = v32
                }
                fmt.Println(fi31, v3)
            }
            _ = fi31
        }
        var v33 uint8 = 128
        v3 -= v3
        _ = v33
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if ((!true) && true) {
            var v35 uint16 = 2
            _ = v35
        }
        v8 = ("hello" + "foo bar")
        _ = fi34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
