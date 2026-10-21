package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool; F1 uint32 }

func main() {
    var v1 uint = 775
    var v2 int8 = -1
    var v3 int = -977
    var v4 int64 = 0
    switch "" {
    case "\n\t":
        v5 := []int{(v3 & v3)}
        _ = v5
    case "hello":
        if (0 > 2147483647) {
            v6 := GenS1{F0: 137}
            for fi7 := 0; fi7 < 5; fi7++ {
                v8 := GenS1{F0: 1}
                v3 = fi7
                for fi9 := 0; fi9 < 1; fi9++ {
                    switch v3 {
                    case 9223372036854775807:
                        if (468 != 928) {
                            v10 := GenS1{F0: 2}
                            _ = v10
                        } else {
                            switch fi9 {
                            case -518:
                                fmt.Println(v3, v4, fi9)
                            case 127:
                                v11 := []int{-793, (fi7 | v3), 42, v3}
                                _ = v11
                            }
                            switch (("a" + (("\n\t" + ("\n\t" + "hello")) + "\n\t")) + ("a" + "")) {
                            case "foo bar":
                                var v12 uint = v1
                                _ = v12
                            case "":
                                v3 += fi7
                            }
                        }
                    }
                    _ = fi9
                }
                _ = fi7
                _ = v8
            }
            _ = v6
        } else {
            v2 = (v2 << ((1) & 63))
            v2 = -89
        }
    }
    var v13 uint32 = 2147483648
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 float32 = -2.25
        fmt.Println(v3)
        v16 := []int{335, (fi14 * (v3 & 9223372036854775807)), fi14}
        _ = fi14
        _ = v15
        _ = v16
    }
    switch "a" {
    case "\n\t":
        fmt.Println(v4)
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        v3 = v3
        _ = fi17
    }
    var v18 uint16 = 65535
    var v19 uint8 = 1
    v18 = 324
    if (!false) {
        for fi20 := 0; fi20 < 1; fi20++ {
            fmt.Println(v13, v19)
            _ = fi20
        }
    }
    if false {
        if false {
            if true {
                v4 = v4
            } else {
                v18 -= (v18 ^ v18)
            }
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := GenS2{F0: (v18 + (v18 + (v18 | 255)))}
                _ = fi21
                _ = v22
            }
            switch "a" {
            case "foo bar":
                fmt.Println(v19, v13)
            case "\n\t":
                if (((!false) || false) || true) {
                    var v23 int16 = -565
                    fmt.Println(v19)
                    if false {
                        var v24 float32 = -2.25
                        v25 := GenS3{F0: (((v3 + v3) > v3) && true), F1: (v13 % ((42) | 1))}
                        _ = v24
                        _ = v25
                    }
                    _ = v23
                } else {
                    v1 = v1
                    var v26 uint64 = 411
                    v19 -= 126
                    _ = v26
                }
            default:
                v19 += v19
            }
        }
        fmt.Println(v18, v3, v13)
    }
    fmt.Println(v2)
    for fi27 := 0; fi27 < 5; fi27++ {
        var v28 float64 = float64(0.1)
        v29 := GenS1{F0: 157}
        var v30 float64 = (((v28 * -2.25) + (-2.25 - float64(0.1))) * 1.5)
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    v2 += -82
    v18 = (v18 * (v18 * v18))
    var v31 uint32 = v13
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
