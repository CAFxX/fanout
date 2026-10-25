package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int32 }

func main() {
    var v1 float64 = 3.14159
    v2 := []string{"a", "foo bar"}
    var v3 uint64 = 284
    v4 := []int{-2147483648, 342, 642}
    fmt.Println(v1)
    fmt.Println(v3)
    v5 := GenS1{F0: ("" + ("\n\t" + "hello"))}
    v1 -= (((v1 - v1) * 1.5) - float64(1.5))
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS2{F0: v3, F1: -689}
        v8 := GenS2{F0: (v3 & 837), F1: 953}
        if len(v2) > 0 {
            v2[0] = ((("hello" + ("" + ("" + "a"))) + "\n\t") + "\n\t")
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v9 int32 = 2147483647
    if false {
        switch ("\n\t" + "a") {
        case "a":
            var v10 float64 = -2.25
            _ = v10
        case "hello":
            if (2 >= 247) {
                v11 := GenS1{F0: "\n\t"}
                var v12 uint = 1
                _ = v11
                _ = v12
            } else {
                v13 := GenS1{F0: "foo bar"}
                for fi14 := 0; fi14 < 20; fi14++ {
                    fmt.Println(v3, v1, v1)
                    var v15 string = ""
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
        case "\n\t":
            var v16 int64 = 42
            _ = v16
        default:
            fmt.Println(v1, v3, v3)
        }
        var v17 string = ("" + "")
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = v17
    } else {
        _ = len(v2)
        fmt.Println(v9, v9)
    }
    if false {
        v1 += v1
        var v18 uint8 = 255
        _ = v18
    }
    if (255 != 128) {
        if true {
            if len(v4) > 0 {
                v4[0] = int(v9)
            }
            var v19 float64 = v1
            _ = v19
        }
        switch "a" {
        case "\n\t":
            for fi20 := 0; fi20 < 8; fi20++ {
                for fi21 := 0; fi21 < 20; fi21++ {
                    if (!true) {
                        var v22 uint16 = 255
                        switch ("hello" + "a") {
                        case "foo bar":
                            fmt.Println(v22, fi21)
                        case "":
                            v1 = (float64(1.5) - float64(0.0))
                        case "\n\t":
                            switch ("a" + ("a" + "hello")) {
                            case "":
                                var v23 int32 = v9
                                _ = v23
                            case "\n\t":
                                v22 = v22
                            }
                        default:
                            v22 = 2
                        }
                        _ = v22
                    } else {
                        switch fi20 {
                        case -9223372036854775808:
                            fmt.Println(v1, fi21, v1)
                        case 567:
                            v1 += (v1 * 3.14159)
                        }
                        fmt.Println(fi21, fi20)
                        v3 -= (v3 ^ (^(v3)))
                    }
                    _ = fi21
                }
                fmt.Println(v9, fi20, v1)
                var v24 int8 = 42
                _ = fi20
                _ = v24
            }
        }
    }
    v3 -= (v3 >> ((1) & 63))
    v9 -= v9
    v1 -= v1
    var v25 uint64 = (v3 * v3)
    var v26 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
