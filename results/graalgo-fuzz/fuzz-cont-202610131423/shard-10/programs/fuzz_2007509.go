package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 int32; F2 int32 }
type GenS4 struct { F0 uint8 }
func genFn0() string {
    if (42 <= 1) {
        var v4 uint8 = 0
        v5 := GenS1{F0: -226}
        v6 := GenS2{F0: ("a" + "")}
        _ = v4
        _ = v5
        _ = v6
    }
    v7 := GenS3{F0: 22, F1: -770, F2: -2147483648}
    var v8 int16 = -32768
    v8 = (^(v8))
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint32 = 1
    var v3 int16 = 1
    if (!(3.14159 != 1.5)) {
        if true {
            var v9 uint64 = uint64(v2)
            v10 := GenS3{F0: (v3 >> ((333) & 63)), F1: -1, F2: -2147483648}
            if false {
                var v11 bool = (!false)
                if false {
                    if v11 {
                        switch "hello" {
                        case "hello":
                            v12 := []int{-1, 813, 4294967295}
                            _ = v12
                        case "":
                            v11 = v11
                        case "\n\t":
                            fmt.Println(v9)
                        default:
                            var v13 int64 = -9223372036854775808
                            _ = v13
                        }
                        fmt.Println(v3)
                    } else {
                        if (-1 < (v3 - 1)) {
                            var v14 int16 = v3
                            _ = v14
                        }
                        var v15 int32 = 2147483647
                        _ = v15
                    }
                    if (v11 && v11) {
                        v2 += v2
                        var v16 float64 = 1e10
                        if (!(v11 || false)) {
                            var v17 int64 = 715
                            var v18 uint8 = 1
                            _ = v17
                            _ = v18
                        } else {
                            switch -2147483648 {
                            case 32768:
                                fmt.Println(v3)
                            case 152:
                                switch "\n\t" {
                                case "a":
                                    if v11 {
                                        if ((!v11) || false) {
                                            var v19 int64 = 723
                                            _ = v19
                                        }
                                        v9 -= 42
                                    }
                                case "foo bar":
                                    var v20 string = "\n\t"
                                    _ = v20
                                default:
                                    _ = len(v1)
                                }
                            case -1:
                                var v21 float32 = ((0.1 + 0.1) + 0.0)
                                _ = v21
                            default:
                                var v22 int32 = 42
                                _ = v22
                            }
                        }
                        _ = v16
                    }
                    if v11 {
                        fmt.Println(v11, v11, v11)
                    }
                }
                _ = v11
            }
            _ = v10
            _ = v9
        } else {
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 int8 = 42
                var v25 uint16 = 207
                _ = fi23
                _ = v24
                _ = v25
            }
        }
        v3 += -554
    } else {
        var v26 int8 = -1
        _ = v26
    }
    v27 := []int{9223372036854775807, -1}
    switch -690 {
    case 128:
        var v28 uint = 403
        _ = v28
    case -612:
        var v29 uint16 = 2
        _ = v29
    case 9223372036854775807:
        switch ("" + "hello") {
        case "a":
            v2 -= v2
        default:
            switch "a" {
            case "foo bar":
                v1[("hello" + ((("\n\t" + "a") + "a") + ("\n\t" + (("hello" + "a") + "a"))))] = v2
            case "a":
                v3 = (v3 + v3)
            }
        }
    default:
        for fi30 := 0; fi30 < 17; fi30++ {
            if false {
                fmt.Println(v2, fi30, v3)
            }
            _ = fi30
        }
    }
    var v31 int32 = -1
    if len(v27) >= 2 {
        v27 = v27[1:len(v27)]
    }
    var v32 uint8 = 128
    v33 := GenS4{F0: v32}
    switch ("foo bar" + "hello") {
    case "a":
        v34 := GenS4{F0: v32}
        _ = v34
    case "foo bar":
        var v35 string = (("foo bar" + ("a" + "hello")) + "")
        _ = v35
    }
    var v36 uint = 553
    var v37 bool = false
    var v38 int64 = 9223372036854775807
    var v39 int16 = 1000
    fmt.Println(v39)
    var v40 string = ("" + "foo bar")
    v36 -= (v36 + v36)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v40
}
