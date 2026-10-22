package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int16; F1 int8 }
func genFn0(p0 int) bool {
    p0 = (-(p0))
    for fi6 := 0; fi6 < 13; fi6++ {
        if (42 == 4294967295) {
            p0 -= (p0 & p0)
            switch p0 {
            case -9223372036854775808:
                var v7 uint = 0
                _ = v7
            case -2147483648:
                fmt.Println(p0, p0)
            default:
                var v8 int64 = -661
                _ = v8
            }
            p0 = (^(p0))
        }
        _ = fi6
    }
    v9 := GenS1{F0: 1, F1: 127}
    v10 := GenS2{F0: -582}
    p0 = 206
    _ = p0
    _ = v10
    _ = v9
    return false
}
func genFn1() uint32 {
    var v11 uint8 = 128
    v12 := []int{42}
    if (42 >= 1) {
        if (("" + "foo bar") < "") {
            if true {
                v11 += 255
                var v13 uint = 842
                _ = v13
            }
            if (false && (0 <= 501)) {
                var v14 int = 2147483647
                v14 = v14
                _ = v14
            } else {
                _ = len(v12)
            }
            var v15 int16 = -32768
            _ = v15
        }
        var v16 uint64 = 2
        _ = v16
    }
    var v17 uint16 = 65535
    _ = v11
    _ = v12
    _ = v17
    return 4294967295
}

func main() {
    var v1 int16 = 32767
    v2 := []uint32{2147483647}
    var v3 int8 = -128
    var v4 uint = 2
    v5 := make(map[string]string)
    switch "" {
    case "\n\t":
        var v18 uint16 = 526
        _ = v18
    case "":
        fmt.Println(v4, v1, v3)
    }
    switch "" {
    case "\n\t":
        v3 = v3
    case "":
        var v19 int = -971
        _ = v19
    default:
        var v20 string = "\n\t"
        _ = v20
    }
    v21 := GenS3{F0: -1, F1: v3}
    v1 = (v1 / ((v1) | 1))
    fmt.Println(v1, v1)
    v1 += v1
    fmt.Println(v1, v3, v1)
    var v22 int = -9223372036854775808
    if (("\n\t" != ((("\n\t" + "foo bar") + "a") + "\n\t")) && true) {
        if false {
            var v23 uint32 = 42
            if true {
                v24 := GenS4{F0: (v1 - v1), F1: (v3 * (v3 - (v3 | v3)))}
                var v25 string = "hello"
                v26 := GenS1{F0: 2, F1: 42}
                _ = v24
                _ = v25
                _ = v26
            } else {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v1)
                case "":
                    var v27 bool = (true || true)
                    _ = v27
                }
            }
            var v28 int16 = v1
            _ = v23
            _ = v28
        } else {
            if (-804 <= 42) {
                fmt.Println(v4)
                var v29 int64 = -1
                switch 128 {
                case 4294967295:
                    fmt.Println(v3, v4, v1)
                default:
                    if (((v1 >> ((1) & 63)) > (v1 / ((-32768) | 1))) || false) {
                        for fi30 := 0; fi30 < 3; fi30++ {
                            var v31 int8 = int8(v4)
                            for fi32 := 0; fi32 < 20; fi32++ {
                                delete(v5, ("" + "hello"))
                                v1 += (-(v1))
                                v4 += 65535
                                _ = fi32
                            }
                            var v33 float64 = 3.14159
                            _ = fi30
                            _ = v31
                            _ = v33
                        }
                    }
                }
                _ = v29
            }
            v4 += v4
            v1 = v1
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v1 = (v1 | 1000)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
