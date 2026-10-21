package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int64 }
type GenS2 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 uint32 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint }
func genFn0() int8 {
    var v7 int64 = -934
    var v8 int64 = v7
    for fi9 := 0; fi9 < 2; fi9++ {
        var v10 int8 = 100
        var v11 float32 = float32(-2.25)
        _ = fi9
        _ = v10
        _ = v11
    }
    if (135 <= 2) {
        for fi12 := 0; fi12 < 17; fi12++ {
            for fi13 := 0; fi13 < 19; fi13++ {
                fmt.Println(v8, v8)
                v14 := GenS1{F0: 0, F1: 1}
                _ = fi13
                _ = v14
            }
            v7 = (^(v7))
            _ = fi12
        }
        var v15 uint8 = 236
        v15 += (v15 | (^(v15)))
        _ = v15
    } else {
        var v16 int16 = -1
        v7 = (v8 / ((9223372036854775807) | 1))
        for fi17 := 0; fi17 < 19; fi17++ {
            if true {
                for fi18 := 0; fi18 < 11; fi18++ {
                    v16 = 0
                    v19 := []int{fi17}
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    _ = fi18
                    _ = v19
                }
            }
            _ = fi17
        }
        _ = v16
    }
    var v20 uint64 = 467
    _ = v20
    _ = v7
    _ = v8
    return -111
}
func genFn1(p0 uint16, p1 string) uint64 {
    p0 = (^(p0))
    if false {
        v21 := []int{32768, -629}
        _ = v21
    } else {
        p1 = ((p1 + (p1 + p1)) + (p1 + "hello"))
    }
    if false {
        var v22 bool = false
        p1 = ""
        v23 := []int{0, 2147483647, -9223372036854775808, -172}
        _ = v22
        _ = v23
    } else {
        switch p1 {
        case "hello":
            var v24 int16 = 1
            _ = v24
        }
        for fi25 := 0; fi25 < 7; fi25++ {
            for fi26 := 0; fi26 < 18; fi26++ {
                v27 := GenS1{F0: 216, F1: -29}
                v28 := []int{(fi26 | 327), (fi26 ^ fi26)}
                var v29 uint = 2
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
            }
            v30 := GenS2{F0: -128, F1: 103, F2: 963}
            _ = fi25
            _ = v30
        }
        switch p1 {
        case "a":
            p0 = (p0 | p0)
        }
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}

func main() {
    var v1 float32 = 3.14159
    var v2 uint64 = 969
    var v3 uint64 = (v2 - v2)
    v4 := make(map[int]uint32)
    var v5 float64 = (0.1 + 0.0)
    var v6 float64 = 1.5
    var v31 string = (("hello" + "\n\t") + "hello")
    switch v31 {
    case "\n\t":
        v32 := []int{42}
        _ = v32
    }
    if (true || ((0 != -524) || false)) {
        for fi33 := 0; fi33 < 6; fi33++ {
            v6 = v5
            _ = fi33
        }
        var v34 uint64 = (v2 % ((v3) | 1))
        _ = v34
    } else {
        v5 += (v5 / 1e10)
        var v35 string = ((v31 + ((v31 + (v31 + "a")) + (("hello" + v31) + "\n\t"))) + (("foo bar" + v31) + (("" + v31) + (v31 + "a"))))
        delete(v4, 209)
        _ = v35
    }
    if false {
        var v36 int16 = -1
        _ = v36
    }
    var v37 int16 = 32767
    switch "foo bar" {
    case "":
        v38 := GenS1{F0: 2, F1: -840}
        _ = v38
    case "a":
        for fi39 := 0; fi39 < 7; fi39++ {
            for fi40 := 0; fi40 < 16; fi40++ {
                switch (fi39 << ((2) & 63)) {
                case -147:
                    if (0 < v37) {
                        fmt.Println(v6)
                        fmt.Println(v2, v31, v37)
                        fmt.Println(v37, fi39)
                        switch "\n\t" {
                        case "a":
                            v31 = ((("" + (((v31 + ((("a" + "hello") + "a") + (("hello" + v31) + v31))) + ("\n\t" + ((v31 + v31) + ("hello" + v31)))) + v31)) + v31) + (v31 + "\n\t"))
                        case "hello":
                            var v41 int = 32768
                            _ = v41
                        case "":
                            delete(v4, fi39)
                        }
                        v42 := GenS2{F0: genFn0(), F1: 42, F2: genFn1(255, v31)}
                        _ = v42
                    } else {
                        switch (fi39 | fi40) {
                        case 1:
                            var v43 uint = 18446744073709551615
                            _ = v43
                        case -9223372036854775808:
                            var v44 uint32 = 65535
                            _ = v44
                        case -342:
                            v2 = genFn1(566, v31)
                        default:
                            v45 := GenS3{F0: 6, F1: true, F2: 877}
                            _ = v45
                        }
                        var v46 int = fi40
                        _ = v46
                    }
                    var v47 bool = (48 >= 111)
                    _ = v47
                case 4294967295:
                    v4[fi40] = 2147483647
                default:
                    fmt.Println(v3)
                }
                switch (v31 + "") {
                case "hello":
                    if true {
                        fmt.Println(fi39, v3)
                    }
                case "":
                    var v48 bool = (false && ((v5 * 1.5) > 1e10))
                    v3 = 18446744073709551615
                    _ = v48
                }
                fmt.Println(v31, v2, v3)
                _ = fi40
            }
            if true {
                v49 := GenS4{F0: 213}
                if (42 >= 377) {
                    if true {
                        fmt.Println(v5)
                        for fi50 := 0; fi50 < 15; fi50++ {
                            var v51 int64 = 0
                            _ = fi50
                            _ = v51
                        }
                    } else {
                        switch (((v31 + "foo bar") + "a") + v31) {
                        case "foo bar":
                            var v52 string = "a"
                            _ = v52
                        }
                        var v53 uint8 = 42
                        _ = v53
                    }
                    if true {
                        var v54 float64 = (v6 - (((float64(1e10) + v5) / 3.14159) * ((3.14159 * v6) / float64(v1))))
                        _ = v54
                    } else {
                        v55 := GenS5{F0: -11}
                        fmt.Println(v3)
                        fmt.Println(v3, v1, v31)
                        _ = v55
                    }
                    v56 := GenS2{F0: 1, F1: 1, F2: (v3 ^ genFn1(1, v31))}
                    _ = v56
                }
                _ = len(v4)
                _ = v49
            } else {
                switch (-(fi39)) {
                case 127:
                    v57 := GenS2{F0: genFn0(), F1: 0, F2: v3}
                    _ = v57
                case 1:
                    fmt.Println(v5, v5)
                }
            }
            _ = fi39
        }
    default:
        v1 = float32(float64(1e10))
    }
    v58 := []int{}
    var v59 uint64 = (v2 ^ 79)
    if false {
        fmt.Println(v1, v31)
        if (255 >= 0) {
            if len(v58) > 0 {
                v58[0] = 1
            }
            var v60 int32 = 0
            v4[68] = 70
            _ = v60
        }
    } else {
        switch "a" {
        case "foo bar":
            for fi61 := 0; fi61 < 16; fi61++ {
                var v62 int32 = 1
                _ = fi61
                _ = v62
            }
        }
        v63 := GenS5{F0: 2147483647}
        _ = v63
    }
    _ = len(v4)
    switch 9223372036854775807 {
    case 2147483647:
        v58 = append(v58, 9223372036854775807)
    case 1:
        var v64 bool = false
        _ = v64
    case 32767:
        _ = len(v58)
    }
    v65 := GenS2{F0: 0, F1: 42, F2: v3}
    if true {
        v66 := GenS6{F0: 801}
        v67 := GenS3{F0: 42, F1: true, F2: 1}
        _ = v66
        _ = v67
    } else {
        switch 32767 {
        case 4294967295:
            if (v37 == 1000) {
                var v68 int32 = 0
                v69 := v4[-904]
                _ = v68
                _ = v69
            }
        case -2147483648:
            v31 = ""
        case 32767:
            var v70 int32 = -2147483648
            _ = v70
        }
        _ = len(v58)
        v59 = (v2 % ((1) | 1))
    }
    v6 = (3.14159 * (float64(0.1) * (1e10 - v6)))
    var v71 int8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v37
    _ = v4
    _ = v5
    _ = v58
    _ = v59
    _ = v6
    _ = v65
    _ = v71
}
