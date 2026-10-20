package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 uint; F1 uint32; F2 uint16 }
type GenS4 struct { F0 bool; F1 int64; F2 int64 }
func genFn0(p0 uint8, p1 int16) bool {
    if true {
        var v5 float64 = 1.5
        _ = v5
    } else {
        v6 := []int{128, 31, 32767, 1}
        p1 = -1
        _ = v6
    }
    v7 := GenS1{F0: 4294967295, F1: 30, F2: 18446744073709551615}
    v8 := []int{284, -2147483648, 0, 1}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]bool)
    var v3 string = "hello"
    v4 := make(map[int]bool)
    fmt.Println(v3)
    var v9 int = 42
    if (!(222 > 0)) {
        delete(v2, v9)
    } else {
        v3 = "\n\t"
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(fi10, fi10)
            v9 -= (v9 >> ((2) & 63))
            _ = fi10
        }
        if false {
            var v11 uint32 = 799
            _ = v11
        } else {
            switch (v3 + v3) {
            case "\n\t":
                var v12 int8 = 42
                _ = v12
            case "":
                delete(v4, v9)
            default:
                for fi13 := 0; fi13 < 17; fi13++ {
                    v3 = "a"
                    v14 := GenS2{F0: 782, F1: 32767, F2: true}
                    _ = fi13
                    _ = v14
                }
            }
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        var v16 uint8 = 1
        for fi17 := 0; fi17 < 17; fi17++ {
            var v18 uint64 = 587
            if (973 > 435) {
                fmt.Println(v9, v9)
            } else {
                v3 = v3
                if true {
                    v16 += (v16 - (^(v16)))
                    v18 = (v18 << ((2) & 63))
                    v16 -= (v16 & (^(v16)))
                } else {
                    fmt.Println(v18, v18)
                }
                switch v3 {
                case "a":
                    v18 = 2
                case "foo bar":
                    if false {
                        switch v9 {
                        case 886:
                            fmt.Println(v18, v3, v18)
                        case 32767:
                            if (((!true) && (true && true)) && true) {
                                var v19 uint32 = 29
                                v18 = (^(v18))
                                _ = v19
                            }
                        case -2147483648:
                            v20 := GenS2{F0: 65535, F1: fi17, F2: false}
                            _ = v20
                        default:
                            if true {
                                fmt.Println(fi17)
                                v2[(fi15 | -1)] = (!((!(false && false)) && true))
                            }
                        }
                        fmt.Println(fi15, v3, v3)
                        var v21 int8 = -128
                        _ = v21
                    } else {
                        var v22 int = 176
                        fmt.Println(fi15)
                        v9 = 128
                        _ = v22
                    }
                case "":
                    if (!(v18 < v18)) {
                        v23 := GenS3{F0: 2, F1: 416, F2: 1}
                        v24 := GenS4{F0: true, F1: 461, F2: 0}
                        _ = v23
                        _ = v24
                    } else {
                        var v25 string = v3
                        var v26 bool = true
                        _ = v25
                        _ = v26
                    }
                default:
                    if (42 > 277) {
                        fmt.Println(v9)
                    } else {
                        v9 = int(v18)
                    }
                }
            }
            var v27 float64 = 0.0
            _ = fi17
            _ = v18
            _ = v27
        }
        _ = fi15
        _ = v16
    }
    _ = len(v1)
    v28 := GenS4{F0: (4294967295 != 1), F1: 9223372036854775807, F2: 1}
    v29 := GenS3{F0: 650, F1: 2147483648, F2: 32767}
    v9 += (v9 >> ((18446744073709551615) & 63))
    if (-2147483648 > -2147483648) {
        var v30 uint = 2
        switch (v3 + v3) {
        case "\n\t":
            fmt.Println(v9)
        default:
            fmt.Println(v9)
        }
        var v31 float32 = -2.25
        _ = v30
        _ = v31
    } else {
        v9 = v9
        v9 -= -9223372036854775808
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        _ = len(v1)
        v9 = (fi32 >> ((1) & 63))
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 uint64 = 4294967295
            var v35 int8 = -15
            fmt.Println(fi33, v34, fi32)
            _ = fi33
            _ = v34
            _ = v35
        }
        _ = fi32
    }
    fmt.Println(v3)
    v1[(v9 & v9)] = "\n\t"
    fmt.Println(v9, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v9
}
