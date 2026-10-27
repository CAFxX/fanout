package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint; F2 int }

func main() {
    v1 := []uint32{703}
    var v2 uint8 = 42
    var v3 string = ("a" + ("hello" + "\n\t"))
    var v4 int8 = -128
    v5 := make(map[int]string)
    var v6 uint64 = 18446744073709551615
    if true {
        var v7 uint32 = 2
        _ = v7
    }
    if true {
        v1 = append(v1, 42)
    }
    v8 := GenS1{F0: (v2 - v2)}
    v3 = "foo bar"
    fmt.Println(v3, v4, v6)
    if true {
        if (!(!(1e10 >= 1e10))) {
            v2 += 128
        } else {
            var v9 uint32 = 216
            if (748 >= 127) {
                if (4294967295 == uint(v2)) {
                    if ((v6 - (v6 | (v6 & (v6 << ((959) & 63))))) < v6) {
                        v3 = ("hello" + v3)
                    }
                } else {
                    var v10 float32 = 3.14159
                    v4 -= (-(v4))
                    _ = v10
                }
            }
            _ = v9
        }
        fmt.Println(v4, v2, v2)
        for fi11 := 0; fi11 < 5; fi11++ {
            v6 = v6
            if (852 < 65535) {
                fmt.Println(v6, fi11, fi11)
                var v12 uint = 2
                _ = v12
            } else {
                fmt.Println(v4)
                var v13 string = "a"
                _ = v13
            }
            _ = fi11
        }
    } else {
        if (false && (!false)) {
            _ = len(v5)
            switch 2147483647 {
            case -1:
                v3 = v3
            case 1:
                v14 := GenS1{F0: v2}
                _ = v14
            case 4294967295:
                var v15 float64 = -2.25
                _ = v15
            default:
                switch 394 {
                case 1:
                    switch -490 {
                    case -624:
                        var v16 uint8 = v2
                        _ = v16
                    default:
                        v6 = v6
                    }
                default:
                    _ = len(v5)
                }
            }
            v17 := GenS2{F0: 0, F1: v3}
            _ = v17
        } else {
            fmt.Println(v4)
            v18 := GenS3{F0: false, F1: (^(v4)), F2: v2}
            v4 -= v4
            _ = v18
        }
        fmt.Println(v3, v4, v6)
        v6 -= (v6 & (v6 * (v6 / ((v6) | 1))))
    }
    v4 -= v4
    _ = len(v1)
    if true {
        if false {
            var v19 float64 = 0.0
            fmt.Println(v2, v4)
            _ = v19
        } else {
            switch (v3 + "a") {
            case "a":
                switch v3 {
                case "hello":
                    var v20 int = -2147483648
                    _ = v20
                case "\n\t":
                    var v21 bool = (!false)
                    _ = v21
                }
            case "\n\t":
                v2 -= (^(v2))
            case "foo bar":
                switch 127 {
                case -9223372036854775808:
                    fmt.Println(v6, v3, v6)
                case -543:
                    var v22 uint8 = 1
                    _ = v22
                default:
                    fmt.Println(v4, v4)
                }
            default:
                v6 = v6
            }
            fmt.Println(v2, v4, v3)
            v6 += (^(v6))
        }
        v23 := GenS4{F0: ("a" <= "\n\t"), F1: 312, F2: -576}
        _ = v23
    }
    var v24 uint8 = (v2 + (v2 >> ((255) & 63)))
    var v25 float64 = 0.1
    var v26 float64 = (v25 * v25)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
