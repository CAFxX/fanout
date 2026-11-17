package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int16; F1 int32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int8; F1 bool; F2 uint16 }
func genFn0(p0 uint8) int16 {
    var v5 uint16 = 0
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = 0
        _ = fi6
        _ = v7
    }
    v5 += (^(v5))
    switch 1 {
    case 1:
        fmt.Println(v5)
    case 9223372036854775807:
        v5 = 0
    case -2147483648:
        v8 := []int{}
        _ = v8
    }
    _ = p0
    _ = v5
    return 237
}
func genFn1(p0 uint, p1 int16, p2 uint8) bool {
    var v9 bool = false
    v10 := []int{}
    for fi11 := 0; fi11 < 17; fi11++ {
        for fi12 := 0; fi12 < 4; fi12++ {
            if false {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                var v13 uint16 = 897
                _ = v13
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                p1 = (p1 + (p1 / ((p1) | 1)))
            }
            fmt.Println(fi11, p1, p2)
            _ = fi12
        }
        _ = fi11
    }
    _ = len(v10)
    for fi14 := 0; fi14 < 13; fi14++ {
        if (p2 > p2) {
            if len(v10) > 0 {
                v10[0] = fi14
            }
            var v15 float32 = (3.14159 * float32(0.0))
            if ((p2 >> ((p0) & 63)) == (p2 << ((p0) & 63))) {
                v15 += float32(0.0)
                p1 -= 332
                v15 += float32(((0.0 - -2.25) - 0.1))
            } else {
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
            }
            _ = v15
        }
        switch ("" + ((("hello" + "a") + "hello") + "a")) {
        case "\n\t":
            var v16 uint16 = 888
            _ = v16
        case "a":
            if v9 {
                for fi17 := 0; fi17 < 13; fi17++ {
                    switch "" {
                    case "foo bar":
                        p2 = uint8(p0)
                    case "\n\t":
                        if v9 {
                            v9 = v9
                        }
                    default:
                        v10 = append(v10, fi14)
                    }
                    switch fi17 {
                    case 4294967295:
                        v18 := GenS1{F0: 2147483647}
                        _ = v18
                    }
                    v19 := GenS2{F0: 31}
                    _ = fi17
                    _ = v19
                }
            }
        case "":
            _ = len(v10)
        default:
            v9 = (((2 >= 666) && v9) || (v9 && (true && (false && ((2 > 127) && v9)))))
        }
        _ = fi14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return v9
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    v3 := []uint32{0}
    var v4 int64 = 1
    v20 := GenS3{F0: 4294967295}
    switch 0 {
    case 792:
        v21 := GenS1{F0: 42}
        _ = v21
    case 32768:
        var v22 float64 = -2.25
        _ = v22
    default:
        for fi23 := 0; fi23 < 20; fi23++ {
            v2 = true
            if (32767 >= fi23) {
                switch "hello" {
                case "\n\t":
                    v3 = append(v3, 703)
                case "":
                    fmt.Println(v2)
                case "foo bar":
                    v2 = false
                default:
                    v24 := GenS2{F0: 886}
                    _ = v24
                }
                v4 += v4
                v25 := GenS1{F0: -81}
                _ = v25
            }
            _ = fi23
        }
    }
    fmt.Println(v2, v2)
    v4 -= 42
    var v26 bool = v2
    switch 1 {
    case -559:
        v4 = v4
    case 32768:
        for fi27 := 0; fi27 < 17; fi27++ {
            fmt.Println(v26, fi27, v2)
            switch "" {
            case "a":
                if (v2 || (2 <= 364)) {
                    fmt.Println(v4)
                }
            case "\n\t":
                if v2 {
                    fmt.Println(fi27, v4, fi27)
                    var v28 string = "\n\t"
                    _ = v28
                } else {
                    v29 := GenS4{F0: 863, F1: 0}
                    v26 = (false && v2)
                    v2 = false
                    _ = v29
                }
            }
            switch fi27 {
            case 2147483647:
                v4 -= 9223372036854775807
            case -357:
                if len(v3) > 0 {
                    v3[0] = 2147483648
                }
            default:
                var v30 uint32 = 0
                _ = v30
            }
            _ = fi27
        }
    case 2147483647:
        v1[4294967295] = (1 == 0)
    default:
        var v31 int16 = 739
        _ = v31
    }
    if v26 {
        for fi32 := 0; fi32 < 18; fi32++ {
            switch "a" {
            case "\n\t":
                if ((18446744073709551615 <= 2) && v2) {
                    v33 := GenS3{F0: -123}
                    _ = v33
                }
            default:
                var v34 uint16 = 255
                _ = v34
            }
            _ = fi32
        }
        for fi35 := 0; fi35 < 2; fi35++ {
            _ = len(v3)
            v36 := GenS1{F0: 0}
            _ = fi35
            _ = v36
        }
        for fi37 := 0; fi37 < 4; fi37++ {
            v26 = (127 != 83)
            _ = fi37
        }
    }
    fmt.Println(v2, v26)
    if true {
        for fi38 := 0; fi38 < 8; fi38++ {
            v39 := GenS5{F0: false}
            v26 = (v4 > (v4 / (((v4 / ((v4) | 1))) | 1)))
            v40 := GenS6{F0: 85, F1: (2147483647 >= -1), F2: 1}
            _ = fi38
            _ = v39
            _ = v40
        }
        var v41 int = int(v4)
        _ = v41
    } else {
        v42 := GenS5{F0: (!v2)}
        delete(v1, 2147483647)
        _ = v42
    }
    var v43 uint = 549
    v44 := GenS2{F0: -2147483648}
    var v45 int16 = 0
    var v46 uint8 = 233
    if v2 {
        v1[2147483648] = true
    }
    fmt.Println(v43, v45)
    if ("hello" > "") {
        _ = len(v3)
    } else {
        if len(v3) > 0 {
            v3[0] = 2147483648
        }
        fmt.Println(v4)
    }
    var v47 uint8 = 128
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
}
