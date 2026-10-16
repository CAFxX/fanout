package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 uint64; F1 uint32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int8; F1 uint; F2 uint8 }
type GenS5 struct { F0 uint; F1 bool }
type GenS6 struct { F0 int8; F1 int8; F2 string }
type GenS7 struct { F0 uint64; F1 int32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{56, 0, 108}
    v3 := []int{-2147483648, 127}
    v4 := make(map[uint32]string)
    var v5 uint = 45
    v5 += (v5 ^ 65535)
    switch -9223372036854775808 {
    case -2147483648:
        var v6 uint = (v5 >> (((v5 * 445)) & 63))
        _ = v6
    case 32768:
        if (329 <= 2147483647) {
            v7 := v1[42]
            _ = v7
        } else {
            var v8 int = 32768
            _ = v8
        }
    case 42:
        v1[-125] = 916
    default:
        var v9 uint = (v5 >> ((v5) & 63))
        _ = v9
    }
    v10 := GenS1{F0: 30, F1: 0}
    v11 := GenS2{F0: 912, F1: 0}
    switch 0 {
    case 401:
        v5 -= 671
    case 1:
        v5 = (v5 << ((255) & 63))
    case 4294967295:
        var v12 int64 = 718
        _ = v12
    }
    if (2 >= 703) {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi13 := 0; fi13 < 12; fi13++ {
            _ = len(v3)
            _ = fi13
        }
        v14 := GenS3{F0: 1, F1: "\n\t"}
        _ = v14
    } else {
        v15 := v1[127]
        var v16 float64 = 3.14159
        _ = v15
        _ = v16
    }
    v5 -= 106
    if false {
        v17 := GenS4{F0: -128, F1: v5, F2: 67}
        for fi18 := 0; fi18 < 11; fi18++ {
            v5 = v5
            if true {
                v5 += v5
            } else {
                if len(v2) > 0 {
                    v2[0] = 65535
                }
                v5 += (v5 | (v5 % ((v5) | 1)))
            }
            _ = fi18
        }
        v5 -= v5
        _ = v17
    } else {
        v5 += v5
        v19 := GenS5{F0: 4294967295, F1: true}
        v5 -= v5
        _ = v19
    }
    var v20 bool = false
    if (128 <= 2) {
        v21 := GenS3{F0: 2, F1: ("" + "hello")}
        v20 = true
        var v22 uint16 = 32767
        _ = v21
        _ = v22
    } else {
        _ = len(v4)
        var v23 uint = (v5 ^ 18446744073709551615)
        _ = v23
    }
    v20 = true
    var v24 int8 = 100
    if v20 {
        if false {
            v24 += 0
            var v25 int8 = (v24 / ((v24) | 1))
            _ = v25
        } else {
            for fi26 := 0; fi26 < 13; fi26++ {
                for fi27 := 0; fi27 < 20; fi27++ {
                    v20 = ((((((-538 < -32768) && v20) && (433 <= 42)) || v20) || ((141 != 4294967295) || (true || ((3.14159 == float64(-2.25)) || (v5 >= (v5 % ((v5) | 1))))))) && (!(!false)))
                    if (-2.25 > 1e10) {
                        _ = len(v4)
                    }
                    v20 = true
                    _ = fi27
                }
                v24 = (v24 * v24)
                var v28 uint = (v5 + (v5 / (((^(v5))) | 1)))
                _ = fi26
                _ = v28
            }
            v29 := GenS6{F0: v24, F1: (v24 + (-(v24))), F2: ((((("foo bar" + "\n\t") + "hello") + "a") + "\n\t") + ("" + ("\n\t" + (("hello" + "") + (((("a" + "hello") + ("a" + "")) + ("hello" + "a")) + "foo bar")))))}
            v5 = 0
            _ = v29
        }
    } else {
        if (65535 <= 2147483648) {
            if ((42 < 4294967295) || false) {
                var v30 float32 = 1.5
                v31 := GenS7{F0: 1, F1: 42}
                for fi32 := 0; fi32 < 8; fi32++ {
                    var v33 int8 = (v24 ^ v24)
                    _ = fi32
                    _ = v33
                }
                _ = v30
                _ = v31
            }
        } else {
            var v34 uint8 = 1
            _ = v34
        }
    }
    fmt.Println(v24, v24)
    v1[-9223372036854775808] = 65535
    _ = len(v3)
    switch "\n\t" {
    case "":
        delete(v1, -193)
    default:
        v3 = append(v3, -1)
    }
    fmt.Println(v5, v24, v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
