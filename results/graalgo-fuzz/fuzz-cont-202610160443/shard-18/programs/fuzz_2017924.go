package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 uint32; F1 int; F2 uint64 }
type GenS4 struct { F0 bool; F1 bool; F2 bool }
func genFn0() uint32 {
    var v7 uint = 0
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{}
        v7 = 0
        _ = len(v9)
        _ = fi8
        _ = v9
    }
    v7 = v7
    v10 := GenS1{F0: 65535}
    if false {
        v11 := GenS2{F0: -903, F1: 0}
        v7 = v7
        var v12 int16 = 1000
        _ = v11
        _ = v12
    }
    _ = v10
    _ = v7
    return 625
}
func genFn1(p0 int64) int32 {
    switch ("a" + "foo bar") {
    case "\n\t":
        for fi13 := 0; fi13 < 10; fi13++ {
            p0 = p0
            v14 := []int{}
            switch "a" {
            case "a":
                var v15 uint64 = 0
                _ = v15
            case "foo bar":
                for fi16 := 0; fi16 < 6; fi16++ {
                    if (-1 > 1) {
                        if len(v14) > 0 {
                            v14[0] = fi13
                        }
                        fmt.Println(fi16)
                        v17 := GenS2{F0: -32768, F1: 0}
                        _ = len(v14)
                        var v18 uint32 = 0
                        if true {
                            _ = len(v14)
                        } else {
                            _ = len(v14)
                            if false {
                                v19 := GenS3{F0: genFn0(), F1: fi16, F2: 878}
                                _ = v19
                            } else {
                                var v20 int64 = p0
                                if true {
                                    var v21 uint64 = 985
                                    v22 := GenS4{F0: true, F1: (v18 > (v18 & genFn0())), F2: (!(32767 != (fi16 + (fi13 | (fi13 << ((uint(v21)) & 63))))))}
                                    if (false && (-1 >= 1)) {
                                        var v23 int16 = 1
                                        _ = v23
                                    } else {
                                        if false {
                                            var v24 int8 = 100
                                            switch ("" + "hello") {
                                            case "":
                                                var v25 float32 = 0.0
                                                _ = v25
                                            default:
                                                if len(v14) >= 2 {
                                                    v14 = v14[1:len(v14)]
                                                }
                                            }
                                            _ = v24
                                        } else {
                                            switch "hello" {
                                            case "a":
                                                v26 := GenS2{F0: 1, F1: -128}
                                                _ = v26
                                            case "\n\t":
                                                var v27 bool = false
                                                _ = v27
                                            default:
                                                if (((float32(-2.25) + 1.5) * (3.14159 / 1.0)) > 3.14159) {
                                                    v20 = int64(v21)
                                                    v21 += 61
                                                    p0 = 9223372036854775807
                                                }
                                            }
                                            var v28 int64 = 48
                                            v29 := GenS1{F0: 1}
                                            if ("foo bar" < "a") {
                                                if true {
                                                    switch "" {
                                                    case "foo bar":
                                                        fmt.Println(fi13, fi16, v20)
                                                    case "":
                                                        var v30 int16 = 1000
                                                        _ = v30
                                                    }
                                                    p0 += 9223372036854775807
                                                    var v31 float64 = float64(1.5)
                                                    _ = v31
                                                } else {
                                                    p0 = (p0 / ((-180) | 1))
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                    var v32 int16 = 968
                                                    fmt.Println(v28)
                                                    _ = v32
                                                }
                                                if false {
                                                    fmt.Println(p0)
                                                }
                                            }
                                            _ = v28
                                            _ = v29
                                        }
                                        if false {
                                            switch "a" {
                                            case "a":
                                                v33 := GenS3{F0: v18, F1: fi16, F2: v21}
                                                _ = v33
                                            }
                                            v14 = append(v14, int(v20))
                                        }
                                    }
                                    p0 += (p0 * (v20 * (p0 - p0)))
                                    _ = v21
                                    _ = v22
                                }
                                v20 += p0
                                _ = v20
                            }
                        }
                        _ = v17
                        _ = v18
                    } else {
                        fmt.Println(p0)
                        switch (fi16 | fi13) {
                        case 42:
                            var v34 int8 = -128
                            v35 := GenS2{F0: 1000, F1: v34}
                            v34 -= int8(p0)
                            _ = v34
                            _ = v35
                        case 10:
                            switch fi13 {
                            case -1:
                                v36 := GenS1{F0: 4294967295}
                                _ = v36
                            default:
                                var v37 bool = true
                                _ = v37
                            }
                        case 539:
                            p0 = 440
                        default:
                            fmt.Println(p0)
                        }
                        _ = len(v14)
                    }
                    _ = fi16
                }
            case "hello":
                for fi38 := 0; fi38 < 5; fi38++ {
                    if false {
                        var v39 int16 = 925
                        fmt.Println(fi13, fi13, fi38)
                        _ = v39
                    }
                    _ = fi38
                }
            default:
                fmt.Println(p0, fi13)
            }
            _ = fi13
            _ = v14
        }
    case "foo bar":
        v40 := []int{1}
        _ = v40
    }
    fmt.Println(p0)
    p0 += (p0 | (p0 % ((p0) | 1)))
    _ = p0
    return 0
}

