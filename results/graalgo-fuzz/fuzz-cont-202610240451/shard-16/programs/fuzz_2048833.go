package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32 }
type GenS2 struct { F0 string; F1 int }
func genFn0(p0 uint32, p1 uint16) uint {
    p1 -= (^(p1))
    if (((((3.14159 + 3.14159) - ((float64(0.0) * 3.14159) + -2.25)) + 1.5) - 1e10) <= 0.1) {
        if (1.5 >= 0.0) {
            var v5 bool = true
            for fi6 := 0; fi6 < 15; fi6++ {
                for fi7 := 0; fi7 < 18; fi7++ {
                    v8 := []int{fi6, fi7, (fi7 % ((128) | 1)), fi6}
                    _ = fi7
                    _ = v8
                }
                _ = fi6
            }
            for fi9 := 0; fi9 < 18; fi9++ {
                p0 -= (p0 >> ((18446744073709551615) & 63))
                if (0 > 42) {
                    v10 := []int{}
                    v11 := GenS1{F0: "\n\t", F1: int32(p0)}
                    if len(v10) > 0 {
                        v10[0] = (fi9 % ((1) | 1))
                    }
                    _ = v10
                    _ = v11
                } else {
                    v12 := []int{fi9, fi9}
                    _ = v12
                }
                var v13 int8 = 126
                _ = fi9
                _ = v13
            }
            _ = v5
        }
        if false {
            switch 1 {
            case 9223372036854775807:
                p1 = 2
            case -2147483648:
                var v14 int16 = -1
                _ = v14
            }
            if ((1 == 1) || true) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    fmt.Println(p0)
                    _ = fi15
                }
                v16 := GenS2{F0: (("foo bar" + "\n\t") + ("a" + ("\n\t" + "foo bar"))), F1: 128}
                _ = v16
            } else {
                v17 := []int{-530, 32768}
                switch 42 {
                case 32768:
                    var v18 int16 = -32768
                    _ = v18
                default:
                    p1 -= (p1 / (((p1 * (p1 + p1))) | 1))
                }
                _ = v17
            }
        } else {
            v19 := []int{2147483647, -217, -9223372036854775808, -9223372036854775808}
            switch 32767 {
            case 1:
                var v20 uint32 = (^(p0))
                _ = v20
            default:
                if len(v19) > 0 {
                    v19[0] = -1
                }
            }
            if (-2147483648 <= 2147483647) {
                fmt.Println(p1, p0)
                var v21 string = ("hello" + ("hello" + ("\n\t" + (("\n\t" + "a") + ("hello" + "hello")))))
                _ = v21
            }
            _ = v19
        }
    } else {
        fmt.Println(p1, p1)
        var v22 string = "hello"
        var v23 int64 = 188
        _ = v22
        _ = v23
    }
    _ = p0
    _ = p1
    return 255
}

func main() {
    var v1 bool = true
    var v2 uint16 = 1
    var v3 int16 = -1
    var v4 uint = 4294967295
    v24 := GenS2{F0: "hello", F1: -1}
    var v25 float32 = 1.5
    v1 = (!true)
    v26 := []int{0, 2147483647, -2147483648}
    var v27 float32 = v25
    var v28 int8 = 42
    v29 := GenS1{F0: "hello", F1: 483}
    var v30 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
