package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int; F1 uint8; F2 uint32 }
func genFn0(p0 uint) bool {
    var v7 uint32 = 351
    if false {
        fmt.Println(p0)
        p0 = (p0 * (^(p0)))
    }
    _ = p0
    _ = v7
    return (0 < 32767)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]bool)
    var v3 uint16 = 134
    var v4 uint = 2
    var v5 bool = false
    v6 := make(map[string]bool)
    switch 4294967295 {
    case 32768:
        var v8 int32 = 42
        _ = v8
    case -506:
        if (false || v5) {
            if ((!(!("a" <= ((((("hello" + "\n\t") + "a") + "foo bar") + "\n\t") + "foo bar")))) || v5) {
                switch ("hello" + "a") {
                case "hello":
                    switch 471 {
                    case 127:
                        if false {
                            var v9 int64 = -322
                            _ = v9
                        }
                    case 32767:
                        v5 = false
                    default:
                        fmt.Println(v4, v3)
                    }
                case "a":
                    if ((2 > 0) && (42 != 42)) {
                        if v5 {
                            v5 = true
                        } else {
                            switch ("hello" + ("" + "a")) {
                            case "a":
                                switch "" {
                                case "hello":
                                    v6[""] = v5
                                }
                            case "foo bar":
                                var v10 uint8 = 255
                                _ = v10
                            case "":
                                var v11 bool = v5
                                _ = v11
                            default:
                                v5 = (!((1.5 * -2.25) != -2.25))
                            }
                        }
                    } else {
                        var v12 uint16 = 32767
                        for fi13 := 0; fi13 < 7; fi13++ {
                            v5 = ((v5 && v5) || (1 <= fi13))
                            for fi14 := 0; fi14 < 3; fi14++ {
                                v3 += (v12 * (v3 / ((v3) | 1)))
                                var v15 uint32 = 886
                                var v16 uint64 = 932
                                _ = fi14
                                _ = v15
                                _ = v16
                            }
                            for fi17 := 0; fi17 < 2; fi17++ {
                                var v18 string = "\n\t"
                                v19 := v2[2147483647]
                                v5 = true
                                v20 := v6[""]
                                v21 := GenS1{F0: v5, F1: 1000}
                                v18 = ""
                                _ = fi17
                                _ = v18
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = fi13
                        }
                        _ = v12
                    }
                }
            } else {
                fmt.Println(v4)
                v3 = (v3 << ((v4) & 63))
                var v22 uint16 = 32767
                _ = v22
            }
            var v23 uint64 = 32
            _ = v23
        } else {
            for fi24 := 0; fi24 < 2; fi24++ {
                var v25 int8 = 17
                _ = fi24
                _ = v25
            }
        }
    case 1:
        v4 = (v4 & (v4 + 65535))
    }
    var v26 uint16 = v3
    fmt.Println(v26, v4)
    v27 := []int{-897}
    if v5 {
        v6["foo bar"] = (!(v5 || (!v5)))
    } else {
        fmt.Println(v26, v26)
        var v28 string = ("foo bar" + "a")
        _ = v28
    }
    v29 := GenS1{F0: v5, F1: 42}
    v4 = v4
    delete(v1, 872)
    v27 = append(v27, 32767)
    fmt.Println(v3, v4)
    var v30 int32 = 42
    var v31 float64 = (1.5 * (0.0 + (float64(-2.25) - 3.14159)))
    switch 128 {
    case 127:
        fmt.Println(v31, v26)
    default:
        if v5 {
            var v32 bool = (!v5)
            _ = v32
        }
    }
    switch -180 {
    case -9223372036854775808:
        switch 128 {
        case 42:
            var v33 int8 = 76
            _ = v33
        default:
            v3 = (v3 >> (((^(v4))) & 63))
        }
    default:
        v3 = (v26 >> ((807) & 63))
    }
    v34 := GenS2{F0: 42, F1: 128, F2: 242}
    v35 := GenS1{F0: v5, F1: 32767}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
