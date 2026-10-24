package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int32 }
func genFn0(p0 uint8, p1 uint32) uint16 {
    p0 -= 1
    var v7 int64 = -568
    v8 := []int{587, -2147483648, 862}
    p1 += p1
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 543
}

func main() {
    var v1 uint16 = 104
    var v2 float64 = 3.14159
    var v3 uint8 = 255
    v4 := []int{32767, 0}
    v5 := make(map[string]string)
    var v6 uint16 = v1
    v6 += 203
    switch "foo bar" {
    case "foo bar":
        v9 := GenS1{F0: 129, F1: 42}
        _ = v9
    case "":
        fmt.Println(v2)
    default:
        if true {
            switch "\n\t" {
            case "hello":
                switch 42 {
                case 42:
                    switch 32768 {
                    case 32767:
                        fmt.Println(v3)
                    case 431:
                        var v10 string = "foo bar"
                        _ = v10
                    case -658:
                        v1 += v1
                    default:
                        v3 += v3
                    }
                default:
                    v2 -= 0.0
                }
            case "":
                v11 := GenS1{F0: 2147483647, F1: 18446744073709551615}
                _ = v11
            }
            var v12 int8 = 0
            _ = v12
        } else {
            v13 := GenS2{F0: 540, F1: -9223372036854775808}
            switch 42 {
            case 2147483647:
                var v14 float64 = 3.14159
                _ = v14
            case 4294967295:
                v6 = (v6 & 255)
            }
            var v15 int64 = -1
            _ = v13
            _ = v15
        }
    }
    var v16 uint = 2
    v17 := GenS2{F0: v6, F1: -9223372036854775808}
    fmt.Println(v6)
    for fi18 := 0; fi18 < 4; fi18++ {
        var v19 int64 = -142
        _ = fi18
        _ = v19
    }
    var v20 uint64 = 4294967295
    for fi21 := 0; fi21 < 10; fi21++ {
        for fi22 := 0; fi22 < 15; fi22++ {
            if len(v4) > 0 {
                v4[0] = fi21
            }
            _ = fi22
        }
        _ = fi21
    }
    if true {
        v23 := GenS3{F0: 127}
        if (v1 > (v6 + v6)) {
            switch "a" {
            case "hello":
                v16 -= (v16 ^ (^(v16)))
            default:
                v1 = 0
            }
            var v24 string = "hello"
            switch "hello" {
            case "\n\t":
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 int16 = -1
                    v20 = (^(v20))
                    _ = fi25
                    _ = v26
                }
            default:
                if ((424 >= -1) || (false && false)) {
                    fmt.Println(v20)
                } else {
                    var v27 float64 = v2
                    _ = v27
                }
            }
            _ = v24
        } else {
            v3 += 1
        }
        _ = v23
    }
    switch "a" {
    case "hello":
        if false {
            v28 := GenS4{F0: 0, F1: -2147483648}
            _ = v28
        }
    case "":
        v1 -= (v6 ^ uint16(v16))
    }
    for fi29 := 0; fi29 < 14; fi29++ {
        for fi30 := 0; fi30 < 7; fi30++ {
            if (!false) {
                fmt.Println(v3, fi30, v2)
                var v31 string = "\n\t"
                v31 = (((v31 + v31) + v31) + "a")
                _ = v31
            } else {
                v5["\n\t"] = "foo bar"
            }
            var v32 uint = 578
            switch "\n\t" {
            case "\n\t":
                var v33 float32 = float32(1.5)
                fmt.Println(v1)
                _ = v33
            case "foo bar":
                v20 = 42
            default:
                var v34 uint32 = 401
                _ = v34
            }
            if (((fi29 % (((fi29 + 32767)) | 1)) < (-(fi29))) || (!true)) {
                if false {
                    var v35 uint = (v16 >> ((v16) & 63))
                    _ = v35
                } else {
                    var v36 uint64 = v20
                    _ = v36
                }
                delete(v5, "\n\t")
            }
            _ = fi30
            _ = v32
        }
        switch "\n\t" {
        case "foo bar":
            v37 := GenS4{F0: -1, F1: 200}
            _ = v37
        case "":
            var v38 float64 = float64(1e10)
            _ = v38
        }
        _ = fi29
    }
    var v39 int64 = 9223372036854775807
    var v40 float32 = (0.1 + (0.1 - 3.14159))
    v3 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
