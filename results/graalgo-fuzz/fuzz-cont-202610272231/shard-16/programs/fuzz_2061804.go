package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int64; F1 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int) uint64 {
    v6 := GenS1{F0: 9223372036854775807, F1: -108}
    fmt.Println(p0, p0)
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := []int{(-(p1)), (p1 * (-(p1))), (p1 ^ p1)}
        _ = fi7
        _ = v8
    }
    for fi9 := 0; fi9 < 2; fi9++ {
        for fi10 := 0; fi10 < 19; fi10++ {
            var v11 uint32 = 880
            v12 := []int{(p1 ^ fi10), fi9}
            if (true || false) {
                switch (fi10 % (((p1 * fi10)) | 1)) {
                case 32767:
                    fmt.Println(fi9, p1, fi9)
                }
                fmt.Println(fi9, fi9)
                if true {
                    v13 := GenS1{F0: 1, F1: 127}
                    _ = v13
                }
            } else {
                if (!true) {
                    var v14 int64 = -1
                    var v15 int8 = 79
                    p1 += p1
                    v14 = (v14 * (v14 & (v14 / ((9223372036854775807) | 1))))
                    _ = v14
                    _ = v15
                }
            }
            _ = fi10
            _ = v11
            _ = v12
        }
        v16 := GenS2{F0: 0, F1: 2}
        v17 := []int{}
        _ = fi9
        _ = v16
        _ = v17
    }
    _ = p0
    _ = p1
    _ = v6
    return p0
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 42
    var v3 string = ("\n\t" + "hello")
    var v4 int64 = 42
    var v5 uint32 = 4294967295
    v18 := GenS2{F0: v4, F1: 65535}
    if false {
        v1 = v3
        switch 0 {
        case 127:
            switch "foo bar" {
            case "":
                var v19 int = -9223372036854775808
                _ = v19
            }
        case 78:
            v2 -= (v2 / (((v2 * (v2 << ((351) & 63)))) | 1))
        case 42:
            var v20 int8 = -128
            _ = v20
        }
    } else {
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 bool = true
            var v23 int32 = 394
            switch (v3 + ((v1 + "") + v3)) {
            case "\n\t":
                fmt.Println(v22)
            case "foo bar":
                var v24 int8 = 42
                _ = v24
            case "hello":
                v25 := GenS3{F0: (-(fi21))}
                _ = v25
            }
            _ = fi21
            _ = v22
            _ = v23
        }
        switch "\n\t" {
        case "\n\t":
            v2 = 1
        case "foo bar":
            for fi26 := 0; fi26 < 4; fi26++ {
                v27 := GenS3{F0: fi26}
                _ = fi26
                _ = v27
            }
        }
    }
    v28 := GenS2{F0: (v4 & 1), F1: v2}
    v29 := GenS1{F0: v4, F1: 53}
    if (!(1 >= 18446744073709551615)) {
        for fi30 := 0; fi30 < 11; fi30++ {
            switch "foo bar" {
            case "foo bar":
                for fi31 := 0; fi31 < 5; fi31++ {
                    var v32 int = (fi30 * (-(fi30)))
                    var v33 bool = (fi30 == fi30)
                    var v34 float32 = 3.14159
                    _ = fi31
                    _ = v32
                    _ = v33
                    _ = v34
                }
            default:
                fmt.Println(v2)
            }
            if true {
                for fi35 := 0; fi35 < 11; fi35++ {
                    var v36 uint64 = 57
                    fmt.Println(fi30, v36, fi35)
                    var v37 bool = (2147483647 <= (fi35 / ((9223372036854775807) | 1)))
                    _ = fi35
                    _ = v36
                    _ = v37
                }
            }
            v38 := []int{}
            _ = fi30
            _ = v38
        }
        v3 = (((v3 + v3) + ("hello" + v3)) + ("foo bar" + v3))
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        switch fi39 {
        case 4294967295:
            if true {
                if false {
                    v40 := []int{}
                    var v41 int64 = v4
                    _ = v40
                    _ = v41
                }
                v42 := GenS1{F0: v4, F1: -14}
                v4 += v4
                _ = v42
            } else {
                for fi43 := 0; fi43 < 3; fi43++ {
                    v44 := GenS2{F0: v4, F1: v2}
                    _ = fi43
                    _ = v44
                }
                if (1e10 != -2.25) {
                    v45 := GenS1{F0: (v4 / (((v4 & (v4 ^ (v4 * int64(fi39))))) | 1)), F1: 0}
                    _ = v45
                }
            }
        default:
            fmt.Println(v3, v4, fi39)
        }
        _ = fi39
    }
    fmt.Println(v3, v2)
    fmt.Println(v3)
    if (!(0 >= 491)) {
        v46 := GenS1{F0: 1, F1: 123}
        fmt.Println(v5, v2, v1)
        for fi47 := 0; fi47 < 8; fi47++ {
            fmt.Println(v2, fi47, v5)
            v4 -= v4
            if ((true || (0.0 == (0.1 / 1e10))) || true) {
                v48 := GenS4{F0: (v5 >> ((255) & 63)), F1: 20}
                _ = v48
            }
            _ = fi47
        }
        _ = v46
    }
    if true {
        var v49 int8 = -128
        v50 := GenS3{F0: 127}
        fmt.Println(v4, v5)
        _ = v49
        _ = v50
    } else {
        v51 := []int{0, 42}
        _ = v51
    }
    if (!(1 > 0)) {
        switch -2147483648 {
        case 9223372036854775807:
            fmt.Println(v1)
        case 2147483647:
            if false {
                if (1 == -1) {
                    var v52 int64 = (v4 % ((9223372036854775807) | 1))
                    var v53 int64 = int64(v2)
                    _ = v52
                    _ = v53
                } else {
                    if (!(v4 == v4)) {
                        if false {
                            var v54 int = -2147483648
                            for fi55 := 0; fi55 < 4; fi55++ {
                                v3 = (v1 + (v3 + (v1 + v3)))
                                var v56 bool = true
                                _ = fi55
                                _ = v56
                            }
                            _ = v54
                        }
                    } else {
                        v57 := []int{32767, 128}
                        _ = v57
                    }
                }
                v58 := GenS3{F0: -475}
                v3 = v3
                _ = v58
            }
        case 127:
            for fi59 := 0; fi59 < 1; fi59++ {
                var v60 bool = (-173 != 0)
                v1 = "hello"
                v1 = (v1 + v1)
                _ = fi59
                _ = v60
            }
        default:
            var v61 int64 = v4
            _ = v61
        }
        fmt.Println(v3, v4, v5)
        v62 := GenS3{F0: 2147483647}
        _ = v62
    } else {
        var v63 float32 = 1.5
        _ = v63
    }
    v64 := []int{-24, 1, -2147483648}
    v65 := GenS5{F0: v2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v64
    _ = v65
}
