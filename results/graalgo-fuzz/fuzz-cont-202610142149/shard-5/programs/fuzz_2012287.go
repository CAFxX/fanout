package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint }
func genFn0(p0 uint64, p1 int64, p2 uint8) uint {
    p1 = (p1 * p1)
    fmt.Println(p0)
    v5 := GenS1{F0: 844, F1: 65535, F2: 710}
    switch ("hello" + "\n\t") {
    case "foo bar":
        v6 := []int{1}
        _ = v6
    case "hello":
        v7 := []int{}
        _ = v7
    }
    v8 := GenS2{F0: 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v8
    return 65535
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 int16 = 1000
    v3 := make(map[string]uint32)
    var v4 uint64 = 1
    switch -419 {
    case -1:
        if (false && false) {
            fmt.Println(v2, v2, v1)
        } else {
            for fi9 := 0; fi9 < 17; fi9++ {
                v10 := GenS1{F0: (fi9 | (fi9 | 9223372036854775807)), F1: 2147483648, F2: 547}
                _ = fi9
                _ = v10
            }
        }
    default:
        v2 = 32767
    }
    v4 = (^(v4))
    _ = len(v3)
    v2 -= v2
    if true {
        switch "foo bar" {
        case "hello":
            var v11 uint32 = 4294967295
            _ = v11
        case "":
            switch -434 {
            case -1:
                switch 2147483647 {
                case -693:
                    var v12 uint64 = v4
                    _ = v12
                }
            case -9223372036854775808:
                var v13 uint = 42
                _ = v13
            default:
                if false {
                    v4 = (v4 + (v4 - v4))
                    if (2 >= 255) {
                        var v14 float64 = 1.5
                        for fi15 := 0; fi15 < 5; fi15++ {
                            for fi16 := 0; fi16 < 13; fi16++ {
                                var v17 string = ((("" + ("a" + "foo bar")) + "a") + "\n\t")
                                v17 = (("hello" + ((("" + (v17 + v17)) + (("" + (v17 + v17)) + ("hello" + v17))) + ("hello" + v17))) + "foo bar")
                                _ = fi16
                                _ = v17
                            }
                            _ = fi15
                        }
                        v1 = (v1 << ((42) & 63))
                        _ = v14
                    } else {
                        if false {
                            var v18 uint64 = (^(v4))
                            if true {
                                var v19 int16 = v2
                                _ = v19
                            }
                            _ = v18
                        }
                        v2 -= (v2 + int16(v1))
                        v20 := GenS3{F0: -2147483648}
                        _ = v20
                    }
                }
            }
        case "a":
            fmt.Println(v2)
        default:
            for fi21 := 0; fi21 < 19; fi21++ {
                var v22 string = "hello"
                var v23 uint16 = 32767
                _ = fi21
                _ = v22
                _ = v23
            }
        }
        fmt.Println(v4, v2)
    }
    v24 := GenS2{F0: 2147483647}
    var v25 uint64 = v4
    if true {
        if false {
            var v26 int64 = -695
            var v27 int16 = v2
            _ = v26
            _ = v27
        }
        v28 := []int{606, -2147483648}
        _ = v28
    } else {
        if ((42 < 291) && (!((-(v2)) >= 47))) {
            if false {
                v2 += -1
            } else {
                if (457 <= genFn0((v4 * (v4 & (^(v25)))), (v1 | (v1 & 1)), 255)) {
                    switch 127 {
                    case 128:
                        v29 := []int{}
                        _ = v29
                    case -1:
                        v4 = 136
                    case -2147483648:
                        fmt.Println(v4, v25, v4)
                    }
                    var v30 int16 = int16(v1)
                    var v31 uint32 = 65535
                    _ = v30
                    _ = v31
                }
                var v32 float64 = 0.0
                _ = v32
            }
        }
        v4 = 4294967295
        for fi33 := 0; fi33 < 6; fi33++ {
            fmt.Println(fi33, fi33)
            if (!(false && false)) {
                var v34 int16 = (v2 % (((v2 / ((v2) | 1))) | 1))
                fmt.Println(v2, v4)
                _ = v34
            } else {
                var v35 uint = 236
                var v36 uint16 = 2
                _ = v35
                _ = v36
            }
            _ = fi33
        }
    }
    v37 := GenS4{F0: -87, F1: 267}
    fmt.Println(v2, v2, v4)
    v4 -= (^(v25))
    fmt.Println(v25)
    var v38 uint64 = (v25 * (v4 % ((2) | 1)))
    v2 -= v2
    var v39 float32 = 1.5
    v38 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
