package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint64; F1 int }
type GenS5 struct { F0 int64; F1 uint64 }
func genFn0(p0 int) uint32 {
    v6 := []int{127, -9223372036854775808, p0}
    for fi7 := 0; fi7 < 11; fi7++ {
        switch "" {
        case "\n\t":
            p0 -= (^(fi7))
        default:
            p0 = fi7
        }
        for fi8 := 0; fi8 < 16; fi8++ {
            switch ("" + "") {
            case "hello":
                if false {
                    var v9 float32 = (0.0 + ((((-2.25 * (-2.25 * 3.14159)) + (3.14159 + 1.5)) - 1e10) * 1.5))
                    var v10 int8 = -128
                    _ = len(v6)
                    _ = v10
                    _ = v9
                } else {
                    fmt.Println(fi8, fi7)
                    if true {
                        var v11 uint = 244
                        v12 := GenS1{F0: true}
                        v11 = v11
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        var v13 float64 = (0.1 + 0.0)
                        fmt.Println(fi8, p0)
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                }
            case "a":
                var v14 float32 = 0.1
                _ = v14
            case "\n\t":
                if len(v6) > 0 {
                    v6[0] = -1
                }
            default:
                v15 := GenS2{F0: 408}
                _ = v15
            }
            _ = fi8
        }
        if (100 >= -128) {
            var v16 int64 = -9223372036854775808
            _ = v16
        } else {
            v17 := GenS1{F0: (!false)}
            var v18 int = p0
            for fi19 := 0; fi19 < 5; fi19++ {
                p0 -= -1
                var v20 uint32 = 2
                fmt.Println(fi7, fi7)
                _ = fi19
                _ = v20
            }
            _ = v17
            _ = v18
        }
        _ = fi7
    }
    v6 = append(v6, (p0 % ((p0) | 1)))
    switch p0 {
    case 92:
        var v21 bool = true
        _ = v21
    case -183:
        if true {
            v6 = append(v6, (p0 & (p0 / ((127) | 1))))
            _ = len(v6)
            p0 = 42
        }
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int64 = 42
    v2 := make(map[uint32]bool)
    var v3 string = ("foo bar" + "\n\t")
    var v4 string = (v3 + (v3 + v3))
    var v5 float64 = 1.5
    v4 = v4
    fmt.Println(v4, v3)
    if false {
        fmt.Println(v5)
        switch (v4 + "a") {
        case "hello":
            v22 := []int{128, -2147483648, 964, 9223372036854775807}
            _ = v22
        case "foo bar":
            var v23 int = 803
            _ = v23
        case "":
            v24 := GenS1{F0: (!(1 < -623))}
            _ = v24
        default:
            fmt.Println(v4)
        }
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            switch ("\n\t" + v4) {
            case "hello":
                switch (v4 + v4) {
                case "hello":
                    v26 := GenS1{F0: (1 <= 1)}
                    _ = v26
                case "foo bar":
                    v3 = ("hello" + v3)
                case "\n\t":
                    switch fi25 {
                    case 2147483647:
                        v1 = v1
                    case 735:
                        v27 := GenS2{F0: int32(fi25)}
                        _ = v27
                    case 42:
                        if (2147483647 >= 488) {
                            switch v3 {
                            case "":
                                var v28 bool = (!false)
                                _ = v28
                            default:
                                v29 := GenS2{F0: int32(fi25)}
                                _ = v29
                            }
                        } else {
                            var v30 float32 = ((0.1 * (1e10 / 1.0)) + 1.5)
                            _ = v30
                        }
                    default:
                        v31 := []int{(fi25 % ((352) | 1)), 32767}
                        _ = v31
                    }
                default:
                    v32 := GenS3{F0: 1}
                    _ = v32
                }
            }
            for fi33 := 0; fi33 < 17; fi33++ {
                v34 := GenS3{F0: -65}
                v5 = v5
                v4 = "a"
                _ = fi33
                _ = v34
            }
            _ = fi25
        }
        delete(v2, genFn0(32767))
    }
    v35 := GenS4{F0: 732, F1: -202}
    for fi36 := 0; fi36 < 18; fi36++ {
        v4 = (v3 + (v4 + v3))
        if false {
            var v37 uint32 = genFn0((fi36 >> ((0) & 63)))
            fmt.Println(v4)
            _ = v37
        } else {
            switch fi36 {
            case 0:
                if true {
                    switch ((v3 + "hello") + v3) {
                    case "":
                        var v38 uint8 = 178
                        _ = v38
                    case "a":
                        for fi39 := 0; fi39 < 4; fi39++ {
                            v40 := GenS5{F0: 134, F1: 18446744073709551615}
                            v1 = v1
                            _ = fi39
                            _ = v40
                        }
                    case "\n\t":
                        fmt.Println(fi36, v1)
                    }
                    fmt.Println(v3, v1, v4)
                } else {
                    v4 = "hello"
                }
            default:
                v3 = (v3 + "hello")
            }
        }
        var v41 string = (v4 + "\n\t")
        _ = fi36
        _ = v41
    }
    var v42 int32 = 812
    v43 := GenS2{F0: v42}
    var v44 int8 = -111
    v1 = v1
    for fi45 := 0; fi45 < 3; fi45++ {
        if false {
            var v46 float64 = 1e10
            _ = v46
        }
        var v47 int64 = (v1 % ((int64(fi45)) | 1))
        if (genFn0(-725) == 2147483648) {
            for fi48 := 0; fi48 < 5; fi48++ {
                v49 := []int{fi45, fi48, 2147483647, fi45}
                _ = fi48
                _ = v49
            }
            var v50 int64 = v47
            for fi51 := 0; fi51 < 14; fi51++ {
                fmt.Println(v42)
                v3 = ("hello" + ("" + v4))
                _ = fi51
            }
            _ = v50
        }
        _ = fi45
        _ = v47
    }
    if (4294967295 > 662) {
        v42 += v42
    }
    v1 = (v1 & 368)
    v5 = v5
    var v52 uint64 = 718
    for fi53 := 0; fi53 < 1; fi53++ {
        var v54 uint = 4294967295
        _ = fi53
        _ = v54
    }
    for fi55 := 0; fi55 < 13; fi55++ {
        var v56 bool = true
        fmt.Println(fi55, v4, v3)
        _ = fi55
        _ = v56
    }
    v57 := []int{1}
    v3 = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v57:", len(v57), v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v5
    _ = v52
    _ = v57
}
