package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint8; F1 uint32 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]int)
    var v3 int64 = 0
    var v4 uint16 = 32767
    v5 := []int{42, 42}
    if true {
        var v6 uint32 = 158
        _ = v6
    }
    var v7 uint8 = 0
    v3 -= (v3 ^ v3)
    v7 -= 2
    switch ("" + "hello") {
    case "a":
        v7 = 218
    case "\n\t":
        var v8 int = 1
        _ = v8
    }
    var v9 int64 = v3
    var v10 bool = true
    if false {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 float64 = 0.0
            for fi13 := 0; fi13 < 11; fi13++ {
                if len(v5) > 0 {
                    v5[0] = fi13
                }
                if (v10 && v10) {
                    var v14 bool = false
                    _ = len(v5)
                    switch "hello" {
                    case "\n\t":
                        if false {
                            var v15 int8 = -81
                            v16 := GenS1{F0: 213}
                            var v17 int32 = 566
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                    case "hello":
                        if len(v5) > 0 {
                            v5[0] = fi13
                        }
                    case "foo bar":
                        var v18 int32 = 361
                        _ = v18
                    }
                    _ = v14
                } else {
                    switch fi13 {
                    case 32767:
                        if ((!(1 >= 42)) || v10) {
                            var v19 uint = 4294967295
                            v19 -= (^(v19))
                            if ((v4 | 1) == (v4 % ((uint16(v7)) | 1))) {
                                fmt.Println(v19, v19, v9)
                            } else {
                                v9 = (v3 + v3)
                            }
                            v20 := GenS1{F0: v4}
                            switch "foo bar" {
                            case "":
                                v19 = v19
                            }
                            _ = v19
                            _ = v20
                        } else {
                            if len(v5) > 0 {
                                v5[0] = (fi13 % ((-418) | 1))
                            }
                        }
                    case 454:
                        var v21 float32 = ((-2.25 + 3.14159) / 1.5)
                        _ = v21
                    case 1:
                        var v22 uint16 = v4
                        _ = v22
                    default:
                        var v23 uint8 = 128
                        _ = v23
                    }
                    v5 = append(v5, -2147483648)
                }
                v3 -= (v3 ^ (v3 - v3))
                _ = fi13
            }
            v24 := GenS2{F0: 100}
            _ = fi11
            _ = v12
            _ = v24
        }
        if ((v10 || v10) && (42 != 2147483648)) {
            var v25 int32 = 2147483647
            switch "foo bar" {
            case "\n\t":
                fmt.Println(v25, v3, v3)
            case "":
                v9 -= v3
            case "foo bar":
                v25 -= v25
            }
            switch ("a" + "") {
            case "\n\t":
                if (!((v9 | 42) > (v9 + v3))) {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        for fi27 := 0; fi27 < 8; fi27++ {
                            v28 := GenS1{F0: v4}
                            _ = fi27
                            _ = v28
                        }
                        var v29 uint32 = 2147483647
                        _ = fi26
                        _ = v29
                    }
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        v25 = v25
                    default:
                        fmt.Println(v7)
                    }
                    for fi30 := 0; fi30 < 11; fi30++ {
                        for fi31 := 0; fi31 < 16; fi31++ {
                            if (-71 == -1) {
                                var v32 uint = 2
                                v5 = append(v5, (fi30 % (((fi31 << (((v32 * v32)) & 63))) | 1)))
                                v33 := GenS3{F0: false}
                                _ = v32
                                _ = v33
                            }
                            var v34 int = (fi30 ^ (fi30 << ((2) & 63)))
                            _ = fi31
                            _ = v34
                        }
                        fmt.Println(v7, fi30)
                        if (!v10) {
                            delete(v1, (-(fi30)))
                        }
                        _ = fi30
                    }
                } else {
                    if ((("a" > "foo bar") || v10) && ((!v10) || (0 < -19))) {
                        v10 = ("" <= "")
                        var v35 int32 = v25
                        _ = v35
                    }
                    v9 = (v9 + 1)
                    v10 = ("" <= "hello")
                }
            }
            _ = v25
        }
        switch (("" + (("hello" + "\n\t") + "a")) + ("hello" + (("foo bar" + (("hello" + "a") + "foo bar")) + ""))) {
        case "":
            delete(v2, "a")
        case "a":
            v10 = (v10 || v10)
        case "\n\t":
            var v36 bool = (float32(1.5) >= float32(0.0))
            _ = v36
        default:
            v37 := GenS4{F0: (v7 % ((255) | 1)), F1: 2}
            _ = v37
        }
    } else {
        for fi38 := 0; fi38 < 2; fi38++ {
            var v39 float32 = 0.1
            v7 -= 1
            _ = fi38
            _ = v39
        }
        for fi40 := 0; fi40 < 10; fi40++ {
            fmt.Println(v10, v10)
            _ = fi40
        }
        fmt.Println(v7, v4, v9)
    }
    if true {
        delete(v1, 670)
        var v41 string = ""
        _ = v41
    } else {
        var v42 float32 = 0.0
        var v43 string = ""
        _ = v42
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
