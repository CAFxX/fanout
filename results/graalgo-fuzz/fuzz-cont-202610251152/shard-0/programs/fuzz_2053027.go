package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
func genFn0(p0 bool, p1 int8, p2 uint8) uint16 {
    var v5 uint = 0
    var v6 uint = 2
    for fi7 := 0; fi7 < 7; fi7++ {
        p0 = (("a" + "a") >= ("a" + (("foo bar" + "hello") + "\n\t")))
        var v8 uint = (^(v5))
        var v9 float32 = -2.25
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 string = ("hello" + "foo bar")
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v5
    _ = v6
    return 763
}
func genFn1() string {
    for fi11 := 0; fi11 < 12; fi11++ {
        v12 := GenS1{F0: (fi11 % (((fi11 << ((2) & 63))) | 1)), F1: "\n\t"}
        _ = fi11
        _ = v12
    }
    var v13 int8 = 69
    if true {
        v14 := []int{346, 9223372036854775807, -9223372036854775808, 32767}
        _ = v14
    } else {
        if ((!(1 < 2)) || ("foo bar" != "foo bar")) {
            v15 := []int{}
            v13 = v13
            switch "\n\t" {
            case "":
                v13 -= 100
            }
            _ = v15
        }
        v13 = (v13 >> ((255) & 63))
    }
    var v16 uint32 = 4294967295
    fmt.Println(v16)
    _ = v13
    _ = v16
    return ("\n\t" + "\n\t")
}

func main() {
    v1 := []uint32{2, 1, 125, 2147483648}
    var v2 int16 = 32767
    var v3 uint16 = 65535
    var v4 int32 = 0
    var v17 string = ("hello" + ("\n\t" + (("foo bar" + "\n\t") + (("a" + ("" + "hello")) + "foo bar"))))
    var v18 uint16 = (v3 | 255)
    var v19 uint8 = 255
    v1 = append(v1, 65535)
    for fi20 := 0; fi20 < 8; fi20++ {
        v18 = (v18 / ((genFn0(false, 0, v19)) | 1))
        _ = fi20
    }
    for fi21 := 0; fi21 < 9; fi21++ {
        var v22 bool = false
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 17; fi23++ {
        v4 += 1
        _ = fi23
    }
    v19 = v19
    v17 = ("a" + "")
    var v24 uint64 = 2
    v3 += (v3 | 379)
    v25 := GenS1{F0: 9223372036854775807, F1: "\n\t"}
    fmt.Println(v24)
    if (0 < -1) {
        var v26 uint32 = 65535
        _ = v26
    } else {
        v18 -= (v3 >> ((1) & 63))
        switch "foo bar" {
        case "\n\t":
            switch "foo bar" {
            case "foo bar":
                var v27 string = ("a" + v17)
                _ = v27
            case "hello":
                v17 = v17
            default:
                if false {
                    switch v17 {
                    case "":
                        var v28 int = 2147483647
                        _ = v28
                    }
                }
            }
        case "hello":
            v24 -= (v24 % (((v24 + (v24 * 940))) | 1))
        case "":
            fmt.Println(v2, v2, v19)
        default:
            var v29 int32 = (v4 & (v4 << ((4294967295) & 63)))
            _ = v29
        }
    }
    if len(v1) > 0 {
        v1[0] = 4294967295
    }
    v30 := GenS1{F0: 4294967295, F1: (v17 + "")}
    var v31 bool = true
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
