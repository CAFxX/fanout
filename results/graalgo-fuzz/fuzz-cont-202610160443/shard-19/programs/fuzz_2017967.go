package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint8; F1 uint64 }
type GenS5 struct { F0 int; F1 int64 }
func genFn0(p0 int32, p1 uint32) string {
    switch -881 {
    case 2147483647:
        if false {
            p1 += p1
        } else {
            var v6 string = ""
            _ = v6
        }
    case -689:
        if false {
            v7 := []int{128, 4294967295}
            if ((p1 >> ((0) & 63)) > p1) {
                if true {
                    var v8 float32 = ((0.0 * -2.25) - (((1e10 + -2.25) + -2.25) / ((float32(3.14159) / 0.1) + 3.14159)))
                    if len(v7) > 0 {
                        v7[0] = -331
                    }
                    switch ("" + ("hello" + "hello")) {
                    case "hello":
                        for fi9 := 0; fi9 < 17; fi9++ {
                            v10 := GenS1{F0: 2, F1: 0}
                            for fi11 := 0; fi11 < 15; fi11++ {
                                var v12 uint64 = 94
                                fmt.Println(v8, v12, p1)
                                _ = fi11
                                _ = v12
                            }
                            _ = fi9
                            _ = v10
                        }
                    }
                    _ = v8
                } else {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        p0 = p0
                        _ = fi13
                    }
                    p0 = p0
                }
                switch 42 {
                case -2147483648:
                    p1 = (p1 / ((p1) | 1))
                case 4294967295:
                    v14 := GenS1{F0: 65535, F1: -32768}
                    _ = v14
                }
                v15 := GenS2{F0: int64(p0), F1: 4294967295}
                _ = v15
            }
            _ = v7
        } else {
            if ("foo bar" == ("\n\t" + "a")) {
                switch "\n\t" {
                case "":
                    p1 = (p1 - (^(p1)))
                case "a":
                    var v16 uint = 0
                    _ = v16
                case "hello":
                    v17 := []int{9223372036854775807, 42}
                    _ = v17
                default:
                    for fi18 := 0; fi18 < 9; fi18++ {
                        var v19 string = "foo bar"
                        p0 -= p0
                        for fi20 := 0; fi20 < 1; fi20++ {
                            v21 := GenS1{F0: 32767, F1: 1}
                            _ = fi20
                            _ = v21
                        }
                        _ = fi18
                        _ = v19
                    }
                }
                fmt.Println(p1, p1)
            } else {
                var v22 int8 = 100
                for fi23 := 0; fi23 < 11; fi23++ {
                    if true {
                        var v24 string = (("" + ("foo bar" + ("hello" + ""))) + ("hello" + "a"))
                        p0 += 42
                        var v25 uint16 = 2
                        _ = v24
                        _ = v25
                    }
                    var v26 int8 = v22
                    var v27 int16 = -766
                    _ = fi23
                    _ = v26
                    _ = v27
                }
                _ = v22
            }
            v28 := []int{632, -1, 32767}
            if false {
                switch "foo bar" {
                case "":
                    var v29 int16 = 419
                    _ = v29
                case "\n\t":
                    var v30 int16 = -85
                    _ = v30
                }
                v31 := GenS3{F0: 195}
                _ = v31
            }
            _ = v28
        }
    case 128:
        var v32 int = 42
        _ = v32
    default:
        if true {
            var v33 float32 = 1.5
            v34 := GenS1{F0: 42, F1: 1}
            p1 = 2
            _ = v33
            _ = v34
        }
    }
    v35 := GenS4{F0: 142, F1: 42}
    _ = p0
    _ = p1
    _ = v35
    return "a"
}
func genFn1(p0 int, p1 int8, p2 int) uint64 {
    v36 := []int{(-(p0)), p2, (p0 / ((p2) | 1)), (p2 & -597)}
    p1 = p1
    for fi37 := 0; fi37 < 10; fi37++ {
        if len(v36) > 0 {
            v36[0] = p2
        }
        var v38 int64 = -808
        v39 := GenS5{F0: -9223372036854775808, F1: -1}
        _ = fi37
        _ = v38
        _ = v39
    }
    var v40 int = 32768
    var v41 float32 = 0.1
    _ = p0
    _ = p1
    _ = p2
    _ = v36
    _ = v40
    _ = v41
    return 1
}

func main() {
    var v1 string = "\n\t"
    var v2 uint = 18446744073709551615
    var v3 uint = 2
    var v4 int8 = int8(v2)
    var v5 uint16 = 32767
    for fi42 := 0; fi42 < 16; fi42++ {
        var v43 int16 = -624
        _ = fi42
        _ = v43
    }
    switch "" {
    case "foo bar":
        v2 += v3
    case "\n\t":
        var v44 uint8 = 0
        _ = v44
    default:
        v5 -= 32767
    }
    var v45 int32 = 1
    v1 = (v1 + (("foo bar" + (v1 + ("hello" + ""))) + (v1 + v1)))
    v46 := []int{0, -934, 32767, 32767}
    fmt.Println(v45)
    for fi47 := 0; fi47 < 16; fi47++ {
        v48 := GenS2{F0: -538, F1: v2}
        _ = fi47
        _ = v48
    }
    var v49 string = (v1 + v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v49
    _ = v5
}
