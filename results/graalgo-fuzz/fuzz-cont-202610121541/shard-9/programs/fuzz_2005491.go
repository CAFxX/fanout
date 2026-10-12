package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32; F1 int }
func genFn0(p0 int) int16 {
    if (!(("hello" + "") < "\n\t")) {
        if true {
            var v6 uint16 = 32767
            fmt.Println(v6, v6)
            v6 += 967
            _ = v6
        }
    }
    v7 := []int{}
    var v8 string = ("\n\t" + "")
    if len(v7) > 0 {
        v7[0] = 4294967295
    }
    _ = p0
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 uint = 65535
    var v2 float32 = (0.1 + float32(1e10))
    v3 := make(map[uint32]bool)
    var v4 float32 = 0.1
    v5 := make(map[string]string)
    for fi9 := 0; fi9 < 20; fi9++ {
        fmt.Println(fi9, fi9)
        v10 := GenS1{F0: ("a" + (("foo bar" + "hello") + ""))}
        var v11 int8 = 42
        _ = fi9
        _ = v10
        _ = v11
    }
    v4 -= 3.14159
    v12 := []int{}
    var v13 uint = (v1 / (((v1 << (((v1 >> (((^(v1))) & 63))) & 63))) | 1))
    switch ((("\n\t" + "foo bar") + "") + "") {
    case "a":
        if false {
            _ = len(v3)
            for fi14 := 0; fi14 < 15; fi14++ {
                for fi15 := 0; fi15 < 1; fi15++ {
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                    _ = fi15
                }
                fmt.Println(v4, v4)
                _ = fi14
            }
            v4 = 0.1
        }
    case "":
        switch "" {
        case "foo bar":
            if true {
                _ = len(v12)
                if (!(false || (!true))) {
                    for fi16 := 0; fi16 < 10; fi16++ {
                        var v17 uint16 = 2
                        var v18 int32 = -2147483648
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                    switch 136 {
                    case -9223372036854775808:
                        if true {
                            switch 884 {
                            case 128:
                                v19 := GenS1{F0: ("a" + "hello")}
                                _ = v19
                            }
                            switch "foo bar" {
                            case "hello":
                                for fi20 := 0; fi20 < 6; fi20++ {
                                    fmt.Println(v1, v1)
                                    _ = fi20
                                }
                            case "a":
                                var v21 uint = v1
                                _ = v21
                            case "":
                                for fi22 := 0; fi22 < 17; fi22++ {
                                    if len(v12) >= 2 {
                                        v12 = v12[1:len(v12)]
                                    }
                                    for fi23 := 0; fi23 < 17; fi23++ {
                                        v12 = append(v12, fi22)
                                        v4 += v4
                                        v13 = v13
                                        _ = fi23
                                    }
                                    _ = fi22
                                }
                            default:
                                fmt.Println(v2)
                            }
                        }
                    case 1:
                        v12 = append(v12, 128)
                    case -370:
                        for fi24 := 0; fi24 < 12; fi24++ {
                            fmt.Println(v2, v4)
                            v12 = append(v12, -202)
                            _ = fi24
                        }
                    }
                    var v25 string = ""
                    _ = v25
                } else {
                    fmt.Println(v13, v1, v1)
                }
            }
        case "hello":
            v26 := GenS1{F0: "a"}
            _ = v26
        }
    }
    v13 = (v13 / (((v13 << (((^(v13))) & 63))) | 1))
    v27 := GenS1{F0: "a"}
    v2 = 0.0
    v28 := GenS2{F0: 2147483647, F1: 0}
    v4 -= v4
    switch "" {
    case "hello":
        switch ("a" + "") {
        case "\n\t":
            _ = len(v5)
        case "hello":
            v12 = append(v12, -1)
        }
    }
    if (-1 > 42) {
        if len(v12) > 0 {
            v12[0] = 265
        }
        delete(v3, 113)
    }
    if false {
        var v29 int = 728
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
