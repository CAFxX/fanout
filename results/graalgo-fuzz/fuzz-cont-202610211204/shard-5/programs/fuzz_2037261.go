package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 uint16 }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 uint; F1 uint16; F2 string }
func genFn0(p0 int, p1 int32, p2 uint16) uint64 {
    var v4 int32 = (p1 & (-(p1)))
    v5 := GenS1{F0: (-(p1)), F1: -737, F2: p2}
    fmt.Println(p1, v4)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 2
}
func genFn1(p0 bool) int {
    v6 := GenS1{F0: 2147483647, F1: -649, F2: 443}
    p0 = (-195 != 1)
    v7 := []int{}
    _ = p0
    _ = v6
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint32 = 42
    var v2 float64 = 1.5
    var v3 uint64 = 1
    fmt.Println(v3, v1, v2)
    var v8 string = "a"
    switch (v8 + v8) {
    case "hello":
        var v9 string = v8
        _ = v9
    }
    v10 := []int{-2147483648, 759}
    v8 = "hello"
    if false {
        var v11 float32 = 0.1
        _ = v11
    }
    v2 -= (3.14159 / (float64(3.14159) + v2))
    if len(v10) >= 2 {
        v10 = v10[1:len(v10)]
    }
    switch 42 {
    case 32767:
        fmt.Println(v1)
    case -781:
        fmt.Println(v1, v2)
    }
    var v12 uint8 = 59
    if false {
        v1 += v1
    } else {
        v10 = append(v10, 586)
        for fi13 := 0; fi13 < 7; fi13++ {
            v14 := GenS1{F0: int32(v3), F1: 570, F2: 42}
            switch (v8 + v8) {
            case "hello":
                fmt.Println(v2, v12, v12)
            case "a":
                v12 = (v12 << ((268) & 63))
            case "foo bar":
                var v15 uint16 = 576
                _ = v15
            }
            if (!true) {
                switch v8 {
                case "a":
                    var v16 uint8 = v12
                    _ = v16
                }
                v17 := GenS1{F0: int32(v1), F1: 42, F2: 0}
                _ = v17
            } else {
                v10 = append(v10, int(v1))
                if false {
                    switch (("" + v8) + v8) {
                    case "hello":
                        v12 += (v12 / (((v12 + v12)) | 1))
                    }
                    if true {
                        switch v8 {
                        case "\n\t":
                            if len(v10) > 0 {
                                v10[0] = (fi13 | fi13)
                            }
                        case "a":
                            var v18 uint = 375
                            _ = v18
                        }
                    } else {
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v8 = ((v8 + v8) + v8)
                            var v20 bool = (42 < 602)
                            fmt.Println(fi13, v8)
                            _ = fi19
                            _ = v20
                        }
                        v12 += v12
                    }
                    if false {
                        fmt.Println(v12)
                    }
                }
            }
            _ = fi13
            _ = v14
        }
    }
    for fi21 := 0; fi21 < 2; fi21++ {
        if true {
            switch int(v1) {
            case -2147483648:
                switch "a" {
                case "hello":
                    fmt.Println(v3, v8)
                default:
                    var v22 int8 = 42
                    _ = v22
                }
            default:
                var v23 bool = (((!("hello" >= (v8 + ""))) && false) && false)
                _ = v23
            }
        } else {
            v24 := GenS2{F0: 42, F1: (v1 | (^(v1)))}
            _ = v24
        }
        if len(v10) > 0 {
            v10[0] = (^(fi21))
        }
        _ = fi21
    }
    v25 := GenS2{F0: 2147483647, F1: (v1 & 65535)}
    for fi26 := 0; fi26 < 19; fi26++ {
        fmt.Println(v1)
        for fi27 := 0; fi27 < 17; fi27++ {
            if (!false) {
                v10 = append(v10, fi26)
                v28 := GenS3{F0: 679, F1: 255, F2: v8}
                _ = v28
            } else {
                var v29 int8 = 1
                var v30 float64 = float64(((0.0 - 1e10) - 3.14159))
                _ = v29
                _ = v30
            }
            v1 -= uint32(v12)
            if len(v10) > 0 {
                v10[0] = genFn1(false)
            }
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = fi27
        }
        _ = fi26
    }
    v12 += 2
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v8
}
