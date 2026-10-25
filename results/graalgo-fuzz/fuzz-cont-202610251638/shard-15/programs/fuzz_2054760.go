package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int; F1 int }
func genFn0(p0 uint16, p1 int, p2 int64) bool {
    if (p2 == (p2 & 140)) {
        if true {
            var v7 uint = 255
            _ = v7
        } else {
            var v8 string = ""
            _ = v8
        }
    } else {
        p0 = 42
    }
    if false {
        v9 := []int{2147483647, int(p2), p1}
        _ = v9
    }
    var v10 int32 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (!true)
}
func genFn1(p0 uint, p1 uint, p2 uint64) uint32 {
    v11 := []int{0, -474, 127, 32767}
    switch "\n\t" {
    case "\n\t":
        switch "hello" {
        case "hello":
            if len(v11) >= 2 {
                v11 = v11[1:len(v11)]
            }
        case "":
            v11 = append(v11, 4294967295)
        case "a":
            p2 = uint64(p1)
        default:
            _ = len(v11)
        }
    case "foo bar":
        p0 = (p0 / ((p0) | 1))
    }
    p2 += p2
    for fi12 := 0; fi12 < 14; fi12++ {
        switch "a" {
        case "\n\t":
            v13 := GenS1{F0: (46 < -1)}
            _ = v13
        }
        var v14 uint = p1
        _ = fi12
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 uint32 = 214
    var v2 float64 = 3.14159
    var v3 int64 = -9223372036854775808
    var v4 uint8 = 142
    var v5 uint32 = 42
    var v6 float32 = (1e10 + -2.25)
    v15 := GenS2{F0: 723}
    fmt.Println(v1, v6)
    var v16 uint32 = v5
    fmt.Println(v3, v5, v3)
    v17 := GenS3{F0: 152, F1: -1}
    if (v6 > v6) {
        v1 = genFn1(18446744073709551615, 696, 42)
        if ((v4 % ((42) | 1)) == v4) {
            if true {
                v18 := []int{595, 204}
                v19 := GenS4{F0: 4294967295, F1: 2147483647}
                var v20 uint8 = 127
                _ = v18
                _ = v19
                _ = v20
            }
        }
        v16 -= 2147483648
    } else {
        var v21 uint32 = (^(v5))
        if ((v6 / v6) >= 1e10) {
            if true {
                var v22 string = ("\n\t" + (("foo bar" + ("" + "foo bar")) + ("" + (("" + (("hello" + "") + ("\n\t" + "hello"))) + ("foo bar" + "")))))
                v23 := []int{1, -9223372036854775808, 0}
                _ = v22
                _ = v23
            }
        } else {
            var v24 uint16 = 633
            _ = v24
        }
        var v25 float64 = 1e10
        _ = v21
        _ = v25
    }
    switch 2147483647 {
    case 225:
        fmt.Println(v2)
    default:
        switch "\n\t" {
        case "hello":
            var v26 string = ("hello" + "hello")
            _ = v26
        case "a":
            v27 := GenS3{F0: 0, F1: -643}
            _ = v27
        case "foo bar":
            if (!(!(v3 > -1))) {
                v28 := []int{-2147483648, 128, -9223372036854775808}
                fmt.Println(v3, v2, v5)
                _ = v28
            } else {
                var v29 int64 = (v3 & (v3 / ((v3) | 1)))
                for fi30 := 0; fi30 < 18; fi30++ {
                    v1 = v16
                    var v31 int32 = 2147483647
                    _ = fi30
                    _ = v31
                }
                _ = v29
            }
        default:
            for fi32 := 0; fi32 < 20; fi32++ {
                if (!(false || ((2147483647 > 1) || true))) {
                    v1 -= (v5 / (((v1 & v5)) | 1))
                    v5 -= (v1 << ((uint(v4)) & 63))
                }
                _ = fi32
            }
        }
    }
    if (((v5 + genFn1(174, 18446744073709551615, 1)) > v1) || (!true)) {
        v6 = (v6 * float32(float64(v6)))
    } else {
        if (32767 <= 255) {
            var v33 int = 736
            if (((((v2 / 0.1) + v2) / ((v2 / 0.1) * v2)) >= (1e10 - (v2 * (float64(v6) + ((v2 / (-2.25 + v2)) / v2))))) && false) {
                v34 := []int{(v33 >> ((1) & 63))}
                v2 += v2
                _ = v34
            }
            _ = v33
        }
    }
    v2 += 0.0
    var v35 int64 = (v3 % (((v3 / ((v3) | 1))) | 1))
    for fi36 := 0; fi36 < 7; fi36++ {
        v1 = 1
        v1 = v5
        _ = fi36
    }
    v37 := []int{}
    switch 1 {
    case 127:
        var v38 uint8 = v4
        _ = v38
    case 32768:
        v16 = (^(v5))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
