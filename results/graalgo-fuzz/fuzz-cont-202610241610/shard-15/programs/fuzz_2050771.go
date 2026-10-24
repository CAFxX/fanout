package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 int64; F1 string }
func genFn0() int32 {
    var v5 int64 = 9223372036854775807
    if false {
        fmt.Println(v5)
        if true {
            if true {
                v5 = v5
            } else {
                var v6 int32 = 391
                switch -844 {
                case 127:
                    fmt.Println(v5)
                case -796:
                    for fi7 := 0; fi7 < 1; fi7++ {
                        var v8 int = (-(fi7))
                        if false {
                            if true {
                                v6 -= 88
                                v8 += (fi7 >> ((4294967295) & 63))
                                for fi9 := 0; fi9 < 7; fi9++ {
                                    v8 = (fi9 / ((32768) | 1))
                                    _ = fi9
                                }
                            } else {
                                for fi10 := 0; fi10 < 3; fi10++ {
                                    v11 := GenS1{F0: ("a" + ("a" + "")), F1: -276}
                                    _ = fi10
                                    _ = v11
                                }
                            }
                            v12 := []int{v8}
                            _ = v12
                        }
                        _ = fi7
                        _ = v8
                    }
                case 4294967295:
                    v5 -= -959
                }
                v13 := []int{}
                _ = v13
                _ = v6
            }
            var v14 bool = false
            switch "a" {
            case "":
                if (65535 != 2) {
                    var v15 uint16 = 1
                    var v16 uint = 4294967295
                    _ = v15
                    _ = v16
                } else {
                    var v17 bool = (1000 == 42)
                    for fi18 := 0; fi18 < 13; fi18++ {
                        fmt.Println(fi18, v14)
                        v14 = (("" + "a") != (((("hello" + "") + "a") + "foo bar") + "a"))
                        fmt.Println(v14, fi18, v5)
                        _ = fi18
                    }
                    _ = v17
                }
            case "foo bar":
                v19 := GenS2{F0: (!false)}
                _ = v19
            case "hello":
                v20 := []int{4294967295, 2147483647, 128}
                _ = v20
            default:
                v5 += (v5 | (v5 >> ((42) & 63)))
            }
            _ = v14
        }
    }
    v21 := []int{42}
    _ = len(v21)
    switch 0 {
    case -9223372036854775808:
        var v22 uint32 = 4294967295
        _ = v22
    case -243:
        v5 = v5
    }
    _ = v21
    _ = v5
    return 897
}

func main() {
    var v1 int = 128
    var v2 float64 = -2.25
    var v3 uint16 = 2
    var v4 float32 = (3.14159 * (1e10 * 0.0))
    for fi23 := 0; fi23 < 18; fi23++ {
        v24 := GenS2{F0: false}
        if true {
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 uint32 = 499
                _ = fi25
                _ = v26
            }
        }
        var v27 int16 = -124
        _ = fi23
        _ = v24
        _ = v27
    }
    if true {
        v28 := GenS3{F0: 32767, F1: (v1 / ((v1) | 1))}
        switch v1 {
        case 2147483647:
            v2 += (float64(3.14159) + (v2 * v2))
        }
        v29 := []int{801, 32767}
        _ = v28
        _ = v29
    } else {
        v30 := []int{}
        fmt.Println(v1)
        var v31 string = "a"
        _ = v30
        _ = v31
    }
    v1 = 32767
    v32 := []int{(^(v1))}
    v33 := GenS1{F0: "", F1: 32767}
    v4 = float32(float64(-2.25))
    v34 := GenS4{F0: 9223372036854775807, F1: "\n\t"}
    v4 -= v4
    fmt.Println(v1, v1, v3)
    v3 = (v3 >> ((1) & 63))
    v4 += ((((3.14159 - (v4 - v4)) / (-2.25 - v4)) / 1.0) / 0.1)
    v4 -= (3.14159 + v4)
    _ = len(v32)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
