package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 uint32; F1 int16; F2 int }
type GenS3 struct { F0 string; F1 uint64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 int }
func genFn0(p0 bool, p1 uint64, p2 string) string {
    switch 32768 {
    case 32767:
        var v6 bool = true
        _ = v6
    case 0:
        var v7 string = p2
        _ = v7
    case 127:
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint32 = 4294967295
            var v10 int16 = -1
            fmt.Println(fi8, v9)
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 uint8 = 63
        _ = v11
    }
    var v12 float64 = 1e10
    if p0 {
        v13 := GenS1{F0: 158, F1: 856, F2: p2}
        _ = v13
    } else {
        for fi14 := 0; fi14 < 7; fi14++ {
            if true {
                switch (((p2 + p2) + (p2 + ("a" + ((p2 + "a") + p2)))) + ((((p2 + p2) + "a") + "foo bar") + p2)) {
                case "a":
                    p2 = p2
                case "hello":
                    switch "\n\t" {
                    case "hello":
                        switch (p2 + ("" + "foo bar")) {
                        case "hello":
                            v15 := []int{(fi14 % ((fi14) | 1))}
                            _ = v15
                        case "\n\t":
                            v16 := GenS1{F0: 128, F1: 0, F2: (p2 + p2)}
                            _ = v16
                        case "":
                            var v17 int32 = -1
                            _ = v17
                        }
                    }
                case "\n\t":
                    switch p2 {
                    case "\n\t":
                        var v18 int16 = 32767
                        _ = v18
                    case "hello":
                        if (100 < 100) {
                            switch "a" {
                            case "a":
                                v19 := []int{fi14, 128, fi14}
                                _ = v19
                            case "":
                                switch -2147483648 {
                                case 2147483647:
                                    switch (fi14 | (-(fi14))) {
                                    case 4294967295:
                                        var v20 uint = 2
                                        _ = v20
                                    case -9223372036854775808:
                                        var v21 uint = 19
                                        _ = v21
                                    case 127:
                                        switch ("foo bar" + "foo bar") {
                                        case "":
                                            v22 := []int{(^(fi14)), -67, fi14}
                                            _ = v22
                                        default:
                                            fmt.Println(p0)
                                        }
                                    default:
                                        v23 := []int{(-(fi14))}
                                        _ = v23
                                    }
                                default:
                                    for fi24 := 0; fi24 < 8; fi24++ {
                                        p2 = ("foo bar" + "hello")
                                        var v25 float32 = float32(v12)
                                        p2 = p2
                                        _ = fi24
                                        _ = v25
                                    }
                                }
                            }
                            var v26 int64 = 9223372036854775807
                            _ = v26
                        } else {
                            p2 = (p2 + (p2 + ""))
                        }
                    default:
                        var v27 int = -749
                        _ = v27
                    }
                }
                v28 := []int{fi14, fi14}
                _ = v28
            } else {
                p2 = "foo bar"
                var v29 int64 = 42
                p2 = (((p2 + "foo bar") + p2) + p2)
                _ = v29
            }
            var v30 float64 = v12
            _ = fi14
            _ = v30
        }
        var v31 float32 = (1.5 * float32((v12 / 1.5)))
        _ = v31
    }
    switch -396 {
    case -9223372036854775808:
        v32 := GenS2{F0: 65535, F1: -520, F2: 0}
        _ = v32
    case 1:
        p2 = p2
    case 9223372036854775807:
        var v33 int16 = 1
        _ = v33
    default:
        v34 := []int{664}
        _ = v34
    }
    v35 := []int{9223372036854775807, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v35
    return (p2 + "foo bar")
}
func genFn1() bool {
    switch 128 {
    case 989:
        switch -2147483648 {
        case 935:
            switch -357 {
            case 936:
                var v36 string = "\n\t"
                _ = v36
            default:
                var v37 uint = 1
                _ = v37
            }
        case 764:
            v38 := GenS1{F0: 142, F1: 2147483647, F2: ("" + "")}
            _ = v38
        default:
            var v39 int16 = -1
            _ = v39
        }
    case 88:
        var v40 float32 = -2.25
        _ = v40
    default:
        for fi41 := 0; fi41 < 20; fi41++ {
            for fi42 := 0; fi42 < 11; fi42++ {
                var v43 uint64 = 553
                v43 = v43
                _ = fi42
                _ = v43
            }
            var v44 string = "\n\t"
            if true {
                v44 = (v44 + "hello")
            } else {
                var v45 int = 127
                _ = v45
            }
            _ = fi41
            _ = v44
        }
    }
    if false {
        var v46 uint8 = 128
        switch 32768 {
        case -528:
            if false {
                v47 := []int{}
                v46 += 0
                _ = v47
            } else {
                var v48 uint16 = 627
                _ = v48
            }
        }
        v49 := GenS3{F0: "hello", F1: 18446744073709551615}
        _ = v46
        _ = v49
    }
    return false
}

func main() {
    var v1 int32 = -2147483648
    var v2 int32 = v1
    var v3 bool = false
    v4 := []int{-823, -2147483648}
    var v5 uint32 = 42
    v50 := GenS2{F0: (v5 ^ v5), F1: 1, F2: 2147483647}
    if (-2147483648 <= v1) {
        fmt.Println(v1, v2, v3)
    } else {
        v2 -= (v1 + v1)
    }
    _ = len(v4)
    var v51 float64 = float64(0.0)
    var v52 uint64 = 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v53 := GenS3{F0: "\n\t", F1: v52}
    v54 := GenS4{F0: 1}
    if v3 {
        v52 = v52
    }
    for fi55 := 0; fi55 < 1; fi55++ {
        var v56 float64 = 0.0
        v5 -= v5
        v5 = (v5 << ((765) & 63))
        _ = fi55
        _ = v56
    }
    switch "hello" {
    case "foo bar":
        var v57 uint64 = (v52 % ((v52) | 1))
        _ = v57
    case "":
        var v58 int32 = v1
        _ = v58
    case "a":
        for fi59 := 0; fi59 < 11; fi59++ {
            v60 := GenS5{F0: int32(fi59), F1: (fi59 & -9223372036854775808)}
            var v61 int = (-(fi59))
            _ = fi59
            _ = v60
            _ = v61
        }
    default:
        switch 32768 {
        case 4294967295:
            var v62 string = "foo bar"
            _ = v62
        case 1:
            v63 := GenS5{F0: 656, F1: -9223372036854775808}
            _ = v63
        }
    }
    v64 := GenS2{F0: v5, F1: -32768, F2: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v54
    _ = v64
}
