package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 uint16 }
type GenS2 struct { F0 bool; F1 int; F2 uint }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool; F1 int64; F2 int16 }
type GenS5 struct { F0 string; F1 uint32 }

func main() {
    var v1 uint8 = 151
    var v2 bool = (false || (!true))
    v3 := make(map[uint32]string)
    var v4 int32 = 883
    v5 := make(map[uint32]bool)
    var v6 int8 = 1
    for fi7 := 0; fi7 < 10; fi7++ {
        if (32767 < -358) {
            if ((v2 || (v1 != 149)) && v2) {
                v4 -= 2147483647
            } else {
                fmt.Println(v1)
            }
            if (!false) {
                v6 -= v6
            }
            for fi8 := 0; fi8 < 17; fi8++ {
                v9 := []int{fi7}
                v10 := GenS1{F0: v6, F1: 255, F2: 65535}
                if v2 {
                    var v11 int64 = 9223372036854775807
                    if v2 {
                        v12 := GenS2{F0: v2, F1: (fi8 - (fi8 / (((fi8 & 689)) | 1))), F2: 4294967295}
                        _ = v12
                    } else {
                        var v13 string = ("foo bar" + "")
                        _ = v13
                    }
                    _ = v11
                } else {
                    v14 := v5[427]
                    _ = v14
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        }
        switch "hello" {
        case "\n\t":
            var v15 float64 = 0.0
            _ = v15
        case "":
            switch ("a" + ("hello" + ("a" + ("" + "")))) {
            case "foo bar":
                v6 += v6
            case "a":
                v16 := v5[727]
                _ = v16
            case "hello":
                switch fi7 {
                case 42:
                    var v17 int16 = -32768
                    _ = v17
                case 0:
                    for fi18 := 0; fi18 < 14; fi18++ {
                        fmt.Println(v6)
                        _ = fi18
                    }
                }
            default:
                v6 -= (v6 % ((v6) | 1))
            }
        case "hello":
            fmt.Println(v1)
        }
        var v19 float32 = 3.14159
        _ = fi7
        _ = v19
    }
    var v20 int8 = v6
    switch 2147483647 {
    case 128:
        fmt.Println(v1)
    case 32768:
        fmt.Println(v4)
    case -887:
        v21 := GenS3{F0: 88}
        _ = v21
    }
    v22 := []int{-9223372036854775808}
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 5; fi24++ {
            fmt.Println(v2, fi24, v20)
            v25 := GenS4{F0: v2, F1: -41, F2: 42}
            if (false && false) {
                var v26 uint8 = v1
                fmt.Println(fi24, fi24, fi24)
                switch (((("\n\t" + "foo bar") + ("a" + ("a" + "hello"))) + "") + ("\n\t" + "hello")) {
                case "":
                    v26 -= (v26 * 226)
                default:
                    var v27 uint32 = 4294967295
                    _ = v27
                }
                _ = v26
            } else {
                fmt.Println(v2, fi23, fi24)
                v4 -= 0
            }
            _ = fi24
            _ = v25
        }
        v28 := GenS5{F0: "\n\t", F1: 2147483647}
        v4 = v4
        _ = fi23
        _ = v28
    }
    v4 = v4
    fmt.Println(v6)
    v22 = append(v22, 42)
    for fi29 := 0; fi29 < 9; fi29++ {
        for fi30 := 0; fi30 < 18; fi30++ {
            fmt.Println(v2, v6, v4)
            if len(v22) >= 2 {
                v22 = v22[1:len(v22)]
            }
            var v31 uint = 0
            _ = fi30
            _ = v31
        }
        var v32 int16 = -1
        _ = fi29
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
