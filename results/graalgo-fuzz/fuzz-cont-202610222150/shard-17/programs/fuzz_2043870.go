package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := []string{("a" + "foo bar"), ("\n\t" + "hello"), "\n\t", ("hello" + ("foo bar" + ("hello" + "hello")))}
    v2 := make(map[uint32]bool)
    var v3 uint8 = 255
    var v4 int8 = 42
    var v5 uint8 = (v3 & 2)
    for fi6 := 0; fi6 < 7; fi6++ {
        if false {
            for fi7 := 0; fi7 < 9; fi7++ {
                if (!(0 > -174)) {
                    var v8 int32 = -509
                    _ = v8
                } else {
                    if ((v3 == v3) && false) {
                        if true {
                            switch (fi7 | (fi6 - (fi6 / (((-(fi6))) | 1)))) {
                            case 1:
                                v4 += (v4 * v4)
                            }
                        }
                    } else {
                        v9 := GenS1{F0: 4294967295, F1: true}
                        fmt.Println(fi6, v5, fi7)
                        v4 += (v4 ^ v4)
                        _ = v9
                    }
                }
                v10 := GenS1{F0: 4294967295, F1: (true || (65535 == 345))}
                v1 = append(v1, (("" + "foo bar") + ""))
                _ = fi7
                _ = v10
            }
        } else {
            switch "\n\t" {
            case "\n\t":
                fmt.Println(v4, fi6)
            default:
                if true {
                    v11 := GenS1{F0: 42, F1: true}
                    var v12 int = (fi6 - (fi6 & -1))
                    _ = v11
                    _ = v12
                }
            }
            v13 := GenS1{F0: 42, F1: true}
            _ = v13
        }
        _ = fi6
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v4 += v4
        var v15 int64 = 468
        for fi16 := 0; fi16 < 2; fi16++ {
            if false {
                var v17 string = ("foo bar" + ("" + ((((((("\n\t" + "") + (("" + ("foo bar" + "hello")) + (("hello" + "foo bar") + ("" + (("\n\t" + "\n\t") + ("a" + ("" + "hello"))))))) + ("foo bar" + ("foo bar" + (("a" + (((("" + "\n\t") + ("\n\t" + ("\n\t" + ("" + "")))) + ("\n\t" + ("foo bar" + "\n\t"))) + "hello")) + ((((("a" + ("" + ("a" + (("a" + "hello") + "hello")))) + "hello") + "\n\t") + (("hello" + "a") + "\n\t")) + ("hello" + (("hello" + "\n\t") + ("\n\t" + ("\n\t" + "\n\t"))))))))) + "\n\t") + "foo bar") + ("hello" + ("hello" + ""))) + "a")))
                _ = v17
            }
            v3 = 2
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    v5 -= (^(v3))
    v4 = (-(v4))
    v5 = 1
    var v18 int = 0
    v1 = append(v1, "\n\t")
    var v19 uint8 = v5
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