func main() {
    v1 := []string{("" + ""), "hello", "foo bar"}
    v2 := make(map[int]uint32)
    v3 := []string{("a" + "foo bar"), "", "a"}
    v4 := []uint32{4294967295}
    v5 := make(map[uint32]string)
    v6 := make(map[uint32]string)
    var v41 string = ("foo bar" + "foo bar")
    var v42 float32 = -2.25
    v43 := v5[2147483647]
    v42 = v42
    if false {
        v41 = v41
    } else {
        v41 = ("foo bar" + v43)
    }
    v42 = (v42 / 1.5)
    v42 = v42
    v44 := GenS2{F0: 1, F1: 127}
    var v45 uint32 = 2147483647
    for fi46 := 0; fi46 < 6; fi46++ {
        switch fi46 {
        case 128:
            var v47 float64 = -2.25
            _ = v47
        case 862:
            switch 127 {
            case 32767:
                var v48 uint16 = 65535
                _ = v48
            }
        }
        var v49 float64 = (1.5 + 0.1)
        if true {
            v4 = append(v4, (^(v45)))
        } else {
            v49 -= (1e10 + (v49 - 1e10))
        }
        _ = fi46
        _ = v49
    }
    v42 = v42
    switch (v43 + "foo bar") {
    case "a":
        var v50 int = -1
        _ = v50
    case "\n\t":
        v43 = "foo bar"
    case "":
        var v51 string = ("foo bar" + ("foo bar" + "\n\t"))
        _ = v51
    }
    switch -931 {
    case -740:
        fmt.Println(v42, v43, v41)
    default:
        if (!(42 >= -395)) {
            var v52 int64 = 0
            _ = v52
        }
    }
    switch ("a" + "a") {
    case "\n\t":
        switch -588 {
        case 32768:
            v42 += 3.14159
        default:
            if true {
                v45 = v45
                var v53 float32 = v42
                _ = v53
            } else {
                var v54 int32 = 95
                var v55 int32 = genFn1(0)
                _ = v54
                _ = v55
            }
        }
    default:
        for fi56 := 0; fi56 < 20; fi56++ {
            for fi57 := 0; fi57 < 12; fi57++ {
                v43 = ("a" + (("a" + v43) + v41))
                v58 := GenS3{F0: v45, F1: (fi56 / ((202) | 1)), F2: 42}
                v43 = v41
                _ = fi57
                _ = v58
            }
            switch v41 {
            case "":
                v59 := v2[fi56]
                _ = v59
            case "a":
                var v60 int16 = -1
                _ = v60
            case "hello":
                v5[v45] = (v41 + v41)
            default:
                var v61 int64 = 9223372036854775807
                _ = v61
            }
            _ = fi56
        }
    }
    if false {
        v62 := GenS4{F0: (40 != 65535), F1: true, F2: (-589 <= 32767)}
        if (!(!(!(!(1 >= -1))))) {
            v63 := GenS3{F0: (v45 - (v45 >> ((65535) & 63))), F1: 32767, F2: 4294967295}
            _ = len(v5)
            _ = v63
        }
        _ = v62
    }
    v45 = v45
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
