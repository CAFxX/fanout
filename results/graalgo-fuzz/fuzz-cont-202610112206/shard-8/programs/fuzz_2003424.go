package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := []string{"foo bar", "\n\t", ("hello" + ("\n\t" + "foo bar")), "a"}
    v2 := make(map[int]int)
    v3 := make(map[string]string)
    v4 := make(map[int]bool)
    var v5 bool = false
    v5 = (v5 || (!(v5 && (!(-32768 == 0)))))
    if true {
        for fi6 := 0; fi6 < 14; fi6++ {
            var v7 uint16 = 255
            var v8 int = fi6
            _ = fi6
            _ = v7
            _ = v8
        }
        v5 = v5
        v5 = v5
    }
    switch ("\n\t" + "") {
    case "hello":
        if v5 {
            switch ("hello" + "foo bar") {
            case "hello":
                var v9 int64 = -428
                _ = v9
            case "a":
                switch "hello" {
                case "foo bar":
                    v5 = true
                case "a":
                    if (2 != 782) {
                        v5 = (-1 <= -429)
                        if (((-2.25 - (1.5 * (-2.25 * -2.25))) / (3.14159 - 0.0)) > ((0.1 * 3.14159) + (0.0 + 1e10))) {
                            var v10 bool = (!true)
                            for fi11 := 0; fi11 < 4; fi11++ {
                                switch ("hello" + ("\n\t" + ("" + (("a" + "\n\t") + "a")))) {
                                case "hello":
                                    var v12 float32 = (float32(float64(float32((1e10 / (3.14159 + (3.14159 / 3.14159)))))) / 1e10)
                                    _ = v12
                                default:
                                    v4[(fi11 % (((^(fi11))) | 1))] = (!(18446744073709551615 != 2))
                                }
                                _ = fi11
                            }
                            _ = v10
                        }
                        switch 48 {
                        case -1:
                            v5 = false
                        default:
                            v1 = append(v1, "")
                        }
                    } else {
                        switch 32767 {
                        case 42:
                            var v13 int8 = 100
                            _ = v13
                        case -774:
                            if v5 {
                                switch ("hello" + "a") {
                                case "\n\t":
                                    fmt.Println(v5)
                                default:
                                    v14 := GenS1{F0: 42, F1: 32767, F2: 0}
                                    _ = v14
                                }
                                v15 := GenS2{F0: 1, F1: 2}
                                _ = v15
                            } else {
                                v5 = (!(1 < 255))
                            }
                        case 32768:
                            v5 = true
                        default:
                            var v16 int8 = 18
                            _ = v16
                        }
                    }
                case "\n\t":
                    fmt.Println(v5)
                }
            case "foo bar":
                v17 := GenS2{F0: 0, F1: 2}
                _ = v17
            default:
                var v18 bool = v5
                _ = v18
            }
            var v19 uint16 = 437
            v5 = v5
            _ = v19
        }
    case "foo bar":
        _ = len(v2)
    case "":
        v5 = true
    }
    _ = len(v1)
    var v20 int64 = 808
    var v21 uint = 1
    delete(v2, 32768)
    if (581 < 2) {
        var v22 int16 = 1000
        _ = v22
    } else {
        switch 127 {
        case 42:
            _ = len(v1)
        case 413:
            switch ("hello" + "") {
            case "foo bar":
                v23 := GenS2{F0: 42, F1: 932}
                _ = v23
            case "\n\t":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
        var v24 uint64 = 2
        _ = v24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
