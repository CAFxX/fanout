package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int; F1 int64; F2 int8 }
func genFn0(p0 uint16, p1 uint16, p2 int32) int8 {
    p2 = 417
    p2 -= p2
    fmt.Println(p0)
    switch 9223372036854775807 {
    case -813:
        if false {
            v5 := GenS1{F0: 2147483648, F1: -1}
            if true {
                if (0 != 2147483647) {
                    var v6 string = "foo bar"
                    v7 := GenS1{F0: 320, F1: (p2 | (p2 << ((65535) & 63)))}
                    v8 := []int{42, 2147483647, 42}
                    _ = v6
                    _ = v7
                    _ = v8
                }
            } else {
                fmt.Println(p0)
            }
            for fi9 := 0; fi9 < 13; fi9++ {
                var v10 uint32 = 65535
                v11 := GenS1{F0: uint32(p0), F1: p2}
                if (!true) {
                    var v12 int64 = 42
                    if true {
                        p2 += (p2 | p2)
                    } else {
                        for fi13 := 0; fi13 < 1; fi13++ {
                            switch "a" {
                            case "\n\t":
                                var v14 int = -9223372036854775808
                                _ = v14
                            }
                            if false {
                                var v15 uint = 591
                                _ = v15
                            }
                            _ = fi13
                        }
                        var v16 float64 = 0.0
                        _ = v16
                    }
                    v17 := []int{}
                    _ = v12
                    _ = v17
                }
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v5
        }
    case 2147483647:
        if (!false) {
            if (false && true) {
                v18 := GenS1{F0: 65535, F1: (-(p2))}
                var v19 string = ""
                _ = v18
                _ = v19
            }
        } else {
            if false {
                p2 = p2
            }
        }
    case -2147483648:
        if true {
            v20 := []int{int(p2)}
            var v21 float64 = 3.14159
            _ = v20
            _ = v21
        } else {
            p2 -= (p2 % ((p2) | 1))
            v22 := []int{-1, -972, 1}
            _ = v22
        }
    default:
        fmt.Println(p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return -128
}

func main() {
    var v1 string = ("foo bar" + "")
    v2 := make(map[int]int)
    var v3 int64 = 42
    var v4 int64 = v3
    v23 := GenS1{F0: 33, F1: -2147483648}
    if true {
        fmt.Println(v3)
    }
    v3 += v4
    v2[42] = -2147483648
    var v24 int8 = -128
    for fi25 := 0; fi25 < 6; fi25++ {
        var v26 float64 = (0.0 - float64((((1.5 / 3.14159) / 1.5) / -2.25)))
        var v27 uint = 255
        var v28 int64 = (v4 & v4)
        _ = fi25
        _ = v26
        _ = v27
        _ = v28
    }
    var v29 uint = 42
    v30 := GenS1{F0: 0, F1: 2147483647}
    v3 -= v3
    var v31 bool = (255 < 127)
    switch ("foo bar" + "") {
    case "a":
        if ((v24 - v24) > (v24 * int8(v3))) {
            switch v1 {
            case "a":
                for fi32 := 0; fi32 < 19; fi32++ {
                    v33 := GenS2{F0: (fi32 | fi32), F1: (-(v3)), F2: v24}
                    _ = fi32
                    _ = v33
                }
            case "foo bar":
                if (v24 <= (v24 ^ (v24 << (((^(v29))) & 63)))) {
                    for fi34 := 0; fi34 < 7; fi34++ {
                        var v35 int8 = (v24 << (((v29 >> ((v29) & 63))) & 63))
                        var v36 uint16 = 65535
                        v31 = (!((2 == 18446744073709551615) && v31))
                        _ = fi34
                        _ = v35
                        _ = v36
                    }
                }
            }
            for fi37 := 0; fi37 < 3; fi37++ {
                for fi38 := 0; fi38 < 20; fi38++ {
                    v3 += v4
                    var v39 int32 = -700
                    fmt.Println(v39, v31)
                    v31 = (255 <= (v29 / ((v29) | 1)))
                    switch fi38 {
                    case 1:
                        v1 = v1
                    }
                    _ = fi38
                    _ = v39
                }
                _ = fi37
            }
        }
    case "":
        v24 = v24
    }
    var v40 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v40
}
