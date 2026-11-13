package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 uint32; F1 uint; F2 bool }
type GenS3 struct { F0 uint32; F1 int64; F2 uint32 }
type GenS4 struct { F0 int32 }
func genFn0(p0 string, p1 string, p2 uint) int64 {
    switch -200 {
    case -977:
        var v7 int64 = 652
        _ = v7
    case -1:
        var v8 int64 = 0
        _ = v8
    default:
        if false {
            var v9 int16 = 659
            _ = v9
        }
    }
    var v10 float64 = 1.5
    switch 32768 {
    case 9223372036854775807:
        switch 459 {
        case -9223372036854775808:
            p0 = p1
        case 4294967295:
            p2 = (p2 >> ((596) & 63))
        case 32768:
            p2 = 0
        }
    case 311:
        p0 = p0
    case -1:
        var v11 string = ("hello" + p1)
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 0
}
func genFn1(p0 uint64, p1 int) string {
    p0 += 194
    p0 -= 748
    var v12 uint16 = 32767
    v13 := GenS1{F0: (v12 - v12), F1: (0 < 161)}
    if true {
        if true {
            v14 := GenS1{F0: (v12 & (v12 & (v12 * (v12 | (v12 - 42))))), F1: ((p0 - p0) != (p0 | (p0 | p0)))}
            _ = v14
        }
        if (0 <= 67) {
            var v15 float32 = float32(0.0)
            v12 = (v12 * (^(v12)))
            switch "hello" {
            case "\n\t":
                v16 := []int{127, p1, (-(p1)), (p1 >> ((2) & 63))}
                _ = v16
            case "foo bar":
                p1 += p1
            case "a":
                v17 := []int{(p1 >> ((2) & 63))}
                _ = v17
            default:
                v18 := []int{(p1 % (((p1 / ((840) | 1))) | 1)), p1}
                _ = v18
            }
            _ = v15
        }
    } else {
        switch p1 {
        case -1:
            v19 := []int{p1, (p1 & (p1 / ((p1) | 1))), 4294967295}
            _ = v19
        case 32768:
            v12 -= v12
        }
        var v20 bool = (-1 >= -490)
        if true {
            v21 := GenS1{F0: v12, F1: (!true)}
            for fi22 := 0; fi22 < 2; fi22++ {
                for fi23 := 0; fi23 < 1; fi23++ {
                    var v24 float64 = 1e10
                    _ = fi23
                    _ = v24
                }
                _ = fi22
            }
            _ = v21
        } else {
            v25 := GenS2{F0: 2147483648, F1: uint(p0), F2: (v20 && (!(677 != 677)))}
            var v26 uint64 = 544
            var v27 uint64 = (p0 | 356)
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v12
    _ = v13
    return "foo bar"
}

func main() {
    var v1 float64 = -2.25
    var v2 int = -127
    var v3 uint16 = 21
    var v4 uint16 = (^(v3))
    var v5 uint32 = 65535
    var v6 string = ("" + "hello")
    v28 := GenS2{F0: v5, F1: 270, F2: true}
    var v29 uint8 = uint8(v3)
    v6 = ""
    if true {
        v30 := GenS1{F0: v4, F1: false}
        v31 := []int{(v2 + v2), v2, 128, v2}
        if ((false || false) && (!(float32(v1) < 1.5))) {
            var v32 float64 = float64((1.5 * 1e10))
            for fi33 := 0; fi33 < 8; fi33++ {
                if len(v31) > 0 {
                    v31[0] = (fi33 + (fi33 - 127))
                }
                var v34 uint16 = v4
                for fi35 := 0; fi35 < 20; fi35++ {
                    _ = len(v31)
                    fmt.Println(v4, v34)
                    if true {
                        fmt.Println(v4, v1, v5)
                        if len(v31) > 0 {
                            v31[0] = fi35
                        }
                    } else {
                        v36 := GenS3{F0: 927, F1: 9223372036854775807, F2: v5}
                        var v37 uint = 42
                        fmt.Println(v5, v4)
                        _ = v36
                        _ = v37
                    }
                    _ = fi35
                }
                _ = fi33
                _ = v34
            }
            _ = v32
        }
        _ = v30
        _ = v31
    } else {
        fmt.Println(v29)
    }
    v38 := GenS4{F0: 0}
    var v39 int32 = -196
    v40 := GenS4{F0: (v39 + v39)}
    switch (v2 | 32768) {
    case 42:
        var v41 uint = 255
        _ = v41
    case -2147483648:
        for fi42 := 0; fi42 < 3; fi42++ {
            var v43 bool = true
            v44 := GenS2{F0: (^(v5)), F1: 819, F2: v43}
            _ = fi42
            _ = v43
            _ = v44
        }
    case -872:
        var v45 uint32 = (v5 & v5)
        _ = v45
    }
    v39 -= 42
    var v46 uint16 = (v4 | (^(v4)))
    v47 := GenS1{F0: v3, F1: true}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
