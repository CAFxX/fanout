package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int; F1 int16; F2 int16 }
func genFn0(p0 uint16) int32 {
    if false {
        fmt.Println(p0)
        v6 := GenS1{F0: 1, F1: 220, F2: p0}
        _ = v6
    } else {
        var v7 float64 = 0.0
        fmt.Println(p0, p0)
        _ = v7
    }
    var v8 uint32 = 59
    v8 -= v8
    _ = p0
    _ = v8
    return 42
}

func main() {
    var v1 uint64 = 21
    v2 := []uint32{862, 42, 1, 295}
    var v3 int = -2147483648
    v4 := []string{"", "foo bar", ((("hello" + "a") + "\n\t") + "\n\t"), "\n\t"}
    var v5 uint64 = 145
    for fi9 := 0; fi9 < 19; fi9++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        for fi10 := 0; fi10 < 16; fi10++ {
            var v11 int64 = 42
            var v12 int = (fi9 / (((-(fi9))) | 1))
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi9
    }
    var v13 bool = ((false && (false || ("foo bar" != ""))) || (true && (-32768 > 32767)))
    for fi14 := 0; fi14 < 8; fi14++ {
        v13 = false
        for fi15 := 0; fi15 < 13; fi15++ {
            if (!(!v13)) {
                if ((128 < 42) && (v1 >= (v1 ^ 18446744073709551615))) {
                    v16 := GenS1{F0: v1, F1: 0, F2: 355}
                    fmt.Println(v1, v3)
                    _ = v16
                } else {
                    switch v3 {
                    case -9223372036854775808:
                        v3 -= v3
                    case 4294967295:
                        switch "foo bar" {
                        case "\n\t":
                            fmt.Println(v1)
                        case "hello":
                            v13 = true
                        case "a":
                            v17 := GenS2{F0: fi14, F1: 1, F2: 1}
                            _ = v17
                        default:
                            v1 -= (v1 + v5)
                        }
                    case 127:
                        fmt.Println(v1, v13, v13)
                    }
                    v3 = 42
                    var v18 int16 = 1
                    _ = v18
                }
                fmt.Println(v13, v13)
                var v19 float64 = -2.25
                _ = v19
            }
            _ = fi15
        }
        fmt.Println(v5)
        _ = fi14
    }
    v5 = 2
    for fi20 := 0; fi20 < 13; fi20++ {
        for fi21 := 0; fi21 < 12; fi21++ {
            var v22 uint = 130
            if v13 {
                fmt.Println(v22)
            } else {
                v13 = (-9223372036854775808 > -221)
                var v23 int = 223
                if false {
                    v1 -= (v1 - v1)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                v22 += v22
                _ = v23
            }
            var v24 int8 = 127
            _ = fi21
            _ = v22
            _ = v24
        }
        for fi25 := 0; fi25 < 3; fi25++ {
            fmt.Println(fi20)
            var v26 int = (fi25 >> ((4294967295) & 63))
            var v27 uint16 = 0
            _ = fi25
            _ = v26
            _ = v27
        }
        _ = fi20
    }
    if v13 {
        _ = len(v2)
    } else {
        for fi28 := 0; fi28 < 7; fi28++ {
            v5 -= 4294967295
            var v29 uint16 = 65535
            _ = fi28
            _ = v29
        }
        fmt.Println(v13, v3)
        _ = len(v4)
    }
    v13 = false
    v5 = (v1 >> ((255) & 63))
    var v30 string = ""
    fmt.Println(v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
