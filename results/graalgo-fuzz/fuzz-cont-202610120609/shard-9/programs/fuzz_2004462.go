package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint16; F1 int64; F2 uint8 }
func genFn0() uint8 {
    var v7 float64 = ((((3.14159 * -2.25) * 0.1) - 1.5) - 1.5)
    fmt.Println(v7)
    _ = v7
    return 2
}

func main() {
    v1 := []int{}
    v2 := []string{((("\n\t" + ("hello" + ("a" + ("\n\t" + "\n\t")))) + "\n\t") + "")}
    var v3 float64 = float64(3.14159)
    var v4 uint = 176
    v5 := make(map[int]string)
    v6 := []string{"a", "foo bar", "a", "foo bar"}
    var v8 int8 = 43
    var v9 int16 = -409
    v8 += (v8 + v8)
    if false {
        switch ("\n\t" + "") {
        case "foo bar":
            var v10 uint32 = 0
            _ = v10
        }
        if false {
            fmt.Println(v9)
            v9 += v9
        } else {
            var v11 uint32 = 42
            v12 := GenS1{F0: 42, F1: "hello"}
            _ = v11
            _ = v12
        }
        var v13 int64 = 9223372036854775807
        _ = v13
    } else {
        if true {
            switch "\n\t" {
            case "":
                fmt.Println(v3)
            default:
                var v14 float64 = ((float64((1e10 + ((0.1 - 0.1) - 1.5))) + v3) - v3)
                _ = v14
            }
        }
        switch 9223372036854775807 {
        case 32767:
            if (!false) {
                v15 := GenS2{F0: true, F1: 1}
                v8 -= (v8 % ((v8) | 1))
                for fi16 := 0; fi16 < 16; fi16++ {
                    var v17 float64 = v3
                    _ = fi16
                    _ = v17
                }
                _ = v15
            } else {
                switch 128 {
                case 127:
                    var v18 int = 32768
                    _ = v18
                }
                var v19 uint32 = 2147483648
                v1 = append(v1, -708)
                _ = v19
            }
        case 4294967295:
            for fi20 := 0; fi20 < 18; fi20++ {
                if (!true) {
                    fmt.Println(fi20, fi20, v8)
                }
                _ = fi20
            }
        default:
            v4 -= (^(v4))
        }
    }
    var v21 uint8 = 115
    v22 := GenS2{F0: true, F1: 382}
    v23 := GenS3{F0: false}
    var v24 string = "\n\t"
    _ = len(v5)
    if true {
        fmt.Println(v8, v24)
    } else {
        v8 -= (v8 | v8)
        v24 = ("a" + ("" + (((v24 + v24) + v24) + ((v24 + v24) + "foo bar"))))
        fmt.Println(v21, v3, v21)
    }
    switch 127 {
    case 32768:
        if len(v2) > 0 {
            v2[0] = (v24 + (((v24 + (v24 + v24)) + (("hello" + "hello") + v24)) + v24))
        }
    default:
        v25 := GenS4{F0: 89, F1: 9223372036854775807, F2: v21}
        _ = v25
    }
    if true {
        v3 += v3
        var v26 int64 = 9223372036854775807
        _ = v26
    } else {
        v21 = v21
    }
    var v27 string = "foo bar"
    switch -2147483648 {
    case 32767:
        var v28 uint8 = (^(v21))
        _ = v28
    }
    var v29 int = 501
    _ = len(v5)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
