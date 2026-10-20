package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 180
    v3 := make(map[uint32]uint32)
    var v4 uint8 = 0
    var v5 uint = 0
    v4 = v2
    for fi6 := 0; fi6 < 8; fi6++ {
        switch (fi6 - fi6) {
        case 4294967295:
            switch (fi6 >> ((v5) & 63)) {
            case -195:
                if true {
                    v5 = v5
                } else {
                    v7 := GenS1{F0: 349}
                    fmt.Println(fi6, fi6)
                    _ = v7
                }
            case -101:
                v8 := GenS1{F0: 172}
                _ = v8
            case 128:
                for fi9 := 0; fi9 < 16; fi9++ {
                    var v10 float64 = (0.0 - ((0.0 * float64(float32(0.1))) + (1e10 + ((0.1 * float64((1e10 + (0.0 * 3.14159)))) - 0.0))))
                    if false {
                        if true {
                            v4 -= v2
                            fmt.Println(v5)
                            var v11 int = (fi6 | fi9)
                            _ = v11
                        }
                        v12 := v1[(fi9 % (((fi9 * (fi9 ^ fi9))) | 1))]
                        _ = v12
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                v5 += 65535
            }
        default:
            v13 := GenS1{F0: 42}
            _ = v13
        }
        _ = fi6
    }
    if ("hello" > "") {
        fmt.Println(v5, v4)
        var v14 int8 = 100
        _ = v14
    } else {
        v15 := GenS1{F0: 83}
        _ = v15
    }
    for fi16 := 0; fi16 < 6; fi16++ {
        var v17 uint32 = 42
        v18 := []int{2147483647}
        var v19 string = (("" + "a") + ("" + "hello"))
        _ = fi16
        _ = v17
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 16; fi20++ {
        var v21 float32 = -2.25
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 string = ("\n\t" + ("a" + "a"))
            v2 += (v2 - 128)
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    v5 += 255
    v4 += (v2 % ((v2) | 1))
    var v24 uint16 = 65535
    v25 := GenS2{F0: -2147483648, F1: -1}
    if (((-2.25 / 0.1) * 3.14159) > 3.14159) {
        var v26 float64 = 1.5
        for fi27 := 0; fi27 < 5; fi27++ {
            for fi28 := 0; fi28 < 18; fi28++ {
                if true {
                    v24 = v24
                } else {
                    v29 := []int{fi28, (fi28 - -9223372036854775808), fi27}
                    if ((v5 << (((v5 ^ (v5 - (v5 - 65535)))) & 63)) != v5) {
                        if (v5 == (v5 ^ (v5 & 963))) {
                            v30 := GenS1{F0: uint32(v4)}
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v26, v24)
                            case "a":
                                v2 -= (^(v2))
                            default:
                                _ = len(v1)
                            }
                            if len(v29) > 0 {
                                v29[0] = (-(fi27))
                            }
                            _ = v30
                        } else {
                            v24 += 1
                        }
                        switch fi28 {
                        case -922:
                            if true {
                                fmt.Println(v26)
                                _ = len(v29)
                                var v31 uint16 = 65535
                                _ = v31
                            }
                        case 9223372036854775807:
                            v2 += v4
                        }
                        var v32 uint64 = 4294967295
                        _ = v32
                    } else {
                        fmt.Println(fi28, fi27, v5)
                    }
                    _ = v29
                }
                fmt.Println(v26, v4, fi28)
                var v33 uint = v5
                _ = fi28
                _ = v33
            }
            _ = fi27
        }
        v34 := GenS2{F0: -2147483648, F1: 0}
        _ = v26
        _ = v34
    } else {
        if false {
            fmt.Println(v5, v5)
            fmt.Println(v4, v5, v24)
        }
        fmt.Println(v24, v2)
        if false {
            var v35 uint64 = 208
            v36 := []int{9223372036854775807, 2147483647, 343, 128}
            _ = v35
            _ = v36
        }
    }
    var v37 float64 = -2.25
    var v38 int64 = 9223372036854775807
    var v39 int64 = 9223372036854775807
    if true {
        fmt.Println(v24, v4)
    }
    v4 += (v2 << ((v5) & 63))
    for fi40 := 0; fi40 < 18; fi40++ {
        v41 := []int{fi40, 42}
        switch "hello" {
        case "\n\t":
            fmt.Println(fi40, v39, v24)
        default:
            var v42 uint32 = 2
            _ = v42
        }
        _ = fi40
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
