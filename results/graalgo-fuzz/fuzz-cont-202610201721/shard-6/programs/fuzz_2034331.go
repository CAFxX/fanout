package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int64; F1 uint32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int; F1 uint32; F2 int64 }
type GenS8 struct { F0 int; F1 uint64 }
type GenS9 struct { F0 string; F1 uint16; F2 bool }
func genFn0(p0 int, p1 uint, p2 uint) int {
    switch "foo bar" {
    case "\n\t":
        for fi7 := 0; fi7 < 15; fi7++ {
            var v8 int8 = 1
            fmt.Println(fi7, p0)
            v9 := []int{fi7, 9223372036854775807}
            _ = fi7
            _ = v8
            _ = v9
        }
    case "foo bar":
        p1 -= p1
    default:
        var v10 float32 = 1e10
        _ = v10
    }
    v11 := GenS1{F0: -1, F1: 163}
    v12 := GenS1{F0: 1, F1: (p1 % ((42) | 1))}
    v13 := []int{(p0 | p0)}
    for fi14 := 0; fi14 < 5; fi14++ {
        p2 = p2
        var v15 int64 = 42
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return -545
}
func genFn1(p0 int64, p1 int, p2 string) int8 {
    v16 := GenS2{F0: 2147483648}
    fmt.Println(p1, p2)
    var v17 uint32 = 2
    v18 := GenS3{F0: 0, F1: 0}
    switch 42 {
    case 32767:
        v19 := GenS2{F0: (v17 << ((42) & 63))}
        _ = v19
    case 32768:
        var v20 bool = (32767 >= 42)
        _ = v20
    default:
        for fi21 := 0; fi21 < 17; fi21++ {
            for fi22 := 0; fi22 < 20; fi22++ {
                var v23 int8 = 100
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    _ = v17
    _ = v18
    return 61
}

func main() {
    var v1 uint8 = 255
    v2 := []string{("foo bar" + ("" + (("" + ("\n\t" + "")) + "hello")))}
    var v3 string = ("a" + "hello")
    v4 := []uint32{4294967295}
    var v5 int64 = 0
    var v6 uint16 = 852
    v3 = v3
    for fi24 := 0; fi24 < 13; fi24++ {
        v3 = (((v3 + v3) + v3) + ((v3 + "\n\t") + (("\n\t" + "hello") + ("hello" + v3))))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        fmt.Println(v1)
        _ = fi24
    }
    if len(v4) > 0 {
        v4[0] = 612
    }
    var v25 uint16 = 42
    for fi26 := 0; fi26 < 3; fi26++ {
        fmt.Println(v3, v5)
        var v27 uint64 = 1
        v5 = (v5 * (v5 << ((1) & 63)))
        _ = fi26
        _ = v27
    }
    var v28 int = 369
    switch (v28 | (v28 % ((v28) | 1))) {
    case 0:
        for fi29 := 0; fi29 < 10; fi29++ {
            fmt.Println(fi29)
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            } else {
                v25 = (v6 << ((42) & 63))
                _ = len(v4)
                fmt.Println(v1, v5)
            }
            _ = fi29
        }
    case -2147483648:
        if true {
            v1 = 172
            if len(v2) > 0 {
                v2[0] = v3
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                switch -862 {
                case -2147483648:
                    var v31 uint32 = 0
                    _ = v31
                }
                _ = fi30
            }
        }
    case -1:
        fmt.Println(v28, v3, v28)
    }
    v2 = append(v2, v3)
    if true {
        v2 = append(v2, "a")
        var v32 int32 = 42
        v33 := GenS4{F0: v25}
        _ = v32
        _ = v33
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi34 := 0; fi34 < 15; fi34++ {
        if true {
            var v35 uint16 = (^(v6))
            var v36 int16 = 0
            _ = v35
            _ = v36
        } else {
            if true {
                for fi37 := 0; fi37 < 7; fi37++ {
                    v6 = (v25 >> ((152) & 63))
                    v2 = append(v2, v3)
                    var v38 int = fi37
                    fmt.Println(v28)
                    _ = fi37
                    _ = v38
                }
            }
            for fi39 := 0; fi39 < 13; fi39++ {
                v40 := GenS5{F0: (v5 % (((v5 | int64(fi34))) | 1)), F1: 143}
                switch (fi39 * v28) {
                case 2147483647:
                    fmt.Println(v25, v5, fi39)
                case -635:
                    var v41 int64 = (v5 ^ v5)
                    _ = v41
                case -2147483648:
                    var v42 int32 = -2147483648
                    _ = v42
                }
                _ = fi39
                _ = v40
            }
        }
        _ = len(v2)
        var v43 bool = false
        _ = fi34
        _ = v43
    }
    if (!(-1 != 1)) {
        v44 := GenS2{F0: 65535}
        v4 = append(v4, 2147483647)
        if (59 < v1) {
            var v45 int = -38
            if true {
                if false {
                    fmt.Println(v25, v6, v3)
                }
                v2 = append(v2, v3)
                if true {
                    var v46 uint64 = 461
                    _ = v46
                }
            } else {
                var v47 uint8 = v1
                v48 := GenS6{F0: v1}
                _ = v47
                _ = v48
            }
            v49 := GenS6{F0: v1}
            _ = v45
            _ = v49
        }
        _ = v44
    }
    fmt.Println(v1)
    v28 = int(v6)
    for fi50 := 0; fi50 < 6; fi50++ {
        switch fi50 {
        case -1:
            v1 = (^(v1))
        case -2147483648:
            if false {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                for fi51 := 0; fi51 < 19; fi51++ {
                    v1 += v1
                    fmt.Println(v6, fi50, fi50)
                    switch (fi50 | fi50) {
                    case -760:
                        switch ((v3 + "hello") + v3) {
                        case "a":
                            if (v5 < (^(v5))) {
                                if true {
                                    v4 = append(v4, 2147483648)
                                    var v52 uint16 = (^(v6))
                                    if (false || (381 > 42)) {
                                        if (-2147483648 == -240) {
                                            _ = len(v2)
                                            var v53 int64 = v5
                                            fmt.Println(v1)
                                            _ = v53
                                        }
                                        var v54 int = fi50
                                        v25 -= (v52 & 2)
                                        _ = v54
                                    }
                                    _ = v52
                                }
                                if len(v4) > 0 {
                                    v4[0] = 42
                                }
                                v55 := GenS7{F0: v28, F1: 65535, F2: int64(v1)}
                                _ = v55
                            }
                        case "hello":
                            fmt.Println(v1, v28)
                        case "foo bar":
                            switch genFn0(v28, 4294967295, 337) {
                            case -2147483648:
                                _ = len(v4)
                            case -9223372036854775808:
                                var v56 int16 = 0
                                _ = v56
                            default:
                                fmt.Println(v6)
                            }
                        default:
                            _ = len(v2)
                        }
                    case 9223372036854775807:
                        v6 = v25
                    case 42:
                        if false {
                            var v57 string = (v3 + v3)
                            v4 = append(v4, 2147483648)
                            _ = v57
                        } else {
                            switch v28 {
                            case 32768:
                                v6 += v25
                            default:
                                switch "\n\t" {
                                case "":
                                    v3 = "a"
                                case "hello":
                                    var v58 uint16 = (v6 - v6)
                                    _ = v58
                                case "a":
                                    if false {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch (fi51 % ((32767) | 1)) {
                                        case 128:
                                            var v59 int = (fi51 ^ 2147483647)
                                            _ = v59
                                        case 127:
                                            v60 := GenS8{F0: (fi50 % (((fi50 | (-(fi51)))) | 1)), F1: 784}
                                            _ = v60
                                        case 32768:
                                            v61 := GenS1{F0: -2147483648, F1: 0}
                                            _ = v61
                                        default:
                                            v62 := GenS4{F0: v25}
                                            _ = v62
                                        }
                                        v25 -= v25
                                    } else {
                                        v1 = v1
                                        v63 := GenS9{F0: v3, F1: v25, F2: false}
                                        _ = v63
                                    }
                                }
                            }
                        }
                    default:
                        var v64 uint32 = 2147483648
                        _ = v64
                    }
                    _ = fi51
                }
            } else {
                var v65 uint64 = uint64(v25)
                v6 += 677
                _ = v65
            }
        case 49:
            var v66 int8 = -25
            _ = v66
        default:
            var v67 float64 = 0.1
            _ = v67
        }
        _ = len(v2)
        if (!true) {
            switch genFn0(-461, 636, 367) {
            case -9223372036854775808:
                v3 = ((v3 + "\n\t") + v3)
            }
            v3 = (v3 + "a")
            v3 = (v3 + v3)
        } else {
            _ = len(v4)
        }
        _ = fi50
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
