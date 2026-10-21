package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 int; F2 uint8 }
type GenS4 struct { F0 bool; F1 int }
type GenS5 struct { F0 string; F1 int32; F2 uint32 }
type GenS6 struct { F0 uint8; F1 int32; F2 bool }
type GenS7 struct { F0 uint16; F1 int64 }
type GenS8 struct { F0 uint32; F1 int }
type GenS9 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS10 struct { F0 string; F1 int64 }
type GenS11 struct { F0 string; F1 int8 }
func genFn0(p0 uint) string {
    var v5 uint8 = uint8(p0)
    v6 := GenS1{F0: 42}
    if true {
        for fi7 := 0; fi7 < 19; fi7++ {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 int32 = 2147483647
                v10 := GenS2{F0: -1}
                var v11 string = ("" + "")
                v12 := []int{fi8, 128, 9223372036854775807, fi7}
                var v13 int16 = 0
                v13 += (v13 % (((-(v13))) | 1))
                switch ((v11 + (v11 + (("foo bar" + ((v11 + v11) + "a")) + v11))) + v11) {
                case "foo bar":
                    switch -747 {
                    case 0:
                        v14 := GenS2{F0: -1}
                        _ = v14
                    case -384:
                        switch v11 {
                        case "foo bar":
                            v9 -= 229
                        case "hello":
                            var v15 float32 = 0.0
                            var v16 uint = (^(p0))
                            _ = v15
                            _ = v16
                        case "a":
                            v9 += 0
                        }
                    }
                case "a":
                    v17 := GenS1{F0: -9223372036854775808}
                    _ = v17
                }
                _ = fi8
                _ = v10
                _ = v11
                _ = v12
                _ = v13
                _ = v9
            }
            for fi18 := 0; fi18 < 2; fi18++ {
                var v19 uint = (p0 + (p0 ^ p0))
                p0 += (p0 << (((v19 | (p0 + (p0 & (v19 ^ 2))))) & 63))
                var v20 float64 = (3.14159 - (3.14159 - float64(1.5)))
                _ = fi18
                _ = v19
                _ = v20
            }
            _ = fi7
        }
        v21 := GenS3{F0: false, F1: 4294967295, F2: 70}
        _ = v21
    }
    if false {
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v5 = 127
                var v24 int64 = 9223372036854775807
                v24 += (v24 << ((p0) & 63))
                _ = fi23
                _ = v24
            }
            _ = fi22
        }
    } else {
        for fi25 := 0; fi25 < 4; fi25++ {
            v26 := GenS4{F0: true, F1: -558}
            if true {
                fmt.Println(p0, fi25)
                v5 = v5
                var v27 int8 = int8(p0)
                _ = v27
            } else {
                v28 := []int{-1, (fi25 ^ (fi25 * fi25))}
                switch ("" + ("a" + (("a" + ("hello" + "foo bar")) + ("\n\t" + "foo bar")))) {
                case "a":
                    var v29 string = ("a" + ("\n\t" + ("hello" + "hello")))
                    _ = v29
                case "hello":
                    for fi30 := 0; fi30 < 10; fi30++ {
                        _ = len(v28)
                        if len(v28) >= 2 {
                            v28 = v28[1:len(v28)]
                        }
                        _ = fi30
                    }
                case "":
                    v28 = append(v28, (fi25 & (fi25 / ((fi25) | 1))))
                default:
                    switch fi25 {
                    case 592:
                        for fi31 := 0; fi31 < 11; fi31++ {
                            v5 -= v5
                            fmt.Println(fi25)
                            var v32 uint16 = 65535
                            _ = fi31
                            _ = v32
                        }
                    }
                }
                _ = v28
            }
            var v33 int16 = 1000
            _ = fi25
            _ = v26
            _ = v33
        }
        var v34 int8 = 1
        _ = v34
    }
    _ = p0
    _ = v5
    _ = v6
    return "foo bar"
}
func genFn1() uint64 {
    var v35 uint8 = 1
    v35 = v35
    v36 := []int{499, 0, 127}
    v35 -= v35
    _ = v35
    _ = v36
    return 42
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 string = "a"
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    for fi37 := 0; fi37 < 2; fi37++ {
        if ((false && false) || (!true)) {
            v2 = ((v2 + ("foo bar" + "hello")) + v2)
        }
        for fi38 := 0; fi38 < 10; fi38++ {
            v39 := []int{fi38, (fi38 >> ((1) & 63))}
            var v40 uint64 = v1
            switch (fi38 >> ((71) & 63)) {
            case 32767:
                if (!(!true)) {
                    var v41 float32 = float32(0.0)
                    if false {
                        fmt.Println(v1)
                    }
                    if true {
                        fmt.Println(fi38, v41, v1)
                        v2 = (v2 + v2)
                        delete(v3, (v2 + v2))
                    }
                    _ = v41
                }
                v42 := GenS5{F0: ("a" + (("hello" + (((((((("\n\t" + v2) + "a") + v2) + v2) + (v2 + v2)) + v2) + (v2 + v2)) + "a")) + v2)), F1: 42, F2: 0}
                v43 := GenS6{F0: 2, F1: 2147483647, F2: false}
                _ = v42
                _ = v43
            default:
                var v44 float32 = (1e10 / (1.5 / 1.5))
                _ = v44
            }
            var v45 bool = false
            _ = fi38
            _ = v39
            _ = v40
            _ = v45
        }
        _ = fi37
    }
    switch 356 {
    case 127:
        v46 := GenS1{F0: 9223372036854775807}
        _ = v46
    case 128:
        switch ("foo bar" + ("" + v2)) {
        case "hello":
            var v47 int = -2147483648
            _ = v47
        case "\n\t":
            if true {
                v48 := GenS7{F0: 358, F1: -9223372036854775808}
                _ = v48
            } else {
                var v49 bool = (!true)
                fmt.Println(v49, v49)
                var v50 int8 = 1
                _ = v49
                _ = v50
            }
        default:
            var v51 int16 = -50
            _ = v51
        }
    default:
        fmt.Println(v2)
    }
    switch v2 {
    case "\n\t":
        var v52 int32 = 42
        _ = v52
    case "foo bar":
        v53 := []int{-2147483648, 939, -1, -445}
        _ = v53
    case "":
        if (v2 > v2) {
            var v54 uint16 = 2
            var v55 uint32 = 0
            v56 := GenS4{F0: false, F1: -2147483648}
            _ = v54
            _ = v55
            _ = v56
        }
    default:
        fmt.Println(v1)
    }
    var v57 uint16 = 24
    v58 := []int{}
    v2 = ("" + ("" + (v2 + "\n\t")))
    var v59 uint16 = (v57 + v57)
    switch "\n\t" {
    case "a":
        v60 := GenS8{F0: 2147483648, F1: 42}
        _ = v60
    case "\n\t":
        if (127 < -128) {
            switch ("" + "a") {
            case "a":
                _ = len(v58)
            case "\n\t":
                fmt.Println(v2, v1, v2)
            default:
                v57 = 32767
            }
            switch 392 {
            case 128:
                var v61 int = 510
                _ = v61
            case -2147483648:
                var v62 int32 = 2147483647
                _ = v62
            default:
                v1 -= 1
            }
            v58 = append(v58, 2147483647)
        } else {
            var v63 uint16 = uint16(v1)
            _ = len(v4)
            if len(v58) >= 2 {
                v58 = v58[1:len(v58)]
            }
            _ = v63
        }
    default:
        v64 := GenS8{F0: 784, F1: 711}
        _ = v64
    }
    var v65 float32 = -2.25
    v66 := GenS9{F0: uint32(v57), F1: 0, F2: 928}
    for fi67 := 0; fi67 < 1; fi67++ {
        if len(v58) > 0 {
            v58[0] = (fi67 >> ((2) & 63))
        }
        var v68 int32 = -735
        v69 := GenS10{F0: v2, F1: 407}
        _ = fi67
        _ = v68
        _ = v69
    }
    var v70 bool = (2 > uint8(v1))
    var v71 uint8 = 42
    v71 -= v71
    v72 := GenS11{F0: v2, F1: -128}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v58:", len(v58), v58)
    fmt.Println("v59:", v59)
    fmt.Println("v65:", v65)
    fmt.Println("v70:", v70)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v57
    _ = v58
    _ = v59
    _ = v65
    _ = v66
    _ = v70
    _ = v71
    _ = v72
}
