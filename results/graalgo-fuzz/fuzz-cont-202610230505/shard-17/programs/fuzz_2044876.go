package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }

func main() {
    var v1 int64 = 1
    v2 := []uint32{2}
    var v3 float64 = 0.0
    var v4 bool = (2147483647 >= 0)
    if (2147483647 < 2) {
        var v5 string = "a"
        _ = v5
    }
    switch "\n\t" {
    case "a":
        var v6 uint8 = 243
        _ = v6
    case "hello":
        v1 = 9223372036854775807
    case "foo bar":
        for fi7 := 0; fi7 < 2; fi7++ {
            for fi8 := 0; fi8 < 9; fi8++ {
                var v9 int8 = -23
                fmt.Println(v4, v9, v4)
                _ = fi8
                _ = v9
            }
            v4 = false
            _ = fi7
        }
    }
    v4 = false
    v10 := GenS1{F0: -2147483648}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            v4 = false
            _ = fi12
        }
        var v13 uint8 = 1
        var v14 uint16 = 770
        _ = fi11
        _ = v13
        _ = v14
    }
    if ((2 == 0) && true) {
        if v4 {
            switch "foo bar" {
            case "hello":
                var v15 bool = ((1.5 * -2.25) == (float32(v3) / 1.5))
                _ = v15
            case "a":
                if v4 {
                    var v16 float32 = (0.1 * 1e10)
                    v1 = 42
                    v3 = v3
                    _ = v16
                } else {
                    var v17 float32 = 0.0
                    v3 = v3
                    _ = v17
                }
            case "foo bar":
                var v18 uint32 = 2
                _ = v18
            default:
                v3 -= v3
            }
        }
        if true {
            v3 = (((1e10 / v3) * ((0.0 - (v3 / v3)) + v3)) / float64(float32((float64((0.1 * 0.0)) + (-2.25 * v3)))))
            var v19 int32 = 0
            _ = v19
        }
        for fi20 := 0; fi20 < 3; fi20++ {
            v4 = (0 != 65535)
            if (-938 < 0) {
                fmt.Println(v4, v4, v4)
                var v21 uint8 = 2
                _ = v21
            } else {
                switch "hello" {
                case "hello":
                    _ = len(v2)
                case "a":
                    v22 := GenS1{F0: 0}
                    _ = v22
                }
                var v23 uint8 = 2
                _ = v23
            }
            _ = fi20
        }
    }
    var v24 uint = 2
    var v25 uint64 = 2
    fmt.Println(v4, v24)
    v1 = v1
    var v26 int32 = 42
    fmt.Println(v3, v4, v4)
    var v27 int64 = 42
    if false {
        v28 := GenS1{F0: (v26 % ((v26) | 1))}
        _ = v28
    } else {
        if v4 {
            v25 = v25
            v3 = float64(((float32(v3) + 0.1) - ((((0.1 + -2.25) * 3.14159) / 1.0) - (0.1 + (float32(float64(-2.25)) / 1.0)))))
        } else {
            switch 42 {
            case -9223372036854775808:
                fmt.Println(v27, v26, v3)
            case -2147483648:
                switch 128 {
                case -2147483648:
                    var v29 int16 = 0
                    _ = v29
                case 893:
                    v4 = v4
                case 0:
                    var v30 float64 = (1e10 + v3)
                    _ = v30
                default:
                    fmt.Println(v4, v24, v24)
                }
            case 2147483647:
                switch "\n\t" {
                case "a":
                    var v31 uint8 = 243
                    _ = v31
                default:
                    v26 = (v26 - -2147483648)
                }
            default:
                if v4 {
                    var v32 bool = v4
                    if len(v2) > 0 {
                        v2[0] = 4294967295
                    }
                    _ = v32
                }
            }
        }
        for fi33 := 0; fi33 < 15; fi33++ {
            var v34 uint16 = 32767
            if v4 {
                switch "hello" {
                case "\n\t":
                    switch ((("a" + ("foo bar" + "foo bar")) + "a") + "hello") {
                    case "a":
                        for fi35 := 0; fi35 < 7; fi35++ {
                            v2 = append(v2, 107)
                            v36 := GenS1{F0: int32(v25)}
                            _ = fi35
                            _ = v36
                        }
                    }
                case "hello":
                    for fi37 := 0; fi37 < 3; fi37++ {
                        if len(v2) > 0 {
                            v2[0] = 1
                        }
                        _ = fi37
                    }
                }
            } else {
                if false {
                    v2 = append(v2, 4294967295)
                } else {
                    fmt.Println(v27)
                }
                if (!(v25 != (v25 + (v25 + v25)))) {
                    switch fi33 {
                    case -2147483648:
                        fmt.Println(fi33, v25, v26)
                    case -1:
                        v24 -= (v24 * 2)
                    case 127:
                        var v38 float64 = (v3 / v3)
                        _ = v38
                    }
                    var v39 int16 = 518
                    _ = v39
                }
            }
            v25 += (v25 + (v25 >> ((v24) & 63)))
            _ = fi33
            _ = v34
        }
        v4 = v4
    }
    v24 -= 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
