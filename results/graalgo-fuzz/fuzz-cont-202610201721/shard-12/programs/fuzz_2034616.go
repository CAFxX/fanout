package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 int16 }
func genFn0(p0 string, p1 uint8) int64 {
    var v7 uint8 = 42
    v8 := []int{128}
    var v9 string = p0
    v7 = (v7 ^ 140)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 980
}
func genFn1(p0 uint16, p1 int32) int64 {
    if true {
        if true {
            v10 := GenS1{F0: 175}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := []int{42}
                v13 := GenS1{F0: 128}
                _ = len(v12)
                _ = fi11
                _ = v12
                _ = v13
            }
            for fi14 := 0; fi14 < 14; fi14++ {
                fmt.Println(p0)
                _ = fi14
            }
            fmt.Println(p0)
        }
        v15 := []int{127, -989, 0, -9223372036854775808}
        switch 481 {
        case 42:
            var v16 bool = true
            _ = v16
        case 2147483647:
            fmt.Println(p1, p0)
        default:
            if true {
                var v17 uint16 = 329
                p1 = p1
                _ = v17
            } else {
                v18 := GenS2{F0: 377, F1: 1000}
                _ = v18
            }
        }
        _ = v15
    } else {
        if false {
            var v19 float64 = ((3.14159 + 1e10) / 1.0)
            switch -2147483648 {
            case -411:
                var v20 int64 = 9223372036854775807
                _ = v20
            case 192:
                switch "hello" {
                case "\n\t":
                    v21 := GenS1{F0: 167}
                    _ = v21
                case "foo bar":
                    fmt.Println(v19, v19, p1)
                case "hello":
                    var v22 int64 = 1
                    _ = v22
                default:
                    switch -1 {
                    case -1:
                        p1 -= p1
                    case -9223372036854775808:
                        v19 -= ((0.0 / ((float64((1.5 - 0.0)) + (((v19 * v19) + float64(0.1)) * v19)) - float64((0.0 - float32(v19))))) * float64(0.1))
                    case 2147483647:
                        v23 := []int{32767, 984, 42}
                        _ = v23
                    }
                }
            default:
                p1 = 2147483647
            }
            switch ("hello" + ("\n\t" + "a")) {
            case "foo bar":
                var v24 uint8 = 2
                _ = v24
            }
            _ = v19
        }
        switch -594 {
        case -9223372036854775808:
            fmt.Println(p0, p1)
        }
        var v25 string = "foo bar"
        _ = v25
    }
    p0 += 234
    _ = p0
    _ = p1
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int32 = -1
    var v3 uint32 = 894
    v4 := []uint32{}
    v5 := make(map[uint32]bool)
    var v6 int = 2147483647
    switch 128 {
    case -263:
        v2 = 817
    default:
        v3 = v3
    }
    if true {
        var v26 uint64 = 1
        v27 := GenS2{F0: (v3 - v3), F1: -1}
        _ = v26
        _ = v27
    } else {
        var v28 uint32 = (v3 / (((v3 % ((v3) | 1))) | 1))
        v3 += v28
        _ = v28
    }
    v2 -= v2
    v6 -= 127
    switch v6 {
    case 32767:
        v5[(v3 + (^(v3)))] = true
    case 32768:
        v4 = append(v4, (^(v3)))
    }
    if false {
        var v29 string = "hello"
        _ = v29
    } else {
        v2 = (v2 / ((v2) | 1))
        v3 = (v3 | (v3 >> ((0) & 63)))
        for fi30 := 0; fi30 < 12; fi30++ {
            if (!true) {
                v31 := GenS2{F0: 1, F1: 1000}
                var v32 uint16 = 32767
                var v33 int64 = -1
                _ = v31
                _ = v32
                _ = v33
            } else {
                if (2 <= 0) {
                    for fi34 := 0; fi34 < 10; fi34++ {
                        var v35 float64 = 1.5
                        switch "" {
                        case "\n\t":
                            var v36 uint8 = 255
                            _ = v36
                        }
                        _ = fi34
                        _ = v35
                    }
                    v37 := GenS2{F0: (^(v3)), F1: -1}
                    v4 = append(v4, v3)
                    _ = v37
                }
                v3 += 2147483647
                v4 = append(v4, (v3 << ((42) & 63)))
            }
            _ = fi30
        }
    }
    fmt.Println(v3, v2, v3)
    var v38 float32 = ((3.14159 * 1e10) / 3.14159)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
