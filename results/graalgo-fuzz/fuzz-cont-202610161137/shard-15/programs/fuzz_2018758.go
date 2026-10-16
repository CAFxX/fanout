package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint16; F1 uint64 }
func genFn0(p0 int64, p1 int32) int64 {
    if true {
        var v6 int16 = 42
        fmt.Println(p1, v6, p1)
        _ = v6
    }
    v7 := []int{0, 1, 127}
    _ = p0
    _ = p1
    _ = v7
    return -9223372036854775808
}
func genFn1(p0 int) string {
    if true {
        v8 := []int{p0}
        var v9 bool = true
        _ = v8
        _ = v9
    } else {
        p0 -= (p0 - 9223372036854775807)
        var v10 bool = (-1 > 2147483647)
        _ = v10
    }
    var v11 int8 = -128
    _ = p0
    _ = v11
    return "\n\t"
}

func main() {
    v1 := []int{4294967295}
    var v2 string = "hello"
    var v3 bool = ((false && false) || (!false))
    var v4 int = 32767
    v5 := []int{v4, 42}
    if v3 {
        var v12 uint64 = 42
        v1 = append(v1, v4)
        _ = v12
    } else {
        var v13 int16 = 534
        _ = len(v1)
        v14 := GenS1{F0: ("" + "\n\t")}
        _ = v13
        _ = v14
    }
    if len(v1) > 0 {
        v1[0] = 127
    }
    var v15 bool = (((v2 + ((("" + (v2 + "")) + (v2 + v2)) + (v2 + ""))) + "a") == "hello")
    if (724 >= 395) {
        fmt.Println(v15, v15, v4)
    } else {
        v3 = v3
    }
    switch (v4 / (((v4 * 128)) | 1)) {
    case 1:
        var v16 uint64 = 491
        _ = v16
    case 2147483647:
        for fi17 := 0; fi17 < 18; fi17++ {
            for fi18 := 0; fi18 < 20; fi18++ {
                if (!v15) {
                    if (false && (2 >= 127)) {
                        var v19 int32 = -1
                        switch -1 {
                        case 1:
                            var v20 int32 = v19
                            _ = v20
                        case 0:
                            fmt.Println(fi18)
                        case -2147483648:
                            v15 = v3
                            v21 := GenS1{F0: "hello"}
                            var v22 uint16 = 121
                            _ = v21
                            _ = v22
                        }
                        var v23 uint = 591
                        _ = v19
                        _ = v23
                    } else {
                        if ((-2147483648 == 125) || (!(!((2 >= 0) && (false && v3))))) {
                            if false {
                                fmt.Println(fi17, fi17, v15)
                                fmt.Println(v15)
                            }
                            var v24 float64 = 0.0
                            switch "foo bar" {
                            case "":
                                v2 = "a"
                            }
                            _ = v24
                        }
                        v3 = (!(1.5 < ((float64(((0.0 * 1.5) - (0.1 - (1e10 * float32(1e10))))) * float64(0.1)) - (((((3.14159 * 0.0) * 0.1) / 1.0) * float64(((float32(1e10) * 1.5) * 3.14159))) * 1e10))))
                        v25 := GenS2{F0: 763, F1: (v2 + ("\n\t" + ("a" + v2)))}
                        v26 := GenS1{F0: v2}
                        v4 = 0
                        _ = v25
                        _ = v26
                    }
                } else {
                    var v27 uint16 = 255
                    _ = v27
                }
                if len(v5) > 0 {
                    v5[0] = -509
                }
                _ = fi18
            }
            _ = fi17
        }
    default:
        v1 = append(v1, (v4 ^ (^(v4))))
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        var v29 int = 503
        v30 := GenS3{F0: 0, F1: 18446744073709551615}
        _ = fi28
        _ = v29
        _ = v30
    }
    v2 = ((v2 + (v2 + ((v2 + "foo bar") + v2))) + (("hello" + v2) + (v2 + (((("hello" + v2) + "\n\t") + v2) + "foo bar"))))
    v31 := GenS1{F0: "\n\t"}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
}
