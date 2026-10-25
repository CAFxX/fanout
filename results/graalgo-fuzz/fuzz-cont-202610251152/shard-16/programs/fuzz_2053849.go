package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint8; F1 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 int, p1 int16) int64 {
    if (true && true) {
        if false {
            if false {
                for fi5 := 0; fi5 < 17; fi5++ {
                    v6 := GenS1{F0: 65535, F1: (p0 << ((374) & 63))}
                    fmt.Println(p0)
                    fmt.Println(p1, p0)
                    _ = fi5
                    _ = v6
                }
                if (789 >= 1) {
                    switch "\n\t" {
                    case "":
                        fmt.Println(p0, p0)
                    default:
                        p0 -= p0
                    }
                    v7 := []int{(p0 + p0)}
                    switch "" {
                    case "a":
                        fmt.Println(p0)
                    case "hello":
                        for fi8 := 0; fi8 < 20; fi8++ {
                            fmt.Println(p1)
                            _ = fi8
                        }
                    default:
                        var v9 int8 = 0
                        _ = v9
                    }
                    _ = v7
                }
                v10 := GenS1{F0: 4294967295, F1: 4294967295}
                _ = v10
            } else {
                p1 = p1
            }
            fmt.Println(p1)
        } else {
            var v11 int8 = 0
            v12 := GenS1{F0: 691, F1: (p0 | 464)}
            _ = v11
            _ = v12
        }
        if (-173 != 726) {
            switch 32767 {
            case 0:
                var v13 int32 = -1
                _ = v13
            case -2147483648:
                var v14 int16 = p1
                _ = v14
            }
            var v15 int = -629
            _ = v15
        }
    }
    var v16 int8 = -128
    _ = p0
    _ = p1
    _ = v16
    return 9223372036854775807
}

func main() {
    var v1 int64 = 1
    var v2 uint16 = 741
    var v3 int8 = 127
    var v4 bool = (true || false)
    fmt.Println(v4, v4, v4)
    v17 := GenS1{F0: 4294967295, F1: -141}
    for fi18 := 0; fi18 < 8; fi18++ {
        v4 = false
        v19 := GenS2{F0: 109, F1: ""}
        var v20 uint16 = (v2 & (^(v2)))
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := []int{}
    v22 := GenS1{F0: 42, F1: -171}
    fmt.Println(v4, v2, v2)
    switch "hello" {
    case "hello":
        v2 = (v2 & (v2 << ((0) & 63)))
    default:
        v23 := GenS3{F0: (v1 ^ -9223372036854775808)}
        _ = v23
    }
    v2 -= 708
    if (128 == 127) {
        switch -823 {
        case -1:
            v3 -= (v3 >> ((0) & 63))
        }
        var v24 uint = 282
        _ = v24
    } else {
        var v25 uint32 = 0
        fmt.Println(v25, v3)
        _ = v25
    }
    _ = len(v21)
    if (2 < 1) {
        v1 = genFn0(128, 1)
        v3 += -128
        if (18446744073709551615 <= 1) {
            v26 := GenS2{F0: 2, F1: "hello"}
            _ = v26
        } else {
            _ = len(v21)
        }
    } else {
        if (!true) {
            var v27 float32 = (-2.25 - 0.0)
            var v28 bool = (!(!(true && ("\n\t" <= ""))))
            _ = v27
            _ = v28
        }
    }
    var v29 int16 = 32767
    fmt.Println(v2, v2)
    for fi30 := 0; fi30 < 12; fi30++ {
        v31 := GenS3{F0: 375}
        _ = fi30
        _ = v31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v29
    _ = v3
    _ = v4
}
