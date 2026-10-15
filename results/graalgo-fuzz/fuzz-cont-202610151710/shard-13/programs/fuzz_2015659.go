package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint16; F1 int64; F2 uint64 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[int]string)
    var v3 float64 = 0.0
    var v4 uint32 = 1
    v1 = (v1 ^ (v1 << ((65535) & 63)))
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 8; fi6++ {
            switch "hello" {
            case "":
                if (!true) {
                    var v7 int32 = -569
                    v1 = (^(v1))
                    _ = v7
                } else {
                    if (int8(fi5) < 42) {
                        if (!false) {
                            switch 2147483647 {
                            case -565:
                                var v8 int = fi6
                                _ = v8
                            }
                        } else {
                            var v9 uint = 355
                            if (255 <= 2) {
                                fmt.Println(v9, v3)
                            } else {
                                v9 = v9
                                v10 := v2[-2147483648]
                                _ = v10
                            }
                            switch "\n\t" {
                            case "hello":
                                if (1000 >= 1) {
                                    fmt.Println(v4, v9)
                                    fmt.Println(v3)
                                    fmt.Println(v1, v4, fi6)
                                } else {
                                    v9 = v9
                                    fmt.Println(v3)
                                }
                            case "a":
                                v11 := GenS1{F0: "\n\t"}
                                var v12 uint8 = 127
                                var v13 uint16 = uint16(v1)
                                _ = v11
                                _ = v12
                                _ = v13
                            case "foo bar":
                                var v14 int64 = 9223372036854775807
                                _ = v14
                            }
                            _ = v9
                        }
                        var v15 bool = (("" <= ("\n\t" + "")) && true)
                        switch 2147483647 {
                        case -2147483648:
                            var v16 float64 = (((1e10 + v3) - 0.1) / v3)
                            var v17 uint64 = (v1 + (v1 / (((v1 - 4294967295)) | 1)))
                            var v18 uint = 255
                            if (v15 || v15) {
                                v19 := GenS1{F0: "hello"}
                                var v20 uint64 = 1
                                _ = v19
                                _ = v20
                            }
                            _ = v16
                            _ = v17
                            _ = v18
                        default:
                            switch (("" + "hello") + "foo bar") {
                            case "a":
                                var v21 string = ("foo bar" + ("hello" + (("\n\t" + "a") + (("hello" + "a") + "foo bar"))))
                                _ = v21
                            default:
                                switch (("hello" + "hello") + "a") {
                                case "\n\t":
                                    var v22 float32 = 0.1
                                    _ = v22
                                }
                            }
                        }
                        _ = v15
                    }
                    var v23 int16 = 32767
                    _ = v23
                }
                if (2 > 255) {
                    switch "foo bar" {
                    case "foo bar":
                        v3 = v3
                    case "":
                        var v24 bool = false
                        _ = v24
                    }
                    fmt.Println(v3, v4)
                } else {
                    v25 := GenS1{F0: ("foo bar" + "")}
                    v26 := v2[fi5]
                    v27 := GenS1{F0: v26}
                    var v28 float32 = 0.1
                    _ = v25
                    _ = v26
                    _ = v27
                    _ = v28
                }
            case "\n\t":
                v4 -= 352
            }
            v29 := GenS1{F0: "\n\t"}
            v30 := GenS1{F0: "\n\t"}
            _ = len(v2)
            var v31 uint64 = v1
            _ = fi6
            _ = v29
            _ = v30
            _ = v31
        }
        if false {
            v32 := GenS1{F0: ("hello" + ("" + ("" + "")))}
            var v33 uint16 = 42
            _ = v32
            _ = v33
        }
        _ = fi5
    }
    v34 := []int{-2147483648}
    v35 := GenS1{F0: "foo bar"}
    for fi36 := 0; fi36 < 14; fi36++ {
        var v37 uint16 = 0
        v38 := GenS2{F0: 4294967295}
        switch "foo bar" {
        case "\n\t":
            _ = len(v34)
        default:
            var v39 uint64 = v1
            _ = v39
        }
        _ = fi36
        _ = v37
        _ = v38
    }
    v40 := GenS2{F0: 18446744073709551615}
    var v41 uint8 = 0
    v4 = 279
    switch 469 {
    case 932:
        var v42 int16 = 1000
        _ = v42
    default:
        switch "foo bar" {
        case "a":
            v43 := GenS3{F0: 0, F1: 0, F2: (v1 << ((4294967295) & 63))}
            _ = v43
        case "":
            v34 = append(v34, -2147483648)
        default:
            var v44 float64 = v3
            _ = v44
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v40
    _ = v41
}
