package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint }
func genFn0(p0 string) string {
    for fi4 := 0; fi4 < 19; fi4++ {
        fmt.Println(fi4)
        for fi5 := 0; fi5 < 14; fi5++ {
            if ((!(4294967295 <= (fi4 / ((2147483647) | 1)))) && true) {
                v6 := []int{}
                p0 = "a"
                var v7 uint8 = 128
                v8 := GenS1{F0: (fi5 >> ((18446744073709551615) & 63))}
                _ = v6
                _ = v7
                _ = v8
            }
            if (false && false) {
                var v9 uint16 = 42
                _ = v9
            } else {
                p0 = ("a" + ((p0 + (p0 + ("" + "\n\t"))) + ((p0 + (p0 + (((p0 + "") + p0) + p0))) + "\n\t")))
                v10 := []int{9223372036854775807, 1}
                v11 := GenS1{F0: fi5}
                _ = v10
                _ = v11
            }
            switch "foo bar" {
            case "a":
                var v12 int32 = 42
                _ = v12
            case "hello":
                v13 := []int{(fi5 << ((65535) & 63))}
                _ = v13
            }
            v14 := []int{fi5, (fi5 * (fi4 + fi5)), fi5}
            _ = fi5
            _ = v14
        }
        _ = fi4
    }
    v15 := []int{2147483647, 9223372036854775807, 997}
    var v16 float64 = 0.1
    _ = p0
    _ = v15
    _ = v16
    return p0
}
func genFn1() uint32 {
    for fi17 := 0; fi17 < 14; fi17++ {
        switch (fi17 / (((fi17 >> ((346) & 63))) | 1)) {
        case 4294967295:
            switch "hello" {
            case "\n\t":
                switch (-(fi17)) {
                case 1:
                    var v18 int32 = 2147483647
                    _ = v18
                case -2147483648:
                    switch (fi17 / (((fi17 | fi17)) | 1)) {
                    case -9223372036854775808:
                        switch ("a" + ("" + ("\n\t" + ((("\n\t" + (("a" + "a") + "hello")) + ("a" + ("\n\t" + "\n\t"))) + ("hello" + ("a" + "a")))))) {
                        case "a":
                            for fi19 := 0; fi19 < 1; fi19++ {
                                var v20 float64 = (0.0 - 3.14159)
                                switch 32767 {
                                case 4294967295:
                                    var v21 float64 = ((v20 - v20) - (float64((-2.25 / -2.25)) / v20))
                                    _ = v21
                                case 32768:
                                    v22 := GenS2{F0: 1000, F1: 128}
                                    _ = v22
                                case -230:
                                    var v23 int32 = 42
                                    _ = v23
                                }
                                v20 += v20
                                _ = fi19
                                _ = v20
                            }
                        default:
                            var v24 bool = (!(42 != 127))
                            _ = v24
                        }
                    case 4294967295:
                        var v25 int32 = 1000
                        _ = v25
                    case 2147483647:
                        var v26 int8 = 100
                        _ = v26
                    default:
                        var v27 int32 = 544
                        _ = v27
                    }
                }
            default:
                var v28 uint = 4294967295
                _ = v28
            }
        }
        if true {
            var v29 bool = (false || false)
            _ = v29
        } else {
            if (!(2 != 2)) {
                var v30 int = fi17
                var v31 uint64 = 2
                for fi32 := 0; fi32 < 18; fi32++ {
                    v30 += fi17
                    v33 := []int{2147483647}
                    switch "foo bar" {
                    case "foo bar":
                        var v34 int16 = 32767
                        _ = v34
                    }
                    _ = fi32
                    _ = v33
                }
                _ = v30
                _ = v31
            }
        }
        fmt.Println(fi17)
        _ = fi17
    }
    for fi35 := 0; fi35 < 1; fi35++ {
        v36 := []int{fi35}
        switch "a" {
        case "a":
            var v37 int32 = -2147483648
            _ = v37
        }
        _ = fi35
        _ = v36
    }
    for fi38 := 0; fi38 < 19; fi38++ {
        var v39 int = -504
        _ = fi38
        _ = v39
    }
    v40 := []int{-233, -787, 42, 9223372036854775807}
    _ = v40
    return 42
}

