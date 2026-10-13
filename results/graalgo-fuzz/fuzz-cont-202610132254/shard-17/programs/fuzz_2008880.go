package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int64; F1 int32 }
func genFn0(p0 uint) int32 {
    var v6 int16 = 1
    v7 := []int{9223372036854775807, -2147483648, -2147483648, 22}
    if true {
        v8 := GenS1{F0: 0, F1: 2, F2: 42}
        fmt.Println(v6, v6)
        for fi9 := 0; fi9 < 17; fi9++ {
            p0 -= (p0 * p0)
            v10 := GenS2{F0: 1}
            for fi11 := 0; fi11 < 12; fi11++ {
                p0 -= (p0 % (((p0 / ((2) | 1))) | 1))
                _ = fi11
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := GenS1{F0: 495, F1: 4294967295, F2: -2147483648}
        p0 = (p0 * (p0 >> (((^(p0))) & 63)))
        _ = fi12
        _ = v13
    }
    _ = p0
    _ = v6
    _ = v7
    return -272
}
func genFn1(p0 int8, p1 int8) int8 {
    var v14 string = ("hello" + "hello")
    switch "foo bar" {
    case "hello":
        v15 := []int{-664, -1, 32768, 42}
        _ = v15
    case "foo bar":
        v16 := []int{42}
        _ = v16
    }
    v17 := []int{42, -987}
    for fi18 := 0; fi18 < 3; fi18++ {
        if (!((127 == 201) && true)) {
            fmt.Println(fi18, v14, p1)
        }
        _ = fi18
    }
    _ = p0
    _ = p1
    _ = v14
    _ = v17
    return p0
}

func main() {
    var v1 uint = 422
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]int)
    var v4 uint32 = 192
    var v5 uint16 = 195
    v3["foo bar"] = 0
    v19 := GenS2{F0: 0}
    if false {
        for fi20 := 0; fi20 < 15; fi20++ {
            for fi21 := 0; fi21 < 16; fi21++ {
                delete(v2, 0)
                _ = fi21
            }
            var v22 uint32 = (v4 + v4)
            _ = fi20
            _ = v22
        }
        var v23 int8 = 83
        _ = v23
    }
    if true {
        for fi24 := 0; fi24 < 3; fi24++ {
            if true {
                v5 += v5
                var v25 int32 = 800
                v26 := GenS2{F0: 127}
                _ = v25
                _ = v26
            } else {
                v4 = v4
                v27 := v2[v4]
                var v28 uint16 = v5
                _ = v27
                _ = v28
            }
            switch ("hello" + ("foo bar" + "hello")) {
            case "\n\t":
                for fi29 := 0; fi29 < 3; fi29++ {
                    if true {
                        var v30 int64 = 1
                        fmt.Println(v30, fi29)
                        v31 := []int{fi29, 937, fi24}
                        _ = v30
                        _ = v31
                    }
                    var v32 uint32 = (v4 >> ((v1) & 63))
                    var v33 int = (fi29 % ((32768) | 1))
                    fmt.Println(v5, fi24)
                    _ = fi29
                    _ = v32
                    _ = v33
                }
            case "":
                var v34 int = (fi24 - (fi24 >> ((v1) & 63)))
                _ = v34
            case "a":
                v35 := []int{fi24}
                _ = v35
            default:
                for fi36 := 0; fi36 < 20; fi36++ {
                    if (1.5 == 1.5) {
                        v37 := GenS3{F0: 1}
                        _ = v37
                    } else {
                        if (!false) {
                            fmt.Println(fi24, v1)
                            if false {
                                v1 = (v1 % ((v1) | 1))
                            }
                            if false {
                                v4 = (v4 << (((v1 - uint(v4))) & 63))
                                var v38 float32 = (((float32((float64(0.0) + float64(1.5))) + (0.1 - 1e10)) + 3.14159) + (0.0 * -2.25))
                                if true {
                                    fmt.Println(v4, v4)
                                    v5 -= (v5 | 0)
                                }
                                _ = v38
                            } else {
                                if true {
                                    var v39 uint32 = 2
                                    v40 := []int{}
                                    fmt.Println(v4)
                                    _ = v39
                                    _ = v40
                                } else {
                                    if (42 >= 1) {
                                        var v41 int = (fi24 >> ((v1) & 63))
                                        _ = v41
                                    }
                                }
                                v42 := []int{(fi24 + 127), 296, fi24}
                                if (false || (((0.1 + (0.0 * 3.14159)) + 0.1) != 1e10)) {
                                    _ = len(v3)
                                    var v43 uint32 = 2147483647
                                    _ = v43
                                }
                                _ = v42
                            }
                        }
                    }
                    var v44 bool = true
                    _ = fi36
                    _ = v44
                }
            }
            v1 -= v1
            _ = fi24
        }
    } else {
        for fi45 := 0; fi45 < 20; fi45++ {
            switch ((("\n\t" + (((("foo bar" + "foo bar") + "\n\t") + "foo bar") + (("a" + ("foo bar" + "")) + "foo bar"))) + ("a" + "\n\t")) + ("" + "\n\t")) {
            case "":
                _ = len(v2)
            case "foo bar":
                delete(v3, "\n\t")
            case "hello":
                v46 := GenS2{F0: 100}
                _ = v46
            default:
                var v47 int16 = 1000
                _ = v47
            }
            _ = fi45
        }
    }
    v48 := GenS4{F0: 0, F1: -706}
    var v49 float32 = ((0.0 / (0.0 + float32(0.1))) * 3.14159)
    if false {
        switch ("" + "") {
        case "foo bar":
            for fi50 := 0; fi50 < 8; fi50++ {
                v5 = v5
                fmt.Println(fi50)
                _ = fi50
            }
        case "\n\t":
            var v51 string = "hello"
            _ = v51
        case "hello":
            var v52 uint8 = uint8(v1)
            _ = v52
        }
    } else {
        for fi53 := 0; fi53 < 10; fi53++ {
            for fi54 := 0; fi54 < 17; fi54++ {
                v55 := v2[v4]
                _ = fi54
                _ = v55
            }
            if true {
                switch "foo bar" {
                case "a":
                    if ((!(("a" + "hello") > ("" + "a"))) || (false && false)) {
                        v49 += float32(0.1)
                        switch (fi53 - (^(fi53))) {
                        case 128:
                            for fi56 := 0; fi56 < 5; fi56++ {
                                var v57 float64 = 0.0
                                v5 = (v5 << ((1) & 63))
                                _ = fi56
                                _ = v57
                            }
                        }
                    }
                case "\n\t":
                    v58 := []int{fi53}
                    _ = v58
                case "":
                    v59 := []int{(-(fi53)), (fi53 * (-(fi53))), 128}
                    _ = v59
                default:
                    v5 = v5
                }
                for fi60 := 0; fi60 < 15; fi60++ {
                    switch (fi53 - (fi53 ^ -957)) {
                    case -2147483648:
                        v5 += v5
                    case 464:
                        v61 := GenS1{F0: (v4 % ((2147483648) | 1)), F1: 4294967295, F2: 2147483647}
                        _ = v61
                    case 128:
                        v62 := GenS1{F0: (v4 / (((^(v4))) | 1)), F1: 824, F2: -2147483648}
                        v3["foo bar"] = (fi60 << (((v1 << (((v1 / ((v1) | 1))) & 63))) & 63))
                        _ = v62
                    }
                    v4 -= 1
                    _ = fi60
                }
            }
            _ = fi53
        }
    }
    var v63 uint = v1
    var v64 uint64 = 1
    var v65 uint = v1
    v66 := []int{}
    for fi67 := 0; fi67 < 18; fi67++ {
        v68 := GenS3{F0: uint8(v5)}
        fmt.Println(v5)
        _ = fi67
        _ = v68
    }
    v1 += v65
    var v69 bool = false
    var v70 uint16 = v5
    var v71 uint64 = v64
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("slicelen v66:", len(v66), v66)
    fmt.Println("v69:", v69)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v63
    _ = v64
    _ = v65
    _ = v66
    _ = v69
    _ = v70
    _ = v71
}
