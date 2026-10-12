package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS4 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int8, p1 string, p2 int32) uint64 {
    switch -440 {
    case 32768:
        var v6 int8 = (^(p0))
        _ = v6
    case -975:
        var v7 int8 = p0
        _ = v7
    case 330:
        p2 = p2
    }
    v8 := GenS1{F0: 844}
    p0 -= (p0 ^ p0)
    for fi9 := 0; fi9 < 9; fi9++ {
        p1 = p1
        fmt.Println(p2, p0, fi9)
        for fi10 := 0; fi10 < 7; fi10++ {
            if false {
                v11 := []int{}
                switch int(p2) {
                case 54:
                    v11 = append(v11, fi10)
                default:
                    p2 = p2
                }
                _ = v11
            }
            if (!(fi10 > -1)) {
                v12 := GenS2{F0: 32767, F1: (p1 + ((p1 + p1) + p1))}
                _ = v12
            } else {
                if (true || false) {
                    v13 := GenS2{F0: 42, F1: (("foo bar" + ("hello" + "a")) + p1)}
                    p1 = ("\n\t" + "foo bar")
                    var v14 int8 = -128
                    _ = v13
                    _ = v14
                } else {
                    switch ("a" + "a") {
                    case "":
                        v15 := []int{897, (fi9 % ((fi10) | 1)), (fi9 ^ (fi9 / (((fi9 >> ((4294967295) & 63))) | 1)))}
                        _ = v15
                    case "foo bar":
                        v16 := GenS3{F0: 18446744073709551615, F1: 25, F2: 4294967295}
                        _ = v16
                    case "a":
                        var v17 int16 = 332
                        _ = v17
                    }
                    v18 := GenS1{F0: p2}
                    p2 = (p2 | -1)
                    _ = v18
                }
                fmt.Println(fi9, p0)
            }
            v19 := []int{(fi9 & (^(fi9))), fi9, fi10, fi10}
            _ = fi10
            _ = v19
        }
        _ = fi9
    }
    fmt.Println(p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 432
}

func main() {
    v1 := []string{"hello", (("foo bar" + "\n\t") + ""), "a", ("" + "")}
    v2 := []string{"\n\t"}
    v3 := []int{-590}
    var v4 int32 = 437
    v5 := make(map[uint32]int)
    for fi20 := 0; fi20 < 8; fi20++ {
        v21 := GenS4{F0: 0, F1: 4294967295}
        v22 := GenS1{F0: -1}
        _ = fi20
        _ = v21
        _ = v22
    }
    if (!(78 <= 0)) {
        v23 := GenS4{F0: 2147483648, F1: 963}
        delete(v5, 655)
        _ = v23
    }
    v4 -= 1
    v4 -= v4
    switch 127 {
    case 4294967295:
        switch "\n\t" {
        case "foo bar":
            v4 = (v4 | v4)
        case "hello":
            v4 += (v4 | (v4 >> ((65535) & 63)))
        default:
            if false {
                var v24 int = -865
                _ = v24
            }
        }
    case 9223372036854775807:
        v4 += v4
    case 42:
        for fi25 := 0; fi25 < 19; fi25++ {
            if (4294967295 == 131) {
                for fi26 := 0; fi26 < 16; fi26++ {
                    var v27 uint8 = 196
                    switch int(v27) {
                    case 4294967295:
                        fmt.Println(v27, v27, fi26)
                    case 0:
                        if ((fi25 >= fi26) && true) {
                            var v28 float64 = 0.0
                            var v29 string = ("a" + "\n\t")
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = fi26
                    _ = v27
                }
                var v30 uint = 65535
                _ = v30
            }
            v4 = (v4 ^ v4)
            _ = fi25
        }
    }
    if false {
        for fi31 := 0; fi31 < 9; fi31++ {
            delete(v5, 137)
            for fi32 := 0; fi32 < 15; fi32++ {
                var v33 int32 = (v4 | (v4 ^ (v4 / (((v4 << ((0) & 63))) | 1))))
                fmt.Println(fi32)
                v33 += v33
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
        var v34 uint32 = 4294967295
        switch -1 {
        case -2147483648:
            fmt.Println(v4, v4)
        }
        _ = v34
    }
    for fi35 := 0; fi35 < 7; fi35++ {
        var v36 int64 = 42
        _ = fi35
        _ = v36
    }
    v4 = v4
    switch 738 {
    case 454:
        var v37 int32 = (-(v4))
        _ = v37
    case 42:
        v4 = (v4 - v4)
    default:
        if (!(!true)) {
            v4 += (v4 * (v4 & (v4 % ((-1) | 1))))
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
