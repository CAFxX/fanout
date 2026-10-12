package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 int64; F2 uint16 }
func genFn0(p0 int) int8 {
    p0 = p0
    for fi4 := 0; fi4 < 20; fi4++ {
        p0 = p0
        _ = fi4
    }
    p0 += (p0 / ((-2147483648) | 1))
    if false {
        if true {
            fmt.Println(p0)
            for fi5 := 0; fi5 < 12; fi5++ {
                p0 = p0
                p0 -= (p0 - 127)
                _ = fi5
            }
            switch 32767 {
            case -9223372036854775808:
                fmt.Println(p0)
            default:
                var v6 uint64 = 4294967295
                _ = v6
            }
        }
    }
    v7 := []int{32768, p0, (p0 / ((p0) | 1))}
    _ = p0
    _ = v7
    return -1
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 string = "a"
    var v3 uint16 = 385
    v8 := GenS1{F0: -659}
    var v9 uint16 = (v3 - (v3 % ((0) | 1)))
    v10 := []int{42, 422, 88, -556}
    var v11 int64 = 42
    if (float32(-2.25) != 1.5) {
        var v12 int8 = -104
        fmt.Println(v9, v12)
        if false {
            v13 := GenS1{F0: 946}
            _ = v13
        } else {
            v14 := GenS1{F0: v11}
            var v15 uint8 = 164
            _ = v14
            _ = v15
        }
        _ = v12
    } else {
        if (((v2 + (v2 + "a")) + "") >= v2) {
            var v16 int8 = 100
            _ = v16
        } else {
            v17 := GenS1{F0: (v11 >> ((65535) & 63))}
            if (103 <= 128) {
                var v18 int32 = 2147483647
                v19 := GenS1{F0: v11}
                v2 = (((v2 + v2) + ("hello" + v2)) + v2)
                _ = v18
                _ = v19
            }
            v11 = 0
            _ = v17
        }
    }
    fmt.Println(v3, v2)
    v11 = (v11 & int64(v9))
    v20 := GenS1{F0: v11}
    var v21 bool = (false && (v3 == v9))
    switch 128 {
    case 2147483647:
        if true {
            if (!v21) {
                switch -316 {
                case 2147483647:
                    var v22 uint = 867
                    _ = v22
                case 9223372036854775807:
                    v23 := GenS2{F0: genFn0(-1)}
                    _ = v23
                case 278:
                    v11 -= -1
                default:
                    switch v2 {
                    case "":
                        v2 = "foo bar"
                    case "\n\t":
                        if v21 {
                            v24 := GenS1{F0: 42}
                            _ = v24
                        } else {
                            v25 := GenS3{F0: 533, F1: v11, F2: (v9 * v9)}
                            if (v2 != "\n\t") {
                                var v26 int64 = (-(v11))
                                _ = v26
                            }
                            _ = v25
                        }
                    case "hello":
                        _ = len(v1)
                    }
                }
                v2 = ("foo bar" + "hello")
                v3 += v9
            }
            var v27 int8 = 127
            v21 = v21
            _ = v27
        } else {
            var v28 bool = ((982 == 798) || v21)
            _ = v28
        }
    case -74:
        var v29 uint16 = 205
        _ = v29
    }
    var v30 int32 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v30
    _ = v8
    _ = v9
}
