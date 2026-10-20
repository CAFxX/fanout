package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int64 }
func genFn0(p0 uint16, p1 int32) string {
    v7 := []int{0, -2147483648, 0}
    v8 := GenS1{F0: false, F1: 636}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := []string{"a", "foo bar"}
    var v2 int16 = 1000
    v3 := make(map[string]string)
    var v4 uint16 = 65535
    v5 := make(map[uint32]string)
    var v6 uint16 = (v4 | v4)
    switch 127 {
    case 9223372036854775807:
        if false {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    case 0:
        for fi9 := 0; fi9 < 9; fi9++ {
            if true {
                if (true && true) {
                    if true {
                        switch "hello" {
                        case "":
                            if (!(-1 < 326)) {
                                v2 = v2
                                v6 += (v6 >> ((1) & 63))
                            } else {
                                _ = len(v1)
                                fmt.Println(v2, v2)
                            }
                        default:
                            fmt.Println(v6, v6, fi9)
                        }
                    } else {
                        if (!true) {
                            switch "" {
                            case "hello":
                                v4 = v4
                            case "foo bar":
                                _ = len(v1)
                            case "":
                                if true {
                                    v3[""] = ""
                                }
                            }
                            var v10 uint64 = 2
                            _ = v10
                        }
                    }
                }
                v11 := v5[669]
                var v12 uint64 = 926
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    case -1:
        if false {
            var v13 uint8 = 0
            v4 = v4
            v6 += v4
            _ = v13
        }
    default:
        v14 := GenS1{F0: (("hello" + (("\n\t" + (("a" + "foo bar") + "foo bar")) + "hello")) > (("hello" + "hello") + "hello")), F1: (v2 << ((388) & 63))}
        _ = v14
    }
    if (0.1 < ((float32(0.1) - 0.0) + (3.14159 - 0.1))) {
        var v15 uint = 984
        switch "" {
        case "\n\t":
            var v16 int16 = (v2 & (v2 >> ((v15) & 63)))
            _ = v16
        case "foo bar":
            v4 = v4
        }
        _ = v15
    } else {
        v4 -= (v6 / ((v6) | 1))
        switch ("hello" + ("\n\t" + (("foo bar" + "hello") + "hello"))) {
        case "\n\t":
            v2 -= (v2 - -32768)
        case "hello":
            fmt.Println(v6, v6)
        case "a":
            if true {
                switch "hello" {
                case "\n\t":
                    v2 = -386
                case "hello":
                    switch "a" {
                    case "":
                        var v17 int16 = v2
                        _ = v17
                    case "foo bar":
                        v18 := GenS2{F0: -9223372036854775808}
                        _ = v18
                    default:
                        v1 = append(v1, "\n\t")
                    }
                default:
                    var v19 int8 = 0
                    _ = v19
                }
            } else {
                var v20 int64 = 0
                v6 -= 939
                var v21 int64 = (v20 << ((1) & 63))
                _ = v20
                _ = v21
            }
        default:
            var v22 int16 = v2
            _ = v22
        }
    }
    v3["hello"] = "foo bar"
    fmt.Println(v2, v6, v4)
    if false {
        fmt.Println(v2, v2)
    }
    v2 = v2
    var v23 int16 = (v2 + 32767)
    var v24 uint = 65535
    var v25 string = "\n\t"
    var v26 uint16 = 32767
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
