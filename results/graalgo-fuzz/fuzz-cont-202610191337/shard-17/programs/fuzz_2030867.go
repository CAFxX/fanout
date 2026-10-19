package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int; F1 string; F2 uint32 }
func genFn0(p0 uint8) bool {
    var v5 bool = false
    p0 -= 127
    v6 := []int{9223372036854775807, 32768}
    _ = p0
    _ = v5
    _ = v6
    return true
}

func main() {
    var v1 int = 1
    var v2 float64 = -2.25
    var v3 uint = 4294967295
    var v4 int32 = 2147483647
    var v7 int16 = -32768
    fmt.Println(v4, v1)
    var v8 uint32 = 2
    v9 := []int{v1}
    _ = len(v9)
    if (((v1 % (((v1 * (v1 >> ((v3) & 63)))) | 1)) == (v1 ^ v1)) && (false && (100 < int8(v3)))) {
        if (100 <= 1) {
            switch (^(v1)) {
            case -655:
                v7 = (v7 ^ (v7 & v7))
            default:
                v10 := GenS1{F0: 1, F1: v7}
                _ = v10
            }
        } else {
            v8 -= 1
            v8 = 278
            var v11 uint16 = 633
            _ = v11
        }
    } else {
        var v12 uint16 = 424
        for fi13 := 0; fi13 < 13; fi13++ {
            switch "\n\t" {
            case "foo bar":
                v4 -= (^(v4))
            case "a":
                if (((-983 != 42) || true) || (!(true || (2147483648 != (v8 - v8))))) {
                    v8 += v8
                    var v14 int64 = 936
                    _ = v14
                }
            }
            v9 = append(v9, (fi13 ^ (v1 / ((127) | 1))))
            _ = fi13
        }
        _ = v12
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        if (v8 == v8) {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 uint64 = 47
                var v18 string = "foo bar"
                v19 := GenS1{F0: 195, F1: v7}
                _ = fi16
                _ = v17
                _ = v18
                _ = v19
            }
        }
        v20 := GenS1{F0: 2, F1: (v7 % (((v7 / ((579) | 1))) | 1))}
        _ = fi15
        _ = v20
    }
    switch (v1 >> ((v3) & 63)) {
    case -9223372036854775808:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    var v21 float64 = ((v2 * -2.25) * 0.1)
    if (false && false) {
        if len(v9) > 0 {
            v9[0] = (v1 - (v1 * -30))
        }
        if true {
            v4 = v4
        } else {
            for fi22 := 0; fi22 < 14; fi22++ {
                v7 += (v7 + v7)
                v3 = v3
                _ = fi22
            }
            var v23 uint16 = 866
            var v24 uint = (v3 % (((v3 + 255)) | 1))
            _ = v23
            _ = v24
        }
    } else {
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
        v25 := GenS1{F0: 233, F1: -32768}
        _ = v25
    }
    if ((("hello" + ("foo bar" + "")) <= (((("hello" + "foo bar") + "hello") + "") + "")) || true) {
        if ((v2 < (-2.25 - v2)) && (-9223372036854775808 != 0)) {
            v3 += (v3 - (v3 + (v3 - 18446744073709551615)))
            v26 := GenS2{F0: (v4 / (((v4 + (v4 >> ((v3) & 63)))) | 1)), F1: 32767}
            var v27 uint32 = v8
            _ = v26
            _ = v27
        } else {
            v28 := GenS3{F0: v1, F1: "hello", F2: (v8 % (((v8 % ((v8) | 1))) | 1))}
            var v29 int8 = -4
            _ = v28
            _ = v29
        }
    }
    var v30 float64 = 3.14159
    var v31 uint8 = 127
    v9 = append(v9, (v1 | (v1 - (v1 / (((v1 * (-(v1)))) | 1)))))
    v30 = float64(-2.25)
    v32 := GenS3{F0: (v1 - (v1 | v1)), F1: ("foo bar" + "a"), F2: v8}
    var v33 uint64 = 732
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
