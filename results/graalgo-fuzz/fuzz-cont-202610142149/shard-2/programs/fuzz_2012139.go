package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint, p1 int64, p2 int) uint {
    v6 := []int{42}
    if len(v6) > 0 {
        v6[0] = p2
    }
    p2 = 32767
    p1 += 42
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (p0 + (^(p0)))
}
func genFn1(p0 uint16, p1 int64) uint16 {
    switch 706 {
    case 127:
        p0 = 19
    default:
        p0 += (p0 * (p0 + p0))
    }
    switch "\n\t" {
    case "a":
        var v7 int8 = -1
        _ = v7
    case "foo bar":
        p0 = (p0 & (p0 ^ (p0 + (p0 | p0))))
    case "":
        var v8 string = "foo bar"
        _ = v8
    }
    p0 = p0
    _ = p0
    _ = p1
    return (^(p0))
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int = 9
    var v3 float32 = 1.5
    var v4 uint = 215
    v5 := make(map[uint32]string)
    if true {
        if true {
            if (!true) {
                v9 := []int{v2, -9223372036854775808, v2}
                v2 = 0
                _ = v9
            } else {
                fmt.Println(v4, v2, v3)
            }
            v10 := []int{v2, v2, v2}
            switch ("" + "\n\t") {
            case "foo bar":
                switch v2 {
                case 42:
                    v4 -= 4294967295
                case 62:
                    var v11 string = "foo bar"
                    _ = v11
                }
            case "":
                v12 := v1[32767]
                _ = v12
            case "\n\t":
                v13 := GenS1{F0: 18446744073709551615, F1: 0, F2: 42}
                _ = v13
            default:
                switch -9223372036854775808 {
                case 1:
                    switch ("\n\t" + (("a" + ("a" + "")) + "hello")) {
                    case "a":
                        var v14 uint8 = 245
                        _ = v14
                    case "":
                        fmt.Println(v3, v3)
                    default:
                        for fi15 := 0; fi15 < 6; fi15++ {
                            v16 := GenS1{F0: 2, F1: 4294967295, F2: 42}
                            _ = len(v5)
                            fmt.Println(v2, v2, fi15)
                            _ = fi15
                            _ = v16
                        }
                    }
                }
            }
            _ = v10
        }
    } else {
        var v17 uint16 = 32767
        _ = v17
    }
    if true {
        v4 = (^(v4))
        v3 += v3
        if true {
            v4 = v4
        } else {
            var v18 int32 = -287
            _ = v18
        }
    }
    fmt.Println(v2, v2)
    v19 := GenS1{F0: 0, F1: 18446744073709551615, F2: 1}
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS1{F0: 1, F1: 1, F2: 888}
        _ = fi20
        _ = v21
    }
    v22 := []int{}
    if len(v22) > 0 {
        v22[0] = (v2 / ((v2) | 1))
    }
    v23 := v1[128]
    var v24 int16 = 32767
    for fi25 := 0; fi25 < 2; fi25++ {
        fmt.Println(v24, fi25)
        fmt.Println(fi25)
        _ = fi25
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        var v27 float32 = v3
        v4 = 42
        v27 += v3
        _ = fi26
        _ = v27
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int = fi28
        for fi30 := 0; fi30 < 17; fi30++ {
            var v31 int64 = 1
            var v32 uint16 = 65535
            v24 += -714
            _ = fi30
            _ = v31
            _ = v32
        }
        if (!(0 < 89)) {
            var v33 float32 = (v3 * (3.14159 / 3.14159))
            _ = v33
        }
        _ = fi28
        _ = v29
    }
    v22 = append(v22, -1)
    v34 := GenS2{F0: 2}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
