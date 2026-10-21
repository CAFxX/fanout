package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint; F2 uint }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 uint8; F1 int; F2 bool }
type GenS5 struct { F0 bool; F1 int }
type GenS6 struct { F0 uint16; F1 int16 }
func genFn0(p0 int, p1 int32, p2 int) string {
    var v5 int = (p2 * 1)
    var v6 int16 = 32767
    var v7 uint = 255
    p2 -= (^(v5))
    if (18446744073709551615 <= 42) {
        var v8 uint16 = 453
        switch "\n\t" {
        case "foo bar":
            v8 = v8
        }
        if (2 >= 668) {
            v5 -= int(v8)
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := []int{(v5 & (p0 << (((v7 ^ (v7 / (((v7 + v7)) | 1)))) & 63))), (^(p2)), int(v7), (^(p2))}
                var v11 int16 = -234
                _ = fi9
                _ = v10
                _ = v11
            }
            if ((v8 << ((v7) & 63)) >= (v8 >> ((1) & 63))) {
                v12 := []int{p2}
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                v14 := GenS1{F0: (^(p1)), F1: (v7 * (v7 - v7)), F2: 329}
                v15 := GenS2{F0: v8, F1: p1}
                _ = fi13
                _ = v14
                _ = v15
            }
            v16 := GenS2{F0: (v8 * (^(v8))), F1: (p1 - p1)}
            _ = v16
        }
        _ = v8
    } else {
        v17 := []int{}
        for fi18 := 0; fi18 < 12; fi18++ {
            v6 -= 1
            if (true && true) {
                var v19 uint8 = 120
                _ = v19
            } else {
                switch p0 {
                case -1:
                    for fi20 := 0; fi20 < 19; fi20++ {
                        p0 += 2147483647
                        _ = fi20
                    }
                default:
                    var v21 uint32 = 2147483647
                    _ = v21
                }
                p0 = (p0 * p2)
            }
            v22 := GenS2{F0: 32767, F1: p1}
            _ = fi18
            _ = v22
        }
        _ = v17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (("a" + "foo bar") + ("foo bar" + "hello"))
}

func main() {
    var v1 bool = false
    v2 := []string{("foo bar" + "hello"), (("" + (("" + "foo bar") + "foo bar")) + "a"), ""}
    v3 := make(map[string]string)
    var v4 int32 = -670
    fmt.Println(v1, v1)
    v23 := GenS1{F0: v4, F1: 1, F2: 524}
    for fi24 := 0; fi24 < 13; fi24++ {
        for fi25 := 0; fi25 < 5; fi25++ {
            fmt.Println(fi25)
            _ = fi25
        }
        delete(v3, "hello")
        if (279 == 0) {
            switch fi24 {
            case 9223372036854775807:
                v26 := GenS1{F0: (v4 % (((v4 * v4)) | 1)), F1: 2, F2: 65535}
                _ = v26
            default:
                v1 = true
            }
            switch 734 {
            case 0:
                if true {
                    var v27 int16 = 32767
                    v4 -= (-(v4))
                    var v28 uint64 = 1
                    _ = v27
                    _ = v28
                } else {
                    _ = len(v2)
                }
            case 32767:
                switch "" {
                case "a":
                    var v29 bool = v1
                    _ = v29
                }
            }
            v30 := GenS3{F0: 2, F1: (v4 >> ((42) & 63))}
            _ = v30
        } else {
            if (42 >= 255) {
                if (42 >= 263) {
                    delete(v3, ("foo bar" + ("foo bar" + "a")))
                    _ = len(v3)
                } else {
                    v1 = v1
                }
                for fi31 := 0; fi31 < 1; fi31++ {
                    v32 := GenS2{F0: 255, F1: 0}
                    var v33 int8 = 9
                    _ = fi31
                    _ = v32
                    _ = v33
                }
            } else {
                fmt.Println(fi24, v4, v1)
            }
            v4 = (v4 % (((v4 - (v4 >> ((0) & 63)))) | 1))
            if v1 {
                v34 := v3["foo bar"]
                var v35 uint8 = 0
                if v1 {
                    _ = len(v3)
                    _ = len(v3)
                    if (v1 && v1) {
                        var v36 string = v34
                        _ = v36
                    }
                }
                _ = v34
                _ = v35
            } else {
                v37 := GenS4{F0: 0, F1: -9223372036854775808, F2: (!(!v1))}
                v38 := GenS2{F0: 32767, F1: 1}
                _ = v37
                _ = v38
            }
        }
        _ = fi24
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        v40 := GenS5{F0: false, F1: (fi39 - fi39)}
        fmt.Println(v4, fi39, v1)
        _ = fi39
        _ = v40
    }
    var v41 int32 = (v4 - v4)
    switch "" {
    case "\n\t":
        v42 := v3[(("\n\t" + ("foo bar" + "\n\t")) + (("hello" + "") + "\n\t"))]
        _ = v42
    default:
        v43 := GenS6{F0: 992, F1: -1}
        _ = v43
    }
    v1 = true
    fmt.Println(v1)
    v1 = v1
    var v44 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v41
    _ = v44
}
