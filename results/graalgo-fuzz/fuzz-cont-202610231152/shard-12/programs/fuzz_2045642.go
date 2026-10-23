package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint; F1 string }
type GenS4 struct { F0 int32 }
func genFn0() uint {
    var v6 uint32 = 42
    if (!((float32(0.1) / 1.0) < float32(((((0.1 - (1.5 - float64(3.14159))) + ((float64(((3.14159 + (0.1 / -2.25)) / float32(((1.5 - float64(0.0)) + (-2.25 - ((3.14159 / 1.0) / (((-2.25 - 3.14159) * (0.1 / (1.5 - ((0.0 / 1e10) + float64((0.1 * 3.14159)))))) * 1.5))))))) - ((1e10 + (0.0 - 0.0)) + -2.25)) / (float64((1e10 + 0.1)) - 1.5))) - -2.25) * (0.0 - (float64(0.0) / 1.5)))))) {
        v6 += v6
        v7 := GenS1{F0: -835}
        v8 := []int{2147483647, 866}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 20; fi9++ {
            v10 := []int{(fi9 - fi9)}
            v10 = append(v10, fi9)
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: 9223372036854775807}
        fmt.Println(v6)
        _ = v11
    }
    if (100 == -128) {
        v6 = (v6 % ((v6) | 1))
    } else {
        switch "\n\t" {
        case "\n\t":
            var v12 int16 = 1
            _ = v12
        case "":
            v13 := GenS1{F0: 32767}
            _ = v13
        case "a":
            fmt.Println(v6)
        default:
            if true {
                var v14 float64 = 0.1
                if true {
                    var v15 float64 = v14
                    for fi16 := 0; fi16 < 11; fi16++ {
                        fmt.Println(fi16)
                        if true {
                            fmt.Println(fi16, v6)
                        } else {
                            v6 = (v6 - v6)
                            if true {
                                fmt.Println(v15, v14, v6)
                            }
                        }
                        _ = fi16
                    }
                    var v17 int8 = 42
                    _ = v15
                    _ = v17
                }
                _ = v14
            }
        }
        v6 += 140
        if false {
            fmt.Println(v6)
            fmt.Println(v6)
        }
    }
    _ = v6
    return 42
}

func main() {
    var v1 int = 32767
    v2 := []uint32{80, 1, 2, 2147483648}
    var v3 float64 = ((float64(0.0) - 1e10) * (float64((1.5 - (0.0 - (1.5 * 1.5)))) / 1e10))
    var v4 uint32 = 494
    var v5 float32 = -2.25
    switch "hello" {
    case "":
        v18 := GenS2{F0: 2147483647, F1: -128}
        _ = v18
    case "foo bar":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "\n\t":
        var v19 float32 = ((-2.25 / float32(v3)) + (float32(v3) * 1.5))
        _ = v19
    }
    v1 = (v1 | -2147483648)
    if (-587 != 1) {
        fmt.Println(v5)
        switch "a" {
        case "":
            switch 1 {
            case 703:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 9223372036854775807:
                switch int(v4) {
                case 127:
                    v5 += float32(v3)
                case -1:
                    if false {
                        for fi20 := 0; fi20 < 19; fi20++ {
                            _ = len(v2)
                            var v21 bool = false
                            _ = fi20
                            _ = v21
                        }
                        v5 = (-2.25 / -2.25)
                    }
                case 32768:
                    var v22 uint8 = 53
                    _ = v22
                default:
                    fmt.Println(v5)
                }
            default:
                for fi23 := 0; fi23 < 20; fi23++ {
                    var v24 bool = (false || false)
                    fmt.Println(v1, v3, v3)
                    switch "" {
                    case "":
                        v25 := GenS3{F0: 255, F1: "\n\t"}
                        _ = v25
                    case "\n\t":
                        switch (("\n\t" + "\n\t") + "\n\t") {
                        case "foo bar":
                            _ = len(v2)
                        case "\n\t":
                            fmt.Println(fi23, v1)
                        }
                    }
                    _ = fi23
                    _ = v24
                }
            }
        }
    } else {
        v4 -= (v4 | (v4 / ((1) | 1)))
        switch "hello" {
        case "":
            v26 := GenS1{F0: v1}
            _ = v26
        case "a":
            var v27 int16 = 1000
            _ = v27
        case "\n\t":
            v3 += (v3 / v3)
        default:
            for fi28 := 0; fi28 < 7; fi28++ {
                v5 = v5
                v4 = v4
                var v29 uint32 = 138
                _ = fi28
                _ = v29
            }
        }
        fmt.Println(v4)
    }
    v30 := GenS4{F0: -1}
    fmt.Println(v1)
    if true {
        v31 := GenS4{F0: 1}
        _ = v31
    } else {
        switch "a" {
        case "hello":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            if len(v2) > 0 {
                v2[0] = (v4 >> ((65535) & 63))
            }
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        default:
            v1 -= (v1 / (((v1 & (v1 % ((v1) | 1)))) | 1))
        }
    }
    fmt.Println(v5)
    fmt.Println(v1)
    v2 = append(v2, 2147483647)
    fmt.Println(v1)
    for fi32 := 0; fi32 < 12; fi32++ {
        if false {
            for fi33 := 0; fi33 < 12; fi33++ {
                _ = len(v2)
                fmt.Println(fi32)
                _ = fi33
            }
        } else {
            var v34 uint64 = 799
            var v35 bool = (1 == 0)
            var v36 float64 = 1e10
            _ = v34
            _ = v35
            _ = v36
        }
        if ((v3 / v3) != float64(1.5)) {
            v5 += ((v5 - ((float32(1e10) * v5) - v5)) + float32((1.5 + (1e10 - ((0.0 * float64(v5)) + v3)))))
        }
        _ = fi32
    }
    var v37 int8 = 100
    v3 += v3
    v37 -= v37
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v37
    _ = v4
    _ = v5
}
