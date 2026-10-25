package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint32; F1 int }
type GenS5 struct { F0 int; F1 string }
func genFn0(p0 string) bool {
    v5 := []int{-1, -9223372036854775808, 2147483647, -4}
    for fi6 := 0; fi6 < 19; fi6++ {
        p0 = p0
        var v7 float32 = 1.5
        v7 += float32((-2.25 / 1.5))
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = v5
    return (true && true)
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{843}
    var v3 int = 9223372036854775807
    var v4 int8 = -1
    fmt.Println(v3, v4)
    v4 = (-(v4))
    v8 := GenS1{F0: 42, F1: 1, F2: (v4 * (v4 * 127))}
    var v9 int32 = -2147483648
    fmt.Println(v9, v3)
    if true {
        for fi10 := 0; fi10 < 10; fi10++ {
            switch ("" + "hello") {
            case "hello":
                var v11 int8 = v4
                _ = v11
            case "foo bar":
                v12 := v1["\n\t"]
                _ = v12
            }
            if true {
                v4 = (v4 | v4)
                v13 := GenS1{F0: 2, F1: (v9 ^ (v9 & v9)), F2: v4}
                fmt.Println(v9, v9)
                _ = v13
            }
            v9 = (^(v9))
            _ = fi10
        }
        v9 += (v9 * v9)
        v14 := GenS2{F0: 384}
        _ = v14
    } else {
        for fi15 := 0; fi15 < 20; fi15++ {
            if false {
                v9 = v9
                v16 := GenS3{F0: 42, F1: 42}
                if (!true) {
                    v9 = (v9 >> ((2) & 63))
                    v3 -= fi15
                }
                _ = v16
            } else {
                if (0.0 != 3.14159) {
                    v17 := GenS4{F0: 676, F1: (fi15 << ((1) & 63))}
                    switch "hello" {
                    case "":
                        for fi18 := 0; fi18 < 15; fi18++ {
                            switch 128 {
                            case 0:
                                if (true || false) {
                                    var v19 string = ("hello" + "foo bar")
                                    if true {
                                        switch 0 {
                                        case -591:
                                            fmt.Println(v3)
                                        case 128:
                                            var v20 bool = false
                                            _ = v20
                                        }
                                        var v21 uint16 = 999
                                        fmt.Println(v3)
                                        _ = v21
                                    } else {
                                        var v22 int16 = 51
                                        _ = v22
                                    }
                                    _ = v19
                                }
                            case -53:
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(fi15)
                                case "hello":
                                    v9 = 1
                                }
                            default:
                                var v23 uint8 = 128
                                _ = v23
                            }
                            var v24 uint8 = 128
                            _ = len(v1)
                            var v25 int8 = v4
                            fmt.Println(fi18, v9, v24)
                            var v26 uint32 = 0
                            _ = fi18
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                    case "foo bar":
                        for fi27 := 0; fi27 < 12; fi27++ {
                            switch "hello" {
                            case "a":
                                if (!true) {
                                    var v28 uint16 = 42
                                    fmt.Println(fi27, v3, fi27)
                                    switch "" {
                                    case "foo bar":
                                        v3 = (fi27 % (((fi15 * -2147483648)) | 1))
                                    default:
                                        switch 32767 {
                                        case 127:
                                            var v29 uint = 425
                                            _ = v29
                                        default:
                                            fmt.Println(v28, fi27)
                                        }
                                    }
                                    _ = v28
                                } else {
                                    var v30 int64 = 0
                                    fmt.Println(fi15)
                                    _ = v30
                                }
                            case "\n\t":
                                var v31 int32 = (v9 + (v9 % (((v9 >> ((412) & 63))) | 1)))
                                _ = v31
                            case "hello":
                                var v32 int16 = 0
                                _ = v32
                            }
                            var v33 int = 1
                            v34 := GenS3{F0: fi15, F1: 4294967295}
                            _ = fi27
                            _ = v33
                            _ = v34
                        }
                    case "\n\t":
                        var v35 string = ("\n\t" + ("" + ("\n\t" + "a")))
                        _ = v35
                    default:
                        for fi36 := 0; fi36 < 13; fi36++ {
                            var v37 int32 = v9
                            _ = fi36
                            _ = v37
                        }
                    }
                    for fi38 := 0; fi38 < 16; fi38++ {
                        v4 = 42
                        _ = fi38
                    }
                    _ = v17
                }
                v39 := v1["a"]
                _ = v39
            }
            _ = fi15
        }
        v40 := GenS5{F0: (^(v3)), F1: ""}
        _ = v40
    }
    for fi41 := 0; fi41 < 3; fi41++ {
        fmt.Println(v9, v9, v9)
        _ = fi41
    }
    if (false || false) {
        v3 = (v3 >> ((42) & 63))
    } else {
        var v42 uint16 = 1
        var v43 int16 = 1000
        v1["\n\t"] = (v42 > v42)
        _ = v42
        _ = v43
    }
    _ = len(v2)
    v4 = (v4 - v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
