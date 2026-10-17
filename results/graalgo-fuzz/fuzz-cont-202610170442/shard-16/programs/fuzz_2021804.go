package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 bool; F2 int64 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 bool; F1 int64 }
type GenS5 struct { F0 uint; F1 int16; F2 int }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 bool; F1 uint32; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint16 }
func genFn0(p0 int16, p1 int8) uint16 {
    v6 := GenS1{F0: 47}
    p1 -= (^(p1))
    _ = p0
    _ = p1
    _ = v6
    return 0
}
func genFn1(p0 bool, p1 uint16, p2 string) int32 {
    v7 := GenS1{F0: 0}
    switch ((p2 + (("" + p2) + (((p2 + p2) + ("" + p2)) + p2))) + p2) {
    case "":
        for fi8 := 0; fi8 < 1; fi8++ {
            if ((false && ((fi8 ^ 9223372036854775807) < fi8)) || p0) {
                if true {
                    p2 = "\n\t"
                    var v9 int64 = 1
                    p0 = true
                    _ = v9
                } else {
                    var v10 float64 = -2.25
                    _ = v10
                }
                if ((-(fi8)) > -9223372036854775808) {
                    var v11 uint = 255
                    _ = v11
                }
                var v12 uint16 = genFn0(-717, 71)
                _ = v12
            }
            v13 := GenS2{F0: 42, F1: false, F2: 9223372036854775807}
            _ = fi8
            _ = v13
        }
    case "\n\t":
        switch 32768 {
        case 4294967295:
            v14 := []int{9223372036854775807, -907, -2147483648, 32768}
            _ = v14
        }
    case "hello":
        for fi15 := 0; fi15 < 20; fi15++ {
            fmt.Println(p0)
            v16 := GenS3{F0: -998, F1: 0}
            fmt.Println(p0)
            _ = fi15
            _ = v16
        }
    default:
        for fi17 := 0; fi17 < 5; fi17++ {
            var v18 float64 = 1.5
            if false {
                var v19 bool = p0
                for fi20 := 0; fi20 < 12; fi20++ {
                    v21 := []int{32768}
                    switch "" {
                    case "":
                        switch (fi20 + -566) {
                        case 9223372036854775807:
                            fmt.Println(v18)
                        case -2147483648:
                            var v22 int32 = -1
                            _ = v22
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "hello":
                            var v23 int8 = int8(fi20)
                            _ = v23
                        }
                    case "a":
                        var v24 int8 = 32
                        _ = v24
                    default:
                        var v25 int8 = int8(p1)
                        _ = v25
                    }
                    v26 := GenS4{F0: (fi17 >= fi17), F1: -336}
                    v27 := GenS4{F0: (135 > 127), F1: 42}
                    _ = fi20
                    _ = v21
                    _ = v26
                    _ = v27
                }
                _ = v19
            } else {
                var v28 uint8 = 105
                for fi29 := 0; fi29 < 3; fi29++ {
                    var v30 uint = 2
                    v28 = 255
                    var v31 int32 = 0
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                switch fi17 {
                case 9223372036854775807:
                    v32 := GenS1{F0: 60}
                    _ = v32
                default:
                    v33 := GenS1{F0: v28}
                    _ = v33
                }
                _ = v28
            }
            p2 = "foo bar"
            _ = fi17
            _ = v18
        }
    }
    switch "foo bar" {
    case "a":
        v34 := GenS5{F0: 255, F1: -32768, F2: -2147483648}
        _ = v34
    case "foo bar":
        fmt.Println(p1, p2)
    default:
        var v35 uint = 4294967295
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 42
}

func main() {
    var v1 bool = true
    var v2 uint = 255
    v3 := []int{127, 480, 544, 210}
    v4 := []int{128, 42, 128, -878}
    var v5 int64 = 9223372036854775807
    var v36 int32 = 157
    if ((2 >= 42) && v1) {
        var v37 uint16 = genFn0(-636, 42)
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = v37
    } else {
        v38 := GenS4{F0: v1, F1: v5}
        _ = v38
    }
    if ("foo bar" == "hello") {
        if true {
            var v39 uint32 = 65535
            switch -9223372036854775808 {
            case -198:
                v40 := GenS2{F0: 1, F1: (1 > 76), F2: (-(v5))}
                _ = v40
            case 32767:
                var v41 float32 = 0.1
                _ = v41
            }
            _ = v39
        } else {
            v2 = (v2 ^ 1)
            var v42 int = 390
            if true {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            _ = v42
        }
        if false {
            if v1 {
                v5 -= (v5 / ((v5) | 1))
            }
            for fi43 := 0; fi43 < 10; fi43++ {
                v5 = -128
                for fi44 := 0; fi44 < 12; fi44++ {
                    v3 = append(v3, (fi43 << ((v2) & 63)))
                    _ = fi44
                }
                for fi45 := 0; fi45 < 13; fi45++ {
                    v36 = v36
                    switch "\n\t" {
                    case "":
                        switch (fi43 - fi43) {
                        case 4294967295:
                            v46 := GenS3{F0: 1000, F1: 1}
                            _ = v46
                        }
                    case "hello":
                        var v47 int32 = -911
                        v1 = true
                        _ = v47
                    case "foo bar":
                        switch "hello" {
                        case "":
                            var v48 uint32 = 42
                            _ = v48
                        case "foo bar":
                            var v49 int64 = (v5 - v5)
                            _ = v49
                        case "a":
                            var v50 int32 = (v36 * (^(v36)))
                            _ = v50
                        }
                    default:
                        var v51 int8 = -124
                        _ = v51
                    }
                    v52 := GenS1{F0: 1}
                    _ = fi45
                    _ = v52
                }
                _ = fi43
            }
        } else {
            v36 -= -2147483648
            v53 := GenS4{F0: false, F1: (v5 & v5)}
            for fi54 := 0; fi54 < 1; fi54++ {
                fmt.Println(v5)
                if v1 {
                    v55 := GenS6{F0: 0}
                    v2 = 1
                    var v56 int = (^(fi54))
                    _ = v55
                    _ = v56
                } else {
                    v57 := GenS7{F0: v1, F1: 467, F2: v2}
                    if ((v5 / (((v5 % ((-304) | 1))) | 1)) > 1) {
                        fmt.Println(v5)
                    }
                    _ = v57
                }
                v5 = v5
                _ = fi54
            }
            _ = v53
        }
    }
    v5 -= (v5 & (^(v5)))
    var v58 int32 = (v36 >> ((v2) & 63))
    if v1 {
        v59 := GenS8{F0: v5}
        var v60 int64 = v5
        var v61 uint = (^(v2))
        _ = v59
        _ = v60
        _ = v61
    }
    v62 := GenS8{F0: v5}
    var v63 int64 = 42
    fmt.Println(v58)
    if (((v1 || (!v1)) && false) && ("" < "hello")) {
        v64 := GenS1{F0: 26}
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
        _ = v64
    }
    v1 = (60 >= 100)
    var v65 int8 = -32
    var v66 bool = (!v1)
    switch -144 {
    case -1:
        fmt.Println(v58, v1, v65)
    default:
        var v67 uint8 = 17
        _ = v67
    }
    v68 := GenS9{F0: 32767}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v63:", v63)
    fmt.Println("v65:", v65)
    fmt.Println("v66:", v66)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v58
    _ = v62
    _ = v63
    _ = v65
    _ = v66
    _ = v68
}
