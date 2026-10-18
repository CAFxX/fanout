package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int, p1 uint64, p2 string) uint32 {
    var v6 int16 = 723
    if true {
        var v7 int32 = -851
        switch p2 {
        case "\n\t":
            fmt.Println(p0)
        }
        var v8 uint16 = 1
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615, F2: -679}
    if true {
        var v10 int8 = 42
        v11 := []int{(p0 / ((p0) | 1))}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v9
    return 53
}

func main() {
    var v1 int32 = 550
    var v2 uint32 = 2147483648
    v3 := make(map[string]int)
    var v4 int64 = -1
    v5 := []uint32{v2, v2, v2, (v2 + 1)}
    switch ("\n\t" + (("a" + ("a" + "a")) + "a")) {
    case "a":
        v12 := GenS1{F0: 0, F1: 2, F2: v1}
        _ = v12
    case "":
        v3[""] = 971
    }
    if false {
        if (4294967295 > 42) {
            fmt.Println(v2, v4, v1)
            var v13 int16 = -32768
            _ = v13
        } else {
            for fi14 := 0; fi14 < 18; fi14++ {
                if false {
                    switch -9223372036854775808 {
                    case -288:
                        v1 -= v1
                    case -9223372036854775808:
                        _ = len(v5)
                    case 128:
                        for fi15 := 0; fi15 < 20; fi15++ {
                            v16 := GenS1{F0: -1, F1: 537, F2: (v1 | (v1 % (((v1 << ((2) & 63))) | 1)))}
                            _ = fi15
                            _ = v16
                        }
                    default:
                        var v17 float32 = 0.1
                        _ = v17
                    }
                }
                switch (((("a" + ("foo bar" + ("hello" + "foo bar"))) + ("hello" + ("foo bar" + "foo bar"))) + (("\n\t" + ("\n\t" + "foo bar")) + "a")) + "") {
                case "hello":
                    var v18 int32 = v1
                    _ = v18
                case "foo bar":
                    if len(v5) > 0 {
                        v5[0] = v2
                    }
                }
                _ = fi14
            }
            fmt.Println(v4, v1)
        }
    } else {
        if (("" + (("a" + (("a" + ("hello" + "\n\t")) + "foo bar")) + ("\n\t" + "\n\t"))) <= "a") {
            v19 := v3["\n\t"]
            v20 := GenS1{F0: (^(v4)), F1: 1, F2: v1}
            _ = v19
            _ = v20
        } else {
            v2 = v2
            switch 42 {
            case 2147483647:
                var v21 float64 = -2.25
                _ = v21
            case -1:
                v22 := GenS1{F0: (v4 / ((v4) | 1)), F1: 0, F2: 42}
                _ = v22
            default:
                var v23 float32 = (1.5 * float32(float64(-2.25)))
                _ = v23
            }
            var v24 uint32 = v2
            _ = v24
        }
        v25 := GenS1{F0: (v4 * v4), F1: 518, F2: int32(v2)}
        _ = v25
    }
    v2 -= 559
    if true {
        v2 = v2
        if true {
            switch -989 {
            case -9223372036854775808:
                var v26 uint32 = v2
                _ = v26
            }
            for fi27 := 0; fi27 < 8; fi27++ {
                fmt.Println(fi27)
                _ = fi27
            }
            _ = len(v5)
        }
    }
    if false {
        v28 := GenS2{F0: 0}
        var v29 string = ("hello" + ("a" + "foo bar"))
        _ = v28
        _ = v29
    } else {
        v1 -= (v1 & (v1 - (v1 + -1)))
    }
    var v30 uint64 = 2
    var v31 int = 9223372036854775807
    var v32 uint16 = 2
    v1 = (v1 >> ((433) & 63))
    var v33 uint = 89
    v31 = v31
    v34 := GenS3{F0: (v30 / (((v30 % ((v30) | 1))) | 1)), F1: 1}
    if false {
        if len(v5) > 0 {
            v5[0] = (v2 & genFn0(-1, v30, ""))
        }
        v35 := v3[("\n\t" + ("" + ""))]
        _ = v35
    }
    v5 = append(v5, genFn0((v31 % ((v31) | 1)), v30, ""))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
