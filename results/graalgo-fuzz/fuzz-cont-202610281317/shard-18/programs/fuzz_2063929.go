package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16; F2 bool }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int32; F1 uint32; F2 string }
type GenS4 struct { F0 bool; F1 string; F2 uint8 }
func genFn0(p0 int8, p1 int64, p2 uint16) string {
    v6 := []int{2147483647, 4294967295, 32767}
    var v7 int16 = 42
    for fi8 := 0; fi8 < 17; fi8++ {
        if true {
            switch (fi8 - (^(fi8))) {
            case -446:
                fmt.Println(p2, p0, p0)
            case 9223372036854775807:
                var v9 uint64 = 180
                _ = v9
            case -882:
                switch "foo bar" {
                case "foo bar":
                    v6 = append(v6, fi8)
                case "hello":
                    fmt.Println(fi8, p2, fi8)
                case "":
                    v10 := GenS1{F0: (p2 >> ((769) & 63)), F1: (v7 / ((32767) | 1)), F2: false}
                    _ = v10
                default:
                    var v11 uint8 = 128
                    _ = v11
                }
            }
            fmt.Println(fi8, p1, p0)
            p2 = p2
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "foo bar"
}

func main() {
    v1 := []int{}
    var v2 bool = (42 >= 18446744073709551615)
    var v3 bool = v2
    v4 := []int{124, -353, 32768}
    v5 := []uint32{2147483648}
    if ((!v3) && (18446744073709551615 >= 18446744073709551615)) {
        var v12 int8 = 100
        _ = v12
    }
    fmt.Println(v2, v2)
    var v13 int8 = 42
    var v14 uint8 = 128
    for fi15 := 0; fi15 < 16; fi15++ {
        v16 := GenS2{F0: 58, F1: (true && (v3 && (!(4294967295 > 255))))}
        _ = fi15
        _ = v16
    }
    v3 = false
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 int8 = v13
            v20 := GenS1{F0: 1, F1: -1, F2: v2}
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 int8 = (v13 & -48)
        if (0 >= 1) {
            if len(v1) > 0 {
                v1[0] = (fi17 & (fi17 + -2147483648))
            }
            var v22 int = fi17
            v23 := GenS2{F0: 42, F1: v3}
            _ = v22
            _ = v23
        } else {
            var v24 uint16 = 65535
            var v25 int16 = 0
            v26 := GenS2{F0: 90, F1: true}
            _ = v24
            _ = v25
            _ = v26
        }
        _ = fi17
        _ = v21
    }
    v14 += (^(v14))
    v27 := GenS3{F0: 2147483647, F1: 688, F2: "a"}
    v28 := GenS3{F0: -999, F1: 232, F2: ""}
    fmt.Println(v3)
    if (2147483647 >= -1) {
        var v29 uint8 = (^(v14))
        _ = v29
    } else {
        fmt.Println(v3, v14, v13)
        fmt.Println(v13)
        if v3 {
            v30 := GenS4{F0: (65535 > 100), F1: ((("hello" + "") + "") + "foo bar"), F2: 105}
            switch 433 {
            case 2147483647:
                fmt.Println(v13, v14, v13)
            }
            _ = v30
        }
    }
    fmt.Println(v2, v2, v14)
    v3 = v3
    if v2 {
        v14 -= v14
        v13 -= (v13 % ((v13) | 1))
        v31 := GenS1{F0: 438, F1: 1, F2: (v3 || v3)}
        _ = v31
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
