package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16; F1 uint64 }
type GenS7 struct { F0 int64; F1 int8; F2 uint8 }
type GenS8 struct { F0 uint32; F1 uint16; F2 int }
func genFn0(p0 string, p1 uint32) int {
    var v6 int = -686
    for fi7 := 0; fi7 < 1; fi7++ {
        v8 := GenS1{F0: 720}
        switch "a" {
        case "foo bar":
            var v9 uint64 = 588
            _ = v9
        }
        var v10 float64 = (float64(0.0) + 3.14159)
        _ = fi7
        _ = v10
        _ = v8
    }
    v11 := GenS1{F0: 255}
    if (!true) {
        v12 := []int{v6, 198, 4294967295}
        v13 := GenS2{F0: 32767}
        p1 = (p1 * (p1 - (p1 - (p1 % (((^(p1))) | 1)))))
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v6
    return v6
}
func genFn1() int {
    var v14 uint16 = 601
    v15 := GenS3{F0: 255}
    v16 := GenS1{F0: (v14 + (v14 & v14))}
    _ = v14
    _ = v15
    _ = v16
    return -2147483648
}

func main() {
    var v1 uint8 = 42
    var v2 uint = 18446744073709551615
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v5 := []uint32{}
    for fi17 := 0; fi17 < 3; fi17++ {
        v1 -= (^(v1))
        _ = fi17
    }
    if len(v5) > 0 {
        v5[0] = 0
    }
    switch 32768 {
    case -1:
        fmt.Println(v2, v2)
    case 4294967295:
        v18 := GenS4{F0: 42, F1: 9223372036854775807}
        _ = v18
    case -675:
        switch "hello" {
        case "a":
            if (!false) {
                if false {
                    var v19 int = 42
                    for fi20 := 0; fi20 < 12; fi20++ {
                        switch fi20 {
                        case 42:
                            delete(v3, 2)
                        default:
                            var v21 uint64 = uint64(v1)
                            _ = v21
                        }
                        v22 := GenS5{F0: (fi20 + (v19 & fi20))}
                        _ = fi20
                        _ = v22
                    }
                    _ = v19
                }
            } else {
                fmt.Println(v1, v1)
                fmt.Println(v1, v1)
            }
        case "":
            v3[2147483647] = "foo bar"
        default:
            var v23 int = genFn1()
            _ = v23
        }
    default:
        v2 -= v2
    }
    v24 := GenS6{F0: 42, F1: 228}
    var v25 uint64 = 4294967295
    v25 = v25
    v25 -= 44
    var v26 float32 = -2.25
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (-1 < 1) {
        var v27 uint = (v2 * v2)
        switch 818 {
        case -367:
            v28 := GenS2{F0: 42}
            _ = v28
        default:
            if true {
                v1 = v1
                if (((false || (906 >= 9223372036854775807)) || (-9223372036854775808 <= -9223372036854775808)) || false) {
                    if (!((!false) && (!(v27 >= (v2 + (v2 & (v27 | (v27 & uint(v1))))))))) {
                        if false {
                            v1 = (^(v1))
                        }
                        var v29 int16 = -910
                        var v30 int32 = 998
                        _ = v29
                        _ = v30
                    }
                } else {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        v2 = v2
                        var v32 int8 = 1
                        _ = fi31
                        _ = v32
                    }
                    for fi33 := 0; fi33 < 14; fi33++ {
                        v25 += 18446744073709551615
                        for fi34 := 0; fi34 < 17; fi34++ {
                            var v35 int64 = 312
                            v27 = (v2 / ((4294967295) | 1))
                            _ = fi34
                            _ = v35
                        }
                        var v36 float64 = (0.0 - 1e10)
                        _ = fi33
                        _ = v36
                    }
                    switch "foo bar" {
                    case "":
                        v37 := GenS5{F0: 740}
                        _ = v37
                    case "\n\t":
                        switch -2147483648 {
                        case 32767:
                            var v38 uint32 = 624
                            _ = v38
                        case 9223372036854775807:
                            var v39 int8 = 127
                            _ = v39
                        }
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    default:
                        var v40 uint32 = 572
                        _ = v40
                    }
                }
                for fi41 := 0; fi41 < 8; fi41++ {
                    v1 = v1
                    var v42 uint32 = 0
                    var v43 string = ("\n\t" + "foo bar")
                    _ = fi41
                    _ = v42
                    _ = v43
                }
            }
        }
        switch 127 {
        case 0:
            fmt.Println(v1, v25)
        case 330:
            var v44 uint8 = v1
            _ = v44
        case 4294967295:
            fmt.Println(v26, v1)
        }
        _ = v27
    } else {
        switch "" {
        case "\n\t":
            var v45 int = 946
            _ = v45
        case "a":
            var v46 int32 = 1
            _ = v46
        case "hello":
            var v47 float32 = v26
            _ = v47
        }
        v48 := GenS3{F0: (v1 | 42)}
        for fi49 := 0; fi49 < 16; fi49++ {
            v2 = 255
            _ = fi49
        }
        _ = v48
    }
    for fi50 := 0; fi50 < 16; fi50++ {
        for fi51 := 0; fi51 < 11; fi51++ {
            v4[(fi50 + (fi50 >> ((v2) & 63)))] = -765
            switch ("hello" + ("" + "\n\t")) {
            case "hello":
                v1 = v1
            case "\n\t":
                var v52 int16 = 833
                _ = v52
            case "a":
                var v53 int16 = -1
                _ = v53
            default:
                var v54 uint16 = 2
                switch (^(fi51)) {
                case -1:
                    var v55 uint8 = uint8(v2)
                    _ = v55
                default:
                    v56 := v4[(fi51 & 2147483647)]
                    v57 := GenS7{F0: -1, F1: -45, F2: 154}
                    var v58 uint16 = (v54 & (v54 >> (((^(v2))) & 63)))
                    _ = v56
                    _ = v57
                    _ = v58
                }
                _ = v54
            }
            _ = fi51
        }
        var v59 uint32 = 297
        var v60 uint32 = 2147483648
        _ = fi50
        _ = v59
        _ = v60
    }
    v61 := GenS5{F0: -702}
    v62 := GenS8{F0: 1, F1: 154, F2: -596}
    v3[764] = ""
    for fi63 := 0; fi63 < 11; fi63++ {
        var v64 int32 = -1
        _ = fi63
        _ = v64
    }
    for fi65 := 0; fi65 < 8; fi65++ {
        var v66 int32 = -977
        var v67 uint8 = 255
        _ = fi65
        _ = v66
        _ = v67
    }
    switch "\n\t" {
    case "":
        if (v25 > v25) {
            var v68 bool = true
            fmt.Println(v25, v1)
            _ = v68
        } else {
            var v69 bool = (762 != 2147483647)
            _ = v69
        }
    case "\n\t":
        delete(v4, 4294967295)
    case "hello":
        v70 := GenS7{F0: 1, F1: 20, F2: v1}
        _ = v70
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v61
    _ = v62
}
