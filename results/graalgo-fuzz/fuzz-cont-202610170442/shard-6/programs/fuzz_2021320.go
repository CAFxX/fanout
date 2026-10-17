package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 uint8; F1 int8 }
func genFn0(p0 int64, p1 int8) bool {
    var v7 uint16 = 65535
    var v8 uint32 = 880
    switch -4 {
    case 32768:
        if (float64((1.5 + 1e10)) == 0.0) {
            switch "foo bar" {
            case "\n\t":
                v9 := []int{}
                _ = v9
            }
            switch 9223372036854775807 {
            case 60:
                v10 := []int{9223372036854775807}
                _ = v10
            default:
                p1 = (p1 & 0)
            }
            p1 -= p1
        }
    case 445:
        p0 = (-(p0))
    case 128:
        v11 := []int{128, 42}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return (false || (!(42 == 65535)))
}

func main() {
    var v1 float64 = 0.0
    var v2 uint64 = 126
    var v3 int8 = -128
    var v4 int16 = 42
    var v5 uint8 = 127
    v6 := []string{("\n\t" + ""), ("hello" + (("\n\t" + "foo bar") + "a")), ((("a" + ("hello" + ("" + ("hello" + "a")))) + ("hello" + ("\n\t" + "a"))) + "foo bar"), "a"}
    var v12 bool = ((-2 < (v3 << ((1) & 63))) && true)
    v3 += v3
    fmt.Println(v4, v1)
    for fi13 := 0; fi13 < 14; fi13++ {
        v2 -= 2
        if v12 {
            fmt.Println(fi13, v3, v4)
            if true {
                v3 -= 0
                for fi14 := 0; fi14 < 13; fi14++ {
                    fmt.Println(v5)
                    var v15 uint16 = 255
                    _ = fi14
                    _ = v15
                }
            }
            v2 += 1
        } else {
            for fi16 := 0; fi16 < 4; fi16++ {
                switch "\n\t" {
                case "\n\t":
                    switch (fi16 - (fi16 << ((1) & 63))) {
                    case 127:
                        v2 = v2
                    case 9223372036854775807:
                        v17 := GenS1{F0: 4294967295, F1: v12}
                        _ = v17
                    case -1:
                        switch fi13 {
                        case -23:
                            if v12 {
                                var v18 uint64 = uint64(v5)
                                _ = v18
                            }
                        }
                    default:
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                }
                if len(v6) > 0 {
                    v6[0] = (("\n\t" + "") + "a")
                }
                v3 = v3
                _ = fi16
            }
        }
        fmt.Println(v12)
        _ = fi13
    }
    for fi19 := 0; fi19 < 17; fi19++ {
        v4 += -722
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := GenS1{F0: 255, F1: (-499 != 1)}
            _ = len(v6)
            v6 = append(v6, "\n\t")
            fmt.Println(v4, v2)
            v22 := GenS2{F0: 64, F1: 29}
            _ = fi20
            _ = v21
            _ = v22
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi19
    }
    var v23 int = 32767
    if len(v6) > 0 {
        v6[0] = ("foo bar" + (("a" + "\n\t") + "foo bar"))
    }
    var v24 uint8 = (v5 - v5)
    if v12 {
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v5, v1)
            fmt.Println(v23, v23)
            _ = fi25
        }
        v4 = (v4 ^ -32768)
        var v26 float64 = v1
        _ = v26
    }
    var v27 int8 = (v3 << ((952) & 63))
    var v28 int8 = (^(v3))
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    var v29 int64 = 552
    v3 = v28
    var v30 float64 = (v1 / v1)
    v12 = (-128 > v28)
    v5 += v5
    fmt.Println(v24)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
