package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint; F1 uint16 }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 uint64, p1 uint16, p2 int8) bool {
    if true {
        v7 := []int{-9223372036854775808, 0}
        p1 -= (p1 - (p1 - (p1 | 720)))
        _ = len(v7)
        _ = v7
    } else {
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 int8 = p2
            _ = fi8
            _ = v9
        }
    }
    p2 = (p2 + int8(p1))
    for fi10 := 0; fi10 < 7; fi10++ {
        v11 := []int{}
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    return false
}

func main() {
    var v1 int8 = 1
    var v2 uint64 = 0
    var v3 uint8 = 2
    v4 := make(map[int]bool)
    v5 := []string{}
    var v6 uint8 = v3
    v6 = v3
    v12 := v4[42]
    v13 := GenS1{F0: (v3 + (v6 - (v6 + v6))), F1: 565, F2: -2147483648}
    if true {
        v14 := GenS2{F0: 668}
        for fi15 := 0; fi15 < 9; fi15++ {
            fmt.Println(v6, v1)
            fmt.Println(v6, v3)
            _ = fi15
        }
        fmt.Println(v1)
        _ = v14
    } else {
        v16 := GenS1{F0: uint8(v2), F1: 255, F2: int32(v2)}
        v17 := GenS3{F0: v3}
        if v12 {
            v18 := GenS4{F0: -1}
            _ = v18
        } else {
            var v19 int8 = v1
            _ = v19
        }
        _ = v16
        _ = v17
    }
    var v20 int = 2147483647
    switch ("" + "a") {
    case "\n\t":
        var v21 uint64 = (v2 / (((v2 % (((v2 | (v2 >> ((385) & 63)))) | 1))) | 1))
        _ = v21
    case "":
        v22 := GenS4{F0: -1}
        _ = v22
    case "foo bar":
        for fi23 := 0; fi23 < 4; fi23++ {
            fmt.Println(v12, v12, v3)
            _ = fi23
        }
    }
    switch v20 {
    case -1:
        v24 := GenS5{F0: 4294967295, F1: uint16(v3)}
        _ = v24
    case 0:
        var v25 float32 = (1.5 - 3.14159)
        _ = v25
    case 9223372036854775807:
        v12 = false
    default:
        var v26 uint8 = (v3 - v3)
        _ = v26
    }
    if ((("hello" + "foo bar") < "foo bar") && v12) {
        for fi27 := 0; fi27 < 11; fi27++ {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            v2 -= 42
            v5 = append(v5, ("a" + "hello"))
            _ = fi27
        }
        var v28 string = "hello"
        if v12 {
            if false {
                v29 := GenS4{F0: 1}
                v3 = v6
                v3 -= v3
                _ = v29
            } else {
                for fi30 := 0; fi30 < 14; fi30++ {
                    v3 -= v3
                    var v31 uint8 = v6
                    switch ("hello" + "\n\t") {
                    case "foo bar":
                        v2 = v2
                    case "a":
                        if v12 {
                            v20 = v20
                            if (((v31 + (^(v31))) <= v3) && v12) {
                                switch (v28 + "a") {
                                case "":
                                    v31 = v31
                                case "a":
                                    _ = len(v5)
                                case "\n\t":
                                    var v32 uint = 495
                                    _ = v32
                                default:
                                    fmt.Println(v20, v31, v20)
                                }
                                var v33 uint64 = (^(v2))
                                _ = v33
                            }
                            v31 -= (^(v3))
                        } else {
                            var v34 float64 = -2.25
                            v1 -= (v1 % ((100) | 1))
                            _ = v34
                        }
                    case "":
                        for fi35 := 0; fi35 < 16; fi35++ {
                            v31 = (v31 & (v6 * (v6 | v3)))
                            v3 = v6
                            _ = fi35
                        }
                    }
                    _ = fi30
                    _ = v31
                }
                v36 := GenS1{F0: (^(v3)), F1: 1, F2: 0}
                v3 = v6
                _ = v36
            }
        } else {
            switch v20 {
            case 127:
                fmt.Println(v28, v6, v12)
            default:
                v6 = 0
            }
            for fi37 := 0; fi37 < 17; fi37++ {
                switch (fi37 - int(v2)) {
                case 32768:
                    fmt.Println(v3, v12)
                default:
                    var v38 uint16 = 657
                    _ = v38
                }
                _ = fi37
            }
            v39 := GenS6{F0: 1, F1: v12}
            _ = v39
        }
        _ = v28
    } else {
        fmt.Println(v20, v3)
    }
    v40 := GenS5{F0: 1, F1: 32767}
    v2 -= v2
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
