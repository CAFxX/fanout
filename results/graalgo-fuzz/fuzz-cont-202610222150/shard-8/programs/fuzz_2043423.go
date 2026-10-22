package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 int16; F1 uint }
type GenS4 struct { F0 int32; F1 bool }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    var v3 bool = true
    v3 = v3
    v3 = v3
    var v4 bool = (1 <= -577)
    if ((-820 > -9223372036854775808) || (v3 && (227 >= 578))) {
        if v3 {
            for fi5 := 0; fi5 < 11; fi5++ {
                v6 := []int{fi5}
                var v7 int64 = -816
                v8 := GenS1{F0: 65535}
                _ = fi5
                _ = v6
                _ = v7
                _ = v8
            }
            v4 = ((v4 && (v4 || (2 == 2))) && v4)
        } else {
            var v9 int64 = 9223372036854775807
            v10 := []int{}
            _ = v10
            _ = v9
        }
    } else {
        v11 := GenS2{F0: ("foo bar" + "\n\t"), F1: 101}
        switch -375 {
        case -293:
            var v12 int32 = 0
            _ = v12
        case 4294967295:
            v3 = (((((0.1 + (1e10 / (1e10 - 0.1))) + 3.14159) + (0.0 + 1e10)) == (0.0 - -2.25)) && v3)
        case 0:
            switch -2147483648 {
            case 127:
                var v13 string = ("a" + ("hello" + "\n\t"))
                _ = v13
            case -9223372036854775808:
                v1[127] = "\n\t"
            }
        }
        fmt.Println(v4, v4)
        _ = v11
    }
    switch ("\n\t" + "a") {
    case "":
        fmt.Println(v4, v4)
    default:
        var v14 float64 = ((0.0 - (0.1 + 0.1)) / (1.5 + 0.1))
        _ = v14
    }
    v15 := GenS2{F0: "hello", F1: -128}
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            switch "hello" {
            case "hello":
                v4 = (1 != 42)
            default:
                var v17 int32 = -483
                _ = v17
            }
            _ = fi16
        }
    }
    v18 := []int{32768}
    fmt.Println(v3, v3)
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    delete(v2, "a")
    switch "hello" {
    case "\n\t":
        delete(v2, ("\n\t" + ("foo bar" + "")))
    case "hello":
        v4 = v3
    default:
        v19 := GenS1{F0: 65535}
        _ = v19
    }
    var v20 uint16 = 65535
    var v21 float64 = -2.25
    for fi22 := 0; fi22 < 14; fi22++ {
        for fi23 := 0; fi23 < 8; fi23++ {
            switch "hello" {
            case "":
                if (!(2147483647 < 42)) {
                    v24 := GenS2{F0: ("foo bar" + ("" + (("" + ("\n\t" + "a")) + ("hello" + "hello")))), F1: -128}
                    v4 = v4
                    _ = v24
                }
            case "hello":
                if ((2147483647 >= 1) && v3) {
                    switch "a" {
                    case "hello":
                        switch "hello" {
                        case "foo bar":
                            _ = len(v18)
                        case "hello":
                            fmt.Println(v21, fi23, v3)
                        default:
                            var v25 bool = (65535 >= 4294967295)
                            _ = v25
                        }
                    case "a":
                        v18 = append(v18, (-(fi22)))
                    }
                } else {
                    fmt.Println(v3, fi22)
                }
            case "\n\t":
                if (-2.25 <= 0.0) {
                    if (!(v21 < ((v21 + v21) / v21))) {
                        if (18446744073709551615 != 824) {
                            v21 = 0.0
                            fmt.Println(v20, v3)
                            if v4 {
                                v26 := GenS3{F0: 133, F1: 65535}
                                v27 := GenS4{F0: 2147483647, F1: (v3 && true)}
                                var v28 int32 = 0
                                _ = v26
                                _ = v27
                                _ = v28
                            } else {
                                var v29 uint8 = 128
                                _ = v29
                            }
                        } else {
                            v20 = v20
                            v2["hello"] = (false && (!((0.0 / (3.14159 - -2.25)) != -2.25)))
                        }
                        var v30 int = (^(fi22))
                        _ = v30
                    } else {
                        switch (("hello" + ("" + "foo bar")) + "foo bar") {
                        case "a":
                            var v31 uint8 = 0
                            _ = v31
                        default:
                            v20 = (v20 * (^(v20)))
                        }
                        v4 = ((!((((float32((((0.1 + 0.0) * 0.1) / (1e10 * (float64(3.14159) / (v21 / ((v21 / 1.0) / float64(1e10))))))) * 1e10) + 0.1) >= (1.5 + 0.1)) || v3)) || v4)
                    }
                    _ = len(v2)
                    v4 = true
                    v21 = 0.0
                }
            }
            _ = fi23
        }
        _ = fi22
    }
    v18 = append(v18, 9223372036854775807)
    var v32 string = "a"
    var v33 string = v32
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v18
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
