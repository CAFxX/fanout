package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int }
func genFn0(p0 uint16) int64 {
    p0 += p0
    p0 += p0
    var v5 uint8 = 0
    _ = p0
    _ = v5
    return -1
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint64 = 598
    var v3 uint16 = 32767
    var v4 uint8 = 128
    for fi6 := 0; fi6 < 12; fi6++ {
        fmt.Println(v3, fi6, v4)
        fmt.Println(fi6, fi6)
        _ = fi6
    }
    v7 := []int{}
    var v8 float64 = ((0.0 / 3.14159) - (3.14159 + 0.1))
    v7 = append(v7, 374)
    var v9 uint32 = 2
    v10 := GenS1{F0: -844}
    switch 32767 {
    case 577:
        v11 := GenS1{F0: 9223372036854775807}
        _ = v11
    case 2147483647:
        v12 := GenS1{F0: int64(v2)}
        _ = v12
    case 9223372036854775807:
        v13 := GenS2{F0: uint(v9)}
        _ = v13
    }
    if false {
        switch ("a" + ((((((("foo bar" + "foo bar") + "a") + "foo bar") + "") + "") + ("a" + "foo bar")) + ("foo bar" + "foo bar"))) {
        case "hello":
            v14 := GenS3{F0: -2147483648}
            _ = v14
        case "foo bar":
            v3 += (v3 / (((v3 * v3)) | 1))
        case "\n\t":
            v15 := GenS4{F0: 90}
            _ = v15
        }
        if true {
            v4 = (^(v4))
            v2 = 204
            switch 0 {
            case -2147483648:
                v1["\n\t"] = (!(9223372036854775807 >= 32768))
            case 42:
                v3 = 255
            case 544:
                for fi16 := 0; fi16 < 6; fi16++ {
                    v17 := GenS4{F0: -145}
                    _ = fi16
                    _ = v17
                }
            }
        }
        switch 808 {
        case 42:
            if len(v7) > 0 {
                v7[0] = -843
            }
        case -502:
            if (0.1 >= ((v8 * v8) + (float64((((-2.25 * 0.1) - ((3.14159 / 1.5) - 0.0)) / 1e10)) / 0.1))) {
                var v18 float64 = v8
                _ = v18
            } else {
                switch 42 {
                case 127:
                    var v19 uint8 = 218
                    _ = v19
                case -1:
                    if ((3.14159 != v8) && ((3.14159 / ((0.0 / (float32(v8) + 1.5)) * 0.0)) < 0.1)) {
                        var v20 uint64 = v2
                        v2 -= 41
                        _ = v20
                    } else {
                        if len(v7) >= 2 {
                            v7 = v7[1:len(v7)]
                        }
                        var v21 float32 = ((float32(-2.25) + 1e10) * 0.1)
                        _ = v21
                    }
                }
                switch ("a" + ((("foo bar" + ((("\n\t" + "foo bar") + "\n\t") + "\n\t")) + "") + ("" + "hello"))) {
                case "":
                    if false {
                        fmt.Println(v2, v9, v8)
                    }
                case "hello":
                    v22 := GenS2{F0: 42}
                    _ = v22
                case "foo bar":
                    if (!(92 >= 444)) {
                        delete(v1, ("\n\t" + (("a" + (("\n\t" + (("foo bar" + ("\n\t" + "foo bar")) + "")) + ("hello" + ""))) + ("" + "foo bar"))))
                        v2 = (^(v2))
                    } else {
                        v1[(("foo bar" + "") + ("" + "hello"))] = ("hello" < "\n\t")
                        if true {
                            fmt.Println(v2)
                            switch "foo bar" {
                            case "":
                                v2 += v2
                            default:
                                if (!false) {
                                    var v23 uint8 = uint8(v3)
                                    var v24 uint32 = 4294967295
                                    _ = v23
                                    _ = v24
                                } else {
                                    fmt.Println(v9, v9)
                                    v3 += 2
                                }
                            }
                        } else {
                            v25 := GenS1{F0: 0}
                            if false {
                                v26 := GenS4{F0: 459}
                                if len(v7) >= 2 {
                                    v7 = v7[1:len(v7)]
                                }
                                v1[(("hello" + "") + "\n\t")] = false
                                _ = v26
                            }
                            v8 += (v8 / v8)
                            _ = v25
                        }
                    }
                default:
                    for fi27 := 0; fi27 < 19; fi27++ {
                        var v28 float64 = (((v8 - v8) * float64(0.1)) * v8)
                        _ = fi27
                        _ = v28
                    }
                }
                v2 += (v2 & 767)
            }
        case 2147483647:
            fmt.Println(v9, v3)
        }
    } else {
        v2 += 907
    }
    var v29 int8 = 1
    if (80 >= 1) {
        if ((v29 - v29) < v29) {
            _ = len(v7)
        }
        if true {
            v30 := GenS1{F0: int64(v9)}
            _ = v30
        }
    }
    if len(v7) > 0 {
        v7[0] = 32767
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
