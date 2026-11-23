package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64 }
type GenS2 struct { F0 int16; F1 string; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint32; F2 string }
func genFn0(p0 string) int32 {
    var v4 uint64 = 0
    v4 += (v4 << ((42) & 63))
    _ = p0
    _ = v4
    return -94
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    var v3 string = "foo bar"
    for fi5 := 0; fi5 < 16; fi5++ {
        var v6 int8 = 0
        if false {
            v7 := v2[v3]
            if (!(18446744073709551615 > 288)) {
                v6 -= v6
                v8 := GenS1{F0: -671, F1: -9223372036854775808}
                _ = v8
            }
            _ = v7
        } else {
            _ = len(v1)
            switch ("\n\t" + ((v3 + ((v3 + v3) + ("hello" + v3))) + v3)) {
            case "\n\t":
                switch -456 {
                case 139:
                    if ((2147483647 != 0) || (!(127 > 58))) {
                        v3 = ""
                        v3 = "hello"
                    }
                case 545:
                    v9 := GenS1{F0: 129, F1: -9223372036854775808}
                    _ = v9
                case 128:
                    v10 := GenS2{F0: 138, F1: ("a" + v3), F2: 624}
                    _ = v10
                default:
                    if (!true) {
                        if false {
                            if false {
                                v3 = v3
                            }
                        }
                    } else {
                        var v11 uint = 0
                        var v12 int32 = 288
                        v11 += 42
                        _ = v11
                        _ = v12
                    }
                }
            case "foo bar":
                var v13 float32 = 0.1
                _ = v13
            default:
                if false {
                    switch (-(fi5)) {
                    case 4294967295:
                        if (false && ((42 >= 664) && true)) {
                            v14 := []int{0, fi5}
                            v15 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807}
                            _ = v14
                            _ = v15
                        }
                    case 42:
                        switch 0 {
                        case 32768:
                            v3 = (v3 + "")
                        case 42:
                            switch ("a" + (v3 + "")) {
                            case "":
                                fmt.Println(fi5)
                            case "foo bar":
                                if true {
                                    v16 := []int{fi5}
                                    v17 := GenS1{F0: 1, F1: 613}
                                    _ = v16
                                    _ = v17
                                } else {
                                    var v18 int16 = 320
                                    var v19 uint8 = 61
                                    _ = v18
                                    _ = v19
                                }
                            case "\n\t":
                                var v20 uint64 = 42
                                _ = v20
                            default:
                                switch 42 {
                                case 32768:
                                    var v21 uint = 42
                                    _ = v21
                                default:
                                    v3 = (v3 + v3)
                                }
                            }
                        case -750:
                            fmt.Println(v3)
                        }
                    case 127:
                        v6 = (v6 >> ((4294967295) & 63))
                    }
                }
            }
            var v22 int16 = 42
            _ = v22
        }
        _ = fi5
        _ = v6
    }
    v23 := []int{128, -109, 127}
    switch -2147483648 {
    case 4294967295:
        switch "\n\t" {
        case "\n\t":
            v3 = "a"
        default:
            v3 = (((v3 + (v3 + "hello")) + "") + "a")
        }
    case 128:
        var v24 int = -709
        _ = v24
    case 32768:
        var v25 bool = false
        _ = v25
    }
    var v26 int64 = 42
    _ = len(v23)
    v3 = (v3 + v3)
    fmt.Println(v26, v3)
    var v27 uint64 = 772
    var v28 float32 = (0.0 * 0.0)
    switch v3 {
    case "hello":
        v29 := GenS3{F0: (^(v26)), F1: 2147483648, F2: "hello"}
        _ = v29
    case "a":
        v26 += 42
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
