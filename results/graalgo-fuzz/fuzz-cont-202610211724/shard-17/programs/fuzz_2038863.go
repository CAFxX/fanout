package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint64 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int32, p1 uint8) uint32 {
    v5 := []int{0, 128, -2147483648}
    v5 = append(v5, 1)
    switch 907 {
    case 0:
        var v6 uint8 = p1
        _ = v6
    default:
        switch 42 {
        case -2147483648:
            p0 -= p0
        }
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v5
    return 551
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []string{"a", (("foo bar" + ("hello" + "hello")) + ""), (("foo bar" + "") + ((((("hello" + "hello") + ("" + "a")) + "\n\t") + (("" + "a") + "foo bar")) + "a"))}
    var v3 string = "\n\t"
    var v4 uint32 = 1
    var v7 int = 42
    fmt.Println(v7, v7)
    var v8 uint = 42
    switch v7 {
    case 42:
        v9 := GenS1{F0: -131, F1: 0}
        _ = v9
    default:
        if len(v2) > 0 {
            v2[0] = (((v3 + "foo bar") + "\n\t") + "foo bar")
        }
    }
    var v10 int64 = -9223372036854775808
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 uint64 = 1
        v12 = (v12 / (((v12 / (((v12 * v12)) | 1))) | 1))
        _ = fi11
        _ = v12
    }
    fmt.Println(v3)
    var v13 int8 = -98
    switch v7 {
    case 2147483647:
        if (v4 != 744) {
            var v14 int8 = -128
            v3 = v3
            _ = v14
        } else {
            switch ((((v3 + "hello") + v3) + v3) + "hello") {
            case "":
                if false {
                    for fi15 := 0; fi15 < 11; fi15++ {
                        fmt.Println(v13, v13)
                        v10 += v10
                        _ = fi15
                    }
                }
            case "hello":
                var v16 float32 = float32(1e10)
                _ = v16
            case "\n\t":
                var v17 int = (v7 * (v7 << (((v8 + v8)) & 63)))
                _ = v17
            }
        }
    default:
        if (611 != 32767) {
            if false {
                fmt.Println(v13)
            } else {
                v18 := GenS2{F0: v8, F1: 575, F2: 18446744073709551615}
                if false {
                    v3 = (v3 + v3)
                    if (!true) {
                        var v19 uint32 = genFn0(-998, 2)
                        v10 = v10
                        _ = v19
                    }
                } else {
                    v20 := GenS1{F0: -9223372036854775808, F1: v10}
                    _ = v20
                }
                _ = v18
            }
            v13 += -1
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                var v22 int8 = v13
                var v23 uint32 = (v4 ^ (v4 << ((47) & 63)))
                var v24 int8 = v13
                _ = fi21
                _ = v22
                _ = v23
                _ = v24
            }
            var v25 int64 = (v10 << (((v8 ^ (v8 << ((v8) & 63)))) & 63))
            _ = len(v2)
            _ = v25
        }
    }
    fmt.Println(v7, v8)
    v26 := GenS3{F0: v4}
    for fi27 := 0; fi27 < 6; fi27++ {
        if (1 <= -973) {
            var v28 uint8 = 1
            switch v3 {
            case "":
                var v29 uint64 = uint64(v28)
                _ = v29
            case "foo bar":
                var v30 uint32 = (v4 * (^(v4)))
                _ = v30
            default:
                var v31 string = v3
                _ = v31
            }
            _ = v28
        }
        _ = fi27
    }
    var v32 int8 = v13
    v32 = (v32 % ((1) | 1))
    if true {
        var v33 int16 = 1000
        fmt.Println(v3, v13, v3)
        _ = v33
    } else {
        var v34 int16 = 0
        _ = v34
    }
    v4 -= v4
    delete(v1, v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v32
    _ = v4
    _ = v7
    _ = v8
}
