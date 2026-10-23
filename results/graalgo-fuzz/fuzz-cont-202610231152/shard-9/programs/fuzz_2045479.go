package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 int8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32; F1 uint16 }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint; F1 int32 }
type GenS9 struct { F0 int8 }
func genFn0(p0 string, p1 uint16) string {
    var v6 int64 = 20
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(p1)
        var v8 float64 = 0.1
        var v9 int64 = (-(v6))
        _ = fi7
        _ = v8
        _ = v9
    }
    for fi10 := 0; fi10 < 14; fi10++ {
        v11 := GenS1{F0: 32767}
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        if true {
            for fi13 := 0; fi13 < 14; fi13++ {
                var v14 int8 = -128
                _ = fi13
                _ = v14
            }
            var v15 uint = 1
            var v16 string = ""
            _ = v15
            _ = v16
        } else {
            switch "a" {
            case "foo bar":
                var v17 int = fi12
                _ = v17
            case "a":
                var v18 int = fi12
                _ = v18
            default:
                p0 = "\n\t"
            }
            var v19 uint8 = 127
            for fi20 := 0; fi20 < 3; fi20++ {
                v6 = -559
                p1 = (p1 / (((p1 ^ (^(p1)))) | 1))
                if (false || (true || (2147483647 == 975))) {
                    v21 := GenS1{F0: 1}
                    p1 -= p1
                    _ = v21
                } else {
                    switch "\n\t" {
                    case "a":
                        v22 := GenS1{F0: -870}
                        _ = v22
                    case "hello":
                        v23 := GenS2{F0: -285, F1: 127}
                        _ = v23
                    case "foo bar":
                        switch 1 {
                        case 2147483647:
                            v6 -= v6
                        default:
                            v24 := []int{9223372036854775807, (fi12 | (-(fi20))), -1}
                            _ = v24
                        }
                    }
                    switch fi12 {
                    case -984:
                        switch fi20 {
                        case 4294967295:
                            v25 := []int{9223372036854775807, fi20}
                            switch (((p0 + p0) + ((p0 + p0) + ((((p0 + ((("\n\t" + p0) + "foo bar") + p0)) + ((p0 + (((p0 + p0) + p0) + ("\n\t" + (p0 + p0)))) + "a")) + p0) + ""))) + p0) {
                            case "foo bar":
                                var v26 string = ""
                                _ = v26
                            case "a":
                                v27 := GenS1{F0: -72}
                                _ = v27
                            case "":
                                var v28 int = (fi20 >> ((2) & 63))
                                _ = v28
                            default:
                                v25 = append(v25, (fi12 ^ (fi20 - fi12)))
                            }
                            _ = len(v25)
                            _ = v25
                        case 602:
                            if (-1 <= int64(fi12)) {
                                var v29 uint = uint(v19)
                                _ = v29
                            }
                        case 128:
                            fmt.Println(p1, p1)
                        }
                    case 945:
                        var v30 float32 = 3.14159
                        _ = v30
                    case 588:
                        if false {
                            if (-128 == -1) {
                                fmt.Println(fi20, p0)
                                var v31 string = p0
                                switch fi20 {
                                case 4294967295:
                                    if (false || true) {
                                        v32 := GenS3{F0: 544, F1: p0}
                                        v33 := GenS4{F0: 545, F1: (^(p1))}
                                        var v34 float64 = -2.25
                                        v34 = (float64(3.14159) - v34)
                                        _ = v32
                                        _ = v33
                                        _ = v34
                                    } else {
                                        v35 := []int{584, 337}
                                        fmt.Println(v19, v6, v19)
                                        _ = v35
                                    }
                                    var v36 uint8 = 127
                                    var v37 uint16 = p1
                                    _ = v36
                                    _ = v37
                                default:
                                    switch fi12 {
                                    case 2147483647:
                                        if true {
                                            if (true || true) {
                                                if true {
                                                    v19 = (v19 + v19)
                                                    var v38 uint8 = v19
                                                    v39 := GenS5{F0: p1, F1: 4294967295, F2: 784}
                                                    var v40 uint8 = v19
                                                    _ = v38
                                                    _ = v39
                                                    _ = v40
                                                } else {
                                                    if (false || true) {
                                                        var v41 uint64 = 0
                                                        switch (^(fi20)) {
                                                        case -1:
                                                            var v42 uint32 = 377
                                                            _ = v42
                                                        default:
                                                            var v43 int32 = -293
                                                            _ = v43
                                                        }
                                                        v44 := GenS1{F0: 1000}
                                                        _ = v41
                                                        _ = v44
                                                    } else {
                                                        v45 := GenS3{F0: 2147483648, F1: v31}
                                                        fmt.Println(p1)
                                                        _ = v45
                                                    }
                                                    var v46 int64 = (v6 & v6)
                                                    v47 := GenS6{F0: (p1 - uint16(v19))}
                                                    _ = v46
                                                    _ = v47
                                                }
                                                var v48 int = int(v19)
                                                _ = v48
                                            } else {
                                                fmt.Println(v19, p0, p1)
                                                v49 := []int{fi12, (fi20 % ((fi20) | 1))}
                                                _ = v49
                                            }
                                            if true {
                                                fmt.Println(fi20, p0)
                                            } else {
                                                var v50 int64 = (v6 >> ((0) & 63))
                                                _ = v50
                                            }
                                        }
                                    case 108:
                                        v51 := GenS7{F0: -2147483648}
                                        _ = v51
                                    }
                                }
                                _ = v31
                            }
                            fmt.Println(v6, fi12, p0)
                        } else {
                            if false {
                                v52 := []int{(fi20 >> ((359) & 63))}
                                var v53 uint32 = 443
                                switch (fi20 / ((fi20) | 1)) {
                                case 1:
                                    v6 = v6
                                    fmt.Println(v53, v19)
                                    v53 = v53
                                case 4294967295:
                                    v6 = v6
                                default:
                                    v19 = 0
                                    var v54 int16 = 0
                                    _ = v54
                                }
                                _ = v52
                                _ = v53
                            }
                            var v55 int32 = -234
                            _ = v55
                        }
                    }
                    v56 := GenS8{F0: uint(p1), F1: -620}
                    _ = v56
                }
                _ = fi20
            }
            _ = v19
        }
        fmt.Println(fi12, fi12)
        _ = fi12
    }
    var v57 int64 = v6
    _ = p0
    _ = p1
    _ = v57
    _ = v6
    return p0
}
func genFn1(p0 uint64, p1 uint32, p2 string) uint16 {
    for fi58 := 0; fi58 < 12; fi58++ {
        v59 := []int{(fi58 | fi58), (fi58 & 42), fi58}
        switch ((((p2 + p2) + ("hello" + p2)) + p2) + p2) {
        case "a":
            p1 -= p1
        }
        var v60 bool = (4294967295 >= 4294967295)
        _ = fi58
        _ = v59
        _ = v60
    }
    for fi61 := 0; fi61 < 14; fi61++ {
        p2 = "foo bar"
        v62 := GenS9{F0: 0}
        _ = fi61
        _ = v62
    }
    _ = p0
    _ = p1
    _ = p2
    return uint16(p0)
}

