package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int64; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint64; F2 int32 }
func genFn0(p0 int64, p1 int64, p2 int8) uint8 {
    fmt.Println(p1, p1)
    if (true && (-2147483648 != 0)) {
        var v5 float32 = (1.5 + ((1e10 + 0.1) * (((1e10 / 1.0) - -2.25) / (-2.25 + -2.25))))
        v6 := GenS1{F0: 1, F1: 95}
        fmt.Println(p0)
        _ = v5
        _ = v6
    } else {
        p2 = 100
        fmt.Println(p2, p1, p1)
        v7 := GenS1{F0: -2147483648, F1: p2}
        _ = v7
    }
    v8 := []int{42, 32768, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[uint32]uint32)
    var v3 uint = 84
    var v4 int32 = 42
    for fi9 := 0; fi9 < 16; fi9++ {
        v1 = ""
        delete(v2, 42)
        _ = fi9
    }
    if false {
        var v10 float32 = -2.25
        var v11 float32 = 3.14159
        _ = v10
        _ = v11
    }
    v12 := []int{127}
    var v13 uint8 = 42
    for fi14 := 0; fi14 < 7; fi14++ {
        v13 = (v13 | (v13 | (v13 ^ (v13 ^ 1))))
        v15 := GenS2{F0: -635, F1: 0}
        _ = fi14
        _ = v15
    }
    if true {
        switch 0 {
        case 127:
            v1 = ((v1 + v1) + "hello")
        case 9223372036854775807:
            v3 = v3
        case 166:
            var v16 uint16 = 114
            _ = v16
        }
    } else {
        switch v1 {
        case "hello":
            fmt.Println(v3, v3, v4)
        default:
            var v17 int64 = 972
            _ = v17
        }
        v18 := GenS3{F0: v4}
        if false {
            for fi19 := 0; fi19 < 14; fi19++ {
                if (!(!(-721 >= -9223372036854775808))) {
                    _ = len(v2)
                } else {
                    switch -308 {
                    case -9223372036854775808:
                        var v20 int = (-(fi19))
                        _ = v20
                    }
                    var v21 string = ((("hello" + v1) + v1) + v1)
                    if false {
                        v21 = ((v1 + (v21 + (("foo bar" + v1) + v21))) + "\n\t")
                        if true {
                            v3 -= v3
                            fmt.Println(v13)
                            v22 := GenS2{F0: 498, F1: 4294967295}
                            _ = v22
                        } else {
                            v23 := GenS4{F0: -32768}
                            _ = v23
                        }
                    }
                    _ = v21
                }
                v13 = (v13 - (^(v13)))
                _ = fi19
            }
            fmt.Println(v13)
        }
        _ = v18
    }
    if false {
        switch v1 {
        case "\n\t":
            var v24 float64 = ((3.14159 + (0.1 - float64((1.5 * (1e10 + -2.25))))) / 1.5)
            _ = v24
        default:
            for fi25 := 0; fi25 < 7; fi25++ {
                var v26 uint64 = 389
                _ = fi25
                _ = v26
            }
        }
        v13 -= (v13 + genFn0(-249, 0, 100))
    }
    fmt.Println(v1, v3)
    fmt.Println(v1)
    for fi27 := 0; fi27 < 8; fi27++ {
        var v28 float64 = -2.25
        switch v1 {
        case "foo bar":
            v29 := v2[0]
            _ = v29
        case "a":
            var v30 int32 = 1
            _ = v30
        }
        _ = fi27
        _ = v28
    }
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint = v3
        var v33 string = v1
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v1 = v1
    if false {
        v34 := GenS5{F0: -408, F1: 42, F2: (v4 & -2147483648)}
        _ = v34
    }
    var v35 int16 = -562
    v3 += v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v35
    _ = v4
}
