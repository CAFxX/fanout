package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32; F1 uint; F2 int32 }
func genFn0(p0 uint, p1 int32, p2 string) int8 {
    p2 = ("a" + p2)
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := GenS1{F0: 2, F1: "hello"}
        fmt.Println(p2, fi5, p0)
        var v7 uint8 = 171
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return int8(p0)
}

func main() {
    var v1 uint64 = 1
    var v2 uint16 = 0
    var v3 uint32 = 2147483648
    v4 := []int{128, 128, 128}
    v3 = v3
    switch "" {
    case "\n\t":
        v4 = append(v4, 4294967295)
    case "hello":
        var v8 int8 = 1
        _ = v8
    default:
        if len(v4) > 0 {
            v4[0] = -1
        }
    }
    fmt.Println(v3, v3, v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v3 += v3
    v2 -= (v2 + (v2 ^ (v2 << ((42) & 63))))
    v2 += v2
    var v9 int64 = 522
    fmt.Println(v1, v3)
    switch 1 {
    case 470:
        if (-878 <= 101) {
            if (42 <= 1) {
                for fi10 := 0; fi10 < 19; fi10++ {
                    if false {
                        var v11 int8 = 26
                        _ = v11
                    }
                    _ = fi10
                }
                v9 += 42
            } else {
                var v12 uint = 0
                _ = v12
            }
            var v13 uint64 = (^(v1))
            if (false || true) {
                switch 128 {
                case -9223372036854775808:
                    fmt.Println(v3, v13, v3)
                default:
                    if true {
                        v13 = (v13 / (((^(v1))) | 1))
                        var v14 string = "a"
                        fmt.Println(v14, v13)
                        _ = v14
                    } else {
                        fmt.Println(v13, v3, v13)
                        var v15 uint32 = 40
                        v16 := GenS1{F0: 4294967295, F1: ""}
                        _ = v15
                        _ = v16
                    }
                }
            } else {
                v3 -= 0
                switch 32768 {
                case 0:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                case 32767:
                    _ = len(v4)
                }
            }
            _ = v13
        }
    case 32767:
        if len(v4) > 0 {
            v4[0] = 192
        }
    case 4294967295:
        v17 := GenS2{F0: -2147483648, F1: 524, F2: 2147483647}
        _ = v17
    default:
        if false {
            var v18 uint16 = (^(v2))
            v19 := GenS1{F0: v3, F1: ""}
            _ = v18
            _ = v19
        } else {
            var v20 string = "\n\t"
            _ = v20
        }
    }
    if (42 < v1) {
        fmt.Println(v1, v2)
    }
    if true {
        v3 -= v3
    } else {
        if len(v4) > 0 {
            v4[0] = 128
        }
    }
    fmt.Println(v2, v9)
    var v21 int32 = 0
    v22 := GenS2{F0: 0, F1: 42, F2: (v21 * v21)}
    if (false || true) {
        var v23 uint16 = (v2 / ((v2) | 1))
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v9
}
