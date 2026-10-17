package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 int16; F1 string; F2 int16 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int64; F1 int64; F2 uint }
func genFn0(p0 uint, p1 uint64) uint8 {
    for fi5 := 0; fi5 < 8; fi5++ {
        p1 = (p1 / ((p1) | 1))
        v6 := GenS1{F0: "foo bar", F1: p0}
        if (0.1 != (0.1 - (1e10 + (-2.25 * (float64((0.0 / 1.5)) + -2.25))))) {
            v7 := GenS2{F0: 0, F1: -1}
            fmt.Println(p1)
            _ = v7
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0)
    if (true && true) {
        var v8 int64 = -332
        _ = v8
    } else {
        v9 := GenS1{F0: ("hello" + ("" + "\n\t")), F1: p0}
        _ = v9
    }
    if true {
        var v10 uint = 251
        switch 2147483647 {
        case 4294967295:
            var v11 float32 = 0.0
            _ = v11
        }
        v12 := []int{9223372036854775807, 0, -912, 42}
        _ = v10
        _ = v12
    } else {
        v13 := []int{-1, -374, -746}
        p1 += (p1 * (p1 - uint64(p0)))
        _ = v13
    }
    v14 := GenS1{F0: (("foo bar" + "\n\t") + "foo bar"), F1: (p0 | p0)}
    _ = p0
    _ = p1
    _ = v14
    return 119
}
func genFn1(p0 int8) uint8 {
    p0 = (p0 ^ -104)
    var v15 int64 = 761
    _ = p0
    _ = v15
    return 255
}

func main() {
    var v1 uint = 42
    var v2 uint32 = 42
    v3 := make(map[string]uint32)
    var v4 uint8 = 42
    if true {
        v1 -= v1
        if true {
            v16 := []int{}
            for fi17 := 0; fi17 < 5; fi17++ {
                fmt.Println(v4)
                for fi18 := 0; fi18 < 8; fi18++ {
                    if true {
                        if (0.0 >= 3.14159) {
                            if true {
                                v19 := GenS3{F0: "a", F1: 1}
                                _ = v19
                            } else {
                                v4 = v4
                            }
                        }
                        v4 -= 1
                    }
                    v2 -= v2
                    _ = fi18
                }
                v20 := v3["foo bar"]
                _ = fi17
                _ = v20
            }
            var v21 uint16 = 32767
            _ = v16
            _ = v21
        }
        if ((0.0 > float32(1.5)) && (20 > 32767)) {
            fmt.Println(v1, v4, v2)
            var v22 int8 = 0
            _ = v22
        } else {
            v4 = v4
            v1 -= v1
        }
    } else {
        v2 = v2
        var v23 string = "\n\t"
        _ = v23
    }
    fmt.Println(v4, v1)
    switch "foo bar" {
    case "\n\t":
        v24 := GenS1{F0: "a", F1: 18446744073709551615}
        _ = v24
    case "foo bar":
        var v25 int32 = -1
        _ = v25
    default:
        if ((32767 >= -32768) || ((0 < 1) || (3.14159 >= ((((float64(3.14159) + float64((0.1 / float32(((1e10 - 0.0) - (((float64(1e10) / (1.5 + ((1e10 - 1.5) + (3.14159 / 1e10)))) * 0.1) / 1.0)))))) - 3.14159) / 1.0) + 1e10)))) {
            for fi26 := 0; fi26 < 17; fi26++ {
                fmt.Println(v2, fi26)
                if true {
                    var v27 float64 = 3.14159
                    _ = v27
                } else {
                    v1 = (v1 >> ((902) & 63))
                    for fi28 := 0; fi28 < 13; fi28++ {
                        var v29 uint16 = 42
                        var v30 int8 = -1
                        switch (^(fi26)) {
                        case -855:
                            v29 -= v29
                        }
                        if true {
                            v2 -= 42
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v29, v2)
                            case "hello":
                                v31 := GenS1{F0: "a", F1: (v1 ^ 500)}
                                _ = v31
                            }
                            v30 += (v30 * v30)
                        }
                        if ((!(177 >= uint(v2))) && false) {
                            v32 := GenS4{F0: 32767, F1: "hello", F2: 952}
                            _ = v32
                        } else {
                            if (false && (-2.25 == 1e10)) {
                                switch fi28 {
                                case 2147483647:
                                    switch (fi28 % (((fi26 / ((fi26) | 1))) | 1)) {
                                    case 2147483647:
                                        var v33 uint = uint(v2)
                                        _ = v33
                                    case 32767:
                                        delete(v3, ("a" + "hello"))
                                    case 128:
                                        var v34 bool = (128 != (v4 << ((v1) & 63)))
                                        _ = v34
                                    default:
                                        v35 := GenS2{F0: 2, F1: 20}
                                        _ = v35
                                    }
                                case 197:
                                    fmt.Println(v30)
                                case 9223372036854775807:
                                    fmt.Println(fi28, v30)
                                default:
                                    v4 -= v4
                                    if false {
                                        v36 := GenS5{F0: v2}
                                        _ = v36
                                    } else {
                                        if true {
                                            v29 -= 65535
                                            v4 = uint8(v1)
                                            fmt.Println(v29, fi28)
                                        }
                                        switch fi26 {
                                        case 434:
                                            switch fi26 {
                                            case -9223372036854775808:
                                                var v37 bool = (v4 < (v4 - (v4 ^ genFn0((^(v1)), 2))))
                                                _ = v37
                                            }
                                        case -9223372036854775808:
                                            switch (fi28 % ((fi26) | 1)) {
                                            case 42:
                                                v2 += (v2 >> ((2) & 63))
                                                if (1 >= -1) {
                                                    if true {
                                                        var v38 int16 = 1000
                                                        var v39 uint64 = uint64(v2)
                                                        _ = v38
                                                        _ = v39
                                                    } else {
                                                        delete(v3, "")
                                                        switch "hello" {
                                                        case "hello":
                                                            v40 := []int{(fi28 % (((fi26 % ((fi26) | 1))) | 1)), (fi26 | (-(fi26))), (^(fi26)), (fi28 << ((0) & 63))}
                                                            _ = v40
                                                        case "a":
                                                            var v41 uint16 = (^(v29))
                                                            _ = v41
                                                        case "foo bar":
                                                            if ((v30 == (-(v30))) || ((v1 << ((v1) & 63)) <= 255)) {
                                                                var v42 int32 = 2147483647
                                                                _ = v42
                                                            } else {
                                                                fmt.Println(v1, fi28)
                                                                v4 = (^(v4))
                                                                v43 := GenS5{F0: v2}
                                                                _ = v43
                                                            }
                                                        }
                                                    }
                                                    if (true || false) {
                                                        v44 := GenS3{F0: "", F1: v1}
                                                        switch fi28 {
                                                        case 0:
                                                            var v45 int64 = 9223372036854775807
                                                            _ = v45
                                                        case -9223372036854775808:
                                                            _ = len(v3)
                                                        case 268:
                                                            v4 -= (v4 << ((v1) & 63))
                                                        }
                                                        _ = v44
                                                    } else {
                                                        if ((!true) && (2147483647 < 2147483647)) {
                                                            fmt.Println(fi26)
                                                            var v46 uint = 18446744073709551615
                                                            _ = v46
                                                        }
                                                        v4 -= v4
                                                    }
                                                }
                                            }
                                        }
                                        var v47 uint = (v1 | (v1 | 977))
                                        _ = v47
                                    }
                                    var v48 uint = (v1 * (v1 ^ v1))
                                    _ = v48
                                }
                                var v49 uint64 = 18446744073709551615
                                _ = v49
                            }
                            delete(v3, ("a" + "foo bar"))
                            var v50 int32 = 2147483647
                            _ = v50
                        }
                        if false {
                            if ((42 != 1) || false) {
                                switch (fi28 | (-(fi28))) {
                                case 32767:
                                    v51 := []int{}
                                    _ = v51
                                case 1:
                                    switch "\n\t" {
                                    case "a":
                                        var v52 bool = (uint(v2) >= (v1 % ((339) | 1)))
                                        _ = v52
                                    case "":
                                        v2 = (v2 - 2147483647)
                                    case "foo bar":
                                        var v53 uint8 = (v4 & (v4 * v4))
                                        _ = v53
                                    }
                                default:
                                    v54 := GenS6{F0: 0, F1: 42, F2: 331}
                                    _ = v54
                                }
                                v55 := GenS5{F0: (v2 << ((uint(v4)) & 63))}
                                _ = v55
                            }
                            v2 += 2147483648
                            var v56 float64 = 1e10
                            _ = v56
                        } else {
                            v57 := GenS6{F0: 9223372036854775807, F1: int64(v29), F2: uint(v4)}
                            _ = v57
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    v1 -= (v1 / ((uint(v4)) | 1))
                }
                _ = fi26
            }
            var v58 int32 = 2147483647
            _ = v58
        } else {
            v59 := GenS1{F0: ("foo bar" + ""), F1: v1}
            v2 += 259
            _ = v59
        }
    }
    if false {
        var v60 uint16 = 985
        _ = v60
    }
    for fi61 := 0; fi61 < 15; fi61++ {
        fmt.Println(v1, v4, v1)
        delete(v3, "")
        _ = fi61
    }
    fmt.Println(v1, v2, v1)
    v4 -= v4
    v2 = (v2 & (v2 ^ v2))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
