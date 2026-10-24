package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 string, p1 int, p2 int8) int16 {
    v5 := GenS1{F0: -1, F1: 2}
    if (false || (true && false)) {
        p1 += p1
    } else {
        for fi6 := 0; fi6 < 6; fi6++ {
            var v7 uint16 = 32767
            _ = fi6
            _ = v7
        }
    }
    if false {
        p0 = ""
        switch (p1 - -346) {
        case 128:
            p0 = (((p0 + p0) + "foo bar") + p0)
        case 4294967295:
            for fi8 := 0; fi8 < 2; fi8++ {
                var v9 uint32 = 307
                _ = fi8
                _ = v9
            }
        case -527:
            var v10 string = p0
            _ = v10
        default:
            for fi11 := 0; fi11 < 13; fi11++ {
                p2 = -128
                v12 := []int{}
                if true {
                    var v13 string = "\n\t"
                    v14 := GenS1{F0: 0, F1: 650}
                    _ = v13
                    _ = v14
                } else {
                    var v15 int = (p1 & p1)
                    _ = v15
                }
                _ = fi11
                _ = v12
            }
        }
    }
    if true {
        switch -758 {
        case 128:
            for fi16 := 0; fi16 < 20; fi16++ {
                v17 := []int{fi16, p1, (fi16 >> ((4294967295) & 63)), p1}
                v18 := GenS1{F0: -1, F1: 65535}
                _ = fi16
                _ = v17
                _ = v18
            }
        case -441:
            p1 += 9223372036854775807
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 1000
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 0
    v3 := []string{}
    v4 := make(map[int]string)
    fmt.Println(v2, v1)
    v1 -= v1
    v2 = 42
    for fi19 := 0; fi19 < 14; fi19++ {
        if ((1.5 / (-2.25 * (float64(0.0) + 0.1))) > 1.5) {
            for fi20 := 0; fi20 < 14; fi20++ {
                v21 := GenS1{F0: 0, F1: 204}
                switch (fi20 / (((fi20 / (((fi19 - (fi19 - fi19))) | 1))) | 1)) {
                case 4294967295:
                    _ = len(v3)
                case -9223372036854775808:
                    v22 := GenS1{F0: -9223372036854775808, F1: 923}
                    _ = v22
                case -2147483648:
                    v4[fi20] = "\n\t"
                default:
                    var v23 uint32 = 58
                    _ = v23
                }
                _ = fi20
                _ = v21
            }
            var v24 float64 = 1.5
            if false {
                v2 = v2
            }
            _ = v24
        }
        switch 127 {
        case 1:
            delete(v4, fi19)
        case 32768:
            var v25 uint32 = 42
            _ = v25
        }
        v26 := GenS1{F0: -9223372036854775808, F1: 878}
        _ = fi19
        _ = v26
    }
    switch "\n\t" {
    case "hello":
        var v27 int8 = -58
        _ = v27
    }
    v1 += v1
    v2 = (v2 / (((^(v2))) | 1))
    v4[-841] = "\n\t"
    fmt.Println(v1, v2)
    if false {
        if (!false) {
            for fi28 := 0; fi28 < 19; fi28++ {
                fmt.Println(fi28, v2)
                _ = fi28
            }
            _ = len(v3)
            var v29 int = 9223372036854775807
            _ = v29
        } else {
            if len(v3) > 0 {
                v3[0] = "hello"
            }
            v1 = 2
        }
        _ = len(v4)
        var v30 int8 = 42
        _ = v30
    } else {
        var v31 int32 = -1
        v2 = v2
        _ = v31
    }
    var v32 bool = (false && false)
    v32 = true
    var v33 bool = (("a" + ("\n\t" + "foo bar")) >= "foo bar")
    _ = len(v4)
    if ((0 < 2147483648) || (!(("a" + ("foo bar" + ("" + "a"))) >= ("\n\t" + "\n\t")))) {
        if (v32 && (626 == 42)) {
            v1 = v1
        }
        fmt.Println(v33)
        v34 := GenS2{F0: 32767, F1: (^(v1))}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
