package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 string }
func genFn0(p0 string, p1 uint32, p2 uint32) int16 {
    p0 = p0
    var v6 uint32 = 2
    var v7 int64 = 644
    p2 -= 42
    var v8 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return 726
}
func genFn1(p0 uint16) bool {
    var v9 int16 = 32767
    p0 -= 0
    v10 := GenS1{F0: 1}
    p0 -= 0
    p0 = 1
    _ = p0
    _ = v10
    _ = v9
    return true
}

func main() {
    v1 := []uint32{1}
    var v2 uint32 = 2147483648
    var v3 uint16 = 1
    var v4 float64 = -2.25
    v5 := make(map[uint32]int)
    v3 = 65535
    if false {
        var v11 float32 = 1.5
        var v12 float64 = (1.5 - ((v4 / -2.25) / (0.0 + v4)))
        switch "foo bar" {
        case "hello":
            _ = len(v1)
        case "\n\t":
            var v13 uint32 = (v2 - v2)
            _ = v13
        default:
            var v14 uint = 363
            _ = v14
        }
        _ = v11
        _ = v12
    }
    var v15 uint = 695
    if ((true || (!true)) && (1.5 <= ((0.1 * 1.5) + (3.14159 / ((float32(v4) * ((0.0 - (0.1 + (0.1 * 0.0))) * (((0.1 / 3.14159) * -2.25) + 1e10))) * (3.14159 / 3.14159)))))) {
        v16 := GenS2{F0: (((((("foo bar" + ("" + "a")) + "a") + ("a" + "")) + "a") + "foo bar") + "a")}
        if (1 >= 127) {
            var v17 float64 = (v4 + 3.14159)
            fmt.Println(v4, v17, v17)
            v4 -= ((v4 / 1.0) + ((v4 / v17) - v17))
            _ = v17
        }
        fmt.Println(v2, v3)
        _ = v16
    }
    v15 = uint(v3)
    if (18446744073709551615 < 2) {
        for fi18 := 0; fi18 < 10; fi18++ {
            switch (fi18 % ((636) | 1)) {
            case 128:
                v4 = (1e10 * (v4 + v4))
            case 4294967295:
                v19 := GenS3{F0: -746}
                _ = v19
            }
            fmt.Println(v15, v2, v4)
            _ = fi18
        }
    }
    v15 = 1
    for fi20 := 0; fi20 < 3; fi20++ {
        switch ("" + "\n\t") {
        case "hello":
            var v21 uint64 = 2
            _ = v21
        case "foo bar":
            v15 += (v15 >> ((v15) & 63))
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = (v2 & v2)
            }
        }
        v3 += v3
        for fi22 := 0; fi22 < 4; fi22++ {
            v23 := v5[v2]
            if false {
                var v24 uint32 = v2
                v2 = uint32(v3)
                switch "" {
                case "a":
                    v25 := GenS3{F0: -934}
                    _ = v25
                case "foo bar":
                    v24 += (v24 - (v24 + (v24 + v2)))
                }
                _ = v24
            }
            _ = fi22
            _ = v23
        }
        _ = fi20
    }
    switch "\n\t" {
    case "":
        var v26 uint8 = 255
        _ = v26
    case "foo bar":
        for fi27 := 0; fi27 < 16; fi27++ {
            fmt.Println(v3)
            if (true || false) {
                fmt.Println(fi27)
                var v28 int = fi27
                _ = v28
            } else {
                v29 := GenS4{F0: 0, F1: "hello"}
                if true {
                    var v30 uint = v15
                    _ = v30
                } else {
                    v5[511] = (^(fi27))
                }
                _ = v29
            }
            _ = fi27
        }
    default:
        v2 = v2
    }
    v1 = append(v1, v2)
    switch 237 {
    case 32768:
        var v31 float32 = ((-2.25 + (float32(float64(0.0)) * 3.14159)) / 1e10)
        _ = v31
    default:
        var v32 float64 = float64((0.0 / 1.0))
        _ = v32
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
