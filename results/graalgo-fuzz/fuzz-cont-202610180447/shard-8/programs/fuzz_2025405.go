package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint }
type GenS3 struct { F0 uint8; F1 uint8 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint16 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    var v3 int32 = 733
    var v4 uint16 = 255
    var v5 int16 = 482
    for fi6 := 0; fi6 < 4; fi6++ {
        switch "a" {
        case "hello":
            switch (fi6 << ((593) & 63)) {
            case 391:
                fmt.Println(fi6, v4)
            case -739:
                v7 := GenS1{F0: 65535, F1: -87}
                _ = v7
            }
        case "\n\t":
            v8 := GenS1{F0: 604, F1: fi6}
            _ = v8
        }
        for fi9 := 0; fi9 < 14; fi9++ {
            switch "" {
            case "foo bar":
                v4 += (v4 + (^(v4)))
            case "a":
                v3 = v3
            }
            fmt.Println(v5, fi9)
            var v10 int8 = 33
            _ = fi9
            _ = v10
        }
        v11 := v2[2147483647]
        _ = fi6
        _ = v11
    }
    v12 := GenS2{F0: 0, F1: (v4 & (^(v4))), F2: 2}
    for fi13 := 0; fi13 < 7; fi13++ {
        var v14 float64 = -2.25
        v15 := GenS3{F0: 0, F1: 217}
        _ = fi13
        _ = v14
        _ = v15
    }
    if true {
        v16 := []int{int(v3), 128, -978, 2147483647}
        _ = v16
    }
    switch "\n\t" {
    case "hello":
        var v17 uint16 = v4
        _ = v17
    case "foo bar":
        v1[-9223372036854775808] = 0
    case "":
        if (!true) {
            if ((v5 << ((724) & 63)) >= v5) {
                v18 := GenS4{F0: "\n\t"}
                _ = v18
            }
            fmt.Println(v4)
            v19 := GenS1{F0: 872, F1: 4294967295}
            _ = v19
        } else {
            v1[127] = 980
            v3 = (v3 / ((v3) | 1))
            switch 0 {
            case 127:
                v20 := []int{-562, -1, -1, 9223372036854775807}
                _ = v20
            default:
                if (true || false) {
                    for fi21 := 0; fi21 < 16; fi21++ {
                        for fi22 := 0; fi22 < 20; fi22++ {
                            fmt.Println(v4, v5, fi21)
                            switch (fi22 >> ((356) & 63)) {
                            case -1:
                                var v23 int32 = (v3 - 1)
                                _ = v23
                            case 127:
                                v5 = (v5 * int16(fi22))
                            default:
                                fmt.Println(v3, fi21, v3)
                            }
                            _ = fi22
                        }
                        _ = fi21
                    }
                    switch (("\n\t" + "hello") + ("foo bar" + "a")) {
                    case "a":
                        fmt.Println(v3, v4)
                    case "foo bar":
                        switch "hello" {
                        case "hello":
                            for fi24 := 0; fi24 < 10; fi24++ {
                                v4 = v4
                                v4 = (^(v4))
                                _ = fi24
                            }
                        case "foo bar":
                            v4 -= 255
                        case "a":
                            for fi25 := 0; fi25 < 1; fi25++ {
                                v3 -= 42
                                _ = fi25
                            }
                        default:
                            v26 := GenS3{F0: 243, F1: 140}
                            _ = v26
                        }
                    default:
                        v27 := GenS4{F0: "a"}
                        _ = v27
                    }
                }
            }
        }
    }
    if true {
        v5 += v5
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        for fi29 := 0; fi29 < 7; fi29++ {
            var v30 int8 = -128
            switch (fi28 & (fi28 >> ((255) & 63))) {
            case -347:
                fmt.Println(v3, fi28, v30)
            case 426:
                var v31 uint16 = 2
                _ = v31
            case 389:
                var v32 bool = false
                _ = v32
            default:
                v33 := []int{32768, fi29}
                _ = v33
            }
            var v34 int = int(v3)
            _ = fi29
            _ = v30
            _ = v34
        }
        if (!((2 < 1) && true)) {
            if false {
                v3 += int32(v5)
                var v35 uint = 192
                _ = v35
            }
        } else {
            var v36 int = (fi28 & (fi28 ^ fi28))
            _ = v36
        }
        _ = fi28
    }
    v37 := GenS5{F0: 2147483647, F1: 61, F2: v4}
    switch 9223372036854775807 {
    case 58:
        if false {
            var v38 uint8 = 42
            _ = v38
        }
    case 127:
        fmt.Println(v3, v4, v4)
    default:
        v39 := GenS5{F0: 42, F1: 1, F2: 255}
        _ = v39
    }
    if true {
        fmt.Println(v4, v4, v3)
    } else {
        v2[2] = ""
    }
    var v40 float32 = 1e10
    v3 = v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v40
    _ = v5
}
