package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int32; F2 int32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int8; F2 int64 }
type GenS5 struct { F0 int64 }
func genFn0() int16 {
    if (!false) {
        var v6 uint = 609
        v7 := GenS1{F0: 919}
        _ = v6
        _ = v7
    }
    var v8 int64 = -1
    for fi9 := 0; fi9 < 6; fi9++ {
        if (false || false) {
            var v10 uint = 902
            if (v8 == v8) {
                v11 := []int{fi9}
                var v12 int8 = 42
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
                _ = v11
                _ = v12
            } else {
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 string = ((("foo bar" + ("a" + ("" + ""))) + "foo bar") + "a")
                    v15 := GenS1{F0: v8}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 uint64 = 1
                v17 := GenS2{F0: 32767, F1: 1, F2: -650}
                _ = v16
                _ = v17
            }
            v10 -= v10
            _ = v10
        }
        _ = fi9
    }
    v8 = v8
    _ = v8
    return -103
}
func genFn1(p0 uint, p1 uint) int32 {
    if false {
        switch "\n\t" {
        case "foo bar":
            var v18 uint32 = uint32(p0)
            _ = v18
        default:
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 bool = (!false)
                _ = fi19
                _ = v20
            }
        }
    } else {
        switch -2147483648 {
        case 9223372036854775807:
            switch -715 {
            case -9223372036854775808:
                if (-128 <= 42) {
                    var v21 float32 = 3.14159
                    _ = v21
                }
            case 9223372036854775807:
                var v22 int16 = 42
                _ = v22
            case -2147483648:
                p0 -= (^(p0))
            }
        case 32767:
            v23 := GenS2{F0: 42, F1: -1, F2: 365}
            _ = v23
        }
    }
    if false {
        fmt.Println(p1, p1)
    } else {
        if (255 >= 231) {
            switch 42 {
            case 4294967295:
                if ((((((-2.25 / -2.25) - 0.1) * (1e10 * 1.5)) / 1.0) + 1.5) == (1e10 * 0.0)) {
                    v24 := GenS2{F0: 255, F1: 874, F2: 2147483647}
                    v25 := []int{42}
                    _ = v24
                    _ = v25
                } else {
                    p0 -= (p1 ^ p1)
                    if true {
                        v26 := []int{-9223372036854775808, 81}
                        var v27 int32 = 1
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                        _ = v26
                        _ = v27
                    }
                    var v28 int32 = 0
                    _ = v28
                }
            default:
                p0 -= 255
            }
            if ((!true) || false) {
                if false {
                    v29 := GenS2{F0: 620, F1: 0, F2: 42}
                    var v30 int64 = -9223372036854775808
                    _ = v29
                    _ = v30
                }
            }
        } else {
            v31 := GenS1{F0: 42}
            v32 := GenS2{F0: 65535, F1: 2147483647, F2: 1}
            _ = v31
            _ = v32
        }
        v33 := []int{-279, 32768, 127, 1}
        _ = v33
    }
    var v34 uint32 = 65535
    _ = p0
    _ = p1
    _ = v34
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[int]bool)
    v3 := make(map[int]int)
    v4 := make(map[int]uint32)
    var v5 uint32 = 1
    if true {
        var v35 uint64 = 4294967295
        if true {
            if true {
                v35 = 59
            } else {
                v36 := v3[32768]
                _ = v36
            }
            var v37 uint8 = 131
            fmt.Println(v37, v5, v5)
            _ = v37
        } else {
            v38 := []int{42, 2147483647}
            for fi39 := 0; fi39 < 10; fi39++ {
                v38 = append(v38, fi39)
                _ = fi39
            }
            v40 := GenS1{F0: 652}
            _ = v38
            _ = v40
        }
        _ = v35
    } else {
        if true {
            v41 := GenS1{F0: 9223372036854775807}
            _ = v41
        } else {
            if (32767 <= uint16(v5)) {
                if false {
                    var v42 uint = 0
                    v43 := []int{-1, 4294967295, -680, -9223372036854775808}
                    _ = len(v43)
                    _ = v42
                    _ = v43
                } else {
                    v44 := []int{962}
                    _ = v44
                }
            }
            var v45 bool = false
            var v46 int = -973
            _ = v45
            _ = v46
        }
        fmt.Println(v5)
        v47 := GenS2{F0: 32767, F1: 1, F2: -1}
        _ = v47
    }
    var v48 int32 = 172
    fmt.Println(v48, v5)
    v49 := v2[722]
    var v50 string = "hello"
    v51 := GenS2{F0: 997, F1: (-(v48)), F2: (-(v48))}
    v52 := GenS2{F0: 484, F1: v48, F2: (v48 + v48)}
    if false {
        if v49 {
            switch v50 {
            case "hello":
                v49 = true
            case "foo bar":
                v49 = ((v5 + v5) >= (v5 | v5))
            case "":
                v49 = v49
            }
        }
        v48 -= 0
        if v49 {
            if v49 {
                var v53 uint64 = 2
                _ = v53
            } else {
                v54 := []int{2147483647, -9223372036854775808, -98, -9223372036854775808}
                _ = v54
            }
        }
    } else {
        fmt.Println(v48)
    }
    v55 := []int{42, 1, -2147483648}
    v56 := GenS3{F0: 1}
    v48 = 1
    v57 := GenS4{F0: (v5 | v5), F1: 42, F2: -9223372036854775808}
    var v58 uint8 = 214
    if (-2.25 < ((0.1 - 0.0) - ((3.14159 + 0.0) / (3.14159 / (1e10 + -2.25))))) {
        var v59 uint8 = (v58 | 42)
        switch "hello" {
        case "foo bar":
            var v60 uint8 = 128
            _ = v60
        case "":
            v58 = 42
        case "a":
            if ((-1 >= 9223372036854775807) && true) {
                v55 = append(v55, -1)
                if (!(!v49)) {
                    v48 = v48
                }
                v61 := v2[247]
                _ = v61
            } else {
                v49 = v49
            }
        default:
            _ = len(v2)
        }
        _ = v59
    }
    var v62 float64 = 0.1
    var v63 float32 = 1.5
    var v64 string = (v50 + ((v50 + v50) + v50))
    v65 := GenS5{F0: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("slicelen v55:", len(v55), v55)
    fmt.Println("v58:", v58)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
