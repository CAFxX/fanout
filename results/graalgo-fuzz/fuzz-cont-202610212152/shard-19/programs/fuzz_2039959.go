package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int64 }
func genFn0(p0 int16) bool {
    v6 := GenS1{F0: (p0 % ((p0) | 1)), F1: 1}
    v7 := GenS1{F0: p0, F1: 42}
    _ = p0
    _ = v6
    _ = v7
    return (!((1.5 <= 0.1) && false))
}
func genFn1(p0 uint8, p1 int64) int64 {
    for fi8 := 0; fi8 < 4; fi8++ {
        fmt.Println(p0, p1)
        var v9 uint = 574
        for fi10 := 0; fi10 < 17; fi10++ {
            if ((-(fi10)) == (fi10 + fi10)) {
                if false {
                    var v11 float32 = 1e10
                    _ = v11
                } else {
                    v12 := []int{int(p1), (fi10 % ((673) | 1))}
                    if true {
                        p1 -= p1
                        var v13 int = 258
                        var v14 int = -368
                        v13 = v14
                        p1 = 0
                        v15 := GenS1{F0: 101, F1: 4294967295}
                        v16 := GenS2{F0: uint64(p0)}
                        if false {
                            var v17 int64 = p1
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            _ = v17
                        } else {
                            fmt.Println(p1)
                            v14 = (^(v13))
                            fmt.Println(v13, v9, v14)
                        }
                        _ = v13
                        _ = v14
                        _ = v15
                        _ = v16
                    }
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = v12
                }
            }
            var v18 float32 = 1.5
            _ = fi10
            _ = v18
        }
        _ = fi8
        _ = v9
    }
    var v19 float32 = 1e10
    _ = p0
    _ = p1
    _ = v19
    return (p1 >> ((65535) & 63))
}

func main() {
    v1 := []uint32{455, 2147483648, 1}
    var v2 bool = true
    var v3 int = 9223372036854775807
    var v4 int32 = -847
    v5 := []uint32{1}
    v1 = append(v1, 2)
    v3 += (v3 / ((v3) | 1))
    var v20 int64 = 0
    var v21 float32 = (1.5 * ((0.1 * 1.5) + 1e10))
    var v22 string = "hello"
    if v2 {
        var v23 uint = 800
        switch 32767 {
        case 1:
            if v2 {
                if v2 {
                    fmt.Println(v4)
                    switch ("" + v22) {
                    case "\n\t":
                        var v24 uint = (v23 * v23)
                        _ = v24
                    case "":
                        v25 := GenS3{F0: v3}
                        _ = v25
                    case "a":
                        v2 = false
                    default:
                        switch v3 {
                        case 1:
                            switch -2147483648 {
                            case -2147483648:
                                v26 := GenS2{F0: 18446744073709551615}
                                _ = v26
                            case -404:
                                v4 = v4
                            case 4294967295:
                                v20 = (v20 | v20)
                            }
                        case 2147483647:
                            fmt.Println(v22, v20, v3)
                        }
                    }
                }
            } else {
                if (0.1 == (3.14159 / v21)) {
                    var v27 uint8 = 138
                    _ = v27
                }
                v21 += (((v21 - (v21 * 0.0)) - float32(0.1)) * 3.14159)
                for fi28 := 0; fi28 < 2; fi28++ {
                    v29 := GenS4{F0: -687, F1: v20}
                    var v30 uint = 573
                    _ = fi28
                    _ = v29
                    _ = v30
                }
            }
        }
        _ = v23
    }
    var v31 uint16 = 820
    if true {
        for fi32 := 0; fi32 < 1; fi32++ {
            if v2 {
                fmt.Println(v31)
                var v33 uint32 = 65535
                _ = v33
            }
            var v34 float32 = 1.5
            v20 -= (^(v20))
            _ = fi32
            _ = v34
        }
    } else {
        v21 = v21
        v5 = append(v5, 1)
        var v35 int32 = 2147483647
        _ = v35
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
