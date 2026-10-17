package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int; F1 uint16 }
type GenS4 struct { F0 uint8 }
func genFn0(p0 string, p1 uint, p2 int16) int16 {
    fmt.Println(p2, p0)
    var v5 uint64 = 1
    p0 = ((p0 + p0) + p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p2
}
func genFn1() int8 {
    v6 := []int{188, -2147483648, 128}
    var v7 uint8 = 255
    var v8 int64 = -1
    v7 = (v7 & 24)
    for fi9 := 0; fi9 < 6; fi9++ {
        switch "a" {
        case "foo bar":
            switch (fi9 << ((255) & 63)) {
            case 4294967295:
                var v10 int16 = 365
                _ = v10
            case 97:
                for fi11 := 0; fi11 < 2; fi11++ {
                    var v12 uint8 = 128
                    _ = fi11
                    _ = v12
                }
            case -254:
                for fi13 := 0; fi13 < 2; fi13++ {
                    v8 += v8
                    var v14 uint = 232
                    _ = fi13
                    _ = v14
                }
            }
        case "hello":
            if len(v6) > 0 {
                v6[0] = fi9
            }
        case "\n\t":
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        default:
            v7 -= 255
        }
        var v15 bool = false
        _ = fi9
        _ = v15
    }
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    var v1 float32 = 1.5
    var v2 float32 = (1.5 + (1e10 + v1))
    var v3 uint8 = 0
    v4 := make(map[int]string)
    v1 = (float32(0.0) / 1.5)
    for fi16 := 0; fi16 < 1; fi16++ {
        var v17 int32 = 1
        fmt.Println(v2)
        _ = fi16
        _ = v17
    }
    var v18 bool = false
    for fi19 := 0; fi19 < 9; fi19++ {
        var v20 int64 = 723
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 13; fi21++ {
        if v18 {
            for fi22 := 0; fi22 < 20; fi22++ {
                v1 = v2
                switch 417 {
                case 2147483647:
                    v2 += v2
                }
                _ = fi22
            }
            v23 := GenS1{F0: 2, F1: 1}
            v24 := GenS2{F0: 0}
            _ = v23
            _ = v24
        } else {
            if (v18 && v18) {
                switch fi21 {
                case 128:
                    v2 = 1e10
                case 0:
                    v2 = v1
                case 42:
                    for fi25 := 0; fi25 < 2; fi25++ {
                        fmt.Println(fi25)
                        if (false && false) {
                            v3 = (v3 << ((4294967295) & 63))
                            v26 := GenS1{F0: 857, F1: 42}
                            v27 := GenS1{F0: 0, F1: 18446744073709551615}
                            _ = v26
                            _ = v27
                        }
                        var v28 uint64 = 0
                        v29 := v4[fi21]
                        _ = fi25
                        _ = v28
                        _ = v29
                    }
                default:
                    var v30 float64 = ((0.1 - (0.1 + 1e10)) / 1.0)
                    _ = v30
                }
                var v31 int16 = -32768
                v31 -= v31
                _ = v31
            }
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 bool = false
                v1 = v1
                _ = fi32
                _ = v33
            }
            switch fi21 {
            case 313:
                var v34 uint8 = (v3 / ((v3) | 1))
                _ = v34
            }
        }
        if (!(-1 > genFn1())) {
            v35 := GenS1{F0: 867, F1: 809}
            switch "foo bar" {
            case "a":
                var v36 int = (fi21 << ((2) & 63))
                _ = v36
            default:
                fmt.Println(v2, v18)
            }
            _ = v35
        }
        _ = fi21
    }
    v37 := []int{-2147483648, 4294967295}
    var v38 uint = 304
    if v18 {
        var v39 uint16 = 65535
        for fi40 := 0; fi40 < 15; fi40++ {
            var v41 int8 = -40
            _ = fi40
            _ = v41
        }
        var v42 uint16 = (^(v39))
        _ = v39
        _ = v42
    } else {
        v3 -= v3
        var v43 int32 = 1
        _ = v43
    }
    fmt.Println(v18, v18)
    for fi44 := 0; fi44 < 10; fi44++ {
        v45 := GenS1{F0: 677, F1: 629}
        switch "\n\t" {
        case "foo bar":
            _ = len(v4)
        case "\n\t":
            for fi46 := 0; fi46 < 17; fi46++ {
                if v18 {
                    v47 := GenS3{F0: fi44, F1: 255}
                    if (0 < 82) {
                        v1 += v1
                        v2 = v2
                    }
                    _ = v47
                } else {
                    switch 0 {
                    case 0:
                        fmt.Println(v18)
                    case 128:
                        var v48 uint16 = 0
                        _ = v48
                    default:
                        v37 = append(v37, fi44)
                    }
                }
                v49 := GenS4{F0: v3}
                _ = fi46
                _ = v49
            }
        }
        _ = fi44
        _ = v45
    }
    switch 254 {
    case 128:
        if v18 {
            var v50 uint16 = uint16(v3)
            v18 = v18
            switch 467 {
            case 264:
                if ("foo bar" <= "a") {
                    if true {
                        fmt.Println(v38, v3)
                        var v51 string = "foo bar"
                        v1 = v1
                        _ = v51
                    }
                    var v52 uint64 = 18446744073709551615
                    var v53 uint64 = v52
                    _ = v52
                    _ = v53
                }
            case -2147483648:
                var v54 string = ("\n\t" + (("a" + "") + "a"))
                _ = v54
            case 4294967295:
                v50 -= v50
            default:
                var v55 uint64 = 18446744073709551615
                _ = v55
            }
            _ = v50
        } else {
            v56 := GenS1{F0: 2147483648, F1: 4294967295}
            var v57 bool = v18
            _ = v56
            _ = v57
        }
    default:
        fmt.Println(v18)
    }
    v18 = (690 > 42)
    if (1e10 < (((1e10 * 3.14159) * 1.5) - (3.14159 - -2.25))) {
        switch "foo bar" {
        case "":
            v2 -= ((float32(1.5) * v2) / v1)
        case "hello":
            v4[127] = ("" + (("\n\t" + "foo bar") + ""))
        case "\n\t":
            v58 := GenS1{F0: 65535, F1: 261}
            _ = v58
        default:
            if true {
                if (v18 && (0 != -260)) {
                    var v59 int32 = 1
                    v60 := GenS2{F0: (-(v59))}
                    _ = v59
                    _ = v60
                } else {
                    for fi61 := 0; fi61 < 9; fi61++ {
                        var v62 int = fi61
                        _ = fi61
                        _ = v62
                    }
                    v3 = (v3 - 1)
                }
                for fi63 := 0; fi63 < 12; fi63++ {
                    for fi64 := 0; fi64 < 10; fi64++ {
                        var v65 int8 = 92
                        var v66 int64 = -1
                        _ = len(v37)
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    fmt.Println(v3, fi63, v3)
                    v18 = v18
                    _ = fi63
                }
            }
        }
    }
    var v67 bool = (v18 || v18)
    fmt.Println(v38)
    var v68 uint = v38
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v67:", v67)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v67
    _ = v68
}
