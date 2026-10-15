package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 bool; F2 uint16 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 uint8; F1 bool }
type GenS6 struct { F0 int32; F1 uint64; F2 string }
func genFn0() uint64 {
    var v4 int64 = -577
    switch "a" {
    case "foo bar":
        v5 := []int{-590, 127, -9223372036854775808, 14}
        _ = v5
    case "hello":
        if true {
            var v6 uint64 = 42
            v7 := GenS1{F0: v6}
            _ = v6
            _ = v7
        }
    case "a":
        var v8 uint64 = 18446744073709551615
        _ = v8
    default:
        fmt.Println(v4)
    }
    _ = v4
    return 18446744073709551615
}

func main() {
    var v1 uint = 2
    var v2 int = 9223372036854775807
    v3 := []int{}
    v2 = v2
    _ = len(v3)
    switch ("foo bar" + "foo bar") {
    case "foo bar":
        if (false || true) {
            var v9 uint16 = 65535
            switch ("a" + (("hello" + "hello") + "\n\t")) {
            case "\n\t":
                v2 += (v2 & (v2 / (((v2 / ((v2) | 1))) | 1)))
            case "":
                for fi10 := 0; fi10 < 4; fi10++ {
                    var v11 int = v2
                    var v12 int16 = 0
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "a":
                v13 := GenS1{F0: 2}
                _ = v13
            default:
                var v14 uint16 = (v9 | 965)
                _ = v14
            }
            v15 := GenS2{F0: (v2 * v2), F1: (274 != 793), F2: (v9 - v9)}
            _ = v15
            _ = v9
        } else {
            v16 := GenS3{F0: 0, F1: v1}
            if false {
                var v17 int = v2
                v18 := GenS3{F0: 1000, F1: (v1 % (((v1 / (((v1 & 0)) | 1))) | 1))}
                _ = v17
                _ = v18
            }
            _ = v16
        }
    case "a":
        if (((0.1 * 0.0) > 0.1) || ((true && ((v1 + v1) <= v1)) || true)) {
            v2 += 213
            if false {
                v1 -= v1
                for fi19 := 0; fi19 < 8; fi19++ {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        v1 -= (v1 >> ((v1) & 63))
                        var v21 int32 = -2147483648
                        _ = fi20
                        _ = v21
                    }
                    var v22 uint8 = 1
                    v23 := GenS3{F0: -593, F1: v1}
                    _ = fi19
                    _ = v22
                    _ = v23
                }
                if ((!(-2147483648 < -1)) || false) {
                    v1 -= (v1 * 255)
                    var v24 uint = v1
                    _ = v24
                }
            }
        } else {
            v25 := GenS4{F0: 255, F1: (v1 ^ v1)}
            v2 -= v2
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 string = "foo bar"
                _ = fi26
                _ = v27
            }
            _ = v25
        }
    default:
        v28 := GenS5{F0: 2, F1: true}
        _ = v28
    }
    var v29 uint64 = 18446744073709551615
    if false {
        switch (v2 + v2) {
        case 128:
            if (v29 != (v29 << (((^(v1))) & 63))) {
                if true {
                    switch ("hello" + "\n\t") {
                    case "":
                        v1 = 42
                    }
                }
            }
        case 2147483647:
            v30 := GenS1{F0: v29}
            _ = v30
        case 289:
            switch (v2 + v2) {
            case -1:
                var v31 uint8 = 127
                _ = v31
            }
        default:
            fmt.Println(v1, v1, v2)
        }
        var v32 uint16 = 587
        _ = v32
    }
    var v33 uint = v1
    if (!(42 < 2)) {
        for fi34 := 0; fi34 < 15; fi34++ {
            v29 = v29
            _ = fi34
        }
        _ = len(v3)
    } else {
        var v35 bool = (197 <= 42)
        v36 := GenS1{F0: v29}
        _ = v35
        _ = v36
    }
    v2 -= v2
    for fi37 := 0; fi37 < 3; fi37++ {
        v38 := GenS3{F0: 1, F1: (v1 - (v1 | v1))}
        v29 = uint64(v1)
        _ = fi37
        _ = v38
    }
    if false {
        _ = len(v3)
        var v39 int16 = 42
        if len(v3) > 0 {
            v3[0] = 1
        }
        _ = v39
    } else {
        var v40 uint16 = uint16(v33)
        v41 := GenS6{F0: -487, F1: (v29 * v29), F2: "\n\t"}
        var v42 float64 = (1.5 - float64((1.5 - 3.14159)))
        _ = v40
        _ = v41
        _ = v42
    }
    fmt.Println(v33)
    var v43 int = 268
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v33
    _ = v43
}
