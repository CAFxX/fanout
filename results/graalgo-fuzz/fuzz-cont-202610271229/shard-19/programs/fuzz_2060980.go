package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint64; F1 uint }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 uint32; F1 int; F2 uint64 }
type GenS8 struct { F0 bool; F1 int32; F2 int32 }
func genFn0(p0 uint8, p1 int8) int8 {
    for fi6 := 0; fi6 < 3; fi6++ {
        var v7 uint = 719
        var v8 float64 = 3.14159
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := []int{32767, 974, 869, -9223372036854775808}
    v10 := GenS1{F0: true, F1: (p1 << ((278) & 63))}
    for fi11 := 0; fi11 < 9; fi11++ {
        fmt.Println(p0, p0)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return p1
}
func genFn1(p0 uint64) uint {
    if false {
        v12 := GenS2{F0: "foo bar"}
        switch ((("\n\t" + ("foo bar" + ("\n\t" + "\n\t"))) + ("" + "a")) + "\n\t") {
        case "\n\t":
            if true {
                v13 := GenS3{F0: 9}
                if true {
                    var v14 uint64 = p0
                    var v15 uint8 = 208
                    var v16 int8 = 51
                    _ = v14
                    _ = v15
                    _ = v16
                }
                switch -49 {
                case 810:
                    var v17 uint32 = 4294967295
                    _ = v17
                case 32768:
                    for fi18 := 0; fi18 < 12; fi18++ {
                        fmt.Println(p0)
                        _ = fi18
                    }
                case 127:
                    var v19 uint32 = 2147483648
                    _ = v19
                }
                _ = v13
            } else {
                v20 := GenS4{F0: 4294967295, F1: 0}
                _ = v20
            }
        case "foo bar":
            v21 := []int{-268, 127, 4294967295, 0}
            _ = v21
        case "hello":
            v22 := GenS5{F0: 1, F1: 18446744073709551615}
            _ = v22
        default:
            v23 := GenS4{F0: -337, F1: -262}
            _ = v23
        }
        var v24 uint8 = 127
        _ = v12
        _ = v24
    } else {
        p0 = 1
        for fi25 := 0; fi25 < 13; fi25++ {
            v26 := []int{fi25}
            _ = fi25
            _ = v26
        }
        p0 = p0
    }
    if true {
        for fi27 := 0; fi27 < 2; fi27++ {
            v28 := GenS6{F0: -628}
            p0 = p0
            fmt.Println(fi27, p0)
            _ = fi27
            _ = v28
        }
    } else {
        v29 := []int{}
        var v30 int16 = -1
        p0 = (^(p0))
        _ = v29
        _ = v30
    }
    _ = p0
    return 657
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = 1
    var v3 float64 = -2.25
    var v4 uint = 2
    var v5 uint16 = 961
    for fi31 := 0; fi31 < 6; fi31++ {
        var v32 uint = uint(v5)
        v4 = v32
        v5 = (v5 % (((v5 | 65535)) | 1))
        _ = fi31
        _ = v32
    }
    if (true || false) {
        v2 -= (v2 ^ (v2 ^ int32(v5)))
        switch "a" {
        case "":
            var v33 float64 = v3
            _ = v33
        case "\n\t":
            if (!(!false)) {
                var v34 bool = true
                _ = v34
            }
        default:
            fmt.Println(v5)
        }
    }
    var v35 float64 = (v3 / -2.25)
    _ = len(v1)
    v36 := GenS7{F0: 0, F1: 32768, F2: 896}
    var v37 uint = (v4 + genFn1(2))
    fmt.Println(v2)
    var v38 int64 = 1
    v39 := []int{127}
    var v40 uint = 255
    if len(v39) >= 2 {
        v39 = v39[1:len(v39)]
    }
    var v41 uint64 = 533
    for fi42 := 0; fi42 < 20; fi42++ {
        fmt.Println(v38)
        _ = fi42
    }
    v43 := GenS8{F0: (1.5 == 3.14159), F1: (v2 << (((v37 / ((0) | 1))) & 63)), F2: v2}
    for fi44 := 0; fi44 < 11; fi44++ {
        if (!(!(v2 > 553))) {
            v45 := GenS4{F0: fi44, F1: 2147483647}
            _ = v45
        } else {
            fmt.Println(v40, v3)
            fmt.Println(v4, v5, v35)
            fmt.Println(v35)
        }
        _ = fi44
    }
    v3 += v35
    var v46 int16 = -801
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v43
    _ = v46
    _ = v5
}
