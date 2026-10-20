package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint16 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int, p1 uint16, p2 uint8) string {
    var v6 uint16 = p1
    v7 := []int{}
    var v8 bool = ("a" != (("hello" + ("foo bar" + "\n\t")) + (("foo bar" + "\n\t") + (("a" + "\n\t") + "\n\t"))))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    return (("" + "a") + ("" + ("foo bar" + "\n\t")))
}

func main() {
    var v1 float64 = (3.14159 + 0.1)
    v2 := []string{(("hello" + "") + "")}
    var v3 bool = (81 > 284)
    v4 := make(map[int]int)
    var v5 uint64 = 912
    if (!v3) {
        switch int(v5) {
        case -543:
            for fi9 := 0; fi9 < 19; fi9++ {
                switch 32768 {
                case 4294967295:
                    v1 += ((v1 - ((1.5 * 3.14159) * v1)) - (v1 + float64(1e10)))
                case -1:
                    var v10 int8 = 0
                    _ = v10
                case -952:
                    switch fi9 {
                    case 32768:
                        _ = len(v4)
                    case -9223372036854775808:
                        v11 := GenS1{F0: -674, F1: fi9, F2: 65535}
                        _ = v11
                    }
                default:
                    if v3 {
                        v12 := GenS1{F0: 42, F1: fi9, F2: 97}
                        fmt.Println(v1)
                        _ = v12
                    }
                }
                _ = fi9
            }
        case 88:
            v4[-615] = 127
        case 693:
            v3 = true
        default:
            var v13 uint64 = 4294967295
            _ = v13
        }
        fmt.Println(v5)
        var v14 uint = 38
        _ = v14
    }
    var v15 string = (("" + ("a" + "foo bar")) + "hello")
    v16 := GenS2{F0: -1}
    if len(v2) > 0 {
        v2[0] = "hello"
    }
    v17 := GenS1{F0: 872, F1: 42, F2: 42}
    var v18 uint8 = 42
    switch -9223372036854775808 {
    case 127:
        var v19 uint16 = 1
        _ = v19
    case -970:
        for fi20 := 0; fi20 < 19; fi20++ {
            var v21 uint8 = v18
            switch v15 {
            case "":
                var v22 int8 = 82
                _ = v22
            case "hello":
                v23 := GenS1{F0: 233, F1: (fi20 / (((fi20 | (fi20 / ((int(v21)) | 1)))) | 1)), F2: 65535}
                _ = v23
            case "\n\t":
                v15 = v15
            default:
                var v24 int32 = 0
                _ = v24
            }
            switch 9223372036854775807 {
            case 1:
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            case 9223372036854775807:
                var v25 int8 = 54
                _ = v25
            default:
                for fi26 := 0; fi26 < 18; fi26++ {
                    if false {
                        var v27 uint8 = v21
                        _ = v27
                    } else {
                        if v3 {
                            switch 4294967295 {
                            case 1:
                                if (v5 == 42) {
                                    fmt.Println(v1, v18)
                                } else {
                                    fmt.Println(v1)
                                }
                            case 9223372036854775807:
                                v15 = ("foo bar" + "foo bar")
                            case 0:
                                v28 := GenS2{F0: 1}
                                _ = v28
                            default:
                                switch 9223372036854775807 {
                                case -2147483648:
                                    switch v15 {
                                    case "":
                                        var v29 string = (((v15 + v15) + v15) + v15)
                                        _ = v29
                                    case "\n\t":
                                        v18 = v18
                                    case "hello":
                                        var v30 string = "foo bar"
                                        _ = v30
                                    default:
                                        if true {
                                            v5 += 0
                                            fmt.Println(v3)
                                        }
                                    }
                                default:
                                    _ = len(v2)
                                }
                                var v31 int64 = 0
                                _ = v31
                            }
                            if v3 {
                                fmt.Println(v15, v18, v18)
                            } else {
                                if (((v1 + (v1 / v1)) < v1) && (100 <= int8(v5))) {
                                    v21 -= 218
                                }
                                switch v15 {
                                case "a":
                                    var v32 float64 = v1
                                    _ = v32
                                default:
                                    v1 += v1
                                }
                            }
                        }
                        fmt.Println(v15)
                    }
                    v5 = (v5 >> ((736) & 63))
                    fmt.Println(v5, v5)
                    var v33 bool = false
                    _ = fi26
                    _ = v33
                }
            }
            _ = fi20
            _ = v21
        }
    default:
        fmt.Println(v5, v18)
    }
    var v34 int64 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
