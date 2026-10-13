package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint64 }

func main() {
    var v1 float64 = 0.1
    v2 := []string{"foo bar"}
    var v3 int32 = -526
    v4 := make(map[uint32]uint32)
    var v5 uint = 4294967295
    if true {
        for fi6 := 0; fi6 < 8; fi6++ {
            var v7 bool = false
            _ = fi6
            _ = v7
        }
        if len(v2) > 0 {
            v2[0] = ""
        }
    } else {
        var v8 string = "hello"
        var v9 uint16 = 224
        _ = v8
        _ = v9
    }
    if (-1 == (-(v3))) {
        var v10 uint = (^(v5))
        _ = v10
    }
    fmt.Println(v5, v3)
    v4[4294967295] = 2147483647
    switch ("a" + "") {
    case "hello":
        var v11 int64 = 0
        _ = v11
    case "foo bar":
        switch 32767 {
        case 128:
            fmt.Println(v3, v3, v3)
        case 32768:
            var v12 float64 = (3.14159 + 0.1)
            _ = v12
        case -9223372036854775808:
            v5 = (^(v5))
        default:
            v1 = 1.5
        }
    default:
        fmt.Println(v1, v3)
    }
    v1 += v1
    v1 -= v1
    if (0 <= 42) {
        fmt.Println(v5, v5, v3)
        if true {
            for fi13 := 0; fi13 < 10; fi13++ {
                if false {
                    fmt.Println(v1, v3)
                } else {
                    v1 -= (0.1 + v1)
                }
                var v14 uint8 = 255
                var v15 int8 = -52
                _ = fi13
                _ = v14
                _ = v15
            }
            if true {
                _ = len(v2)
                v2 = append(v2, "foo bar")
            }
            fmt.Println(v3, v1)
        } else {
            v16 := GenS1{F0: 42, F1: 42}
            if (810 <= 32767) {
                for fi17 := 0; fi17 < 7; fi17++ {
                    for fi18 := 0; fi18 < 20; fi18++ {
                        switch -9223372036854775808 {
                        case 0:
                            fmt.Println(v3, fi18, v5)
                        }
                        _ = fi18
                    }
                    fmt.Println(v3, v1)
                    _ = fi17
                }
                if false {
                    if true {
                        var v19 string = ("foo bar" + ((("foo bar" + "") + ("\n\t" + "\n\t")) + ("a" + "hello")))
                        _ = v19
                    } else {
                        _ = len(v4)
                    }
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if (v1 > (3.14159 - 1e10)) {
                if false {
                    switch "foo bar" {
                    case "a":
                        if true {
                            var v20 bool = (false && (((false && (0 < 18446744073709551615)) && (42 == 61)) && (true && true)))
                            fmt.Println(v3, v5, v20)
                            _ = v20
                        }
                    default:
                        if len(v2) > 0 {
                            v2[0] = "foo bar"
                        }
                    }
                    for fi21 := 0; fi21 < 12; fi21++ {
                        fmt.Println(v1, fi21)
                        _ = fi21
                    }
                    v1 = float64(float32(-2.25))
                }
                v3 += 42
                fmt.Println(v5, v1)
            } else {
                var v22 float64 = 3.14159
                for fi23 := 0; fi23 < 18; fi23++ {
                    fmt.Println(v1, v22)
                    for fi24 := 0; fi24 < 2; fi24++ {
                        _ = len(v4)
                        fmt.Println(v22)
                        _ = fi24
                    }
                    _ = fi23
                }
                var v25 int16 = 32767
                _ = v22
                _ = v25
            }
            _ = v16
        }
        if true {
            v26 := GenS1{F0: 2147483647, F1: 4294967295}
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 2; fi27++ {
            fmt.Println(v1, v5, v5)
            if false {
                v1 -= v1
                switch "foo bar" {
                case "a":
                    _ = len(v4)
                case "foo bar":
                    v4[111] = 2
                }
                v28 := GenS2{F0: 0, F1: (^(fi27))}
                _ = v28
            }
            v3 += v3
            _ = fi27
        }
        if (((94 <= 9223372036854775807) && (0 >= 1)) || (42 >= 536)) {
            for fi29 := 0; fi29 < 7; fi29++ {
                fmt.Println(v1, v5)
                _ = fi29
            }
            v30 := GenS3{F0: -1}
            _ = v30
        }
        v1 += v1
    }
    fmt.Println(v3, v3)
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS1{F0: 65535, F1: 32768}
        _ = fi31
        _ = v32
    }
    var v33 float64 = v1
    fmt.Println(v3, v3, v3)
    switch -9223372036854775808 {
    case 127:
        for fi34 := 0; fi34 < 18; fi34++ {
            v1 = v33
            var v35 uint64 = 4294967295
            _ = fi34
            _ = v35
        }
    case 32768:
        v33 -= v33
    case -2147483648:
        v36 := GenS4{F0: -653, F1: 1}
        _ = v36
    }
    var v37 uint32 = 42
    for fi38 := 0; fi38 < 16; fi38++ {
        v5 += v5
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v4
    _ = v5
}