func main() {
    v1 := make(map[int]string)
    var v2 int = -2147483648
    var v3 int = v2
    if true {
        for fi41 := 0; fi41 < 16; fi41++ {
            v2 = 280
            for fi42 := 0; fi42 < 20; fi42++ {
                switch "a" {
                case "foo bar":
                    switch "a" {
                    case "foo bar":
                        v2 = (v2 / ((v2) | 1))
                    }
                }
                _ = fi42
            }
            var v43 uint64 = 0
            _ = fi41
            _ = v43
        }
        fmt.Println(v3, v2)
    }
    if (1 == -50) {
        v44 := []int{v2}
        _ = v44
    }
    v1[-9223372036854775808] = ("a" + ("hello" + "a"))
    v45 := []int{v2, (v3 % (((v3 >> ((48) & 63))) | 1)), v2}
    v46 := GenS3{F0: -2147483648}
    for fi47 := 0; fi47 < 12; fi47++ {
        v3 = v2
        for fi48 := 0; fi48 < 1; fi48++ {
            v1[162] = (("hello" + "a") + ("" + "hello"))
            fmt.Println(v3, v2, v2)
            v49 := GenS3{F0: 786}
            _ = fi48
            _ = v49
        }
        var v50 int32 = 212
        _ = fi47
        _ = v50
    }
    for fi51 := 0; fi51 < 18; fi51++ {
        v52 := v1[(v2 >> ((150) & 63))]
        v53 := GenS3{F0: 679}
        _ = fi51
        _ = v52
        _ = v53
    }
    fmt.Println(v3, v3)
    var v54 uint16 = 255
    delete(v1, (v2 | v3))
    if true {
        v2 = 9223372036854775807
        v55 := GenS4{F0: 65535, F1: 626}
        for fi56 := 0; fi56 < 14; fi56++ {
            v57 := GenS2{F0: -519, F1: 174}
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                var v58 int64 = -126
                _ = v58
            default:
                if len(v45) >= 2 {
                    v45 = v45[1:len(v45)]
                }
            }
            _ = len(v45)
            _ = fi56
            _ = v57
        }
        _ = v55
    }
    switch (v2 - 59) {
    case 924:
        _ = len(v1)
    default:
        var v59 bool = (!((false && true) || (false && false)))
        _ = v59
    }
    v54 = v54
    fmt.Println(v54, v2)
    if true {
        if ((0 != 938) || (2147483647 >= 2)) {
            fmt.Println(v3, v2)
        } else {
            v60 := GenS4{F0: 65535, F1: 85}
            for fi61 := 0; fi61 < 11; fi61++ {
                if (42 != 913) {
                    fmt.Println(v54, v2)
                    var v62 uint32 = 0
                    var v63 int8 = -1
                    _ = v62
                    _ = v63
                } else {
                    if false {
                        var v64 int32 = -163
                        v3 = (fi61 * (fi61 / ((v3) | 1)))
                        _ = v64
                    }
                }
                for fi65 := 0; fi65 < 17; fi65++ {
                    var v66 int16 = -1
                    _ = fi65
                    _ = v66
                }
                fmt.Println(v3, v54)
                _ = fi61
            }
            v2 = v2
            _ = v60
        }
    } else {
        v54 = 32767
        var v67 uint32 = genFn1()
        fmt.Println(v67, v67, v67)
        _ = v67
    }
    var v68 uint64 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v45:", len(v45), v45)
    fmt.Println("v54:", v54)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v45
    _ = v46
    _ = v54
    _ = v68
}
