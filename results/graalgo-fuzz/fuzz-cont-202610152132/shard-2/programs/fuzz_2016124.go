package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int32; F1 uint }
func genFn0(p0 uint16, p1 int64, p2 string) int64 {
    p2 = (((p2 + (p2 + p2)) + "foo bar") + ((p2 + ("" + "a")) + (p2 + "\n\t")))
    switch p2 {
    case "a":
        var v6 uint64 = 42
        _ = v6
    case "":
        fmt.Println(p0, p0)
    case "\n\t":
        p1 += p1
    }
    p2 = p2
    var v7 string = "hello"
    v8 := []int{969, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}

func main() {
    var v1 string = ("a" + "")
    var v2 uint64 = 42
    v3 := make(map[uint32]bool)
    var v4 uint8 = 162
    v5 := []string{""}
    var v9 uint64 = v2
    v2 = v2
    var v10 string = v1
    v1 = ("foo bar" + v1)
    for fi11 := 0; fi11 < 6; fi11++ {
        if len(v5) > 0 {
            v5[0] = v1
        }
        var v12 uint16 = 2
        var v13 string = v1
        _ = fi11
        _ = v12
        _ = v13
    }
    v14 := v3[65535]
    var v15 uint64 = v2
    v16 := GenS1{F0: 164, F1: 35}
    if v14 {
        switch 4294967295 {
        case 42:
            var v17 uint64 = (v2 ^ 723)
            _ = v17
        case 470:
            fmt.Println(v10)
        default:
            switch (("\n\t" + v10) + v10) {
            case "a":
                v4 -= 42
            case "foo bar":
                fmt.Println(v14)
            case "\n\t":
                _ = len(v5)
            }
        }
        fmt.Println(v1, v2)
        v14 = true
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        v19 := GenS1{F0: (v4 << ((1) & 63)), F1: (v4 * (v4 / ((v4) | 1)))}
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint64 = (v9 >> ((65535) & 63))
            var v22 bool = (false || ((1.5 * (0.0 * (1e10 + (((0.0 / float32((1.5 - 3.14159))) * 1.5) + (0.0 / 1.0))))) != 0.1))
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi18
        _ = v19
    }
    v23 := GenS1{F0: v4, F1: 61}
    if (1e10 < 3.14159) {
        fmt.Println(v4, v2, v10)
        v1 = v10
    } else {
        v24 := GenS1{F0: v4, F1: 42}
        if true {
            v9 = 0
        }
        for fi25 := 0; fi25 < 18; fi25++ {
            for fi26 := 0; fi26 < 9; fi26++ {
                v4 += v4
                if v14 {
                    fmt.Println(v10, v15)
                }
                var v27 bool = ((32767 < 0) && (true || v14))
                _ = fi26
                _ = v27
            }
            var v28 float32 = ((-2.25 - 3.14159) * float32(1.5))
            _ = fi25
            _ = v28
        }
        _ = v24
    }
    v29 := v3[65535]
    for fi30 := 0; fi30 < 13; fi30++ {
        v10 = (v10 + "hello")
        switch "foo bar" {
        case "foo bar":
            _ = len(v5)
        }
        _ = fi30
    }
    switch (v1 + (v1 + (v1 + v1))) {
    case "hello":
        var v31 float64 = 1e10
        _ = v31
    case "\n\t":
        var v32 uint32 = 0
        _ = v32
    case "a":
        var v33 int64 = 9223372036854775807
        _ = v33
    default:
        fmt.Println(v9, v4)
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        if (2147483647 < 1) {
            v35 := GenS2{F0: v15, F1: (v4 * 128), F2: 2147483648}
            _ = v35
        } else {
            switch (fi34 | (fi34 % (((fi34 ^ fi34)) | 1))) {
            case -2147483648:
                v36 := GenS2{F0: v2, F1: (v4 << ((255) & 63)), F2: uint32(v4)}
                _ = v36
            case -405:
                var v37 string = ""
                _ = v37
            case 42:
                if true {
                    var v38 int16 = 1
                    v39 := GenS3{F0: 0}
                    _ = v38
                    _ = v39
                }
            default:
                v40 := GenS3{F0: -1}
                _ = v40
            }
            v41 := GenS4{F0: 2147483647}
            _ = v41
        }
        _ = len(v5)
        _ = fi34
    }
    if true {
        switch -964 {
        case 42:
            var v42 uint8 = 130
            _ = v42
        default:
            v2 -= (v2 + 531)
        }
        v43 := GenS5{F0: 966}
        _ = v43
    }
    v44 := GenS6{F0: -463, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v23
    _ = v29
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v9
}
