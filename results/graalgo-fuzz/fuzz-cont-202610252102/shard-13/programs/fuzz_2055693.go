package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 uint64; F1 int64; F2 uint }

func main() {
    var v1 uint32 = 2
    var v2 float32 = 1e10
    var v3 int16 = 0
    v3 = (v3 * v3)
    for fi4 := 0; fi4 < 3; fi4++ {
        var v5 float32 = ((v2 / float32(3.14159)) + v2)
        if false {
            v6 := GenS1{F0: 0, F1: 1}
            _ = v6
        } else {
            fmt.Println(v3, v3)
            v2 = float32((1e10 - 1.5))
            v3 = (v3 | (v3 + (v3 - 1000)))
        }
        if false {
            var v7 int = 127
            switch (("" + "hello") + "\n\t") {
            case "hello":
                v2 -= v2
            default:
                var v8 uint8 = 1
                _ = v8
            }
            if false {
                v9 := []int{(^(fi4)), (fi4 * (v7 - (fi4 - (v7 & (^(v7))))))}
                switch "" {
                case "":
                    v2 = float32(0.1)
                case "a":
                    v10 := GenS1{F0: 0, F1: 1}
                    _ = v10
                default:
                    v11 := GenS1{F0: 42, F1: 42}
                    _ = v11
                }
                _ = v9
            }
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    v2 = v2
    v12 := []int{}
    v1 = 507
    if (255 >= 769) {
        var v13 int32 = -2147483648
        _ = v13
    } else {
        fmt.Println(v3)
        if (-2147483648 < 478) {
            var v14 int8 = 1
            _ = v14
        } else {
            if (0 < 2) {
                v15 := GenS2{F0: 32767, F1: 1}
                _ = v15
            } else {
                var v16 uint = 625
                fmt.Println(v16)
                _ = v16
            }
            v2 += v2
            if len(v12) > 0 {
                v12[0] = 128
            }
        }
    }
    var v17 bool = true
    v18 := GenS2{F0: 2, F1: 65535}
    for fi19 := 0; fi19 < 9; fi19++ {
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 int8 = 0
            var v22 int = fi19
            _ = fi20
            _ = v21
            _ = v22
        }
        if false {
            v2 -= ((3.14159 * 0.0) - v2)
            if (201 == 109) {
                var v23 uint64 = 0
                _ = len(v12)
                _ = v23
            }
            v3 -= v3
        } else {
            var v24 int32 = -2147483648
            var v25 float32 = v2
            v26 := GenS1{F0: -9223372036854775808, F1: 51}
            _ = v24
            _ = v25
            _ = v26
        }
        switch "foo bar" {
        case "a":
            var v27 uint = 19
            _ = v27
        case "foo bar":
            v28 := GenS3{F0: 0, F1: 629}
            _ = v28
        case "\n\t":
            fmt.Println(v2, v2, fi19)
        }
        _ = fi19
    }
    v29 := GenS4{F0: 0, F1: -9223372036854775808, F2: 255}
    var v30 int8 = 127
    switch (("hello" + "\n\t") + ("" + "a")) {
    case "a":
        if len(v12) > 0 {
            v12[0] = 32768
        }
    }
    var v31 int8 = -1
    v1 += (v1 | v1)
    if (!v17) {
        if ((((v2 + (1.5 / v2)) + (v2 + (v2 + ((v2 / v2) - 0.0)))) / v2) < -2.25) {
            v17 = ((float32((-2.25 / -2.25)) * ((0.0 / float32(((0.0 * 3.14159) - 0.1))) / (v2 - float32(1e10)))) == (3.14159 - v2))
            fmt.Println(v1, v3, v2)
        }
    }
    v32 := GenS4{F0: 311, F1: -9223372036854775808, F2: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    _ = v1
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
}
