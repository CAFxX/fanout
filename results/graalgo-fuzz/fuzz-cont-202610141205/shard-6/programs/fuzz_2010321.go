package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 uint; F1 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int) bool {
    for fi4 := 0; fi4 < 9; fi4++ {
        p0 += (fi4 ^ 9223372036854775807)
        var v5 int16 = 0
        switch ("a" + "a") {
        case "foo bar":
            p0 = 32767
        case "":
            var v6 int64 = -1
            _ = v6
        default:
            var v7 uint16 = 75
            _ = v7
        }
        _ = fi4
        _ = v5
    }
    switch "hello" {
    case "foo bar":
        if true {
            if (!(!true)) {
                switch p0 {
                case 0:
                    if false {
                        p0 += (p0 >> ((2) & 63))
                        var v8 uint16 = 255
                        v8 = 2
                        _ = v8
                    }
                case 32768:
                    if false {
                        for fi9 := 0; fi9 < 11; fi9++ {
                            var v10 string = ""
                            v10 = ("foo bar" + v10)
                            if (!(v10 <= v10)) {
                                for fi11 := 0; fi11 < 9; fi11++ {
                                    fmt.Println(fi9, p0)
                                    var v12 int64 = 1
                                    _ = fi11
                                    _ = v12
                                }
                                v13 := GenS1{F0: 128}
                                v10 = v10
                                _ = v13
                            } else {
                                var v14 int32 = 0
                                v14 = v14
                                _ = v14
                            }
                            _ = fi9
                            _ = v10
                        }
                        v15 := GenS1{F0: 171}
                        _ = v15
                    }
                default:
                    for fi16 := 0; fi16 < 5; fi16++ {
                        for fi17 := 0; fi17 < 3; fi17++ {
                            p0 -= -1
                            _ = fi17
                        }
                        _ = fi16
                    }
                }
            } else {
                if (!(p0 == p0)) {
                    p0 -= (p0 >> ((1) & 63))
                    var v18 int = -445
                    _ = v18
                } else {
                    switch ((("a" + "") + ("\n\t" + "\n\t")) + "foo bar") {
                    case "hello":
                        if (0 > 0) {
                            p0 = p0
                            switch "\n\t" {
                            case "hello":
                                var v19 int64 = 9223372036854775807
                                _ = v19
                            case "foo bar":
                                var v20 float32 = 1e10
                                _ = v20
                            case "\n\t":
                                v21 := GenS2{F0: 65535}
                                _ = v21
                            }
                        } else {
                            var v22 int64 = 306
                            var v23 int32 = 0
                            _ = v22
                            _ = v23
                        }
                    case "a":
                        var v24 uint = 345
                        _ = v24
                    case "\n\t":
                        var v25 uint16 = 110
                        _ = v25
                    }
                    v26 := []int{-200, (p0 >> ((2) & 63)), 127, p0}
                    _ = v26
                }
                for fi27 := 0; fi27 < 10; fi27++ {
                    p0 = p0
                    _ = fi27
                }
            }
            if true {
                p0 = p0
                v28 := GenS3{F0: -350, F1: 49}
                _ = v28
            }
        } else {
            if (1 == 1) {
                for fi29 := 0; fi29 < 18; fi29++ {
                    fmt.Println(fi29, p0)
                    var v30 uint8 = 171
                    _ = fi29
                    _ = v30
                }
                switch "a" {
                case "foo bar":
                    v31 := []int{p0, -537}
                    _ = v31
                case "a":
                    var v32 float32 = 1e10
                    _ = v32
                case "\n\t":
                    switch p0 {
                    case 481:
                        for fi33 := 0; fi33 < 10; fi33++ {
                            for fi34 := 0; fi34 < 12; fi34++ {
                                if (128 > 188) {
                                    fmt.Println(fi33, p0, fi34)
                                    p0 = -358
                                } else {
                                    var v35 uint32 = 965
                                    if false {
                                        v35 = v35
                                    } else {
                                        v36 := []int{(fi34 ^ 738)}
                                        if len(v36) > 0 {
                                            v36[0] = fi33
                                        }
                                        if (0.1 >= float32(0.0)) {
                                            if true {
                                                v35 -= 65535
                                            }
                                            if (-2.25 < -2.25) {
                                                v37 := GenS3{F0: -1, F1: 0}
                                                _ = v37
                                            }
                                            if false {
                                                v35 = 65535
                                                if len(v36) >= 2 {
                                                    v36 = v36[1:len(v36)]
                                                }
                                            } else {
                                                var v38 int16 = 32767
                                                var v39 int64 = -9223372036854775808
                                                v40 := GenS4{F0: 70, F1: v39}
                                                _ = v38
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        _ = v36
                                    }
                                    v35 -= 65535
                                    v41 := GenS1{F0: 89}
                                    _ = v35
                                    _ = v41
                                }
                                _ = fi34
                            }
                            if false {
                                var v42 bool = true
                                _ = v42
                            }
                            _ = fi33
                        }
                    case -9223372036854775808:
                        fmt.Println(p0)
                    default:
                        v43 := []int{p0, p0, -2147483648}
                        _ = v43
                    }
                }
                v44 := GenS1{F0: 127}
                _ = v44
            }
            for fi45 := 0; fi45 < 18; fi45++ {
                fmt.Println(p0, fi45)
                _ = fi45
            }
            fmt.Println(p0)
        }
    default:
        var v46 uint16 = 0
        _ = v46
    }
    _ = p0
    return (1 <= 42)
}
func genFn1(p0 uint64) uint64 {
    var v47 int16 = 32767
    for fi48 := 0; fi48 < 3; fi48++ {
        if true {
            p0 = p0
        }
        v49 := []int{fi48, fi48, -1, fi48}
        p0 = p0
        _ = fi48
        _ = v49
    }
    switch "\n\t" {
    case "\n\t":
        if true {
            switch 684 {
            case 1:
                if true {
                    for fi50 := 0; fi50 < 10; fi50++ {
                        v51 := GenS4{F0: 255, F1: -538}
                        var v52 int64 = 810
                        v53 := GenS1{F0: 181}
                        _ = fi50
                        _ = v51
                        _ = v52
                        _ = v53
                    }
                    fmt.Println(v47, v47)
                }
            }
        } else {
            fmt.Println(v47, v47)
            switch 32768 {
            case -2147483648:
                for fi54 := 0; fi54 < 1; fi54++ {
                    v47 += -32768
                    if false {
                        p0 += (p0 << ((557) & 63))
                    }
                    _ = fi54
                }
            case -9223372036854775808:
                v47 -= (v47 ^ v47)
            case 509:
                fmt.Println(p0)
            }
        }
    }
    _ = p0
    _ = v47
    return p0
}

func main() {
    v1 := []uint32{2147483647, 710}
    var v2 int8 = 42
    var v3 uint64 = 0
    if ((!(false || (-932 <= -2147483648))) || (true && true)) {
        fmt.Println(v2)
        var v55 uint16 = 42
        v1 = append(v1, 42)
        _ = v55
    }
    if (((("hello" + "") + (("hello" + "hello") + "\n\t")) + "foo bar") <= ("hello" + ("a" + ("a" + "\n\t")))) {
        switch -1 {
        case 824:
            fmt.Println(v3, v2)
        case -822:
            var v56 int16 = -759
            _ = v56
        }
        v57 := GenS5{F0: 42}
        switch "" {
        case "foo bar":
            v58 := GenS5{F0: -2147483648}
            _ = v58
        case "":
            for fi59 := 0; fi59 < 12; fi59++ {
                var v60 int8 = 100
                var v61 string = ("\n\t" + "a")
                _ = fi59
                _ = v60
                _ = v61
            }
        }
        _ = v57
    }
    v3 = v3
    for fi62 := 0; fi62 < 14; fi62++ {
        _ = len(v1)
        _ = fi62
    }
    var v63 uint32 = 65535
    var v64 uint16 = 645
    var v65 uint = 446
    v65 = v65
    var v66 string = "a"
    if len(v1) > 0 {
        v1[0] = 65535
    }
    for fi67 := 0; fi67 < 18; fi67++ {
        var v68 int = (fi67 - 776)
        v65 = (v65 ^ (v65 ^ (v65 & uint(v3))))
        _ = fi67
        _ = v68
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v63
    _ = v64
    _ = v65
    _ = v66
}
