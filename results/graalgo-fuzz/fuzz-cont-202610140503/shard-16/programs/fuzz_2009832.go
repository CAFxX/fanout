package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint16 }
type GenS3 struct { F0 uint8; F1 int; F2 int16 }
type GenS4 struct { F0 int64; F1 uint64 }
type GenS5 struct { F0 uint16; F1 int32 }
func genFn0(p0 int64) int {
    v5 := []int{}
    if (255 <= 701) {
        var v6 int64 = (p0 >> ((857) & 63))
        switch -217 {
        case 42:
            v6 = (p0 | (v6 % (((v6 / ((v6) | 1))) | 1)))
        case 958:
            var v7 string = ("a" + "")
            _ = v7
        case -9223372036854775808:
            p0 += v6
        }
        v6 += (v6 & 9223372036854775807)
        _ = v6
    } else {
        p0 = p0
    }
    _ = p0
    _ = v5
    return 0
}
func genFn1(p0 uint8, p1 int16) uint32 {
    var v8 int64 = 0
    switch -993 {
    case 763:
        p0 = p0
    case 913:
        v9 := []int{-2147483648, 1}
        _ = v9
    default:
        p1 = -702
    }
    _ = p0
    _ = p1
    _ = v8
    return 2147483648
}

func main() {
    var v1 int8 = -17
    var v2 uint16 = 255
    v3 := []int{-927, 1}
    var v4 uint8 = 2
    for fi10 := 0; fi10 < 18; fi10++ {
        if false {
            for fi11 := 0; fi11 < 16; fi11++ {
                if true {
                    var v12 int16 = 623
                    var v13 uint = 18446744073709551615
                    var v14 uint16 = v2
                    _ = v12
                    _ = v13
                    _ = v14
                }
                var v15 uint16 = 798
                _ = fi11
                _ = v15
            }
        } else {
            if (2147483647 == genFn1(v4, -1)) {
                for fi16 := 0; fi16 < 9; fi16++ {
                    switch "foo bar" {
                    case "a":
                        v2 = v2
                    case "":
                        v1 -= 0
                    }
                    v4 -= (v4 & (v4 + (^(v4))))
                    _ = fi16
                }
            }
            v4 = 2
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 float32 = 1.5
                if ((((fi10 | int(v2)) == (fi10 * (fi17 << ((255) & 63)))) || (-32768 < 256)) && true) {
                    if false {
                        if (!true) {
                            if true {
                                if false {
                                    v4 += v4
                                    if len(v3) > 0 {
                                        v3[0] = 4294967295
                                    }
                                    if (v2 <= v2) {
                                        var v19 string = ("hello" + "hello")
                                        var v20 string = "a"
                                        fmt.Println(v2, v18, fi10)
                                        _ = v19
                                        _ = v20
                                    } else {
                                        switch "\n\t" {
                                        case "":
                                            if false {
                                                if false {
                                                    v2 -= 944
                                                    v21 := GenS1{F0: fi17}
                                                    v22 := GenS1{F0: (fi17 ^ (-(fi17)))}
                                                    _ = v21
                                                    _ = v22
                                                }
                                            } else {
                                                v23 := GenS1{F0: 32767}
                                                _ = v23
                                            }
                                        }
                                    }
                                    switch "hello" {
                                    case "a":
                                        _ = len(v3)
                                    }
                                } else {
                                    switch "\n\t" {
                                    case "hello":
                                        var v24 int16 = 32767
                                        _ = v24
                                    case "a":
                                        if ((-2.25 + (v18 * (-2.25 * v18))) > v18) {
                                            var v25 uint = 65535
                                            _ = v25
                                        } else {
                                            v4 -= v4
                                            var v26 uint32 = 0
                                            switch "foo bar" {
                                            case "a":
                                                fmt.Println(v1, v4)
                                            default:
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            }
                                            fmt.Println(v26)
                                            _ = v26
                                        }
                                    case "foo bar":
                                        v27 := GenS1{F0: 0}
                                        _ = v27
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                if (false || (663 <= 324)) {
                                    v28 := GenS2{F0: 255, F1: 1}
                                    _ = v28
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        v2 = v2
                                        if true {
                                            _ = len(v3)
                                        }
                                        v2 += 2
                                    case "a":
                                        switch ("a" + "") {
                                        case "":
                                            switch "a" {
                                            case "a":
                                                v29 := GenS3{F0: (^(v4)), F1: fi10, F2: -893}
                                                _ = v29
                                            case "\n\t":
                                                var v30 string = "a"
                                                _ = v30
                                            case "":
                                                var v31 uint64 = 4294967295
                                                _ = v31
                                            default:
                                                v1 -= v1
                                            }
                                        case "a":
                                            if (-2.25 < (float32(-2.25) - v18)) {
                                                var v32 uint32 = 65535
                                                var v33 int8 = v1
                                                _ = v32
                                                _ = v33
                                            } else {
                                                if len(v3) > 0 {
                                                    v3[0] = fi10
                                                }
                                                _ = len(v3)
                                                var v34 int16 = int16(v2)
                                                _ = v34
                                            }
                                        }
                                    }
                                    _ = len(v3)
                                }
                                fmt.Println(v4, v1, fi10)
                            }
                            if false {
                                switch 128 {
                                case 32768:
                                    switch fi17 {
                                    case -9223372036854775808:
                                        var v35 uint = 42
                                        _ = v35
                                    case 128:
                                        switch (("foo bar" + "a") + "a") {
                                        case "foo bar":
                                            switch int(v4) {
                                            case 127:
                                                v1 = v1
                                            }
                                        default:
                                            var v36 int32 = 42
                                            _ = v36
                                        }
                                    case -737:
                                        switch ("hello" + "foo bar") {
                                        case "":
                                            if len(v3) > 0 {
                                                v3[0] = (fi10 / ((fi17) | 1))
                                            }
                                        case "\n\t":
                                            var v37 uint64 = 18446744073709551615
                                            _ = v37
                                        }
                                    }
                                case 1:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                case -490:
                                    v4 = (v4 + v4)
                                default:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                var v38 uint16 = (^(v2))
                                _ = v38
                            } else {
                                fmt.Println(v2, fi17)
                                v39 := GenS2{F0: (v4 & (^(v4))), F1: 42}
                                _ = v39
                            }
                            var v40 float32 = v18
                            _ = v40
                        }
                    } else {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v4 += 42
                    v18 = (v18 / (float32(1e10) / 3.14159))
                }
                _ = fi17
                _ = v18
            }
        }
        var v41 int64 = 250
        _ = fi10
        _ = v41
    }
    v42 := GenS4{F0: 9223372036854775807, F1: 319}
    var v43 string = "\n\t"
    var v44 float64 = 1.5
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    for fi45 := 0; fi45 < 18; fi45++ {
        var v46 uint = 712
        var v47 uint = v46
        var v48 uint16 = (v2 & (v2 * 32767))
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    var v49 bool = (true || (uint32(v4) >= 42))
    v44 = (v44 + v44)
    fmt.Println(v2, v44, v49)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v50 := GenS2{F0: 128, F1: (v2 ^ uint16(v4))}
    fmt.Println(v2, v44)
    v44 = v44
    for fi51 := 0; fi51 < 5; fi51++ {
        for fi52 := 0; fi52 < 10; fi52++ {
            _ = len(v3)
            var v53 string = v43
            switch ("hello" + "a") {
            case "":
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            case "hello":
                switch (fi52 / (((fi52 << ((4294967295) & 63))) | 1)) {
                case 32767:
                    if v49 {
                        var v54 uint8 = (v4 + (^(v4)))
                        if v49 {
                            v43 = "hello"
                            v54 = v54
                            v4 += 1
                        }
                        _ = v54
                    }
                }
            case "foo bar":
                v55 := GenS5{F0: (v2 ^ 1), F1: 2147483647}
                _ = v55
            default:
                v53 = v53
            }
            var v56 int8 = 100
            v4 += 1
            v53 = "a"
            _ = fi52
            _ = v53
            _ = v56
        }
        _ = fi51
    }
    v2 -= (v2 / ((v2) | 1))
    var v57 int8 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v49:", v49)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v49
    _ = v50
    _ = v57
}
