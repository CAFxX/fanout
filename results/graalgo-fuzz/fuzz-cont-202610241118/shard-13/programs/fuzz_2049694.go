package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 uint32; F2 int8 }
type GenS4 struct { F0 int64; F1 uint; F2 uint }
type GenS5 struct { F0 int64; F1 int; F2 uint16 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 int16; F1 int8; F2 uint8 }
func genFn0(p0 uint64, p1 uint, p2 uint32) bool {
    var v4 uint32 = (p2 & uint32(p1))
    v5 := GenS1{F0: 42}
    fmt.Println(p2, p1, p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return false
}

func main() {
    var v1 int8 = 42
    var v2 uint64 = 4294967295
    var v3 uint = 255
    v6 := GenS1{F0: (v3 << ((v3) & 63))}
    switch 9223372036854775807 {
    case 0:
        fmt.Println(v3, v1)
    }
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 int8 = (-(v1))
        v3 += v3
        _ = fi7
        _ = v8
    }
    if (-595 <= 9223372036854775807) {
        v9 := []int{-724}
        fmt.Println(v3, v2)
        v10 := GenS2{F0: 33}
        _ = v10
        _ = v9
    }
    if (4294967295 <= 65535) {
        switch "" {
        case "foo bar":
            v2 = (v2 - (v2 - (^(v2))))
        case "":
            var v11 uint32 = 2
            _ = v11
        case "\n\t":
            if (!false) {
                var v12 uint = (v3 * v3)
                for fi13 := 0; fi13 < 3; fi13++ {
                    for fi14 := 0; fi14 < 14; fi14++ {
                        if (42 > 1000) {
                            var v15 uint32 = 2147483648
                            v16 := GenS2{F0: 128}
                            var v17 int32 = 42
                            fmt.Println(v17, fi13)
                            if (!true) {
                                if false {
                                    var v18 int64 = 42
                                    switch "hello" {
                                    case "hello":
                                        fmt.Println(v18, v3)
                                    default:
                                        fmt.Println(v17, v1)
                                    }
                                    _ = v18
                                }
                                var v19 uint = v12
                                _ = v19
                            }
                            switch (((("foo bar" + "foo bar") + "a") + "hello") + "a") {
                            case "hello":
                                var v20 int64 = -9223372036854775808
                                if true {
                                    v21 := GenS1{F0: (v12 + (v3 * v12))}
                                    v1 += v1
                                    _ = v21
                                }
                                _ = v20
                            case "\n\t":
                                if false {
                                    v22 := []int{127, (fi14 ^ fi13), (fi13 + 32768)}
                                    var v23 int32 = 1
                                    var v24 float64 = 1.5
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                            default:
                                fmt.Println(v12, v12, v15)
                            }
                            switch fi13 {
                            case -466:
                                var v25 float64 = 3.14159
                                _ = v25
                            case -9223372036854775808:
                                if ((v17 % (((v17 >> (((^(v12))) & 63))) | 1)) >= v17) {
                                    if true {
                                        v26 := []int{(fi14 | (fi14 | fi13)), (fi14 ^ fi14), 0}
                                        if true {
                                            v27 := GenS1{F0: (v12 | uint(v2))}
                                            if ((0.1 > -2.25) && (!(v2 < v2))) {
                                                v28 := GenS3{F0: -443, F1: v15, F2: v1}
                                                _ = v28
                                            }
                                            _ = v27
                                        } else {
                                            var v29 int8 = 127
                                            switch "a" {
                                            case "foo bar":
                                                switch "hello" {
                                                case "":
                                                    fmt.Println(v3, v17)
                                                case "hello":
                                                    v15 = (v15 + (v15 - v15))
                                                }
                                            }
                                            _ = v29
                                        }
                                        _ = v26
                                    }
                                    v30 := GenS4{F0: 281, F1: (^(v3)), F2: 932}
                                    _ = v30
                                }
                            case -469:
                                if ((uint32(v12) == 1) || false) {
                                    if false {
                                        var v31 uint16 = 1
                                        var v32 int8 = -17
                                        v33 := []int{}
                                        _ = v31
                                        _ = v32
                                        _ = v33
                                    } else {
                                        v1 -= (v1 - (v1 & v1))
                                    }
                                    v12 = 336
                                    if true {
                                        if true {
                                            var v34 uint = (v12 << ((v3) & 63))
                                            v12 = v3
                                            v3 = (v34 ^ (v34 ^ (v34 & (v34 ^ 0))))
                                            v35 := []int{fi13, (-(fi13)), (fi14 + 67), fi13}
                                            _ = v34
                                            _ = v35
                                        }
                                        v1 += v1
                                        if (int8(v17) == v1) {
                                            var v36 uint = v12
                                            fmt.Println(v2, v3)
                                            _ = v36
                                        }
                                    } else {
                                        switch 9223372036854775807 {
                                        case 32767:
                                            var v37 uint16 = 32767
                                            _ = v37
                                        case 0:
                                            var v38 string = ("\n\t" + "\n\t")
                                            _ = v38
                                        case 32768:
                                            v2 = 42
                                        default:
                                            if (!true) {
                                                v12 = 672
                                                v2 = 49
                                            } else {
                                                v39 := GenS5{F0: 161, F1: fi14, F2: 65535}
                                                _ = v39
                                            }
                                        }
                                        v15 = (v15 ^ 0)
                                        var v40 int64 = -1
                                        _ = v40
                                    }
                                    var v41 int64 = 393
                                    _ = v41
                                } else {
                                    v42 := GenS6{F0: v17, F1: 1}
                                    v43 := GenS2{F0: 42}
                                    v44 := []int{-9223372036854775808}
                                    _ = v42
                                    _ = v43
                                    _ = v44
                                }
                            }
                            _ = v15
                            _ = v16
                            _ = v17
                        }
                        _ = fi14
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                fmt.Println(v1, v3)
            }
        default:
            v2 = v2
        }
        if false {
            switch "hello" {
            case "a":
                for fi45 := 0; fi45 < 11; fi45++ {
                    switch (fi45 ^ fi45) {
                    case -2147483648:
                        v1 -= (v1 ^ v1)
                    default:
                        var v46 int = (fi45 >> ((0) & 63))
                        _ = v46
                    }
                    var v47 int16 = 42
                    v48 := GenS6{F0: -1, F1: uint8(v3)}
                    _ = fi45
                    _ = v47
                    _ = v48
                }
            case "foo bar":
                if (1 > 2147483647) {
                    v49 := []int{1, int(v1)}
                    fmt.Println(v1)
                    for fi50 := 0; fi50 < 17; fi50++ {
                        v51 := GenS7{F0: -32768, F1: (v1 | 127), F2: 255}
                        _ = len(v49)
                        var v52 int64 = -1
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                } else {
                    v53 := []int{4294967295, 1}
                    _ = v53
                }
            case "":
                v1 = (v1 >> ((2) & 63))
            }
        }
        var v54 uint32 = 97
        _ = v54
    } else {
        v2 -= 42
    }
    v1 = (^(v1))
    v3 -= (v3 & 42)
    var v55 int64 = 42
    v1 += v1
    var v56 float32 = float32(3.14159)
    var v57 uint64 = (v2 - (v2 >> ((42) & 63)))
    v58 := []int{-862, 32767, -9223372036854775808, 32768}
    if false {
        fmt.Println(v57)
        switch 127 {
        case 42:
            var v59 int8 = (-(v1))
            _ = v59
        case -826:
            var v60 float64 = 0.0
            _ = v60
        case -209:
            v56 = v56
        default:
            switch 577 {
            case -1:
                var v61 uint8 = 36
                _ = v61
            case 2147483647:
                fmt.Println(v2, v56, v57)
            case 42:
                v62 := GenS3{F0: 32767, F1: 788, F2: v1}
                _ = v62
            }
        }
        var v63 uint8 = 37
        _ = v63
    } else {
        _ = len(v58)
        fmt.Println(v2, v2)
        v57 = v57
    }
    fmt.Println(v57)
    for fi64 := 0; fi64 < 1; fi64++ {
        var v65 int32 = 775
        _ = fi64
        _ = v65
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v55:", v55)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("slicelen v58:", len(v58), v58)
    _ = v1
    _ = v2
    _ = v3
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
