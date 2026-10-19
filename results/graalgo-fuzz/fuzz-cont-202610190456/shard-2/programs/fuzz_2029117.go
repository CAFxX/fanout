package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 int32; F1 uint32 }
type GenS3 struct { F0 int }
func genFn0(p0 uint) uint32 {
    p0 = p0
    var v7 int16 = 42
    v7 -= 1
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint8 = 18
    var v2 int64 = 328
    var v3 uint16 = 42
    v4 := []string{"a"}
    var v5 int8 = 127
    v6 := []uint32{42, 2, 276, 2}
    v1 += uint8(v3)
    v8 := GenS1{F0: "", F1: ""}
    var v9 bool = (-2.25 <= 1e10)
    v10 := GenS1{F0: "hello", F1: "foo bar"}
    if (18446744073709551615 <= 1) {
        v11 := GenS1{F0: ("" + ("hello" + "a")), F1: "\n\t"}
        _ = v11
    } else {
        v12 := GenS2{F0: -714, F1: 1}
        var v13 bool = true
        _ = v12
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        _ = len(v6)
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 string = ""
            if (2147483648 > 2147483648) {
                switch -2147483648 {
                case 300:
                    v17 := GenS1{F0: v16, F1: v16}
                    _ = v17
                case 128:
                    v6 = append(v6, 4294967295)
                case -9223372036854775808:
                    var v18 uint = 42
                    _ = v18
                default:
                    var v19 uint = 42
                    _ = v19
                }
                v20 := GenS1{F0: (("\n\t" + (v16 + v16)) + v16), F1: "foo bar"}
                fmt.Println(fi14, v2)
                _ = v20
            }
            v21 := GenS1{F0: v16, F1: (v16 + (v16 + v16))}
            _ = fi15
            _ = v16
            _ = v21
        }
        v3 -= 32767
        _ = fi14
    }
    if v9 {
        v22 := GenS3{F0: 32767}
        var v23 uint8 = 90
        _ = v22
        _ = v23
    }
    v24 := GenS1{F0: "", F1: "foo bar"}
    v5 += -128
    fmt.Println(v2, v9)
    if false {
        v25 := GenS2{F0: 2147483647, F1: 2147483647}
        _ = v25
    }
    var v26 uint8 = (v1 & (v1 | 206))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
