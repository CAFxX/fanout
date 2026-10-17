package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 string; F1 uint; F2 uint64 }
func genFn0() string {
    var v7 uint64 = 412
    var v8 uint16 = 2
    _ = v7
    _ = v8
    return "foo bar"
}
func genFn1(p0 uint8, p1 uint16) int32 {
    p1 += (p1 | 255)
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 2147483647
}

func main() {
    var v1 uint16 = 202
    var v2 float64 = (float64((((1e10 - 1e10) * float32(float64((1.5 + 1e10)))) / 0.1)) - 1.5)
    var v3 uint64 = 405
    v4 := make(map[uint32]bool)
    var v5 int8 = 1
    v6 := []int{128, -213}
    for fi9 := 0; fi9 < 14; fi9++ {
        var v10 uint16 = (v1 & (v1 % ((v1) | 1)))
        _ = fi9
        _ = v10
    }
    switch 826 {
    case 323:
        var v11 uint8 = 188
        _ = v11
    case -9223372036854775808:
        switch -847 {
        case 1:
            var v12 float64 = v2
            _ = v12
        case -167:
            v13 := GenS1{F0: "a", F1: v3}
            _ = v13
        case 42:
            fmt.Println(v2, v1, v1)
        }
    default:
        var v14 int16 = 1000
        _ = v14
    }
    _ = len(v4)
    if (79 > 295) {
        if false {
            if false {
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint64 = (v3 % ((v3) | 1))
                    var v17 uint8 = uint8(v16)
                    v18 := GenS1{F0: ("a" + ("hello" + "")), F1: (v3 / ((uint64(v1)) | 1))}
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            } else {
                fmt.Println(v2, v3)
            }
            v4[65535] = true
        } else {
            for fi19 := 0; fi19 < 5; fi19++ {
                v5 = 42
                if false {
                    v4[2] = (42 > 174)
                    fmt.Println(v2)
                } else {
                    var v20 float64 = 0.1
                    v2 -= v20
                    _ = v20
                }
                _ = fi19
            }
        }
    }
    if true {
        if false {
            if ((3.14159 * 3.14159) <= (1e10 - ((0.1 * 0.1) - float32(1e10)))) {
                for fi21 := 0; fi21 < 14; fi21++ {
                    v1 += 0
                    _ = fi21
                }
                v5 = (v5 | (v5 % ((v5) | 1)))
            }
        } else {
            var v22 uint16 = v1
            delete(v4, 2147483647)
            v23 := GenS2{F0: 65535}
            _ = v22
            _ = v23
        }
        var v24 uint16 = v1
        v4[2147483647] = (!(1e10 != 0.1))
        _ = v24
    } else {
        for fi25 := 0; fi25 < 18; fi25++ {
            var v26 string = "\n\t"
            var v27 bool = false
            _ = fi25
            _ = v26
            _ = v27
        }
        for fi28 := 0; fi28 < 12; fi28++ {
            switch "" {
            case "hello":
                delete(v4, 313)
            case "":
                v29 := v4[271]
                _ = v29
            case "\n\t":
                v30 := GenS2{F0: v1}
                _ = v30
            }
            if false {
                v31 := GenS2{F0: v1}
                _ = v31
            } else {
                fmt.Println(v3)
            }
            _ = fi28
        }
        v5 = v5
    }
    for fi32 := 0; fi32 < 5; fi32++ {
        if false {
            var v33 uint64 = (v3 / (((v3 | 18446744073709551615)) | 1))
            v1 = 1
            var v34 uint64 = (v3 % ((1) | 1))
            _ = v33
            _ = v34
        }
        _ = fi32
    }
    if (false && false) {
        switch "\n\t" {
        case "\n\t":
            var v35 uint8 = 2
            _ = v35
        case "foo bar":
            v5 += (v5 + 0)
        case "hello":
            if (!true) {
                v5 += (-(v5))
            }
        }
    }
    v1 += 580
    v36 := GenS3{F0: "", F1: 389, F2: (v3 << ((65535) & 63))}
    for fi37 := 0; fi37 < 2; fi37++ {
        v2 -= v2
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