func main() {
    var v1 int32 = -2147483648
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    var v4 int64 = -1
    var v5 int = 127
    v1 = (v1 >> ((218) & 63))
    var v63 string = "\n\t"
    var v64 int64 = v4
    fmt.Println(v5)
    v63 = (v63 + "foo bar")
    for fi65 := 0; fi65 < 8; fi65++ {
        switch (v5 + fi65) {
        case 32767:
            v66 := []int{}
            _ = v66
        }
        var v67 int32 = 642
        var v68 string = (v63 + ("foo bar" + ("\n\t" + "a")))
        _ = fi65
        _ = v67
        _ = v68
    }
    for fi69 := 0; fi69 < 16; fi69++ {
        var v70 uint8 = 39
        v70 = v70
        v64 = v4
        _ = fi69
        _ = v70
    }
    var v71 uint = 0
    if (!(v64 >= (^(v4)))) {
        for fi72 := 0; fi72 < 8; fi72++ {
            v73 := v2[fi72]
            v5 = fi72
            v74 := GenS2{F0: v1, F1: 127}
            _ = fi72
            _ = v73
            _ = v74
        }
    } else {
        var v75 int16 = -417
        var v76 uint32 = 2
        var v77 int8 = -111
        _ = v75
        _ = v76
        _ = v77
    }
    if (true || false) {
        var v78 uint = (v71 & v71)
        _ = v78
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v63
    _ = v64
    _ = v71
}
