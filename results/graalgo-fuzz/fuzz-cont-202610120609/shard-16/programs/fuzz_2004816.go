package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 bool; F1 uint32 }
type GenS6 struct { F0 bool; F1 uint }
type GenS7 struct { F0 uint32; F1 uint32; F2 uint }

func main() {
    var v1 int16 = -125
    var v2 uint64 = 1
    v3 := []uint32{2}
    var v4 float64 = (((0.0 - ((3.14159 + (-2.25 * 1.5)) - float64((((1e10 * -2.25) * 1e10) * 0.0)))) + 1e10) - 0.0)
    v5 := []int{-1, 1}
    var v6 int16 = v1
    if len(v5) > 0 {
        v5[0] = 32768
    }
    v7 := GenS1{F0: "a"}
    var v8 int = 49
    if true {
        if false {
            fmt.Println(v4, v1, v6)
            _ = len(v3)
        } else {
            v9 := GenS2{F0: -455, F1: -931}
            v10 := GenS3{F0: v1, F1: 0}
            v3 = append(v3, 969)
            _ = v10
            _ = v9
        }
    }
    if ((false || (64 != -12)) || (v4 > v4)) {
        var v11 uint64 = v2
        if false {
            var v12 int = v8
            _ = v12
        }
        v2 = v2
        _ = v11
    } else {
        var v13 float32 = 1.5
        v14 := GenS4{F0: -456}
        _ = v13
        _ = v14
    }
    v15 := GenS5{F0: false, F1: 346}
    v6 -= 161
    v16 := GenS6{F0: (true && false), F1: 65535}
    var v17 uint8 = 231
    if ((!false) && true) {
        switch "foo bar" {
        case "a":
            v4 += v4
        case "\n\t":
            if true {
                if false {
                    if (100 < -128) {
                        v18 := GenS1{F0: "\n\t"}
                        var v19 uint32 = uint32(v17)
                        var v20 uint8 = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                    var v21 int32 = -1
                    _ = v21
                }
                var v22 bool = true
                var v23 uint32 = 42
                _ = v22
                _ = v23
            } else {
                switch "\n\t" {
                case "\n\t":
                    if (((int(v2) <= v8) || (("" + "\n\t") <= "a")) || false) {
                        _ = len(v3)
                    }
                case "a":
                    var v24 bool = ((77 == 2) || false)
                    _ = v24
                default:
                    var v25 uint16 = 32767
                    _ = v25
                }
            }
        default:
            fmt.Println(v1)
        }
    }
    switch 32768 {
    case -147:
        for fi26 := 0; fi26 < 17; fi26++ {
            v27 := GenS6{F0: (-98 > 42), F1: 356}
            _ = fi26
            _ = v27
        }
    case -2147483648:
        var v28 bool = (0 >= v17)
        _ = v28
    case 4294967295:
        if true {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            if (32767 < 2) {
                var v29 int32 = 0
                v5 = append(v5, (v8 + v8))
                _ = v29
            }
            v30 := GenS3{F0: 608, F1: (v17 | (v17 + v17))}
            _ = v30
        }
    default:
        switch (v8 / (((^(v8))) | 1)) {
        case 32767:
            if (!(false && ((v1 <= -903) && false))) {
                var v31 uint16 = 2
                _ = v31
            }
        default:
            fmt.Println(v2, v4, v17)
        }
    }
    for fi32 := 0; fi32 < 12; fi32++ {
        v2 = v2
        _ = fi32
    }
    v33 := GenS1{F0: "hello"}
    v34 := GenS7{F0: 959, F1: 42, F2: 997}
    switch (v8 + v8) {
    case 2147483647:
        var v35 int = 929
        _ = v35
    case 9223372036854775807:
        var v36 int64 = -9223372036854775808
        _ = v36
    case 80:
        v6 = v1
    default:
        v2 = 4294967295
    }
    var v37 uint32 = 2147483647
    var v38 uint32 = v37
    var v39 uint32 = v37
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
