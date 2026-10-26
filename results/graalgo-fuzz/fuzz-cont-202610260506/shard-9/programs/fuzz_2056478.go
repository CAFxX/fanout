package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 string }

func main() {
    var v1 uint32 = 4294967295
    var v2 int16 = 909
    v3 := []int{-9223372036854775808, -521, 0}
    fmt.Println(v1, v2)
    var v4 int16 = (v2 * v2)
    v5 := GenS1{F0: 2, F1: 797}
    for fi6 := 0; fi6 < 15; fi6++ {
        var v7 float64 = 1e10
        var v8 int8 = 127
        if true {
            v8 = v8
            if false {
                if len(v3) > 0 {
                    v3[0] = fi6
                }
                v4 = 1000
                var v9 uint = 0
                _ = v9
            } else {
                v10 := GenS1{F0: 754, F1: -1}
                _ = v10
            }
            v1 = (v1 % (((v1 / (((^(v1))) | 1))) | 1))
        } else {
            v11 := GenS1{F0: 255, F1: -1}
            _ = v11
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    v2 += int16(v1)
    fmt.Println(v2)
    fmt.Println(v4)
    for fi12 := 0; fi12 < 16; fi12++ {
        if len(v3) > 0 {
            v3[0] = fi12
        }
        _ = fi12
    }
    v3 = append(v3, 32767)
    switch "\n\t" {
    case "foo bar":
        if (-32768 != v4) {
            switch 2147483647 {
            case 4294967295:
                for fi13 := 0; fi13 < 3; fi13++ {
                    v1 -= 42
                    _ = fi13
                }
            case 32768:
                for fi14 := 0; fi14 < 4; fi14++ {
                    v1 -= v1
                    v3 = append(v3, fi14)
                    _ = fi14
                }
            case 2147483647:
                v15 := GenS1{F0: 42, F1: 1}
                _ = v15
            default:
                var v16 bool = (!(121 < 2))
                _ = v16
            }
        }
    case "":
        if (-2.25 == -2.25) {
            v4 -= v4
        }
    case "a":
        v4 = 32767
    }
    _ = len(v3)
    v17 := GenS2{F0: ("a" + "a")}
    if false {
        v3 = append(v3, 965)
        for fi18 := 0; fi18 < 8; fi18++ {
            if true {
                var v19 float64 = 1.5
                _ = v19
            }
            for fi20 := 0; fi20 < 1; fi20++ {
                v21 := GenS2{F0: ("" + "foo bar")}
                _ = fi20
                _ = v21
            }
            _ = fi18
        }
        for fi22 := 0; fi22 < 16; fi22++ {
            v1 = (^(v1))
            switch "\n\t" {
            case "foo bar":
                if false {
                    v4 = v2
                    v1 = 42
                } else {
                    v23 := GenS2{F0: "\n\t"}
                    var v24 uint32 = v1
                    v3 = append(v3, 9223372036854775807)
                    _ = v23
                    _ = v24
                }
            case "\n\t":
                var v25 int8 = -128
                _ = v25
            case "hello":
                fmt.Println(fi22)
            default:
                v3 = append(v3, fi22)
            }
            _ = fi22
        }
    } else {
        var v26 int64 = -1
        var v27 uint = 577
        if (1 < 42) {
            var v28 int16 = (^(v4))
            var v29 int16 = (v28 & 140)
            _ = v28
            _ = v29
        }
        _ = v26
        _ = v27
    }
    var v30 uint = 2
    var v31 int32 = 2147483647
    switch -1 {
    case 127:
        var v32 float32 = (-2.25 * float32(0.0))
        _ = v32
    default:
        if false {
            v30 -= v30
        } else {
            v4 += (v2 << ((v30) & 63))
            for fi33 := 0; fi33 < 14; fi33++ {
                _ = len(v3)
                _ = fi33
            }
            var v34 string = ("" + (("hello" + (("foo bar" + ("a" + "")) + "\n\t")) + "\n\t"))
            _ = v34
        }
    }
    var v35 bool = false
    v30 -= 165
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v4
    _ = v5
}
