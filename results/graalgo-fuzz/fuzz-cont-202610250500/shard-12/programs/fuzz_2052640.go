package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 bool }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 int64; F1 uint; F2 uint64 }

func main() {
    v1 := []string{"hello"}
    var v2 int32 = -1
    var v3 int16 = 42
    var v4 uint16 = 468
    v5 := make(map[string]bool)
    v4 -= v4
    _ = len(v1)
    v2 = v2
    switch "\n\t" {
    case "\n\t":
        var v6 string = "a"
        _ = v6
    case "foo bar":
        if true {
            switch "foo bar" {
            case "":
                _ = len(v5)
            case "\n\t":
                for fi7 := 0; fi7 < 20; fi7++ {
                    if false {
                        if len(v1) > 0 {
                            v1[0] = ("\n\t" + "hello")
                        }
                        for fi8 := 0; fi8 < 3; fi8++ {
                            v9 := v5[("" + "\n\t")]
                            _ = fi8
                            _ = v9
                        }
                        _ = len(v1)
                    }
                    v3 += v3
                    _ = fi7
                }
            case "hello":
                if false {
                    for fi10 := 0; fi10 < 2; fi10++ {
                        if (false || (1 <= -684)) {
                            fmt.Println(v4, v2, v4)
                            var v11 uint16 = (v4 << ((159) & 63))
                            var v12 uint32 = 0
                            _ = v11
                            _ = v12
                        }
                        if false {
                            switch ("hello" + (("\n\t" + ("" + ("" + ("a" + "hello")))) + ("\n\t" + ("" + ((("\n\t" + "a") + (("foo bar" + "hello") + "\n\t")) + (("hello" + "foo bar") + "foo bar")))))) {
                            case "\n\t":
                                delete(v5, "")
                            default:
                                fmt.Println(v2)
                            }
                        }
                        _ = fi10
                    }
                    var v13 uint = 1
                    _ = v13
                }
            default:
                v2 -= 2147483647
            }
            for fi14 := 0; fi14 < 13; fi14++ {
                v1 = append(v1, "foo bar")
                _ = fi14
            }
            v4 += (v4 | (v4 + v4))
        } else {
            fmt.Println(v3, v2, v3)
        }
    case "":
        var v15 string = ""
        _ = v15
    default:
        if (!(42 >= -735)) {
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS1{F0: 4294967295, F1: -9223372036854775808, F2: (true || false)}
                var v18 string = "\n\t"
                _ = fi16
                _ = v17
                _ = v18
            }
            switch 4294967295 {
            case -9223372036854775808:
                v2 = 0
            case 9223372036854775807:
                _ = len(v5)
            case -1:
                v4 = v4
            }
        } else {
            fmt.Println(v2)
            v19 := GenS2{F0: 42, F1: (v2 / ((v2) | 1))}
            _ = v19
        }
    }
    v4 += (v4 << ((1) & 63))
    v2 -= v2
    _ = len(v5)
    if (!false) {
        var v20 uint = 255
        _ = v20
    } else {
        var v21 float64 = (1.5 + ((0.1 * 1.5) * 1.5))
        var v22 int16 = v3
        _ = v21
        _ = v22
    }
    var v23 int32 = v2
    v23 -= (^(v23))
    fmt.Println(v3, v2)
    v3 = (v3 & 32767)
    v24 := GenS3{F0: -9223372036854775808, F1: 1, F2: 318}
    var v25 bool = ((v4 | v4) < 1)
    if v25 {
        for fi26 := 0; fi26 < 17; fi26++ {
            delete(v5, ("" + "\n\t"))
            _ = fi26
        }
        v1 = append(v1, ("a" + ((("\n\t" + ("" + "hello")) + "hello") + ("foo bar" + ("foo bar" + "foo bar")))))
    }
    v23 = 344
    v27 := GenS2{F0: 2147483648, F1: v23}
    v2 = -583
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
