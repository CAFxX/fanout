package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 uint8 }
func genFn0(p0 uint16, p1 int64, p2 uint16) string {
    p1 += 1
    v4 := GenS1{F0: p0, F1: 648}
    if true {
        switch -2147483648 {
        case 1:
            v5 := []int{-1, 32767, 42, -1}
            _ = v5
        }
    } else {
        v6 := GenS1{F0: (^(p0)), F1: -438}
        if false {
            v7 := []int{-60}
            for fi8 := 0; fi8 < 19; fi8++ {
                v9 := GenS1{F0: (p0 / ((0) | 1)), F1: fi8}
                if (!false) {
                    var v10 uint64 = 993
                    var v11 int = -102
                    v12 := GenS1{F0: (p2 / ((p0) | 1)), F1: (fi8 + -1)}
                    _ = v10
                    _ = v11
                    _ = v12
                } else {
                    if len(v7) > 0 {
                        v7[0] = fi8
                    }
                    fmt.Println(p0, fi8)
                }
                _ = fi8
                _ = v9
            }
            var v13 uint16 = (p2 / (((^(p2))) | 1))
            _ = v13
            _ = v7
        } else {
            p1 = (p1 & 9223372036854775807)
        }
        _ = v6
    }
    if false {
        var v14 uint64 = 850
        _ = v14
    } else {
        if (!false) {
            var v15 bool = (!false)
            var v16 bool = (-480 != 4294967295)
            _ = v15
            _ = v16
        } else {
            p1 = p1
        }
        switch ((("foo bar" + "foo bar") + "a") + "\n\t") {
        case "foo bar":
            if ((1 != 2147483647) || (false || false)) {
                var v17 uint16 = (p0 << ((570) & 63))
                var v18 int8 = 0
                _ = v17
                _ = v18
            } else {
                switch -944 {
                case 253:
                    var v19 int16 = -609
                    _ = v19
                }
            }
        case "hello":
            var v20 int8 = 78
            _ = v20
        case "a":
            if true {
                var v21 int16 = 983
                fmt.Println(p0, p1, p1)
                var v22 int8 = -96
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 5; fi23++ {
                    if true {
                        var v24 int8 = 127
                        var v25 string = (("hello" + "foo bar") + ("foo bar" + ""))
                        switch ("" + v25) {
                        case "foo bar":
                            var v26 bool = false
                            _ = v26
                        case "":
                            p0 = p0
                        }
                        _ = v24
                        _ = v25
                    }
                    var v27 float64 = 0.0
                    _ = fi23
                    _ = v27
                }
            }
        default:
            var v28 uint = 65535
            _ = v28
        }
        fmt.Println(p0, p0)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return ("hello" + ((((((("hello" + "foo bar") + "") + "hello") + ("hello" + "")) + "hello") + (("" + "a") + (("hello" + "foo bar") + "hello"))) + "\n\t"))
}

func main() {
    var v1 float64 = -2.25
    var v2 float32 = -2.25
    var v3 uint32 = 2147483648
    var v29 int64 = 9223372036854775807
    if false {
        if (0 != 18446744073709551615) {
            v30 := GenS2{F0: 0}
            v31 := GenS2{F0: 116}
            _ = v30
            _ = v31
        }
        var v32 int8 = 100
        _ = v32
    } else {
        v3 -= (v3 ^ (^(v3)))
        if (!((0.1 - ((v1 * (0.0 / v1)) * (0.0 - -2.25))) == v1)) {
            for fi33 := 0; fi33 < 18; fi33++ {
                switch fi33 {
                case 9223372036854775807:
                    var v34 int = (fi33 << ((0) & 63))
                    _ = v34
                case 27:
                    var v35 int = 267
                    _ = v35
                case 128:
                    fmt.Println(v1)
                default:
                    for fi36 := 0; fi36 < 9; fi36++ {
                        v1 += v1
                        var v37 int8 = 42
                        _ = fi36
                        _ = v37
                    }
                }
                v38 := []int{(^(fi33)), fi33, (^(fi33)), 961}
                v39 := GenS1{F0: 780, F1: fi33}
                _ = fi33
                _ = v38
                _ = v39
            }
        } else {
            fmt.Println(v3)
            v40 := []int{-359}
            _ = v40
        }
    }
    fmt.Println(v2)
    v2 = (v2 / v2)
    var v41 bool = (((false || false) && (((v3 & 1) < v3) && true)) && false)
    if true {
        v41 = false
    }
    v42 := []int{4294967295}
    v1 = v1
    var v43 int = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v42:", len(v42), v42)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
