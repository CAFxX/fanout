package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 bool }
type GenS2 struct { F0 string; F1 uint8 }
type GenS3 struct { F0 string }
func genFn0(p0 uint64) int64 {
    fmt.Println(p0)
    v4 := []int{}
    var v5 int32 = -28
    if true {
        v6 := GenS1{F0: (^(p0)), F1: 113, F2: (!(!true))}
        p0 = (p0 + (p0 * p0))
        _ = v6
    }
    _ = p0
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 bool = (!(((3.14159 - ((-2.25 - (0.0 / 3.14159)) * (1.5 / 3.14159))) >= 3.14159) && false))
    var v3 int16 = -490
    var v7 int = 128
    v8 := []int{v7}
    fmt.Println(v2, v1)
    if v2 {
        switch "foo bar" {
        case "\n\t":
            v9 := GenS2{F0: ("a" + "\n\t"), F1: 112}
            _ = v9
        default:
            if (985 >= 797) {
                v2 = true
                if (((-2.25 != (0.0 / 1.0)) && v2) && ((1 == 271) && (!((v1 % ((v1) | 1)) <= (^(v1)))))) {
                    fmt.Println(v7, v1)
                    if false {
                        if (-32768 == 42) {
                            switch 54 {
                            case 9223372036854775807:
                                var v10 bool = v2
                                _ = v10
                            }
                            if (!(v3 > (v3 ^ v3))) {
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    var v12 int16 = v3
                                    var v13 uint16 = 2
                                    _ = fi11
                                    _ = v12
                                    _ = v13
                                }
                                v8 = append(v8, -2147483648)
                                var v14 int32 = 703
                                _ = v14
                            } else {
                                v1 = 42
                            }
                        } else {
                            var v15 int64 = 9223372036854775807
                            var v16 int64 = (v1 << ((988) & 63))
                            _ = v15
                            _ = v16
                        }
                        v1 = (v1 << ((96) & 63))
                    }
                    v17 := GenS3{F0: ""}
                    _ = v17
                }
            } else {
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                var v18 uint64 = 42
                v3 -= 1000
                _ = v18
            }
        }
        v3 = 42
        var v19 int32 = 0
        _ = v19
    } else {
        fmt.Println(v2)
    }
    var v20 uint = 2
    var v21 uint64 = 4294967295
    _ = len(v8)
    if ((!(v7 > v7)) && v2) {
        if len(v8) > 0 {
            v8[0] = 127
        }
        fmt.Println(v1, v2)
    } else {
        switch "a" {
        case "":
            v3 = v3
        case "foo bar":
            v3 = (v3 | (v3 & (v3 - v3)))
        case "a":
            var v22 float64 = 0.0
            _ = v22
        }
        fmt.Println(v21, v3, v2)
        for fi23 := 0; fi23 < 3; fi23++ {
            v21 = 2
            _ = fi23
        }
    }
    var v24 uint16 = 32767
    var v25 uint = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v25
    _ = v3
    _ = v7
    _ = v8
}
