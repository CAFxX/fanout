package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 int64 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint8, p1 int) int32 {
    v7 := []int{(^(p1))}
    var v8 int = p1
    switch v8 {
    case 199:
        if true {
            var v9 float64 = ((0.0 + 0.1) * 0.0)
            _ = v9
        } else {
            if (((0.0 < 3.14159) && true) || false) {
                switch (v8 << ((18446744073709551615) & 63)) {
                case -9223372036854775808:
                    var v10 int64 = 1
                    _ = v10
                default:
                    if true {
                        if (!true) {
                            v8 = -1
                            var v11 float32 = 3.14159
                            if len(v7) >= 2 {
                                v7 = v7[1:len(v7)]
                            }
                            _ = v11
                        }
                        var v12 uint16 = 255
                        _ = v12
                    } else {
                        v8 -= 9223372036854775807
                    }
                }
                fmt.Println(p0, p0, p1)
                if len(v7) > 0 {
                    v7[0] = v8
                }
            }
            v13 := GenS1{F0: 0, F1: 967, F2: 1}
            _ = v13
        }
    case 103:
        switch (v8 - v8) {
        case -9223372036854775808:
            var v14 uint32 = 65535
            _ = v14
        case 2147483647:
            fmt.Println(p0, p1, p0)
        case 9223372036854775807:
            p0 = (p0 + p0)
        default:
            switch "\n\t" {
            case "foo bar":
                if len(v7) >= 2 {
                    v7 = v7[1:len(v7)]
                }
            case "hello":
                p0 -= p0
            case "\n\t":
                v7 = append(v7, 4294967295)
            }
        }
    case 4294967295:
        if false {
            var v15 uint = 4294967295
            var v16 uint8 = (p0 ^ 2)
            _ = v15
            _ = v16
        }
    default:
        var v17 uint64 = 42
        _ = v17
    }
    switch "hello" {
    case "a":
        var v18 int16 = 0
        _ = v18
    default:
        fmt.Println(p0)
    }
    var v19 float32 = (float32((((-2.25 - 3.14159) * -2.25) / -2.25)) * (1.5 / (((((3.14159 / (-2.25 / 3.14159)) * ((-2.25 + 1e10) + 1e10)) * 0.1) + 0.1) - ((-2.25 - (1.5 / 1.5)) - -2.25))))
    _ = p0
    _ = p1
    _ = v19
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := make(map[int]uint32)
    var v2 float32 = float32(1e10)
    v3 := make(map[int]uint32)
    var v4 int16 = 1000
    var v5 uint64 = 42
    v6 := make(map[uint32]uint32)
    for fi20 := 0; fi20 < 5; fi20++ {
        for fi21 := 0; fi21 < 7; fi21++ {
            var v22 int8 = -106
            if ((125 > -287) || false) {
                v5 -= v5
            } else {
                v22 = 127
                switch "foo bar" {
                case "hello":
                    if (!(0 < 1)) {
                        delete(v3, fi21)
                    } else {
                        var v23 string = ""
                        var v24 bool = (!true)
                        var v25 float64 = (-2.25 + 1.5)
                        _ = v23
                        _ = v24
                        _ = v25
                    }
                default:
                    var v26 uint = 4294967295
                    v26 += (v26 ^ v26)
                    _ = v26
                }
            }
            switch (fi20 - (^(fi20))) {
            case 32767:
                fmt.Println(v22)
            }
            _ = fi21
            _ = v22
        }
        _ = fi20
    }
    v27 := GenS2{F0: 0}
    v5 = (^(v5))
    if (true || true) {
        v4 = (-(v4))
        v28 := GenS1{F0: v5, F1: (^(v5)), F2: -567}
        _ = v28
    } else {
        v4 += v4
    }
    if ("" <= "a") {
        var v29 uint64 = v5
        v29 = (v29 * v29)
        _ = v29
    } else {
        v2 = v2
    }
    _ = len(v1)
    if true {
        fmt.Println(v4, v5, v4)
    }
    var v30 int64 = 1
    if true {
        if (255 >= 127) {
            if (!(!true)) {
                v30 = -1
                delete(v3, -9223372036854775808)
            }
            for fi31 := 0; fi31 < 3; fi31++ {
                var v32 float32 = float32(1.5)
                fmt.Println(v2, fi31)
                _ = fi31
                _ = v32
            }
            var v33 uint8 = 255
            _ = v33
        }
        var v34 uint16 = 0
        _ = v34
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
