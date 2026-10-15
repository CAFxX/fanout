package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int; F1 int8 }
func genFn0(p0 bool, p1 int32) bool {
    p0 = true
    v7 := []int{-1, 42, -9223372036854775808, 994}
    v8 := GenS1{F0: 63, F1: -9223372036854775808}
    switch 127 {
    case -1:
        var v9 uint8 = 146
        _ = v9
    case 2147483647:
        var v10 float64 = 3.14159
        _ = v10
    case 1:
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return p0
}

func main() {
    var v1 bool = false
    var v2 uint64 = 94
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    v5 := make(map[uint32]bool)
    var v6 uint64 = 81
    v11 := GenS2{F0: 1}
    switch 127 {
    case 32768:
        v1 = false
    case 42:
        switch 42 {
        case 82:
            fmt.Println(v6)
        case 2147483647:
            v1 = (v1 && (v1 || false))
        }
    case 127:
        var v12 float64 = 1e10
        _ = v12
    default:
        var v13 uint32 = 653
        _ = v13
    }
    switch 127 {
    case 273:
        var v14 float64 = (0.1 * 1.5)
        _ = v14
    case 702:
        if true {
            v15 := []int{42, -9223372036854775808, 4294967295}
            var v16 int8 = 1
            for fi17 := 0; fi17 < 14; fi17++ {
                v18 := GenS3{F0: (v6 << ((189) & 63)), F1: 128}
                var v19 bool = true
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = v15
            _ = v16
        } else {
            var v20 string = "\n\t"
            _ = v20
        }
    default:
        fmt.Println(v1, v2)
    }
    var v21 uint32 = 2
    v6 = 18446744073709551615
    if ((2147483647 <= int32(v6)) && v1) {
        for fi22 := 0; fi22 < 20; fi22++ {
            switch "foo bar" {
            case "":
                fmt.Println(v2)
            default:
                delete(v3, v21)
            }
            v21 = v21
            _ = fi22
        }
        _ = len(v3)
        v6 -= v6
    }
    v23 := GenS3{F0: (v2 >> ((18446744073709551615) & 63)), F1: 128}
    switch 32767 {
    case -9223372036854775808:
        _ = len(v5)
    case 128:
        var v24 string = ((("\n\t" + "") + "hello") + "")
        _ = v24
    }
    v21 = 964
    v25 := GenS1{F0: 127, F1: -9223372036854775808}
    switch 1 {
    case 252:
        var v26 uint64 = v6
        _ = v26
    case -853:
        for fi27 := 0; fi27 < 12; fi27++ {
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 uint = 1
                switch fi28 {
                case 297:
                    var v30 int16 = 1000
                    _ = v30
                case 852:
                    if v1 {
                        if (-1 <= -1) {
                            v31 := GenS2{F0: -617}
                            var v32 int64 = 1
                            v33 := GenS1{F0: 241, F1: fi28}
                            _ = v31
                            _ = v32
                            _ = v33
                        }
                        v34 := GenS4{F0: v21}
                        _ = v34
                    } else {
                        var v35 int16 = -1
                        v1 = (-2.25 < 0.0)
                        v6 -= uint64(v21)
                        var v36 bool = (!v1)
                        _ = v35
                        _ = v36
                    }
                    var v37 float32 = (1.5 - 0.0)
                    v38 := GenS3{F0: 0, F1: 0}
                    _ = v37
                    _ = v38
                default:
                    switch "foo bar" {
                    case "\n\t":
                        var v39 float64 = 1e10
                        _ = v39
                    default:
                        v3[2147483647] = v21
                    }
                }
                switch -1 {
                case 4294967295:
                    _ = len(v5)
                case 9223372036854775807:
                    fmt.Println(fi27, v29, fi27)
                case 42:
                    v2 = (v2 | (v2 * v6))
                }
                var v40 uint16 = 42
                _ = fi28
                _ = v29
                _ = v40
            }
            if v1 {
                switch "a" {
                case "hello":
                    var v41 int32 = -1
                    _ = v41
                case "":
                    v42 := []int{128, 262, fi27, 128}
                    _ = v42
                case "foo bar":
                    if v1 {
                        fmt.Println(fi27, v1)
                    } else {
                        var v43 float64 = -2.25
                        _ = v43
                    }
                default:
                    v1 = (float32((3.14159 / (1e10 - -2.25))) == ((3.14159 - 1e10) - float32(1.5)))
                }
                var v44 float64 = 1e10
                v6 += (v2 % ((v2) | 1))
                _ = v44
            } else {
                for fi45 := 0; fi45 < 1; fi45++ {
                    switch "foo bar" {
                    case "a":
                        v46 := v4[(-(fi45))]
                        _ = v46
                    case "":
                        v47 := []int{fi27, fi27, (fi45 >> ((2) & 63)), (fi27 << ((uint(v21)) & 63))}
                        _ = v47
                    default:
                        v21 = v21
                    }
                    v48 := GenS5{F0: "hello"}
                    _ = fi45
                    _ = v48
                }
                if (!true) {
                    var v49 int64 = -99
                    v50 := GenS3{F0: v6, F1: 125}
                    _ = v49
                    _ = v50
                }
                v1 = v1
            }
            _ = fi27
        }
    case -267:
        for fi51 := 0; fi51 < 1; fi51++ {
            if false {
                for fi52 := 0; fi52 < 2; fi52++ {
                    v21 -= (v21 & v21)
                    v53 := []int{(fi51 % (((^(fi51))) | 1))}
                    v1 = (32767 > 74)
                    _ = fi52
                    _ = v53
                }
                for fi54 := 0; fi54 < 9; fi54++ {
                    v6 += 42
                    var v55 uint64 = v2
                    var v56 int16 = 742
                    _ = fi54
                    _ = v55
                    _ = v56
                }
            } else {
                v1 = v1
                fmt.Println(fi51)
                for fi57 := 0; fi57 < 2; fi57++ {
                    switch 510 {
                    case -2147483648:
                        var v58 float64 = 0.1
                        _ = v58
                    }
                    v59 := GenS6{F0: 127, F1: -105}
                    fmt.Println(v2, v1, v2)
                    _ = fi57
                    _ = v59
                }
            }
            _ = fi51
        }
    default:
        fmt.Println(v1)
    }
    if false {
        var v60 uint64 = 0
        v1 = false
        v1 = ((!false) || false)
        _ = v60
    } else {
        for fi61 := 0; fi61 < 3; fi61++ {
            fmt.Println(v1)
            var v62 int8 = 42
            _ = fi61
            _ = v62
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
