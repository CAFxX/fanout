package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int; F1 int }
type GenS3 struct { F0 int64 }

func main() {
    v1 := []string{((("a" + "\n\t") + ((("\n\t" + ("a" + "\n\t")) + "foo bar") + "\n\t")) + "\n\t"), "hello", "a", (("a" + "\n\t") + "a")}
    var v2 int32 = 962
    v3 := []string{"foo bar", ((("\n\t" + ("" + "hello")) + "") + "\n\t"), "\n\t", ("\n\t" + ((("foo bar" + (("a" + "a") + "foo bar")) + "a") + ("" + "\n\t")))}
    v4 := []string{"foo bar", "foo bar", ("hello" + "\n\t")}
    var v5 float64 = ((0.1 - float64((1e10 * 1e10))) * (-2.25 / 1.5))
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 float64 = v5
        if true {
            var v8 int16 = 1
            v9 := GenS1{F0: -9223372036854775808, F1: (-(v2))}
            _ = v8
            _ = v9
        } else {
            var v10 string = "foo bar"
            _ = v10
        }
        _ = fi6
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v11 := GenS1{F0: -800, F1: (v2 * v2)}
    if false {
        if false {
            switch (("a" + ("a" + "a")) + (("a" + "hello") + "hello")) {
            case "foo bar":
                var v12 uint = 255
                _ = v12
            case "":
                if true {
                    if false {
                        if false {
                            if (true || ((false || false) && true)) {
                                var v13 string = ("" + "")
                                v14 := GenS2{F0: 1, F1: -662}
                                _ = v13
                                _ = v14
                            }
                            if (true && false) {
                                var v15 string = (("\n\t" + "") + "\n\t")
                                _ = v15
                            } else {
                                var v16 uint32 = 42
                                v5 = v5
                                v3 = append(v3, "a")
                                _ = v16
                            }
                            v17 := GenS1{F0: int64(v2), F1: v2}
                            _ = v17
                        } else {
                            var v18 float64 = v5
                            _ = v18
                        }
                        if true {
                            v5 -= 1.5
                        }
                    }
                }
            case "hello":
                if false {
                    if (!(-113 > 42)) {
                        for fi19 := 0; fi19 < 20; fi19++ {
                            if false {
                                switch ((("foo bar" + "foo bar") + ("a" + (("\n\t" + "") + ("\n\t" + "foo bar")))) + "\n\t") {
                                case "":
                                    v2 = 0
                                }
                                v5 = (v5 / 1e10)
                            }
                            for fi20 := 0; fi20 < 2; fi20++ {
                                v5 += v5
                                _ = fi20
                            }
                            _ = fi19
                        }
                    }
                    v4 = append(v4, "\n\t")
                } else {
                    v2 -= (v2 % ((v2) | 1))
                }
            default:
                v4 = append(v4, "foo bar")
            }
            v21 := GenS3{F0: 0}
            var v22 uint16 = 188
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v2)
            v2 += (v2 % (((^(v2))) | 1))
            _ = fi23
        }
    } else {
        v5 = 3.14159
    }
    v2 -= (v2 >> ((870) & 63))
    v24 := GenS3{F0: 394}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = 1
    for fi25 := 0; fi25 < 12; fi25++ {
        switch (fi25 + fi25) {
        case 42:
            fmt.Println(v2, v5, v5)
        default:
            v2 -= v2
        }
        _ = fi25
    }
    switch "foo bar" {
    case "hello":
        var v26 float64 = ((1e10 / (v5 * v5)) / v5)
        _ = v26
    default:
        fmt.Println(v2)
    }
    if len(v4) > 0 {
        v4[0] = "a"
    }
    fmt.Println(v2)
    v27 := GenS3{F0: 9223372036854775807}
    var v28 string = (("\n\t" + "") + "\n\t")
    v29 := GenS2{F0: 451, F1: 1}
    v30 := GenS3{F0: 29}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
