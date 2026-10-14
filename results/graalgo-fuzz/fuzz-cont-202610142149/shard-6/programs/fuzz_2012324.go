package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16; F1 bool; F2 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 string }

func main() {
    v1 := make(map[int]string)
    v2 := []string{"a"}
    var v3 uint64 = 14
    var v4 uint16 = 65535
    var v5 bool = (false && false)
    v6 := []uint32{uint32(v4), 2}
    v3 = (v3 & 42)
    for fi7 := 0; fi7 < 8; fi7++ {
        for fi8 := 0; fi8 < 3; fi8++ {
            v5 = v5
            if (!v5) {
                fmt.Println(v3, v4)
                v9 := GenS1{F0: 1}
                var v10 uint8 = 128
                _ = v10
                _ = v9
            } else {
                var v11 string = ""
                v12 := GenS2{F0: 1000, F1: false, F2: -207}
                _ = v11
                _ = v12
            }
            var v13 float64 = 1.5
            _ = fi8
            _ = v13
        }
        v14 := GenS1{F0: 1000}
        _ = fi7
        _ = v14
    }
    v5 = (!v5)
    fmt.Println(v5, v4, v4)
    v5 = true
    fmt.Println(v4, v5)
    switch "" {
    case "\n\t":
        var v15 uint16 = (v4 + v4)
        _ = v15
    default:
        var v16 int = 42
        _ = v16
    }
    if true {
        if true {
            var v17 bool = (128 != 127)
            _ = v17
        }
    }
    if v5 {
        fmt.Println(v4, v4)
        if v5 {
            v3 = v3
            v4 = (v4 % (((v4 + v4)) | 1))
            switch 2147483647 {
            case -1:
                v5 = true
            case 32767:
                v1[4294967295] = "hello"
            default:
                for fi18 := 0; fi18 < 7; fi18++ {
                    fmt.Println(v4, fi18)
                    var v19 uint = 255
                    for fi20 := 0; fi20 < 13; fi20++ {
                        if (v5 && v5) {
                            if (178 > 2) {
                                var v21 int8 = 100
                                v22 := GenS2{F0: 1, F1: false, F2: 42}
                                _ = v21
                                _ = v22
                            } else {
                                switch "foo bar" {
                                case "a":
                                    v23 := GenS3{F0: -860, F1: 4294967295, F2: "\n\t"}
                                    _ = v23
                                case "foo bar":
                                    switch fi20 {
                                    case 477:
                                        switch "hello" {
                                        case "\n\t":
                                            v4 = (v4 & v4)
                                        case "a":
                                            fmt.Println(v4, v5, v19)
                                            v4 -= (v4 + (v4 & (v4 & (v4 ^ (v4 << ((18446744073709551615) & 63))))))
                                            var v24 string = ""
                                            _ = v24
                                        case "":
                                            switch ("hello" + "") {
                                            case "hello":
                                                v1[fi20] = "a"
                                            case "\n\t":
                                                v25 := GenS2{F0: -1, F1: false, F2: 2147483647}
                                                _ = v25
                                            }
                                        }
                                    case 127:
                                        var v26 bool = (18446744073709551615 == (^(v3)))
                                        _ = v26
                                    case -565:
                                        fmt.Println(v5)
                                    }
                                case "\n\t":
                                    _ = len(v1)
                                }
                                var v27 float64 = -2.25
                                if (!false) {
                                    v28 := GenS1{F0: 42}
                                    v29 := GenS2{F0: 0, F1: true, F2: 42}
                                    _ = v28
                                    _ = v29
                                } else {
                                    var v30 int16 = -943
                                    _ = v30
                                }
                                _ = v27
                            }
                        } else {
                            v31 := v1[(fi20 << ((227) & 63))]
                            _ = v31
                        }
                        _ = fi20
                    }
                    _ = fi18
                    _ = v19
                }
            }
        }
        var v32 string = (("a" + "\n\t") + "hello")
        _ = v32
    }
    var v33 uint16 = (^(v4))
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    fmt.Println(v4, v4, v33)
    switch (("a" + (("" + ("" + "")) + "hello")) + (("foo bar" + "hello") + ("a" + ((("a" + (("foo bar" + ((("hello" + ("a" + "hello")) + "foo bar") + "a")) + ("\n\t" + "foo bar"))) + (("\n\t" + ("\n\t" + ("a" + "\n\t"))) + (("\n\t" + "hello") + ""))) + "a")))) {
    case "foo bar":
        if (v5 && ("\n\t" > "foo bar")) {
            v1[-2147483648] = "a"
            v4 -= v4
            var v34 uint8 = 42
            _ = v34
        } else {
            if v5 {
                var v35 uint16 = (v33 % ((42) | 1))
                if v5 {
                    _ = len(v6)
                } else {
                    v5 = false
                }
                _ = v35
            } else {
                v3 = 42
            }
            v33 = v4
            fmt.Println(v5, v33, v33)
        }
    }
    v4 += 58
    for fi36 := 0; fi36 < 11; fi36++ {
        if true {
            v37 := GenS3{F0: -1, F1: 18446744073709551615, F2: "a"}
            switch "foo bar" {
            case "foo bar":
                _ = len(v1)
            default:
                v4 -= (v33 ^ (v4 | 42))
            }
            _ = v37
        }
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
