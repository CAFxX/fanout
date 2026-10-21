package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 bool }
func genFn0(p0 int64, p1 int64, p2 int32) string {
    v6 := []int{}
    v7 := GenS1{F0: -2147483648, F1: true, F2: 846}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return "hello"
}

func main() {
    var v1 uint = 923
    var v2 bool = false
    var v3 int16 = -1
    var v4 string = ("" + "\n\t")
    var v5 float32 = (-2.25 - 1e10)
    if v2 {
        v8 := []int{42, 2147483647}
        for fi9 := 0; fi9 < 6; fi9++ {
            fmt.Println(v4, v3, v2)
            _ = fi9
        }
        for fi10 := 0; fi10 < 13; fi10++ {
            if (v4 >= v4) {
                _ = len(v8)
                _ = len(v8)
                _ = len(v8)
            } else {
                for fi11 := 0; fi11 < 2; fi11++ {
                    if len(v8) > 0 {
                        v8[0] = (-(fi11))
                    }
                    v12 := GenS2{F0: v4}
                    v5 = (((3.14159 * v5) / (-2.25 * (float32(0.1) - v5))) / 1.0)
                    _ = fi11
                    _ = v12
                }
                fmt.Println(v4, fi10, v4)
                v3 += (v3 * -371)
            }
            fmt.Println(v3)
            switch fi10 {
            case -9223372036854775808:
                var v13 uint = 2
                _ = v13
            case 4294967295:
                if (!false) {
                    var v14 uint16 = 2
                    v5 += v5
                    for fi15 := 0; fi15 < 7; fi15++ {
                        v1 = (^(v1))
                        _ = fi15
                    }
                    _ = v14
                }
            case 32767:
                v3 = (v3 * (v3 >> ((v1) & 63)))
            default:
                var v16 int64 = 9223372036854775807
                _ = v16
            }
            _ = fi10
        }
        _ = v8
    } else {
        switch 4294967295 {
        case -15:
            v5 += v5
        }
        var v17 int = 128
        _ = v17
    }
    switch 433 {
    case -939:
        v3 += 42
    default:
        v18 := GenS3{F0: 0, F1: (-1 != 0)}
        _ = v18
    }
    v19 := []int{4294967295}
    v2 = v2
    for fi20 := 0; fi20 < 1; fi20++ {
        if len(v19) >= 2 {
            v19 = v19[1:len(v19)]
        }
        _ = fi20
    }
    v2 = (!(-2147483648 <= int32(v3)))
    if false {
        v5 += v5
    }
    if v2 {
        switch v4 {
        case "\n\t":
            for fi21 := 0; fi21 < 7; fi21++ {
                v22 := GenS3{F0: 42, F1: v2}
                for fi23 := 0; fi23 < 8; fi23++ {
                    var v24 uint16 = 42
                    v2 = true
                    var v25 int8 = -105
                    fmt.Println(v5)
                    _ = fi23
                    _ = v24
                    _ = v25
                }
                _ = fi21
                _ = v22
            }
        case "":
            for fi26 := 0; fi26 < 9; fi26++ {
                v3 += v3
                _ = fi26
            }
        case "foo bar":
            v5 = 0.1
        default:
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        }
    } else {
        v27 := GenS1{F0: -763, F1: false, F2: (v1 + v1)}
        _ = v27
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
