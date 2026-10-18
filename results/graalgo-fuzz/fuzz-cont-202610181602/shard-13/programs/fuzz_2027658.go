package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint }
func genFn0(p0 uint, p1 uint32, p2 int32) string {
    p1 = (^(p1))
    var v5 uint8 = 0
    v5 += (v5 >> (((p0 >> (((p0 << ((p0) & 63))) & 63))) & 63))
    var v6 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (("a" + (("" + "") + "foo bar")) + "\n\t")
}

func main() {
    v1 := []string{}
    var v2 int8 = 42
    var v3 float32 = 0.0
    v4 := make(map[int]bool)
    v2 = 1
    var v7 uint64 = 158
    v8 := GenS1{F0: -244, F1: -353}
    switch 127 {
    case 127:
        var v9 string = "hello"
        _ = v9
    case -2147483648:
        v10 := GenS2{F0: v2}
        _ = v10
    case 42:
        var v11 int32 = 0
        _ = v11
    }
    switch 10 {
    case 32767:
        _ = len(v1)
    case 128:
        v2 = v2
    default:
        v12 := GenS3{F0: 2}
        _ = v12
    }
    var v13 int32 = -1
    _ = len(v4)
    fmt.Println(v13, v3)
    var v14 int32 = v13
    if false {
        v15 := GenS1{F0: 9223372036854775807, F1: -32768}
        _ = v15
    } else {
        for fi16 := 0; fi16 < 3; fi16++ {
            fmt.Println(v3)
            _ = fi16
        }
        if ((v2 - 0) <= 12) {
            for fi17 := 0; fi17 < 13; fi17++ {
                var v18 uint16 = 199
                _ = fi17
                _ = v18
            }
            if (560 > 18446744073709551615) {
                switch -9223372036854775808 {
                case 901:
                    var v19 uint64 = 710
                    _ = v19
                case -2147483648:
                    if len(v1) > 0 {
                        v1[0] = ("a" + (((("a" + (("" + "a") + "foo bar")) + "\n\t") + ("foo bar" + "a")) + ("a" + "foo bar")))
                    }
                case 4294967295:
                    fmt.Println(v14)
                }
            } else {
                var v20 string = "\n\t"
                delete(v4, 128)
                if true {
                    v13 = (v13 | -867)
                    v2 = v2
                    var v21 float32 = ((0.0 / (float32(0.1) * v3)) / v3)
                    _ = v21
                }
                _ = v20
            }
            v7 -= (v7 + (v7 / (((^(v7))) | 1)))
        }
    }
    switch 127 {
    case 2147483647:
        v1 = append(v1, "")
    case 127:
        if (1 >= 0) {
            var v22 uint16 = 0
            _ = v22
        }
    case 42:
        v23 := GenS3{F0: 1}
        _ = v23
    default:
        for fi24 := 0; fi24 < 12; fi24++ {
            var v25 uint = 65535
            v3 = (v3 * ((0.1 - float32(1e10)) * -2.25))
            _ = fi24
            _ = v25
        }
    }
    fmt.Println(v7)
    v26 := v4[4294967295]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
