package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
func genFn0(p0 uint16, p1 int64, p2 bool) int {
    fmt.Println(p1, p0, p2)
    switch 190 {
    case -105:
        if p2 {
            p0 = p0
            var v6 int64 = (p1 / ((p1) | 1))
            _ = v6
        } else {
            if ((42 >= 491) && p2) {
                var v7 int16 = 1
                v8 := GenS1{F0: -1, F1: (p0 >> ((4294967295) & 63))}
                p2 = true
                _ = v7
                _ = v8
            } else {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int8 = 1
                    p1 = p1
                    p0 -= p0
                    _ = fi9
                    _ = v10
                }
                var v11 float32 = 1.5
                _ = v11
            }
        }
    default:
        switch ("" + "a") {
        case "foo bar":
            if (!(p2 && (2 <= 255))) {
                switch -9223372036854775808 {
                case 41:
                    switch (("a" + "foo bar") + "foo bar") {
                    case "\n\t":
                        p2 = (p2 && false)
                    case "":
                        for fi12 := 0; fi12 < 10; fi12++ {
                            fmt.Println(p1)
                            _ = fi12
                        }
                    }
                case 9223372036854775807:
                    v13 := GenS1{F0: 0, F1: p0}
                    _ = v13
                }
            }
        case "":
            var v14 float32 = float32(1.5)
            _ = v14
        case "a":
            fmt.Println(p1, p0)
        default:
            if (0 >= -960) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    for fi16 := 0; fi16 < 4; fi16++ {
                        p1 += -9223372036854775808
                        var v17 int8 = 112
                        _ = fi16
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 12; fi18++ {
                        switch "\n\t" {
                        case "\n\t":
                            v19 := []int{fi15, 1}
                            _ = v19
                        case "foo bar":
                            var v20 string = ("" + (("" + "foo bar") + ("\n\t" + "hello")))
                            _ = v20
                        default:
                            var v21 bool = p2
                            _ = v21
                        }
                        switch fi18 {
                        case 1:
                            var v22 string = "hello"
                            _ = v22
                        case -568:
                            p2 = true
                        case 192:
                            var v23 bool = (p2 && (true || true))
                            _ = v23
                        }
                        if p2 {
                            var v24 int32 = -2147483648
                            _ = v24
                        }
                        _ = fi18
                    }
                    _ = fi15
                }
            }
        }
    }
    v25 := []int{32768}
    for fi26 := 0; fi26 < 19; fi26++ {
        if (733 <= 0) {
            switch fi26 {
            case 4294967295:
                fmt.Println(p1, p1, p1)
            }
        }
        _ = fi26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v25
    return 0
}
func genFn1() int16 {
    for fi27 := 0; fi27 < 9; fi27++ {
        var v28 uint = 65535
        v29 := []int{}
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := []int{0, 9223372036854775807, -905}
    _ = v30
    return 42
}

func main() {
    var v1 uint8 = 2
    var v2 uint8 = 83
    v3 := []uint32{}
    var v4 uint16 = 2
    v5 := []uint32{}
    v2 += v1
    v1 = (v1 % (((^(v2))) | 1))
    v4 = v4
    var v31 string = "foo bar"
    var v32 int16 = -488
    for fi33 := 0; fi33 < 9; fi33++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v34 float32 = 3.14159
        _ = fi33
        _ = v34
    }
    switch v31 {
    case "a":
        v2 = 255
    case "":
        fmt.Println(v31, v31, v1)
    }
    v31 = (v31 + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
}
