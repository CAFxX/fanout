package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int32; F1 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint32; F1 uint16; F2 bool }
func genFn0(p0 uint64, p1 uint8) bool {
    if (0 >= -2147483648) {
        var v7 int32 = -661
        for fi8 := 0; fi8 < 5; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                v10 := GenS1{F0: 0}
                fmt.Println(v7, fi8)
                _ = fi9
                _ = v10
            }
            var v11 int = -9223372036854775808
            _ = fi8
            _ = v11
        }
        _ = v7
    }
    v12 := []int{-1, 1}
    _ = p0
    _ = p1
    _ = v12
    return (!false)
}
func genFn1(p0 int8, p1 bool, p2 uint64) bool {
    if true {
        p0 -= (p0 | (-(p0)))
    } else {
        p1 = true
        for fi13 := 0; fi13 < 8; fi13++ {
            var v14 int64 = 42
            if (190 == 127) {
                v15 := GenS1{F0: 1000}
                p1 = false
                _ = v15
            }
            for fi16 := 0; fi16 < 13; fi16++ {
                switch (^(fi13)) {
                case 455:
                    p2 -= p2
                }
                fmt.Println(fi16, fi16, v14)
                v17 := []int{fi16, fi13}
                _ = fi16
                _ = v17
            }
            _ = fi13
            _ = v14
        }
        if false {
            var v18 uint16 = 1
            v18 = 70
            if (int(v18) > 4294967295) {
                fmt.Println(p2, p2, p0)
            }
            _ = v18
        } else {
            p0 += p0
            var v19 float64 = 0.0
            _ = v19
        }
    }
    v20 := []int{888, -749}
    v21 := GenS2{F0: 560, F1: 2}
    var v22 float64 = ((-2.25 / 0.1) * (float64(1.5) * (3.14159 * 0.0)))
    p0 += (p0 | 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    _ = v21
    _ = v22
    return (p2 >= (p2 >> ((170) & 63)))
}

func main() {
    var v1 int64 = 1
    var v2 uint8 = 128
    v3 := make(map[uint32]int)
    var v4 int = -538
    v5 := make(map[uint32]bool)
    v6 := make(map[string]uint32)
    var v23 bool = (2147483647 != 0)
    for fi24 := 0; fi24 < 13; fi24++ {
        v25 := GenS1{F0: 0}
        v1 += v1
        _ = fi24
        _ = v25
    }
    var v26 uint64 = 4294967295
    v2 = v2
    for fi27 := 0; fi27 < 18; fi27++ {
        switch 597 {
        case 128:
            if (int64(fi27) == v1) {
                var v28 float32 = 1.5
                var v29 float64 = 0.0
                var v30 int64 = v1
                _ = v28
                _ = v29
                _ = v30
            }
        case 42:
            v26 = 18446744073709551615
        default:
            for fi31 := 0; fi31 < 6; fi31++ {
                v32 := GenS3{F0: 2147483647}
                fmt.Println(fi31, fi31)
                v33 := []int{(-(v4)), 9223372036854775807, (v4 << ((0) & 63)), (-(fi27))}
                _ = len(v33)
                if v23 {
                    fmt.Println(v23)
                    switch -267 {
                    case -857:
                        switch "a" {
                        case "\n\t":
                            _ = len(v6)
                        case "":
                            delete(v5, 631)
                        }
                    }
                    var v34 uint16 = 42
                    _ = v34
                } else {
                    var v35 float64 = ((1e10 + 1.5) * (((float64(0.1) * (3.14159 / -2.25)) / float64(1.5)) - 1.5))
                    _ = v35
                }
                switch (v4 >> ((42) & 63)) {
                case 32768:
                    if true {
                        var v36 int32 = -2147483648
                        fmt.Println(v23, fi31, fi31)
                        _ = v36
                    }
                case 128:
                    var v37 string = ("hello" + "a")
                    _ = v37
                }
                _ = fi31
                _ = v32
                _ = v33
            }
        }
        _ = fi27
    }
    v1 = v1
    for fi38 := 0; fi38 < 20; fi38++ {
        if (v1 != (v1 | 1)) {
            var v39 float32 = -2.25
            switch "foo bar" {
            case "\n\t":
                var v40 int32 = 2147483647
                _ = v40
            case "foo bar":
                v41 := v6["\n\t"]
                _ = v41
            }
            _ = v39
        } else {
            var v42 uint8 = v2
            v43 := GenS2{F0: 0, F1: v26}
            v44 := []int{827, (-(fi38)), (fi38 >> ((uint(v26)) & 63))}
            _ = v42
            _ = v43
            _ = v44
        }
        switch fi38 {
        case -9223372036854775808:
            v45 := GenS1{F0: -736}
            _ = v45
        case -1:
            for fi46 := 0; fi46 < 10; fi46++ {
                v47 := GenS3{F0: -334}
                _ = fi46
                _ = v47
            }
        default:
            if (v26 <= (^(v26))) {
                if v23 {
                    var v48 float64 = 1e10
                    v48 = 3.14159
                    _ = len(v6)
                    _ = v48
                } else {
                    v26 = 459
                }
                v4 += v4
                v49 := GenS3{F0: 80}
                _ = v49
            }
        }
        v23 = (0.1 != float32((0.0 / 1.0)))
        _ = fi38
    }
    var v50 uint8 = (v2 / (((^(v2))) | 1))
    for fi51 := 0; fi51 < 11; fi51++ {
        if false {
            var v52 uint64 = (v26 * v26)
            var v53 uint64 = v52
            var v54 int64 = (v1 % ((v1) | 1))
            _ = v52
            _ = v53
            _ = v54
        }
        switch v4 {
        case 128:
            v4 = (^(fi51))
        }
        v23 = false
        _ = fi51
    }
    if false {
        for fi55 := 0; fi55 < 20; fi55++ {
            var v56 int = 554
            for fi57 := 0; fi57 < 19; fi57++ {
                v23 = v23
                _ = fi57
            }
            var v58 int = fi55
            _ = fi55
            _ = v56
            _ = v58
        }
        v59 := GenS3{F0: -2147483648}
        _ = v59
    }
    v1 -= v1
    if (-1 > 1) {
        switch (v4 ^ 32768) {
        case 880:
            v2 = v50
        case 19:
            switch (v4 % (((v4 % ((v4) | 1))) | 1)) {
            case -762:
                v60 := GenS4{F0: 0, F1: 2, F2: v23}
                _ = v60
            case -9223372036854775808:
                if ("hello" >= "") {
                    v26 = 0
                }
            case -1:
                if (v23 && v23) {
                    v61 := []int{v4, (^(v4)), 0, v4}
                    _ = v61
                }
            default:
                if (v26 != 2) {
                    if (!((("foo bar" + ("\n\t" + "foo bar")) + "hello") <= ("hello" + ""))) {
                        var v62 uint = 0
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 5; fi63++ {
                            fmt.Println(v1)
                            _ = fi63
                        }
                        v64 := GenS4{F0: 2, F1: 32767, F2: ((v26 <= 2) && v23)}
                        _ = v64
                    }
                } else {
                    v50 = v2
                    var v65 string = (("" + ("\n\t" + "a")) + (("\n\t" + ("\n\t" + "a")) + "hello"))
                    _ = v65
                }
            }
        default:
            delete(v3, 236)
        }
    } else {
        var v66 string = "a"
        _ = v66
    }
    v67 := []int{(v4 - -9223372036854775808), v4, 861}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v67:", len(v67), v67)
    _ = v1
    _ = v2
    _ = v23
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v6
    _ = v67
}
