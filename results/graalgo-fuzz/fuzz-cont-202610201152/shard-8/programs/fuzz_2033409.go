package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 string; F2 uint16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int64) int {
    v6 := GenS1{F0: 613}
    if (-9223372036854775808 >= (p0 / (((p0 & (p0 / ((-9223372036854775808) | 1)))) | 1))) {
        v7 := GenS1{F0: 2}
        switch "a" {
        case "\n\t":
            switch 0 {
            case -2147483648:
                v8 := GenS1{F0: 1}
                _ = v8
            }
        case "foo bar":
            var v9 bool = false
            _ = v9
        case "hello":
            v10 := []int{2147483647}
            _ = v10
        default:
            fmt.Println(p0)
        }
        v11 := []int{4294967295, -890, 32767, 32768}
        _ = v11
        _ = v7
    } else {
        fmt.Println(p0)
        for fi12 := 0; fi12 < 17; fi12++ {
            p0 += p0
            v13 := GenS2{F0: 128}
            var v14 int = (fi12 + (fi12 ^ -9223372036854775808))
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    _ = p0
    _ = v6
    return -21
}
func genFn1(p0 uint16) uint16 {
    p0 = 2
    var v15 bool = (!true)
    if v15 {
        v16 := []int{42, 32768, 84, 2147483647}
        _ = v16
    }
    _ = p0
    _ = v15
    return p0
}

func main() {
    var v1 bool = false
    v2 := make(map[int]string)
    v3 := []uint32{42, 42, 90}
    var v4 uint16 = 255
    var v5 string = (((("a" + "a") + "foo bar") + "foo bar") + ("hello" + ((("\n\t" + ("hello" + ("a" + "foo bar"))) + (("\n\t" + "foo bar") + "\n\t")) + "hello")))
    var v17 int8 = 100
    switch 127 {
    case -1:
        var v18 int32 = -1
        _ = v18
    case 128:
        if (-618 >= -1) {
            v17 -= (v17 - (v17 << ((112) & 63)))
            v4 -= (^(v4))
        }
    }
    fmt.Println(v17)
    v19 := GenS3{F0: 0, F1: v5, F2: (v4 >> ((uint(v4)) & 63))}
    for fi20 := 0; fi20 < 4; fi20++ {
        v3 = append(v3, 755)
        v3 = append(v3, 736)
        for fi21 := 0; fi21 < 18; fi21++ {
            v1 = v1
            _ = fi21
        }
        _ = fi20
    }
    if v1 {
        fmt.Println(v17, v4, v17)
        var v22 int64 = -348
        _ = v22
    } else {
        v17 = 100
        for fi23 := 0; fi23 < 20; fi23++ {
            var v24 uint16 = v4
            v5 = ("foo bar" + v5)
            if v1 {
                v17 -= v17
            }
            _ = fi23
            _ = v24
        }
    }
    var v25 int = 9223372036854775807
    var v26 int64 = 9223372036854775807
    var v27 uint = 2
    switch v5 {
    case "hello":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        switch v5 {
        case "\n\t":
            if (!v1) {
                v1 = false
                for fi28 := 0; fi28 < 6; fi28++ {
                    var v29 uint8 = 127
                    v2[(v25 << ((v27) & 63))] = v5
                    if (false || (2147483647 > 1)) {
                        v4 += genFn1((v4 & (v4 / ((v4) | 1))))
                    }
                    _ = fi28
                    _ = v29
                }
            } else {
                if len(v3) > 0 {
                    v3[0] = 605
                }
                var v30 int32 = 0
                _ = v30
            }
        case "foo bar":
            fmt.Println(v25, v1, v1)
        case "a":
            fmt.Println(v26, v27, v26)
        }
    }
    v27 = (v27 ^ (v27 | v27))
    if ((v26 + (-(v26))) > (v26 * 831)) {
        for fi31 := 0; fi31 < 4; fi31++ {
            for fi32 := 0; fi32 < 9; fi32++ {
                if (!false) {
                    if (int(v4) <= 0) {
                        v26 = -9223372036854775808
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                        v3 = append(v3, 2)
                    } else {
                        fmt.Println(v4)
                        v5 = (v5 + v5)
                        fmt.Println(v17, v27)
                        v33 := GenS1{F0: 4294967295}
                        _ = v33
                    }
                }
                v3 = append(v3, 42)
                _ = fi32
            }
            var v34 int64 = (^(v26))
            var v35 int32 = -2147483648
            _ = fi31
            _ = v34
            _ = v35
        }
    } else {
        var v36 int = -499
        var v37 string = (v5 + v5)
        if ((v25 % ((v36) | 1)) <= v36) {
            fmt.Println(v1, v37, v25)
            fmt.Println(v27)
        } else {
            for fi38 := 0; fi38 < 13; fi38++ {
                switch v37 {
                case "\n\t":
                    if v1 {
                        v25 = 643
                        v25 = 42
                    } else {
                        var v39 uint64 = 2
                        fmt.Println(v5, v5)
                        if false {
                            var v40 int32 = -840
                            if (1.5 != -2.25) {
                                v41 := GenS4{F0: (v4 ^ (v4 & genFn1((v4 >> (((v27 / ((v27) | 1))) & 63))))), F1: v5}
                                var v42 uint = 42
                                _ = v41
                                _ = v42
                            } else {
                                v43 := GenS5{F0: 42}
                                _ = v43
                            }
                            v3 = append(v3, 622)
                            _ = v40
                        } else {
                            var v44 int64 = v26
                            _ = v44
                        }
                        _ = v39
                    }
                case "":
                    switch ("a" + (v37 + "a")) {
                    case "\n\t":
                        v45 := GenS1{F0: (v27 - v27)}
                        _ = v45
                    case "hello":
                        fmt.Println(v1, v36, v5)
                    case "foo bar":
                        for fi46 := 0; fi46 < 20; fi46++ {
                            var v47 uint64 = 206
                            v47 -= v47
                            _ = fi46
                            _ = v47
                        }
                    default:
                        v25 += -403
                    }
                case "hello":
                    v25 = 759
                default:
                    v5 = ("\n\t" + v37)
                }
                v26 += 690
                v25 += (v25 / (((v25 >> (((v27 ^ (v27 ^ (v27 - 255)))) & 63))) | 1))
                _ = fi38
            }
        }
        _ = v36
        _ = v37
    }
    if (false && v1) {
        if (232 != 255) {
            v17 -= v17
            v17 = (-(v17))
            var v48 uint32 = 4294967295
            _ = v48
        }
        v25 = v25
        v49 := GenS6{F0: uint32(v27), F1: 128}
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
